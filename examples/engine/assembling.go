package state

type assembleConfig struct {
	forceInclude bool // include everything, regardless of update status
}

func (engine *Engine) assembleGearScore(gearScoreID GearScoreID, check *recursionCheck, config assembleConfig) (gearScore, bool, bool) {
	if check != nil {
		if alreadyExists := check.gearScore[gearScoreID]; alreadyExists {
			return gearScore{}, false, false
		} else {
			check.gearScore[gearScoreID] = true
		}
	}

	gearScoreData, hasUpdated := engine.Patch.GearScore[gearScoreID]
	if !hasUpdated {
		gearScoreData = engine.State.GearScore[gearScoreID]
	}

	if cachedGearScore, ok := engine.assembleCache.gearScore[gearScoreData.ID]; ok && !config.forceInclude {
		return cachedGearScore.gearScore, cachedGearScore.hasUpdated || config.forceInclude, cachedGearScore.hasUpdated
	}

	var element gearScore

	element.ID = gearScoreData.ID
	element.OperationKind = gearScoreData.OperationKind
	element.Level = gearScoreData.Level
	element.Score = gearScoreData.Score

	engine.assembleCache.gearScore[element.ID] = gearScoreCacheElement{hasUpdated: hasUpdated, gearScore: element}

	return element, hasUpdated || config.forceInclude, hasUpdated
}

func (engine *Engine) assemblePosition(positionID PositionID, check *recursionCheck, config assembleConfig) (position, bool, bool) {
	if check != nil {
		if alreadyExists := check.position[positionID]; alreadyExists {
			return position{}, false, false
		} else {
			check.position[positionID] = true
		}
	}

	positionData, hasUpdated := engine.Patch.Position[positionID]
	if !hasUpdated {
		positionData = engine.State.Position[positionID]
	}

	if cachedPosition, ok := engine.assembleCache.position[positionData.ID]; ok && !config.forceInclude {
		return cachedPosition.position, cachedPosition.hasUpdated || config.forceInclude, cachedPosition.hasUpdated
	}

	var element position

	element.ID = positionData.ID
	element.OperationKind = positionData.OperationKind
	element.X = positionData.X
	element.Y = positionData.Y

	engine.assembleCache.position[element.ID] = positionCacheElement{hasUpdated: hasUpdated, position: element}

	return element, hasUpdated || config.forceInclude, hasUpdated
}

func (engine *Engine) assembleEquipmentSet(equipmentSetID EquipmentSetID, check *recursionCheck, config assembleConfig) (equipmentSet, bool, bool) {
	if check != nil {
		if alreadyExists := check.equipmentSet[equipmentSetID]; alreadyExists {
			return equipmentSet{}, false, false
		} else {
			check.equipmentSet[equipmentSetID] = true
		}
	}

	equipmentSetData, hasUpdated := engine.Patch.EquipmentSet[equipmentSetID]
	if !hasUpdated {
		equipmentSetData = engine.State.EquipmentSet[equipmentSetID]
	}

	if cachedEquipmentSet, ok := engine.assembleCache.equipmentSet[equipmentSetData.ID]; ok && !config.forceInclude {
		return cachedEquipmentSet.equipmentSet, cachedEquipmentSet.hasUpdated || config.forceInclude, cachedEquipmentSet.hasUpdated
	}

	var element equipmentSet

	for _, equipmentSetEquipmentRefID := range mergeEquipmentSetEquipmentRefIDs(engine.State.EquipmentSet[equipmentSetData.ID].Equipment, engine.Patch.EquipmentSet[equipmentSetData.ID].Equipment) {
		if treeEquipmentSetEquipmentRef, include, childHasUpdated := engine.assembleEquipmentSetEquipmentRef(equipmentSetEquipmentRefID, check, config); include {
			if childHasUpdated {
				hasUpdated = true
			}
			if element.Equipment == nil {
				element.Equipment = make(map[ItemID]itemReference)
			}
			element.Equipment[treeEquipmentSetEquipmentRef.ElementID] = treeEquipmentSetEquipmentRef
		}
	}

	element.ID = equipmentSetData.ID
	element.OperationKind = equipmentSetData.OperationKind
	element.Name = equipmentSetData.Name

	engine.assembleCache.equipmentSet[element.ID] = equipmentSetCacheElement{hasUpdated: hasUpdated, equipmentSet: element}

	return element, hasUpdated || config.forceInclude, hasUpdated
}

func (engine *Engine) assembleItem(itemID ItemID, check *recursionCheck, config assembleConfig) (item, bool, bool) {
	if check != nil {
		if alreadyExists := check.item[itemID]; alreadyExists {
			return item{}, false, false
		} else {
			check.item[itemID] = true
		}
	}

	itemData, hasUpdated := engine.Patch.Item[itemID]
	if !hasUpdated {
		itemData = engine.State.Item[itemID]
	}

	if cachedItem, ok := engine.assembleCache.item[itemData.ID]; ok && !config.forceInclude {
		return cachedItem.item, cachedItem.hasUpdated || config.forceInclude, cachedItem.hasUpdated
	}

	var element item

	if treeItemBoundToRef, include, childHasUpdated := engine.assembleItemBoundToRef(itemID, check, config); include {
		if childHasUpdated {
			hasUpdated = true
		}
		element.BoundTo = treeItemBoundToRef
	}

	if treeGearScore, include, childHasUpdated := engine.assembleGearScore(itemData.GearScore, check, config); include {
		if childHasUpdated {
			hasUpdated = true
		}
		element.GearScore = &treeGearScore
	}

	anyOfPlayer_PositionContainer := engine.anyOfPlayer_Position(itemData.Origin).anyOfPlayer_Position
	if anyOfPlayer_PositionContainer.ElementKind == ElementKindPlayer {
		playerID := anyOfPlayer_PositionContainer.Player
		if treePlayer, include, childHasUpdated := engine.assemblePlayer(playerID, check, config); include {
			if childHasUpdated {
				hasUpdated = true
			}
			element.Origin = &treePlayer
		}
	} else if anyOfPlayer_PositionContainer.ElementKind == ElementKindPosition {
		positionID := anyOfPlayer_PositionContainer.Position
		if treePosition, include, childHasUpdated := engine.assemblePosition(positionID, check, config); include {
			if childHasUpdated {
				hasUpdated = true
			}
			element.Origin = &treePosition
		}
	}

	element.ID = itemData.ID
	element.OperationKind = itemData.OperationKind
	element.Name = itemData.Name

	engine.assembleCache.item[element.ID] = itemCacheElement{hasUpdated: hasUpdated, item: element}

	return element, hasUpdated || config.forceInclude, hasUpdated
}

func (engine *Engine) assembleZoneItem(zoneItemID ZoneItemID, check *recursionCheck, config assembleConfig) (zoneItem, bool, bool) {
	if check != nil {
		if alreadyExists := check.zoneItem[zoneItemID]; alreadyExists {
			return zoneItem{}, false, false
		} else {
			check.zoneItem[zoneItemID] = true
		}
	}

	zoneItemData, hasUpdated := engine.Patch.ZoneItem[zoneItemID]
	if !hasUpdated {
		zoneItemData = engine.State.ZoneItem[zoneItemID]
	}

	if cachedZoneItem, ok := engine.assembleCache.zoneItem[zoneItemData.ID]; ok && !config.forceInclude {
		return cachedZoneItem.zoneItem, cachedZoneItem.hasUpdated || config.forceInclude, cachedZoneItem.hasUpdated
	}

	var element zoneItem

	if treeItem, include, childHasUpdated := engine.assembleItem(zoneItemData.Item, check, config); include {
		if childHasUpdated {
			hasUpdated = true
		}
		element.Item = &treeItem
	}

	if treePosition, include, childHasUpdated := engine.assemblePosition(zoneItemData.Position, check, config); include {
		if childHasUpdated {
			hasUpdated = true
		}
		element.Position = &treePosition
	}

	element.ID = zoneItemData.ID
	element.OperationKind = zoneItemData.OperationKind

	engine.assembleCache.zoneItem[element.ID] = zoneItemCacheElement{hasUpdated: hasUpdated, zoneItem: element}

	return element, hasUpdated || config.forceInclude, hasUpdated
}

func (engine *Engine) assemblePlayer(playerID PlayerID, check *recursionCheck, config assembleConfig) (player, bool, bool) {
	if check != nil {
		if alreadyExists := check.player[playerID]; alreadyExists {
			return player{}, false, false
		} else {
			check.player[playerID] = true
		}
	}

	playerData, hasUpdated := engine.Patch.Player[playerID]
	if !hasUpdated {
		playerData = engine.State.Player[playerID]
	}

	if cachedPlayer, ok := engine.assembleCache.player[playerData.ID]; ok && !config.forceInclude {
		return cachedPlayer.player, cachedPlayer.hasUpdated || config.forceInclude, cachedPlayer.hasUpdated
	}

	var element player

	for _, playerEquipmentSetRefID := range mergePlayerEquipmentSetRefIDs(engine.State.Player[playerData.ID].EquipmentSets, engine.Patch.Player[playerData.ID].EquipmentSets) {
		if treePlayerEquipmentSetRef, include, childHasUpdated := engine.assemblePlayerEquipmentSetRef(playerEquipmentSetRefID, check, config); include {
			if childHasUpdated {
				hasUpdated = true
			}
			if element.EquipmentSets == nil {
				element.EquipmentSets = make(map[EquipmentSetID]equipmentSetReference)
			}
			element.EquipmentSets[treePlayerEquipmentSetRef.ElementID] = treePlayerEquipmentSetRef
		}
	}

	if treeGearScore, include, childHasUpdated := engine.assembleGearScore(playerData.GearScore, check, config); include {
		if childHasUpdated {
			hasUpdated = true
		}
		element.GearScore = &treeGearScore
	}

	for _, playerGuildMemberRefID := range mergePlayerGuildMemberRefIDs(engine.State.Player[playerData.ID].GuildMembers, engine.Patch.Player[playerData.ID].GuildMembers) {
		if treePlayerGuildMemberRef, include, childHasUpdated := engine.assemblePlayerGuildMemberRef(playerGuildMemberRefID, check, config); include {
			if childHasUpdated {
				hasUpdated = true
			}
			if element.GuildMembers == nil {
				element.GuildMembers = make(map[PlayerID]playerReference)
			}
			element.GuildMembers[treePlayerGuildMemberRef.ElementID] = treePlayerGuildMemberRef
		}
	}

	for _, itemID := range mergeItemIDs(engine.State.Player[playerData.ID].Items, engine.Patch.Player[playerData.ID].Items) {
		if treeItem, include, childHasUpdated := engine.assembleItem(itemID, check, config); include {
			if childHasUpdated {
				hasUpdated = true
			}
			if element.Items == nil {
				element.Items = make(map[ItemID]item)
			}
			element.Items[treeItem.ID] = treeItem
		}
	}

	if treePosition, include, childHasUpdated := engine.assemblePosition(playerData.Position, check, config); include {
		if childHasUpdated {
			hasUpdated = true
		}
		element.Position = &treePosition
	}

	if treePlayerTargetRef, include, childHasUpdated := engine.assemblePlayerTargetRef(playerID, check, config); include {
		if childHasUpdated {
			hasUpdated = true
		}
		element.Target = treePlayerTargetRef
	}

	for _, playerTargetedByRefID := range mergePlayerTargetedByRefIDs(engine.State.Player[playerData.ID].TargetedBy, engine.Patch.Player[playerData.ID].TargetedBy) {
		if treePlayerTargetedByRef, include, childHasUpdated := engine.assemblePlayerTargetedByRef(playerTargetedByRefID, check, config); include {
			if childHasUpdated {
				hasUpdated = true
			}
			if element.TargetedBy == nil {
				element.TargetedBy = make(map[int]anyOfPlayer_ZoneItemReference)
			}
			element.TargetedBy[treePlayerTargetedByRef.ElementID] = treePlayerTargetedByRef
		}
	}

	element.ID = playerData.ID
	element.OperationKind = playerData.OperationKind

	engine.assembleCache.player[element.ID] = playerCacheElement{hasUpdated: hasUpdated, player: element}

	return element, hasUpdated || config.forceInclude, hasUpdated
}

func (engine *Engine) assembleZone(zoneID ZoneID, check *recursionCheck, config assembleConfig) (zone, bool, bool) {
	if check != nil {
		if alreadyExists := check.zone[zoneID]; alreadyExists {
			return zone{}, false, false
		} else {
			check.zone[zoneID] = true
		}
	}

	zoneData, hasUpdated := engine.Patch.Zone[zoneID]
	if !hasUpdated {
		zoneData = engine.State.Zone[zoneID]
	}

	if cachedZone, ok := engine.assembleCache.zone[zoneData.ID]; ok && !config.forceInclude {
		return cachedZone.zone, cachedZone.hasUpdated || config.forceInclude, cachedZone.hasUpdated
	}

	var element zone

	for _, anyOfItem_Player_ZoneItemID := range mergeAnyOfItem_Player_ZoneItemIDs(engine.State.Zone[zoneData.ID].Interactables, engine.Patch.Zone[zoneData.ID].Interactables) {
		anyOfItem_Player_ZoneItemContainer := engine.anyOfItem_Player_ZoneItem(anyOfItem_Player_ZoneItemID).anyOfItem_Player_ZoneItem
		if anyOfItem_Player_ZoneItemContainer.ElementKind == ElementKindItem {
			itemID := anyOfItem_Player_ZoneItemContainer.Item
			if treeItem, include, childHasUpdated := engine.assembleItem(itemID, check, config); include {
				if childHasUpdated {
					hasUpdated = true
				}
				if element.Interactables == nil {
					element.Interactables = make(map[int]interface{})
				}
				element.Interactables[int(treeItem.ID)] = treeItem
			}
		} else if anyOfItem_Player_ZoneItemContainer.ElementKind == ElementKindPlayer {
			playerID := anyOfItem_Player_ZoneItemContainer.Player
			if treePlayer, include, childHasUpdated := engine.assemblePlayer(playerID, check, config); include {
				if childHasUpdated {
					hasUpdated = true
				}
				if element.Interactables == nil {
					element.Interactables = make(map[int]interface{})
				}
				element.Interactables[int(treePlayer.ID)] = treePlayer
			}
		} else if anyOfItem_Player_ZoneItemContainer.ElementKind == ElementKindZoneItem {
			zoneItemID := anyOfItem_Player_ZoneItemContainer.ZoneItem
			if treeZoneItem, include, childHasUpdated := engine.assembleZoneItem(zoneItemID, check, config); include {
				if childHasUpdated {
					hasUpdated = true
				}
				if element.Interactables == nil {
					element.Interactables = make(map[int]interface{})
				}
				element.Interactables[int(treeZoneItem.ID)] = treeZoneItem
			}
		}
	}

	for _, zoneItemID := range mergeZoneItemIDs(engine.State.Zone[zoneData.ID].Items, engine.Patch.Zone[zoneData.ID].Items) {
		if treeZoneItem, include, childHasUpdated := engine.assembleZoneItem(zoneItemID, check, config); include {
			if childHasUpdated {
				hasUpdated = true
			}
			if element.Items == nil {
				element.Items = make(map[ZoneItemID]zoneItem)
			}
			element.Items[treeZoneItem.ID] = treeZoneItem
		}
	}

	for _, playerID := range mergePlayerIDs(engine.State.Zone[zoneData.ID].Players, engine.Patch.Zone[zoneData.ID].Players) {
		if treePlayer, include, childHasUpdated := engine.assemblePlayer(playerID, check, config); include {
			if childHasUpdated {
				hasUpdated = true
			}
			if element.Players == nil {
				element.Players = make(map[PlayerID]player)
			}
			element.Players[treePlayer.ID] = treePlayer
		}
	}

	element.ID = zoneData.ID
	element.OperationKind = zoneData.OperationKind
	element.Tags = zoneData.Tags

	engine.assembleCache.zone[element.ID] = zoneCacheElement{hasUpdated: hasUpdated, zone: element}

	return element, hasUpdated || config.forceInclude, hasUpdated
}

func (engine *Engine) assemblePlayerTargetRef(playerID PlayerID, check *recursionCheck, config assembleConfig) (*anyOfPlayer_ZoneItemReference, bool, bool) {
	statePlayer := engine.State.Player[playerID]
	patchPlayer, playerIsInPatch := engine.Patch.Player[playerID]

	// ref not set at all
	if statePlayer.Target == 0 && (!playerIsInPatch || patchPlayer.Target == 0) {
		return nil, false, false
	}

	// force include
	if config.forceInclude {
		var ref PlayerTargetRef
		if patchPlayer.ID == 0 {
			ref = engine.playerTargetRef(statePlayer.Target)
		} else {
			ref = engine.playerTargetRef(patchPlayer.Target)
		}
		anyContainer := engine.anyOfPlayer_ZoneItem(ref.playerTargetRef.ReferencedElementID)
		if anyContainer.anyOfPlayer_ZoneItem.ElementKind == ElementKindPlayer {
			referencedElement := engine.Player(anyContainer.anyOfPlayer_ZoneItem.Player).player
			return &anyOfPlayer_ZoneItemReference{ref.playerTargetRef.OperationKind, int(referencedElement.ID), ElementKindPlayer, ReferencedDataUnchanged, referencedElement.Path}, true, false
		} else if anyContainer.anyOfPlayer_ZoneItem.ElementKind == ElementKindZoneItem {
			referencedElement := engine.ZoneItem(anyContainer.anyOfPlayer_ZoneItem.ZoneItem).zoneItem
			return &anyOfPlayer_ZoneItemReference{ref.playerTargetRef.OperationKind, int(referencedElement.ID), ElementKindZoneItem, ReferencedDataUnchanged, referencedElement.Path}, true, false
		}
	}

	// ref was definitely created
	if statePlayer.Target == 0 && (playerIsInPatch && patchPlayer.Target != 0) {
		ref := engine.playerTargetRef(patchPlayer.Target)
		anyContainer := engine.anyOfPlayer_ZoneItem(ref.playerTargetRef.ReferencedElementID)
		if anyContainer.anyOfPlayer_ZoneItem.ElementKind == ElementKindPlayer {
			if check == nil {
				check = newRecursionCheck()
			}
			referencedElement := engine.Player(anyContainer.anyOfPlayer_ZoneItem.Player).player
			referencedDataStatus := ReferencedDataUnchanged
			_, _, hasUpdatedDownstream := engine.assemblePlayer(referencedElement.ID, check, config)
			if hasUpdatedDownstream {
				referencedDataStatus = ReferencedDataModified
			}
			return &anyOfPlayer_ZoneItemReference{OperationKindUpdate, int(referencedElement.ID), ElementKindPlayer, referencedDataStatus, referencedElement.Path}, true, referencedDataStatus == ReferencedDataModified
		} else if anyContainer.anyOfPlayer_ZoneItem.ElementKind == ElementKindZoneItem {
			if check == nil {
				check = newRecursionCheck()
			}
			referencedElement := engine.ZoneItem(anyContainer.anyOfPlayer_ZoneItem.ZoneItem).zoneItem
			referencedDataStatus := ReferencedDataUnchanged
			_, _, hasUpdatedDownstream := engine.assembleZoneItem(referencedElement.ID, check, config)
			if hasUpdatedDownstream {
				referencedDataStatus = ReferencedDataModified
			}
			return &anyOfPlayer_ZoneItemReference{OperationKindUpdate, int(referencedElement.ID), ElementKindZoneItem, referencedDataStatus, referencedElement.Path}, true, referencedDataStatus == ReferencedDataModified
		}
	}

	// ref was definitely removed
	if statePlayer.Target != 0 && (playerIsInPatch && patchPlayer.Target == 0) {
		ref := engine.playerTargetRef(statePlayer.Target)
		anyContainer := engine.anyOfPlayer_ZoneItem(ref.playerTargetRef.ReferencedElementID)
		if anyContainer.anyOfPlayer_ZoneItem.ElementKind == ElementKindPlayer {
			if check == nil {
				check = newRecursionCheck()
			}
			referencedElement := engine.Player(anyContainer.anyOfPlayer_ZoneItem.Player).player
			referencedDataStatus := ReferencedDataUnchanged
			if _, _, hasUpdatedDownstream := engine.assemblePlayer(referencedElement.ID, check, config); hasUpdatedDownstream {
				referencedDataStatus = ReferencedDataModified
			}
			return &anyOfPlayer_ZoneItemReference{OperationKindDelete, int(referencedElement.ID), ElementKindPlayer, referencedDataStatus, referencedElement.Path}, true, referencedDataStatus == ReferencedDataModified
		} else if anyContainer.anyOfPlayer_ZoneItem.ElementKind == ElementKindZoneItem {
			if check == nil {
				check = newRecursionCheck()
			}
			referencedElement := engine.ZoneItem(anyContainer.anyOfPlayer_ZoneItem.ZoneItem).zoneItem
			referencedDataStatus := ReferencedDataUnchanged
			if _, _, hasUpdatedDownstream := engine.assembleZoneItem(referencedElement.ID, check, config); hasUpdatedDownstream {
				referencedDataStatus = ReferencedDataModified
			}
			return &anyOfPlayer_ZoneItemReference{OperationKindDelete, int(referencedElement.ID), ElementKindZoneItem, referencedDataStatus, referencedElement.Path}, true, referencedDataStatus == ReferencedDataModified
		}
	}

	// immediate replacement of refs
	if statePlayer.Target != 0 && (playerIsInPatch && patchPlayer.Target != 0) {
		if statePlayer.Target != patchPlayer.Target {
			ref := engine.playerTargetRef(patchPlayer.Target)
			anyContainer := engine.anyOfPlayer_ZoneItem(ref.playerTargetRef.ReferencedElementID)
			if anyContainer.anyOfPlayer_ZoneItem.ElementKind == ElementKindPlayer {
				if check == nil {
					check = newRecursionCheck()
				}
				referencedElement := engine.Player(anyContainer.anyOfPlayer_ZoneItem.Player).player
				referencedDataStatus := ReferencedDataUnchanged
				_, _, hasUpdatedDownstream := engine.assemblePlayer(referencedElement.ID, check, config)
				if hasUpdatedDownstream {
					referencedDataStatus = ReferencedDataModified
				}
				return &anyOfPlayer_ZoneItemReference{OperationKindUpdate, int(referencedElement.ID), ElementKindPlayer, referencedDataStatus, referencedElement.Path}, true, referencedDataStatus == ReferencedDataModified
			} else if anyContainer.anyOfPlayer_ZoneItem.ElementKind == ElementKindZoneItem {
				if check == nil {
					check = newRecursionCheck()
				}
				referencedElement := engine.ZoneItem(anyContainer.anyOfPlayer_ZoneItem.ZoneItem).zoneItem
				referencedDataStatus := ReferencedDataUnchanged
				_, _, hasUpdatedDownstream := engine.assembleZoneItem(referencedElement.ID, check, config)
				if hasUpdatedDownstream {
					referencedDataStatus = ReferencedDataModified
				}
				return &anyOfPlayer_ZoneItemReference{OperationKindUpdate, int(referencedElement.ID), ElementKindZoneItem, referencedDataStatus, referencedElement.Path}, true, referencedDataStatus == ReferencedDataModified
			}
		}
	}

	// element got updated - OperationKindUpdate
	if statePlayer.Target != 0 {
		ref := engine.playerTargetRef(statePlayer.Target)
		anyContainer := engine.anyOfPlayer_ZoneItem(ref.playerTargetRef.ReferencedElementID)
		if anyContainer.anyOfPlayer_ZoneItem.ElementKind == ElementKindPlayer {
			if check == nil {
				check = newRecursionCheck()
			}
			referencedElement := engine.Player(anyContainer.anyOfPlayer_ZoneItem.Player).player
			if _, _, hasUpdatedDownstream := engine.assemblePlayer(anyContainer.anyOfPlayer_ZoneItem.Player, check, config); hasUpdatedDownstream {
				return &anyOfPlayer_ZoneItemReference{OperationKindUnchanged, int(anyContainer.anyOfPlayer_ZoneItem.Player), ElementKindPlayer, ReferencedDataModified, referencedElement.Path}, true, true
			}
		} else if anyContainer.anyOfPlayer_ZoneItem.ElementKind == ElementKindZoneItem {
			if check == nil {
				check = newRecursionCheck()
			}
			referencedElement := engine.ZoneItem(anyContainer.anyOfPlayer_ZoneItem.ZoneItem).zoneItem
			if _, _, hasUpdatedDownstream := engine.assembleZoneItem(anyContainer.anyOfPlayer_ZoneItem.ZoneItem, check, config); hasUpdatedDownstream {
				return &anyOfPlayer_ZoneItemReference{OperationKindUnchanged, int(anyContainer.anyOfPlayer_ZoneItem.ZoneItem), ElementKindZoneItem, ReferencedDataModified, referencedElement.Path}, true, true
			}
		}
	}

	return nil, false, false
}

func (engine *Engine) assembleItemBoundToRef(itemID ItemID, check *recursionCheck, config assembleConfig) (*playerReference, bool, bool) {
	stateItem := engine.State.Item[itemID]
	patchItem, itemIsInPatch := engine.Patch.Item[itemID]

	// ref not set at all
	if stateItem.BoundTo == 0 && (!itemIsInPatch || patchItem.BoundTo == 0) {
		return nil, false, false
	}

	// force include
	if config.forceInclude {
		var ref ItemBoundToRef
		if patchItem.ID == 0 {
			ref = engine.itemBoundToRef(stateItem.BoundTo)
		} else {
			ref = engine.itemBoundToRef(patchItem.BoundTo)
		}
		referencedElement := engine.Player(ref.itemBoundToRef.ReferencedElementID).player
		return &playerReference{ref.itemBoundToRef.OperationKind, referencedElement.ID, ElementKindPlayer, ReferencedDataUnchanged, referencedElement.Path}, true, false
	}

	// ref was definitely created
	if stateItem.BoundTo == 0 && (itemIsInPatch && patchItem.BoundTo != 0) {
		ref := engine.itemBoundToRef(patchItem.BoundTo)
		if check == nil {
			check = newRecursionCheck()
		}
		referencedElement := engine.Player(ref.itemBoundToRef.ReferencedElementID).player
		referencedDataStatus := ReferencedDataUnchanged
		_, _, hasUpdatedDownstream := engine.assemblePlayer(referencedElement.ID, check, config)
		if hasUpdatedDownstream {
			referencedDataStatus = ReferencedDataModified
		}
		return &playerReference{OperationKindUpdate, referencedElement.ID, ElementKindPlayer, referencedDataStatus, referencedElement.Path}, true, referencedDataStatus == ReferencedDataModified
	}

	// ref was definitely removed
	if stateItem.BoundTo != 0 && (itemIsInPatch && patchItem.BoundTo == 0) {
		ref := engine.itemBoundToRef(stateItem.BoundTo)
		if check == nil {
			check = newRecursionCheck()
		}
		referencedElement := engine.Player(ref.itemBoundToRef.ReferencedElementID).player
		referencedDataStatus := ReferencedDataUnchanged
		if _, _, hasUpdatedDownstream := engine.assemblePlayer(referencedElement.ID, check, config); hasUpdatedDownstream {
			referencedDataStatus = ReferencedDataModified
		}
		return &playerReference{OperationKindDelete, referencedElement.ID, ElementKindPlayer, referencedDataStatus, referencedElement.Path}, true, referencedDataStatus == ReferencedDataModified
	}

	// immediate replacement of refs
	if stateItem.BoundTo != 0 && (itemIsInPatch && patchItem.BoundTo != 0) {
		if stateItem.BoundTo != patchItem.BoundTo {
			ref := engine.itemBoundToRef(patchItem.BoundTo)
			if check == nil {
				check = newRecursionCheck()
			}
			referencedElement := engine.Player(ref.itemBoundToRef.ReferencedElementID).player
			referencedDataStatus := ReferencedDataUnchanged
			_, _, hasUpdatedDownstream := engine.assemblePlayer(referencedElement.ID, check, config)
			if hasUpdatedDownstream {
				referencedDataStatus = ReferencedDataModified
			}
			return &playerReference{OperationKindUpdate, referencedElement.ID, ElementKindPlayer, referencedDataStatus, referencedElement.Path}, true, referencedDataStatus == ReferencedDataModified
		}
	}

	// OperationKindUpdate element got updated
	if stateItem.BoundTo != 0 {
		ref := engine.itemBoundToRef(stateItem.BoundTo)
		if check == nil {
			check = newRecursionCheck()
		}
		referencedElement := engine.Player(ref.itemBoundToRef.ReferencedElementID).player
		if _, _, hasUpdatedDownstream := engine.assemblePlayer(ref.ID(), check, config); hasUpdatedDownstream {
			return &playerReference{OperationKindUnchanged, ref.ID(), ElementKindPlayer, ReferencedDataModified, referencedElement.Path}, true, true
		}
	}

	return nil, false, false
}

func (engine *Engine) assemblePlayerTargetedByRef(refID PlayerTargetedByRefID, check *recursionCheck, config assembleConfig) (anyOfPlayer_ZoneItemReference, bool, bool) {
	if config.forceInclude {
		ref := engine.playerTargetedByRef(refID).playerTargetedByRef
		anyContainer := engine.anyOfPlayer_ZoneItem(ref.ReferencedElementID)
		if anyContainer.anyOfPlayer_ZoneItem.ElementKind == ElementKindPlayer {
			referencedElement := engine.Player(anyContainer.anyOfPlayer_ZoneItem.Player).player
			return anyOfPlayer_ZoneItemReference{ref.OperationKind, int(referencedElement.ID), ElementKindPlayer, ReferencedDataUnchanged, referencedElement.Path}, true, false
		} else if anyContainer.anyOfPlayer_ZoneItem.ElementKind == ElementKindZoneItem {
			referencedElement := engine.ZoneItem(anyContainer.anyOfPlayer_ZoneItem.ZoneItem).zoneItem
			return anyOfPlayer_ZoneItemReference{ref.OperationKind, int(referencedElement.ID), ElementKindZoneItem, ReferencedDataUnchanged, referencedElement.Path}, true, false
		}
	}

	if patchRef, hasUpdated := engine.Patch.PlayerTargetedByRef[refID]; hasUpdated {
		anyContainer := engine.anyOfPlayer_ZoneItem(patchRef.ReferencedElementID)
		if anyContainer.anyOfPlayer_ZoneItem.ElementKind == ElementKindPlayer {
			if check == nil {
				check = newRecursionCheck()
			}
			referencedElement := engine.Player(anyContainer.anyOfPlayer_ZoneItem.Player).player
			_, _, hasUpdatedDownstream := engine.assemblePlayer(referencedElement.ID, check, config)
			referencedDataStatus := ReferencedDataUnchanged
			if hasUpdatedDownstream {
				referencedDataStatus = ReferencedDataModified
			}
			return anyOfPlayer_ZoneItemReference{patchRef.OperationKind, int(referencedElement.ID), ElementKindPlayer, referencedDataStatus, referencedElement.Path}, true, patchRef.OperationKind == OperationKindUpdate || referencedDataStatus == ReferencedDataModified
		} else if anyContainer.anyOfPlayer_ZoneItem.ElementKind == ElementKindZoneItem {
			if check == nil {
				check = newRecursionCheck()
			}
			referencedElement := engine.ZoneItem(anyContainer.anyOfPlayer_ZoneItem.ZoneItem).zoneItem
			_, _, hasUpdatedDownstream := engine.assembleZoneItem(referencedElement.ID, check, config)
			referencedDataStatus := ReferencedDataUnchanged
			if hasUpdatedDownstream {
				referencedDataStatus = ReferencedDataModified
			}
			return anyOfPlayer_ZoneItemReference{patchRef.OperationKind, int(referencedElement.ID), ElementKindZoneItem, referencedDataStatus, referencedElement.Path}, true, patchRef.OperationKind == OperationKindUpdate || referencedDataStatus == ReferencedDataModified
		}
	}

	ref := engine.playerTargetedByRef(refID).playerTargetedByRef
	if check == nil {
		check = newRecursionCheck()
	}
	anyContainer := engine.anyOfPlayer_ZoneItem(ref.ReferencedElementID)
	if anyContainer.anyOfPlayer_ZoneItem.ElementKind == ElementKindPlayer {
		referencedElement := engine.Player(anyContainer.anyOfPlayer_ZoneItem.Player).player
		if _, _, hasUpdatedDownstream := engine.assemblePlayer(anyContainer.anyOfPlayer_ZoneItem.Player, check, config); hasUpdatedDownstream {
			return anyOfPlayer_ZoneItemReference{OperationKindUnchanged, int(anyContainer.anyOfPlayer_ZoneItem.Player), ElementKindPlayer, ReferencedDataModified, referencedElement.Path}, true, true
		}
	} else if anyContainer.anyOfPlayer_ZoneItem.ElementKind == ElementKindZoneItem {
		referencedElement := engine.ZoneItem(anyContainer.anyOfPlayer_ZoneItem.ZoneItem).zoneItem
		if _, _, hasUpdatedDownstream := engine.assembleZoneItem(anyContainer.anyOfPlayer_ZoneItem.ZoneItem, check, config); hasUpdatedDownstream {
			return anyOfPlayer_ZoneItemReference{OperationKindUnchanged, int(anyContainer.anyOfPlayer_ZoneItem.ZoneItem), ElementKindZoneItem, ReferencedDataModified, referencedElement.Path}, true, true
		}
	}

	return anyOfPlayer_ZoneItemReference{}, false, false
}

func (engine *Engine) assemblePlayerGuildMemberRef(refID PlayerGuildMemberRefID, check *recursionCheck, config assembleConfig) (playerReference, bool, bool) {
	if config.forceInclude {
		ref := engine.playerGuildMemberRef(refID).playerGuildMemberRef
		referencedElement := engine.Player(ref.ReferencedElementID).player
		return playerReference{ref.OperationKind, ref.ReferencedElementID, ElementKindPlayer, ReferencedDataUnchanged, referencedElement.Path}, true, false
	}

	if patchRef, hasUpdated := engine.Patch.PlayerGuildMemberRef[refID]; hasUpdated {
		if check == nil {
			check = newRecursionCheck()
		}
		referencedElement := engine.Player(patchRef.ReferencedElementID).player
		_, _, hasUpdatedDownstream := engine.assemblePlayer(referencedElement.ID, check, config)
		referencedDataStatus := ReferencedDataUnchanged
		if hasUpdatedDownstream {
			referencedDataStatus = ReferencedDataModified
		}
		return playerReference{patchRef.OperationKind, patchRef.ReferencedElementID, ElementKindPlayer, referencedDataStatus, referencedElement.Path}, true, patchRef.OperationKind == OperationKindUpdate || referencedDataStatus == ReferencedDataModified
	}

	ref := engine.playerGuildMemberRef(refID).playerGuildMemberRef
	if check == nil {
		check = newRecursionCheck()
	}
	referencedElement := engine.Player(ref.ReferencedElementID).player
	if _, _, hasUpdatedDownstream := engine.assemblePlayer(ref.ReferencedElementID, check, config); hasUpdatedDownstream {
		return playerReference{OperationKindUnchanged, ref.ReferencedElementID, ElementKindPlayer, ReferencedDataModified, referencedElement.Path}, true, true
	}

	return playerReference{}, false, false
}

func (engine *Engine) assemblePlayerEquipmentSetRef(refID PlayerEquipmentSetRefID, check *recursionCheck, config assembleConfig) (equipmentSetReference, bool, bool) {
	if config.forceInclude {
		ref := engine.playerEquipmentSetRef(refID).playerEquipmentSetRef
		referencedElement := engine.EquipmentSet(ref.ReferencedElementID).equipmentSet
		return equipmentSetReference{ref.OperationKind, ref.ReferencedElementID, ElementKindEquipmentSet, ReferencedDataUnchanged, referencedElement.Path}, true, false
	}

	if patchRef, hasUpdated := engine.Patch.PlayerEquipmentSetRef[refID]; hasUpdated {
		if check == nil {
			check = newRecursionCheck()
		}
		referencedElement := engine.EquipmentSet(patchRef.ReferencedElementID).equipmentSet
		_, _, hasUpdatedDownstream := engine.assembleEquipmentSet(referencedElement.ID, check, config)
		referencedDataStatus := ReferencedDataUnchanged
		if hasUpdatedDownstream {
			referencedDataStatus = ReferencedDataModified
		}
		return equipmentSetReference{patchRef.OperationKind, patchRef.ReferencedElementID, ElementKindEquipmentSet, referencedDataStatus, referencedElement.Path}, true, patchRef.OperationKind == OperationKindUpdate || referencedDataStatus == ReferencedDataModified
	}

	ref := engine.playerEquipmentSetRef(refID).playerEquipmentSetRef
	if check == nil {
		check = newRecursionCheck()
	}
	referencedElement := engine.EquipmentSet(ref.ReferencedElementID).equipmentSet
	if _, _, hasUpdatedDownstream := engine.assembleEquipmentSet(ref.ReferencedElementID, check, config); hasUpdatedDownstream {
		return equipmentSetReference{OperationKindUnchanged, ref.ReferencedElementID, ElementKindEquipmentSet, ReferencedDataModified, referencedElement.Path}, true, true
	}

	return equipmentSetReference{}, false, false
}

func (engine *Engine) assembleEquipmentSetEquipmentRef(refID EquipmentSetEquipmentRefID, check *recursionCheck, config assembleConfig) (itemReference, bool, bool) {
	if config.forceInclude {
		ref := engine.equipmentSetEquipmentRef(refID).equipmentSetEquipmentRef
		referencedElement := engine.Item(ref.ReferencedElementID).item
		return itemReference{ref.OperationKind, ref.ReferencedElementID, ElementKindItem, ReferencedDataUnchanged, referencedElement.Path}, true, false
	}

	if patchRef, hasUpdated := engine.Patch.EquipmentSetEquipmentRef[refID]; hasUpdated {
		if check == nil {
			check = newRecursionCheck()
		}
		referencedElement := engine.Item(patchRef.ReferencedElementID).item
		_, _, hasUpdatedDownstream := engine.assembleItem(referencedElement.ID, check, config)
		referencedDataStatus := ReferencedDataUnchanged
		if hasUpdatedDownstream {
			referencedDataStatus = ReferencedDataModified
		}
		return itemReference{patchRef.OperationKind, patchRef.ReferencedElementID, ElementKindItem, referencedDataStatus, referencedElement.Path}, true, patchRef.OperationKind == OperationKindUpdate || referencedDataStatus == ReferencedDataModified
	}

	ref := engine.equipmentSetEquipmentRef(refID).equipmentSetEquipmentRef
	if check == nil {
		check = newRecursionCheck()
	}
	referencedElement := engine.Item(ref.ReferencedElementID).item
	if _, _, hasUpdatedDownstream := engine.assembleItem(ref.ReferencedElementID, check, config); hasUpdatedDownstream {
		return itemReference{OperationKindUnchanged, ref.ReferencedElementID, ElementKindItem, ReferencedDataModified, referencedElement.Path}, true, true
	}

	return itemReference{}, false, false
}

func (engine *Engine) assembleTree(assembleEntireTree bool) Tree {

	for key := range engine.assembleCache.equipmentSet {
		delete(engine.assembleCache.equipmentSet, key)
	}
	for key := range engine.assembleCache.gearScore {
		delete(engine.assembleCache.gearScore, key)
	}
	for key := range engine.assembleCache.item {
		delete(engine.assembleCache.item, key)
	}
	for key := range engine.assembleCache.player {
		delete(engine.assembleCache.player, key)
	}
	for key := range engine.assembleCache.position {
		delete(engine.assembleCache.position, key)
	}
	for key := range engine.assembleCache.zone {
		delete(engine.assembleCache.zone, key)
	}
	for key := range engine.assembleCache.zoneItem {
		delete(engine.assembleCache.zoneItem, key)
	}

	for key := range engine.Tree.EquipmentSet {
		delete(engine.Tree.EquipmentSet, key)
	}
	for key := range engine.Tree.GearScore {
		delete(engine.Tree.GearScore, key)
	}
	for key := range engine.Tree.Item {
		delete(engine.Tree.Item, key)
	}
	for key := range engine.Tree.Player {
		delete(engine.Tree.Player, key)
	}
	for key := range engine.Tree.Position {
		delete(engine.Tree.Position, key)
	}
	for key := range engine.Tree.Zone {
		delete(engine.Tree.Zone, key)
	}
	for key := range engine.Tree.ZoneItem {
		delete(engine.Tree.ZoneItem, key)
	}

	config := assembleConfig{
		forceInclude: assembleEntireTree,
	}

	for _, equipmentSetData := range engine.Patch.EquipmentSet {
		if !equipmentSetData.HasParent {
			equipmentSet, include, _ := engine.assembleEquipmentSet(equipmentSetData.ID, nil, config)
			if include {
				engine.Tree.EquipmentSet[equipmentSetData.ID] = equipmentSet
			}
		}
	}
	for _, gearScoreData := range engine.Patch.GearScore {
		if !gearScoreData.HasParent {
			gearScore, include, _ := engine.assembleGearScore(gearScoreData.ID, nil, config)
			if include {
				engine.Tree.GearScore[gearScoreData.ID] = gearScore
			}
		}
	}
	for _, itemData := range engine.Patch.Item {
		if !itemData.HasParent {
			item, include, _ := engine.assembleItem(itemData.ID, nil, config)
			if include {
				engine.Tree.Item[itemData.ID] = item
			}
		}
	}
	for _, playerData := range engine.Patch.Player {
		if !playerData.HasParent {
			player, include, _ := engine.assemblePlayer(playerData.ID, nil, config)
			if include {
				engine.Tree.Player[playerData.ID] = player
			}
		}
	}
	for _, positionData := range engine.Patch.Position {
		if !positionData.HasParent {
			position, include, _ := engine.assemblePosition(positionData.ID, nil, config)
			if include {
				engine.Tree.Position[positionData.ID] = position
			}
		}
	}
	for _, zoneData := range engine.Patch.Zone {
		if !zoneData.HasParent {
			zone, include, _ := engine.assembleZone(zoneData.ID, nil, config)
			if include {
				engine.Tree.Zone[zoneData.ID] = zone
			}
		}
	}
	for _, zoneItemData := range engine.Patch.ZoneItem {
		if !zoneItemData.HasParent {
			zoneItem, include, _ := engine.assembleZoneItem(zoneItemData.ID, nil, config)
			if include {
				engine.Tree.ZoneItem[zoneItemData.ID] = zoneItem
			}
		}
	}

	for _, equipmentSetData := range engine.State.EquipmentSet {
		if !equipmentSetData.HasParent {
			if _, ok := engine.Tree.EquipmentSet[equipmentSetData.ID]; !ok {
				equipmentSet, include, _ := engine.assembleEquipmentSet(equipmentSetData.ID, nil, config)
				if include {
					engine.Tree.EquipmentSet[equipmentSetData.ID] = equipmentSet
				}
			}
		}
	}
	for _, gearScoreData := range engine.State.GearScore {
		if !gearScoreData.HasParent {
			if _, ok := engine.Tree.GearScore[gearScoreData.ID]; !ok {
				gearScore, include, _ := engine.assembleGearScore(gearScoreData.ID, nil, config)
				if include {
					engine.Tree.GearScore[gearScoreData.ID] = gearScore
				}
			}
		}
	}
	for _, itemData := range engine.State.Item {
		if !itemData.HasParent {
			if _, ok := engine.Tree.Item[itemData.ID]; !ok {
				item, include, _ := engine.assembleItem(itemData.ID, nil, config)
				if include {
					engine.Tree.Item[itemData.ID] = item
				}
			}
		}
	}
	for _, playerData := range engine.State.Player {
		if !playerData.HasParent {
			if _, ok := engine.Tree.Player[playerData.ID]; !ok {
				player, include, _ := engine.assemblePlayer(playerData.ID, nil, config)
				if include {
					engine.Tree.Player[playerData.ID] = player
				}
			}
		}
	}
	for _, positionData := range engine.State.Position {
		if !positionData.HasParent {
			if _, ok := engine.Tree.Position[positionData.ID]; !ok {
				position, include, _ := engine.assemblePosition(positionData.ID, nil, config)
				if include {
					engine.Tree.Position[positionData.ID] = position
				}
			}
		}
	}
	for _, zoneData := range engine.State.Zone {
		if !zoneData.HasParent {
			if _, ok := engine.Tree.Zone[zoneData.ID]; !ok {
				zone, include, _ := engine.assembleZone(zoneData.ID, nil, config)
				if include {
					engine.Tree.Zone[zoneData.ID] = zone
				}
			}
		}
	}
	for _, zoneItemData := range engine.State.ZoneItem {
		if !zoneItemData.HasParent {
			if _, ok := engine.Tree.ZoneItem[zoneItemData.ID]; !ok {
				zoneItem, include, _ := engine.assembleZoneItem(zoneItemData.ID, nil, config)
				if include {
					engine.Tree.ZoneItem[zoneItemData.ID] = zoneItem
				}
			}
		}
	}

	return engine.Tree
}
