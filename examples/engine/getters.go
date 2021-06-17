package state

func (engine *Engine) EveryPlayer() []player {
	playerIDs := engine.allPlayerIDs()
	var players []player
	for _, playerID := range playerIDs {
		players = append(players, engine.Player(playerID))
	}
	return players
}

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

func (engine *Engine) EveryGearScore() []gearScore {
	gearScoreIDs := engine.allGearScoreIDs()
	var gearScores []gearScore
	for _, gearScoreID := range gearScoreIDs {
		gearScores = append(gearScores, engine.GearScore(gearScoreID))
	}
	return gearScores
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

func (engine *Engine) EveryItem() []item {
	itemIDs := engine.allItemIDs()
	var items []item
	for _, itemID := range itemIDs {
		items = append(items, engine.Item(itemID))
	}
	return items
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

func (_item item) Origin() anyOfPlayer_Position {
	item := _item.item.engine.Item(_item.item.ID)
	return item.item.engine.anyOfPlayer_Position(item.item.Origin)
}

func (engine *Engine) EveryPosition() []position {
	positionIDs := engine.allPositionIDs()
	var positions []position
	for _, positionID := range positionIDs {
		positions = append(positions, engine.Position(positionID))
	}
	return positions
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

func (engine *Engine) EveryZoneItem() []zoneItem {
	zoneItemIDs := engine.allZoneItemIDs()
	var zoneItems []zoneItem
	for _, zoneItemID := range zoneItemIDs {
		zoneItems = append(zoneItems, engine.ZoneItem(zoneItemID))
	}
	return zoneItems
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

func (engine *Engine) EveryZone() []zone {
	zoneIDs := engine.allZoneIDs()
	var zones []zone
	for _, zoneID := range zoneIDs {
		zones = append(zones, engine.Zone(zoneID))
	}
	return zones
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

func (_zone zone) Interactables() []anyOfItem_Player_ZoneItem {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	var interactables []anyOfItem_Player_ZoneItem
	for _, anyOfItem_Player_ZoneItemID := range zone.zone.Interactables {
		interactables = append(interactables, zone.zone.engine.anyOfItem_Player_ZoneItem(anyOfItem_Player_ZoneItemID))
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

func (_playerTargetRef playerTargetRef) ID() AnyOfPlayer_ZoneItemID {
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

func (_playerTargetedByRef playerTargetedByRef) ID() AnyOfPlayer_ZoneItemID {
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

func (_anyOfPlayer_Position anyOfPlayer_Position) ID() AnyOfPlayer_PositionID {
	return _anyOfPlayer_Position.anyOfPlayer_Position.ID
}

func (_anyOfPlayer_Position anyOfPlayer_Position) Player() player {
	anyOfPlayer_Position := _anyOfPlayer_Position.anyOfPlayer_Position.engine.anyOfPlayer_Position(_anyOfPlayer_Position.anyOfPlayer_Position.ID)
	return anyOfPlayer_Position.anyOfPlayer_Position.engine.Player(anyOfPlayer_Position.anyOfPlayer_Position.Player)
}

func (_anyOfPlayer_Position anyOfPlayer_Position) Position() position {
	anyOfPlayer_Position := _anyOfPlayer_Position.anyOfPlayer_Position.engine.anyOfPlayer_Position(_anyOfPlayer_Position.anyOfPlayer_Position.ID)
	return anyOfPlayer_Position.anyOfPlayer_Position.engine.Position(anyOfPlayer_Position.anyOfPlayer_Position.Position)
}

func (engine *Engine) anyOfPlayer_Position(anyOfPlayer_PositionID AnyOfPlayer_PositionID) anyOfPlayer_Position {
	patchingAnyOfPlayer_Position, ok := engine.Patch.AnyOfPlayer_Position[anyOfPlayer_PositionID]
	if ok {
		return anyOfPlayer_Position{anyOfPlayer_Position: patchingAnyOfPlayer_Position}
	}
	currentAnyOfPlayer_Position, ok := engine.State.AnyOfPlayer_Position[anyOfPlayer_PositionID]
	if ok {
		return anyOfPlayer_Position{anyOfPlayer_Position: currentAnyOfPlayer_Position}
	}
	return anyOfPlayer_Position{anyOfPlayer_Position: anyOfPlayer_PositionCore{OperationKind: OperationKindDelete, engine: engine}}
}

func (_anyOfPlayer_ZoneItem anyOfPlayer_ZoneItem) ID() AnyOfPlayer_ZoneItemID {
	return _anyOfPlayer_ZoneItem.anyOfPlayer_ZoneItem.ID
}

func (_anyOfPlayer_ZoneItem anyOfPlayer_ZoneItem) Player() player {
	anyOfPlayer_ZoneItem := _anyOfPlayer_ZoneItem.anyOfPlayer_ZoneItem.engine.anyOfPlayer_ZoneItem(_anyOfPlayer_ZoneItem.anyOfPlayer_ZoneItem.ID)
	return anyOfPlayer_ZoneItem.anyOfPlayer_ZoneItem.engine.Player(anyOfPlayer_ZoneItem.anyOfPlayer_ZoneItem.Player)
}

func (_anyOfPlayer_ZoneItem anyOfPlayer_ZoneItem) ZoneItem() zoneItem {
	anyOfPlayer_ZoneItem := _anyOfPlayer_ZoneItem.anyOfPlayer_ZoneItem.engine.anyOfPlayer_ZoneItem(_anyOfPlayer_ZoneItem.anyOfPlayer_ZoneItem.ID)
	return anyOfPlayer_ZoneItem.anyOfPlayer_ZoneItem.engine.ZoneItem(anyOfPlayer_ZoneItem.anyOfPlayer_ZoneItem.ZoneItem)
}

func (engine *Engine) anyOfPlayer_ZoneItem(anyOfPlayer_ZoneItemID AnyOfPlayer_ZoneItemID) anyOfPlayer_ZoneItem {
	patchingAnyOfPlayer_ZoneItem, ok := engine.Patch.AnyOfPlayer_ZoneItem[anyOfPlayer_ZoneItemID]
	if ok {
		return anyOfPlayer_ZoneItem{anyOfPlayer_ZoneItem: patchingAnyOfPlayer_ZoneItem}
	}
	currentAnyOfPlayer_ZoneItem, ok := engine.State.AnyOfPlayer_ZoneItem[anyOfPlayer_ZoneItemID]
	if ok {
		return anyOfPlayer_ZoneItem{anyOfPlayer_ZoneItem: currentAnyOfPlayer_ZoneItem}
	}
	return anyOfPlayer_ZoneItem{anyOfPlayer_ZoneItem: anyOfPlayer_ZoneItemCore{OperationKind: OperationKindDelete, engine: engine}}
}

func (engine *Engine) anyOfItem_Player_ZoneItem(anyOfItem_Player_ZoneItemID AnyOfItem_Player_ZoneItemID) anyOfItem_Player_ZoneItem {
	patchingAnyOfItem_Player_ZoneItem, ok := engine.Patch.AnyOfItem_Player_ZoneItem[anyOfItem_Player_ZoneItemID]
	if ok {
		return anyOfItem_Player_ZoneItem{anyOfItem_Player_ZoneItem: patchingAnyOfItem_Player_ZoneItem}
	}
	currentAnyOfItem_Player_ZoneItem, ok := engine.State.AnyOfItem_Player_ZoneItem[anyOfItem_Player_ZoneItemID]
	if ok {
		return anyOfItem_Player_ZoneItem{anyOfItem_Player_ZoneItem: currentAnyOfItem_Player_ZoneItem}
	}
	return anyOfItem_Player_ZoneItem{anyOfItem_Player_ZoneItem: anyOfItem_Player_ZoneItemCore{OperationKind: OperationKindDelete, engine: engine}}
}

func (_anyOfItem_Player_ZoneItem anyOfItem_Player_ZoneItem) ID() AnyOfItem_Player_ZoneItemID {
	return _anyOfItem_Player_ZoneItem.anyOfItem_Player_ZoneItem.ID
}

func (_anyOfItem_Player_ZoneItem anyOfItem_Player_ZoneItem) Player() player {
	anyOfItem_Player_ZoneItem := _anyOfItem_Player_ZoneItem.anyOfItem_Player_ZoneItem.engine.anyOfItem_Player_ZoneItem(_anyOfItem_Player_ZoneItem.anyOfItem_Player_ZoneItem.ID)
	return anyOfItem_Player_ZoneItem.anyOfItem_Player_ZoneItem.engine.Player(anyOfItem_Player_ZoneItem.anyOfItem_Player_ZoneItem.Player)
}

func (_anyOfItem_Player_ZoneItem anyOfItem_Player_ZoneItem) ZoneItem() zoneItem {
	anyOfItem_Player_ZoneItem := _anyOfItem_Player_ZoneItem.anyOfItem_Player_ZoneItem.engine.anyOfItem_Player_ZoneItem(_anyOfItem_Player_ZoneItem.anyOfItem_Player_ZoneItem.ID)
	return anyOfItem_Player_ZoneItem.anyOfItem_Player_ZoneItem.engine.ZoneItem(anyOfItem_Player_ZoneItem.anyOfItem_Player_ZoneItem.ZoneItem)
}

func (_anyOfItem_Player_ZoneItem anyOfItem_Player_ZoneItem) Item() item {
	anyOfItem_Player_ZoneItem := _anyOfItem_Player_ZoneItem.anyOfItem_Player_ZoneItem.engine.anyOfItem_Player_ZoneItem(_anyOfItem_Player_ZoneItem.anyOfItem_Player_ZoneItem.ID)
	return anyOfItem_Player_ZoneItem.anyOfItem_Player_ZoneItem.engine.Item(anyOfItem_Player_ZoneItem.anyOfItem_Player_ZoneItem.Item)
}
