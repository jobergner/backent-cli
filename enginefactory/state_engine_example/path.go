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
