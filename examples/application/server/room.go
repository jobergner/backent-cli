package state

import (
	"fmt"
	"log"
	"time"
)

type Room struct {
	clients                 map[*Client]bool
	clientMessageChannel    chan Message
	pendingResponsesChannel chan Message
	registerChannel         chan *Client
	unregisterChannel       chan *Client
	incomingClients         map[*Client]bool
	state                   *Engine
	actions                 Actions
	sideEffects             SideEffects
	fps                     int
}

func newRoom(actions Actions, sideEffects SideEffects, fps int) *Room {
	return &Room{
		clients:                 make(map[*Client]bool),
		clientMessageChannel:    make(chan Message, 1024),
		pendingResponsesChannel: make(chan Message, 1024),
		unregisterChannel:       make(chan *Client),
		registerChannel:         make(chan *Client),
		incomingClients:         make(map[*Client]bool),
		state:                   newEngine(),
		sideEffects:             sideEffects,
		actions:                 actions,
		fps:                     fps,
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
	// we check if the client still exists because
	// closing a closed channel panics
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
			client.conn.Close()
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
			client.conn.Close()
			r.unregisterClient(client)
		}
	}

	return nil
}

func (r *Room) processFrame() {
Exit:
	for {
		select {
		case msg := <-r.clientMessageChannel:
			response, err := r.processClientMessage(msg)
			if err != nil {
				log.Println("error processing client message:", err)
				continue
			}

			// TODO: why?
			if response.client == nil {
				continue
			}

			select {
			case r.pendingResponsesChannel <- response:
			default:
				log.Printf("pending responses channel full, continuing with tick")
				// remaining messages in clientMessageChannel will be processed next tick
				break Exit
			}

		default:
			// clientMessageChannel is empty for now
			break Exit
		}
	}

	if r.sideEffects.OnFrameTick != nil {
		r.sideEffects.OnFrameTick(r.state)
	}
}

func (r *Room) publishPatch() error {
	tree := r.state.assembleTree(false)

	patchBytes, err := tree.MarshalJSON()
	if err != nil {
		return fmt.Errorf("error marshalling tree for patch: %s", err)
	}

	// TODO: if patch is empty -> find better way for evaluation
	if len(patchBytes) == 2 {
		return nil
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
Exit:
	for {
		select {
		case pendingResponse := <-r.pendingResponsesChannel:

			response, err := pendingResponse.MarshalJSON()
			if err != nil {
				log.Printf("error marshalling pending response message: %s", err)
				continue
			}

			select {
			case pendingResponse.client.messageChannel <- response:
			default:
				log.Printf("client's message buffer full -> dropping client %s", pendingResponse.client.id)
				pendingResponse.client.conn.Close()
				r.unregisterClient(pendingResponse.client)
			}

		default:
			// pendingResponsesChannel is empty for now
			break Exit
		}
	}
}

func (r *Room) tick() error {
	r.processFrame()

	r.handlePendingResponses()

	err := r.publishPatch()
	if err != nil {
		log.Println(err)
		// do not update state when we failed to publish it
		return err
	}

	r.state.UpdateState()

	return nil
}

func (r *Room) run() {
	ticker := time.NewTicker(time.Second / time.Duration(r.fps))

	// manipulation of the clients maps happens in this routine only
	// the channels are used for registering incoming clients
	// and unregistering clients whose connection dropped.
	// If reasons occur which require a client to be dropped in
	// this thread we do it by directly manipulating the maps.
	for {
		select {
		case client := <-r.registerChannel:
			r.registerClient(client)

		case client := <-r.unregisterChannel:
			r.unregisterClient(client)

		case <-ticker.C:
			err := r.tick()
			if err != nil {
				break
			}

			err = r.handleIncomingClients()
			if err != nil {
				log.Println(err)
			}

		}
	}

}

func (r *Room) Deploy() {
	if r.sideEffects.OnDeploy != nil {
		r.sideEffects.OnDeploy(r.state)
	}

	go r.run()
}
