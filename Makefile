default: develop

develop:
	@PORT=7777 air

build:
	@go mod download
	@GOOS=linux GOARCH=amd64 go build