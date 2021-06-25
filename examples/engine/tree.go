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

type ZoneItem struct {
	ID            ZoneItemID    `json:"id"`
	Item          *Item         `json:"item"`
	Position      *Position     `json:"position"`
	OperationKind OperationKind `json:"operationKind"`
}
type ZoneItemReference struct {
	OperationKind        OperationKind        `json:"operationKind"`
	ElementID            ZoneItemID           `json:"id"`
	ElementKind          ElementKind          `json:"elementKind"`
	ReferencedDataStatus ReferencedDataStatus `json:"referencedDataStatus"`
	ElementPath          string               `json:"elementPath"`
	ZoneItem             *ZoneItem            `json:"zoneItem"`
}

type Item struct {
	ID            ItemID           `json:"id"`
	BoundTo       *PlayerReference `json:"boundTo"`
	GearScore     *GearScore       `json:"gearScore"`
	Name          string           `json:"name"`
	Origin        interface{}      `json:"origin"`
	OperationKind OperationKind    `json:"operationKind"`
}
type ItemReference struct {
	OperationKind        OperationKind        `json:"operationKind"`
	ElementID            ItemID               `json:"id"`
	ElementKind          ElementKind          `json:"elementKind"`
	ReferencedDataStatus ReferencedDataStatus `json:"referencedDataStatus"`
	ElementPath          string               `json:"elementPath"`
	Item                 *Item                `json:"item"`
}

type EquipmentSet struct {
	ID            EquipmentSetID           `json:"id"`
	Equipment     map[ItemID]ItemReference `json:"equipment"`
	Name          string                   `json:"name"`
	OperationKind OperationKind            `json:"operationKind"`
}
type EquipmentSetReference struct {
	OperationKind        OperationKind        `json:"operationKind"`
	ElementID            EquipmentSetID       `json:"id"`
	ElementKind          ElementKind          `json:"elementKind"`
	ReferencedDataStatus ReferencedDataStatus `json:"referencedDataStatus"`
	ElementPath          string               `json:"elementPath"`
	EquipmentSet         *EquipmentSet        `json:"equipmentSet"`
}

type Position struct {
	ID            PositionID    `json:"id"`
	X             float64       `json:"x"`
	Y             float64       `json:"y"`
	OperationKind OperationKind `json:"operationKind"`
}
type PositionReference struct {
	OperationKind        OperationKind        `json:"operationKind"`
	ElementID            PositionID           `json:"id"`
	ElementKind          ElementKind          `json:"elementKind"`
	ReferencedDataStatus ReferencedDataStatus `json:"referencedDataStatus"`
	ElementPath          string               `json:"elementPath"`
	Position             *Position            `json:"position"`
}

type GearScore struct {
	ID            GearScoreID   `json:"id"`
	Level         int           `json:"level"`
	Score         int           `json:"score"`
	OperationKind OperationKind `json:"operationKind"`
}
type GearScoreReference struct {
	OperationKind        OperationKind        `json:"operationKind"`
	ElementID            GearScoreID          `json:"id"`
	ElementKind          ElementKind          `json:"elementKind"`
	ReferencedDataStatus ReferencedDataStatus `json:"referencedDataStatus"`
	ElementPath          string               `json:"elementPath"`
	GearScore            *GearScore           `json:"gearScore"`
}

type Player struct {
	ID            PlayerID                                 `json:"id"`
	EquipmentSets map[EquipmentSetID]EquipmentSetReference `json:"equipmentSets"`
	GearScore     *GearScore                               `json:"gearScore"`
	GuildMembers  map[PlayerID]PlayerReference             `json:"guildMembers"`
	Items         map[ItemID]Item                          `json:"items"`
	Position      *Position                                `json:"position"`
	Target        *AnyOfPlayer_ZoneItemReference           `json:"target"`
	TargetedBy    map[int]AnyOfPlayer_ZoneItemReference    `json:"targetedBy"`
	OperationKind OperationKind                            `json:"operationKind"`
}
type PlayerReference struct {
	OperationKind        OperationKind        `json:"operationKind"`
	ElementID            PlayerID             `json:"id"`
	ElementKind          ElementKind          `json:"elementKind"`
	ReferencedDataStatus ReferencedDataStatus `json:"referencedDataStatus"`
	ElementPath          string               `json:"elementPath"`
	Player               *Player              `json:"player"`
}

type Zone struct {
	ID            ZoneID                  `json:"id"`
	Interactables map[int]interface{}     `json:"interactables"`
	Items         map[ZoneItemID]ZoneItem `json:"items"`
	Players       map[PlayerID]Player     `json:"players"`
	Tags          []string                `json:"tags"`
	OperationKind OperationKind           `json:"operationKind"`
}
type ZoneReference struct {
	OperationKind        OperationKind        `json:"operationKind"`
	ElementID            ZoneID               `json:"id"`
	ElementKind          ElementKind          `json:"elementKind"`
	ReferencedDataStatus ReferencedDataStatus `json:"referencedDataStatus"`
	ElementPath          string               `json:"elementPath"`
	Zone                 *Zone                `json:"zone"`
}

type AnyOfPlayer_ZoneItemReference struct {
	OperationKind        OperationKind        `json:"operationKind"`
	ElementID            int                  `json:"id"`
	ElementKind          ElementKind          `json:"elementKind"`
	ReferencedDataStatus ReferencedDataStatus `json:"referencedDataStatus"`
	ElementPath          string               `json:"elementPath"`
	Element              interface{}          `json:"element"`
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

type assembleCache struct {
	equipmentSet map[EquipmentSetID]equipmentSetCacheElement
	gearScore    map[GearScoreID]gearScoreCacheElement
	item         map[ItemID]itemCacheElement
	player       map[PlayerID]playerCacheElement
	position     map[PositionID]positionCacheElement
	zone         map[ZoneID]zoneCacheElement
	zoneItem     map[ZoneItemID]zoneItemCacheElement
}

func newAssembleCache() assembleCache {
	return assembleCache{
		equipmentSet: make(map[EquipmentSetID]equipmentSetCacheElement),
		gearScore:    make(map[GearScoreID]gearScoreCacheElement),
		item:         make(map[ItemID]itemCacheElement),
		player:       make(map[PlayerID]playerCacheElement),
		position:     make(map[PositionID]positionCacheElement),
		zone:         make(map[ZoneID]zoneCacheElement),
		zoneItem:     make(map[ZoneItemID]zoneItemCacheElement),
	}
}

type equipmentSetCacheElement struct {
	hasUpdated   bool
	equipmentSet EquipmentSet
}
type gearScoreCacheElement struct {
	hasUpdated bool
	gearScore  GearScore
}
type itemCacheElement struct {
	hasUpdated bool
	item       Item
}
type playerCacheElement struct {
	hasUpdated bool
	player     Player
}
type positionCacheElement struct {
	hasUpdated bool
	position   Position
}
type zoneCacheElement struct {
	hasUpdated bool
	zone       Zone
}
type zoneItemCacheElement struct {
	hasUpdated bool
	zoneItem   ZoneItem
}
