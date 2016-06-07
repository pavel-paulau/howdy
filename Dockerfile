FROM alpine:3.4

MAINTAINER Pavel Paulau <Pavel.Paulau@gmail.com>

EXPOSE 8081

# Install dependencies
RUN apk --update add ca-certificates

# Run bot in foreground
COPY howdy /usr/local/bin/howdy

CMD ["/usr/local/bin/howdy"]
