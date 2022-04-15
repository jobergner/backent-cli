package state

type anyOfPlayer_PositionRef struct {
	anyOfPlayer_PositionWrapper AnyOfPlayer_Position
	anyOfPlayer_Position        anyOfPlayer_PositionCore
}

func (_any anyOfPlayer_PositionRef) Kind() ElementKind {
	return _any.anyOfPlayer_PositionWrapper.Kind()
}

func (_any anyOfPlayer_PositionRef) Player() Player {
	return _any.anyOfPlayer_PositionWrapper.Player()
}

func (_any anyOfPlayer_PositionRef) Position() Position {
	return _any.anyOfPlayer_PositionWrapper.Position()
}

type anyOfPlayer_ZoneItemRef struct {
	anyOfPlayer_ZoneItemWrapper AnyOfPlayer_ZoneItem
	anyOfPlayer_ZoneItem        anyOfPlayer_ZoneItemCore
}

func (_any anyOfPlayer_ZoneItemRef) Kind() ElementKind {
	return _any.anyOfPlayer_ZoneItemWrapper.Kind()
}

func (_any anyOfPlayer_ZoneItemRef) Player() Player {
	return _any.anyOfPlayer_ZoneItemWrapper.Player()
}

func (_any anyOfPlayer_ZoneItemRef) ZoneItem() ZoneItem {
	return _any.anyOfPlayer_ZoneItemWrapper.ZoneItem()
}

type anyOfItem_Player_ZoneItemRef struct {
	anyOfItem_Player_ZoneItemWrapper AnyOfItem_Player_ZoneItem
	anyOfItem_Player_ZoneItem        anyOfItem_Player_ZoneItemCore
}

func (_any anyOfItem_Player_ZoneItemRef) Kind() ElementKind {
	return _any.anyOfItem_Player_ZoneItemWrapper.Kind()
}

func (_any anyOfItem_Player_ZoneItemRef) Item() Item {
	return _any.anyOfItem_Player_ZoneItemWrapper.Item()
}

func (_any anyOfItem_Player_ZoneItemRef) Player() Player {
	return _any.anyOfItem_Player_ZoneItemWrapper.Player()
}

func (_any anyOfItem_Player_ZoneItemRef) ZoneItem() ZoneItem {
	return _any.anyOfItem_Player_ZoneItemWrapper.ZoneItem()
}

func (_any AnyOfPlayer_ZoneItem) Kind() ElementKind {
	any := _any.anyOfPlayer_ZoneItem.engine.anyOfPlayer_ZoneItem(_any.anyOfPlayer_ZoneItem.ID)
	return any.anyOfPlayer_ZoneItem.ElementKind
}

func (_any AnyOfPlayer_ZoneItem) BeZoneItem() ZoneItem {
	any := _any.anyOfPlayer_ZoneItem.engine.anyOfPlayer_ZoneItem(_any.anyOfPlayer_ZoneItem.ID)
	if any.anyOfPlayer_ZoneItem.ElementKind == ElementKindZoneItem || any.anyOfPlayer_ZoneItem.OperationKind == OperationKindDelete {
		return any.ZoneItem()
	}
	zoneItem := any.anyOfPlayer_ZoneItem.engine.createZoneItem(any.anyOfPlayer_ZoneItem.ParentElementPath, any.anyOfPlayer_ZoneItem.FieldIdentifier)
	any.anyOfPlayer_ZoneItem.beZoneItem(zoneItem.ID(), true)
	return zoneItem
}

func (_any anyOfPlayer_ZoneItemCore) beZoneItem(zoneItemID ZoneItemID, deleteCurrentChild bool) {
	any := _any.engine.anyOfPlayer_ZoneItem(_any.ID).anyOfPlayer_ZoneItem
	any.engine.deleteAnyOfPlayer_ZoneItem(any.ID, deleteCurrentChild)
	any = any.engine.createAnyOfPlayer_ZoneItem(any.ID.ParentID, int(zoneItemID), ElementKindZoneItem, any.ParentElementPath, any.FieldIdentifier).anyOfPlayer_ZoneItem
	switch any.FieldIdentifier {
	}
	any.Meta.sign(any.engine.BroadcastingClientID)
	any.engine.Patch.AnyOfPlayer_ZoneItem[any.ID] = any
}

func (_any AnyOfPlayer_ZoneItem) BePlayer() Player {
	any := _any.anyOfPlayer_ZoneItem.engine.anyOfPlayer_ZoneItem(_any.anyOfPlayer_ZoneItem.ID)
	if any.anyOfPlayer_ZoneItem.ElementKind == ElementKindPlayer || any.anyOfPlayer_ZoneItem.OperationKind == OperationKindDelete {
		return any.Player()
	}
	player := any.anyOfPlayer_ZoneItem.engine.createPlayer(any.anyOfPlayer_ZoneItem.ParentElementPath, any.anyOfPlayer_ZoneItem.FieldIdentifier)
	any.anyOfPlayer_ZoneItem.bePlayer(player.ID(), true)
	return player
}

func (_any anyOfPlayer_ZoneItemCore) bePlayer(playerID PlayerID, deleteCurrentChild bool) {
	any := _any.engine.anyOfPlayer_ZoneItem(_any.ID).anyOfPlayer_ZoneItem
	any.engine.deleteAnyOfPlayer_ZoneItem(any.ID, deleteCurrentChild)
	any = any.engine.createAnyOfPlayer_ZoneItem(any.ID.ParentID, int(playerID), ElementKindPlayer, any.ParentElementPath, any.FieldIdentifier).anyOfPlayer_ZoneItem
	switch any.FieldIdentifier {
	}
	any.Meta.sign(any.engine.BroadcastingClientID)
	any.engine.Patch.AnyOfPlayer_ZoneItem[any.ID] = any
}

func (_any anyOfPlayer_ZoneItemCore) deleteChild() {
	any := _any.engine.anyOfPlayer_ZoneItem(_any.ID).anyOfPlayer_ZoneItem
	switch any.ElementKind {
	case ElementKindPlayer:
		any.engine.deletePlayer(PlayerID(any.ChildID))
	case ElementKindZoneItem:
		any.engine.deleteZoneItem(ZoneItemID(any.ChildID))
	}
}

func (_any AnyOfPlayer_Position) Kind() ElementKind {
	any := _any.anyOfPlayer_Position.engine.anyOfPlayer_Position(_any.anyOfPlayer_Position.ID)
	return any.anyOfPlayer_Position.ElementKind
}

func (_any AnyOfPlayer_Position) BePosition() Position {
	any := _any.anyOfPlayer_Position.engine.anyOfPlayer_Position(_any.anyOfPlayer_Position.ID)
	if any.anyOfPlayer_Position.ElementKind == ElementKindPosition || any.anyOfPlayer_Position.OperationKind == OperationKindDelete {
		return any.Position()
	}
	position := any.anyOfPlayer_Position.engine.createPosition(any.anyOfPlayer_Position.ParentElementPath, any.anyOfPlayer_Position.FieldIdentifier)
	any.anyOfPlayer_Position.bePosition(position.ID(), true)
	return position
}

func (_any anyOfPlayer_PositionCore) bePosition(positionID PositionID, deleteCurrentChild bool) {
	any := _any.engine.anyOfPlayer_Position(_any.ID).anyOfPlayer_Position
	any.engine.deleteAnyOfPlayer_Position(any.ID, deleteCurrentChild)
	any = any.engine.createAnyOfPlayer_Position(any.ID.ParentID, int(positionID), ElementKindPosition, any.ParentElementPath, any.FieldIdentifier).anyOfPlayer_Position
	switch any.FieldIdentifier {
	case item_originIdentifier:
		item := any.engine.Item(ItemID(any.ID.ParentID)).item
		item.Origin = any.ID
		item.Meta.sign(item.engine.BroadcastingClientID)
		item.engine.Patch.Item[item.ID] = item
		// we do not set OperationKindUpdate on purpose as it technicaly has not any updated values
		// however it still has to be put in patch
	}
	any.Meta.sign(any.engine.BroadcastingClientID)
	any.engine.Patch.AnyOfPlayer_Position[any.ID] = any
}

func (_any anyOfPlayer_PositionCore) deleteChild() {
	any := _any.engine.anyOfPlayer_Position(_any.ID).anyOfPlayer_Position
	switch any.ElementKind {
	case ElementKindPlayer:
		any.engine.deletePlayer(PlayerID(any.ChildID))
	case ElementKindPosition:
		any.engine.deletePosition(PositionID(any.ChildID))
	}
}

func (_any AnyOfPlayer_Position) BePlayer() Player {
	any := _any.anyOfPlayer_Position.engine.anyOfPlayer_Position(_any.anyOfPlayer_Position.ID)
	if any.anyOfPlayer_Position.ElementKind == ElementKindPlayer || any.anyOfPlayer_Position.OperationKind == OperationKindDelete {
		return any.Player()
	}
	player := any.anyOfPlayer_Position.engine.createPlayer(any.anyOfPlayer_Position.ParentElementPath, any.anyOfPlayer_Position.FieldIdentifier)
	any.anyOfPlayer_Position.bePlayer(player.ID(), true)
	return player
}

func (_any anyOfPlayer_PositionCore) bePlayer(playerID PlayerID, deleteCurrentChild bool) {
	any := _any.engine.anyOfPlayer_Position(_any.ID).anyOfPlayer_Position
	any.engine.deleteAnyOfPlayer_Position(any.ID, deleteCurrentChild)
	any = any.engine.createAnyOfPlayer_Position(any.ID.ParentID, int(playerID), ElementKindPlayer, any.ParentElementPath, any.FieldIdentifier).anyOfPlayer_Position
	switch any.FieldIdentifier {
	case item_originIdentifier:
		item := any.engine.Item(ItemID(any.ID.ParentID)).item
		item.Origin = any.ID
		item.Meta.sign(item.engine.BroadcastingClientID)
		item.engine.Patch.Item[item.ID] = item
	}
	any.Meta.sign(any.engine.BroadcastingClientID)
	any.engine.Patch.AnyOfPlayer_Position[any.ID] = any
}

func (_any AnyOfItem_Player_ZoneItem) Kind() ElementKind {
	any := _any.anyOfItem_Player_ZoneItem.engine.anyOfItem_Player_ZoneItem(_any.anyOfItem_Player_ZoneItem.ID)
	return any.anyOfItem_Player_ZoneItem.ElementKind
}

func (_any AnyOfItem_Player_ZoneItem) BeZoneItem() ZoneItem {
	any := _any.anyOfItem_Player_ZoneItem.engine.anyOfItem_Player_ZoneItem(_any.anyOfItem_Player_ZoneItem.ID)
	if any.anyOfItem_Player_ZoneItem.ElementKind == ElementKindZoneItem || any.anyOfItem_Player_ZoneItem.OperationKind == OperationKindDelete {
		return any.ZoneItem()
	}
	zoneItem := any.anyOfItem_Player_ZoneItem.engine.createZoneItem(any.anyOfItem_Player_ZoneItem.ParentElementPath, any.anyOfItem_Player_ZoneItem.FieldIdentifier)
	any.anyOfItem_Player_ZoneItem.beZoneItem(zoneItem.ID(), true)
	return zoneItem
}

func (_any anyOfItem_Player_ZoneItemCore) beZoneItem(zoneItemID ZoneItemID, deleteCurrentChild bool) {
	any := _any.engine.anyOfItem_Player_ZoneItem(_any.ID).anyOfItem_Player_ZoneItem
	any.engine.deleteAnyOfItem_Player_ZoneItem(any.ID, deleteCurrentChild)
	any = any.engine.createAnyOfItem_Player_ZoneItem(any.ID.ParentID, int(zoneItemID), ElementKindZoneItem, any.ParentElementPath, any.FieldIdentifier).anyOfItem_Player_ZoneItem
	switch any.FieldIdentifier {
	}
	any.Meta.sign(any.engine.BroadcastingClientID)
	any.engine.Patch.AnyOfItem_Player_ZoneItem[any.ID] = any
}

func (_any AnyOfItem_Player_ZoneItem) BePlayer() Player {
	any := _any.anyOfItem_Player_ZoneItem.engine.anyOfItem_Player_ZoneItem(_any.anyOfItem_Player_ZoneItem.ID)
	if any.anyOfItem_Player_ZoneItem.ElementKind == ElementKindPlayer || any.anyOfItem_Player_ZoneItem.OperationKind == OperationKindDelete {
		return any.Player()
	}
	player := any.anyOfItem_Player_ZoneItem.engine.createPlayer(any.anyOfItem_Player_ZoneItem.ParentElementPath, any.anyOfItem_Player_ZoneItem.FieldIdentifier)
	any.anyOfItem_Player_ZoneItem.bePlayer(player.ID(), true)
	return player
}

func (_any anyOfItem_Player_ZoneItemCore) bePlayer(playerID PlayerID, deleteCurrentChild bool) {
	any := _any.engine.anyOfItem_Player_ZoneItem(_any.ID).anyOfItem_Player_ZoneItem
	any.engine.deleteAnyOfItem_Player_ZoneItem(any.ID, deleteCurrentChild)
	any = any.engine.createAnyOfItem_Player_ZoneItem(any.ID.ParentID, int(playerID), ElementKindPlayer, any.ParentElementPath, any.FieldIdentifier).anyOfItem_Player_ZoneItem
	switch any.FieldIdentifier {
	}
	any.Meta.sign(any.engine.BroadcastingClientID)
	any.engine.Patch.AnyOfItem_Player_ZoneItem[any.ID] = any
}

func (_any AnyOfItem_Player_ZoneItem) BeItem() Item {
	any := _any.anyOfItem_Player_ZoneItem.engine.anyOfItem_Player_ZoneItem(_any.anyOfItem_Player_ZoneItem.ID)
	if any.anyOfItem_Player_ZoneItem.ElementKind == ElementKindItem || any.anyOfItem_Player_ZoneItem.OperationKind == OperationKindDelete {
		return any.Item()
	}
	item := any.anyOfItem_Player_ZoneItem.engine.createItem(any.anyOfItem_Player_ZoneItem.ParentElementPath, any.anyOfItem_Player_ZoneItem.FieldIdentifier)
	any.anyOfItem_Player_ZoneItem.beItem(item.ID(), true)
	return item
}

func (_any anyOfItem_Player_ZoneItemCore) beItem(itemID ItemID, deleteCurrentChild bool) {
	any := _any.engine.anyOfItem_Player_ZoneItem(_any.ID).anyOfItem_Player_ZoneItem
	any.engine.deleteAnyOfItem_Player_ZoneItem(any.ID, deleteCurrentChild)
	any = any.engine.createAnyOfItem_Player_ZoneItem(any.ID.ParentID, int(itemID), ElementKindItem, any.ParentElementPath, any.FieldIdentifier).anyOfItem_Player_ZoneItem
	switch any.FieldIdentifier {
	}
	any.Meta.sign(any.engine.BroadcastingClientID)
	any.engine.Patch.AnyOfItem_Player_ZoneItem[any.ID] = any
}

func (_any anyOfItem_Player_ZoneItemCore) deleteChild() {
	any := _any.engine.anyOfItem_Player_ZoneItem(_any.ID).anyOfItem_Player_ZoneItem
	switch any.ElementKind {
	case ElementKindItem:
		any.engine.deleteItem(ItemID(any.ChildID))
	case ElementKindPlayer:
		any.engine.deletePlayer(PlayerID(any.ChildID))
	case ElementKindZoneItem:
		any.engine.deleteZoneItem(ZoneItemID(any.ChildID))
	}
}
