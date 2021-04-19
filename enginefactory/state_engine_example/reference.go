package state

type itemBoundToRef struct {
	id       PlayerID
	parentID ItemID
}

func (ref itemBoundToRef) IsSet(se *Engine) bool {
	return ref.id != 0
}

func (ref itemBoundToRef) Unset(se *Engine) {
	ref.id = 0
	item := se.Item(ref.parentID).item
	item.BoundTo = ref
	se.updateItem(item)
}

func (ref itemBoundToRef) Set(se *Engine, id PlayerID) {
	ref.id = id
	item := se.Item(ref.parentID).item
	item.BoundTo = ref
	se.updateItem(item)
}

func (ref itemBoundToRef) Get(se *Engine) player {
	return se.Player(ref.id)
}

type gearScoreSliceRef struct {
	id       GearScoreID
	parentID int
}

func (ref gearScoreSliceRef) Get(se *Engine) gearScore {
	return se.GearScore(ref.id)
}

type itemSliceRef struct {
	id       ItemID
	parentID int
}

func (ref itemSliceRef) Get(se *Engine) item {
	return se.Item(ref.id)
}

type playerSliceRef struct {
	id       PlayerID
	parentID int
}

func (ref playerSliceRef) Get(se *Engine) player {
	return se.Player(ref.id)
}

type positionSliceRef struct {
	id       PositionID
	parentID int
}

func (ref positionSliceRef) Get(se *Engine) position {
	return se.Position(ref.id)
}

type zoneSliceRef struct {
	id       ZoneID
	parentID int
}

func (ref zoneSliceRef) Get(se *Engine) zone {
	return se.Zone(ref.id)
}

type zoneItemSliceRef struct {
	id       ZoneItemID
	parentID int
}

func (ref zoneItemSliceRef) Get(se *Engine) zoneItem {
	return se.ZoneItem(ref.id)
}
