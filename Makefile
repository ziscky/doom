BINARY=doom
build:
    go get gopkg.in/urfave/cli.v1
	go build -o ${BINARY} main.go
install:
	mv doom /usr/bin
