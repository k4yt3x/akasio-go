build:
	go build -o bin/akasio ./cmd/akasio/main.go

run:
	go run ./cmd/akasio/main.go

install:
	sudo cp bin/akasio /usr/local/bin/akasio
