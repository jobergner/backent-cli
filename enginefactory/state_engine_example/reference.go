package state

func (_ref itemBoundToRef) IsSet(se *Engine) bool {
	ref := se.itemBoundToRef(_ref.itemBoundToRef.ID)
	return ref.itemBoundToRef.ID != 0
}

func (_ref itemBoundToRef) Unset(se *Engine) {
	ref := se.itemBoundToRef(_ref.itemBoundToRef.ID)
	se.deleteItemBoundToRef(ref.itemBoundToRef.ID)
	item := se.Item(ref.itemBoundToRef.ParentID).item
	if item.OperationKind_ == OperationKindDelete {
		return
	}
	item.BoundTo = 0
	item.OperationKind_ = OperationKindUpdate
	se.Patch.Item[item.ID] = item
}

func (_ref itemBoundToRef) Get(se *Engine) player {
	ref := se.itemBoundToRef(_ref.itemBoundToRef.ID)
	return se.Player(ref.itemBoundToRef.ReferencedElementID)
}

func (_ref playerGuildMemberRef) Get(se *Engine) player {
	ref := se.playerGuildMemberRef(_ref.playerGuildMemberRef.ID)
	return se.Player(ref.playerGuildMemberRef.ReferencedElementID)
}

func (_ref equipmentSetEquipmentRef) Get(se *Engine) item {
	ref := se.equipmentSetEquipmentRef(_ref.equipmentSetEquipmentRef.ID)
	return se.Item(ref.equipmentSetEquipmentRef.ReferencedElementID)
}
