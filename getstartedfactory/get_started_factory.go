package getstartedfactory

import (
	"bytes"

	. "github.com/jobergner/backent-cli/factoryutils"

	. "github.com/dave/jennifer/jen"
	"github.com/jobergner/backent-cli/ast"
)

type GetStartedFactory struct {
	config *ast.AST
	buf    *bytes.Buffer
}

func newGetStartedFactory(config *ast.AST) *GetStartedFactory {
	return &GetStartedFactory{
		config: config,
		buf:    &bytes.Buffer{},
	}
}

func WriteGetStarted(moduleName string, useExample bool, stateConfigData, actionsConfigData, responsesConfigData map[interface{}]interface{}) string {
	config := ast.Parse(stateConfigData, actionsConfigData, responsesConfigData)

	if useExample {
		g := newGetStartedFactory(config).
			writePackageName().
			writeImport(moduleName)

		g.buf.WriteString(exampleMainFile)

		return g.buf.String()
	}

	g := newGetStartedFactory(config).
		writePackageName().
		writeImport(moduleName).
		writeMainFunc()

	return g.buf.String()
}

func (g *GetStartedFactory) writePackageName() *GetStartedFactory {
	g.buf.WriteString("package main\n")
	return g
}

// TODO fix
func (g *GetStartedFactory) writeImport(moduleName string) *GetStartedFactory {
	g.buf.WriteString(`
import (
	state "` + moduleName + `"
)`)
	return g
}

func (g *GetStartedFactory) writeMainFunc() *GetStartedFactory {
	decls := NewDeclSet()

	decls.File.Const().Id("fps").Op("=").Lit(30)

	decls.File.Var().Id("sideEffects").Op("=").Id("state").Dot("SideEffects").Values(Dict{
		Id("OnDeploy"):    Func().Params(Id("engine").Id("*state.Engine")).Block(),
		Id("OnFrameTick"): Func().Params(Id("engine").Id("*state.Engine")).Block(),
	}).Line()

	decls.File.Var().Id("actions").Op("=").Id("state").Dot("Actions").Values(
		Line().Add(
			ForEachActionInAST(g.config, func(action ast.Action) *Statement {
				if action.Response == nil {
					return Id(Title(action.Name)).Op(":").Func().Params(Id("params").Id("state").Dot(Title(action.Name)+"Params"), Id("engine").Id("*state.Engine")).Block().Id(",")
				}
				responseName := Id("state").Dot(Title(action.Name) + "Response")
				return Id(Title(action.Name)).Op(":").Func().Params(Id("params").Id("state").Dot(Title(action.Name)+"Params"), Id("engine").Id("*state.Engine")).Add(responseName).Block(
					Return(responseName).Values(),
				).Id(",")
			}),
		),
	)

	decls.File.Func().Id("main").Params().Block(
		Id("err").Op(":=").Id("state").Dot("Start").Call(Id("actions"), Id("sideEffects"), Id("fps"), Lit(3496)),
		If(Id("err").Op("!=").Nil()).Block(
			Panic(Id("err")),
		),
	)

	decls.Render(g.buf)
	return g
}

const exampleMainFile = `
const fps = 30

var sideEffects = state.SideEffects{
	OnDeploy: func(engine *state.Engine) {
		engine.CreateNpc().SetName("Mottled Boar")
		engine.CreateNpc().SetName("Scorpid Worker")
		engine.CreatePlayer().SetName("Thralltheorc")
	},
	OnFrameTick: func(engine *state.Engine) {},
}

var actions = state.Actions{
	AddFriend: func(params state.AddFriendParams, engine *state.Engine) state.AddFriendResponse {
		player := engine.Player(params.Player)
		player.AddFriendsList(params.NewFriend)
		return state.AddFriendResponse{
			NewNumberOfFriends: len(player.FriendsList()),
		}
	},
	AddItemToPlayer: func(params state.AddItemToPlayerParams, engine *state.Engine) state.AddItemToPlayerResponse {
		player := engine.Player(params.Player)
		item := player.AddItem().SetName(params.ItemName)
		item.SetFirstLootedBy(player.ID())
		return state.AddItemToPlayerResponse{
			ItemPath: item.Path(),
		}
	},
	CreatePlayer: func(params state.CreatePlayerParams, engine *state.Engine) state.CreatePlayerResponse {
		player := engine.CreatePlayer().SetName(params.Name)
		return state.CreatePlayerResponse{
			PlayerPath: player.Path(),
		}
	},
	DeletePlayer: func(params state.DeletePlayerParams, engine *state.Engine) {
		engine.DeletePlayer(params.Player)
	},
	MoveNpc: func(params state.MoveNpcParams, engine *state.Engine) {
		npc := engine.Npc(params.Npc)
		npc.Location().SetX(params.NewX).SetY(params.NewY)
	},
	MovePlayer: func(params state.MovePlayerParams, engine *state.Engine) {
		player := engine.Player(params.Player)
		player.Location().SetX(params.NewX).SetY(params.NewY)
	},
	PlayerLeaveCombat: func(params state.PlayerLeaveCombatParams, engine *state.Engine) state.PlayerLeaveCombatResponse {
		player := engine.Player(params.Player)
		inCombatRef, isSet := player.InCombatWith()
		if isSet {
			inCombatRef.Unset()
		}
		return state.PlayerLeaveCombatResponse{
			CombatWon: true,
		}
	},
	RemoveFriend: func(params state.RemoveFriendParams, engine *state.Engine) {
		player := engine.Player(params.Player)
		player.RemoveFriendsList(params.FriendToRemove)
	},
	RemoveItemFromPlayer: func(params state.RemoveItemFromPlayerParams, engine *state.Engine) {
		player := engine.Player(params.Player)
		player.RemoveItems(params.Item)
	},
	SetPlayerCombat: func(params state.SetPlayerCombatParams, engine *state.Engine) state.SetPlayerCombatResponse {
		player := engine.Player(params.Player)
		if state.ElementKind(params.EnemyKind) == state.ElementKindNpc {
			enemyNpc := engine.Npc(state.NpcID(params.EnemyID))
			player.SetInCombatWithNpc(enemyNpc.ID())
			return state.SetPlayerCombatResponse{
				EnemyEntityKind: string(state.ElementKindNpc),
				EnemyEntityPath: enemyNpc.Path(),
			}
		}
		enemyPlayer := engine.Player(state.PlayerID(params.EnemyID))
		player.SetInCombatWithPlayer(enemyPlayer.ID())
		return state.SetPlayerCombatResponse{
			EnemyEntityKind: string(state.ElementKindNpc),
			EnemyEntityPath: enemyPlayer.Path(),
		}
	},
}

func main() {
	err := state.Start(actions, sideEffects, fps, 3496)
	if err != nil {
		panic(err)
	}
}
`
