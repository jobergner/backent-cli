package statemachine

type Tree struct {
	Player    map[PlayerID]_player       `json:"player"`
	Zone      map[ZoneID]_zone           `json:"zone"`
	ZoneItem  map[ZoneItemID]_zoneItem   `json:"zoneItem"`
	Item      map[ItemID]_item           `json:"item"`
	Position  map[PositionID]_position   `json:"position"`
	GearScore map[GearScoreID]_gearScore `json:"gearScore"`
}

func newTree() Tree {
	return Tree{
		Player:    make(map[PlayerID]_player),
		Zone:      make(map[ZoneID]_zone),
		ZoneItem:  make(map[ZoneItemID]_zoneItem),
		Position:  make(map[PositionID]_position),
		Item:      make(map[ItemID]_item),
		GearScore: make(map[GearScoreID]_gearScore),
	}
}

type _zoneItem struct {
	ID            ZoneItemID    `json:"id"`
	Position      *_position    `json:"position"`
	Item          *_item        `json:"item"`
	OperationKind OperationKind `json:"operationKind"`
}

type _item struct {
	ID            ItemID        `json:"id"`
	GearScore     *_gearScore   `json:"gearScore"`
	OperationKind OperationKind `json:"operationKind"`
}

type _position struct {
	ID            PositionID    `json:"id"`
	X             float64       `json:"x"`
	Y             float64       `json:"y"`
	OperationKind OperationKind `json:"operationKind"`
}

type _gearScore struct {
	ID            GearScoreID   `json:"id"`
	Level         int           `json:"level"`
	Score         int           `json:"score"`
	OperationKind OperationKind `json:"operationKind"`
}

type _player struct {
	ID            PlayerID      `json:"id"`
	Items         []_item       `json:"items"`
	GearScore     *_gearScore   `json:"gearScore"`
	Position      *_position    `json:"position"`
	OperationKind OperationKind `json:"operationKind"`
}

type _zone struct {
	ID            ZoneID        `json:"id"`
	Players       []_player     `json:"players"`
	Items         []_zoneItem   `json:"items"`
	OperationKind OperationKind `json:"operationKind"`
}
