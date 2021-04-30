package state

type EquipmentSetID int
type GearScoreID int
type ItemID int
type PlayerID int
type PositionID int
type ZoneID int
type ZoneItemID int
type PlayerGuildMemberRefID int
type ItemBoundToRefID int
type EquipmentSetEquipmentRefID int
type PlayerEquipmentSetRefID int
type AnyOfPlayerZoneItemID int
type AnyOfPlayerZoneID int
type PlayerTargetRefID int
type PlayerTargetedByRefID int

type State struct {
	AnyOfPlayerZone          map[AnyOfPlayerZoneID]anyOfPlayerZoneCore                   `json:"anyOfPlayerZone"`
	AnyOfPlayerZoneItem      map[AnyOfPlayerZoneItemID]anyOfPlayerZoneItemCore           `json:"anyOfPlayerZoneItem"`
	EquipmentSet             map[EquipmentSetID]equipmentSetCore                         `json:"equipmentSet"`
	EquipmentSetEquipmentRef map[EquipmentSetEquipmentRefID]equipmentSetEquipmentRefCore `json:"equipmentSetEquipmentRef"`
	GearScore                map[GearScoreID]gearScoreCore                               `json:"gearScore"`
	Item                     map[ItemID]itemCore                                         `json:"item"`
	ItemBoundToRef           map[ItemBoundToRefID]itemBoundToRefCore                     `json:"itemBoundToRef"`
	Player                   map[PlayerID]playerCore                                     `json:"player"`
	PlayerEquipmentSetRef    map[PlayerEquipmentSetRefID]playerEquipmentSetRefCore       `json:"PlayerEquipmentSerRef"`
	PlayerGuildMemberRef     map[PlayerGuildMemberRefID]playerGuildMemberRefCore         `json:"playerGuildMemberRef"`
	PlayerTargetRef          map[PlayerTargetRefID]playerTargetRefCore                   `json:"playerTargetRef"`
	PlayerTargetedByRef      map[PlayerTargetedByRefID]playerTargetedByRefCore           `json:"playerTargetedByRef"`
	Position                 map[PositionID]positionCore                                 `json:"position"`
	Zone                     map[ZoneID]zoneCore                                         `json:"zone"`
	ZoneItem                 map[ZoneItemID]zoneItemCore                                 `json:"zoneItem"`
}

func newState() State {
	return State{
		AnyOfPlayerZone:          make(map[AnyOfPlayerZoneID]anyOfPlayerZoneCore),
		AnyOfPlayerZoneItem:      make(map[AnyOfPlayerZoneItemID]anyOfPlayerZoneItemCore),
		EquipmentSet:             make(map[EquipmentSetID]equipmentSetCore),
		EquipmentSetEquipmentRef: make(map[EquipmentSetEquipmentRefID]equipmentSetEquipmentRefCore),
		GearScore:                make(map[GearScoreID]gearScoreCore),
		Item:                     make(map[ItemID]itemCore),
		ItemBoundToRef:           make(map[ItemBoundToRefID]itemBoundToRefCore),
		Player:                   make(map[PlayerID]playerCore),
		PlayerEquipmentSetRef:    make(map[PlayerEquipmentSetRefID]playerEquipmentSetRefCore),
		PlayerGuildMemberRef:     make(map[PlayerGuildMemberRefID]playerGuildMemberRefCore),
		PlayerTargetRef:          make(map[PlayerTargetRefID]playerTargetRefCore),
		PlayerTargetedByRef:      make(map[PlayerTargetedByRefID]playerTargetedByRefCore),
		Position:                 make(map[PositionID]positionCore),
		Zone:                     make(map[ZoneID]zoneCore),
		ZoneItem:                 make(map[ZoneItemID]zoneItemCore),
	}
}

type zoneCore struct {
	ID             ZoneID        `json:"id"`
	Items          []ZoneItemID  `json:"items"`
	Players        []PlayerID    `json:"players"`
	Tags           []string      `json:"tags"`
	OperationKind_ OperationKind `json:"operationKind_"`
	engine         *Engine
}

type zone struct{ zone zoneCore }

type zoneItemCore struct {
	ID             ZoneItemID    `json:"id"`
	Item           ItemID        `json:"item"`
	Position       PositionID    `json:"position"`
	OperationKind_ OperationKind `json:"operationKind_"`
	HasParent_     bool          `json:"hasParent_"`
	engine         *Engine
}

type zoneItem struct{ zoneItem zoneItemCore }

type itemCore struct {
	ID             ItemID            `json:"id"`
	BoundTo        ItemBoundToRefID  `json:"boundTo"`
	GearScore      GearScoreID       `json:"gearScore"`
	Name           string            `json:"name"`
	Origin         AnyOfPlayerZoneID `json:"origin"`
	OperationKind_ OperationKind     `json:"operationKind_"`
	HasParent_     bool              `json:"hasParent_"`
	engine         *Engine
}

type item struct{ item itemCore }

type playerCore struct {
	ID             PlayerID                  `json:"id"`
	GearScore      GearScoreID               `json:"gearScore"`
	EquipmentSets  []PlayerEquipmentSetRefID `json:"equipmentSets"`
	GuildMembers   []PlayerGuildMemberRefID  `json:"guildMembers"`
	Items          []ItemID                  `json:"items"`
	Position       PositionID                `json:"position"`
	Target         PlayerTargetRefID         `json:"target"`
	TargetedBy     []PlayerTargetedByRefID   `json:"targetedBy"`
	OperationKind_ OperationKind             `json:"operationKind_"`
	HasParent_     bool                      `json:"hasParent_"`
	engine         *Engine
}

type player struct{ player playerCore }

type gearScoreCore struct {
	ID             GearScoreID   `json:"id"`
	Level          int           `json:"level"`
	Score          int           `json:"score"`
	OperationKind_ OperationKind `json:"operationKind_"`
	HasParent_     bool          `json:"hasParent_"`
	engine         *Engine
}

type gearScore struct{ gearScore gearScoreCore }

type positionCore struct {
	ID             PositionID    `json:"id"`
	X              float64       `json:"x"`
	Y              float64       `json:"y"`
	OperationKind_ OperationKind `json:"operationKind_"`
	HasParent_     bool          `json:"hasParent_"`
	engine         *Engine
}

type position struct{ position positionCore }

type equipmentSetCore struct {
	ID             EquipmentSetID               `json:"id"`
	Name           string                       `json:"name"`
	Equipment      []EquipmentSetEquipmentRefID `json:"equipment"`
	OperationKind_ OperationKind                `json:"operationKind_"`
	engine         *Engine
}

type equipmentSet struct{ equipmentSet equipmentSetCore }

type itemBoundToRefCore struct {
	ID                  ItemBoundToRefID `json:"id"`
	ParentID            ItemID           `json:"parentID"`
	ReferencedElementID PlayerID         `json:"referencedElementID"`
	OperationKind_      OperationKind    `json:"operationKind_"`
	engine              *Engine
}

type itemBoundToRef struct{ itemBoundToRef itemBoundToRefCore }

type playerGuildMemberRefCore struct {
	ID                  PlayerGuildMemberRefID `json:"id"`
	ParentID            PlayerID               `json:"parentID"`
	ReferencedElementID PlayerID               `json:"referencedElementID"`
	OperationKind_      OperationKind          `json:"operationKind_"`
	engine              *Engine
}

type playerGuildMemberRef struct{ playerGuildMemberRef playerGuildMemberRefCore }

type equipmentSetEquipmentRefCore struct {
	ID                  EquipmentSetEquipmentRefID `json:"id"`
	ParentID            EquipmentSetID             `json:"parentID"`
	ReferencedElementID ItemID                     `json:"referencedElementID"`
	OperationKind_      OperationKind              `json:"operationKind_"`
	engine              *Engine
}

type equipmentSetEquipmentRef struct{ equipmentSetEquipmentRef equipmentSetEquipmentRefCore }

type playerEquipmentSetRefCore struct {
	ID                  PlayerEquipmentSetRefID `json:"id"`
	ParentID            PlayerID                `json:"parentID"`
	ReferencedElementID EquipmentSetID          `json:"referencedElementID"`
	OperationKind_      OperationKind           `json:"operationKind_"`
	engine              *Engine
}

type playerEquipmentSetRef struct{ playerEquipmentSetRef playerEquipmentSetRefCore }

type anyOfPlayerZoneCore struct {
	ID             AnyOfPlayerZoneID `json:"id"`
	Player         PlayerID          `json:"player"`
	Zone           ZoneID            `json:"zone"`
	OperationKind_ OperationKind     `json:"operationKind_"`
	engine         *Engine
}

type anyOfPlayerZone struct{ anyOfPlayerZone anyOfPlayerZoneCore }

type anyOfPlayerZoneItemCore struct {
	ID             AnyOfPlayerZoneItemID `json:"id"`
	Player         PlayerID              `json:"player"`
	ZoneItem       ZoneItem              `json:"zoneItem"`
	OperationKind_ OperationKind         `json:"operationKind_"`
	engine         *Engine
}

type anyOfPlayerZoneItem struct{ anyOfPlayerZoneItem anyOfPlayerZoneItemCore }

type playerTargetRef struct{ playerTargetRef playerTargetRefCore }

type playerTargetRefCore struct {
	ID                  PlayerTargetRefID     `json:"id"`
	ParentID            PlayerID              `json:"parentID"`
	ReferencedElementID AnyOfPlayerZoneItemID `json:"referencedElementID"`
	OperationKind_      OperationKind         `json:"operationKind_"`
	engine              *Engine
}

type playerTargetedByRef struct{ playerTargetedByRef playerTargetedByRefCore }

type playerTargetedByRefCore struct {
	ID                  PlayerTargetedByRefID `json:"id"`
	ParentID            PlayerID              `json:"parentID"`
	ReferencedElementID AnyOfPlayerZoneItemID `json:"referencedElementID"`
	OperationKind_      OperationKind         `json:"operationKind_"`
	engine              *Engine
}
