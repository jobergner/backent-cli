package state

import (
	"fmt"
	"log"
	"time"
)

type Room struct {
	clients              map[*Client]bool
	clientMessageChannel chan message
	registerChannel      chan *Client
	unregisterChannel    chan *Client
	incomingClients      map[*Client]bool
	pendingResponses     []message
	state                *Engine
	actions              actions
	onDeploy             func(*Engine)
	onFrameTick          func(*Engine)
}

func newRoom(a actions, onDeploy func(*Engine), onFrameTick func(*Engine)) *Room {
	return &Room{
		clients:              make(map[*Client]bool),
		clientMessageChannel: make(chan message, 1024),
		registerChannel:      make(chan *Client),
		unregisterChannel:    make(chan *Client),
		incomingClients:      make(map[*Client]bool),
		state:                newEngine(),
		onDeploy:             onDeploy,
		onFrameTick:          onFrameTick,
		actions:              a,
	}
}

func (r *Room) registerClient(client *Client) {
	r.incomingClients[client] = true
}

func (r *Room) promoteIncomingClient(client *Client) {
	r.clients[client] = true
	delete(r.incomingClients, client)
}

func (r *Room) unregisterClient(client *Client) {
	log.Printf("unregistering client %s", client.id)
	// TODO: panic close of closed channel?
	close(client.messageChannel)
	delete(r.clients, client)
	delete(r.incomingClients, client)
}

func (r *Room) broadcastPatchToClients(patchBytes []byte) {
	for client := range r.clients {
		select {
		case client.messageChannel <- patchBytes:
		default:
			log.Printf("client's message buffer full -> dropping client %s", client.id)
			r.unregisterClient(client)
		}
	}
}

func (r *Room) handleIncomingClients() error {
	if len(r.incomingClients) == 0 {
		return nil
	}
	tree := r.state.assembleTree(true)
	stateBytes, err := tree.MarshalJSON()
	if err != nil {
		return fmt.Errorf("error marshalling tree for init request: %s", err)
	}

	for client := range r.incomingClients {
		select {
		case client.messageChannel <- stateBytes:
			r.promoteIncomingClient(client)
		default:
			log.Printf("client's message buffer full -> dropping client %s", client.id)
			r.unregisterClient(client)
		}
	}

	return nil
}

func (r *Room) processFrame() error {
Exit:
	for {
		select {
		case msg := <-r.clientMessageChannel:
			response, err := r.processClientMessage(msg)
			if err != nil {
				log.Println("error processing client message:", err)
				continue
			}
			if response.client == nil {
				continue
			}
			r.pendingResponses = append(r.pendingResponses, response)
		default:
			break Exit
		}
	}

	r.onFrameTick(r.state)

	return nil
}

func (r *Room) publishPatch() error {
	tree := r.state.assembleTree(false)
	patchBytes, err := tree.MarshalJSON()
	if err != nil {
		return fmt.Errorf("error marshalling tree for patch: %s", err)
	}
	r.broadcastPatchToClients(patchBytes)
	return nil
}

func (r *Room) handlePendingResponses() {
	for _, pendingResponse := range r.pendingResponses {
		select {
		case pendingResponse.client.messageChannel <- pendingResponse.Content:
		default:
			log.Printf("client's message buffer full -> dropping client %s", pendingResponse.client.id)
			r.unregisterClient(pendingResponse.client)
		}
	}
	r.pendingResponses = r.pendingResponses[:0]
}

func (r *Room) process() {
	err := r.processFrame()
	if err != nil {
		log.Println(err)
	}
	err = r.publishPatch()
	if err != nil {
		log.Println(err)
	}
	r.state.UpdateState()
	err = r.handleIncomingClients()
	if err != nil {
		log.Println(err)
	}
	r.handlePendingResponses()
}

func (r *Room) run() {
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case client := <-r.registerChannel:
			r.registerClient(client)
		case client := <-r.unregisterChannel:
			r.unregisterClient(client)
		case <-ticker.C:
			r.process()
		}
	}
}

func (r *Room) Deploy() {
	r.onDeploy(r.state)
	go r.run()
}
