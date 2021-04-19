package state

type Tree struct {
	GearScore map[GearScoreID]GearScore `json:"gearScore"`
	Item      map[ItemID]Item           `json:"item"`
	Player    map[PlayerID]Player       `json:"player"`
	Position  map[PositionID]Position   `json:"position"`
	Zone      map[ZoneID]Zone           `json:"zone"`
	ZoneItem  map[ZoneItemID]ZoneItem   `json:"zoneItem"`
}

func newTree() Tree {
	return Tree{
		GearScore: make(map[GearScoreID]GearScore),
		Item:      make(map[ItemID]Item),
		Player:    make(map[PlayerID]Player),
		Position:  make(map[PositionID]Position),
		Zone:      make(map[ZoneID]Zone),
		ZoneItem:  make(map[ZoneItemID]ZoneItem),
	}
}

type ElementReference struct {
	ID          int
	ElementKind ElementKind
}

type ZoneItem struct {
	ID             ZoneItemID    `json:"id"`
	Item           *Item         `json:"item"`
	Position       *Position     `json:"position"`
	OperationKind_ OperationKind `json:"operationKind_"`
}

type Item struct {
	ID             ItemID           `json:"id"`
	BoundTo        ElementReference `json:"boundTo"`
	GearScore      *GearScore       `json:"gearScore"`
	OperationKind_ OperationKind    `json:"operationKind_"`
}

type Position struct {
	ID             PositionID    `json:"id"`
	X              float64       `json:"x"`
	Y              float64       `json:"y"`
	OperationKind_ OperationKind `json:"operationKind_"`
}

type GearScore struct {
	ID             GearScoreID   `json:"id"`
	Level          int           `json:"level"`
	Score          int           `json:"score"`
	OperationKind_ OperationKind `json:"operationKind_"`
}

type Player struct {
	ID             PlayerID           `json:"id"`
	GearScore      *GearScore         `json:"gearScore"`
	GuildMembers   []ElementReference `json:"guildMembers"`
	Items          []Item             `json:"items"`
	Position       *Position          `json:"position"`
	OperationKind_ OperationKind      `json:"operationKind_"`
}

type Zone struct {
	ID             ZoneID        `json:"id"`
	Items          []ZoneItem    `json:"items"`
	Players        []Player      `json:"players"`
	Tags           []string      `json:"tags"`
	OperationKind_ OperationKind `json:"operationKind_"`
}
