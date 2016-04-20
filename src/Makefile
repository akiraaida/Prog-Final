all: build prepare run

prepare:
	go build server.go

run:
	nohup ./server &

build:
	go get golang.org/x/net/html

clean:
	rm server
