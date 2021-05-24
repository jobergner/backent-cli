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
type AnyOfItemPlayerZoneItemID int
type AnyOfPlayerZoneItemID int
type AnyOfPlayerPositionID int
type PlayerTargetRefID int
type PlayerTargetedByRefID int

type State struct {
	EquipmentSet             map[EquipmentSetID]equipmentSetCore                         `json:"equipmentSet"`
	GearScore                map[GearScoreID]gearScoreCore                               `json:"gearScore"`
	Item                     map[ItemID]itemCore                                         `json:"item"`
	Player                   map[PlayerID]playerCore                                     `json:"player"`
	Position                 map[PositionID]positionCore                                 `json:"position"`
	Zone                     map[ZoneID]zoneCore                                         `json:"zone"`
	ZoneItem                 map[ZoneItemID]zoneItemCore                                 `json:"zoneItem"`
	EquipmentSetEquipmentRef map[EquipmentSetEquipmentRefID]equipmentSetEquipmentRefCore `json:"equipmentSetEquipmentRef"`
	ItemBoundToRef           map[ItemBoundToRefID]itemBoundToRefCore                     `json:"itemBoundToRef"`
	PlayerEquipmentSetRef    map[PlayerEquipmentSetRefID]playerEquipmentSetRefCore       `json:"playerEquipmentSetRef"`
	PlayerGuildMemberRef     map[PlayerGuildMemberRefID]playerGuildMemberRefCore         `json:"playerGuildMemberRef"`
	PlayerTargetRef          map[PlayerTargetRefID]playerTargetRefCore                   `json:"playerTargetRef"`
	PlayerTargetedByRef      map[PlayerTargetedByRefID]playerTargetedByRefCore           `json:"playerTargetedByRef"`
	AnyOfPlayerPosition      map[AnyOfPlayerPositionID]anyOfPlayerPositionCore           `json:"anyOfPlayerPosition"`
	AnyOfPlayerZoneItem      map[AnyOfPlayerZoneItemID]anyOfPlayerZoneItemCore           `json:"anyOfPlayerZoneItem"`
	AnyOfItemPlayerZoneItem  map[AnyOfItemPlayerZoneItemID]anyOfItemPlayerZoneItemCore   `json:"anyOfItemPlayerZoneItem"`
}

func newState() State {
	return State{
		EquipmentSet:             make(map[EquipmentSetID]equipmentSetCore),
		GearScore:                make(map[GearScoreID]gearScoreCore),
		Item:                     make(map[ItemID]itemCore),
		Player:                   make(map[PlayerID]playerCore),
		Position:                 make(map[PositionID]positionCore),
		Zone:                     make(map[ZoneID]zoneCore),
		ZoneItem:                 make(map[ZoneItemID]zoneItemCore),
		EquipmentSetEquipmentRef: make(map[EquipmentSetEquipmentRefID]equipmentSetEquipmentRefCore),
		ItemBoundToRef:           make(map[ItemBoundToRefID]itemBoundToRefCore),
		PlayerEquipmentSetRef:    make(map[PlayerEquipmentSetRefID]playerEquipmentSetRefCore),
		PlayerGuildMemberRef:     make(map[PlayerGuildMemberRefID]playerGuildMemberRefCore),
		PlayerTargetRef:          make(map[PlayerTargetRefID]playerTargetRefCore),
		PlayerTargetedByRef:      make(map[PlayerTargetedByRefID]playerTargetedByRefCore),
		AnyOfPlayerPosition:      make(map[AnyOfPlayerPositionID]anyOfPlayerPositionCore),
		AnyOfPlayerZoneItem:      make(map[AnyOfPlayerZoneItemID]anyOfPlayerZoneItemCore),
		AnyOfItemPlayerZoneItem:  make(map[AnyOfItemPlayerZoneItemID]anyOfItemPlayerZoneItemCore),
	}
}

type zoneCore struct {
	ID            ZoneID                      `json:"id"`
	Interactables []AnyOfItemPlayerZoneItemID `json:"interactables"`
	Items         []ZoneItemID                `json:"items"`
	Players       []PlayerID                  `json:"players"`
	Tags          []string                    `json:"tags"`
	OperationKind OperationKind               `json:"operationKind"`
	engine        *Engine
}

type zone struct{ zone zoneCore }

type zoneItemCore struct {
	ID            ZoneItemID    `json:"id"`
	Item          ItemID        `json:"item"`
	Position      PositionID    `json:"position"`
	OperationKind OperationKind `json:"operationKind"`
	HasParent     bool          `json:"hasParent"`
	engine        *Engine
}

type zoneItem struct{ zoneItem zoneItemCore }

type itemCore struct {
	ID            ItemID                `json:"id"`
	BoundTo       ItemBoundToRefID      `json:"boundTo"`
	GearScore     GearScoreID           `json:"gearScore"`
	Name          string                `json:"name"`
	Origin        AnyOfPlayerPositionID `json:"origin"`
	OperationKind OperationKind         `json:"operationKind"`
	HasParent     bool                  `json:"hasParent"`
	engine        *Engine
}

type item struct{ item itemCore }

type playerCore struct {
	ID            PlayerID                  `json:"id"`
	EquipmentSets []PlayerEquipmentSetRefID `json:"equipmentSets"`
	GearScore     GearScoreID               `json:"gearScore"`
	GuildMembers  []PlayerGuildMemberRefID  `json:"guildMembers"`
	Items         []ItemID                  `json:"items"`
	Position      PositionID                `json:"position"`
	Target        PlayerTargetRefID         `json:"target"`
	TargetedBy    []PlayerTargetedByRefID   `json:"targetedBy"`
	OperationKind OperationKind             `json:"operationKind"`
	HasParent     bool                      `json:"hasParent"`
	engine        *Engine
}

type player struct{ player playerCore }

type gearScoreCore struct {
	ID            GearScoreID   `json:"id"`
	Level         int           `json:"level"`
	Score         int           `json:"score"`
	OperationKind OperationKind `json:"operationKind"`
	HasParent     bool          `json:"hasParent"`
	engine        *Engine
}

type gearScore struct{ gearScore gearScoreCore }

type positionCore struct {
	ID            PositionID    `json:"id"`
	X             float64       `json:"x"`
	Y             float64       `json:"y"`
	OperationKind OperationKind `json:"operationKind"`
	HasParent     bool          `json:"hasParent"`
	engine        *Engine
}

type position struct{ position positionCore }

type equipmentSetCore struct {
	ID            EquipmentSetID               `json:"id"`
	Equipment     []EquipmentSetEquipmentRefID `json:"equipment"`
	Name          string                       `json:"name"`
	OperationKind OperationKind                `json:"operationKind"`
	engine        *Engine
}

type equipmentSet struct{ equipmentSet equipmentSetCore }

type itemBoundToRefCore struct {
	ID                  ItemBoundToRefID `json:"id"`
	ParentID            ItemID           `json:"parentID"`
	ReferencedElementID PlayerID         `json:"referencedElementID"`
	OperationKind       OperationKind    `json:"operationKind"`
	engine              *Engine
}

type itemBoundToRef struct{ itemBoundToRef itemBoundToRefCore }

type playerGuildMemberRefCore struct {
	ID                  PlayerGuildMemberRefID `json:"id"`
	ParentID            PlayerID               `json:"parentID"`
	ReferencedElementID PlayerID               `json:"referencedElementID"`
	OperationKind       OperationKind          `json:"operationKind"`
	engine              *Engine
}

type playerGuildMemberRef struct{ playerGuildMemberRef playerGuildMemberRefCore }

type equipmentSetEquipmentRefCore struct {
	ID                  EquipmentSetEquipmentRefID `json:"id"`
	ParentID            EquipmentSetID             `json:"parentID"`
	ReferencedElementID ItemID                     `json:"referencedElementID"`
	OperationKind       OperationKind              `json:"operationKind"`
	engine              *Engine
}

type equipmentSetEquipmentRef struct{ equipmentSetEquipmentRef equipmentSetEquipmentRefCore }

type playerEquipmentSetRefCore struct {
	ID                  PlayerEquipmentSetRefID `json:"id"`
	ParentID            PlayerID                `json:"parentID"`
	ReferencedElementID EquipmentSetID          `json:"referencedElementID"`
	OperationKind       OperationKind           `json:"operationKind"`
	engine              *Engine
}

type playerEquipmentSetRef struct{ playerEquipmentSetRef playerEquipmentSetRefCore }

type anyOfPlayerPositionCore struct {
	ID            AnyOfPlayerPositionID `json:"id"`
	ElementKind   ElementKind           `json:"elementKind"`
	Player        PlayerID              `json:"player"`
	Position      PositionID            `json:"position"`
	OperationKind OperationKind         `json:"operationKind"`
	engine        *Engine
}

type anyOfPlayerPosition struct{ anyOfPlayerPosition anyOfPlayerPositionCore }

type anyOfPlayerZoneItemCore struct {
	ID            AnyOfPlayerZoneItemID `json:"id"`
	ElementKind   ElementKind           `json:"elementKind"`
	Player        PlayerID              `json:"player"`
	ZoneItem      ZoneItemID            `json:"zoneItem"`
	OperationKind OperationKind         `json:"operationKind"`
	engine        *Engine
}

type anyOfPlayerZoneItem struct{ anyOfPlayerZoneItem anyOfPlayerZoneItemCore }

type anyOfItemPlayerZoneItemCore struct {
	ID            AnyOfItemPlayerZoneItemID `json:"id"`
	ElementKind   ElementKind               `json:"elementKind"`
	Item          ItemID                    `json:"item"`
	Player        PlayerID                  `json:"player"`
	ZoneItem      ZoneItemID                `json:"zoneItem"`
	OperationKind OperationKind             `json:"operationKind"`
	engine        *Engine
}

type anyOfItemPlayerZoneItem struct{ anyOfItemPlayerZoneItem anyOfItemPlayerZoneItemCore }

type playerTargetRefCore struct {
	ID                  PlayerTargetRefID     `json:"id"`
	ParentID            PlayerID              `json:"parentID"`
	ReferencedElementID AnyOfPlayerZoneItemID `json:"referencedElementID"`
	OperationKind       OperationKind         `json:"operationKind"`
	engine              *Engine
}

type playerTargetRef struct{ playerTargetRef playerTargetRefCore }

type playerTargetedByRefCore struct {
	ID                  PlayerTargetedByRefID `json:"id"`
	ParentID            PlayerID              `json:"parentID"`
	ReferencedElementID AnyOfPlayerZoneItemID `json:"referencedElementID"`
	OperationKind       OperationKind         `json:"operationKind"`
	engine              *Engine
}

type playerTargetedByRef struct{ playerTargetedByRef playerTargetedByRefCore }
