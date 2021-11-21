package state

func (engine *Engine) assembleGearScorePath(element *gearScore, p path, pIndex int, includedElements map[int]bool) {

	gearScoreData, ok := engine.Patch.GearScore[element.ID]
	if !ok {
		gearScoreData = engine.State.GearScore[element.ID]
	}

	element.OperationKind = gearScoreData.OperationKind
	element.Level = gearScoreData.Level
	element.Score = gearScoreData.Score

	_ = gearScoreData
}

func (engine *Engine) assemblePositionPath(element *position, p path, pIndex int, includedElements map[int]bool) {

	positionData, ok := engine.Patch.Position[element.ID]
	if !ok {
		positionData = engine.State.Position[element.ID]
	}

	element.OperationKind = positionData.OperationKind
	element.X = positionData.X
	element.Y = positionData.Y

	_ = positionData
}

func (engine *Engine) assembleEquipmentSetPath(element *equipmentSet, p path, pIndex int, includedElements map[int]bool) {

	equipmentSetData, ok := engine.Patch.EquipmentSet[element.ID]
	if !ok {
		equipmentSetData = engine.State.EquipmentSet[element.ID]
	}

	element.OperationKind = equipmentSetData.OperationKind
	element.Name = equipmentSetData.Name

	if pIndex+1 == len(p) {
		return
	}

	switch p[pIndex+1].identifier {
	case equipmentSet_equipmentIdentifier:
		var refID EquipmentSetEquipmentRefID
		for _, refID = range equipmentSetData.Equipment {
			if int(engine.Patch.EquipmentSetEquipmentRef[refID].ReferencedElementID) == p[pIndex+1].id {
				break
			}
			if int(engine.State.EquipmentSetEquipmentRef[refID].ReferencedElementID) == p[pIndex+1].id {
				break
			}
		}
		ref := engine.equipmentSetEquipmentRef(refID).equipmentSetEquipmentRef
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
	}

	_ = equipmentSetData
}

func (engine *Engine) assembleItemPath(element *item, p path, pIndex int, includedElements map[int]bool) {

	itemData, ok := engine.Patch.Item[element.ID]
	if !ok {
		itemData = engine.State.Item[element.ID]
	}

	element.OperationKind = itemData.OperationKind
	element.Name = itemData.Name

	if pIndex+1 == len(p) {
		return
	}

	switch p[pIndex+1].identifier {
	case item_boundToIdentifier:
		ref := engine.itemBoundToRef(ItemBoundToRefID(p[pIndex+1].id)).itemBoundToRef
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
			child = &gearScore{ID: itemData.GearScore}
		}
		engine.assembleGearScorePath(child, p, pIndex+1, includedElements)
		element.GearScore = child
	case item_originIdentifier:
		switch p[pIndex+1].kind {
		case ElementKindPlayer:
			child, ok := element.Origin.(*player)
			if !ok || child == nil {
				child = &player{ID: PlayerID(p[pIndex+1].id)}
			}
			engine.assemblePlayerPath(child, p, pIndex+1, includedElements)
			element.Origin = child
		case ElementKindPosition:
			child, ok := element.Origin.(*position)
			if !ok || child == nil {
				child = &position{ID: PositionID(p[pIndex+1].id)}
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

	element.OperationKind = zoneItemData.OperationKind

	if pIndex+1 == len(p) {
		return
	}

	switch p[pIndex+1].identifier {
	case zoneItem_itemIdentifier:
		child := element.Item
		if child == nil {
			child = &item{ID: zoneItemData.Item}
		}
		engine.assembleItemPath(child, p, pIndex+1, includedElements)
		element.Item = child
	case zoneItem_positionIdentifier:
		child := element.Position
		if child == nil {
			child = &position{ID: zoneItemData.Position}
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

	element.OperationKind = playerData.OperationKind

	if pIndex+1 == len(p) {
		return
	}

	switch p[pIndex+1].identifier {
	case player_equipmentSetsIdentifier:
		var refID PlayerEquipmentSetRefID
		for _, refID = range playerData.EquipmentSets {
			if int(engine.Patch.PlayerEquipmentSetRef[refID].ReferencedElementID) == p[pIndex+1].id {
				break
			}
			if int(engine.State.PlayerEquipmentSetRef[refID].ReferencedElementID) == p[pIndex+1].id {
				break
			}
		}
		ref := engine.playerEquipmentSetRef(refID).playerEquipmentSetRef
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
			child = &gearScore{ID: playerData.GearScore}
		}
		engine.assembleGearScorePath(child, p, pIndex+1, includedElements)
		element.GearScore = child
	case player_guildMembersIdentifier:
		var refID PlayerGuildMemberRefID
		for _, refID = range playerData.GuildMembers {
			if int(engine.Patch.PlayerGuildMemberRef[refID].ReferencedElementID) == p[pIndex+1].id {
				break
			}
			if int(engine.State.PlayerGuildMemberRef[refID].ReferencedElementID) == p[pIndex+1].id {
				break
			}
		}
		ref := engine.playerGuildMemberRef(refID).playerGuildMemberRef
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
		child, ok := element.Items[ItemID(p[pIndex+1].id)]
		if !ok {
			child = item{ID: ItemID(p[pIndex+1].id)}
		}
		engine.assembleItemPath(&child, p, pIndex+1, includedElements)
		element.Items[child.ID] = child
	case player_positionIdentifier:
		child := element.Position
		if child == nil {
			child = &position{ID: playerData.Position}
		}
		engine.assemblePositionPath(child, p, pIndex+1, includedElements)
		element.Position = child
	case player_targetIdentifier:
		ref := engine.playerTargetRef(playerData.Target).playerTargetRef
		referencedDataStatus := ReferencedDataUnchanged
		if _, ok := includedElements[p[pIndex+1].id]; ok {
			referencedDataStatus = ReferencedDataModified
		}
		switch p[pIndex+1].kind {
		case ElementKindPlayer:
			referencedElement := engine.Player(PlayerID(p[pIndex+1].id)).player
			treeRef := elementReference{
				OperationKind:        ref.OperationKind,
				ElementID:            p[pIndex+1].id,
				ElementKind:          ElementKindPlayer,
				ReferencedDataStatus: referencedDataStatus,
				ElementPath:          referencedElement.Path,
			}
			element.Target = &treeRef
		case ElementKindZoneItem:
			referencedElement := engine.ZoneItem(ZoneItemID(p[pIndex+1].id)).zoneItem
			treeRef := elementReference{
				OperationKind:        ref.OperationKind,
				ElementID:            p[pIndex+1].id,
				ElementKind:          ElementKindPlayer,
				ReferencedDataStatus: referencedDataStatus,
				ElementPath:          referencedElement.Path,
			}
			element.Target = &treeRef
		}
	case player_targetedByIdentifier:
		if element.TargetedBy == nil {
			element.TargetedBy = make(map[int]elementReference)
		}
		var refID PlayerTargetedByRefID
		for _, refID = range playerData.TargetedBy {
			if int(engine.Patch.PlayerTargetedByRef[refID].ReferencedElementID) == p[pIndex+1].id {
				break
			}
			if int(engine.State.PlayerTargetedByRef[refID].ReferencedElementID) == p[pIndex+1].id {
				break
			}
		}
		ref := engine.playerTargetedByRef(refID).playerTargetedByRef
		referencedDataStatus := ReferencedDataUnchanged
		if _, ok := includedElements[p[pIndex+1].id]; ok {
			referencedDataStatus = ReferencedDataModified
		}
		switch p[pIndex+1].kind {
		case ElementKindPlayer:
			referencedElement := engine.Player(PlayerID(p[pIndex+1].id)).player
			treeRef := elementReference{
				OperationKind:        ref.OperationKind,
				ElementID:            p[pIndex+1].id,
				ElementKind:          ElementKindPlayer,
				ReferencedDataStatus: referencedDataStatus,
				ElementPath:          referencedElement.Path,
			}
			element.TargetedBy[p[pIndex+1].id] = treeRef
		case ElementKindZoneItem:
			referencedElement := engine.ZoneItem(ZoneItemID(p[pIndex+1].id)).zoneItem
			treeRef := elementReference{
				OperationKind:        ref.OperationKind,
				ElementID:            p[pIndex+1].id,
				ElementKind:          ElementKindPlayer,
				ReferencedDataStatus: referencedDataStatus,
				ElementPath:          referencedElement.Path,
			}
			element.TargetedBy[p[pIndex+1].id] = treeRef
		}
	}

	_ = playerData
}

func (engine *Engine) assembleZonePath(element *zone, p path, pIndex int, includedElements map[int]bool) {

	zoneData, ok := engine.Patch.Zone[element.ID]
	if !ok {
		zoneData = engine.State.Zone[element.ID]
	}

	element.OperationKind = zoneData.OperationKind
	element.Tags = zoneData.Tags[:]

	if pIndex+1 == len(p) {
		return
	}

	switch p[pIndex+1].identifier {
	case zone_interactablesIdentifier:
		if element.Interactables == nil {
			element.Interactables = make(map[int]interface{})
		}
		switch p[pIndex+1].kind {
		case ElementKindItem:
			child, ok := element.Interactables[p[pIndex+1].id].(item)
			if !ok {
				child = item{ID: ItemID(p[pIndex+1].id)}
			}
			engine.assembleItemPath(&child, p, pIndex+1, includedElements)
			element.Interactables[p[pIndex+1].id] = child
		case ElementKindPlayer:
			child, ok := element.Interactables[p[pIndex+1].id].(player)
			if !ok {
				child = player{ID: PlayerID(p[pIndex+1].id)}
			}
			engine.assemblePlayerPath(&child, p, pIndex+1, includedElements)
			element.Interactables[p[pIndex+1].id] = child
		case ElementKindZoneItem:
			child, ok := element.Interactables[p[pIndex+1].id].(zoneItem)
			if !ok {
				child = zoneItem{ID: ZoneItemID(p[pIndex+1].id)}
			}
			engine.assembleZoneItemPath(&child, p, pIndex+1, includedElements)
			element.Interactables[p[pIndex+1].id] = child
		}
	case zone_itemsIdentifier:
		if element.Items == nil {
			element.Items = make(map[ZoneItemID]zoneItem)
		}
		child, ok := element.Items[ZoneItemID(p[pIndex+1].id)]
		if !ok {
			child = zoneItem{ID: ZoneItemID(p[pIndex+1].id)}
		}
		engine.assembleZoneItemPath(&child, p, pIndex+1, includedElements)
		element.Items[child.ID] = child
	case zone_playersIdentifier:
		if element.Players == nil {
			element.Players = make(map[PlayerID]player)
		}
		child, ok := element.Players[PlayerID(p[pIndex+1].id)]
		if !ok {
			child = player{ID: PlayerID(p[pIndex+1].id)}
		}
		engine.assemblePlayerPath(&child, p, pIndex+1, includedElements)
		element.Players[child.ID] = child
	}

	_ = zoneData
}

// 1. get all basic elements and references out of patch, put their paths in updatedPaths
// 2. go through all paths and put ids in includedElements map, save len(includedElements)
// 3. get all references out of STATE, check if they reference element in includedElements, if TRUE put reference path into updatedByReferecePaths
// 4. if len(updatedByReferecePaths) != 0: go through all updatedByReferecePaths and put ids in includedElements map, ELSE continue with 6.
// 5. back to step 3.

// TODO PROBLEM?? if a reference is Set and then Unset and Set again, does that mess with pathuilding, as referenceID might be 0 in patch or state
// TODO what happens if you call SetPlayer? will the path be built with a player-update or a position-delete??

func (engine *Engine) assembleUpdateTree() Tree {

	for key := range engine.Tree.EquipmentSet {
		delete(engine.Tree.EquipmentSet, key)
	}
	for key := range engine.Tree.GearScore {
		delete(engine.Tree.GearScore, key)
	}
	for key := range engine.Tree.Item {
		delete(engine.Tree.Item, key)
	}
	for key := range engine.Tree.Player {
		delete(engine.Tree.Player, key)
	}
	for key := range engine.Tree.Position {
		delete(engine.Tree.Position, key)
	}
	for key := range engine.Tree.Zone {
		delete(engine.Tree.Zone, key)
	}
	for key := range engine.Tree.ZoneItem {
		delete(engine.Tree.ZoneItem, key)
	}

	updatedPaths := make(map[int]path)
	// TODO possibly big performance boost
	// updatedElements := make(map[int]bool)

	for _, equipmentSet := range engine.Patch.EquipmentSet {
		updatedPaths[int(equipmentSet.ID)] = equipmentSet.path
	}
	for _, gearScore := range engine.Patch.GearScore {
		updatedPaths[int(gearScore.ID)] = gearScore.path
	}
	for _, item := range engine.Patch.Item {
		updatedPaths[int(item.ID)] = item.path
	}
	for _, player := range engine.Patch.Player {
		updatedPaths[int(player.ID)] = player.path
	}
	for _, position := range engine.Patch.Position {
		updatedPaths[int(position.ID)] = position.path
	}
	for _, zone := range engine.Patch.Zone {
		updatedPaths[int(zone.ID)] = zone.path
	}
	for _, zoneItem := range engine.Patch.ZoneItem {
		updatedPaths[int(zoneItem.ID)] = zoneItem.path
	}
	for _, equipmentSetEquipmentRef := range engine.Patch.EquipmentSetEquipmentRef {
		updatedPaths[int(equipmentSetEquipmentRef.ID)] = equipmentSetEquipmentRef.path
	}
	for _, itemBoundToRef := range engine.Patch.ItemBoundToRef {
		updatedPaths[int(itemBoundToRef.ID)] = itemBoundToRef.path
	}
	for _, playerEquipmentSetRef := range engine.Patch.PlayerEquipmentSetRef {
		updatedPaths[int(playerEquipmentSetRef.ID)] = playerEquipmentSetRef.path
	}
	for _, playerGuildMemberRef := range engine.Patch.PlayerGuildMemberRef {
		updatedPaths[int(playerGuildMemberRef.ID)] = playerGuildMemberRef.path
	}
	for _, playerTargetRef := range engine.Patch.PlayerTargetRef {
		// TODO because paths contain the ids of referenced elements (not ids of references) changing a reference
		// includes the newly referenced element in the path. this is why the referenced element is considered updated
		updatedPaths[int(playerTargetRef.ID)] = playerTargetRef.path
	}
	for _, playerTargetedByRef := range engine.Patch.PlayerTargetedByRef {
		updatedPaths[int(playerTargetedByRef.ID)] = playerTargetedByRef.path
	}

	includedElements := make(map[int]bool)

	prevousLength := len(updatedPaths)
	for {
		for _, p := range updatedPaths {
			for _, seg := range p {
				includedElements[seg.id] = true
			}
		}
		// TODO in case a reference is newly created and references an element that
		// enters includedElements at a later iteration, I also need to loop
		// over patch references again

		// TODO failing test cases for all todos

		for _, equipmentSetEquipmentRef := range engine.State.EquipmentSetEquipmentRef {
			if _, ok := includedElements[int(equipmentSetEquipmentRef.ReferencedElementID)]; !ok {
				updatedPaths[int(equipmentSetEquipmentRef.ID)] = equipmentSetEquipmentRef.path
			}
		}
		for _, itemBoundToRef := range engine.State.ItemBoundToRef {
			if _, ok := includedElements[int(itemBoundToRef.ReferencedElementID)]; !ok {
				updatedPaths[int(itemBoundToRef.ID)] = itemBoundToRef.path
			}
		}
		for _, playerEquipmentSetRef := range engine.State.PlayerEquipmentSetRef {
			if _, ok := includedElements[int(playerEquipmentSetRef.ReferencedElementID)]; !ok {
				updatedPaths[int(playerEquipmentSetRef.ID)] = playerEquipmentSetRef.path
			}
		}
		for _, playerGuildMemberRef := range engine.State.PlayerGuildMemberRef {
			if _, ok := includedElements[int(playerGuildMemberRef.ReferencedElementID)]; !ok {
				updatedPaths[int(playerGuildMemberRef.ID)] = playerGuildMemberRef.path
			}
		}
		for _, playerTargetRef := range engine.State.PlayerTargetRef {
			anyContainer := engine.anyOfPlayer_ZoneItem(playerTargetRef.ReferencedElementID)
			switch anyContainer.anyOfPlayer_ZoneItem.ElementKind {
			case ElementKindPlayer:
				if _, ok := includedElements[int(anyContainer.anyOfPlayer_ZoneItem.Player)]; !ok {
					updatedPaths[int(playerTargetRef.ID)] = playerTargetRef.path
				}
			case ElementKindZoneItem:
				if _, ok := includedElements[int(anyContainer.anyOfPlayer_ZoneItem.Player)]; !ok {
					updatedPaths[int(playerTargetRef.ID)] = playerTargetRef.path
				}
			}
		}
		for _, playerTargetedByRef := range engine.State.PlayerTargetedByRef {
			anyContainer := engine.anyOfPlayer_ZoneItem(playerTargetedByRef.ReferencedElementID)
			switch anyContainer.anyOfPlayer_ZoneItem.ElementKind {
			case ElementKindPlayer:
				if _, ok := includedElements[int(anyContainer.anyOfPlayer_ZoneItem.Player)]; !ok {
					updatedPaths[int(playerTargetedByRef.ID)] = playerTargetedByRef.path
				}
			case ElementKindZoneItem:
				if _, ok := includedElements[int(anyContainer.anyOfPlayer_ZoneItem.Player)]; !ok {
					updatedPaths[int(playerTargetedByRef.ID)] = playerTargetedByRef.path
				}
			}
		}

		if prevousLength == len(updatedPaths) {
			break
		}

		prevousLength = len(updatedPaths)
	}

	for _, elementPath := range updatedPaths {
		switch elementPath[0].identifier {
		case equipmentSetIdentifier:
			child, ok := engine.Tree.EquipmentSet[EquipmentSetID(elementPath[0].id)]
			if !ok {
				child = equipmentSet{ID: EquipmentSetID(elementPath[0].id)}
			}
			engine.assembleEquipmentSetPath(&child, elementPath, 0, includedElements)
			engine.Tree.EquipmentSet[EquipmentSetID(elementPath[0].id)] = child
		case gearScoreIdentifier:
			child, ok := engine.Tree.GearScore[GearScoreID(elementPath[0].id)]
			if !ok {
				child = gearScore{ID: GearScoreID(elementPath[0].id)}
			}
			engine.assembleGearScorePath(&child, elementPath, 0, includedElements)
			engine.Tree.GearScore[GearScoreID(elementPath[0].id)] = child
		case itemIdentifier:
			child, ok := engine.Tree.Item[ItemID(elementPath[0].id)]
			if !ok {
				child = item{ID: ItemID(elementPath[0].id)}
			}
			engine.assembleItemPath(&child, elementPath, 0, includedElements)
			engine.Tree.Item[ItemID(elementPath[0].id)] = child
		case playerIdentifier:
			child, ok := engine.Tree.Player[PlayerID(elementPath[0].id)]
			if !ok {
				child = player{ID: PlayerID(elementPath[0].id)}
			}
			engine.assemblePlayerPath(&child, elementPath, 0, includedElements)
			engine.Tree.Player[PlayerID(elementPath[0].id)] = child
		case positionIdentifier:
			child, ok := engine.Tree.Position[PositionID(elementPath[0].id)]
			if !ok {
				child = position{ID: PositionID(elementPath[0].id)}
			}
			engine.assemblePositionPath(&child, elementPath, 0, includedElements)
			engine.Tree.Position[PositionID(elementPath[0].id)] = child
		case zoneIdentifier:
			child, ok := engine.Tree.Zone[ZoneID(elementPath[0].id)]
			if !ok {
				child = zone{ID: ZoneID(elementPath[0].id)}
			}
			engine.assembleZonePath(&child, elementPath, 0, includedElements)
			engine.Tree.Zone[ZoneID(elementPath[0].id)] = child
		case zoneItemIdentifier:
			child, ok := engine.Tree.ZoneItem[ZoneItemID(elementPath[0].id)]
			if !ok {
				child = zoneItem{ID: ZoneItemID(elementPath[0].id)}
			}
			engine.assembleZoneItemPath(&child, elementPath, 0, includedElements)
			engine.Tree.ZoneItem[ZoneItemID(elementPath[0].id)] = child
		}
	}

	return engine.Tree
}
