package server

import (
	"github.com/jobergner/backent-cli/examples/logging"
	"github.com/jobergner/backent-cli/examples/message"
	"github.com/rs/zerolog/log"
)

func (r *Room) tick() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.controller.OnFrameTick(r.state)

	err := r.publishPatch()
	if err != nil {
		return
	}

	r.state.UpdateState()

	r.handleIncomingClients()
}

func (r *Room) publishPatch() error {
	if r.state.Patch.IsEmpty() {
		return nil
	}

	r.state.AssembleUpdateTree()

	patchBytes, err := r.state.Tree.MarshalJSON()
	if err != nil {
		log.Err(err).Msg("failed marshalling patch")
		return err
	}

	stateUpdateMsg := Message{
		Kind:    message.MessageKindUpdate,
		Content: patchBytes,
	}

	stateUpdateBytes, err := stateUpdateMsg.MarshalJSON()
	if err != nil {
		log.Err(err).Str(logging.MessageKind, string(stateUpdateMsg.Kind)).Msg("failed marshalling message")
		return err
	}

	r.broadcastPatchToClients(stateUpdateBytes)

	return nil
}

func (r *Room) broadcastPatchToClients(stateUpdateBytes []byte) {
	for client := range r.clients.clients {

		select {
		case client.messageChannel <- stateUpdateBytes:
		default:
			log.Warn().Str(logging.ClientID, client.id).Msg(logging.ClientBufferFull)
			client.closeConnection(logging.ClientBufferFull)
		}

	}
}

func (r *Room) handleIncomingClients() {
	if len(r.clients.incomingClients) == 0 {
		return
	}

	stateBytes, err := r.state.State.MarshalJSON()
	if err != nil {
		log.Err(err).Msg("failed marshalling state")
		return
	}

	currentStateMsg := Message{
		Kind:    message.MessageKindCurrentState,
		Content: stateBytes,
	}

	currentStateMessageBytes, err := currentStateMsg.MarshalJSON()
	if err != nil {
		log.Err(err).Str(logging.MessageKind, string(currentStateMsg.Kind)).Msg("failed marshalling message")
		return
	}

	for client := range r.clients.incomingClients {
		select {
		case client.messageChannel <- currentStateMessageBytes:
			r.clients.promote(client)
		default:
			log.Warn().Str(logging.ClientID, client.id).Msg(logging.ClientBufferFull)
			client.closeConnection(logging.ClientBufferFull)
		}
	}
}
