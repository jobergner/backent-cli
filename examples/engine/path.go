package state

import (
	"fmt"
	"strconv"
)

type treeFieldIdentifier int

// TODO: consider IOTA
const (
	attackEventIdentifier  treeFieldIdentifier = 1
	equipmentSetIdentifier treeFieldIdentifier = 2
	gearScoreIdentifier    treeFieldIdentifier = 3
	itemIdentifier         treeFieldIdentifier = 4
	playerIdentifier       treeFieldIdentifier = 5
	positionIdentifier     treeFieldIdentifier = 6
	zoneIdentifier         treeFieldIdentifier = 7
	zoneItemIdentifier     treeFieldIdentifier = 8

	attackEvent_targetIdentifier treeFieldIdentifier = 9

	equipmentSet_equipmentIdentifier treeFieldIdentifier = 10
	equipmentSet_nameIdentifier      treeFieldIdentifier = 11

	gearScore_levelIdentifier treeFieldIdentifier = 12
	gearScore_scoreIdentifier treeFieldIdentifier = 13

	item_boundToIdentifier   treeFieldIdentifier = 14
	item_gearScoreIdentifier treeFieldIdentifier = 15
	item_nameIdenfitier      treeFieldIdentifier = 16
	item_originIdentifier    treeFieldIdentifier = 17

	player_actionIdentifier        treeFieldIdentifier = 18
	player_equipmentSetsIdentifier treeFieldIdentifier = 19
	player_gearScoreIdentifier     treeFieldIdentifier = 20
	player_guildMembersIdentifier  treeFieldIdentifier = 21
	player_itemsIdentifier         treeFieldIdentifier = 22
	player_positionIdentifier      treeFieldIdentifier = 23
	player_targetIdentifier        treeFieldIdentifier = 24
	player_targetedByIdentifier    treeFieldIdentifier = 25

	position_xIdentifier treeFieldIdentifier = 26
	position_yIdentifier treeFieldIdentifier = 27

	zone_interactablesIdentifier treeFieldIdentifier = 28
	zone_itemsIdentifier         treeFieldIdentifier = 29
	zone_playersIdentifier       treeFieldIdentifier = 30
	zone_tagsIdentifier          treeFieldIdentifier = 31

	zoneItem_itemIdentifier     treeFieldIdentifier = 32
	zoneItem_positionIdentifier treeFieldIdentifier = 33
)

func (t treeFieldIdentifier) toString() string {
	switch t {
	case attackEventIdentifier:
		return "attackEvent"
	case equipmentSetIdentifier:
		return "equipmentSet"
	case gearScoreIdentifier:
		return "gearScore"
	case itemIdentifier:
		return "item"
	case playerIdentifier:
		return "player"
	case positionIdentifier:
		return "position"
	case zoneIdentifier:
		return "zone"
	case zoneItemIdentifier:
		return "zoneItem"
	case attackEvent_targetIdentifier:
		return "target"
	case equipmentSet_equipmentIdentifier:
		return "equipment"
	case equipmentSet_nameIdentifier:
		return "name"
	case gearScore_levelIdentifier:
		return "level"
	case gearScore_scoreIdentifier:
		return "score"
	case item_boundToIdentifier:
		return "boundTo"
	case item_gearScoreIdentifier:
		return "gearScore"
	case item_nameIdenfitier:
		return "name"
	case item_originIdentifier:
		return "origin"
	case player_actionIdentifier:
		return "action"
	case player_equipmentSetsIdentifier:
		return "equipmentSets"
	case player_gearScoreIdentifier:
		return "gearScore"
	case player_guildMembersIdentifier:
		return "guildMembers"
	case player_itemsIdentifier:
		return "items"
	case player_positionIdentifier:
		return "position"
	case player_targetIdentifier:
		return "target"
	case player_targetedByIdentifier:
		return "targetedBy"
	case position_xIdentifier:
		return "x"
	case position_yIdentifier:
		return "y"
	case zone_interactablesIdentifier:
		return "interactables"
	case zone_itemsIdentifier:
		return "items"
	case zone_playersIdentifier:
		return "players"
	case zone_tagsIdentifier:
		return "tags"
	case zoneItem_itemIdentifier:
		return "item"
	case zoneItem_positionIdentifier:
		return "position"
	}

	panic(fmt.Sprintf("no string found for identifier <%d>", t))
}

type segment struct {
	id         int // is 0 when segment is of reference
	identifier treeFieldIdentifier
	kind       ElementKind
	refID      ComplexID // is ComplexID{} if segment is of non-reference
}

type path []segment

func newPath() path {
	return make(path, 0)
}

func (p path) extendAndCopy(fieldIdentifier treeFieldIdentifier, id int, kind ElementKind, refID ComplexID) path {
	newPath := make(path, len(p), len(p)+1)
	copy(newPath, p)
	newPath = append(newPath, segment{id, fieldIdentifier, kind, refID})
	return newPath
}

func (p path) toJSONPath() string {
	jsonPath := "$"

	for _, seg := range p {
		// first segment is always a name of a type
		jsonPath += "." + seg.identifier.toString()
		if isSliceFieldIdentifier(seg.identifier) {
			jsonPath += "[" + strconv.Itoa(seg.id) + "]"
		}
	}

	return jsonPath
}

func isSliceFieldIdentifier(fieldIdentifier treeFieldIdentifier) bool {
	switch fieldIdentifier {
	case attackEventIdentifier:
		return true
	case equipmentSetIdentifier:
		return true
	case gearScoreIdentifier:
		return true
	case itemIdentifier:
		return true
	case playerIdentifier:
		return true
	case positionIdentifier:
		return true
	case zoneIdentifier:
		return true
	case zoneItemIdentifier:
		return true
	case equipmentSet_equipmentIdentifier:
		return true
	case player_actionIdentifier:
		return true
	case player_equipmentSetsIdentifier:
		return true
	case player_guildMembersIdentifier:
		return true
	case player_itemsIdentifier:
		return true
	case player_targetedByIdentifier:
		return true
	case zone_interactablesIdentifier:
		return true
	case zone_itemsIdentifier:
		return true
	case zone_playersIdentifier:
		return true
	}
	return false
}
