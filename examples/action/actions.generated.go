package action

import (
	"github.com/jobergner/backent-cli/examples/message"
	"github.com/jobergner/backent-cli/examples/state"
)

const (
	MessageKindAction_addItemToPlayer message.Kind = "addItemToPlayer"
	MessageKindAction_movePlayer      message.Kind = "movePlayer"
	MessageKindAction_spawnZoneItems  message.Kind = "spawnZoneItems"
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

type AddItemToPlayerResponse struct {
	PlayerPath string `json:"playerPath"`
}

type SpawnZoneItemsResponse struct {
	NewZoneItemPaths []string `json:"newZoneItemPaths"`
}

//easyjson:skip
type AddItemToPlayerAction struct {
	Broadcast func(params AddItemToPlayerParams, engine *state.Engine, roomName, clientID string)
	Emit      func(params AddItemToPlayerParams, engine *state.Engine, roomName, clientID string) AddItemToPlayerResponse
}

//easyjson:skip
type MovePlayerAction struct {
	Broadcast func(params MovePlayerParams, engine *state.Engine, roomName, clientID string)
	Emit      func(params MovePlayerParams, engine *state.Engine, roomName, clientID string)
}

//easyjson:skip
type SpawnZoneItemsAction struct {
	Broadcast func(params SpawnZoneItemsParams, engine *state.Engine, roomName, clientID string)
	Emit      func(params SpawnZoneItemsParams, engine *state.Engine, roomName, clientID string) SpawnZoneItemsResponse
}

//easyjson:skip
type Actions struct {
	AddItemToPlayer AddItemToPlayerAction
	MovePlayer      MovePlayerAction
	SpawnZoneItems  SpawnZoneItemsAction
}
