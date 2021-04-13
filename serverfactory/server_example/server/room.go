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
	state                *Engine
	actions              actions
	onDeploy             func(*Engine)
	onFrameTick          func(*Engine)
}

func newRoom(a actions, onDeploy func(*Engine), onFrameTick func(*Engine)) *Room {
	return &Room{
		clients:              make(map[*Client]bool),
		clientMessageChannel: make(chan message),
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
	if _, ok := r.clients[client]; ok {
		close(client.messageChannel)
		delete(r.clients, client)
		delete(r.incomingClients, client)
	}
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
	stateBytes, err := r.state.State.MarshalJSON()
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

func (r *Room) runProcessingFrames() {
	ticker := time.NewTicker(time.Second)
	for {
		<-ticker.C
		patchBytes, err := r.state.Patch.MarshalJSON()
		if err != nil {
			log.Println(err)
		}
//TODO: state is being manipulated by actions and here (2 different routines) 
		r.state.UpdateState()
		err = r.answerInitRequests()
		if err != nil {
			log.Println(err)
		}
		r.promoteIncomingClients()
		err = r.broadcastPatchToClients(patchBytes)
		if err != nil {
			log.Println(err)
		}
	}
}

func (r *Room) runWatchClientMessages() {
	for {
		msg := <-r.clientMessageChannel
		err := r.processClientMessage(msg)
		if err != nil {
			log.Println(err)
		}
	}
}

func (r *Room) Deploy() {
	r.onDeploy(r.state)
	go r.runHandleConnections()
	go r.runWatchClientMessages()
	go r.runProcessingFrames()
}
