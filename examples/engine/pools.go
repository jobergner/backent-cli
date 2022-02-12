package state

import (
	"sync"
)

var zoneItemCheckPool = sync.Pool{
	New: func() interface{} { return make(map[ZoneItemID]bool) },
}
var zoneItemIDSlicePool = sync.Pool{
	New: func() interface{} { return make([]ZoneItemID, 0) },
}

var zoneCheckPool = sync.Pool{
	New: func() interface{} { return make(map[ZoneID]bool) },
}
var zoneIDSlicePool = sync.Pool{
	New: func() interface{} { return make([]ZoneID, 0) },
}

var playerCheckPool = sync.Pool{
	New: func() interface{} { return make(map[PlayerID]bool) },
}
var playerIDSlicePool = sync.Pool{
	New: func() interface{} { return make([]PlayerID, 0) },
}

var positionCheckPool = sync.Pool{
	New: func() interface{} { return make(map[PositionID]bool) },
}
var positionIDSlicePool = sync.Pool{
	New: func() interface{} { return make([]PositionID, 0) },
}

var itemCheckPool = sync.Pool{
	New: func() interface{} { return make(map[ItemID]bool) },
}
var itemIDSlicePool = sync.Pool{
	New: func() interface{} { return make([]ItemID, 0) },
}

var gearScoreCheckPool = sync.Pool{
	New: func() interface{} { return make(map[GearScoreID]bool) },
}
var gearScoreIDSlicePool = sync.Pool{
	New: func() interface{} { return make([]GearScoreID, 0) },
}

var attackEventCheckPool = sync.Pool{
	New: func() interface{} { return make(map[AttackEventID]bool) },
}
var attackEventIDSlicePool = sync.Pool{
	New: func() interface{} { return make([]AttackEventID, 0) },
}

var equipmentSetCheckPool = sync.Pool{
	New: func() interface{} { return make(map[EquipmentSetID]bool) },
}
var equipmentSetIDSlicePool = sync.Pool{
	New: func() interface{} { return make([]EquipmentSetID, 0) },
}

var playerTargetedByRefCheckPool = sync.Pool{
	New: func() interface{} { return make(map[PlayerTargetedByRefID]bool) },
}
var playerTargetedByRefIDSlicePool = sync.Pool{
	New: func() interface{} { return make([]PlayerTargetedByRefID, 0) },
}

var playerTargetRefCheckPool = sync.Pool{
	New: func() interface{} { return make(map[PlayerTargetRefID]bool) },
}
var playerTargetRefIDSlicePool = sync.Pool{
	New: func() interface{} { return make([]PlayerTargetRefID, 0) },
}

var attackEventTargetRefCheckPool = sync.Pool{
	New: func() interface{} { return make(map[AttackEventTargetRefID]bool) },
}
var attackEventTargetRefIDSlicePool = sync.Pool{
	New: func() interface{} { return make([]AttackEventTargetRefID, 0) },
}

var itemBoundToRefCheckPool = sync.Pool{
	New: func() interface{} { return make(map[ItemBoundToRefID]bool) },
}
var itemBoundToRefIDSlicePool = sync.Pool{
	New: func() interface{} { return make([]ItemBoundToRefID, 0) },
}

var playerGuildMemberRefCheckPool = sync.Pool{
	New: func() interface{} { return make(map[PlayerGuildMemberRefID]bool) },
}
var playerGuildMemberRefIDSlicePool = sync.Pool{
	New: func() interface{} { return make([]PlayerGuildMemberRefID, 0) },
}

var playerEquipmentSetRefCheckPool = sync.Pool{
	New: func() interface{} { return make(map[PlayerEquipmentSetRefID]bool) },
}
var playerEquipmentSetRefIDSlicePool = sync.Pool{
	New: func() interface{} { return make([]PlayerEquipmentSetRefID, 0) },
}

var equipmentSetEquipmentRefCheckPool = sync.Pool{
	New: func() interface{} { return make(map[EquipmentSetEquipmentRefID]bool) },
}
var equipmentSetEquipmentRefIDSlicePool = sync.Pool{
	New: func() interface{} { return make([]EquipmentSetEquipmentRefID, 0) },
}
