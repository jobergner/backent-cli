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

func (_ref playerTargetRef) IsSet() bool {
	ref := _ref.playerTargetRef.engine.playerTargetRef(_ref.playerTargetRef.ID)
	return ref.playerTargetRef.ID != 0
}

func (_ref playerTargetRef) Unset() {
	ref := _ref.playerTargetRef.engine.playerTargetRef(_ref.playerTargetRef.ID)
	_ref.playerTargetRef.engine.deletePlayerTargetRef(ref.playerTargetRef.ID)
	player := _ref.playerTargetRef.engine.Player(ref.playerTargetRef.ParentID).player
	if player.OperationKind_ == OperationKindDelete {
		return
	}
	player.Target = 0
	player.OperationKind_ = OperationKindUpdate
	_ref.playerTargetRef.engine.Patch.Player[player.ID] = player
}

func (_ref playerTargetRef) Get() anyOfPlayerZoneItem {
	ref := _ref.playerTargetRef.engine.playerTargetRef(_ref.playerTargetRef.ID)
	return _ref.playerTargetRef.engine.anyOfPlayerZoneItem(ref.playerTargetRef.ReferencedElementID)
}

func (_ref playerTargetedByRef) Get() anyOfPlayerZoneItem {
	ref := _ref.playerTargetedByRef.engine.playerTargetedByRef(_ref.playerTargetedByRef.ID)
	return _ref.playerTargetedByRef.engine.anyOfPlayerZoneItem(ref.playerTargetedByRef.ReferencedElementID)
}

func (engine *Engine) dereferenceItemBoundToRefs(playerID PlayerID) {
	for _, refID := range engine.allItemBoundToRefIDs() {
		ref := engine.itemBoundToRef(refID)
		if ref.itemBoundToRef.ReferencedElementID == playerID {
			ref.Unset()
		}
	}
}

func (engine *Engine) dereferencePlayerGuildMemberRefs(playerID PlayerID) {
	for _, refID := range engine.allPlayerGuildMemberRefIDs() {
		ref := engine.playerGuildMemberRef(refID)
		if ref.playerGuildMemberRef.ReferencedElementID == playerID {
			parent := engine.Player(ref.playerGuildMemberRef.ParentID)
			parent.RemoveGuildMembers(playerID)
		}
	}
}

func (engine *Engine) dereferencePlayerEquipmentSetRefs(equipmentSetID EquipmentSetID) {
	for _, refID := range engine.allPlayerEquipmentSetRefIDs() {
		ref := engine.playerEquipmentSetRef(refID)
		if ref.playerEquipmentSetRef.ReferencedElementID == equipmentSetID {
			parent := engine.Player(ref.playerEquipmentSetRef.ParentID)
			parent.RemoveEquipmentSets(equipmentSetID)
		}
	}
}

func (engine *Engine) dereferenceEquipmentSetEquipmentRef(itemID ItemID) {
	for _, refID := range engine.allEquipmentSetEquipmentRefIDs() {
		ref := engine.equipmentSetEquipmentRef(refID)
		if ref.equipmentSetEquipmentRef.ReferencedElementID == itemID {
			parent := engine.EquipmentSet(ref.equipmentSetEquipmentRef.ParentID)
			parent.RemoveEquipment(itemID)
		}
	}
}

func (engine *Engine) dereferencePlayerTargetPlayerRefs(playerID PlayerID) {
	for _, refID := range engine.allPlayerTargetRefIDs() {
		ref := engine.playerTargetRef(refID)
		anyContainer := ref.Get()
		if anyContainer.anyOfPlayerZoneItem.ElementKind != ElementKindPlayer {
			return
		}
		if anyContainer.anyOfPlayerZoneItem.Player == playerID {
			ref.Unset()
		}
	}
}

func (engine *Engine) dereferencePlayerTargetZoneItemRefs(zoneItemID ZoneItemID) {
	for _, refID := range engine.allPlayerTargetRefIDs() {
		ref := engine.playerTargetRef(refID)
		anyContainer := ref.Get()
		if anyContainer.anyOfPlayerZoneItem.ElementKind != ElementKindZoneItem {
			return
		}
		if anyContainer.anyOfPlayerZoneItem.ZoneItem == zoneItemID {
			ref.Unset()
		}
	}
}

func (engine *Engine) dereferencePlayerTargetedByPlayerRefs(playerID PlayerID) {
	for _, refID := range engine.allPlayerTargetedByRefIDs() {
		ref := engine.playerTargetedByRef(refID)
		anyContainer := ref.Get()
		if anyContainer.anyOfPlayerZoneItem.ElementKind != ElementKindPlayer {
			continue
		}
		if anyContainer.anyOfPlayerZoneItem.Player == playerID {
			parent := engine.Player(ref.playerTargetedByRef.ParentID)
			parent.RemoveTargetedByPlayer(playerID)
		}
	}
}

func (engine *Engine) dereferencePlayerTargetedByZoneItemRefs(zoneItemID ZoneItemID) {
	for _, refID := range engine.allPlayerTargetedByRefIDs() {
		ref := engine.playerTargetedByRef(refID)
		anyContainer := ref.Get()
		if anyContainer.anyOfPlayerZoneItem.ElementKind != ElementKindZoneItem {
			continue
		}
		if anyContainer.anyOfPlayerZoneItem.ZoneItem == zoneItemID {
			parent := engine.Player(ref.playerTargetedByRef.ParentID)
			parent.RemoveTargetedByZoneItem(zoneItemID)
		}
	}
}
