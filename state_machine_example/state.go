package statemachine

type EntityKind string

const (
	EntityKindGearScore EntityKind = "gearScore"
	EntityKindItem                 = "item"
	EntityKindPlayer               = "player"
	EntityKindPosition             = "position"
	EntityKindZone                 = "zone"
	EntityKindZoneItem             = "zoneItem"
)

type GearScoreID int
type ItemID int
type PlayerID int
type PositionID int
type ZoneID int
type ZoneItemID int

type State struct {
	GearScore map[GearScoreID]gearScoreCore `json:"gearScore"`
	Item      map[ItemID]itemCore           `json:"item"`
	Player    map[PlayerID]playerCore       `json:"player"`
	Position  map[PositionID]positionCore   `json:"position"`
	Zone      map[ZoneID]zoneCore           `json:"zone"`
	ZoneItem  map[ZoneItemID]zoneItemCore   `json:"zoneItem"`
}

func newState() State {
	return State{
		GearScore: make(map[GearScoreID]gearScoreCore),
		Item:      make(map[ItemID]itemCore),
		Player:    make(map[PlayerID]playerCore),
		Position:  make(map[PositionID]positionCore),
		Zone:      make(map[ZoneID]zoneCore),
		ZoneItem:  make(map[ZoneItemID]zoneItemCore),
	}
}

type zoneCore struct {
	ID            ZoneID        `json:"id"`
	Items         []ZoneItemID  `json:"items"`
	Players       []PlayerID    `json:"players"`
	OperationKind OperationKind `json:"operationKind"`
}

type Zone struct{ zone zoneCore }

type zoneItemCore struct {
	ID            ZoneItemID    `json:"id"`
	Item          ItemID        `json:"item"`
	Position      PositionID    `json:"position"`
	OperationKind OperationKind `json:"operationKind"`
	Parentage     Parentage     `json:"parentage"`
}

type ZoneItem struct{ zoneItem zoneItemCore }

type itemCore struct {
	ID            ItemID        `json:"id"`
	GearScore     GearScoreID   `json:"gearScore"`
	OperationKind OperationKind `json:"operationKind"`
	Parentage     Parentage     `json:"parentage"`
}

type Item struct{ item itemCore }

type playerCore struct {
	ID            PlayerID      `json:"id"`
	GearScore     GearScoreID   `json:"gearScore"`
	Items         []ItemID      `json:"items"`
	Position      PositionID    `json:"position"`
	OperationKind OperationKind `json:"operationKind"`
	Parentage     Parentage     `json:"parentage"`
}

type Player struct{ player playerCore }

type gearScoreCore struct {
	ID            GearScoreID   `json:"id"`
	Level         int           `json:"level"`
	Score         int           `json:"score"`
	OperationKind OperationKind `json:"operationKind"`
	Parentage     Parentage     `json:"parentage"`
}

type GearScore struct{ gearScore gearScoreCore }

type positionCore struct {
	ID            PositionID    `json:"id"`
	X             float64       `json:"x"`
	Y             float64       `json:"y"`
	OperationKind OperationKind `json:"operationKind"`
	Parentage     Parentage     `json:"parentage"`
}

type Position struct{ position positionCore }
