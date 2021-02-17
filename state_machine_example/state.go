package statemachine

const (
	entityKindPlayer    entityKind = "player"
	entityKindZone                 = "zone"
	entityKindZoneItem             = "zoneItem"
	entityKindPosition             = "position"
	entityKindItem                 = "item"
	entityKindGearScore            = "gearScore"
)

type zoneID int
type zoneItemID int
type positionID int
type playerID int
type itemID int
type gearScoreID int

type state struct {
	player    map[playerID]player
	zone      map[zoneID]zone
	zoneItem  map[zoneItemID]zoneItem
	position  map[positionID]position
	item      map[itemID]item
	gearScore map[gearScoreID]gearScore
}

func newState() state {
	return state{
		player:    make(map[playerID]player),
		zone:      make(map[zoneID]zone),
		zoneItem:  make(map[zoneItemID]zoneItem),
		position:  make(map[positionID]position),
		item:      make(map[itemID]item),
		gearScore: make(map[gearScoreID]gearScore),
	}
}

type zone struct {
	id            zoneID
	players       []playerID
	items         []zoneItemID
	operationKind operationKind
}

type zoneItem struct {
	id            zoneItemID
	position      positionID
	item          itemID
	operationKind operationKind
	parentage     parentage
}

type item struct {
	id            itemID
	gearScore     gearScoreID
	operationKind operationKind
	parentage     parentage
}

type player struct {
	id            playerID
	items         []itemID
	gearScore     gearScoreID
	position      positionID
	operationKind operationKind
	parentage     parentage
}

type gearScore struct {
	id            gearScoreID
	level         int
	score         int
	operationKind operationKind
	parentage     parentage
}

type position struct {
	id            positionID
	x             float64
	y             float64
	operationKind operationKind
	parentage     parentage
}
