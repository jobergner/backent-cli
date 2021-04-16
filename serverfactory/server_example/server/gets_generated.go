package state

import (
	"errors"
	"log"
	"net/http"
)

const (
	messageKindAction_MovePlayer messageKind = iota + 1
	messageKindAction_addItemToPlayer
	messageKindAction_spawnZoneItems
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
	MovePlayer      func(MovePlayerParams, *Engine)
	addItemToPlayer func(AddItemToPlayerParams, *Engine)
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
	case messageKindAction_MovePlayer:
		var params MovePlayerParams
		err := params.UnmarshalJSON(msg.Content)
		if err != nil {
			return err
		}
		r.actions.MovePlayer(params, r.state)
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
	movePlayer func(MovePlayerParams, *Engine),
	addItemToPlayer func(AddItemToPlayerParams, *Engine),
	spawnZoneItemsParams func(SpawnZoneItemsParams, *Engine),
	onDeploy func(*Engine),
	onFrameTick func(*Engine),
) {
	log.Println("Hello World")
	a := actions{movePlayer, addItemToPlayer, spawnZoneItemsParams}
	setupRoutes(a, onDeploy, onFrameTick)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
