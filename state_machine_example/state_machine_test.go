package statemachine

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStateMachine(t *testing.T) {
	t.Run("creates elements", func(t *testing.T) {
		sm := newStateMachine()
		gearScore := sm.CreateGearScore()
		_, ok := sm.Patch.GearScore[gearScore.gearScore.ID]
		assert.True(t, ok)
	})
	t.Run("gets elements", func(t *testing.T) {
		sm := newStateMachine()
		gearScore := sm.CreateGearScore()
		_gearScore := sm.GetGearScore(gearScore.gearScore.ID)
		assert.NotZero(t, _gearScore.gearScore.ID)
	})
	t.Run("sets elements", func(t *testing.T) {
		sm := newStateMachine()
		gearScore := sm.CreateGearScore()
		_gearScore := sm.Patch.GearScore[gearScore.gearScore.ID]
		assert.Zero(t, _gearScore.Level)
		gearScore.SetLevel(10, sm)
		_gearScore = sm.Patch.GearScore[gearScore.gearScore.ID]
		assert.NotZero(t, _gearScore.Level)
	})
	t.Run("deletes elements", func(t *testing.T) {
		sm := newStateMachine()
		gearScore := sm.CreateGearScore()
		sm.DeleteGearScore(gearScore.gearScore.ID)
		_gearScore := sm.Patch.GearScore[gearScore.gearScore.ID]
		assert.Equal(t, OperationKind(OperationKindDelete), _gearScore.OperationKind)
	})
	t.Run("adds elements", func(t *testing.T) {
		sm := newStateMachine()
		player := sm.CreatePlayer()
		item := player.AddItem(sm)
		playerItem := player.GetItems(sm)[0]
		assert.NotZero(t, playerItem.item.ID)
		_, ok := sm.Patch.Item[item.item.ID]
		assert.True(t, ok)
	})
	t.Run("removes elements", func(t *testing.T) {
		sm := newStateMachine()
		player := sm.CreatePlayer()
		item := player.AddItem(sm)
		player.RemoveItem(item.item.ID, sm)
		_item := sm.Patch.Item[item.item.ID]
		assert.Equal(t, OperationKind(OperationKindDelete), _item.OperationKind)
	})
}

func TestUpdateState(t *testing.T) {
	t.Run("creates elements", func(t *testing.T) {
		sm := newStateMachine()
		gearScore := sm.CreateGearScore()
		sm.UpdateState()
		_, ok := sm.Patch.GearScore[gearScore.gearScore.ID]
		assert.False(t, ok)
		_, ok = sm.State.GearScore[gearScore.gearScore.ID]
		assert.True(t, ok)
	})
	t.Run("gets elements", func(t *testing.T) {
		sm := newStateMachine()
		gearScore := sm.CreateGearScore()
		sm.UpdateState()
		assert.Zero(t, gearScore.GetLevel(sm))
		gearScore.SetLevel(1, sm)
		assert.Equal(t, gearScore.GetLevel(sm), 1)
	})
	t.Run("sets elements", func(t *testing.T) {
		sm := newStateMachine()
		gearScore := sm.CreateGearScore().SetLevel(1, sm)
		sm.UpdateState()
		_gearScore := sm.State.GearScore[gearScore.gearScore.ID]
		assert.Equal(t, _gearScore.Level, 1)
		_, ok := sm.Patch.GearScore[gearScore.gearScore.ID]
		assert.False(t, ok)
	})
	t.Run("deletes elements", func(t *testing.T) {
		sm := newStateMachine()
		gearScore := sm.CreateGearScore()
		sm.UpdateState()
		sm.DeleteGearScore(gearScore.gearScore.ID)
		sm.UpdateState()
		_, ok := sm.State.GearScore[gearScore.gearScore.ID]
		assert.False(t, ok)
	})
	t.Run("adds elements", func(t *testing.T) {
		sm := newStateMachine()
		player := sm.CreatePlayer()
		item := player.AddItem(sm)
		sm.UpdateState()
		_, ok := sm.State.Item[item.item.ID]
		assert.True(t, ok)
		_player, ok := sm.State.Player[player.player.ID]
		_itemID := _player.Items[0]
		assert.NotZero(t, _itemID)
	})
	t.Run("removes elements", func(t *testing.T) {
		sm := newStateMachine()
		player := sm.CreatePlayer()
		item := player.AddItem(sm)
		sm.UpdateState()
		player.RemoveItem(item.item.ID, sm)
		sm.UpdateState()
		_, ok := sm.State.Item[item.item.ID]
		assert.False(t, ok)
	})
}

func TestTree(t *testing.T) {
	t.Run("assembles expected tree", func(t *testing.T) {
		sm := newStateMachine()
		zone := sm.CreateZone()
		player1 := zone.AddPlayer(sm)
		sm.UpdateState()
		player1.GetGearScore(sm).SetLevel(1, sm)

		actual := sm.assembleTree()

		expected := newTree()
		expected.Zone = map[ZoneID]_zone{
			zone.zone.ID: {
				ID: zone.zone.ID,
				Players: []_player{
					{
						ID: player1.player.ID,
						GearScore: &_gearScore{
							ID: player1.GetGearScore(sm).gearScore.ID,
						},
						OperationKind: OperationKindUpdate,
					},
				},
				OperationKind: OperationKindUpdate,
			},
		}

		assert.Equal(t, expected, actual)
		fmt.Println("DAW")
		fmt.Printf("%+v\n", expected)
	})
	// t.Run("assembles expected tree", func(t *testing.T) {
	// 	sm := newStateMachine()

	// 	actual := newTree().assemble(sm.Patch)

	// 	expected := newTree()

	// 	assert.Equal(t, expected, actual)
	// 	fmt.Println("DAW")
	// 	fmt.Printf("%+v\n", expected)
	// })
}
