package server

import (
	"github.com/jobergner/backent-cli/examples/action"
	"github.com/jobergner/backent-cli/examples/logging"
	"github.com/jobergner/backent-cli/examples/message"
	"github.com/rs/zerolog/log"
)

func (r *Room) processClientMessage(msg Message) (response Message) {
	switch msg.Kind {
	case action.MessageKindAction_addItemToPlayer:
		var params action.AddItemToPlayerParams
		err := params.UnmarshalJSON(msg.Content)
		if err != nil {
			log.Err(err).Str(logging.MessageKind, string(msg.Kind)).Str(logging.MessageContent, string(msg.Content)).Msg("failed unmarshalling params")
			return Message{msg.ID, message.MessageKindError, []byte("invalid message"), msg.client}
		}
		if r.actions.AddItemToPlayer.Broadcast != nil {
			r.state.BroadcastingClientID = msg.client.id
			r.actions.AddItemToPlayer.Broadcast(params, r.state, r.name, msg.client.id)
			r.state.BroadcastingClientID = ""
		}
		if r.actions.AddItemToPlayer.Emit == nil {
			return Message{ID: msg.ID, Kind: message.MessageKindNoResponse}
		}
		res := r.actions.AddItemToPlayer.Emit(params, r.state, r.name, msg.client.id)
		resContent, err := res.MarshalJSON()
		if err != nil {
			log.Err(err).Str(logging.MessageKind, string(msg.Kind)).Msg("failed marshalling response content")
			return Message{msg.ID, message.MessageKindError, []byte("invalid message"), msg.client}
		}
		return Message{msg.ID, msg.Kind, resContent, msg.client}
	case action.MessageKindAction_movePlayer:
		var params action.MovePlayerParams
		err := params.UnmarshalJSON(msg.Content)
		if err != nil {
			log.Err(err).Str(logging.MessageKind, string(msg.Kind)).Str(logging.MessageContent, string(msg.Content)).Msg("failed unmarshalling params")
			return Message{msg.ID, message.MessageKindError, []byte("invalid message"), msg.client}
		}
		if r.actions.MovePlayer.Broadcast != nil {
			r.state.BroadcastingClientID = msg.client.id
			r.actions.MovePlayer.Broadcast(params, r.state, r.name, msg.client.id)
			r.state.BroadcastingClientID = ""
		}
		if r.actions.MovePlayer.Emit == nil {
			return Message{ID: msg.ID, Kind: message.MessageKindNoResponse}
		}
		r.actions.MovePlayer.Emit(params, r.state, r.name, msg.client.id)
		return Message{ID: msg.ID, Kind: message.MessageKindNoResponse}
	case action.MessageKindAction_spawnZoneItems:
		var params action.SpawnZoneItemsParams
		err := params.UnmarshalJSON(msg.Content)
		if err != nil {
			log.Err(err).Str(logging.MessageKind, string(msg.Kind)).Str(logging.MessageContent, string(msg.Content)).Msg("failed unmarshalling params")
			return Message{msg.ID, message.MessageKindError, []byte("invalid message"), msg.client}
		}
		if r.actions.SpawnZoneItems.Broadcast != nil {
			r.state.BroadcastingClientID = msg.client.id
			r.actions.SpawnZoneItems.Broadcast(params, r.state, r.name, msg.client.id)
			r.state.BroadcastingClientID = ""
		}
		if r.actions.SpawnZoneItems.Emit == nil {
			return Message{ID: msg.ID, Kind: message.MessageKindNoResponse}
		}
		res := r.actions.SpawnZoneItems.Emit(params, r.state, r.name, msg.client.id)
		resContent, err := res.MarshalJSON()
		if err != nil {
			log.Err(err).Str(logging.MessageKind, string(msg.Kind)).Msg("failed marshalling response content")
			return Message{msg.ID, message.MessageKindError, []byte("invalid message"), msg.client}
		}
		return Message{msg.ID, msg.Kind, resContent, msg.client}
	default:
		log.Warn().Str(logging.MessageKind, string(msg.Kind)).Msg("unknown message kind")
		return Message{msg.ID, message.MessageKindError, []byte("invalid message"), msg.client}
	}
}
