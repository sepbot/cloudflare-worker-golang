.PHONY: build

build: export GOOS=js
build: export GOARCH=wasm
build:
	rm -rf .cache dist

	cp "${GOROOT}/misc/wasm/wasm_exec.js" .
	npm install
	npm run build

	go build -o dist/go.wasm go.go

	cp metadata.json dist/metadata.json
	sed -i 's/PRIVATE_KEY_PLACEHOLDER/${PRIVATE_KEY}/g' dist/metadata.json

	curl -X PUT \
		"https://api.cloudflare.com/client/v4/accounts/${CF_ACCOUNT_ID}/workers/scripts/${CF_WORKER_NAME}" \
		-H "Authorization: Bearer ${CF_API_TOKEN}" \
		-F "metadata=@dist/metadata.json;type=application/json" \
		-F "script=@dist/worker.js;type=application/javascript" \
		-F "wasm=@dist/go.wasm;type=application/wasm" > /dev/null # the response is too verbose
