GO_FILES := $(shell find . -name '*.go')

.PHONY: default
default: generate

.PHONY: web
web: web/app.wasm

web/app.wasm: $(GO_FILES)
	GOARCH=wasm GOOS=js go build -o web/app.wasm

.PHONY: bin
bin: bin/generate bin/serve

bin/generate: $(GO_FILES)
	go build -o bin/generate ./cmd/generate

bin/serve: $(GO_FILES)
	go build -o bin/serve ./cmd/serve

.PHONY: clean
clean:
	rm -f ./bin/generate ./bin/serve ./web/app.wasm
	rm -rf ./docs
	mkdir ./docs

.PHONY: generate
generate: bin/generate web/app.wasm
	cp -r web docs/
	./bin/generate

.PHONY: serve
serve: web/app.wasm bin/serve
	lsof -nti:8000 | xargs kill -9
	./bin/serve
