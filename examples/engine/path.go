package state

import "strconv"

type treeFieldIdentifier string

const (
	attackEventIdentifier  = "attackEvent"
	equipmentSetIdentifier = "equipmentSet"
	gearScoreIdentifier    = "gearScore"
	itemIdentifier         = "item"
	playerIdentifier       = "player"
	positionIdentifier     = "position"
	zoneIdentifier         = "zone"
	zoneItemIdentifier     = "zoneItem"

	attackEvent_targetIdentifier = "attackEvent_target"

	equipmentSet_equipmentIdentifier = "equipmentSet_equipment"

	item_boundToIdentifier   = "item_boundTo"
	item_gearScoreIdentifier = "item_gearScore"
	item_originIdentifier    = "item_origin"

	player_actionIdentifier        = "player_action"
	player_equipmentSetsIdentifier = "player_equipmentSets"
	player_gearScoreIdentifier     = "player_gearScore"
	player_guildMembersIdentifier  = "player_guildMembers"
	player_itemsIdentifier         = "player_items"
	player_positionIdentifier      = "player_position"
	player_targetIdentifier        = "player_target"
	player_targetedByIdentifier    = "player_targetedBy"

	zone_interactablesIdentifier = "zone_interactables"
	zone_itemsIdentifier         = "zone_items"
	zone_playersIdentifier       = "zone_players"

	zoneItem_itemIdentifier     = "zoneItem_item"
	zoneItem_positionIdentifier = "zoneItem_position"
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
	case equipmentSet_equipmentIdentifier:
		return "equipment"
	case item_boundToIdentifier:
		return "boundTo"
	case item_gearScoreIdentifier:
		return "gearScore"
	case item_originIdentifier:
		return "origin"
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
