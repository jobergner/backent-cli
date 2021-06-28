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

type Actions struct {
	AddItemToPlayer func(AddItemToPlayerParams, *Engine) AddItemToPlayerResponse
	MovePlayer      func(MovePlayerParams, *Engine)
	SpawnZoneItems  func(SpawnZoneItemsParams, *Engine) SpawnZoneItemsResponse
}

type SideEffects struct {
	OnDeploy    func(*Engine)
	OnFrameTick func(*Engine)
}

func (r *Room) processClientMessage(msg Message) (Message, error) {
	switch MessageKind(msg.Kind) {
	case MessageKindAction_addItemToPlayer:
		if r.actions.AddItemToPlayer == nil {
			break
		}
		var params AddItemToPlayerParams
		err := params.UnmarshalJSON(msg.Content)
		if err != nil {
			return Message{MessageKindError, messageUnmarshallingError(msg.Content, err), msg.client}, err
		}
		res := r.actions.AddItemToPlayer(params, r.state)
		resContent, err := res.MarshalJSON()
		if err != nil {
			return Message{MessageKindError, responseMarshallingError(msg.Content, err), msg.client}, err
		}
		return Message{msg.Kind, resContent, msg.client}, nil
	case MessageKindAction_movePlayer:
		if r.actions.MovePlayer == nil {
			break
		}
		var params MovePlayerParams
		err := params.UnmarshalJSON(msg.Content)
		if err != nil {
			return Message{MessageKindError, messageUnmarshallingError(msg.Content, err), msg.client}, err
		}
		r.actions.MovePlayer(params, r.state)
		return Message{}, nil
	case MessageKindAction_spawnZoneItems:
		if r.actions.SpawnZoneItems == nil {
			break
		}
		var params SpawnZoneItemsParams
		err := params.UnmarshalJSON(msg.Content)
		if err != nil {
			return Message{MessageKindError, messageUnmarshallingError(msg.Content, err), msg.client}, err
		}
		res := r.actions.SpawnZoneItems(params, r.state)
		resContent, err := res.MarshalJSON()
		if err != nil {
			return Message{MessageKindError, responseMarshallingError(msg.Content, err), msg.client}, err
		}
		return Message{msg.Kind, resContent, msg.client}, nil
	default:
		return Message{MessageKindError, []byte("unknown message kind " + msg.Kind), msg.client}, fmt.Errorf("unknown message kind in: %s", printMessage(msg))
	}

	return Message{}, nil
}
