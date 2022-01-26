build:
	go build -o bin/akasio ./cmd/akasio/main.go
	strip -s bin/akasio

run:
	go run ./cmd/akasio/main.go

install:
	cp -v bin/akasio /usr/local/bin/akasio

compile:
	GOOS=linux GOARCH=amd64 go build -o bin/akasio-linux-amd64 ./cmd/akasio/main.go
	GOOS=windows GOARCH=amd64 go build -o bin/akasio-windows-amd64.exe ./cmd/akasio/main.go
	GOOS=darwin GOARCH=amd64 go build -o bin/akasio-darwin-amd64 ./cmd/akasio/main.go
	strip -s bin/akasio-linux-amd64

upload:
	sudo podman build -t ghcr.io/k4yt3x/akasio-go:1.1.1 .
	sudo podman push ghcr.io/k4yt3x/akasio-go:1.1.1
