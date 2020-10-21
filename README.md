cloudflare worker that calls into golang code to perform computational heavy processing
i saw a 15x performance increase compared to the native js implementation of the same algo

to run `make build` you will need the following env vars:

```
GOROOT=
CF_WORKER_NAME=
CF_ACCOUNT_ID=
CF_API_TOKEN=
PRIVATE_KEY=
```
