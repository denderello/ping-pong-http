FROM golang:1.6-alpine

RUN apk add --update \
    make

COPY . $GOPATH/src/github.com/denderello/ping-pong-http
WORKDIR $GOPATH/src/github.com/denderello/ping-pong-http

RUN make install

ENTRYPOINT ["ping-pong-http"]

EXPOSE 9000
