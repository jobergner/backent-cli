package statemachine

func (t *Tree) assembleFrom(s State) {
	for _, player := range s.Player {
		t.assembleByEnityKind(player.Parentage, s)
	}
	for _, zoneItem := range s.ZoneItem {
		t.assembleByEnityKind(zoneItem.Parentage, s)
	}
	for _, position := range s.Position {
		t.assembleByEnityKind(position.Parentage, s)
	}
	for _, item := range s.Item {
		t.assembleByEnityKind(item.Parentage, s)
	}
	for _, gearScore := range s.GearScore {
		t.assembleByEnityKind(gearScore.Parentage, s)
	}
}

func (t *Tree) assembleByEnityKind(parentage Parentage, s State) {
	greatestAncestor := parentage[0]
	switch greatestAncestor.Kind {
	case EntityKindZoneItem:
		_zoneItem := t.assembleZoneItem(parentage[1:], ZoneItemID(greatestAncestor.ID), s)
		t.ZoneItem[ZoneItemID(greatestAncestor.ID)] = _zoneItem
	}
}

func (t Tree) assembleZone(parentage Parentage, zoneID ZoneID, s State) _zone {
	_zone := t.Zone[zoneID]
	zone := s.Zone[zoneID]
	_zone.ID = zone.ID
	_zone.OperationKind = zone.OperationKind

	nextDescendant := parentage[0]
	switch nextDescendant.Kind {
	case EntityKindPlayer:
		_player := t.assemblePlayer(parentage[1:], PlayerID(nextDescendant.ID), s)
		_zone.Players = append(_zone.Players, _player)
	case EntityKindZoneItem:
		_zoneItem := t.assembleZoneItem(parentage[1:], ZoneItemID(nextDescendant.ID), s)
		_zone.Items = append(_zone.Items, _zoneItem)
	}
	return _zone
}

func (t Tree) assemblePlayer(parentage Parentage, playerID PlayerID, s State) _player {
	_player := t.Player[playerID]
	player := s.Player[playerID]
	_player.ID = player.ID
	_player.OperationKind = player.OperationKind

	nextDescendant := parentage[0]
	switch nextDescendant.Kind {
	case EntityKindPosition:
		_position := t.assemblePosition(PositionID(nextDescendant.ID), s)
		_player.Position = &_position
	case EntityKindGearScore:
		_gearScore := t.assembleGearScore(GearScoreID(nextDescendant.ID), s)
		_player.GearScore = &_gearScore
	case EntityKindItem:
		_item := t.assembleItem(parentage[1:], ItemID(nextDescendant.ID), s)
		_player.Items = append(_player.Items, _item)
	}
	return _player
}

func (t Tree) assembleZoneItem(parentage Parentage, zoneItemID ZoneItemID, s State) _zoneItem {
	_zoneItem := t.ZoneItem[zoneItemID]
	zoneItem := s.ZoneItem[zoneItemID]
	_zoneItem.ID = zoneItem.ID
	_zoneItem.OperationKind = zoneItem.OperationKind

	nextDescendant := parentage[0]
	switch nextDescendant.Kind {
	case EntityKindPosition:
		_position := t.assemblePosition(PositionID(nextDescendant.ID), s)
		_zoneItem.Position = &_position
	case EntityKindItem:
		_item := t.assembleItem(parentage[1:], ItemID(nextDescendant.ID), s)
		_zoneItem.Item = &_item
	}
	return _zoneItem
}

func (t Tree) assemblePosition(positionID PositionID, s State) _position {
	_position := t.Position[positionID]
	position := s.Position[positionID]
	_position.ID = position.ID
	_position.OperationKind = position.OperationKind

	_position.X = position.X
	_position.Y = position.Y
	return _position
}

func (t Tree) assembleItem(parentage Parentage, itemID ItemID, s State) _item {
	_item := t.Item[itemID]
	item := s.Item[itemID]
	_item.ID = item.ID
	_item.OperationKind = item.OperationKind

	nextDescendant := parentage[0]
	switch nextDescendant.Kind {
	case EntityKindGearScore:
		_gearScore := t.assembleGearScore(GearScoreID(nextDescendant.ID), s)
		_item.GearScore = &_gearScore
	}
	return _item
}

func (t Tree) assembleGearScore(gearScoreID GearScoreID, s State) _gearScore {
	_gearScore := t.GearScore[gearScoreID]
	gearScore := s.GearScore[gearScoreID]
	_gearScore.ID = gearScore.ID
	_gearScore.OperationKind = gearScore.OperationKind

	_gearScore.Level = gearScore.Level
	_gearScore.Score = gearScore.Score
	return _gearScore
}
