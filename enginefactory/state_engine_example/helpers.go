package state

func (se Engine) allGearScoreIDs() []GearScoreID {
	var stateGearScoreIDs []GearScoreID
	for gearScoreID := range se.State.GearScore {
		stateGearScoreIDs = append(stateGearScoreIDs, gearScoreID)
	}
	var patchGearScoreIDs []GearScoreID
	for gearScoreID := range se.Patch.GearScore {
		patchGearScoreIDs = append(patchGearScoreIDs, gearScoreID)
	}
	return mergeGearScoreIDs(stateGearScoreIDs, patchGearScoreIDs)
}

func (se Engine) allItemIDs() []ItemID {
	var stateItemIDs []ItemID
	for itemID := range se.State.Item {
		stateItemIDs = append(stateItemIDs, itemID)
	}
	var patchItemIDs []ItemID
	for itemID := range se.Patch.Item {
		patchItemIDs = append(patchItemIDs, itemID)
	}
	return mergeItemIDs(stateItemIDs, patchItemIDs)
}

func (se Engine) allPlayerIDs() []PlayerID {
	var statePlayerIDs []PlayerID
	for playerID := range se.State.Player {
		statePlayerIDs = append(statePlayerIDs, playerID)
	}
	var patchPlayerIDs []PlayerID
	for playerID := range se.Patch.Player {
		patchPlayerIDs = append(patchPlayerIDs, playerID)
	}
	return mergePlayerIDs(statePlayerIDs, patchPlayerIDs)
}

func (se Engine) allPositionIDs() []PositionID {
	var statePositionIDs []PositionID
	for positionID := range se.State.Position {
		statePositionIDs = append(statePositionIDs, positionID)
	}
	var patchPositionIDs []PositionID
	for positionID := range se.Patch.Position {
		patchPositionIDs = append(patchPositionIDs, positionID)
	}
	return mergePositionIDs(statePositionIDs, patchPositionIDs)
}

func (se Engine) allZoneIDs() []ZoneID {
	var stateZoneIDs []ZoneID
	for zoneID := range se.State.Zone {
		stateZoneIDs = append(stateZoneIDs, zoneID)
	}
	var patchZoneIDs []ZoneID
	for zoneID := range se.Patch.Zone {
		patchZoneIDs = append(patchZoneIDs, zoneID)
	}
	return mergeZoneIDs(stateZoneIDs, patchZoneIDs)
}

func (se Engine) allZoneItemIDs() []ZoneItemID {
	var stateZoneItemIDs []ZoneItemID
	for zoneItemID := range se.State.ZoneItem {
		stateZoneItemIDs = append(stateZoneItemIDs, zoneItemID)
	}
	var patchZoneItemIDs []ZoneItemID
	for zoneItemID := range se.Patch.ZoneItem {
		patchZoneItemIDs = append(patchZoneItemIDs, zoneItemID)
	}
	return mergeZoneItemIDs(stateZoneItemIDs, patchZoneItemIDs)
}

func (se Engine) allItemBoundToRefIDs() []ItemBoundToRefID {
	var stateItemBoundToRefIDs []ItemBoundToRefID
	for itemBoundToRefID := range se.State.ItemBoundToRef {
		stateItemBoundToRefIDs = append(stateItemBoundToRefIDs, itemBoundToRefID)
	}
	var patchItemBoundToRefIDs []ItemBoundToRefID
	for itemBoundToRefID := range se.Patch.ItemBoundToRef {
		patchItemBoundToRefIDs = append(patchItemBoundToRefIDs, itemBoundToRefID)
	}
	return mergeItemBoundToRefIDs(stateItemBoundToRefIDs, patchItemBoundToRefIDs)
}

func (se Engine) allPlayerGuildMemberRefIDs() []PlayerGuildMemberRefID {
	var statePlayerGuildMemberRefIDs []PlayerGuildMemberRefID
	for playerGuildMemberRefID := range se.State.PlayerGuildMemberRef {
		statePlayerGuildMemberRefIDs = append(statePlayerGuildMemberRefIDs, playerGuildMemberRefID)
	}
	var patchPlayerGuildMemberRefIDs []PlayerGuildMemberRefID
	for playerGuildMemberRefID := range se.Patch.PlayerGuildMemberRef {
		patchPlayerGuildMemberRefIDs = append(patchPlayerGuildMemberRefIDs, playerGuildMemberRefID)
	}
	return mergePlayerGuildMemberRefIDs(statePlayerGuildMemberRefIDs, patchPlayerGuildMemberRefIDs)
}

func (se Engine) allPlayerEquipmentSetRefIDs() []PlayerEquipmentSetRefID {
	var statePlayerEquipmentSetRefIDs []PlayerEquipmentSetRefID
	for playerEquipmentSetRefID := range se.State.PlayerEquipmentSetRef {
		statePlayerEquipmentSetRefIDs = append(statePlayerEquipmentSetRefIDs, playerEquipmentSetRefID)
	}
	var patchPlayerEquipmentSetRefIDs []PlayerEquipmentSetRefID
	for playerEquipmentSetRefID := range se.Patch.PlayerEquipmentSetRef {
		patchPlayerEquipmentSetRefIDs = append(patchPlayerEquipmentSetRefIDs, playerEquipmentSetRefID)
	}
	return mergePlayerEquipmentSetRefIDs(statePlayerEquipmentSetRefIDs, patchPlayerEquipmentSetRefIDs)
}

func (se Engine) allEquipmentSetEquipmentRefIDs() []EquipmentSetEquipmentRefID {
	var stateEquipmentSetEquipmentRefIDs []EquipmentSetEquipmentRefID
	for equipmentSetEquipmentRefID := range se.State.EquipmentSetEquipmentRef {
		stateEquipmentSetEquipmentRefIDs = append(stateEquipmentSetEquipmentRefIDs, equipmentSetEquipmentRefID)
	}
	var patchEquipmentSetEquipmentRefIDs []EquipmentSetEquipmentRefID
	for equipmentSetEquipmentRefID := range se.Patch.EquipmentSetEquipmentRef {
		patchEquipmentSetEquipmentRefIDs = append(patchEquipmentSetEquipmentRefIDs, equipmentSetEquipmentRefID)
	}
	return mergeEquipmentSetEquipmentRefIDs(stateEquipmentSetEquipmentRefIDs, patchEquipmentSetEquipmentRefIDs)
}

func mergeGearScoreIDs(currentIDs, nextIDs []GearScoreID) []GearScoreID {
	ids := make([]GearScoreID, len(currentIDs))
	copy(ids, currentIDs)
	var j int

	for _, currentID := range currentIDs {
		if len(nextIDs) <= j || currentID != nextIDs[j] {
			continue
		}
		j += 1
	}

	for _, nextID := range nextIDs[j:] {
		ids = append(ids, nextID)
	}

	return ids
}

func mergeItemIDs(currentIDs, nextIDs []ItemID) []ItemID {
	ids := make([]ItemID, len(currentIDs))
	copy(ids, currentIDs)
	var j int

	for _, currentID := range currentIDs {
		if len(nextIDs) <= j || currentID != nextIDs[j] {
			continue
		}
		j += 1
	}

	for _, nextID := range nextIDs[j:] {
		ids = append(ids, nextID)
	}

	return ids
}

func mergePlayerIDs(currentIDs, nextIDs []PlayerID) []PlayerID {
	ids := make([]PlayerID, len(currentIDs))
	copy(ids, currentIDs)
	var j int

	for _, currentID := range currentIDs {
		if len(nextIDs) <= j || currentID != nextIDs[j] {
			continue
		}
		j += 1
	}

	for _, nextID := range nextIDs[j:] {
		ids = append(ids, nextID)
	}

	return ids
}

func mergePositionIDs(currentIDs, nextIDs []PositionID) []PositionID {
	ids := make([]PositionID, len(currentIDs))
	copy(ids, currentIDs)
	var j int

	for _, currentID := range currentIDs {
		if len(nextIDs) <= j || currentID != nextIDs[j] {
			continue
		}
		j += 1
	}

	for _, nextID := range nextIDs[j:] {
		ids = append(ids, nextID)
	}

	return ids
}

func mergeZoneIDs(currentIDs, nextIDs []ZoneID) []ZoneID {
	ids := make([]ZoneID, len(currentIDs))
	copy(ids, currentIDs)
	var j int

	for _, currentID := range currentIDs {
		if len(nextIDs) <= j || currentID != nextIDs[j] {
			continue
		}
		j += 1
	}

	for _, nextID := range nextIDs[j:] {
		ids = append(ids, nextID)
	}

	return ids
}

func mergeZoneItemIDs(currentIDs, nextIDs []ZoneItemID) []ZoneItemID {
	ids := make([]ZoneItemID, len(currentIDs))
	copy(ids, currentIDs)
	var j int

	for _, currentID := range currentIDs {
		if len(nextIDs) <= j || currentID != nextIDs[j] {
			continue
		}
		j += 1
	}

	for _, nextID := range nextIDs[j:] {
		ids = append(ids, nextID)
	}

	return ids
}

func mergeEquipmentSetIDs(currentIDs, nextIDs []EquipmentSetID) []EquipmentSetID {
	ids := make([]EquipmentSetID, len(currentIDs))
	copy(ids, currentIDs)
	var j int

	for _, currentID := range currentIDs {
		if len(nextIDs) <= j || currentID != nextIDs[j] {
			continue
		}
		j += 1
	}

	for _, nextID := range nextIDs[j:] {
		ids = append(ids, nextID)
	}

	return ids
}

func mergeItemBoundToRefIDs(currentIDs, nextIDs []ItemBoundToRefID) []ItemBoundToRefID {
	ids := make([]ItemBoundToRefID, len(currentIDs))
	copy(ids, currentIDs)
	var j int

	for _, currentID := range currentIDs {
		if len(nextIDs) <= j || currentID != nextIDs[j] {
			continue
		}
		j += 1
	}

	for _, nextID := range nextIDs[j:] {
		ids = append(ids, nextID)
	}

	return ids
}

func mergeEquipmentSetEquipmentRefIDs(currentIDs, nextIDs []EquipmentSetEquipmentRefID) []EquipmentSetEquipmentRefID {
	ids := make([]EquipmentSetEquipmentRefID, len(currentIDs))
	copy(ids, currentIDs)
	var j int

	for _, currentID := range currentIDs {
		if len(nextIDs) <= j || currentID != nextIDs[j] {
			continue
		}
		j += 1
	}

	for _, nextID := range nextIDs[j:] {
		ids = append(ids, nextID)
	}

	return ids
}

func mergePlayerGuildMemberRefIDs(currentIDs, nextIDs []PlayerGuildMemberRefID) []PlayerGuildMemberRefID {
	ids := make([]PlayerGuildMemberRefID, len(currentIDs))
	copy(ids, currentIDs)
	var j int

	for _, currentID := range currentIDs {
		if len(nextIDs) <= j || currentID != nextIDs[j] {
			continue
		}
		j += 1
	}

	for _, nextID := range nextIDs[j:] {
		ids = append(ids, nextID)
	}

	return ids
}

func mergePlayerEquipmentSetRefIDs(currentIDs, nextIDs []PlayerEquipmentSetRefID) []PlayerEquipmentSetRefID {
	ids := make([]PlayerEquipmentSetRefID, len(currentIDs))
	copy(ids, currentIDs)
	var j int

	for _, currentID := range currentIDs {
		if len(nextIDs) <= j || currentID != nextIDs[j] {
			continue
		}
		j += 1
	}

	for _, nextID := range nextIDs[j:] {
		ids = append(ids, nextID)
	}

	return ids
}

func (se *Engine) dereferenceItemBoundToRefs(playerID PlayerID) {
	for _, refID := range se.allItemBoundToRefIDs() {
		ref := se.itemBoundToRef(refID)
		if ref.itemBoundToRef.ReferencedElementID == playerID {
			ref.Unset(se)
		}
	}
}

func (se *Engine) dereferencePlayerGuildMemberRefs(playerID PlayerID) {
	for _, refID := range se.allPlayerGuildMemberRefIDs() {
		ref := se.playerGuildMemberRef(refID)
		if ref.playerGuildMemberRef.ReferencedElementID == playerID {
			parent := se.Player(ref.playerGuildMemberRef.ParentID)
			parent.RemoveGuildMembers(se, playerID)
		}
	}
}

func (se *Engine) dereferencePlayerEquipmentSetRefs(equipmentSetID EquipmentSetID) {
	for _, refID := range se.allPlayerEquipmentSetRefIDs() {
		ref := se.playerEquipmentSetRef(refID)
		if ref.playerEquipmentSetRef.ReferencedElementID == equipmentSetID {
			parent := se.Player(ref.playerEquipmentSetRef.ParentID)
			parent.RemoveEquipmentSets(se, equipmentSetID)
		}
	}
}

func (se *Engine) dereferenceEquipmentSetEquipmentRef(itemID ItemID) {
	for _, refID := range se.allEquipmentSetEquipmentRefIDs() {
		ref := se.equipmentSetEquipmentRef(refID)
		if ref.equipmentSetEquipmentRef.ReferencedElementID == itemID {
			parent := se.EquipmentSet(ref.equipmentSetEquipmentRef.ParentID)
			parent.RemoveEquipment(se, itemID)
		}
	}
}
