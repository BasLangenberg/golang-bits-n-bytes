clean:
	rm -f html/*.wasm

build:
	tinygo build -o ./html/wasm.wasm -target wasm ./callback/main.go

serve:
	go run server.go