package state

import (
	"testing"

	"github.com/jobergner/backent-cli/pkg/testutils"

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
	t.Run("gets element and checks of they exist", func(t *testing.T) {
		se := newEngine()
		gearScore := se.CreateGearScore()
		_, gearScoreExists := se.GearScore(gearScore.ID()).Exists()
		assert.True(t, gearScoreExists)

		_, itemExists := se.Item(ItemID(1)).Exists()
		assert.False(t, itemExists)

		player := se.CreatePlayer()
		se.DeletePlayer(player.ID())
		_, playerExists := player.Exists()
		assert.False(t, playerExists)
	})
	t.Run("gets every element", func(t *testing.T) {
		se := newEngine()
		se.CreateGearScore()
		se.CreateGearScore()
		se.CreatePlayer()
		assert.Equal(t, 2, len(se.EveryGearScore()))
	})
	t.Run("gets slice of elements", func(t *testing.T) {
		se := newEngine()
		player := se.CreatePlayer()
		player.AddItem()
		player.AddItem()
		assert.Equal(t, 2, len(player.Items()))
	})
	t.Run("gets slice of elements excluding elements which have OperationKindDelete", func(t *testing.T) {
		se := newEngine()
		player := se.CreatePlayer()
		item1 := player.AddItem()
		player.AddItem()
		player.RemoveItems(item1.ID())
		assert.Equal(t, 1, len(player.Items()))
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
	t.Run("removes elements of any type", func(t *testing.T) {
		se := newEngine()
		zone := se.CreateZone()
		zone.AddInteractableItem()
		zone.AddInteractablePlayer()
		zoneItem := zone.AddInteractableZoneItem()
		zone.RemoveInteractablesZoneItem(zoneItem.ID())
		assert.Equal(t, len(zone.Interactables()), 2)
	})
	t.Run("removes elements of any type reference", func(t *testing.T) {
		se := newEngine()
		player1 := se.CreatePlayer()
		zoneItem1 := se.CreateZoneItem()
		player2 := se.CreatePlayer()
		player1.AddTargetedByZoneItem(zoneItem1.ID())
		player1.AddTargetedByPlayer(player2.ID())
		player1.RemoveTargetedByPlayer(player2.ID())
		assert.Equal(t, len(player1.TargetedBy()), 1)
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
		_, isSet := se.Item(item.ID()).BoundTo().IsSet()
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
		item := se.CreateItem()
		player := se.CreatePlayer()
		player2 := se.CreatePlayer()
		item.SetBoundTo(player.ID())
		item.SetBoundTo(player2.ID())
		assert.Equal(t, 1, len(se.Patch.ItemBoundToRef))
	})
	t.Run("does not add reference to slice when element is already referenced", func(t *testing.T) {
		se := newEngine()
		player1 := se.CreatePlayer()
		player2 := se.CreatePlayer()
		player1.AddGuildMember(player2.ID())
		assert.Equal(t, 1, len(player1.GuildMembers()))
		player1.AddGuildMember(player2.ID())
		assert.Equal(t, 1, len(player1.GuildMembers()))
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
	t.Run("does not set elements when value has not changed", func(t *testing.T) {
		se := newEngine()
		gearScore := se.CreateGearScore().SetLevel(1)
		se.UpdateState()
		gearScore.SetLevel(1)
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
	actualTree := se.assembleTree(false)

	if !assert.ObjectsAreEqualValues(expectedTree, actualTree) {
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
				zne := se.CreateZone()
				player1 := zne.AddPlayer()
				player2 := zne.AddPlayer()

				expectedTree.Zone = map[ZoneID]zone{
					zne.ID(): {
						ID: zne.ID(),
						Players: map[PlayerID]player{
							player1.ID(): {
								ID: player1.ID(),
								GearScore: &gearScore{
									ID:            player1.GearScore().ID(),
									OperationKind: OperationKindUpdate,
								},
								OperationKind: OperationKindUpdate,
								Position: &position{
									ID:            player1.Position().ID(),
									OperationKind: OperationKindUpdate,
								},
							},
							player2.ID(): {
								ID: player2.ID(),
								GearScore: &gearScore{
									ID:            player2.GearScore().ID(),
									OperationKind: OperationKindUpdate,
								},
								OperationKind: OperationKindUpdate,
								Position: &position{
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
				zne := se.CreateZone()
				player1 := zne.AddPlayer()
				_ = zne.AddPlayer()
				se.UpdateState()
				player1.GearScore().SetLevel(1)

				expectedTree.Zone = map[ZoneID]zone{
					zne.ID(): {
						ID: zne.ID(),
						Players: map[PlayerID]player{
							player1.ID(): {
								ID: player1.ID(),
								GearScore: &gearScore{
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
				zne := se.CreateZone()
				player1 := zne.AddPlayer()
				_ = zne.AddPlayer()
				se.UpdateState()
				player1item1 := player1.AddItem()

				expectedTree.Zone = map[ZoneID]zone{
					zne.ID(): {
						ID: zne.ID(),
						Players: map[PlayerID]player{
							player1.ID(): {
								ID: player1.ID(),
								Items: map[ItemID]item{
									player1item1.ID(): {
										ID:            player1item1.ID(),
										OperationKind: OperationKindUpdate,
										GearScore: &gearScore{
											ID:            player1item1.GearScore().ID(),
											OperationKind: OperationKindUpdate,
										},
										Origin: &player{
											ID:            player1item1.Origin().Player().ID(),
											OperationKind: OperationKindUpdate,
											GearScore: &gearScore{
												ID:            player1item1.Origin().Player().GearScore().ID(),
												OperationKind: OperationKindUpdate,
											},
											Position: &position{
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
				zne := se.CreateZone()
				player1 := zne.AddPlayer()
				_ = zne.AddPlayer()
				_ = player1.AddItem()
				player1item2 := player1.AddItem()

				se.UpdateState()

				player1.RemoveItems(player1item2.ID())

				expectedTree.Zone = map[ZoneID]zone{
					zne.ID(): {
						ID: zne.ID(),
						Players: map[PlayerID]player{
							player1.ID(): {
								ID: player1.ID(),
								Items: map[ItemID]item{
									player1item2.ID(): {
										ID:            player1item2.ID(),
										OperationKind: OperationKindDelete,
										GearScore: &gearScore{
											ID:            player1item2.item.GearScore,
											OperationKind: OperationKindDelete,
										},
										Origin: &player{
											ID:            player1item2.Origin().Player().ID(),
											OperationKind: OperationKindDelete,
											GearScore: &gearScore{
												ID:            player1item2.Origin().Player().GearScore().ID(),
												OperationKind: OperationKindDelete,
											},
											Position: &position{
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
				itm := se.CreateItem().SetName("myItem")
				plyr := se.CreatePlayer()
				itm.SetBoundTo(plyr.ID())

				se.UpdateState()

				playerItem := plyr.AddItem()

				expectedTree.Item = map[ItemID]item{
					itm.ID(): {
						ID:            itm.ID(),
						Name:          "myItem",
						BoundTo:       &playerReference{OperationKindUnchanged, plyr.ID(), ElementKindPlayer, ReferencedDataModified, newPath(playerIdentifier).id(int(plyr.ID())).toJSONPath(), nil},
						OperationKind: OperationKindUnchanged,
					},
				}
				expectedTree.Player = map[PlayerID]player{
					plyr.ID(): {
						ID: plyr.ID(),
						Items: map[ItemID]item{
							playerItem.ID(): {
								ID:            playerItem.ID(),
								OperationKind: OperationKindUpdate,
								GearScore: &gearScore{
									ID:            playerItem.GearScore().ID(),
									OperationKind: OperationKindUpdate,
								},
								Origin: &player{
									ID:            playerItem.Origin().Player().ID(),
									OperationKind: OperationKindUpdate,
									GearScore: &gearScore{
										ID:            playerItem.Origin().Player().GearScore().ID(),
										OperationKind: OperationKindUpdate,
									},
									Position: &position{
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
				player1 := se.CreatePlayer()
				player2 := se.CreatePlayer()
				player3 := se.CreatePlayer()

				player2.AddGuildMember(player1.ID())
				player3.AddGuildMember(player1.ID())

				se.UpdateState()

				itm := player1.AddItem()
				player1.AddGuildMember(player2.ID())

				expectedTree.Player = map[PlayerID]player{
					player1.ID(): {
						ID: player1.ID(),
						Items: map[ItemID]item{
							itm.ID(): {
								ID:        itm.ID(),
								BoundTo:   nil,
								GearScore: &gearScore{ID: itm.GearScore().ID(), OperationKind: OperationKindUpdate},
								Origin: &player{
									ID:            itm.Origin().Player().ID(),
									OperationKind: OperationKindUpdate,
									GearScore: &gearScore{
										ID:            itm.Origin().Player().GearScore().ID(),
										OperationKind: OperationKindUpdate,
									},
									Position: &position{
										ID:            itm.Origin().Player().Position().ID(),
										OperationKind: OperationKindUpdate,
									},
								},
								OperationKind: OperationKindUpdate,
							},
						},
						GuildMembers: map[PlayerID]playerReference{
							player2.ID(): {
								OperationKind:        OperationKindUpdate,
								ElementID:            player2.ID(),
								ElementKind:          ElementKindPlayer,
								ReferencedDataStatus: ReferencedDataModified,
								ElementPath:          newPath(playerIdentifier).id(int(player2.ID())).toJSONPath(),
								Player: &player{
									ID:            player2.ID(),
									OperationKind: OperationKindUnchanged,
									GearScore: &gearScore{
										ID:            player2.GearScore().ID(),
										OperationKind: OperationKindUnchanged,
									},
									Position: &position{
										ID:            player2.Position().ID(),
										OperationKind: OperationKindUnchanged,
									},
									GuildMembers: map[PlayerID]playerReference{
										player1.ID(): {
											OperationKind:        OperationKindUnchanged,
											ElementID:            player1.ID(),
											ElementKind:          ElementKindPlayer,
											ReferencedDataStatus: ReferencedDataModified,
											ElementPath:          newPath(playerIdentifier).id(int(player1.ID())).toJSONPath(),
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
						GuildMembers: map[PlayerID]playerReference{
							player1.ID(): {
								OperationKind:        OperationKindUnchanged,
								ElementID:            player1.ID(),
								ElementKind:          ElementKindPlayer,
								ReferencedDataStatus: ReferencedDataModified,
								ElementPath:          newPath(playerIdentifier).id(int(player1.ID())).toJSONPath(),
							},
						},
					},
					player3.ID(): {
						ID:            player3.ID(),
						OperationKind: OperationKindUnchanged,
						GuildMembers: map[PlayerID]playerReference{
							player1.ID(): {
								OperationKind:        OperationKindUnchanged,
								ElementID:            player1.ID(),
								ElementKind:          ElementKindPlayer,
								ReferencedDataStatus: ReferencedDataModified,
								ElementPath:          newPath(playerIdentifier).id(int(player1.ID())).toJSONPath(),
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
				plyr := se.CreatePlayer()

				se.UpdateState()

				plyr.AddGuildMember(plyr.ID())

				expectedTree.Player = map[PlayerID]player{
					plyr.ID(): {
						ID: plyr.ID(),
						GuildMembers: map[PlayerID]playerReference{
							plyr.ID(): {
								OperationKind:        OperationKindUnchanged,
								ElementID:            plyr.ID(),
								ElementKind:          ElementKindPlayer,
								ReferencedDataStatus: ReferencedDataModified,
								ElementPath:          newPath(playerIdentifier).id(int(plyr.ID())).toJSONPath(),
							},
							plyr.ID(): {
								OperationKind:        OperationKindUpdate,
								ElementID:            plyr.ID(),
								ElementKind:          ElementKindPlayer,
								ReferencedDataStatus: ReferencedDataModified,
								ElementPath:          newPath(playerIdentifier).id(int(plyr.ID())).toJSONPath(),
								Player: &player{
									ID:            plyr.ID(),
									OperationKind: OperationKindUpdate,
									GearScore: &gearScore{
										ID:            plyr.GearScore().ID(),
										OperationKind: OperationKindUnchanged,
									},
									Position: &position{
										ID:            plyr.Position().ID(),
										OperationKind: OperationKindUnchanged,
									},
									GuildMembers: map[PlayerID]playerReference{
										plyr.ID(): {
											OperationKind:        OperationKindUnchanged,
											ElementID:            plyr.ID(),
											ElementKind:          ElementKindPlayer,
											ReferencedDataStatus: ReferencedDataUnchanged, // TODO: should be modified, but won't fix for now
											ElementPath:          newPath(playerIdentifier).id(int(plyr.ID())).toJSONPath(),
										},
										plyr.ID(): {
											OperationKind:        OperationKindUpdate,
											ElementID:            plyr.ID(),
											ElementKind:          ElementKindPlayer,
											ReferencedDataStatus: ReferencedDataUnchanged, // TODO: should be modified, but won't fix for now
											ElementPath:          newPath(playerIdentifier).id(int(plyr.ID())).toJSONPath(),
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
				plyr := se.CreatePlayer()
				itm := se.CreateItem()
				eqSet := se.CreateEquipmentSet()

				plyr.AddEquipmentSet(eqSet.ID())
				eqSet.AddEquipment(itm.ID())
				itm.SetBoundTo(plyr.ID())

				se.UpdateState()

				itm.SetName("myName")

				expectedTree.Item = map[ItemID]item{
					itm.ID(): {
						ID:   itm.ID(),
						Name: "myName",
						BoundTo: &playerReference{
							OperationKind:        OperationKindUnchanged,
							ElementID:            plyr.ID(),
							ElementKind:          ElementKindPlayer,
							ReferencedDataStatus: ReferencedDataModified,
							ElementPath:          newPath(playerIdentifier).id(int(plyr.ID())).toJSONPath(),
						},
						OperationKind: OperationKindUpdate,
					},
				}
				expectedTree.Player = map[PlayerID]player{
					plyr.ID(): {
						ID: plyr.ID(),
						EquipmentSets: map[EquipmentSetID]equipmentSetReference{
							eqSet.ID(): {
								OperationKind:        OperationKindUnchanged,
								ElementID:            eqSet.ID(),
								ElementKind:          ElementKindEquipmentSet,
								ReferencedDataStatus: ReferencedDataModified,
								ElementPath:          newPath(equipmentSetIdentifier).id(int(eqSet.ID())).toJSONPath(),
							},
						},
						OperationKind: OperationKindUnchanged,
					},
				}
				expectedTree.EquipmentSet = map[EquipmentSetID]equipmentSet{
					eqSet.ID(): {
						ID: eqSet.ID(),
						Equipment: map[ItemID]itemReference{
							itm.ID(): {
								OperationKind:        OperationKindUnchanged,
								ElementID:            itm.ID(),
								ElementKind:          ElementKindItem,
								ReferencedDataStatus: ReferencedDataModified,
								ElementPath:          newPath(itemIdentifier).id(int(itm.ID())).toJSONPath(),
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
				player1 := se.CreatePlayer()
				player2 := se.CreatePlayer()
				itm := se.CreateItem()

				itm.SetBoundTo(player1.ID())
				player2.AddGuildMember(player1.ID())

				se.UpdateState()

				player1.GearScore().SetLevel(1)

				expectedTree.Item = map[ItemID]item{
					itm.ID(): {
						ID: itm.ID(),
						BoundTo: &playerReference{
							OperationKind:        OperationKindUnchanged,
							ElementID:            player1.ID(),
							ElementKind:          ElementKindPlayer,
							ReferencedDataStatus: ReferencedDataModified,
							ElementPath:          newPath(playerIdentifier).id(int(player1.ID())).toJSONPath(),
						},
						OperationKind: OperationKindUnchanged,
					},
				}
				expectedTree.Player = map[PlayerID]player{
					player1.ID(): {
						ID:            player1.ID(),
						OperationKind: OperationKindUnchanged,
						GearScore: &gearScore{
							ID:            player1.GearScore().ID(),
							Level:         1,
							OperationKind: OperationKindUpdate,
						},
					},
					player2.ID(): {
						ID:            player2.ID(),
						OperationKind: OperationKindUnchanged,
						GuildMembers: map[PlayerID]playerReference{
							player1.ID(): {
								OperationKind:        OperationKindUnchanged,
								ElementID:            player1.ID(),
								ElementKind:          ElementKindPlayer,
								ReferencedDataStatus: ReferencedDataModified,
								ElementPath:          newPath(playerIdentifier).id(int(player1.ID())).toJSONPath(),
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
				player1 := se.CreatePlayer()
				itm := se.CreateItem()

				itm.SetBoundTo(player1.ID())

				se.UpdateState()

				itm.GearScore().SetLevel(1)

				expectedTree.Item = map[ItemID]item{
					itm.ID(): {
						ID: itm.ID(),
						GearScore: &gearScore{
							ID:            itm.GearScore().ID(),
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
				player1 := se.CreatePlayer()
				player2 := se.CreatePlayer()
				itm := se.CreateItem()

				se.UpdateState()
				itm.SetBoundTo(player1.ID())
				player2.AddGuildMember(player1.ID())

				player1.GearScore().SetLevel(1)

				expectedTree.Item = map[ItemID]item{
					itm.ID(): {
						ID: itm.ID(),
						BoundTo: &playerReference{
							OperationKind:        OperationKindUpdate,
							ElementID:            player1.ID(),
							ElementKind:          ElementKindPlayer,
							ReferencedDataStatus: ReferencedDataModified,
							ElementPath:          newPath(playerIdentifier).id(int(player1.ID())).toJSONPath(),
							Player: &player{
								ID:            player1.ID(),
								OperationKind: OperationKindUnchanged,
								GearScore: &gearScore{
									ID:            player1.GearScore().ID(),
									Level:         1,
									OperationKind: OperationKindUpdate,
								},
								Position: &position{
									ID:            player1.Position().ID(),
									OperationKind: OperationKindUnchanged,
								},
							},
						},
						OperationKind: OperationKindUpdate,
					},
				}
				expectedTree.Player = map[PlayerID]player{
					player1.ID(): {
						ID:            player1.ID(),
						OperationKind: OperationKindUnchanged,
						GearScore: &gearScore{
							ID:            player1.GearScore().ID(),
							Level:         1,
							OperationKind: OperationKindUpdate,
						},
					},
					player2.ID(): {
						ID: player2.ID(),
						GuildMembers: map[PlayerID]playerReference{
							player1.ID(): {
								OperationKind:        OperationKindUpdate,
								ElementID:            player1.ID(),
								ElementKind:          ElementKindPlayer,
								ReferencedDataStatus: ReferencedDataModified,
								ElementPath:          newPath(playerIdentifier).id(int(player1.ID())).toJSONPath(),
								Player: &player{
									ID:            player1.ID(),
									OperationKind: OperationKindUnchanged,
									GearScore: &gearScore{
										ID:            player1.GearScore().ID(),
										Level:         1,
										OperationKind: OperationKindUpdate,
									},
									Position: &position{
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

				ref, _ := item1.BoundTo().IsSet()
				ref.Unset()

				item2.SetBoundTo(player1.ID())
				player2.AddGuildMember(player1.ID())

				expectedTree.Item = map[ItemID]item{
					item1.ID(): {
						ID:   item1.ID(),
						Name: "item1",
						BoundTo: &playerReference{
							OperationKind:        OperationKindDelete,
							ElementID:            player1.ID(),
							ElementKind:          ElementKindPlayer,
							ReferencedDataStatus: ReferencedDataUnchanged,
							ElementPath:          newPath(playerIdentifier).id(int(player1.ID())).toJSONPath(),
						},
						OperationKind: OperationKindUpdate,
					},
					item2.ID(): {
						ID:   item2.ID(),
						Name: "item2",
						BoundTo: &playerReference{
							OperationKind:        OperationKindUpdate,
							ElementID:            player1.ID(),
							ElementKind:          ElementKindPlayer,
							ReferencedDataStatus: ReferencedDataUnchanged,
							ElementPath:          newPath(playerIdentifier).id(int(player1.ID())).toJSONPath(),
							Player: &player{
								ID:            player1.ID(),
								OperationKind: OperationKindUnchanged,
								GearScore: &gearScore{
									ID:            player1.GearScore().ID(),
									OperationKind: OperationKindUnchanged,
								},
								Position: &position{
									ID:            player1.Position().ID(),
									OperationKind: OperationKindUnchanged,
								},
							},
						},
						OperationKind: OperationKindUpdate,
					},
				}
				expectedTree.Player = map[PlayerID]player{
					player2.ID(): {
						ID:            player2.ID(),
						OperationKind: OperationKindUpdate,
						GuildMembers: map[PlayerID]playerReference{
							player1.ID(): {
								OperationKind:        OperationKindUpdate,
								ElementID:            player1.ID(),
								ElementKind:          ElementKindPlayer,
								ReferencedDataStatus: ReferencedDataUnchanged,
								ElementPath:          newPath(playerIdentifier).id(int(player1.ID())).toJSONPath(),
								Player: &player{
									ID:            player1.ID(),
									OperationKind: OperationKindUnchanged,
									GearScore: &gearScore{
										ID:            player1.GearScore().ID(),
										OperationKind: OperationKindUnchanged,
									},
									Position: &position{
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
				itm := se.CreateItem().SetName("myItem")
				player2 := se.CreatePlayer()
				player2.GearScore().SetLevel(2)
				plyr := se.CreatePlayer()
				plyr.Position().SetX(10)
				plyr.GearScore().SetLevel(8)
				plyr.AddGuildMember(player2.ID())

				se.UpdateState()
				itm.SetBoundTo(plyr.ID())

				expectedTree.Item = map[ItemID]item{
					itm.ID(): {
						ID:   itm.ID(),
						Name: "myItem",
						BoundTo: &playerReference{
							OperationKindUpdate,
							plyr.ID(),
							ElementKindPlayer,
							ReferencedDataUnchanged,
							newPath(playerIdentifier).id(int(plyr.ID())).toJSONPath(),
							&player{
								ID:            plyr.ID(),
								OperationKind: OperationKindUnchanged,
								GearScore: &gearScore{
									ID:            plyr.GearScore().ID(),
									OperationKind: OperationKindUnchanged,
									Level:         8,
								},
								Position: &position{
									ID:            plyr.Position().ID(),
									OperationKind: OperationKindUnchanged,
									X:             10,
								},
								GuildMembers: map[PlayerID]playerReference{
									player2.ID(): {
										ElementID:            player2.ID(),
										OperationKind:        OperationKindUnchanged,
										ElementKind:          ElementKindPlayer,
										ReferencedDataStatus: ReferencedDataUnchanged,
										ElementPath:          newPath(playerIdentifier).id(int(player2.ID())).toJSONPath(),
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
				item1 := se.CreateItem()

				item2 := se.CreateItem()
				item2.Origin().SetPosition()

				expectedTree.Item = map[ItemID]item{
					item1.ID(): {
						ID: item1.ID(),
						GearScore: &gearScore{
							ID:            item1.GearScore().ID(),
							OperationKind: OperationKindUpdate,
						},
						Origin: &player{
							ID:            item1.Origin().Player().ID(),
							OperationKind: OperationKindUpdate,
							GearScore: &gearScore{
								ID:            item1.Origin().Player().GearScore().ID(),
								OperationKind: OperationKindUpdate,
							},
							Position: &position{
								ID:            item1.Origin().Player().Position().ID(),
								OperationKind: OperationKindUpdate,
							},
						},
						OperationKind: OperationKindUpdate,
					},
					item2.ID(): {
						ID: item2.ID(),
						GearScore: &gearScore{
							ID:            item2.GearScore().ID(),
							OperationKind: OperationKindUpdate,
						},
						Origin: &position{
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
				zne := se.CreateZone()
				itm := zne.AddInteractableItem()
				zi := zne.AddInteractableZoneItem()

				expectedTree.Zone = map[ZoneID]zone{
					zne.ID(): {
						ID:            zne.ID(),
						OperationKind: OperationKindUpdate,
						Interactables: map[int]interface{}{
							int(itm.ID()): item{
								ID: itm.ID(),
								GearScore: &gearScore{
									ID:            itm.GearScore().ID(),
									OperationKind: OperationKindUpdate,
								},
								Origin: &player{
									ID:            itm.Origin().Player().ID(),
									OperationKind: OperationKindUpdate,
									GearScore: &gearScore{
										ID:            itm.Origin().Player().GearScore().ID(),
										OperationKind: OperationKindUpdate,
									},
									Position: &position{
										ID:            itm.Origin().Player().Position().ID(),
										OperationKind: OperationKindUpdate,
									},
								},
								OperationKind: OperationKindUpdate,
							},
							int(zi.ID()): zoneItem{
								ID:            zi.ID(),
								OperationKind: OperationKindUpdate,
								Position: &position{
									ID:            zi.Position().ID(),
									OperationKind: OperationKindUpdate,
								},
								Item: &item{
									ID:            zi.Item().ID(),
									OperationKind: OperationKindUpdate,
									GearScore: &gearScore{
										ID:            zi.Item().GearScore().ID(),
										OperationKind: OperationKindUpdate,
									},
									Origin: &player{
										ID:            zi.Item().Origin().Player().ID(),
										OperationKind: OperationKindUpdate,
										GearScore: &gearScore{
											ID:            zi.Item().Origin().Player().GearScore().ID(),
											OperationKind: OperationKindUpdate,
										},
										Position: &position{
											ID:            zi.Item().Origin().Player().Position().ID(),
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
				player1 := se.CreatePlayer()
				player2 := se.CreatePlayer()
				player1.SetTargetPlayer(player2.ID())

				expectedTree.Player = map[PlayerID]player{
					player1.ID(): {
						ID:            player1.ID(),
						OperationKind: OperationKindUpdate,
						GearScore: &gearScore{
							ID:            player1.GearScore().ID(),
							OperationKind: OperationKindUpdate,
						},
						Position: &position{
							ID:            player1.Position().ID(),
							OperationKind: OperationKindUpdate,
						},
						Target: &anyOfPlayer_ZoneItemReference{
							OperationKind:        OperationKindUpdate,
							ElementID:            int(player2.ID()),
							ElementKind:          ElementKindPlayer,
							ReferencedDataStatus: ReferencedDataModified,
							ElementPath:          newPath(playerIdentifier).id(int(player2.ID())).toJSONPath(),
							Element: &player{
								ID:            player2.ID(),
								OperationKind: OperationKindUpdate,
								GearScore: &gearScore{
									ID:            player2.GearScore().ID(),
									OperationKind: OperationKindUpdate,
								},
								Position: &position{
									ID:            player2.Position().ID(),
									OperationKind: OperationKindUpdate,
								},
							},
						},
					},
					player2.ID(): {
						ID:            player2.ID(),
						OperationKind: OperationKindUpdate,
						GearScore: &gearScore{
							ID:            player2.GearScore().ID(),
							OperationKind: OperationKindUpdate,
						},
						Position: &position{
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
				player1 := se.CreatePlayer()
				player2 := se.CreatePlayer()
				player1.AddTargetedByPlayer(player2.ID())

				expectedTree.Player = map[PlayerID]player{
					player1.ID(): {
						ID:            player1.ID(),
						OperationKind: OperationKindUpdate,
						GearScore: &gearScore{
							ID:            player1.GearScore().ID(),
							OperationKind: OperationKindUpdate,
						},
						Position: &position{
							ID:            player1.Position().ID(),
							OperationKind: OperationKindUpdate,
						},
						TargetedBy: map[int]anyOfPlayer_ZoneItemReference{
							int(player2.ID()): {
								OperationKind:        OperationKindUpdate,
								ElementID:            int(player2.ID()),
								ElementKind:          ElementKindPlayer,
								ReferencedDataStatus: ReferencedDataModified,
								ElementPath:          newPath(playerIdentifier).id(int(player2.ID())).toJSONPath(),
								Element: &player{
									ID:            player2.ID(),
									OperationKind: OperationKindUpdate,
									GearScore: &gearScore{
										ID:            player2.GearScore().ID(),
										OperationKind: OperationKindUpdate,
									},
									Position: &position{
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
						GearScore: &gearScore{
							ID:            player2.GearScore().ID(),
							OperationKind: OperationKindUpdate,
						},
						Position: &position{
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
				zne := se.CreateZone()
				player1 := zne.AddPlayer()
				player2 := se.CreatePlayer()
				player2.SetTargetPlayer(player1.ID())

				se.UpdateState()
				player1.GearScore().SetLevel(1)

				expectedTree.Player = map[PlayerID]player{
					player2.ID(): {
						ID:            player2.ID(),
						OperationKind: OperationKindUnchanged,
						Target: &anyOfPlayer_ZoneItemReference{
							OperationKind:        OperationKindUnchanged,
							ElementID:            int(player1.ID()),
							ElementKind:          ElementKindPlayer,
							ReferencedDataStatus: ReferencedDataModified,
							ElementPath:          newPath(zoneIdentifier).id(int(zne.ID())).players().id(int(player1.ID())).toJSONPath(),
						},
					},
				}
				expectedTree.Zone = map[ZoneID]zone{
					zne.ID(): {
						ID:            zne.ID(),
						OperationKind: OperationKindUnchanged,
						Players: map[PlayerID]player{
							player1.ID(): {
								ID:            player1.ID(),
								OperationKind: OperationKindUnchanged,
								GearScore: &gearScore{
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
				zne := se.CreateZone()
				plyr := zne.AddInteractablePlayer()
				itm := se.CreateItem()
				itm.SetBoundTo(plyr.ID())

				se.UpdateState()
				plyr.GearScore().SetLevel(1)

				expectedTree.Item = map[ItemID]item{
					itm.ID(): {
						ID:            itm.ID(),
						OperationKind: OperationKindUnchanged,
						BoundTo: &playerReference{
							OperationKind:        OperationKindUnchanged,
							ElementID:            plyr.ID(),
							ElementKind:          ElementKindPlayer,
							ReferencedDataStatus: ReferencedDataModified,
							ElementPath:          newPath(zoneIdentifier).id(int(zne.ID())).interactables().id(int(plyr.ID())).toJSONPath(),
						},
					},
				}
				expectedTree.Zone = map[ZoneID]zone{
					zne.ID(): {
						ID:            zne.ID(),
						OperationKind: OperationKindUnchanged,
						Interactables: map[int]interface{}{
							int(plyr.ID()): player{
								ID:            plyr.ID(),
								OperationKind: OperationKindUnchanged,
								GearScore: &gearScore{
									ID:            plyr.GearScore().ID(),
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
	t.Run("do not delete ZoneItem when calling player.SetTargetPlayer", func(t *testing.T) {
		newTreeTest(
			func(se *Engine, expectedTree *Tree) {
				player1 := se.CreatePlayer()
				player2 := se.CreatePlayer()
				zoneItem1 := se.CreateZoneItem()
				player1.SetTargetZoneItem(zoneItem1.ID())
				se.UpdateState()
				player1.SetTargetPlayer(player2.ID())

				expectedTree.Player = map[PlayerID]player{
					player1.ID(): {
						ID:            player1.ID(),
						OperationKind: OperationKindUpdate,
						Target: &anyOfPlayer_ZoneItemReference{
							OperationKind:        OperationKindUpdate,
							ElementID:            int(player2.ID()),
							ElementKind:          ElementKindPlayer,
							ReferencedDataStatus: ReferencedDataUnchanged,
							ElementPath:          newPath(playerIdentifier).id(int(player2.ID())).toJSONPath(),
							Element: &player{
								ID:            player2.ID(),
								OperationKind: OperationKindUnchanged,
								GearScore: &gearScore{
									ID:            player2.GearScore().ID(),
									OperationKind: OperationKindUnchanged,
								},
								Position: &position{
									ID:            player2.Position().ID(),
									OperationKind: OperationKindUnchanged,
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
