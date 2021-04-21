package state

func (se *Engine) updateGearScore(gearScore gearScoreCore) {
	gearScore.OperationKind_ = OperationKindUpdate
	se.Patch.GearScore[gearScore.ID] = gearScore
}
func (se *Engine) updateItem(item itemCore) {
	item.OperationKind_ = OperationKindUpdate
	se.Patch.Item[item.ID] = item
}

// TODO rename?
func (se *Engine) updatePlayer(player playerCore, updatedPlayerIDs []PlayerID) {
	player.OperationKind_ = OperationKindUpdate
	se.Patch.Player[player.ID] = player
	updatedPlayerIDs = append(updatedPlayerIDs, player.ID)
	for _, itemID := range se.allItemIDs() {
		_item := se.Item(itemID)
		if _item.BoundTo(se).ID(se) == player.ID {
			se.updateItem(_item.item)
		}
	}
	for _, playerID := range se.allPlayerIDs() {
		_player := se.Player(playerID)
	OUTER:
		for _, ref := range _player.GuildMembers(se) {
			if ref.playerGuildMemberRef.ReferencedElementID == playerID {
				for _, updatedPlayerID := range updatedPlayerIDs {
					if ref.playerGuildMemberRef.ReferencedElementID == updatedPlayerID {
						continue OUTER
					}
				}
				se.updatePlayer(ref.Get(se).player, updatedPlayerIDs)
			}
		}
	}
}
func (se *Engine) updatePosition(position positionCore) {
	position.OperationKind_ = OperationKindUpdate
	se.Patch.Position[position.ID] = position
}
func (se *Engine) updateZone(zone zoneCore) {
	zone.OperationKind_ = OperationKindUpdate
	se.Patch.Zone[zone.ID] = zone
}
func (se *Engine) updateZoneItem(zoneItem zoneItemCore) {
	zoneItem.OperationKind_ = OperationKindUpdate
	se.Patch.ZoneItem[zoneItem.ID] = zoneItem
}
