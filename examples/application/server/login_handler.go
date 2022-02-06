package state

import (
	"fmt"
	"sync"
)

type LoginHandler struct {
	mu             *sync.Mutex
	clients        map[*Client]struct{}
	rooms          map[string]*Room
	clientMessages chan Message
	signals        LoginSignals
}

func newLoginHandler(signals LoginSignals) *LoginHandler {
	// TODO thread safe
	return &LoginHandler{
		clients:        make(map[*Client]struct{}),
		rooms:          make(map[string]*Room),
		clientMessages: make(chan Message),
		signals:        signals,
	}
}

func (l *LoginHandler) CreateRoom(name string, actions Actions, sideEffects SideEffects, fps int) (*Room, error) {
	if _, ok := l.rooms[name]; ok {
		return nil, fmt.Errorf("room with name \"%s\" already exists", name)
	}

	room := newRoom(actions, sideEffects, fps)

	l.rooms[name] = room

	return room, nil
}

func (l *LoginHandler) DeleteRoom(name string) error {
	room, ok := l.rooms[name]
	if !ok {
		return fmt.Errorf("room with name \"%s\" does not exist", name)
	}

	room.terminate <- struct{}{}
	<-room.terminated

	return nil
}

func (l *LoginHandler) processMessageSync(msg Message) {
	l.mu.Lock()
	defer l.mu.Unlock()

	msg.client.room.mu.Lock()
	l.signals.OnGlobalMessage(msg, msg.client.room.state, msg.client, l)
	msg.client.room.mu.Unlock()
}
