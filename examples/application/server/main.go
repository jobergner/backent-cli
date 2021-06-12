package state

import (
	"fmt"
	"log"
	"net/http"

	"nhooyr.io/websocket"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func wsEndpoint(w http.ResponseWriter, r *http.Request, room *Room) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	websocketConnection, err := websocket.Accept(w, r, &websocket.AcceptOptions{InsecureSkipVerify: true})
	if err != nil {
		log.Println(err)
		return
	}

	c, err := newClient(NewConnection(websocketConnection, r))
	if err != nil {
		log.Println(err)
		return
	}
	c.assignToRoom(room)
	room.registerChannel <- c

	go c.runReadMessages()
	go c.runWriteMessages()

	// wait until client disconnects
	<-r.Context().Done()
}

func setupRoutes(a actions, onDeploy func(*Engine), onFrameTick func(*Engine)) {
	room := newRoom(a, onDeploy, onFrameTick)
	room.Deploy()

	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) { wsEndpoint(w, r, room) })
}
