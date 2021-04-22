package state

func (se *Engine) updateGearScoreUpstream(gearScore gearScoreCore) {
	gearScore.OperationKind_ = OperationKindUpdate
	se.Patch.GearScore[gearScore.ID] = gearScore
}
func (se *Engine) updateItemUpstream(item itemCore) {
	item.OperationKind_ = OperationKindUpdate
	se.Patch.Item[item.ID] = item
}
func (se *Engine) updatePlayerUpstream(player playerCore, updatedPlayerIDs []PlayerID) {
	for _, itemID := range se.allItemIDs() {
		_item := se.Item(itemID)
		if _item.BoundTo(se).ID(se) == player.ID {
			se.updateItemUpstream(_item.item)
		}
	}
	for _, playerID := range se.allPlayerIDs() {
		_player := se.Player(playerID)
	OUTER:
		for _, ref := range _player.GuildMembers(se) {
			if ref.playerGuildMemberRef.ReferencedElementID == player.ID {
				for _, updatedPlayerID := range updatedPlayerIDs {
					if ref.playerGuildMemberRef.ReferencedElementID == updatedPlayerID {
						continue OUTER
					}
				}
				updatedPlayerIDs = append(updatedPlayerIDs, player.ID)
				se.updatePlayerUpstream(_player.player, updatedPlayerIDs)
			}
		}
	}
	player.OperationKind_ = OperationKindUpdate
	se.Patch.Player[player.ID] = player
}
func (se *Engine) updatePositionUpstream(position positionCore) {
	position.OperationKind_ = OperationKindUpdate
	se.Patch.Position[position.ID] = position
}
func (se *Engine) updateZoneUpstream(zone zoneCore) {
	zone.OperationKind_ = OperationKindUpdate
	se.Patch.Zone[zone.ID] = zone
}
func (se *Engine) updateZoneItemUpstream(zoneItem zoneItemCore) {
	zoneItem.OperationKind_ = OperationKindUpdate
	se.Patch.ZoneItem[zoneItem.ID] = zoneItem
}
