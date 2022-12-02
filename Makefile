all:
	GOOS=js GOARCH=wasm go build -o ./static/app.wasm app.go
	# cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ./static/wasm_exec.js

serve:
	caddy file-server --listen :8080 --root ./static

clean:
	rm -f ./static/app.wasm
