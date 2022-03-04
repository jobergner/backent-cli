package state

func (engine *Engine) assembleGearScorePath(element *gearScore, p path, pIndex int, includedElements map[int]bool) {

	gearScoreData, ok := engine.Patch.GearScore[element.ID]
	if !ok {
		gearScoreData = engine.State.GearScore[element.ID]
	}

	// to prevent element OperationKindUnchanged overwriting OperKindUpdate
	// set by basic child elements
	if element.OperationKind != OperationKindUpdate && element.OperationKind != OperationKindDelete {
		element.OperationKind = gearScoreData.OperationKind
	}

	if pIndex+1 == len(p) {
		return
	}

	nextSeg := p[pIndex+1]

	switch nextSeg.identifier {
	case gearScore_levelIdentifier:
		child := engine.intValue(gearScoreData.Level)
		element.OperationKind = child.OperationKind
		element.Level = &child.Value
	case gearScore_scoreIdentifier:
		child := engine.intValue(gearScoreData.Score)
		element.OperationKind = child.OperationKind
		element.Score = &child.Value
	}

	_ = gearScoreData
}

func (engine *Engine) assemblePositionPath(element *position, p path, pIndex int, includedElements map[int]bool) {

	positionData, ok := engine.Patch.Position[element.ID]
	if !ok {
		positionData = engine.State.Position[element.ID]
	}

	if element.OperationKind != OperationKindUpdate && element.OperationKind != OperationKindDelete {
		element.OperationKind = positionData.OperationKind
	}

	if pIndex+1 == len(p) {
		return
	}

	nextSeg := p[pIndex+1]

	switch nextSeg.identifier {
	case position_xIdentifier:
		child := engine.floatValue(positionData.X)
		element.OperationKind = child.OperationKind
		element.X = &child.Value
	case position_yIdentifier:
		child := engine.floatValue(positionData.Y)
		element.OperationKind = child.OperationKind
		element.Y = &child.Value
	}

	_ = positionData
}

func (engine *Engine) assembleEquipmentSetPath(element *equipmentSet, p path, pIndex int, includedElements map[int]bool) {

	equipmentSetData, ok := engine.Patch.EquipmentSet[element.ID]
	if !ok {
		equipmentSetData = engine.State.EquipmentSet[element.ID]
	}

	if element.OperationKind != OperationKindUpdate && element.OperationKind != OperationKindDelete {
		element.OperationKind = equipmentSetData.OperationKind
	}

	if pIndex+1 == len(p) {
		return
	}

	nextSeg := p[pIndex+1]

	switch nextSeg.identifier {
	case equipmentSet_equipmentIdentifier:
		ref := engine.equipmentSetEquipmentRef(EquipmentSetEquipmentRefID(nextSeg.refID)).equipmentSetEquipmentRef
		referencedDataStatus := ReferencedDataUnchanged
		if _, ok := includedElements[int(ref.ReferencedElementID)]; ok {
			referencedDataStatus = ReferencedDataModified
		}
		referencedElement := engine.Item(ref.ReferencedElementID).item
		treeRef := elementReference{
			OperationKind:        ref.OperationKind,
			ElementID:            int(ref.ReferencedElementID),
			ElementKind:          ElementKindItem,
			ReferencedDataStatus: referencedDataStatus,
			ElementPath:          referencedElement.Path,
		}
		if element.Equipment == nil {
			element.Equipment = make(map[ItemID]elementReference)
		}
		element.Equipment[referencedElement.ID] = treeRef
	case equipmentSet_nameIdentifier:
		child := engine.stringValue(equipmentSetData.Name)
		element.OperationKind = child.OperationKind
		element.Name = &child.Value
	}

	_ = equipmentSetData
}

func (engine *Engine) assembleAttackEventPath(element *attackEvent, p path, pIndex int, includedElements map[int]bool) {

	attackEventData, ok := engine.Patch.AttackEvent[element.ID]
	if !ok {
		attackEventData = engine.State.AttackEvent[element.ID]
	}

	if element.OperationKind != OperationKindUpdate && element.OperationKind != OperationKindDelete {
		element.OperationKind = attackEventData.OperationKind
	}

	if pIndex+1 == len(p) {
		return
	}

	nextSeg := p[pIndex+1]

	switch nextSeg.identifier {
	case attackEvent_targetIdentifier:
		ref := engine.attackEventTargetRef(AttackEventTargetRefID(nextSeg.refID)).attackEventTargetRef
		if element.Target != nil && ref.OperationKind == OperationKindDelete {
			break
		}
		referencedDataStatus := ReferencedDataUnchanged
		if _, ok := includedElements[int(ref.ReferencedElementID)]; ok {
			referencedDataStatus = ReferencedDataModified
		}
		referencedElement := engine.Player(ref.ReferencedElementID).player
		treeRef := elementReference{
			OperationKind:        ref.OperationKind,
			ElementID:            int(ref.ReferencedElementID),
			ElementKind:          ElementKindPlayer,
			ReferencedDataStatus: referencedDataStatus,
			ElementPath:          referencedElement.Path,
		}
		element.Target = &treeRef
	}

	_ = attackEventData
}

func (engine *Engine) assembleItemPath(element *item, p path, pIndex int, includedElements map[int]bool) {

	itemData, ok := engine.Patch.Item[element.ID]
	if !ok {
		itemData = engine.State.Item[element.ID]
	}

	if element.OperationKind != OperationKindUpdate && element.OperationKind != OperationKindDelete {
		element.OperationKind = itemData.OperationKind
	}

	if pIndex+1 == len(p) {
		return
	}

	nextSeg := p[pIndex+1]

	switch nextSeg.identifier {
	case item_boundToIdentifier:
		ref := engine.itemBoundToRef(ItemBoundToRefID(nextSeg.refID)).itemBoundToRef
		if element.BoundTo != nil && ref.OperationKind == OperationKindDelete {
			break
		}
		referencedDataStatus := ReferencedDataUnchanged
		if _, ok := includedElements[int(ref.ReferencedElementID)]; ok {
			referencedDataStatus = ReferencedDataModified
		}
		referencedElement := engine.Player(ref.ReferencedElementID).player
		treeRef := elementReference{
			OperationKind:        ref.OperationKind,
			ElementID:            int(ref.ReferencedElementID),
			ElementKind:          ElementKindPlayer,
			ReferencedDataStatus: referencedDataStatus,
			ElementPath:          referencedElement.Path,
		}
		element.BoundTo = &treeRef
	case item_gearScoreIdentifier:
		child := element.GearScore
		if child == nil {
			child = &gearScore{ID: GearScoreID(nextSeg.id)}
		}
		engine.assembleGearScorePath(child, p, pIndex+1, includedElements)
		element.GearScore = child
	case item_nameIdenfitier:
		child := engine.stringValue(itemData.Name)
		element.OperationKind = child.OperationKind
		element.Name = &child.Value
	case item_originIdentifier:
		switch nextSeg.kind {
		case ElementKindPlayer:
			child, ok := element.Origin.(*player)
			if !ok || child == nil {
				child = &player{ID: PlayerID(nextSeg.id)}
			}
			engine.assemblePlayerPath(child, p, pIndex+1, includedElements)
			element.Origin = child
		case ElementKindPosition:
			child, ok := element.Origin.(*position)
			if !ok || child == nil {
				child = &position{ID: PositionID(nextSeg.id)}
			}
			engine.assemblePositionPath(child, p, pIndex+1, includedElements)
			element.Origin = child
		}
	}

	_ = itemData
}

func (engine *Engine) assembleZoneItemPath(element *zoneItem, p path, pIndex int, includedElements map[int]bool) {

	zoneItemData, ok := engine.Patch.ZoneItem[element.ID]
	if !ok {
		zoneItemData = engine.State.ZoneItem[element.ID]
	}

	if element.OperationKind != OperationKindUpdate && element.OperationKind != OperationKindDelete {
		element.OperationKind = zoneItemData.OperationKind
	}

	if pIndex+1 == len(p) {
		return
	}

	nextSeg := p[pIndex+1]

	switch nextSeg.identifier {
	case zoneItem_itemIdentifier:
		child := element.Item
		if child == nil {
			child = &item{ID: ItemID(nextSeg.id)}
		}
		engine.assembleItemPath(child, p, pIndex+1, includedElements)
		element.Item = child
	case zoneItem_positionIdentifier:
		child := element.Position
		if child == nil {
			child = &position{ID: PositionID(nextSeg.id)}
		}
		engine.assemblePositionPath(child, p, pIndex+1, includedElements)
		element.Position = child
	}

	_ = zoneItemData
}

func (engine *Engine) assemblePlayerPath(element *player, p path, pIndex int, includedElements map[int]bool) {

	playerData, ok := engine.Patch.Player[element.ID]
	if !ok {
		playerData = engine.State.Player[element.ID]
	}

	if element.OperationKind != OperationKindUpdate && element.OperationKind != OperationKindDelete {
		element.OperationKind = playerData.OperationKind
	}

	if pIndex+1 == len(p) {
		return
	}

	nextSeg := p[pIndex+1]

	switch nextSeg.identifier {
	case player_actionIdentifier:
		if element.Action == nil {
			element.Action = make(map[AttackEventID]attackEvent)
		}
		child, ok := element.Action[AttackEventID(nextSeg.id)]
		if !ok {
			child = attackEvent{ID: AttackEventID(nextSeg.id)}
		}
		engine.assembleAttackEventPath(&child, p, pIndex+1, includedElements)
		element.Action[child.ID] = child
	case player_equipmentSetsIdentifier:
		ref := engine.playerEquipmentSetRef(PlayerEquipmentSetRefID(nextSeg.refID)).playerEquipmentSetRef
		referencedDataStatus := ReferencedDataUnchanged
		if _, ok := includedElements[int(ref.ReferencedElementID)]; ok {
			referencedDataStatus = ReferencedDataModified
		}
		referencedElement := engine.EquipmentSet(ref.ReferencedElementID).equipmentSet
		treeRef := elementReference{
			OperationKind:        ref.OperationKind,
			ElementID:            int(ref.ReferencedElementID),
			ElementKind:          ElementKindEquipmentSet,
			ReferencedDataStatus: referencedDataStatus,
			ElementPath:          referencedElement.Path,
		}
		if element.EquipmentSets == nil {
			element.EquipmentSets = make(map[EquipmentSetID]elementReference)
		}
		element.EquipmentSets[referencedElement.ID] = treeRef
	case player_gearScoreIdentifier:
		child := element.GearScore
		if child == nil {
			child = &gearScore{ID: GearScoreID(nextSeg.id)}
		}
		engine.assembleGearScorePath(child, p, pIndex+1, includedElements)
		element.GearScore = child
	case player_guildMembersIdentifier:
		ref := engine.playerGuildMemberRef(PlayerGuildMemberRefID(nextSeg.refID)).playerGuildMemberRef
		referencedDataStatus := ReferencedDataUnchanged
		if _, ok := includedElements[int(ref.ReferencedElementID)]; ok {
			referencedDataStatus = ReferencedDataModified
		}
		referencedElement := engine.Player(ref.ReferencedElementID).player
		treeRef := elementReference{
			OperationKind:        ref.OperationKind,
			ElementID:            int(ref.ReferencedElementID),
			ElementKind:          ElementKindPlayer,
			ReferencedDataStatus: referencedDataStatus,
			ElementPath:          referencedElement.Path,
		}
		if element.GuildMembers == nil {
			element.GuildMembers = make(map[PlayerID]elementReference)
		}
		element.GuildMembers[referencedElement.ID] = treeRef
	case player_itemsIdentifier:
		if element.Items == nil {
			element.Items = make(map[ItemID]item)
		}
		child, ok := element.Items[ItemID(nextSeg.id)]
		if !ok {
			child = item{ID: ItemID(nextSeg.id)}
		}
		engine.assembleItemPath(&child, p, pIndex+1, includedElements)
		element.Items[child.ID] = child
	case player_positionIdentifier:
		child := element.Position
		if child == nil {
			child = &position{ID: PositionID(nextSeg.id)}
		}
		engine.assemblePositionPath(child, p, pIndex+1, includedElements)
		element.Position = child
	case player_targetIdentifier:
		ref := engine.playerTargetRef(PlayerTargetRefID(nextSeg.refID)).playerTargetRef
		if element.Target != nil && ref.OperationKind == OperationKindDelete {
			break
		}
		referencedDataStatus := ReferencedDataUnchanged
		if _, ok := includedElements[ref.ReferencedElementID.ChildID]; ok {
			referencedDataStatus = ReferencedDataModified
		}
		switch nextSeg.kind {
		case ElementKindPlayer:
			referencedElement := engine.Player(PlayerID(ref.ReferencedElementID.ChildID)).player
			treeRef := elementReference{
				OperationKind:        ref.OperationKind,
				ElementID:            ref.ReferencedElementID.ChildID,
				ElementKind:          ElementKindPlayer,
				ReferencedDataStatus: referencedDataStatus,
				ElementPath:          referencedElement.Path,
			}
			element.Target = &treeRef
		case ElementKindZoneItem:
			referencedElement := engine.ZoneItem(ZoneItemID(ref.ReferencedElementID.ChildID)).zoneItem
			treeRef := elementReference{
				OperationKind:        ref.OperationKind,
				ElementID:            ref.ReferencedElementID.ChildID,
				ElementKind:          ElementKindZoneItem,
				ReferencedDataStatus: referencedDataStatus,
				ElementPath:          referencedElement.Path,
			}
			element.Target = &treeRef
		}
	case player_targetedByIdentifier:
		if element.TargetedBy == nil {
			element.TargetedBy = make(map[int]elementReference)
		}
		ref := engine.playerTargetedByRef(PlayerTargetedByRefID(nextSeg.refID)).playerTargetedByRef
		referencedDataStatus := ReferencedDataUnchanged
		if _, ok := includedElements[ref.ReferencedElementID.ChildID]; ok {
			referencedDataStatus = ReferencedDataModified
		}
		switch nextSeg.kind {
		case ElementKindPlayer:
			referencedElement := engine.Player(PlayerID(ref.ReferencedElementID.ChildID)).player
			treeRef := elementReference{
				OperationKind:        ref.OperationKind,
				ElementID:            ref.ReferencedElementID.ChildID,
				ElementKind:          ElementKindPlayer,
				ReferencedDataStatus: referencedDataStatus,
				ElementPath:          referencedElement.Path,
			}
			element.TargetedBy[ref.ReferencedElementID.ChildID] = treeRef
		case ElementKindZoneItem:
			referencedElement := engine.ZoneItem(ZoneItemID(ref.ReferencedElementID.ChildID)).zoneItem
			treeRef := elementReference{
				OperationKind:        ref.OperationKind,
				ElementID:            ref.ReferencedElementID.ChildID,
				ElementKind:          ElementKindZoneItem,
				ReferencedDataStatus: referencedDataStatus,
				ElementPath:          referencedElement.Path,
			}
			element.TargetedBy[ref.ReferencedElementID.ChildID] = treeRef
		}
	}

	_ = playerData
}

func (engine *Engine) assembleZonePath(element *zone, p path, pIndex int, includedElements map[int]bool) {

	zoneData, ok := engine.Patch.Zone[element.ID]
	if !ok {
		zoneData = engine.State.Zone[element.ID]
	}

	if element.OperationKind != OperationKindUpdate && element.OperationKind != OperationKindDelete {
		element.OperationKind = zoneData.OperationKind
	}

	if pIndex+1 == len(p) {
		return
	}

	nextSeg := p[pIndex+1]

	switch nextSeg.identifier {
	case zone_interactablesIdentifier:
		if element.Interactables == nil {
			element.Interactables = make(map[int]interface{})
		}
		switch nextSeg.kind {
		case ElementKindItem:
			child, ok := element.Interactables[nextSeg.id].(item)
			if !ok {
				child = item{ID: ItemID(nextSeg.id)}
			}
			engine.assembleItemPath(&child, p, pIndex+1, includedElements)
			element.Interactables[nextSeg.id] = child
		case ElementKindPlayer:
			child, ok := element.Interactables[nextSeg.id].(player)
			if !ok {
				child = player{ID: PlayerID(nextSeg.id)}
			}
			engine.assemblePlayerPath(&child, p, pIndex+1, includedElements)
			element.Interactables[nextSeg.id] = child
		case ElementKindZoneItem:
			child, ok := element.Interactables[nextSeg.id].(zoneItem)
			if !ok {
				child = zoneItem{ID: ZoneItemID(nextSeg.id)}
			}
			engine.assembleZoneItemPath(&child, p, pIndex+1, includedElements)
			element.Interactables[nextSeg.id] = child
		}
	case zone_itemsIdentifier:
		if element.Items == nil {
			element.Items = make(map[ZoneItemID]zoneItem)
		}
		child, ok := element.Items[ZoneItemID(nextSeg.id)]
		if !ok {
			child = zoneItem{ID: ZoneItemID(nextSeg.id)}
		}
		engine.assembleZoneItemPath(&child, p, pIndex+1, includedElements)
		element.Items[child.ID] = child
	case zone_playersIdentifier:
		if element.Players == nil {
			element.Players = make(map[PlayerID]player)
		}
		child, ok := element.Players[PlayerID(nextSeg.id)]
		if !ok {
			child = player{ID: PlayerID(nextSeg.id)}
		}
		engine.assemblePlayerPath(&child, p, pIndex+1, includedElements)
		element.Players[child.ID] = child
	case zone_tagsIdentifier:
		if element.Tags == nil {
			element.Tags = make([]string, 0, len(zoneData.Tags))
		}
		child := engine.stringValue(StringValueID(nextSeg.id))
		element.OperationKind = child.OperationKind
		element.Tags = append(element.Tags, child.Value)
	}

	_ = zoneData
}
