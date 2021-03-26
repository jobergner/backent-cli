package state

type Tree struct {
	GearScore map[GearScoreID]tGearScore `json:"gearScore"`
	Item      map[ItemID]tItem           `json:"item"`
	Player    map[PlayerID]tPlayer       `json:"player"`
	Position  map[PositionID]tPosition   `json:"position"`
	Zone      map[ZoneID]tZone           `json:"zone"`
	ZoneItem  map[ZoneItemID]tZoneItem   `json:"zoneItem"`
}

func newTree() Tree {
	return Tree{
		GearScore: make(map[GearScoreID]tGearScore),
		Item:      make(map[ItemID]tItem),
		Player:    make(map[PlayerID]tPlayer),
		Position:  make(map[PositionID]tPosition),
		Zone:      make(map[ZoneID]tZone),
		ZoneItem:  make(map[ZoneItemID]tZoneItem),
	}
}

type tZoneItem struct {
	ID             ZoneItemID    `json:"id"`
	Item           *tItem        `json:"item"`
	Position       *tPosition    `json:"position"`
	OperationKind_ OperationKind `json:"operationKind_"`
}

type tItem struct {
	ID             ItemID        `json:"id"`
	GearScore      *tGearScore   `json:"gearScore"`
	OperationKind_ OperationKind `json:"operationKind_"`
}

type tPosition struct {
	ID             PositionID    `json:"id"`
	X              float64       `json:"x"`
	Y              float64       `json:"y"`
	OperationKind_ OperationKind `json:"operationKind_"`
}

type tGearScore struct {
	ID             GearScoreID   `json:"id"`
	Level          int           `json:"level"`
	Score          int           `json:"score"`
	OperationKind_ OperationKind `json:"operationKind_"`
}

type tPlayer struct {
	ID             PlayerID      `json:"id"`
	GearScore      *tGearScore   `json:"gearScore"`
	Items          []tItem       `json:"items"`
	Position       *tPosition    `json:"position"`
	OperationKind_ OperationKind `json:"operationKind_"`
}

type tZone struct {
	ID             ZoneID        `json:"id"`
	Items          []tZoneItem   `json:"items"`
	Players        []tPlayer     `json:"players"`
	Tags           []string      `json:"tags"`
	OperationKind_ OperationKind `json:"operationKind_"`
}
