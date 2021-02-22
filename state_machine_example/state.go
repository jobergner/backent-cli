package statemachine

const (
	EntityKindPlayer    EntityKind = "player"
	EntityKindZone                 = "zone"
	EntityKindZoneItem             = "zoneItem"
	EntityKindPosition             = "position"
	EntityKindItem                 = "item"
	EntityKindGearScore            = "gearScore"
)

type ZoneID int
type ZoneItemID int
type PositionID int
type PlayerID int
type ItemID int
type GearScoreID int

type State struct {
	Player    map[PlayerID]playerCore       `json:"player"`
	Zone      map[ZoneID]zoneCore           `json:"zone"`
	ZoneItem  map[ZoneItemID]zoneItemCore   `json:"zoneItem"`
	Position  map[PositionID]positionCore   `json:"position"`
	Item      map[ItemID]itemCore           `json:"item"`
	GearScore map[GearScoreID]gearScoreCore `json:"gearScore"`
}

func newState() State {
	return State{
		Player:    make(map[PlayerID]playerCore),
		Zone:      make(map[ZoneID]zoneCore),
		ZoneItem:  make(map[ZoneItemID]zoneItemCore),
		Position:  make(map[PositionID]positionCore),
		Item:      make(map[ItemID]itemCore),
		GearScore: make(map[GearScoreID]gearScoreCore),
	}
}

type zoneCore struct {
	ID            ZoneID        `json:"id"`
	Players       []PlayerID    `json:"players"`
	Items         []ZoneItemID  `json:"items"`
	OperationKind OperationKind `json:"operationKind"`
}

type Zone struct{ zone zoneCore }

type zoneItemCore struct {
	ID            ZoneItemID    `json:"id"`
	Position      PositionID    `json:"position"`
	Item          ItemID        `json:"item"`
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
	Items         []ItemID      `json:"items"`
	GearScore     GearScoreID   `json:"gearScore"`
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
