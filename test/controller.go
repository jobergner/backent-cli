package endtoend

import (
	"github.com/jobergner/backent-cli/examples/message"
	"github.com/jobergner/backent-cli/examples/server"
	"github.com/jobergner/backent-cli/examples/state"
)

//go:generate mockgen -destination=mock_controller.go -package=endtoend . Controller
type Controller interface {
	OnSuperMessage(msg server.Message, room *server.Room, client *server.Client, lobby *server.Lobby)
	OnClientConnect(client *server.Client, lobby *server.Lobby)
	OnClientDisconnect(room *server.Room, clientID string, lobby *server.Lobby)
	OnCreation(lobby *server.Lobby)
	OnFrameTick(engine *state.Engine)
	AddItemToPlayerBroadcast(params message.AddItemToPlayerParams, engine *state.Engine, roomName, clientID string)
	AddItemToPlayerEmit(params message.AddItemToPlayerParams, engine *state.Engine, roomName, clientID string) message.AddItemToPlayerResponse
	MovePlayerBroadcast(params message.MovePlayerParams, engine *state.Engine, roomName, clientID string)
	MovePlayerEmit(params message.MovePlayerParams, engine *state.Engine, roomName, clientID string)
	SpawnZoneItemsBroadcast(params message.SpawnZoneItemsParams, engine *state.Engine, roomName, clientID string)
	SpawnZoneItemsEmit(params message.SpawnZoneItemsParams, engine *state.Engine, roomName, clientID string) message.SpawnZoneItemsResponse
}
