package state

func (engine *Engine) CreateGearScore() gearScore {
	return engine.createGearScore(false)
}

func (engine *Engine) createGearScore(hasParent bool) gearScore {
	var element gearScoreCore
	element.engine = engine
	element.ID = GearScoreID(engine.GenerateID())
	element.HasParent = hasParent
	element.OperationKind = OperationKindUpdate
	engine.Patch.GearScore[element.ID] = element
	return gearScore{gearScore: element}
}

func (engine *Engine) CreatePosition() position {
	return engine.createPosition(false)
}

func (engine *Engine) createPosition(hasParent bool) position {
	var element positionCore
	element.engine = engine
	element.ID = PositionID(engine.GenerateID())
	element.HasParent = hasParent
	element.OperationKind = OperationKindUpdate
	engine.Patch.Position[element.ID] = element
	return position{position: element}
}

func (engine *Engine) CreateItem() item {
	return engine.createItem(false)
}

func (engine *Engine) createItem(hasParent bool) item {
	var element itemCore
	element.engine = engine
	element.ID = ItemID(engine.GenerateID())
	element.HasParent = hasParent
	elementGearScore := engine.createGearScore(true)
	element.GearScore = elementGearScore.gearScore.ID
	elementOrigin := engine.createAnyOfPlayer_Position(true)
	element.Origin = elementOrigin.anyOfPlayer_Position.ID
	element.OperationKind = OperationKindUpdate
	engine.Patch.Item[element.ID] = element
	return item{item: element}
}

func (engine *Engine) CreateZoneItem() zoneItem {
	return engine.createZoneItem(false)
}

func (engine *Engine) createZoneItem(hasParent bool) zoneItem {
	var element zoneItemCore
	element.engine = engine
	element.ID = ZoneItemID(engine.GenerateID())
	element.HasParent = hasParent
	elementItem := engine.createItem(true)
	element.Item = elementItem.item.ID
	elementPosition := engine.createPosition(true)
	element.Position = elementPosition.position.ID
	element.OperationKind = OperationKindUpdate
	engine.Patch.ZoneItem[element.ID] = element
	return zoneItem{zoneItem: element}
}

func (engine *Engine) CreatePlayer() player {
	return engine.createPlayer(false)
}

func (engine *Engine) createPlayer(hasParent bool) player {
	var element playerCore
	element.engine = engine
	element.ID = PlayerID(engine.GenerateID())
	element.HasParent = hasParent
	elementGearScore := engine.createGearScore(true)
	element.GearScore = elementGearScore.gearScore.ID
	elementPosition := engine.createPosition(true)
	element.Position = elementPosition.position.ID
	element.OperationKind = OperationKindUpdate
	engine.Patch.Player[element.ID] = element
	return player{player: element}
}

func (engine *Engine) CreateZone() zone {
	return engine.createZone()
}

func (engine *Engine) createZone() zone {
	var element zoneCore
	element.engine = engine
	element.ID = ZoneID(engine.GenerateID())
	element.OperationKind = OperationKindUpdate
	engine.Patch.Zone[element.ID] = element
	return zone{zone: element}
}

func (engine *Engine) createItemBoundToRef(referencedElementID PlayerID, parentID ItemID) itemBoundToRefCore {
	var element itemBoundToRefCore
	element.engine = engine
	element.ReferencedElementID = referencedElementID
	element.ParentID = parentID
	element.ID = ItemBoundToRefID(engine.GenerateID())
	element.OperationKind = OperationKindUpdate
	engine.Patch.ItemBoundToRef[element.ID] = element
	return element
}

func (engine *Engine) createPlayerGuildMemberRef(referencedElementID PlayerID, parentID PlayerID) playerGuildMemberRefCore {
	var element playerGuildMemberRefCore
	element.engine = engine
	element.ReferencedElementID = referencedElementID
	element.ParentID = parentID
	element.ID = PlayerGuildMemberRefID(engine.GenerateID())
	element.OperationKind = OperationKindUpdate
	engine.Patch.PlayerGuildMemberRef[element.ID] = element
	return element
}

func (engine *Engine) CreateEquipmentSet() equipmentSet {
	return engine.createEquipmentSet()
}

func (engine *Engine) createEquipmentSet() equipmentSet {
	var element equipmentSetCore
	element.engine = engine
	element.ID = EquipmentSetID(engine.GenerateID())
	element.OperationKind = OperationKindUpdate
	engine.Patch.EquipmentSet[element.ID] = element
	return equipmentSet{equipmentSet: element}
}

func (engine *Engine) createEquipmentSetEquipmentRef(referencedElementID ItemID, parentID EquipmentSetID) equipmentSetEquipmentRefCore {
	var element equipmentSetEquipmentRefCore
	element.engine = engine
	element.ReferencedElementID = referencedElementID
	element.ParentID = parentID
	element.ID = EquipmentSetEquipmentRefID(engine.GenerateID())
	element.OperationKind = OperationKindUpdate
	engine.Patch.EquipmentSetEquipmentRef[element.ID] = element
	return element
}

func (engine *Engine) createPlayerEquipmentSetRef(referencedElementID EquipmentSetID, parentID PlayerID) playerEquipmentSetRefCore {
	var element playerEquipmentSetRefCore
	element.engine = engine
	element.ReferencedElementID = referencedElementID
	element.ParentID = parentID
	element.ID = PlayerEquipmentSetRefID(engine.GenerateID())
	element.OperationKind = OperationKindUpdate
	engine.Patch.PlayerEquipmentSetRef[element.ID] = element
	return element
}

func (engine *Engine) createPlayerTargetRef(referencedElementID AnyOfPlayer_ZoneItemID, parentID PlayerID) playerTargetRefCore {
	var element playerTargetRefCore
	element.engine = engine
	element.ReferencedElementID = referencedElementID
	element.ParentID = parentID
	element.ID = PlayerTargetRefID(engine.GenerateID())
	element.OperationKind = OperationKindUpdate
	engine.Patch.PlayerTargetRef[element.ID] = element
	return element
}

func (engine *Engine) createPlayerTargetedByRef(referencedElementID AnyOfPlayer_ZoneItemID, parentID PlayerID) playerTargetedByRefCore {
	var element playerTargetedByRefCore
	element.engine = engine
	element.ReferencedElementID = referencedElementID
	element.ParentID = parentID
	element.ID = PlayerTargetedByRefID(engine.GenerateID())
	element.OperationKind = OperationKindUpdate
	engine.Patch.PlayerTargetedByRef[element.ID] = element
	return element
}

func (engine *Engine) createAnyOfPlayer_ZoneItem(setDefaultValue bool) anyOfPlayer_ZoneItem {
	var element anyOfPlayer_ZoneItemCore
	element.engine = engine
	element.ID = AnyOfPlayer_ZoneItemID(engine.GenerateID())
	if setDefaultValue {
		elementPlayer := engine.createPlayer(true)
		element.Player = elementPlayer.player.ID
		element.ElementKind = ElementKindPlayer
	}
	element.OperationKind = OperationKindUpdate
	engine.Patch.AnyOfPlayer_ZoneItem[element.ID] = element
	return anyOfPlayer_ZoneItem{anyOfPlayer_ZoneItem: element}
}

func (engine *Engine) createAnyOfPlayer_Position(setDefaultValue bool) anyOfPlayer_Position {
	var element anyOfPlayer_PositionCore
	element.engine = engine
	element.ID = AnyOfPlayer_PositionID(engine.GenerateID())
	if setDefaultValue {
		elementPlayer := engine.createPlayer(true)
		element.Player = elementPlayer.player.ID
		element.ElementKind = ElementKindPlayer
	}
	element.OperationKind = OperationKindUpdate
	engine.Patch.AnyOfPlayer_Position[element.ID] = element
	return anyOfPlayer_Position{anyOfPlayer_Position: element}
}

func (engine *Engine) createAnyOfItem_Player_ZoneItem(setDefaultValue bool) anyOfItem_Player_ZoneItem {
	var element anyOfItem_Player_ZoneItemCore
	element.engine = engine
	element.ID = AnyOfItem_Player_ZoneItemID(engine.GenerateID())
	if setDefaultValue {
		elementItem := engine.createItem(true)
		element.Item = elementItem.item.ID
		element.ElementKind = ElementKindItem
	}
	element.OperationKind = OperationKindUpdate
	engine.Patch.AnyOfItem_Player_ZoneItem[element.ID] = element
	return anyOfItem_Player_ZoneItem{anyOfItem_Player_ZoneItem: element}
}