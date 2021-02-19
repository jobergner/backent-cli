package statemachine

type tree struct {
	player    map[playerID]_player
	zone      map[zoneID]_zone
	zoneItem  map[zoneItemID]_zoneItem
	item      map[itemID]_item
	position  map[positionID]_position
	gearScore map[gearScoreID]_gearScore
}

type _zoneItem struct {
	id            zoneItemID
	position      *_position
	item          *_item
	operationKind operationKind
}

type _item struct {
	id            itemID
	gearScore     *_gearScore
	operationKind operationKind
}

type _position struct {
	id            positionID
	x             float64
	y             float64
	operationKind operationKind
}

type _gearScore struct {
	id            gearScoreID
	level         int
	score         int
	operationKind operationKind
}

type _player struct {
	id            playerID
	items         []_item
	gearScore     *_gearScore
	position      *_position
	operationKind operationKind
}

type _zone struct {
	id            zoneID
	players       []_player
	items         []_zoneItem
	operationKind operationKind
}
