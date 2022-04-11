package server

import (
	"sync"

	"github.com/jobergner/backent-cli/examples/logging"
	"github.com/rs/zerolog/log"
)

type Lobby struct {
	mu         sync.Mutex
	Rooms      map[string]*Room
	controller Controller
	fps        int
}

func newLobby(controller Controller, fps int) *Lobby {
	l := &Lobby{
		Rooms:      make(map[string]*Room),
		controller: controller,
		fps:        fps,
	}

	l.signalCreation()

	return l
}

func (l *Lobby) CreateRoom(name string) *Room {
	if room, ok := l.Rooms[name]; ok {
		log.Warn().Str(logging.RoomName, name).Msg("attempted to create room which already exists")
		return room
	}

	room := newRoom(l.controller, name)

	l.Rooms[name] = room

	room.Deploy(l.controller, l.fps)

	return room
}

func (l *Lobby) DeleteRoom(name string) {
	room, ok := l.Rooms[name]
	if !ok {
		log.Warn().Str(logging.RoomName, name).Msg("attempted to delete room which does not exists")
		return
	}

	room.mode = RoomModeTerminating

	delete(l.Rooms, name)
}

func (l *Lobby) addClient(client *Client) {
	l.signalClientConnect(client)
}

func (l *Lobby) deleteClient(client *Client) {
	if client.room != nil {
		client.room.clients.remove(client)
	}

	l.signalClientDisconnect(client)
}

func (l *Lobby) processMessageSync(msg Message) {

	l.mu.Lock()
	defer l.mu.Unlock()

	l.controller.OnSuperMessage(msg, msg.client.room, msg.client, l)
}

func (l *Lobby) signalClientDisconnect(client *Client) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.controller.OnClientDisconnect(client.room, client.id, l)
}

func (l *Lobby) signalClientConnect(client *Client) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.controller.OnClientConnect(client, l)
}

func (l *Lobby) signalCreation() {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.controller.OnCreation(l)
}
