package state

import (
	"fmt"
	"sync"
)

type LoginHandler struct {
	mu             *sync.Mutex
	clients        map[*Client]struct{}
	Rooms          map[string]*Room
	clientMessages chan Message
	signals        LoginSignals
}

func newLoginHandler(signals LoginSignals) *LoginHandler {
	// TODO thread safe
	return &LoginHandler{
		clients:        make(map[*Client]struct{}),
		Rooms:          make(map[string]*Room),
		clientMessages: make(chan Message),
		signals:        signals,
	}
}

func (l *LoginHandler) createRoom(name string, actions Actions, sideEffects SideEffects, fps int) (*Room, error) {
	if _, ok := l.Rooms[name]; ok {
		return nil, fmt.Errorf("room with name \"%s\" already exists", name)
	}

	room := newRoom(actions)

	l.Rooms[name] = room

	room.Deploy(sideEffects, fps)

	return room, nil
}

func (l *LoginHandler) deleteRoom(name string) error {
	room, ok := l.Rooms[name]
	if !ok {
		return fmt.Errorf("room with name \"%s\" does not exist", name)
	}

	return nil
}

func (l *LoginHandler) processMessageSync(msg Message) {
	l.mu.Lock()
	defer l.mu.Unlock()

	msg.client.room.mu.Lock()

	if msg.client.room == nil {
		l.signals.OnGlobalMessage(msg, nil, msg.client, l)
	} else {
		l.signals.OnGlobalMessage(msg, msg.client.room.state, msg.client, l)
	}

	msg.client.room.mu.Unlock()
}

func (l *LoginHandler) handleClientDisconnect(client *Client) {
	l.mu.Lock()
	defer l.mu.Unlock()

	client.room.mu.Lock()

	if client.room == nil {
		l.signals.OnClientDisconnect(nil, client.id, l)
	} else {
		l.signals.OnClientDisconnect(client.room.state, client.id, l)
	}

	client.room.mu.Unlock()
}

func (l *LoginHandler) handleClientConnect(client *Client) {
	l.mu.Lock()
	defer l.mu.Unlock()

	client.room.mu.Lock()
	l.signals.OnClientConnect(client, l)
	client.room.mu.Unlock()
}

// TODO do clients live in rooms and loginHandler simultaneously? prob yes
// what do we do when deleting room? notify handler of all clients losing their rooms
// do clients of deleted room get their client.room removed?

// loginHandler:
// CreateRoom
// DeleteRoom

// Client:
// LeaveRoom()
// JoinRoom (*Room)
// ID() Client.ID
// SendMessage()
// Room()
