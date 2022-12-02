all:
	GOOS=js GOARCH=wasm go build -o ./static/app.wasm -ldflags="-s -w" -tags=js -gcflags="-m" -trimpath app.go

clean:
	rm -f ./static/app.wasm
