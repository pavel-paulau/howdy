.PHONY: build clean fmt

build:
	go build -v && nrsc howdy app

bot:
	go build -v -o bot ./demo

fmt:
	find . -name "*.go" -not -path "./vendor/*" | xargs gofmt -w -s

docker:
	go build -v -tags netgo && strip howdy && upx -q6 howdy && nrsc howdy app
	docker build --rm -t howdy .

clean:
	rm -fr howdy build bot
