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
	Player    map[PlayerID]PlayerCore
	Zone      map[ZoneID]ZoneCore
	ZoneItem  map[ZoneItemID]ZoneItemCore
	Position  map[PositionID]PositionCore
	Item      map[ItemID]ItemCore
	GearScore map[GearScoreID]GearScoreCore
}

func newState() State {
	return State{
		Player:    make(map[PlayerID]PlayerCore),
		Zone:      make(map[ZoneID]ZoneCore),
		ZoneItem:  make(map[ZoneItemID]ZoneItemCore),
		Position:  make(map[PositionID]PositionCore),
		Item:      make(map[ItemID]ItemCore),
		GearScore: make(map[GearScoreID]GearScoreCore),
	}
}

type ZoneCore struct {
	ID            ZoneID
	Players       []PlayerID
	Items         []ZoneItemID
	OperationKind OperationKind
}

type Zone struct{ zone ZoneCore }

type ZoneItemCore struct {
	ID            ZoneItemID
	Position      PositionID
	Item          ItemID
	OperationKind OperationKind
	Parentage     Parentage
}

type ZoneItem struct{ zoneItem ZoneItemCore }

type ItemCore struct {
	ID            ItemID
	GearScore     GearScoreID
	OperationKind OperationKind
	Parentage     Parentage
}

type Item struct{ item ItemCore }

type PlayerCore struct {
	ID            PlayerID
	Items         []ItemID
	GearScore     GearScoreID
	Position      PositionID
	OperationKind OperationKind
	Parentage     Parentage
}

type Player struct{ player PlayerCore }

type GearScoreCore struct {
	ID            GearScoreID
	Level         int
	Score         int
	OperationKind OperationKind
	Parentage     Parentage
}

type GearScore struct{ gearScore GearScoreCore }

type PositionCore struct {
	ID            PositionID
	X             float64
	Y             float64
	OperationKind OperationKind
	Parentage     Parentage
}

type Position struct{ position PositionCore }
