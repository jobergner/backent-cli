package server

import (
	"sync"
	"time"

	"github.com/jobergner/backent-cli/examples/logging"
	"github.com/jobergner/backent-cli/examples/message"
	"github.com/jobergner/backent-cli/examples/state"
	"github.com/rs/zerolog/log"
)

type RoomMode int

const (
	RoomModeIdle RoomMode = iota
	RoomModeRunning
	RoomModeTerminating
)

type Room struct {
	name       string
	mu         sync.Mutex
	clients    *clientRegistrar
	state      *state.Engine
	controller Controller
	mode       RoomMode
}

func newRoom(controller Controller, name string) *Room {
	return &Room{
		name:       name,
		clients:    newClientRegistar(),
		state:      state.NewEngine(),
		controller: controller,
	}
}

func (r *Room) Name() string {
	return r.name
}

func (r *Room) RemoveClient(client *Client) {
	r.clients.remove(client)
}

func (r *Room) AddClient(client *Client) {
	client.room = r
	r.clients.add(client)
}

func (r *Room) AlterState(fn func(*state.Engine)) {
	r.mu.Lock()
	defer r.mu.Unlock()
	fn(r.state)
}

func (r *Room) RangeClients(fn func(client *Client)) {
	for c := range r.clients.incomingClients {
		fn(c)
	}
	for c := range r.clients.clients {
		fn(c)
	}
}

func (r *Room) processMessageSync(msg Message) {
	r.mu.Lock()
	defer r.mu.Unlock()

	response := r.processClientMessage(msg)

	if response.Kind == message.MessageKindNoResponse {
		return
	}

	responseBytes, err := response.MarshalJSON()
	if err != nil {
		log.Err(err).Str(logging.Message, string(responseBytes)).Str(logging.MessageKind, string(response.Kind)).Msg("failed marshalling response")
		return
	}

	select {
	case response.client.messageChannel <- responseBytes:
	default:
		log.Warn().Str(logging.ClientID, response.client.id).Msg(logging.ClientBufferFull)
		response.client.closeConnection()
	}
}

func (r *Room) run(controller Controller, fps int) {
	ticker := time.NewTicker(time.Second / time.Duration(fps))

	for {
		<-ticker.C

		r.tickSync(controller)

		if r.mode == RoomModeTerminating {
			break
		}
	}

	log.Debug().Str(logging.RoomName, r.name).Msg("terminating room")
}

func (r *Room) Deploy(controller Controller, fps int) {
	r.mu.Lock()
	log.Debug().Str(logging.RoomName, r.name).Msg("onDeploy")
	controller.OnDeploy(r.state)
	r.mu.Unlock()

	go r.run(controller, fps)
}
