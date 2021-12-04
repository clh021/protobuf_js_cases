
all: main

main: main.go
	go build main.go

gen:
	protoc --go_out=. message/window.proto