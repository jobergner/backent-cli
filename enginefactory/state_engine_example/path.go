package state

import "strconv"

type pathTrack struct {
	_iterations  int
	equipmentSet map[EquipmentSetID]path
	gearScore    map[GearScoreID]path
	item         map[ItemID]path
	player       map[PlayerID]path
	position     map[PositionID]path
	zone         map[ZoneID]path
	zoneItem     map[ZoneItemID]path
}

func newPathTrack() pathTrack {
	return pathTrack{
		equipmentSet: make(map[EquipmentSetID]path),
		gearScore:    make(map[GearScoreID]path),
		item:         make(map[ItemID]path),
		player:       make(map[PlayerID]path),
		position:     make(map[PositionID]path),
		zone:         make(map[ZoneID]path),
		zoneItem:     make(map[ZoneItemID]path),
	}
}

const (
	itemsIdentifier         int = -1
	gearScoreIdentifier     int = -2
	positionIdentifier      int = -3
	targetIdentifier        int = -4
	playersIdentifier       int = -5
	interactablesIdentifier int = -6
	itemIdentifier          int = -7
	originIdentifier        int = -8
	equipmentSetIdentifier  int = -9
	playerIdentifier        int = -10
	zoneIdentifier          int = -11
	zoneItemIdentifier      int = -12
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
	newPath := make([]int, len(p), len(p)+1)
	copy(newPath, p)
	newPath = append(newPath, itemsIdentifier)
	return newPath
}

func (p path) gearScore() path {
	newPath := make([]int, len(p), len(p)+1)
	copy(newPath, p)
	newPath = append(newPath, gearScoreIdentifier)
	return newPath
}

func (p path) position() path {
	newPath := make([]int, len(p), len(p)+1)
	copy(newPath, p)
	newPath = append(newPath, positionIdentifier)
	return newPath
}

func (p path) target() path {
	newPath := make([]int, len(p), len(p)+1)
	copy(newPath, p)
	newPath = append(newPath, targetIdentifier)
	return newPath
}

func (p path) players() path {
	newPath := make([]int, len(p), len(p)+1)
	copy(newPath, p)
	newPath = append(newPath, playersIdentifier)
	return newPath
}

func (p path) interactables() path {
	newPath := make([]int, len(p), len(p)+1)
	copy(newPath, p)
	newPath = append(newPath, interactablesIdentifier)
	return newPath
}

func (p path) item() path {
	newPath := make([]int, len(p), len(p)+1)
	copy(newPath, p)
	newPath = append(newPath, itemIdentifier)
	return newPath
}

func (p path) origin() path {
	newPath := make([]int, len(p), len(p)+1)
	copy(newPath, p)
	newPath = append(newPath, originIdentifier)
	return newPath
}

func (p path) index(i int) path {
	newPath := make([]int, len(p), len(p)+1)
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
		if existingPath, pathExists := se.PathTrack.item[itemID]; !pathExists || !existingPath.equals(p) {
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
		if existingPath, pathExists := se.PathTrack.zoneItem[zoneItemID]; !pathExists || !existingPath.equals(p) {
			itemsPath = p.items().index(i)
		} else {
			itemsPath = existingPath
		}
		se.walkZoneItem(zoneItemID, itemsPath)
	}

	for i, playerID := range mergePlayerIDs(se.State.Zone[zoneData.ID].Players, se.Patch.Zone[zoneData.ID].Players) {
		var playersPath path
		if existingPath, pathExists := se.PathTrack.player[playerID]; !pathExists || !existingPath.equals(p) {
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

	for id, equipmentSetData := range se.Patch.EquipmentSet {
		se.walkEquipmentSet(equipmentSetData.ID, newPath(equipmentSetIdentifier, int(id)))
		walkedCheck.equipmentSet[equipmentSetData.ID] = true
	}
	for id, gearScoreData := range se.Patch.GearScore {
		if !gearScoreData.HasParent {
			se.walkGearScore(gearScoreData.ID, newPath(gearScoreIdentifier, int(id)))
			walkedCheck.gearScore[gearScoreData.ID] = true
		}
	}
	for id, itemData := range se.Patch.Item {
		if !itemData.HasParent {
			se.walkItem(itemData.ID, newPath(itemIdentifier, int(id)))
			walkedCheck.item[itemData.ID] = true
		}
	}
	for id, playerData := range se.Patch.Player {
		if !playerData.HasParent {
			se.walkPlayer(playerData.ID, newPath(playerIdentifier, int(id)))
			walkedCheck.player[playerData.ID] = true
		}
	}
	for id, positionData := range se.Patch.Position {
		if !positionData.HasParent {
			se.walkPosition(positionData.ID, newPath(positionIdentifier, int(id)))
			walkedCheck.position[positionData.ID] = true
		}
	}
	for id, zoneData := range se.Patch.Zone {
		se.walkZone(zoneData.ID, newPath(zoneIdentifier, int(id)))
		walkedCheck.zone[zoneData.ID] = true
	}
	for id, zoneItemData := range se.Patch.ZoneItem {
		if !zoneItemData.HasParent {
			se.walkZoneItem(zoneItemData.ID, newPath(zoneItemIdentifier, int(id)))
			walkedCheck.zoneItem[zoneItemData.ID] = true
		}
	}

	for id, equipmentSetData := range se.State.EquipmentSet {
		if _, ok := walkedCheck.equipmentSet[equipmentSetData.ID]; !ok {
			se.walkEquipmentSet(equipmentSetData.ID, newPath(equipmentSetIdentifier, int(id)))
		}
	}
	for id, gearScoreData := range se.State.GearScore {
		if !gearScoreData.HasParent {
			if _, ok := walkedCheck.gearScore[gearScoreData.ID]; !ok {
				se.walkGearScore(gearScoreData.ID, newPath(gearScoreIdentifier, int(id)))
			}
		}
	}
	for id, itemData := range se.State.Item {
		if !itemData.HasParent {
			if _, ok := walkedCheck.item[itemData.ID]; !ok {
				se.walkItem(itemData.ID, newPath(itemIdentifier, int(id)))
			}
		}
	}
	for id, playerData := range se.State.Player {
		if !playerData.HasParent {
			if _, ok := walkedCheck.player[playerData.ID]; !ok {
				se.walkPlayer(playerData.ID, newPath(playerIdentifier, int(id)))
			}
		}
	}
	for id, positionData := range se.State.Position {
		if !positionData.HasParent {
			if _, ok := walkedCheck.position[positionData.ID]; !ok {
				se.walkPosition(positionData.ID, newPath(positionIdentifier, int(id)))
			}
		}
	}
	for id, zoneData := range se.State.Zone {
		if _, ok := walkedCheck.zone[zoneData.ID]; !ok {
			se.walkZone(zoneData.ID, newPath(zoneIdentifier, int(id)))
		}
	}
	for id, zoneItemData := range se.State.ZoneItem {
		if !zoneItemData.HasParent {
			if _, ok := walkedCheck.zoneItem[zoneItemData.ID]; !ok {
				se.walkZoneItem(zoneItemData.ID, newPath(zoneItemIdentifier, int(id)))
			}
		}
	}

	se.PathTrack._iterations += 1
	if se.PathTrack._iterations == 100 {
		for key := range se.PathTrack.equipmentSet {
			delete(se.PathTrack.equipmentSet, key)
		}
		for key := range se.PathTrack.gearScore {
			delete(se.PathTrack.gearScore, key)
		}
		for key := range se.PathTrack.item {
			delete(se.PathTrack.item, key)
		}
		for key := range se.PathTrack.player {
			delete(se.PathTrack.player, key)
		}
		for key := range se.PathTrack.position {
			delete(se.PathTrack.position, key)
		}
		for key := range se.PathTrack.zone {
			delete(se.PathTrack.zone, key)
		}
		for key := range se.PathTrack.zoneItem {
			delete(se.PathTrack.zoneItem, key)
		}
	}
}

func (p path) toJSONPath() string {
	jsonPath := "$"

	for i, seg := range p {
		if seg < 0 {
			jsonPath += "." + pathIdentifierToString(seg)
		} else if i == 1 {
			jsonPath += "." + strconv.Itoa(seg)
		} else {
			jsonPath += "[" + strconv.Itoa(seg) + "]"
		}
	}

	return jsonPath
}

func pathIdentifierToString(identifier int) string {
	switch identifier {
	case itemsIdentifier:
		return "items"
	case gearScoreIdentifier:
		return "gearScore"
	case positionIdentifier:
		return "position"
	case targetIdentifier:
		return "target"
	case playersIdentifier:
		return "players"
	case interactablesIdentifier:
		return "interactables"
	case itemIdentifier:
		return "item"
	case originIdentifier:
		return "origin"
	case equipmentSetIdentifier:
		return "equipmentSet"
	case playerIdentifier:
		return "player"
	case zoneIdentifier:
		return "zone"
	case zoneItemIdentifier:
		return "zoneItem"
	}
	return ""
}
