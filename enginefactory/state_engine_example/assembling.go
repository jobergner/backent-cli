package state

func (se *Engine) assembleGearScore(gearScoreID GearScoreID) (GearScore, bool) {
	gearScoreData, hasUpdated := se.Patch.GearScore[gearScoreID]
	if !hasUpdated {
		return GearScore{}, false
	}

	var gearScore GearScore

	gearScore.ID = gearScoreData.ID
	gearScore.OperationKind_ = gearScoreData.OperationKind_
	gearScore.Level = gearScoreData.Level
	gearScore.Score = gearScoreData.Score
	return gearScore, true
}

func (se *Engine) assemblePosition(positionID PositionID) (Position, bool) {
	positionData, hasUpdated := se.Patch.Position[positionID]
	if !hasUpdated {
		return Position{}, false
	}

	var position Position

	position.ID = positionData.ID
	position.OperationKind_ = positionData.OperationKind_
	position.X = positionData.X
	position.Y = positionData.Y
	return position, true
}

func (se *Engine) assembleItem(itemID ItemID) (Item, bool) {
	itemData, hasUpdated := se.Patch.Item[itemID]
	if !hasUpdated {
		itemData = se.State.Item[itemID]
	}

	var item Item

	if treeGearScore, gearScoreHasUpdated := se.assembleGearScore(itemData.GearScore); gearScoreHasUpdated {
		hasUpdated = true
		item.GearScore = &treeGearScore
	}

	item.ID = itemData.ID
	item.OperationKind_ = itemData.OperationKind_
	// item.BoundTo = se.evalItemBoundToElementRef(itemData)
	item.Name = itemData.Name
	return item, hasUpdated
}

func (se *Engine) assembleZoneItem(zoneItemID ZoneItemID) (ZoneItem, bool) {
	zoneItemData, hasUpdated := se.Patch.ZoneItem[zoneItemID]
	if !hasUpdated {
		zoneItemData = se.State.ZoneItem[zoneItemID]
	}

	var zoneItem ZoneItem

	if treeItem, itemHasUpdated := se.assembleItem(zoneItemData.Item); itemHasUpdated {
		hasUpdated = true
		zoneItem.Item = &treeItem
	}
	if treePosition, positionHasUpdated := se.assemblePosition(zoneItemData.Position); positionHasUpdated {
		hasUpdated = true
		zoneItem.Position = &treePosition
	}

	zoneItem.ID = zoneItemData.ID
	zoneItem.OperationKind_ = zoneItemData.OperationKind_
	return zoneItem, hasUpdated

}

func (se *Engine) assemblePlayer(playerID PlayerID) (Player, bool) {
	playerData, hasUpdated := se.Patch.Player[playerID]
	if !hasUpdated {
		playerData = se.State.Player[playerID]
	}

	var player Player

	if treeGearScore, gearScoreHasUpdated := se.assembleGearScore(playerData.GearScore); gearScoreHasUpdated {
		hasUpdated = true
		player.GearScore = &treeGearScore
	}
	for _, itemID := range mergeItemIDs(se.State.Player[playerData.ID].Items, se.Patch.Player[playerData.ID].Items) {
		if treeItem, itemHasUpdated := se.assembleItem(itemID); itemHasUpdated {
			hasUpdated = true
			player.Items = append(player.Items, treeItem)
		}
	}
	if treePosition, positionHasUpdated := se.assemblePosition(playerData.Position); positionHasUpdated {
		hasUpdated = true
		player.Position = &treePosition
	}

	player.ID = playerData.ID
	player.OperationKind_ = playerData.OperationKind_
	// player.GuildMembers = se.evalPlayerGuildMembersElementRefs(playerData)
	return player, hasUpdated
}

func (se *Engine) assembleZone(zoneID ZoneID) (Zone, bool) {
	zoneData, hasUpdated := se.Patch.Zone[zoneID]
	if !hasUpdated {
		zoneData = se.State.Zone[zoneID]
	}

	var zone Zone

	for _, zoneItemID := range mergeZoneItemIDs(se.State.Zone[zoneData.ID].Items, se.Patch.Zone[zoneData.ID].Items) {
		if treeZoneItem, zoneItemHasUpdated := se.assembleZoneItem(zoneItemID); zoneItemHasUpdated {
			hasUpdated = true
			zone.Items = append(zone.Items, treeZoneItem)
		}
	}
	for _, playerID := range mergePlayerIDs(se.State.Zone[zoneData.ID].Players, se.Patch.Zone[zoneData.ID].Players) {
		if treePlayer, playerHasUpdated := se.assemblePlayer(playerID); playerHasUpdated {
			hasUpdated = true
			zone.Players = append(zone.Players, treePlayer)
		}
	}

	zone.ID = zoneData.ID
	zone.OperationKind_ = zoneData.OperationKind_
	zone.Tags = zoneData.Tags
	return zone, hasUpdated
}

func (se *Engine) assembleTree() Tree {
	tree := newTree()
	for _, gearScoreData := range se.Patch.GearScore {
		if !gearScoreData.HasParent_ {
			gearScore, hasUpdated := se.assembleGearScore(gearScoreData.ID)
			if hasUpdated {
				tree.GearScore[gearScoreData.ID] = gearScore
			}
		}
	}
	for _, itemData := range se.Patch.Item {
		if !itemData.HasParent_ {
			item, hasUpdated := se.assembleItem(itemData.ID)
			if hasUpdated {
				tree.Item[itemData.ID] = item
			}
		}
	}
	for _, playerData := range se.Patch.Player {
		if !playerData.HasParent_ {
			player, hasUpdated := se.assemblePlayer(playerData.ID)
			if hasUpdated {
				tree.Player[playerData.ID] = player
			}
		}
	}
	for _, positionData := range se.Patch.Position {
		if !positionData.HasParent_ {
			position, hasUpdated := se.assemblePosition(positionData.ID)
			if hasUpdated {
				tree.Position[positionData.ID] = position
			}
		}
	}
	for _, zoneData := range se.Patch.Zone {
		zone, hasUpdated := se.assembleZone(zoneData.ID)
		if hasUpdated {
			tree.Zone[zoneData.ID] = zone
		}
	}
	for _, zoneItemData := range se.Patch.ZoneItem {
		if !zoneItemData.HasParent_ {
			zoneItem, hasUpdated := se.assembleZoneItem(zoneItemData.ID)
			if hasUpdated {
				tree.ZoneItem[zoneItemData.ID] = zoneItem
			}
		}
	}

	for _, gearScoreData := range se.State.GearScore {
		if !gearScoreData.HasParent_ {
			if _, ok := tree.GearScore[gearScoreData.ID]; !ok {
				gearScore, hasUpdated := se.assembleGearScore(gearScoreData.ID)
				if hasUpdated {
					tree.GearScore[gearScoreData.ID] = gearScore
				}
			}
		}
	}
	for _, itemData := range se.State.Item {
		if !itemData.HasParent_ {
			if _, ok := tree.Item[itemData.ID]; !ok {
				item, hasUpdated := se.assembleItem(itemData.ID)
				if hasUpdated {
					tree.Item[itemData.ID] = item
				}
			}
		}
	}
	for _, playerData := range se.State.Player {
		if !playerData.HasParent_ {
			if _, ok := tree.Player[playerData.ID]; !ok {
				player, hasUpdated := se.assemblePlayer(playerData.ID)
				if hasUpdated {
					tree.Player[playerData.ID] = player
				}
			}
		}
	}
	for _, positionData := range se.State.Position {
		if !positionData.HasParent_ {
			if _, ok := tree.Position[positionData.ID]; !ok {
				position, hasUpdated := se.assemblePosition(positionData.ID)
				if hasUpdated {
					tree.Position[positionData.ID] = position
				}
			}
		}
	}
	for _, zoneData := range se.State.Zone {
		if _, ok := tree.Zone[zoneData.ID]; !ok {
			zone, hasUpdated := se.assembleZone(zoneData.ID)
			if hasUpdated {
				tree.Zone[zoneData.ID] = zone
			}
		}
	}
	for _, zoneItemData := range se.State.ZoneItem {
		if !zoneItemData.HasParent_ {
			if _, ok := tree.ZoneItem[zoneItemData.ID]; !ok {
				zoneItem, hasUpdated := se.assembleZoneItem(zoneItemData.ID)
				if hasUpdated {
					tree.ZoneItem[zoneItemData.ID] = zoneItem
				}
			}
		}
	}

	return tree
}
