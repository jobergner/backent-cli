package state

import (
	"fmt"
	"log"
)

func (r *Room) tickSync(sideEffects SideEffects) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if sideEffects.OnFrameTick != nil {
		sideEffects.OnFrameTick(r.state)
	}

	err := r.publishPatch()
	if err != nil {
		return
	}

	r.state.UpdateState()

	err = r.handleIncomingClients()
	if err != nil {
		return
	}
}

func (r *Room) publishPatch() error {
	patchBytes, err := r.state.Patch.MarshalJSON()
	if err != nil {
		return fmt.Errorf("error marshalling patch: %s", err)
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

func (r *Room) broadcastPatchToClients(stateUpdateBytes []byte) {
	for client := range r.clients {
		select {
		case client.messageChannel <- stateUpdateBytes:
		default:
			log.Printf("client's message buffer full -> dropping client %s", client.id)
			r.unregisterClientAsync(client)
		}
	}
}

func (r *Room) handleIncomingClients() error {
	if len(r.incomingClients) == 0 {
		return nil
	}

	stateBytes, err := r.state.State.MarshalJSON()
	if err != nil {
		return fmt.Errorf("error marshalling tree for init request: %s", err)
	}

	currentStateMsg := Message{
		Kind:    MessageKindCurrentState,
		Content: stateBytes,
	}

	currentStateMessageBytes, err := currentStateMsg.MarshalJSON()
	if err != nil {
		return fmt.Errorf("error marshalling message for init request: %s", err)
	}

	for client := range r.incomingClients {
		select {
		case client.messageChannel <- currentStateMessageBytes:
			r.promoteClientAsync(client)
		default:
			log.Printf("client's message buffer full -> dropping client %s", client.id)
			r.unregisterClientAsync(client)
		}
	}

	return nil
}

func (r *Room) promoteClientAsync(client *Client) {
	r.clients[client] = struct{}{}
	delete(r.incomingClients, client)
}

func (r *Room) unregisterClientAsync(client *Client) {
	if _, ok := r.clients[client]; ok {

		log.Printf("unregistering client %s", client.id)

		delete(r.clients, client)

		client.conn.Close()
	} else if _, ok := r.incomingClients[client]; ok {

		log.Printf("unregistering incoming client %s", client.id)

		delete(r.incomingClients, client)

		client.conn.Close()
	}
}
