package integrationtest

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/Java-Jonas/bar-cli/integrationtest/state"
	"github.com/Java-Jonas/bar-cli/testutils"
	"github.com/stretchr/testify/assert"
)

func TestIntegration(t *testing.T) {
	go startServer()

	serverResponseChannel := make(chan state.Message)
	c, ctx, _ := dialServer(serverResponseChannel)

	serverResponse := <-serverResponseChannel
	assert.Equal(t, state.MessageKindCurrentState, serverResponse.Kind)
	expected := `{"player":{"1":{"id":1,"gearScore":{"id":2,"level":1,"operationKind":"UNCHANGED"},"position":{"id":3,"operationKind":"UNCHANGED"},"operationKind":"UNCHANGED"}}}`
	actual := string(serverResponse.Content)
	if expected != actual {
		t.Error(testutils.Diff(actual, expected))

	}

	sendActionAddItemToPlayer(ctx, c)
	serverResponse = <-serverResponseChannel
	assert.Equal(t, state.MessageKindAction_addItemToPlayer, serverResponse.Kind)
	expected = `{"playerPath":"$.player.1"}`
	actual = string(serverResponse.Content)
	if expected != actual {
		t.Error(testutils.Diff(actual, expected))
	}
	serverResponse = <-serverResponseChannel
	assert.Equal(t, state.MessageKindUpdate, serverResponse.Kind)
	expected = `{"player":{"1":{"id":1,"gearScore":{"id":2,"level":2,"operationKind":"UPDATE"},"items":{"4":{"id":4,"gearScore":{"id":5,"operationKind":"UPDATE"},"name":"myItem","origin":{"id":7,"gearScore":{"id":8,"operationKind":"UPDATE"},"position":{"id":9,"operationKind":"UPDATE"},"operationKind":"UPDATE"},"operationKind":"UPDATE"}},"operationKind":"UPDATE"}}}`
	actual = string(serverResponse.Content)
	if expected != actual {
		t.Error(testutils.Diff(actual, expected))
	}

	sendActionMovePlayer(ctx, c)
	serverResponse = <-serverResponseChannel
	assert.Equal(t, state.MessageKindUpdate, serverResponse.Kind)
	expected = `{"player":{"1":{"id":1,"gearScore":{"id":2,"level":3,"operationKind":"UPDATE"},"position":{"id":3,"x":1,"operationKind":"UPDATE"},"operationKind":"UNCHANGED"}}}`
	actual = string(serverResponse.Content)
	if expected != actual {
		t.Error(testutils.Diff(actual, expected))
	}

	sendActionUnknownKind(ctx, c)
	serverResponse = <-serverResponseChannel
	assert.Equal(t, state.MessageKindError, serverResponse.Kind)
	expected = `unknown message kind whoami`
	actual = string(serverResponse.Content)
	if expected != actual {
		t.Error(testutils.Diff(actual, expected))
	}
	serverResponse = <-serverResponseChannel
	assert.Equal(t, state.MessageKindUpdate, serverResponse.Kind)
	expected = `{"player":{"1":{"id":1,"gearScore":{"id":2,"level":4,"operationKind":"UPDATE"},"operationKind":"UNCHANGED"}}}`
	actual = string(serverResponse.Content)
	if expected != actual {
		t.Error(testutils.Diff(actual, expected))
	}

	sendActionBadContent(ctx, c)
	serverResponse = <-serverResponseChannel
	assert.Equal(t, state.MessageKindError, serverResponse.Kind)
	expected = "error when unmarshalling received message content `{ badcontent123# \"playerID\": 0, \"changeX\": 1, \"changeY\": 0}`: parse error: syntax error near offset 2 of 'badcontent...'"
	actual = string(serverResponse.Content)
	if expected != actual {
		t.Error(testutils.Diff(actual, expected))
	}
	serverResponse = <-serverResponseChannel
	assert.Equal(t, state.MessageKindUpdate, serverResponse.Kind)
	expected = `{"player":{"1":{"id":1,"gearScore":{"id":2,"level":5,"operationKind":"UPDATE"},"operationKind":"UNCHANGED"}}}`
	actual = string(serverResponse.Content)
	if expected != actual {
		t.Error(testutils.Diff(actual, expected))
	}

	sendBadAction(ctx, c)
	serverResponse = <-serverResponseChannel
	assert.Equal(t, state.MessageKindError, serverResponse.Kind)
	expected = "error when unmarshalling received message content `\"foo bar\"\n`: parse error: expected { near offset 9 of 'foo bar'"
	actual = string(serverResponse.Content)
	if expected != actual {
		t.Error(testutils.Diff(actual, expected))
	}
	serverResponse = <-serverResponseChannel
	assert.Equal(t, state.MessageKindUpdate, serverResponse.Kind)
	expected = `{"player":{"1":{"id":1,"gearScore":{"id":2,"level":6,"operationKind":"UPDATE"},"operationKind":"UNCHANGED"}}}`
	actual = string(serverResponse.Content)
	if expected != actual {
		t.Error(testutils.Diff(actual, expected))
	}

	resp, err := http.Get("http://localhost:3496/inspect")
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	expected = expectedConfig
	actual = string(body)
	if expected != actual {
		t.Error(testutils.Diff(actual, expected))
	}
	m := make(map[string]interface{})
	err = json.Unmarshal(body, &m)
	if err != nil {
		t.Error("expected inspection result to be unmarshallable, but it was not")
	}
	assert.Equal(t, 3, len(m))
}

const expectedConfig = `{
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
}
`
