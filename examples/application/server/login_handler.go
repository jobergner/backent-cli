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
	actions        Actions
	signals        LoginSignals
	sideEffects    SideEffects
	fps            int
}

func newLoginHandler(signals LoginSignals, actions Actions, sideEffects SideEffects, fps int) *LoginHandler {
	// TODO thread safe
	return &LoginHandler{
		mu:             &sync.Mutex{},
		clients:        make(map[*Client]struct{}),
		Rooms:          make(map[string]*Room),
		clientMessages: make(chan Message),
		signals:        signals,
		actions:        actions,
		sideEffects:    sideEffects,
		fps:            fps,
	}
}

func (l *LoginHandler) CreateRoom(name string) (*Room, error) {
	if _, ok := l.Rooms[name]; ok {
		return nil, fmt.Errorf("room with name \"%s\" already exists", name)
	}

	room := newRoom(l.actions)

	l.Rooms[name] = room

	room.Deploy(l.sideEffects, l.fps)

	return room, nil
}

func (l *LoginHandler) DeleteRoom(name string) error {
	room, ok := l.Rooms[name]
	if !ok {
		return fmt.Errorf("room with name \"%s\" does not exist", name)
	}

	room.mode = RoomModeTerminating

	delete(l.Rooms, name)

	return nil
}

func (l *LoginHandler) addClient(client *Client) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.clients[client] = struct{}{}
}

func (l *LoginHandler) deleteClient(client *Client) {
	l.mu.Lock()
	defer l.mu.Unlock()

	delete(l.clients, client)
}

func (l *LoginHandler) processMessageSync(msg Message) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.signals.OnSuperMessage(msg, msg.client.room.state, msg.client, l)
}

// TODO rename to signalClientDisconnect
func (l *LoginHandler) handleClientDisconnect(client *Client) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.signals.OnClientDisconnect(client.room.state, client.id, l)
}

func (l *LoginHandler) handleClientConnect(client *Client) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.signals.OnClientConnect(client, l)
}

// NOTE: youll probably want to create the player data on connect WITH the client ID
// this way any action sent will automatically include an identifier for the data belonging to
// this client

// ClientSwitch -> run multiple rooms on one server
// Message HTTP endpoint with acces too all information. Find a way that is somewhat clean
// 	- fetch information from all rooms (information for a eg. lobby)
// 	- create/delete rooms
//  - populate rooms with data
//

// OnSuperMessage:
// client can leave room, join room
// client can send message to itself/other clients cross room

// OnClientConnect:
// - client can join room. use room.AlterState to initialize state

// OnClientDisconnect:
// - remove data from room

// Client:
// LeaveRoom()
// JoinRoom (*Room)
// ID() Client.ID
// SendMessage()
// Room()

// TODO do clients live in rooms and loginHandler simultaneously? prob yes
// what do we do when deleting room? notify handler of all clients losing their rooms
// do clients of deleted room get their client.room removed?
