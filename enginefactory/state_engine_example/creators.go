package state

func (engine *Engine) CreateGearScore() gearScore {
	return engine.createGearScore(false)
}

func (engine *Engine) createGearScore(hasParent bool) gearScore {
	var element gearScoreCore
	element.engine = engine
	element.ID = GearScoreID(engine.GenerateID())
	element.HasParent_ = hasParent
	element.OperationKind_ = OperationKindUpdate
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
	element.HasParent_ = hasParent
	element.OperationKind_ = OperationKindUpdate
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
	element.HasParent_ = hasParent
	elementGearScore := engine.createGearScore(true)
	element.GearScore = elementGearScore.gearScore.ID
	element.OperationKind_ = OperationKindUpdate
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
	element.HasParent_ = hasParent
	elementItem := engine.createItem(true)
	element.Item = elementItem.item.ID
	elementPosition := engine.createPosition(true)
	element.Position = elementPosition.position.ID
	element.OperationKind_ = OperationKindUpdate
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
	element.HasParent_ = hasParent
	elementGearScore := engine.createGearScore(true)
	element.GearScore = elementGearScore.gearScore.ID
	elementPosition := engine.createPosition(true)
	element.Position = elementPosition.position.ID
	element.OperationKind_ = OperationKindUpdate
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
	element.OperationKind_ = OperationKindUpdate
	engine.Patch.Zone[element.ID] = element
	return zone{zone: element}
}

func (engine *Engine) createItemBoundToRef(referencedElementID PlayerID, parentID ItemID) itemBoundToRefCore {
	var element itemBoundToRefCore
	element.engine = engine
	element.ReferencedElementID = referencedElementID
	element.ParentID = parentID
	element.ID = ItemBoundToRefID(engine.GenerateID())
	element.OperationKind_ = OperationKindUpdate
	engine.Patch.ItemBoundToRef[element.ID] = element
	return element
}

func (engine *Engine) createPlayerGuildMemberRef(referencedElementID PlayerID, parentID PlayerID) playerGuildMemberRefCore {
	var element playerGuildMemberRefCore
	element.engine = engine
	element.ReferencedElementID = referencedElementID
	element.ParentID = parentID
	element.ID = PlayerGuildMemberRefID(engine.GenerateID())
	element.OperationKind_ = OperationKindUpdate
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
	element.OperationKind_ = OperationKindUpdate
	engine.Patch.EquipmentSet[element.ID] = element
	return equipmentSet{equipmentSet: element}
}

func (engine *Engine) createEquipmentSetEquipmentRef(referencedElementID ItemID, parentID EquipmentSetID) equipmentSetEquipmentRefCore {
	var element equipmentSetEquipmentRefCore
	element.engine = engine
	element.ReferencedElementID = referencedElementID
	element.ParentID = parentID
	element.ID = EquipmentSetEquipmentRefID(engine.GenerateID())
	element.OperationKind_ = OperationKindUpdate
	engine.Patch.EquipmentSetEquipmentRef[element.ID] = element
	return element
}

func (engine *Engine) createPlayerEquipmentSetRef(referencedElementID EquipmentSetID, parentID PlayerID) playerEquipmentSetRefCore {
	var element playerEquipmentSetRefCore
	element.engine = engine
	element.ReferencedElementID = referencedElementID
	element.ParentID = parentID
	element.ID = PlayerEquipmentSetRefID(engine.GenerateID())
	element.OperationKind_ = OperationKindUpdate
	engine.Patch.PlayerEquipmentSetRef[element.ID] = element
	return element
}
