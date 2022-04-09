package endtoend

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/rs/zerolog/log"

	gomock "github.com/golang/mock/gomock"
	"github.com/jobergner/backent-cli/examples/client"
	"github.com/jobergner/backent-cli/examples/message"
	"github.com/jobergner/backent-cli/examples/server"
)

const fps = 3

func startServer(m *MockController, configs ...func(s *server.Server)) func() {
	s := server.NewServer(m, fps)

	for _, c := range configs {
		c(s)
	}

	signalFail := s.Start()
	go func() {
		err := <-signalFail
		if err != http.ErrServerClosed {
			panic(err)
		}
	}()

	kill := make(chan struct{})

	go func() {
		<-kill
		s.HttpServer.Shutdown(context.Background())
	}()

	time.Sleep(time.Microsecond * 100)
	return func() {
		kill <- struct{}{}
		time.Sleep(time.Millisecond * 50)
	}
}

func connectClient(m *MockController) *client.Client {
	c, err := client.NewClient(context.Background(), m, fps)
	if err != nil {
		panic(err)
	}

	go func() {
		for range c.ReadUpdate() {
		}
	}()

	return c
}

func TestEndToEnd(t *testing.T) {
	log.Logger = log.Logger.Hook(SeverityHook{})

	t.Run("Lobby calls OnCreation when server gets created", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := NewMockController(ctrl)

		m.EXPECT().OnCreation(gomock.Any())

		kill := startServer(m)
		kill()
	})
	t.Run("Lobby calls OnClientConnect when client connects", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := NewMockController(ctrl)

		m.EXPECT().OnCreation(gomock.Any())
		m.EXPECT().OnClientConnect(gomock.Any(), gomock.Any())

		kill := startServer(m)

		connectClient(m)

		kill()
	})
	t.Run("Rooms calls OnFrameTick every fps", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := NewMockController(ctrl)

		m.EXPECT().OnCreation(gomock.Any())
		m.EXPECT().OnClientConnect(gomock.Any(), gomock.Any())
		m.EXPECT().OnFrameTick(gomock.Any()).MinTimes(2)

		kill := startServer(m, func(s *server.Server) {
			s.Lobby.CreateRoom("foo")
		})

		connectClient(m)

		time.Sleep(time.Second / fps * 3)
		kill()
	})
	t.Run("Room calls actions when triggered by client", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := NewMockController(ctrl)

		params := message.AddItemToPlayerParams{
			NewName: "name",
		}

		m.EXPECT().OnCreation(gomock.Any())
		m.EXPECT().OnClientConnect(gomock.Any(), gomock.Any()).Do(func(client *server.Client, lobby *server.Lobby) {
			r, _ := lobby.Rooms["foo"]
			r.AddClient(client)
		})
		m.EXPECT().OnFrameTick(gomock.Any()).AnyTimes()
		m.EXPECT().AddItemToPlayerBroadcast(params, gomock.Any(), "", gomock.Any())
		m.EXPECT().AddItemToPlayerBroadcast(params, gomock.Any(), "foo", gomock.Any())
		m.EXPECT().AddItemToPlayerEmit(params, gomock.Any(), "foo", gomock.Any())

		kill := startServer(m, func(s *server.Server) {
			s.Lobby.CreateRoom("foo")
		})

		c := connectClient(m)
		_, err := c.AddItemToPlayer(params)
		if err != nil {
			panic(err)
		}

		kill()
	})
}
