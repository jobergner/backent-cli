package state

type ReferencedDataStatus string

const (
	ReferencedDataModified  ReferencedDataStatus = "MODIFIED"
	ReferencedDataUnchanged ReferencedDataStatus = "UNCHANGED"
)

type ElementKind string

const (
	ElementKindEquipmentSet ElementKind = "EquipmentSet"
	ElementKindGearScore    ElementKind = "GearScore"
	ElementKindItem         ElementKind = "Item"
	ElementKindPlayer       ElementKind = "Player"
	ElementKindPosition     ElementKind = "Position"
	ElementKindZone         ElementKind = "Zone"
	ElementKindZoneItem     ElementKind = "ZoneItem"
)

type Tree struct {
	EquipmentSet map[EquipmentSetID]EquipmentSet `json:"equipmentSet"`
	GearScore    map[GearScoreID]GearScore       `json:"gearScore"`
	Item         map[ItemID]Item                 `json:"item"`
	Player       map[PlayerID]Player             `json:"player"`
	Position     map[PositionID]Position         `json:"position"`
	Zone         map[ZoneID]Zone                 `json:"zone"`
	ZoneItem     map[ZoneItemID]ZoneItem         `json:"zoneItem"`
}

func newTree() Tree {
	return Tree{
		EquipmentSet: make(map[EquipmentSetID]EquipmentSet),
		GearScore:    make(map[GearScoreID]GearScore),
		Item:         make(map[ItemID]Item),
		Player:       make(map[PlayerID]Player),
		Position:     make(map[PositionID]Position),
		Zone:         make(map[ZoneID]Zone),
		ZoneItem:     make(map[ZoneItemID]ZoneItem),
	}
}

type ElementReference struct {
	OperationKind        OperationKind        `json:"operationKind_"`
	ElementID            int                  `json:"id"`
	ElementKind          ElementKind          `json:"elementKind"`
	ReferencedDataStatus ReferencedDataStatus `json:"referencedDataStatus"`
}

type ZoneItem struct {
	ID             ZoneItemID    `json:"id"`
	Item           *Item         `json:"item"`
	Position       *Position     `json:"position"`
	OperationKind_ OperationKind `json:"operationKind_"`
}

type Item struct {
	ID             ItemID            `json:"id"`
	BoundTo        *ElementReference `json:"boundTo"`
	GearScore      *GearScore        `json:"gearScore"`
	Name           string            `json:"name"`
	OperationKind_ OperationKind     `json:"operationKind_"`
}

type EquipmentSet struct {
	ID             EquipmentSetID     `json:"id"`
	Name           string             `json:"name"`
	Equipment      []ElementReference `json:"equipment"`
	OperationKind_ OperationKind      `json:"operationKind_"`
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
	EquipmentSets  []ElementReference `json:"equipmentSets"`
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

type recursionCheck struct {
	equipmentSet map[EquipmentSetID]bool
	gearScore    map[GearScoreID]bool
	item         map[ItemID]bool
	player       map[PlayerID]bool
	position     map[PositionID]bool
	zone         map[ZoneID]bool
	zoneItem     map[ZoneItemID]bool
}

func newRecursionCheck() *recursionCheck {
	return &recursionCheck{
		equipmentSet: make(map[EquipmentSetID]bool),
		gearScore:    make(map[GearScoreID]bool),
		item:         make(map[ItemID]bool),
		player:       make(map[PlayerID]bool),
		position:     make(map[PositionID]bool),
		zone:         make(map[ZoneID]bool),
		zoneItem:     make(map[ZoneItemID]bool),
	}
}

type pathTrack struct {
	equipmentSet map[EquipmentSetID]path
	gearScore    map[GearScoreID]path
	item         map[ItemID]path
	player       map[PlayerID]path
	position     map[PositionID]path
	zone         map[ZoneID]path
	zoneItem     map[ZoneItemID]path
}

func newPathTrack() pathTrack {
	return pathTrack{
		equipmentSet: make(map[EquipmentSetID]path),
		gearScore:    make(map[GearScoreID]path),
		item:         make(map[ItemID]path),
		player:       make(map[PlayerID]path),
		position:     make(map[PositionID]path),
		zone:         make(map[ZoneID]path),
		zoneItem:     make(map[ZoneItemID]path),
	}
}
