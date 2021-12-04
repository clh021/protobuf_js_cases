package main

import (
	"fmt"
	"main/protos/window"
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
)

var data []byte

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			return
		}
		fmt.Println("read done")
		fmt.Println("write %s", data)
		if err = conn.WriteMessage(websocket.BinaryMessage, data); err != nil {
			fmt.Println("err %s", err)
			return
		}
		fmt.Println("write done")
	}
}

func main() {
	// protoc 对应包名 protobuf-compiler 有 bug，应使用 https://github.com/protocolbuffers/protobuf/releases/download/v3.19.1/protoc-3.19.1-linux-x86_64.zip
	var chromiumId int32
	chromiumId = 1
	winId1 := &window.WindowId{
		ChromiumId: &chromiumId,
	}
	winList := &window.MainWindowList{
		Ok:      true,
		Windows: []*window.WindowId{winId1},
	}
	fmt.Println("winList: %s", winList)
	winData, _ := proto.Marshal(winList)

	fmt.Println("winList: %s", winData)

	// msg := &message.Message{
	// 	Id: proto.Int32(17),
	// 	Author: &message.Message_Person{
	// 		Id:   proto.Int32(1),
	// 		Name: proto.String("othree"),
	// 	},
	// 	Text: proto.String("Hi, this is message."),
	// }

	// fmt.Println(msg.GetAuthor().GetName() + ": " + msg.GetText())

	// data, _ = proto.Marshal(msg)

	// fmt.Println("%s", data)

	// http.HandleFunc("/ws", handler)
	// if err := http.ListenAndServe("127.0.0.1:1337", nil); err != nil {
	// 	log.Fatal(err)
	// }
}
