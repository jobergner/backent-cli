package state

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEngine(t *testing.T) {
	t.Run("creates elements", func(t *testing.T) {
		se := newEngine()
		gearScore := se.CreateGearScore()
		_, ok := se.Patch.GearScore[gearScore.ID(se)]
		assert.True(t, ok)
	})
	t.Run("gets elements", func(t *testing.T) {
		se := newEngine()
		gearScore := se.CreateGearScore()
		_gearScore := se.GearScore(gearScore.ID(se))
		assert.NotZero(t, _gearScore.ID(se))
	})
	t.Run("sets elements", func(t *testing.T) {
		se := newEngine()
		gearScore := se.CreateGearScore()
		_gearScore := se.Patch.GearScore[gearScore.ID(se)]
		assert.Zero(t, _gearScore.Level)
		gearScore.SetLevel(se, 10)
		_gearScore = se.Patch.GearScore[gearScore.ID(se)]
		assert.NotZero(t, _gearScore.Level)
	})
	t.Run("deletes elements", func(t *testing.T) {
		se := newEngine()
		gearScore := se.CreateGearScore()
		se.deleteGearScore(gearScore.ID(se))
		_gearScore := se.Patch.GearScore[gearScore.ID(se)]
		assert.Equal(t, OperationKind(OperationKindDelete), _gearScore.OperationKind)
	})
	t.Run("adds elements", func(t *testing.T) {
		se := newEngine()
		player := se.CreatePlayer()
		item := player.AddItem(se)
		playerItem := player.Items(se)[0]
		assert.NotZero(t, playerItem.ID(se))
		_, ok := se.Patch.Item[item.ID(se)]
		assert.True(t, ok)
	})
	t.Run("removes elements", func(t *testing.T) {
		se := newEngine()
		player := se.CreatePlayer()
		item := player.AddItem(se)
		player.RemoveItems(se, item.ID(se))
		_item := se.Patch.Item[item.ID(se)]
		assert.Equal(t, OperationKind(OperationKindDelete), _item.OperationKind)
	})
}

func TestUpdateState(t *testing.T) {
	t.Run("clears patch", func(t *testing.T) {
		se := newEngine()
		se.CreateGearScore()
		se.UpdateState()
		assert.Equal(t, len(se.Patch.GearScore), 0)
	})
	t.Run("creates elements", func(t *testing.T) {
		se := newEngine()
		gearScore := se.CreateGearScore()
		se.UpdateState()
		_, ok := se.Patch.GearScore[gearScore.ID(se)]
		assert.False(t, ok)
		_, ok = se.State.GearScore[gearScore.ID(se)]
		assert.True(t, ok)
	})
	t.Run("gets elements", func(t *testing.T) {
		se := newEngine()
		gearScore := se.CreateGearScore()
		se.UpdateState()
		assert.Zero(t, gearScore.Level(se))
		gearScore.SetLevel(se, 1)
		assert.Equal(t, gearScore.Level(se), 1)
	})
	t.Run("sets elements", func(t *testing.T) {
		se := newEngine()
		gearScore := se.CreateGearScore().SetLevel(se, 1)
		se.UpdateState()
		_gearScore := se.State.GearScore[gearScore.ID(se)]
		assert.Equal(t, _gearScore.Level, 1)
		_, ok := se.Patch.GearScore[gearScore.ID(se)]
		assert.False(t, ok)
	})
	t.Run("deletes elements", func(t *testing.T) {
		se := newEngine()
		gearScore := se.CreateGearScore()
		se.UpdateState()
		se.deleteGearScore(gearScore.ID(se))
		se.UpdateState()
		_, ok := se.State.GearScore[gearScore.ID(se)]
		assert.False(t, ok)
	})
	t.Run("does not delete on illegal delete element with parent", func(t *testing.T) {
		// todo
	})
	t.Run("adds elements", func(t *testing.T) {
		se := newEngine()
		player := se.CreatePlayer()
		item := player.AddItem(se)
		se.UpdateState()
		_, ok := se.State.Item[item.ID(se)]
		assert.True(t, ok)
		_player, ok := se.State.Player[player.ID(se)]
		_itemID := _player.Items[0]
		assert.NotZero(t, _itemID)
	})
	t.Run("removes elements", func(t *testing.T) {
		se := newEngine()
		player := se.CreatePlayer()
		item := player.AddItem(se)
		se.UpdateState()
		player.RemoveItems(se, item.ID(se))
		se.UpdateState()
		_, ok := se.State.Item[item.ID(se)]
		assert.False(t, ok)
	})
}

func TestActionsOnDeletedItems(t *testing.T) {
	t.Run("does not set attribute on element which is set to be deleted", func(t *testing.T) {
		se := newEngine()
		gearScore := se.CreateGearScore()
		assert.Equal(t, 0, gearScore.Level(se))
		se.DeleteGearScore(gearScore.ID(se))
		gearScore.SetLevel(se, 1)
		assert.Equal(t, 0, gearScore.Level(se))
	})
	t.Run("does not add child on element which is set to be deleted", func(t *testing.T) {
		se := newEngine()
		player := se.CreatePlayer()
		se.DeletePlayer(player.ID(se))
		item := player.AddItem(se)
		assert.Equal(t, OperationKind(OperationKindDelete), item.item.OperationKind)
		se.UpdateState()
		assert.Equal(t, 0, len(se.Player(player.ID(se)).Items(se)))
	})
	t.Run("does not remove child on element which is set to be deleted", func(t *testing.T) {
		se := newEngine()
		player := se.CreatePlayer()
		item := player.AddItem(se)
		se.UpdateState()
		assert.Equal(t, 1, len(se.Player(player.ID(se)).Items(se)))
		se.DeletePlayer(player.ID(se))
		player.RemoveItems(se, item.ID(se))
		assert.Equal(t, 1, len(se.Player(player.ID(se)).Items(se)))
	})
	t.Run("does not delete element which is a child of another element", func(t *testing.T) {
		se := newEngine()
		player := se.CreatePlayer()
		item := player.AddItem(se)
		se.DeleteItem(item.ID(se))
		assert.Equal(t, 1, len(se.Player(player.ID(se)).Items(se)))
	})
}

func TestTree(t *testing.T) {
	t.Run("assembles elements in a tree", func(t *testing.T) {
		se := newEngine()
		zone := se.CreateZone()
		player1 := zone.AddPlayer(se)
		player2 := zone.AddPlayer(se)

		actual := se.assembleTree()

		expected := newTree()
		expected.Zone = map[ZoneID]_zone{
			zone.ID(se): {
				ID: zone.ID(se),
				Players: []_player{
					{
						ID: player1.ID(se),
						GearScore: &_gearScore{
							ID:            player1.GearScore(se).ID(se),
							OperationKind: OperationKindUpdate,
						},
						OperationKind: OperationKindUpdate,
						Position: &_position{
							ID:            player1.Position(se).ID(se),
							OperationKind: OperationKindUpdate,
						},
					},
					{
						ID: player2.ID(se),
						GearScore: &_gearScore{
							ID:            player2.GearScore(se).ID(se),
							OperationKind: OperationKindUpdate,
						},
						OperationKind: OperationKindUpdate,
						Position: &_position{
							ID:            player2.Position(se).ID(se),
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
		se := newEngine()
		zone := se.CreateZone()
		player1 := zone.AddPlayer(se)
		_ = zone.AddPlayer(se)
		se.UpdateState()
		player1.GearScore(se).SetLevel(se, 1)

		actual := se.assembleTree()

		expected := newTree()
		expected.Zone = map[ZoneID]_zone{
			zone.ID(se): {
				ID: zone.ID(se),
				Players: []_player{
					{
						ID: player1.ID(se),
						GearScore: &_gearScore{
							ID:            player1.GearScore(se).ID(se),
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
		se := newEngine()
		zone := se.CreateZone()
		player1 := zone.AddPlayer(se)
		_ = zone.AddPlayer(se)
		se.UpdateState()
		player1item1 := player1.AddItem(se)

		actual := se.assembleTree()

		expected := newTree()
		expected.Zone = map[ZoneID]_zone{
			zone.ID(se): {
				ID: zone.ID(se),
				Players: []_player{
					{
						ID: player1.ID(se),
						Items: []_item{
							{
								ID:            player1item1.ID(se),
								OperationKind: OperationKindUpdate,
								GearScore: &_gearScore{
									ID:            player1item1.GearScore(se).ID(se),
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
		se := newEngine()
		zone := se.CreateZone()
		player1 := zone.AddPlayer(se)
		_ = zone.AddPlayer(se)
		_ = player1.AddItem(se)
		player1item2 := player1.AddItem(se)

		se.UpdateState()

		player1.RemoveItems(se, player1item2.ID(se))
		actual := se.assembleTree()

		expected := newTree()
		expected.Zone = map[ZoneID]_zone{
			zone.ID(se): {
				ID: zone.ID(se),
				Players: []_player{
					{
						ID: player1.ID(se),
						Items: []_item{
							{
								ID:            player1item2.ID(se),
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