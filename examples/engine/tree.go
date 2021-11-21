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
	EquipmentSet map[EquipmentSetID]equipmentSet `json:"equipmentSet"`
	GearScore    map[GearScoreID]gearScore       `json:"gearScore"`
	Item         map[ItemID]item                 `json:"item"`
	Player       map[PlayerID]player             `json:"player"`
	Position     map[PositionID]position         `json:"position"`
	Zone         map[ZoneID]zone                 `json:"zone"`
	ZoneItem     map[ZoneItemID]zoneItem         `json:"zoneItem"`
}

func newTree() Tree {
	return Tree{
		EquipmentSet: make(map[EquipmentSetID]equipmentSet),
		GearScore:    make(map[GearScoreID]gearScore),
		Item:         make(map[ItemID]item),
		Player:       make(map[PlayerID]player),
		Position:     make(map[PositionID]position),
		Zone:         make(map[ZoneID]zone),
		ZoneItem:     make(map[ZoneItemID]zoneItem),
	}
}

type zoneItem struct {
	ID            ZoneItemID    `json:"id"`
	Item          *item         `json:"item"`
	Position      *position     `json:"position"`
	OperationKind OperationKind `json:"operationKind"`
}

type item struct {
	ID            ItemID            `json:"id"`
	BoundTo       *elementReference `json:"boundTo"`
	GearScore     *gearScore        `json:"gearScore"`
	Name          string            `json:"name"`
	Origin        interface{}       `json:"origin"`
	OperationKind OperationKind     `json:"operationKind"`
}

type equipmentSet struct {
	ID            EquipmentSetID              `json:"id"`
	Equipment     map[ItemID]elementReference `json:"equipment"`
	Name          string                      `json:"name"`
	OperationKind OperationKind               `json:"operationKind"`
}

type position struct {
	ID            PositionID    `json:"id"`
	X             float64       `json:"x"`
	Y             float64       `json:"y"`
	OperationKind OperationKind `json:"operationKind"`
}

type gearScore struct {
	ID            GearScoreID   `json:"id"`
	Level         int           `json:"level"`
	Score         int           `json:"score"`
	OperationKind OperationKind `json:"operationKind"`
}

type player struct {
	ID            PlayerID                            `json:"id"`
	EquipmentSets map[EquipmentSetID]elementReference `json:"equipmentSets"`
	GearScore     *gearScore                          `json:"gearScore"`
	GuildMembers  map[PlayerID]elementReference       `json:"guildMembers"`
	Items         map[ItemID]item                     `json:"items"`
	Position      *position                           `json:"position"`
	Target        *elementReference                   `json:"target"`
	TargetedBy    map[int]elementReference            `json:"targetedBy"`
	OperationKind OperationKind                       `json:"operationKind"`
}

type zone struct {
	ID            ZoneID                  `json:"id"`
	Interactables map[int]interface{}     `json:"interactables"`
	Items         map[ZoneItemID]zoneItem `json:"items"`
	Players       map[PlayerID]player     `json:"players"`
	Tags          []string                `json:"tags"`
	OperationKind OperationKind           `json:"operationKind"`
}

type elementReference struct {
	OperationKind        OperationKind        `json:"operationKind"`
	ElementID            int                  `json:"id"`
	ElementKind          ElementKind          `json:"elementKind"`
	ReferencedDataStatus ReferencedDataStatus `json:"referencedDataStatus"`
	ElementPath          string               `json:"elementPath"`
}

type dataAggregator struct {
	equipmentSet map[EquipmentSetID]equipmentSetCore
	gearScore    map[GearScoreID]gearScoreCore
	item         map[ItemID]itemCore
	player       map[PlayerID]playerCore
	position     map[PositionID]positionCore
	zone         map[ZoneID]zoneCore
	zoneItem     map[ZoneItemID]zoneItemCore
}

func newDataAggregator() *dataAggregator {
	return &dataAggregator{
		equipmentSet: make(map[EquipmentSetID]equipmentSetCore),
		gearScore:    make(map[GearScoreID]gearScoreCore),
		item:         make(map[ItemID]itemCore),
		player:       make(map[PlayerID]playerCore),
		position:     make(map[PositionID]positionCore),
		zone:         make(map[ZoneID]zoneCore),
		zoneItem:     make(map[ZoneItemID]zoneItemCore),
	}
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
	equipmentSet equipmentSet
}
type gearScoreCacheElement struct {
	hasUpdated bool
	gearScore  gearScore
}
type itemCacheElement struct {
	hasUpdated bool
	item       item
}
type playerCacheElement struct {
	hasUpdated bool
	player     player
}
type positionCacheElement struct {
	hasUpdated bool
	position   position
}
type zoneCacheElement struct {
	hasUpdated bool
	zone       zone
}
type zoneItemCacheElement struct {
	hasUpdated bool
	zoneItem   zoneItem
}
