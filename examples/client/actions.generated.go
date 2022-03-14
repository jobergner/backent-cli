package client

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jobergner/backent-cli/examples/action"
)

func (c *Client) AddItemToPlayer(params action.AddItemToPlayerParams) (action.AddItemToPlayerResponse, error) {
	c.mu.Lock()
	c.actions.AddItemToPlayer.Broadcast(params, c.engine, "", c.id)
	c.mu.Unlock()

	msgContent, err := params.MarshalJSON()
	if err != nil {
		return action.AddItemToPlayerResponse{}, err
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return action.AddItemToPlayerResponse{}, err
	}

	idString := id.String()

	msg := Message{idString, action.MessageKindAction_addItemToPlayer, msgContent}

	msgBytes, err := msg.MarshalJSON()
	if err != nil {
		return action.AddItemToPlayerResponse{}, err
	}

	responseChan := make(chan []byte)
	c.router.add(idString, responseChan)
	defer c.router.remove(idString)

	c.messageChannel <- msgBytes

	select {
	case <-time.After(2 * time.Second):
		return action.AddItemToPlayerResponse{}, errors.New("timeout")
	case responseBytes := <-responseChan:
		var res action.AddItemToPlayerResponse

		err := res.UnmarshalJSON(responseBytes)
		if err != nil {
			return action.AddItemToPlayerResponse{}, err
		}

		return res, nil
	}
}

func (c *Client) MovePlayer(params action.MovePlayerParams) error {
	c.mu.Lock()
	c.actions.MovePlayer.Broadcast(params, c.engine, "", c.id)
	c.mu.Unlock()

	msgContent, err := params.MarshalJSON()
	if err != nil {
		return err
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	idString := id.String()

	msg := Message{idString, action.MessageKindAction_addItemToPlayer, msgContent}

	msgBytes, err := msg.MarshalJSON()
	if err != nil {
		return err
	}

	c.messageChannel <- msgBytes

	return nil
}

func (c *Client) SpawnZoneItems(params action.SpawnZoneItemsParams) (action.SpawnZoneItemsResponse, error) {
	c.mu.Lock()
	c.actions.SpawnZoneItems.Broadcast(params, c.engine, "", c.id)
	c.mu.Unlock()

	msgContent, err := params.MarshalJSON()
	if err != nil {
		return action.SpawnZoneItemsResponse{}, err
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return action.SpawnZoneItemsResponse{}, err
	}

	idString := id.String()

	msg := Message{idString, action.MessageKindAction_addItemToPlayer, msgContent}

	msgBytes, err := msg.MarshalJSON()
	if err != nil {
		return action.SpawnZoneItemsResponse{}, err
	}

	responseChan := make(chan []byte)
	c.router.add(idString, responseChan)
	defer c.router.remove(idString)

	c.messageChannel <- msgBytes

	select {
	case <-time.After(2 * time.Second):
		return action.SpawnZoneItemsResponse{}, errors.New("timeout")
	case responseBytes := <-responseChan:
		var res action.SpawnZoneItemsResponse

		err := res.UnmarshalJSON(responseBytes)
		if err != nil {
			return action.SpawnZoneItemsResponse{}, err
		}

		return res, nil
	}
}
