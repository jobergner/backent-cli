package statemachine

type Tree struct {
	Player    map[PlayerID]_player
	Zone      map[ZoneID]_zone
	ZoneItem  map[ZoneItemID]_zoneItem
	Item      map[ItemID]_item
	Position  map[PositionID]_position
	GearScore map[GearScoreID]_gearScore
}

type _zoneItem struct {
	ID            ZoneItemID
	Position      *_position
	Item          *_item
	OperationKind OperationKind
}

type _item struct {
	ID            ItemID
	GearScore     *_gearScore
	OperationKind OperationKind
}

type _position struct {
	ID            PositionID
	X             float64
	Y             float64
	OperationKind OperationKind
}

type _gearScore struct {
	ID            GearScoreID
	Level         int
	Score         int
	OperationKind OperationKind
}

type _player struct {
	ID            PlayerID
	Items         []_item
	GearScore     *_gearScore
	Position      *_position
	OperationKind OperationKind
}

type _zone struct {
	ID            ZoneID
	Players       []_player
	Items         []_zoneItem
	OperationKind OperationKind
}
