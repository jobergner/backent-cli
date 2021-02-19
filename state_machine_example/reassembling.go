package statemachine

func (t *tree) reassembleFrom(s state) {
	for _, player := range s.player {
		t.reassembleByEnityKind(player.parentage, s)
	}
	for _, zoneItem := range s.zoneItem {
		t.reassembleByEnityKind(zoneItem.parentage, s)
	}
	for _, position := range s.position {
		t.reassembleByEnityKind(position.parentage, s)
	}
	for _, item := range s.item {
		t.reassembleByEnityKind(item.parentage, s)
	}
	for _, gearScore := range s.gearScore {
		t.reassembleByEnityKind(gearScore.parentage, s)
	}
}

func (t *tree) reassembleByEnityKind(parentage parentage, s state) {
	greatestAncestor := parentage[0]
	switch greatestAncestor.kind {
	case entityKindZoneItem:
		_zoneItem := t.reassembleZoneItem(parentage[1:], zoneItemID(greatestAncestor.id), s)
		t.zoneItem[zoneItemID(greatestAncestor.id)] = _zoneItem
	}
}

func (t tree) reassembleZone(parentage parentage, zoneID zoneID, s state) _zone {
	_zone := t.zone[zoneID]
	zone := s.zone[zoneID]
	_zone.id = zone.id
	_zone.operationKind = zone.operationKind

	nextDescendant := parentage[0]
	switch nextDescendant.kind {
	case entityKindPlayer:
		_player := t.reassemblePlayer(parentage[1:], playerID(nextDescendant.id), s)
		_zone.players = append(_zone.players, _player)
	case entityKindZoneItem:
		_zoneItem := t.reassembleZoneItem(parentage[1:], zoneItemID(nextDescendant.id), s)
		_zone.items = append(_zone.items, _zoneItem)
	}
	return _zone
}

func (t tree) reassemblePlayer(parentage parentage, playerID playerID, s state) _player {
	_player := t.player[playerID]
	player := s.player[playerID]
	_player.id = player.id
	_player.operationKind = player.operationKind

	nextDescendant := parentage[0]
	switch nextDescendant.kind {
	case entityKindPosition:
		_position := t.reassemblePosition(positionID(nextDescendant.id), s)
		_player.position = &_position
	case entityKindGearScore:
		_gearScore := t.reassembleGearScore(gearScoreID(nextDescendant.id), s)
		_player.gearScore = &_gearScore
	case entityKindItem:
		_item := t.reassembleItem(parentage[1:], itemID(nextDescendant.id), s)
		_player.items = append(_player.items, _item)
	}
	return _player
}

func (t tree) reassembleZoneItem(parentage parentage, zoneItemID zoneItemID, s state) _zoneItem {
	_zoneItem := t.zoneItem[zoneItemID]
	zoneItem := s.zoneItem[zoneItemID]
	_zoneItem.id = zoneItem.id
	_zoneItem.operationKind = zoneItem.operationKind

	nextDescendant := parentage[0]
	switch nextDescendant.kind {
	case entityKindPosition:
		_position := t.reassemblePosition(positionID(nextDescendant.id), s)
		_zoneItem.position = &_position
	case entityKindItem:
		_item := t.reassembleItem(parentage[1:], itemID(nextDescendant.id), s)
		_zoneItem.item = &_item
	}
	return _zoneItem
}

func (t tree) reassemblePosition(positionID positionID, s state) _position {
	_position := t.position[positionID]
	position := s.position[positionID]
	_position.id = position.id
	_position.operationKind = position.operationKind

	_position.x = position.x
	_position.y = position.y
	return _position
}

func (t tree) reassembleItem(parentage parentage, itemID itemID, s state) _item {
	_item := t.item[itemID]
	item := s.item[itemID]
	_item.id = item.id
	_item.operationKind = item.operationKind

	nextDescendant := parentage[0]
	switch nextDescendant.kind {
	case entityKindGearScore:
		_gearScore := t.reassembleGearScore(gearScoreID(nextDescendant.id), s)
		_item.gearScore = &_gearScore
	}
	return _item
}

func (t tree) reassembleGearScore(gearScoreID gearScoreID, s state) _gearScore {
	_gearScore := t.gearScore[gearScoreID]
	gearScore := s.gearScore[gearScoreID]
	_gearScore.id = gearScore.id
	_gearScore.operationKind = gearScore.operationKind

	_gearScore.level = gearScore.level
	_gearScore.score = gearScore.score
	return _gearScore
}
