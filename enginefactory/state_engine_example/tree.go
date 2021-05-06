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

type treeElement interface {
	kind() ElementKind
}

type ZoneItem struct {
	ID             ZoneItemID    `json:"id"`
	Item           *Item         `json:"item"`
	Position       *Position     `json:"position"`
	OperationKind_ OperationKind `json:"operationKind_"`
}
type ZoneItemReference struct {
	OperationKind        OperationKind        `json:"operationKind_"`
	ElementID            ZoneItemID           `json:"id"`
	ElementKind          ElementKind          `json:"elementKind"`
	ReferencedDataStatus ReferencedDataStatus `json:"referencedDataStatus"`
	ElementPath          string               `json:"elementPath"`
	ZoneItem             *ZoneItem            `json:"zoneItem"`
}

func (zoneItem ZoneItem) kind() ElementKind {
	return ElementKindZoneItem
}

type Item struct {
	ID             ItemID           `json:"id"`
	BoundTo        *PlayerReference `json:"boundTo"`
	GearScore      *GearScore       `json:"gearScore"`
	Name           string           `json:"name"`
	Origin         treeElement      `json:"origin"`
	OperationKind_ OperationKind    `json:"operationKind_"`
}
type ItemReference struct {
	OperationKind        OperationKind        `json:"operationKind_"`
	ElementID            ItemID               `json:"id"`
	ElementKind          ElementKind          `json:"elementKind"`
	ReferencedDataStatus ReferencedDataStatus `json:"referencedDataStatus"`
	ElementPath          string               `json:"elementPath"`
	Item                 *Item                `json:"item"`
}

func (item Item) kind() ElementKind {
	return ElementKindItem
}

type EquipmentSet struct {
	ID             EquipmentSetID  `json:"id"`
	Name           string          `json:"name"`
	Equipment      []ItemReference `json:"equipment"`
	OperationKind_ OperationKind   `json:"operationKind_"`
}
type EquipmentSetReference struct {
	OperationKind        OperationKind        `json:"operationKind_"`
	ElementID            EquipmentSetID       `json:"id"`
	ElementKind          ElementKind          `json:"elementKind"`
	ReferencedDataStatus ReferencedDataStatus `json:"referencedDataStatus"`
	ElementPath          string               `json:"elementPath"`
	EquipmentSet         *EquipmentSet        `json:"equipmentSet"`
}

func (equipmentSet EquipmentSet) kind() ElementKind {
	return ElementKindEquipmentSet
}

type Position struct {
	ID             PositionID    `json:"id"`
	X              float64       `json:"x"`
	Y              float64       `json:"y"`
	OperationKind_ OperationKind `json:"operationKind_"`
}
type PositionReference struct {
	OperationKind        OperationKind        `json:"operationKind_"`
	ElementID            PositionID           `json:"id"`
	ElementKind          ElementKind          `json:"elementKind"`
	ReferencedDataStatus ReferencedDataStatus `json:"referencedDataStatus"`
	ElementPath          string               `json:"elementPath"`
	Position             *Position            `json:"position"`
}

func (position Position) kind() ElementKind {
	return ElementKindPosition
}

type GearScore struct {
	ID             GearScoreID   `json:"id"`
	Level          int           `json:"level"`
	Score          int           `json:"score"`
	OperationKind_ OperationKind `json:"operationKind_"`
}
type GearScoreReference struct {
	OperationKind        OperationKind        `json:"operationKind_"`
	ElementID            GearScoreID          `json:"id"`
	ElementKind          ElementKind          `json:"elementKind"`
	ReferencedDataStatus ReferencedDataStatus `json:"referencedDataStatus"`
	ElementPath          string               `json:"elementPath"`
	GearScore            *GearScore           `json:"gearScore"`
}

func (gearScore GearScore) kind() ElementKind {
	return ElementKindGearScore
}

type Player struct {
	ID             PlayerID                       `json:"id"`
	EquipmentSets  []EquipmentSetReference        `json:"equipmentSets"`
	GearScore      *GearScore                     `json:"gearScore"`
	GuildMembers   []PlayerReference              `json:"guildMembers"`
	Items          []Item                         `json:"items"`
	Position       *Position                      `json:"position"`
	Target         *AnyOfPlayerZoneItemReference  `json:"target"`
	TargetedBy     []AnyOfPlayerZoneItemReference `json:"targetedBy"`
	OperationKind_ OperationKind                  `json:"operationKind_"`
}
type PlayerReference struct {
	OperationKind        OperationKind        `json:"operationKind_"`
	ElementID            PlayerID             `json:"id"`
	ElementKind          ElementKind          `json:"elementKind"`
	ReferencedDataStatus ReferencedDataStatus `json:"referencedDataStatus"`
	ElementPath          string               `json:"elementPath"`
	Player               *Player              `json:"player"`
}

func (player Player) kind() ElementKind {
	return ElementKindPlayer
}

type Zone struct {
	ID             ZoneID        `json:"id"`
	Interactables  []treeElement `json:"interactables"`
	Items          []ZoneItem    `json:"items"`
	Players        []Player      `json:"players"`
	Tags           []string      `json:"tags"`
	OperationKind_ OperationKind `json:"operationKind_"`
}
type ZoneReference struct {
	OperationKind        OperationKind        `json:"operationKind_"`
	ElementID            ZoneID               `json:"id"`
	ElementKind          ElementKind          `json:"elementKind"`
	ReferencedDataStatus ReferencedDataStatus `json:"referencedDataStatus"`
	ElementPath          string               `json:"elementPath"`
	Zone                 *Zone                `json:"zone"`
}

func (zone Zone) kind() ElementKind {
	return ElementKindZone
}

type AnyOfPlayerZoneItemReference struct {
	OperationKind        OperationKind        `json:"operationKind_"`
	ElementID            int                  `json:"id"`
	ElementKind          ElementKind          `json:"elementKind"`
	ReferencedDataStatus ReferencedDataStatus `json:"referencedDataStatus"`
	ElementPath          string               `json:"elementPath"`
	Element              treeElement          `json:"element"`
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
