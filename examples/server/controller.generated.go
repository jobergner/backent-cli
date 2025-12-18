package server

import (
	"github.com/jobergner/backent-cli/examples/message"
	"github.com/jobergner/backent-cli/examples/state"
)

type Controller interface {
	OnSuperMessage(msg Message, room *Room, client *Client, lobby *Lobby) Message
	OnClientConnect(client *Client, lobby *Lobby)
	OnClientDisconnect(room *Room, clientID string, lobby *Lobby)
	OnCreation(lobby *Lobby)
	OnFrameTick(engine *state.Engine)
	AddItemToPlayer(params message.AddItemToPlayerParams, engine *state.Engine, roomName, clientID string) message.AddItemToPlayerResponse
	MovePlayer(params message.MovePlayerParams, engine *state.Engine, roomName, clientID string)
	SpawnZoneItems(params message.SpawnZoneItemsParams, engine *state.Engine, roomName, clientID string) message.SpawnZoneItemsResponse
}
