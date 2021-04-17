package state

import (
	"errors"
	"log"
	"net/http"
)

const (
	messageKindAction_addItemToPlayer messageKind = 1
	messageKindAction_movePlayer      messageKind = 2
	messageKindAction_spawnZoneItems  messageKind = 3
)

type MovePlayerParams struct {
	ChangeX  float64  `json:"changeX"`
	ChangeY  float64  `json:"changeY"`
	PlayerID PlayerID `json:"playerID"`
}

type AddItemToPlayerParams struct {
	Item     Item     `json:"item"`
	PlayerID PlayerID `json:"playerID"`
}

type SpawnZoneItemsParams struct {
	Items []Item `json:"items"`
}

type actions struct {
	addItemToPlayer func(AddItemToPlayerParams, *Engine)
	movePlayer      func(MovePlayerParams, *Engine)
	spawnZoneItems  func(SpawnZoneItemsParams, *Engine)
}

func (r *Room) processClientMessage(msg message) error {
	switch messageKind(msg.Kind) {
	case messageKindAction_addItemToPlayer:
		var params AddItemToPlayerParams
		err := params.UnmarshalJSON(msg.Content)
		if err != nil {
			return err
		}
		r.actions.addItemToPlayer(params, r.state)
	case messageKindAction_movePlayer:
		var params MovePlayerParams
		err := params.UnmarshalJSON(msg.Content)
		if err != nil {
			return err
		}
		r.actions.movePlayer(params, r.state)
	case messageKindAction_spawnZoneItems:
		var params SpawnZoneItemsParams
		err := params.UnmarshalJSON(msg.Content)
		if err != nil {
			return err
		}
		r.actions.spawnZoneItems(params, r.state)
	default:
		return errors.New("unknown message kind")
	}

	return nil
}

func Start(
	addItemToPlayer func(AddItemToPlayerParams, *Engine),
	movePlayer func(MovePlayerParams, *Engine),
	spawnZoneItems func(SpawnZoneItemsParams, *Engine),
	onDeploy func(*Engine),
	onFrameTick func(*Engine),
) {
	a := actions{addItemToPlayer, movePlayer, spawnZoneItems}
	setupRoutes(a, onDeploy, onFrameTick)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
