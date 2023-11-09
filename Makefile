.PHONY: local wasm clean

local:
	go build

wasm:
	env GOOS=js GOARCH=wasm go build -o www/zombie-dash.wasm
	cp $(shell go env GOROOT)/misc/wasm/wasm_exec.js www

clean:
	rm -f zombie-dash
	rm -f www/zombie-dash.wasm
	rm -f www/wasm_exec.js
