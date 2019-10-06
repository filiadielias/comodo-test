test:
	@GO111MODULE=on go test -v ./...

build:
	@GO111MODULE=on go mod tidy
	@GO111MODULE=on go test -v ./...
	@GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
