package state

import (
	"fmt"
	"net/http"
)

const (
	messageKindAction_addItemToPlayer messageKind = "addItemToPlayer"
	messageKindAction_movePlayer      messageKind = "movePlayer"
	messageKindAction_spawnZoneItems  messageKind = "spawnZoneItems"
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

func (r *Room) processClientMessage(msg message) (message, error) {
	switch messageKind(msg.Kind) {
	case messageKindAction_addItemToPlayer:
		var params AddItemToPlayerParams
		err := params.UnmarshalJSON(msg.Content)
		if err != nil {
			return message{}, err
		}
		res := r.actions.addItemToPlayer(params, r.state)
		resContent, err := res.MarshalJSON()
		if err != nil {
			return message{}, err
		}
		return message{msg.Kind, resContent, msg.client}, nil
	case messageKindAction_movePlayer:
		var params MovePlayerParams
		err := params.UnmarshalJSON(msg.Content)
		if err != nil {
			return message{}, err
		}
		r.actions.movePlayer(params, r.state)
		return message{}, nil
	case messageKindAction_spawnZoneItems:
		var params SpawnZoneItemsParams
		err := params.UnmarshalJSON(msg.Content)
		if err != nil {
			return message{}, err
		}
		res := r.actions.spawnZoneItems(params, r.state)
		resContent, err := res.MarshalJSON()
		if err != nil {
			return message{}, err
		}
		return message{msg.Kind, resContent, msg.client}, nil
	default:
		return message{}, fmt.Errorf("unknown message kind in: %s", printMessage(msg))
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
