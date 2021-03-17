package statemachine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStateMachine(t *testing.T) {
	t.Run("creates elements", func(t *testing.T) {
		sm := newStateMachine()
		gearScore := sm.CreateGearScore()
		_, ok := sm.Patch.GearScore[gearScore.GetID(sm)]
		assert.True(t, ok)
	})
	t.Run("gets elements", func(t *testing.T) {
		sm := newStateMachine()
		gearScore := sm.CreateGearScore()
		_gearScore := sm.GetGearScore(gearScore.GetID(sm))
		assert.NotZero(t, _gearScore.GetID(sm))
	})
	t.Run("sets elements", func(t *testing.T) {
		sm := newStateMachine()
		gearScore := sm.CreateGearScore()
		_gearScore := sm.Patch.GearScore[gearScore.GetID(sm)]
		assert.Zero(t, _gearScore.Level)
		gearScore.SetLevel(sm, 10)
		_gearScore = sm.Patch.GearScore[gearScore.GetID(sm)]
		assert.NotZero(t, _gearScore.Level)
	})
	t.Run("deletes elements", func(t *testing.T) {
		sm := newStateMachine()
		gearScore := sm.CreateGearScore()
		sm.deleteGearScore(gearScore.GetID(sm))
		_gearScore := sm.Patch.GearScore[gearScore.GetID(sm)]
		assert.Equal(t, OperationKind(OperationKindDelete), _gearScore.OperationKind)
	})
	t.Run("adds elements", func(t *testing.T) {
		sm := newStateMachine()
		player := sm.CreatePlayer()
		item := player.AddItem(sm)
		playerItem := player.GetItems(sm)[0]
		assert.NotZero(t, playerItem.GetID(sm))
		_, ok := sm.Patch.Item[item.GetID(sm)]
		assert.True(t, ok)
	})
	t.Run("removes elements", func(t *testing.T) {
		sm := newStateMachine()
		player := sm.CreatePlayer()
		item := player.AddItem(sm)
		player.RemoveItems(sm, item.GetID(sm))
		_item := sm.Patch.Item[item.GetID(sm)]
		assert.Equal(t, OperationKind(OperationKindDelete), _item.OperationKind)
	})
}

func TestUpdateState(t *testing.T) {
	t.Run("clears patch", func(t *testing.T) {
		sm := newStateMachine()
		sm.CreateGearScore()
		sm.UpdateState()
		assert.Equal(t, len(sm.Patch.GearScore), 0)
	})
	t.Run("creates elements", func(t *testing.T) {
		sm := newStateMachine()
		gearScore := sm.CreateGearScore()
		sm.UpdateState()
		_, ok := sm.Patch.GearScore[gearScore.GetID(sm)]
		assert.False(t, ok)
		_, ok = sm.State.GearScore[gearScore.GetID(sm)]
		assert.True(t, ok)
	})
	t.Run("gets elements", func(t *testing.T) {
		sm := newStateMachine()
		gearScore := sm.CreateGearScore()
		sm.UpdateState()
		assert.Zero(t, gearScore.GetLevel(sm))
		gearScore.SetLevel(sm, 1)
		assert.Equal(t, gearScore.GetLevel(sm), 1)
	})
	t.Run("sets elements", func(t *testing.T) {
		sm := newStateMachine()
		gearScore := sm.CreateGearScore().SetLevel(sm, 1)
		sm.UpdateState()
		_gearScore := sm.State.GearScore[gearScore.GetID(sm)]
		assert.Equal(t, _gearScore.Level, 1)
		_, ok := sm.Patch.GearScore[gearScore.GetID(sm)]
		assert.False(t, ok)
	})
	t.Run("deletes elements", func(t *testing.T) {
		sm := newStateMachine()
		gearScore := sm.CreateGearScore()
		sm.UpdateState()
		sm.deleteGearScore(gearScore.GetID(sm))
		sm.UpdateState()
		_, ok := sm.State.GearScore[gearScore.GetID(sm)]
		assert.False(t, ok)
	})
	t.Run("does not delete on illegal delete element with parent", func(t *testing.T) {
		// todo
	})
	t.Run("adds elements", func(t *testing.T) {
		sm := newStateMachine()
		player := sm.CreatePlayer()
		item := player.AddItem(sm)
		sm.UpdateState()
		_, ok := sm.State.Item[item.GetID(sm)]
		assert.True(t, ok)
		_player, ok := sm.State.Player[player.GetID(sm)]
		_itemID := _player.Items[0]
		assert.NotZero(t, _itemID)
	})
	t.Run("removes elements", func(t *testing.T) {
		sm := newStateMachine()
		player := sm.CreatePlayer()
		item := player.AddItem(sm)
		sm.UpdateState()
		player.RemoveItems(sm, item.GetID(sm))
		sm.UpdateState()
		_, ok := sm.State.Item[item.GetID(sm)]
		assert.False(t, ok)
	})
}

func TestActionsOnDeletedItems(t *testing.T) {
	t.Run("does not set attribute on element which is set to be deleted", func(t *testing.T) {
		sm := newStateMachine()
		gearScore := sm.CreateGearScore()
		assert.Equal(t, 0, gearScore.GetLevel(sm))
		sm.DeleteGearScore(gearScore.GetID(sm))
		gearScore.SetLevel(sm, 1)
		assert.Equal(t, 0, gearScore.GetLevel(sm))
	})
	t.Run("does not add child on element which is set to be deleted", func(t *testing.T) {
		sm := newStateMachine()
		player := sm.CreatePlayer()
		sm.DeletePlayer(player.GetID(sm))
		item := player.AddItem(sm)
		assert.Equal(t, OperationKind(OperationKindDelete), item.item.OperationKind)
		sm.UpdateState()
		assert.Equal(t, 0, len(sm.GetPlayer(player.GetID(sm)).GetItems(sm)))
	})
	t.Run("does not remove child on element which is set to be deleted", func(t *testing.T) {
		sm := newStateMachine()
		player := sm.CreatePlayer()
		item := player.AddItem(sm)
		sm.UpdateState()
		assert.Equal(t, 1, len(sm.GetPlayer(player.GetID(sm)).GetItems(sm)))
		sm.DeletePlayer(player.GetID(sm))
		player.RemoveItems(sm, item.GetID(sm))
		assert.Equal(t, 1, len(sm.GetPlayer(player.GetID(sm)).GetItems(sm)))
	})
	t.Run("does not delete element which is a child of another element", func(t *testing.T) {
		sm := newStateMachine()
		player := sm.CreatePlayer()
		item := player.AddItem(sm)
		sm.DeleteItem(item.GetID(sm))
		assert.Equal(t, 1, len(sm.GetPlayer(player.GetID(sm)).GetItems(sm)))
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
			zone.GetID(sm): {
				ID: zone.GetID(sm),
				Players: []_player{
					{
						ID: player1.GetID(sm),
						GearScore: &_gearScore{
							ID:            player1.GetGearScore(sm).GetID(sm),
							OperationKind: OperationKindUpdate,
						},
						OperationKind: OperationKindUpdate,
						Position: &_position{
							ID:            player1.GetPosition(sm).GetID(sm),
							OperationKind: OperationKindUpdate,
						},
					},
					{
						ID: player2.GetID(sm),
						GearScore: &_gearScore{
							ID:            player2.GetGearScore(sm).GetID(sm),
							OperationKind: OperationKindUpdate,
						},
						OperationKind: OperationKindUpdate,
						Position: &_position{
							ID:            player2.GetPosition(sm).GetID(sm),
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
		player1.GetGearScore(sm).SetLevel(sm, 1)

		actual := sm.assembleTree()

		expected := newTree()
		expected.Zone = map[ZoneID]_zone{
			zone.GetID(sm): {
				ID: zone.GetID(sm),
				Players: []_player{
					{
						ID: player1.GetID(sm),
						GearScore: &_gearScore{
							ID:            player1.GetGearScore(sm).GetID(sm),
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
			zone.GetID(sm): {
				ID: zone.GetID(sm),
				Players: []_player{
					{
						ID: player1.GetID(sm),
						Items: []_item{
							{
								ID:            player1item1.GetID(sm),
								OperationKind: OperationKindUpdate,
								GearScore: &_gearScore{
									ID:            player1item1.GetGearScore(sm).GetID(sm),
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

		player1.RemoveItems(sm, player1item2.GetID(sm))
		actual := sm.assembleTree()

		expected := newTree()
		expected.Zone = map[ZoneID]_zone{
			zone.GetID(sm): {
				ID: zone.GetID(sm),
				Players: []_player{
					{
						ID: player1.GetID(sm),
						Items: []_item{
							{
								ID:            player1item2.GetID(sm),
								OperationKind: OperationKindDelete,
								GearScore: &_gearScore{
									ID:            player1item2.item.GearScore,
									OperationKind: OperationKindDelete,
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
}
