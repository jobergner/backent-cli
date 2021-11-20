package state

import "fmt"

func (engine *Engine) planGearScoreTree(elementID GearScoreID, p path, pIndex int, includedElements map[int]bool) {
	gearScoreData, ok := engine.Patch.GearScore[elementID]
	if !ok {
		gearScoreData = engine.State.GearScore[elementID]
	}

	includedElements[int(elementID)] = true

	_ = gearScoreData
}

func (engine *Engine) assembleGearScorePath(element *gearScore, p path, pIndex int, includedElements map[int]bool) {
	includedElements[int(element.ID)] = true

	gearScoreData, ok := engine.Patch.GearScore[element.ID]
	if !ok {
		gearScoreData = engine.State.GearScore[element.ID]
	}

	element.OperationKind = gearScoreData.OperationKind
	element.Level = gearScoreData.Level
	element.Score = gearScoreData.Score

	_ = gearScoreData
}

func (engine *Engine) planPositionTree(elementID PositionID, p path, pIndex int, includedElements map[int]bool) {
	positionData, ok := engine.Patch.Position[elementID]
	if !ok {
		positionData = engine.State.Position[elementID]
	}

	includedElements[int(elementID)] = true

	_ = positionData
}

func (engine *Engine) assemblePositionPath(element *position, p path, pIndex int, includedElements map[int]bool) {
	includedElements[int(element.ID)] = true

	positionData, ok := engine.Patch.Position[element.ID]
	if !ok {
		positionData = engine.State.Position[element.ID]
	}

	element.OperationKind = positionData.OperationKind
	element.X = positionData.X
	element.Y = positionData.Y

	_ = positionData
}

func (engine *Engine) planEquipmentSetTree(elementID EquipmentSetID, p path, pIndex int, includedElements map[int]bool) {
	equipmentSetData, ok := engine.Patch.EquipmentSet[elementID]
	if !ok {
		equipmentSetData = engine.State.EquipmentSet[elementID]
	}

	includedElements[int(elementID)] = true

	_ = equipmentSetData
}

func (engine *Engine) assembleEquipmentSetPath(element *equipmentSet, p path, pIndex int, includedElements map[int]bool) {
	includedElements[int(element.ID)] = true

	equipmentSetData, ok := engine.Patch.EquipmentSet[element.ID]
	if !ok {
		equipmentSetData = engine.State.EquipmentSet[element.ID]
	}

	element.OperationKind = equipmentSetData.OperationKind
	element.Name = equipmentSetData.Name

	_ = equipmentSetData
}

func (engine *Engine) planItemTree(elementID ItemID, p path, pIndex int, includedElements map[int]bool) {
	itemData, ok := engine.Patch.Item[elementID]
	if !ok {
		itemData = engine.State.Item[elementID]
	}

	includedElements[int(elementID)] = true

	switch p[pIndex] {
	case gearScoreIdentifier:
		engine.planGearScoreTree(itemData.GearScore, p, pIndex+1, includedElements)
	case originIdentifier:
		anyOfPlayer_PositionContainer := engine.anyOfPlayer_Position(itemData.Origin).anyOfPlayer_Position
		switch anyOfPlayer_PositionContainer.ElementKind {
		case ElementKindPlayer:
			engine.planPlayerTree(anyOfPlayer_PositionContainer.Player, p, pIndex+1, includedElements)
		case ElementKindPosition:
			engine.planPositionTree(anyOfPlayer_PositionContainer.Position, p, pIndex+1, includedElements)
		}
	}

	_ = itemData
}

func (engine *Engine) assembleItemPath(element *item, p path, pIndex int, includedElements map[int]bool) {
	includedElements[int(element.ID)] = true

	itemData, ok := engine.Patch.Item[element.ID]
	if !ok {
		itemData = engine.State.Item[element.ID]
	}

	element.OperationKind = itemData.OperationKind
	element.Name = itemData.Name

	if pIndex == len(p) {
		return
	}

	switch p[pIndex] {
	case gearScoreIdentifier:
		child := element.GearScore
		if child == nil {
			child = &gearScore{ID: itemData.GearScore}
		}
		engine.assembleGearScorePath(child, p, pIndex+1, includedElements)
		element.GearScore = child
	case originIdentifier:
		anyOfPlayer_PositionContainer := engine.anyOfPlayer_Position(itemData.Origin).anyOfPlayer_Position
		switch anyOfPlayer_PositionContainer.ElementKind {
		case ElementKindPlayer:
			if element.Origin == nil {
				element.Origin = &player{ID: anyOfPlayer_PositionContainer.Player}
			}
			child := element.Origin.(*player)
			engine.assemblePlayerPath(child, p, pIndex+1, includedElements)
			element.Origin = child
		case ElementKindPosition:
			if element.Origin == nil {
				element.Origin = &position{ID: anyOfPlayer_PositionContainer.Position}
			}
			child := element.Origin.(*position)
			engine.assemblePositionPath(child, p, pIndex+1, includedElements)
			element.Origin = child
		}
	}

	_ = itemData
}

func (engine *Engine) planZoneItemTree(elementID ZoneItemID, p path, pIndex int, includedElements map[int]bool) {
	zoneItemData, ok := engine.Patch.ZoneItem[elementID]
	if !ok {
		zoneItemData = engine.State.ZoneItem[elementID]
	}

	includedElements[int(elementID)] = true

	switch p[pIndex] {
	case itemIdentifier:
		engine.planItemTree(zoneItemData.Item, p, pIndex+1, includedElements)
	case positionIdentifier:
		engine.planPositionTree(zoneItemData.Position, p, pIndex+1, includedElements)
	}

	_ = zoneItemData
}

func (engine *Engine) assembleZoneItemPath(element *zoneItem, p path, pIndex int, includedElements map[int]bool) {
	includedElements[int(element.ID)] = true

	zoneItemData, ok := engine.Patch.ZoneItem[element.ID]
	if !ok {
		zoneItemData = engine.State.ZoneItem[element.ID]
	}

	element.OperationKind = zoneItemData.OperationKind

	if pIndex == len(p) {
		return
	}

	switch p[pIndex] {
	case itemIdentifier:
		child := element.Item
		if child == nil {
			child = &item{ID: zoneItemData.Item}
		}
		engine.assembleItemPath(child, p, pIndex+1, includedElements)
		element.Item = child
	case positionIdentifier:
		child := element.Position
		if child == nil {
			child = &position{ID: zoneItemData.Position}
		}
		engine.assemblePositionPath(child, p, pIndex+1, includedElements)
		element.Position = child
	}

	_ = zoneItemData
}

func (engine *Engine) planPlayerTree(elementID PlayerID, p path, pIndex int, includedElements map[int]bool) {
	playerData, ok := engine.Patch.Player[elementID]
	if !ok {
		playerData = engine.State.Player[elementID]
	}

	includedElements[int(elementID)] = true

	switch p[pIndex] {
	case gearScoreIdentifier:
		engine.planGearScoreTree(playerData.GearScore, p, pIndex+1, includedElements)
	case itemsIdentifier:
		engine.planItemTree(ItemID(p[pIndex+1]), p, pIndex+2, includedElements)
	case positionIdentifier:
		engine.planPositionTree(playerData.Position, p, pIndex+1, includedElements)
	}

	_ = playerData
}

func (engine *Engine) assemblePlayerPath(element *player, p path, pIndex int, includedElements map[int]bool) {
	includedElements[int(element.ID)] = true

	playerData, ok := engine.Patch.Player[element.ID]
	if !ok {
		playerData = engine.State.Player[element.ID]
	}

	element.OperationKind = playerData.OperationKind

	if pIndex == len(p) {
		return
	}

	switch p[pIndex] {
	case gearScoreIdentifier:
		child := element.GearScore
		if child == nil {
			child = &gearScore{ID: playerData.GearScore}
		}
		engine.assembleGearScorePath(child, p, pIndex+1, includedElements)
		element.GearScore = child
	case itemsIdentifier:
		if element.Items == nil {
			element.Items = make(map[ItemID]item)
		}
		child, ok := element.Items[ItemID(p[pIndex+1])]
		if !ok {
			child = item{ID: ItemID(p[pIndex+1])}
		}
		engine.assembleItemPath(&child, p, pIndex+2, includedElements)
		element.Items[child.ID] = child
	case positionIdentifier:
		child := element.Position
		if child == nil {
			child = &position{ID: playerData.Position}
		}
		engine.assemblePositionPath(child, p, pIndex+1, includedElements)
		element.Position = child
	}

	_ = playerData
}

func (engine *Engine) planZoneTree(elementID ZoneID, p path, pIndex int, includedElements map[int]bool) {
	zoneData, ok := engine.Patch.Zone[elementID]
	if !ok {
		zoneData = engine.State.Zone[elementID]
	}

	includedElements[int(elementID)] = true

	switch p[pIndex] {
	case interactablesIdentifier:
		anyOfItem_Player_ZoneItemContainer := engine.anyOfItem_Player_ZoneItem(AnyOfItem_Player_ZoneItemID(p[pIndex+1])).anyOfItem_Player_ZoneItem
		switch anyOfItem_Player_ZoneItemContainer.ElementKind {
		case ElementKindItem:
			engine.planItemTree(ItemID(p[pIndex+1]), p, pIndex+2, includedElements)
		case ElementKindPlayer:
			engine.planPlayerTree(PlayerID(p[pIndex+1]), p, pIndex+2, includedElements)
		case ElementKindZoneItem:
			engine.planZoneItemTree(ZoneItemID(p[pIndex+1]), p, pIndex+2, includedElements)
		}
	case itemsIdentifier:
		engine.planItemTree(ItemID(p[pIndex+1]), p, pIndex+2, includedElements)
	case playersIdentifier:
		engine.planPlayerTree(PlayerID(p[pIndex+1]), p, pIndex+2, includedElements)
	}

	_ = zoneData
}

func (engine *Engine) assembleZonePath(element *zone, p path, pIndex int, includedElements map[int]bool) {
	includedElements[int(element.ID)] = true

	zoneData, ok := engine.Patch.Zone[element.ID]
	if !ok {
		zoneData = engine.State.Zone[element.ID]
	}

	element.OperationKind = zoneData.OperationKind
	element.Tags = zoneData.Tags[:]

	if pIndex == len(p) {
		return
	}

	switch p[pIndex] {
	case interactablesIdentifier:
		if element.Interactables == nil {
			element.Interactables = make(map[int]interface{})
		}
		anyOfItem_Player_ZoneItemContainer := engine.anyOfItem_Player_ZoneItem(AnyOfItem_Player_ZoneItemID(p[pIndex+1])).anyOfItem_Player_ZoneItem
		switch anyOfItem_Player_ZoneItemContainer.ElementKind {
		case ElementKindItem:
			child := element.Interactables[p[pIndex+1]].(*item)
			if child == nil {
				child = &item{ID: anyOfItem_Player_ZoneItemContainer.Item}
			}
			engine.assembleItemPath(child, p, pIndex+2, includedElements)
			element.Interactables[p[pIndex+1]] = child
		case ElementKindPlayer:
			child := element.Interactables[p[pIndex+1]].(*player)
			if child == nil {
				child = &player{ID: anyOfItem_Player_ZoneItemContainer.Player}
			}
			engine.assemblePlayerPath(child, p, pIndex+2, includedElements)
			element.Interactables[p[pIndex+1]] = child
		case ElementKindZoneItem:
			child := element.Interactables[p[pIndex+1]].(*zoneItem)
			if child == nil {
				child = &zoneItem{ID: anyOfItem_Player_ZoneItemContainer.ZoneItem}
			}
			engine.assembleZoneItemPath(child, p, pIndex+2, includedElements)
			element.Interactables[p[pIndex+1]] = child
		}
	case itemsIdentifier:
		if element.Items == nil {
			element.Items = make(map[ZoneItemID]zoneItem)
		}
		child, ok := element.Items[ZoneItemID(p[pIndex+1])]
		if !ok {
			child = zoneItem{ID: ZoneItemID(p[pIndex+1])}
		}
		engine.assembleZoneItemPath(&child, p, pIndex+2, includedElements)
		element.Items[child.ID] = child
	case playersIdentifier:
		if element.Players == nil {
			element.Players = make(map[PlayerID]player)
		}
		child, ok := element.Players[PlayerID(p[pIndex+1])]
		if !ok {
			child = player{ID: PlayerID(p[pIndex+1])}
		}
		engine.assemblePlayerPath(&child, p, pIndex+2, includedElements)
		element.Players[child.ID] = child
	}

	_ = zoneData
}

// assemble basic elements and references which have updated + correct referencedDataModified state
// 1. get updatedPaths (paths of all elements which have updated (feed all updated elements + updated references in planTree))
// 2. get updatedElements (a list of all elements) from result 1.
// 3. save result 1. for later as essentialPaths
// 4. get all references out of State which reference elements in result 2. and add paths to updatedPaths
// 5. save length of list of all included elements
// 6. expand updatedElements by planning tree with new ipdatedPaths
// 7. compare new length of updatedElements with previous, if increased return to step 4.
// 8. finally assemble tree with essential paths by passing updatedElements

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
	for _, equipmentSetEquipmentRef := range engine.State.EquipmentSetEquipmentRef {
		if _, ok := includedElements[int(equipmentSetEquipmentRef.ReferencedElementID)]; ok {
			updatedPaths[int(equipmentSetEquipmentRef.ID)] = equipmentSetEquipmentRef.path
		}
	}
	for _, equipmentSetEquipmentRef := range engine.Patch.EquipmentSetEquipmentRef {
		updatedPaths[int(equipmentSetEquipmentRef.ID)] = equipmentSetEquipmentRef.path
	}
	for _, itemBoundToRef := range engine.State.ItemBoundToRef {
		if _, ok := includedElements[int(itemBoundToRef.ReferencedElementID)]; ok {
			updatedPaths[int(itemBoundToRef.ID)] = itemBoundToRef.path
		}
	}
	for _, itemBoundToRef := range engine.Patch.ItemBoundToRef {
		updatedPaths[int(itemBoundToRef.ID)] = itemBoundToRef.path
	}
	for _, playerEquipmentSetRef := range engine.State.PlayerEquipmentSetRef {
		if _, ok := includedElements[int(playerEquipmentSetRef.ReferencedElementID)]; ok {
			updatedPaths[int(playerEquipmentSetRef.ID)] = playerEquipmentSetRef.path
		}
	}
	for _, playerEquipmentSetRef := range engine.Patch.PlayerEquipmentSetRef {
		updatedPaths[int(playerEquipmentSetRef.ID)] = playerEquipmentSetRef.path
	}
	for _, playerGuildMemberRef := range engine.State.PlayerGuildMemberRef {
		if _, ok := includedElements[int(playerGuildMemberRef.ReferencedElementID)]; ok {
			updatedPaths[int(playerGuildMemberRef.ID)] = playerGuildMemberRef.path
		}
	}
	for _, playerGuildMemberRef := range engine.Patch.PlayerGuildMemberRef {
		updatedPaths[int(playerGuildMemberRef.ID)] = playerGuildMemberRef.path
	}
	for _, playerTargetRef := range engine.State.PlayerTargetRef {
		if _, ok := includedElements[int(playerTargetRef.ReferencedElementID)]; ok {
			updatedPaths[int(playerTargetRef.ID)] = playerTargetRef.path
		}
	}
	for _, playerTargetRef := range engine.Patch.PlayerTargetRef {
		updatedPaths[int(playerTargetRef.ID)] = playerTargetRef.path
	}
	for _, playerTargetedByRef := range engine.State.PlayerTargetedByRef {
		if _, ok := includedElements[int(playerTargetedByRef.ReferencedElementID)]; ok {
			updatedPaths[int(playerTargetedByRef.ID)] = playerTargetedByRef.path
		}
	}
	for _, playerTargetedByRef := range engine.Patch.PlayerTargetedByRef {
		updatedPaths[int(playerTargetedByRef.ID)] = playerTargetedByRef.path
	}

	includedElements := newIncludedElements()

	for _, elementPath := range updatedPaths {
		switch elementPath[0] {
		case equipmentSetIdentifier:
			child, ok := engine.Tree.EquipmentSet[EquipmentSetID(elementPath[1])]
			if !ok {
				child = equipmentSet{ID: EquipmentSetID(elementPath[1])}
			}
			engine.assembleEquipmentSetPath(&child, elementPath, 2, includedElements)
			engine.Tree.EquipmentSet[EquipmentSetID(elementPath[1])] = child
		case gearScoreIdentifier:
			child, ok := engine.Tree.GearScore[GearScoreID(elementPath[1])]
			if !ok {
				child = gearScore{ID: GearScoreID(elementPath[1])}
			}
			engine.assembleGearScorePath(&child, elementPath, 2, includedElements)
			engine.Tree.GearScore[GearScoreID(elementPath[1])] = child
		case itemIdentifier:
			child, ok := engine.Tree.Item[ItemID(elementPath[1])]
			if !ok {
				child = item{ID: ItemID(elementPath[1])}
			}
			engine.assembleItemPath(&child, elementPath, 2, includedElements)
			engine.Tree.Item[ItemID(elementPath[1])] = child
		case playerIdentifier:
			child, ok := engine.Tree.Player[PlayerID(elementPath[1])]
			if !ok {
				child = player{ID: PlayerID(elementPath[1])}
			}
			engine.assemblePlayerPath(&child, elementPath, 2, includedElements)
			engine.Tree.Player[PlayerID(elementPath[1])] = child
		case positionIdentifier:
			child, ok := engine.Tree.Position[PositionID(elementPath[1])]
			if !ok {
				child = position{ID: PositionID(elementPath[1])}
			}
			engine.assemblePositionPath(&child, elementPath, 2, includedElements)
			engine.Tree.Position[PositionID(elementPath[1])] = child
		case zoneIdentifier:
			child, ok := engine.Tree.Zone[ZoneID(elementPath[1])]
			if !ok {
				child = zone{ID: ZoneID(elementPath[1])}
			}
			engine.assembleZonePath(&child, elementPath, 2, includedElements)
			engine.Tree.Zone[ZoneID(elementPath[1])] = child
		case zoneItemIdentifier:
			child, ok := engine.Tree.ZoneItem[ZoneItemID(elementPath[1])]
			if !ok {
				child = zoneItem{ID: ZoneItemID(elementPath[1])}
			}
			engine.assembleZoneItemPath(&child, elementPath, 2, includedElements)
			engine.Tree.ZoneItem[ZoneItemID(elementPath[1])] = child
		}
	}

	for _, elementPath := range updatedPaths {
		switch elementPath[0] {
		case equipmentSetIdentifier:
			child := engine.Tree.EquipmentSet[EquipmentSetID(elementPath[1])]
			engine.assembleEquipmentSetReferences(&child, elementPath, 2, includedElements)
			engine.Tree.EquipmentSet[EquipmentSetID(elementPath[1])] = child
		case gearScoreIdentifier:
			child := engine.Tree.GearScore[GearScoreID(elementPath[1])]
			engine.assembleGearScoreReferences(&child, elementPath, 2, includedElements)
			engine.Tree.GearScore[GearScoreID(elementPath[1])] = child
		case itemIdentifier:
			fmt.Println("@", elementPath, engine.Tree.Item)
			child := engine.Tree.Item[ItemID(elementPath[1])]
			engine.assembleItemReferences(&child, elementPath, 2, includedElements)
			engine.Tree.Item[ItemID(elementPath[1])] = child
		case playerIdentifier:
			child := engine.Tree.Player[PlayerID(elementPath[1])]
			engine.assemblePlayerReferences(&child, elementPath, 2, includedElements)
			engine.Tree.Player[PlayerID(elementPath[1])] = child
		case positionIdentifier:
			child := engine.Tree.Position[PositionID(elementPath[1])]
			engine.assemblePositionReferences(&child, elementPath, 2, includedElements)
			engine.Tree.Position[PositionID(elementPath[1])] = child
		case zoneIdentifier:
			child := engine.Tree.Zone[ZoneID(elementPath[1])]
			engine.assembleZoneReferences(&child, elementPath, 2, includedElements)
			engine.Tree.Zone[ZoneID(elementPath[1])] = child
		case zoneItemIdentifier:
			child := engine.Tree.ZoneItem[ZoneItemID(elementPath[1])]
			engine.assembleZoneItemReferences(&child, elementPath, 2, includedElements)
			engine.Tree.ZoneItem[ZoneItemID(elementPath[1])] = child
		}
	}

	return engine.Tree
}

func (engine *Engine) assembleItemReferences(element *item, p path, pIndex int, includedElements map[int]bool) {
	itemData, ok := engine.Patch.Item[element.ID]
	if !ok {
		itemData = engine.State.Item[element.ID]
	}

	switch p[pIndex] {
	case boundToIdentifier:
		ref := engine.itemBoundToRef(itemData.BoundTo).itemBoundToRef
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
	case gearScoreIdentifier:
		engine.assembleGearScoreReferences(element.GearScore, p, pIndex+1, includedElements)
	case originIdentifier:
		switch v := element.Origin.(type) {
		case *player:
			engine.assemblePlayerReferences(v, p, pIndex+1, includedElements)
		case *position:
			engine.assemblePositionReferences(v, p, pIndex+1, includedElements)
		}
	}

	_ = itemData
}

func (engine *Engine) assembleGearScoreReferences(element *gearScore, p path, pIndex int, includedElements map[int]bool) {
}

func (engine *Engine) assemblePositionReferences(element *position, p path, pIndex int, includedElements map[int]bool) {
}

func (engine *Engine) assemblePlayerReferences(element *player, p path, pIndex int, includedElements map[int]bool) {
	playerData, ok := engine.Patch.Player[element.ID]
	if !ok {
		playerData = engine.State.Player[element.ID]
	}

	switch p[pIndex] {
	case equipmentSetsIdentifier:
		ref := engine.playerEquipmentSetRef(PlayerEquipmentSetRefID(playerData.EquipmentSets[p[pIndex+1]])).playerEquipmentSetRef
		referencedDataStatus := ReferencedDataUnchanged
		if _, ok := includedElements[int(ref.ReferencedElementID)]; ok {
			referencedDataStatus = ReferencedDataModified
		}
		referencedElement := engine.EquipmentSet(ref.ReferencedElementID).equipmentSet
		treeRef := elementReference{
			OperationKind:        ref.OperationKind,
			ElementID:            int(ref.ReferencedElementID),
			ElementKind:          ElementKindPlayer,
			ReferencedDataStatus: referencedDataStatus,
			ElementPath:          referencedElement.Path,
		}
		if element.EquipmentSets == nil {
			element.EquipmentSets = make(map[EquipmentSetID]elementReference)
		}
		element.EquipmentSets[referencedElement.ID] = treeRef
	case gearScoreIdentifier:
		engine.assembleGearScoreReferences(element.GearScore, p, pIndex+1, includedElements)
	case guildMembersIdentifier:
		ref := engine.playerGuildMemberRef(PlayerGuildMemberRefID(playerData.GuildMembers[p[pIndex+1]])).playerGuildMemberRef
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
	case itemsIdentifier:
		child := element.Items[ItemID(p[pIndex+1])]
		engine.assembleItemReferences(&child, p, pIndex+2, includedElements)
		element.Items[ItemID(p[pIndex+1])] = child
	case positionIdentifier:
		engine.assemblePositionReferences(element.Position, p, pIndex+1, includedElements)
	case targetIdentifier:
		ref := engine.playerTargetRef(playerData.Target).playerTargetRef
		anyContainer := engine.anyOfPlayer_ZoneItem(ref.ReferencedElementID).anyOfPlayer_ZoneItem
		switch anyContainer.ElementKind {
		case ElementKindPlayer:
			referencedDataStatus := ReferencedDataUnchanged
			if _, ok := includedElements[int(anyContainer.Player)]; ok {
				referencedDataStatus = ReferencedDataModified
			}
			referencedElement := engine.Player(anyContainer.Player).player
			treeRef := elementReference{
				OperationKind:        playerData.OperationKind,
				ElementID:            int(anyContainer.Player),
				ElementKind:          ElementKindPlayer,
				ReferencedDataStatus: referencedDataStatus,
				ElementPath:          referencedElement.Path,
			}
			element.Target = &treeRef
		case ElementKindZoneItem:
			referencedDataStatus := ReferencedDataUnchanged
			if _, ok := includedElements[int(anyContainer.ZoneItem)]; ok {
				referencedDataStatus = ReferencedDataModified
			}
			referencedElement := engine.ZoneItem(anyContainer.ZoneItem).zoneItem
			treeRef := elementReference{
				OperationKind:        playerData.OperationKind,
				ElementID:            int(anyContainer.ZoneItem),
				ElementKind:          ElementKindPlayer,
				ReferencedDataStatus: referencedDataStatus,
				ElementPath:          referencedElement.Path,
			}
			element.Target = &treeRef
		}
	case targetedByIdentifier:
		ref := engine.playerTargetedByRef(playerData.TargetedBy[p[pIndex+1]]).playerTargetedByRef
		anyContainer := engine.anyOfPlayer_ZoneItem(ref.ReferencedElementID).anyOfPlayer_ZoneItem
		switch anyContainer.ElementKind {
		case ElementKindPlayer:
			referencedDataStatus := ReferencedDataUnchanged
			if _, ok := includedElements[int(anyContainer.Player)]; ok {
				referencedDataStatus = ReferencedDataModified
			}
			referencedElement := engine.Player(anyContainer.Player).player
			treeRef := elementReference{
				OperationKind:        playerData.OperationKind,
				ElementID:            int(anyContainer.Player),
				ElementKind:          ElementKindPlayer,
				ReferencedDataStatus: referencedDataStatus,
				ElementPath:          referencedElement.Path,
			}
			element.Target = &treeRef
		case ElementKindZoneItem:
			referencedDataStatus := ReferencedDataUnchanged
			if _, ok := includedElements[int(anyContainer.ZoneItem)]; ok {
				referencedDataStatus = ReferencedDataModified
			}
			referencedElement := engine.ZoneItem(anyContainer.ZoneItem).zoneItem
			treeRef := elementReference{
				OperationKind:        playerData.OperationKind,
				ElementID:            int(anyContainer.ZoneItem),
				ElementKind:          ElementKindPlayer,
				ReferencedDataStatus: referencedDataStatus,
				ElementPath:          referencedElement.Path,
			}
			element.Target = &treeRef
		}
	}

	_ = playerData
}

func (engine *Engine) assembleZoneReferences(element *zone, p path, pIndex int, includedElements map[int]bool) {
	zoneData, ok := engine.Patch.Zone[element.ID]
	if !ok {
		zoneData = engine.State.Zone[element.ID]
	}

	switch p[pIndex] {
	case interactablesIdentifier:
		switch v := element.Interactables[p[pIndex+1]].(type) {
		case *item:
			engine.assembleItemReferences(v, p, pIndex+2, includedElements)
			element.Interactables[p[pIndex+1]] = v
		case *player:
			engine.assemblePlayerReferences(v, p, pIndex+2, includedElements)
			element.Interactables[p[pIndex+1]] = v
		case *zoneItem:
			engine.assembleZoneItemReferences(v, p, pIndex+2, includedElements)
			element.Interactables[p[pIndex+1]] = v
		}
	case itemsIdentifier:
		child := element.Items[ZoneItemID(p[pIndex+1])]
		engine.assembleZoneItemReferences(&child, p, pIndex+2, includedElements)
		element.Items[ZoneItemID(p[pIndex+1])] = child
	case playersIdentifier:
		child := element.Players[PlayerID(p[pIndex+1])]
		engine.assemblePlayerReferences(&child, p, pIndex+2, includedElements)
		element.Players[PlayerID(p[pIndex+1])] = child
	}

	_ = zoneData
}

func (engine *Engine) assembleZoneItemReferences(element *zoneItem, p path, pIndex int, includedElements map[int]bool) {
	zoneItemData, ok := engine.Patch.ZoneItem[element.ID]
	if !ok {
		zoneItemData = engine.State.ZoneItem[element.ID]
	}

	switch p[pIndex] {
	case itemIdentifier:
	case gearScoreIdentifier:
		engine.assembleItemReferences(element.Item, p, pIndex+1, includedElements)
	case positionIdentifier:
		engine.assemblePositionReferences(element.Position, p, pIndex+1, includedElements)
	}

	_ = zoneItemData
}

func (engine *Engine) assembleEquipmentSetReferences(element *equipmentSet, p path, pIndex int, includedElements map[int]bool) {
	equipmentSetData, ok := engine.Patch.EquipmentSet[element.ID]
	if !ok {
		equipmentSetData = engine.State.EquipmentSet[element.ID]
	}

	switch p[pIndex] {
	case equipmentIdentifier:
		ref := engine.equipmentSetEquipmentRef(EquipmentSetEquipmentRefID(equipmentSetData.Equipment[p[pIndex+1]])).equipmentSetEquipmentRef
		referencedDataStatus := ReferencedDataUnchanged
		if _, ok := includedElements[int(ref.ReferencedElementID)]; ok {
			referencedDataStatus = ReferencedDataModified
		}
		referencedElement := engine.Item(ref.ReferencedElementID).item
		treeRef := elementReference{
			OperationKind:        ref.OperationKind,
			ElementID:            int(ref.ReferencedElementID),
			ElementKind:          ElementKindPlayer,
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
