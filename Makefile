all: build

build:
	go build -o qmpv

clean:
	go clean
	rm qmpv -rf

install: build
	install -m 777 qmpv /usr/local/bin/

.PHONY: build clean install
