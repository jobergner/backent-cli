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
	if player.OperationKind == OperationKindDelete {
		return
	}
	engine.dereferenceAttackEventTargetRefs(playerID)
	engine.dereferenceItemBoundToRefs(playerID)
	engine.dereferencePlayerGuildMemberRefs(playerID)
	engine.dereferencePlayerTargetRefsPlayer(playerID)
	engine.dereferencePlayerTargetedByRefsPlayer(playerID)
	for _, actionID := range player.Action {
		engine.deleteAttackEvent(actionID)
	}
	for _, equipmentSetID := range player.EquipmentSets {
		engine.deletePlayerEquipmentSetRef(equipmentSetID)
	}
	engine.deleteGearScore(player.GearScore)
	for _, guildMemberID := range player.GuildMembers {
		engine.deletePlayerGuildMemberRef(guildMemberID)
	}
	for _, itemID := range player.Items {
		engine.deleteItem(itemID)
	}
	engine.deletePosition(player.Position)
	engine.deletePlayerTargetRef(player.Target)
	for _, targetedByID := range player.TargetedBy {
		engine.deletePlayerTargetedByRef(targetedByID)
	}
	if _, ok := engine.State.Player[playerID]; ok {
		player.OperationKind = OperationKindDelete
		player.Meta.sign(player.engine.broadcastingClientID)
		engine.Patch.Player[player.ID] = player
	} else {
		delete(engine.Patch.Player, playerID)
	}
}

func (engine *Engine) deleteBoolValue(boolValueID BoolValueID) {
	boolValue := engine.boolValue(boolValueID)
	if boolValue.OperationKind == OperationKindDelete {
		return
	}
	if _, ok := engine.State.BoolValue[boolValueID]; ok {
		boolValue.OperationKind = OperationKindDelete
		boolValue.Meta.sign(boolValue.engine.broadcastingClientID)
		engine.Patch.BoolValue[boolValue.ID] = boolValue
	} else {
		delete(engine.Patch.BoolValue, boolValueID)
	}
}

func (engine *Engine) deleteIntValue(intValueID IntValueID) {
	intValue := engine.intValue(intValueID)
	if intValue.OperationKind == OperationKindDelete {
		return
	}
	if _, ok := engine.State.IntValue[intValueID]; ok {
		intValue.OperationKind = OperationKindDelete
		intValue.Meta.sign(intValue.engine.broadcastingClientID)
		engine.Patch.IntValue[intValue.ID] = intValue
	} else {
		delete(engine.Patch.IntValue, intValueID)
	}
}

func (engine *Engine) deleteFloatValue(floatValueID FloatValueID) {
	floatValue := engine.floatValue(floatValueID)
	if floatValue.OperationKind == OperationKindDelete {
		return
	}
	if _, ok := engine.State.FloatValue[floatValueID]; ok {
		floatValue.OperationKind = OperationKindDelete
		floatValue.Meta.sign(floatValue.engine.broadcastingClientID)
		engine.Patch.FloatValue[floatValue.ID] = floatValue
	} else {
		delete(engine.Patch.FloatValue, floatValueID)
	}
}

func (engine *Engine) deleteStringValue(stringValueID StringValueID) {
	stringValue := engine.stringValue(stringValueID)
	if stringValue.OperationKind == OperationKindDelete {
		return
	}
	if _, ok := engine.State.StringValue[stringValueID]; ok {
		stringValue.OperationKind = OperationKindDelete
		stringValue.Meta.sign(stringValue.engine.broadcastingClientID)
		engine.Patch.StringValue[stringValue.ID] = stringValue
	} else {
		delete(engine.Patch.StringValue, stringValueID)
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
	if gearScore.OperationKind == OperationKindDelete {
		return
	}
	engine.deleteIntValue(gearScore.Level)
	engine.deleteIntValue(gearScore.Score)
	if _, ok := engine.State.GearScore[gearScoreID]; ok {
		gearScore.OperationKind = OperationKindDelete
		gearScore.Meta.sign(gearScore.engine.broadcastingClientID)
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
	if position.OperationKind == OperationKindDelete {
		return
	}
	engine.deleteFloatValue(position.X)
	engine.deleteFloatValue(position.Y)
	if _, ok := engine.State.Position[positionID]; ok {
		position.OperationKind = OperationKindDelete
		position.Meta.sign(position.engine.broadcastingClientID)
		engine.Patch.Position[position.ID] = position
	} else {
		delete(engine.Patch.Position, positionID)
	}
}

func (engine *Engine) DeleteAttackEvent(attackEventID AttackEventID) {
	attackEvent := engine.AttackEvent(attackEventID).attackEvent
	if attackEvent.HasParent {
		return
	}
	engine.deleteAttackEvent(attackEventID)
}
func (engine *Engine) deleteAttackEvent(attackEventID AttackEventID) {
	attackEvent := engine.AttackEvent(attackEventID).attackEvent
	if attackEvent.OperationKind == OperationKindDelete {
		return
	}
	engine.deleteAttackEventTargetRef(attackEvent.Target)
	if _, ok := engine.State.AttackEvent[attackEventID]; ok {
		attackEvent.OperationKind = OperationKindDelete
		attackEvent.Meta.sign(attackEvent.engine.broadcastingClientID)
		engine.Patch.AttackEvent[attackEvent.ID] = attackEvent
	} else {
		delete(engine.Patch.AttackEvent, attackEventID)
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
	if item.OperationKind == OperationKindDelete {
		return
	}
	engine.dereferenceEquipmentSetEquipmentRefs(itemID)
	engine.deleteItemBoundToRef(item.BoundTo)
	engine.deleteGearScore(item.GearScore)
	engine.deleteStringValue(item.Name)
	engine.deleteAnyOfPlayer_Position(item.Origin, true)
	if _, ok := engine.State.Item[itemID]; ok {
		item.OperationKind = OperationKindDelete
		item.Meta.sign(item.engine.broadcastingClientID)
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
	if zoneItem.OperationKind == OperationKindDelete {
		return
	}
	engine.dereferencePlayerTargetRefsZoneItem(zoneItemID)
	engine.dereferencePlayerTargetedByRefsZoneItem(zoneItemID)
	engine.deleteItem(zoneItem.Item)
	engine.deletePosition(zoneItem.Position)
	if _, ok := engine.State.ZoneItem[zoneItemID]; ok {
		zoneItem.OperationKind = OperationKindDelete
		zoneItem.Meta.sign(zoneItem.engine.broadcastingClientID)
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
	if zone.OperationKind == OperationKindDelete {
		return
	}
	for _, interactableID := range zone.Interactables {
		engine.deleteAnyOfItem_Player_ZoneItem(interactableID, true)
	}
	for _, itemID := range zone.Items {
		engine.deleteZoneItem(itemID)
	}
	for _, playerID := range zone.Players {
		engine.deletePlayer(playerID)
	}
	for _, tagID := range zone.Tags {
		engine.deleteStringValue(tagID)
	}
	if _, ok := engine.State.Zone[zoneID]; ok {
		zone.OperationKind = OperationKindDelete
		zone.Meta.sign(zone.engine.broadcastingClientID)
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
	if equipmentSet.OperationKind == OperationKindDelete {
		return
	}
	engine.dereferencePlayerEquipmentSetRefs(equipmentSetID)
	for _, equipmentID := range equipmentSet.Equipment {
		engine.deleteEquipmentSetEquipmentRef(equipmentID)
	}
	engine.deleteStringValue(equipmentSet.Name)
	if _, ok := engine.State.EquipmentSet[equipmentSetID]; ok {
		equipmentSet.OperationKind = OperationKindDelete
		equipmentSet.Meta.sign(equipmentSet.engine.broadcastingClientID)
		engine.Patch.EquipmentSet[equipmentSet.ID] = equipmentSet
	} else {
		delete(engine.Patch.EquipmentSet, equipmentSetID)
	}
}

func (engine *Engine) deletePlayerGuildMemberRef(playerGuildMemberRefID PlayerGuildMemberRefID) {
	playerGuildMemberRef := engine.playerGuildMemberRef(playerGuildMemberRefID).playerGuildMemberRef
	if playerGuildMemberRef.OperationKind == OperationKindDelete {
		return
	}
	if _, ok := engine.State.PlayerGuildMemberRef[playerGuildMemberRefID]; ok {
		playerGuildMemberRef.OperationKind = OperationKindDelete
		playerGuildMemberRef.Meta.sign(playerGuildMemberRef.engine.broadcastingClientID)
		engine.Patch.PlayerGuildMemberRef[playerGuildMemberRef.ID] = playerGuildMemberRef
	} else {
		delete(engine.Patch.PlayerGuildMemberRef, playerGuildMemberRefID)
	}
}

func (engine *Engine) deletePlayerEquipmentSetRef(playerEquipmentSetRefID PlayerEquipmentSetRefID) {
	playerEquipmentSetRef := engine.playerEquipmentSetRef(playerEquipmentSetRefID).playerEquipmentSetRef
	if playerEquipmentSetRef.OperationKind == OperationKindDelete {
		return
	}
	if _, ok := engine.State.PlayerEquipmentSetRef[playerEquipmentSetRefID]; ok {
		playerEquipmentSetRef.OperationKind = OperationKindDelete
		playerEquipmentSetRef.Meta.sign(playerEquipmentSetRef.engine.broadcastingClientID)
		engine.Patch.PlayerEquipmentSetRef[playerEquipmentSetRef.ID] = playerEquipmentSetRef
	} else {
		delete(engine.Patch.PlayerEquipmentSetRef, playerEquipmentSetRefID)
	}
}

func (engine *Engine) deleteItemBoundToRef(itemBoundToRefID ItemBoundToRefID) {
	itemBoundToRef := engine.itemBoundToRef(itemBoundToRefID).itemBoundToRef
	if itemBoundToRef.OperationKind == OperationKindDelete {
		return
	}
	if _, ok := engine.State.ItemBoundToRef[itemBoundToRefID]; ok {
		itemBoundToRef.OperationKind = OperationKindDelete
		itemBoundToRef.Meta.sign(itemBoundToRef.engine.broadcastingClientID)
		engine.Patch.ItemBoundToRef[itemBoundToRef.ID] = itemBoundToRef
	} else {
		delete(engine.Patch.ItemBoundToRef, itemBoundToRefID)
	}
}

func (engine *Engine) deleteAttackEventTargetRef(attackEventTargetRefID AttackEventTargetRefID) {
	attackEventTargetRef := engine.attackEventTargetRef(attackEventTargetRefID).attackEventTargetRef
	if attackEventTargetRef.OperationKind == OperationKindDelete {
		return
	}
	if _, ok := engine.State.AttackEventTargetRef[attackEventTargetRefID]; ok {
		attackEventTargetRef.OperationKind = OperationKindDelete
		attackEventTargetRef.Meta.sign(attackEventTargetRef.engine.broadcastingClientID)
		engine.Patch.AttackEventTargetRef[attackEventTargetRef.ID] = attackEventTargetRef
	} else {
		delete(engine.Patch.AttackEventTargetRef, attackEventTargetRefID)
	}
}

func (engine *Engine) deleteEquipmentSetEquipmentRef(equipmentSetEquipmentRefID EquipmentSetEquipmentRefID) {
	equipmentSetEquipmentRef := engine.equipmentSetEquipmentRef(equipmentSetEquipmentRefID).equipmentSetEquipmentRef
	if equipmentSetEquipmentRef.OperationKind == OperationKindDelete {
		return
	}
	if _, ok := engine.State.EquipmentSetEquipmentRef[equipmentSetEquipmentRefID]; ok {
		equipmentSetEquipmentRef.OperationKind = OperationKindDelete
		equipmentSetEquipmentRef.Meta.sign(equipmentSetEquipmentRef.engine.broadcastingClientID)
		engine.Patch.EquipmentSetEquipmentRef[equipmentSetEquipmentRef.ID] = equipmentSetEquipmentRef
	} else {
		delete(engine.Patch.EquipmentSetEquipmentRef, equipmentSetEquipmentRefID)
	}
}

func (engine *Engine) deletePlayerTargetRef(playerTargetRefID PlayerTargetRefID) {
	playerTargetRef := engine.playerTargetRef(playerTargetRefID).playerTargetRef
	if playerTargetRef.OperationKind == OperationKindDelete {
		return
	}
	engine.deleteAnyOfPlayer_ZoneItem(playerTargetRef.ReferencedElementID, false)
	if _, ok := engine.State.PlayerTargetRef[playerTargetRefID]; ok {
		playerTargetRef.OperationKind = OperationKindDelete
		playerTargetRef.Meta.sign(playerTargetRef.engine.broadcastingClientID)
		engine.Patch.PlayerTargetRef[playerTargetRef.ID] = playerTargetRef
	} else {
		delete(engine.Patch.PlayerTargetRef, playerTargetRefID)
	}
}

func (engine *Engine) deletePlayerTargetedByRef(playerTargetedByRefID PlayerTargetedByRefID) {
	playerTargetedByRef := engine.playerTargetedByRef(playerTargetedByRefID).playerTargetedByRef
	if playerTargetedByRef.OperationKind == OperationKindDelete {
		return
	}
	engine.deleteAnyOfPlayer_ZoneItem(playerTargetedByRef.ReferencedElementID, false)
	if _, ok := engine.State.PlayerTargetedByRef[playerTargetedByRefID]; ok {
		playerTargetedByRef.OperationKind = OperationKindDelete
		playerTargetedByRef.Meta.sign(playerTargetedByRef.engine.broadcastingClientID)
		engine.Patch.PlayerTargetedByRef[playerTargetedByRef.ID] = playerTargetedByRef
	} else {
		delete(engine.Patch.PlayerTargetedByRef, playerTargetedByRefID)
	}
}

func (engine *Engine) deleteAnyOfPlayer_ZoneItem(anyOfPlayer_ZoneItemID AnyOfPlayer_ZoneItemID, deleteChild bool) {
	anyOfPlayer_ZoneItem := engine.anyOfPlayer_ZoneItem(anyOfPlayer_ZoneItemID).anyOfPlayer_ZoneItem
	if anyOfPlayer_ZoneItem.OperationKind == OperationKindDelete {
		return
	}
	if deleteChild {
		anyOfPlayer_ZoneItem.deleteChild()
	}
	if _, ok := engine.State.AnyOfPlayer_ZoneItem[anyOfPlayer_ZoneItemID]; ok {
		anyOfPlayer_ZoneItem.OperationKind = OperationKindDelete
		anyOfPlayer_ZoneItem.Meta.sign(anyOfPlayer_ZoneItem.engine.broadcastingClientID)
		engine.Patch.AnyOfPlayer_ZoneItem[anyOfPlayer_ZoneItem.ID] = anyOfPlayer_ZoneItem
	} else {
		delete(engine.Patch.AnyOfPlayer_ZoneItem, anyOfPlayer_ZoneItemID)
	}
}

func (engine *Engine) deleteAnyOfPlayer_Position(anyOfPlayer_PositionID AnyOfPlayer_PositionID, deleteChild bool) {
	anyOfPlayer_Position := engine.anyOfPlayer_Position(anyOfPlayer_PositionID).anyOfPlayer_Position
	if anyOfPlayer_Position.OperationKind == OperationKindDelete {
		return
	}
	if deleteChild {
		anyOfPlayer_Position.deleteChild()
	}
	if _, ok := engine.State.AnyOfPlayer_Position[anyOfPlayer_PositionID]; ok {
		anyOfPlayer_Position.OperationKind = OperationKindDelete
		anyOfPlayer_Position.Meta.sign(anyOfPlayer_Position.engine.broadcastingClientID)
		engine.Patch.AnyOfPlayer_Position[anyOfPlayer_Position.ID] = anyOfPlayer_Position
	} else {
		delete(engine.Patch.AnyOfPlayer_Position, anyOfPlayer_PositionID)
	}
}

func (engine *Engine) deleteAnyOfItem_Player_ZoneItem(anyOfItem_Player_ZoneItemID AnyOfItem_Player_ZoneItemID, deleteChild bool) {
	anyOfItem_Player_ZoneItem := engine.anyOfItem_Player_ZoneItem(anyOfItem_Player_ZoneItemID).anyOfItem_Player_ZoneItem
	if anyOfItem_Player_ZoneItem.OperationKind == OperationKindDelete {
		return
	}
	if deleteChild {
		anyOfItem_Player_ZoneItem.deleteChild()
	}
	if _, ok := engine.State.AnyOfItem_Player_ZoneItem[anyOfItem_Player_ZoneItemID]; ok {
		anyOfItem_Player_ZoneItem.OperationKind = OperationKindDelete
		anyOfItem_Player_ZoneItem.Meta.sign(anyOfItem_Player_ZoneItem.engine.broadcastingClientID)
		engine.Patch.AnyOfItem_Player_ZoneItem[anyOfItem_Player_ZoneItem.ID] = anyOfItem_Player_ZoneItem
	} else {
		delete(engine.Patch.AnyOfItem_Player_ZoneItem, anyOfItem_Player_ZoneItemID)
	}
}
