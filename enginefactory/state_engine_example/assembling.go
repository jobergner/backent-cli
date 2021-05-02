package state

type recursionCheck struct {
	equipmentSet map[EquipmentSetID]bool
	gearScore    map[GearScoreID]bool
	item         map[ItemID]bool
	player       map[PlayerID]bool
	position     map[PositionID]bool
	zone         map[ZoneID]bool
	zoneItem     map[ZoneItemID]bool
}

func newRecursionCheck() *recursionCheck {
	return &recursionCheck{
		equipmentSet: make(map[EquipmentSetID]bool),
		gearScore:    make(map[GearScoreID]bool),
		item:         make(map[ItemID]bool),
		player:       make(map[PlayerID]bool),
		position:     make(map[PositionID]bool),
		zone:         make(map[ZoneID]bool),
		zoneItem:     make(map[ZoneItemID]bool),
	}
}

type assembleConfig struct {
	forceInclude   bool
	recursionCheck *recursionCheck
}

func (se *Engine) assembleGearScore(gearScoreID GearScoreID, config assembleConfig) (GearScore, bool) {
	if config.recursionCheck != nil {
		if alreadyExists := config.recursionCheck.gearScore[gearScoreID]; alreadyExists {
			return GearScore{}, false
		} else {
			config.recursionCheck.gearScore[gearScoreID] = true
		}
	}

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

func (se *Engine) assemblePosition(positionID PositionID, config assembleConfig) (Position, bool) {
	if config.recursionCheck != nil {
		if alreadyExists := config.recursionCheck.position[positionID]; alreadyExists {
			return Position{}, false
		} else {
			config.recursionCheck.position[positionID] = true
		}
	}

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

func (se *Engine) assembleEquipmentSet(equipmentSetID EquipmentSetID, config assembleConfig) (EquipmentSet, bool) {
	if config.recursionCheck != nil {
		if alreadyExists := config.recursionCheck.equipmentSet[equipmentSetID]; alreadyExists {
			return EquipmentSet{}, false
		} else {
			config.recursionCheck.equipmentSet[equipmentSetID] = true
		}
	}

	equipmentSetData, hasUpdated := se.Patch.EquipmentSet[equipmentSetID]
	if !hasUpdated {
		equipmentSetData = se.State.EquipmentSet[equipmentSetID]
	}

	var equipmentSet EquipmentSet

	for _, refID := range mergeEquipmentSetEquipmentRefIDs(se.State.EquipmentSet[equipmentSetID].Equipment, se.Patch.EquipmentSet[equipmentSetID].Equipment) {
		if ref, refHasUpdated := se.assembleEquipmentSetEquipmentRef(refID, config); refHasUpdated {
			hasUpdated = true
			equipmentSet.Equipment = append(equipmentSet.Equipment, ref)
		}
	}

	equipmentSet.ID = equipmentSetData.ID
	equipmentSet.OperationKind_ = equipmentSetData.OperationKind_
	equipmentSet.Name = equipmentSetData.Name
	return equipmentSet, hasUpdated
}

func (se *Engine) assembleItem(itemID ItemID, config assembleConfig) (Item, bool) {
	if config.recursionCheck != nil {
		if alreadyExists := config.recursionCheck.item[itemID]; alreadyExists {
			return Item{}, false
		} else {
			config.recursionCheck.item[itemID] = true
		}
	}

	itemData, hasUpdated := se.Patch.Item[itemID]
	if !hasUpdated {
		itemData = se.State.Item[itemID]
	}

	var item Item

	if refs, refHasUpdated := se.assembleItemBoundToRef(itemID, config); refHasUpdated {
		item.BoundTo = refs
		hasUpdated = true
	}
	if treeGearScore, gearScoreHasUpdated := se.assembleGearScore(itemData.GearScore, config); gearScoreHasUpdated {
		hasUpdated = true
		item.GearScore = &treeGearScore
	}

	item.ID = itemData.ID
	item.OperationKind_ = itemData.OperationKind_
	item.Name = itemData.Name
	return item, hasUpdated
}

func (se *Engine) assembleZoneItem(zoneItemID ZoneItemID, config assembleConfig) (ZoneItem, bool) {
	if config.recursionCheck != nil {
		if alreadyExists := config.recursionCheck.zoneItem[zoneItemID]; alreadyExists {
			return ZoneItem{}, false
		} else {
			config.recursionCheck.zoneItem[zoneItemID] = true
		}
	}

	zoneItemData, hasUpdated := se.Patch.ZoneItem[zoneItemID]
	if !hasUpdated {
		zoneItemData = se.State.ZoneItem[zoneItemID]
	}

	var zoneItem ZoneItem

	if treeItem, itemHasUpdated := se.assembleItem(zoneItemData.Item, config); itemHasUpdated {
		hasUpdated = true
		zoneItem.Item = &treeItem
	}
	if treePosition, positionHasUpdated := se.assemblePosition(zoneItemData.Position, config); positionHasUpdated {
		hasUpdated = true
		zoneItem.Position = &treePosition
	}

	zoneItem.ID = zoneItemData.ID
	zoneItem.OperationKind_ = zoneItemData.OperationKind_
	return zoneItem, hasUpdated

}

func (se *Engine) assemblePlayer(playerID PlayerID, config assembleConfig) (Player, bool) {
	if config.recursionCheck != nil {
		if alreadyExists := config.recursionCheck.player[playerID]; alreadyExists {
			return Player{}, false
		} else {
			config.recursionCheck.player[playerID] = true
		}
	}

	playerData, hasUpdated := se.Patch.Player[playerID]
	if !hasUpdated {
		playerData = se.State.Player[playerID]
	}

	var player Player

	for _, refID := range mergePlayerEquipmentSetRefIDs(se.State.Player[playerID].EquipmentSets, se.Patch.Player[playerID].EquipmentSets) {
		if ref, refHasUpdated := se.assemblePlayerEquipmentSetRef(refID, config); refHasUpdated {
			hasUpdated = true
			player.EquipmentSets = append(player.EquipmentSets, ref)
		}
	}
	if treeGearScore, gearScoreHasUpdated := se.assembleGearScore(playerData.GearScore, config); gearScoreHasUpdated {
		hasUpdated = true
		player.GearScore = &treeGearScore
	}
	for _, refID := range mergePlayerGuildMemberRefIDs(se.State.Player[playerID].GuildMembers, se.Patch.Player[playerID].GuildMembers) {
		if ref, refHasUpdated := se.assemblePlayerGuildMemberRef(refID, config); refHasUpdated {
			hasUpdated = true
			player.GuildMembers = append(player.GuildMembers, ref)
		}
	}
	for _, itemID := range mergeItemIDs(se.State.Player[playerData.ID].Items, se.Patch.Player[playerData.ID].Items) {
		if treeItem, itemHasUpdated := se.assembleItem(itemID, config); itemHasUpdated {
			hasUpdated = true
			player.Items = append(player.Items, treeItem)
		}
	}
	if treePosition, positionHasUpdated := se.assemblePosition(playerData.Position, config); positionHasUpdated {
		hasUpdated = true
		player.Position = &treePosition
	}

	player.ID = playerData.ID
	player.OperationKind_ = playerData.OperationKind_
	return player, hasUpdated
}

func (se *Engine) assembleZone(zoneID ZoneID, config assembleConfig) (Zone, bool) {
	if config.recursionCheck != nil {
		if alreadyExists := config.recursionCheck.zone[zoneID]; alreadyExists {
			return Zone{}, false
		} else {
			config.recursionCheck.zone[zoneID] = true
		}
	}

	zoneData, hasUpdated := se.Patch.Zone[zoneID]
	if !hasUpdated {
		zoneData = se.State.Zone[zoneID]
	}

	var zone Zone

	for _, zoneItemID := range mergeZoneItemIDs(se.State.Zone[zoneData.ID].Items, se.Patch.Zone[zoneData.ID].Items) {
		if treeZoneItem, zoneItemHasUpdated := se.assembleZoneItem(zoneItemID, config); zoneItemHasUpdated {
			hasUpdated = true
			zone.Items = append(zone.Items, treeZoneItem)
		}
	}
	for _, playerID := range mergePlayerIDs(se.State.Zone[zoneData.ID].Players, se.Patch.Zone[zoneData.ID].Players) {
		if treePlayer, playerHasUpdated := se.assemblePlayer(playerID, config); playerHasUpdated {
			hasUpdated = true
			zone.Players = append(zone.Players, treePlayer)
		}
	}

	zone.ID = zoneData.ID
	zone.OperationKind_ = zoneData.OperationKind_
	zone.Tags = zoneData.Tags
	return zone, hasUpdated
}

func (se *Engine) assembleItemBoundToRef(itemID ItemID, config assembleConfig) (*ElementReference, bool) {
	stateItem := se.State.Item[itemID]
	patchItem, itemIsInPatch := se.Patch.Item[itemID]

	// ref not set at all
	if stateItem.BoundTo == 0 && (!itemIsInPatch || patchItem.BoundTo == 0) {
		return nil, false
	}

	// immediate replacement of refs
	if stateItem.BoundTo != 0 && (itemIsInPatch && patchItem.BoundTo != 0) {
		if stateItem.BoundTo != patchItem.BoundTo {
			referencedElement := se.Player(se.itemBoundToRef(patchItem.BoundTo).itemBoundToRef.ReferencedElementID).player
			if config.recursionCheck == nil {
				config.recursionCheck = newRecursionCheck()
			}
			referencedDataStatus := ReferencedDataUnchanged
			if _, hasUpdatedDownstream := se.assemblePlayer(referencedElement.ID, config); hasUpdatedDownstream {
				referencedDataStatus = ReferencedDataModified
			}
			return &ElementReference{OperationKindUpdate, int(referencedElement.ID), ElementKindPlayer, referencedDataStatus}, true
		}
	}

	// ref was definitely removed
	if stateItem.BoundTo != 0 && (itemIsInPatch && patchItem.BoundTo == 0) {
		referencedElement := se.Player(se.itemBoundToRef(stateItem.BoundTo).itemBoundToRef.ReferencedElementID).player
		if config.recursionCheck == nil {
			config.recursionCheck = newRecursionCheck()
		}
		referencedDataStatus := ReferencedDataUnchanged
		if _, hasUpdatedDownstream := se.assemblePlayer(referencedElement.ID, config); hasUpdatedDownstream {
			referencedDataStatus = ReferencedDataModified
		}
		return &ElementReference{OperationKindDelete, int(referencedElement.ID), ElementKindPlayer, referencedDataStatus}, true
	}

	// ref was definitely created
	if stateItem.BoundTo == 0 && (itemIsInPatch && patchItem.BoundTo != 0) {
		referencedElement := se.Player(se.itemBoundToRef(patchItem.BoundTo).itemBoundToRef.ReferencedElementID).player
		if config.recursionCheck == nil {
			config.recursionCheck = newRecursionCheck()
		}
		referencedDataStatus := ReferencedDataUnchanged
		if _, hasUpdatedDownstream := se.assemblePlayer(referencedElement.ID, config); hasUpdatedDownstream {
			referencedDataStatus = ReferencedDataModified
		}
		return &ElementReference{OperationKindUpdate, int(referencedElement.ID), ElementKindPlayer, referencedDataStatus}, true
	}

	// OperationKindUpdate element got updated
	if stateItem.BoundTo != 0 {
		ref := se.itemBoundToRef(stateItem.BoundTo)
		if config.recursionCheck == nil {
			config.recursionCheck = newRecursionCheck()
		}
		if _, hasUpdatedDownstream := se.assemblePlayer(ref.ID(), config); hasUpdatedDownstream {
			return &ElementReference{OperationKindUnchanged, int(ref.ID()), ElementKindPlayer, ReferencedDataModified}, true
		}
	}

	return nil, false
}

func (se *Engine) assemblePlayerGuildMemberRef(refID PlayerGuildMemberRefID, config assembleConfig) (ElementReference, bool) {
	if patchRef, hasUpdated := se.Patch.PlayerGuildMemberRef[refID]; hasUpdated {
		referencedElement := se.Player(patchRef.ReferencedElementID).player
		if config.recursionCheck == nil {
			config.recursionCheck = newRecursionCheck()
		}
		referencedDataStatus := ReferencedDataUnchanged
		if _, hasUpdatedDownstream := se.assemblePlayer(referencedElement.ID, config); hasUpdatedDownstream {
			referencedDataStatus = ReferencedDataModified
		}
		return ElementReference{patchRef.OperationKind_, int(patchRef.ReferencedElementID), ElementKindPlayer, referencedDataStatus}, true
	}

	ref := se.playerGuildMemberRef(refID).playerGuildMemberRef
	if config.recursionCheck == nil {
		config.recursionCheck = newRecursionCheck()
	}
	if _, hasUpdatedDownstream := se.assemblePlayer(ref.ReferencedElementID, config); hasUpdatedDownstream {
		return ElementReference{OperationKindUnchanged, int(ref.ReferencedElementID), ElementKindPlayer, ReferencedDataModified}, true
	}

	return ElementReference{}, false
}

func (se *Engine) assemblePlayerEquipmentSetRef(refID PlayerEquipmentSetRefID, config assembleConfig) (ElementReference, bool) {
	if patchRef, hasUpdated := se.Patch.PlayerEquipmentSetRef[refID]; hasUpdated {
		referencedElement := se.EquipmentSet(patchRef.ReferencedElementID).equipmentSet
		if config.recursionCheck == nil {
			config.recursionCheck = newRecursionCheck()
		}
		referencedDataStatus := ReferencedDataUnchanged
		if _, hasUpdatedDownstream := se.assembleEquipmentSet(referencedElement.ID, config); hasUpdatedDownstream {
			referencedDataStatus = ReferencedDataModified
		}
		return ElementReference{patchRef.OperationKind_, int(patchRef.ReferencedElementID), ElementKindEquipmentSet, referencedDataStatus}, true
	}

	ref := se.playerEquipmentSetRef(refID).playerEquipmentSetRef
	if config.recursionCheck == nil {
		config.recursionCheck = newRecursionCheck()
	}
	if _, hasUpdatedDownstream := se.assembleEquipmentSet(ref.ReferencedElementID, config); hasUpdatedDownstream {
		return ElementReference{OperationKindUnchanged, int(ref.ReferencedElementID), ElementKindEquipmentSet, ReferencedDataModified}, true
	}

	return ElementReference{}, false
}

func (se *Engine) assembleEquipmentSetEquipmentRef(refID EquipmentSetEquipmentRefID, config assembleConfig) (ElementReference, bool) {
	if patchRef, hasUpdated := se.Patch.EquipmentSetEquipmentRef[refID]; hasUpdated {
		referencedElement := se.Item(patchRef.ReferencedElementID).item
		if config.recursionCheck == nil {
			config.recursionCheck = newRecursionCheck()
		}
		referencedDataStatus := ReferencedDataUnchanged
		if _, hasUpdatedDownstream := se.assembleItem(referencedElement.ID, config); hasUpdatedDownstream {
			referencedDataStatus = ReferencedDataModified
		}
		return ElementReference{patchRef.OperationKind_, int(patchRef.ReferencedElementID), ElementKindItem, referencedDataStatus}, true
	}

	ref := se.equipmentSetEquipmentRef(refID).equipmentSetEquipmentRef
	if config.recursionCheck == nil {
		config.recursionCheck = newRecursionCheck()
	}
	if _, hasUpdatedDownstream := se.assembleItem(ref.ReferencedElementID, config); hasUpdatedDownstream {
		return ElementReference{OperationKindUnchanged, int(ref.ReferencedElementID), ElementKindItem, ReferencedDataModified}, true
	}

	return ElementReference{}, false
}

func (se *Engine) assembleTree() Tree {

	config := assembleConfig{
		forceInclude:   false,
		recursionCheck: nil,
	}

	for _, equipmentSetData := range se.Patch.EquipmentSet {
		equipmentSet, hasUpdated := se.assembleEquipmentSet(equipmentSetData.ID, config)
		if hasUpdated {
			se.Tree.EquipmentSet[equipmentSetData.ID] = equipmentSet
		}
	}
	for _, gearScoreData := range se.Patch.GearScore {
		if !gearScoreData.HasParent_ {
			gearScore, hasUpdated := se.assembleGearScore(gearScoreData.ID, config)
			if hasUpdated {
				se.Tree.GearScore[gearScoreData.ID] = gearScore
			}
		}
	}
	for _, itemData := range se.Patch.Item {
		if !itemData.HasParent_ {
			item, hasUpdated := se.assembleItem(itemData.ID, config)
			if hasUpdated {
				se.Tree.Item[itemData.ID] = item
			}
		}
	}
	for _, playerData := range se.Patch.Player {
		if !playerData.HasParent_ {
			player, hasUpdated := se.assemblePlayer(playerData.ID, config)
			if hasUpdated {
				se.Tree.Player[playerData.ID] = player
			}
		}
	}
	for _, positionData := range se.Patch.Position {
		if !positionData.HasParent_ {
			position, hasUpdated := se.assemblePosition(positionData.ID, config)
			if hasUpdated {
				se.Tree.Position[positionData.ID] = position
			}
		}
	}
	for _, zoneData := range se.Patch.Zone {
		zone, hasUpdated := se.assembleZone(zoneData.ID, config)
		if hasUpdated {
			se.Tree.Zone[zoneData.ID] = zone
		}
	}
	for _, zoneItemData := range se.Patch.ZoneItem {
		if !zoneItemData.HasParent_ {
			zoneItem, hasUpdated := se.assembleZoneItem(zoneItemData.ID, config)
			if hasUpdated {
				se.Tree.ZoneItem[zoneItemData.ID] = zoneItem
			}
		}
	}

	for _, equipmentSetData := range se.State.EquipmentSet {
		if _, ok := se.Tree.EquipmentSet[equipmentSetData.ID]; !ok {
			equipmentSet, hasUpdated := se.assembleEquipmentSet(equipmentSetData.ID, config)
			if hasUpdated {
				se.Tree.EquipmentSet[equipmentSetData.ID] = equipmentSet
			}
		}
	}
	for _, gearScoreData := range se.State.GearScore {
		if !gearScoreData.HasParent_ {
			if _, ok := se.Tree.GearScore[gearScoreData.ID]; !ok {
				gearScore, hasUpdated := se.assembleGearScore(gearScoreData.ID, config)
				if hasUpdated {
					se.Tree.GearScore[gearScoreData.ID] = gearScore
				}
			}
		}
	}
	for _, itemData := range se.State.Item {
		if !itemData.HasParent_ {
			if _, ok := se.Tree.Item[itemData.ID]; !ok {
				item, hasUpdated := se.assembleItem(itemData.ID, config)
				if hasUpdated {
					se.Tree.Item[itemData.ID] = item
				}
			}
		}
	}
	for _, playerData := range se.State.Player {
		if !playerData.HasParent_ {
			if _, ok := se.Tree.Player[playerData.ID]; !ok {
				player, hasUpdated := se.assemblePlayer(playerData.ID, config)
				if hasUpdated {
					se.Tree.Player[playerData.ID] = player
				}
			}
		}
	}
	for _, positionData := range se.State.Position {
		if !positionData.HasParent_ {
			if _, ok := se.Tree.Position[positionData.ID]; !ok {
				position, hasUpdated := se.assemblePosition(positionData.ID, config)
				if hasUpdated {
					se.Tree.Position[positionData.ID] = position
				}
			}
		}
	}
	for _, zoneData := range se.State.Zone {
		if _, ok := se.Tree.Zone[zoneData.ID]; !ok {
			zone, hasUpdated := se.assembleZone(zoneData.ID, config)
			if hasUpdated {
				se.Tree.Zone[zoneData.ID] = zone
			}
		}
	}
	for _, zoneItemData := range se.State.ZoneItem {
		if !zoneItemData.HasParent_ {
			if _, ok := se.Tree.ZoneItem[zoneItemData.ID]; !ok {
				zoneItem, hasUpdated := se.assembleZoneItem(zoneItemData.ID, config)
				if hasUpdated {
					se.Tree.ZoneItem[zoneItemData.ID] = zoneItem
				}
			}
		}
	}

	return se.Tree
}
