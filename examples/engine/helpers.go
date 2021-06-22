package state

import (
	"sync"
)

var zoneItemCheckPool = sync.Pool{
	New: func() interface{} { return make(map[ZoneItemID]bool) },
}
var zoneItemSlicePool = sync.Pool{
	New: func() interface{} { return make([]ZoneItemID, 10) },
}

func deduplicateZoneItemIDs(a []ZoneItemID, b []ZoneItemID) []ZoneItemID {

	check := zoneItemCheckPool.Get().(map[ZoneItemID]bool)
	for k := range check {
		delete(check, k)
	}
	deduped := zoneItemSlicePool.Get().([]ZoneItemID)[:0]

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
	zoneItemSlicePool.Put(deduped)

	return deduped
}

var zoneCheckPool = sync.Pool{
	New: func() interface{} { return make(map[ZoneID]bool) },
}
var zoneSlicePool = sync.Pool{
	New: func() interface{} { return make([]ZoneID, 10) },
}

func deduplicateZoneIDs(a []ZoneID, b []ZoneID) []ZoneID {

	check := zoneCheckPool.Get().(map[ZoneID]bool)
	for k := range check {
		delete(check, k)
	}
	deduped := zoneSlicePool.Get().([]ZoneID)[:0]
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
	zoneSlicePool.Put(deduped)

	return deduped
}

var playerCheckPool = sync.Pool{
	New: func() interface{} { return make(map[PlayerID]bool) },
}
var playerSlicePool = sync.Pool{
	New: func() interface{} { return make([]PlayerID, 10) },
}

func deduplicatePlayerIDs(a []PlayerID, b []PlayerID) []PlayerID {

	check := playerCheckPool.Get().(map[PlayerID]bool)
	for k := range check {
		delete(check, k)
	}
	deduped := playerSlicePool.Get().([]PlayerID)[:0]
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
	playerSlicePool.Put(deduped)

	return deduped
}

var positionCheckPool = sync.Pool{
	New: func() interface{} { return make(map[PositionID]bool) },
}
var positionSlicePool = sync.Pool{
	New: func() interface{} { return make([]PositionID, 10) },
}

func deduplicatePositionIDs(a []PositionID, b []PositionID) []PositionID {

	check := positionCheckPool.Get().(map[PositionID]bool)
	for k := range check {
		delete(check, k)
	}
	deduped := positionSlicePool.Get().([]PositionID)[:0]
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
	positionSlicePool.Put(deduped)

	return deduped
}

var itemCheckPool = sync.Pool{
	New: func() interface{} { return make(map[ItemID]bool) },
}
var itemSlicePool = sync.Pool{
	New: func() interface{} { return make([]ItemID, 10) },
}

func deduplicateItemIDs(a []ItemID, b []ItemID) []ItemID {

	check := itemCheckPool.Get().(map[ItemID]bool)
	for k := range check {
		delete(check, k)
	}
	deduped := itemSlicePool.Get().([]ItemID)[:0]
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
	itemSlicePool.Put(deduped)

	return deduped
}

var gearScoreCheckPool = sync.Pool{
	New: func() interface{} { return make(map[GearScoreID]bool) },
}
var gearScoreSlicePool = sync.Pool{
	New: func() interface{} { return make([]GearScoreID, 10) },
}

func deduplicateGearScoreIDs(a []GearScoreID, b []GearScoreID) []GearScoreID {

	check := gearScoreCheckPool.Get().(map[GearScoreID]bool)
	for k := range check {
		delete(check, k)
	}
	deduped := gearScoreSlicePool.Get().([]GearScoreID)[:0]
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
	gearScoreSlicePool.Put(deduped)

	return deduped
}

var equipmentSetCheckPool = sync.Pool{
	New: func() interface{} { return make(map[EquipmentSetID]bool) },
}
var equipmentSetSlicePool = sync.Pool{
	New: func() interface{} { return make([]EquipmentSetID, 10) },
}

func deduplicateEquipmentSetIDs(a []EquipmentSetID, b []EquipmentSetID) []EquipmentSetID {

	check := equipmentSetCheckPool.Get().(map[EquipmentSetID]bool)
	for k := range check {
		delete(check, k)
	}
	deduped := equipmentSetSlicePool.Get().([]EquipmentSetID)[:0]
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
	equipmentSetSlicePool.Put(deduped)

	return deduped
}

var playerTargetedByRefCheckPool = sync.Pool{
	New: func() interface{} { return make(map[PlayerTargetedByRefID]bool) },
}
var playerTargetedByRefSlicePool = sync.Pool{
	New: func() interface{} { return make([]PlayerTargetedByRefID, 10) },
}

func deduplicatePlayerTargetedByRefIDs(a []PlayerTargetedByRefID, b []PlayerTargetedByRefID) []PlayerTargetedByRefID {

	check := playerTargetedByRefCheckPool.Get().(map[PlayerTargetedByRefID]bool)
	for k := range check {
		delete(check, k)
	}
	deduped := playerTargetedByRefSlicePool.Get().([]PlayerTargetedByRefID)[:0]
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
	playerTargetedByRefSlicePool.Put(deduped)

	return deduped
}

var playerTargetRefCheckPool = sync.Pool{
	New: func() interface{} { return make(map[PlayerTargetRefID]bool) },
}
var playerTargetRefSlicePool = sync.Pool{
	New: func() interface{} { return make([]PlayerTargetRefID, 10) },
}

func deduplicatePlayerTargetRefIDs(a []PlayerTargetRefID, b []PlayerTargetRefID) []PlayerTargetRefID {

	check := playerTargetRefCheckPool.Get().(map[PlayerTargetRefID]bool)
	for k := range check {
		delete(check, k)
	}
	deduped := playerTargetRefSlicePool.Get().([]PlayerTargetRefID)[:0]
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
	playerTargetRefSlicePool.Put(deduped)

	return deduped
}

var itemBoundToRefCheckPool = sync.Pool{
	New: func() interface{} { return make(map[ItemBoundToRefID]bool) },
}
var itemBoundToRefSlicePool = sync.Pool{
	New: func() interface{} { return make([]ItemBoundToRefID, 10) },
}

func deduplicateItemBoundToRefIDs(a []ItemBoundToRefID, b []ItemBoundToRefID) []ItemBoundToRefID {

	check := itemBoundToRefCheckPool.Get().(map[ItemBoundToRefID]bool)
	for k := range check {
		delete(check, k)
	}
	deduped := itemBoundToRefSlicePool.Get().([]ItemBoundToRefID)[:0]
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
	itemBoundToRefSlicePool.Put(deduped)

	return deduped
}

var playerGuildMemberRefCheckPool = sync.Pool{
	New: func() interface{} { return make(map[PlayerGuildMemberRefID]bool) },
}
var playerGuildMemberRefSlicePool = sync.Pool{
	New: func() interface{} { return make([]PlayerGuildMemberRefID, 10) },
}

func deduplicatePlayerGuildMemberRefIDs(a []PlayerGuildMemberRefID, b []PlayerGuildMemberRefID) []PlayerGuildMemberRefID {

	check := playerGuildMemberRefCheckPool.Get().(map[PlayerGuildMemberRefID]bool)
	for k := range check {
		delete(check, k)
	}
	deduped := playerGuildMemberRefSlicePool.Get().([]PlayerGuildMemberRefID)[:0]
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
	playerGuildMemberRefSlicePool.Put(deduped)

	return deduped
}

var playerEquipmentSetRefCheckPool = sync.Pool{
	New: func() interface{} { return make(map[PlayerEquipmentSetRefID]bool) },
}
var playerEquipmentSetRefSlicePool = sync.Pool{
	New: func() interface{} { return make([]PlayerEquipmentSetRefID, 10) },
}

func deduplicatePlayerEquipmentSetRefIDs(a []PlayerEquipmentSetRefID, b []PlayerEquipmentSetRefID) []PlayerEquipmentSetRefID {

	check := playerEquipmentSetRefCheckPool.Get().(map[PlayerEquipmentSetRefID]bool)
	for k := range check {
		delete(check, k)
	}
	deduped := playerEquipmentSetRefSlicePool.Get().([]PlayerEquipmentSetRefID)[:0]
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
	playerEquipmentSetRefSlicePool.Put(deduped)

	return deduped
}

var equipmentSetEquipmentRefCheckPool = sync.Pool{
	New: func() interface{} { return make(map[EquipmentSetEquipmentRefID]bool) },
}
var equipmentSetEquipmentRefSlicePool = sync.Pool{
	New: func() interface{} { return make([]EquipmentSetEquipmentRefID, 10) },
}

func deduplicateEquipmentSetEquipmentRefIDs(a []EquipmentSetEquipmentRefID, b []EquipmentSetEquipmentRefID) []EquipmentSetEquipmentRefID {

	check := equipmentSetEquipmentRefCheckPool.Get().(map[EquipmentSetEquipmentRefID]bool)
	for k := range check {
		delete(check, k)
	}
	deduped := equipmentSetEquipmentRefSlicePool.Get().([]EquipmentSetEquipmentRefID)[:0]
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
	equipmentSetEquipmentRefSlicePool.Put(deduped)

	return deduped
}

func (engine Engine) allEquipmentSetIDs() []EquipmentSetID {
	stateEquipmentSetIDs := equipmentSetSlicePool.Get().([]EquipmentSetID)[:0]
	for equipmentSetID := range engine.State.EquipmentSet {
		stateEquipmentSetIDs = append(stateEquipmentSetIDs, equipmentSetID)
	}
	patchEquipmentSetIDs := equipmentSetSlicePool.Get().([]EquipmentSetID)[:0]
	for equipmentSetID := range engine.Patch.EquipmentSet {
		patchEquipmentSetIDs = append(patchEquipmentSetIDs, equipmentSetID)
	}
	dedupedIDs := deduplicateEquipmentSetIDs(stateEquipmentSetIDs, patchEquipmentSetIDs)

	equipmentSetSlicePool.Put(stateEquipmentSetIDs)
	equipmentSetSlicePool.Put(patchEquipmentSetIDs)

	return dedupedIDs
}

func (engine Engine) allGearScoreIDs() []GearScoreID {
	stateGearScoreIDs := gearScoreSlicePool.Get().([]GearScoreID)[:0]
	for gearScoreID := range engine.State.GearScore {
		stateGearScoreIDs = append(stateGearScoreIDs, gearScoreID)
	}
	patchGearScoreIDs := gearScoreSlicePool.Get().([]GearScoreID)[:0]
	for gearScoreID := range engine.Patch.GearScore {
		patchGearScoreIDs = append(patchGearScoreIDs, gearScoreID)
	}
	dedupedIDs := deduplicateGearScoreIDs(stateGearScoreIDs, patchGearScoreIDs)

	gearScoreSlicePool.Put(stateGearScoreIDs)
	gearScoreSlicePool.Put(patchGearScoreIDs)

	return dedupedIDs
}

func (engine Engine) allItemIDs() []ItemID {
	stateItemIDs := itemSlicePool.Get().([]ItemID)[:0]
	for itemID := range engine.State.Item {
		stateItemIDs = append(stateItemIDs, itemID)
	}
	patchItemIDs := itemSlicePool.Get().([]ItemID)[:0]
	for itemID := range engine.Patch.Item {
		patchItemIDs = append(patchItemIDs, itemID)
	}
	dedupedIDs := deduplicateItemIDs(stateItemIDs, patchItemIDs)

	itemSlicePool.Put(stateItemIDs)
	itemSlicePool.Put(patchItemIDs)

	return dedupedIDs
}

func (engine Engine) allPositionIDs() []PositionID {
	statePositionIDs := positionSlicePool.Get().([]PositionID)[:0]
	for positionID := range engine.State.Position {
		statePositionIDs = append(statePositionIDs, positionID)
	}
	patchPositionIDs := positionSlicePool.Get().([]PositionID)[:0]
	for positionID := range engine.Patch.Position {
		patchPositionIDs = append(patchPositionIDs, positionID)
	}
	dedupedIDs := deduplicatePositionIDs(statePositionIDs, patchPositionIDs)

	positionSlicePool.Put(statePositionIDs)
	positionSlicePool.Put(patchPositionIDs)

	return dedupedIDs
}

func (engine Engine) allZoneIDs() []ZoneID {
	stateZoneIDs := zoneSlicePool.Get().([]ZoneID)[:0]
	for zoneID := range engine.State.Zone {
		stateZoneIDs = append(stateZoneIDs, zoneID)
	}
	patchZoneIDs := zoneSlicePool.Get().([]ZoneID)[:0]
	for zoneID := range engine.Patch.Zone {
		patchZoneIDs = append(patchZoneIDs, zoneID)
	}
	dedupedIDs := deduplicateZoneIDs(stateZoneIDs, patchZoneIDs)

	zoneSlicePool.Put(stateZoneIDs)
	zoneSlicePool.Put(patchZoneIDs)

	return dedupedIDs
}

func (engine Engine) allZoneItemIDs() []ZoneItemID {
	stateZoneItemIDs := zoneItemSlicePool.Get().([]ZoneItemID)[:0]
	for zoneItemID := range engine.State.ZoneItem {
		stateZoneItemIDs = append(stateZoneItemIDs, zoneItemID)
	}
	patchZoneItemIDs := zoneItemSlicePool.Get().([]ZoneItemID)[:0]
	for zoneItemID := range engine.Patch.ZoneItem {
		patchZoneItemIDs = append(patchZoneItemIDs, zoneItemID)
	}
	dedupedIDs := deduplicateZoneItemIDs(stateZoneItemIDs, patchZoneItemIDs)

	zoneItemSlicePool.Put(stateZoneItemIDs)
	zoneItemSlicePool.Put(patchZoneItemIDs)

	return dedupedIDs
}

func (engine Engine) allPlayerIDs() []PlayerID {
	statePlayerIDs := playerSlicePool.Get().([]PlayerID)[:0]
	for playerID := range engine.State.Player {
		statePlayerIDs = append(statePlayerIDs, playerID)
	}
	patchPlayerIDs := playerSlicePool.Get().([]PlayerID)[:0]
	for playerID := range engine.Patch.Player {
		patchPlayerIDs = append(patchPlayerIDs, playerID)
	}
	dedupedIDs := deduplicatePlayerIDs(statePlayerIDs, patchPlayerIDs)

	playerSlicePool.Put(statePlayerIDs)
	playerSlicePool.Put(patchPlayerIDs)

	return dedupedIDs
}

func (engine Engine) allPlayerTargetedByRefIDs() []PlayerTargetedByRefID {
	statePlayerTargetedByRefIDs := playerTargetedByRefSlicePool.Get().([]PlayerTargetedByRefID)[:0]
	for playerTargetedByRefID := range engine.State.PlayerTargetedByRef {
		statePlayerTargetedByRefIDs = append(statePlayerTargetedByRefIDs, playerTargetedByRefID)
	}
	patchPlayerTargetedByRefIDs := playerTargetedByRefSlicePool.Get().([]PlayerTargetedByRefID)[:0]
	for playerTargetedByRefID := range engine.Patch.PlayerTargetedByRef {
		patchPlayerTargetedByRefIDs = append(patchPlayerTargetedByRefIDs, playerTargetedByRefID)
	}
	dedupedIDs := deduplicatePlayerTargetedByRefIDs(statePlayerTargetedByRefIDs, patchPlayerTargetedByRefIDs)

	playerTargetedByRefSlicePool.Put(statePlayerTargetedByRefIDs)
	playerTargetedByRefSlicePool.Put(patchPlayerTargetedByRefIDs)

	return dedupedIDs
}

func (engine Engine) allPlayerTargetRefIDs() []PlayerTargetRefID {
	statePlayerTargetRefIDs := playerTargetRefSlicePool.Get().([]PlayerTargetRefID)[:0]
	for playerTargetRefID := range engine.State.PlayerTargetRef {
		statePlayerTargetRefIDs = append(statePlayerTargetRefIDs, playerTargetRefID)
	}
	patchPlayerTargetRefIDs := playerTargetRefSlicePool.Get().([]PlayerTargetRefID)[:0]
	for playerTargetRefID := range engine.Patch.PlayerTargetRef {
		patchPlayerTargetRefIDs = append(patchPlayerTargetRefIDs, playerTargetRefID)
	}
	dedupedIDs := deduplicatePlayerTargetRefIDs(statePlayerTargetRefIDs, patchPlayerTargetRefIDs)

	playerTargetRefSlicePool.Put(statePlayerTargetRefIDs)
	playerTargetRefSlicePool.Put(patchPlayerTargetRefIDs)

	return dedupedIDs
}

func (engine Engine) allItemBoundToRefIDs() []ItemBoundToRefID {
	stateItemBoundToRefIDs := itemBoundToRefSlicePool.Get().([]ItemBoundToRefID)[:0]
	for itemBoundToRefID := range engine.State.ItemBoundToRef {
		stateItemBoundToRefIDs = append(stateItemBoundToRefIDs, itemBoundToRefID)
	}
	patchItemBoundToRefIDs := itemBoundToRefSlicePool.Get().([]ItemBoundToRefID)[:0]
	for itemBoundToRefID := range engine.Patch.ItemBoundToRef {
		patchItemBoundToRefIDs = append(patchItemBoundToRefIDs, itemBoundToRefID)
	}
	dedupedIDs := deduplicateItemBoundToRefIDs(stateItemBoundToRefIDs, patchItemBoundToRefIDs)

	itemBoundToRefSlicePool.Put(stateItemBoundToRefIDs)
	itemBoundToRefSlicePool.Put(patchItemBoundToRefIDs)

	return dedupedIDs
}

func (engine Engine) allPlayerGuildMemberRefIDs() []PlayerGuildMemberRefID {
	statePlayerGuildMemberRefIDs := playerGuildMemberRefSlicePool.Get().([]PlayerGuildMemberRefID)[:0]
	for playerGuildMemberRefID := range engine.State.PlayerGuildMemberRef {
		statePlayerGuildMemberRefIDs = append(statePlayerGuildMemberRefIDs, playerGuildMemberRefID)
	}
	patchPlayerGuildMemberRefIDs := playerGuildMemberRefSlicePool.Get().([]PlayerGuildMemberRefID)[:0]
	for playerGuildMemberRefID := range engine.Patch.PlayerGuildMemberRef {
		patchPlayerGuildMemberRefIDs = append(patchPlayerGuildMemberRefIDs, playerGuildMemberRefID)
	}
	dedupedIDs := deduplicatePlayerGuildMemberRefIDs(statePlayerGuildMemberRefIDs, patchPlayerGuildMemberRefIDs)

	playerGuildMemberRefSlicePool.Put(statePlayerGuildMemberRefIDs)
	playerGuildMemberRefSlicePool.Put(patchPlayerGuildMemberRefIDs)

	return dedupedIDs
}

func (engine Engine) allPlayerEquipmentSetRefIDs() []PlayerEquipmentSetRefID {
	statePlayerEquipmentSetRefIDs := playerEquipmentSetRefSlicePool.Get().([]PlayerEquipmentSetRefID)[:0]
	for playerEquipmentSetRefID := range engine.State.PlayerEquipmentSetRef {
		statePlayerEquipmentSetRefIDs = append(statePlayerEquipmentSetRefIDs, playerEquipmentSetRefID)
	}
	patchPlayerEquipmentSetRefIDs := playerEquipmentSetRefSlicePool.Get().([]PlayerEquipmentSetRefID)[:0]
	for playerEquipmentSetRefID := range engine.Patch.PlayerEquipmentSetRef {
		patchPlayerEquipmentSetRefIDs = append(patchPlayerEquipmentSetRefIDs, playerEquipmentSetRefID)
	}
	dedupedIDs := deduplicatePlayerEquipmentSetRefIDs(statePlayerEquipmentSetRefIDs, patchPlayerEquipmentSetRefIDs)

	playerEquipmentSetRefSlicePool.Put(statePlayerEquipmentSetRefIDs)
	playerEquipmentSetRefSlicePool.Put(patchPlayerEquipmentSetRefIDs)

	return dedupedIDs
}

func (engine Engine) allEquipmentSetEquipmentRefIDs() []EquipmentSetEquipmentRefID {
	stateEquipmentSetEquipmentRefIDs := equipmentSetEquipmentRefSlicePool.Get().([]EquipmentSetEquipmentRefID)[:0]
	for equipmentSetEquipmentRefID := range engine.State.EquipmentSetEquipmentRef {
		stateEquipmentSetEquipmentRefIDs = append(stateEquipmentSetEquipmentRefIDs, equipmentSetEquipmentRefID)
	}
	patchEquipmentSetEquipmentRefIDs := equipmentSetEquipmentRefSlicePool.Get().([]EquipmentSetEquipmentRefID)[:0]
	for equipmentSetEquipmentRefID := range engine.Patch.EquipmentSetEquipmentRef {
		patchEquipmentSetEquipmentRefIDs = append(patchEquipmentSetEquipmentRefIDs, equipmentSetEquipmentRefID)
	}
	dedupedIDs := deduplicateEquipmentSetEquipmentRefIDs(stateEquipmentSetEquipmentRefIDs, patchEquipmentSetEquipmentRefIDs)

	equipmentSetEquipmentRefSlicePool.Put(stateEquipmentSetEquipmentRefIDs)
	equipmentSetEquipmentRefSlicePool.Put(patchEquipmentSetEquipmentRefIDs)

	return dedupedIDs
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
