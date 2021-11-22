package state

func (engine *Engine) CreateEquipmentSet() EquipmentSet {
	return engine.createEquipmentSet(newPath(), equipmentSetIdentifier)
}

func (engine *Engine) createEquipmentSet(p path, fieldIdentifier treeFieldIdentifier) EquipmentSet {
	var element equipmentSetCore
	element.engine = engine
	element.ID = EquipmentSetID(engine.GenerateID())
	element.path = p.extendAndCopy(fieldIdentifier, int(element.ID), ElementKindEquipmentSet, 0)
	element.Path = element.path.toJSONPath()
	element.OperationKind = OperationKindUpdate
	element.HasParent = len(element.path) > 1
	engine.Patch.EquipmentSet[element.ID] = element
	return EquipmentSet{equipmentSet: element}
}

func (engine *Engine) CreateGearScore() GearScore {
	return engine.createGearScore(newPath(), gearScoreIdentifier)
}

func (engine *Engine) createGearScore(p path, fieldIdentifier treeFieldIdentifier) GearScore {
	var element gearScoreCore
	element.engine = engine
	element.ID = GearScoreID(engine.GenerateID())
	element.path = p.extendAndCopy(fieldIdentifier, int(element.ID), ElementKindGearScore, 0)
	element.Path = element.path.toJSONPath()
	element.OperationKind = OperationKindUpdate
	element.HasParent = len(element.path) > 1
	engine.Patch.GearScore[element.ID] = element
	return GearScore{gearScore: element}
}

func (engine *Engine) CreatePosition() Position {
	return engine.createPosition(newPath(), positionIdentifier)
}

func (engine *Engine) createPosition(p path, fieldIdentifier treeFieldIdentifier) Position {
	var element positionCore
	element.engine = engine
	element.ID = PositionID(engine.GenerateID())
	element.path = p.extendAndCopy(fieldIdentifier, int(element.ID), ElementKindPosition, 0)
	element.Path = element.path.toJSONPath()
	element.OperationKind = OperationKindUpdate
	element.HasParent = len(element.path) > 1
	engine.Patch.Position[element.ID] = element
	return Position{position: element}
}

func (engine *Engine) CreateItem() Item {
	return engine.createItem(newPath(), itemIdentifier)
}

func (engine *Engine) createItem(p path, fieldIdentifier treeFieldIdentifier) Item {
	var element itemCore
	element.engine = engine
	element.ID = ItemID(engine.GenerateID())
	element.path = p.extendAndCopy(fieldIdentifier, int(element.ID), ElementKindItem, 0)
	element.Path = element.path.toJSONPath()
	elementGearScore := engine.createGearScore(element.path, item_gearScoreIdentifier)
	element.GearScore = elementGearScore.gearScore.ID
	elementOrigin := engine.createAnyOfPlayer_Position(true, element.path, item_originIdentifier)
	element.Origin = elementOrigin.anyOfPlayer_Position.ID
	element.OperationKind = OperationKindUpdate
	element.HasParent = len(element.path) > 1
	engine.Patch.Item[element.ID] = element
	return Item{item: element}
}

func (engine *Engine) CreateZoneItem() ZoneItem {
	return engine.createZoneItem(newPath(), zoneItemIdentifier)
}

func (engine *Engine) createZoneItem(p path, fieldIdentifier treeFieldIdentifier) ZoneItem {
	var element zoneItemCore
	element.engine = engine
	element.ID = ZoneItemID(engine.GenerateID())
	element.path = p.extendAndCopy(fieldIdentifier, int(element.ID), ElementKindZoneItem, 0)
	element.Path = element.path.toJSONPath()
	elementItem := engine.createItem(element.path, zoneItem_itemIdentifier)
	element.Item = elementItem.item.ID
	elementPosition := engine.createPosition(element.path, zoneItem_positionIdentifier)
	element.Position = elementPosition.position.ID
	element.OperationKind = OperationKindUpdate
	element.HasParent = len(element.path) > 1
	engine.Patch.ZoneItem[element.ID] = element
	return ZoneItem{zoneItem: element}
}

func (engine *Engine) CreatePlayer() Player {
	return engine.createPlayer(newPath(), playerIdentifier)
}

func (engine *Engine) createPlayer(p path, fieldIdentifier treeFieldIdentifier) Player {
	var element playerCore
	element.engine = engine
	element.ID = PlayerID(engine.GenerateID())
	element.path = p.extendAndCopy(fieldIdentifier, int(element.ID), ElementKindPlayer, 0)
	element.Path = element.path.toJSONPath()
	elementGearScore := engine.createGearScore(element.path, player_gearScoreIdentifier)
	element.GearScore = elementGearScore.gearScore.ID
	elementPosition := engine.createPosition(element.path, player_positionIdentifier)
	element.Position = elementPosition.position.ID
	element.OperationKind = OperationKindUpdate
	element.HasParent = len(element.path) > 1
	engine.Patch.Player[element.ID] = element
	return Player{player: element}
}

func (engine *Engine) CreateZone() Zone {
	return engine.createZone(newPath(), zoneIdentifier)
}

func (engine *Engine) createZone(p path, fieldIdentifier treeFieldIdentifier) Zone {
	var element zoneCore
	element.engine = engine
	element.ID = ZoneID(engine.GenerateID())
	element.path = p.extendAndCopy(fieldIdentifier, int(element.ID), ElementKindZone, 0)
	element.Path = element.path.toJSONPath()
	element.OperationKind = OperationKindUpdate
	element.HasParent = len(element.path) > 1
	engine.Patch.Zone[element.ID] = element
	return Zone{zone: element}
}

func (engine *Engine) createItemBoundToRef(p path, fieldIdentifier treeFieldIdentifier, referencedElementID PlayerID, parentID ItemID) itemBoundToRefCore {
	var element itemBoundToRefCore
	element.engine = engine
	element.ReferencedElementID = referencedElementID
	element.ParentID = parentID
	element.ID = ItemBoundToRefID(engine.GenerateID())
	element.path = p.extendAndCopy(fieldIdentifier, int(element.ID), ElementKindPlayer, int(element.ID))
	element.OperationKind = OperationKindUpdate
	engine.Patch.ItemBoundToRef[element.ID] = element
	return element
}

func (engine *Engine) createPlayerGuildMemberRef(p path, fieldIdentifier treeFieldIdentifier, referencedElementID PlayerID, parentID PlayerID) playerGuildMemberRefCore {
	var element playerGuildMemberRefCore
	element.engine = engine
	element.ReferencedElementID = referencedElementID
	element.ParentID = parentID
	element.ID = PlayerGuildMemberRefID(engine.GenerateID())
	element.path = p.extendAndCopy(fieldIdentifier, int(element.ID), ElementKindPlayer, int(element.ID))
	element.OperationKind = OperationKindUpdate
	engine.Patch.PlayerGuildMemberRef[element.ID] = element
	return element
}

func (engine *Engine) createEquipmentSetEquipmentRef(p path, fieldIdentifier treeFieldIdentifier, referencedElementID ItemID, parentID EquipmentSetID) equipmentSetEquipmentRefCore {
	var element equipmentSetEquipmentRefCore
	element.engine = engine
	element.ReferencedElementID = referencedElementID
	element.ParentID = parentID
	element.ID = EquipmentSetEquipmentRefID(engine.GenerateID())
	element.path = p.extendAndCopy(fieldIdentifier, int(element.ID), ElementKindItem, int(element.ID))
	element.OperationKind = OperationKindUpdate
	engine.Patch.EquipmentSetEquipmentRef[element.ID] = element
	return element
}

func (engine *Engine) createPlayerEquipmentSetRef(p path, fieldIdentifier treeFieldIdentifier, referencedElementID EquipmentSetID, parentID PlayerID) playerEquipmentSetRefCore {
	var element playerEquipmentSetRefCore
	element.engine = engine
	element.ReferencedElementID = referencedElementID
	element.ParentID = parentID
	element.ID = PlayerEquipmentSetRefID(engine.GenerateID())
	element.path = p.extendAndCopy(fieldIdentifier, int(element.ID), ElementKindEquipmentSet, int(element.ID))
	element.OperationKind = OperationKindUpdate
	engine.Patch.PlayerEquipmentSetRef[element.ID] = element
	return element
}

func (engine *Engine) createPlayerTargetRef(p path, fieldIdentifier treeFieldIdentifier, referencedElementID AnyOfPlayer_ZoneItemID, parentID PlayerID, childKind ElementKind, childID int) playerTargetRefCore {
	var element playerTargetRefCore
	element.engine = engine
	element.ReferencedElementID = referencedElementID
	element.ParentID = parentID
	element.ID = PlayerTargetRefID(engine.GenerateID())
	element.path = p.extendAndCopy(fieldIdentifier, childID, childKind, int(element.ID))
	element.OperationKind = OperationKindUpdate
	engine.Patch.PlayerTargetRef[element.ID] = element
	return element
}

func (engine *Engine) createPlayerTargetedByRef(p path, fieldIdentifier treeFieldIdentifier, referencedElementID AnyOfPlayer_ZoneItemID, parentID PlayerID, childKind ElementKind, childID int) playerTargetedByRefCore {
	var element playerTargetedByRefCore
	element.engine = engine
	element.ReferencedElementID = referencedElementID
	element.ParentID = parentID
	element.ID = PlayerTargetedByRefID(engine.GenerateID())
	element.path = p.extendAndCopy(fieldIdentifier, childID, childKind, int(element.ID))
	element.OperationKind = OperationKindUpdate
	engine.Patch.PlayerTargetedByRef[element.ID] = element
	return element
}

func (engine *Engine) createAnyOfPlayer_ZoneItem(setDefaultValue bool, p path, fieldIdentifier treeFieldIdentifier) AnyOfPlayer_ZoneItem {
	var element anyOfPlayer_ZoneItemCore
	element.engine = engine
	element.ID = AnyOfPlayer_ZoneItemID(engine.GenerateID())
	if setDefaultValue {
		elementPlayer := engine.createPlayer(p, fieldIdentifier)
		element.Player = elementPlayer.player.ID
		element.ElementKind = ElementKindPlayer
	}
	element.OperationKind = OperationKindUpdate
	element.ParentElementPath = p
	element.FieldIdentifier = fieldIdentifier
	engine.Patch.AnyOfPlayer_ZoneItem[element.ID] = element
	return AnyOfPlayer_ZoneItem{anyOfPlayer_ZoneItem: element}
}

func (engine *Engine) createAnyOfPlayer_Position(setDefaultValue bool, p path, fieldIdentifier treeFieldIdentifier) AnyOfPlayer_Position {
	var element anyOfPlayer_PositionCore
	element.engine = engine
	element.ID = AnyOfPlayer_PositionID(engine.GenerateID())
	if setDefaultValue {
		elementPlayer := engine.createPlayer(p, fieldIdentifier)
		element.Player = elementPlayer.player.ID
		element.ElementKind = ElementKindPlayer
	}
	element.OperationKind = OperationKindUpdate
	element.ParentElementPath = p
	element.FieldIdentifier = fieldIdentifier
	engine.Patch.AnyOfPlayer_Position[element.ID] = element
	return AnyOfPlayer_Position{anyOfPlayer_Position: element}
}

func (engine *Engine) createAnyOfItem_Player_ZoneItem(setDefaultValue bool, p path, fieldIdentifier treeFieldIdentifier) AnyOfItem_Player_ZoneItem {
	var element anyOfItem_Player_ZoneItemCore
	element.engine = engine
	element.ID = AnyOfItem_Player_ZoneItemID(engine.GenerateID())
	if setDefaultValue {
		elementItem := engine.createItem(p, fieldIdentifier)
		element.Item = elementItem.item.ID
		element.ElementKind = ElementKindItem
	}
	element.OperationKind = OperationKindUpdate
	element.ParentElementPath = p
	element.FieldIdentifier = fieldIdentifier
	engine.Patch.AnyOfItem_Player_ZoneItem[element.ID] = element
	return AnyOfItem_Player_ZoneItem{anyOfItem_Player_ZoneItem: element}
}
