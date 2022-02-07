package main

import (
	"fmt"
	"log"
	"main/message"
	"main/protos/window"
	"net/http"

	"google.golang.org/protobuf/proto"
	// "github.com/golang/protobuf/proto"
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
func server() {
	msg := &message.Message{
		Id: proto.Int32(17),
		Author: &message.Message_Person{
			Id:   proto.Int32(1),
			Name: proto.String("othree"),
		},
		Text: proto.String("Hi, this is message."),
	}

	fmt.Println(msg.GetAuthor().GetName() + ": " + msg.GetText())

	// data, _ = proto.Marshal(&msg)

	// fmt.Println("%s", data)

	http.HandleFunc("/ws", handler)
	if err := http.ListenAndServe("127.0.0.1:1337", nil); err != nil {
		log.Fatal(err)
	}
}

func test() {
	var testInt int32
	testInt = 150
	test := &window.Test1{
		A: &testInt,
	}
	fmt.Println("testInt: ", test)
	testData, _ := proto.Marshal(test)
	fmt.Println("testInt: ", testData)

	var testString string
	testString = "test"
	test2 := &window.Test2{
		B: &testString,
	}
	fmt.Println("testStr: ", test2)
	test2Data, _ := proto.Marshal(test2)
	fmt.Println("testStr: ", test2Data)
}

// 测试验证消息，case 建立的目的
func run() {
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
	fmt.Println("winList: ", winList)
	winData, _ := proto.Marshal(winList)
	fmt.Println("winList: ", winData)

	var wins []*window.WindowId
	wins = append(wins, winId1)
	winList2 := &window.MainWindowList{
		Ok:      true,
		Windows: wins,
	}
	fmt.Println("winList2: ", winList2)
	winData2, _ := proto.Marshal(winList2)
	fmt.Println("winList2: ", winData2)

	winTitle := "test"
	winProps := &window.MainWindowProps{
		Id:    winId1,
		Ok:    true,
		Title: &winTitle,
	}
	fmt.Println("winProps: ", winProps)
	winPropsData, _ := proto.Marshal(winProps)

	fmt.Println("winProps: ", winPropsData)
}
func main() {
	run() // 查看 case Test Result
	// test()
}
