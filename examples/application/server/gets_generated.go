package state

import (
	"fmt"
	"net/http"
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

type actions struct {
	addItemToPlayer func(AddItemToPlayerParams, *Engine) AddItemToPlayerResponse
	movePlayer      func(MovePlayerParams, *Engine)
	spawnZoneItems  func(SpawnZoneItemsParams, *Engine) SpawnZoneItemsResponse
}

func (r *Room) processClientMessage(msg Message) (Message, error) {
	switch MessageKind(msg.Kind) {
	case MessageKindAction_addItemToPlayer:
		var params AddItemToPlayerParams
		err := params.UnmarshalJSON(msg.Content)
		if err != nil {
			return Message{MessageKindError, messageUnmarshallingError(msg.Content, err), msg.client}, err
		}
		res := r.actions.addItemToPlayer(params, r.state)
		resContent, err := res.MarshalJSON()
		if err != nil {
			return Message{MessageKindError, responseMarshallingError(msg.Content, err), msg.client}, err
		}
		return Message{msg.Kind, resContent, msg.client}, nil
	case MessageKindAction_movePlayer:
		var params MovePlayerParams
		err := params.UnmarshalJSON(msg.Content)
		if err != nil {
			return Message{MessageKindError, messageUnmarshallingError(msg.Content, err), msg.client}, err
		}
		r.actions.movePlayer(params, r.state)
		return Message{}, nil
	case MessageKindAction_spawnZoneItems:
		var params SpawnZoneItemsParams
		err := params.UnmarshalJSON(msg.Content)
		if err != nil {
			return Message{MessageKindError, messageUnmarshallingError(msg.Content, err), msg.client}, err
		}
		res := r.actions.spawnZoneItems(params, r.state)
		resContent, err := res.MarshalJSON()
		if err != nil {
			return Message{MessageKindError, responseMarshallingError(msg.Content, err), msg.client}, err
		}
		return Message{msg.Kind, resContent, msg.client}, nil
	default:
		return Message{MessageKindError, []byte("unknown message kind " + msg.Kind), msg.client}, fmt.Errorf("unknown message kind in: %s", printMessage(msg))
	}
}

func Start(
	addItemToPlayer func(AddItemToPlayerParams, *Engine) AddItemToPlayerResponse,
	movePlayer func(MovePlayerParams, *Engine),
	spawnZoneItems func(SpawnZoneItemsParams, *Engine) SpawnZoneItemsResponse,
	onDeploy func(*Engine),
	onFrameTick func(*Engine),
) error {
	a := actions{addItemToPlayer, movePlayer, spawnZoneItems}
	setupRoutes(a, onDeploy, onFrameTick)
	err := http.ListenAndServe(":8080", nil)
	return err
}
