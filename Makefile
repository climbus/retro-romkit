build:
	go build -o bin/tosec-manager ./cmd/cli/main.go

run: build
	./bin/tosec-manager

test:
	go test -v ./...
