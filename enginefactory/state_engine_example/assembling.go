package state

type assembleConfig struct {
	forceInclude bool // include everything, regardless of update status
}

func (engine *Engine) assembleGearScore(gearScoreID GearScoreID, check *recursionCheck, config assembleConfig) (GearScore, bool, bool) {
	if check != nil {
		if alreadyExists := check.gearScore[gearScoreID]; alreadyExists {
			return GearScore{}, false, true
		} else {
			check.gearScore[gearScoreID] = true
		}
	}

	gearScoreData, hasUpdated := engine.Patch.GearScore[gearScoreID]
	if !hasUpdated {
		gearScoreData, _ = engine.State.GearScore[gearScoreID]
	}

	var gearScore GearScore

	gearScore.ID = gearScoreData.ID
	gearScore.OperationKind = gearScoreData.OperationKind
	gearScore.Level = gearScoreData.Level
	gearScore.Score = gearScoreData.Score
	return gearScore, hasUpdated || config.forceInclude, hasUpdated
}

func (engine *Engine) assemblePosition(positionID PositionID, check *recursionCheck, config assembleConfig) (Position, bool, bool) {
	if check != nil {
		if alreadyExists := check.position[positionID]; alreadyExists {
			return Position{}, false, true
		} else {
			check.position[positionID] = true
		}
	}

	positionData, hasUpdated := engine.Patch.Position[positionID]
	if !hasUpdated {
		positionData, _ = engine.State.Position[positionID]
	}

	var position Position

	position.ID = positionData.ID
	position.OperationKind = positionData.OperationKind
	position.X = positionData.X
	position.Y = positionData.Y
	return position, hasUpdated || config.forceInclude, hasUpdated
}

func (engine *Engine) assembleEquipmentSet(equipmentSetID EquipmentSetID, check *recursionCheck, config assembleConfig) (EquipmentSet, bool, bool) {
	if check != nil {
		if alreadyExists := check.equipmentSet[equipmentSetID]; alreadyExists {
			return EquipmentSet{}, false, true
		} else {
			check.equipmentSet[equipmentSetID] = true
		}
	}

	equipmentSetData, hasUpdated := engine.Patch.EquipmentSet[equipmentSetID]
	if !hasUpdated {
		equipmentSetData = engine.State.EquipmentSet[equipmentSetID]
	}

	var equipmentSet EquipmentSet

	for _, refID := range mergeEquipmentSetEquipmentRefIDs(engine.State.EquipmentSet[equipmentSetID].Equipment, engine.Patch.EquipmentSet[equipmentSetID].Equipment) {
		if ref, include, refHasUpdated := engine.assembleEquipmentSetEquipmentRef(refID, check, config); include {
			if refHasUpdated {
				hasUpdated = true
			}
			equipmentSet.Equipment = append(equipmentSet.Equipment, ref)
		}
	}

	equipmentSet.ID = equipmentSetData.ID
	equipmentSet.OperationKind = equipmentSetData.OperationKind
	equipmentSet.Name = equipmentSetData.Name
	return equipmentSet, hasUpdated || config.forceInclude, hasUpdated
}

func (engine *Engine) assembleItem(itemID ItemID, check *recursionCheck, config assembleConfig) (Item, bool, bool) {
	if check != nil {
		if alreadyExists := check.item[itemID]; alreadyExists {
			return Item{}, false, true
		} else {
			check.item[itemID] = true
		}
	}

	itemData, hasUpdated := engine.Patch.Item[itemID]
	if !hasUpdated {
		itemData = engine.State.Item[itemID]
	}

	var item Item

	if ref, include, refHasUpdated := engine.assembleItemBoundToRef(itemID, check, config); include {
		if refHasUpdated {
			hasUpdated = true
		}
		item.BoundTo = ref
	}

	if treeGearScore, include, gearScoreHasUpdated := engine.assembleGearScore(itemData.GearScore, check, config); include {
		if gearScoreHasUpdated {
			hasUpdated = true
		}
		item.GearScore = &treeGearScore
	}

	if anyContainer := engine.anyOfPlayerPosition(itemData.Origin).anyOfPlayerPosition; anyContainer.ElementKind == ElementKindPlayer {
		playerID := anyContainer.Player
		if treePlayer, include, playerHasUpdated := engine.assemblePlayer(playerID, check, config); include {
			if playerHasUpdated {
				hasUpdated = true
			}
			item.Origin = &treePlayer
		}
	} else if anyContainer.ElementKind == ElementKindPosition {
		positionID := anyContainer.Position
		if treePosition, include, positionHasUpdated := engine.assemblePosition(positionID, check, config); include {
			if positionHasUpdated {
				hasUpdated = true
			}
			item.Origin = &treePosition
		}
	}

	if treeGearScore, include, gearScoreHasUpdated := engine.assembleGearScore(itemData.GearScore, check, config); include {
		if gearScoreHasUpdated {
			hasUpdated = true
		}
		item.GearScore = &treeGearScore
	}

	item.ID = itemData.ID
	item.OperationKind = itemData.OperationKind
	item.Name = itemData.Name
	return item, hasUpdated || config.forceInclude, hasUpdated
}

func (engine *Engine) assembleZoneItem(zoneItemID ZoneItemID, check *recursionCheck, config assembleConfig) (ZoneItem, bool, bool) {
	if check != nil {
		if alreadyExists := check.zoneItem[zoneItemID]; alreadyExists {
			return ZoneItem{}, false, true
		} else {
			check.zoneItem[zoneItemID] = true
		}
	}

	zoneItemData, hasUpdated := engine.Patch.ZoneItem[zoneItemID]
	if !hasUpdated {
		zoneItemData = engine.State.ZoneItem[zoneItemID]
	}

	var zoneItem ZoneItem

	if treeItem, include, itemHasUpdated := engine.assembleItem(zoneItemData.Item, check, config); include {
		if itemHasUpdated {
			hasUpdated = true
		}
		zoneItem.Item = &treeItem
	}

	if treePosition, include, positionHasUpdated := engine.assemblePosition(zoneItemData.Position, check, config); include {
		if positionHasUpdated {
			hasUpdated = true
		}
		zoneItem.Position = &treePosition
	}

	zoneItem.ID = zoneItemData.ID
	zoneItem.OperationKind = zoneItemData.OperationKind
	return zoneItem, hasUpdated || config.forceInclude, hasUpdated
}

func (engine *Engine) assemblePlayer(playerID PlayerID, check *recursionCheck, config assembleConfig) (Player, bool, bool) {
	if check != nil {
		if alreadyExists := check.player[playerID]; alreadyExists {
			return Player{}, false, true
		} else {
			check.player[playerID] = true
		}
	}

	playerData, hasUpdated := engine.Patch.Player[playerID]
	if !hasUpdated {
		playerData = engine.State.Player[playerID]
	}

	var player Player

	for _, refID := range mergePlayerEquipmentSetRefIDs(engine.State.Player[playerID].EquipmentSets, engine.Patch.Player[playerID].EquipmentSets) {
		if ref, include, refHasUpdated := engine.assemblePlayerEquipmentSetRef(refID, check, config); include {
			if refHasUpdated {
				hasUpdated = true
			}
			player.EquipmentSets = append(player.EquipmentSets, ref)
		}
	}

	if treeGearScore, include, gearScoreHasUpdated := engine.assembleGearScore(playerData.GearScore, check, config); include {
		if gearScoreHasUpdated {
			hasUpdated = true
		}
		player.GearScore = &treeGearScore
	}

	for _, refID := range mergePlayerGuildMemberRefIDs(engine.State.Player[playerID].GuildMembers, engine.Patch.Player[playerID].GuildMembers) {
		if ref, include, refHasUpdated := engine.assemblePlayerGuildMemberRef(refID, check, config); include {
			if refHasUpdated {
				hasUpdated = true
			}
			player.GuildMembers = append(player.GuildMembers, ref)
		}
	}

	for _, itemID := range mergeItemIDs(engine.State.Player[playerData.ID].Items, engine.Patch.Player[playerData.ID].Items) {
		if treeItem, include, itemHasUpdated := engine.assembleItem(itemID, check, config); include {
			if itemHasUpdated {
				hasUpdated = true
			}
			player.Items = append(player.Items, treeItem)
		}
	}

	if treePosition, include, positionHasUpdated := engine.assemblePosition(playerData.Position, check, config); include {
		if positionHasUpdated {
			hasUpdated = true
		}
		player.Position = &treePosition
	}

	if ref, include, refHasUpdated := engine.assemblePlayerTargetRef(playerID, check, config); include {
		if refHasUpdated {
			hasUpdated = true
		}
		player.Target = ref
	}

	player.ID = playerData.ID
	player.OperationKind = playerData.OperationKind
	return player, hasUpdated || config.forceInclude, hasUpdated
}

func (engine *Engine) assembleZone(zoneID ZoneID, check *recursionCheck, config assembleConfig) (Zone, bool, bool) {
	if check != nil {
		if alreadyExists := check.zone[zoneID]; alreadyExists {
			return Zone{}, false, true
		} else {
			check.zone[zoneID] = true
		}
	}

	zoneData, hasUpdated := engine.Patch.Zone[zoneID]
	if !hasUpdated {
		zoneData = engine.State.Zone[zoneID]
	}

	var zone Zone

	for _, anyID := range mergeAnyOfItemPlayerZoneItemIDs(engine.State.Zone[zoneData.ID].Interactables, engine.Patch.Zone[zoneData.ID].Interactables) {
		if anyContainer := engine.anyOfItemPlayerZoneItem(anyID).anyOfItemPlayerZoneItem; anyContainer.ElementKind == ElementKindItem {
			itemID := anyContainer.Item
			if treeItem, include, itemHasUpdated := engine.assembleItem(itemID, check, config); include {
				if itemHasUpdated {
					hasUpdated = true
				}
				zone.Interactables = append(zone.Interactables, treeItem)
			}
		} else if anyContainer.ElementKind == ElementKindPlayer {
			playerID := anyContainer.Player
			if treePlayer, include, playerHasUpdated := engine.assemblePlayer(playerID, check, config); include {
				if playerHasUpdated {
					hasUpdated = true
				}
				zone.Interactables = append(zone.Interactables, treePlayer)
			}
		} else if anyContainer.ElementKind == ElementKindZoneItem {
			zoneItemID := anyContainer.ZoneItem
			if treeZoneItem, include, zoneItemHasUpdated := engine.assembleZoneItem(zoneItemID, check, config); include {
				if zoneItemHasUpdated {
					hasUpdated = true
				}
				zone.Interactables = append(zone.Interactables, treeZoneItem)
			}
		}
	}

	for _, zoneItemID := range mergeZoneItemIDs(engine.State.Zone[zoneData.ID].Items, engine.Patch.Zone[zoneData.ID].Items) {
		if treeZoneItem, include, zoneItemHasUpdated := engine.assembleZoneItem(zoneItemID, check, config); include {
			if zoneItemHasUpdated {
				hasUpdated = true
			}
			zone.Items = append(zone.Items, treeZoneItem)
		}
	}

	for _, playerID := range mergePlayerIDs(engine.State.Zone[zoneData.ID].Players, engine.Patch.Zone[zoneData.ID].Players) {
		if treePlayer, include, playerHasUpdated := engine.assemblePlayer(playerID, check, config); include {
			if playerHasUpdated {
				hasUpdated = true
			}
			zone.Players = append(zone.Players, treePlayer)
		}
	}

	zone.ID = zoneData.ID
	zone.OperationKind = zoneData.OperationKind
	zone.Tags = zoneData.Tags
	return zone, hasUpdated || config.forceInclude, hasUpdated
}

func (engine *Engine) assemblePlayerTargetRef(playerID PlayerID, check *recursionCheck, config assembleConfig) (*AnyOfPlayerZoneItemReference, bool, bool) {
	statePlayer := engine.State.Player[playerID]
	patchPlayer, playerIsInPatch := engine.Patch.Player[playerID]

	// ref not set at all
	if statePlayer.Target == 0 && (!playerIsInPatch || patchPlayer.Target == 0) {
		return nil, false, false
	}

	// force include
	if config.forceInclude {
		ref := engine.playerTargetRef(patchPlayer.Target)
		if anyContainer := engine.anyOfPlayerZoneItem(ref.playerTargetRef.ReferencedElementID); anyContainer.anyOfPlayerZoneItem.ElementKind == ElementKindPlayer {
			referencedElement := engine.Player(anyContainer.anyOfPlayerZoneItem.Player).player
			if check == nil {
				check = newRecursionCheck()
			}
			referencedDataStatus := ReferencedDataUnchanged
			if _, _, hasUpdatedDownstream := engine.assemblePlayer(referencedElement.ID, check, config); hasUpdatedDownstream {
				referencedDataStatus = ReferencedDataModified
			}
			path, _ := engine.PathTrack.player[referencedElement.ID]
			return &AnyOfPlayerZoneItemReference{ref.playerTargetRef.OperationKind, int(referencedElement.ID), ElementKindPlayer, referencedDataStatus, path.toJSONPath(), nil}, true, ref.playerTargetRef.OperationKind == OperationKindUpdate || referencedDataStatus == ReferencedDataModified
		} else if anyContainer.anyOfPlayerZoneItem.ElementKind == ElementKindZoneItem {
			referencedElement := engine.ZoneItem(anyContainer.anyOfPlayerZoneItem.ZoneItem).zoneItem
			if check == nil {
				check = newRecursionCheck()
			}
			referencedDataStatus := ReferencedDataUnchanged
			if _, _, hasUpdatedDownstream := engine.assembleZoneItem(referencedElement.ID, check, config); hasUpdatedDownstream {
				referencedDataStatus = ReferencedDataModified
			}
			path, _ := engine.PathTrack.zoneItem[referencedElement.ID]
			return &AnyOfPlayerZoneItemReference{ref.playerTargetRef.OperationKind, int(referencedElement.ID), ElementKindZoneItem, referencedDataStatus, path.toJSONPath(), nil}, true, ref.playerTargetRef.OperationKind == OperationKindUpdate || referencedDataStatus == ReferencedDataModified
		}
	}

	// ref was definitely created
	if statePlayer.Target == 0 && (playerIsInPatch && patchPlayer.Target != 0) {
		config.forceInclude = true
		ref := engine.playerTargetRef(patchPlayer.Target)
		if anyContainer := engine.anyOfPlayerZoneItem(ref.playerTargetRef.ReferencedElementID); anyContainer.anyOfPlayerZoneItem.ElementKind == ElementKindPlayer {
			referencedElement := engine.Player(anyContainer.anyOfPlayerZoneItem.Player).player
			if check == nil {
				check = newRecursionCheck()
			}
			referencedDataStatus := ReferencedDataUnchanged
			element, _, hasUpdatedDownstream := engine.assemblePlayer(referencedElement.ID, check, config)
			if hasUpdatedDownstream {
				referencedDataStatus = ReferencedDataModified
			}
			path, _ := engine.PathTrack.player[referencedElement.ID]
			return &AnyOfPlayerZoneItemReference{OperationKindUpdate, int(referencedElement.ID), ElementKindPlayer, referencedDataStatus, path.toJSONPath(), &element}, true, referencedDataStatus == ReferencedDataModified
		} else if anyContainer.anyOfPlayerZoneItem.ElementKind == ElementKindZoneItem {
			referencedElement := engine.ZoneItem(anyContainer.anyOfPlayerZoneItem.ZoneItem).zoneItem
			if check == nil {
				check = newRecursionCheck()
			}
			referencedDataStatus := ReferencedDataUnchanged
			element, _, hasUpdatedDownstream := engine.assembleZoneItem(referencedElement.ID, check, config)
			if hasUpdatedDownstream {
				referencedDataStatus = ReferencedDataModified
			}
			path, _ := engine.PathTrack.zoneItem[referencedElement.ID]
			return &AnyOfPlayerZoneItemReference{OperationKindUpdate, int(referencedElement.ID), ElementKindZoneItem, referencedDataStatus, path.toJSONPath(), &element}, true, referencedDataStatus == ReferencedDataModified
		}
	}

	// ref was definitely removed
	if statePlayer.Target != 0 && (playerIsInPatch && patchPlayer.Target == 0) {
		ref := engine.playerTargetRef(patchPlayer.Target)
		if anyContainer := engine.anyOfPlayerZoneItem(ref.playerTargetRef.ReferencedElementID); anyContainer.anyOfPlayerZoneItem.ElementKind == ElementKindPlayer {
			referencedElement := engine.Player(anyContainer.anyOfPlayerZoneItem.Player).player
			if check == nil {
				check = newRecursionCheck()
			}
			referencedDataStatus := ReferencedDataUnchanged
			if _, _, hasUpdatedDownstream := engine.assemblePlayer(referencedElement.ID, check, config); hasUpdatedDownstream {
				referencedDataStatus = ReferencedDataModified
			}
			path, _ := engine.PathTrack.player[referencedElement.ID]
			return &AnyOfPlayerZoneItemReference{OperationKindDelete, int(referencedElement.ID), ElementKindPlayer, referencedDataStatus, path.toJSONPath(), nil}, true, referencedDataStatus == ReferencedDataModified
		} else if anyContainer.anyOfPlayerZoneItem.ElementKind == ElementKindZoneItem {
			referencedElement := engine.ZoneItem(anyContainer.anyOfPlayerZoneItem.ZoneItem).zoneItem
			if check == nil {
				check = newRecursionCheck()
			}
			referencedDataStatus := ReferencedDataUnchanged
			if _, _, hasUpdatedDownstream := engine.assembleZoneItem(referencedElement.ID, check, config); hasUpdatedDownstream {
				referencedDataStatus = ReferencedDataModified
			}
			path, _ := engine.PathTrack.zoneItem[referencedElement.ID]
			return &AnyOfPlayerZoneItemReference{OperationKindDelete, int(referencedElement.ID), ElementKindZoneItem, referencedDataStatus, path.toJSONPath(), nil}, true, referencedDataStatus == ReferencedDataModified
		}
	}

	// immediate replacement of refs
	if statePlayer.Target != 0 && (playerIsInPatch && patchPlayer.Target != 0) {
		if statePlayer.Target != patchPlayer.Target {
			ref := engine.playerTargetRef(patchPlayer.Target)
			if anyContainer := engine.anyOfPlayerZoneItem(ref.playerTargetRef.ReferencedElementID); anyContainer.anyOfPlayerZoneItem.ElementKind == ElementKindPlayer {
				referencedElement := engine.Player(anyContainer.anyOfPlayerZoneItem.Player).player
				if check == nil {
					check = newRecursionCheck()
				}
				referencedDataStatus := ReferencedDataUnchanged
				if _, _, hasUpdatedDownstream := engine.assemblePlayer(referencedElement.ID, check, config); hasUpdatedDownstream {
					referencedDataStatus = ReferencedDataModified
				}
				path, _ := engine.PathTrack.player[referencedElement.ID]
				return &AnyOfPlayerZoneItemReference{OperationKindUpdate, int(referencedElement.ID), ElementKindPlayer, referencedDataStatus, path.toJSONPath(), nil}, true, referencedDataStatus == ReferencedDataModified
			} else if anyContainer.anyOfPlayerZoneItem.ElementKind == ElementKindZoneItem {
				referencedElement := engine.ZoneItem(anyContainer.anyOfPlayerZoneItem.ZoneItem).zoneItem
				if check == nil {
					check = newRecursionCheck()
				}
				referencedDataStatus := ReferencedDataUnchanged
				if _, _, hasUpdatedDownstream := engine.assembleZoneItem(referencedElement.ID, check, config); hasUpdatedDownstream {
					referencedDataStatus = ReferencedDataModified
				}
				path, _ := engine.PathTrack.zoneItem[referencedElement.ID]
				return &AnyOfPlayerZoneItemReference{OperationKindUpdate, int(referencedElement.ID), ElementKindZoneItem, referencedDataStatus, path.toJSONPath(), nil}, true, referencedDataStatus == ReferencedDataModified
			}
		}
	}

	// OperationKindUpdate element got updated
	if statePlayer.Target != 0 {
		ref := engine.playerTargetRef(patchPlayer.Target)
		if anyContainer := engine.anyOfPlayerZoneItem(ref.playerTargetRef.ReferencedElementID); anyContainer.anyOfPlayerZoneItem.ElementKind == ElementKindPlayer {
			if check == nil {
				check = newRecursionCheck()
			}
			if _, _, hasUpdatedDownstream := engine.assemblePlayer(anyContainer.anyOfPlayerZoneItem.Player, check, config); hasUpdatedDownstream {
				path, _ := engine.PathTrack.player[anyContainer.anyOfPlayerZoneItem.Player]
				return &AnyOfPlayerZoneItemReference{OperationKindUnchanged, int(anyContainer.anyOfPlayerZoneItem.Player), ElementKindPlayer, ReferencedDataModified, path.toJSONPath(), nil}, true, true
			}
		} else if anyContainer.anyOfPlayerZoneItem.ElementKind == ElementKindZoneItem {
			if check == nil {
				check = newRecursionCheck()
			}
			if _, _, hasUpdatedDownstream := engine.assembleZoneItem(anyContainer.anyOfPlayerZoneItem.ZoneItem, check, config); hasUpdatedDownstream {
				path, _ := engine.PathTrack.zoneItem[anyContainer.anyOfPlayerZoneItem.ZoneItem]
				return &AnyOfPlayerZoneItemReference{OperationKindUnchanged, int(anyContainer.anyOfPlayerZoneItem.ZoneItem), ElementKindZoneItem, ReferencedDataModified, path.toJSONPath(), nil}, true, true
			}
		}
	}

	return nil, false, false
}

func (engine *Engine) assembleItemBoundToRef(itemID ItemID, check *recursionCheck, config assembleConfig) (*PlayerReference, bool, bool) {
	stateItem := engine.State.Item[itemID]
	patchItem, itemIsInPatch := engine.Patch.Item[itemID]

	// ref not set at all
	if stateItem.BoundTo == 0 && (!itemIsInPatch || patchItem.BoundTo == 0) {
		return nil, false, false
	}

	// force include
	if config.forceInclude {
		ref := engine.itemBoundToRef(patchItem.BoundTo)
		referencedElement := engine.Player(ref.itemBoundToRef.ReferencedElementID).player
		if check == nil {
			check = newRecursionCheck()
		}
		referencedDataStatus := ReferencedDataUnchanged
		if _, _, hasUpdatedDownstream := engine.assemblePlayer(referencedElement.ID, check, config); hasUpdatedDownstream {
			referencedDataStatus = ReferencedDataModified
		}
		path, _ := engine.PathTrack.player[referencedElement.ID]
		return &PlayerReference{ref.itemBoundToRef.OperationKind, referencedElement.ID, ElementKindPlayer, referencedDataStatus, path.toJSONPath(), nil}, true, ref.itemBoundToRef.OperationKind == OperationKindUpdate || referencedDataStatus == ReferencedDataModified
	}

	// ref was definitely created
	if stateItem.BoundTo == 0 && (itemIsInPatch && patchItem.BoundTo != 0) {
		config.forceInclude = true
		ref := engine.itemBoundToRef(patchItem.BoundTo)
		referencedElement := engine.Player(ref.itemBoundToRef.ReferencedElementID).player
		if check == nil {
			check = newRecursionCheck()
		}
		referencedDataStatus := ReferencedDataUnchanged
		element, _, hasUpdatedDownstream := engine.assemblePlayer(referencedElement.ID, check, config)
		if hasUpdatedDownstream {
			referencedDataStatus = ReferencedDataModified
		}
		path, _ := engine.PathTrack.player[referencedElement.ID]
		return &PlayerReference{OperationKindUpdate, referencedElement.ID, ElementKindPlayer, referencedDataStatus, path.toJSONPath(), &element}, true, referencedDataStatus == ReferencedDataModified
	}

	// ref was definitely removed
	if stateItem.BoundTo != 0 && (itemIsInPatch && patchItem.BoundTo == 0) {
		ref := engine.itemBoundToRef(stateItem.BoundTo)
		referencedElement := engine.Player(ref.itemBoundToRef.ReferencedElementID).player
		if check == nil {
			check = newRecursionCheck()
		}
		referencedDataStatus := ReferencedDataUnchanged
		if _, _, hasUpdatedDownstream := engine.assemblePlayer(referencedElement.ID, check, config); hasUpdatedDownstream {
			referencedDataStatus = ReferencedDataModified
		}
		path, _ := engine.PathTrack.player[referencedElement.ID]
		return &PlayerReference{OperationKindDelete, referencedElement.ID, ElementKindPlayer, referencedDataStatus, path.toJSONPath(), nil}, true, referencedDataStatus == ReferencedDataModified
	}

	// immediate replacement of refs
	if stateItem.BoundTo != 0 && (itemIsInPatch && patchItem.BoundTo != 0) {
		if stateItem.BoundTo != patchItem.BoundTo {
			ref := engine.itemBoundToRef(patchItem.BoundTo)
			referencedElement := engine.Player(ref.itemBoundToRef.ReferencedElementID).player
			if check == nil {
				check = newRecursionCheck()
			}
			referencedDataStatus := ReferencedDataUnchanged
			if _, _, hasUpdatedDownstream := engine.assemblePlayer(referencedElement.ID, check, config); hasUpdatedDownstream {
				referencedDataStatus = ReferencedDataModified
			}
			path, _ := engine.PathTrack.player[referencedElement.ID]
			return &PlayerReference{OperationKindUpdate, referencedElement.ID, ElementKindPlayer, referencedDataStatus, path.toJSONPath(), nil}, true, referencedDataStatus == ReferencedDataModified
		}
	}

	// OperationKindUpdate element got updated
	if stateItem.BoundTo != 0 {
		ref := engine.itemBoundToRef(stateItem.BoundTo)
		if check == nil {
			check = newRecursionCheck()
		}
		if _, _, hasUpdatedDownstream := engine.assemblePlayer(ref.ID(), check, config); hasUpdatedDownstream {
			path, _ := engine.PathTrack.player[ref.itemBoundToRef.ReferencedElementID]
			return &PlayerReference{OperationKindUnchanged, ref.ID(), ElementKindPlayer, ReferencedDataModified, path.toJSONPath(), nil}, true, true
		}
	}

	return nil, false, false
}

func (engine *Engine) assemblePlayerTargetedByRef(refID PlayerTargetedByRefID, check *recursionCheck, config assembleConfig) (AnyOfPlayerZoneItemReference, bool, bool) {
	if config.forceInclude {
		ref := engine.playerTargetedByRef(refID).playerTargetedByRef
		if anyContainer := engine.anyOfPlayerZoneItem(ref.ReferencedElementID); anyContainer.anyOfPlayerZoneItem.ElementKind == ElementKindPlayer {
			referencedElement := engine.Player(anyContainer.anyOfPlayerZoneItem.Player).player
			if check == nil {
				check = newRecursionCheck()
			}
			referencedDataStatus := ReferencedDataUnchanged
			if _, _, hasUpdatedDownstream := engine.assemblePlayer(referencedElement.ID, check, config); hasUpdatedDownstream {
				referencedDataStatus = ReferencedDataModified
			}
			path, _ := engine.PathTrack.player[referencedElement.ID]
			return AnyOfPlayerZoneItemReference{ref.OperationKind, int(referencedElement.ID), ElementKindPlayer, referencedDataStatus, path.toJSONPath(), nil}, true, ref.OperationKind == OperationKindUpdate || referencedDataStatus == ReferencedDataModified
		} else if anyContainer.anyOfPlayerZoneItem.ElementKind == ElementKindZoneItem {
			referencedElement := engine.ZoneItem(anyContainer.anyOfPlayerZoneItem.ZoneItem).zoneItem
			if check == nil {
				check = newRecursionCheck()
			}
			referencedDataStatus := ReferencedDataUnchanged
			if _, _, hasUpdatedDownstream := engine.assembleZoneItem(referencedElement.ID, check, config); hasUpdatedDownstream {
				referencedDataStatus = ReferencedDataModified
			}
			path, _ := engine.PathTrack.zoneItem[referencedElement.ID]
			return AnyOfPlayerZoneItemReference{ref.OperationKind, int(referencedElement.ID), ElementKindZoneItem, referencedDataStatus, path.toJSONPath(), nil}, true, ref.OperationKind == OperationKindUpdate || referencedDataStatus == ReferencedDataModified
		}
	}

	if patchRef, hasUpdated := engine.Patch.PlayerTargetedByRef[refID]; hasUpdated {
		if patchRef.OperationKind == OperationKindUpdate {
			config.forceInclude = true
		}
		if anyContainer := engine.anyOfPlayerZoneItem(patchRef.ReferencedElementID); anyContainer.anyOfPlayerZoneItem.ElementKind == ElementKindPlayer {
			referencedElement := engine.Player(anyContainer.anyOfPlayerZoneItem.Player).player
			if check == nil {
				check = newRecursionCheck()
			}
			element, _, hasUpdatedDownstream := engine.assemblePlayer(referencedElement.ID, check, config)
			referencedDataStatus := ReferencedDataUnchanged
			if hasUpdatedDownstream {
				referencedDataStatus = ReferencedDataModified
			}
			var el *Player
			if patchRef.OperationKind == OperationKindUpdate {
				el = &element
			}
			path, _ := engine.PathTrack.player[referencedElement.ID]
			return AnyOfPlayerZoneItemReference{patchRef.OperationKind, int(referencedElement.ID), ElementKindPlayer, referencedDataStatus, path.toJSONPath(), el}, true, patchRef.OperationKind == OperationKindUpdate || referencedDataStatus == ReferencedDataModified
		} else if anyContainer.anyOfPlayerZoneItem.ElementKind == ElementKindZoneItem {
			referencedElement := engine.ZoneItem(anyContainer.anyOfPlayerZoneItem.ZoneItem).zoneItem
			if check == nil {
				check = newRecursionCheck()
			}
			element, _, hasUpdatedDownstream := engine.assembleZoneItem(referencedElement.ID, check, config)
			referencedDataStatus := ReferencedDataUnchanged
			if hasUpdatedDownstream {
				referencedDataStatus = ReferencedDataModified
			}
			var el *ZoneItem
			if patchRef.OperationKind == OperationKindUpdate {
				el = &element
			}
			path, _ := engine.PathTrack.zoneItem[referencedElement.ID]
			return AnyOfPlayerZoneItemReference{patchRef.OperationKind, int(referencedElement.ID), ElementKindZoneItem, referencedDataStatus, path.toJSONPath(), el}, true, patchRef.OperationKind == OperationKindUpdate || referencedDataStatus == ReferencedDataModified
		}
	}

	ref := engine.playerTargetedByRef(refID).playerTargetedByRef
	if check == nil {
		check = newRecursionCheck()
	}
	if anyContainer := engine.anyOfPlayerZoneItem(ref.ReferencedElementID); anyContainer.anyOfPlayerZoneItem.ElementKind == ElementKindPlayer {
		if _, _, hasUpdatedDownstream := engine.assemblePlayer(anyContainer.anyOfPlayerZoneItem.Player, check, config); hasUpdatedDownstream {
			path, _ := engine.PathTrack.player[anyContainer.anyOfPlayerZoneItem.Player]
			return AnyOfPlayerZoneItemReference{OperationKindUnchanged, int(anyContainer.anyOfPlayerZoneItem.Player), ElementKindPlayer, ReferencedDataModified, path.toJSONPath(), nil}, true, true
		}
	} else if anyContainer.anyOfPlayerZoneItem.ElementKind == ElementKindZoneItem {
		if _, _, hasUpdatedDownstream := engine.assembleZoneItem(anyContainer.anyOfPlayerZoneItem.ZoneItem, check, config); hasUpdatedDownstream {
			path, _ := engine.PathTrack.zoneItem[anyContainer.anyOfPlayerZoneItem.ZoneItem]
			return AnyOfPlayerZoneItemReference{OperationKindUnchanged, int(anyContainer.anyOfPlayerZoneItem.ZoneItem), ElementKindZoneItem, ReferencedDataModified, path.toJSONPath(), nil}, true, true
		}
	}

	return AnyOfPlayerZoneItemReference{}, false, false
}

func (engine *Engine) assemblePlayerGuildMemberRef(refID PlayerGuildMemberRefID, check *recursionCheck, config assembleConfig) (PlayerReference, bool, bool) {
	if config.forceInclude {
		ref := engine.playerGuildMemberRef(refID).playerGuildMemberRef
		referencedElement := engine.Player(ref.ReferencedElementID).player
		if check == nil {
			check = newRecursionCheck()
		}
		referencedDataStatus := ReferencedDataUnchanged
		if _, _, hasUpdatedDownstream := engine.assemblePlayer(referencedElement.ID, check, config); hasUpdatedDownstream {
			referencedDataStatus = ReferencedDataModified
		}
		path, _ := engine.PathTrack.player[referencedElement.ID]
		return PlayerReference{ref.OperationKind, ref.ReferencedElementID, ElementKindPlayer, referencedDataStatus, path.toJSONPath(), nil}, true, ref.OperationKind == OperationKindUpdate || referencedDataStatus == ReferencedDataModified
	}

	if patchRef, hasUpdated := engine.Patch.PlayerGuildMemberRef[refID]; hasUpdated {
		if patchRef.OperationKind == OperationKindUpdate {
			config.forceInclude = true
		}
		referencedElement := engine.Player(patchRef.ReferencedElementID).player
		if check == nil {
			check = newRecursionCheck()
		}
		element, _, hasUpdatedDownstream := engine.assemblePlayer(referencedElement.ID, check, config)
		referencedDataStatus := ReferencedDataUnchanged
		if hasUpdatedDownstream {
			referencedDataStatus = ReferencedDataModified
		}
		var el *Player
		if patchRef.OperationKind == OperationKindUpdate {
			el = &element
		}
		path, _ := engine.PathTrack.player[referencedElement.ID]
		return PlayerReference{patchRef.OperationKind, patchRef.ReferencedElementID, ElementKindPlayer, referencedDataStatus, path.toJSONPath(), el}, true, patchRef.OperationKind == OperationKindUpdate || referencedDataStatus == ReferencedDataModified
	}

	ref := engine.playerGuildMemberRef(refID).playerGuildMemberRef
	if check == nil {
		check = newRecursionCheck()
	}
	if _, _, hasUpdatedDownstream := engine.assemblePlayer(ref.ReferencedElementID, check, config); hasUpdatedDownstream {
		path, _ := engine.PathTrack.player[ref.ReferencedElementID]
		return PlayerReference{OperationKindUnchanged, ref.ReferencedElementID, ElementKindPlayer, ReferencedDataModified, path.toJSONPath(), nil}, true, true
	}

	return PlayerReference{}, false, false
}

func (engine *Engine) assemblePlayerEquipmentSetRef(refID PlayerEquipmentSetRefID, check *recursionCheck, config assembleConfig) (EquipmentSetReference, bool, bool) {
	if config.forceInclude {
		ref := engine.playerEquipmentSetRef(refID).playerEquipmentSetRef
		referencedElement := engine.EquipmentSet(ref.ReferencedElementID).equipmentSet
		if check == nil {
			check = newRecursionCheck()
		}
		referencedDataStatus := ReferencedDataUnchanged
		if _, _, hasUpdatedDownstream := engine.assembleEquipmentSet(referencedElement.ID, check, config); hasUpdatedDownstream {
			referencedDataStatus = ReferencedDataModified
		}
		path, _ := engine.PathTrack.equipmentSet[referencedElement.ID]
		return EquipmentSetReference{ref.OperationKind, ref.ReferencedElementID, ElementKindEquipmentSet, referencedDataStatus, path.toJSONPath(), nil}, true, ref.OperationKind == OperationKindUpdate || referencedDataStatus == ReferencedDataModified
	}

	if patchRef, hasUpdated := engine.Patch.PlayerEquipmentSetRef[refID]; hasUpdated {
		if patchRef.OperationKind == OperationKindUpdate {
			config.forceInclude = true
		}
		referencedElement := engine.EquipmentSet(patchRef.ReferencedElementID).equipmentSet
		if check == nil {
			check = newRecursionCheck()
		}
		element, _, hasUpdatedDownstream := engine.assembleEquipmentSet(referencedElement.ID, check, config)
		referencedDataStatus := ReferencedDataUnchanged
		if hasUpdatedDownstream {
			referencedDataStatus = ReferencedDataModified
		}
		var el *EquipmentSet
		if patchRef.OperationKind == OperationKindUpdate {
			el = &element
		}
		path, _ := engine.PathTrack.equipmentSet[referencedElement.ID]
		return EquipmentSetReference{patchRef.OperationKind, patchRef.ReferencedElementID, ElementKindEquipmentSet, referencedDataStatus, path.toJSONPath(), el}, true, patchRef.OperationKind == OperationKindUpdate || referencedDataStatus == ReferencedDataModified
	}

	ref := engine.playerEquipmentSetRef(refID).playerEquipmentSetRef
	if check == nil {
		check = newRecursionCheck()
	}
	if _, _, hasUpdatedDownstream := engine.assembleEquipmentSet(ref.ReferencedElementID, check, config); hasUpdatedDownstream {
		path, _ := engine.PathTrack.equipmentSet[ref.ReferencedElementID]
		return EquipmentSetReference{OperationKindUnchanged, ref.ReferencedElementID, ElementKindEquipmentSet, ReferencedDataModified, path.toJSONPath(), nil}, true, true
	}

	return EquipmentSetReference{}, false, false
}

func (engine *Engine) assembleEquipmentSetEquipmentRef(refID EquipmentSetEquipmentRefID, check *recursionCheck, config assembleConfig) (ItemReference, bool, bool) {
	if config.forceInclude {
		ref := engine.equipmentSetEquipmentRef(refID).equipmentSetEquipmentRef
		referencedElement := engine.Item(ref.ReferencedElementID).item
		if check == nil {
			check = newRecursionCheck()
		}
		referencedDataStatus := ReferencedDataUnchanged
		if _, _, hasUpdatedDownstream := engine.assembleItem(referencedElement.ID, check, config); hasUpdatedDownstream {
			referencedDataStatus = ReferencedDataModified
		}
		path, _ := engine.PathTrack.item[referencedElement.ID]
		return ItemReference{ref.OperationKind, ref.ReferencedElementID, ElementKindItem, referencedDataStatus, path.toJSONPath(), nil}, true, ref.OperationKind == OperationKindUpdate || referencedDataStatus == ReferencedDataModified
	}

	if patchRef, hasUpdated := engine.Patch.EquipmentSetEquipmentRef[refID]; hasUpdated {
		if patchRef.OperationKind == OperationKindUpdate {
			config.forceInclude = true
		}
		referencedElement := engine.Item(patchRef.ReferencedElementID).item
		if check == nil {
			check = newRecursionCheck()
		}
		element, _, hasUpdatedDownstream := engine.assembleItem(referencedElement.ID, check, config)
		referencedDataStatus := ReferencedDataUnchanged
		if hasUpdatedDownstream {
			referencedDataStatus = ReferencedDataModified
		}
		var el *Item
		if patchRef.OperationKind == OperationKindUpdate {
			el = &element
		}
		path, _ := engine.PathTrack.item[referencedElement.ID]
		return ItemReference{patchRef.OperationKind, patchRef.ReferencedElementID, ElementKindItem, referencedDataStatus, path.toJSONPath(), el}, true, patchRef.OperationKind == OperationKindUpdate || referencedDataStatus == ReferencedDataModified
	}

	ref := engine.equipmentSetEquipmentRef(refID).equipmentSetEquipmentRef
	if check == nil {
		check = newRecursionCheck()
	}
	if _, _, hasUpdatedDownstream := engine.assembleItem(ref.ReferencedElementID, check, config); hasUpdatedDownstream {
		path, _ := engine.PathTrack.item[ref.ReferencedElementID]
		return ItemReference{OperationKindUnchanged, ref.ReferencedElementID, ElementKindItem, ReferencedDataModified, path.toJSONPath(), nil}, true, true
	}

	return ItemReference{}, false, false
}

func (engine *Engine) assembleTree() Tree {

	config := assembleConfig{
		forceInclude: false,
	}

	for _, equipmentSetData := range engine.Patch.EquipmentSet {
		equipmentSet, include, _ := engine.assembleEquipmentSet(equipmentSetData.ID, nil, config)
		if include {
			engine.Tree.EquipmentSet[equipmentSetData.ID] = equipmentSet
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
		zone, include, _ := engine.assembleZone(zoneData.ID, nil, config)
		if include {
			engine.Tree.Zone[zoneData.ID] = zone
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
		if _, ok := engine.Tree.EquipmentSet[equipmentSetData.ID]; !ok {
			equipmentSet, include, _ := engine.assembleEquipmentSet(equipmentSetData.ID, nil, config)
			if include {
				engine.Tree.EquipmentSet[equipmentSetData.ID] = equipmentSet
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
		if _, ok := engine.Tree.Zone[zoneData.ID]; !ok {
			zone, include, _ := engine.assembleZone(zoneData.ID, nil, config)
			if include {
				engine.Tree.Zone[zoneData.ID] = zone
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
