package state

type ReferencedDataStatus string

const (
	ReferencedDataModified  ReferencedDataStatus = "MODIFIED"
	ReferencedDataUnchanged ReferencedDataStatus = "UNCHANGED"
)

type ElementKind string

const (
	ElementKindAttackEvent  ElementKind = "AttackEvent"
	ElementKindEquipmentSet ElementKind = "EquipmentSet"
	ElementKindGearScore    ElementKind = "GearScore"
	ElementKindItem         ElementKind = "Item"
	ElementKindPlayer       ElementKind = "Player"
	ElementKindPosition     ElementKind = "Position"
	ElementKindZone         ElementKind = "Zone"
	ElementKindZoneItem     ElementKind = "ZoneItem"
)

type Tree struct {
	AttackEvent  map[AttackEventID]attackEvent   `json:"attackEvent"`
	EquipmentSet map[EquipmentSetID]equipmentSet `json:"equipmentSet"`
	GearScore    map[GearScoreID]gearScore       `json:"gearScore"`
	Item         map[ItemID]item                 `json:"item"`
	Player       map[PlayerID]player             `json:"player"`
	Position     map[PositionID]position         `json:"position"`
	Zone         map[ZoneID]zone                 `json:"zone"`
	ZoneItem     map[ZoneItemID]zoneItem         `json:"zoneItem"`
}

func newTree() *Tree {
	return &Tree{
		AttackEvent:  make(map[AttackEventID]attackEvent),
		EquipmentSet: make(map[EquipmentSetID]equipmentSet),
		GearScore:    make(map[GearScoreID]gearScore),
		Item:         make(map[ItemID]item),
		Player:       make(map[PlayerID]player),
		Position:     make(map[PositionID]position),
		Zone:         make(map[ZoneID]zone),
		ZoneItem:     make(map[ZoneItemID]zoneItem),
	}
}

func (t *Tree) clear() {
	for key := range t.AttackEvent {
		delete(t.AttackEvent, key)
	}
	for key := range t.EquipmentSet {
		delete(t.EquipmentSet, key)
	}
	for key := range t.GearScore {
		delete(t.GearScore, key)
	}
	for key := range t.Item {
		delete(t.Item, key)
	}
	for key := range t.Player {
		delete(t.Player, key)
	}
	for key := range t.Position {
		delete(t.Position, key)
	}
	for key := range t.Zone {
		delete(t.Zone, key)
	}
	for key := range t.ZoneItem {
		delete(t.ZoneItem, key)
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

type attackEvent struct {
	ID            AttackEventID     `json:"id"`
	Target        *elementReference `json:"target"`
	Name          string            `json:"name"`
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
