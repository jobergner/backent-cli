package state

func (_ref ItemBoundToRef) IsSet() (ItemBoundToRef, bool) {
	ref := _ref.itemBoundToRef.engine.itemBoundToRef(_ref.itemBoundToRef.ID)
	return ref, ref.itemBoundToRef.ID != 0
}

func (_ref ItemBoundToRef) Unset() {
	ref := _ref.itemBoundToRef.engine.itemBoundToRef(_ref.itemBoundToRef.ID)
	if ref.itemBoundToRef.OperationKind == OperationKindDelete {
		return
	}
	ref.itemBoundToRef.engine.deleteItemBoundToRef(ref.itemBoundToRef.ID)
	parent := ref.itemBoundToRef.engine.Item(ref.itemBoundToRef.ParentID).item
	if parent.OperationKind == OperationKindDelete {
		return
	}
	parent.BoundTo = 0
	parent.OperationKind = OperationKindUpdate
	ref.itemBoundToRef.engine.Patch.Item[parent.ID] = parent
}

func (_ref ItemBoundToRef) Get() Player {
	ref := _ref.itemBoundToRef.engine.itemBoundToRef(_ref.itemBoundToRef.ID)
	return ref.itemBoundToRef.engine.Player(ref.itemBoundToRef.ReferencedElementID)
}

func (_ref AttackEventTargetRef) IsSet() (AttackEventTargetRef, bool) {
	ref := _ref.attackEventTargetRef.engine.attackEventTargetRef(_ref.attackEventTargetRef.ID)
	return ref, ref.attackEventTargetRef.ID != 0
}

func (_ref AttackEventTargetRef) Unset() {
	ref := _ref.attackEventTargetRef.engine.attackEventTargetRef(_ref.attackEventTargetRef.ID)
	if ref.attackEventTargetRef.OperationKind == OperationKindDelete {
		return
	}
	ref.attackEventTargetRef.engine.deleteAttackEventTargetRef(ref.attackEventTargetRef.ID)
	parent := ref.attackEventTargetRef.engine.AttackEvent(ref.attackEventTargetRef.ParentID).attackEvent
	if parent.OperationKind == OperationKindDelete {
		return
	}
	parent.Target = 0
	parent.OperationKind = OperationKindUpdate
	ref.attackEventTargetRef.engine.Patch.AttackEvent[parent.ID] = parent
}

func (_ref AttackEventTargetRef) Get() Player {
	ref := _ref.attackEventTargetRef.engine.attackEventTargetRef(_ref.attackEventTargetRef.ID)
	return ref.attackEventTargetRef.engine.Player(ref.attackEventTargetRef.ReferencedElementID)
}

func (_ref PlayerGuildMemberRef) Get() Player {
	ref := _ref.playerGuildMemberRef.engine.playerGuildMemberRef(_ref.playerGuildMemberRef.ID)
	return ref.playerGuildMemberRef.engine.Player(ref.playerGuildMemberRef.ReferencedElementID)
}

func (_ref PlayerEquipmentSetRef) Get() EquipmentSet {
	ref := _ref.playerEquipmentSetRef.engine.playerEquipmentSetRef(_ref.playerEquipmentSetRef.ID)
	return ref.playerEquipmentSetRef.engine.EquipmentSet(ref.playerEquipmentSetRef.ReferencedElementID)
}

func (_ref EquipmentSetEquipmentRef) Get() Item {
	ref := _ref.equipmentSetEquipmentRef.engine.equipmentSetEquipmentRef(_ref.equipmentSetEquipmentRef.ID)
	return ref.equipmentSetEquipmentRef.engine.Item(ref.equipmentSetEquipmentRef.ReferencedElementID)
}

func (_ref PlayerTargetRef) IsSet() (PlayerTargetRef, bool) {
	ref := _ref.playerTargetRef.engine.playerTargetRef(_ref.playerTargetRef.ID)
	return ref, ref.playerTargetRef.ID != 0
}

func (_ref PlayerTargetRef) Unset() {
	ref := _ref.playerTargetRef.engine.playerTargetRef(_ref.playerTargetRef.ID)
	if ref.playerTargetRef.OperationKind == OperationKindDelete {
		return
	}
	ref.playerTargetRef.engine.deletePlayerTargetRef(ref.playerTargetRef.ID)
	parent := ref.playerTargetRef.engine.Player(ref.playerTargetRef.ParentID).player
	if parent.OperationKind == OperationKindDelete {
		return
	}
	parent.Target = 0
	parent.OperationKind = OperationKindUpdate
	ref.playerTargetRef.engine.Patch.Player[parent.ID] = parent
}

func (_ref PlayerTargetRef) Get() AnyOfPlayer_ZoneItemRef {
	ref := _ref.playerTargetRef.engine.playerTargetRef(_ref.playerTargetRef.ID)
	anyContainer := ref.playerTargetRef.engine.anyOfPlayer_ZoneItem(ref.playerTargetRef.ReferencedElementID)
	return anyContainer
}

func (_ref PlayerTargetedByRef) Get() AnyOfPlayer_ZoneItemRef {
	ref := _ref.playerTargetedByRef.engine.playerTargetedByRef(_ref.playerTargetedByRef.ID)
	anyContainer := ref.playerTargetedByRef.engine.anyOfPlayer_ZoneItem(ref.playerTargetedByRef.ReferencedElementID)
	return anyContainer
}

// dereferencing is the response to an element, of which other elements held references to, being deleted
// we do this so elements don't hold references of elements which no longer exist

// this will never be used as events can't be referenced
// however we'll keep it here just to have more consistency
func (engine *Engine) dereferenceAttackEventTargetRefs(playerID PlayerID) {
	allAttackEventTargetRefIDs := engine.allAttackEventTargetRefIDs()
	for _, refID := range allAttackEventTargetRefIDs {
		ref := engine.attackEventTargetRef(refID)
		if ref.attackEventTargetRef.ReferencedElementID == playerID {
			ref.Unset()
		}
	}
	attackEventTargetRefIDSlicePool.Put(allAttackEventTargetRefIDs)
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
			parent.RemoveGuildMember(playerID)
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
			parent.RemoveEquipmentSet(equipmentSetID)
		}
	}
	playerEquipmentSetRefIDSlicePool.Put(allPlayerEquipmentSetRefIDs)
}

func (engine *Engine) dereferencePlayerTargetRefsPlayer(playerID PlayerID) {
	allPlayerTargetRefIDs := engine.allPlayerTargetRefIDs()
	for _, refID := range allPlayerTargetRefIDs {
		ref := engine.playerTargetRef(refID)
		anyContainer := ref.playerTargetRef.engine.anyOfPlayer_ZoneItem(ref.playerTargetRef.ReferencedElementID)
		if anyContainer.anyOfPlayer_ZoneItem.ElementKind != ElementKindPlayer {
			continue
		}
		if PlayerID(anyContainer.anyOfPlayer_ZoneItem.ChildID) == playerID {
			ref.Unset()
		}
	}
	playerTargetRefIDSlicePool.Put(allPlayerTargetRefIDs)
}

func (engine *Engine) dereferencePlayerTargetRefsZoneItem(zoneItemID ZoneItemID) {
	allPlayerTargetRefIDs := engine.allPlayerTargetRefIDs()
	for _, refID := range allPlayerTargetRefIDs {
		ref := engine.playerTargetRef(refID)
		anyContainer := ref.playerTargetRef.engine.anyOfPlayer_ZoneItem(ref.playerTargetRef.ReferencedElementID)
		if anyContainer.anyOfPlayer_ZoneItem.ElementKind != ElementKindZoneItem {
			continue
		}
		if ZoneItemID(anyContainer.anyOfPlayer_ZoneItem.ChildID) == zoneItemID {
			ref.Unset()
		}
	}
	playerTargetRefIDSlicePool.Put(allPlayerTargetRefIDs)
}

func (engine *Engine) dereferencePlayerTargetedByRefsPlayer(playerID PlayerID) {
	allPlayerTargetedByRefIDs := engine.allPlayerTargetedByRefIDs()
	for _, refID := range allPlayerTargetedByRefIDs {
		ref := engine.playerTargetedByRef(refID)
		anyContainer := ref.playerTargetedByRef.engine.anyOfPlayer_ZoneItem(ref.playerTargetedByRef.ReferencedElementID)
		if anyContainer.anyOfPlayer_ZoneItem.ElementKind != ElementKindPlayer {
			continue
		}
		if PlayerID(anyContainer.anyOfPlayer_ZoneItem.ChildID) == playerID {
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
		anyContainer := ref.playerTargetedByRef.engine.anyOfPlayer_ZoneItem(ref.playerTargetedByRef.ReferencedElementID)
		if anyContainer.anyOfPlayer_ZoneItem.ElementKind != ElementKindZoneItem {
			continue
		}
		if ZoneItemID(anyContainer.anyOfPlayer_ZoneItem.ChildID) == zoneItemID {
			parent := engine.Player(ref.playerTargetedByRef.ParentID)
			parent.RemoveTargetedByZoneItem(zoneItemID)
		}
	}
	playerTargetedByRefIDSlicePool.Put(allPlayerTargetedByRefIDs)
}
