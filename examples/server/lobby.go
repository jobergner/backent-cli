package server

import (
	"sync"

	"github.com/jobergner/backent-cli/examples/action"
	"github.com/jobergner/backent-cli/examples/logging"
	"github.com/rs/zerolog/log"
)

type Lobby struct {
	mu     sync.Mutex
	Rooms  map[string]*Room
	config applicationConfig
}

func newLoginHandler(signals LobbySignals, actions action.Actions, sideEffects SideEffects, fps int) *Lobby {
	return &Lobby{
		Rooms: make(map[string]*Room),
		config: applicationConfig{
			signals:     signals,
			actions:     actions,
			sideEffects: sideEffects,
			fps:         fps,
		},
	}
}

func (l *Lobby) CreateRoom(name string) *Room {
	if room, ok := l.Rooms[name]; ok {
		log.Warn().Str(logging.RoomName, name).Msg("attempted to create room which already exists")
		return room
	}

	room := newRoom(l.config.actions, name)

	l.Rooms[name] = room

	room.Deploy(l.config.sideEffects, l.config.fps)

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
	if l.config.signals.OnSuperMessage == nil {
		return
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	log.Debug().Str(logging.ClientID, msg.client.id).Msg("OnSuperMessage")
	l.config.signals.OnSuperMessage(msg, msg.client.room, msg.client, l)
}

func (l *Lobby) signalClientDisconnect(client *Client) {
	if l.config.signals.OnClientDisconnect == nil {
		return
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	log.Debug().Str(logging.ClientID, client.id).Msg("OnClientDisconnect")
	l.config.signals.OnClientDisconnect(client.room, client.id, l)
}

func (l *Lobby) signalClientConnect(client *Client) {
	if l.config.signals.OnClientConnect == nil {
		return
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	log.Debug().Str(logging.ClientID, client.id).Msg("OnClientConnect")
	l.config.signals.OnClientConnect(client, l)
}
