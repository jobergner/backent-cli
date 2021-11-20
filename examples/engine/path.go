package state

import "strconv"

const (
	equipmentSetIdentifier  int = -1
	gearScoreIdentifier     int = -2
	itemIdentifier          int = -3
	originIdentifier        int = -4
	playerIdentifier        int = -5
	itemsIdentifier         int = -6
	positionIdentifier      int = -7
	zoneIdentifier          int = -8
	interactablesIdentifier int = -9
	playersIdentifier       int = -10
	zoneItemIdentifier      int = -11
	boundToIdentifier       int = -12
	equipmentIdentifier     int = -13
	equipmentSetsIdentifier int = -14
	guildMembersIdentifier  int = -15
	targetIdentifier        int = -16
	targetedByIdentifier    int = -17
)

type path []int

func newPath(elementIdentifier int) path {
	return []int{elementIdentifier}
}

func (p path) equipmentSet() path {
	newPath := make([]int, len(p), len(p)+1)
	copy(newPath, p)
	newPath = append(newPath, equipmentSetIdentifier)
	return newPath
}

func (p path) items() path {
	newPath := make([]int, len(p), len(p)+1)
	copy(newPath, p)
	newPath = append(newPath, itemsIdentifier)
	return newPath
}

func (p path) player() path {
	newPath := make([]int, len(p), len(p)+1)
	copy(newPath, p)
	newPath = append(newPath, playerIdentifier)
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

func (p path) zone() path {
	newPath := make([]int, len(p), len(p)+1)
	copy(newPath, p)
	newPath = append(newPath, zoneIdentifier)
	return newPath
}

func (p path) zoneItem() path {
	newPath := make([]int, len(p), len(p)+1)
	copy(newPath, p)
	newPath = append(newPath, zoneItemIdentifier)
	return newPath
}

func (p path) origin() path {
	newPath := make([]int, len(p), len(p)+1)
	copy(newPath, p)
	newPath = append(newPath, originIdentifier)
	return newPath
}

func (p path) boundTo() path {
	newPath := make([]int, len(p), len(p)+1)
	copy(newPath, p)
	newPath = append(newPath, boundToIdentifier)
	return newPath
}

func (p path) equipment() path {
	newPath := make([]int, len(p), len(p)+1)
	copy(newPath, p)
	newPath = append(newPath, equipmentIdentifier)
	return newPath
}

func (p path) equipmentSets() path {
	newPath := make([]int, len(p), len(p)+1)
	copy(newPath, p)
	newPath = append(newPath, equipmentSetsIdentifier)
	return newPath
}

func (p path) guildMembers() path {
	newPath := make([]int, len(p), len(p)+1)
	copy(newPath, p)
	newPath = append(newPath, guildMembersIdentifier)
	return newPath
}

func (p path) target() path {
	newPath := make([]int, len(p), len(p)+1)
	copy(newPath, p)
	newPath = append(newPath, targetIdentifier)
	return newPath
}

func (p path) targetedBy() path {
	newPath := make([]int, len(p), len(p)+1)
	copy(newPath, p)
	newPath = append(newPath, targetedByIdentifier)
	return newPath
}

func (p path) id(id int) path {
	newPath := make([]int, len(p), len(p)+1)
	copy(newPath, p)
	newPath = append(newPath, id)
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
	case equipmentSetIdentifier:
		return "equipmentSet"
	case gearScoreIdentifier:
		return "gearScore"
	case itemIdentifier:
		return "item"
	case originIdentifier:
		return "origin"
	case playerIdentifier:
		return "player"
	case itemsIdentifier:
		return "items"
	case positionIdentifier:
		return "position"
	case zoneIdentifier:
		return "zone"
	case interactablesIdentifier:
		return "interactables"
	case playersIdentifier:
		return "players"
	case zoneItemIdentifier:
		return "zoneItem"
	}
	return ""
}
