package main

import (
	"errors"
	"log"
	"time"
)

type Room struct {
	clients              map[*Client]bool
	clientMessageChannel chan []byte
	registerChannel      chan *Client
	unregisterChannel    chan *Client
	state                *Engine
	actions              actions
}

func newRoom() *Room {
	return &Room{
		clientMessageChannel: make(chan []byte),
		registerChannel:      make(chan *Client),
		unregisterChannel:    make(chan *Client),
		clients:              make(map[*Client]bool),
		state:                newEngine(),
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
			return errors.New("client dropped")
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

func (r *Room) runBroadcastState() {
	ticker := time.NewTicker(time.Second)
	for {
		<-ticker.C
		err := r.broadcastStateToClients()
		if err != nil {
			log.Println(err)
		}
	}
}

func (r *Room) runWatchClientMessages() {
	for {
		msg := <-r.clientMessageChannel
		err := r.handleClientMessage(msg)
		if err != nil {
			log.Println(err)
		}
	}
}

func (r *Room) Deploy() {
	go r.runHandleConnections()
	go r.runWatchClientMessages()
	go r.runBroadcastState()
}
