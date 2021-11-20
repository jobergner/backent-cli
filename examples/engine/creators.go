package state

func (engine *Engine) CreateEquipmentSet() EquipmentSet {
	return engine.createEquipmentSet(newPath(equipmentSetIdentifier), true)
}

func (engine *Engine) createEquipmentSet(p path, extendWithID bool) EquipmentSet {
	var element equipmentSetCore
	element.engine = engine
	element.ID = EquipmentSetID(engine.GenerateID())
	element.path = p
	if extendWithID {
		element.path = element.path.id(int(element.ID))
	}
	element.Path = element.path.toJSONPath()
	element.OperationKind = OperationKindUpdate
	element.HasParent = len(p) > 1
	engine.Patch.EquipmentSet[element.ID] = element
	return EquipmentSet{equipmentSet: element}
}

func (engine *Engine) CreateGearScore() GearScore {
	return engine.createGearScore(newPath(gearScoreIdentifier), true)
}

func (engine *Engine) createGearScore(p path, extendWithID bool) GearScore {
	var element gearScoreCore
	element.engine = engine
	element.ID = GearScoreID(engine.GenerateID())
	element.path = p
	if extendWithID {
		element.path = element.path.id(int(element.ID))
	}
	element.Path = element.path.toJSONPath()
	element.OperationKind = OperationKindUpdate
	element.HasParent = len(p) > 1
	engine.Patch.GearScore[element.ID] = element
	return GearScore{gearScore: element}
}

func (engine *Engine) CreatePosition() Position {
	return engine.createPosition(newPath(positionIdentifier), true)
}

func (engine *Engine) createPosition(p path, extendWithID bool) Position {
	var element positionCore
	element.engine = engine
	element.ID = PositionID(engine.GenerateID())
	element.path = p
	if extendWithID {
		element.path = element.path.id(int(element.ID))
	}
	element.Path = element.path.toJSONPath()
	element.OperationKind = OperationKindUpdate
	element.HasParent = len(p) > 1
	engine.Patch.Position[element.ID] = element
	return Position{position: element}
}

func (engine *Engine) CreateItem() Item {
	return engine.createItem(newPath(itemIdentifier), true)
}

func (engine *Engine) createItem(p path, extendWithID bool) Item {
	var element itemCore
	element.engine = engine
	element.ID = ItemID(engine.GenerateID())
	element.path = p
	if extendWithID {
		element.path = element.path.id(int(element.ID))
	}
	element.Path = element.path.toJSONPath()
	elementGearScore := engine.createGearScore(element.path.gearScore(), false)
	element.GearScore = elementGearScore.gearScore.ID
	elementOrigin := engine.createAnyOfPlayer_Position(true, element.path.origin())
	element.Origin = elementOrigin.anyOfPlayer_Position.ID
	element.OperationKind = OperationKindUpdate
	element.HasParent = len(p) > 1
	engine.Patch.Item[element.ID] = element
	return Item{item: element}
}

func (engine *Engine) CreateZoneItem() ZoneItem {
	return engine.createZoneItem(newPath(zoneItemIdentifier), true)
}

func (engine *Engine) createZoneItem(p path, extendWithID bool) ZoneItem {
	var element zoneItemCore
	element.engine = engine
	element.ID = ZoneItemID(engine.GenerateID())
	element.path = p
	if extendWithID {
		element.path = element.path.id(int(element.ID))
	}
	element.Path = element.path.toJSONPath()
	elementItem := engine.createItem(element.path.item(), false)
	element.Item = elementItem.item.ID
	elementPosition := engine.createPosition(element.path.position(), false)
	element.Position = elementPosition.position.ID
	element.OperationKind = OperationKindUpdate
	element.HasParent = len(p) > 1
	engine.Patch.ZoneItem[element.ID] = element
	return ZoneItem{zoneItem: element}
}

func (engine *Engine) CreatePlayer() Player {
	return engine.createPlayer(newPath(playerIdentifier), true)
}

func (engine *Engine) createPlayer(p path, extendWithID bool) Player {
	var element playerCore
	element.engine = engine
	element.ID = PlayerID(engine.GenerateID())
	element.path = p
	if extendWithID {
		element.path = element.path.id(int(element.ID))
	}
	element.Path = element.path.toJSONPath()
	elementGearScore := engine.createGearScore(element.path.gearScore(), false)
	element.GearScore = elementGearScore.gearScore.ID
	elementPosition := engine.createPosition(element.path.position(), false)
	element.Position = elementPosition.position.ID
	element.OperationKind = OperationKindUpdate
	element.HasParent = len(p) > 1
	engine.Patch.Player[element.ID] = element
	return Player{player: element}
}

func (engine *Engine) CreateZone() Zone {
	return engine.createZone(newPath(zoneIdentifier), true)
}

func (engine *Engine) createZone(p path, extendWithID bool) Zone {
	var element zoneCore
	element.engine = engine
	element.ID = ZoneID(engine.GenerateID())
	element.path = p
	if extendWithID {
		element.path = element.path.id(int(element.ID))
	}
	element.Path = element.path.toJSONPath()
	element.OperationKind = OperationKindUpdate
	element.HasParent = len(p) > 1
	engine.Patch.Zone[element.ID] = element
	return Zone{zone: element}
}

func (engine *Engine) createItemBoundToRef(p path, extendWithID bool, referencedElementID PlayerID, parentID ItemID) itemBoundToRefCore {
	var element itemBoundToRefCore
	element.engine = engine
	element.ReferencedElementID = referencedElementID
	element.ParentID = parentID
	element.ID = ItemBoundToRefID(engine.GenerateID())
	element.path = p
	if extendWithID {
		element.path = element.path.id(int(element.ID))
	}
	element.OperationKind = OperationKindUpdate
	engine.Patch.ItemBoundToRef[element.ID] = element
	return element
}

func (engine *Engine) createPlayerGuildMemberRef(p path, extendWithID bool, referencedElementID PlayerID, parentID PlayerID) playerGuildMemberRefCore {
	var element playerGuildMemberRefCore
	element.engine = engine
	element.ReferencedElementID = referencedElementID
	element.ParentID = parentID
	element.ID = PlayerGuildMemberRefID(engine.GenerateID())
	element.path = p
	if extendWithID {
		element.path = element.path.id(int(element.ID))
	}
	element.OperationKind = OperationKindUpdate
	engine.Patch.PlayerGuildMemberRef[element.ID] = element
	return element
}

func (engine *Engine) createEquipmentSetEquipmentRef(p path, extendWithID bool, referencedElementID ItemID, parentID EquipmentSetID) equipmentSetEquipmentRefCore {
	var element equipmentSetEquipmentRefCore
	element.engine = engine
	element.ReferencedElementID = referencedElementID
	element.ParentID = parentID
	element.ID = EquipmentSetEquipmentRefID(engine.GenerateID())
	element.path = p
	if extendWithID {
		element.path = element.path.id(int(element.ID))
	}
	element.OperationKind = OperationKindUpdate
	engine.Patch.EquipmentSetEquipmentRef[element.ID] = element
	return element
}

func (engine *Engine) createPlayerEquipmentSetRef(p path, extendWithID bool, referencedElementID EquipmentSetID, parentID PlayerID) playerEquipmentSetRefCore {
	var element playerEquipmentSetRefCore
	element.engine = engine
	element.ReferencedElementID = referencedElementID
	element.ParentID = parentID
	element.ID = PlayerEquipmentSetRefID(engine.GenerateID())
	element.path = p
	if extendWithID {
		element.path = element.path.id(int(element.ID))
	}
	element.OperationKind = OperationKindUpdate
	engine.Patch.PlayerEquipmentSetRef[element.ID] = element
	return element
}

func (engine *Engine) createPlayerTargetRef(p path, extendWithID bool, referencedElementID AnyOfPlayer_ZoneItemID, parentID PlayerID) playerTargetRefCore {
	var element playerTargetRefCore
	element.engine = engine
	element.ReferencedElementID = referencedElementID
	element.ParentID = parentID
	element.ID = PlayerTargetRefID(engine.GenerateID())
	element.path = p
	if extendWithID {
		element.path = element.path.id(int(element.ID))
	}
	element.OperationKind = OperationKindUpdate
	engine.Patch.PlayerTargetRef[element.ID] = element
	return element
}

func (engine *Engine) createPlayerTargetedByRef(p path, extendWithID bool, referencedElementID AnyOfPlayer_ZoneItemID, parentID PlayerID) playerTargetedByRefCore {
	var element playerTargetedByRefCore
	element.engine = engine
	element.ReferencedElementID = referencedElementID
	element.ParentID = parentID
	element.ID = PlayerTargetedByRefID(engine.GenerateID())
	element.path = p
	if extendWithID {
		element.path = element.path.id(int(element.ID))
	}
	element.OperationKind = OperationKindUpdate
	engine.Patch.PlayerTargetedByRef[element.ID] = element
	return element
}

func (engine *Engine) createAnyOfPlayer_ZoneItem(setDefaultValue bool, childElementPath path) AnyOfPlayer_ZoneItem {
	var element anyOfPlayer_ZoneItemCore
	element.engine = engine
	element.ID = AnyOfPlayer_ZoneItemID(engine.GenerateID())
	if setDefaultValue {
		elementPlayer := engine.createPlayer(childElementPath, false)
		element.Player = elementPlayer.player.ID
		element.ElementKind = ElementKindPlayer
	}
	element.OperationKind = OperationKindUpdate
	element.ChildElementPath = childElementPath
	engine.Patch.AnyOfPlayer_ZoneItem[element.ID] = element
	return AnyOfPlayer_ZoneItem{anyOfPlayer_ZoneItem: element}
}

func (engine *Engine) createAnyOfPlayer_Position(setDefaultValue bool, childElementPath path) AnyOfPlayer_Position {
	var element anyOfPlayer_PositionCore
	element.engine = engine
	element.ID = AnyOfPlayer_PositionID(engine.GenerateID())
	if setDefaultValue {
		elementPlayer := engine.createPlayer(childElementPath, false)
		element.Player = elementPlayer.player.ID
		element.ElementKind = ElementKindPlayer
	}
	element.OperationKind = OperationKindUpdate
	element.ChildElementPath = childElementPath
	engine.Patch.AnyOfPlayer_Position[element.ID] = element
	return AnyOfPlayer_Position{anyOfPlayer_Position: element}
}

func (engine *Engine) createAnyOfItem_Player_ZoneItem(setDefaultValue bool, childElementPath path) AnyOfItem_Player_ZoneItem {
	var element anyOfItem_Player_ZoneItemCore
	element.engine = engine
	element.ID = AnyOfItem_Player_ZoneItemID(engine.GenerateID())
	if setDefaultValue {
		elementItem := engine.createItem(childElementPath, false)
		element.Item = elementItem.item.ID
		element.ElementKind = ElementKindItem
	}
	element.OperationKind = OperationKindUpdate
	element.ChildElementPath = childElementPath
	engine.Patch.AnyOfItem_Player_ZoneItem[element.ID] = element
	return AnyOfItem_Player_ZoneItem{anyOfItem_Player_ZoneItem: element}
}
