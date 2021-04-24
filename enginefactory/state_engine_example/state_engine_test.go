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
		assert.Equal(t, OperationKindDelete, _gearScore.OperationKind_)
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
		assert.Equal(t, OperationKindDelete, _item.OperationKind_)
	})
	t.Run("does not allow creation if element by 'getting' an non-existing one", func(t *testing.T) {
		se := newEngine()
		gearScore := se.GearScore(GearScoreID(999))
		gearScore.SetLevel(se, 2)
		gearScoreLevel := se.Patch.GearScore[GearScoreID(0)].Level
		assert.NotEqual(t, gearScoreLevel, 2)
	})
}

func TestReferences(t *testing.T) {
	t.Run("deletes reference off element if referenced element gets deleted (1/3)", func(t *testing.T) {
		se := newEngine()
		player1 := se.CreatePlayer()
		player2 := se.CreatePlayer()
		player1.AddGuildMember(se, player2.ID(se))

		se.UpdateState()
		se.DeletePlayer(player2.ID(se))
		player1_updated, ok := se.Patch.Player[player1.ID(se)]
		assert.True(t, ok)
		assert.Equal(t, 0, len(player1_updated.GuildMembers))
	})
	t.Run("deletes reference off element if referenced element gets deleted (2/3)", func(t *testing.T) {
		se := newEngine()
		player := se.CreatePlayer()
		item := se.CreateItem()
		item.SetBoundTo(se, player.ID(se))
		se.UpdateState()

		se.deletePlayer(player.ID(se))
		_, ok := se.Patch.Item[item.ID(se)]
		assert.True(t, ok)
		_, isSet := se.Item(item.ID(se)).BoundTo(se)
		assert.False(t, isSet)
	})
	t.Run("deletes reference off element if referenced element gets deleted (3/3)", func(t *testing.T) {
		se := newEngine()
		zone := se.CreateZone()
		player1 := se.CreatePlayer()
		player2 := zone.AddPlayer(se)
		player1.AddGuildMember(se, player2.ID(se))

		se.UpdateState()
		zone.RemovePlayers(se, player2.ID(se))
		player1_updated, ok := se.Patch.Player[player1.ID(se)]
		assert.True(t, ok)
		assert.Equal(t, 0, len(player1_updated.GuildMembers))
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
		assert.Equal(t, OperationKindDelete, item.item.OperationKind_)
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
		expected.Zone = map[ZoneID]Zone{
			zone.ID(se): {
				ID: zone.ID(se),
				Players: []Player{
					{
						ID: player1.ID(se),
						GearScore: &GearScore{
							ID:             player1.GearScore(se).ID(se),
							OperationKind_: OperationKindUpdate,
						},
						OperationKind_: OperationKindUpdate,
						Position: &Position{
							ID:             player1.Position(se).ID(se),
							OperationKind_: OperationKindUpdate,
						},
					},
					{
						ID: player2.ID(se),
						GearScore: &GearScore{
							ID:             player2.GearScore(se).ID(se),
							OperationKind_: OperationKindUpdate,
						},
						OperationKind_: OperationKindUpdate,
						Position: &Position{
							ID:             player2.Position(se).ID(se),
							OperationKind_: OperationKindUpdate,
						},
					},
				},
				OperationKind_: OperationKindUpdate,
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
		expected.Zone = map[ZoneID]Zone{
			zone.ID(se): {
				ID: zone.ID(se),
				Players: []Player{
					{
						ID: player1.ID(se),
						GearScore: &GearScore{
							ID:             player1.GearScore(se).ID(se),
							Level:          1,
							OperationKind_: OperationKindUpdate,
						},
						OperationKind_: OperationKindUpdate,
					},
				},
				OperationKind_: OperationKindUpdate,
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
		expected.Zone = map[ZoneID]Zone{
			zone.ID(se): {
				ID: zone.ID(se),
				Players: []Player{
					{
						ID: player1.ID(se),
						Items: []Item{
							{
								ID:             player1item1.ID(se),
								OperationKind_: OperationKindUpdate,
								GearScore: &GearScore{
									ID:             player1item1.GearScore(se).ID(se),
									OperationKind_: OperationKindUpdate,
								},
							},
						},
						OperationKind_: OperationKindUpdate,
					},
				},
				OperationKind_: OperationKindUpdate,
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
		expected.Zone = map[ZoneID]Zone{
			zone.ID(se): {
				ID: zone.ID(se),
				Players: []Player{
					{
						ID: player1.ID(se),
						Items: []Item{
							{
								ID:             player1item2.ID(se),
								OperationKind_: OperationKindDelete,
								GearScore: &GearScore{
									ID:             player1item2.item.GearScore,
									OperationKind_: OperationKindDelete,
								},
							},
						},
						OperationKind_: OperationKindUpdate,
					},
				},
				OperationKind_: OperationKindUpdate,
			},
		}

		_actual, _ := actual.MarshalJSON()
		_expected, _ := expected.MarshalJSON()

		assert.Equal(t, string(_expected), string(_actual))
	})
}

func TestDiffPlayerIDs(t *testing.T) {
	t.Run("", func(t *testing.T) {
		inputCurrentIDs := []PlayerID{}
		inputNextIDs := []PlayerID{}

		expected := []PlayerID{}

		actual := mergePlayerIDs(inputCurrentIDs, inputNextIDs)

		assert.ElementsMatch(t, expected, actual)
	})
	t.Run("", func(t *testing.T) {
		inputCurrentIDs := []PlayerID{1, 2, 3}
		inputNextIDs := []PlayerID{1, 2, 3}

		expected := []PlayerID{1, 2, 3}

		actual := mergePlayerIDs(inputCurrentIDs, inputNextIDs)

		assert.Equal(t, expected, actual)
	})
	t.Run("", func(t *testing.T) {
		inputCurrentIDs := []PlayerID{1, 2, 3}
		inputNextIDs := []PlayerID{1, 3}

		expected := []PlayerID{1, 2, 3}

		actual := mergePlayerIDs(inputCurrentIDs, inputNextIDs)

		assert.Equal(t, expected, actual)
	})
	t.Run("", func(t *testing.T) {
		inputCurrentIDs := []PlayerID{1, 2, 3}
		inputNextIDs := []PlayerID{1, 3, 4}

		expected := []PlayerID{1, 2, 3, 4}

		actual := mergePlayerIDs(inputCurrentIDs, inputNextIDs)

		assert.Equal(t, expected, actual)
	})
	t.Run("", func(t *testing.T) {
		inputCurrentIDs := []PlayerID{1, 2, 3}
		inputNextIDs := []PlayerID{}

		expected := []PlayerID{1, 2, 3}

		actual := mergePlayerIDs(inputCurrentIDs, inputNextIDs)

		assert.Equal(t, expected, actual)
	})
	t.Run("", func(t *testing.T) {
		inputCurrentIDs := []PlayerID{}
		inputNextIDs := []PlayerID{1, 2, 3}

		expected := []PlayerID{1, 2, 3}

		actual := mergePlayerIDs(inputCurrentIDs, inputNextIDs)

		assert.Equal(t, expected, actual)
	})
}
