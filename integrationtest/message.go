package main

import (
	"github.com/Java-Jonas/bar-cli/integrationtest/state"
)

// i just copy pasted this from the server example

type messageKind int

type message struct {
	Kind    messageKind `json:"kind"`
	Content []byte      `json:"content"`
}

const (
	messageKindAction_addItemToPlayer messageKind = 1
	messageKindAction_movePlayer      messageKind = 2
	messageKindAction_spawnZoneItems  messageKind = 3
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
