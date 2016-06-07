.PHONY: build clean fmt

build:
	go build -v && nrsc howdy app

fmt:
	find . -name "*.go" | xargs gofmt -w -s

docker:
	go build -v -tags netgo && strip howdy && upx -q6 howdy && nrsc howdy app
	docker build --rm -t howdy .

clean:
	rm -fr howdy build
