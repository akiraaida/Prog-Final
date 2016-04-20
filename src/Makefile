all: build prepare run

prepare:
	go build server.go

run:
	./server

build:
	go get golang.org/x/net/html

clean:
	rm server
