BINARY=bin/zhare 
INSTALL_DIR="${HOME}/.local/bin/"

build:
	@go build -ldflags="-s -w" -o ${BINARY} main.go

run: build
	@bin/zhare

install: build
	@cp -v ${BINARY} ${INSTALL_DIR}
