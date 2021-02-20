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
	Player    map[PlayerID]Player
	Zone      map[ZoneID]Zone
	ZoneItem  map[ZoneItemID]ZoneItem
	Position  map[PositionID]Position
	Item      map[ItemID]Item
	GearScore map[GearScoreID]GearScore
}

func newState() State {
	return State{
		Player:    make(map[PlayerID]Player),
		Zone:      make(map[ZoneID]Zone),
		ZoneItem:  make(map[ZoneItemID]ZoneItem),
		Position:  make(map[PositionID]Position),
		Item:      make(map[ItemID]Item),
		GearScore: make(map[GearScoreID]GearScore),
	}
}

type Zone struct {
	ID            ZoneID
	Players       []PlayerID
	Items         []ZoneItemID
	OperationKind OperationKind
}

type ZoneItem struct {
	ID            ZoneItemID
	Position      PositionID
	Item          ItemID
	OperationKind OperationKind
	Parentage     Parentage
}

type Item struct {
	ID            ItemID
	GearScore     GearScoreID
	OperationKind OperationKind
	Parentage     Parentage
}

type Player struct {
	ID            PlayerID
	Items         []ItemID
	GearScore     GearScoreID
	Position      PositionID
	OperationKind OperationKind
	Parentage     Parentage
}

type GearScore struct {
	ID            GearScoreID
	Level         int
	Score         int
	OperationKind OperationKind
	Parentage     Parentage
}

type Position struct {
	ID            PositionID
	X             float64
	Y             float64
	OperationKind OperationKind
	Parentage     Parentage
}
