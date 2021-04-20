package state

func (_zone zone) AddPlayer(se *Engine) player {
	zone := se.Zone(_zone.zone.ID)
	if zone.zone.OperationKind_ == OperationKindDelete {
		return player{player: playerCore{OperationKind_: OperationKindDelete}}
	}
	player := se.createPlayer(true)
	zone.zone.Players = append(zone.zone.Players, player.player.ID)
	se.updateZone(zone.zone)
	return player
}

func (_zone zone) AddItem(se *Engine) zoneItem {
	zone := se.Zone(_zone.zone.ID)
	if zone.zone.OperationKind_ == OperationKindDelete {
		return zoneItem{zoneItem: zoneItemCore{OperationKind_: OperationKindDelete}}
	}
	zoneItem := se.createZoneItem(true)
	zone.zone.Items = append(zone.zone.Items, zoneItem.zoneItem.ID)
	se.updateZone(zone.zone)
	return zoneItem
}

func (_zone zone) AddTags(se *Engine, tags ...string) {
	zone := se.Zone(_zone.zone.ID)
	if zone.zone.OperationKind_ == OperationKindDelete {
		return
	}
	zone.zone.Tags = append(zone.zone.Tags, tags...)
	se.updateZone(zone.zone)
}

func (_player player) AddItem(se *Engine) item {
	player := se.Player(_player.player.ID)
	if player.player.OperationKind_ == OperationKindDelete {
		return item{item: itemCore{OperationKind_: OperationKindDelete}}
	}
	item := se.createItem(true)
	player.player.Items = append(player.player.Items, item.item.ID)
	se.updatePlayer(player.player)
	return item
}

func (_player player) AddGuildMember(se *Engine, playerID PlayerID) {
	player := se.Player(_player.player.ID)
	if player.player.OperationKind_ == OperationKindDelete {
		return
	}
	player.player.GuildMembers = append(player.player.GuildMembers, playerGuildMembersSliceRef{playerID, player.player.ID})
	se.updatePlayer(player.player)
}
