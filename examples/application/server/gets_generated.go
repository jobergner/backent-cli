package state

import (
	"fmt"
)

const (
	MessageKindAction_addItemToPlayer MessageKind = "addItemToPlayer"
	MessageKindAction_movePlayer      MessageKind = "movePlayer"
	MessageKindAction_spawnZoneItems  MessageKind = "spawnZoneItems"
)

type MovePlayerParams struct {
	ChangeX float64  `json:"changeX"`
	ChangeY float64  `json:"changeY"`
	Player  PlayerID `json:"player"`
}

type AddItemToPlayerParams struct {
	Item    ItemID `json:"item"`
	NewName string `json:"newName"`
}

type SpawnZoneItemsParams struct {
	Items []ItemID `json:"items"`
}

type AddItemToPlayerResponse struct {
	PlayerPath string `json:"playerPath"`
}

type SpawnZoneItemsResponse struct {
	NewZoneItemPaths []string `json:"newZoneItemPaths"`
}

type AddItemToPlayerAction struct {
	Broadcast func(params AddItemToPlayerParams, engine *Engine, roomName, clientID string)
	Emit      func(params AddItemToPlayerParams, engine *Engine, roomName, clientID string) AddItemToPlayerResponse
}
type MovePlayerAction struct {
	Broadcast func(params MovePlayerParams, engine *Engine, roomName, clientID string)
	Emit      func(params MovePlayerParams, engine *Engine, roomName, clientID string)
}
type SpawnZoneItemsAction struct {
	Broadcast func(params SpawnZoneItemsParams, engine *Engine, roomName, clientID string)
	Emit      func(params SpawnZoneItemsParams, engine *Engine, roomName, clientID string) SpawnZoneItemsResponse
}

type Actions struct {
	AddItemToPlayer AddItemToPlayerAction
	MovePlayer      MovePlayerAction
	SpawnZoneItems  SpawnZoneItemsAction
}

// TODO logging should happen here
// TODO its kinda odd that we return a value and an error on fail
func (r *Room) processClientMessage(msg Message) (Message, error) {
	switch MessageKind(msg.Kind) {
	case MessageKindAction_addItemToPlayer:
		var params AddItemToPlayerParams
		err := params.UnmarshalJSON(msg.Content)
		if err != nil {
			return Message{msg.ID, MessageKindError, messageUnmarshallingError(msg.Content, err), msg.client}, err
		}
		if r.actions.AddItemToPlayer.Broadcast != nil {
			r.actions.AddItemToPlayer.Broadcast(params, r.state, r.name, msg.client.id)
		}
		if r.actions.AddItemToPlayer.Emit == nil {
			break
		}
		res := r.actions.AddItemToPlayer.Emit(params, r.state, r.name, msg.client.id)
		resContent, err := res.MarshalJSON()
		if err != nil {
			return Message{msg.ID, MessageKindError, responseMarshallingError(msg.Content, err), msg.client}, err
		}
		return Message{msg.ID, msg.Kind, resContent, msg.client}, nil
	case MessageKindAction_movePlayer:
		var params MovePlayerParams
		err := params.UnmarshalJSON(msg.Content)
		if err != nil {
			return Message{msg.ID, MessageKindError, messageUnmarshallingError(msg.Content, err), msg.client}, err
		}
		if r.actions.MovePlayer.Broadcast != nil {
			r.actions.MovePlayer.Broadcast(params, r.state, r.name, msg.client.id)
		}
		if r.actions.MovePlayer.Emit == nil {
			break
		}
		r.actions.MovePlayer.Emit(params, r.state, r.name, msg.client.id)
		return Message{ID: msg.ID}, nil
	case MessageKindAction_spawnZoneItems:
		var params SpawnZoneItemsParams
		err := params.UnmarshalJSON(msg.Content)
		if err != nil {
			return Message{msg.ID, MessageKindError, messageUnmarshallingError(msg.Content, err), msg.client}, err
		}
		if r.actions.SpawnZoneItems.Broadcast != nil {
			r.actions.SpawnZoneItems.Broadcast(params, r.state, r.name, msg.client.id)
		}
		if r.actions.SpawnZoneItems.Emit == nil {
			break
		}
		res := r.actions.SpawnZoneItems.Emit(params, r.state, r.name, msg.client.id)
		resContent, err := res.MarshalJSON()
		if err != nil {
			return Message{msg.ID, MessageKindError, responseMarshallingError(msg.Content, err), msg.client}, err
		}
		return Message{msg.ID, msg.Kind, resContent, msg.client}, nil
	default:
		return Message{msg.ID, MessageKindError, []byte("unknown message kind " + msg.Kind), msg.client}, fmt.Errorf("unknown message kind in: %s", printMessage(msg))
	}

	return Message{ID: msg.ID}, nil
}
