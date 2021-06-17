package state

func deduplicateZoneItemIDs(a []ZoneItemID, b []ZoneItemID) []ZoneItemID {

	check := make(map[ZoneItemID]bool)
	deduped := make([]ZoneItemID, 0)
	for _, val := range a {
		check[val] = true
	}
	for _, val := range b {
		check[val] = true
	}

	for val := range check {
		deduped = append(deduped, val)
	}

	return deduped
}

func deduplicateZoneIDs(a []ZoneID, b []ZoneID) []ZoneID {

	check := make(map[ZoneID]bool)
	deduped := make([]ZoneID, 0)
	for _, val := range a {
		check[val] = true
	}
	for _, val := range b {
		check[val] = true
	}

	for val := range check {
		deduped = append(deduped, val)
	}

	return deduped
}

func deduplicatePositionIDs(a []PositionID, b []PositionID) []PositionID {

	check := make(map[PositionID]bool)
	deduped := make([]PositionID, 0)
	for _, val := range a {
		check[val] = true
	}
	for _, val := range b {
		check[val] = true
	}

	for val := range check {
		deduped = append(deduped, val)
	}

	return deduped
}

func deduplicateItemIDs(a []ItemID, b []ItemID) []ItemID {

	check := make(map[ItemID]bool)
	deduped := make([]ItemID, 0)
	for _, val := range a {
		check[val] = true
	}
	for _, val := range b {
		check[val] = true
	}

	for val := range check {
		deduped = append(deduped, val)
	}

	return deduped
}

func deduplicateGearScoreIDs(a []GearScoreID, b []GearScoreID) []GearScoreID {

	check := make(map[GearScoreID]bool)
	deduped := make([]GearScoreID, 0)
	for _, val := range a {
		check[val] = true
	}
	for _, val := range b {
		check[val] = true
	}

	for val := range check {
		deduped = append(deduped, val)
	}

	return deduped
}

func deduplicateEquipmentSetIDs(a []EquipmentSetID, b []EquipmentSetID) []EquipmentSetID {

	check := make(map[EquipmentSetID]bool)
	deduped := make([]EquipmentSetID, 0)
	for _, val := range a {
		check[val] = true
	}
	for _, val := range b {
		check[val] = true
	}

	for val := range check {
		deduped = append(deduped, val)
	}

	return deduped
}

func deduplicatePlayerIDs(a []PlayerID, b []PlayerID) []PlayerID {

	check := make(map[PlayerID]bool)
	deduped := make([]PlayerID, 0)
	for _, val := range a {
		check[val] = true
	}
	for _, val := range b {
		check[val] = true
	}

	for val := range check {
		deduped = append(deduped, val)
	}

	return deduped
}

func deduplicatePlayerTargetedByRefIDs(a []PlayerTargetedByRefID, b []PlayerTargetedByRefID) []PlayerTargetedByRefID {

	check := make(map[PlayerTargetedByRefID]bool)
	deduped := make([]PlayerTargetedByRefID, 0)
	for _, val := range a {
		check[val] = true
	}
	for _, val := range b {
		check[val] = true
	}

	for val := range check {
		deduped = append(deduped, val)
	}

	return deduped
}

func deduplicatePlayerTargetRefIDs(a []PlayerTargetRefID, b []PlayerTargetRefID) []PlayerTargetRefID {

	check := make(map[PlayerTargetRefID]bool)
	deduped := make([]PlayerTargetRefID, 0)
	for _, val := range a {
		check[val] = true
	}
	for _, val := range b {
		check[val] = true
	}

	for val := range check {
		deduped = append(deduped, val)
	}

	return deduped
}

func deduplicateItemBoundToRefIDs(a []ItemBoundToRefID, b []ItemBoundToRefID) []ItemBoundToRefID {

	check := make(map[ItemBoundToRefID]bool)
	deduped := make([]ItemBoundToRefID, 0)
	for _, val := range a {
		check[val] = true
	}
	for _, val := range b {
		check[val] = true
	}

	for val := range check {
		deduped = append(deduped, val)
	}

	return deduped
}

func deduplicatePlayerGuildMemberRefIDs(a []PlayerGuildMemberRefID, b []PlayerGuildMemberRefID) []PlayerGuildMemberRefID {

	check := make(map[PlayerGuildMemberRefID]bool)
	deduped := make([]PlayerGuildMemberRefID, 0)
	for _, val := range a {
		check[val] = true
	}
	for _, val := range b {
		check[val] = true
	}

	for val := range check {
		deduped = append(deduped, val)
	}

	return deduped
}

func deduplicatePlayerEquipmentSetRefIDs(a []PlayerEquipmentSetRefID, b []PlayerEquipmentSetRefID) []PlayerEquipmentSetRefID {

	check := make(map[PlayerEquipmentSetRefID]bool)
	deduped := make([]PlayerEquipmentSetRefID, 0)
	for _, val := range a {
		check[val] = true
	}
	for _, val := range b {
		check[val] = true
	}

	for val := range check {
		deduped = append(deduped, val)
	}

	return deduped
}

func deduplicateEquipmentSetEquipmentRefIDs(a []EquipmentSetEquipmentRefID, b []EquipmentSetEquipmentRefID) []EquipmentSetEquipmentRefID {

	check := make(map[EquipmentSetEquipmentRefID]bool)
	deduped := make([]EquipmentSetEquipmentRefID, 0)
	for _, val := range a {
		check[val] = true
	}
	for _, val := range b {
		check[val] = true
	}

	for val := range check {
		deduped = append(deduped, val)
	}

	return deduped
}

func (engine Engine) allEquipmentSetIDs() []EquipmentSetID {
	var stateEquipmentSetIDs []EquipmentSetID
	for equipmentSetID := range engine.State.EquipmentSet {
		stateEquipmentSetIDs = append(stateEquipmentSetIDs, equipmentSetID)
	}
	var patchEquipmentSetIDs []EquipmentSetID
	for equipmentSetID := range engine.Patch.EquipmentSet {
		patchEquipmentSetIDs = append(patchEquipmentSetIDs, equipmentSetID)
	}
	return deduplicateEquipmentSetIDs(stateEquipmentSetIDs, patchEquipmentSetIDs)
}

func (engine Engine) allGearScoreIDs() []GearScoreID {
	var stateGearScoreIDs []GearScoreID
	for playerID := range engine.State.GearScore {
		stateGearScoreIDs = append(stateGearScoreIDs, playerID)
	}
	var patchGearScoreIDs []GearScoreID
	for playerID := range engine.Patch.GearScore {
		patchGearScoreIDs = append(patchGearScoreIDs, playerID)
	}
	return deduplicateGearScoreIDs(stateGearScoreIDs, patchGearScoreIDs)
}

func (engine Engine) allItemIDs() []ItemID {
	var stateItemIDs []ItemID
	for itemID := range engine.State.Item {
		stateItemIDs = append(stateItemIDs, itemID)
	}
	var patchItemIDs []ItemID
	for itemID := range engine.Patch.Item {
		patchItemIDs = append(patchItemIDs, itemID)
	}
	return deduplicateItemIDs(stateItemIDs, patchItemIDs)
}

func (engine Engine) allPositionIDs() []PositionID {
	var statePositionIDs []PositionID
	for positionID := range engine.State.Position {
		statePositionIDs = append(statePositionIDs, positionID)
	}
	var patchPositionIDs []PositionID
	for positionID := range engine.Patch.Position {
		patchPositionIDs = append(patchPositionIDs, positionID)
	}
	return deduplicatePositionIDs(statePositionIDs, patchPositionIDs)
}

func (engine Engine) allZoneIDs() []ZoneID {
	var stateZoneIDs []ZoneID
	for positionID := range engine.State.Zone {
		stateZoneIDs = append(stateZoneIDs, positionID)
	}
	var patchZoneIDs []ZoneID
	for positionID := range engine.Patch.Zone {
		patchZoneIDs = append(patchZoneIDs, positionID)
	}
	return deduplicateZoneIDs(stateZoneIDs, patchZoneIDs)
}

func (engine Engine) allZoneItemIDs() []ZoneItemID {
	var stateZoneItemIDs []ZoneItemID
	for positionID := range engine.State.ZoneItem {
		stateZoneItemIDs = append(stateZoneItemIDs, positionID)
	}
	var patchZoneItemIDs []ZoneItemID
	for positionID := range engine.Patch.ZoneItem {
		patchZoneItemIDs = append(patchZoneItemIDs, positionID)
	}
	return deduplicateZoneItemIDs(stateZoneItemIDs, patchZoneItemIDs)
}

func (engine Engine) allPlayerIDs() []PlayerID {
	var statePlayerIDs []PlayerID
	for playerID := range engine.State.Player {
		statePlayerIDs = append(statePlayerIDs, playerID)
	}
	var patchPlayerIDs []PlayerID
	for playerID := range engine.Patch.Player {
		patchPlayerIDs = append(patchPlayerIDs, playerID)
	}
	return deduplicatePlayerIDs(statePlayerIDs, patchPlayerIDs)
}

func (engine Engine) allPlayerTargetedByRefIDs() []PlayerTargetedByRefID {
	var statePlayerTargetedByRefIDs []PlayerTargetedByRefID
	for playerTargetedByRefID := range engine.State.PlayerTargetedByRef {
		statePlayerTargetedByRefIDs = append(statePlayerTargetedByRefIDs, playerTargetedByRefID)
	}
	var patchPlayerTargetedByRefIDs []PlayerTargetedByRefID
	for playerTargetedByRefID := range engine.Patch.PlayerTargetedByRef {
		patchPlayerTargetedByRefIDs = append(patchPlayerTargetedByRefIDs, playerTargetedByRefID)
	}
	return deduplicatePlayerTargetedByRefIDs(statePlayerTargetedByRefIDs, patchPlayerTargetedByRefIDs)
}

func (engine Engine) allPlayerTargetRefIDs() []PlayerTargetRefID {
	var statePlayerTargetRefIDs []PlayerTargetRefID
	for playerTargetRefID := range engine.State.PlayerTargetRef {
		statePlayerTargetRefIDs = append(statePlayerTargetRefIDs, playerTargetRefID)
	}
	var patchPlayerTargetRefIDs []PlayerTargetRefID
	for playerTargetRefID := range engine.Patch.PlayerTargetRef {
		patchPlayerTargetRefIDs = append(patchPlayerTargetRefIDs, playerTargetRefID)
	}
	return deduplicatePlayerTargetRefIDs(statePlayerTargetRefIDs, patchPlayerTargetRefIDs)
}

func (engine Engine) allItemBoundToRefIDs() []ItemBoundToRefID {
	var stateItemBoundToRefIDs []ItemBoundToRefID
	for itemBoundToRefID := range engine.State.ItemBoundToRef {
		stateItemBoundToRefIDs = append(stateItemBoundToRefIDs, itemBoundToRefID)
	}
	var patchItemBoundToRefIDs []ItemBoundToRefID
	for itemBoundToRefID := range engine.Patch.ItemBoundToRef {
		patchItemBoundToRefIDs = append(patchItemBoundToRefIDs, itemBoundToRefID)
	}
	return deduplicateItemBoundToRefIDs(stateItemBoundToRefIDs, patchItemBoundToRefIDs)
}

func (engine Engine) allPlayerGuildMemberRefIDs() []PlayerGuildMemberRefID {
	var statePlayerGuildMemberRefIDs []PlayerGuildMemberRefID
	for playerGuildMemberRefID := range engine.State.PlayerGuildMemberRef {
		statePlayerGuildMemberRefIDs = append(statePlayerGuildMemberRefIDs, playerGuildMemberRefID)
	}
	var patchPlayerGuildMemberRefIDs []PlayerGuildMemberRefID
	for playerGuildMemberRefID := range engine.Patch.PlayerGuildMemberRef {
		patchPlayerGuildMemberRefIDs = append(patchPlayerGuildMemberRefIDs, playerGuildMemberRefID)
	}
	return deduplicatePlayerGuildMemberRefIDs(statePlayerGuildMemberRefIDs, patchPlayerGuildMemberRefIDs)
}

func (engine Engine) allPlayerEquipmentSetRefIDs() []PlayerEquipmentSetRefID {
	var statePlayerEquipmentSetRefIDs []PlayerEquipmentSetRefID
	for playerEquipmentSetRefID := range engine.State.PlayerEquipmentSetRef {
		statePlayerEquipmentSetRefIDs = append(statePlayerEquipmentSetRefIDs, playerEquipmentSetRefID)
	}
	var patchPlayerEquipmentSetRefIDs []PlayerEquipmentSetRefID
	for playerEquipmentSetRefID := range engine.Patch.PlayerEquipmentSetRef {
		patchPlayerEquipmentSetRefIDs = append(patchPlayerEquipmentSetRefIDs, playerEquipmentSetRefID)
	}
	return deduplicatePlayerEquipmentSetRefIDs(statePlayerEquipmentSetRefIDs, patchPlayerEquipmentSetRefIDs)
}

func (engine Engine) allEquipmentSetEquipmentRefIDs() []EquipmentSetEquipmentRefID {
	var stateEquipmentSetEquipmentRefIDs []EquipmentSetEquipmentRefID
	for equipmentSetEquipmentRefID := range engine.State.EquipmentSetEquipmentRef {
		stateEquipmentSetEquipmentRefIDs = append(stateEquipmentSetEquipmentRefIDs, equipmentSetEquipmentRefID)
	}
	var patchEquipmentSetEquipmentRefIDs []EquipmentSetEquipmentRefID
	for equipmentSetEquipmentRefID := range engine.Patch.EquipmentSetEquipmentRef {
		patchEquipmentSetEquipmentRefIDs = append(patchEquipmentSetEquipmentRefIDs, equipmentSetEquipmentRefID)
	}
	return deduplicateEquipmentSetEquipmentRefIDs(stateEquipmentSetEquipmentRefIDs, patchEquipmentSetEquipmentRefIDs)
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

func mergePlayerTargetedByRefIDs(currentIDs, nextIDs []PlayerTargetedByRefID) []PlayerTargetedByRefID {
	ids := make([]PlayerTargetedByRefID, len(currentIDs))
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

func mergePlayerTargetRefIDs(currentIDs, nextIDs []PlayerTargetRefID) []PlayerTargetRefID {
	ids := make([]PlayerTargetRefID, len(currentIDs))
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

func mergeAnyOfPlayer_ZoneItemIDs(currentIDs, nextIDs []AnyOfPlayer_ZoneItemID) []AnyOfPlayer_ZoneItemID {
	ids := make([]AnyOfPlayer_ZoneItemID, len(currentIDs))
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

func mergeAnyOfPlayer_PositionIDs(currentIDs, nextIDs []AnyOfPlayer_PositionID) []AnyOfPlayer_PositionID {
	ids := make([]AnyOfPlayer_PositionID, len(currentIDs))
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

func mergeAnyOfItem_Player_ZoneItemIDs(currentIDs, nextIDs []AnyOfItem_Player_ZoneItemID) []AnyOfItem_Player_ZoneItemID {
	ids := make([]AnyOfItem_Player_ZoneItemID, len(currentIDs))
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
