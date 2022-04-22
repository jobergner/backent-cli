package server

import (
	"fmt"
	"net/http"

	"github.com/jobergner/backent-cli/examples/connect"
	"github.com/rs/zerolog/log"
	"nhooyr.io/websocket"
)

// easyjson:skip
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

	server.HttpServer.Handler = handler

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

	// TODO I dont think this works as I think it should work
	// wait until connection closes
	<-client.conn.Context().Done()

	log.Debug().Msg("client context done")
	s.Lobby.deleteClient(client)
}

func (s *Server) Start() chan error {
	log.Info().Msgf("server running on port %s\n", s.HttpServer.Addr)

	serverError := make(chan error)

	go func() {
		err := s.HttpServer.ListenAndServe()
		serverError <- err
	}()

	return serverError
}
