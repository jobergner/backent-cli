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

	err = wsjson.Write(ctx, c, "hi")
	if err != nil {
		panic(err)
	}

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
