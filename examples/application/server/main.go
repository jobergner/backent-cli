package state

import (
	"fmt"
	"log"
	"net/http"

	"nhooyr.io/websocket"
)

func homePageHandler(w http.ResponseWriter, r *http.Request) {
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

func setupRoutes(actions Actions, sideEffects SideEffects, fps int) {
	room := newRoom(actions, sideEffects, fps)
	room.Deploy()

	http.HandleFunc("/", homePageHandler)
	http.HandleFunc("/inspect", inspectHandler)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) { wsEndpoint(w, r, room) })
}

func Start(actions Actions, sideEffects SideEffects, fps int, port int) error {
	if fps < 1 {
		setupRoutes(actions, sideEffects, 1)
	} else {
		setupRoutes(actions, sideEffects, fps)
	}
	fmt.Printf("backent running on port %d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	return err
}
