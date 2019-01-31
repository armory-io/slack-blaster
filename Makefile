COMMIT=$(shell git rev-parse HEAD)
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)

LDFLAGS ?= -ldflags "-X main.COMMIT=${COMMIT} -X main.BRANCH=${BRANCH}"

pre:
	mkdir -p ./build

fmt:
	go fmt ./...

build: pre
	mkdir -p ./build/macos
	mkdir -p ./build/linux
	GOOS=darwin GOARCH=amd64 go build -v -o ./build/macos/slack-blaster-mac ${LDFLAGS}
	GOOS=linux GOARCH=amd64 go build -v -o ./build/linux/slack-blaster-linux ${LDFLAGS}

clean:
	rm -rf ./build
