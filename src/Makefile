all: build prepare run

prepare:
	go build main.go

run:
	./main

build:
	go get golang.org/x/net/html

clean:
	rm main