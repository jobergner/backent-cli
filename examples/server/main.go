package server

import (
	"fmt"
	"net/http"

	"github.com/jobergner/backent-cli/examples/connect"
	"github.com/rs/zerolog/log"
	"nhooyr.io/websocket"
)

type Server struct {
	HttpServer *http.Server
	Lobby      *Lobby
}

func NewServer(controller Controller, fps int, configs ...func(*http.Server, *http.ServeMux)) *Server {
	server := Server{
		HttpServer: new(http.Server),
		Lobby:      newLobby(controller, fps),
	}

	handler := http.NewServeMux()

	for _, c := range configs {
		c(server.HttpServer, handler)
	}

	if server.HttpServer.Addr == "" {
		server.HttpServer.Addr = fmt.Sprintf(":%d", 8080)
	}

	handler.HandleFunc("/", homePageHandler)
	handler.HandleFunc("/ws", server.wsEndpoint)

	return &server
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func (s *Server) wsEndpoint(w http.ResponseWriter, r *http.Request) {

	websocketConnection, err := websocket.Accept(w, r, &websocket.AcceptOptions{InsecureSkipVerify: true})
	if err != nil {
		log.Err(err).Msg("failed creating websocket connection")
		return
	}

	client, err := newClient(connect.NewConnection(websocketConnection, r.Context()), s.Lobby)
	if err != nil {
		return
	}

	s.Lobby.addClient(client)

	go client.runReadMessages()
	go client.runWriteMessages()

	// wait until connection closes
	<-r.Context().Done()
	s.Lobby.deleteClient(client)
}

func (s *Server) Start(controller Controller, fps int, port int) chan error {
	log.Info().Msgf("backent running on port %d\n", port)

	serverError := make(chan error, 1)

	go func() {
		err := s.HttpServer.ListenAndServe()
		serverError <- err
	}()

	return serverError
}
