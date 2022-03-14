package server

import (
	"github.com/jobergner/backent-cli/examples/action"
	"github.com/jobergner/backent-cli/examples/state"
)

//easyjson:skip
type applicationConfig struct {
	actions     action.Actions
	signals     LobbySignals
	sideEffects SideEffects
	fps         int
}

//easyjson:skip
type SideEffects struct {
	OnDeploy    func(engine *state.Engine)
	OnFrameTick func(engine *state.Engine)
}

//easyjson:skip
type LobbySignals struct {
	OnSuperMessage     func(msg Message, room *Room, client *Client, loginHandler *Lobby)
	OnClientConnect    func(client *Client, loginHandler *Lobby)
	OnClientDisconnect func(room *Room, clientID string, loginHandler *Lobby)
}
