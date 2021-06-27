package state

import (
	"fmt"
	"log"
	"time"
)

type Room struct {
	clients              map[*Client]bool
	clientMessageChannel chan Message
	registerChannel      chan *Client
	unregisterChannel    chan *Client
	incomingClients      map[*Client]bool
	pendingResponses     []Message
	state                *Engine
	actions              actions
	onDeploy             func(*Engine)
	onFrameTick          func(*Engine)
}

func newRoom(a actions, onDeploy func(*Engine), onFrameTick func(*Engine)) *Room {
	return &Room{
		clients:              make(map[*Client]bool),
		clientMessageChannel: make(chan Message, 1024),
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
	if _, ok := r.clients[client]; ok {
		log.Printf("unregistering client %s", client.id)
		close(client.messageChannel)
		delete(r.clients, client)
	} else if _, ok := r.incomingClients[client]; ok {
		log.Printf("unregistering incoming client %s", client.id)
		close(client.messageChannel)
		delete(r.incomingClients, client)
	}
}

func (r *Room) broadcastPatchToClients(stateUpdateBytes []byte) {
	for client := range r.clients {
		select {
		case client.messageChannel <- stateUpdateBytes:
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

	currentStateMsg := Message{
		Kind:    MessageKindCurrentState,
		Content: stateBytes,
	}
	response, err := currentStateMsg.MarshalJSON()
	if err != nil {
		return fmt.Errorf("error marshalling response message for init request: %s", err)
	}

	for client := range r.incomingClients {
		select {
		case client.messageChannel <- response:
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
			if len(msg.Content) == 0 {
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

	stateUpdateMsg := Message{
		Kind:    MessageKindUpdate,
		Content: patchBytes,
	}
	stateUpdateBytes, err := stateUpdateMsg.MarshalJSON()
	if err != nil {
		return fmt.Errorf("error marshalling state update message: %s", err)
	}

	r.broadcastPatchToClients(stateUpdateBytes)
	return nil
}

func (r *Room) handlePendingResponses() {
	for _, pendingResponse := range r.pendingResponses {
		response, err := pendingResponse.MarshalJSON()
		if err != nil {
			log.Printf("error marshalling response message pending response: %s", err)
			continue
		}
		select {
		case pendingResponse.client.messageChannel <- response:
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
