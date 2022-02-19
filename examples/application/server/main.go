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

func wsEndpoint(w http.ResponseWriter, r *http.Request, loginHandler *LoginHandler) {

	websocketConnection, err := websocket.Accept(w, r, &websocket.AcceptOptions{InsecureSkipVerify: true})
	if err != nil {
		log.Println(err)
		return
	}

	client, err := newClient(NewConnection(websocketConnection, r))
	if err != nil {
		log.Println(err)
		return
	}

	loginHandler.addClient(client)

	go client.runReadMessages()
	go client.runWriteMessages()

	// wait until connection closes
	<-r.Context().Done()
	loginHandler.deleteClient(client)
}

func setupRoutes(loginHandler *LoginHandler) {
	http.HandleFunc("/", homePageHandler)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) { wsEndpoint(w, r, loginHandler) })
}

func Start(signals LoginSignals, actions Actions, sideEffects SideEffects, fps int, port int) error {
	// TODO: what does this even do??
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	loginHandler := newLoginHandler(signals, actions, sideEffects, fps)

	setupRoutes(loginHandler)

	fmt.Printf("backent running on port %d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	return err
}
