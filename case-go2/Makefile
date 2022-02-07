
all: main

main: main.go
	go build main.go

gen:
	protoc --go_out=. protos/window.proto
	protoc --go_out=. protos/hello.proto

run:
	go run main.go