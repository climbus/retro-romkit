build:
	go build -o bin/romkit ./cmd/cli/main.go

run: build
	romkit
test:
	go test -v ./...
