package state

func deduplicateZoneItemIDs(a []ZoneItemID, b []ZoneItemID) []ZoneItemID {

	check := zoneItemCheckPool.Get().(map[ZoneItemID]bool)
	for k := range check {
		delete(check, k)
	}
	deduped := zoneItemIDSlicePool.Get().([]ZoneItemID)[:0]

	for _, val := range a {
		check[val] = true
	}
	for _, val := range b {
		check[val] = true
	}

	for val := range check {
		deduped = append(deduped, val)
	}

	zoneItemCheckPool.Put(check)

	return deduped
}

func deduplicateZoneIDs(a []ZoneID, b []ZoneID) []ZoneID {

	check := zoneCheckPool.Get().(map[ZoneID]bool)
	for k := range check {
		delete(check, k)
	}
	deduped := zoneIDSlicePool.Get().([]ZoneID)[:0]
	for _, val := range a {
		check[val] = true
	}
	for _, val := range b {
		check[val] = true
	}

	for val := range check {
		deduped = append(deduped, val)
	}

	zoneCheckPool.Put(check)

	return deduped
}

func deduplicatePlayerIDs(a []PlayerID, b []PlayerID) []PlayerID {

	check := playerCheckPool.Get().(map[PlayerID]bool)
	for k := range check {
		delete(check, k)
	}
	deduped := playerIDSlicePool.Get().([]PlayerID)[:0]
	for _, val := range a {
		check[val] = true
	}
	for _, val := range b {
		check[val] = true
	}

	for val := range check {
		deduped = append(deduped, val)
	}

	playerCheckPool.Put(check)

	return deduped
}

func deduplicatePositionIDs(a []PositionID, b []PositionID) []PositionID {

	check := positionCheckPool.Get().(map[PositionID]bool)
	for k := range check {
		delete(check, k)
	}
	deduped := positionIDSlicePool.Get().([]PositionID)[:0]
	for _, val := range a {
		check[val] = true
	}
	for _, val := range b {
		check[val] = true
	}

	for val := range check {
		deduped = append(deduped, val)
	}

	positionCheckPool.Put(check)

	return deduped
}

func deduplicateItemIDs(a []ItemID, b []ItemID) []ItemID {

	check := itemCheckPool.Get().(map[ItemID]bool)
	for k := range check {
		delete(check, k)
	}
	deduped := itemIDSlicePool.Get().([]ItemID)[:0]
	for _, val := range a {
		check[val] = true
	}
	for _, val := range b {
		check[val] = true
	}

	for val := range check {
		deduped = append(deduped, val)
	}

	itemCheckPool.Put(check)

	return deduped
}

func deduplicateGearScoreIDs(a []GearScoreID, b []GearScoreID) []GearScoreID {

	check := gearScoreCheckPool.Get().(map[GearScoreID]bool)
	for k := range check {
		delete(check, k)
	}
	deduped := gearScoreIDSlicePool.Get().([]GearScoreID)[:0]
	for _, val := range a {
		check[val] = true
	}
	for _, val := range b {
		check[val] = true
	}

	for val := range check {
		deduped = append(deduped, val)
	}

	gearScoreCheckPool.Put(check)

	return deduped
}

func deduplicateEquipmentSetIDs(a []EquipmentSetID, b []EquipmentSetID) []EquipmentSetID {

	check := equipmentSetCheckPool.Get().(map[EquipmentSetID]bool)
	for k := range check {
		delete(check, k)
	}
	deduped := equipmentSetIDSlicePool.Get().([]EquipmentSetID)[:0]
	for _, val := range a {
		check[val] = true
	}
	for _, val := range b {
		check[val] = true
	}

	for val := range check {
		deduped = append(deduped, val)
	}

	equipmentSetCheckPool.Put(check)

	return deduped
}

func deduplicatePlayerTargetedByRefIDs(a []PlayerTargetedByRefID, b []PlayerTargetedByRefID) []PlayerTargetedByRefID {

	check := playerTargetedByRefCheckPool.Get().(map[PlayerTargetedByRefID]bool)
	for k := range check {
		delete(check, k)
	}
	deduped := playerTargetedByRefIDSlicePool.Get().([]PlayerTargetedByRefID)[:0]
	for _, val := range a {
		check[val] = true
	}
	for _, val := range b {
		check[val] = true
	}

	for val := range check {
		deduped = append(deduped, val)
	}

	playerTargetedByRefCheckPool.Put(check)

	return deduped
}

func deduplicatePlayerTargetRefIDs(a []PlayerTargetRefID, b []PlayerTargetRefID) []PlayerTargetRefID {

	check := playerTargetRefCheckPool.Get().(map[PlayerTargetRefID]bool)
	for k := range check {
		delete(check, k)
	}
	deduped := playerTargetRefIDSlicePool.Get().([]PlayerTargetRefID)[:0]
	for _, val := range a {
		check[val] = true
	}
	for _, val := range b {
		check[val] = true
	}

	for val := range check {
		deduped = append(deduped, val)
	}

	playerTargetRefCheckPool.Put(check)

	return deduped
}

func deduplicateItemBoundToRefIDs(a []ItemBoundToRefID, b []ItemBoundToRefID) []ItemBoundToRefID {

	check := itemBoundToRefCheckPool.Get().(map[ItemBoundToRefID]bool)
	for k := range check {
		delete(check, k)
	}
	deduped := itemBoundToRefIDSlicePool.Get().([]ItemBoundToRefID)[:0]
	for _, val := range a {
		check[val] = true
	}
	for _, val := range b {
		check[val] = true
	}

	for val := range check {
		deduped = append(deduped, val)
	}

	itemBoundToRefCheckPool.Put(check)

	return deduped
}

func deduplicatePlayerGuildMemberRefIDs(a []PlayerGuildMemberRefID, b []PlayerGuildMemberRefID) []PlayerGuildMemberRefID {

	check := playerGuildMemberRefCheckPool.Get().(map[PlayerGuildMemberRefID]bool)
	for k := range check {
		delete(check, k)
	}
	deduped := playerGuildMemberRefIDSlicePool.Get().([]PlayerGuildMemberRefID)[:0]
	for _, val := range a {
		check[val] = true
	}
	for _, val := range b {
		check[val] = true
	}

	for val := range check {
		deduped = append(deduped, val)
	}

	playerGuildMemberRefCheckPool.Put(check)

	return deduped
}

func deduplicatePlayerEquipmentSetRefIDs(a []PlayerEquipmentSetRefID, b []PlayerEquipmentSetRefID) []PlayerEquipmentSetRefID {

	check := playerEquipmentSetRefCheckPool.Get().(map[PlayerEquipmentSetRefID]bool)
	for k := range check {
		delete(check, k)
	}
	deduped := playerEquipmentSetRefIDSlicePool.Get().([]PlayerEquipmentSetRefID)[:0]
	for _, val := range a {
		check[val] = true
	}
	for _, val := range b {
		check[val] = true
	}

	for val := range check {
		deduped = append(deduped, val)
	}

	playerEquipmentSetRefCheckPool.Put(check)

	return deduped
}

func deduplicateEquipmentSetEquipmentRefIDs(a []EquipmentSetEquipmentRefID, b []EquipmentSetEquipmentRefID) []EquipmentSetEquipmentRefID {

	check := equipmentSetEquipmentRefCheckPool.Get().(map[EquipmentSetEquipmentRefID]bool)
	for k := range check {
		delete(check, k)
	}
	deduped := equipmentSetEquipmentRefIDSlicePool.Get().([]EquipmentSetEquipmentRefID)[:0]
	for _, val := range a {
		check[val] = true
	}
	for _, val := range b {
		check[val] = true
	}

	for val := range check {
		deduped = append(deduped, val)
	}

	equipmentSetEquipmentRefCheckPool.Put(check)

	return deduped
}

func (engine Engine) allEquipmentSetIDs() []EquipmentSetID {
	stateEquipmentSetIDs := equipmentSetIDSlicePool.Get().([]EquipmentSetID)[:0]
	for equipmentSetID := range engine.State.EquipmentSet {
		stateEquipmentSetIDs = append(stateEquipmentSetIDs, equipmentSetID)
	}
	patchEquipmentSetIDs := equipmentSetIDSlicePool.Get().([]EquipmentSetID)[:0]
	for equipmentSetID := range engine.Patch.EquipmentSet {
		patchEquipmentSetIDs = append(patchEquipmentSetIDs, equipmentSetID)
	}
	dedupedIDs := deduplicateEquipmentSetIDs(stateEquipmentSetIDs, patchEquipmentSetIDs)

	equipmentSetIDSlicePool.Put(stateEquipmentSetIDs)
	equipmentSetIDSlicePool.Put(patchEquipmentSetIDs)

	return dedupedIDs
}

func (engine Engine) allGearScoreIDs() []GearScoreID {
	stateGearScoreIDs := gearScoreIDSlicePool.Get().([]GearScoreID)[:0]
	for gearScoreID := range engine.State.GearScore {
		stateGearScoreIDs = append(stateGearScoreIDs, gearScoreID)
	}
	patchGearScoreIDs := gearScoreIDSlicePool.Get().([]GearScoreID)[:0]
	for gearScoreID := range engine.Patch.GearScore {
		patchGearScoreIDs = append(patchGearScoreIDs, gearScoreID)
	}
	dedupedIDs := deduplicateGearScoreIDs(stateGearScoreIDs, patchGearScoreIDs)

	gearScoreIDSlicePool.Put(stateGearScoreIDs)
	gearScoreIDSlicePool.Put(patchGearScoreIDs)

	return dedupedIDs
}

func (engine Engine) allItemIDs() []ItemID {
	stateItemIDs := itemIDSlicePool.Get().([]ItemID)[:0]
	for itemID := range engine.State.Item {
		stateItemIDs = append(stateItemIDs, itemID)
	}
	patchItemIDs := itemIDSlicePool.Get().([]ItemID)[:0]
	for itemID := range engine.Patch.Item {
		patchItemIDs = append(patchItemIDs, itemID)
	}
	dedupedIDs := deduplicateItemIDs(stateItemIDs, patchItemIDs)

	itemIDSlicePool.Put(stateItemIDs)
	itemIDSlicePool.Put(patchItemIDs)

	return dedupedIDs
}

func (engine Engine) allPositionIDs() []PositionID {
	statePositionIDs := positionIDSlicePool.Get().([]PositionID)[:0]
	for positionID := range engine.State.Position {
		statePositionIDs = append(statePositionIDs, positionID)
	}
	patchPositionIDs := positionIDSlicePool.Get().([]PositionID)[:0]
	for positionID := range engine.Patch.Position {
		patchPositionIDs = append(patchPositionIDs, positionID)
	}
	dedupedIDs := deduplicatePositionIDs(statePositionIDs, patchPositionIDs)

	positionIDSlicePool.Put(statePositionIDs)
	positionIDSlicePool.Put(patchPositionIDs)

	return dedupedIDs
}

func (engine Engine) allZoneIDs() []ZoneID {
	stateZoneIDs := zoneIDSlicePool.Get().([]ZoneID)[:0]
	for zoneID := range engine.State.Zone {
		stateZoneIDs = append(stateZoneIDs, zoneID)
	}
	patchZoneIDs := zoneIDSlicePool.Get().([]ZoneID)[:0]
	for zoneID := range engine.Patch.Zone {
		patchZoneIDs = append(patchZoneIDs, zoneID)
	}
	dedupedIDs := deduplicateZoneIDs(stateZoneIDs, patchZoneIDs)

	zoneIDSlicePool.Put(stateZoneIDs)
	zoneIDSlicePool.Put(patchZoneIDs)

	return dedupedIDs
}

func (engine Engine) allZoneItemIDs() []ZoneItemID {
	stateZoneItemIDs := zoneItemIDSlicePool.Get().([]ZoneItemID)[:0]
	for zoneItemID := range engine.State.ZoneItem {
		stateZoneItemIDs = append(stateZoneItemIDs, zoneItemID)
	}
	patchZoneItemIDs := zoneItemIDSlicePool.Get().([]ZoneItemID)[:0]
	for zoneItemID := range engine.Patch.ZoneItem {
		patchZoneItemIDs = append(patchZoneItemIDs, zoneItemID)
	}
	dedupedIDs := deduplicateZoneItemIDs(stateZoneItemIDs, patchZoneItemIDs)

	zoneItemIDSlicePool.Put(stateZoneItemIDs)
	zoneItemIDSlicePool.Put(patchZoneItemIDs)

	return dedupedIDs
}

func (engine Engine) allPlayerIDs() []PlayerID {
	statePlayerIDs := playerIDSlicePool.Get().([]PlayerID)[:0]
	for playerID := range engine.State.Player {
		statePlayerIDs = append(statePlayerIDs, playerID)
	}
	patchPlayerIDs := playerIDSlicePool.Get().([]PlayerID)[:0]
	for playerID := range engine.Patch.Player {
		patchPlayerIDs = append(patchPlayerIDs, playerID)
	}
	dedupedIDs := deduplicatePlayerIDs(statePlayerIDs, patchPlayerIDs)

	playerIDSlicePool.Put(statePlayerIDs)
	playerIDSlicePool.Put(patchPlayerIDs)

	return dedupedIDs
}

func (engine Engine) allPlayerTargetedByRefIDs() []PlayerTargetedByRefID {
	statePlayerTargetedByRefIDs := playerTargetedByRefIDSlicePool.Get().([]PlayerTargetedByRefID)[:0]
	for playerTargetedByRefID := range engine.State.PlayerTargetedByRef {
		statePlayerTargetedByRefIDs = append(statePlayerTargetedByRefIDs, playerTargetedByRefID)
	}
	patchPlayerTargetedByRefIDs := playerTargetedByRefIDSlicePool.Get().([]PlayerTargetedByRefID)[:0]
	for playerTargetedByRefID := range engine.Patch.PlayerTargetedByRef {
		patchPlayerTargetedByRefIDs = append(patchPlayerTargetedByRefIDs, playerTargetedByRefID)
	}
	dedupedIDs := deduplicatePlayerTargetedByRefIDs(statePlayerTargetedByRefIDs, patchPlayerTargetedByRefIDs)

	playerTargetedByRefIDSlicePool.Put(statePlayerTargetedByRefIDs)
	playerTargetedByRefIDSlicePool.Put(patchPlayerTargetedByRefIDs)

	return dedupedIDs
}

func (engine Engine) allPlayerTargetRefIDs() []PlayerTargetRefID {
	statePlayerTargetRefIDs := playerTargetRefIDSlicePool.Get().([]PlayerTargetRefID)[:0]
	for playerTargetRefID := range engine.State.PlayerTargetRef {
		statePlayerTargetRefIDs = append(statePlayerTargetRefIDs, playerTargetRefID)
	}
	patchPlayerTargetRefIDs := playerTargetRefIDSlicePool.Get().([]PlayerTargetRefID)[:0]
	for playerTargetRefID := range engine.Patch.PlayerTargetRef {
		patchPlayerTargetRefIDs = append(patchPlayerTargetRefIDs, playerTargetRefID)
	}
	dedupedIDs := deduplicatePlayerTargetRefIDs(statePlayerTargetRefIDs, patchPlayerTargetRefIDs)

	playerTargetRefIDSlicePool.Put(statePlayerTargetRefIDs)
	playerTargetRefIDSlicePool.Put(patchPlayerTargetRefIDs)

	return dedupedIDs
}

func (engine Engine) allItemBoundToRefIDs() []ItemBoundToRefID {
	stateItemBoundToRefIDs := itemBoundToRefIDSlicePool.Get().([]ItemBoundToRefID)[:0]
	for itemBoundToRefID := range engine.State.ItemBoundToRef {
		stateItemBoundToRefIDs = append(stateItemBoundToRefIDs, itemBoundToRefID)
	}
	patchItemBoundToRefIDs := itemBoundToRefIDSlicePool.Get().([]ItemBoundToRefID)[:0]
	for itemBoundToRefID := range engine.Patch.ItemBoundToRef {
		patchItemBoundToRefIDs = append(patchItemBoundToRefIDs, itemBoundToRefID)
	}
	dedupedIDs := deduplicateItemBoundToRefIDs(stateItemBoundToRefIDs, patchItemBoundToRefIDs)

	itemBoundToRefIDSlicePool.Put(stateItemBoundToRefIDs)
	itemBoundToRefIDSlicePool.Put(patchItemBoundToRefIDs)

	return dedupedIDs
}

func (engine Engine) allPlayerGuildMemberRefIDs() []PlayerGuildMemberRefID {
	statePlayerGuildMemberRefIDs := playerGuildMemberRefIDSlicePool.Get().([]PlayerGuildMemberRefID)[:0]
	for playerGuildMemberRefID := range engine.State.PlayerGuildMemberRef {
		statePlayerGuildMemberRefIDs = append(statePlayerGuildMemberRefIDs, playerGuildMemberRefID)
	}
	patchPlayerGuildMemberRefIDs := playerGuildMemberRefIDSlicePool.Get().([]PlayerGuildMemberRefID)[:0]
	for playerGuildMemberRefID := range engine.Patch.PlayerGuildMemberRef {
		patchPlayerGuildMemberRefIDs = append(patchPlayerGuildMemberRefIDs, playerGuildMemberRefID)
	}
	dedupedIDs := deduplicatePlayerGuildMemberRefIDs(statePlayerGuildMemberRefIDs, patchPlayerGuildMemberRefIDs)

	playerGuildMemberRefIDSlicePool.Put(statePlayerGuildMemberRefIDs)
	playerGuildMemberRefIDSlicePool.Put(patchPlayerGuildMemberRefIDs)

	return dedupedIDs
}

func (engine Engine) allPlayerEquipmentSetRefIDs() []PlayerEquipmentSetRefID {
	statePlayerEquipmentSetRefIDs := playerEquipmentSetRefIDSlicePool.Get().([]PlayerEquipmentSetRefID)[:0]
	for playerEquipmentSetRefID := range engine.State.PlayerEquipmentSetRef {
		statePlayerEquipmentSetRefIDs = append(statePlayerEquipmentSetRefIDs, playerEquipmentSetRefID)
	}
	patchPlayerEquipmentSetRefIDs := playerEquipmentSetRefIDSlicePool.Get().([]PlayerEquipmentSetRefID)[:0]
	for playerEquipmentSetRefID := range engine.Patch.PlayerEquipmentSetRef {
		patchPlayerEquipmentSetRefIDs = append(patchPlayerEquipmentSetRefIDs, playerEquipmentSetRefID)
	}
	dedupedIDs := deduplicatePlayerEquipmentSetRefIDs(statePlayerEquipmentSetRefIDs, patchPlayerEquipmentSetRefIDs)

	playerEquipmentSetRefIDSlicePool.Put(statePlayerEquipmentSetRefIDs)
	playerEquipmentSetRefIDSlicePool.Put(patchPlayerEquipmentSetRefIDs)

	return dedupedIDs
}

func (engine Engine) allEquipmentSetEquipmentRefIDs() []EquipmentSetEquipmentRefID {
	stateEquipmentSetEquipmentRefIDs := equipmentSetEquipmentRefIDSlicePool.Get().([]EquipmentSetEquipmentRefID)[:0]
	for equipmentSetEquipmentRefID := range engine.State.EquipmentSetEquipmentRef {
		stateEquipmentSetEquipmentRefIDs = append(stateEquipmentSetEquipmentRefIDs, equipmentSetEquipmentRefID)
	}
	patchEquipmentSetEquipmentRefIDs := equipmentSetEquipmentRefIDSlicePool.Get().([]EquipmentSetEquipmentRefID)[:0]
	for equipmentSetEquipmentRefID := range engine.Patch.EquipmentSetEquipmentRef {
		patchEquipmentSetEquipmentRefIDs = append(patchEquipmentSetEquipmentRefIDs, equipmentSetEquipmentRefID)
	}
	dedupedIDs := deduplicateEquipmentSetEquipmentRefIDs(stateEquipmentSetEquipmentRefIDs, patchEquipmentSetEquipmentRefIDs)

	equipmentSetEquipmentRefIDSlicePool.Put(stateEquipmentSetEquipmentRefIDs)
	equipmentSetEquipmentRefIDSlicePool.Put(patchEquipmentSetEquipmentRefIDs)

	return dedupedIDs
}
