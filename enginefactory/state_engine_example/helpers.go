package state

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

func (se Engine) allGearScoreIDs() []GearScoreID {
	var stateGearScoreIDs []GearScoreID
	for gearScoreID := range se.State.GearScore {
		stateGearScoreIDs = append(stateGearScoreIDs, gearScoreID)
	}
	var patchGearScoreIDs []GearScoreID
	for gearScoreID := range se.Patch.GearScore {
		patchGearScoreIDs = append(patchGearScoreIDs, gearScoreID)
	}
	return deduplicateGearScoreIDs(stateGearScoreIDs, patchGearScoreIDs)
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
	return deduplicateItemIDs(stateItemIDs, patchItemIDs)
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
	return deduplicatePlayerIDs(statePlayerIDs, patchPlayerIDs)
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
	return deduplicatePositionIDs(statePositionIDs, patchPositionIDs)
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
	return deduplicateZoneIDs(stateZoneIDs, patchZoneIDs)
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
	return deduplicateZoneItemIDs(stateZoneItemIDs, patchZoneItemIDs)
}

// TODO extensive testing
func (se *Engine) evalItemBoundToElementRef(itemData itemCore) *ElementReference {
	if itemData.BoundTo.id != 0 {
		_, hasUpdated := se.Patch.Player[itemData.BoundTo.id]
		justGotSet := se.State.Item[itemData.ID].BoundTo.id == 0
		operationKind := OperationKindRefUnchanged
		if hasUpdated || justGotSet {
			operationKind = OperationKindUpdate
		}
		return &ElementReference{ID: int(itemData.BoundTo.id), ElementKind: ElementKindPlayer, OperationKind_: operationKind}
	} else if se.State.Item[itemData.ID].BoundTo.id != 0 {
		return &ElementReference{ID: 0, ElementKind: ElementKindPlayer, OperationKind_: OperationKindDelete}
	}
	return nil
}

// TODO extensive testing
func (se *Engine) evalPlayerGuildMembersElementRefs(playerData playerCore) []ElementReference {
	var refs []ElementReference
	var evalProgressionIndex int
	for _, previousGuildMember := range se.State.Player[playerData.ID].GuildMembers {
		refGotDeleted := true
		for _, currentGuildMember := range playerData.GuildMembers {
			if previousGuildMember.id == currentGuildMember.id {
				refGotDeleted = false
			}
		}
		if refGotDeleted {
			ref := ElementReference{ID: 0, ElementKind: ElementKindPlayer, OperationKind_: OperationKindDelete}
			refs = append(refs, ref)
		} else {
			ref := ElementReference{ID: int(previousGuildMember.id), ElementKind: ElementKindPlayer, OperationKind_: OperationKindRefUnchanged}
			refs = append(refs, ref)
			evalProgressionIndex += 1
		}
	}
	for _, guildMember := range playerData.GuildMembers[evalProgressionIndex:] {
		ref := ElementReference{ID: int(guildMember.id), ElementKind: ElementKindPlayer, OperationKind_: OperationKindUpdate}
		refs = append(refs, ref)
	}
	return refs
}
