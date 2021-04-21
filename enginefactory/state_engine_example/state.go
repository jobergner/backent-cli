package state

type GearScoreID int
type ItemID int
type PlayerID int
type PositionID int
type ZoneID int
type ZoneItemID int
type PlayerGuildMemberRefID int
type ItemBoundToRefID int

type State struct {
	GearScore            map[GearScoreID]gearScoreCore                       `json:"gearScore"`
	Item                 map[ItemID]itemCore                                 `json:"item"`
	ItemBoundToRef       map[ItemBoundToRefID]itemBoundToRefCore             `json:"itemBoundToRef"`
	Player               map[PlayerID]playerCore                             `json:"player"`
	PlayerGuildMemberRef map[PlayerGuildMemberRefID]playerGuildMemberRefCore `json:"playerGuildMemberRef"`
	Position             map[PositionID]positionCore                         `json:"position"`
	Zone                 map[ZoneID]zoneCore                                 `json:"zone"`
	ZoneItem             map[ZoneItemID]zoneItemCore                         `json:"zoneItem"`
}

func newState() State {
	return State{
		GearScore: make(map[GearScoreID]gearScoreCore),
		Item:      make(map[ItemID]itemCore),
		Player:    make(map[PlayerID]playerCore),
		Position:  make(map[PositionID]positionCore),
		Zone:      make(map[ZoneID]zoneCore),
		ZoneItem:  make(map[ZoneItemID]zoneItemCore),
	}
}

type zoneCore struct {
	ID             ZoneID        `json:"id"`
	Items          []ZoneItemID  `json:"items"`
	Players        []PlayerID    `json:"players"`
	Tags           []string      `json:"tags"`
	OperationKind_ OperationKind `json:"operationKind_"`
}

type zone struct{ zone zoneCore }

type zoneItemCore struct {
	ID             ZoneItemID    `json:"id"`
	Item           ItemID        `json:"item"`
	Position       PositionID    `json:"position"`
	OperationKind_ OperationKind `json:"operationKind_"`
	HasParent_     bool          `json:"hasParent_"`
}

type zoneItem struct{ zoneItem zoneItemCore }

type itemCore struct {
	ID             ItemID           `json:"id"`
	BoundTo        ItemBoundToRefID `json:"boundTo"`
	GearScore      GearScoreID      `json:"gearScore"`
	Name           string           `json:"name"`
	OperationKind_ OperationKind    `json:"operationKind_"`
	HasParent_     bool             `json:"hasParent_"`
}

type item struct{ item itemCore }

type playerCore struct {
	ID             PlayerID                 `json:"id"`
	GearScore      GearScoreID              `json:"gearScore"`
	GuildMembers   []PlayerGuildMemberRefID `json:"guildMembers"`
	Items          []ItemID                 `json:"items"`
	Position       PositionID               `json:"position"`
	OperationKind_ OperationKind            `json:"operationKind_"`
	HasParent_     bool                     `json:"hasParent_"`
}

type player struct{ player playerCore }

type gearScoreCore struct {
	ID             GearScoreID   `json:"id"`
	Level          int           `json:"level"`
	Score          int           `json:"score"`
	OperationKind_ OperationKind `json:"operationKind_"`
	HasParent_     bool          `json:"hasParent_"`
}

type gearScore struct{ gearScore gearScoreCore }

type positionCore struct {
	ID             PositionID    `json:"id"`
	X              float64       `json:"x"`
	Y              float64       `json:"y"`
	OperationKind_ OperationKind `json:"operationKind_"`
	HasParent_     bool          `json:"hasParent_"`
}

type position struct{ position positionCore }

type itemBoundToRefCore struct {
	ID                  ItemBoundToRefID `json:"id"`
	ParentID            ItemID           `json:"parentID"`
	ReferencedElementID PlayerID         `json:"referencedElementID"`
	OperationKind_      OperationKind    `json:"operationKind_"`
}

type itemBoundToRef struct{ itemBoundToRef itemBoundToRefCore }

type playerGuildMemberRefCore struct {
	ID                  PlayerGuildMemberRefID `json:"id"`
	ParentID            PlayerID               `json:"parentID"`
	ReferencedElementID PlayerID               `json:"referencedElementID"`
	OperationKind_      OperationKind          `json:"operationKind_"`
}

type playerGuildMemberRef struct{ playerGuildMemberRef playerGuildMemberRefCore }
