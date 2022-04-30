package state

import (
	"testing"

	"github.com/jobergner/backent-cli/pkg/factory/testutils"

	"github.com/stretchr/testify/assert"
)

func TestEngine(t *testing.T) {
	t.Run("creates elements", func(t *testing.T) {
		se := NewEngine()
		gearScore := se.CreateGearScore()
		_, ok := se.Patch.GearScore[gearScore.ID()]
		assert.True(t, ok)
	})
	t.Run("gets elements", func(t *testing.T) {
		se := NewEngine()
		gearScore := se.CreateGearScore()
		_gearScore := se.GearScore(gearScore.ID())
		assert.NotZero(t, _gearScore.ID())
	})
	t.Run("gets element and checks of they exist", func(t *testing.T) {
		se := NewEngine()
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
		se := NewEngine()
		se.CreateGearScore()
		se.CreateGearScore()
		se.CreatePlayer()
		assert.Equal(t, 2, len(se.EveryGearScore()))
	})
	t.Run("gets slice of elements", func(t *testing.T) {
		se := NewEngine()
		player := se.CreatePlayer()
		player.AddItem()
		player.AddItem()
		assert.Equal(t, 2, len(player.Items()))
	})
	t.Run("gets slice of elements excluding elements which have OperationKindDelete", func(t *testing.T) {
		se := NewEngine()
		player := se.CreatePlayer()
		item1 := player.AddItem()
		player.AddItem()
		player.RemoveItems(item1.ID())
		assert.Equal(t, 1, len(player.Items()))
	})
	t.Run("sets elements", func(t *testing.T) {
		se := NewEngine()
		gearScore := se.CreateGearScore()
		gearScore.SetLevel(10)
		intValue := se.Patch.IntValue[gearScore.gearScore.Level]
		assert.Equal(t, int64(10), intValue.Value)
	})
	t.Run("deletes elements", func(t *testing.T) {
		se := NewEngine()
		gearScore := se.CreateGearScore()
		se.UpdateState()
		se.deleteGearScore(gearScore.ID())
		_gearScore := se.Patch.GearScore[gearScore.ID()]
		assert.Equal(t, OperationKindDelete, _gearScore.OperationKind)
	})
	t.Run("adds elements", func(t *testing.T) {
		se := NewEngine()
		player := se.CreatePlayer()
		item := player.AddItem()
		playerItem := player.Items()[0]
		assert.NotZero(t, playerItem.ID())
		_, ok := se.Patch.Item[item.ID()]
		assert.True(t, ok)
	})
	t.Run("removes elements", func(t *testing.T) {
		se := NewEngine()
		player := se.CreatePlayer()
		item := player.AddItem()
		se.UpdateState()
		player.RemoveItems(item.ID())
		_item := se.Patch.Item[item.ID()]
		assert.Equal(t, OperationKindDelete, _item.OperationKind)
	})
	t.Run("removes elements of any type", func(t *testing.T) {
		se := NewEngine()
		zone := se.CreateZone()
		zone.AddInteractableItem()
		zone.AddInteractablePlayer()
		zoneItem := zone.AddInteractableZoneItem()
		zone.RemoveInteractablesZoneItem(zoneItem.ID())
		assert.Equal(t, len(zone.Interactables()), 2)
	})
	t.Run("removes elements of any type reference", func(t *testing.T) {
		se := NewEngine()
		player1 := se.CreatePlayer()
		zoneItem1 := se.CreateZoneItem()
		player2 := se.CreatePlayer()
		player1.AddTargetedByZoneItem(zoneItem1.ID())
		player1.AddTargetedByPlayer(player2.ID())
		player1.RemoveTargetedByPlayer(player2.ID())
		assert.Equal(t, len(player1.TargetedBy()), 1)
	})
	t.Run("removes events on update state", func(t *testing.T) {
		se := NewEngine()
		player := se.CreatePlayer()
		player2 := se.CreatePlayer()
		player.AddAction().SetTarget(player2.ID())
		assert.Equal(t, 0, len(se.State.AttackEventTargetRef))
		assert.Equal(t, 1, len(se.Patch.AttackEventTargetRef))
		se.UpdateState()
		assert.Equal(t, 0, len(se.State.AttackEventTargetRef))
		assert.Equal(t, 0, len(se.Patch.AttackEventTargetRef))
	})
	t.Run("creates origin, sets it from player to position", func(t *testing.T) {
		se := NewEngine()
		item1 := se.CreateItem()
		assert.Equal(t, ElementKindPlayer, item1.Origin().anyOfPlayer_Position.ElementKind)
		item1.Origin().Player().Position().SetX(1)
		assert.Equal(t, float64(1), item1.Origin().Player().Position().X())
		item2 := se.CreateItem()
		assert.Equal(t, 2, len(se.Patch.AnyOfPlayer_Position))
		assert.Equal(t, 2, len(se.Patch.Player))
		item2.Origin().BePosition()
		item2.Origin().Position().SetY(2)
		assert.Equal(t, 1, len(se.Patch.Player))
		assert.Equal(t, 2, len(se.Patch.Position)) //position of item1.origin(player).position and this one
		assert.Equal(t, float64(2), item2.Origin().Position().Y())
	})
	t.Run("signs elements as exoected", func(t *testing.T) {
		se := NewEngine()
		item := se.CreateItem()
		se.BroadcastingClientID = "foo"
		item.SetName("newname0")
		assert.Equal(t, "foo", se.Patch.StringValue[se.Item(item.ID()).item.Name].Meta.BroadcastedBy)
		assert.Equal(t, false, se.Patch.StringValue[se.Item(item.ID()).item.Name].Meta.TouchedByMany)
		item.SetName("newname2")
		assert.Equal(t, "foo", se.Patch.StringValue[se.Item(item.ID()).item.Name].Meta.BroadcastedBy)
		assert.Equal(t, false, se.Patch.StringValue[se.Item(item.ID()).item.Name].Meta.TouchedByMany)
		se.BroadcastingClientID = "bar"
		item.SetName("newname3")
		assert.Equal(t, "", se.Patch.StringValue[se.Item(item.ID()).item.Name].Meta.BroadcastedBy)
		assert.Equal(t, true, se.Patch.StringValue[se.Item(item.ID()).item.Name].Meta.TouchedByMany)
		se.UpdateState()
		assert.Equal(t, "", se.State.StringValue[se.Item(item.ID()).item.Name].Meta.BroadcastedBy)
		assert.Equal(t, false, se.State.StringValue[se.Item(item.ID()).item.Name].Meta.TouchedByMany)
	})
	t.Run("does not panic when getting parent of element without parent", func(t *testing.T) {
		se := NewEngine()
		itm := se.CreateItem()
		itm.ParentPlayer()
	})
}

func TestReferences(t *testing.T) {
	t.Run("deletes reference off element if referenced element gets deleted (1/3)", func(t *testing.T) {
		se := NewEngine()
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
		se := NewEngine()
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
		se := NewEngine()
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
		se := NewEngine()
		item := se.CreateItem()
		player := se.CreatePlayer()
		player2 := se.CreatePlayer()
		item.SetBoundTo(player.ID())
		item.SetBoundTo(player2.ID())
		assert.Equal(t, 1, len(se.Patch.ItemBoundToRef))
	})
	t.Run("does not add reference to slice when element is already referenced", func(t *testing.T) {
		se := NewEngine()
		player1 := se.CreatePlayer()
		player2 := se.CreatePlayer()
		player1.AddGuildMember(player2.ID())
		assert.Equal(t, 1, len(player1.GuildMembers()))
		player1.AddGuildMember(player2.ID())
		assert.Equal(t, 1, len(player1.GuildMembers()))
	})
	t.Run("does not add reference to slice when element is already referenced", func(t *testing.T) {
		se := NewEngine()
		player1 := se.CreatePlayer()
		player1.Position().SetX(2)
		se.CreatePlayer()
		players := se.QueryPlayers(func(p Player) bool { return p.Position().X() == 2 })
		assert.Equal(t, player1.ID(), players[0].ID())
	})
}

func TestUpdateState(t *testing.T) {
	t.Run("clears patch", func(t *testing.T) {
		se := NewEngine()
		se.CreateGearScore()
		se.UpdateState()
		assert.Equal(t, len(se.Patch.GearScore), 0)
	})
	t.Run("creates elements", func(t *testing.T) {
		se := NewEngine()
		gearScore := se.CreateGearScore()
		se.UpdateState()
		_, ok := se.Patch.GearScore[gearScore.ID()]
		assert.False(t, ok)
		_, ok = se.State.GearScore[gearScore.ID()]
		assert.True(t, ok)
	})
	t.Run("gets elements", func(t *testing.T) {
		se := NewEngine()
		gearScore := se.CreateGearScore()
		se.UpdateState()
		assert.Zero(t, gearScore.Level())
		gearScore.SetLevel(1)
		assert.Equal(t, gearScore.Level(), int64(1))
	})
	t.Run("sets elements", func(t *testing.T) {
		se := NewEngine()
		gearScore := se.CreateGearScore().SetLevel(1)
		se.UpdateState()
		_gearScore := se.State.GearScore[gearScore.ID()]
		assert.Equal(t, int64(1), se.State.IntValue[_gearScore.Level].Value)
		_, ok := se.Patch.IntValue[gearScore.gearScore.Level]
		assert.False(t, ok)
	})
	t.Run("does not set elements when value has not changed", func(t *testing.T) {
		se := NewEngine()
		gearScore := se.CreateGearScore().SetLevel(1)
		se.UpdateState()
		gearScore.SetLevel(1)
		_, ok := se.Patch.GearScore[gearScore.ID()]
		assert.False(t, ok)
	})
	t.Run("deletes elements", func(t *testing.T) {
		se := NewEngine()
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
		se := NewEngine()
		player := se.CreatePlayer()
		item := player.AddItem()
		se.UpdateState()
		_, ok := se.State.Item[item.ID()]
		assert.True(t, ok)
		_player := se.State.Player[player.ID()]
		assert.NotZero(t, len(_player.Items))
	})
	t.Run("removes elements", func(t *testing.T) {
		se := NewEngine()
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
		se := NewEngine()
		gearScore := se.CreateGearScore()
		assert.Equal(t, int64(0), gearScore.Level())
		se.DeleteGearScore(gearScore.ID())
		gearScore.SetLevel(1)
		assert.Equal(t, int64(0), gearScore.Level())
	})
	t.Run("does not set attribute on element which is set to be deleted", func(t *testing.T) {
		se := NewEngine()
		gearScore := se.CreateGearScore()
		assert.Equal(t, int64(0), gearScore.Level())
		se.UpdateState()
		se.DeleteGearScore(gearScore.ID())
		gearScore.SetLevel(1)
		assert.Equal(t, int64(0), gearScore.Level())
	})
	t.Run("does not add child on element which is set to be deleted", func(t *testing.T) {
		se := NewEngine()
		player := se.CreatePlayer()
		se.DeletePlayer(player.ID())
		item := player.AddItem()
		assert.Equal(t, OperationKindDelete, item.item.OperationKind)
		assert.Equal(t, 0, len(se.Player(player.ID()).Items()))
	})
	t.Run("does not remove child on element which is set to be deleted", func(t *testing.T) {
		se := NewEngine()
		player := se.CreatePlayer()
		item := player.AddItem()
		se.UpdateState()
		assert.Equal(t, 1, len(se.Player(player.ID()).Items()))
		se.DeletePlayer(player.ID())
		player.RemoveItems(item.ID())
		assert.Equal(t, 1, len(se.Player(player.ID()).Items()))
	})
	t.Run("does not delete element which is a child of another element", func(t *testing.T) {
		se := NewEngine()
		player := se.CreatePlayer()
		item := player.AddItem()
		se.DeleteItem(item.ID())
		assert.Equal(t, 1, len(se.Player(player.ID()).Items()))
	})
}

func newTreeTest(define func(*Engine, *Tree), onFail func(errText string), assembleEntireTree bool) {
	se := NewEngine()
	expectedTree := newTree()
	define(se, expectedTree)
	if assembleEntireTree {
		se.AssembleFullTree()
	} else {
		se.AssembleUpdateTree()
	}
	actualTree := se.Tree

	if !assert.ObjectsAreEqualValues(expectedTree, actualTree) {
		actual, _ := actualTree.MarshalJSON()
		expected, _ := expectedTree.MarshalJSON()
		actualString := string(actual)
		expectedString := string(expected)
		onFail(testutils.DiffJSON(actualString, expectedString))
	}
}

func TestTree(t *testing.T) {

	stringPtr := func(s string) *string {
		return &s
	}

	intPtr := func(i int64) *int64 {
		return &i
	}

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
									Level:         new(int64),
									Score:         new(int64),
									OperationKind: OperationKindUpdate,
								},
								OperationKind: OperationKindUpdate,
								Position: &position{
									ID:            player1.Position().ID(),
									X:             new(float64),
									Y:             new(float64),
									OperationKind: OperationKindUpdate,
								},
							},
							player2.ID(): {
								ID: player2.ID(),
								GearScore: &gearScore{
									ID:            player2.GearScore().ID(),
									Level:         new(int64),
									Score:         new(int64),
									OperationKind: OperationKindUpdate,
								},
								OperationKind: OperationKindUpdate,
								Position: &position{
									ID:            player2.Position().ID(),
									X:             new(float64),
									Y:             new(float64),
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
			false,
		)
	})
	t.Run("assembles basic elements in a tree", func(t *testing.T) {
		newTreeTest(
			func(se *Engine, expectedTree *Tree) {
				zne := se.CreateZone()
				// TODO fix how slices of basic types work, or what was i thinking when implementing it
				// maybe do map[type]OperationKind (wont't work if slice contains same string multiple times)
				zne.AddTag("foo")
				zne.AddTag("bar")
				itm := zne.AddItem()

				se.UpdateState()
				zne.RemoveTags("bar")
				zne.AddTag("baz")
				itm.Item().SetName("newName")

				expectedTree.Zone = map[ZoneID]zone{
					zne.ID(): {
						ID:            zne.ID(),
						OperationKind: OperationKindUpdate,
						Tags:          []string{"baz"},
						Items: map[ZoneItemID]zoneItem{
							itm.ID(): {
								ID:            itm.ID(),
								OperationKind: OperationKindUnchanged,
								Item: &item{
									ID:            itm.Item().ID(),
									OperationKind: OperationKindUpdate,
									Name:          stringPtr("newName"),
								},
							},
						},
					},
				}
			},
			func(errText string) {
				t.Errorf(errText)
			},
			false,
		)
	})
	t.Run("assembles correctly when using be<Type> method after state update", func(t *testing.T) {
		newTreeTest(
			func(se *Engine, expectedTree *Tree) {
				itm := se.CreateItem()
				se.UpdateState()

				itm.Origin().BePosition()

				expectedTree.Item = map[ItemID]item{
					itm.ID(): {
						ID:            itm.ID(),
						OperationKind: OperationKindUnchanged,
						Origin: &position{
							ID:            itm.Origin().Position().ID(),
							X:             new(float64),
							Y:             new(float64),
							OperationKind: OperationKindUpdate,
						},
					},
				}
			},
			func(errText string) {
				t.Errorf(errText)
			},
			false,
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
									Level:         intPtr(1),
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
			false,
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
										Name:          new(string),
										OperationKind: OperationKindUpdate,
										GearScore: &gearScore{
											ID:            player1item1.GearScore().ID(),
											Level:         new(int64),
											Score:         new(int64),
											OperationKind: OperationKindUpdate,
										},
										Origin: &player{
											ID:            player1item1.Origin().Player().ID(),
											OperationKind: OperationKindUpdate,
											GearScore: &gearScore{
												ID:            player1item1.Origin().Player().GearScore().ID(),
												Level:         new(int64),
												Score:         new(int64),
												OperationKind: OperationKindUpdate,
											},
											Position: &position{
												ID:            player1item1.Origin().Player().Position().ID(),
												X:             new(float64),
												Y:             new(float64),
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
			false,
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
										Name:          new(string),
										OperationKind: OperationKindDelete,
										GearScore: &gearScore{
											ID:            player1item2.item.GearScore,
											Level:         new(int64),
											Score:         new(int64),
											OperationKind: OperationKindDelete,
										},
										Origin: &player{
											ID:            player1item2.Origin().Player().ID(),
											OperationKind: OperationKindDelete,
											GearScore: &gearScore{
												ID:            player1item2.Origin().Player().GearScore().ID(),
												Level:         new(int64),
												Score:         new(int64),
												OperationKind: OperationKindDelete,
											},
											Position: &position{
												ID:            player1item2.Origin().Player().Position().ID(),
												X:             new(float64),
												Y:             new(float64),
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
			false,
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
						ID: itm.ID(),
						BoundTo: &elementReference{
							OperationKindUnchanged,
							int(plyr.ID()),
							ElementKindPlayer,
							ReferencedDataModified,
							newPath().extendAndCopy(playerIdentifier, int(plyr.ID()), ElementKindPlayer, ComplexID(itm.item.BoundTo)).toJSONPath()},
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
								Name:          new(string),
								GearScore: &gearScore{
									ID:            playerItem.GearScore().ID(),
									Level:         new(int64),
									Score:         new(int64),
									OperationKind: OperationKindUpdate,
								},
								Origin: &player{
									ID:            playerItem.Origin().Player().ID(),
									OperationKind: OperationKindUpdate,
									GearScore: &gearScore{
										ID:            playerItem.Origin().Player().GearScore().ID(),
										Level:         new(int64),
										Score:         new(int64),
										OperationKind: OperationKindUpdate,
									},
									Position: &position{
										ID:            playerItem.Origin().Player().Position().ID(),
										X:             new(float64),
										Y:             new(float64),
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
			false,
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
								ID:      itm.ID(),
								Name:    new(string),
								BoundTo: nil,
								GearScore: &gearScore{
									ID:            itm.GearScore().ID(),
									Level:         new(int64),
									Score:         new(int64),
									OperationKind: OperationKindUpdate,
								},
								Origin: &player{
									ID:            itm.Origin().Player().ID(),
									OperationKind: OperationKindUpdate,
									GearScore: &gearScore{
										ID:            itm.Origin().Player().GearScore().ID(),
										Level:         new(int64),
										Score:         new(int64),
										OperationKind: OperationKindUpdate,
									},
									Position: &position{
										ID:            itm.Origin().Player().Position().ID(),
										X:             new(float64),
										Y:             new(float64),
										OperationKind: OperationKindUpdate,
									},
								},
								OperationKind: OperationKindUpdate,
							},
						},
						GuildMembers: map[PlayerID]elementReference{
							player2.ID(): {
								OperationKind:        OperationKindUpdate,
								ElementID:            int(player2.ID()),
								ElementKind:          ElementKindPlayer,
								ReferencedDataStatus: ReferencedDataModified,
								ElementPath:          newPath().extendAndCopy(playerIdentifier, int(player2.ID()), ElementKindPlayer, ComplexID{}).toJSONPath(),
							},
						},
						OperationKind: OperationKindUpdate,
					},
					player2.ID(): {
						ID:            player2.ID(),
						OperationKind: OperationKindUnchanged,
						GuildMembers: map[PlayerID]elementReference{
							player1.ID(): {
								OperationKind:        OperationKindUnchanged,
								ElementID:            int(player1.ID()),
								ElementKind:          ElementKindPlayer,
								ReferencedDataStatus: ReferencedDataModified,
								ElementPath:          newPath().extendAndCopy(playerIdentifier, int(player1.ID()), ElementKindPlayer, ComplexID{}).toJSONPath(),
							},
						},
					},
					player3.ID(): {
						ID:            player3.ID(),
						OperationKind: OperationKindUnchanged,
						GuildMembers: map[PlayerID]elementReference{
							player1.ID(): {
								OperationKind:        OperationKindUnchanged,
								ElementID:            int(player1.ID()),
								ElementKind:          ElementKindPlayer,
								ReferencedDataStatus: ReferencedDataModified,
								ElementPath:          newPath().extendAndCopy(playerIdentifier, int(player1.ID()), ElementKindPlayer, ComplexID{}).toJSONPath(),
							},
						},
					},
				}

			},
			func(errText string) {
				t.Errorf(errText)
			},
			false,
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
						GuildMembers: map[PlayerID]elementReference{
							plyr.ID(): {
								OperationKind:        OperationKindUpdate,
								ElementID:            int(plyr.ID()),
								ElementKind:          ElementKindPlayer,
								ReferencedDataStatus: ReferencedDataModified,
								ElementPath:          newPath().extendAndCopy(playerIdentifier, int(plyr.ID()), ElementKindPlayer, ComplexID{}).toJSONPath(),
							},
						},
						OperationKind: OperationKindUpdate,
					},
				}

			},
			func(errText string) {
				t.Errorf(errText)
			},
			false,
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
						Name: stringPtr("myName"),
						BoundTo: &elementReference{
							OperationKind:        OperationKindUnchanged,
							ElementID:            int(plyr.ID()),
							ElementKind:          ElementKindPlayer,
							ReferencedDataStatus: ReferencedDataModified,
							ElementPath:          newPath().extendAndCopy(playerIdentifier, int(plyr.ID()), ElementKindPlayer, ComplexID{}).toJSONPath(),
						},
						OperationKind: OperationKindUpdate,
					},
				}
				expectedTree.Player = map[PlayerID]player{
					plyr.ID(): {
						ID: plyr.ID(),
						EquipmentSets: map[EquipmentSetID]elementReference{
							eqSet.ID(): {
								OperationKind:        OperationKindUnchanged,
								ElementID:            int(eqSet.ID()),
								ElementKind:          ElementKindEquipmentSet,
								ReferencedDataStatus: ReferencedDataModified,
								ElementPath:          newPath().extendAndCopy(equipmentSetIdentifier, int(eqSet.ID()), ElementKindEquipmentSet, ComplexID{}).toJSONPath(),
							},
						},
						OperationKind: OperationKindUnchanged,
					},
				}
				expectedTree.EquipmentSet = map[EquipmentSetID]equipmentSet{
					eqSet.ID(): {
						ID: eqSet.ID(),
						Equipment: map[ItemID]elementReference{
							itm.ID(): {
								OperationKind:        OperationKindUnchanged,
								ElementID:            int(itm.ID()),
								ElementKind:          ElementKindItem,
								ReferencedDataStatus: ReferencedDataModified,
								ElementPath:          newPath().extendAndCopy(itemIdentifier, int(itm.ID()), ElementKindItem, ComplexID{}).toJSONPath(),
							},
						},
						OperationKind: OperationKindUnchanged,
					},
				}

			},
			func(errText string) {
				t.Errorf(errText)
			},
			false,
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
						BoundTo: &elementReference{
							OperationKind:        OperationKindUnchanged,
							ElementID:            int(player1.ID()),
							ElementKind:          ElementKindPlayer,
							ReferencedDataStatus: ReferencedDataModified,
							ElementPath:          newPath().extendAndCopy(playerIdentifier, int(player1.ID()), ElementKindPlayer, ComplexID{}).toJSONPath(),
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
							Level:         intPtr(1),
							OperationKind: OperationKindUpdate,
						},
					},
					player2.ID(): {
						ID:            player2.ID(),
						OperationKind: OperationKindUnchanged,
						GuildMembers: map[PlayerID]elementReference{
							player1.ID(): {
								OperationKind:        OperationKindUnchanged,
								ElementID:            int(player1.ID()),
								ElementKind:          ElementKindPlayer,
								ReferencedDataStatus: ReferencedDataModified,
								ElementPath:          newPath().extendAndCopy(playerIdentifier, int(player1.ID()), ElementKindPlayer, ComplexID{}).toJSONPath(),
							},
						},
					},
				}
			},
			func(errText string) {
				t.Errorf(errText)
			},
			false,
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
							Level:         intPtr(1),
							OperationKind: OperationKindUpdate,
						},
						OperationKind: OperationKindUnchanged,
					},
				}
			},
			func(errText string) {
				t.Errorf(errText)
			},
			false,
		)
	})
	t.Run("updates reference in element if referenced element gets deleted", func(t *testing.T) {
		newTreeTest(
			func(se *Engine, expectedTree *Tree) {
				player1 := se.CreatePlayer()
				itm := se.CreateItem()

				itm.SetBoundTo(player1.ID())

				se.UpdateState()

				se.DeletePlayer(player1.ID())

				expectedTree.Player = map[PlayerID]player{
					player1.ID(): {
						ID: player1.ID(),
						GearScore: &gearScore{
							ID:            player1.GearScore().ID(),
							Level:         new(int64),
							Score:         new(int64),
							OperationKind: OperationKindDelete,
						},
						Position: &position{
							ID:            player1.Position().ID(),
							X:             new(float64),
							Y:             new(float64),
							OperationKind: OperationKindDelete,
						},
						OperationKind: OperationKindDelete,
					},
				}
				expectedTree.Item = map[ItemID]item{
					itm.ID(): {
						ID:            itm.ID(),
						OperationKind: OperationKindUpdate,
						BoundTo: &elementReference{
							OperationKind:        OperationKindDelete,
							ElementID:            int(player1.ID()),
							ElementKind:          ElementKindPlayer,
							ReferencedDataStatus: ReferencedDataModified,
							ElementPath:          newPath().extendAndCopy(playerIdentifier, int(player1.ID()), ElementKindPlayer, ComplexID{}).toJSONPath(),
						},
					},
				}
			},
			func(errText string) {
				t.Errorf(errText)
			},
			false,
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
						BoundTo: &elementReference{
							OperationKind:        OperationKindUpdate,
							ElementID:            int(player1.ID()),
							ElementKind:          ElementKindPlayer,
							ReferencedDataStatus: ReferencedDataModified,
							ElementPath:          newPath().extendAndCopy(playerIdentifier, int(player1.ID()), ElementKindPlayer, ComplexID{}).toJSONPath(),
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
							Level:         intPtr(1),
							OperationKind: OperationKindUpdate,
						},
					},
					player2.ID(): {
						ID: player2.ID(),
						GuildMembers: map[PlayerID]elementReference{
							player1.ID(): {
								OperationKind:        OperationKindUpdate,
								ElementID:            int(player1.ID()),
								ElementKind:          ElementKindPlayer,
								ReferencedDataStatus: ReferencedDataModified,
								ElementPath:          newPath().extendAndCopy(playerIdentifier, int(player1.ID()), ElementKindPlayer, ComplexID{}).toJSONPath(),
							},
						},
						OperationKind: OperationKindUpdate,
					},
				}
			},
			func(errText string) {
				t.Errorf(errText)
			},
			false,
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
						ID: item1.ID(),
						BoundTo: &elementReference{
							OperationKind:        OperationKindDelete,
							ElementID:            int(player1.ID()),
							ElementKind:          ElementKindPlayer,
							ReferencedDataStatus: ReferencedDataUnchanged,
							ElementPath:          newPath().extendAndCopy(playerIdentifier, int(player1.ID()), ElementKindPlayer, ComplexID{}).toJSONPath(),
						},
						OperationKind: OperationKindUpdate,
					},
					item2.ID(): {
						ID: item2.ID(),
						BoundTo: &elementReference{
							OperationKind:        OperationKindUpdate,
							ElementID:            int(player1.ID()),
							ElementKind:          ElementKindPlayer,
							ReferencedDataStatus: ReferencedDataUnchanged,
							ElementPath:          newPath().extendAndCopy(playerIdentifier, int(player1.ID()), ElementKindPlayer, ComplexID{}).toJSONPath(),
						},
						OperationKind: OperationKindUpdate,
					},
				}
				expectedTree.Player = map[PlayerID]player{
					player2.ID(): {
						ID:            player2.ID(),
						OperationKind: OperationKindUpdate,
						GuildMembers: map[PlayerID]elementReference{
							player1.ID(): {
								OperationKind:        OperationKindUpdate,
								ElementID:            int(player1.ID()),
								ElementKind:          ElementKindPlayer,
								ReferencedDataStatus: ReferencedDataUnchanged,
								ElementPath:          newPath().extendAndCopy(playerIdentifier, int(player1.ID()), ElementKindPlayer, ComplexID{}).toJSONPath(),
							},
						},
					},
				}
			},
			func(errText string) {
				t.Errorf(errText)
			},
			false,
		)
	})
	t.Run("builds any kind", func(t *testing.T) {
		newTreeTest(
			func(se *Engine, expectedTree *Tree) {
				item1 := se.CreateItem()

				item2 := se.CreateItem()
				item2.Origin().BePosition()

				expectedTree.Item = map[ItemID]item{
					item1.ID(): {
						ID:   item1.ID(),
						Name: new(string),
						GearScore: &gearScore{
							ID:            item1.GearScore().ID(),
							Level:         new(int64),
							Score:         new(int64),
							OperationKind: OperationKindUpdate,
						},
						Origin: &player{
							ID:            item1.Origin().Player().ID(),
							OperationKind: OperationKindUpdate,
							GearScore: &gearScore{
								ID:            item1.Origin().Player().GearScore().ID(),
								Level:         new(int64),
								Score:         new(int64),
								OperationKind: OperationKindUpdate,
							},
							Position: &position{
								ID:            item1.Origin().Player().Position().ID(),
								X:             new(float64),
								Y:             new(float64),
								OperationKind: OperationKindUpdate,
							},
						},
						OperationKind: OperationKindUpdate,
					},
					item2.ID(): {
						ID:   item2.ID(),
						Name: new(string),
						GearScore: &gearScore{
							Level:         new(int64),
							Score:         new(int64),
							ID:            item2.GearScore().ID(),
							OperationKind: OperationKindUpdate,
						},
						Origin: &position{
							ID:            item2.Origin().Position().ID(),
							X:             new(float64),
							Y:             new(float64),
							OperationKind: OperationKindUpdate,
						},
						OperationKind: OperationKindUpdate,
					},
				}
			},
			func(errText string) {
				t.Errorf(errText)
			},
			false,
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
								ID:   itm.ID(),
								Name: new(string),
								GearScore: &gearScore{
									ID:            itm.GearScore().ID(),
									Level:         new(int64),
									Score:         new(int64),
									OperationKind: OperationKindUpdate,
								},
								Origin: &player{
									ID:            itm.Origin().Player().ID(),
									OperationKind: OperationKindUpdate,
									GearScore: &gearScore{
										ID:            itm.Origin().Player().GearScore().ID(),
										Level:         new(int64),
										Score:         new(int64),
										OperationKind: OperationKindUpdate,
									},
									Position: &position{
										ID:            itm.Origin().Player().Position().ID(),
										X:             new(float64),
										Y:             new(float64),
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
									X:             new(float64),
									Y:             new(float64),
									OperationKind: OperationKindUpdate,
								},
								Item: &item{
									ID:            zi.Item().ID(),
									Name:          new(string),
									OperationKind: OperationKindUpdate,
									GearScore: &gearScore{
										ID:            zi.Item().GearScore().ID(),
										Level:         new(int64),
										Score:         new(int64),
										OperationKind: OperationKindUpdate,
									},
									Origin: &player{
										ID:            zi.Item().Origin().Player().ID(),
										OperationKind: OperationKindUpdate,
										GearScore: &gearScore{
											ID:            zi.Item().Origin().Player().GearScore().ID(),
											Level:         new(int64),
											Score:         new(int64),
											OperationKind: OperationKindUpdate,
										},
										Position: &position{
											ID:            zi.Item().Origin().Player().Position().ID(),
											X:             new(float64),
											Y:             new(float64),
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
			false,
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
							Level:         new(int64),
							Score:         new(int64),
							OperationKind: OperationKindUpdate,
						},
						Position: &position{
							ID:            player1.Position().ID(),
							X:             new(float64),
							Y:             new(float64),
							OperationKind: OperationKindUpdate,
						},
						Target: &elementReference{
							OperationKind:        OperationKindUpdate,
							ElementID:            int(player2.ID()),
							ElementKind:          ElementKindPlayer,
							ReferencedDataStatus: ReferencedDataModified,
							ElementPath:          newPath().extendAndCopy(playerIdentifier, int(player2.ID()), ElementKindPlayer, ComplexID{}).toJSONPath(),
						},
					},
					player2.ID(): {
						ID:            player2.ID(),
						OperationKind: OperationKindUpdate,
						GearScore: &gearScore{
							ID:            player2.GearScore().ID(),
							Level:         new(int64),
							Score:         new(int64),
							OperationKind: OperationKindUpdate,
						},
						Position: &position{
							ID:            player2.Position().ID(),
							X:             new(float64),
							Y:             new(float64),
							OperationKind: OperationKindUpdate,
						},
					},
				}
			},
			func(errText string) {
				t.Errorf(errText)
			},
			false,
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
							Level:         new(int64),
							Score:         new(int64),
							OperationKind: OperationKindUpdate,
						},
						Position: &position{
							ID:            player1.Position().ID(),
							X:             new(float64),
							Y:             new(float64),
							OperationKind: OperationKindUpdate,
						},
						TargetedBy: map[int]elementReference{
							int(player2.ID()): {
								OperationKind:        OperationKindUpdate,
								ElementID:            int(player2.ID()),
								ElementKind:          ElementKindPlayer,
								ReferencedDataStatus: ReferencedDataModified,
								ElementPath:          newPath().extendAndCopy(playerIdentifier, int(player2.ID()), ElementKindPlayer, ComplexID{}).toJSONPath(),
							},
						},
					},
					player2.ID(): {
						ID:            player2.ID(),
						OperationKind: OperationKindUpdate,
						GearScore: &gearScore{
							ID:            player2.GearScore().ID(),
							Level:         new(int64),
							Score:         new(int64),
							OperationKind: OperationKindUpdate,
						},
						Position: &position{
							ID:            player2.Position().ID(),
							X:             new(float64),
							Y:             new(float64),
							OperationKind: OperationKindUpdate,
						},
					},
				}
			},
			func(errText string) {
				t.Errorf(errText)
			},
			false,
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
						Target: &elementReference{
							OperationKind:        OperationKindUnchanged,
							ElementID:            int(player1.ID()),
							ElementKind:          ElementKindPlayer,
							ReferencedDataStatus: ReferencedDataModified,
							ElementPath:          newPath().extendAndCopy(zoneIdentifier, int(zne.ID()), ElementKindZone, ComplexID{}).extendAndCopy(zone_playersIdentifier, int(player1.ID()), ElementKindPlayer, ComplexID{}).toJSONPath(),
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
									Level:         intPtr(1),
								},
							},
						},
					},
				}
			},
			func(errText string) {
				t.Errorf(errText)
			},
			false,
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
						BoundTo: &elementReference{
							OperationKind:        OperationKindUnchanged,
							ElementID:            int(plyr.ID()),
							ElementKind:          ElementKindPlayer,
							ReferencedDataStatus: ReferencedDataModified,
							ElementPath:          newPath().extendAndCopy(zoneIdentifier, int(zne.ID()), ElementKindZone, ComplexID{}).extendAndCopy(zone_interactablesIdentifier, int(plyr.ID()), ElementKindPlayer, ComplexID{}).toJSONPath(),
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
									Level:         intPtr(1),
								},
							},
						},
					},
				}
			},
			func(errText string) {
				t.Errorf(errText)
			},
			false,
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
						Target: &elementReference{
							OperationKind:        OperationKindUpdate,
							ElementID:            int(player2.ID()),
							ElementKind:          ElementKindPlayer,
							ReferencedDataStatus: ReferencedDataUnchanged,
							ElementPath:          newPath().extendAndCopy(playerIdentifier, int(player2.ID()), ElementKindPlayer, ComplexID(player1.Target().playerTargetRef.ID)).toJSONPath(),
						},
					},
				}
			},
			func(errText string) {
				t.Errorf(errText)
			},
			false,
		)
	})
	t.Run("includes reference to updated element through multiple reference links", func(t *testing.T) {
		newTreeTest(
			func(se *Engine, expectedTree *Tree) {
				player1 := se.CreatePlayer()
				player2 := se.CreatePlayer()
				item1 := se.CreateItem()
				equipmentSet1 := se.CreateEquipmentSet()

				equipmentSet1.AddEquipment(item1.ID())
				item1.SetBoundTo(player1.ID())
				player1.AddGuildMember(player2.ID())

				se.UpdateState()

				player2.GearScore().SetLevel(1)

				expectedTree.EquipmentSet = map[EquipmentSetID]equipmentSet{
					equipmentSet1.ID(): {
						ID:            equipmentSet1.ID(),
						OperationKind: OperationKindUnchanged,
						Equipment: map[ItemID]elementReference{
							item1.ID(): {
								OperationKind:        OperationKindUnchanged,
								ElementID:            int(item1.ID()),
								ElementKind:          ElementKindItem,
								ReferencedDataStatus: ReferencedDataModified,
								ElementPath:          newPath().extendAndCopy(itemIdentifier, int(item1.ID()), ElementKindItem, ComplexID{}).toJSONPath(),
							},
						},
					},
				}
				expectedTree.Item = map[ItemID]item{
					item1.ID(): {
						ID:            item1.ID(),
						OperationKind: OperationKindUnchanged,
						BoundTo: &elementReference{
							OperationKind:        OperationKindUnchanged,
							ElementID:            int(player1.ID()),
							ElementKind:          ElementKindPlayer,
							ReferencedDataStatus: ReferencedDataModified,
							ElementPath:          newPath().extendAndCopy(playerIdentifier, int(player1.ID()), ElementKindPlayer, ComplexID{}).toJSONPath(),
						},
					},
				}
				expectedTree.Player = map[PlayerID]player{
					player1.ID(): {
						ID:            player1.ID(),
						OperationKind: OperationKindUnchanged,
						GuildMembers: map[PlayerID]elementReference{
							player2.ID(): {
								OperationKind:        OperationKindUnchanged,
								ElementID:            int(player2.ID()),
								ElementKind:          ElementKindPlayer,
								ReferencedDataStatus: ReferencedDataModified,
								ElementPath:          newPath().extendAndCopy(playerIdentifier, int(player2.ID()), ElementKindPlayer, ComplexID{}).toJSONPath(),
							},
						},
					},
					player2.ID(): {
						ID:            player2.ID(),
						OperationKind: OperationKindUnchanged,
						GearScore: &gearScore{
							OperationKind: OperationKindUpdate,
							ID:            player2.GearScore().ID(),
							Level:         intPtr(1),
						},
					},
				}
			},
			func(errText string) {
				t.Errorf(errText)
			},
			false,
		)
	})
	t.Run("do not include reference if reference got unset", func(t *testing.T) {
		newTreeTest(
			func(se *Engine, expectedTree *Tree) {
				player1 := se.CreatePlayer()
				item1 := se.CreateItem()
				item2 := se.CreateItem()

				item1.SetBoundTo(player1.ID())

				se.UpdateState()

				item1.BoundTo().Unset()
				item2.SetBoundTo(player1.ID())
				item2.BoundTo().Unset()

				player1.GearScore().SetLevel(1)

				expectedTree.Item = map[ItemID]item{
					item1.ID(): {
						ID:            item1.ID(),
						OperationKind: OperationKindUpdate,
						BoundTo: &elementReference{
							ElementKind:          ElementKindPlayer,
							ElementPath:          newPath().extendAndCopy(playerIdentifier, int(player1.ID()), ElementKindPlayer, ComplexID{}).toJSONPath(),
							ElementID:            int(player1.ID()),
							OperationKind:        OperationKindDelete,
							ReferencedDataStatus: ReferencedDataModified,
						},
					},
					item2.ID(): {
						ID:            item2.ID(),
						OperationKind: OperationKindUpdate,
					},
				}
				expectedTree.Player = map[PlayerID]player{
					player1.ID(): {
						ID:            player1.ID(),
						OperationKind: OperationKindUnchanged,
						GearScore: &gearScore{
							OperationKind: OperationKindUpdate,
							ID:            player1.GearScore().ID(),
							Level:         intPtr(1),
						},
					},
				}
			},
			func(errText string) {
				t.Errorf(errText)
			},
			false,
		)
	})
	t.Run("assembles entire tree correctly", func(t *testing.T) {
		newTreeTest(
			func(se *Engine, expectedTree *Tree) {
				plyr := se.CreatePlayer()
				itm := plyr.AddItem().SetName("item0").SetBoundTo(plyr.ID())
				se.UpdateState()

				expectedTree.Player = map[PlayerID]player{
					plyr.ID(): {
						ID:            plyr.ID(),
						OperationKind: OperationKindUnchanged,
						GearScore: &gearScore{
							ID:            plyr.GearScore().ID(),
							Level:         new(int64),
							Score:         new(int64),
							OperationKind: OperationKindUnchanged,
						},
						Position: &position{
							ID:            plyr.Position().ID(),
							X:             new(float64),
							Y:             new(float64),
							OperationKind: OperationKindUnchanged,
						},
						Items: map[ItemID]item{
							itm.ID(): {
								ID:            itm.ID(),
								Name:          stringPtr("item0"),
								OperationKind: OperationKindUnchanged,
								GearScore: &gearScore{
									ID:            itm.GearScore().ID(),
									Level:         new(int64),
									Score:         new(int64),
									OperationKind: OperationKindUnchanged,
								},
								BoundTo: &elementReference{
									OperationKind:        OperationKindUnchanged,
									ElementID:            int(plyr.ID()),
									ElementKind:          ElementKindPlayer,
									ReferencedDataStatus: ReferencedDataUnchanged,
									ElementPath:          newPath().extendAndCopy(playerIdentifier, int(plyr.ID()), ElementKindPlayer, ComplexID(itm.BoundTo().itemBoundToRef.ID)).toJSONPath(),
								},
								Origin: &player{
									ID:            itm.Origin().Player().ID(),
									OperationKind: OperationKindUnchanged,
									Position: &position{
										ID:            itm.Origin().Player().Position().ID(),
										X:             new(float64),
										Y:             new(float64),
										OperationKind: OperationKindUnchanged,
									},
									GearScore: &gearScore{
										ID:            itm.Origin().Player().GearScore().ID(),
										Level:         new(int64),
										Score:         new(int64),
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
			true,
		)
	})
}
