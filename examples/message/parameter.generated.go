package message

import (
	"github.com/jobergner/backent-cli/examples/state"
)

type MovePlayerParams struct {
	ChangeX float64        `json:"changeX"`
	ChangeY float64        `json:"changeY"`
	Player  state.PlayerID `json:"player"`
}

type AddItemToPlayerParams struct {
	Item    state.ItemID `json:"item"`
	NewName string       `json:"newName"`
}

type SpawnZoneItemsParams struct {
	Items []state.ItemID `json:"items"`
}
