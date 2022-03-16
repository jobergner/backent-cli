package server

import (
	"github.com/jobergner/backent-cli/examples/logging"
	"github.com/jobergner/backent-cli/examples/message"
	"github.com/rs/zerolog/log"
)

func (r *Room) processClientMessage(msg Message) (response Message) {
	switch msg.Kind {
	case message.MessageKindAction_addItemToPlayer:
		var params message.AddItemToPlayerParams
		err := params.UnmarshalJSON(msg.Content)
		if err != nil {
			log.Err(err).Str(logging.MessageKind, string(msg.Kind)).Str(logging.MessageContent, string(msg.Content)).Msg("failed unmarshalling params")
			return Message{msg.ID, message.MessageKindError, []byte("invalid message"), msg.client}
		}

		r.state.BroadcastingClientID = msg.client.id
		r.controller.AddItemToPlayerBroadcast(params, r.state, r.name, msg.client.id)
		r.state.BroadcastingClientID = ""

		res := r.controller.AddItemToPlayerEmit(params, r.state, r.name, msg.client.id)
		resContent, err := res.MarshalJSON()
		if err != nil {
			log.Err(err).Str(logging.MessageKind, string(msg.Kind)).Msg("failed marshalling response content")
			return Message{msg.ID, message.MessageKindError, []byte("invalid message"), msg.client}
		}

		return Message{msg.ID, msg.Kind, resContent, msg.client}
	case message.MessageKindAction_movePlayer:
		var params message.MovePlayerParams
		err := params.UnmarshalJSON(msg.Content)
		if err != nil {
			log.Err(err).Str(logging.MessageKind, string(msg.Kind)).Str(logging.MessageContent, string(msg.Content)).Msg("failed unmarshalling params")
			return Message{msg.ID, message.MessageKindError, []byte("invalid message"), msg.client}
		}

		r.state.BroadcastingClientID = msg.client.id
		r.controller.MovePlayerBroadcast(params, r.state, r.name, msg.client.id)
		r.state.BroadcastingClientID = ""

		r.controller.MovePlayerEmit(params, r.state, r.name, msg.client.id)
		return Message{ID: msg.ID, Kind: message.MessageKindNoResponse}
	case message.MessageKindAction_spawnZoneItems:
		var params message.SpawnZoneItemsParams
		err := params.UnmarshalJSON(msg.Content)
		if err != nil {
			log.Err(err).Str(logging.MessageKind, string(msg.Kind)).Str(logging.MessageContent, string(msg.Content)).Msg("failed unmarshalling params")
			return Message{msg.ID, message.MessageKindError, []byte("invalid message"), msg.client}
		}

		r.state.BroadcastingClientID = msg.client.id
		r.controller.SpawnZoneItemsBroadcast(params, r.state, r.name, msg.client.id)
		r.state.BroadcastingClientID = ""

		res := r.controller.SpawnZoneItemsEmit(params, r.state, r.name, msg.client.id)
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
