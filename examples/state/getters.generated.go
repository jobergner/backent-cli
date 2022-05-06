package state

import (
	"sort"
)

func (engine *Engine) boolValue(boolValueID BoolValueID) boolValue {
	patchingBoolValue, ok := engine.Patch.BoolValue[boolValueID]
	if ok {
		return patchingBoolValue
	}
	return engine.State.BoolValue[boolValueID]
}

func (engine *Engine) intValue(intValueID IntValueID) intValue {
	patchingIntValue, ok := engine.Patch.IntValue[intValueID]
	if ok {
		return patchingIntValue
	}
	return engine.State.IntValue[intValueID]
}

func (engine *Engine) floatValue(floatValueID FloatValueID) floatValue {
	patchingFloatValue, ok := engine.Patch.FloatValue[floatValueID]
	if ok {
		return patchingFloatValue
	}
	return engine.State.FloatValue[floatValueID]
}

func (engine *Engine) stringValue(stringValueID StringValueID) stringValue {
	patchingStringValue, ok := engine.Patch.StringValue[stringValueID]
	if ok {
		return patchingStringValue
	}
	return engine.State.StringValue[stringValueID]
}

func (engine *Engine) QueryPlayers(matcher func(Player) bool) []Player {
	playerIDs := engine.allPlayerIDs()
	sort.Slice(playerIDs, func(i, j int) bool {
		return playerIDs[i] < playerIDs[j]
	})
	var players []Player
	for _, playerID := range playerIDs {
		player := engine.Player(playerID)
		if matcher(player) {
			players = append(players, player)
		}
	}
	playerIDSlicePool.Put(playerIDs)
	return players
}

func (engine *Engine) EveryPlayer() []Player {
	playerIDs := engine.allPlayerIDs()
	sort.Slice(playerIDs, func(i, j int) bool {
		return playerIDs[i] < playerIDs[j]
	})
	var players []Player
	for _, playerID := range playerIDs {
		player := engine.Player(playerID)
		if player.player.HasParent {
			continue
		}
		players = append(players, player)
	}
	playerIDSlicePool.Put(playerIDs)
	return players
}

func (engine *Engine) Player(playerID PlayerID) Player {
	patchingPlayer, ok := engine.Patch.Player[playerID]
	if ok {
		return Player{player: patchingPlayer}
	}
	currentPlayer, ok := engine.State.Player[playerID]
	if ok {
		return Player{player: currentPlayer}
	}
	return Player{player: playerCore{OperationKind: OperationKindDelete, engine: engine}}
}

func (_player Player) ParentItem() Item {
	player := _player.player.engine.Player(_player.player.ID)
	if !player.player.HasParent {
		return Item{item: itemCore{OperationKind: OperationKindDelete, engine: player.player.engine}}
	}
	parentSeg := player.player.Path[len(player.player.Path)-2]
	return player.player.engine.Item(ItemID(parentSeg.ID))
}

func (_player Player) ParentZone() Zone {
	player := _player.player.engine.Player(_player.player.ID)
	if !player.player.HasParent {
		return Zone{zone: zoneCore{OperationKind: OperationKindDelete, engine: player.player.engine}}
	}
	parentSeg := player.player.Path[len(player.player.Path)-2]
	return player.player.engine.Zone(ZoneID(parentSeg.ID))
}

func (_player Player) ParentKind() (ElementKind, bool) {
	if !_player.player.HasParent {
		return "", false
	}
	return _player.player.Path[len(_player.player.Path)-2].Kind, true
}

func (_player Player) ID() PlayerID {
	return _player.player.ID
}

func (_player Player) Exists() (Player, bool) {
	player := _player.player.engine.Player(_player.player.ID)
	return player, player.player.OperationKind != OperationKindDelete
}

func (_player Player) Path() string {
	return _player.player.JSONPath
}

func (_player Player) Target() PlayerTargetRef {
	player := _player.player.engine.Player(_player.player.ID)
	return player.player.engine.playerTargetRef(player.player.Target)
}

func (_player Player) TargetedBy() []PlayerTargetedByRef {
	player := _player.player.engine.Player(_player.player.ID)
	var targetedBy []PlayerTargetedByRef
	for _, refID := range player.player.TargetedBy {
		targetedBy = append(targetedBy, player.player.engine.playerTargetedByRef(refID))
	}
	return targetedBy
}

func (_player Player) Action() []AttackEvent {
	player := _player.player.engine.Player(_player.player.ID)
	var action []AttackEvent
	for _, attackEventID := range player.player.Action {
		action = append(action, player.player.engine.AttackEvent(attackEventID))
	}
	return action
}

func (_player Player) Items() []Item {
	player := _player.player.engine.Player(_player.player.ID)
	var items []Item
	for _, itemID := range player.player.Items {
		items = append(items, player.player.engine.Item(itemID))
	}
	return items
}

func (_player Player) GearScore() GearScore {
	player := _player.player.engine.Player(_player.player.ID)
	return player.player.engine.GearScore(player.player.GearScore)
}

func (_player Player) GuildMembers() []PlayerGuildMemberRef {
	player := _player.player.engine.Player(_player.player.ID)
	var guildMembers []PlayerGuildMemberRef
	for _, refID := range player.player.GuildMembers {
		guildMembers = append(guildMembers, player.player.engine.playerGuildMemberRef(refID))
	}
	return guildMembers
}

func (_player Player) EquipmentSets() []PlayerEquipmentSetRef {
	player := _player.player.engine.Player(_player.player.ID)
	var equipmentSets []PlayerEquipmentSetRef
	for _, refID := range player.player.EquipmentSets {
		equipmentSets = append(equipmentSets, player.player.engine.playerEquipmentSetRef(refID))
	}
	return equipmentSets
}

func (_player Player) Position() Position {
	player := _player.player.engine.Player(_player.player.ID)
	return player.player.engine.Position(player.player.Position)
}

func (engine *Engine) QueryGearScores(matcher func(GearScore) bool) []GearScore {
	gearScoreIDs := engine.allGearScoreIDs()
	sort.Slice(gearScoreIDs, func(i, j int) bool {
		return gearScoreIDs[i] < gearScoreIDs[j]
	})
	var gearScores []GearScore
	for _, gearScoreID := range gearScoreIDs {
		gearScore := engine.GearScore(gearScoreID)
		if matcher(gearScore) {
			gearScores = append(gearScores, gearScore)
		}
	}
	gearScoreIDSlicePool.Put(gearScoreIDs)
	return gearScores
}

func (engine *Engine) EveryGearScore() []GearScore {
	gearScoreIDs := engine.allGearScoreIDs()
	sort.Slice(gearScoreIDs, func(i, j int) bool {
		return gearScoreIDs[i] < gearScoreIDs[j]
	})
	var gearScores []GearScore
	for _, gearScoreID := range gearScoreIDs {
		gearScore := engine.GearScore(gearScoreID)
		if gearScore.gearScore.HasParent {
			continue
		}
		gearScores = append(gearScores, gearScore)
	}
	gearScoreIDSlicePool.Put(gearScoreIDs)
	return gearScores
}

func (engine *Engine) GearScore(gearScoreID GearScoreID) GearScore {
	patchingGearScore, ok := engine.Patch.GearScore[gearScoreID]
	if ok {
		return GearScore{gearScore: patchingGearScore}
	}
	currentGearScore, ok := engine.State.GearScore[gearScoreID]
	if ok {
		return GearScore{gearScore: currentGearScore}
	}
	return GearScore{gearScore: gearScoreCore{OperationKind: OperationKindDelete, engine: engine}}
}

func (_gearScore GearScore) ParentItem() Item {
	gearScore := _gearScore.gearScore.engine.GearScore(_gearScore.gearScore.ID)
	if !gearScore.gearScore.HasParent {
		return Item{item: itemCore{OperationKind: OperationKindDelete, engine: gearScore.gearScore.engine}}
	}
	parentSeg := gearScore.gearScore.Path[len(gearScore.gearScore.Path)-2]
	return gearScore.gearScore.engine.Item(ItemID(parentSeg.ID))
}

func (_gearScore GearScore) ParentPlayer() Player {
	gearScore := _gearScore.gearScore.engine.GearScore(_gearScore.gearScore.ID)
	if !gearScore.gearScore.HasParent {
		return Player{player: playerCore{OperationKind: OperationKindDelete, engine: gearScore.gearScore.engine}}
	}
	parentSeg := gearScore.gearScore.Path[len(gearScore.gearScore.Path)-2]
	return gearScore.gearScore.engine.Player(PlayerID(parentSeg.ID))
}

func (_gearScore GearScore) ParentKind() (ElementKind, bool) {
	if !_gearScore.gearScore.HasParent {
		return "", false
	}
	return _gearScore.gearScore.Path[len(_gearScore.gearScore.Path)-2].Kind, true
}

func (_gearScore GearScore) ID() GearScoreID {
	return _gearScore.gearScore.ID
}

func (_gearScore GearScore) Exists() (GearScore, bool) {
	gearScore := _gearScore.gearScore.engine.GearScore(_gearScore.gearScore.ID)
	return gearScore, gearScore.gearScore.OperationKind != OperationKindDelete
}

func (_gearScore GearScore) Path() string {
	return _gearScore.gearScore.JSONPath
}

func (_gearScore GearScore) Level() int64 {
	gearScore := _gearScore.gearScore.engine.GearScore(_gearScore.gearScore.ID)
	return gearScore.gearScore.engine.intValue(gearScore.gearScore.Level).Value
}

func (_gearScore GearScore) Score() int64 {
	gearScore := _gearScore.gearScore.engine.GearScore(_gearScore.gearScore.ID)
	return gearScore.gearScore.engine.intValue(gearScore.gearScore.Score).Value
}

func (engine *Engine) QueryItems(matcher func(Item) bool) []Item {
	itemIDs := engine.allItemIDs()
	sort.Slice(itemIDs, func(i, j int) bool {
		return itemIDs[i] < itemIDs[j]
	})
	var items []Item
	for _, itemID := range itemIDs {
		item := engine.Item(itemID)
		if matcher(item) {
			items = append(items, item)
		}
	}
	itemIDSlicePool.Put(itemIDs)
	return items
}

func (engine *Engine) EveryItem() []Item {
	itemIDs := engine.allItemIDs()
	sort.Slice(itemIDs, func(i, j int) bool {
		return itemIDs[i] < itemIDs[j]
	})
	var items []Item
	for _, itemID := range itemIDs {
		item := engine.Item(itemID)
		if item.item.HasParent {
			continue
		}
		items = append(items, item)
	}
	itemIDSlicePool.Put(itemIDs)
	return items
}

func (engine *Engine) Item(itemID ItemID) Item {
	patchingItem, ok := engine.Patch.Item[itemID]
	if ok {
		return Item{item: patchingItem}
	}
	currentItem, ok := engine.State.Item[itemID]
	if ok {
		return Item{item: currentItem}
	}
	return Item{item: itemCore{OperationKind: OperationKindDelete, engine: engine}}
}

func (_item Item) ParentPlayer() Player {
	item := _item.item.engine.Item(_item.item.ID)
	if !item.item.HasParent {
		return Player{player: playerCore{OperationKind: OperationKindDelete, engine: item.item.engine}}
	}
	parentSeg := item.item.Path[len(item.item.Path)-2]
	return item.item.engine.Player(PlayerID(parentSeg.ID))
}

func (_item Item) ParentZone() Zone {
	item := _item.item.engine.Item(_item.item.ID)
	if !item.item.HasParent {
		return Zone{zone: zoneCore{OperationKind: OperationKindDelete, engine: item.item.engine}}
	}
	parentSeg := item.item.Path[len(item.item.Path)-2]
	return item.item.engine.Zone(ZoneID(parentSeg.ID))
}

func (_item Item) ParentZoneItem() ZoneItem {
	item := _item.item.engine.Item(_item.item.ID)
	if !item.item.HasParent {
		return ZoneItem{zoneItem: zoneItemCore{OperationKind: OperationKindDelete, engine: item.item.engine}}
	}
	parentSeg := item.item.Path[len(item.item.Path)-2]
	return item.item.engine.ZoneItem(ZoneItemID(parentSeg.ID))
}

func (_item Item) ParentKind() (ElementKind, bool) {
	if !_item.item.HasParent {
		return "", false
	}
	return _item.item.Path[len(_item.item.Path)-2].Kind, true
}

func (_item Item) ID() ItemID {
	return _item.item.ID
}

func (_item Item) Exists() (Item, bool) {
	item := _item.item.engine.Item(_item.item.ID)
	return item, item.item.OperationKind != OperationKindDelete
}

func (_item Item) Path() string {
	return _item.item.JSONPath
}
func (_item Item) Name() string {
	item := _item.item.engine.Item(_item.item.ID)
	return item.item.engine.stringValue(item.item.Name).Value
}

func (_item Item) GearScore() GearScore {
	item := _item.item.engine.Item(_item.item.ID)
	return item.item.engine.GearScore(item.item.GearScore)
}

func (_item Item) BoundTo() ItemBoundToRef {
	item := _item.item.engine.Item(_item.item.ID)
	return item.item.engine.itemBoundToRef(item.item.BoundTo)
}

func (_item Item) Origin() AnyOfPlayer_Position {
	item := _item.item.engine.Item(_item.item.ID)
	return item.item.engine.anyOfPlayer_Position(item.item.Origin)
}

func (engine *Engine) QueryAttackEvents(matcher func(AttackEvent) bool) []AttackEvent {
	attackEventIDs := engine.allAttackEventIDs()
	sort.Slice(attackEventIDs, func(i, j int) bool {
		return attackEventIDs[i] < attackEventIDs[j]
	})
	var attackEvents []AttackEvent
	for _, attackEventID := range attackEventIDs {
		attackEvent := engine.AttackEvent(attackEventID)
		if matcher(attackEvent) {
			attackEvents = append(attackEvents, attackEvent)
		}
	}
	attackEventIDSlicePool.Put(attackEventIDs)
	return attackEvents
}

func (engine *Engine) EveryAttackEvent() []AttackEvent {
	attackEventIDs := engine.allAttackEventIDs()
	sort.Slice(attackEventIDs, func(i, j int) bool {
		return attackEventIDs[i] < attackEventIDs[j]
	})
	var attackEvents []AttackEvent
	for _, attackEventID := range attackEventIDs {
		attackEvent := engine.AttackEvent(attackEventID)
		if attackEvent.attackEvent.HasParent {
			continue
		}
		attackEvents = append(attackEvents, attackEvent)
	}
	attackEventIDSlicePool.Put(attackEventIDs)
	return attackEvents
}

func (engine *Engine) AttackEvent(attackEventID AttackEventID) AttackEvent {
	patchingAttackEvent, ok := engine.Patch.AttackEvent[attackEventID]
	if ok {
		return AttackEvent{attackEvent: patchingAttackEvent}
	}
	currentAttackEvent, ok := engine.State.AttackEvent[attackEventID]
	if ok {
		return AttackEvent{attackEvent: currentAttackEvent}
	}
	return AttackEvent{attackEvent: attackEventCore{OperationKind: OperationKindDelete, engine: engine}}
}

func (_attackEvent AttackEvent) ParentPlayer() Player {
	attackEvent := _attackEvent.attackEvent.engine.AttackEvent(_attackEvent.attackEvent.ID)
	if !attackEvent.attackEvent.HasParent {
		return Player{player: playerCore{OperationKind: OperationKindDelete, engine: attackEvent.attackEvent.engine}}
	}
	parentSeg := attackEvent.attackEvent.Path[len(attackEvent.attackEvent.Path)-2]
	return attackEvent.attackEvent.engine.Player(PlayerID(parentSeg.ID))
}

func (_attackEvent AttackEvent) ParentKind() (ElementKind, bool) {
	if !_attackEvent.attackEvent.HasParent {
		return "", false
	}
	return _attackEvent.attackEvent.Path[len(_attackEvent.attackEvent.Path)-2].Kind, true
}

func (_attackEvent AttackEvent) ID() AttackEventID {
	return _attackEvent.attackEvent.ID
}

func (_attackEvent AttackEvent) Exists() (AttackEvent, bool) {
	attackEvent := _attackEvent.attackEvent.engine.AttackEvent(_attackEvent.attackEvent.ID)
	return attackEvent, attackEvent.attackEvent.OperationKind != OperationKindDelete
}

func (_attackEvent AttackEvent) Path() string {
	return _attackEvent.attackEvent.JSONPath
}

func (_attackEvent AttackEvent) Target() AttackEventTargetRef {
	attackEvent := _attackEvent.attackEvent.engine.AttackEvent(_attackEvent.attackEvent.ID)
	return attackEvent.attackEvent.engine.attackEventTargetRef(attackEvent.attackEvent.Target)
}

func (engine *Engine) QueryPositions(matcher func(Position) bool) []Position {
	positionIDs := engine.allPositionIDs()
	sort.Slice(positionIDs, func(i, j int) bool {
		return positionIDs[i] < positionIDs[j]
	})
	var positions []Position
	for _, positionID := range positionIDs {
		position := engine.Position(positionID)
		if matcher(position) {
			positions = append(positions, position)
		}
	}
	positionIDSlicePool.Put(positionIDs)
	return positions
}

func (engine *Engine) EveryPosition() []Position {
	positionIDs := engine.allPositionIDs()
	sort.Slice(positionIDs, func(i, j int) bool {
		return positionIDs[i] < positionIDs[j]
	})
	var positions []Position
	for _, positionID := range positionIDs {
		position := engine.Position(positionID)
		if position.position.HasParent {
			continue
		}
		positions = append(positions, position)
	}
	positionIDSlicePool.Put(positionIDs)
	return positions
}

func (engine *Engine) Position(positionID PositionID) Position {
	patchingPosition, ok := engine.Patch.Position[positionID]
	if ok {
		return Position{position: patchingPosition}
	}
	currentPosition, ok := engine.State.Position[positionID]
	if ok {
		return Position{position: currentPosition}
	}
	return Position{position: positionCore{OperationKind: OperationKindDelete, engine: engine}}
}

func (_position Position) ParentItem() Item {
	position := _position.position.engine.Position(_position.position.ID)
	if !position.position.HasParent {
		return Item{item: itemCore{OperationKind: OperationKindDelete, engine: position.position.engine}}
	}
	parentSeg := position.position.Path[len(position.position.Path)-2]
	return position.position.engine.Item(ItemID(parentSeg.ID))
}

func (_position Position) ParentPlayer() Player {
	position := _position.position.engine.Position(_position.position.ID)
	if !position.position.HasParent {
		return Player{player: playerCore{OperationKind: OperationKindDelete, engine: position.position.engine}}
	}
	parentSeg := position.position.Path[len(position.position.Path)-2]
	return position.position.engine.Player(PlayerID(parentSeg.ID))
}

func (_position Position) ParentZoneItem() ZoneItem {
	position := _position.position.engine.Position(_position.position.ID)
	if !position.position.HasParent {
		return ZoneItem{zoneItem: zoneItemCore{OperationKind: OperationKindDelete, engine: position.position.engine}}
	}
	parentSeg := position.position.Path[len(position.position.Path)-2]
	return position.position.engine.ZoneItem(ZoneItemID(parentSeg.ID))
}

func (_position Position) Exists() (Position, bool) {
	position := _position.position.engine.Position(_position.position.ID)
	return position, position.position.OperationKind != OperationKindDelete
}

func (_position Position) ParentKind() (ElementKind, bool) {
	if !_position.position.HasParent {
		return "", false
	}
	return _position.position.Path[len(_position.position.Path)-2].Kind, true
}

func (_position Position) ID() PositionID {
	return _position.position.ID
}

func (_position Position) Path() string {
	return _position.position.JSONPath
}

func (_position Position) X() float64 {
	position := _position.position.engine.Position(_position.position.ID)
	return position.position.engine.floatValue(position.position.X).Value
}

func (_position Position) Y() float64 {
	position := _position.position.engine.Position(_position.position.ID)
	return position.position.engine.floatValue(position.position.Y).Value
}

func (engine *Engine) QueryZoneItems(matcher func(ZoneItem) bool) []ZoneItem {
	zoneItemIDs := engine.allZoneItemIDs()
	sort.Slice(zoneItemIDs, func(i, j int) bool {
		return zoneItemIDs[i] < zoneItemIDs[j]
	})
	var zoneItems []ZoneItem
	for _, zoneItemID := range zoneItemIDs {
		zoneItem := engine.ZoneItem(zoneItemID)
		if matcher(zoneItem) {
			zoneItems = append(zoneItems, zoneItem)
		}
	}
	zoneItemIDSlicePool.Put(zoneItemIDs)
	return zoneItems
}

func (engine *Engine) EveryZoneItem() []ZoneItem {
	zoneItemIDs := engine.allZoneItemIDs()
	sort.Slice(zoneItemIDs, func(i, j int) bool {
		return zoneItemIDs[i] < zoneItemIDs[j]
	})
	var zoneItems []ZoneItem
	for _, zoneItemID := range zoneItemIDs {
		zoneItem := engine.ZoneItem(zoneItemID)
		if zoneItem.zoneItem.HasParent {
			continue
		}
		zoneItems = append(zoneItems, zoneItem)
	}
	zoneItemIDSlicePool.Put(zoneItemIDs)
	return zoneItems
}

func (engine *Engine) ZoneItem(zoneItemID ZoneItemID) ZoneItem {
	patchingZoneItem, ok := engine.Patch.ZoneItem[zoneItemID]
	if ok {
		return ZoneItem{zoneItem: patchingZoneItem}
	}
	currentZoneItem, ok := engine.State.ZoneItem[zoneItemID]
	if ok {
		return ZoneItem{zoneItem: currentZoneItem}
	}
	return ZoneItem{zoneItem: zoneItemCore{OperationKind: OperationKindDelete, engine: engine}}
}

func (_zoneItem ZoneItem) ParentZone() Zone {
	zoneItem := _zoneItem.zoneItem.engine.ZoneItem(_zoneItem.zoneItem.ID)
	if !zoneItem.zoneItem.HasParent {
		return Zone{zone: zoneCore{OperationKind: OperationKindDelete, engine: zoneItem.zoneItem.engine}}
	}
	parentSeg := zoneItem.zoneItem.Path[len(zoneItem.zoneItem.Path)-2]
	return zoneItem.zoneItem.engine.Zone(ZoneID(parentSeg.ID))
}

func (_zoneItem ZoneItem) ParentKind() (ElementKind, bool) {
	if !_zoneItem.zoneItem.HasParent {
		return "", false
	}
	return _zoneItem.zoneItem.Path[len(_zoneItem.zoneItem.Path)-2].Kind, true
}

func (_zoneItem ZoneItem) ID() ZoneItemID {
	return _zoneItem.zoneItem.ID
}

func (_zoneItem ZoneItem) Exists() (ZoneItem, bool) {
	zoneItem := _zoneItem.zoneItem.engine.ZoneItem(_zoneItem.zoneItem.ID)
	return zoneItem, zoneItem.zoneItem.OperationKind != OperationKindDelete
}

func (_zoneItem ZoneItem) Path() string {
	return _zoneItem.zoneItem.JSONPath
}

func (_zoneItem ZoneItem) Position() Position {
	zoneItem := _zoneItem.zoneItem.engine.ZoneItem(_zoneItem.zoneItem.ID)
	return zoneItem.zoneItem.engine.Position(zoneItem.zoneItem.Position)
}

func (_zoneItem ZoneItem) Item() Item {
	zoneItem := _zoneItem.zoneItem.engine.ZoneItem(_zoneItem.zoneItem.ID)
	return zoneItem.zoneItem.engine.Item(zoneItem.zoneItem.Item)
}

func (engine *Engine) QueryZones(matcher func(Zone) bool) []Zone {
	zoneIDs := engine.allZoneIDs()
	sort.Slice(zoneIDs, func(i, j int) bool {
		return zoneIDs[i] < zoneIDs[j]
	})
	var zones []Zone
	for _, zoneID := range zoneIDs {
		zone := engine.Zone(zoneID)
		if matcher(zone) {
			zones = append(zones, zone)
		}
	}
	zoneIDSlicePool.Put(zoneIDs)
	return zones
}

func (engine *Engine) EveryZone() []Zone {
	zoneIDs := engine.allZoneIDs()
	sort.Slice(zoneIDs, func(i, j int) bool {
		return zoneIDs[i] < zoneIDs[j]
	})
	var zones []Zone
	for _, zoneID := range zoneIDs {
		zone := engine.Zone(zoneID)
		zones = append(zones, zone)
	}
	zoneIDSlicePool.Put(zoneIDs)
	return zones
}

func (engine *Engine) Zone(zoneID ZoneID) Zone {
	patchingZone, ok := engine.Patch.Zone[zoneID]
	if ok {
		return Zone{zone: patchingZone}
	}
	currentZone, ok := engine.State.Zone[zoneID]
	if ok {
		return Zone{zone: currentZone}
	}
	return Zone{zone: zoneCore{OperationKind: OperationKindDelete, engine: engine}}
}

func (_zone Zone) ID() ZoneID {
	return _zone.zone.ID
}

func (_zone Zone) Exists() (Zone, bool) {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	return zone, zone.zone.OperationKind != OperationKindDelete
}

func (_zone Zone) Path() string {
	return _zone.zone.JSONPath
}

func (_zone Zone) Players() []Player {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	var players []Player
	for _, playerID := range zone.zone.Players {
		players = append(players, zone.zone.engine.Player(playerID))
	}
	return players
}

func (_zone Zone) Interactables() []AnyOfItem_Player_ZoneItemSliceElement {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	var interactables []AnyOfItem_Player_ZoneItemSliceElement
	for _, anyOfItem_Player_ZoneItemID := range zone.zone.Interactables {
		interactables = append(interactables, zone.zone.engine.anyOfItem_Player_ZoneItem(anyOfItem_Player_ZoneItemID))
	}
	return interactables
}

func (_zone Zone) Items() []ZoneItem {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	var items []ZoneItem
	for _, zoneItemID := range zone.zone.Items {
		items = append(items, zone.zone.engine.ZoneItem(zoneItemID))
	}
	return items
}

func (_zone Zone) Tags() []string {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	var tags []string
	for _, stringValueID := range zone.zone.Tags {
		tags = append(tags, zone.zone.engine.stringValue(stringValueID).Value)
	}
	return tags
}

func (_itemBoundToRef ItemBoundToRef) ID() PlayerID {
	return _itemBoundToRef.itemBoundToRef.ReferencedElementID
}

func (engine *Engine) itemBoundToRef(itemBoundToRefID ItemBoundToRefID) ItemBoundToRef {
	patchingItemBoundToRef, ok := engine.Patch.ItemBoundToRef[itemBoundToRefID]
	if ok {
		return ItemBoundToRef{itemBoundToRef: patchingItemBoundToRef}
	}
	currentItemBoundToRef, ok := engine.State.ItemBoundToRef[itemBoundToRefID]
	if ok {
		return ItemBoundToRef{itemBoundToRef: currentItemBoundToRef}
	}
	return ItemBoundToRef{itemBoundToRef: itemBoundToRefCore{OperationKind: OperationKindDelete, engine: engine}}
}

func (_attackEventTargetRef AttackEventTargetRef) ID() PlayerID {
	return _attackEventTargetRef.attackEventTargetRef.ReferencedElementID
}

func (engine *Engine) attackEventTargetRef(attackEventTargetRefID AttackEventTargetRefID) AttackEventTargetRef {
	patchingAttackEventTargetRef, ok := engine.Patch.AttackEventTargetRef[attackEventTargetRefID]
	if ok {
		return AttackEventTargetRef{attackEventTargetRef: patchingAttackEventTargetRef}
	}
	currentAttackEventTargetRef, ok := engine.State.AttackEventTargetRef[attackEventTargetRefID]
	if ok {
		return AttackEventTargetRef{attackEventTargetRef: currentAttackEventTargetRef}
	}
	return AttackEventTargetRef{attackEventTargetRef: attackEventTargetRefCore{OperationKind: OperationKindDelete, engine: engine}}
}

func (_playerGuildMemberRef PlayerGuildMemberRef) ID() PlayerID {
	return _playerGuildMemberRef.playerGuildMemberRef.ReferencedElementID
}

func (engine *Engine) playerGuildMemberRef(playerGuildMemberRefID PlayerGuildMemberRefID) PlayerGuildMemberRef {
	patchingPlayerGuildMemberRef, ok := engine.Patch.PlayerGuildMemberRef[playerGuildMemberRefID]
	if ok {
		return PlayerGuildMemberRef{playerGuildMemberRef: patchingPlayerGuildMemberRef}
	}
	currentPlayerGuildMemberRef, ok := engine.State.PlayerGuildMemberRef[playerGuildMemberRefID]
	if ok {
		return PlayerGuildMemberRef{playerGuildMemberRef: currentPlayerGuildMemberRef}
	}
	return PlayerGuildMemberRef{playerGuildMemberRef: playerGuildMemberRefCore{OperationKind: OperationKindDelete, engine: engine}}
}

func (_playerEquipmentSetRef PlayerEquipmentSetRef) ID() EquipmentSetID {
	return _playerEquipmentSetRef.playerEquipmentSetRef.ReferencedElementID
}

func (engine *Engine) QueryEquipmentSets(matcher func(EquipmentSet) bool) []EquipmentSet {
	equipmentSetIDs := engine.allEquipmentSetIDs()
	sort.Slice(equipmentSetIDs, func(i, j int) bool {
		return equipmentSetIDs[i] < equipmentSetIDs[j]
	})
	var equipmentSets []EquipmentSet
	for _, equipmentSetID := range equipmentSetIDs {
		equipmentSet := engine.EquipmentSet(equipmentSetID)
		if matcher(equipmentSet) {
			equipmentSets = append(equipmentSets, equipmentSet)
		}
	}
	equipmentSetIDSlicePool.Put(equipmentSetIDs)
	return equipmentSets
}

func (engine *Engine) EveryEquipmentSet() []EquipmentSet {
	equipmentSetIDs := engine.allEquipmentSetIDs()
	sort.Slice(equipmentSetIDs, func(i, j int) bool {
		return equipmentSetIDs[i] < equipmentSetIDs[j]
	})
	var equipmentSets []EquipmentSet
	for _, equipmentSetID := range equipmentSetIDs {
		equipmentSet := engine.EquipmentSet(equipmentSetID)
		equipmentSets = append(equipmentSets, equipmentSet)
	}
	equipmentSetIDSlicePool.Put(equipmentSetIDs)
	return equipmentSets
}

func (engine *Engine) EquipmentSet(equipmentSetID EquipmentSetID) EquipmentSet {
	patchingEquipmentSet, ok := engine.Patch.EquipmentSet[equipmentSetID]
	if ok {
		return EquipmentSet{equipmentSet: patchingEquipmentSet}
	}
	currentEquipmentSet, ok := engine.State.EquipmentSet[equipmentSetID]
	if ok {
		return EquipmentSet{equipmentSet: currentEquipmentSet}
	}
	return EquipmentSet{equipmentSet: equipmentSetCore{OperationKind: OperationKindDelete, engine: engine}}
}

func (_equipmentSet EquipmentSet) ID() EquipmentSetID {
	return _equipmentSet.equipmentSet.ID
}

func (_equipmentSet EquipmentSet) Exists() (EquipmentSet, bool) {
	equipmentSet := _equipmentSet.equipmentSet.engine.EquipmentSet(_equipmentSet.equipmentSet.ID)
	return equipmentSet, equipmentSet.equipmentSet.OperationKind != OperationKindDelete
}

func (_equipmentSet EquipmentSet) Path() string {
	return _equipmentSet.equipmentSet.JSONPath
}

func (_equipmentSet EquipmentSet) Name() string {
	equipmentSet := _equipmentSet.equipmentSet.engine.EquipmentSet(_equipmentSet.equipmentSet.ID)
	return equipmentSet.equipmentSet.engine.stringValue(equipmentSet.equipmentSet.Name).Value
}

func (_equipmentSet EquipmentSet) Equipment() []EquipmentSetEquipmentRef {
	equipmentSet := _equipmentSet.equipmentSet.engine.EquipmentSet(_equipmentSet.equipmentSet.ID)
	var equipment []EquipmentSetEquipmentRef
	for _, refID := range equipmentSet.equipmentSet.Equipment {
		equipment = append(equipment, equipmentSet.equipmentSet.engine.equipmentSetEquipmentRef(refID))
	}
	return equipment
}

func (engine *Engine) playerEquipmentSetRef(playerEquipmentSetRefID PlayerEquipmentSetRefID) PlayerEquipmentSetRef {
	patchingPlayerEquipmentSetRef, ok := engine.Patch.PlayerEquipmentSetRef[playerEquipmentSetRefID]
	if ok {
		return PlayerEquipmentSetRef{playerEquipmentSetRef: patchingPlayerEquipmentSetRef}
	}
	currentPlayerEquipmentSetRef, ok := engine.State.PlayerEquipmentSetRef[playerEquipmentSetRefID]
	if ok {
		return PlayerEquipmentSetRef{playerEquipmentSetRef: currentPlayerEquipmentSetRef}
	}
	return PlayerEquipmentSetRef{playerEquipmentSetRef: playerEquipmentSetRefCore{OperationKind: OperationKindDelete, engine: engine}}
}

func (_equipmentSetEquipmentRef EquipmentSetEquipmentRef) ID() ItemID {
	return _equipmentSetEquipmentRef.equipmentSetEquipmentRef.ReferencedElementID
}

func (engine *Engine) equipmentSetEquipmentRef(equipmentSetEquipmentRefID EquipmentSetEquipmentRefID) EquipmentSetEquipmentRef {
	patchingEquipmentSetEquipmentRef, ok := engine.Patch.EquipmentSetEquipmentRef[equipmentSetEquipmentRefID]
	if ok {
		return EquipmentSetEquipmentRef{equipmentSetEquipmentRef: patchingEquipmentSetEquipmentRef}
	}
	currentEquipmentSetEquipmentRef, ok := engine.State.EquipmentSetEquipmentRef[equipmentSetEquipmentRefID]
	if ok {
		return EquipmentSetEquipmentRef{equipmentSetEquipmentRef: currentEquipmentSetEquipmentRef}
	}
	return EquipmentSetEquipmentRef{equipmentSetEquipmentRef: equipmentSetEquipmentRefCore{OperationKind: OperationKindDelete, engine: engine}}
}

func (_playerTargetRef PlayerTargetRef) ID() AnyOfPlayer_ZoneItemID {
	return _playerTargetRef.playerTargetRef.ReferencedElementID
}

func (engine *Engine) playerTargetRef(playerTargetRefID PlayerTargetRefID) PlayerTargetRef {
	patchingPlayerTargetRef, ok := engine.Patch.PlayerTargetRef[playerTargetRefID]
	if ok {
		return PlayerTargetRef{playerTargetRef: patchingPlayerTargetRef}
	}
	currentPlayerTargetRef, ok := engine.State.PlayerTargetRef[playerTargetRefID]
	if ok {
		return PlayerTargetRef{playerTargetRef: currentPlayerTargetRef}
	}
	return PlayerTargetRef{playerTargetRef: playerTargetRefCore{OperationKind: OperationKindDelete, engine: engine}}
}

func (_playerTargetedByRef PlayerTargetedByRef) ID() AnyOfPlayer_ZoneItemID {
	return _playerTargetedByRef.playerTargetedByRef.ReferencedElementID
}

func (engine *Engine) playerTargetedByRef(playerTargetedByRefID PlayerTargetedByRefID) PlayerTargetedByRef {
	patchingPlayerTargetedByRef, ok := engine.Patch.PlayerTargetedByRef[playerTargetedByRefID]
	if ok {
		return PlayerTargetedByRef{playerTargetedByRef: patchingPlayerTargetedByRef}
	}
	currentPlayerTargetedByRef, ok := engine.State.PlayerTargetedByRef[playerTargetedByRefID]
	if ok {
		return PlayerTargetedByRef{playerTargetedByRef: currentPlayerTargetedByRef}
	}
	return PlayerTargetedByRef{playerTargetedByRef: playerTargetedByRefCore{OperationKind: OperationKindDelete, engine: engine}}
}

func (_anyOfPlayer_Position AnyOfPlayer_Position) ID() AnyOfPlayer_PositionID {
	return _anyOfPlayer_Position.anyOfPlayer_Position.ID
}

func (_anyOfPlayer_Position AnyOfPlayer_Position) Player() Player {
	anyOfPlayer_Position := _anyOfPlayer_Position.anyOfPlayer_Position.engine.anyOfPlayer_Position(_anyOfPlayer_Position.anyOfPlayer_Position.ID)
	return anyOfPlayer_Position.anyOfPlayer_Position.engine.Player(PlayerID(anyOfPlayer_Position.anyOfPlayer_Position.ChildID))
}

func (_anyOfPlayer_Position AnyOfPlayer_Position) Position() Position {
	anyOfPlayer_Position := _anyOfPlayer_Position.anyOfPlayer_Position.engine.anyOfPlayer_Position(_anyOfPlayer_Position.anyOfPlayer_Position.ID)
	return anyOfPlayer_Position.anyOfPlayer_Position.engine.Position(PositionID(anyOfPlayer_Position.anyOfPlayer_Position.ChildID))
}

func (engine *Engine) anyOfPlayer_Position(anyOfPlayer_PositionID AnyOfPlayer_PositionID) AnyOfPlayer_Position {
	patchingAnyOfPlayer_Position, ok := engine.Patch.AnyOfPlayer_Position[anyOfPlayer_PositionID]
	if ok {
		return AnyOfPlayer_Position{anyOfPlayer_Position: patchingAnyOfPlayer_Position}
	}
	currentAnyOfPlayer_Position, ok := engine.State.AnyOfPlayer_Position[anyOfPlayer_PositionID]
	if ok {
		return AnyOfPlayer_Position{anyOfPlayer_Position: currentAnyOfPlayer_Position}
	}
	return AnyOfPlayer_Position{anyOfPlayer_Position: anyOfPlayer_PositionCore{OperationKind: OperationKindDelete, engine: engine}}
}

func (_anyOfPlayer_ZoneItem AnyOfPlayer_ZoneItem) ID() AnyOfPlayer_ZoneItemID {
	return _anyOfPlayer_ZoneItem.anyOfPlayer_ZoneItem.ID
}

func (_anyOfPlayer_ZoneItem AnyOfPlayer_ZoneItem) Player() Player {
	anyOfPlayer_ZoneItem := _anyOfPlayer_ZoneItem.anyOfPlayer_ZoneItem.engine.anyOfPlayer_ZoneItem(_anyOfPlayer_ZoneItem.anyOfPlayer_ZoneItem.ID)
	return anyOfPlayer_ZoneItem.anyOfPlayer_ZoneItem.engine.Player(PlayerID(anyOfPlayer_ZoneItem.anyOfPlayer_ZoneItem.ChildID))
}

func (_anyOfPlayer_ZoneItem AnyOfPlayer_ZoneItem) ZoneItem() ZoneItem {
	anyOfPlayer_ZoneItem := _anyOfPlayer_ZoneItem.anyOfPlayer_ZoneItem.engine.anyOfPlayer_ZoneItem(_anyOfPlayer_ZoneItem.anyOfPlayer_ZoneItem.ID)
	return anyOfPlayer_ZoneItem.anyOfPlayer_ZoneItem.engine.ZoneItem(ZoneItemID(anyOfPlayer_ZoneItem.anyOfPlayer_ZoneItem.ChildID))
}

func (engine *Engine) anyOfPlayer_ZoneItem(anyOfPlayer_ZoneItemID AnyOfPlayer_ZoneItemID) AnyOfPlayer_ZoneItem {
	patchingAnyOfPlayer_ZoneItem, ok := engine.Patch.AnyOfPlayer_ZoneItem[anyOfPlayer_ZoneItemID]
	if ok {
		return AnyOfPlayer_ZoneItem{anyOfPlayer_ZoneItem: patchingAnyOfPlayer_ZoneItem}
	}
	currentAnyOfPlayer_ZoneItem, ok := engine.State.AnyOfPlayer_ZoneItem[anyOfPlayer_ZoneItemID]
	if ok {
		return AnyOfPlayer_ZoneItem{anyOfPlayer_ZoneItem: currentAnyOfPlayer_ZoneItem}
	}
	return AnyOfPlayer_ZoneItem{anyOfPlayer_ZoneItem: anyOfPlayer_ZoneItemCore{OperationKind: OperationKindDelete, engine: engine}}
}

func (engine *Engine) anyOfItem_Player_ZoneItem(anyOfItem_Player_ZoneItemID AnyOfItem_Player_ZoneItemID) AnyOfItem_Player_ZoneItem {
	patchingAnyOfItem_Player_ZoneItem, ok := engine.Patch.AnyOfItem_Player_ZoneItem[anyOfItem_Player_ZoneItemID]
	if ok {
		return AnyOfItem_Player_ZoneItem{anyOfItem_Player_ZoneItem: patchingAnyOfItem_Player_ZoneItem}
	}
	currentAnyOfItem_Player_ZoneItem, ok := engine.State.AnyOfItem_Player_ZoneItem[anyOfItem_Player_ZoneItemID]
	if ok {
		return AnyOfItem_Player_ZoneItem{anyOfItem_Player_ZoneItem: currentAnyOfItem_Player_ZoneItem}
	}
	return AnyOfItem_Player_ZoneItem{anyOfItem_Player_ZoneItem: anyOfItem_Player_ZoneItemCore{OperationKind: OperationKindDelete, engine: engine}}
}

func (_anyOfItem_Player_ZoneItem AnyOfItem_Player_ZoneItem) ID() AnyOfItem_Player_ZoneItemID {
	return _anyOfItem_Player_ZoneItem.anyOfItem_Player_ZoneItem.ID
}

func (_anyOfItem_Player_ZoneItem AnyOfItem_Player_ZoneItem) Player() Player {
	anyOfItem_Player_ZoneItem := _anyOfItem_Player_ZoneItem.anyOfItem_Player_ZoneItem.engine.anyOfItem_Player_ZoneItem(_anyOfItem_Player_ZoneItem.anyOfItem_Player_ZoneItem.ID)
	return anyOfItem_Player_ZoneItem.anyOfItem_Player_ZoneItem.engine.Player(PlayerID(anyOfItem_Player_ZoneItem.anyOfItem_Player_ZoneItem.ChildID))
}

func (_anyOfItem_Player_ZoneItem AnyOfItem_Player_ZoneItem) ZoneItem() ZoneItem {
	anyOfItem_Player_ZoneItem := _anyOfItem_Player_ZoneItem.anyOfItem_Player_ZoneItem.engine.anyOfItem_Player_ZoneItem(_anyOfItem_Player_ZoneItem.anyOfItem_Player_ZoneItem.ID)
	return anyOfItem_Player_ZoneItem.anyOfItem_Player_ZoneItem.engine.ZoneItem(ZoneItemID(anyOfItem_Player_ZoneItem.anyOfItem_Player_ZoneItem.ChildID))
}

func (_anyOfItem_Player_ZoneItem AnyOfItem_Player_ZoneItem) Item() Item {
	anyOfItem_Player_ZoneItem := _anyOfItem_Player_ZoneItem.anyOfItem_Player_ZoneItem.engine.anyOfItem_Player_ZoneItem(_anyOfItem_Player_ZoneItem.anyOfItem_Player_ZoneItem.ID)
	return anyOfItem_Player_ZoneItem.anyOfItem_Player_ZoneItem.engine.Item(ItemID(anyOfItem_Player_ZoneItem.anyOfItem_Player_ZoneItem.ChildID))
}
