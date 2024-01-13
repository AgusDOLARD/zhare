build:
	@go build -o bin/zhare main.go

run: build
	@bin/zhare
