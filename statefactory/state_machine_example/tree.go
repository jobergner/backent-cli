package statemachine

type Tree struct {
	GearScore map[GearScoreID]_gearScore `json:"gearScore"`
	Item      map[ItemID]_item           `json:"item"`
	Player    map[PlayerID]_player       `json:"player"`
	Position  map[PositionID]_position   `json:"position"`
	Zone      map[ZoneID]_zone           `json:"zone"`
	ZoneItem  map[ZoneItemID]_zoneItem   `json:"zoneItem"`
}

func newTree() Tree {
	return Tree{
		GearScore: make(map[GearScoreID]_gearScore),
		Item:      make(map[ItemID]_item),
		Player:    make(map[PlayerID]_player),
		Position:  make(map[PositionID]_position),
		Zone:      make(map[ZoneID]_zone),
		ZoneItem:  make(map[ZoneItemID]_zoneItem),
	}
}

type _zoneItem struct {
	ID            ZoneItemID    `json:"id"`
	Item          *_item        `json:"item"`
	Position      *_position    `json:"position"`
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
	GearScore     *_gearScore   `json:"gearScore"`
	Items         []_item       `json:"items"`
	Position      *_position    `json:"position"`
	OperationKind OperationKind `json:"operationKind"`
}

type _zone struct {
	ID            ZoneID        `json:"id"`
	Items         []_zoneItem   `json:"items"`
	Players       []_player     `json:"players"`
	Tags          []string      `json:"tags"`
	OperationKind OperationKind `json:"operationKind"`
}
