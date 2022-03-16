package server

import (
	"fmt"
	"net/http"

	"github.com/jobergner/backent-cli/examples/connect"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"nhooyr.io/websocket"
)

func init() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func wsEndpoint(w http.ResponseWriter, r *http.Request, lobby *Lobby) {

	websocketConnection, err := websocket.Accept(w, r, &websocket.AcceptOptions{InsecureSkipVerify: true})
	if err != nil {
		log.Err(err).Msg("failed creating websocket connection")
		return
	}

	client, err := newClient(connect.NewConnection(websocketConnection, r.Context()), lobby)
	if err != nil {
		return
	}

	lobby.addClient(client)

	go client.runReadMessages()
	go client.runWriteMessages()

	// wait until connection closes
	<-r.Context().Done()
	lobby.deleteClient(client)
}

func setupRoutes(loginHandler *Lobby) {
	http.HandleFunc("/", homePageHandler)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) { wsEndpoint(w, r, loginHandler) })
}

func Start(controller Controller, fps int, port int) error {
	loginHandler := newLoginHandler(controller, fps)

	setupRoutes(loginHandler)

	log.Info().Msgf("backent running on port %d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	return err
}
