package server

import (
	"github.com/jobergner/backent-cli/examples/message"
	"github.com/jobergner/backent-cli/examples/state"
)

type Controller interface {
	OnSuperMessage(msg Message, room *Room, client *Client, lobby *Lobby)
	OnClientConnect(client *Client, lobby *Lobby)
	OnClientDisconnect(room *Room, clientID string, lobby *Lobby)
	OnCreation(lobby *Lobby)
	OnFrameTick(engine *state.Engine)
	AddItemToPlayerBroadcast(params message.AddItemToPlayerParams, engine *state.Engine, roomName, clientID string)
	AddItemToPlayerEmit(params message.AddItemToPlayerParams, engine *state.Engine, roomName, clientID string) message.AddItemToPlayerResponse
	MovePlayerBroadcast(params message.MovePlayerParams, engine *state.Engine, roomName, clientID string)
	MovePlayerEmit(params message.MovePlayerParams, engine *state.Engine, roomName, clientID string)
	SpawnZoneItemsBroadcast(params message.SpawnZoneItemsParams, engine *state.Engine, roomName, clientID string)
	SpawnZoneItemsEmit(params message.SpawnZoneItemsParams, engine *state.Engine, roomName, clientID string) message.SpawnZoneItemsResponse
}
