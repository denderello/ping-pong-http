# Ping Pong HTTP

[![Docker Repository on Quay](https://quay.io/repository/denderello/ping-pong-http/status "Docker Repository on Quay")](https://quay.io/repository/denderello/ping-pong-http)
[![Go Report Card](https://goreportcard.com/badge/denderello/ping-pong-http "Go Report Card")](https://goreportcard.com/report/denderello/ping-pong-http)
[![Build Status](https://travis-ci.org/denderello/ping-pong-http.svg?branch=master)](https://travis-ci.org/denderello/ping-pong-http)

`ping-pong-http` is a little http server that responds on ping requests with
pong message in a JSON format.

The basic idea of this project is to try out some topics:
 - How to integrate github.com/gorilla/mux
 - How to handle graceful shutdowns using gopkg.in/tylerb/graceful.v1
 - How to make logging injectable instead of using the `log` everywhere
 - How to handle multiple API versions in one server

## Running server

You can run the server using the compiled binary:
```
ping-pong-http
```

### Docker Support

Instead of compiling and running the binary yourself you can also use the public
available docker image hosted on [Quay](https://quay.io):
```
docker run --rm -it -p 9000:9000 quay.io/denderello/ping-pong-http
```
