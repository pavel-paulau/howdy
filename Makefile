.PHONY: build clean fmt

build:
	go build -v && nrsc howdy app

fmt:
	find . -name "*.go" | xargs gofmt -w -s

clean:
	rm -fr howdy
