package state

func (engine *Engine) DeletePlayer(playerID PlayerID) {
	player := engine.Player(playerID).player
	if player.HasParent {
		return
	}
	engine.deletePlayer(playerID)
}
func (engine *Engine) deletePlayer(playerID PlayerID) {
	player := engine.Player(playerID).player
	engine.dereferenceItemBoundToRefs(playerID)
	engine.dereferencePlayerGuildMemberRefs(playerID)
	engine.dereferencePlayerTargetRefsPlayer(playerID)
	engine.dereferencePlayerTargetedByRefsPlayer(playerID)
	engine.deleteGearScore(player.GearScore)
	for _, guildMember := range player.GuildMembers {
		engine.deletePlayerGuildMemberRef(guildMember)
	}
	for _, itemID := range player.Items {
		engine.deleteItem(itemID)
	}
	engine.deletePosition(player.Position)
	engine.deletePlayerTargetRef(player.Target)
	for _, targetedBy := range player.TargetedBy {
		engine.deletePlayerTargetedByRef(targetedBy)
	}
	if _, ok := engine.State.Player[playerID]; ok {
		player.OperationKind = OperationKindDelete
		engine.Patch.Player[player.ID] = player
	} else {
		delete(engine.Patch.Player, playerID)
	}
}

func (engine *Engine) DeleteGearScore(gearScoreID GearScoreID) {
	gearScore := engine.GearScore(gearScoreID).gearScore
	if gearScore.HasParent {
		return
	}
	engine.deleteGearScore(gearScoreID)
}
func (engine *Engine) deleteGearScore(gearScoreID GearScoreID) {
	gearScore := engine.GearScore(gearScoreID).gearScore
	if _, ok := engine.State.GearScore[gearScoreID]; ok {
		gearScore.OperationKind = OperationKindDelete
		engine.Patch.GearScore[gearScore.ID] = gearScore
	} else {
		delete(engine.Patch.GearScore, gearScoreID)
	}
}

func (engine *Engine) DeletePosition(positionID PositionID) {
	position := engine.Position(positionID).position
	if position.HasParent {
		return
	}
	engine.deletePosition(positionID)
}
func (engine *Engine) deletePosition(positionID PositionID) {
	position := engine.Position(positionID).position
	if _, ok := engine.State.Position[positionID]; ok {
		position.OperationKind = OperationKindDelete
		engine.Patch.Position[position.ID] = position
	} else {
		delete(engine.Patch.Position, positionID)
	}
}

func (engine *Engine) DeleteItem(itemID ItemID) {
	item := engine.Item(itemID).item
	if item.HasParent {
		return
	}
	engine.deleteItem(itemID)
}
func (engine *Engine) deleteItem(itemID ItemID) {
	item := engine.Item(itemID).item
	engine.deleteItemBoundToRef(item.BoundTo)
	engine.deleteGearScore(item.GearScore)
	engine.deleteAnyOfPlayerPosition(item.Origin, true)
	if _, ok := engine.State.Item[itemID]; ok {
		item.OperationKind = OperationKindDelete
		engine.Patch.Item[item.ID] = item
	} else {
		delete(engine.Patch.Item, itemID)
	}
}

func (engine *Engine) DeleteZoneItem(zoneItemID ZoneItemID) {
	zoneItem := engine.ZoneItem(zoneItemID).zoneItem
	if zoneItem.HasParent {
		return
	}
	engine.deleteZoneItem(zoneItemID)
}
func (engine *Engine) deleteZoneItem(zoneItemID ZoneItemID) {
	zoneItem := engine.ZoneItem(zoneItemID).zoneItem
	engine.dereferencePlayerTargetRefsZoneItem(zoneItemID)
	engine.dereferencePlayerTargetedByRefsZoneItem(zoneItemID)
	engine.deleteItem(zoneItem.Item)
	engine.deletePosition(zoneItem.Position)
	if _, ok := engine.State.ZoneItem[zoneItemID]; ok {
		zoneItem.OperationKind = OperationKindDelete
		engine.Patch.ZoneItem[zoneItem.ID] = zoneItem
	} else {
		delete(engine.Patch.ZoneItem, zoneItemID)
	}
}

func (engine *Engine) DeleteZone(zoneID ZoneID) {
	engine.deleteZone(zoneID)
}
func (engine *Engine) deleteZone(zoneID ZoneID) {
	zone := engine.Zone(zoneID).zone
	for _, zoneItemID := range zone.Items {
		engine.deleteZoneItem(zoneItemID)
	}
	for _, playerID := range zone.Players {
		engine.deletePlayer(playerID)
	}
	if _, ok := engine.State.Zone[zoneID]; ok {
		zone.OperationKind = OperationKindDelete
		engine.Patch.Zone[zone.ID] = zone
	} else {
		delete(engine.Patch.Zone, zoneID)
	}
}

func (engine *Engine) DeleteEquipmentSet(equipmentSetID EquipmentSetID) {
	engine.deleteEquipmentSet(equipmentSetID)
}
func (engine *Engine) deleteEquipmentSet(equipmentSetID EquipmentSetID) {
	equipmentSet := engine.EquipmentSet(equipmentSetID).equipmentSet
	engine.dereferencePlayerEquipmentSetRefs(equipmentSetID)
	for _, equipmentSet := range equipmentSet.Equipment {
		engine.deleteEquipmentSetEquipmentRef(equipmentSet)
	}
	if _, ok := engine.State.EquipmentSet[equipmentSetID]; ok {
		equipmentSet.OperationKind = OperationKindDelete
		engine.Patch.EquipmentSet[equipmentSet.ID] = equipmentSet
	} else {
		delete(engine.Patch.EquipmentSet, equipmentSetID)
	}
}

func (engine *Engine) deletePlayerGuildMemberRef(playerGuildMemberRefID PlayerGuildMemberRefID) {
	playerGuildMemberRef := engine.playerGuildMemberRef(playerGuildMemberRefID).playerGuildMemberRef
	if _, ok := engine.State.PlayerGuildMemberRef[playerGuildMemberRefID]; ok {
		playerGuildMemberRef.OperationKind = OperationKindDelete
		engine.Patch.PlayerGuildMemberRef[playerGuildMemberRef.ID] = playerGuildMemberRef
	} else {
		delete(engine.Patch.PlayerGuildMemberRef, playerGuildMemberRefID)
	}
}

func (engine *Engine) deletePlayerEquipmentSetRef(playerEquipmentSetRefID PlayerEquipmentSetRefID) {
	playerEquipmentSetRef := engine.playerEquipmentSetRef(playerEquipmentSetRefID).playerEquipmentSetRef
	if _, ok := engine.State.PlayerEquipmentSetRef[playerEquipmentSetRefID]; ok {
		playerEquipmentSetRef.OperationKind = OperationKindDelete
		engine.Patch.PlayerEquipmentSetRef[playerEquipmentSetRef.ID] = playerEquipmentSetRef
	} else {
		delete(engine.Patch.PlayerEquipmentSetRef, playerEquipmentSetRefID)
	}
}

func (engine *Engine) deleteItemBoundToRef(itemBoundToRefID ItemBoundToRefID) {
	itemBoundToRef := engine.itemBoundToRef(itemBoundToRefID).itemBoundToRef
	if _, ok := engine.State.ItemBoundToRef[itemBoundToRefID]; ok {
		itemBoundToRef.OperationKind = OperationKindDelete
		engine.Patch.ItemBoundToRef[itemBoundToRef.ID] = itemBoundToRef
	} else {
		delete(engine.Patch.ItemBoundToRef, itemBoundToRefID)
	}
}

func (engine *Engine) deleteEquipmentSetEquipmentRef(equipmentSetEquipmentRefID EquipmentSetEquipmentRefID) {
	equipmentSetEquipmentRef := engine.equipmentSetEquipmentRef(equipmentSetEquipmentRefID).equipmentSetEquipmentRef
	if _, ok := engine.State.EquipmentSetEquipmentRef[equipmentSetEquipmentRefID]; ok {
		equipmentSetEquipmentRef.OperationKind = OperationKindDelete
		engine.Patch.EquipmentSetEquipmentRef[equipmentSetEquipmentRef.ID] = equipmentSetEquipmentRef
	} else {
		delete(engine.Patch.EquipmentSetEquipmentRef, equipmentSetEquipmentRefID)
	}
}

func (engine *Engine) deletePlayerTargetRef(playerTargetRefID PlayerTargetRefID) {
	playerTargetRef := engine.playerTargetRef(playerTargetRefID).playerTargetRef
	engine.deleteAnyOfPlayerZoneItem(playerTargetRef.ReferencedElementID, false)
	if _, ok := engine.State.PlayerTargetRef[playerTargetRefID]; ok {
		playerTargetRef.OperationKind = OperationKindDelete
		engine.Patch.PlayerTargetRef[playerTargetRef.ID] = playerTargetRef
	} else {
		delete(engine.Patch.PlayerTargetRef, playerTargetRefID)
	}
}

func (engine *Engine) deletePlayerTargetedByRef(playerTargetedByRefID PlayerTargetedByRefID) {
	playerTargetedByRef := engine.playerTargetedByRef(playerTargetedByRefID).playerTargetedByRef
	engine.deleteAnyOfPlayerZoneItem(playerTargetedByRef.ReferencedElementID, false)
	if _, ok := engine.State.PlayerTargetedByRef[playerTargetedByRefID]; ok {
		playerTargetedByRef.OperationKind = OperationKindDelete
		engine.Patch.PlayerTargetedByRef[playerTargetedByRef.ID] = playerTargetedByRef
	} else {
		delete(engine.Patch.PlayerTargetedByRef, playerTargetedByRefID)
	}
}

func (engine *Engine) deleteAnyOfPlayerZoneItem(anyOfPlayerZoneItemID AnyOfPlayerZoneItemID, deleteChild bool) {
	anyOfPlayerZoneItem := engine.anyOfPlayerZoneItem(anyOfPlayerZoneItemID).anyOfPlayerZoneItem
	if deleteChild {
		anyOfPlayerZoneItem.deleteChild()
	}
	if _, ok := engine.State.AnyOfPlayerZoneItem[anyOfPlayerZoneItemID]; ok {
		anyOfPlayerZoneItem.OperationKind = OperationKindDelete
		engine.Patch.AnyOfPlayerZoneItem[anyOfPlayerZoneItem.ID] = anyOfPlayerZoneItem
	} else {
		delete(engine.Patch.AnyOfPlayerZoneItem, anyOfPlayerZoneItemID)
	}
}

func (engine *Engine) deleteAnyOfPlayerPosition(anyOfPlayerZoneID AnyOfPlayerPositionID, deleteChild bool) {
	anyOfPlayerPosition := engine.anyOfPlayerPosition(anyOfPlayerZoneID).anyOfPlayerPosition
	if deleteChild {
		anyOfPlayerPosition.deleteChild()
	}
	if _, ok := engine.State.AnyOfPlayerPosition[anyOfPlayerZoneID]; ok {
		anyOfPlayerPosition.OperationKind = OperationKindDelete
		engine.Patch.AnyOfPlayerPosition[anyOfPlayerPosition.ID] = anyOfPlayerPosition
	} else {
		delete(engine.Patch.AnyOfPlayerPosition, anyOfPlayerZoneID)
	}
}

func (engine *Engine) deleteAnyOfItemPlayerZoneItem(anyOfItemPlayerZoneItemID AnyOfItemPlayerZoneItemID, deleteChild bool) {
	anyOfItemPlayerZoneItem := engine.anyOfItemPlayerZoneItem(anyOfItemPlayerZoneItemID).anyOfItemPlayerZoneItem
	if deleteChild {
		anyOfItemPlayerZoneItem.deleteChild()
	}
	if _, ok := engine.State.AnyOfItemPlayerZoneItem[anyOfItemPlayerZoneItemID]; ok {
		anyOfItemPlayerZoneItem.OperationKind = OperationKindDelete
		engine.Patch.AnyOfItemPlayerZoneItem[anyOfItemPlayerZoneItem.ID] = anyOfItemPlayerZoneItem
	} else {
		delete(engine.Patch.AnyOfItemPlayerZoneItem, anyOfItemPlayerZoneItemID)
	}
}
