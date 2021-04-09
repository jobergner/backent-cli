package state

import (
	"log"
)

const (
	messageKindAction_MovePlayer messageKind = iota + messageKindInit
	messageKindAction_addItemToPlayer
	messageKindAction_spawnZoneItems
)

type _MovePlayerParams struct {
	PlayerID PlayerID `json:"playerID"`
	ChangeX  float64  `json:"changeX"`
	ChangeY  float64  `json:"changeY"`
}

type _addItemToPlayer struct {
	Item     tItem    `json:"item"`
	PlayerID PlayerID `json:"playerID"`
}

type _spawnZoneItemsParams struct {
	Items []tItem `json:"items"`
}

type actions struct {
	movePlayerParams     func(PlayerID, float64, float64)
	addItemToPlayer      func(tItem, PlayerID)
	spawnZoneItemsParams func([]tItem)
}

func (r *Room) handleClientMessage(msg []byte) error {
	// r.state = r.state + 1

	log.Println(r.state)
	return nil
}
