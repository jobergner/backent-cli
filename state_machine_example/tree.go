package statemachine

type Tree struct {
	Player    map[PlayerID]_Player
	Zone      map[ZoneID]_Zone
	ZoneItem  map[ZoneItemID]_ZoneItem
	Item      map[ItemID]_Item
	Position  map[PositionID]_Position
	GearScore map[GearScoreID]_GearScore
}

type _ZoneItem struct {
	ID            ZoneItemID
	Position      *_Position
	Item          *_Item
	OperationKind OperationKind
}

type _Item struct {
	ID            ItemID
	GearScore     *_GearScore
	OperationKind OperationKind
}

type _Position struct {
	ID            PositionID
	X             float64
	Y             float64
	OperationKind OperationKind
}

type _GearScore struct {
	ID            GearScoreID
	Level         int
	Score         int
	OperationKind OperationKind
}

type _Player struct {
	ID            PlayerID
	Items         []_Item
	GearScore     *_GearScore
	Position      *_Position
	OperationKind OperationKind
}

type _Zone struct {
	ID            ZoneID
	Players       []_Player
	Items         []_ZoneItem
	OperationKind OperationKind
}
