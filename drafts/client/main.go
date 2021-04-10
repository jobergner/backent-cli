package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
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

var newx = 1.1

func runSendMSG(ctx context.Context, con *websocket.Conn) {
	ticker := time.NewTicker(time.Second)
	for {
		<-ticker.C
		newx += 1
		msg := message{
			Kind:    1,
			Content: []byte(`{"playerID": 2, "changeX": ` + fmt.Sprintf("%f", newx) + `, "changeY": 0}`),
		}
		err := wsjson.Write(ctx, con, msg)
		if err != nil {
			panic(err)
		}
	}
}
