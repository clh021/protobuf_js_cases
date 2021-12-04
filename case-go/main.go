package main

import (
	"fmt"
	"log"
	"main/message"
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
	winId1 := &message.WindowId{
		ChromiumId: 1,
	}
	winList := &message.MainWindowList{
		Ok:      true,
		Windows: []*message.WindowId{winId1},
	}
	fmt.Println("winList: %s", winList)
	winData, _ := proto.Marshal(winList)

	fmt.Println("winList: %s", winData)

	msg := &message.Message{
		Id: proto.Int32(17),
		Author: &message.Message_Person{
			Id:   proto.Int32(1),
			Name: proto.String("othree"),
		},
		Text: proto.String("Hi, this is message."),
	}

	fmt.Println(msg.GetAuthor().GetName() + ": " + msg.GetText())

	data, _ = proto.Marshal(msg)

	fmt.Println("%s", data)

	http.HandleFunc("/ws", handler)
	if err := http.ListenAndServe("127.0.0.1:1337", nil); err != nil {
		log.Fatal(err)
	}
}
