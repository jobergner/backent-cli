package state

import (
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
		clientMessageChannel: make(chan message, 264),
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

func (r *Room) unregisterClient(client *Client) {
	close(client.messageChannel)
	delete(r.clients, client)
	delete(r.incomingClients, client)
}

func (r *Room) broadcastPatchToClients(patchBytes []byte) error {
	for client := range r.clients {
		select {
		case client.messageChannel <- patchBytes:
		default:
			r.unregisterClient(client)
			// TODO what do?
			log.Println("client dropped")
		}
	}

	return nil
}

func (r *Room) runHandleConnections() {
	for {
		select {
		case client := <-r.registerChannel:
			r.registerClient(client)
		case client := <-r.unregisterChannel:
			r.unregisterClient(client)
		}
	}
}

func (r *Room) answerInitRequests() error {
	if len(r.incomingClients) == 0 {
		return nil
	}
	tree := r.state.assembleTree(true)
	stateBytes, err := tree.MarshalJSON()
	if err != nil {
		return err
	}

	for client := range r.incomingClients {
		select {
		case client.messageChannel <- stateBytes:
		default:
			r.unregisterClient(client)
			// TODO what do?
			log.Println("client dropped")
		}
	}

	return nil
}

func (r *Room) promoteIncomingClients() {
	for client := range r.incomingClients {
		r.clients[client] = true
		delete(r.incomingClients, client)
	}
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
	r.state.walkTree()
	tree := r.state.assembleTree(false)
	patchBytes, err := tree.MarshalJSON()
	if err != nil {
		return err
	}
	err = r.broadcastPatchToClients(patchBytes)
	if err != nil {
		return err
	}
	return nil
}

func (r *Room) handleIncomingClients() error {
	err := r.answerInitRequests()
	if err != nil {
		return err
	}
	r.promoteIncomingClients()
	return nil
}

func (r *Room) handlePendingResponses() {
	for _, pendingResponse := range r.pendingResponses {
		select {
		case pendingResponse.client.messageChannel <- pendingResponse.Content:
		default:
			r.unregisterClient(pendingResponse.client)
			// TODO what do?
			log.Println("client dropped")
		}
	}
}

func (r *Room) runProcessingFrames() {
	ticker := time.NewTicker(time.Second)
	for {
		<-ticker.C
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
}

func (r *Room) Deploy() {
	r.onDeploy(r.state)
	go r.runHandleConnections()
	go r.runProcessingFrames()
}
