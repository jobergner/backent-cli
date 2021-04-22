package state

type updatedRefsMap map[int]bool

func (se *Engine) updateItemBoundToRef(ref itemBoundToRef) {
	ref.itemBoundToRef.OperationKind_ = OperationKindUpdate
	se.Patch.ItemBoundToRef[ref.itemBoundToRef.ID] = ref.itemBoundToRef
}
func (se *Engine) updatePlayerGuildMemberRef(ref playerGuildMemberRef) {
	ref.playerGuildMemberRef.OperationKind_ = OperationKindUpdate
	se.Patch.PlayerGuildMemberRef[ref.playerGuildMemberRef.ID] = ref.playerGuildMemberRef
	se.ReferenceWatch[int(ref.playerGuildMemberRef.ID)] = true
	for _, refID := range se.allItemBoundToRefIDs() {
		_ref := se.itemBoundToRef(refID)
		if _ref.itemBoundToRef.ReferencedElementID == ref.playerGuildMemberRef.ParentID {
			if _, ok := se.ReferenceWatch[int(_ref.itemBoundToRef.ID)]; !ok {
				se.updateItemBoundToRef(_ref)
			}
		}
	}
	for _, refID := range se.allPlayerGuildMemberRefIDs() {
		_ref := se.playerGuildMemberRef(refID)
		if _ref.playerGuildMemberRef.ReferencedElementID == ref.playerGuildMemberRef.ParentID {
			if _, ok := se.ReferenceWatch[int(_ref.playerGuildMemberRef.ID)]; !ok {
				se.updatePlayerGuildMemberRef(_ref)
			}
		}
	}
}
