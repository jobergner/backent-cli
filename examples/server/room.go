package main

import (
	"log"
	"time"
)

type Room struct {
	clients           map[*Client]bool
	broadcastChannel  chan []byte
	registerChannel   chan *Client
	unregisterChannel chan *Client
	state             int
}

func newRoom() *Room {
	return &Room{
		broadcastChannel:  make(chan []byte),
		registerChannel:   make(chan *Client),
		unregisterChannel: make(chan *Client),
		clients:           make(map[*Client]bool),
		state:             0,
	}
}

func (r *Room) registerClient(client *Client) {
	r.clients[client] = true
}

func (r *Room) unregisterClient(client *Client) {
	if _, ok := r.clients[client]; ok {
		close(client.messageChannel)
		delete(r.clients, client)
	}
}

func (r *Room) broadcastStateToClients() error {
	for client := range r.clients {
		select {
		case client.messageChannel <- []byte("true"):
		default:
			r.unregisterClient(client)
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

func (r *Room) runBroadcastRoomState() {
	ticker := time.NewTicker(time.Second)
	for {
		<-ticker.C
		err := r.broadcastStateToClients()
		if err != nil {
			log.Println(err)
		}
	}
}

func (r *Room) handleClientBroadcast(message []byte) error {
	r.state = r.state + 1
	log.Println(r.state)
	return nil
}

func (r *Room) runWatchClientBroadcasts() {
	for {
		message := <-r.broadcastChannel
		err := r.handleClientBroadcast(message)
		if err != nil {
			log.Println(err)
		}
	}
}

func (r *Room) Deploy() {
	go r.runHandleConnections()
	go r.runWatchClientBroadcasts()
	go r.runBroadcastRoomState()
}
