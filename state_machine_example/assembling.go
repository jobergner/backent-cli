package statemachine

// this is kind of a mess atm
// the implementation completely missed the idea
// this code is supposed to only run in the wasm client
// its purpose is to assemble the data of the incoming patch in a tree
// so it would be called on the stateMachine in the client
// sm.Patch = incomingPatch
// dataTree := sm.assembleTree()
// sm.UpdateState()
// ...
// the tree is assembled from the patch and fills in the missing parents of elements
// with the elements it has in it's state.
// the implementation would look like this
// func (sm *StateMachine) assemble() Tree {
// the inteesting thing of the tree is that it really only holds updated data and their parents,
// and will omit children of elements that haven't updated (with the use of pointers)
//
// the implementation should also not be too hard to write.
// I just need to loop through each elementKind and of State AND Patch
// and check for len(parentage) == 0
// from there I build ALL children, BUT have 2 returned values:
// 1. the element
// 2. a boolean value whether the actually was an item taken from the Patch, and not only from the State
// (alternative is the item always gets returned as pointer, and is just nil in case)
// if the boolean is false, the built element get's dircarded

func (t *Tree) assemble(s State) *Tree {
	for _, gearScore := range s.GearScore {
		t.assembleUpstreamFromGearScore(gearScore.ID, s)
	}
	for _, item := range s.Item {
		t.assembleUpstreamFromItem(item.Parentage, item.ID, s)
	}
	for _, player := range s.Player {
		t.assembleUpstreamFromPlayer(player.Parentage, player.ID, s)
	}
	for _, position := range s.Position {
		t.assembleUpstreamFromPosition(position.ID, s)
	}
	for _, zoneItem := range s.ZoneItem {
		t.assembleUpstreamFromZoneItem(zoneItem.Parentage, zoneItem.ID, s)
	}
	// needs zone here too, just add it

	return t
}

func (t Tree) assembleUpstreamFromPlayer(parentage Parentage, playerID PlayerID, s State) _player {
	_player := t.Player[playerID]
	player := s.Player[playerID]
	_player.ID = player.ID
	_player.OperationKind = player.OperationKind

	nextDescendant := parentage[0]
	switch nextDescendant.Kind {
	case EntityKindPosition:
		_position := t.assembleUpstreamFromPosition(PositionID(nextDescendant.ID), s)
		_player.Position = &_position
	case EntityKindGearScore:
		_gearScore := t.assembleUpstreamFromGearScore(GearScoreID(nextDescendant.ID), s)
		_player.GearScore = &_gearScore
	case EntityKindItem:
		_item := t.assembleUpstreamFromItem(parentage[1:], ItemID(nextDescendant.ID), s)
		_player.Items = append(_player.Items, _item)
	}
	return _player
}

func (t Tree) assembleUpstreamFromZoneItem(parentage Parentage, zoneItemID ZoneItemID, s State) _zoneItem {
	_zoneItem := t.ZoneItem[zoneItemID]
	zoneItem := s.ZoneItem[zoneItemID]
	_zoneItem.ID = zoneItem.ID
	_zoneItem.OperationKind = zoneItem.OperationKind

	nextDescendant := parentage[0]
	switch nextDescendant.Kind {
	case EntityKindPosition:
		_position := t.assembleUpstreamFromPosition(PositionID(nextDescendant.ID), s)
		_zoneItem.Position = &_position
	case EntityKindItem:
		_item := t.assembleUpstreamFromItem(parentage[1:], ItemID(nextDescendant.ID), s)
		_zoneItem.Item = &_item
	}
	return _zoneItem
}

func (t Tree) assembleUpstreamFromPosition(positionID PositionID, s State) _position {
	_position := t.Position[positionID]
	position := s.Position[positionID]
	_position.ID = position.ID
	_position.OperationKind = position.OperationKind

	_position.X = position.X
	_position.Y = position.Y
	return _position
}

func (t Tree) assembleUpstreamFromItem(itemID ItemID, s State) _item {
	_item := t.Item[itemID]
	item := s.Item[itemID]
	_item.ID = item.ID
	_item.OperationKind = item.OperationKind

	nextDescendant := parentage[0]
	switch nextDescendant.Kind {
	case EntityKindGearScore:
		_gearScore := t.assembleUpstreamFromGearScore(GearScoreID(nextDescendant.ID), s)
		_item.GearScore = &_gearScore
	}
	return _item
}

func (t Tree) assembleUpstreamFromGearScore(gearScoreID GearScoreID, s State) _gearScore {
	_gearScore := t.GearScore[gearScoreID]
	gearScore := s.GearScore[gearScoreID]
	_gearScore.ID = gearScore.ID
	_gearScore.OperationKind = gearScore.OperationKind

	_gearScore.Level = gearScore.Level
	_gearScore.Score = gearScore.Score

	if len(gearScore.Parentage) == 0 {
		return _gearScore
	}

	nextParent := gearScore.Parentage[len(gearScore.Parentage)-1]
	switch nextParent.Kind {
	case EntityKindItem:
		_gearScore := t.assembleUpstreamFromItem(ItemID(nextParent.ID), s)
		_item.GearScore = &_gearScore
	}

	return _gearScore
}
