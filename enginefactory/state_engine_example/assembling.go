package state

type assembleConfig struct {
	forceInclude bool // include everything, regardless of update status
}

func (se *Engine) assembleGearScore(gearScoreID GearScoreID, check *recursionCheck, config assembleConfig) (GearScore, bool, bool) {
	gearScoreData, hasUpdated := se.Patch.GearScore[gearScoreID]
	if !hasUpdated {
		gearScoreData, _ = se.State.GearScore[gearScoreID]
	}

	if check != nil {
		if alreadyExists := check.gearScore[gearScoreID]; alreadyExists {
			return GearScore{}, false, hasUpdated
		} else {
			check.gearScore[gearScoreID] = true
		}
	}

	var gearScore GearScore

	gearScore.ID = gearScoreData.ID
	gearScore.OperationKind_ = gearScoreData.OperationKind_
	gearScore.Level = gearScoreData.Level
	gearScore.Score = gearScoreData.Score
	return gearScore, hasUpdated || config.forceInclude, hasUpdated
}

func (se *Engine) assemblePosition(positionID PositionID, check *recursionCheck, config assembleConfig) (Position, bool, bool) {
	positionData, hasUpdated := se.Patch.Position[positionID]
	if !hasUpdated {
		positionData, _ = se.State.Position[positionID]
	}

	if check != nil {
		if alreadyExists := check.position[positionID]; alreadyExists {
			return Position{}, false, hasUpdated
		} else {
			check.position[positionID] = true
		}
	}

	var position Position

	position.ID = positionData.ID
	position.OperationKind_ = positionData.OperationKind_
	position.X = positionData.X
	position.Y = positionData.Y
	return position, hasUpdated || config.forceInclude, hasUpdated
}

func (se *Engine) assembleEquipmentSet(equipmentSetID EquipmentSetID, check *recursionCheck, config assembleConfig) (EquipmentSet, bool, bool) {
	equipmentSetData, hasUpdated := se.Patch.EquipmentSet[equipmentSetID]
	if !hasUpdated {
		equipmentSetData = se.State.EquipmentSet[equipmentSetID]
	}

	if check != nil {
		if alreadyExists := check.equipmentSet[equipmentSetID]; alreadyExists {
			return EquipmentSet{}, false, hasUpdated
		} else {
			check.equipmentSet[equipmentSetID] = true
		}
	}

	var equipmentSet EquipmentSet

	for _, refID := range mergeEquipmentSetEquipmentRefIDs(se.State.EquipmentSet[equipmentSetID].Equipment, se.Patch.EquipmentSet[equipmentSetID].Equipment) {
		if ref, include, refHasUpdated := se.assembleEquipmentSetEquipmentRef(refID, check, config); include {
			if refHasUpdated {
				hasUpdated = true
			}
			equipmentSet.Equipment = append(equipmentSet.Equipment, ref)
		}
	}

	equipmentSet.ID = equipmentSetData.ID
	equipmentSet.OperationKind_ = equipmentSetData.OperationKind_
	equipmentSet.Name = equipmentSetData.Name
	return equipmentSet, hasUpdated || config.forceInclude, hasUpdated
}

func (se *Engine) assembleItem(itemID ItemID, check *recursionCheck, config assembleConfig) (Item, bool, bool) {
	itemData, hasUpdated := se.Patch.Item[itemID]
	if !hasUpdated {
		itemData = se.State.Item[itemID]
	}

	if check != nil {
		if alreadyExists := check.item[itemID]; alreadyExists {
			return Item{}, false, hasUpdated
		} else {
			check.item[itemID] = true
		}
	}

	var item Item

	if refs, include, refHasUpdated := se.assembleItemBoundToRef(itemID, check, config); include {
		if refHasUpdated {
			hasUpdated = true
		}
		item.BoundTo = refs
	}
	if treeGearScore, include, gearScoreHasUpdated := se.assembleGearScore(itemData.GearScore, check, config); include {
		if gearScoreHasUpdated {
			hasUpdated = true
		}
		item.GearScore = &treeGearScore
	}

	item.ID = itemData.ID
	item.OperationKind_ = itemData.OperationKind_
	item.Name = itemData.Name
	return item, hasUpdated || config.forceInclude, hasUpdated
}

func (se *Engine) assembleZoneItem(zoneItemID ZoneItemID, check *recursionCheck, config assembleConfig) (ZoneItem, bool, bool) {
	zoneItemData, hasUpdated := se.Patch.ZoneItem[zoneItemID]
	if !hasUpdated {
		zoneItemData = se.State.ZoneItem[zoneItemID]
	}

	if check != nil {
		if alreadyExists := check.zoneItem[zoneItemID]; alreadyExists {
			return ZoneItem{}, false, hasUpdated
		} else {
			check.zoneItem[zoneItemID] = true
		}
	}

	var zoneItem ZoneItem

	if treeItem, include, itemHasUpdated := se.assembleItem(zoneItemData.Item, check, config); include {
		if itemHasUpdated {
			hasUpdated = true
		}
		zoneItem.Item = &treeItem
	}
	if treePosition, include, positionHasUpdated := se.assemblePosition(zoneItemData.Position, check, config); include {
		if positionHasUpdated {
			hasUpdated = true
		}
		zoneItem.Position = &treePosition
	}

	zoneItem.ID = zoneItemData.ID
	zoneItem.OperationKind_ = zoneItemData.OperationKind_
	return zoneItem, hasUpdated || config.forceInclude, hasUpdated
}

func (se *Engine) assemblePlayer(playerID PlayerID, check *recursionCheck, config assembleConfig) (Player, bool, bool) {
	playerData, hasUpdated := se.Patch.Player[playerID]
	if !hasUpdated {
		playerData = se.State.Player[playerID]
	}

	if check != nil {
		if alreadyExists := check.player[playerID]; alreadyExists {
			return Player{}, false, hasUpdated
		} else {
			check.player[playerID] = true
		}
	}

	var player Player

	for _, refID := range mergePlayerEquipmentSetRefIDs(se.State.Player[playerID].EquipmentSets, se.Patch.Player[playerID].EquipmentSets) {
		if ref, include, refHasUpdated := se.assemblePlayerEquipmentSetRef(refID, check, config); include {
			if refHasUpdated {
				hasUpdated = true
			}
			player.EquipmentSets = append(player.EquipmentSets, ref)
		}
	}
	if treeGearScore, include, gearScoreHasUpdated := se.assembleGearScore(playerData.GearScore, check, config); include {
		if gearScoreHasUpdated {
			hasUpdated = true
		}
		player.GearScore = &treeGearScore
	}
	for _, refID := range mergePlayerGuildMemberRefIDs(se.State.Player[playerID].GuildMembers, se.Patch.Player[playerID].GuildMembers) {
		if ref, include, refHasUpdated := se.assemblePlayerGuildMemberRef(refID, check, config); include {
			if refHasUpdated {
				hasUpdated = true
			}
			player.GuildMembers = append(player.GuildMembers, ref)
		}
	}
	for _, itemID := range mergeItemIDs(se.State.Player[playerData.ID].Items, se.Patch.Player[playerData.ID].Items) {
		if treeItem, include, itemHasUpdated := se.assembleItem(itemID, check, config); include {
			if itemHasUpdated {
				hasUpdated = true
			}
			player.Items = append(player.Items, treeItem)
		}
	}
	if treePosition, include, positionHasUpdated := se.assemblePosition(playerData.Position, check, config); include {
		if positionHasUpdated {
			hasUpdated = true
		}
		player.Position = &treePosition
	}

	player.ID = playerData.ID
	player.OperationKind_ = playerData.OperationKind_
	return player, hasUpdated || config.forceInclude, hasUpdated
}

func (se *Engine) assembleZone(zoneID ZoneID, check *recursionCheck, config assembleConfig) (Zone, bool, bool) {
	zoneData, hasUpdated := se.Patch.Zone[zoneID]
	if !hasUpdated {
		zoneData = se.State.Zone[zoneID]
	}

	if check != nil {
		if alreadyExists := check.zone[zoneID]; alreadyExists {
			return Zone{}, false, hasUpdated
		} else {
			check.zone[zoneID] = true
		}
	}

	var zone Zone

	for _, zoneItemID := range mergeZoneItemIDs(se.State.Zone[zoneData.ID].Items, se.Patch.Zone[zoneData.ID].Items) {
		if treeZoneItem, include, zoneItemHasUpdated := se.assembleZoneItem(zoneItemID, check, config); include {
			if zoneItemHasUpdated {
				hasUpdated = true
			}
			zone.Items = append(zone.Items, treeZoneItem)
		}
	}
	for _, playerID := range mergePlayerIDs(se.State.Zone[zoneData.ID].Players, se.Patch.Zone[zoneData.ID].Players) {
		if treePlayer, include, playerHasUpdated := se.assemblePlayer(playerID, check, config); include {
			if playerHasUpdated {
				hasUpdated = true
			}
			zone.Players = append(zone.Players, treePlayer)
		}
	}

	zone.ID = zoneData.ID
	zone.OperationKind_ = zoneData.OperationKind_
	zone.Tags = zoneData.Tags
	return zone, hasUpdated || config.forceInclude, hasUpdated
}

func (se *Engine) assembleItemBoundToRef(itemID ItemID, check *recursionCheck, config assembleConfig) (*PlayerReference, bool, bool) {
	stateItem := se.State.Item[itemID]
	patchItem, itemIsInPatch := se.Patch.Item[itemID]

	// ref not set at all
	if stateItem.BoundTo == 0 && (!itemIsInPatch || patchItem.BoundTo == 0) {
		return nil, false, false
	}

	// force include
	if config.forceInclude {
		ref := se.itemBoundToRef(patchItem.BoundTo)
		referencedElement := se.Player(ref.itemBoundToRef.ReferencedElementID).player
		if check == nil {
			check = newRecursionCheck()
		}
		referencedDataStatus := ReferencedDataUnchanged
		if _, _, hasUpdatedDownstream := se.assemblePlayer(referencedElement.ID, check, config); hasUpdatedDownstream {
			referencedDataStatus = ReferencedDataModified
		}
		path, _ := se.PathTrack.player[referencedElement.ID]
		return &PlayerReference{ref.itemBoundToRef.OperationKind_, referencedElement.ID, ElementKindPlayer, referencedDataStatus, path.toJSONPath(), nil}, true, ref.itemBoundToRef.OperationKind_ == OperationKindUpdate || referencedDataStatus == ReferencedDataModified
	}

	// ref was definitely created
	if stateItem.BoundTo == 0 && (itemIsInPatch && patchItem.BoundTo != 0) {
		config.forceInclude = true
		ref := se.itemBoundToRef(patchItem.BoundTo)
		referencedElement := se.Player(ref.itemBoundToRef.ReferencedElementID).player
		if check == nil {
			check = newRecursionCheck()
		}
		referencedDataStatus := ReferencedDataUnchanged
		element, _, hasUpdatedDownstream := se.assemblePlayer(referencedElement.ID, check, config)
		if hasUpdatedDownstream {
			referencedDataStatus = ReferencedDataModified
		}
		path, _ := se.PathTrack.player[referencedElement.ID]
		return &PlayerReference{OperationKindUpdate, referencedElement.ID, ElementKindPlayer, referencedDataStatus, path.toJSONPath(), &element}, true, referencedDataStatus == ReferencedDataModified
	}

	// ref was definitely removed
	if stateItem.BoundTo != 0 && (itemIsInPatch && patchItem.BoundTo == 0) {
		ref := se.itemBoundToRef(stateItem.BoundTo)
		referencedElement := se.Player(ref.itemBoundToRef.ReferencedElementID).player
		if check == nil {
			check = newRecursionCheck()
		}
		referencedDataStatus := ReferencedDataUnchanged
		if _, _, hasUpdatedDownstream := se.assemblePlayer(referencedElement.ID, check, config); hasUpdatedDownstream {
			referencedDataStatus = ReferencedDataModified
		}
		path, _ := se.PathTrack.player[referencedElement.ID]
		return &PlayerReference{OperationKindDelete, referencedElement.ID, ElementKindPlayer, referencedDataStatus, path.toJSONPath(), nil}, true, referencedDataStatus == ReferencedDataModified
	}

	// immediate replacement of refs
	if stateItem.BoundTo != 0 && (itemIsInPatch && patchItem.BoundTo != 0) {
		if stateItem.BoundTo != patchItem.BoundTo {
			ref := se.itemBoundToRef(patchItem.BoundTo)
			referencedElement := se.Player(ref.itemBoundToRef.ReferencedElementID).player
			if check == nil {
				check = newRecursionCheck()
			}
			referencedDataStatus := ReferencedDataUnchanged
			if _, _, hasUpdatedDownstream := se.assemblePlayer(referencedElement.ID, check, config); hasUpdatedDownstream {
				referencedDataStatus = ReferencedDataModified
			}
			path, _ := se.PathTrack.player[referencedElement.ID]
			return &PlayerReference{OperationKindUpdate, referencedElement.ID, ElementKindPlayer, referencedDataStatus, path.toJSONPath(), nil}, true, referencedDataStatus == ReferencedDataModified
		}
	}

	// OperationKindUpdate element got updated
	if stateItem.BoundTo != 0 {
		ref := se.itemBoundToRef(stateItem.BoundTo)
		if check == nil {
			check = newRecursionCheck()
		}
		if _, _, hasUpdatedDownstream := se.assemblePlayer(ref.ID(), check, config); hasUpdatedDownstream {
			path, _ := se.PathTrack.player[ref.itemBoundToRef.ReferencedElementID]
			return &PlayerReference{OperationKindUnchanged, ref.ID(), ElementKindPlayer, ReferencedDataModified, path.toJSONPath(), nil}, true, true
		}
	}

	return nil, false, false
}

func (se *Engine) assemblePlayerGuildMemberRef(refID PlayerGuildMemberRefID, check *recursionCheck, config assembleConfig) (PlayerReference, bool, bool) {
	if config.forceInclude {
		ref := se.playerGuildMemberRef(refID).playerGuildMemberRef
		referencedElement := se.Player(ref.ReferencedElementID).player
		if check == nil {
			check = newRecursionCheck()
		}
		referencedDataStatus := ReferencedDataUnchanged
		if _, _, hasUpdatedDownstream := se.assemblePlayer(referencedElement.ID, check, config); hasUpdatedDownstream {
			referencedDataStatus = ReferencedDataModified
		}
		path, _ := se.PathTrack.player[referencedElement.ID]
		return PlayerReference{ref.OperationKind_, ref.ReferencedElementID, ElementKindPlayer, referencedDataStatus, path.toJSONPath(), nil}, true, ref.OperationKind_ == OperationKindUpdate || referencedDataStatus == ReferencedDataModified
	}

	if patchRef, hasUpdated := se.Patch.PlayerGuildMemberRef[refID]; hasUpdated {
		if patchRef.OperationKind_ == OperationKindUpdate {
			config.forceInclude = true
		}
		referencedElement := se.Player(patchRef.ReferencedElementID).player
		if check == nil {
			check = newRecursionCheck()
		}
		element, _, hasUpdatedDownstream := se.assemblePlayer(referencedElement.ID, check, config)
		referencedDataStatus := ReferencedDataUnchanged
		if hasUpdatedDownstream {
			referencedDataStatus = ReferencedDataModified
		}
		var el *Player
		if patchRef.OperationKind_ == OperationKindUpdate {
			el = &element
		}
		path, _ := se.PathTrack.player[referencedElement.ID]
		return PlayerReference{patchRef.OperationKind_, patchRef.ReferencedElementID, ElementKindPlayer, referencedDataStatus, path.toJSONPath(), el}, true, patchRef.OperationKind_ == OperationKindUpdate || referencedDataStatus == ReferencedDataModified
	}

	ref := se.playerGuildMemberRef(refID).playerGuildMemberRef
	if check == nil {
		check = newRecursionCheck()
	}
	if _, _, hasUpdatedDownstream := se.assemblePlayer(ref.ReferencedElementID, check, config); hasUpdatedDownstream {
		path, _ := se.PathTrack.player[ref.ReferencedElementID]
		return PlayerReference{OperationKindUnchanged, ref.ReferencedElementID, ElementKindPlayer, ReferencedDataModified, path.toJSONPath(), nil}, true, true
	}

	return PlayerReference{}, false, false
}

func (se *Engine) assemblePlayerEquipmentSetRef(refID PlayerEquipmentSetRefID, check *recursionCheck, config assembleConfig) (EquipmentSetReference, bool, bool) {
	if config.forceInclude {
		ref := se.playerEquipmentSetRef(refID).playerEquipmentSetRef
		referencedElement := se.EquipmentSet(ref.ReferencedElementID).equipmentSet
		if check == nil {
			check = newRecursionCheck()
		}
		referencedDataStatus := ReferencedDataUnchanged
		if _, _, hasUpdatedDownstream := se.assembleEquipmentSet(referencedElement.ID, check, config); hasUpdatedDownstream {
			referencedDataStatus = ReferencedDataModified
		}
		path, _ := se.PathTrack.equipmentSet[referencedElement.ID]
		return EquipmentSetReference{ref.OperationKind_, ref.ReferencedElementID, ElementKindEquipmentSet, referencedDataStatus, path.toJSONPath(), nil}, true, ref.OperationKind_ == OperationKindUpdate || referencedDataStatus == ReferencedDataModified
	}

	if patchRef, hasUpdated := se.Patch.PlayerEquipmentSetRef[refID]; hasUpdated {
		if patchRef.OperationKind_ == OperationKindUpdate {
			config.forceInclude = true
		}
		referencedElement := se.EquipmentSet(patchRef.ReferencedElementID).equipmentSet
		if check == nil {
			check = newRecursionCheck()
		}
		element, _, hasUpdatedDownstream := se.assembleEquipmentSet(referencedElement.ID, check, config)
		referencedDataStatus := ReferencedDataUnchanged
		if hasUpdatedDownstream {
			referencedDataStatus = ReferencedDataModified
		}
		var el *EquipmentSet
		if patchRef.OperationKind_ == OperationKindUpdate {
			el = &element
		}
		path, _ := se.PathTrack.equipmentSet[referencedElement.ID]
		return EquipmentSetReference{patchRef.OperationKind_, patchRef.ReferencedElementID, ElementKindEquipmentSet, referencedDataStatus, path.toJSONPath(), el}, true, patchRef.OperationKind_ == OperationKindUpdate || referencedDataStatus == ReferencedDataModified
	}

	ref := se.playerEquipmentSetRef(refID).playerEquipmentSetRef
	if check == nil {
		check = newRecursionCheck()
	}
	if _, _, hasUpdatedDownstream := se.assembleEquipmentSet(ref.ReferencedElementID, check, config); hasUpdatedDownstream {
		path, _ := se.PathTrack.equipmentSet[ref.ReferencedElementID]
		return EquipmentSetReference{OperationKindUnchanged, ref.ReferencedElementID, ElementKindEquipmentSet, ReferencedDataModified, path.toJSONPath(), nil}, true, true
	}

	return EquipmentSetReference{}, false, false
}

func (se *Engine) assembleEquipmentSetEquipmentRef(refID EquipmentSetEquipmentRefID, check *recursionCheck, config assembleConfig) (ItemReference, bool, bool) {
	if config.forceInclude {
		ref := se.equipmentSetEquipmentRef(refID).equipmentSetEquipmentRef
		referencedElement := se.Item(ref.ReferencedElementID).item
		if check == nil {
			check = newRecursionCheck()
		}
		referencedDataStatus := ReferencedDataUnchanged
		if _, _, hasUpdatedDownstream := se.assembleItem(referencedElement.ID, check, config); hasUpdatedDownstream {
			referencedDataStatus = ReferencedDataModified
		}
		path, _ := se.PathTrack.item[referencedElement.ID]
		return ItemReference{ref.OperationKind_, ref.ReferencedElementID, ElementKindItem, referencedDataStatus, path.toJSONPath(), nil}, true, ref.OperationKind_ == OperationKindUpdate || referencedDataStatus == ReferencedDataModified
	}

	if patchRef, hasUpdated := se.Patch.EquipmentSetEquipmentRef[refID]; hasUpdated {
		if patchRef.OperationKind_ == OperationKindUpdate {
			config.forceInclude = true
		}
		referencedElement := se.Item(patchRef.ReferencedElementID).item
		if check == nil {
			check = newRecursionCheck()
		}
		element, _, hasUpdatedDownstream := se.assembleItem(referencedElement.ID, check, config)
		referencedDataStatus := ReferencedDataUnchanged
		if hasUpdatedDownstream {
			referencedDataStatus = ReferencedDataModified
		}
		var el *Item
		if patchRef.OperationKind_ == OperationKindUpdate {
			el = &element
		}
		path, _ := se.PathTrack.item[referencedElement.ID]
		return ItemReference{patchRef.OperationKind_, patchRef.ReferencedElementID, ElementKindItem, referencedDataStatus, path.toJSONPath(), el}, true, patchRef.OperationKind_ == OperationKindUpdate || referencedDataStatus == ReferencedDataModified
	}

	ref := se.equipmentSetEquipmentRef(refID).equipmentSetEquipmentRef
	if check == nil {
		check = newRecursionCheck()
	}
	if _, _, hasUpdatedDownstream := se.assembleItem(ref.ReferencedElementID, check, config); hasUpdatedDownstream {
		path, _ := se.PathTrack.item[ref.ReferencedElementID]
		return ItemReference{OperationKindUnchanged, ref.ReferencedElementID, ElementKindItem, ReferencedDataModified, path.toJSONPath(), nil}, true, true
	}

	return ItemReference{}, false, false
}

func (se *Engine) assembleTree() Tree {

	config := assembleConfig{
		forceInclude: false,
	}

	for _, equipmentSetData := range se.Patch.EquipmentSet {
		equipmentSet, include, _ := se.assembleEquipmentSet(equipmentSetData.ID, nil, config)
		if include {
			se.Tree.EquipmentSet[equipmentSetData.ID] = equipmentSet
		}
	}
	for _, gearScoreData := range se.Patch.GearScore {
		if !gearScoreData.HasParent_ {
			gearScore, include, _ := se.assembleGearScore(gearScoreData.ID, nil, config)
			if include {
				se.Tree.GearScore[gearScoreData.ID] = gearScore
			}
		}
	}
	for _, itemData := range se.Patch.Item {
		if !itemData.HasParent_ {
			item, include, _ := se.assembleItem(itemData.ID, nil, config)
			if include {
				se.Tree.Item[itemData.ID] = item
			}
		}
	}
	for _, playerData := range se.Patch.Player {
		if !playerData.HasParent_ {
			player, include, _ := se.assemblePlayer(playerData.ID, nil, config)
			if include {
				se.Tree.Player[playerData.ID] = player
			}
		}
	}
	for _, positionData := range se.Patch.Position {
		if !positionData.HasParent_ {
			position, include, _ := se.assemblePosition(positionData.ID, nil, config)
			if include {
				se.Tree.Position[positionData.ID] = position
			}
		}
	}
	for _, zoneData := range se.Patch.Zone {
		zone, include, _ := se.assembleZone(zoneData.ID, nil, config)
		if include {
			se.Tree.Zone[zoneData.ID] = zone
		}
	}
	for _, zoneItemData := range se.Patch.ZoneItem {
		if !zoneItemData.HasParent_ {
			zoneItem, include, _ := se.assembleZoneItem(zoneItemData.ID, nil, config)
			if include {
				se.Tree.ZoneItem[zoneItemData.ID] = zoneItem
			}
		}
	}

	for _, equipmentSetData := range se.State.EquipmentSet {
		if _, ok := se.Tree.EquipmentSet[equipmentSetData.ID]; !ok {
			equipmentSet, include, _ := se.assembleEquipmentSet(equipmentSetData.ID, nil, config)
			if include {
				se.Tree.EquipmentSet[equipmentSetData.ID] = equipmentSet
			}
		}
	}
	for _, gearScoreData := range se.State.GearScore {
		if !gearScoreData.HasParent_ {
			if _, ok := se.Tree.GearScore[gearScoreData.ID]; !ok {
				gearScore, include, _ := se.assembleGearScore(gearScoreData.ID, nil, config)
				if include {
					se.Tree.GearScore[gearScoreData.ID] = gearScore
				}
			}
		}
	}
	for _, itemData := range se.State.Item {
		if !itemData.HasParent_ {
			if _, ok := se.Tree.Item[itemData.ID]; !ok {
				item, include, _ := se.assembleItem(itemData.ID, nil, config)
				if include {
					se.Tree.Item[itemData.ID] = item
				}
			}
		}
	}
	for _, playerData := range se.State.Player {
		if !playerData.HasParent_ {
			if _, ok := se.Tree.Player[playerData.ID]; !ok {
				player, include, _ := se.assemblePlayer(playerData.ID, nil, config)
				if include {
					se.Tree.Player[playerData.ID] = player
				}
			}
		}
	}
	for _, positionData := range se.State.Position {
		if !positionData.HasParent_ {
			if _, ok := se.Tree.Position[positionData.ID]; !ok {
				position, include, _ := se.assemblePosition(positionData.ID, nil, config)
				if include {
					se.Tree.Position[positionData.ID] = position
				}
			}
		}
	}
	for _, zoneData := range se.State.Zone {
		if _, ok := se.Tree.Zone[zoneData.ID]; !ok {
			zone, include, _ := se.assembleZone(zoneData.ID, nil, config)
			if include {
				se.Tree.Zone[zoneData.ID] = zone
			}
		}
	}
	for _, zoneItemData := range se.State.ZoneItem {
		if !zoneItemData.HasParent_ {
			if _, ok := se.Tree.ZoneItem[zoneItemData.ID]; !ok {
				zoneItem, include, _ := se.assembleZoneItem(zoneItemData.ID, nil, config)
				if include {
					se.Tree.ZoneItem[zoneItemData.ID] = zoneItem
				}
			}
		}
	}

	return se.Tree
}
