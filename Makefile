vet: fmt
	go vet ./...

build: vet
	go build ./...

test:
	go test ./... -v -race

fmt:
	go fmt ./...
