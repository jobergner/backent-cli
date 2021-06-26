package main

import (
	"context"
	"fmt"
	"github.com/Java-Jonas/bar-cli/integrationtest/state"
	"log"
	// "os"
	"time"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

func main() {
	go startServer()
	cancel := dialServer()
	time.Sleep(time.Second * 10)
	cancel()
}

func dialServer() context.CancelFunc {
	ctx, cancel := context.WithCancel(context.Background())

	c, _, err := websocket.Dial(ctx, "http://localhost:8080/ws", nil)
	if err != nil {
		panic(err)
	}
	go runReadMessages(c, ctx)

	go runSendMessage(ctx, c)

	return cancel
}

func runReadMessages(conn *websocket.Conn, ctx context.Context) {
	defer fmt.Println("client discontinued")
	for {
		_, message, err := conn.Read(ctx)
		if err != nil {
			log.Println(err)
			break
		}
		log.Println("server response: ", string(message))
	}
}

var newx = 1.1

func runSendMessage(ctx context.Context, con *websocket.Conn) {
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			newx += 1
			msg := message{
				Kind:    messageKindAction_movePlayer,
				Content: []byte(`{"playerID": 2, "changeX": ` + fmt.Sprintf("%f", newx) + `, "changeY": 0}`),
			}
			err := wsjson.Write(ctx, con, msg)
			if err != nil {
				log.Println(err)
			}
			params := state.AddItemToPlayerParams{
				Item:    state.ItemID(0),
				NewName: "myItem",
			}
			b, err := params.MarshalJSON()
			if err != nil {
				log.Println(err)
			}
			msg = message{
				Kind:    messageKindAction_addItemToPlayer,
				Content: b,
			}
			err = wsjson.Write(ctx, con, msg)
			if err != nil {
				log.Println(err)
			}
		case <-ctx.Done():
			break
		}
	}
}
