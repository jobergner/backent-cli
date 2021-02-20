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
	Player    map[PlayerID]playerCore
	Zone      map[ZoneID]zoneCore
	ZoneItem  map[ZoneItemID]zoneItemCore
	Position  map[PositionID]positionCore
	Item      map[ItemID]itemCore
	GearScore map[GearScoreID]gearScoreCore
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
	ID            ZoneID
	Players       []PlayerID
	Items         []ZoneItemID
	OperationKind OperationKind
}

type Zone struct{ zone zoneCore }

type zoneItemCore struct {
	ID            ZoneItemID
	Position      PositionID
	Item          ItemID
	OperationKind OperationKind
	Parentage     Parentage
}

type ZoneItem struct{ zoneItem zoneItemCore }

type itemCore struct {
	ID            ItemID
	GearScore     GearScoreID
	OperationKind OperationKind
	Parentage     Parentage
}

type Item struct{ item itemCore }

type playerCore struct {
	ID            PlayerID
	Items         []ItemID
	GearScore     GearScoreID
	Position      PositionID
	OperationKind OperationKind
	Parentage     Parentage
}

type Player struct{ player playerCore }

type gearScoreCore struct {
	ID            GearScoreID
	Level         int
	Score         int
	OperationKind OperationKind
	Parentage     Parentage
}

type GearScore struct{ gearScore gearScoreCore }

type positionCore struct {
	ID            PositionID
	X             float64
	Y             float64
	OperationKind OperationKind
	Parentage     Parentage
}

type Position struct{ position positionCore }
