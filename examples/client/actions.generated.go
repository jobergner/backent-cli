package client

import (
	"time"

	"github.com/google/uuid"
	"github.com/jobergner/backent-cli/examples/logging"
	"github.com/jobergner/backent-cli/examples/message"
	"github.com/rs/zerolog/log"
)

func (c *Client) AddItemToPlayer(params message.AddItemToPlayerParams) (message.AddItemToPlayerResponse, error) {
	c.mu.Lock()
	c.controller.AddItemToPlayerBroadcast(params, c.engine, "", c.id)
	c.mu.Unlock()

	msgContent, err := params.MarshalJSON()
	if err != nil {
		log.Err(err).Str(logging.MessageKind, string(message.MessageKindAction_addItemToPlayer)).Msg("failed marshalling parameters")
		return message.AddItemToPlayerResponse{}, err
	}

	id, err := uuid.NewRandom()
	if err != nil {
		log.Err(err).Str(logging.MessageKind, string(message.MessageKindAction_addItemToPlayer)).Msg("failed generating message ID")
		return message.AddItemToPlayerResponse{}, err
	}

	idString := id.String()

	msg := Message{idString, message.MessageKindAction_addItemToPlayer, msgContent}

	msgBytes, err := msg.MarshalJSON()
	if err != nil {
		log.Err(err).Str(logging.MessageID, msg.ID).Str(logging.Message, string(msgBytes)).Str(logging.MessageKind, string(message.MessageKindAction_addItemToPlayer)).Msg("failed marshalling message")
		return message.AddItemToPlayerResponse{}, err
	}

	responseChan := make(chan []byte)

	c.router.add(idString, responseChan)
	defer c.router.remove(idString)

	c.messageChannel <- msgBytes

	select {
	case <-time.After(2 * time.Second):
		log.Err(ErrResponseTimeout).Str(logging.MessageID, msg.ID).Msg("timed out waiting for response")
		return message.AddItemToPlayerResponse{}, ErrResponseTimeout

	case responseBytes := <-responseChan:
		var res message.AddItemToPlayerResponse
		err := res.UnmarshalJSON(responseBytes)
		if err != nil {
			log.Err(err).Str(logging.MessageID, msg.ID).Str(logging.MessageKind, string(message.MessageKindAction_addItemToPlayer)).Msg("failed unmarshalling response")
			return message.AddItemToPlayerResponse{}, err
		}

		return res, nil
	}
}

func (c *Client) MovePlayer(params message.MovePlayerParams) error {
	c.mu.Lock()
	c.controller.MovePlayerBroadcast(params, c.engine, "", c.id)
	c.mu.Unlock()

	msgContent, err := params.MarshalJSON()
	if err != nil {
		log.Err(err).Str(logging.MessageKind, string(message.MessageKindAction_movePlayer)).Msg("failed marshalling parameters")
		return err
	}

	id, err := uuid.NewRandom()
	if err != nil {
		log.Err(err).Str(logging.MessageKind, string(message.MessageKindAction_movePlayer)).Msg("failed generating message ID")
		return err
	}

	idString := id.String()

	msg := Message{idString, message.MessageKindAction_movePlayer, msgContent}

	msgBytes, err := msg.MarshalJSON()
	if err != nil {
		log.Err(err).Str(logging.MessageID, msg.ID).Str(logging.Message, string(msgBytes)).Str(logging.MessageKind, string(message.MessageKindAction_movePlayer)).Msg("failed marshalling message")
		return err
	}

	c.messageChannel <- msgBytes

	return nil
}

func (c *Client) SpawnZoneItems(params message.SpawnZoneItemsParams) (message.SpawnZoneItemsResponse, error) {
	c.mu.Lock()
	c.controller.SpawnZoneItemsBroadcast(params, c.engine, "", c.id)
	c.mu.Unlock()

	msgContent, err := params.MarshalJSON()
	if err != nil {
		log.Err(err).Str(logging.MessageKind, string(message.MessageKindAction_spawnZoneItems)).Msg("failed marshalling parameters")
		return message.SpawnZoneItemsResponse{}, err
	}

	id, err := uuid.NewRandom()
	if err != nil {
		log.Err(err).Str(logging.MessageKind, string(message.MessageKindAction_spawnZoneItems)).Msg("failed generating message ID")
		return message.SpawnZoneItemsResponse{}, err
	}

	idString := id.String()

	msg := Message{idString, message.MessageKindAction_spawnZoneItems, msgContent}

	msgBytes, err := msg.MarshalJSON()
	if err != nil {
		log.Err(err).Str(logging.MessageID, msg.ID).Str(logging.Message, string(msgBytes)).Str(logging.MessageKind, string(message.MessageKindAction_spawnZoneItems)).Msg("failed marshalling message")
		return message.SpawnZoneItemsResponse{}, err
	}

	responseChan := make(chan []byte)

	c.router.add(idString, responseChan)
	defer c.router.remove(idString)

	c.messageChannel <- msgBytes

	select {
	case <-time.After(2 * time.Second):
		log.Err(ErrResponseTimeout).Str(logging.MessageID, msg.ID).Msg("timed out waiting for response")
		return message.SpawnZoneItemsResponse{}, ErrResponseTimeout

	case responseBytes := <-responseChan:
		var res message.SpawnZoneItemsResponse
		err := res.UnmarshalJSON(responseBytes)
		if err != nil {
			log.Err(err).Str(logging.MessageID, msg.ID).Str(logging.MessageKind, string(message.MessageKindAction_spawnZoneItems)).Msg("failed unmarshalling response")
			return message.SpawnZoneItemsResponse{}, err
		}

		return res, nil
	}
}
