package state

import (
	"fmt"
	"net/http"
)

const (
	MessageKindAction_addItemToPlayer MessageKind = "addItemToPlayer"
	MessageKindAction_movePlayer      MessageKind = "movePlayer"
	MessageKindAction_spawnZoneItems  MessageKind = "spawnZoneItems"
)

type MovePlayerParams struct {
	ChangeX float64  `json:"changeX"`
	ChangeY float64  `json:"changeY"`
	Player  PlayerID `json:"player"`
}

type AddItemToPlayerParams struct {
	Item    ItemID `json:"item"`
	NewName string `json:"newName"`
}

type SpawnZoneItemsParams struct {
	Items []ItemID `json:"items"`
}

type AddItemToPlayerResponse struct {
	PlayerPath string `json:"playerPath"`
}

type SpawnZoneItemsResponse struct {
	NewZoneItemPaths []string `json:"newZoneItemPaths"`
}

type Actions struct {
	AddItemToPlayer func(params AddItemToPlayerParams, engine *Engine, roomName, clientID string) AddItemToPlayerResponse
	MovePlayer      func(params MovePlayerParams, engien *Engine, roomName, clientID string)
	SpawnZoneItems  func(params SpawnZoneItemsParams, engine *Engine, roomName, clientID string) SpawnZoneItemsResponse
}

type SideEffects struct {
	OnDeploy    func(*Engine)
	OnFrameTick func(*Engine)
}

type LoginSignals struct {
	OnGlobalMessage    func(Message, *Engine, *Client, *LoginHandler)
	OnClientConnect    func(*Client, *LoginHandler)
	OnClientDisconnect func(*Engine, string, *LoginHandler)
}

// TODO logging should happen here
func (r *Room) processClientMessage(msg Message) (Message, error) {
	switch MessageKind(msg.Kind) {
	case MessageKindAction_addItemToPlayer:
		if r.actions.AddItemToPlayer == nil {
			break
		}
		var params AddItemToPlayerParams
		err := params.UnmarshalJSON(msg.Content)
		if err != nil {
			return Message{MessageKindError, messageUnmarshallingError(msg.Content, err), msg.client}, err
		}
		res := r.actions.AddItemToPlayer(params, r.state, r.name, msg.client.id)
		resContent, err := res.MarshalJSON()
		if err != nil {
			return Message{MessageKindError, responseMarshallingError(msg.Content, err), msg.client}, err
		}
		return Message{msg.Kind, resContent, msg.client}, nil
	case MessageKindAction_movePlayer:
		if r.actions.MovePlayer == nil {
			break
		}
		var params MovePlayerParams
		err := params.UnmarshalJSON(msg.Content)
		if err != nil {
			return Message{MessageKindError, messageUnmarshallingError(msg.Content, err), msg.client}, err
		}
		r.actions.MovePlayer(params, r.state, r.name, msg.client.id)
		return Message{}, nil
	case MessageKindAction_spawnZoneItems:
		if r.actions.SpawnZoneItems == nil {
			break
		}
		var params SpawnZoneItemsParams
		err := params.UnmarshalJSON(msg.Content)
		if err != nil {
			return Message{MessageKindError, messageUnmarshallingError(msg.Content, err), msg.client}, err
		}
		res := r.actions.SpawnZoneItems(params, r.state, r.name, msg.client.id)
		resContent, err := res.MarshalJSON()
		if err != nil {
			return Message{MessageKindError, responseMarshallingError(msg.Content, err), msg.client}, err
		}
		return Message{msg.Kind, resContent, msg.client}, nil
	default:
		return Message{MessageKindError, []byte("unknown message kind " + msg.Kind), msg.client}, fmt.Errorf("unknown message kind in: %s", printMessage(msg))
	}

	return Message{}, nil
}

func inspectHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	fmt.Fprintf(w, `{
  "state": {
    "player": {
      "items": "[]item",
      "equipmentSets": "[]*equipmentSet",
      "gearScore": "gearScore",
      "position": "position",
      "guildMembers": "[]*player",
      "target": "*anyOf<player,zoneItem>",
      "targetedBy": "[]*anyOf<player,zoneItem>"
    },
    "zone": {
      "items": "[]zoneItem",
      "players": "[]player",
      "tags": "[]string",
      "interactables": "[]anyOf<item,player,zoneItem>"
    },
    "zoneItem": {
      "position": "position",
      "item": "item"
    },
    "position": {
      "x": "float64",
      "y": "float64"
    },
    "item": {
      "name": "string",
      "gearScore": "gearScore",
      "boundTo": "*player",
      "origin": "anyOf<player,position>"
    },
    "gearScore": {
      "level": "int",
      "score": "int"
    },
    "equipmentSet": {
      "name": "string",
      "equipment": "[]*item"
    }
  },
  "actions": {
    "movePlayer": {
      "player": "playerID",
      "changeX": "float64",
      "changeY": "float64"
    },
    "addItemToPlayer": {
      "item": "itemID",
      "newName": "string"
    },
    "spawnZoneItems": {
      "items": "[]itemID"
    }
  },
  "responses" : {
    "addItemToPlayer": {
      "playerPath": "string"
    },
    "spawnZoneItems": {
      "newZoneItemPaths": "string"
    }
  }
}`)
}
