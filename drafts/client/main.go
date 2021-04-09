package main

import (
	"context"
	"log"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	c, _, err := websocket.Dial(ctx, "http://localhost:8080/ws", nil)
	if err != nil {
		panic(err)
	}
	defer c.Close(websocket.StatusInternalError, "the sky is falling")
	go runReadMessages(c, ctx)

	go runSendMSG(ctx, c)

	time.Sleep(time.Second * 100)

	c.Close(websocket.StatusNormalClosure, "")
}

func runReadMessages(conn *websocket.Conn, ctx context.Context) {
	defer panic("client discontnued")
	for {
		_, message, err := conn.Read(ctx)
		if err != nil {
			log.Println(err)
			break
		}

		log.Println(string(message))
	}
}

type messageKind int

const (
	messageKindInit messageKind = iota + 1
)

type message struct {
	Kind    messageKind `json:"kind"`
	Content []byte      `json:"content"`
}

func runSendMSG(ctx context.Context, con *websocket.Conn) {
	ticker := time.NewTicker(time.Second)
	msg := message{
		Kind:    2,
		Content: []byte(`{"playerID": 1, "changeX": 1.1, "changeY": 1.1}`),
	}
	for {
		<-ticker.C
		err := wsjson.Write(ctx, con, msg)
		if err != nil {
			panic(err)
		}
	}
}
