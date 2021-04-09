package state

type OperationKind string

const (
	OperationKindDelete OperationKind = "DELETE"
	OperationKindUpdate               = "UPDATE"
)

type GearScoreID int
type ItemID int
type PlayerID int
type PositionID int
type ZoneID int
type ZoneItemID int
type State struct {
	GearScore map[GearScoreID]gearScoreCore `json:"gearScore"`
	Item      map[ItemID]itemCore           `json:"item"`
	Player    map[PlayerID]playerCore       `json:"player"`
	Position  map[PositionID]positionCore   `json:"position"`
	Zone      map[ZoneID]zoneCore           `json:"zone"`
	ZoneItem  map[ZoneItemID]zoneItemCore   `json:"zoneItem"`
}

func newState() State {
	return State{GearScore: make(map[GearScoreID]gearScoreCore), Item: make(map[ItemID]itemCore), Player: make(map[PlayerID]playerCore), Position: make(map[PositionID]positionCore), Zone: make(map[ZoneID]zoneCore), ZoneItem: make(map[ZoneItemID]zoneItemCore)}
}

type Engine struct {
	State State
	Patch State
	IDgen int
}

func newEngine() *Engine {
	return &Engine{IDgen: 1, Patch: newState(), State: newState()}
}
func (se *Engine) GenerateID() int {
	newID := se.IDgen
	se.IDgen = se.IDgen + 1
	return newID
}
func (se *Engine) UpdateState() {
	for _, gearScore := range se.Patch.GearScore {
		if gearScore.OperationKind_ == OperationKindDelete {
			delete(se.State.GearScore, gearScore.ID)
		} else {
			se.State.GearScore[gearScore.ID] = gearScore
		}
	}
	for _, item := range se.Patch.Item {
		if item.OperationKind_ == OperationKindDelete {
			delete(se.State.Item, item.ID)
		} else {
			se.State.Item[item.ID] = item
		}
	}
	for _, player := range se.Patch.Player {
		if player.OperationKind_ == OperationKindDelete {
			delete(se.State.Player, player.ID)
		} else {
			se.State.Player[player.ID] = player
		}
	}
	for _, position := range se.Patch.Position {
		if position.OperationKind_ == OperationKindDelete {
			delete(se.State.Position, position.ID)
		} else {
			se.State.Position[position.ID] = position
		}
	}
	for _, zone := range se.Patch.Zone {
		if zone.OperationKind_ == OperationKindDelete {
			delete(se.State.Zone, zone.ID)
		} else {
			se.State.Zone[zone.ID] = zone
		}
	}
	for _, zoneItem := range se.Patch.ZoneItem {
		if zoneItem.OperationKind_ == OperationKindDelete {
			delete(se.State.ZoneItem, zoneItem.ID)
		} else {
			se.State.ZoneItem[zoneItem.ID] = zoneItem
		}
	}
	for key := range se.Patch.GearScore {
		delete(se.Patch.GearScore, key)
	}
	for key := range se.Patch.Item {
		delete(se.Patch.Item, key)
	}
	for key := range se.Patch.Player {
		delete(se.Patch.Player, key)
	}
	for key := range se.Patch.Position {
		delete(se.Patch.Position, key)
	}
	for key := range se.Patch.Zone {
		delete(se.Patch.Zone, key)
	}
	for key := range se.Patch.ZoneItem {
		delete(se.Patch.ZoneItem, key)
	}
}

type gearScoreCore struct {
	ID             GearScoreID   `json:"id"`
	Level          int           `json:"level"`
	Score          int           `json:"score"`
	OperationKind_ OperationKind `json:"operationKind_"`
	HasParent_     bool          `json:"hasParent_"`
}
type GearScore struct{ gearScore gearScoreCore }
type itemCore struct {
	ID             ItemID        `json:"id"`
	GearScore      GearScoreID   `json:"gearScore"`
	OperationKind_ OperationKind `json:"operationKind_"`
	HasParent_     bool          `json:"hasParent_"`
}
type Item struct{ item itemCore }
type playerCore struct {
	ID             PlayerID      `json:"id"`
	GearScore      GearScoreID   `json:"gearScore"`
	Items          []ItemID      `json:"items"`
	Position       PositionID    `json:"position"`
	OperationKind_ OperationKind `json:"operationKind_"`
	HasParent_     bool          `json:"hasParent_"`
}
type Player struct{ player playerCore }
type positionCore struct {
	ID             PositionID    `json:"id"`
	X              float64       `json:"x"`
	Y              float64       `json:"y"`
	OperationKind_ OperationKind `json:"operationKind_"`
	HasParent_     bool          `json:"hasParent_"`
}
type Position struct{ position positionCore }
type zoneCore struct {
	ID             ZoneID        `json:"id"`
	Items          []ZoneItemID  `json:"items"`
	Players        []PlayerID    `json:"players"`
	Tags           []string      `json:"tags"`
	OperationKind_ OperationKind `json:"operationKind_"`
}
type Zone struct{ zone zoneCore }
type zoneItemCore struct {
	ID             ZoneItemID    `json:"id"`
	Item           ItemID        `json:"item"`
	Position       PositionID    `json:"position"`
	OperationKind_ OperationKind `json:"operationKind_"`
	HasParent_     bool          `json:"hasParent_"`
}
type ZoneItem struct{ zoneItem zoneItemCore }

func (_player Player) AddItem(se *Engine) Item {
	player := se.Player(_player.player.ID)
	if player.player.OperationKind_ == OperationKindDelete {
		return Item{item: itemCore{OperationKind_: OperationKindDelete}}
	}
	item := se.createItem(true)
	player.player.Items = append(player.player.Items, item.item.ID)
	player.player.OperationKind_ = OperationKindUpdate
	se.Patch.Player[player.player.ID] = player.player
	return item
}
func (_zone Zone) AddItem(se *Engine) ZoneItem {
	zone := se.Zone(_zone.zone.ID)
	if zone.zone.OperationKind_ == OperationKindDelete {
		return ZoneItem{zoneItem: zoneItemCore{OperationKind_: OperationKindDelete}}
	}
	zoneItem := se.createZoneItem(true)
	zone.zone.Items = append(zone.zone.Items, zoneItem.zoneItem.ID)
	zone.zone.OperationKind_ = OperationKindUpdate
	se.Patch.Zone[zone.zone.ID] = zone.zone
	return zoneItem
}
func (_zone Zone) AddPlayer(se *Engine) Player {
	zone := se.Zone(_zone.zone.ID)
	if zone.zone.OperationKind_ == OperationKindDelete {
		return Player{player: playerCore{OperationKind_: OperationKindDelete}}
	}
	player := se.createPlayer(true)
	zone.zone.Players = append(zone.zone.Players, player.player.ID)
	zone.zone.OperationKind_ = OperationKindUpdate
	se.Patch.Zone[zone.zone.ID] = zone.zone
	return player
}
func (_zone Zone) AddTags(se *Engine, tags ...string) {
	zone := se.Zone(_zone.zone.ID)
	if zone.zone.OperationKind_ == OperationKindDelete {
		return
	}
	zone.zone.Tags = append(zone.zone.Tags, tags...)
	zone.zone.OperationKind_ = OperationKindUpdate
	se.Patch.Zone[zone.zone.ID] = zone.zone
}
func (se *Engine) CreateGearScore() GearScore {
	return se.createGearScore(false)
}
func (se *Engine) createGearScore(hasParent bool) GearScore {
	var gearScore gearScoreCore
	gearScore.ID = GearScoreID(se.GenerateID())
	gearScore.HasParent_ = hasParent
	gearScore.OperationKind_ = OperationKindUpdate
	se.Patch.GearScore[gearScore.ID] = gearScore
	return GearScore{gearScore: gearScore}
}
func (se *Engine) CreateItem() Item {
	return se.createItem(false)
}
func (se *Engine) createItem(hasParent bool) Item {
	var item itemCore
	item.ID = ItemID(se.GenerateID())
	item.HasParent_ = hasParent
	elementGearScore := se.createGearScore(true)
	item.GearScore = elementGearScore.gearScore.ID
	item.OperationKind_ = OperationKindUpdate
	se.Patch.Item[item.ID] = item
	return Item{item: item}
}
func (se *Engine) CreatePlayer() Player {
	return se.createPlayer(false)
}
func (se *Engine) createPlayer(hasParent bool) Player {
	var player playerCore
	player.ID = PlayerID(se.GenerateID())
	player.HasParent_ = hasParent
	elementGearScore := se.createGearScore(true)
	player.GearScore = elementGearScore.gearScore.ID
	elementPosition := se.createPosition(true)
	player.Position = elementPosition.position.ID
	player.OperationKind_ = OperationKindUpdate
	se.Patch.Player[player.ID] = player
	return Player{player: player}
}
func (se *Engine) CreatePosition() Position {
	return se.createPosition(false)
}
func (se *Engine) createPosition(hasParent bool) Position {
	var position positionCore
	position.ID = PositionID(se.GenerateID())
	position.HasParent_ = hasParent
	position.OperationKind_ = OperationKindUpdate
	se.Patch.Position[position.ID] = position
	return Position{position: position}
}
func (se *Engine) CreateZone() Zone {
	return se.createZone()
}
func (se *Engine) createZone() Zone {
	var zone zoneCore
	zone.ID = ZoneID(se.GenerateID())
	zone.OperationKind_ = OperationKindUpdate
	se.Patch.Zone[zone.ID] = zone
	return Zone{zone: zone}
}
func (se *Engine) CreateZoneItem() ZoneItem {
	return se.createZoneItem(false)
}
func (se *Engine) createZoneItem(hasParent bool) ZoneItem {
	var zoneItem zoneItemCore
	zoneItem.ID = ZoneItemID(se.GenerateID())
	zoneItem.HasParent_ = hasParent
	elementItem := se.createItem(true)
	zoneItem.Item = elementItem.item.ID
	elementPosition := se.createPosition(true)
	zoneItem.Position = elementPosition.position.ID
	zoneItem.OperationKind_ = OperationKindUpdate
	se.Patch.ZoneItem[zoneItem.ID] = zoneItem
	return ZoneItem{zoneItem: zoneItem}
}
func (se *Engine) DeleteGearScore(gearScoreID GearScoreID) {
	gearScore := se.GearScore(gearScoreID).gearScore
	if gearScore.HasParent_ {
		return
	}
	se.deleteGearScore(gearScoreID)
}
func (se *Engine) deleteGearScore(gearScoreID GearScoreID) {
	gearScore := se.GearScore(gearScoreID).gearScore
	gearScore.OperationKind_ = OperationKindDelete
	se.Patch.GearScore[gearScore.ID] = gearScore
}
func (se *Engine) DeleteItem(itemID ItemID) {
	item := se.Item(itemID).item
	if item.HasParent_ {
		return
	}
	se.deleteItem(itemID)
}
func (se *Engine) deleteItem(itemID ItemID) {
	item := se.Item(itemID).item
	item.OperationKind_ = OperationKindDelete
	se.Patch.Item[item.ID] = item
	se.deleteGearScore(item.GearScore)
}
func (se *Engine) DeletePlayer(playerID PlayerID) {
	player := se.Player(playerID).player
	if player.HasParent_ {
		return
	}
	se.deletePlayer(playerID)
}
func (se *Engine) deletePlayer(playerID PlayerID) {
	player := se.Player(playerID).player
	player.OperationKind_ = OperationKindDelete
	se.Patch.Player[player.ID] = player
	se.deleteGearScore(player.GearScore)
	for _, itemID := range player.Items {
		se.deleteItem(itemID)
	}
	se.deletePosition(player.Position)
}
func (se *Engine) DeletePosition(positionID PositionID) {
	position := se.Position(positionID).position
	if position.HasParent_ {
		return
	}
	se.deletePosition(positionID)
}
func (se *Engine) deletePosition(positionID PositionID) {
	position := se.Position(positionID).position
	position.OperationKind_ = OperationKindDelete
	se.Patch.Position[position.ID] = position
}
func (se *Engine) DeleteZone(zoneID ZoneID) {
	se.deleteZone(zoneID)
}
func (se *Engine) deleteZone(zoneID ZoneID) {
	zone := se.Zone(zoneID).zone
	zone.OperationKind_ = OperationKindDelete
	se.Patch.Zone[zone.ID] = zone
	for _, zoneItemID := range zone.Items {
		se.deleteZoneItem(zoneItemID)
	}
	for _, playerID := range zone.Players {
		se.deletePlayer(playerID)
	}
}
func (se *Engine) DeleteZoneItem(zoneItemID ZoneItemID) {
	zoneItem := se.ZoneItem(zoneItemID).zoneItem
	if zoneItem.HasParent_ {
		return
	}
	se.deleteZoneItem(zoneItemID)
}
func (se *Engine) deleteZoneItem(zoneItemID ZoneItemID) {
	zoneItem := se.ZoneItem(zoneItemID).zoneItem
	zoneItem.OperationKind_ = OperationKindDelete
	se.Patch.ZoneItem[zoneItem.ID] = zoneItem
	se.deleteItem(zoneItem.Item)
	se.deletePosition(zoneItem.Position)
}
func (se *Engine) GearScore(gearScoreID GearScoreID) GearScore {
	patchingGearScore, ok := se.Patch.GearScore[gearScoreID]
	if ok {
		return GearScore{gearScore: patchingGearScore}
	}
	currentGearScore := se.State.GearScore[gearScoreID]
	return GearScore{gearScore: currentGearScore}
}
func (_gearScore GearScore) ID(se *Engine) GearScoreID {
	return _gearScore.gearScore.ID
}
func (_gearScore GearScore) Level(se *Engine) int {
	gearScore := se.GearScore(_gearScore.gearScore.ID)
	return gearScore.gearScore.Level
}
func (_gearScore GearScore) Score(se *Engine) int {
	gearScore := se.GearScore(_gearScore.gearScore.ID)
	return gearScore.gearScore.Score
}
func (se *Engine) Item(itemID ItemID) Item {
	patchingItem, ok := se.Patch.Item[itemID]
	if ok {
		return Item{item: patchingItem}
	}
	currentItem := se.State.Item[itemID]
	return Item{item: currentItem}
}
func (_item Item) ID(se *Engine) ItemID {
	return _item.item.ID
}
func (_item Item) GearScore(se *Engine) GearScore {
	item := se.Item(_item.item.ID)
	return se.GearScore(item.item.GearScore)
}
func (se *Engine) Player(playerID PlayerID) Player {
	patchingPlayer, ok := se.Patch.Player[playerID]
	if ok {
		return Player{player: patchingPlayer}
	}
	currentPlayer := se.State.Player[playerID]
	return Player{player: currentPlayer}
}
func (_player Player) ID(se *Engine) PlayerID {
	return _player.player.ID
}
func (_player Player) GearScore(se *Engine) GearScore {
	player := se.Player(_player.player.ID)
	return se.GearScore(player.player.GearScore)
}
func (_player Player) Items(se *Engine) []Item {
	player := se.Player(_player.player.ID)
	var items []Item
	for _, itemID := range player.player.Items {
		items = append(items, se.Item(itemID))
	}
	return items
}
func (_player Player) Position(se *Engine) Position {
	player := se.Player(_player.player.ID)
	return se.Position(player.player.Position)
}
func (se *Engine) Position(positionID PositionID) Position {
	patchingPosition, ok := se.Patch.Position[positionID]
	if ok {
		return Position{position: patchingPosition}
	}
	currentPosition := se.State.Position[positionID]
	return Position{position: currentPosition}
}
func (_position Position) ID(se *Engine) PositionID {
	return _position.position.ID
}
func (_position Position) X(se *Engine) float64 {
	position := se.Position(_position.position.ID)
	return position.position.X
}
func (_position Position) Y(se *Engine) float64 {
	position := se.Position(_position.position.ID)
	return position.position.Y
}
func (se *Engine) Zone(zoneID ZoneID) Zone {
	patchingZone, ok := se.Patch.Zone[zoneID]
	if ok {
		return Zone{zone: patchingZone}
	}
	currentZone := se.State.Zone[zoneID]
	return Zone{zone: currentZone}
}
func (_zone Zone) ID(se *Engine) ZoneID {
	return _zone.zone.ID
}
func (_zone Zone) Items(se *Engine) []ZoneItem {
	zone := se.Zone(_zone.zone.ID)
	var items []ZoneItem
	for _, zoneItemID := range zone.zone.Items {
		items = append(items, se.ZoneItem(zoneItemID))
	}
	return items
}
func (_zone Zone) Players(se *Engine) []Player {
	zone := se.Zone(_zone.zone.ID)
	var players []Player
	for _, playerID := range zone.zone.Players {
		players = append(players, se.Player(playerID))
	}
	return players
}
func (_zone Zone) Tags(se *Engine) []string {
	zone := se.Zone(_zone.zone.ID)
	var tags []string
	for _, element := range zone.zone.Tags {
		tags = append(tags, element)
	}
	return tags
}
func (se *Engine) ZoneItem(zoneItemID ZoneItemID) ZoneItem {
	patchingZoneItem, ok := se.Patch.ZoneItem[zoneItemID]
	if ok {
		return ZoneItem{zoneItem: patchingZoneItem}
	}
	currentZoneItem := se.State.ZoneItem[zoneItemID]
	return ZoneItem{zoneItem: currentZoneItem}
}
func (_zoneItem ZoneItem) ID(se *Engine) ZoneItemID {
	return _zoneItem.zoneItem.ID
}
func (_zoneItem ZoneItem) Item(se *Engine) Item {
	zoneItem := se.ZoneItem(_zoneItem.zoneItem.ID)
	return se.Item(zoneItem.zoneItem.Item)
}
func (_zoneItem ZoneItem) Position(se *Engine) Position {
	zoneItem := se.ZoneItem(_zoneItem.zoneItem.ID)
	return se.Position(zoneItem.zoneItem.Position)
}
func (_player Player) RemoveItems(se *Engine, itemsToRemove ...ItemID) Player {
	player := se.Player(_player.player.ID)
	if player.player.OperationKind_ == OperationKindDelete {
		return player
	}
	var wereElementsAltered bool
	var newElements []ItemID
	for _, element := range player.player.Items {
		var toBeRemoved bool
		for _, elementToRemove := range itemsToRemove {
			if element == elementToRemove {
				toBeRemoved = true
				wereElementsAltered = true
				se.deleteItem(element)
				break
			}
		}
		if !toBeRemoved {
			newElements = append(newElements, element)
		}
	}
	if !wereElementsAltered {
		return player
	}
	player.player.Items = newElements
	player.player.OperationKind_ = OperationKindUpdate
	se.Patch.Player[player.player.ID] = player.player
	return player
}
func (_zone Zone) RemoveItems(se *Engine, itemsToRemove ...ZoneItemID) Zone {
	zone := se.Zone(_zone.zone.ID)
	if zone.zone.OperationKind_ == OperationKindDelete {
		return zone
	}
	var wereElementsAltered bool
	var newElements []ZoneItemID
	for _, element := range zone.zone.Items {
		var toBeRemoved bool
		for _, elementToRemove := range itemsToRemove {
			if element == elementToRemove {
				toBeRemoved = true
				wereElementsAltered = true
				se.deleteZoneItem(element)
				break
			}
		}
		if !toBeRemoved {
			newElements = append(newElements, element)
		}
	}
	if !wereElementsAltered {
		return zone
	}
	zone.zone.Items = newElements
	zone.zone.OperationKind_ = OperationKindUpdate
	se.Patch.Zone[zone.zone.ID] = zone.zone
	return zone
}
func (_zone Zone) RemovePlayers(se *Engine, playersToRemove ...PlayerID) Zone {
	zone := se.Zone(_zone.zone.ID)
	if zone.zone.OperationKind_ == OperationKindDelete {
		return zone
	}
	var wereElementsAltered bool
	var newElements []PlayerID
	for _, element := range zone.zone.Players {
		var toBeRemoved bool
		for _, elementToRemove := range playersToRemove {
			if element == elementToRemove {
				toBeRemoved = true
				wereElementsAltered = true
				se.deletePlayer(element)
				break
			}
		}
		if !toBeRemoved {
			newElements = append(newElements, element)
		}
	}
	if !wereElementsAltered {
		return zone
	}
	zone.zone.Players = newElements
	zone.zone.OperationKind_ = OperationKindUpdate
	se.Patch.Zone[zone.zone.ID] = zone.zone
	return zone
}
func (_zone Zone) RemoveTags(se *Engine, tagsToRemove ...string) Zone {
	zone := se.Zone(_zone.zone.ID)
	if zone.zone.OperationKind_ == OperationKindDelete {
		return zone
	}
	var wereElementsAltered bool
	var newElements []string
	for _, element := range zone.zone.Tags {
		var toBeRemoved bool
		for _, elementToRemove := range tagsToRemove {
			if element == elementToRemove {
				toBeRemoved = true
				wereElementsAltered = true
				break
			}
		}
		if !toBeRemoved {
			newElements = append(newElements, element)
		}
	}
	if !wereElementsAltered {
		return zone
	}
	zone.zone.Tags = newElements
	zone.zone.OperationKind_ = OperationKindUpdate
	se.Patch.Zone[zone.zone.ID] = zone.zone
	return zone
}
func (_gearScore GearScore) SetLevel(se *Engine, newLevel int) GearScore {
	gearScore := se.GearScore(_gearScore.gearScore.ID)
	if gearScore.gearScore.OperationKind_ == OperationKindDelete {
		return gearScore
	}
	gearScore.gearScore.Level = newLevel
	gearScore.gearScore.OperationKind_ = OperationKindUpdate
	se.Patch.GearScore[gearScore.gearScore.ID] = gearScore.gearScore
	return gearScore
}
func (_gearScore GearScore) SetScore(se *Engine, newScore int) GearScore {
	gearScore := se.GearScore(_gearScore.gearScore.ID)
	if gearScore.gearScore.OperationKind_ == OperationKindDelete {
		return gearScore
	}
	gearScore.gearScore.Score = newScore
	gearScore.gearScore.OperationKind_ = OperationKindUpdate
	se.Patch.GearScore[gearScore.gearScore.ID] = gearScore.gearScore
	return gearScore
}
func (_position Position) SetX(se *Engine, newX float64) Position {
	position := se.Position(_position.position.ID)
	if position.position.OperationKind_ == OperationKindDelete {
		return position
	}
	position.position.X = newX
	position.position.OperationKind_ = OperationKindUpdate
	se.Patch.Position[position.position.ID] = position.position
	return position
}
func (_position Position) SetY(se *Engine, newY float64) Position {
	position := se.Position(_position.position.ID)
	if position.position.OperationKind_ == OperationKindDelete {
		return position
	}
	position.position.Y = newY
	position.position.OperationKind_ = OperationKindUpdate
	se.Patch.Position[position.position.ID] = position.position
	return position
}

type Tree struct {
	GearScore map[GearScoreID]tGearScore `json:"gearScore"`
	Item      map[ItemID]tItem           `json:"item"`
	Player    map[PlayerID]tPlayer       `json:"player"`
	Position  map[PositionID]tPosition   `json:"position"`
	Zone      map[ZoneID]tZone           `json:"zone"`
	ZoneItem  map[ZoneItemID]tZoneItem   `json:"zoneItem"`
}

func newTree() Tree {
	return Tree{GearScore: make(map[GearScoreID]tGearScore), Item: make(map[ItemID]tItem), Player: make(map[PlayerID]tPlayer), Position: make(map[PositionID]tPosition), Zone: make(map[ZoneID]tZone), ZoneItem: make(map[ZoneItemID]tZoneItem)}
}

type tGearScore struct {
	ID             GearScoreID   `json:"id"`
	Level          int           `json:"level"`
	Score          int           `json:"score"`
	OperationKind_ OperationKind `json:"operationKind_"`
}
type tItem struct {
	ID             ItemID        `json:"id"`
	GearScore      *tGearScore   `json:"gearScore"`
	OperationKind_ OperationKind `json:"operationKind_"`
}
type tPlayer struct {
	ID             PlayerID      `json:"id"`
	GearScore      *tGearScore   `json:"gearScore"`
	Items          []tItem       `json:"items"`
	Position       *tPosition    `json:"position"`
	OperationKind_ OperationKind `json:"operationKind_"`
}
type tPosition struct {
	ID             PositionID    `json:"id"`
	X              float64       `json:"x"`
	Y              float64       `json:"y"`
	OperationKind_ OperationKind `json:"operationKind_"`
}
type tZone struct {
	ID             ZoneID        `json:"id"`
	Items          []tZoneItem   `json:"items"`
	Players        []tPlayer     `json:"players"`
	Tags           []string      `json:"tags"`
	OperationKind_ OperationKind `json:"operationKind_"`
}
type tZoneItem struct {
	ID             ZoneItemID    `json:"id"`
	Item           *tItem        `json:"item"`
	Position       *tPosition    `json:"position"`
	OperationKind_ OperationKind `json:"operationKind_"`
}

func (se *Engine) assembleTree() Tree {
	tree := newTree()
	for _, gearScore := range se.Patch.GearScore {
		if !gearScore.HasParent_ {
			treeGearScore, hasUpdated := se.assembleGearScore(gearScore.ID)
			if hasUpdated {
				tree.GearScore[gearScore.ID] = treeGearScore
			}
		}
	}
	for _, item := range se.Patch.Item {
		if !item.HasParent_ {
			treeItem, hasUpdated := se.assembleItem(item.ID)
			if hasUpdated {
				tree.Item[item.ID] = treeItem
			}
		}
	}
	for _, player := range se.Patch.Player {
		if !player.HasParent_ {
			treePlayer, hasUpdated := se.assemblePlayer(player.ID)
			if hasUpdated {
				tree.Player[player.ID] = treePlayer
			}
		}
	}
	for _, position := range se.Patch.Position {
		if !position.HasParent_ {
			treePosition, hasUpdated := se.assemblePosition(position.ID)
			if hasUpdated {
				tree.Position[position.ID] = treePosition
			}
		}
	}
	for _, zone := range se.Patch.Zone {
		treeZone, hasUpdated := se.assembleZone(zone.ID)
		if hasUpdated {
			tree.Zone[zone.ID] = treeZone
		}
	}
	for _, zoneItem := range se.Patch.ZoneItem {
		if !zoneItem.HasParent_ {
			treeZoneItem, hasUpdated := se.assembleZoneItem(zoneItem.ID)
			if hasUpdated {
				tree.ZoneItem[zoneItem.ID] = treeZoneItem
			}
		}
	}
	for _, gearScore := range se.State.GearScore {
		if !gearScore.HasParent_ {
			if _, ok := tree.GearScore[gearScore.ID]; !ok {
				treeGearScore, hasUpdated := se.assembleGearScore(gearScore.ID)
				if hasUpdated {
					tree.GearScore[gearScore.ID] = treeGearScore
				}
			}
		}
	}
	for _, item := range se.State.Item {
		if !item.HasParent_ {
			if _, ok := tree.Item[item.ID]; !ok {
				treeItem, hasUpdated := se.assembleItem(item.ID)
				if hasUpdated {
					tree.Item[item.ID] = treeItem
				}
			}
		}
	}
	for _, player := range se.State.Player {
		if !player.HasParent_ {
			if _, ok := tree.Player[player.ID]; !ok {
				treePlayer, hasUpdated := se.assemblePlayer(player.ID)
				if hasUpdated {
					tree.Player[player.ID] = treePlayer
				}
			}
		}
	}
	for _, position := range se.State.Position {
		if !position.HasParent_ {
			if _, ok := tree.Position[position.ID]; !ok {
				treePosition, hasUpdated := se.assemblePosition(position.ID)
				if hasUpdated {
					tree.Position[position.ID] = treePosition
				}
			}
		}
	}
	for _, zone := range se.State.Zone {
		if _, ok := tree.Zone[zone.ID]; !ok {
			treeZone, hasUpdated := se.assembleZone(zone.ID)
			if hasUpdated {
				tree.Zone[zone.ID] = treeZone
			}
		}
	}
	for _, zoneItem := range se.State.ZoneItem {
		if !zoneItem.HasParent_ {
			if _, ok := tree.ZoneItem[zoneItem.ID]; !ok {
				treeZoneItem, hasUpdated := se.assembleZoneItem(zoneItem.ID)
				if hasUpdated {
					tree.ZoneItem[zoneItem.ID] = treeZoneItem
				}
			}
		}
	}
	return tree
}
func (se *Engine) assembleGearScore(gearScoreID GearScoreID) (tGearScore, bool) {
	gearScore, hasUpdated := se.Patch.GearScore[gearScoreID]
	if !hasUpdated {
		return tGearScore{}, false
	}
	var treeGearScore tGearScore
	treeGearScore.ID = gearScore.ID
	treeGearScore.OperationKind_ = gearScore.OperationKind_
	treeGearScore.Level = gearScore.Level
	treeGearScore.Score = gearScore.Score
	return treeGearScore, true
}
func (se *Engine) assembleItem(itemID ItemID) (tItem, bool) {
	item, hasUpdated := se.Patch.Item[itemID]
	if !hasUpdated {
		item = se.State.Item[itemID]
	}
	var treeItem tItem
	if treeGearScore, gearScoreHasUpdated := se.assembleGearScore(item.GearScore); gearScoreHasUpdated {
		hasUpdated = true
		treeItem.GearScore = &treeGearScore
	}
	treeItem.ID = item.ID
	treeItem.OperationKind_ = item.OperationKind_
	return treeItem, hasUpdated
}
func (se *Engine) assemblePlayer(playerID PlayerID) (tPlayer, bool) {
	player, hasUpdated := se.Patch.Player[playerID]
	if !hasUpdated {
		player = se.State.Player[playerID]
	}
	var treePlayer tPlayer
	if treeGearScore, gearScoreHasUpdated := se.assembleGearScore(player.GearScore); gearScoreHasUpdated {
		hasUpdated = true
		treePlayer.GearScore = &treeGearScore
	}
	for _, itemID := range deduplicateItemIDs(se.State.Player[player.ID].Items, se.Patch.Player[player.ID].Items) {
		if treeItem, itemHasUpdated := se.assembleItem(itemID); itemHasUpdated {
			hasUpdated = true
			treePlayer.Items = append(treePlayer.Items, treeItem)
		}
	}
	if treePosition, positionHasUpdated := se.assemblePosition(player.Position); positionHasUpdated {
		hasUpdated = true
		treePlayer.Position = &treePosition
	}
	treePlayer.ID = player.ID
	treePlayer.OperationKind_ = player.OperationKind_
	return treePlayer, hasUpdated
}
func (se *Engine) assemblePosition(positionID PositionID) (tPosition, bool) {
	position, hasUpdated := se.Patch.Position[positionID]
	if !hasUpdated {
		return tPosition{}, false
	}
	var treePosition tPosition
	treePosition.ID = position.ID
	treePosition.OperationKind_ = position.OperationKind_
	treePosition.X = position.X
	treePosition.Y = position.Y
	return treePosition, true
}
func (se *Engine) assembleZone(zoneID ZoneID) (tZone, bool) {
	zone, hasUpdated := se.Patch.Zone[zoneID]
	if !hasUpdated {
		zone = se.State.Zone[zoneID]
	}
	var treeZone tZone
	for _, zoneItemID := range deduplicateZoneItemIDs(se.State.Zone[zone.ID].Items, se.Patch.Zone[zone.ID].Items) {
		if treeZoneItem, zoneItemHasUpdated := se.assembleZoneItem(zoneItemID); zoneItemHasUpdated {
			hasUpdated = true
			treeZone.Items = append(treeZone.Items, treeZoneItem)
		}
	}
	for _, playerID := range deduplicatePlayerIDs(se.State.Zone[zone.ID].Players, se.Patch.Zone[zone.ID].Players) {
		if treePlayer, playerHasUpdated := se.assemblePlayer(playerID); playerHasUpdated {
			hasUpdated = true
			treeZone.Players = append(treeZone.Players, treePlayer)
		}
	}
	treeZone.ID = zone.ID
	treeZone.OperationKind_ = zone.OperationKind_
	treeZone.Tags = zone.Tags
	return treeZone, hasUpdated
}
func (se *Engine) assembleZoneItem(zoneItemID ZoneItemID) (tZoneItem, bool) {
	zoneItem, hasUpdated := se.Patch.ZoneItem[zoneItemID]
	if !hasUpdated {
		zoneItem = se.State.ZoneItem[zoneItemID]
	}
	var treeZoneItem tZoneItem
	if treeItem, itemHasUpdated := se.assembleItem(zoneItem.Item); itemHasUpdated {
		hasUpdated = true
		treeZoneItem.Item = &treeItem
	}
	if treePosition, positionHasUpdated := se.assemblePosition(zoneItem.Position); positionHasUpdated {
		hasUpdated = true
		treeZoneItem.Position = &treePosition
	}
	treeZoneItem.ID = zoneItem.ID
	treeZoneItem.OperationKind_ = zoneItem.OperationKind_
	return treeZoneItem, hasUpdated
}
func deduplicateGearScoreIDs(a []GearScoreID, b []GearScoreID) []GearScoreID {
	check := make(map[GearScoreID]bool)
	deduped := make([]GearScoreID, 0)
	for _, val := range a {
		check[val] = true
	}
	for _, val := range b {
		check[val] = true
	}
	for val := range check {
		deduped = append(deduped, val)
	}
	return deduped
}
func deduplicateItemIDs(a []ItemID, b []ItemID) []ItemID {
	check := make(map[ItemID]bool)
	deduped := make([]ItemID, 0)
	for _, val := range a {
		check[val] = true
	}
	for _, val := range b {
		check[val] = true
	}
	for val := range check {
		deduped = append(deduped, val)
	}
	return deduped
}
func deduplicatePlayerIDs(a []PlayerID, b []PlayerID) []PlayerID {
	check := make(map[PlayerID]bool)
	deduped := make([]PlayerID, 0)
	for _, val := range a {
		check[val] = true
	}
	for _, val := range b {
		check[val] = true
	}
	for val := range check {
		deduped = append(deduped, val)
	}
	return deduped
}
func deduplicatePositionIDs(a []PositionID, b []PositionID) []PositionID {
	check := make(map[PositionID]bool)
	deduped := make([]PositionID, 0)
	for _, val := range a {
		check[val] = true
	}
	for _, val := range b {
		check[val] = true
	}
	for val := range check {
		deduped = append(deduped, val)
	}
	return deduped
}
func deduplicateZoneIDs(a []ZoneID, b []ZoneID) []ZoneID {
	check := make(map[ZoneID]bool)
	deduped := make([]ZoneID, 0)
	for _, val := range a {
		check[val] = true
	}
	for _, val := range b {
		check[val] = true
	}
	for val := range check {
		deduped = append(deduped, val)
	}
	return deduped
}
func deduplicateZoneItemIDs(a []ZoneItemID, b []ZoneItemID) []ZoneItemID {
	check := make(map[ZoneItemID]bool)
	deduped := make([]ZoneItemID, 0)
	for _, val := range a {
		check[val] = true
	}
	for _, val := range b {
		check[val] = true
	}
	for val := range check {
		deduped = append(deduped, val)
	}
	return deduped
}
