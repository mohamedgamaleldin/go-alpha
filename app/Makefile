.PHONY: all run build test

all : | build run

build :
	docker build -t go-alpha .

run :
	docker run -p 8080:8080 --name go-alpha go-alpha

test :
	go get -d ./...
	go run main.go