package state

type BoolValueID int
type FloatValueID int
type IntValueID int
type StringValueID int

type AttackEventID int
type EquipmentSetID int
type GearScoreID int
type ItemID int
type PlayerID int
type PositionID int
type ZoneID int
type ZoneItemID int
type AttackEventTargetRefID int
type PlayerGuildMemberRefID int
type ItemBoundToRefID int
type EquipmentSetEquipmentRefID int
type PlayerEquipmentSetRefID int
type AnyOfItem_Player_ZoneItemID int
type AnyOfPlayer_ZoneItemID int
type AnyOfPlayer_PositionID int
type PlayerTargetRefID int
type PlayerTargetedByRefID int

type State struct {
	BoolValue                 map[BoolValueID]boolValue                                     `json:"boolValue"`
	FloatValue                map[FloatValueID]floatValue                                   `json:"floatValue"`
	IntValue                  map[IntValueID]intValue                                       `json:"intValue"`
	StringValue               map[StringValueID]stringValue                                 `json:"stringValue"`
	AttackEvent               map[AttackEventID]attackEventCore                             `json:"attackEvent"`
	EquipmentSet              map[EquipmentSetID]equipmentSetCore                           `json:"equipmentSet"`
	GearScore                 map[GearScoreID]gearScoreCore                                 `json:"gearScore"`
	Item                      map[ItemID]itemCore                                           `json:"item"`
	Player                    map[PlayerID]playerCore                                       `json:"player"`
	Position                  map[PositionID]positionCore                                   `json:"position"`
	Zone                      map[ZoneID]zoneCore                                           `json:"zone"`
	ZoneItem                  map[ZoneItemID]zoneItemCore                                   `json:"zoneItem"`
	AttackEventTargetRef      map[AttackEventTargetRefID]attackEventTargetRefCore           `json:"attackEventTargetRef"`
	EquipmentSetEquipmentRef  map[EquipmentSetEquipmentRefID]equipmentSetEquipmentRefCore   `json:"equipmentSetEquipmentRef"`
	ItemBoundToRef            map[ItemBoundToRefID]itemBoundToRefCore                       `json:"itemBoundToRef"`
	PlayerEquipmentSetRef     map[PlayerEquipmentSetRefID]playerEquipmentSetRefCore         `json:"playerEquipmentSetRef"`
	PlayerGuildMemberRef      map[PlayerGuildMemberRefID]playerGuildMemberRefCore           `json:"playerGuildMemberRef"`
	PlayerTargetRef           map[PlayerTargetRefID]playerTargetRefCore                     `json:"playerTargetRef"`
	PlayerTargetedByRef       map[PlayerTargetedByRefID]playerTargetedByRefCore             `json:"playerTargetedByRef"`
	AnyOfPlayer_Position      map[AnyOfPlayer_PositionID]anyOfPlayer_PositionCore           `json:"anyOfPlayer_Position"`
	AnyOfPlayer_ZoneItem      map[AnyOfPlayer_ZoneItemID]anyOfPlayer_ZoneItemCore           `json:"anyOfPlayer_ZoneItem"`
	AnyOfItem_Player_ZoneItem map[AnyOfItem_Player_ZoneItemID]anyOfItem_Player_ZoneItemCore `json:"anyOfItem_Player_ZoneItem"`
}

func newState() *State {
	return &State{
		BoolValue:                 make(map[BoolValueID]boolValue),
		FloatValue:                make(map[FloatValueID]floatValue),
		IntValue:                  make(map[IntValueID]intValue),
		StringValue:               make(map[StringValueID]stringValue),
		AttackEvent:               make(map[AttackEventID]attackEventCore),
		EquipmentSet:              make(map[EquipmentSetID]equipmentSetCore),
		GearScore:                 make(map[GearScoreID]gearScoreCore),
		Item:                      make(map[ItemID]itemCore),
		Player:                    make(map[PlayerID]playerCore),
		Position:                  make(map[PositionID]positionCore),
		Zone:                      make(map[ZoneID]zoneCore),
		ZoneItem:                  make(map[ZoneItemID]zoneItemCore),
		AttackEventTargetRef:      make(map[AttackEventTargetRefID]attackEventTargetRefCore),
		EquipmentSetEquipmentRef:  make(map[EquipmentSetEquipmentRefID]equipmentSetEquipmentRefCore),
		ItemBoundToRef:            make(map[ItemBoundToRefID]itemBoundToRefCore),
		PlayerEquipmentSetRef:     make(map[PlayerEquipmentSetRefID]playerEquipmentSetRefCore),
		PlayerGuildMemberRef:      make(map[PlayerGuildMemberRefID]playerGuildMemberRefCore),
		PlayerTargetRef:           make(map[PlayerTargetRefID]playerTargetRefCore),
		PlayerTargetedByRef:       make(map[PlayerTargetedByRefID]playerTargetedByRefCore),
		AnyOfPlayer_Position:      make(map[AnyOfPlayer_PositionID]anyOfPlayer_PositionCore),
		AnyOfPlayer_ZoneItem:      make(map[AnyOfPlayer_ZoneItemID]anyOfPlayer_ZoneItemCore),
		AnyOfItem_Player_ZoneItem: make(map[AnyOfItem_Player_ZoneItemID]anyOfItem_Player_ZoneItemCore),
	}
}

type boolValue struct {
	ID            BoolValueID   `json:"id"`
	Value         bool          `json:"value"`
	OperationKind OperationKind `json:"operationKind"`
	JSONPath      string        `json:"jsonPath"`
	Path          path          `json:"path"`
	engine        *Engine
}

type intValue struct {
	ID            IntValueID    `json:"id"`
	Value         int64         `json:"value"`
	OperationKind OperationKind `json:"operationKind"`
	JSONPath      string        `json:"jsonPath"`
	Path          path          `json:"path"`
	engine        *Engine
}

type floatValue struct {
	ID            FloatValueID  `json:"id"`
	Value         float64       `json:"value"`
	OperationKind OperationKind `json:"operationKind"`
	JSONPath      string        `json:"jsonPath"`
	Path          path          `json:"path"`
	engine        *Engine
}

type stringValue struct {
	ID            StringValueID `json:"id"`
	Value         string        `json:"value"`
	OperationKind OperationKind `json:"operationKind"`
	JSONPath      string        `json:"jsonPath"`
	Path          path          `json:"path"`
	engine        *Engine
}

type attackEventCore struct {
	ID            AttackEventID          `json:"id"`
	Target        AttackEventTargetRefID `json:"target"`
	OperationKind OperationKind          `json:"operationKind"`
	HasParent     bool                   `json:"hasParent"`
	JSONPath      string                 `json:"jsonPath"`
	Path          path                   `json:"path"`
	engine        *Engine
}

type AttackEvent struct{ attackEvent attackEventCore }

type zoneCore struct {
	ID            ZoneID                        `json:"id"`
	Interactables []AnyOfItem_Player_ZoneItemID `json:"interactables"`
	Items         []ZoneItemID                  `json:"items"`
	Players       []PlayerID                    `json:"players"`
	Tags          []StringValueID               `json:"tags"`
	OperationKind OperationKind                 `json:"operationKind"`
	HasParent     bool                          `json:"hasParent"`
	JSONPath      string                        `json:"jsonPath"`
	Path          path                          `json:"path"`
	engine        *Engine
}

type Zone struct{ zone zoneCore }

type zoneItemCore struct {
	ID            ZoneItemID    `json:"id"`
	Item          ItemID        `json:"item"`
	Position      PositionID    `json:"position"`
	OperationKind OperationKind `json:"operationKind"`
	HasParent     bool          `json:"hasParent"`
	JSONPath      string        `json:"jsonPath"`
	Path          path          `json:"path"`
	engine        *Engine
}

type ZoneItem struct{ zoneItem zoneItemCore }

type itemCore struct {
	ID            ItemID                 `json:"id"`
	BoundTo       ItemBoundToRefID       `json:"boundTo"`
	GearScore     GearScoreID            `json:"gearScore"`
	Name          StringValueID          `json:"name"`
	Origin        AnyOfPlayer_PositionID `json:"origin"`
	OperationKind OperationKind          `json:"operationKind"`
	HasParent     bool                   `json:"hasParent"`
	JSONPath      string                 `json:"jsonPath"`
	Path          path                   `json:"path"`
	engine        *Engine
}

type Item struct{ item itemCore }

type playerCore struct {
	ID            PlayerID                  `json:"id"`
	Action        []AttackEventID           `json:"action"`
	EquipmentSets []PlayerEquipmentSetRefID `json:"equipmentSets"`
	GearScore     GearScoreID               `json:"gearScore"`
	GuildMembers  []PlayerGuildMemberRefID  `json:"guildMembers"`
	Items         []ItemID                  `json:"items"`
	Position      PositionID                `json:"position"`
	Target        PlayerTargetRefID         `json:"target"`
	TargetedBy    []PlayerTargetedByRefID   `json:"targetedBy"`
	OperationKind OperationKind             `json:"operationKind"`
	HasParent     bool                      `json:"hasParent"`
	JSONPath      string                    `json:"jsonPath"`
	Path          path                      `json:"path"`
	engine        *Engine
}

type Player struct{ player playerCore }

type gearScoreCore struct {
	ID            GearScoreID   `json:"id"`
	Level         IntValueID    `json:"level"`
	Score         IntValueID    `json:"score"`
	OperationKind OperationKind `json:"operationKind"`
	HasParent     bool          `json:"hasParent"`
	JSONPath      string        `json:"jsonPath"`
	Path          path          `json:"path"`
	engine        *Engine
}

type GearScore struct{ gearScore gearScoreCore }

type positionCore struct {
	ID            PositionID    `json:"id"`
	X             FloatValueID  `json:"x"`
	Y             FloatValueID  `json:"y"`
	OperationKind OperationKind `json:"operationKind"`
	HasParent     bool          `json:"hasParent"`
	JSONPath      string        `json:"jsonPath"`
	Path          path          `json:"path"`
	engine        *Engine
}

type Position struct{ position positionCore }

type equipmentSetCore struct {
	ID            EquipmentSetID               `json:"id"`
	Equipment     []EquipmentSetEquipmentRefID `json:"equipment"`
	Name          StringValueID                `json:"name"`
	OperationKind OperationKind                `json:"operationKind"`
	HasParent     bool                         `json:"hasParent"`
	JSONPath      string                       `json:"jsonPath"`
	Path          path                         `json:"path"`
	engine        *Engine
}

type EquipmentSet struct{ equipmentSet equipmentSetCore }

type itemBoundToRefCore struct {
	ID                  ItemBoundToRefID `json:"id"`
	ParentID            ItemID           `json:"parentID"`
	ChildID             int              `json:"childID"`
	ReferencedElementID PlayerID         `json:"referencedElementID"`
	OperationKind       OperationKind    `json:"operationKind"`
	Path                path             `json:"path"`
	engine              *Engine
}

type ItemBoundToRef struct{ itemBoundToRef itemBoundToRefCore }

type attackEventTargetRefCore struct {
	ID                  AttackEventTargetRefID `json:"id"`
	ParentID            AttackEventID          `json:"parentID"`
	ChildID             int                    `json:"childID"`
	ReferencedElementID PlayerID               `json:"referencedElementID"`
	OperationKind       OperationKind          `json:"operationKind"`
	Path                path                   `json:"path"`
	engine              *Engine
}

type AttackEventTargetRef struct{ attackEventTargetRef attackEventTargetRefCore }

type playerGuildMemberRefCore struct {
	ID                  PlayerGuildMemberRefID `json:"id"`
	ParentID            PlayerID               `json:"parentID"`
	ChildID             int                    `json:"childID"`
	ReferencedElementID PlayerID               `json:"referencedElementID"`
	OperationKind       OperationKind          `json:"operationKind"`
	Path                path                   `json:"path"`
	engine              *Engine
}

type PlayerGuildMemberRef struct{ playerGuildMemberRef playerGuildMemberRefCore }

type equipmentSetEquipmentRefCore struct {
	ID                  EquipmentSetEquipmentRefID `json:"id"`
	ParentID            EquipmentSetID             `json:"parentID"`
	ChildID             int                        `json:"childID"`
	ReferencedElementID ItemID                     `json:"referencedElementID"`
	OperationKind       OperationKind              `json:"operationKind"`
	Path                path                       `json:"path"`
	engine              *Engine
}

type EquipmentSetEquipmentRef struct{ equipmentSetEquipmentRef equipmentSetEquipmentRefCore }

type playerEquipmentSetRefCore struct {
	ID                  PlayerEquipmentSetRefID `json:"id"`
	ParentID            PlayerID                `json:"parentID"`
	ChildID             int                     `json:"childID"`
	ReferencedElementID EquipmentSetID          `json:"referencedElementID"`
	OperationKind       OperationKind           `json:"operationKind"`
	Path                path                    `json:"path"`
	engine              *Engine
}

type PlayerEquipmentSetRef struct{ playerEquipmentSetRef playerEquipmentSetRefCore }

type anyOfPlayer_PositionCore struct {
	ID                AnyOfPlayer_PositionID `json:"id"`
	ElementKind       ElementKind            `json:"elementKind"`
	ParentID          int                    `json:"parentID"`
	ChildID           int                    `json:"childID"`
	ParentElementPath path                   `json:"parentElementPath"`
	FieldIdentifier   treeFieldIdentifier    `json:"fieldIdentifier"`
	OperationKind     OperationKind          `json:"operationKind"`
	engine            *Engine
}

type AnyOfPlayer_Position struct{ anyOfPlayer_Position anyOfPlayer_PositionCore }

type anyOfPlayer_ZoneItemCore struct {
	ID                AnyOfPlayer_ZoneItemID `json:"id"`
	ElementKind       ElementKind            `json:"elementKind"`
	ParentID          int                    `json:"parentID"`
	ChildID           int                    `json:"childID"`
	ParentElementPath path                   `json:"parentElementPath"`
	FieldIdentifier   treeFieldIdentifier    `json:"fieldIdentifier"`
	OperationKind     OperationKind          `json:"operationKind"`
	engine            *Engine
}

type AnyOfPlayer_ZoneItem struct{ anyOfPlayer_ZoneItem anyOfPlayer_ZoneItemCore }

type anyOfItem_Player_ZoneItemCore struct {
	ID                AnyOfItem_Player_ZoneItemID `json:"id"`
	ElementKind       ElementKind                 `json:"elementKind"`
	ParentID          int                         `json:"parentID"`
	ChildID           int                         `json:"childID"`
	ParentElementPath path                        `json:"parentElementPath"`
	FieldIdentifier   treeFieldIdentifier         `json:"fieldIdentifier"`
	OperationKind     OperationKind               `json:"operationKind"`
	engine            *Engine
}

type AnyOfItem_Player_ZoneItem struct{ anyOfItem_Player_ZoneItem anyOfItem_Player_ZoneItemCore }

type playerTargetRefCore struct {
	ID                  PlayerTargetRefID      `json:"id"`
	ParentID            PlayerID               `json:"parentID"`
	ChildID             int                    `json:"childID"`
	ReferencedElementID AnyOfPlayer_ZoneItemID `json:"referencedElementID"`
	OperationKind       OperationKind          `json:"operationKind"`
	Path                path                   `json:"path"`
	engine              *Engine
}

type PlayerTargetRef struct{ playerTargetRef playerTargetRefCore }

type playerTargetedByRefCore struct {
	ID                  PlayerTargetedByRefID  `json:"id"`
	ParentID            PlayerID               `json:"parentID"`
	ChildID             int                    `json:"childID"`
	ReferencedElementID AnyOfPlayer_ZoneItemID `json:"referencedElementID"`
	OperationKind       OperationKind          `json:"operationKind"`
	Path                path                   `json:"path"`
	engine              *Engine
}

type PlayerTargetedByRef struct{ playerTargetedByRef playerTargetedByRefCore }

func (s State) IsEmpty() bool {
	if len(s.BoolValue) != 0 {
		return false
	}
	if len(s.FloatValue) != 0 {
		return false
	}
	if len(s.IntValue) != 0 {
		return false
	}
	if len(s.StringValue) != 0 {
		return false
	}
	if len(s.AttackEvent) != 0 {
		return false
	}
	if len(s.EquipmentSet) != 0 {
		return false
	}
	if len(s.GearScore) != 0 {
		return false
	}
	if len(s.Item) != 0 {
		return false
	}
	if len(s.Player) != 0 {
		return false
	}
	if len(s.Position) != 0 {
		return false
	}
	if len(s.Zone) != 0 {
		return false
	}
	if len(s.ZoneItem) != 0 {
		return false
	}
	if len(s.AttackEventTargetRef) != 0 {
		return false
	}
	if len(s.EquipmentSetEquipmentRef) != 0 {
		return false
	}
	if len(s.ItemBoundToRef) != 0 {
		return false
	}
	if len(s.PlayerEquipmentSetRef) != 0 {
		return false
	}
	if len(s.PlayerGuildMemberRef) != 0 {
		return false
	}
	if len(s.PlayerTargetRef) != 0 {
		return false
	}
	if len(s.PlayerTargetedByRef) != 0 {
		return false
	}
	if len(s.AnyOfPlayer_Position) != 0 {
		return false
	}
	if len(s.AnyOfPlayer_ZoneItem) != 0 {
		return false
	}
	if len(s.AnyOfItem_Player_ZoneItem) != 0 {
		return false
	}

	return true
}
