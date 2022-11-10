# Hello World!

..in Go with [gin-gonic](https://github.com/gin-gonic/gin)

## Setup

- Install [Go](https://go.dev/)

## Build Project

```bash
# Clone repo
$ git clone https://github.com/bits0rcerer/hello_worlds
$ cd go_gin-gonic

# Compile
$ go build

# run
$ ./hello_go
```

The web server is now running and listening on port `8080`.

Try with your browser:

- http://localhost:8080/hello
- http://localhost:8080/hello?name=Go
- http://localhost:8080/hello_json
- http://localhost:8080/hello_json?name=Go

## Run Tests

```bash
# Clone repo
$ git clone https://github.com/bits0rcerer/hello_worlds
$ cd go_gin-gonic

# build and run tests
$ go test
```