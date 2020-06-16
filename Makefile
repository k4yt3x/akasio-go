build:
	go build -o bin/akasio ./cmd/akasio/main.go

run:
	go run ./cmd/akasio/main.go

install:
	sudo cp bin/akasio /usr/local/bin/akasio

compile:
	GOOS=linux GOARCH=amd64 go build -o bin/akasio-linux-amd64 ./cmd/akasio/main.go
	GOOS=windows GOARCH=amd64 go build -o bin/akasio-windows-amd64.exe ./cmd/akasio/main.go
	GOOS=darwin GOARCH=amd64 go build -o bin/akasio-darwin-amd64 ./cmd/akasio/main.go