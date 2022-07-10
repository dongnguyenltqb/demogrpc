default: develop

develop:
	@PORT=7777 CERT_PATH=$$PWD/cert/server.crt KEY_PATH=$$PWD/cert/server.key air

test:
	@SERVER_ADDR=localhost:7777 CERT_SERVER_NAME=dongnguyen.dev CERT_PATH=$$PWD/cert/server.crt KEY_PATH=$$PWD/cert/server.key go test client/* -v

build:
	@go mod download
	@GOOS=linux GOARCH=amd64 go build
