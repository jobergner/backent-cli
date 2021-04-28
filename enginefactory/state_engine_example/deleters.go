package state

func (engine *Engine) DeletePlayer(playerID PlayerID) {
	player := engine.Player(playerID).player
	if player.HasParent_ {
		return
	}
	engine.deletePlayer(playerID)
}
func (engine *Engine) deletePlayer(playerID PlayerID) {
	player := engine.Player(playerID).player
	player.OperationKind_ = OperationKindDelete
	engine.Patch.Player[player.ID] = player
	engine.dereferencePlayerGuildMemberRefs(playerID)
	engine.dereferenceItemBoundToRefs(playerID)
	engine.deleteGearScore(player.GearScore)
	for _, guildMember := range player.GuildMembers {
		engine.deletePlayerGuildMemberRef(guildMember)
	}
	for _, itemID := range player.Items {
		engine.deleteItem(itemID)
	}
	engine.deletePosition(player.Position)
}

func (engine *Engine) DeleteGearScore(gearScoreID GearScoreID) {
	gearScore := engine.GearScore(gearScoreID).gearScore
	if gearScore.HasParent_ {
		return
	}
	engine.deleteGearScore(gearScoreID)
}
func (engine *Engine) deleteGearScore(gearScoreID GearScoreID) {
	gearScore := engine.GearScore(gearScoreID).gearScore
	gearScore.OperationKind_ = OperationKindDelete
	engine.Patch.GearScore[gearScore.ID] = gearScore
}

func (engine *Engine) DeletePosition(positionID PositionID) {
	position := engine.Position(positionID).position
	if position.HasParent_ {
		return
	}
	engine.deletePosition(positionID)
}
func (engine *Engine) deletePosition(positionID PositionID) {
	position := engine.Position(positionID).position
	position.OperationKind_ = OperationKindDelete
	engine.Patch.Position[position.ID] = position
}

func (engine *Engine) DeleteItem(itemID ItemID) {
	item := engine.Item(itemID).item
	if item.HasParent_ {
		return
	}
	engine.deleteItem(itemID)
}
func (engine *Engine) deleteItem(itemID ItemID) {
	item := engine.Item(itemID).item
	item.OperationKind_ = OperationKindDelete
	engine.Patch.Item[item.ID] = item
	engine.deleteItemBoundToRef(item.BoundTo)
	engine.deleteGearScore(item.GearScore)
}

func (engine *Engine) DeleteZoneItem(zoneItemID ZoneItemID) {
	zoneItem := engine.ZoneItem(zoneItemID).zoneItem
	if zoneItem.HasParent_ {
		return
	}
	engine.deleteZoneItem(zoneItemID)
}
func (engine *Engine) deleteZoneItem(zoneItemID ZoneItemID) {
	zoneItem := engine.ZoneItem(zoneItemID).zoneItem
	zoneItem.OperationKind_ = OperationKindDelete
	engine.Patch.ZoneItem[zoneItem.ID] = zoneItem
	engine.deleteItem(zoneItem.Item)
	engine.deletePosition(zoneItem.Position)
}

func (engine *Engine) DeleteZone(zoneID ZoneID) {
	engine.deleteZone(zoneID)
}
func (engine *Engine) deleteZone(zoneID ZoneID) {
	zone := engine.Zone(zoneID).zone
	zone.OperationKind_ = OperationKindDelete
	engine.Patch.Zone[zone.ID] = zone
	for _, zoneItemID := range zone.Items {
		engine.deleteZoneItem(zoneItemID)
	}
	for _, playerID := range zone.Players {
		engine.deletePlayer(playerID)
	}
}

func (engine *Engine) DeleteEquipmentSet(equipmentSetID EquipmentSetID) {
	engine.deleteEquipmentSet(equipmentSetID)
}
func (engine *Engine) deleteEquipmentSet(equipmentSetID EquipmentSetID) {
	equipmentSet := engine.EquipmentSet(equipmentSetID).equipmentSet
	equipmentSet.OperationKind_ = OperationKindDelete
	engine.Patch.EquipmentSet[equipmentSet.ID] = equipmentSet
	engine.dereferencePlayerEquipmentSetRefs(equipmentSetID)
	for _, equipmentSet := range equipmentSet.Equipment {
		engine.deleteEquipmentSetEquipmentRef(equipmentSet)
	}
}

func (engine *Engine) deletePlayerGuildMemberRef(playerGuildMemberRefID PlayerGuildMemberRefID) {
	playerGuildMemberRef := engine.playerGuildMemberRef(playerGuildMemberRefID).playerGuildMemberRef
	playerGuildMemberRef.OperationKind_ = OperationKindDelete
	engine.Patch.PlayerGuildMemberRef[playerGuildMemberRef.ID] = playerGuildMemberRef
}

func (engine *Engine) deletePlayerEquipmentSetRef(playerEquipmentSetRefID PlayerEquipmentSetRefID) {
	playerEquipmentSetRef := engine.playerEquipmentSetRef(playerEquipmentSetRefID).playerEquipmentSetRef
	playerEquipmentSetRef.OperationKind_ = OperationKindDelete
	engine.Patch.PlayerEquipmentSetRef[playerEquipmentSetRef.ID] = playerEquipmentSetRef
}

func (engine *Engine) deleteItemBoundToRef(itemBoundToRefID ItemBoundToRefID) {
	itemBoundToRef := engine.itemBoundToRef(itemBoundToRefID).itemBoundToRef
	itemBoundToRef.OperationKind_ = OperationKindDelete
	engine.Patch.ItemBoundToRef[itemBoundToRef.ID] = itemBoundToRef
}

func (engine *Engine) deleteEquipmentSetEquipmentRef(equipmentSetEquipmentRefID EquipmentSetEquipmentRefID) {
	equipmentSetEquipmentRef := engine.equipmentSetEquipmentRef(equipmentSetEquipmentRefID).equipmentSetEquipmentRef
	equipmentSetEquipmentRef.OperationKind_ = OperationKindDelete
	engine.Patch.EquipmentSetEquipmentRef[equipmentSetEquipmentRef.ID] = equipmentSetEquipmentRef
}
