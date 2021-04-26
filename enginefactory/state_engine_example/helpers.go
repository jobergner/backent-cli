package state

func deduplicateItemBoundToRefIDs(a []ItemBoundToRefID, b []ItemBoundToRefID) []ItemBoundToRefID {

	check := make(map[ItemBoundToRefID]bool)
	deduped := make([]ItemBoundToRefID, 0)
	for _, val := range a {
		check[val] = true
	}
	for _, val := range b {
		check[val] = true
	}

	for val := range check {
		deduped = append(deduped, val)
	}

	return deduped
}

func deduplicatePlayerGuildMemberRefIDs(a []PlayerGuildMemberRefID, b []PlayerGuildMemberRefID) []PlayerGuildMemberRefID {

	check := make(map[PlayerGuildMemberRefID]bool)
	deduped := make([]PlayerGuildMemberRefID, 0)
	for _, val := range a {
		check[val] = true
	}
	for _, val := range b {
		check[val] = true
	}

	for val := range check {
		deduped = append(deduped, val)
	}

	return deduped
}

func deduplicatePlayerEquipmentSetRefIDs(a []PlayerEquipmentSetRefID, b []PlayerEquipmentSetRefID) []PlayerEquipmentSetRefID {

	check := make(map[PlayerEquipmentSetRefID]bool)
	deduped := make([]PlayerEquipmentSetRefID, 0)
	for _, val := range a {
		check[val] = true
	}
	for _, val := range b {
		check[val] = true
	}

	for val := range check {
		deduped = append(deduped, val)
	}

	return deduped
}

func deduplicateEquipmentSetEquipmentRefIDs(a []EquipmentSetEquipmentRefID, b []EquipmentSetEquipmentRefID) []EquipmentSetEquipmentRefID {

	check := make(map[EquipmentSetEquipmentRefID]bool)
	deduped := make([]EquipmentSetEquipmentRefID, 0)
	for _, val := range a {
		check[val] = true
	}
	for _, val := range b {
		check[val] = true
	}

	for val := range check {
		deduped = append(deduped, val)
	}

	return deduped
}

func (se Engine) allItemBoundToRefIDs() []ItemBoundToRefID {
	var stateItemBoundToRefIDs []ItemBoundToRefID
	for itemBoundToRefID := range se.State.ItemBoundToRef {
		stateItemBoundToRefIDs = append(stateItemBoundToRefIDs, itemBoundToRefID)
	}
	var patchItemBoundToRefIDs []ItemBoundToRefID
	for itemBoundToRefID := range se.Patch.ItemBoundToRef {
		patchItemBoundToRefIDs = append(patchItemBoundToRefIDs, itemBoundToRefID)
	}
	return deduplicateItemBoundToRefIDs(stateItemBoundToRefIDs, patchItemBoundToRefIDs)
}

func (se Engine) allPlayerGuildMemberRefIDs() []PlayerGuildMemberRefID {
	var statePlayerGuildMemberRefIDs []PlayerGuildMemberRefID
	for playerGuildMemberRefID := range se.State.PlayerGuildMemberRef {
		statePlayerGuildMemberRefIDs = append(statePlayerGuildMemberRefIDs, playerGuildMemberRefID)
	}
	var patchPlayerGuildMemberRefIDs []PlayerGuildMemberRefID
	for playerGuildMemberRefID := range se.Patch.PlayerGuildMemberRef {
		patchPlayerGuildMemberRefIDs = append(patchPlayerGuildMemberRefIDs, playerGuildMemberRefID)
	}
	return deduplicatePlayerGuildMemberRefIDs(statePlayerGuildMemberRefIDs, patchPlayerGuildMemberRefIDs)
}

func (se Engine) allPlayerEquipmentSetRefIDs() []PlayerEquipmentSetRefID {
	var statePlayerEquipmentSetRefIDs []PlayerEquipmentSetRefID
	for playerEquipmentSetRefID := range se.State.PlayerEquipmentSetRef {
		statePlayerEquipmentSetRefIDs = append(statePlayerEquipmentSetRefIDs, playerEquipmentSetRefID)
	}
	var patchPlayerEquipmentSetRefIDs []PlayerEquipmentSetRefID
	for playerEquipmentSetRefID := range se.Patch.PlayerEquipmentSetRef {
		patchPlayerEquipmentSetRefIDs = append(patchPlayerEquipmentSetRefIDs, playerEquipmentSetRefID)
	}
	return deduplicatePlayerEquipmentSetRefIDs(statePlayerEquipmentSetRefIDs, patchPlayerEquipmentSetRefIDs)
}

func (se Engine) allEquipmentSetEquipmentRefIDs() []EquipmentSetEquipmentRefID {
	var stateEquipmentSetEquipmentRefIDs []EquipmentSetEquipmentRefID
	for equipmentSetEquipmentRefID := range se.State.EquipmentSetEquipmentRef {
		stateEquipmentSetEquipmentRefIDs = append(stateEquipmentSetEquipmentRefIDs, equipmentSetEquipmentRefID)
	}
	var patchEquipmentSetEquipmentRefIDs []EquipmentSetEquipmentRefID
	for equipmentSetEquipmentRefID := range se.Patch.EquipmentSetEquipmentRef {
		patchEquipmentSetEquipmentRefIDs = append(patchEquipmentSetEquipmentRefIDs, equipmentSetEquipmentRefID)
	}
	return deduplicateEquipmentSetEquipmentRefIDs(stateEquipmentSetEquipmentRefIDs, patchEquipmentSetEquipmentRefIDs)
}

func mergeGearScoreIDs(currentIDs, nextIDs []GearScoreID) []GearScoreID {
	ids := make([]GearScoreID, len(currentIDs))
	copy(ids, currentIDs)
	var j int

	for _, currentID := range currentIDs {
		if len(nextIDs) <= j || currentID != nextIDs[j] {
			continue
		}
		j += 1
	}

	for _, nextID := range nextIDs[j:] {
		ids = append(ids, nextID)
	}

	return ids
}

func mergeItemIDs(currentIDs, nextIDs []ItemID) []ItemID {
	ids := make([]ItemID, len(currentIDs))
	copy(ids, currentIDs)
	var j int

	for _, currentID := range currentIDs {
		if len(nextIDs) <= j || currentID != nextIDs[j] {
			continue
		}
		j += 1
	}

	for _, nextID := range nextIDs[j:] {
		ids = append(ids, nextID)
	}

	return ids
}

func mergePlayerIDs(currentIDs, nextIDs []PlayerID) []PlayerID {
	ids := make([]PlayerID, len(currentIDs))
	copy(ids, currentIDs)
	var j int

	for _, currentID := range currentIDs {
		if len(nextIDs) <= j || currentID != nextIDs[j] {
			continue
		}
		j += 1
	}

	for _, nextID := range nextIDs[j:] {
		ids = append(ids, nextID)
	}

	return ids
}

func mergePositionIDs(currentIDs, nextIDs []PositionID) []PositionID {
	ids := make([]PositionID, len(currentIDs))
	copy(ids, currentIDs)
	var j int

	for _, currentID := range currentIDs {
		if len(nextIDs) <= j || currentID != nextIDs[j] {
			continue
		}
		j += 1
	}

	for _, nextID := range nextIDs[j:] {
		ids = append(ids, nextID)
	}

	return ids
}

func mergeZoneIDs(currentIDs, nextIDs []ZoneID) []ZoneID {
	ids := make([]ZoneID, len(currentIDs))
	copy(ids, currentIDs)
	var j int

	for _, currentID := range currentIDs {
		if len(nextIDs) <= j || currentID != nextIDs[j] {
			continue
		}
		j += 1
	}

	for _, nextID := range nextIDs[j:] {
		ids = append(ids, nextID)
	}

	return ids
}

func mergeZoneItemIDs(currentIDs, nextIDs []ZoneItemID) []ZoneItemID {
	ids := make([]ZoneItemID, len(currentIDs))
	copy(ids, currentIDs)
	var j int

	for _, currentID := range currentIDs {
		if len(nextIDs) <= j || currentID != nextIDs[j] {
			continue
		}
		j += 1
	}

	for _, nextID := range nextIDs[j:] {
		ids = append(ids, nextID)
	}

	return ids
}

func mergeEquipmentSetIDs(currentIDs, nextIDs []EquipmentSetID) []EquipmentSetID {
	ids := make([]EquipmentSetID, len(currentIDs))
	copy(ids, currentIDs)
	var j int

	for _, currentID := range currentIDs {
		if len(nextIDs) <= j || currentID != nextIDs[j] {
			continue
		}
		j += 1
	}

	for _, nextID := range nextIDs[j:] {
		ids = append(ids, nextID)
	}

	return ids
}

func mergeItemBoundToRefIDs(currentIDs, nextIDs []ItemBoundToRefID) []ItemBoundToRefID {
	ids := make([]ItemBoundToRefID, len(currentIDs))
	copy(ids, currentIDs)
	var j int

	for _, currentID := range currentIDs {
		if len(nextIDs) <= j || currentID != nextIDs[j] {
			continue
		}
		j += 1
	}

	for _, nextID := range nextIDs[j:] {
		ids = append(ids, nextID)
	}

	return ids
}

func mergeEquipmentSetEquipmentRefIDs(currentIDs, nextIDs []EquipmentSetEquipmentRefID) []EquipmentSetEquipmentRefID {
	ids := make([]EquipmentSetEquipmentRefID, len(currentIDs))
	copy(ids, currentIDs)
	var j int

	for _, currentID := range currentIDs {
		if len(nextIDs) <= j || currentID != nextIDs[j] {
			continue
		}
		j += 1
	}

	for _, nextID := range nextIDs[j:] {
		ids = append(ids, nextID)
	}

	return ids
}

func mergePlayerGuildMemberRefIDs(currentIDs, nextIDs []PlayerGuildMemberRefID) []PlayerGuildMemberRefID {
	ids := make([]PlayerGuildMemberRefID, len(currentIDs))
	copy(ids, currentIDs)
	var j int

	for _, currentID := range currentIDs {
		if len(nextIDs) <= j || currentID != nextIDs[j] {
			continue
		}
		j += 1
	}

	for _, nextID := range nextIDs[j:] {
		ids = append(ids, nextID)
	}

	return ids
}

func mergePlayerEquipmentSetRefIDs(currentIDs, nextIDs []PlayerEquipmentSetRefID) []PlayerEquipmentSetRefID {
	ids := make([]PlayerEquipmentSetRefID, len(currentIDs))
	copy(ids, currentIDs)
	var j int

	for _, currentID := range currentIDs {
		if len(nextIDs) <= j || currentID != nextIDs[j] {
			continue
		}
		j += 1
	}

	for _, nextID := range nextIDs[j:] {
		ids = append(ids, nextID)
	}

	return ids
}

func (se *Engine) dereferenceItemBoundToRefs(playerID PlayerID) {
	for _, refID := range se.allItemBoundToRefIDs() {
		ref := se.itemBoundToRef(refID)
		if ref.itemBoundToRef.ReferencedElementID == playerID {
			ref.Unset(se)
		}
	}
}

func (se *Engine) dereferencePlayerGuildMemberRefs(playerID PlayerID) {
	for _, refID := range se.allPlayerGuildMemberRefIDs() {
		ref := se.playerGuildMemberRef(refID)
		if ref.playerGuildMemberRef.ReferencedElementID == playerID {
			parent := se.Player(ref.playerGuildMemberRef.ParentID)
			parent.RemoveGuildMembers(se, playerID)
		}
	}
}

func (se *Engine) dereferencePlayerEquipmentSetRefs(equipmentSetID EquipmentSetID) {
	for _, refID := range se.allPlayerEquipmentSetRefIDs() {
		ref := se.playerEquipmentSetRef(refID)
		if ref.playerEquipmentSetRef.ReferencedElementID == equipmentSetID {
			parent := se.Player(ref.playerEquipmentSetRef.ParentID)
			parent.RemoveEquipmentSets(se, equipmentSetID)
		}
	}
}

func (se *Engine) dereferenceEquipmentSetEquipmentRef(itemID ItemID) {
	for _, refID := range se.allEquipmentSetEquipmentRefIDs() {
		ref := se.equipmentSetEquipmentRef(refID)
		if ref.equipmentSetEquipmentRef.ReferencedElementID == itemID {
			parent := se.EquipmentSet(ref.equipmentSetEquipmentRef.ParentID)
			parent.RemoveEquipment(se, itemID)
		}
	}
}

func (se *Engine) itemBoundToRefToElementRef(itemID ItemID) (*ElementReference, bool) {
	stateItem := se.State.Item[itemID]
	patchItem, itemIsInPatch := se.Patch.Item[itemID]

	// ref not set at all
	if stateItem.BoundTo == 0 && (!itemIsInPatch || patchItem.BoundTo == 0) {
		return nil, false
	}

	// immediate replacement of refs
	if stateItem.BoundTo != 0 && (itemIsInPatch && patchItem.BoundTo != 0) {
		if stateItem.BoundTo != patchItem.BoundTo {
			referencedElement := se.Player(se.itemBoundToRef(patchItem.BoundTo).itemBoundToRef.ReferencedElementID).player
			return &ElementReference{OperationKindUpdate, int(referencedElement.ID), ElementKindPlayer, referencedElement.OperationKind_}, true
		}
	}

	// ref was definitely removed
	if stateItem.BoundTo != 0 && (itemIsInPatch && patchItem.BoundTo == 0) {
		referencedElement := se.Player(se.itemBoundToRef(stateItem.BoundTo).itemBoundToRef.ReferencedElementID).player
		return &ElementReference{OperationKindDelete, int(referencedElement.ID), ElementKindPlayer, referencedElement.OperationKind_}, true
	}

	// ref was definitely created
	if stateItem.BoundTo == 0 && (itemIsInPatch && patchItem.BoundTo != 0) {
		referencedElement := se.Player(se.itemBoundToRef(patchItem.BoundTo).itemBoundToRef.ReferencedElementID).player
		return &ElementReference{OperationKindUpdate, int(referencedElement.ID), ElementKindPlayer, referencedElement.OperationKind_}, true
	}

	// referenced element got updated
	if stateItem.BoundTo != 0 {
		if referencedElement, ok := se.Patch.Player[se.itemBoundToRef(stateItem.BoundTo).itemBoundToRef.ReferencedElementID]; ok {
			return &ElementReference{OperationKindUnchanged, int(referencedElement.ID), ElementKindPlayer, referencedElement.OperationKind_}, true
		}
	}

	return nil, false
}

func (se *Engine) playerGuildMemberRefsToElementRefs(playerID PlayerID) ([]ElementReference, bool) {
	var anyHaveUpdated bool

	statePlayer := se.State.Player[playerID]
	patchPlayer := se.Patch.Player[playerID]

	var refs []ElementReference

	for _, refID := range mergePlayerGuildMemberRefIDs(statePlayer.GuildMembers, patchPlayer.GuildMembers) {

		if patchRef, hasUpdated := se.Patch.PlayerGuildMemberRef[refID]; hasUpdated {
			referencedElement := se.Player(patchRef.ReferencedElementID).player
			refs = append(refs, ElementReference{patchRef.OperationKind_, int(patchRef.ReferencedElementID), ElementKindPlayer, referencedElement.OperationKind_})
			anyHaveUpdated = true
			continue
		}

		ref := se.playerGuildMemberRef(refID).playerGuildMemberRef
		if referencedElement, hasUpdated := se.Patch.Player[ref.ReferencedElementID]; hasUpdated {
			refs = append(refs, ElementReference{OperationKindUpdate, int(ref.ReferencedElementID), ElementKindPlayer, referencedElement.OperationKind_})
			anyHaveUpdated = true
		}
	}

	return refs, anyHaveUpdated
}

func (se *Engine) playerEquipmentSetRefsToElementRefs(playerID PlayerID) ([]ElementReference, bool) {
	var anyHaveUpdated bool

	statePlayer := se.State.Player[playerID]
	patchPlayer := se.Patch.Player[playerID]

	var refs []ElementReference

	for _, refID := range mergePlayerEquipmentSetRefIDs(statePlayer.EquipmentSets, patchPlayer.EquipmentSets) {

		if patchRef, hasUpdated := se.Patch.PlayerEquipmentSetRef[refID]; hasUpdated {
			referencedElement := se.EquipmentSet(patchRef.ReferencedElementID).equipmentSet
			refs = append(refs, ElementReference{patchRef.OperationKind_, int(patchRef.ReferencedElementID), ElementKindPlayer, referencedElement.OperationKind_})
			anyHaveUpdated = true
			continue
		}

		ref := se.playerEquipmentSetRef(refID).playerEquipmentSetRef
		if referencedElement, hasUpdated := se.Patch.EquipmentSet[ref.ReferencedElementID]; hasUpdated {
			refs = append(refs, ElementReference{OperationKindUpdate, int(ref.ReferencedElementID), ElementKindPlayer, referencedElement.OperationKind_})
			anyHaveUpdated = true
		}
	}

	return refs, anyHaveUpdated
}

func (se *Engine) equipmentSetEquipmentRefsToElementRefs(equipmentSetID EquipmentSetID) ([]ElementReference, bool) {
	var anyHaveUpdated bool

	stateEquipmentSet := se.State.EquipmentSet[equipmentSetID]
	patchEquipmentSet := se.Patch.EquipmentSet[equipmentSetID]

	var refs []ElementReference

	for _, refID := range mergeEquipmentSetEquipmentRefIDs(stateEquipmentSet.Equipment, patchEquipmentSet.Equipment) {

		if patchRef, hasUpdated := se.Patch.EquipmentSetEquipmentRef[refID]; hasUpdated {
			referencedElement := se.Item(patchRef.ReferencedElementID).item
			refs = append(refs, ElementReference{patchRef.OperationKind_, int(patchRef.ReferencedElementID), ElementKindEquipmentSet, referencedElement.OperationKind_})
			anyHaveUpdated = true
			continue
		}

		ref := se.equipmentSetEquipmentRef(refID).equipmentSetEquipmentRef
		if referencedElement, hasUpdated := se.Patch.Item[ref.ReferencedElementID]; hasUpdated {
			refs = append(refs, ElementReference{OperationKindUpdate, int(ref.ReferencedElementID), ElementKindEquipmentSet, referencedElement.OperationKind_})
			anyHaveUpdated = true
		}
	}

	return refs, anyHaveUpdated
}
