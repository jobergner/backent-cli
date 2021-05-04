package state

const (
	itemsIdentifier         int = -1
	gearScoreIdentifier     int = -2
	positionIdentifier      int = -3
	targetIdentifier        int = -4
	playersIdentifier       int = -5
	interactablesIdentifier int = -6
	itemIdentifier          int = -7
	originIdentifier        int = -8
)

type path []int

func newPath(elementIdentifier, id int) path {
	return []int{elementIdentifier, id}
}

func newEmptyPath() path {
	var p path
	return p
}

func (p path) items() path {
	newPath := make([]int, len(p)+1)
	copy(newPath, p)
	newPath = append(newPath, itemsIdentifier)
	return newPath
}

func (p path) gearScore() path {
	newPath := make([]int, len(p)+1)
	copy(newPath, p)
	newPath = append(newPath, gearScoreIdentifier)
	return newPath
}

func (p path) position() path {
	newPath := make([]int, len(p)+1)
	copy(newPath, p)
	newPath = append(newPath, positionIdentifier)
	return newPath
}

func (p path) target() path {
	newPath := make([]int, len(p)+1)
	copy(newPath, p)
	newPath = append(newPath, targetIdentifier)
	return newPath
}

func (p path) players() path {
	newPath := make([]int, len(p)+1)
	copy(newPath, p)
	newPath = append(newPath, playersIdentifier)
	return newPath
}

func (p path) interactables() path {
	newPath := make([]int, len(p)+1)
	copy(newPath, p)
	newPath = append(newPath, interactablesIdentifier)
	return newPath
}

func (p path) item() path {
	newPath := make([]int, len(p)+1)
	copy(newPath, p)
	newPath = append(newPath, itemIdentifier)
	return newPath
}

func (p path) origin() path {
	newPath := make([]int, len(p)+1)
	copy(newPath, p)
	newPath = append(newPath, originIdentifier)
	return newPath
}

func (p path) index(i int) path {
	newPath := make([]int, len(p)+1)
	copy(newPath, p)
	newPath = append(newPath, i)
	return newPath
}

func (p path) equals(parentPath path) bool {
	if len(p) != len(parentPath) {
		return false
	}

	for i, segment := range parentPath {
		if segment != p[i] {
			return false
		}
	}

	return true
}

func (se *Engine) walkGearScore(gearScoreID GearScoreID, p path) {
	se.PathTrack.gearScore[gearScoreID] = p
}

func (se *Engine) walkPosition(positionID PositionID, p path) {
	se.PathTrack.position[positionID] = p
}

func (se *Engine) walkEquipmentSet(equipmentSetID EquipmentSetID, p path) {
	se.PathTrack.equipmentSet[equipmentSetID] = p
}

func (se *Engine) walkItem(itemID ItemID, p path) {
	itemData, hasUpdated := se.Patch.Item[itemID]
	if !hasUpdated {
		itemData = se.State.Item[itemID]
	}

	var gearScorePath path
	if existingPath, pathExists := se.PathTrack.gearScore[itemData.GearScore]; !pathExists {
		gearScorePath = p.gearScore()
	} else {
		gearScorePath = existingPath
	}
	se.walkGearScore(itemData.GearScore, gearScorePath)

	se.PathTrack.item[itemID] = p
}

func (se *Engine) walkZoneItem(zoneItemID ZoneItemID, p path) {

	zoneItemData, hasUpdated := se.Patch.ZoneItem[zoneItemID]
	if !hasUpdated {
		zoneItemData = se.State.ZoneItem[zoneItemID]
	}

	var itemPath path
	if existingPath, pathExists := se.PathTrack.item[zoneItemData.Item]; !pathExists {
		itemPath = p.item()
	} else {
		itemPath = existingPath
	}
	se.walkItem(zoneItemData.Item, itemPath)

	var positionPath path
	if existingPath, pathExists := se.PathTrack.position[zoneItemData.Position]; !pathExists {
		positionPath = p.position()
	} else {
		positionPath = existingPath
	}
	se.walkPosition(zoneItemData.Position, positionPath)

	se.PathTrack.zoneItem[zoneItemID] = p
}

func (se *Engine) walkPlayer(playerID PlayerID, p path) {

	playerData, hasUpdated := se.Patch.Player[playerID]
	if !hasUpdated {
		playerData = se.State.Player[playerID]
	}

	var gearScorePath path
	if existingPath, pathExists := se.PathTrack.gearScore[playerData.GearScore]; !pathExists {
		gearScorePath = p.gearScore()
	} else {
		gearScorePath = existingPath
	}
	se.walkGearScore(playerData.GearScore, gearScorePath)

	for i, itemID := range mergeItemIDs(se.State.Player[playerData.ID].Items, se.Patch.Player[playerData.ID].Items) {
		var itemsPath path
		if existingPath, pathExists := se.PathTrack.item[itemID]; !pathExists && !existingPath.equals(p) {
			itemsPath = p.items().index(i)
		} else {
			itemsPath = existingPath
		}
		se.walkItem(itemID, itemsPath)
	}

	var positionPath path
	if existingPath, pathExists := se.PathTrack.position[playerData.Position]; !pathExists {
		positionPath = p.position()
	} else {
		positionPath = existingPath
	}
	se.walkPosition(playerData.Position, positionPath)

	se.PathTrack.player[playerID] = p
}

func (se *Engine) walkZone(zoneID ZoneID, p path) {
	zoneData, hasUpdated := se.Patch.Zone[zoneID]
	if !hasUpdated {
		zoneData = se.State.Zone[zoneID]
	}

	for i, zoneItemID := range mergeZoneItemIDs(se.State.Zone[zoneData.ID].Items, se.Patch.Zone[zoneData.ID].Items) {
		var itemsPath path
		if existingPath, pathExists := se.PathTrack.zoneItem[zoneItemID]; !pathExists && !existingPath.equals(p) {
			itemsPath = p.items().index(i)
		} else {
			itemsPath = existingPath
		}
		se.walkZoneItem(zoneItemID, itemsPath)
	}

	for i, playerID := range mergePlayerIDs(se.State.Zone[zoneData.ID].Players, se.Patch.Zone[zoneData.ID].Players) {
		var playersPath path
		if existingPath, pathExists := se.PathTrack.player[playerID]; !pathExists && !existingPath.equals(p) {
			playersPath = p.players().index(i)
		} else {
			playersPath = existingPath
		}
		se.walkPlayer(playerID, playersPath)
	}

	se.PathTrack.zone[zoneID] = p
}

func (se *Engine) walkTree() {

	walkedCheck := newRecursionCheck()

	for _, equipmentSetData := range se.Patch.EquipmentSet {
		se.walkEquipmentSet(equipmentSetData.ID, newEmptyPath())
		walkedCheck.equipmentSet[equipmentSetData.ID] = true
	}
	for _, gearScoreData := range se.Patch.GearScore {
		if !gearScoreData.HasParent_ {
			se.walkGearScore(gearScoreData.ID, newEmptyPath())
			walkedCheck.gearScore[gearScoreData.ID] = true
		}
	}
	for _, itemData := range se.Patch.Item {
		if !itemData.HasParent_ {
			se.walkItem(itemData.ID, newEmptyPath())
			walkedCheck.item[itemData.ID] = true
		}
	}
	for _, playerData := range se.Patch.Player {
		if !playerData.HasParent_ {
			se.walkPlayer(playerData.ID, newEmptyPath())
			walkedCheck.player[playerData.ID] = true
		}
	}
	for _, positionData := range se.Patch.Position {
		if !positionData.HasParent_ {
			se.walkPosition(positionData.ID, newEmptyPath())
			walkedCheck.position[positionData.ID] = true
		}
	}
	for _, zoneData := range se.Patch.Zone {
		se.walkZone(zoneData.ID, newEmptyPath())
		walkedCheck.zone[zoneData.ID] = true
	}
	for _, zoneItemData := range se.Patch.ZoneItem {
		if !zoneItemData.HasParent_ {
			se.walkZoneItem(zoneItemData.ID, newEmptyPath())
			walkedCheck.zoneItem[zoneItemData.ID] = true
		}
	}

	for _, equipmentSetData := range se.State.EquipmentSet {
		if _, ok := walkedCheck.equipmentSet[equipmentSetData.ID]; !ok {
			se.walkEquipmentSet(equipmentSetData.ID, newEmptyPath())
		}
	}
	for _, gearScoreData := range se.State.GearScore {
		if !gearScoreData.HasParent_ {
			if _, ok := walkedCheck.gearScore[gearScoreData.ID]; !ok {
				se.walkGearScore(gearScoreData.ID, newEmptyPath())
			}
		}
	}
	for _, itemData := range se.State.Item {
		if !itemData.HasParent_ {
			if _, ok := walkedCheck.item[itemData.ID]; !ok {
				se.walkItem(itemData.ID, newEmptyPath())
			}
		}
	}
	for _, playerData := range se.State.Player {
		if !playerData.HasParent_ {
			if _, ok := walkedCheck.player[playerData.ID]; !ok {
				se.walkPlayer(playerData.ID, newEmptyPath())
			}
		}
	}
	for _, positionData := range se.State.Position {
		if !positionData.HasParent_ {
			if _, ok := walkedCheck.position[positionData.ID]; !ok {
				se.walkPosition(positionData.ID, newEmptyPath())
			}
		}
	}
	for _, zoneData := range se.State.Zone {
		if _, ok := walkedCheck.zone[zoneData.ID]; !ok {
			se.walkZone(zoneData.ID, newEmptyPath())
		}
	}
	for _, zoneItemData := range se.State.ZoneItem {
		if !zoneItemData.HasParent_ {
			if _, ok := walkedCheck.zoneItem[zoneItemData.ID]; !ok {
				se.walkZoneItem(zoneItemData.ID, newEmptyPath())
			}
		}
	}

}
