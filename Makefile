.PHONY: build bot

build:
	go-bindata --debug app/...
	go build -v

bot:
	go build -v -o bot ./demo

fmt:
	find . -name "*.go" -not -path "./vendor/*" | xargs gofmt -w -s

docker:
	go-bindata app/...
	go build -v -tags netgo --ldflags "-s" && upx -q6 howdy
	docker build --rm -t howdy .

clean:
	rm -fr howdy build bot bindata.go
