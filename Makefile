all:
	GOOS=js GOARCH=wasm go build -o static/projectname_generator.wasm -ldflags="-s -w" -tags=js -gcflags="-m" -trimpath app.go

clean:
	rm -f projectname_generator.wasm
