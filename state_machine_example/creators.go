package statemachine

func (sm *StateMachine) CreateGearScore(parentage ...ParentInfo) GearScore {
	var gearScore GearScoreCore
	gearScore.ID = GearScoreID(sm.GenerateID())
	gearScore.Parentage = append(gearScore.Parentage, parentage...)
	gearScore.OperationKind = OperationKindUpdate
	sm.Patch.GearScore[gearScore.ID] = gearScore
	return GearScore{gearScore: gearScore}
}

func (sm *StateMachine) CreatePosition(parentage ...ParentInfo) Position {
	var position PositionCore
	position.ID = PositionID(sm.GenerateID())
	position.Parentage = append(position.Parentage, parentage...)
	position.OperationKind = OperationKindUpdate
	sm.Patch.Position[position.ID] = position
	return Position{position: position}
}

func (sm *StateMachine) CreateItem(parentage ...ParentInfo) Item {
	var item ItemCore
	item.ID = ItemID(sm.GenerateID())
	item.Parentage = append(item.Parentage, parentage...)
	elementGearScore := sm.CreateGearScore(append(item.Parentage, ParentInfo{EntityKindItem, int(item.ID)})...)
	item.GearScore = elementGearScore.gearScore.ID
	item.OperationKind = OperationKindUpdate
	sm.Patch.Item[item.ID] = item
	return Item{item: item}
}

func (sm *StateMachine) CreateZoneItem(parentage ...ParentInfo) ZoneItem {
	var zoneItem ZoneItemCore
	zoneItem.ID = ZoneItemID(sm.GenerateID())
	zoneItem.Parentage = append(zoneItem.Parentage, parentage...)
	elementItem := sm.CreateItem(append(zoneItem.Parentage, ParentInfo{EntityKindZoneItem, int(zoneItem.ID)})...)
	zoneItem.Item = elementItem.item.ID
	elementPosition := sm.CreatePosition(append(zoneItem.Parentage, ParentInfo{EntityKindZoneItem, int(zoneItem.ID)})...)
	zoneItem.Position = elementPosition.position.ID
	zoneItem.OperationKind = OperationKindUpdate
	sm.Patch.ZoneItem[zoneItem.ID] = zoneItem
	return ZoneItem{zoneItem: zoneItem}
}

func (sm *StateMachine) CreatePlayer(parentage ...ParentInfo) Player {
	var player PlayerCore
	player.ID = PlayerID(sm.GenerateID())
	player.Parentage = append(player.Parentage, parentage...)
	elementGearScore := sm.CreateGearScore(append(player.Parentage, ParentInfo{EntityKindPlayer, int(player.ID)})...)
	player.GearScore = elementGearScore.gearScore.ID
	elementPosition := sm.CreatePosition(append(player.Parentage, ParentInfo{EntityKindPlayer, int(player.ID)})...)
	player.Position = elementPosition.position.ID
	player.OperationKind = OperationKindUpdate
	sm.Patch.Player[player.ID] = player
	return Player{player: player}
}

func (sm *StateMachine) CreateZone() Zone {
	var zone ZoneCore
	zone.ID = ZoneID(sm.GenerateID())
	zone.OperationKind = OperationKindUpdate
	sm.Patch.Zone[zone.ID] = zone
	return Zone{zone: zone}
}
