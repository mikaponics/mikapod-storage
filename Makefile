# Makefile: Compiling instructions for output an executable for this repository.

prerequisites: # Install all the libraries our project depends on.
	go get -u github.com/mattn/go-sqlite3

build: # Instruction will create a binary executable in our `/bin` folder.
	go build -o bin/mikapod-storage main.go

build-arm: # Instructions will create a binary executable for the Raspberry Pi hardware in our `/bin` folder.
	CGO_ENABLED=1 GOOS=linux GOARCH=arm GOARM=7 go build -o bin/mikapod-storage -v

deliver:
	croc bin/mikapod-storage

deploy-arm: build-arm deliver

run:
	go run main.go
