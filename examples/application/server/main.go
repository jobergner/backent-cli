package state

import (
	"fmt"
	"log"
	"net/http"

	"nhooyr.io/websocket"
)

// NEED:
// - rest API for room management (create/delete)
// - loginHanlder, holds rooms, and client (clients need to be somehwere until a message arrives which assigns them to a room)
// 	methods:
// 		- func(message)

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

	c, err := newClient(NewConnection(websocketConnection, r), newLoginHandler(room.sideEffects))
	if err != nil {
		log.Println(err)
		return
	}

	room.registerClientSync(c)

	go c.runReadMessages()
	go c.runWriteMessages()

	// wait until client disconnects
	<-r.Context().Done()
	c.removeSelfFromSystem()
}

func setupRoutes(actions Actions, sideEffects SideEffects, fps int) {
	room := newRoom(actions, sideEffects, fps)
	room.Deploy()

	http.HandleFunc("/", homePageHandler)
	http.HandleFunc("/inspect", inspectHandler)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) { wsEndpoint(w, r, room) })
	http.HandleFunc("/state", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		tree := room.state.assembleTree(true)
		stateBytes, err := tree.MarshalJSON()
		if err != nil {
			http.Error(w, "Error marshalling tree", 500)
		}
		fmt.Fprint(w, string(stateBytes))
	})
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
