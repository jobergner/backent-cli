package state

func (engine *Engine) walkGearScore(gearScoreID GearScoreID, p path) {
	engine.PathTrack.gearScore[gearScoreID] = p
}

func (engine *Engine) walkPosition(positionID PositionID, p path) {
	engine.PathTrack.position[positionID] = p
}

func (engine *Engine) walkEquipmentSet(equipmentSetID EquipmentSetID, p path) {
	engine.PathTrack.equipmentSet[equipmentSetID] = p
}

func (engine *Engine) walkItem(itemID ItemID, p path) {
	itemData, hasUpdated := engine.Patch.Item[itemID]
	if !hasUpdated {
		itemData = engine.State.Item[itemID]
	}

	var gearScorePath path
	if existingPath, pathExists := engine.PathTrack.gearScore[itemData.GearScore]; !pathExists {
		gearScorePath = p.gearScore()
	} else {
		gearScorePath = existingPath
	}
	engine.walkGearScore(itemData.GearScore, gearScorePath)

	engine.PathTrack.item[itemID] = p
}

func (engine *Engine) walkZoneItem(zoneItemID ZoneItemID, p path) {
	zoneItemData, hasUpdated := engine.Patch.ZoneItem[zoneItemID]
	if !hasUpdated {
		zoneItemData = engine.State.ZoneItem[zoneItemID]
	}

	var itemPath path
	if existingPath, pathExists := engine.PathTrack.item[zoneItemData.Item]; !pathExists {
		itemPath = p.item()
	} else {
		itemPath = existingPath
	}
	engine.walkItem(zoneItemData.Item, itemPath)

	var positionPath path
	if existingPath, pathExists := engine.PathTrack.position[zoneItemData.Position]; !pathExists {
		positionPath = p.position()
	} else {
		positionPath = existingPath
	}
	engine.walkPosition(zoneItemData.Position, positionPath)

	engine.PathTrack.zoneItem[zoneItemID] = p
}

func (engine *Engine) walkPlayer(playerID PlayerID, p path) {
	playerData, hasUpdated := engine.Patch.Player[playerID]
	if !hasUpdated {
		playerData = engine.State.Player[playerID]
	}

	var gearScorePath path
	if existingPath, pathExists := engine.PathTrack.gearScore[playerData.GearScore]; !pathExists {
		gearScorePath = p.gearScore()
	} else {
		gearScorePath = existingPath
	}
	engine.walkGearScore(playerData.GearScore, gearScorePath)

	for i, itemID := range mergeItemIDs(engine.State.Player[playerData.ID].Items, engine.Patch.Player[playerData.ID].Items) {
		var itemsPath path
		if existingPath, pathExists := engine.PathTrack.item[itemID]; !pathExists || !existingPath.equals(p) {
			itemsPath = p.items().index(i)
		} else {
			itemsPath = existingPath
		}
		engine.walkItem(itemID, itemsPath)
	}

	var positionPath path
	if existingPath, pathExists := engine.PathTrack.position[playerData.Position]; !pathExists {
		positionPath = p.position()
	} else {
		positionPath = existingPath
	}
	engine.walkPosition(playerData.Position, positionPath)

	engine.PathTrack.player[playerID] = p
}

func (engine *Engine) walkZone(zoneID ZoneID, p path) {
	zoneData, hasUpdated := engine.Patch.Zone[zoneID]
	if !hasUpdated {
		zoneData = engine.State.Zone[zoneID]
	}

	for i, zoneItemID := range mergeZoneItemIDs(engine.State.Zone[zoneData.ID].Items, engine.Patch.Zone[zoneData.ID].Items) {
		var itemsPath path
		if existingPath, pathExists := engine.PathTrack.zoneItem[zoneItemID]; !pathExists || !existingPath.equals(p) {
			itemsPath = p.items().index(i)
		} else {
			itemsPath = existingPath
		}
		engine.walkZoneItem(zoneItemID, itemsPath)
	}

	for i, playerID := range mergePlayerIDs(engine.State.Zone[zoneData.ID].Players, engine.Patch.Zone[zoneData.ID].Players) {
		var playersPath path
		if existingPath, pathExists := engine.PathTrack.player[playerID]; !pathExists || !existingPath.equals(p) {
			playersPath = p.players().index(i)
		} else {
			playersPath = existingPath
		}
		engine.walkPlayer(playerID, playersPath)
	}

	engine.PathTrack.zone[zoneID] = p
}

func (engine *Engine) walkTree() {
	walkedCheck := newRecursionCheck()

	for id, equipmentSetData := range engine.Patch.EquipmentSet {
		engine.walkEquipmentSet(equipmentSetData.ID, newPath(equipmentSetIdentifier, int(id)))
		walkedCheck.equipmentSet[equipmentSetData.ID] = true
	}
	for id, gearScoreData := range engine.Patch.GearScore {
		if !gearScoreData.HasParent {
			engine.walkGearScore(gearScoreData.ID, newPath(gearScoreIdentifier, int(id)))
			walkedCheck.gearScore[gearScoreData.ID] = true
		}
	}
	for id, itemData := range engine.Patch.Item {
		if !itemData.HasParent {
			engine.walkItem(itemData.ID, newPath(itemIdentifier, int(id)))
			walkedCheck.item[itemData.ID] = true
		}
	}
	for id, playerData := range engine.Patch.Player {
		if !playerData.HasParent {
			engine.walkPlayer(playerData.ID, newPath(playerIdentifier, int(id)))
			walkedCheck.player[playerData.ID] = true
		}
	}
	for id, positionData := range engine.Patch.Position {
		if !positionData.HasParent {
			engine.walkPosition(positionData.ID, newPath(positionIdentifier, int(id)))
			walkedCheck.position[positionData.ID] = true
		}
	}
	for id, zoneData := range engine.Patch.Zone {
		engine.walkZone(zoneData.ID, newPath(zoneIdentifier, int(id)))
		walkedCheck.zone[zoneData.ID] = true
	}
	for id, zoneItemData := range engine.Patch.ZoneItem {
		if !zoneItemData.HasParent {
			engine.walkZoneItem(zoneItemData.ID, newPath(zoneItemIdentifier, int(id)))
			walkedCheck.zoneItem[zoneItemData.ID] = true
		}
	}

	for id, equipmentSetData := range engine.State.EquipmentSet {
		if _, ok := walkedCheck.equipmentSet[equipmentSetData.ID]; !ok {
			engine.walkEquipmentSet(equipmentSetData.ID, newPath(equipmentSetIdentifier, int(id)))
		}
	}
	for id, gearScoreData := range engine.State.GearScore {
		if !gearScoreData.HasParent {
			if _, ok := walkedCheck.gearScore[gearScoreData.ID]; !ok {
				engine.walkGearScore(gearScoreData.ID, newPath(gearScoreIdentifier, int(id)))
			}
		}
	}
	for id, itemData := range engine.State.Item {
		if !itemData.HasParent {
			if _, ok := walkedCheck.item[itemData.ID]; !ok {
				engine.walkItem(itemData.ID, newPath(itemIdentifier, int(id)))
			}
		}
	}
	for id, playerData := range engine.State.Player {
		if !playerData.HasParent {
			if _, ok := walkedCheck.player[playerData.ID]; !ok {
				engine.walkPlayer(playerData.ID, newPath(playerIdentifier, int(id)))
			}
		}
	}
	for id, positionData := range engine.State.Position {
		if !positionData.HasParent {
			if _, ok := walkedCheck.position[positionData.ID]; !ok {
				engine.walkPosition(positionData.ID, newPath(positionIdentifier, int(id)))
			}
		}
	}
	for id, zoneData := range engine.State.Zone {
		if _, ok := walkedCheck.zone[zoneData.ID]; !ok {
			engine.walkZone(zoneData.ID, newPath(zoneIdentifier, int(id)))
		}
	}
	for id, zoneItemData := range engine.State.ZoneItem {
		if !zoneItemData.HasParent {
			if _, ok := walkedCheck.zoneItem[zoneItemData.ID]; !ok {
				engine.walkZoneItem(zoneItemData.ID, newPath(zoneItemIdentifier, int(id)))
			}
		}
	}

	engine.PathTrack._iterations += 1
	if engine.PathTrack._iterations == 100 {
		for key := range engine.PathTrack.equipmentSet {
			delete(engine.PathTrack.equipmentSet, key)
		}
		for key := range engine.PathTrack.gearScore {
			delete(engine.PathTrack.gearScore, key)
		}
		for key := range engine.PathTrack.item {
			delete(engine.PathTrack.item, key)
		}
		for key := range engine.PathTrack.player {
			delete(engine.PathTrack.player, key)
		}
		for key := range engine.PathTrack.position {
			delete(engine.PathTrack.position, key)
		}
		for key := range engine.PathTrack.zone {
			delete(engine.PathTrack.zone, key)
		}
		for key := range engine.PathTrack.zoneItem {
			delete(engine.PathTrack.zoneItem, key)
		}
	}
}
