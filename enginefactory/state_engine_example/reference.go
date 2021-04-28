package state

func (_ref itemBoundToRef) IsSet() bool {
	ref := _ref.itemBoundToRef.engine.itemBoundToRef(_ref.itemBoundToRef.ID)
	return ref.itemBoundToRef.ID != 0
}

func (_ref itemBoundToRef) Unset() {
	ref := _ref.itemBoundToRef.engine.itemBoundToRef(_ref.itemBoundToRef.ID)
	_ref.itemBoundToRef.engine.deleteItemBoundToRef(ref.itemBoundToRef.ID)
	item := _ref.itemBoundToRef.engine.Item(ref.itemBoundToRef.ParentID).item
	if item.OperationKind_ == OperationKindDelete {
		return
	}
	item.BoundTo = 0
	item.OperationKind_ = OperationKindUpdate
	_ref.itemBoundToRef.engine.Patch.Item[item.ID] = item
}

func (_ref itemBoundToRef) Get() player {
	ref := _ref.itemBoundToRef.engine.itemBoundToRef(_ref.itemBoundToRef.ID)
	return _ref.itemBoundToRef.engine.Player(ref.itemBoundToRef.ReferencedElementID)
}

func (_ref playerGuildMemberRef) Get() player {
	ref := _ref.playerGuildMemberRef.engine.playerGuildMemberRef(_ref.playerGuildMemberRef.ID)
	return _ref.playerGuildMemberRef.engine.Player(ref.playerGuildMemberRef.ReferencedElementID)
}

func (_ref equipmentSetEquipmentRef) Get() item {
	ref := _ref.equipmentSetEquipmentRef.engine.equipmentSetEquipmentRef(_ref.equipmentSetEquipmentRef.ID)
	return _ref.equipmentSetEquipmentRef.engine.Item(ref.equipmentSetEquipmentRef.ReferencedElementID)
}
