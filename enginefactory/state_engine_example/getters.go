package state

func (engine *Engine) Player(playerID PlayerID) player {
	patchingPlayer, ok := engine.Patch.Player[playerID]
	if ok {
		return player{player: patchingPlayer}
	}
	currentPlayer, ok := engine.State.Player[playerID]
	if ok {
		return player{player: currentPlayer}
	}
	return player{player: playerCore{OperationKind: OperationKindDelete, engine: engine}}
}

func (_player player) ID() PlayerID {
	return _player.player.ID
}

func (_player player) Target() (playerTargetRef, bool) {
	player := _player.player.engine.Player(_player.player.ID)
	return player.player.engine.playerTargetRef(player.player.Target), player.player.Target != 0
}

func (_player player) TargetedBy() []playerTargetedByRef {
	player := _player.player.engine.Player(_player.player.ID)
	var targetedBy []playerTargetedByRef
	for _, refID := range player.player.TargetedBy {
		targetedBy = append(targetedBy, player.player.engine.playerTargetedByRef(refID))
	}
	return targetedBy
}

func (_player player) Items() []item {
	player := _player.player.engine.Player(_player.player.ID)
	var items []item
	for _, itemID := range player.player.Items {
		items = append(items, player.player.engine.Item(itemID))
	}
	return items
}

func (_player player) GearScore() gearScore {
	player := _player.player.engine.Player(_player.player.ID)
	return player.player.engine.GearScore(player.player.GearScore)
}

func (_player player) GuildMembers() []playerGuildMemberRef {
	player := _player.player.engine.Player(_player.player.ID)
	var guildMembers []playerGuildMemberRef
	for _, refID := range player.player.GuildMembers {
		guildMembers = append(guildMembers, player.player.engine.playerGuildMemberRef(refID))
	}
	return guildMembers
}

func (_player player) EquipmentSets() []playerEquipmentSetRef {
	player := _player.player.engine.Player(_player.player.ID)
	var equipmentSets []playerEquipmentSetRef
	for _, refID := range player.player.EquipmentSets {
		equipmentSets = append(equipmentSets, player.player.engine.playerEquipmentSetRef(refID))
	}
	return equipmentSets
}

func (_player player) Position() position {
	player := _player.player.engine.Player(_player.player.ID)
	return player.player.engine.Position(player.player.Position)
}

func (engine *Engine) GearScore(gearScoreID GearScoreID) gearScore {
	patchingGearScore, ok := engine.Patch.GearScore[gearScoreID]
	if ok {
		return gearScore{gearScore: patchingGearScore}
	}
	currentGearScore, ok := engine.State.GearScore[gearScoreID]
	if ok {
		return gearScore{gearScore: currentGearScore}
	}
	return gearScore{gearScore: gearScoreCore{OperationKind: OperationKindDelete, engine: engine}}
}

func (_gearScore gearScore) ID() GearScoreID {
	return _gearScore.gearScore.ID
}

func (_gearScore gearScore) Level() int {
	gearScore := _gearScore.gearScore.engine.GearScore(_gearScore.gearScore.ID)
	return gearScore.gearScore.Level
}

func (_gearScore gearScore) Score() int {
	gearScore := _gearScore.gearScore.engine.GearScore(_gearScore.gearScore.ID)
	return gearScore.gearScore.Score
}

func (engine *Engine) Item(itemID ItemID) item {
	patchingItem, ok := engine.Patch.Item[itemID]
	if ok {
		return item{item: patchingItem}
	}
	currentItem, ok := engine.State.Item[itemID]
	if ok {
		return item{item: currentItem}
	}
	return item{item: itemCore{OperationKind: OperationKindDelete, engine: engine}}
}

func (_item item) ID() ItemID {
	return _item.item.ID
}
func (_item item) Name() string {
	item := _item.item.engine.Item(_item.item.ID)
	return item.item.Name
}

func (_item item) GearScore() gearScore {
	item := _item.item.engine.Item(_item.item.ID)
	return item.item.engine.GearScore(item.item.GearScore)
}

func (_item item) BoundTo() (itemBoundToRef, bool) {
	item := _item.item.engine.Item(_item.item.ID)
	return item.item.engine.itemBoundToRef(item.item.BoundTo), item.item.BoundTo != 0
}

func (_item item) Origin() anyOfPlayerPosition {
	item := _item.item.engine.Item(_item.item.ID)
	return item.item.engine.anyOfPlayerPosition(item.item.Origin)
}

func (engine *Engine) Position(positionID PositionID) position {
	patchingPosition, ok := engine.Patch.Position[positionID]
	if ok {
		return position{position: patchingPosition}
	}
	currentPosition, ok := engine.State.Position[positionID]
	if ok {
		return position{position: currentPosition}
	}
	return position{position: positionCore{OperationKind: OperationKindDelete, engine: engine}}
}

func (_position position) ID() PositionID {
	return _position.position.ID
}

func (_position position) X() float64 {
	position := _position.position.engine.Position(_position.position.ID)
	return position.position.X
}

func (_position position) Y() float64 {
	position := _position.position.engine.Position(_position.position.ID)
	return position.position.Y
}

func (engine *Engine) ZoneItem(zoneItemID ZoneItemID) zoneItem {
	patchingZoneItem, ok := engine.Patch.ZoneItem[zoneItemID]
	if ok {
		return zoneItem{zoneItem: patchingZoneItem}
	}
	currentZoneItem, ok := engine.State.ZoneItem[zoneItemID]
	if ok {
		return zoneItem{zoneItem: currentZoneItem}
	}
	return zoneItem{zoneItem: zoneItemCore{OperationKind: OperationKindDelete, engine: engine}}
}

func (_zoneItem zoneItem) ID() ZoneItemID {
	return _zoneItem.zoneItem.ID
}

func (_zoneItem zoneItem) Position() position {
	zoneItem := _zoneItem.zoneItem.engine.ZoneItem(_zoneItem.zoneItem.ID)
	return zoneItem.zoneItem.engine.Position(zoneItem.zoneItem.Position)
}

func (_zoneItem zoneItem) Item() item {
	zoneItem := _zoneItem.zoneItem.engine.ZoneItem(_zoneItem.zoneItem.ID)
	return zoneItem.zoneItem.engine.Item(zoneItem.zoneItem.Item)
}

func (engine *Engine) Zone(zoneID ZoneID) zone {
	patchingZone, ok := engine.Patch.Zone[zoneID]
	if ok {
		return zone{zone: patchingZone}
	}
	currentZone, ok := engine.State.Zone[zoneID]
	if ok {
		return zone{zone: currentZone}
	}
	return zone{zone: zoneCore{OperationKind: OperationKindDelete, engine: engine}}
}

func (_zone zone) ID() ZoneID {
	return _zone.zone.ID
}

func (_zone zone) Players() []player {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	var players []player
	for _, playerID := range zone.zone.Players {
		players = append(players, zone.zone.engine.Player(playerID))
	}
	return players
}

func (_zone zone) Interactables() []anyOfItemPlayerZoneItem {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	var interactables []anyOfItemPlayerZoneItem
	for _, anyOfItemPlayerZoneItemID := range zone.zone.Interactables {
		interactables = append(interactables, zone.zone.engine.anyOfItemPlayerZoneItem(anyOfItemPlayerZoneItemID))
	}
	return interactables
}

func (_zone zone) Items() []zoneItem {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	var items []zoneItem
	for _, zoneItemID := range zone.zone.Items {
		items = append(items, zone.zone.engine.ZoneItem(zoneItemID))
	}
	return items
}

func (_zone zone) Tags() []string {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	var tags []string
	for _, element := range zone.zone.Tags {
		tags = append(tags, element)
	}
	return tags
}

func (_itemBoundToRef itemBoundToRef) ID() PlayerID {
	return _itemBoundToRef.itemBoundToRef.ReferencedElementID
}

func (engine *Engine) itemBoundToRef(itemBoundToRefID ItemBoundToRefID) itemBoundToRef {
	patchingItemBoundToRef, ok := engine.Patch.ItemBoundToRef[itemBoundToRefID]
	if ok {
		return itemBoundToRef{itemBoundToRef: patchingItemBoundToRef}
	}
	currentItemBoundToRef, ok := engine.State.ItemBoundToRef[itemBoundToRefID]
	if ok {
		return itemBoundToRef{itemBoundToRef: currentItemBoundToRef}
	}
	return itemBoundToRef{itemBoundToRef: itemBoundToRefCore{OperationKind: OperationKindDelete, engine: engine}}
}

func (_playerGuildMemberRef playerGuildMemberRef) ID() PlayerID {
	return _playerGuildMemberRef.playerGuildMemberRef.ReferencedElementID
}

func (engine *Engine) playerGuildMemberRef(playerGuildMemberRefID PlayerGuildMemberRefID) playerGuildMemberRef {
	patchingPlayerGuildMemberRef, ok := engine.Patch.PlayerGuildMemberRef[playerGuildMemberRefID]
	if ok {
		return playerGuildMemberRef{playerGuildMemberRef: patchingPlayerGuildMemberRef}
	}
	currentPlayerGuildMemberRef, ok := engine.State.PlayerGuildMemberRef[playerGuildMemberRefID]
	if ok {
		return playerGuildMemberRef{playerGuildMemberRef: currentPlayerGuildMemberRef}
	}
	return playerGuildMemberRef{playerGuildMemberRef: playerGuildMemberRefCore{OperationKind: OperationKindDelete, engine: engine}}
}

func (engine *Engine) playerEquipmentSetRef(playerEquipmentSetRefID PlayerEquipmentSetRefID) playerEquipmentSetRef {
	patchingPlayerEquipmentSetRef, ok := engine.Patch.PlayerEquipmentSetRef[playerEquipmentSetRefID]
	if ok {
		return playerEquipmentSetRef{playerEquipmentSetRef: patchingPlayerEquipmentSetRef}
	}
	currentPlayerEquipmentSetRef, ok := engine.State.PlayerEquipmentSetRef[playerEquipmentSetRefID]
	if ok {
		return playerEquipmentSetRef{playerEquipmentSetRef: currentPlayerEquipmentSetRef}
	}
	return playerEquipmentSetRef{playerEquipmentSetRef: playerEquipmentSetRefCore{OperationKind: OperationKindDelete, engine: engine}}
}

func (_playerEquipmentSetRef playerEquipmentSetRef) ID() EquipmentSetID {
	return _playerEquipmentSetRef.playerEquipmentSetRef.ReferencedElementID
}

func (engine *Engine) EquipmentSet(equipmentSetID EquipmentSetID) equipmentSet {
	patchingEquipmentSet, ok := engine.Patch.EquipmentSet[equipmentSetID]
	if ok {
		return equipmentSet{equipmentSet: patchingEquipmentSet}
	}
	currentEquipmentSet, ok := engine.State.EquipmentSet[equipmentSetID]
	if ok {
		return equipmentSet{equipmentSet: currentEquipmentSet}
	}
	return equipmentSet{equipmentSet: equipmentSetCore{OperationKind: OperationKindDelete, engine: engine}}
}

func (_equipmentSet equipmentSet) ID() EquipmentSetID {
	return _equipmentSet.equipmentSet.ID
}

func (_equipmentSet equipmentSet) Name() string {
	equipmentSet := _equipmentSet.equipmentSet.engine.EquipmentSet(_equipmentSet.equipmentSet.ID)
	return equipmentSet.equipmentSet.Name
}

func (_equipmentSet equipmentSet) Equipment() []equipmentSetEquipmentRef {
	equipmentSet := _equipmentSet.equipmentSet.engine.EquipmentSet(_equipmentSet.equipmentSet.ID)
	var equipment []equipmentSetEquipmentRef
	for _, refID := range equipmentSet.equipmentSet.Equipment {
		equipment = append(equipment, equipmentSet.equipmentSet.engine.equipmentSetEquipmentRef(refID))
	}
	return equipment
}

func (_equipmentSetEquipmentRef equipmentSetEquipmentRef) ID() ItemID {
	return _equipmentSetEquipmentRef.equipmentSetEquipmentRef.ReferencedElementID
}

func (engine *Engine) equipmentSetEquipmentRef(equipmentSetEquipmentRefID EquipmentSetEquipmentRefID) equipmentSetEquipmentRef {
	patchingEquipmentSetEquipmentRef, ok := engine.Patch.EquipmentSetEquipmentRef[equipmentSetEquipmentRefID]
	if ok {
		return equipmentSetEquipmentRef{equipmentSetEquipmentRef: patchingEquipmentSetEquipmentRef}
	}
	currentEquipmentSetEquipmentRef, ok := engine.State.EquipmentSetEquipmentRef[equipmentSetEquipmentRefID]
	if ok {
		return equipmentSetEquipmentRef{equipmentSetEquipmentRef: currentEquipmentSetEquipmentRef}
	}
	return equipmentSetEquipmentRef{equipmentSetEquipmentRef: equipmentSetEquipmentRefCore{OperationKind: OperationKindDelete, engine: engine}}
}

func (_playerTargetRef playerTargetRef) ID() AnyOfPlayerZoneItemID {
	return _playerTargetRef.playerTargetRef.ReferencedElementID
}

func (engine *Engine) playerTargetRef(playerTargetRefID PlayerTargetRefID) playerTargetRef {
	patchingPlayerTargetRef, ok := engine.Patch.PlayerTargetRef[playerTargetRefID]
	if ok {
		return playerTargetRef{playerTargetRef: patchingPlayerTargetRef}
	}
	currentPlayerTargetRef, ok := engine.State.PlayerTargetRef[playerTargetRefID]
	if ok {
		return playerTargetRef{playerTargetRef: currentPlayerTargetRef}
	}
	return playerTargetRef{playerTargetRef: playerTargetRefCore{OperationKind: OperationKindDelete, engine: engine}}
}

func (_playerTargetedByRef playerTargetedByRef) ID() AnyOfPlayerZoneItemID {
	return _playerTargetedByRef.playerTargetedByRef.ReferencedElementID
}

func (engine *Engine) playerTargetedByRef(playerTargetedByRefID PlayerTargetedByRefID) playerTargetedByRef {
	patchingPlayerTargetedByRef, ok := engine.Patch.PlayerTargetedByRef[playerTargetedByRefID]
	if ok {
		return playerTargetedByRef{playerTargetedByRef: patchingPlayerTargetedByRef}
	}
	currentPlayerTargetedByRef, ok := engine.State.PlayerTargetedByRef[playerTargetedByRefID]
	if ok {
		return playerTargetedByRef{playerTargetedByRef: currentPlayerTargetedByRef}
	}
	return playerTargetedByRef{playerTargetedByRef: playerTargetedByRefCore{OperationKind: OperationKindDelete, engine: engine}}
}

func (_anyOfPlayerPosition anyOfPlayerPosition) ID() AnyOfPlayerPositionID {
	return _anyOfPlayerPosition.anyOfPlayerPosition.ID
}

func (_anyOfPlayerPosition anyOfPlayerPosition) Player() player {
	anyOfPlayerPosition := _anyOfPlayerPosition.anyOfPlayerPosition.engine.anyOfPlayerPosition(_anyOfPlayerPosition.anyOfPlayerPosition.ID)
	return anyOfPlayerPosition.anyOfPlayerPosition.engine.Player(anyOfPlayerPosition.anyOfPlayerPosition.Player)
}

func (_anyOfPlayerPosition anyOfPlayerPosition) Position() position {
	anyOfPlayerPosition := _anyOfPlayerPosition.anyOfPlayerPosition.engine.anyOfPlayerPosition(_anyOfPlayerPosition.anyOfPlayerPosition.ID)
	return anyOfPlayerPosition.anyOfPlayerPosition.engine.Position(anyOfPlayerPosition.anyOfPlayerPosition.Position)
}

func (engine *Engine) anyOfPlayerPosition(anyOfPlayerPositionID AnyOfPlayerPositionID) anyOfPlayerPosition {
	patchingAnyOfPlayerPosition, ok := engine.Patch.AnyOfPlayerPosition[anyOfPlayerPositionID]
	if ok {
		return anyOfPlayerPosition{anyOfPlayerPosition: patchingAnyOfPlayerPosition}
	}
	currentAnyOfPlayerPosition, ok := engine.State.AnyOfPlayerPosition[anyOfPlayerPositionID]
	if ok {
		return anyOfPlayerPosition{anyOfPlayerPosition: currentAnyOfPlayerPosition}
	}
	return anyOfPlayerPosition{anyOfPlayerPosition: anyOfPlayerPositionCore{OperationKind: OperationKindDelete, engine: engine}}
}

func (_anyOfPlayerZoneItem anyOfPlayerZoneItem) ID() AnyOfPlayerZoneItemID {
	return _anyOfPlayerZoneItem.anyOfPlayerZoneItem.ID
}

func (_anyOfPlayerZoneItem anyOfPlayerZoneItem) Player() player {
	anyOfPlayerZoneItem := _anyOfPlayerZoneItem.anyOfPlayerZoneItem.engine.anyOfPlayerZoneItem(_anyOfPlayerZoneItem.anyOfPlayerZoneItem.ID)
	return anyOfPlayerZoneItem.anyOfPlayerZoneItem.engine.Player(anyOfPlayerZoneItem.anyOfPlayerZoneItem.Player)
}

func (_anyOfPlayerZoneItem anyOfPlayerZoneItem) ZoneItem() zoneItem {
	anyOfPlayerZoneItem := _anyOfPlayerZoneItem.anyOfPlayerZoneItem.engine.anyOfPlayerZoneItem(_anyOfPlayerZoneItem.anyOfPlayerZoneItem.ID)
	return anyOfPlayerZoneItem.anyOfPlayerZoneItem.engine.ZoneItem(anyOfPlayerZoneItem.anyOfPlayerZoneItem.ZoneItem)
}

func (engine *Engine) anyOfPlayerZoneItem(anyOfPlayerZoneItemID AnyOfPlayerZoneItemID) anyOfPlayerZoneItem {
	patchingAnyOfPlayerZoneItem, ok := engine.Patch.AnyOfPlayerZoneItem[anyOfPlayerZoneItemID]
	if ok {
		return anyOfPlayerZoneItem{anyOfPlayerZoneItem: patchingAnyOfPlayerZoneItem}
	}
	currentAnyOfPlayerZoneItem, ok := engine.State.AnyOfPlayerZoneItem[anyOfPlayerZoneItemID]
	if ok {
		return anyOfPlayerZoneItem{anyOfPlayerZoneItem: currentAnyOfPlayerZoneItem}
	}
	return anyOfPlayerZoneItem{anyOfPlayerZoneItem: anyOfPlayerZoneItemCore{OperationKind: OperationKindDelete, engine: engine}}
}

func (engine *Engine) anyOfItemPlayerZoneItem(anyOfItemPlayerZoneItemID AnyOfItemPlayerZoneItemID) anyOfItemPlayerZoneItem {
	patchingAnyOfItemPlayerZoneItem, ok := engine.Patch.AnyOfItemPlayerZoneItem[anyOfItemPlayerZoneItemID]
	if ok {
		return anyOfItemPlayerZoneItem{anyOfItemPlayerZoneItem: patchingAnyOfItemPlayerZoneItem}
	}
	currentAnyOfItemPlayerZoneItem, ok := engine.State.AnyOfItemPlayerZoneItem[anyOfItemPlayerZoneItemID]
	if ok {
		return anyOfItemPlayerZoneItem{anyOfItemPlayerZoneItem: currentAnyOfItemPlayerZoneItem}
	}
	return anyOfItemPlayerZoneItem{anyOfItemPlayerZoneItem: anyOfItemPlayerZoneItemCore{OperationKind: OperationKindDelete, engine: engine}}
}

func (_anyOfItemPlayerZoneItem anyOfItemPlayerZoneItem) ID() AnyOfItemPlayerZoneItemID {
	return _anyOfItemPlayerZoneItem.anyOfItemPlayerZoneItem.ID
}

func (_anyOfItemPlayerZoneItem anyOfItemPlayerZoneItem) Player() player {
	anyOfItemPlayerZoneItem := _anyOfItemPlayerZoneItem.anyOfItemPlayerZoneItem.engine.anyOfItemPlayerZoneItem(_anyOfItemPlayerZoneItem.anyOfItemPlayerZoneItem.ID)
	return anyOfItemPlayerZoneItem.anyOfItemPlayerZoneItem.engine.Player(anyOfItemPlayerZoneItem.anyOfItemPlayerZoneItem.Player)
}

func (_anyOfItemPlayerZoneItem anyOfItemPlayerZoneItem) ZoneItem() zoneItem {
	anyOfItemPlayerZoneItem := _anyOfItemPlayerZoneItem.anyOfItemPlayerZoneItem.engine.anyOfItemPlayerZoneItem(_anyOfItemPlayerZoneItem.anyOfItemPlayerZoneItem.ID)
	return anyOfItemPlayerZoneItem.anyOfItemPlayerZoneItem.engine.ZoneItem(anyOfItemPlayerZoneItem.anyOfItemPlayerZoneItem.ZoneItem)
}

func (_anyOfItemPlayerZoneItem anyOfItemPlayerZoneItem) Item() item {
	anyOfItemPlayerZoneItem := _anyOfItemPlayerZoneItem.anyOfItemPlayerZoneItem.engine.anyOfItemPlayerZoneItem(_anyOfItemPlayerZoneItem.anyOfItemPlayerZoneItem.ID)
	return anyOfItemPlayerZoneItem.anyOfItemPlayerZoneItem.engine.Item(anyOfItemPlayerZoneItem.anyOfItemPlayerZoneItem.Item)
}
