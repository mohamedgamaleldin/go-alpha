.PHONY: all run build test

all : | build run

build :
	docker build -t go-alpha .

run :
	docker run --rm --name go-alpha -p 8080:8080 --network go-alpha_demo_network go-alpha

test :
	go get -d ./...
	go run main.go