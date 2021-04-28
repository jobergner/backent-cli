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

type State struct {
	EquipmentSet             map[EquipmentSetID]equipmentSetCore                         `json:"equipmentSet"`
	EquipmentSetEquipmentRef map[EquipmentSetEquipmentRefID]equipmentSetEquipmentRefCore `json:"equipmentSetEquipmentRef"`
	GearScore                map[GearScoreID]gearScoreCore                               `json:"gearScore"`
	Item                     map[ItemID]itemCore                                         `json:"item"`
	ItemBoundToRef           map[ItemBoundToRefID]itemBoundToRefCore                     `json:"itemBoundToRef"`
	Player                   map[PlayerID]playerCore                                     `json:"player"`
	PlayerEquipmentSetRef    map[PlayerEquipmentSetRefID]playerEquipmentSetRefCore       `json:"PlayerEquipmentSerRef"`
	PlayerGuildMemberRef     map[PlayerGuildMemberRefID]playerGuildMemberRefCore         `json:"playerGuildMemberRef"`
	Position                 map[PositionID]positionCore                                 `json:"position"`
	Zone                     map[ZoneID]zoneCore                                         `json:"zone"`
	ZoneItem                 map[ZoneItemID]zoneItemCore                                 `json:"zoneItem"`
}

func newState() State {
	return State{
		EquipmentSet:             make(map[EquipmentSetID]equipmentSetCore),
		EquipmentSetEquipmentRef: make(map[EquipmentSetEquipmentRefID]equipmentSetEquipmentRefCore),
		GearScore:                make(map[GearScoreID]gearScoreCore),
		Item:                     make(map[ItemID]itemCore),
		ItemBoundToRef:           make(map[ItemBoundToRefID]itemBoundToRefCore),
		Player:                   make(map[PlayerID]playerCore),
		PlayerEquipmentSetRef:    make(map[PlayerEquipmentSetRefID]playerEquipmentSetRefCore),
		PlayerGuildMemberRef:     make(map[PlayerGuildMemberRefID]playerGuildMemberRefCore),
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
	ID             ItemID           `json:"id"`
	BoundTo        ItemBoundToRefID `json:"boundTo"`
	GearScore      GearScoreID      `json:"gearScore"`
	Name           string           `json:"name"`
	OperationKind_ OperationKind    `json:"operationKind_"`
	HasParent_     bool             `json:"hasParent_"`
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
