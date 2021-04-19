package state

type gearScoreSliceRef struct {
	id         GearScoreID
	parentID   int
	parentKind ElementKind
}

func (ref gearScoreSliceRef) Get(se *Engine) gearScore {
	return se.GearScore(ref.id)
}

type gearScoreRef struct {
	id         GearScoreID
	parentID   int
	parentKind ElementKind
}

func (ref gearScoreRef) IsSet(se *Engine) bool {
	return ref.id != 0
}

func (ref gearScoreRef) Unset(se *Engine) {
	ref.id = 0
	se.setElementUpdated(ref.parentID, ref.parentKind)
}

func (ref *gearScoreRef) Set(se *Engine, id GearScoreID) {
	ref.id = id
	se.setElementUpdated(ref.parentID, ref.parentKind)
}

func (ref gearScoreRef) Get(se *Engine) gearScore {
	return se.GearScore(ref.id)
}

type itemSliceRef struct {
	id         ItemID
	parentID   int
	parentKind ElementKind
}

func (ref itemSliceRef) Get(se *Engine) item {
	return se.Item(ref.id)
}

type itemRef struct {
	id         ItemID
	parentID   int
	parentKind ElementKind
}

func (ref itemRef) IsSet(se *Engine) bool {
	return ref.id != 0
}

func (ref itemRef) Unset(se *Engine) {
	ref.id = 0
	se.setElementUpdated(ref.parentID, ref.parentKind)
}

func (ref *itemRef) Set(se *Engine, id ItemID) {
	ref.id = id
	se.setElementUpdated(ref.parentID, ref.parentKind)
}

func (ref itemRef) Get(se *Engine) item {
	return se.Item(ref.id)
}

type playerSliceRef struct {
	id         PlayerID
	parentID   int
	parentKind ElementKind
}

func (ref playerSliceRef) Get(se *Engine) player {
	return se.Player(ref.id)
}

type playerRef struct {
	id         PlayerID
	parentID   int
	parentKind ElementKind
}

func (ref playerRef) IsSet(se *Engine) bool {
	return ref.id != 0
}

func (ref playerRef) Unset(se *Engine) {
	ref.id = 0
	se.setElementUpdated(ref.parentID, ref.parentKind)
}

func (ref *playerRef) Set(se *Engine, id PlayerID) {
	ref.id = id
	se.setElementUpdated(ref.parentID, ref.parentKind)
}

func (ref playerRef) Get(se *Engine) player {
	return se.Player(ref.id)
}

type positionSliceRef struct {
	id         PositionID
	parentID   int
	parentKind ElementKind
}

func (ref positionSliceRef) Get(se *Engine) position {
	return se.Position(ref.id)
}

type positionRef struct {
	id         PositionID
	parentID   int
	parentKind ElementKind
}

func (ref positionRef) IsSet(se *Engine) bool {
	return ref.id != 0
}

func (ref positionRef) Unset(se *Engine) {
	ref.id = 0
	se.setElementUpdated(ref.parentID, ref.parentKind)
}

func (ref *positionRef) Set(se *Engine, id PositionID) {
	ref.id = id
	se.setElementUpdated(ref.parentID, ref.parentKind)
}

func (ref positionRef) Get(se *Engine) position {
	return se.Position(ref.id)
}

type zoneSliceRef struct {
	id         ZoneID
	parentID   int
	parentKind ElementKind
}

func (ref zoneSliceRef) Get(se *Engine) zone {
	return se.Zone(ref.id)
}

type zoneRef struct {
	id         ZoneID
	parentID   int
	parentKind ElementKind
}

func (ref zoneRef) IsSet(se *Engine) bool {
	return ref.id != 0
}

func (ref zoneRef) Unset(se *Engine) {
	ref.id = 0
	se.setElementUpdated(ref.parentID, ref.parentKind)
}

func (ref *zoneRef) Set(se *Engine, id ZoneID) {
	ref.id = id
	se.setElementUpdated(ref.parentID, ref.parentKind)
}

func (ref zoneRef) Get(se *Engine) zone {
	return se.Zone(ref.id)
}

type zoneItemSliceRef struct {
	id         ZoneItemID
	parentID   int
	parentKind ElementKind
}

func (ref zoneItemSliceRef) Get(se *Engine) zoneItem {
	return se.ZoneItem(ref.id)
}

type zoneItemRef struct {
	id         ZoneItemID
	parentID   int
	parentKind ElementKind
}

func (ref zoneItemRef) IsSet(se *Engine) bool {
	return ref.id != 0
}

func (ref zoneItemRef) Unset(se *Engine) {
	ref.id = 0
	se.setElementUpdated(ref.parentID, ref.parentKind)
}

func (ref *zoneItemRef) Set(se *Engine, id ZoneItemID) {
	ref.id = id
	se.setElementUpdated(ref.parentID, ref.parentKind)
}

func (ref zoneItemRef) Get(se *Engine) zoneItem {
	return se.ZoneItem(ref.id)
}

func (se *Engine) setElementUpdated(id int, kind ElementKind) {
	switch kind {
	case ElementKindGearScore:
		gearScore := se.GearScore(GearScoreID(id))
		se.updateGearScore(gearScore.gearScore)
		break
	case ElementKindItem:
		item := se.Item(ItemID(id))
		se.updateItem(item.item)
		break
	case ElementKindPlayer:
		player := se.Player(PlayerID(id))
		se.updatePlayer(player.player)
		break
	case ElementKindPosition:
		position := se.Position(PositionID(id))
		se.updatePosition(position.position)
		break
	case ElementKindZone:
		zone := se.Zone(ZoneID(id))
		se.updateZone(zone.zone)
		break
	case ElementKindZoneItem:
		zoneItem := se.ZoneItem(ZoneItemID(id))
		se.updateZoneItem(zoneItem.zoneItem)
		break
	}
}
