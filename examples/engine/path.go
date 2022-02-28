package state

import "strconv"

type treeFieldIdentifier string

const (
	boolValueIdentifier   treeFieldIdentifier = "boolValue"
	intValueIdentifier    treeFieldIdentifier = "intValue"
	floatValueIdentifier  treeFieldIdentifier = "floatValue"
	stringValueIdentifier treeFieldIdentifier = "stringValue"

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
	equipmentSet_nameIdentifier      treeFieldIdentifier = "equipmentSet_name"

	gearScore_levelIdentifier treeFieldIdentifier = "gearScore_level"
	gearScore_scoreIdentifier treeFieldIdentifier = "gearScore_score"

	item_boundToIdentifier   treeFieldIdentifier = "item_boundTo"
	item_gearScoreIdentifier treeFieldIdentifier = "item_gearScore"
	item_nameIdenfitier      treeFieldIdentifier = "item_name"
	item_originIdentifier    treeFieldIdentifier = "item_origin"

	player_actionIdentifier        treeFieldIdentifier = "player_action"
	player_equipmentSetsIdentifier treeFieldIdentifier = "player_equipmentSets"
	player_gearScoreIdentifier     treeFieldIdentifier = "player_gearScore"
	player_guildMembersIdentifier  treeFieldIdentifier = "player_guildMembers"
	player_itemsIdentifier         treeFieldIdentifier = "player_items"
	player_positionIdentifier      treeFieldIdentifier = "player_position"
	player_targetIdentifier        treeFieldIdentifier = "player_target"
	player_targetedByIdentifier    treeFieldIdentifier = "player_targetedBy"

	position_xIdentifier treeFieldIdentifier = "position_x"
	position_yIdentifier treeFieldIdentifier = "position_y"

	zone_interactablesIdentifier treeFieldIdentifier = "zone_interactables"
	zone_itemsIdentifier         treeFieldIdentifier = "zone_items"
	zone_playersIdentifier       treeFieldIdentifier = "zone_players"
	zone_tagsIdentifier          treeFieldIdentifier = "zone_tags"

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
		jsonPath += "." + string(seg.identifier)
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
