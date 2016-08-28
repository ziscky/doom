BINARY=doom
build:
	go build -o ${BINARY} main.go
install:
	mv doom /usr/bin
