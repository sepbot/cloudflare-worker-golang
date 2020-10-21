//+ build js,wasm

package main

import (
	"crypto/ed25519"
	"encoding/base64"
	"strings"
	"syscall/js"
)

var (
	privateKey ed25519.PrivateKey
	publicKey ed25519.PublicKey
)

func sign(_ js.Value, input []js.Value) interface{} {
	message := []byte(input[0].String())
	callback := input[1]

	messageB64 := make([]byte, base64.RawURLEncoding.EncodedLen(len(message)))
	base64.RawURLEncoding.Encode(messageB64, message)

	signature := base64.RawURLEncoding.EncodeToString(ed25519.Sign(privateKey, messageB64))
	token := strings.Join([]string{string(messageB64), signature}, ".")

	callback.Invoke(token)
	return nil
}

func verify(_ js.Value, input []js.Value) interface{} {
	token := strings.Split(input[0].String(), ".")
	callback := input[1]

	signature, err := base64.RawURLEncoding.DecodeString(token[1])
	if err != nil {
		panic(err)
	}

	callback.Invoke(ed25519.Verify(publicKey, []byte(token[0]), signature))
	return nil
}

func main() {
	var err error

	privateKey, err = base64.RawURLEncoding.DecodeString(js.Global().Get("PRIVATE_KEY").String())
	if err != nil {
		panic(err)
	}

	publicKey = make([]byte, ed25519.PublicKeySize)
	copy(publicKey, privateKey[32:])

	js.Global().Set("sign", js.FuncOf(sign))
	js.Global().Set("verify", js.FuncOf(verify))
	<-make(chan interface{})
}
