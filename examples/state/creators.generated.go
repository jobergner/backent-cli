package state

func (engine *Engine) createBoolValue(p path, fieldIdentifier treeFieldIdentifier, value bool) boolValue {
	var element boolValue
	element.Value = value
	element.engine = engine
	element.ID = BoolValueID(engine.GenerateID())
	element.Path = p.extendAndCopy(fieldIdentifier, int(element.ID), ElementKindBoolValue, 0)
	element.JSONPath = element.Path.toJSONPath()
	element.OperationKind = OperationKindUpdate
	engine.Patch.BoolValue[element.ID] = element
	return element
}

func (engine *Engine) createIntValue(p path, fieldIdentifier treeFieldIdentifier, value int64) intValue {
	var element intValue
	element.Value = value
	element.engine = engine
	element.ID = IntValueID(engine.GenerateID())
	element.Path = p.extendAndCopy(fieldIdentifier, int(element.ID), ElementKindIntValue, 0)
	element.JSONPath = element.Path.toJSONPath()
	element.OperationKind = OperationKindUpdate
	engine.Patch.IntValue[element.ID] = element
	return element
}

func (engine *Engine) createFloatValue(p path, fieldIdentifier treeFieldIdentifier, value float64) floatValue {
	var element floatValue
	element.Value = value
	element.engine = engine
	element.ID = FloatValueID(engine.GenerateID())
	element.Path = p.extendAndCopy(fieldIdentifier, int(element.ID), ElementKindFloatValue, 0)
	element.JSONPath = element.Path.toJSONPath()
	element.OperationKind = OperationKindUpdate
	engine.Patch.FloatValue[element.ID] = element
	return element
}

func (engine *Engine) createStringValue(p path, fieldIdentifier treeFieldIdentifier, value string) stringValue {
	var element stringValue
	element.Value = value
	element.engine = engine
	element.ID = StringValueID(engine.GenerateID())
	element.Path = p.extendAndCopy(fieldIdentifier, int(element.ID), ElementKindStringValue, 0)
	element.JSONPath = element.Path.toJSONPath()
	element.OperationKind = OperationKindUpdate
	engine.Patch.StringValue[element.ID] = element
	return element
}

func (engine *Engine) CreateEquipmentSet() EquipmentSet {
	return engine.createEquipmentSet(newPath(), equipmentSetIdentifier)
}

func (engine *Engine) createEquipmentSet(p path, fieldIdentifier treeFieldIdentifier) EquipmentSet {
	var element equipmentSetCore
	element.engine = engine
	element.ID = EquipmentSetID(engine.GenerateID())
	element.Path = p.extendAndCopy(fieldIdentifier, int(element.ID), ElementKindEquipmentSet, 0)
	element.JSONPath = element.Path.toJSONPath()
	elementName := engine.createStringValue(element.Path, equipmentSet_nameIdentifier, "")
	element.Name = elementName.ID
	element.OperationKind = OperationKindUpdate
	element.HasParent = len(element.Path) > 1
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
	element.Path = p.extendAndCopy(fieldIdentifier, int(element.ID), ElementKindGearScore, 0)
	element.JSONPath = element.Path.toJSONPath()
	elementLevel := engine.createIntValue(element.Path, gearScore_levelIdentifier, 0)
	element.Level = elementLevel.ID
	elementScore := engine.createIntValue(element.Path, gearScore_scoreIdentifier, 0)
	element.Score = elementScore.ID
	element.OperationKind = OperationKindUpdate
	element.HasParent = len(element.Path) > 1
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
	element.Path = p.extendAndCopy(fieldIdentifier, int(element.ID), ElementKindPosition, 0)
	element.JSONPath = element.Path.toJSONPath()
	elementX := engine.createFloatValue(element.Path, position_xIdentifier, 0.0)
	element.X = elementX.ID
	elementY := engine.createFloatValue(element.Path, position_yIdentifier, 0.0)
	element.Y = elementY.ID
	element.OperationKind = OperationKindUpdate
	element.HasParent = len(element.Path) > 1
	engine.Patch.Position[element.ID] = element
	return Position{position: element}
}

func (engine *Engine) CreateAttackEvent() AttackEvent {
	return engine.createAttackEvent(newPath(), attackEventIdentifier)
}

func (engine *Engine) createAttackEvent(p path, fieldIdentifier treeFieldIdentifier) AttackEvent {
	var element attackEventCore
	element.engine = engine
	element.ID = AttackEventID(engine.GenerateID())
	element.Path = p.extendAndCopy(fieldIdentifier, int(element.ID), ElementKindAttackEvent, 0)
	element.JSONPath = element.Path.toJSONPath()
	element.OperationKind = OperationKindUpdate
	element.HasParent = len(element.Path) > 1
	engine.Patch.AttackEvent[element.ID] = element
	return AttackEvent{attackEvent: element}
}

func (engine *Engine) CreateItem() Item {
	return engine.createItem(newPath(), itemIdentifier)
}

func (engine *Engine) createItem(p path, fieldIdentifier treeFieldIdentifier) Item {
	var element itemCore
	element.engine = engine
	element.ID = ItemID(engine.GenerateID())
	element.Path = p.extendAndCopy(fieldIdentifier, int(element.ID), ElementKindItem, 0)
	element.JSONPath = element.Path.toJSONPath()
	elementGearScore := engine.createGearScore(element.Path, item_gearScoreIdentifier)
	element.GearScore = elementGearScore.gearScore.ID
	elementName := engine.createStringValue(element.Path, item_nameIdentifier, "")
	element.Name = elementName.ID
	originElement := engine.createPlayer(element.Path, item_originIdentifier)
	elementOrigin := engine.createAnyOfPlayer_Position(int(element.ID), int(originElement.player.ID), ElementKindPlayer, element.Path, item_originIdentifier)
	element.Origin = elementOrigin.anyOfPlayer_Position.ID
	element.OperationKind = OperationKindUpdate
	element.HasParent = len(element.Path) > 1
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
	element.Path = p.extendAndCopy(fieldIdentifier, int(element.ID), ElementKindZoneItem, 0)
	element.JSONPath = element.Path.toJSONPath()
	elementItem := engine.createItem(element.Path, zoneItem_itemIdentifier)
	element.Item = elementItem.item.ID
	elementPosition := engine.createPosition(element.Path, zoneItem_positionIdentifier)
	element.Position = elementPosition.position.ID
	element.OperationKind = OperationKindUpdate
	element.HasParent = len(element.Path) > 1
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
	element.Path = p.extendAndCopy(fieldIdentifier, int(element.ID), ElementKindPlayer, 0)
	element.JSONPath = element.Path.toJSONPath()
	elementGearScore := engine.createGearScore(element.Path, player_gearScoreIdentifier)
	element.GearScore = elementGearScore.gearScore.ID
	elementPosition := engine.createPosition(element.Path, player_positionIdentifier)
	element.Position = elementPosition.position.ID
	element.OperationKind = OperationKindUpdate
	element.HasParent = len(element.Path) > 1
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
	element.Path = p.extendAndCopy(fieldIdentifier, int(element.ID), ElementKindZone, 0)
	element.JSONPath = element.Path.toJSONPath()
	element.OperationKind = OperationKindUpdate
	element.HasParent = len(element.Path) > 1
	engine.Patch.Zone[element.ID] = element
	return Zone{zone: element}
}

func (engine *Engine) createAttackEventTargetRef(p path, fieldIdentifier treeFieldIdentifier, referencedElementID PlayerID, parentID AttackEventID, childID int) attackEventTargetRefCore {
	var element attackEventTargetRefCore
	element.engine = engine
	element.ReferencedElementID = referencedElementID
	element.ParentID = parentID
	element.ChildID = childID
	element.ID = AttackEventTargetRefID(engine.GenerateID())
	element.Path = p.extendAndCopy(fieldIdentifier, 0, ElementKindPlayer, int(element.ID))
	element.OperationKind = OperationKindUpdate
	engine.Patch.AttackEventTargetRef[element.ID] = element
	return element
}

func (engine *Engine) createItemBoundToRef(p path, fieldIdentifier treeFieldIdentifier, referencedElementID PlayerID, parentID ItemID, childID int) itemBoundToRefCore {
	var element itemBoundToRefCore
	element.engine = engine
	element.ReferencedElementID = referencedElementID
	element.ParentID = parentID
	element.ChildID = childID
	element.ID = ItemBoundToRefID(engine.GenerateID())
	element.Path = p.extendAndCopy(fieldIdentifier, 0, ElementKindPlayer, int(element.ID))
	element.OperationKind = OperationKindUpdate
	engine.Patch.ItemBoundToRef[element.ID] = element
	return element
}

func (engine *Engine) createPlayerGuildMemberRef(p path, fieldIdentifier treeFieldIdentifier, referencedElementID PlayerID, parentID PlayerID, childID int) playerGuildMemberRefCore {
	var element playerGuildMemberRefCore
	element.engine = engine
	element.ReferencedElementID = referencedElementID
	element.ParentID = parentID
	element.ChildID = childID
	element.ID = PlayerGuildMemberRefID(engine.GenerateID())
	element.Path = p.extendAndCopy(fieldIdentifier, 0, ElementKindPlayer, int(element.ID))
	element.OperationKind = OperationKindUpdate
	engine.Patch.PlayerGuildMemberRef[element.ID] = element
	return element
}

func (engine *Engine) createEquipmentSetEquipmentRef(p path, fieldIdentifier treeFieldIdentifier, referencedElementID ItemID, parentID EquipmentSetID, childID int) equipmentSetEquipmentRefCore {
	var element equipmentSetEquipmentRefCore
	element.engine = engine
	element.ReferencedElementID = referencedElementID
	element.ParentID = parentID
	element.ChildID = childID
	element.ID = EquipmentSetEquipmentRefID(engine.GenerateID())
	element.Path = p.extendAndCopy(fieldIdentifier, 0, ElementKindItem, int(element.ID))
	element.OperationKind = OperationKindUpdate
	engine.Patch.EquipmentSetEquipmentRef[element.ID] = element
	return element
}

func (engine *Engine) createPlayerEquipmentSetRef(p path, fieldIdentifier treeFieldIdentifier, referencedElementID EquipmentSetID, parentID PlayerID, childID int) playerEquipmentSetRefCore {
	var element playerEquipmentSetRefCore
	element.engine = engine
	element.ReferencedElementID = referencedElementID
	element.ParentID = parentID
	element.ChildID = childID
	element.ID = PlayerEquipmentSetRefID(engine.GenerateID())
	element.Path = p.extendAndCopy(fieldIdentifier, 0, ElementKindEquipmentSet, int(element.ID))
	element.OperationKind = OperationKindUpdate
	engine.Patch.PlayerEquipmentSetRef[element.ID] = element
	return element
}

func (engine *Engine) createPlayerTargetRef(p path, fieldIdentifier treeFieldIdentifier, referencedElementID AnyOfPlayer_ZoneItemID, parentID PlayerID, childKind ElementKind, childID int) playerTargetRefCore {
	var element playerTargetRefCore
	element.engine = engine
	element.ReferencedElementID = referencedElementID
	element.ParentID = parentID
	element.ChildID = childID
	element.ID = PlayerTargetRefID(engine.GenerateID())
	element.Path = p.extendAndCopy(fieldIdentifier, 0, childKind, int(element.ID))
	element.OperationKind = OperationKindUpdate
	engine.Patch.PlayerTargetRef[element.ID] = element
	return element
}

func (engine *Engine) createPlayerTargetedByRef(p path, fieldIdentifier treeFieldIdentifier, referencedElementID AnyOfPlayer_ZoneItemID, parentID PlayerID, childKind ElementKind, childID int) playerTargetedByRefCore {
	var element playerTargetedByRefCore
	element.engine = engine
	element.ReferencedElementID = referencedElementID
	element.ParentID = parentID
	element.ChildID = childID
	element.ID = PlayerTargetedByRefID(engine.GenerateID())
	element.Path = p.extendAndCopy(fieldIdentifier, 0, childKind, int(element.ID))
	element.OperationKind = OperationKindUpdate
	engine.Patch.PlayerTargetedByRef[element.ID] = element
	return element
}

func (engine *Engine) createAnyOfPlayer_ZoneItem(parentID int, childID int, childKind ElementKind, p path, fieldIdentifier treeFieldIdentifier) AnyOfPlayer_ZoneItem {
	var element anyOfPlayer_ZoneItemCore
	element.engine = engine
	element.ID = AnyOfPlayer_ZoneItemID(engine.GenerateID())
	element.ParentID = parentID
	element.ChildID = childID
	element.ElementKind = childKind
	element.OperationKind = OperationKindUpdate
	element.ParentElementPath = p
	element.FieldIdentifier = fieldIdentifier
	engine.Patch.AnyOfPlayer_ZoneItem[element.ID] = element
	return AnyOfPlayer_ZoneItem{anyOfPlayer_ZoneItem: element}
}

func (engine *Engine) createAnyOfPlayer_Position(parentID int, childID int, childKind ElementKind, p path, fieldIdentifier treeFieldIdentifier) AnyOfPlayer_Position {
	var element anyOfPlayer_PositionCore
	element.engine = engine
	element.ID = AnyOfPlayer_PositionID(engine.GenerateID())
	element.ParentID = parentID
	element.ChildID = childID
	element.ElementKind = childKind
	element.OperationKind = OperationKindUpdate
	element.ParentElementPath = p
	element.FieldIdentifier = fieldIdentifier
	engine.Patch.AnyOfPlayer_Position[element.ID] = element
	return AnyOfPlayer_Position{anyOfPlayer_Position: element}
}

func (engine *Engine) createAnyOfItem_Player_ZoneItem(parentID int, childID int, childKind ElementKind, p path, fieldIdentifier treeFieldIdentifier) AnyOfItem_Player_ZoneItem {
	var element anyOfItem_Player_ZoneItemCore
	element.engine = engine
	element.ID = AnyOfItem_Player_ZoneItemID(engine.GenerateID())
	element.ParentID = parentID
	element.ChildID = childID
	element.ElementKind = childKind
	element.OperationKind = OperationKindUpdate
	element.ParentElementPath = p
	element.FieldIdentifier = fieldIdentifier
	engine.Patch.AnyOfItem_Player_ZoneItem[element.ID] = element
	return AnyOfItem_Player_ZoneItem{anyOfItem_Player_ZoneItem: element}
}
