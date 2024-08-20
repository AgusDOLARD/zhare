BINARY=bin/zhare 
INSTALL_DIR="${HOME}/.local/bin/"
BUILD_FLAGS="-X 'main.version=`git describe --tags --abbrev=0`' -s -w"

build:
	@go build -ldflags=${BUILD_FLAGS} -o ${BINARY} main.go

run: build
	@bin/zhare

install: build
	@cp -v ${BINARY} ${INSTALL_DIR}
