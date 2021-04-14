all: build-demo

build-demo:
	mkdir -p bin
	go build -o bin/velocity-worker cmd/velocity-worker/main.go
