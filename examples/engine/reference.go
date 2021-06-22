package state

func (_ref itemBoundToRef) IsSet() bool {
	ref := _ref.itemBoundToRef.engine.itemBoundToRef(_ref.itemBoundToRef.ID)
	return ref.itemBoundToRef.ID != 0
}

func (_ref itemBoundToRef) Unset() {
	ref := _ref.itemBoundToRef.engine.itemBoundToRef(_ref.itemBoundToRef.ID)
	ref.itemBoundToRef.engine.deleteItemBoundToRef(ref.itemBoundToRef.ID)
	parent := ref.itemBoundToRef.engine.Item(ref.itemBoundToRef.ParentID).item
	if parent.OperationKind == OperationKindDelete {
		return
	}
	parent.BoundTo = 0
	parent.OperationKind = OperationKindUpdate
	ref.itemBoundToRef.engine.Patch.Item[parent.ID] = parent
}

func (_ref itemBoundToRef) Get() player {
	ref := _ref.itemBoundToRef.engine.itemBoundToRef(_ref.itemBoundToRef.ID)
	return ref.itemBoundToRef.engine.Player(ref.itemBoundToRef.ReferencedElementID)
}

func (_ref playerGuildMemberRef) Get() player {
	ref := _ref.playerGuildMemberRef.engine.playerGuildMemberRef(_ref.playerGuildMemberRef.ID)
	return ref.playerGuildMemberRef.engine.Player(ref.playerGuildMemberRef.ReferencedElementID)
}

func (_ref playerEquipmentSetRef) Get() equipmentSet {
	ref := _ref.playerEquipmentSetRef.engine.playerEquipmentSetRef(_ref.playerEquipmentSetRef.ID)
	return ref.playerEquipmentSetRef.engine.EquipmentSet(ref.playerEquipmentSetRef.ReferencedElementID)
}

func (_ref equipmentSetEquipmentRef) Get() item {
	ref := _ref.equipmentSetEquipmentRef.engine.equipmentSetEquipmentRef(_ref.equipmentSetEquipmentRef.ID)
	return ref.equipmentSetEquipmentRef.engine.Item(ref.equipmentSetEquipmentRef.ReferencedElementID)
}

func (_ref playerTargetRef) IsSet() bool {
	ref := _ref.playerTargetRef.engine.playerTargetRef(_ref.playerTargetRef.ID)
	return ref.playerTargetRef.ID != 0
}

func (_ref playerTargetRef) Unset() {
	ref := _ref.playerTargetRef.engine.playerTargetRef(_ref.playerTargetRef.ID)
	ref.playerTargetRef.engine.deletePlayerTargetRef(ref.playerTargetRef.ID)
	parent := ref.playerTargetRef.engine.Player(ref.playerTargetRef.ParentID).player
	if parent.OperationKind == OperationKindDelete {
		return
	}
	parent.Target = 0
	parent.OperationKind = OperationKindUpdate
	ref.playerTargetRef.engine.Patch.Player[parent.ID] = parent
}

func (_ref playerTargetRef) Get() anyOfPlayer_ZoneItemRef {
	ref := _ref.playerTargetRef.engine.playerTargetRef(_ref.playerTargetRef.ID)
	anyContainer := ref.playerTargetRef.engine.anyOfPlayer_ZoneItem(ref.playerTargetRef.ReferencedElementID)
	return anyOfPlayer_ZoneItemRef{anyOfPlayer_ZoneItem: anyContainer.anyOfPlayer_ZoneItem, anyOfPlayer_ZoneItemWrapper: anyContainer}
}

func (_ref playerTargetedByRef) Get() anyOfPlayer_ZoneItemRef {
	ref := _ref.playerTargetedByRef.engine.playerTargetedByRef(_ref.playerTargetedByRef.ID)
	anyContainer := ref.playerTargetedByRef.engine.anyOfPlayer_ZoneItem(ref.playerTargetedByRef.ReferencedElementID)
	return anyOfPlayer_ZoneItemRef{anyOfPlayer_ZoneItem: anyContainer.anyOfPlayer_ZoneItem, anyOfPlayer_ZoneItemWrapper: anyContainer}
}

func (engine *Engine) dereferenceItemBoundToRefs(playerID PlayerID) {
	allItemBoundToRefIDs := engine.allItemBoundToRefIDs()
	for _, refID := range allItemBoundToRefIDs {
		ref := engine.itemBoundToRef(refID)
		if ref.itemBoundToRef.ReferencedElementID == playerID {
			ref.Unset()
		}
	}
	itemBoundToRefIDSlicePool.Put(allItemBoundToRefIDs)
}

func (engine *Engine) dereferenceEquipmentSetEquipmentRefs(itemID ItemID) {
	allEquipmentSetEquipmentRefIDs := engine.allEquipmentSetEquipmentRefIDs()
	for _, refID := range allEquipmentSetEquipmentRefIDs {
		ref := engine.equipmentSetEquipmentRef(refID)
		if ref.equipmentSetEquipmentRef.ReferencedElementID == itemID {
			parent := engine.EquipmentSet(ref.equipmentSetEquipmentRef.ParentID)
			parent.RemoveEquipment(itemID)
		}
	}
	equipmentSetEquipmentRefIDSlicePool.Put(allEquipmentSetEquipmentRefIDs)
}

func (engine *Engine) dereferencePlayerGuildMemberRefs(playerID PlayerID) {
	allPlayerGuildMemberRefIDs := engine.allPlayerGuildMemberRefIDs()
	for _, refID := range allPlayerGuildMemberRefIDs {
		ref := engine.playerGuildMemberRef(refID)
		if ref.playerGuildMemberRef.ReferencedElementID == playerID {
			parent := engine.Player(ref.playerGuildMemberRef.ParentID)
			parent.RemoveGuildMembers(playerID)
		}
	}
	playerGuildMemberRefIDSlicePool.Put(allPlayerGuildMemberRefIDs)
}

func (engine *Engine) dereferencePlayerEquipmentSetRefs(equipmentSetID EquipmentSetID) {
	allPlayerEquipmentSetRefIDs := engine.allPlayerEquipmentSetRefIDs()
	for _, refID := range allPlayerEquipmentSetRefIDs {
		ref := engine.playerEquipmentSetRef(refID)
		if ref.playerEquipmentSetRef.ReferencedElementID == equipmentSetID {
			parent := engine.Player(ref.playerEquipmentSetRef.ParentID)
			parent.RemoveEquipmentSets(equipmentSetID)
		}
	}
	playerEquipmentSetRefIDSlicePool.Put(allPlayerEquipmentSetRefIDs)
}

func (engine *Engine) dereferencePlayerTargetRefsPlayer(playerID PlayerID) {
	allPlayerTargetRefIDs := engine.allPlayerTargetRefIDs()
	for _, refID := range allPlayerTargetRefIDs {
		ref := engine.playerTargetRef(refID)
		anyContainer := ref.Get()
		if anyContainer.anyOfPlayer_ZoneItem.ElementKind != ElementKindPlayer {
			continue
		}
		if anyContainer.anyOfPlayer_ZoneItem.Player == playerID {
			ref.Unset()
		}
	}
	playerTargetRefIDSlicePool.Put(allPlayerTargetRefIDs)
}

func (engine *Engine) dereferencePlayerTargetRefsZoneItem(zoneItemID ZoneItemID) {
	allPlayerTargetRefIDs := engine.allPlayerTargetRefIDs()
	for _, refID := range allPlayerTargetRefIDs {
		ref := engine.playerTargetRef(refID)
		anyContainer := ref.Get()
		if anyContainer.anyOfPlayer_ZoneItem.ElementKind != ElementKindZoneItem {
			continue
		}
		if anyContainer.anyOfPlayer_ZoneItem.ZoneItem == zoneItemID {
			ref.Unset()
		}
	}
	playerTargetRefIDSlicePool.Put(allPlayerTargetRefIDs)
}

func (engine *Engine) dereferencePlayerTargetedByRefsPlayer(playerID PlayerID) {
	allPlayerTargetedByRefIDs := engine.allPlayerTargetedByRefIDs()
	for _, refID := range allPlayerTargetedByRefIDs {
		ref := engine.playerTargetedByRef(refID)
		anyContainer := ref.Get()
		if anyContainer.anyOfPlayer_ZoneItem.ElementKind != ElementKindPlayer {
			continue
		}
		if anyContainer.anyOfPlayer_ZoneItem.Player == playerID {
			parent := engine.Player(ref.playerTargetedByRef.ParentID)
			parent.RemoveTargetedByPlayer(playerID)
		}
	}
	playerTargetedByRefIDSlicePool.Put(allPlayerTargetedByRefIDs)
}

func (engine *Engine) dereferencePlayerTargetedByRefsZoneItem(zoneItemID ZoneItemID) {
	allPlayerTargetedByRefIDs := engine.allPlayerTargetedByRefIDs()
	for _, refID := range allPlayerTargetedByRefIDs {
		ref := engine.playerTargetedByRef(refID)
		anyContainer := ref.Get()
		if anyContainer.anyOfPlayer_ZoneItem.ElementKind != ElementKindZoneItem {
			continue
		}
		if anyContainer.anyOfPlayer_ZoneItem.ZoneItem == zoneItemID {
			parent := engine.Player(ref.playerTargetedByRef.ParentID)
			parent.RemoveTargetedByZoneItem(zoneItemID)
		}
	}
	playerTargetedByRefIDSlicePool.Put(allPlayerTargetedByRefIDs)
}
