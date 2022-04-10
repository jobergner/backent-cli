package endtoend

import (
	"fmt"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/jobergner/backent-cli/examples/message"
	"github.com/jobergner/backent-cli/examples/server"
	"github.com/jobergner/backent-cli/examples/state"
	"github.com/stretchr/testify/assert"
)

func TestEndToEnd(t *testing.T) {
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
		m.EXPECT().OnClientConnect(gomock.Any(), gomock.Any()).Do(func(client *server.Client, lobby *server.Lobby) {
			if len(client.ID()) == 0 {
				panic("empty client ID")
			}
		})

		kill := startServer(m)

		connectClient(m, nil)

		kill()
	})
	t.Run("Rooms calls OnFrameTick every fps", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := NewMockController(ctrl)

		m.EXPECT().OnCreation(gomock.Any())
		m.EXPECT().OnFrameTick(gomock.Any()).MinTimes(2)

		kill := startServer(m, func(s *server.Server) {
			s.Lobby.CreateRoom("foo")
		})

		time.Sleep(time.Second / fps * 3)
		kill()
	})
	t.Run("Room calls actions when triggered by client", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := NewMockController(ctrl)

		params := message.AddItemToPlayerParams{
			NewName: "name",
		}

		expectedResponse := message.AddItemToPlayerResponse{PlayerPath: "foobar"}

		m.EXPECT().OnCreation(gomock.Any())
		m.EXPECT().OnClientConnect(gomock.Any(), gomock.Any()).Do(func(client *server.Client, lobby *server.Lobby) {
			r, _ := lobby.Rooms["foo"]
			r.AddClient(client)
		})
		m.EXPECT().OnFrameTick(gomock.Any()).AnyTimes()
		m.EXPECT().AddItemToPlayerBroadcast(params, gomock.Any(), "", gomock.Any())
		m.EXPECT().AddItemToPlayerBroadcast(params, gomock.Any(), "foo", gomock.Any())
		m.EXPECT().AddItemToPlayerEmit(params, gomock.Any(), "foo", gomock.Any()).DoAndReturn(func(params message.AddItemToPlayerParams, engine *state.Engine, roomName, clientID string) message.AddItemToPlayerResponse {
			return expectedResponse
		})

		kill := startServer(m, func(s *server.Server) {
			s.Lobby.CreateRoom("foo")
		})

		c, _ := connectClient(m, nil)
		res, err := c.AddItemToPlayer(params)
		if err != nil {
			panic(err)
		}
		assert.Equal(t, expectedResponse, res)

		kill()
	})
	t.Run("Client receives all updates from server", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := NewMockController(ctrl)

		m.EXPECT().OnCreation(gomock.Any())
		m.EXPECT().OnClientConnect(gomock.Any(), gomock.Any()).Do(func(client *server.Client, lobby *server.Lobby) {
			r, _ := lobby.Rooms["foo"]
			r.AddClient(client)
		}).Times(3)
		m.EXPECT().OnFrameTick(gomock.Any()).Do(func(engine *state.Engine) {
			p := engine.QueryPositions(func(state.Position) bool { return true })[0]
			p.SetX(p.X() + 1)
		}).AnyTimes()

		kill := startServer(m, func(s *server.Server) {
			s.Lobby.CreateRoom("foo").AlterState(func(engine *state.Engine) {
				engine.CreatePosition().SetX(1)
			})
		})

		addClient := func() {
			// TODO this is not very clean
			var lastXValue float64
			var i int
			onMessageReceived := func(b []byte) {
				var t state.Tree
				err := t.UnmarshalJSON(b)
				if err != nil {
					panic(err)
				}

				for _, p := range t.Position {
					if *p.X <= lastXValue {
						panic(fmt.Sprintf("%f is not higher than %f", *p.X, lastXValue))
					}
					lastXValue = *p.X
					break
				}
				i++
			}

			connectClient(m, onMessageReceived)

			// add some extra ms se we can be sure the last tick passes
			time.Sleep(time.Second/fps*5 + time.Millisecond*30)

			minClientReceiveUpdate := 4

			if i-1 > minClientReceiveUpdate {
				panic(fmt.Sprintf("expected %d client ticks, but got %d", minClientReceiveUpdate, i))
			}
		}

		addClient()
		addClient()
		addClient()

		kill()
	})
	t.Run("Lobby removes client from room when client disconnects", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := NewMockController(ctrl)

		m.EXPECT().OnCreation(gomock.Any())
		m.EXPECT().OnClientConnect(gomock.Any(), gomock.Any())
		m.EXPECT().OnClientDisconnect(gomock.Nil(), gomock.Any(), gomock.Any())

		kill := startServer(m)

		_, cancel := connectClient(m, nil)
		cancel()

		kill()
	})
	t.Run("SuperMessage gets triggered when client sends it", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := NewMockController(ctrl)

		m.EXPECT().OnCreation(gomock.Any())
		m.EXPECT().OnClientConnect(gomock.Any(), gomock.Any())
		m.EXPECT().OnSuperMessage(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any())

		kill := startServer(m)

		c, _ := connectClient(m, nil)

		c.SuperMessage([]byte("foo"))

		kill()
	})
}
