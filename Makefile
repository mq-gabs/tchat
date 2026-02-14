all: build-client build-server

build-client:
	go build -o dist/tchat tchat.com/client

build-server:
	go build -o dist/tchat-server tchat.com/server
