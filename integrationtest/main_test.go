package integrationtest

import (
	"github.com/Java-Jonas/bar-cli/integrationtest/state"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestIntegration(t *testing.T) {
	go startServer()
	serverResponseChannel := make(chan state.Message)
	c, ctx, _ := dialServer(serverResponseChannel)

	log.Println("init")
	serverResponse := <-serverResponseChannel
	assert.Equal(t, state.MessageKindCurrentState, serverResponse.Kind)
	assert.Equal(t, `{"player":{"1":{"id":1,"gearScore":{"id":2,"level":1,"operationKind":"UNCHANGED"},"position":{"id":3,"operationKind":"UNCHANGED"},"operationKind":"UNCHANGED"}}}`, string(serverResponse.Content))

	sendActionAddItemToPlayer(ctx, c)
	serverResponse = <-serverResponseChannel
	assert.Equal(t, state.MessageKindUpdate, serverResponse.Kind)
	assert.Equal(t, `{"player":{"1":{"id":1,"gearScore":{"id":2,"level":2,"operationKind":"UPDATE"},"items":{"4":{"id":4,"gearScore":{"id":5,"operationKind":"UPDATE"},"name":"myItem","origin":{"id":7,"gearScore":{"id":8,"operationKind":"UPDATE"},"position":{"id":9,"operationKind":"UPDATE"},"operationKind":"UPDATE"},"operationKind":"UPDATE"}},"operationKind":"UPDATE"}}}`, string(serverResponse.Content))
	serverResponse = <-serverResponseChannel
	assert.Equal(t, state.MessageKindAction_addItemToPlayer, serverResponse.Kind)
	assert.Equal(t, `{"playerPath":"$.player.1"}`, string(serverResponse.Content))

	sendActionMovePlayer(ctx, c)
	serverResponse = <-serverResponseChannel
	assert.Equal(t, state.MessageKindUpdate, serverResponse.Kind)
	assert.Equal(t, `{"player":{"1":{"id":1,"gearScore":{"id":2,"level":3,"operationKind":"UPDATE"},"position":{"id":3,"x":1,"operationKind":"UPDATE"},"operationKind":"UNCHANGED"}}}`, string(serverResponse.Content))

	sendActionUnknownKind(ctx, c)
	serverResponse = <-serverResponseChannel
	assert.Equal(t, state.MessageKindUpdate, serverResponse.Kind)
	assert.Equal(t, `{"player":{"1":{"id":1,"gearScore":{"id":2,"level":4,"operationKind":"UPDATE"},"operationKind":"UNCHANGED"}}}`, string(serverResponse.Content))
	serverResponse = <-serverResponseChannel
	assert.Equal(t, state.MessageKindError, serverResponse.Kind)
	assert.Equal(t, `unknown message kind whoami`, string(serverResponse.Content))

	sendActionBadContent(ctx, c)
	serverResponse = <-serverResponseChannel
	assert.Equal(t, state.MessageKindUpdate, serverResponse.Kind)
	assert.Equal(t, `{"player":{"1":{"id":1,"gearScore":{"id":2,"level":5,"operationKind":"UPDATE"},"operationKind":"UNCHANGED"}}}`, string(serverResponse.Content))
	serverResponse = <-serverResponseChannel
	assert.Equal(t, state.MessageKindError, serverResponse.Kind)
	assert.Equal(t, "error when unmarshalling received message content `{ badcontent123# \"playerID\": 0, \"changeX\": 1, \"changeY\": 0}`: parse error: syntax error near offset 2 of 'badcontent...'", string(serverResponse.Content))
}
