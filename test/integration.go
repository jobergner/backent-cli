package main

import (
	"context"
	"fmt"

	"github.com/jobergner/backent-cli/examples/client"
	"github.com/jobergner/backent-cli/examples/message"
	"github.com/jobergner/backent-cli/examples/server"
	"github.com/jobergner/backent-cli/examples/state"

	"github.com/rs/zerolog"
)

type Controller struct{}

func (c *Controller) OnSuperMessage(msg server.Message, room *server.Room, client *server.Client, lobby *server.Lobby) {
}
func (c *Controller) OnClientConnect(client *server.Client, lobby *server.Lobby) {
}
func (c *Controller) OnClientDisconnect(room *server.Room, clientID string, lobby *server.Lobby) {
}
func (c *Controller) OnCreation(lobby *server.Lobby) {
}
func (c *Controller) OnFrameTick(engine *state.Engine) {
}
func (c *Controller) AddItemToPlayerBroadcast(params message.AddItemToPlayerParams, engine *state.Engine, roomName, clientID string) {
}
func (c *Controller) AddItemToPlayerEmit(params message.AddItemToPlayerParams, engine *state.Engine, roomName, clientID string) message.AddItemToPlayerResponse {
	return message.AddItemToPlayerResponse{}
}
func (c *Controller) MovePlayerBroadcast(params message.MovePlayerParams, engine *state.Engine, roomName, clientID string) {
}
func (c *Controller) MovePlayerEmit(params message.MovePlayerParams, engine *state.Engine, roomName, clientID string) {
}
func (c *Controller) SpawnZoneItemsBroadcast(params message.SpawnZoneItemsParams, engine *state.Engine, roomName, clientID string) {
}
func (c *Controller) SpawnZoneItemsEmit(params message.SpawnZoneItemsParams, engine *state.Engine, roomName, clientID string) message.SpawnZoneItemsResponse {
	return message.SpawnZoneItemsResponse{}
}

func main() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	ctrl := &Controller{}

	server.Start(ctrl, 1, 8080)

	c, err := client.NewClient(context.TODO(), ctrl, 1)
	if err != nil {
		panic(err)
	}

	patchBytes := c.ReadUpdate()
	fmt.Println("receiving patch", string(patchBytes))
}
