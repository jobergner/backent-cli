package state

func (se *Engine) updateGearScore(gearScore gearScoreCore) {
	gearScore.OperationKind_ = OperationKindUpdate
	se.Patch.GearScore[gearScore.ID] = gearScore
}
func (se *Engine) updateItem(item itemCore) {
	item.OperationKind_ = OperationKindUpdate
	se.Patch.Item[item.ID] = item
}
func (se *Engine) updatePlayer(player playerCore) {
	player.OperationKind_ = OperationKindUpdate
	se.Patch.Player[player.ID] = player
	for _, itemID := range se.allItemIDs() {
		element := se.Item(itemID)
		if element.item.BoundTo.id == player.ID {
			if _, alreadyUpdating := se.Patch.Item[element.item.ID]; !alreadyUpdating {
				se.updateItem(element.item)
			}
		}
	}
	for _, playerID := range se.allPlayerIDs() {
		element := se.Player(playerID)
		for _, ref := range element.player.GuildMembers {
			if ref.id == player.ID {
				if _, alreadyUpdating := se.Patch.Player[element.player.ID]; !alreadyUpdating {
					se.updatePlayer(element.player)
				}
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
