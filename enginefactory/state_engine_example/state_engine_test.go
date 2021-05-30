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
		se.UpdateState()
		se.deleteGearScore(gearScore.ID())
		_gearScore := se.Patch.GearScore[gearScore.ID()]
		assert.Equal(t, OperationKindDelete, _gearScore.OperationKind)
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
		se.UpdateState()
		player.RemoveItems(item.ID())
		_item := se.Patch.Item[item.ID()]
		assert.Equal(t, OperationKindDelete, _item.OperationKind)
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
	t.Run("deletes reference when Set is called on field which already has reference", func(t *testing.T) {
		se := newEngine()
		item := se.createItem(false)
		player := se.createPlayer(false)
		player2 := se.createPlayer(false)
		item.SetBoundTo(player.ID())
		item.SetBoundTo(player2.ID())
		assert.Equal(t, 1, len(se.Patch.ItemBoundToRef))
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
	t.Run("does not set attribute on element which was deleted even before entering State", func(t *testing.T) {
		se := newEngine()
		gearScore := se.CreateGearScore()
		assert.Equal(t, 0, gearScore.Level())
		se.DeleteGearScore(gearScore.ID())
		gearScore.SetLevel(1)
		assert.Equal(t, 0, gearScore.Level())
	})
	t.Run("does not set attribute on element which is set to be deleted", func(t *testing.T) {
		se := newEngine()
		gearScore := se.CreateGearScore()
		assert.Equal(t, 0, gearScore.Level())
		se.UpdateState()
		se.DeleteGearScore(gearScore.ID())
		gearScore.SetLevel(1)
		assert.Equal(t, 0, gearScore.Level())
	})
	t.Run("does not add child on element which is set to be deleted", func(t *testing.T) {
		se := newEngine()
		player := se.CreatePlayer()
		se.DeletePlayer(player.ID())
		item := player.AddItem()
		assert.Equal(t, OperationKindDelete, item.item.OperationKind)
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
	se.walkTree()
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
									ID:            player1.GearScore().ID(),
									OperationKind: OperationKindUpdate,
								},
								OperationKind: OperationKindUpdate,
								Position: &Position{
									ID:            player1.Position().ID(),
									OperationKind: OperationKindUpdate,
								},
							},
							{
								ID: player2.ID(),
								GearScore: &GearScore{
									ID:            player2.GearScore().ID(),
									OperationKind: OperationKindUpdate,
								},
								OperationKind: OperationKindUpdate,
								Position: &Position{
									ID:            player2.Position().ID(),
									OperationKind: OperationKindUpdate,
								},
							},
						},
						OperationKind: OperationKindUpdate,
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
									ID:            player1.GearScore().ID(),
									Level:         1,
									OperationKind: OperationKindUpdate,
								},
								OperationKind: OperationKindUnchanged,
							},
						},
						OperationKind: OperationKindUnchanged,
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
										ID:            player1item1.ID(),
										OperationKind: OperationKindUpdate,
										GearScore: &GearScore{
											ID:            player1item1.GearScore().ID(),
											OperationKind: OperationKindUpdate,
										},
										Origin: &Player{
											ID:            player1item1.Origin().Player().ID(),
											OperationKind: OperationKindUpdate,
											GearScore: &GearScore{
												ID:            player1item1.Origin().Player().GearScore().ID(),
												OperationKind: OperationKindUpdate,
											},
											Position: &Position{
												ID:            player1item1.Origin().Player().Position().ID(),
												OperationKind: OperationKindUpdate,
											},
										},
									},
								},
								OperationKind: OperationKindUpdate,
							},
						},
						OperationKind: OperationKindUnchanged,
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
										ID:            player1item2.ID(),
										OperationKind: OperationKindDelete,
										GearScore: &GearScore{
											ID:            player1item2.item.GearScore,
											OperationKind: OperationKindDelete,
										},
										Origin: &Player{
											ID:            player1item2.Origin().Player().ID(),
											OperationKind: OperationKindDelete,
											GearScore: &GearScore{
												ID:            player1item2.Origin().Player().GearScore().ID(),
												OperationKind: OperationKindDelete,
											},
											Position: &Position{
												ID:            player1item2.Origin().Player().Position().ID(),
												OperationKind: OperationKindDelete,
											},
										},
									},
								},
								OperationKind: OperationKindUpdate,
							},
						},
						OperationKind: OperationKindUnchanged,
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
						ID:            item.ID(),
						Name:          "myItem",
						BoundTo:       &PlayerReference{OperationKindUnchanged, player.ID(), ElementKindPlayer, ReferencedDataModified, newPath(playerIdentifier, int(player.ID())).toJSONPath(), nil},
						OperationKind: OperationKindUnchanged,
					},
				}
				expectedTree.Player = map[PlayerID]Player{
					player.ID(): {
						ID: player.ID(),
						Items: []Item{
							{
								ID:            playerItem.ID(),
								OperationKind: OperationKindUpdate,
								GearScore: &GearScore{
									ID:            playerItem.GearScore().ID(),
									OperationKind: OperationKindUpdate,
								},
								Origin: &Player{
									ID:            playerItem.Origin().Player().ID(),
									OperationKind: OperationKindUpdate,
									GearScore: &GearScore{
										ID:            playerItem.Origin().Player().GearScore().ID(),
										OperationKind: OperationKindUpdate,
									},
									Position: &Position{
										ID:            playerItem.Origin().Player().Position().ID(),
										OperationKind: OperationKindUpdate,
									},
								},
							},
						},
						OperationKind: OperationKindUpdate,
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
								ID:        item.ID(),
								BoundTo:   nil,
								GearScore: &GearScore{ID: item.GearScore().ID(), OperationKind: OperationKindUpdate},
								Origin: &Player{
									ID:            item.Origin().Player().ID(),
									OperationKind: OperationKindUpdate,
									GearScore: &GearScore{
										ID:            item.Origin().Player().GearScore().ID(),
										OperationKind: OperationKindUpdate,
									},
									Position: &Position{
										ID:            item.Origin().Player().Position().ID(),
										OperationKind: OperationKindUpdate,
									},
								},
								OperationKind: OperationKindUpdate,
							},
						},
						GuildMembers: []PlayerReference{
							{
								OperationKind:        OperationKindUpdate,
								ElementID:            player2.ID(),
								ElementKind:          ElementKindPlayer,
								ReferencedDataStatus: ReferencedDataModified,
								ElementPath:          newPath(playerIdentifier, int(player2.ID())).toJSONPath(),
								Player: &Player{
									ID:            player2.ID(),
									OperationKind: OperationKindUnchanged,
									GearScore: &GearScore{
										ID:            player2.GearScore().ID(),
										OperationKind: OperationKindUnchanged,
									},
									Position: &Position{
										ID:            player2.Position().ID(),
										OperationKind: OperationKindUnchanged,
									},
									GuildMembers: []PlayerReference{
										{
											OperationKind:        OperationKindUnchanged,
											ElementID:            player1.ID(),
											ElementKind:          ElementKindPlayer,
											ReferencedDataStatus: ReferencedDataModified,
											ElementPath:          newPath(playerIdentifier, int(player1.ID())).toJSONPath(),
										},
									},
								},
							},
						},
						OperationKind: OperationKindUpdate,
					},
					player2.ID(): {
						ID:            player2.ID(),
						OperationKind: OperationKindUnchanged,
						GuildMembers: []PlayerReference{
							{
								OperationKind:        OperationKindUnchanged,
								ElementID:            player1.ID(),
								ElementKind:          ElementKindPlayer,
								ReferencedDataStatus: ReferencedDataModified,
								ElementPath:          newPath(playerIdentifier, int(player1.ID())).toJSONPath(),
							},
						},
					},
					player3.ID(): {
						ID:            player3.ID(),
						OperationKind: OperationKindUnchanged,
						GuildMembers: []PlayerReference{
							{
								OperationKind:        OperationKindUnchanged,
								ElementID:            player1.ID(),
								ElementKind:          ElementKindPlayer,
								ReferencedDataStatus: ReferencedDataModified,
								ElementPath:          newPath(playerIdentifier, int(player1.ID())).toJSONPath(),
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
						GuildMembers: []PlayerReference{
							{
								OperationKind:        OperationKindUnchanged,
								ElementID:            player.ID(),
								ElementKind:          ElementKindPlayer,
								ReferencedDataStatus: ReferencedDataModified,
								ElementPath:          newPath(playerIdentifier, int(player.ID())).toJSONPath(),
							},
							{
								OperationKind:        OperationKindUpdate,
								ElementID:            player.ID(),
								ElementKind:          ElementKindPlayer,
								ReferencedDataStatus: ReferencedDataModified,
								ElementPath:          newPath(playerIdentifier, int(player.ID())).toJSONPath(),
								Player: &Player{
									ID:            player.ID(),
									OperationKind: OperationKindUpdate,
									GearScore: &GearScore{
										ID:            player.GearScore().ID(),
										OperationKind: OperationKindUnchanged,
									},
									Position: &Position{
										ID:            player.Position().ID(),
										OperationKind: OperationKindUnchanged,
									},
									GuildMembers: []PlayerReference{
										{
											OperationKind:        OperationKindUnchanged,
											ElementID:            player.ID(),
											ElementKind:          ElementKindPlayer,
											ReferencedDataStatus: ReferencedDataModified,
											ElementPath:          newPath(playerIdentifier, int(player.ID())).toJSONPath(),
										},
										{
											OperationKind:        OperationKindUpdate,
											ElementID:            player.ID(),
											ElementKind:          ElementKindPlayer,
											ReferencedDataStatus: ReferencedDataModified,
											ElementPath:          newPath(playerIdentifier, int(player.ID())).toJSONPath(),
										},
									},
								},
							},
						},
						OperationKind: OperationKindUpdate,
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
						BoundTo: &PlayerReference{
							OperationKind:        OperationKindUnchanged,
							ElementID:            player.ID(),
							ElementKind:          ElementKindPlayer,
							ReferencedDataStatus: ReferencedDataModified,
							ElementPath:          newPath(playerIdentifier, int(player.ID())).toJSONPath(),
						},
						OperationKind: OperationKindUpdate,
					},
				}
				expectedTree.Player = map[PlayerID]Player{
					player.ID(): {
						ID: player.ID(),
						EquipmentSets: []EquipmentSetReference{
							{
								OperationKind:        OperationKindUnchanged,
								ElementID:            equipmentSet.ID(),
								ElementKind:          ElementKindEquipmentSet,
								ReferencedDataStatus: ReferencedDataModified,
								ElementPath:          newPath(equipmentSetIdentifier, int(equipmentSet.ID())).toJSONPath(),
							},
						},
						OperationKind: OperationKindUnchanged,
					},
				}
				expectedTree.EquipmentSet = map[EquipmentSetID]EquipmentSet{
					equipmentSet.ID(): {
						ID: equipmentSet.ID(),
						Equipment: []ItemReference{
							{
								OperationKind:        OperationKindUnchanged,
								ElementID:            item.ID(),
								ElementKind:          ElementKindItem,
								ReferencedDataStatus: ReferencedDataModified,
								ElementPath:          newPath(itemIdentifier, int(item.ID())).toJSONPath(),
							},
						},
						OperationKind: OperationKindUnchanged,
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
						BoundTo: &PlayerReference{
							OperationKind:        OperationKindUnchanged,
							ElementID:            player1.ID(),
							ElementKind:          ElementKindPlayer,
							ReferencedDataStatus: ReferencedDataModified,
							ElementPath:          newPath(playerIdentifier, int(player1.ID())).toJSONPath(),
						},
						OperationKind: OperationKindUnchanged,
					},
				}
				expectedTree.Player = map[PlayerID]Player{
					player1.ID(): {
						ID:            player1.ID(),
						OperationKind: OperationKindUnchanged,
						GearScore: &GearScore{
							ID:            player1.GearScore().ID(),
							Level:         1,
							OperationKind: OperationKindUpdate,
						},
					},
					player2.ID(): {
						ID:            player2.ID(),
						OperationKind: OperationKindUnchanged,
						GuildMembers: []PlayerReference{
							{
								OperationKind:        OperationKindUnchanged,
								ElementID:            player1.ID(),
								ElementKind:          ElementKindPlayer,
								ReferencedDataStatus: ReferencedDataModified,
								ElementPath:          newPath(playerIdentifier, int(player1.ID())).toJSONPath(),
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
							ID:            item.GearScore().ID(),
							Level:         1,
							OperationKind: OperationKindUpdate,
						},
						OperationKind: OperationKindUnchanged,
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
						BoundTo: &PlayerReference{
							OperationKind:        OperationKindUpdate,
							ElementID:            player1.ID(),
							ElementKind:          ElementKindPlayer,
							ReferencedDataStatus: ReferencedDataModified,
							ElementPath:          newPath(playerIdentifier, int(player1.ID())).toJSONPath(),
							Player: &Player{
								ID:            player1.ID(),
								OperationKind: OperationKindUnchanged,
								GearScore: &GearScore{
									ID:            player1.GearScore().ID(),
									Level:         1,
									OperationKind: OperationKindUpdate,
								},
								Position: &Position{
									ID:            player1.Position().ID(),
									OperationKind: OperationKindUnchanged,
								},
							},
						},
						OperationKind: OperationKindUpdate,
					},
				}
				expectedTree.Player = map[PlayerID]Player{
					player1.ID(): {
						ID:            player1.ID(),
						OperationKind: OperationKindUnchanged,
						GearScore: &GearScore{
							ID:            player1.GearScore().ID(),
							Level:         1,
							OperationKind: OperationKindUpdate,
						},
					},
					player2.ID(): {
						ID: player2.ID(),
						GuildMembers: []PlayerReference{
							{
								OperationKind:        OperationKindUpdate,
								ElementID:            player1.ID(),
								ElementKind:          ElementKindPlayer,
								ReferencedDataStatus: ReferencedDataModified,
								ElementPath:          newPath(playerIdentifier, int(player1.ID())).toJSONPath(),
								Player: &Player{
									ID:            player1.ID(),
									OperationKind: OperationKindUnchanged,
									GearScore: &GearScore{
										ID:            player1.GearScore().ID(),
										Level:         1,
										OperationKind: OperationKindUpdate,
									},
									Position: &Position{
										ID:            player1.Position().ID(),
										OperationKind: OperationKindUnchanged,
									},
								},
							},
						},
						OperationKind: OperationKindUpdate,
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
						BoundTo: &PlayerReference{
							OperationKind:        OperationKindDelete,
							ElementID:            player1.ID(),
							ElementKind:          ElementKindPlayer,
							ReferencedDataStatus: ReferencedDataUnchanged,
							ElementPath:          newPath(playerIdentifier, int(player1.ID())).toJSONPath(),
						},
						OperationKind: OperationKindUpdate,
					},
					item2.ID(): {
						ID:   item2.ID(),
						Name: "item2",
						BoundTo: &PlayerReference{
							OperationKind:        OperationKindUpdate,
							ElementID:            player1.ID(),
							ElementKind:          ElementKindPlayer,
							ReferencedDataStatus: ReferencedDataUnchanged,
							ElementPath:          newPath(playerIdentifier, int(player1.ID())).toJSONPath(),
							Player: &Player{
								ID:            player1.ID(),
								OperationKind: OperationKindUnchanged,
								GearScore: &GearScore{
									ID:            player1.GearScore().ID(),
									OperationKind: OperationKindUnchanged,
								},
								Position: &Position{
									ID:            player1.Position().ID(),
									OperationKind: OperationKindUnchanged,
								},
							},
						},
						OperationKind: OperationKindUpdate,
					},
				}
				expectedTree.Player = map[PlayerID]Player{
					player2.ID(): {
						ID:            player2.ID(),
						OperationKind: OperationKindUpdate,
						GuildMembers: []PlayerReference{
							{
								OperationKind:        OperationKindUpdate,
								ElementID:            player1.ID(),
								ElementKind:          ElementKindPlayer,
								ReferencedDataStatus: ReferencedDataUnchanged,
								ElementPath:          newPath(playerIdentifier, int(player1.ID())).toJSONPath(),
								Player: &Player{
									ID:            player1.ID(),
									OperationKind: OperationKindUnchanged,
									GearScore: &GearScore{
										ID:            player1.GearScore().ID(),
										OperationKind: OperationKindUnchanged,
									},
									Position: &Position{
										ID:            player1.Position().ID(),
										OperationKind: OperationKindUnchanged,
									},
								},
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
	t.Run("builds entire referenced player when ref is created", func(t *testing.T) {
		newTreeTest(
			func(se *Engine, expectedTree *Tree) {
				item := se.createItem(false).SetName("myItem")
				player2 := se.createPlayer(false)
				player2.GearScore().SetLevel(2)
				player := se.createPlayer(false)
				player.Position().SetX(10)
				player.GearScore().SetLevel(8)
				player.AddGuildMember(player2.ID())

				se.UpdateState()
				item.SetBoundTo(player.ID())

				expectedTree.Item = map[ItemID]Item{
					item.ID(): {
						ID:   item.ID(),
						Name: "myItem",
						BoundTo: &PlayerReference{
							OperationKindUpdate,
							player.ID(),
							ElementKindPlayer,
							ReferencedDataUnchanged,
							newPath(playerIdentifier, int(player.ID())).toJSONPath(),
							&Player{
								ID:            player.ID(),
								OperationKind: OperationKindUnchanged,
								GearScore: &GearScore{
									ID:            player.GearScore().ID(),
									OperationKind: OperationKindUnchanged,
									Level:         8,
								},
								Position: &Position{
									ID:            player.Position().ID(),
									OperationKind: OperationKindUnchanged,
									X:             10,
								},
								GuildMembers: []PlayerReference{
									{
										ElementID:            player2.ID(),
										OperationKind:        OperationKindUnchanged,
										ElementKind:          ElementKindPlayer,
										ReferencedDataStatus: ReferencedDataUnchanged,
										ElementPath:          newPath(playerIdentifier, int(player2.ID())).toJSONPath(),
										Player:               nil,
									},
								},
							},
						},
						OperationKind: OperationKindUpdate,
					},
				}

			},
			func(errText string) {
				t.Errorf(errText)
			},
		)
	})
	t.Run("builds any kind", func(t *testing.T) {
		newTreeTest(
			func(se *Engine, expectedTree *Tree) {
				item1 := se.createItem(false)

				item2 := se.createItem(false)
				item2.Origin().SetPosition()

				expectedTree.Item = map[ItemID]Item{
					item1.ID(): {
						ID: item1.ID(),
						GearScore: &GearScore{
							ID:            item1.GearScore().ID(),
							OperationKind: OperationKindUpdate,
						},
						Origin: &Player{
							ID:            item1.Origin().Player().ID(),
							OperationKind: OperationKindUpdate,
							GearScore: &GearScore{
								ID:            item1.Origin().Player().GearScore().ID(),
								OperationKind: OperationKindUpdate,
							},
							Position: &Position{
								ID:            item1.Origin().Player().Position().ID(),
								OperationKind: OperationKindUpdate,
							},
						},
						OperationKind: OperationKindUpdate,
					},
					item2.ID(): {
						ID: item2.ID(),
						GearScore: &GearScore{
							ID:            item2.GearScore().ID(),
							OperationKind: OperationKindUpdate,
						},
						Origin: &Position{
							ID:            item2.Origin().Position().ID(),
							OperationKind: OperationKindUpdate,
						},
						OperationKind: OperationKindUpdate,
					},
				}
			},
			func(errText string) {
				t.Errorf(errText)
			},
		)
	})
	t.Run("builds []any kinds", func(t *testing.T) {
		newTreeTest(
			func(se *Engine, expectedTree *Tree) {
				zone := se.createZone()
				item := zone.AddInteractableItem()
				zoneItem := zone.AddInteractableZoneItem()

				expectedTree.Zone = map[ZoneID]Zone{
					zone.ID(): {
						ID:            zone.ID(),
						OperationKind: OperationKindUpdate,
						Interactables: []interface{}{
							Item{
								ID: item.ID(),
								GearScore: &GearScore{
									ID:            item.GearScore().ID(),
									OperationKind: OperationKindUpdate,
								},
								Origin: &Player{
									ID:            item.Origin().Player().ID(),
									OperationKind: OperationKindUpdate,
									GearScore: &GearScore{
										ID:            item.Origin().Player().GearScore().ID(),
										OperationKind: OperationKindUpdate,
									},
									Position: &Position{
										ID:            item.Origin().Player().Position().ID(),
										OperationKind: OperationKindUpdate,
									},
								},
								OperationKind: OperationKindUpdate,
							},
							ZoneItem{
								ID:            zoneItem.ID(),
								OperationKind: OperationKindUpdate,
								Position: &Position{
									ID:            zoneItem.Position().ID(),
									OperationKind: OperationKindUpdate,
								},
								Item: &Item{
									ID:            zoneItem.Item().ID(),
									OperationKind: OperationKindUpdate,
									GearScore: &GearScore{
										ID:            zoneItem.Item().GearScore().ID(),
										OperationKind: OperationKindUpdate,
									},
									Origin: &Player{
										ID:            zoneItem.Item().Origin().Player().ID(),
										OperationKind: OperationKindUpdate,
										GearScore: &GearScore{
											ID:            zoneItem.Item().Origin().Player().GearScore().ID(),
											OperationKind: OperationKindUpdate,
										},
										Position: &Position{
											ID:            zoneItem.Item().Origin().Player().Position().ID(),
											OperationKind: OperationKindUpdate,
										},
									},
								},
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
	t.Run("builds *any kind", func(t *testing.T) {
		newTreeTest(
			func(se *Engine, expectedTree *Tree) {
				player1 := se.createPlayer(false)
				player2 := se.createPlayer(false)
				player1.SetTargetPlayer(player2.ID())

				expectedTree.Player = map[PlayerID]Player{
					player1.ID(): {
						ID:            player1.ID(),
						OperationKind: OperationKindUpdate,
						GearScore: &GearScore{
							ID:            player1.GearScore().ID(),
							OperationKind: OperationKindUpdate,
						},
						Position: &Position{
							ID:            player1.Position().ID(),
							OperationKind: OperationKindUpdate,
						},
						Target: &AnyOfPlayerZoneItemReference{
							OperationKind:        OperationKindUpdate,
							ElementID:            int(player2.ID()),
							ElementKind:          ElementKindPlayer,
							ReferencedDataStatus: ReferencedDataModified,
							ElementPath:          newPath(playerIdentifier, int(player2.ID())).toJSONPath(),
							Element: &Player{
								ID:            player2.ID(),
								OperationKind: OperationKindUpdate,
								GearScore: &GearScore{
									ID:            player2.GearScore().ID(),
									OperationKind: OperationKindUpdate,
								},
								Position: &Position{
									ID:            player2.Position().ID(),
									OperationKind: OperationKindUpdate,
								},
							},
						},
					},
					player2.ID(): {
						ID:            player2.ID(),
						OperationKind: OperationKindUpdate,
						GearScore: &GearScore{
							ID:            player2.GearScore().ID(),
							OperationKind: OperationKindUpdate,
						},
						Position: &Position{
							ID:            player2.Position().ID(),
							OperationKind: OperationKindUpdate,
						},
					},
				}
			},
			func(errText string) {
				t.Errorf(errText)
			},
		)
	})
	t.Run("builds []*any kind", func(t *testing.T) {
		newTreeTest(
			func(se *Engine, expectedTree *Tree) {
				player1 := se.createPlayer(false)
				player2 := se.createPlayer(false)
				player1.AddTargetedByPlayer(player2.ID())

				expectedTree.Player = map[PlayerID]Player{
					player1.ID(): {
						ID:            player1.ID(),
						OperationKind: OperationKindUpdate,
						GearScore: &GearScore{
							ID:            player1.GearScore().ID(),
							OperationKind: OperationKindUpdate,
						},
						Position: &Position{
							ID:            player1.Position().ID(),
							OperationKind: OperationKindUpdate,
						},
						TargetedBy: []AnyOfPlayerZoneItemReference{
							{
								OperationKind:        OperationKindUpdate,
								ElementID:            int(player2.ID()),
								ElementKind:          ElementKindPlayer,
								ReferencedDataStatus: ReferencedDataModified,
								ElementPath:          newPath(playerIdentifier, int(player2.ID())).toJSONPath(),
								Element: &Player{
									ID:            player2.ID(),
									OperationKind: OperationKindUpdate,
									GearScore: &GearScore{
										ID:            player2.GearScore().ID(),
										OperationKind: OperationKindUpdate,
									},
									Position: &Position{
										ID:            player2.Position().ID(),
										OperationKind: OperationKindUpdate,
									},
								},
							},
						},
					},
					player2.ID(): {
						ID:            player2.ID(),
						OperationKind: OperationKindUpdate,
						GearScore: &GearScore{
							ID:            player2.GearScore().ID(),
							OperationKind: OperationKindUpdate,
						},
						Position: &Position{
							ID:            player2.Position().ID(),
							OperationKind: OperationKindUpdate,
						},
					},
				}
			},
			func(errText string) {
				t.Errorf(errText)
			},
		)
	})
	t.Run("build tree with path to element of any type", func(t *testing.T) {
		newTreeTest(
			func(se *Engine, expectedTree *Tree) {
				zone := se.createZone()
				player1 := zone.AddPlayer()
				player2 := se.createPlayer(false)
				player2.SetTargetPlayer(player1.ID())

				se.UpdateState()
				player1.GearScore().SetLevel(1)

				expectedTree.Player = map[PlayerID]Player{
					player2.ID(): {
						ID:            player2.ID(),
						OperationKind: OperationKindUnchanged,
						Target: &AnyOfPlayerZoneItemReference{
							OperationKind:        OperationKindUnchanged,
							ElementID:            int(player1.ID()),
							ElementKind:          ElementKindPlayer,
							ReferencedDataStatus: ReferencedDataModified,
							ElementPath:          newPath(zoneIdentifier, int(zone.ID())).players().index(0).toJSONPath(),
						},
					},
				}
				expectedTree.Zone = map[ZoneID]Zone{
					zone.ID(): {
						ID:            zone.ID(),
						OperationKind: OperationKindUnchanged,
						Players: []Player{
							{
								ID:            player1.ID(),
								OperationKind: OperationKindUnchanged,
								GearScore: &GearScore{
									ID:            player1.GearScore().ID(),
									OperationKind: OperationKindUpdate,
									Level:         1,
								},
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
	t.Run("assembles tree with correct path no non-ref any type", func(t *testing.T) {
		newTreeTest(
			func(se *Engine, expectedTree *Tree) {
				zone := se.createZone()
				player := zone.AddInteractablePlayer()
				item := se.createItem(false)
				item.SetBoundTo(player.ID())

				se.UpdateState()
				player.GearScore().SetLevel(1)

				expectedTree.Item = map[ItemID]Item{
					item.ID(): {
						ID:            item.ID(),
						OperationKind: OperationKindUnchanged,
						BoundTo: &PlayerReference{
							OperationKind:        OperationKindUnchanged,
							ElementID:            player.ID(),
							ElementKind:          ElementKindPlayer,
							ReferencedDataStatus: ReferencedDataModified,
							ElementPath:          newPath(zoneIdentifier, int(zone.ID())).interactables().index(0).toJSONPath(),
						},
					},
				}
				expectedTree.Zone = map[ZoneID]Zone{
					zone.ID(): {
						ID:            zone.ID(),
						OperationKind: OperationKindUnchanged,
						Interactables: []interface{}{
							Player{
								ID:            player.ID(),
								OperationKind: OperationKindUnchanged,
								GearScore: &GearScore{
									ID:            player.GearScore().ID(),
									OperationKind: OperationKindUpdate,
									Level:         1,
								},
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
