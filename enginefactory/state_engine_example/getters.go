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
	return player{player: playerCore{OperationKind_: OperationKindDelete}}
}

func (_player player) ID() PlayerID {
	return _player.player.ID
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
	return gearScore{gearScore: gearScoreCore{OperationKind_: OperationKindDelete}}
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
	return item{item: itemCore{OperationKind_: OperationKindDelete}}
}

func (_item item) ID() ItemID {
	return _item.item.ID
}

func (_item item) GearScore() gearScore {
	item := _item.item.engine.Item(_item.item.ID)
	return item.item.engine.GearScore(item.item.GearScore)
}

func (_item item) BoundTo() (itemBoundToRef, bool) {
	item := _item.item.engine.Item(_item.item.ID)
	return item.item.engine.itemBoundToRef(item.item.BoundTo), item.item.BoundTo != 0
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
	return position{position: positionCore{OperationKind_: OperationKindDelete}}
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
	return zoneItem{zoneItem: zoneItemCore{OperationKind_: OperationKindDelete}}
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
	return zone{zone: zoneCore{OperationKind_: OperationKindDelete}}
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

func (ref itemBoundToRef) ID() PlayerID {
	return ref.itemBoundToRef.ReferencedElementID
}

func (engine *Engine) itemBoundToRef(itemBoundToRefID ItemBoundToRefID) itemBoundToRef {
	patchingRef, ok := engine.Patch.ItemBoundToRef[itemBoundToRefID]
	if ok {
		return itemBoundToRef{itemBoundToRef: patchingRef}
	}
	currentRef, ok := engine.State.ItemBoundToRef[itemBoundToRefID]
	if ok {
		return itemBoundToRef{itemBoundToRef: currentRef}
	}
	return itemBoundToRef{itemBoundToRef: itemBoundToRefCore{OperationKind_: OperationKindDelete}}
}

func (ref playerGuildMemberRef) ID() PlayerID {
	return ref.playerGuildMemberRef.ReferencedElementID
}

func (engine *Engine) playerGuildMemberRef(playerGuildMemberRefID PlayerGuildMemberRefID) playerGuildMemberRef {
	patchingRef, ok := engine.Patch.PlayerGuildMemberRef[playerGuildMemberRefID]
	if ok {
		return playerGuildMemberRef{playerGuildMemberRef: patchingRef}
	}
	currentRef, ok := engine.State.PlayerGuildMemberRef[playerGuildMemberRefID]
	if ok {
		return playerGuildMemberRef{playerGuildMemberRef: currentRef}
	}
	return playerGuildMemberRef{playerGuildMemberRef: playerGuildMemberRefCore{OperationKind_: OperationKindDelete}}
}

func (engine *Engine) playerEquipmentSetRef(playerEquipmentSetRefID PlayerEquipmentSetRefID) playerEquipmentSetRef {
	patchingRef, ok := engine.Patch.PlayerEquipmentSetRef[playerEquipmentSetRefID]
	if ok {
		return playerEquipmentSetRef{playerEquipmentSetRef: patchingRef}
	}
	currentRef, ok := engine.State.PlayerEquipmentSetRef[playerEquipmentSetRefID]
	if ok {
		return playerEquipmentSetRef{playerEquipmentSetRef: currentRef}
	}
	return playerEquipmentSetRef{playerEquipmentSetRef: playerEquipmentSetRefCore{OperationKind_: OperationKindDelete}}
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
	return equipmentSet{equipmentSet: equipmentSetCore{OperationKind_: OperationKindDelete}}
}

func (_equipmentSet equipmentSet) ID() EquipmentSetID {
	return _equipmentSet.equipmentSet.ID
}

func (_equipmentSet equipmentSet) Equipment() []equipmentSetEquipmentRef {
	equipmentSet := _equipmentSet.equipmentSet.engine.EquipmentSet(_equipmentSet.equipmentSet.ID)
	var equipment []equipmentSetEquipmentRef
	for _, refID := range equipmentSet.equipmentSet.Equipment {
		equipment = append(equipment, equipmentSet.equipmentSet.engine.equipmentSetEquipmentRef(refID))
	}
	return equipment
}

func (engine *Engine) equipmentSetEquipmentRef(equipmentSetEquipmentRefID EquipmentSetEquipmentRefID) equipmentSetEquipmentRef {
	patchingRef, ok := engine.Patch.EquipmentSetEquipmentRef[equipmentSetEquipmentRefID]
	if ok {
		return equipmentSetEquipmentRef{equipmentSetEquipmentRef: patchingRef}
	}
	currentRef, ok := engine.State.EquipmentSetEquipmentRef[equipmentSetEquipmentRefID]
	if ok {
		return equipmentSetEquipmentRef{equipmentSetEquipmentRef: currentRef}
	}
	return equipmentSetEquipmentRef{equipmentSetEquipmentRef: equipmentSetEquipmentRefCore{OperationKind_: OperationKindDelete}}
}
