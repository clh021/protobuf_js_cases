package main

import (
	"bytes"
	"encoding/gob"
	"log"
	"net/http"

	pb "main/protos/window"

	"golang.org/x/net/context"
	"golang.org/x/net/trace"
	"google.golang.org/grpc/grpclog"
	// "google.golang.org/grpc/encode"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:50052"
)

// 定义helloService并实现约定的接口
type helloService struct{}

// HelloService Hello服务
var HelloService = helloService{}

// SayHello 实现Hello服务接口
func (h helloService) SayHello(ctx context.Context, in *pb.MainWindowList) (*pb.MainWindowList, error) {
	var chromiumId int32
	chromiumId = 1
	winId1 := &pb.WindowId{
		ChromiumId: &chromiumId,
	}

	winList := new(pb.MainWindowList)
	winList.Ok = true
	winList.Windows = append(winList.Windows, winId1)
	return winList, nil
}

func main() {
	var chromiumId int32
	chromiumId = 1
	winId1 := &pb.WindowId{
		ChromiumId: &chromiumId,
	}

	winList := new(pb.MainWindowList)
	winList.Ok = true
	winList.Windows = append(winList.Windows, winId1)

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(&winList)
	if err != nil {
		log.Println("Encode Error: %+v", err)
		log.Fatal("encode error:", err)
	}
	data := buf.Bytes()
	log.Println(data) //[255 211 255 129 3 1 1 12 ...]

	// listen, err := net.Listen("tcp", Address)
	// if err != nil {
	// 	grpclog.Fatalf("failed to listen: %v", err)
	// }

	// // 实例化grpc Server
	// s := grpc.NewServer()

	// // 注册HelloService
	// pb.RegisterTestServer(s, helloService)

	// // 开启trace
	// go startTrace()

	// grpclog.Println("Listen on " + Address)
	// s.Serve(listen)
}

func startTrace() {
	trace.AuthRequest = func(req *http.Request) (any, sensitive bool) {
		return true, true
	}
	go http.ListenAndServe(":50051", nil)
	grpclog.Println("Trace listen on 50051")
}
