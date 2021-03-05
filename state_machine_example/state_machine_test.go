package statemachine

import (
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
	t.Run("assembles elements in a tree", func(t *testing.T) {
		sm := newStateMachine()
		zone := sm.CreateZone()
		player1 := zone.AddPlayer(sm)
		player2 := zone.AddPlayer(sm)

		actual := sm.assembleTree()

		expected := newTree()
		expected.Zone = map[ZoneID]_zone{
			zone.zone.ID: {
				ID: zone.zone.ID,
				Players: []_player{
					{
						ID: player1.player.ID,
						GearScore: &_gearScore{
							ID:            player1.GetGearScore(sm).gearScore.ID,
							OperationKind: OperationKindUpdate,
						},
						OperationKind: OperationKindUpdate,
						Position: &_position{
							ID:            player1.GetPosition(sm).position.ID,
							OperationKind: OperationKindUpdate,
						},
					},
					{
						ID: player2.player.ID,
						GearScore: &_gearScore{
							ID:            player2.GetGearScore(sm).gearScore.ID,
							OperationKind: OperationKindUpdate,
						},
						OperationKind: OperationKindUpdate,
						Position: &_position{
							ID:            player2.GetPosition(sm).position.ID,
							OperationKind: OperationKindUpdate,
						},
					},
				},
				OperationKind: OperationKindUpdate,
			},
		}

		_actual, _ := actual.MarshalJSON()
		_expected, _ := expected.MarshalJSON()

		assert.Equal(t, string(_expected), string(_actual))
	})
	t.Run("assembles tree based on changed GearScore", func(t *testing.T) {
		sm := newStateMachine()
		zone := sm.CreateZone()
		player1 := zone.AddPlayer(sm)
		_ = zone.AddPlayer(sm)
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
							ID:            player1.GetGearScore(sm).gearScore.ID,
							Level:         1,
							OperationKind: OperationKindUpdate,
						},
						OperationKind: OperationKindUpdate,
					},
				},
				OperationKind: OperationKindUpdate,
			},
		}

		_actual, _ := actual.MarshalJSON()
		_expected, _ := expected.MarshalJSON()

		assert.Equal(t, string(_expected), string(_actual))
	})
	t.Run("assembles tree based on added item", func(t *testing.T) {
		sm := newStateMachine()
		zone := sm.CreateZone()
		player1 := zone.AddPlayer(sm)
		_ = zone.AddPlayer(sm)
		sm.UpdateState()
		player1item1 := player1.AddItem(sm)

		actual := sm.assembleTree()

		expected := newTree()
		expected.Zone = map[ZoneID]_zone{
			zone.zone.ID: {
				ID: zone.zone.ID,
				Players: []_player{
					{
						ID: player1.player.ID,
						Items: []_item{
							{
								ID:            player1item1.item.ID,
								OperationKind: OperationKindUpdate,
								GearScore: &_gearScore{
									ID:            player1item1.GetGearScore(sm).gearScore.ID,
									OperationKind: OperationKindUpdate,
								},
							},
						},
						OperationKind: OperationKindUpdate,
					},
				},
				OperationKind: OperationKindUpdate,
			},
		}

		_actual, _ := actual.MarshalJSON()
		_expected, _ := expected.MarshalJSON()

		assert.Equal(t, string(_expected), string(_actual))
	})
	t.Run("assembles tree based on removed item", func(t *testing.T) {
		sm := newStateMachine()
		zone := sm.CreateZone()
		player1 := zone.AddPlayer(sm)
		_ = zone.AddPlayer(sm)
		_ = player1.AddItem(sm)
		player1item2 := player1.AddItem(sm)

		sm.UpdateState()

		player1.RemoveItem(player1item2.item.ID, sm)
		actual := sm.assembleTree()

		expected := newTree()
		expected.Zone = map[ZoneID]_zone{
			zone.zone.ID: {
				ID: zone.zone.ID,
				Players: []_player{
					{
						ID: player1.player.ID,
						Items: []_item{
							{
								ID:            player1item2.item.ID,
								OperationKind: OperationKindDelete,
							},
						},
						OperationKind: OperationKindUpdate,
					},
				},
				OperationKind: OperationKindUpdate,
			},
		}

		_actual, _ := actual.MarshalJSON()
		_expected, _ := expected.MarshalJSON()

		assert.Equal(t, string(_expected), string(_actual))
	})
}
