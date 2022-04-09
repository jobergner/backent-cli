package endtoend

import (
	"path/filepath"
	"runtime"
	"strings"

	"github.com/jobergner/backent-cli/examples/message"
	server "github.com/jobergner/backent-cli/examples/server"
	"github.com/jobergner/backent-cli/examples/state"
	"github.com/rs/zerolog"
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

type SeverityHook struct{}

var _, callerFile, _, _ = runtime.Caller(0)
var ProjectDir = filepath.Clean(filepath.Join(callerFile, "..", ".."))

func (h SeverityHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	stack := make([]string, 0)

	pc := make([]uintptr, 32)
	n := runtime.Callers(2, pc)
	if n == 0 {
		return
	}
	pc = pc[:n]
	frames := runtime.CallersFrames(pc)

	for {
		frame, more := frames.Next()
		file := frame.File
		if strings.HasPrefix(file, ProjectDir) {
			// strip project dir
			file = file[len(ProjectDir)+1:]
			// add line
			stack = append(stack, zerolog.CallerMarshalFunc(file, frame.Line))
		}
		if !more {
			break
		}
	}

	for i, j := 0, len(stack)-1; i < j; i, j = i+1, j-1 {
		stack[i], stack[j] = stack[j], stack[i]
	}

	var callerOrder string
	for i, s := range stack {
		if i == 0 {
			callerOrder += s
			continue
		}
		callerOrder += " -> " + s
	}

	e.Str("stack", callerOrder)
}
