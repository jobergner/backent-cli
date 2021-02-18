package statemachine

type tree struct {
	zoneItem  map[zoneItemID]_zoneItem
	item      map[itemID]_item
	position  map[positionID]_position
	gearScore map[gearScoreID]_gearScore
}

func (t *tree) reassembleFrom(s state) {
	for _, position := range s.position {
		switch position.parentage[len(position.parentage)-1].kind {
		case entityKindPlayer:
		}
	}
}

func (t *tree) reassembleZoneItem(zoneItemID zoneItemID, s state) {
	// check if already exists in tree
	// if yes -> get it out of map, and loop if field with slice exists
	// as all non slice fields should already exist
	// thought: should there be zero values for fields that havent updated?
	// otherwise a giant object would be reconstructed
	// => optimization ??? or...
	// create operationKindZero with value 0
	// this way it'll be clear to the frontend that this element
	// does not require re-rendering/processing

// scrap that. All children referenced by id are pointers. 
// Get reassembling item out of tree, if non existent, create new one
// then check next element in Parentage, only create field that is mentioned
// if slice, just append

	t.zoneItem[zoneItemID] = _zoneItem{
		position: _position{}, // reassemblePosition()
	}
	// if zoneitem has slice field id have to loop over
	// all ids in the zoneItem (update) and append every
	// child with reassembleChild method
}

type _zoneItem struct {
	id            zoneItemID
	position      _position
	item          _item
	operationKind operationKind
}

type _item struct {
	id            itemID
	gearScore     _gearScore
	operationKind operationKind
}

func positionAdapter(p position) _position {
	return _position{
		id:            p.id,
		x:             p.x,
		y:             p.y,
		operationKind: p.operationKind,
	}
}

type _position struct {
	id            positionID
	x             float64
	y             float64
	operationKind operationKind
}

type _gearScore struct {
	id            gearScoreID
	level         int
	score         int
	operationKind operationKind
}
