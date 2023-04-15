package client

import (
	"time"

	"github.com/jobergner/backent-cli/examples/logging"
	"github.com/jobergner/backent-cli/examples/message"
	"github.com/rs/zerolog/log"
)

func (c *Client) AddItemToPlayer(params message.AddItemToPlayerParams) (message.AddItemToPlayerResponse, error) {

	msgContent, err := params.MarshalJSON()
	if err != nil {
		log.Err(err).Str(logging.MessageKind, string(message.MessageKindAction_addItemToPlayer)).Msg("failed marshalling parameters")
		return message.AddItemToPlayerResponse{}, err
	}

	id, err := newMessageID()
	if err != nil {
		return message.AddItemToPlayerResponse{}, err
	}

	msg := Message{id, message.MessageKindAction_addItemToPlayer, msgContent}

	msgBytes, err := msg.MarshalJSON()
	if err != nil {
		log.Err(err).Str(logging.MessageKind, string(message.MessageKindAction_addItemToPlayer)).Int(logging.MessageID, msg.ID).Str(logging.Message, string(msgBytes)).Msg("failed marshalling message")
		return message.AddItemToPlayerResponse{}, err
	}

	responseChan := make(chan []byte)

	c.router.add(id, responseChan)
	defer c.router.remove(id)

	c.messageChannel <- msgBytes

	select {
	case <-time.After(2 * time.Second):
		log.Err(ErrResponseTimeout).Int(logging.MessageID, msg.ID).Msg("timed out waiting for response")
		return message.AddItemToPlayerResponse{}, ErrResponseTimeout

	case responseBytes := <-responseChan:
		var res message.AddItemToPlayerResponse
		err := res.UnmarshalJSON(responseBytes)
		if err != nil {
			log.Err(err).Str(logging.MessageKind, string(message.MessageKindAction_addItemToPlayer)).Int(logging.MessageID, msg.ID).Msg("failed unmarshalling response")
			return message.AddItemToPlayerResponse{}, err
		}

		return res, nil
	}
}

func (c *Client) MovePlayer(params message.MovePlayerParams) error {

	msgContent, err := params.MarshalJSON()
	if err != nil {
		log.Err(err).Str(logging.MessageKind, string(message.MessageKindAction_movePlayer)).Msg("failed marshalling parameters")
		return err
	}

	id, err := newMessageID()
	if err != nil {
		return err
	}

	msg := Message{id, message.MessageKindAction_movePlayer, msgContent}

	msgBytes, err := msg.MarshalJSON()
	if err != nil {
		log.Err(err).Str(logging.MessageKind, string(message.MessageKindAction_movePlayer)).Int(logging.MessageID, msg.ID).Str(logging.Message, string(msgBytes)).Msg("failed marshalling message")
		return err
	}

	c.messageChannel <- msgBytes

	return nil
}

func (c *Client) SpawnZoneItems(params message.SpawnZoneItemsParams) (message.SpawnZoneItemsResponse, error) {

	msgContent, err := params.MarshalJSON()
	if err != nil {
		log.Err(err).Str(logging.MessageKind, string(message.MessageKindAction_spawnZoneItems)).Msg("failed marshalling parameters")
		return message.SpawnZoneItemsResponse{}, err
	}

	id, err := newMessageID()
	if err != nil {
		return message.SpawnZoneItemsResponse{}, err
	}

	msg := Message{id, message.MessageKindAction_spawnZoneItems, msgContent}

	msgBytes, err := msg.MarshalJSON()
	if err != nil {
		log.Err(err).Str(logging.MessageKind, string(message.MessageKindAction_spawnZoneItems)).Int(logging.MessageID, msg.ID).Str(logging.Message, string(msgBytes)).Msg("failed marshalling message")
		return message.SpawnZoneItemsResponse{}, err
	}

	responseChan := make(chan []byte)

	c.router.add(id, responseChan)
	defer c.router.remove(id)

	c.messageChannel <- msgBytes

	select {
	case <-time.After(2 * time.Second):
		log.Err(ErrResponseTimeout).Int(logging.MessageID, msg.ID).Msg("timed out waiting for response")
		return message.SpawnZoneItemsResponse{}, ErrResponseTimeout

	case responseBytes := <-responseChan:
		var res message.SpawnZoneItemsResponse
		err := res.UnmarshalJSON(responseBytes)
		if err != nil {
			log.Err(err).Str(logging.MessageKind, string(message.MessageKindAction_spawnZoneItems)).Int(logging.MessageID, msg.ID).Msg("failed unmarshalling response")
			return message.SpawnZoneItemsResponse{}, err
		}

		return res, nil
	}
}
