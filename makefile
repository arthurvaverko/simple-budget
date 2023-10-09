build-wasm:
	GOARCH=wasm GOOS=js go build -o web/app.wasm

build: build-wasm
	go build -o simple-budget

dev: build-wasm
	 go run main.go

run: build
	./simple-budget