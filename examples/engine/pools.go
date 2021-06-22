package state

import (
	"sync"
)

var zoneItemCheckPool = sync.Pool{
	New: func() interface{} { return make(map[ZoneItemID]bool) },
}
var zoneItemIDSlicePool = sync.Pool{
	New: func() interface{} { return make([]ZoneItemID, 10) },
}

var zoneCheckPool = sync.Pool{
	New: func() interface{} { return make(map[ZoneID]bool) },
}
var zoneIDSlicePool = sync.Pool{
	New: func() interface{} { return make([]ZoneID, 10) },
}

var playerCheckPool = sync.Pool{
	New: func() interface{} { return make(map[PlayerID]bool) },
}
var playerIDSlicePool = sync.Pool{
	New: func() interface{} { return make([]PlayerID, 10) },
}

var positionCheckPool = sync.Pool{
	New: func() interface{} { return make(map[PositionID]bool) },
}
var positionIDSlicePool = sync.Pool{
	New: func() interface{} { return make([]PositionID, 10) },
}

var itemCheckPool = sync.Pool{
	New: func() interface{} { return make(map[ItemID]bool) },
}
var itemIDSlicePool = sync.Pool{
	New: func() interface{} { return make([]ItemID, 10) },
}

var gearScoreCheckPool = sync.Pool{
	New: func() interface{} { return make(map[GearScoreID]bool) },
}
var gearScoreIDSlicePool = sync.Pool{
	New: func() interface{} { return make([]GearScoreID, 10) },
}

var equipmentSetCheckPool = sync.Pool{
	New: func() interface{} { return make(map[EquipmentSetID]bool) },
}
var equipmentSetIDSlicePool = sync.Pool{
	New: func() interface{} { return make([]EquipmentSetID, 10) },
}

var playerTargetedByRefCheckPool = sync.Pool{
	New: func() interface{} { return make(map[PlayerTargetedByRefID]bool) },
}
var playerTargetedByRefIDSlicePool = sync.Pool{
	New: func() interface{} { return make([]PlayerTargetedByRefID, 10) },
}

var playerTargetRefCheckPool = sync.Pool{
	New: func() interface{} { return make(map[PlayerTargetRefID]bool) },
}
var playerTargetRefIDSlicePool = sync.Pool{
	New: func() interface{} { return make([]PlayerTargetRefID, 10) },
}

var itemBoundToRefCheckPool = sync.Pool{
	New: func() interface{} { return make(map[ItemBoundToRefID]bool) },
}
var itemBoundToRefIDSlicePool = sync.Pool{
	New: func() interface{} { return make([]ItemBoundToRefID, 10) },
}

var playerGuildMemberRefCheckPool = sync.Pool{
	New: func() interface{} { return make(map[PlayerGuildMemberRefID]bool) },
}
var playerGuildMemberRefIDSlicePool = sync.Pool{
	New: func() interface{} { return make([]PlayerGuildMemberRefID, 10) },
}

var playerEquipmentSetRefCheckPool = sync.Pool{
	New: func() interface{} { return make(map[PlayerEquipmentSetRefID]bool) },
}
var playerEquipmentSetRefIDSlicePool = sync.Pool{
	New: func() interface{} { return make([]PlayerEquipmentSetRefID, 10) },
}

var equipmentSetEquipmentRefCheckPool = sync.Pool{
	New: func() interface{} { return make(map[EquipmentSetEquipmentRefID]bool) },
}
var equipmentSetEquipmentRefIDSlicePool = sync.Pool{
	New: func() interface{} { return make([]EquipmentSetEquipmentRefID, 10) },
}
