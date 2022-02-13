package state

import (
	"log"
	"sync"
	"time"
)

type RoomMode int

const (
	RoomModeIdle RoomMode = iota
	RoomModeRunning
	RoomModeTerminating
)

type Room struct {
	name            string
	mu              *sync.Mutex
	clients         map[*Client]struct{}
	incomingClients map[*Client]struct{}
	state           *Engine
	actions         Actions
	mode            RoomMode // TODO implement usage
}

func newRoom(actions Actions) *Room {
	return &Room{
		mu:              &sync.Mutex{},
		clients:         make(map[*Client]struct{}),
		incomingClients: make(map[*Client]struct{}),
		state:           newEngine(),
		actions:         actions,
	}
}

func (r *Room) Name() string {
	return r.name
}

func (r *Room) AlterState(fn func(*Room)) {
	r.mu.Lock()
	fn(r)
	r.mu.Unlock()
}

func (r *Room) processMessageSync(msg Message) {
	r.mu.Lock()
	defer r.mu.Unlock()

	response, err := r.processClientMessage(msg)
	if err != nil {
		return
	}

	// actions may not have a response
	// which means `response` is empty here
	// and can be skipped
	if response.client == nil {
		return
	}

	responseBytes, err := response.MarshalJSON()
	if err != nil {
		log.Printf("error marshalling pending response message: %s", err)
		return
	}

	select {
	case response.client.messageChannel <- responseBytes:
	default:
		log.Printf("client's message buffer full -> dropping client %s", response.client.id)
		// TODO need other solution for this. client sould probably also be kicked from login server
		// response.client.conn.Close()
		r.unregisterClientAsync(response.client)
	}
}

func (r *Room) registerClientSync(client *Client) {
	r.mu.Lock()
	defer r.mu.Unlock()

	client.room = r
	r.incomingClients[client] = struct{}{}
}

func (r *Room) unregisterClientSync(client *Client) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.unregisterClientAsync(client)
}

func (r *Room) run(sideEffects SideEffects, fps int) {
	ticker := time.NewTicker(time.Second / time.Duration(fps))

	for {
		<-ticker.C
		r.tickSync(sideEffects)
		// TODO gracefully shutting down?
	}
}

func (r *Room) Deploy(sideEffects SideEffects, fps int) {
	if sideEffects.OnDeploy != nil {
		r.mu.Lock()
		sideEffects.OnDeploy(r.state)
		r.mu.Unlock()
	}

	go r.run(sideEffects, fps)
}
