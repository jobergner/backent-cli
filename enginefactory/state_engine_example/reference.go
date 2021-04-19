package state

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

func (ref *gearScoreRef) Get(se *Engine) gearScore {
	return se.GearScore(ref.id)
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

func (ref *itemRef) Get(se *Engine) item {
	return se.Item(ref.id)
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

func (ref *playerRef) Get(se *Engine) player {
	return se.Player(ref.id)
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

func (ref *positionRef) Get(se *Engine) position {
	return se.Position(ref.id)
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

func (ref *zoneRef) Get(se *Engine) zone {
	return se.Zone(ref.id)
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

func (ref *zoneItemRef) Get(se *Engine) zoneItem {
	return se.ZoneItem(ref.id)
}

func (se *Engine) setElementUpdated(id int, kind ElementKind) {
	switch kind {
	case ElementKindGearScore:
		gearScore := se.GearScore(GearScoreID(id))
		gearScore.gearScore.OperationKind_ = OperationKindUpdate
		se.Patch.GearScore[gearScore.gearScore.ID] = gearScore.gearScore
		break
	case ElementKindItem:
		item := se.Item(ItemID(id))
		item.item.OperationKind_ = OperationKindUpdate
		se.Patch.Item[item.item.ID] = item.item
		break
	case ElementKindPlayer:
		player := se.Player(PlayerID(id))
		player.player.OperationKind_ = OperationKindUpdate
		se.Patch.Player[player.player.ID] = player.player
		break
	case ElementKindPosition:
		position := se.Position(PositionID(id))
		position.position.OperationKind_ = OperationKindUpdate
		se.Patch.Position[position.position.ID] = position.position
		break
	case ElementKindZone:
		zone := se.Zone(ZoneID(id))
		zone.zone.OperationKind_ = OperationKindUpdate
		se.Patch.Zone[zone.zone.ID] = zone.zone
		break
	case ElementKindZoneItem:
		zoneItem := se.ZoneItem(ZoneItemID(id))
		zoneItem.zoneItem.OperationKind_ = OperationKindUpdate
		se.Patch.ZoneItem[zoneItem.zoneItem.ID] = zoneItem.zoneItem
		break
	}
}
