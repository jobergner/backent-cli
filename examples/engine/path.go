package state

import "strconv"

type treeFieldIdentifier string

const (
	attackEventIdentifier  treeFieldIdentifier = "attackEvent"
	equipmentSetIdentifier treeFieldIdentifier = "equipmentSet"
	gearScoreIdentifier    treeFieldIdentifier = "gearScore"
	itemIdentifier         treeFieldIdentifier = "item"
	playerIdentifier       treeFieldIdentifier = "player"
	positionIdentifier     treeFieldIdentifier = "position"
	zoneIdentifier         treeFieldIdentifier = "zone"
	zoneItemIdentifier     treeFieldIdentifier = "zoneItem"

	attackEvent_targetIdentifier treeFieldIdentifier = "attackEvent_target"

	equipmentSet_equipmentIdentifier treeFieldIdentifier = "equipmentSet_equipment"

	item_boundToIdentifier   treeFieldIdentifier = "item_boundTo"
	item_gearScoreIdentifier treeFieldIdentifier = "item_gearScore"
	item_originIdentifier    treeFieldIdentifier = "item_origin"

	player_actionIdentifier        treeFieldIdentifier = "player_action"
	player_equipmentSetsIdentifier treeFieldIdentifier = "player_equipmentSets"
	player_gearScoreIdentifier     treeFieldIdentifier = "player_gearScore"
	player_guildMembersIdentifier  treeFieldIdentifier = "player_guildMembers"
	player_itemsIdentifier         treeFieldIdentifier = "player_items"
	player_positionIdentifier      treeFieldIdentifier = "player_position"
	player_targetIdentifier        treeFieldIdentifier = "player_target"
	player_targetedByIdentifier    treeFieldIdentifier = "player_targetedBy"

	zone_interactablesIdentifier treeFieldIdentifier = "zone_interactables"
	zone_itemsIdentifier         treeFieldIdentifier = "zone_items"
	zone_playersIdentifier       treeFieldIdentifier = "zone_players"

	zoneItem_itemIdentifier     treeFieldIdentifier = "zoneItem_item"
	zoneItem_positionIdentifier treeFieldIdentifier = "zoneItem_position"
)

type segment struct {
	id         int
	identifier treeFieldIdentifier
	kind       ElementKind
	refID      int
}

type path []segment

func newPath() path {
	return make(path, 0)
}

func (p path) extendAndCopy(fieldIdentifier treeFieldIdentifier, id int, kind ElementKind, refID int) path {
	newPath := make(path, len(p), len(p)+1)
	copy(newPath, p)
	newPath = append(newPath, segment{id, fieldIdentifier, kind, refID})
	return newPath
}

func (p path) toJSONPath() string {
	jsonPath := "$"

	for _, seg := range p {
		jsonPath += "." + pathIdentifierToString(seg.identifier)
		if isSliceFieldIdentifier(seg.identifier) {
			jsonPath += "[" + strconv.Itoa(seg.id) + "]"
		}
	}

	return jsonPath
}

func pathIdentifierToString(fieldIdentifier treeFieldIdentifier) string {
	switch fieldIdentifier {
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
	case item_boundToIdentifier:
		return "boundTo"
	case item_gearScoreIdentifier:
		return "gearScore"
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
	case zone_interactablesIdentifier:
		return "interactables"
	case zone_itemsIdentifier:
		return "items"
	case zone_playersIdentifier:
		return "players"
	case zoneItem_itemIdentifier:
		return "item"
	case zoneItem_positionIdentifier:
		return "position"
	}
	return ""
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
