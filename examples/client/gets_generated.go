package state

import (
	"errors"
	"time"

	"github.com/google/uuid"
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

type AddItemToPlayerAction struct {
	Broadcast func(params AddItemToPlayerParams, engine *Engine, clientID string)
	Emit      func(params AddItemToPlayerParams, engine *Engine, clientID string) AddItemToPlayerResponse
}
type MovePlayerAction struct {
	Broadcast func(params MovePlayerParams, engine *Engine, clientID string)
	Emit      func(params MovePlayerParams, engine *Engine, clientID string)
}
type SpawnZoneItemsAction struct {
	Broadcast func(params SpawnZoneItemsParams, engine *Engine, clientID string)
	Emit      func(params SpawnZoneItemsParams, engine *Engine, clientID string) SpawnZoneItemsResponse
}

type Actions struct {
	AddItemToPlayer AddItemToPlayerAction
	MovePlayer      MovePlayerAction
	SpawnZoneItems  SpawnZoneItemsAction
}

func (c *Client) AddItemToPlayer(params AddItemToPlayerParams) (AddItemToPlayerResponse, error) {
	c.actions.AddItemToPlayer.Broadcast(params, c.engine, c.id)

	msgContent, err := params.MarshalJSON()
	if err != nil {
		return AddItemToPlayerResponse{}, err
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return AddItemToPlayerResponse{}, err
	}

	idString := id.String()

	msg := Message{idString, MessageKindAction_addItemToPlayer, msgContent}

	msgBytes, err := msg.MarshalJSON()
	if err != nil {
		return AddItemToPlayerResponse{}, err
	}

	responseChan := make(chan []byte)
	c.router.add(idString, responseChan)
	defer c.router.remove(idString)

	c.messageChannel <- msgBytes

	select {
	case <-time.After(2 * time.Second):
		return AddItemToPlayerResponse{}, errors.New("timeout")
	case responseBytes := <-responseChan:
		var res AddItemToPlayerResponse

		err := res.UnmarshalJSON(responseBytes)
		if err != nil {
			return AddItemToPlayerResponse{}, err
		}

		return res, nil
	}
}

func (c *Client) MovePlayer(params MovePlayerParams) error {
	c.actions.MovePlayer.Broadcast(params, c.engine, c.id)

	msgContent, err := params.MarshalJSON()
	if err != nil {
		return err
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	idString := id.String()

	msg := Message{idString, MessageKindAction_addItemToPlayer, msgContent}

	msgBytes, err := msg.MarshalJSON()
	if err != nil {
		return err
	}

	c.messageChannel <- msgBytes

	return nil
}

func (c *Client) SpawnZoneItems(params SpawnZoneItemsParams) (SpawnZoneItemsResponse, error) {
	c.actions.SpawnZoneItems.Broadcast(params, c.engine, c.id)

	msgContent, err := params.MarshalJSON()
	if err != nil {
		return SpawnZoneItemsResponse{}, err
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return SpawnZoneItemsResponse{}, err
	}

	idString := id.String()

	msg := Message{idString, MessageKindAction_addItemToPlayer, msgContent}

	msgBytes, err := msg.MarshalJSON()
	if err != nil {
		return SpawnZoneItemsResponse{}, err
	}

	responseChan := make(chan []byte)
	c.router.add(idString, responseChan)
	defer c.router.remove(idString)

	c.messageChannel <- msgBytes

	select {
	case <-time.After(2 * time.Second):
		return SpawnZoneItemsResponse{}, errors.New("timeout")
	case responseBytes := <-responseChan:
		var res SpawnZoneItemsResponse

		err := res.UnmarshalJSON(responseBytes)
		if err != nil {
			return SpawnZoneItemsResponse{}, err
		}

		return res, nil
	}
}
