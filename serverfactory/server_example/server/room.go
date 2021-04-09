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
	initRequests         map[*Client]bool
	state                *Engine
	actions              actions
}

func newRoom() *Room {
	return &Room{
		clientMessageChannel: make(chan message),
		registerChannel:      make(chan *Client),
		unregisterChannel:    make(chan *Client),
		clients:              make(map[*Client]bool),
		state:                newEngine(),
	}
}

func (r *Room) registerClient(client *Client) {
	r.clients[client] = true
}

func (r *Room) requestInit(client *Client) {
	r.initRequests[client] = true
}

func (r *Room) unregisterClient(client *Client) {
	if _, ok := r.clients[client]; ok {
		close(client.messageChannel)
		delete(r.clients, client)
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

func (r *Room) handleInitRequests() error {
	stateBytes, err := r.state.State.MarshalJSON()
	if err != nil {
		return err
	}

	for client := range r.initRequests {
		select {
		case client.messageChannel <- stateBytes:
		default:
			r.unregisterClient(client)
			// TODO what do?
			log.Println("client dropped")
		}
		delete(r.initRequests, client)
	}

	return nil
}

func (r *Room) runBroadcastPatch() {
	ticker := time.NewTicker(time.Second)
	for {
		<-ticker.C
		patchBytes, err := r.state.Patch.MarshalJSON()
		if err != nil {
			log.Println(err)
		}
		r.state.UpdateState()
		err = r.handleInitRequests()
		if err != nil {
			log.Println(err)
		}
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
	go r.runHandleConnections()
	go r.runWatchClientMessages()
	go r.runBroadcastPatch()
}
