
all: main

main: main.go
	go build main.go

gen:
	protoc --go_out=. protos/window.proto

run:
	go run main.go