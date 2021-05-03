package state

import (
	"bar-cli/testutils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEngine(t *testing.T) {
	t.Run("creates elements", func(t *testing.T) {
		se := newEngine()
		gearScore := se.CreateGearScore()
		_, ok := se.Patch.GearScore[gearScore.ID()]
		assert.True(t, ok)
	})
	t.Run("gets elements", func(t *testing.T) {
		se := newEngine()
		gearScore := se.CreateGearScore()
		_gearScore := se.GearScore(gearScore.ID())
		assert.NotZero(t, _gearScore.ID())
	})
	t.Run("sets elements", func(t *testing.T) {
		se := newEngine()
		gearScore := se.CreateGearScore()
		_gearScore := se.Patch.GearScore[gearScore.ID()]
		assert.Zero(t, _gearScore.Level)
		gearScore.SetLevel(10)
		_gearScore = se.Patch.GearScore[gearScore.ID()]
		assert.NotZero(t, _gearScore.Level)
	})
	t.Run("deletes elements", func(t *testing.T) {
		se := newEngine()
		gearScore := se.CreateGearScore()
		se.deleteGearScore(gearScore.ID())
		_gearScore := se.Patch.GearScore[gearScore.ID()]
		assert.Equal(t, OperationKindDelete, _gearScore.OperationKind_)
	})
	t.Run("adds elements", func(t *testing.T) {
		se := newEngine()
		player := se.CreatePlayer()
		item := player.AddItem()
		playerItem := player.Items()[0]
		assert.NotZero(t, playerItem.ID())
		_, ok := se.Patch.Item[item.ID()]
		assert.True(t, ok)
	})
	t.Run("removes elements", func(t *testing.T) {
		se := newEngine()
		player := se.CreatePlayer()
		item := player.AddItem()
		player.RemoveItems(item.ID())
		_item := se.Patch.Item[item.ID()]
		assert.Equal(t, OperationKindDelete, _item.OperationKind_)
	})
}

func TestReferences(t *testing.T) {
	t.Run("deletes reference off element if referenced element gets deleted (1/3)", func(t *testing.T) {
		se := newEngine()
		player1 := se.CreatePlayer()
		player2 := se.CreatePlayer()
		player1.AddGuildMember(player2.ID())

		se.UpdateState()
		se.DeletePlayer(player2.ID())
		player1_updated, ok := se.Patch.Player[player1.ID()]
		assert.True(t, ok)
		assert.Equal(t, 0, len(player1_updated.GuildMembers))
	})
	t.Run("deletes reference off element if referenced element gets deleted (2/3)", func(t *testing.T) {
		se := newEngine()
		player := se.CreatePlayer()
		item := se.CreateItem()
		item.SetBoundTo(player.ID())
		se.UpdateState()

		se.deletePlayer(player.ID())
		_, ok := se.Patch.Item[item.ID()]
		assert.True(t, ok)
		_, isSet := se.Item(item.ID()).BoundTo()
		assert.False(t, isSet)
	})
	t.Run("deletes reference off element if referenced element gets deleted (3/3)", func(t *testing.T) {
		se := newEngine()
		zone := se.CreateZone()
		player1 := se.CreatePlayer()
		player2 := zone.AddPlayer()
		player1.AddGuildMember(player2.ID())

		se.UpdateState()
		zone.RemovePlayers(player2.ID())
		player1_updated, ok := se.Patch.Player[player1.ID()]
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
		_, ok := se.Patch.GearScore[gearScore.ID()]
		assert.False(t, ok)
		_, ok = se.State.GearScore[gearScore.ID()]
		assert.True(t, ok)
	})
	t.Run("gets elements", func(t *testing.T) {
		se := newEngine()
		gearScore := se.CreateGearScore()
		se.UpdateState()
		assert.Zero(t, gearScore.Level())
		gearScore.SetLevel(1)
		assert.Equal(t, gearScore.Level(), 1)
	})
	t.Run("sets elements", func(t *testing.T) {
		se := newEngine()
		gearScore := se.CreateGearScore().SetLevel(1)
		se.UpdateState()
		_gearScore := se.State.GearScore[gearScore.ID()]
		assert.Equal(t, _gearScore.Level, 1)
		_, ok := se.Patch.GearScore[gearScore.ID()]
		assert.False(t, ok)
	})
	t.Run("deletes elements", func(t *testing.T) {
		se := newEngine()
		gearScore := se.CreateGearScore()
		se.UpdateState()
		se.deleteGearScore(gearScore.ID())
		se.UpdateState()
		_, ok := se.State.GearScore[gearScore.ID()]
		assert.False(t, ok)
	})
	t.Run("does not delete on illegal delete element with parent", func(t *testing.T) {
		// todo
	})
	t.Run("adds elements", func(t *testing.T) {
		se := newEngine()
		player := se.CreatePlayer()
		item := player.AddItem()
		se.UpdateState()
		_, ok := se.State.Item[item.ID()]
		assert.True(t, ok)
		_player, ok := se.State.Player[player.ID()]
		_itemID := _player.Items[0]
		assert.NotZero(t, _itemID)
	})
	t.Run("removes elements", func(t *testing.T) {
		se := newEngine()
		player := se.CreatePlayer()
		item := player.AddItem()
		se.UpdateState()
		player.RemoveItems(item.ID())
		se.UpdateState()
		_, ok := se.State.Item[item.ID()]
		assert.False(t, ok)
	})
}

func TestActionsOnDeletedItems(t *testing.T) {
	t.Run("does not set attribute on element which is set to be deleted", func(t *testing.T) {
		se := newEngine()
		gearScore := se.CreateGearScore()
		assert.Equal(t, 0, gearScore.Level())
		se.DeleteGearScore(gearScore.ID())
		gearScore.SetLevel(1)
		assert.Equal(t, 0, gearScore.Level())
	})
	t.Run("does not add child on element which is set to be deleted", func(t *testing.T) {
		se := newEngine()
		player := se.CreatePlayer()
		se.DeletePlayer(player.ID())
		item := player.AddItem()
		assert.Equal(t, OperationKindDelete, item.item.OperationKind_)
		assert.Equal(t, 0, len(se.Player(player.ID()).Items()))
	})
	t.Run("does not remove child on element which is set to be deleted", func(t *testing.T) {
		se := newEngine()
		player := se.CreatePlayer()
		item := player.AddItem()
		se.UpdateState()
		assert.Equal(t, 1, len(se.Player(player.ID()).Items()))
		se.DeletePlayer(player.ID())
		player.RemoveItems(item.ID())
		assert.Equal(t, 1, len(se.Player(player.ID()).Items()))
	})
	t.Run("does not delete element which is a child of another element", func(t *testing.T) {
		se := newEngine()
		player := se.CreatePlayer()
		item := player.AddItem()
		se.DeleteItem(item.ID())
		assert.Equal(t, 1, len(se.Player(player.ID()).Items()))
	})
}

func newTreeTest(define func(*Engine, *Tree), onFail func(errText string)) {
	se := newEngine()
	expectedTree := newTree()
	define(se, &expectedTree)
	actualTree := se.assembleTree()

	if !assert.ObjectsAreEqual(expectedTree, actualTree) {
		actual, _ := actualTree.MarshalJSON()
		expected, _ := expectedTree.MarshalJSON()
		actualString := string(actual)
		expectedString := string(expected)
		onFail(testutils.Diff(actualString, expectedString))
	}
}

func TestTree(t *testing.T) {
	t.Run("assembles elements in a tree", func(t *testing.T) {
		newTreeTest(
			func(se *Engine, expectedTree *Tree) {
				zone := se.CreateZone()
				player1 := zone.AddPlayer()
				player2 := zone.AddPlayer()

				expectedTree.Zone = map[ZoneID]Zone{
					zone.ID(): {
						ID: zone.ID(),
						Players: []Player{
							{
								ID: player1.ID(),
								GearScore: &GearScore{
									ID:             player1.GearScore().ID(),
									OperationKind_: OperationKindUpdate,
								},
								OperationKind_: OperationKindUpdate,
								Position: &Position{
									ID:             player1.Position().ID(),
									OperationKind_: OperationKindUpdate,
								},
							},
							{
								ID: player2.ID(),
								GearScore: &GearScore{
									ID:             player2.GearScore().ID(),
									OperationKind_: OperationKindUpdate,
								},
								OperationKind_: OperationKindUpdate,
								Position: &Position{
									ID:             player2.Position().ID(),
									OperationKind_: OperationKindUpdate,
								},
							},
						},
						OperationKind_: OperationKindUpdate,
					},
				}
			},
			func(errText string) {
				t.Errorf(errText)
			},
		)
	})
	t.Run("assembles tree based on changed GearScore", func(t *testing.T) {
		newTreeTest(
			func(se *Engine, expectedTree *Tree) {
				zone := se.CreateZone()
				player1 := zone.AddPlayer()
				_ = zone.AddPlayer()
				se.UpdateState()
				player1.GearScore().SetLevel(1)

				expectedTree.Zone = map[ZoneID]Zone{
					zone.ID(): {
						ID: zone.ID(),
						Players: []Player{
							{
								ID: player1.ID(),
								GearScore: &GearScore{
									ID:             player1.GearScore().ID(),
									Level:          1,
									OperationKind_: OperationKindUpdate,
								},
								OperationKind_: OperationKindUnchanged,
							},
						},
						OperationKind_: OperationKindUnchanged,
					},
				}

			},
			func(errText string) {
				t.Errorf(errText)
			},
		)
	})
	t.Run("assembles tree based on added item", func(t *testing.T) {
		newTreeTest(
			func(se *Engine, expectedTree *Tree) {
				zone := se.CreateZone()
				player1 := zone.AddPlayer()
				_ = zone.AddPlayer()
				se.UpdateState()
				player1item1 := player1.AddItem()

				expectedTree.Zone = map[ZoneID]Zone{
					zone.ID(): {
						ID: zone.ID(),
						Players: []Player{
							{
								ID: player1.ID(),
								Items: []Item{
									{
										ID:             player1item1.ID(),
										OperationKind_: OperationKindUpdate,
										GearScore: &GearScore{
											ID:             player1item1.GearScore().ID(),
											OperationKind_: OperationKindUpdate,
										},
									},
								},
								OperationKind_: OperationKindUpdate,
							},
						},
						OperationKind_: OperationKindUnchanged,
					},
				}
			},
			func(errText string) {
				t.Errorf(errText)
			},
		)
	})
	t.Run("assembles tree based on removed item", func(t *testing.T) {
		newTreeTest(
			func(se *Engine, expectedTree *Tree) {
				zone := se.CreateZone()
				player1 := zone.AddPlayer()
				_ = zone.AddPlayer()
				_ = player1.AddItem()
				player1item2 := player1.AddItem()

				se.UpdateState()

				player1.RemoveItems(player1item2.ID())

				expectedTree.Zone = map[ZoneID]Zone{
					zone.ID(): {
						ID: zone.ID(),
						Players: []Player{
							{
								ID: player1.ID(),
								Items: []Item{
									{
										ID:             player1item2.ID(),
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
						OperationKind_: OperationKindUnchanged,
					},
				}
			},
			func(errText string) {
				t.Errorf(errText)
			},
		)
	})
	t.Run("includes element which has reference of updating element", func(t *testing.T) {
		newTreeTest(
			func(se *Engine, expectedTree *Tree) {
				item := se.createItem(false).SetName("myItem")
				player := se.createPlayer(false)
				item.SetBoundTo(player.ID())

				se.UpdateState()

				playerItem := player.AddItem()

				expectedTree.Item = map[ItemID]Item{
					item.ID(): {
						ID:             item.ID(),
						Name:           "myItem",
						BoundTo:        &ElementReference{OperationKindUnchanged, int(player.ID()), ElementKindPlayer, ReferencedDataModified},
						OperationKind_: OperationKindUnchanged,
					},
				}
				expectedTree.Player = map[PlayerID]Player{
					player.ID(): {
						ID: player.ID(),
						Items: []Item{
							{
								ID:             playerItem.ID(),
								OperationKind_: OperationKindUpdate,
								GearScore: &GearScore{
									ID:             playerItem.GearScore().ID(),
									OperationKind_: OperationKindUpdate,
								},
							},
						},
						OperationKind_: OperationKindUpdate,
					},
				}

			},
			func(errText string) {
				t.Errorf(errText)
			},
		)
	})
	t.Run("includes elements which have references of updating elements", func(t *testing.T) {
		newTreeTest(
			func(se *Engine, expectedTree *Tree) {
				player1 := se.createPlayer(false)
				player2 := se.createPlayer(false)
				player3 := se.createPlayer(false)

				player2.AddGuildMember(player1.ID())
				player3.AddGuildMember(player1.ID())

				se.UpdateState()

				item := player1.AddItem()
				player1.AddGuildMember(player2.ID())

				expectedTree.Player = map[PlayerID]Player{
					player1.ID(): {
						ID: player1.ID(),
						Items: []Item{
							{
								ID:             item.ID(),
								BoundTo:        nil,
								GearScore:      &GearScore{ID: item.GearScore().ID(), OperationKind_: OperationKindUpdate},
								OperationKind_: OperationKindUpdate,
							},
						},
						GuildMembers: []ElementReference{
							{
								OperationKind:        OperationKindUpdate,
								ElementID:            int(player2.ID()),
								ElementKind:          ElementKindPlayer,
								ReferencedDataStatus: ReferencedDataModified,
							},
						},
						OperationKind_: OperationKindUpdate,
					},
					player2.ID(): {
						ID:             player2.ID(),
						OperationKind_: OperationKindUnchanged,
						GuildMembers: []ElementReference{
							{
								OperationKind:        OperationKindUnchanged,
								ElementID:            int(player1.ID()),
								ElementKind:          ElementKindPlayer,
								ReferencedDataStatus: ReferencedDataModified,
							},
						},
					},
					player3.ID(): {
						ID:             player3.ID(),
						OperationKind_: OperationKindUnchanged,
						GuildMembers: []ElementReference{
							{
								OperationKind:        OperationKindUnchanged,
								ElementID:            int(player1.ID()),
								ElementKind:          ElementKindPlayer,
								ReferencedDataStatus: ReferencedDataModified,
							},
						},
					},
				}

			},
			func(errText string) {
				t.Errorf(errText)
			},
		)
	})
	t.Run("does not break when element references itself", func(t *testing.T) {
		newTreeTest(
			func(se *Engine, expectedTree *Tree) {
				player := se.createPlayer(false)
				player.AddGuildMember(player.ID())

				se.UpdateState()

				player.AddGuildMember(player.ID())

				expectedTree.Player = map[PlayerID]Player{
					player.ID(): {
						ID: player.ID(),
						GuildMembers: []ElementReference{
							{
								OperationKind:        OperationKindUnchanged,
								ElementID:            int(player.ID()),
								ElementKind:          ElementKindPlayer,
								ReferencedDataStatus: ReferencedDataModified,
							},
							{
								OperationKind:        OperationKindUpdate,
								ElementID:            int(player.ID()),
								ElementKind:          ElementKindPlayer,
								ReferencedDataStatus: ReferencedDataModified,
							},
						},
						OperationKind_: OperationKindUpdate,
					},
				}

			},
			func(errText string) {
				t.Errorf(errText)
			},
		)
	})
	t.Run("includes all elements in a reference chain", func(t *testing.T) {
		newTreeTest(
			func(se *Engine, expectedTree *Tree) {
				player := se.createPlayer(false)
				item := se.createItem(false)
				equipmentSet := se.createEquipmentSet()

				player.AddEquipmentSet(equipmentSet.ID())
				equipmentSet.AddEquipment(item.ID())
				item.SetBoundTo(player.ID())

				se.UpdateState()

				item.SetName("myName")

				expectedTree.Item = map[ItemID]Item{
					item.ID(): {
						ID:   item.ID(),
						Name: "myName",
						BoundTo: &ElementReference{
							OperationKind:        OperationKindUnchanged,
							ElementID:            int(player.ID()),
							ElementKind:          ElementKindPlayer,
							ReferencedDataStatus: ReferencedDataModified,
						},
						OperationKind_: OperationKindUpdate,
					},
				}
				expectedTree.Player = map[PlayerID]Player{
					player.ID(): {
						ID: player.ID(),
						EquipmentSets: []ElementReference{
							{
								OperationKind:        OperationKindUnchanged,
								ElementID:            int(equipmentSet.ID()),
								ElementKind:          ElementKindEquipmentSet,
								ReferencedDataStatus: ReferencedDataModified,
							},
						},
						OperationKind_: OperationKindUnchanged,
					},
				}
				expectedTree.EquipmentSet = map[EquipmentSetID]EquipmentSet{
					equipmentSet.ID(): {
						ID: equipmentSet.ID(),
						Equipment: []ElementReference{
							{
								OperationKind:        OperationKindUnchanged,
								ElementID:            int(item.ID()),
								ElementKind:          ElementKindItem,
								ReferencedDataStatus: ReferencedDataModified,
							},
						},
						OperationKind_: OperationKindUnchanged,
					},
				}

			},
			func(errText string) {
				t.Errorf(errText)
			},
		)
	})
	t.Run("recursively travels tree to find if any downstream data has updated", func(t *testing.T) {
		newTreeTest(
			func(se *Engine, expectedTree *Tree) {
				player1 := se.createPlayer(false)
				player2 := se.createPlayer(false)
				item := se.createItem(false)

				item.SetBoundTo(player1.ID())
				player2.AddGuildMember(player1.ID())

				se.UpdateState()

				player1.GearScore().SetLevel(1)

				expectedTree.Item = map[ItemID]Item{
					item.ID(): {
						ID: item.ID(),
						BoundTo: &ElementReference{
							OperationKind:        OperationKindUnchanged,
							ElementID:            int(player1.ID()),
							ElementKind:          ElementKindPlayer,
							ReferencedDataStatus: ReferencedDataModified,
						},
						OperationKind_: OperationKindUnchanged,
					},
				}
				expectedTree.Player = map[PlayerID]Player{
					player1.ID(): {
						ID:             player1.ID(),
						OperationKind_: OperationKindUnchanged,
						GearScore: &GearScore{
							ID:             player1.GearScore().ID(),
							Level:          1,
							OperationKind_: OperationKindUpdate,
						},
					},
					player2.ID(): {
						ID:             player2.ID(),
						OperationKind_: OperationKindUnchanged,
						GuildMembers: []ElementReference{
							{
								OperationKind:        OperationKindUnchanged,
								ElementID:            int(player1.ID()),
								ElementKind:          ElementKindPlayer,
								ReferencedDataStatus: ReferencedDataModified,
							},
						},
					},
				}
			},
			func(errText string) {
				t.Errorf(errText)
			},
		)
	})
	t.Run("does not include references if nothing related to them got updated", func(t *testing.T) {
		newTreeTest(
			func(se *Engine, expectedTree *Tree) {
				player1 := se.createPlayer(false)
				item := se.createItem(false)

				item.SetBoundTo(player1.ID())

				se.UpdateState()

				item.GearScore().SetLevel(1)

				expectedTree.Item = map[ItemID]Item{
					item.ID(): {
						ID: item.ID(),
						GearScore: &GearScore{
							ID:             item.GearScore().ID(),
							Level:          1,
							OperationKind_: OperationKindUpdate,
						},
						OperationKind_: OperationKindUnchanged,
					},
				}
			},
			func(errText string) {
				t.Errorf(errText)
			},
		)
	})
	t.Run("considers downstream updated data even if reference got assigned after state update", func(t *testing.T) {
		newTreeTest(
			func(se *Engine, expectedTree *Tree) {
				player1 := se.createPlayer(false)
				player2 := se.createPlayer(false)
				item := se.createItem(false)

				se.UpdateState()
				item.SetBoundTo(player1.ID())
				player2.AddGuildMember(player1.ID())

				player1.GearScore().SetLevel(1)

				expectedTree.Item = map[ItemID]Item{
					item.ID(): {
						ID: item.ID(),
						BoundTo: &ElementReference{
							OperationKind:        OperationKindUpdate,
							ElementID:            int(player1.ID()),
							ElementKind:          ElementKindPlayer,
							ReferencedDataStatus: ReferencedDataModified,
						},
						OperationKind_: OperationKindUpdate,
					},
				}
				expectedTree.Player = map[PlayerID]Player{
					player1.ID(): {
						ID:             player1.ID(),
						OperationKind_: OperationKindUnchanged,
						GearScore: &GearScore{
							ID:             player1.GearScore().ID(),
							Level:          1,
							OperationKind_: OperationKindUpdate,
						},
					},
					player2.ID(): {
						ID: player2.ID(),
						GuildMembers: []ElementReference{
							{
								OperationKind:        OperationKindUpdate,
								ElementID:            int(player1.ID()),
								ElementKind:          ElementKindPlayer,
								ReferencedDataStatus: ReferencedDataModified,
							},
						},
						OperationKind_: OperationKindUpdate,
					},
				}
			},
			func(errText string) {
				t.Errorf(errText)
			},
		)
	})
	t.Run("has ReferencedDataUnchanged when data was not changed", func(t *testing.T) {
		newTreeTest(
			func(se *Engine, expectedTree *Tree) {
				item1 := se.CreateItem()
				item1.SetName("item1")
				player1 := se.CreatePlayer()
				item1.SetBoundTo(player1.ID())

				item2 := se.CreateItem()
				item2.SetName("item2")

				player2 := se.CreatePlayer()

				se.UpdateState()

				ref, _ := item1.BoundTo()
				ref.Unset()

				item2.SetBoundTo(player1.ID())
				player2.AddGuildMember(player1.ID())

				expectedTree.Item = map[ItemID]Item{
					item1.ID(): {
						ID:   item1.ID(),
						Name: "item1",
						BoundTo: &ElementReference{
							OperationKind:        OperationKindDelete,
							ElementID:            int(player1.ID()),
							ElementKind:          ElementKindPlayer,
							ReferencedDataStatus: ReferencedDataUnchanged,
						},
						OperationKind_: OperationKindUpdate,
					},
					item2.ID(): {
						ID:   item2.ID(),
						Name: "item2",
						BoundTo: &ElementReference{
							OperationKind:        OperationKindUpdate,
							ElementID:            int(player1.ID()),
							ElementKind:          ElementKindPlayer,
							ReferencedDataStatus: ReferencedDataUnchanged,
						},
						OperationKind_: OperationKindUpdate,
					},
				}
				expectedTree.Player = map[PlayerID]Player{
					player2.ID(): {
						ID:             player2.ID(),
						OperationKind_: OperationKindUpdate,
						GuildMembers: []ElementReference{
							{
								OperationKind:        OperationKindUpdate,
								ElementID:            int(player1.ID()),
								ElementKind:          ElementKindPlayer,
								ReferencedDataStatus: ReferencedDataUnchanged,
							},
						},
					},
				}
			},
			func(errText string) {
				t.Errorf(errText)
			},
		)
	})
}

func TestMergePlayerIDs(t *testing.T) {
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
	t.Run("", func(t *testing.T) {
		inputCurrentIDs := []PlayerID{1, 2, 3, 4}
		inputNextIDs := []PlayerID{1, 3}

		expected := []PlayerID{1, 2, 3, 4}

		actual := mergePlayerIDs(inputCurrentIDs, inputNextIDs)

		assert.Equal(t, expected, actual)
	})
}
