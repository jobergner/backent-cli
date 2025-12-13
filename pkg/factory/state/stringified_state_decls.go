// THIS IS A GENERATED FILE. DO NOT EDIT.
package state

const _AddPlayer_Zone_func string = `func (_zone Zone) AddPlayer() Player {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	if zone.zone.OperationKind == OperationKindDelete {
		return Player{player: playerCore{OperationKind: OperationKindDelete, engine: zone.zone.engine}}
	}
	if _, ok := zone.zone.engine.Patch.Zone[zone.zone.ID]; !ok {
		cp := make([]PlayerID, len(zone.zone.Players))
		copy(cp, zone.zone.Players)
		zone.zone.Players = cp
	}
	player := zone.zone.engine.createPlayer(zone.zone.Path, zone_playersIdentifier)
	zone.zone.Players = append(zone.zone.Players, player.player.ID)
	zone.zone.OperationKind = OperationKindUpdate
	zone.zone.engine.Patch.Zone[zone.zone.ID] = zone.zone
	return player
}`
const _AddItem_Zone_func string = `func (_zone Zone) AddItem() ZoneItem {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	if zone.zone.OperationKind == OperationKindDelete {
		return ZoneItem{zoneItem: zoneItemCore{OperationKind: OperationKindDelete, engine: zone.zone.engine}}
	}
	if _, ok := zone.zone.engine.Patch.Zone[zone.zone.ID]; !ok {
		cp := make([]ZoneItemID, len(zone.zone.Items))
		copy(cp, zone.zone.Items)
		zone.zone.Items = cp
	}
	zoneItem := zone.zone.engine.createZoneItem(zone.zone.Path, zone_itemsIdentifier)
	zone.zone.Items = append(zone.zone.Items, zoneItem.zoneItem.ID)
	zone.zone.OperationKind = OperationKindUpdate
	zone.zone.engine.Patch.Zone[zone.zone.ID] = zone.zone
	return zoneItem
}`
const _AddInteractablePlayer_Zone_func string = `func (_zone Zone) AddInteractablePlayer() Player {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	if zone.zone.OperationKind == OperationKindDelete {
		return Player{player: playerCore{OperationKind: OperationKindDelete, engine: zone.zone.engine}}
	}
	if _, ok := zone.zone.engine.Patch.Zone[zone.zone.ID]; !ok {
		cp := make([]AnyOfItem_Player_ZoneItemID, len(zone.zone.Interactables))
		copy(cp, zone.zone.Interactables)
		zone.zone.Interactables = cp
	}
	player := zone.zone.engine.createPlayer(zone.zone.Path, zone_interactablesIdentifier)
	anyContainer := zone.zone.engine.createAnyOfItem_Player_ZoneItem(int(zone.zone.ID), int(player.player.ID), ElementKindPlayer, zone.zone.Path, zone_interactablesIdentifier).anyOfItem_Player_ZoneItem
	zone.zone.Interactables = append(zone.zone.Interactables, anyContainer.ID)
	zone.zone.OperationKind = OperationKindUpdate
	zone.zone.engine.Patch.Zone[zone.zone.ID] = zone.zone
	return player
}`
const _AddInteractableZoneItem_Zone_func string = `func (_zone Zone) AddInteractableZoneItem() ZoneItem {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	if zone.zone.OperationKind == OperationKindDelete {
		return ZoneItem{zoneItem: zoneItemCore{OperationKind: OperationKindDelete, engine: zone.zone.engine}}
	}
	if _, ok := zone.zone.engine.Patch.Zone[zone.zone.ID]; !ok {
		cp := make([]AnyOfItem_Player_ZoneItemID, len(zone.zone.Interactables))
		copy(cp, zone.zone.Interactables)
		zone.zone.Interactables = cp
	}
	zoneItem := zone.zone.engine.createZoneItem(zone.zone.Path, zone_interactablesIdentifier)
	anyContainer := zone.zone.engine.createAnyOfItem_Player_ZoneItem(int(zone.zone.ID), int(zoneItem.zoneItem.ID), ElementKindZoneItem, zone.zone.Path, zone_interactablesIdentifier).anyOfItem_Player_ZoneItem
	zone.zone.Interactables = append(zone.zone.Interactables, anyContainer.ID)
	zone.zone.OperationKind = OperationKindUpdate
	zone.zone.engine.Patch.Zone[zone.zone.ID] = zone.zone
	return zoneItem
}`
const _AddInteractableItem_Zone_func string = `func (_zone Zone) AddInteractableItem() Item {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	if zone.zone.OperationKind == OperationKindDelete {
		return Item{item: itemCore{OperationKind: OperationKindDelete, engine: zone.zone.engine}}
	}
	if _, ok := zone.zone.engine.Patch.Zone[zone.zone.ID]; !ok {
		cp := make([]AnyOfItem_Player_ZoneItemID, len(zone.zone.Interactables))
		copy(cp, zone.zone.Interactables)
		zone.zone.Interactables = cp
	}
	item := zone.zone.engine.createItem(zone.zone.Path, zone_interactablesIdentifier)
	anyContainer := zone.zone.engine.createAnyOfItem_Player_ZoneItem(int(zone.zone.ID), int(item.item.ID), ElementKindItem, zone.zone.Path, zone_interactablesIdentifier).anyOfItem_Player_ZoneItem
	zone.zone.Interactables = append(zone.zone.Interactables, anyContainer.ID)
	zone.zone.OperationKind = OperationKindUpdate
	zone.zone.engine.Patch.Zone[zone.zone.ID] = zone.zone
	return item
}`
const _AddTag_Zone_func string = `func (_zone Zone) AddTag(tag string) {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	if zone.zone.OperationKind == OperationKindDelete {
		return
	}
	if _, ok := zone.zone.engine.Patch.Zone[zone.zone.ID]; !ok {
		cp := make([]StringValueID, len(zone.zone.Tags))
		copy(cp, zone.zone.Tags)
		zone.zone.Tags = cp
	}
	tagValue := zone.zone.engine.createStringValue(zone.zone.Path, zone_tagsIdentifier, tag)
	zone.zone.Tags = append(zone.zone.Tags, tagValue.ID)
	zone.zone.OperationKind = OperationKindUpdate
	zone.zone.engine.Patch.Zone[zone.zone.ID] = zone.zone
}`
const _AddItem_Player_func string = `func (_player Player) AddItem() Item {
	player := _player.player.engine.Player(_player.player.ID)
	if player.player.OperationKind == OperationKindDelete {
		return Item{item: itemCore{OperationKind: OperationKindDelete, engine: player.player.engine}}
	}
	if _, ok := player.player.engine.Patch.Player[player.player.ID]; !ok {
		cp := make([]ItemID, len(player.player.Items))
		copy(cp, player.player.Items)
		player.player.Items = cp
	}
	item := player.player.engine.createItem(player.player.Path, player_itemsIdentifier)
	player.player.Items = append(player.player.Items, item.item.ID)
	player.player.OperationKind = OperationKindUpdate
	player.player.engine.Patch.Player[player.player.ID] = player.player
	return item
}`
const _AddAction_Player_func string = `func (_player Player) AddAction() AttackEvent {
	player := _player.player.engine.Player(_player.player.ID)
	if player.player.OperationKind == OperationKindDelete {
		return AttackEvent{attackEvent: attackEventCore{OperationKind: OperationKindDelete, engine: player.player.engine}}
	}
	if _, ok := player.player.engine.Patch.Player[player.player.ID]; !ok {
		cp := make([]AttackEventID, len(player.player.Action))
		copy(cp, player.player.Action)
		player.player.Action = cp
	}
	attackEvent := player.player.engine.createAttackEvent(player.player.Path, player_actionIdentifier)
	player.player.Action = append(player.player.Action, attackEvent.attackEvent.ID)
	player.player.OperationKind = OperationKindUpdate
	player.player.engine.Patch.Player[player.player.ID] = player.player
	return attackEvent
}`
const _AddGuildMember_Player_func string = `func (_player Player) AddGuildMember(playerID PlayerID) {
	player := _player.player.engine.Player(_player.player.ID)
	if player.player.OperationKind == OperationKindDelete {
		return
	}
	if player.player.engine.Player(playerID).player.OperationKind == OperationKindDelete {
		return
	}
	if _, ok := player.player.engine.Patch.Player[player.player.ID]; !ok {
		cp := make([]PlayerGuildMemberRefID, len(player.player.GuildMembers))
		copy(cp, player.player.GuildMembers)
		player.player.GuildMembers = cp
	}
	for _, currentRefID := range player.player.GuildMembers {
		currentRef := player.player.engine.playerGuildMemberRef(currentRefID)
		if currentRef.playerGuildMemberRef.ReferencedElementID == playerID {
			return
		}
	}
	ref := player.player.engine.createPlayerGuildMemberRef(player.player.Path, player_guildMembersIdentifier, playerID, player.player.ID, int(playerID))
	player.player.GuildMembers = append(player.player.GuildMembers, ref.ID)
	player.player.OperationKind = OperationKindUpdate
	player.player.engine.Patch.Player[player.player.ID] = player.player
}`
const _AddTargetedByPlayer_Player_func string = `func (_player Player) AddTargetedByPlayer(playerID PlayerID) {
	player := _player.player.engine.Player(_player.player.ID)
	if player.player.OperationKind == OperationKindDelete {
		return
	}
	if player.player.engine.Player(playerID).player.OperationKind == OperationKindDelete {
		return
	}
	if _, ok := player.player.engine.Patch.Player[player.player.ID]; !ok {
		cp := make([]PlayerTargetedByRefID, len(player.player.TargetedBy))
		copy(cp, player.player.TargetedBy)
		player.player.TargetedBy = cp
	}
	for _, currentRefID := range player.player.TargetedBy {
		currentRef := player.player.engine.playerTargetedByRef(currentRefID)
		anyContainer := player.player.engine.anyOfPlayer_ZoneItem(currentRef.playerTargetedByRef.ReferencedElementID)
		if PlayerID(anyContainer.anyOfPlayer_ZoneItem.ChildID) == playerID {
			return
		}
	}
	anyContainer := player.player.engine.createAnyOfPlayer_ZoneItem(int(player.player.ID), int(playerID), ElementKindPlayer, player.player.Path, player_targetedByIdentifier).anyOfPlayer_ZoneItem
	ref := player.player.engine.createPlayerTargetedByRef(player.player.Path, player_targetedByIdentifier, anyContainer.ID, player.player.ID, ElementKindPlayer, int(playerID))
	player.player.TargetedBy = append(player.player.TargetedBy, ref.ID)
	player.player.OperationKind = OperationKindUpdate
	player.player.engine.Patch.Player[player.player.ID] = player.player
}`
const _AddTargetedByZoneItem_Player_func string = `func (_player Player) AddTargetedByZoneItem(zoneItemID ZoneItemID) {
	player := _player.player.engine.Player(_player.player.ID)
	if player.player.OperationKind == OperationKindDelete {
		return
	}
	if player.player.engine.ZoneItem(zoneItemID).zoneItem.OperationKind == OperationKindDelete {
		return
	}
	if _, ok := player.player.engine.Patch.Player[player.player.ID]; !ok {
		cp := make([]PlayerTargetedByRefID, len(player.player.TargetedBy))
		copy(cp, player.player.TargetedBy)
		player.player.TargetedBy = cp
	}
	for _, currentRefID := range player.player.TargetedBy {
		currentRef := player.player.engine.playerTargetedByRef(currentRefID)
		anyContainer := player.player.engine.anyOfPlayer_ZoneItem(currentRef.playerTargetedByRef.ReferencedElementID)
		if ZoneItemID(anyContainer.anyOfPlayer_ZoneItem.ChildID) == zoneItemID {
			return
		}
	}
	anyContainer := player.player.engine.createAnyOfPlayer_ZoneItem(int(player.player.ID), int(zoneItemID), ElementKindZoneItem, player.player.Path, player_targetedByIdentifier).anyOfPlayer_ZoneItem
	ref := player.player.engine.createPlayerTargetedByRef(player.player.Path, player_targetedByIdentifier, anyContainer.ID, player.player.ID, ElementKindZoneItem, int(zoneItemID))
	player.player.TargetedBy = append(player.player.TargetedBy, ref.ID)
	player.player.OperationKind = OperationKindUpdate
	player.player.engine.Patch.Player[player.player.ID] = player.player
}`
const _AddEquipmentSet_Player_func string = `func (_player Player) AddEquipmentSet(equipmentSetID EquipmentSetID) {
	player := _player.player.engine.Player(_player.player.ID)
	if player.player.OperationKind == OperationKindDelete {
		return
	}
	if player.player.engine.EquipmentSet(equipmentSetID).equipmentSet.OperationKind == OperationKindDelete {
		return
	}
	if _, ok := player.player.engine.Patch.Player[player.player.ID]; !ok {
		cp := make([]PlayerEquipmentSetRefID, len(player.player.EquipmentSets))
		copy(cp, player.player.EquipmentSets)
		player.player.EquipmentSets = cp
	}
	for _, currentRefID := range player.player.EquipmentSets {
		currentRef := player.player.engine.playerEquipmentSetRef(currentRefID)
		if currentRef.playerEquipmentSetRef.ReferencedElementID == equipmentSetID {
			return
		}
	}
	ref := player.player.engine.createPlayerEquipmentSetRef(player.player.Path, player_equipmentSetsIdentifier, equipmentSetID, player.player.ID, int(equipmentSetID))
	player.player.EquipmentSets = append(player.player.EquipmentSets, ref.ID)
	player.player.OperationKind = OperationKindUpdate
	player.player.engine.Patch.Player[player.player.ID] = player.player
}`
const _AddEquipment_EquipmentSet_func string = `func (_equipmentSet EquipmentSet) AddEquipment(itemID ItemID) {
	equipmentSet := _equipmentSet.equipmentSet.engine.EquipmentSet(_equipmentSet.equipmentSet.ID)
	if equipmentSet.equipmentSet.OperationKind == OperationKindDelete {
		return
	}
	if equipmentSet.equipmentSet.engine.Item(itemID).item.OperationKind == OperationKindDelete {
		return
	}
	if _, ok := equipmentSet.equipmentSet.engine.Patch.EquipmentSet[equipmentSet.equipmentSet.ID]; !ok {
		cp := make([]EquipmentSetEquipmentRefID, len(equipmentSet.equipmentSet.Equipment))
		copy(cp, equipmentSet.equipmentSet.Equipment)
		equipmentSet.equipmentSet.Equipment = cp
	}
	for _, currentRefID := range equipmentSet.equipmentSet.Equipment {
		currentRef := equipmentSet.equipmentSet.engine.equipmentSetEquipmentRef(currentRefID)
		if currentRef.equipmentSetEquipmentRef.ReferencedElementID == itemID {
			return
		}
	}
	ref := equipmentSet.equipmentSet.engine.createEquipmentSetEquipmentRef(equipmentSet.equipmentSet.Path, equipmentSet_equipmentIdentifier, itemID, equipmentSet.equipmentSet.ID, int(itemID))
	equipmentSet.equipmentSet.Equipment = append(equipmentSet.equipmentSet.Equipment, ref.ID)
	equipmentSet.equipmentSet.OperationKind = OperationKindUpdate
	equipmentSet.equipmentSet.engine.Patch.EquipmentSet[equipmentSet.equipmentSet.ID] = equipmentSet.equipmentSet
}`
const _AnyOfPlayer_PositionRef_type string = `type AnyOfPlayer_PositionRef interface {
	Kind() ElementKind
	Player() Player
	Position() Position
}`
const _AnyOfPlayer_ZoneItemRef_type string = `type AnyOfPlayer_ZoneItemRef interface {
	Kind() ElementKind
	Player() Player
	ZoneItem() ZoneItem
}`
const _AnyOfItem_Player_ZoneItemRef_type string = `type AnyOfItem_Player_ZoneItemRef interface {
	Kind() ElementKind
	Item() Item
	Player() Player
	ZoneItem() ZoneItem
}`
const _AnyOfPlayer_PositionSliceElement_type string = `type AnyOfPlayer_PositionSliceElement interface {
	Kind() ElementKind
	Player() Player
	Position() Position
}`
const _AnyOfPlayer_ZoneItemSliceElement_type string = `type AnyOfPlayer_ZoneItemSliceElement interface {
	Kind() ElementKind
	Player() Player
	ZoneItem() ZoneItem
}`
const _AnyOfItem_Player_ZoneItemSliceElement_type string = `type AnyOfItem_Player_ZoneItemSliceElement interface {
	Kind() ElementKind
	Item() Item
	Player() Player
	ZoneItem() ZoneItem
}`
const _Kind_AnyOfPlayer_ZoneItem_func string = `func (_any AnyOfPlayer_ZoneItem) Kind() ElementKind {
	any := _any.anyOfPlayer_ZoneItem.engine.anyOfPlayer_ZoneItem(_any.anyOfPlayer_ZoneItem.ID)
	return any.anyOfPlayer_ZoneItem.ElementKind
}`
const _BeZoneItem_AnyOfPlayer_ZoneItem_func string = `func (_any AnyOfPlayer_ZoneItem) BeZoneItem() ZoneItem {
	any := _any.anyOfPlayer_ZoneItem.engine.anyOfPlayer_ZoneItem(_any.anyOfPlayer_ZoneItem.ID)
	if any.anyOfPlayer_ZoneItem.ElementKind == ElementKindZoneItem || any.anyOfPlayer_ZoneItem.OperationKind == OperationKindDelete {
		return any.ZoneItem()
	}
	zoneItem := any.anyOfPlayer_ZoneItem.engine.createZoneItem(any.anyOfPlayer_ZoneItem.ParentElementPath, any.anyOfPlayer_ZoneItem.FieldIdentifier)
	any.anyOfPlayer_ZoneItem.beZoneItem(zoneItem.ID(), true)
	return zoneItem
}`
const beZoneItem_anyOfPlayer_ZoneItemCore_func string = `func (_any anyOfPlayer_ZoneItemCore) beZoneItem(zoneItemID ZoneItemID, deleteCurrentChild bool) {
	any := _any.engine.anyOfPlayer_ZoneItem(_any.ID).anyOfPlayer_ZoneItem
	any.engine.deleteAnyOfPlayer_ZoneItem(any.ID, deleteCurrentChild)
	any = any.engine.createAnyOfPlayer_ZoneItem(any.ParentID, int(zoneItemID), ElementKindZoneItem, any.ParentElementPath, any.FieldIdentifier).anyOfPlayer_ZoneItem
	switch any.FieldIdentifier {
	}
	any.engine.Patch.AnyOfPlayer_ZoneItem[any.ID] = any
}`
const _BePlayer_AnyOfPlayer_ZoneItem_func string = `func (_any AnyOfPlayer_ZoneItem) BePlayer() Player {
	any := _any.anyOfPlayer_ZoneItem.engine.anyOfPlayer_ZoneItem(_any.anyOfPlayer_ZoneItem.ID)
	if any.anyOfPlayer_ZoneItem.ElementKind == ElementKindPlayer || any.anyOfPlayer_ZoneItem.OperationKind == OperationKindDelete {
		return any.Player()
	}
	player := any.anyOfPlayer_ZoneItem.engine.createPlayer(any.anyOfPlayer_ZoneItem.ParentElementPath, any.anyOfPlayer_ZoneItem.FieldIdentifier)
	any.anyOfPlayer_ZoneItem.bePlayer(player.ID(), true)
	return player
}`
const bePlayer_anyOfPlayer_ZoneItemCore_func string = `func (_any anyOfPlayer_ZoneItemCore) bePlayer(playerID PlayerID, deleteCurrentChild bool) {
	any := _any.engine.anyOfPlayer_ZoneItem(_any.ID).anyOfPlayer_ZoneItem
	any.engine.deleteAnyOfPlayer_ZoneItem(any.ID, deleteCurrentChild)
	any = any.engine.createAnyOfPlayer_ZoneItem(any.ParentID, int(playerID), ElementKindPlayer, any.ParentElementPath, any.FieldIdentifier).anyOfPlayer_ZoneItem
	switch any.FieldIdentifier {
	}
	any.engine.Patch.AnyOfPlayer_ZoneItem[any.ID] = any
}`
const deleteChild_anyOfPlayer_ZoneItemCore_func string = `func (_any anyOfPlayer_ZoneItemCore) deleteChild() {
	any := _any.engine.anyOfPlayer_ZoneItem(_any.ID).anyOfPlayer_ZoneItem
	switch any.ElementKind {
	case ElementKindPlayer:
		any.engine.deletePlayer(PlayerID(any.ChildID))
	case ElementKindZoneItem:
		any.engine.deleteZoneItem(ZoneItemID(any.ChildID))
	}
}`
const _Kind_AnyOfPlayer_Position_func string = `func (_any AnyOfPlayer_Position) Kind() ElementKind {
	any := _any.anyOfPlayer_Position.engine.anyOfPlayer_Position(_any.anyOfPlayer_Position.ID)
	return any.anyOfPlayer_Position.ElementKind
}`
const _BePosition_AnyOfPlayer_Position_func string = `func (_any AnyOfPlayer_Position) BePosition() Position {
	any := _any.anyOfPlayer_Position.engine.anyOfPlayer_Position(_any.anyOfPlayer_Position.ID)
	if any.anyOfPlayer_Position.ElementKind == ElementKindPosition || any.anyOfPlayer_Position.OperationKind == OperationKindDelete {
		return any.Position()
	}
	position := any.anyOfPlayer_Position.engine.createPosition(any.anyOfPlayer_Position.ParentElementPath, any.anyOfPlayer_Position.FieldIdentifier)
	any.anyOfPlayer_Position.bePosition(position.ID(), true)
	return position
}`
const bePosition_anyOfPlayer_PositionCore_func string = `func (_any anyOfPlayer_PositionCore) bePosition(positionID PositionID, deleteCurrentChild bool) {
	any := _any.engine.anyOfPlayer_Position(_any.ID).anyOfPlayer_Position
	any.engine.deleteAnyOfPlayer_Position(any.ID, deleteCurrentChild)
	any = any.engine.createAnyOfPlayer_Position(any.ParentID, int(positionID), ElementKindPosition, any.ParentElementPath, any.FieldIdentifier).anyOfPlayer_Position
	switch any.FieldIdentifier {
	case item_originIdentifier:
		item := any.engine.Item(ItemID(any.ParentID)).item
		item.Origin = any.ID
		item.engine.Patch.Item[item.ID] = item
	}
	any.engine.Patch.AnyOfPlayer_Position[any.ID] = any
}`
const deleteChild_anyOfPlayer_PositionCore_func string = `func (_any anyOfPlayer_PositionCore) deleteChild() {
	any := _any.engine.anyOfPlayer_Position(_any.ID).anyOfPlayer_Position
	switch any.ElementKind {
	case ElementKindPlayer:
		any.engine.deletePlayer(PlayerID(any.ChildID))
	case ElementKindPosition:
		any.engine.deletePosition(PositionID(any.ChildID))
	}
}`
const _BePlayer_AnyOfPlayer_Position_func string = `func (_any AnyOfPlayer_Position) BePlayer() Player {
	any := _any.anyOfPlayer_Position.engine.anyOfPlayer_Position(_any.anyOfPlayer_Position.ID)
	if any.anyOfPlayer_Position.ElementKind == ElementKindPlayer || any.anyOfPlayer_Position.OperationKind == OperationKindDelete {
		return any.Player()
	}
	player := any.anyOfPlayer_Position.engine.createPlayer(any.anyOfPlayer_Position.ParentElementPath, any.anyOfPlayer_Position.FieldIdentifier)
	any.anyOfPlayer_Position.bePlayer(player.ID(), true)
	return player
}`
const bePlayer_anyOfPlayer_PositionCore_func string = `func (_any anyOfPlayer_PositionCore) bePlayer(playerID PlayerID, deleteCurrentChild bool) {
	any := _any.engine.anyOfPlayer_Position(_any.ID).anyOfPlayer_Position
	any.engine.deleteAnyOfPlayer_Position(any.ID, deleteCurrentChild)
	any = any.engine.createAnyOfPlayer_Position(any.ParentID, int(playerID), ElementKindPlayer, any.ParentElementPath, any.FieldIdentifier).anyOfPlayer_Position
	switch any.FieldIdentifier {
	case item_originIdentifier:
		item := any.engine.Item(ItemID(any.ParentID)).item
		item.Origin = any.ID
		item.engine.Patch.Item[item.ID] = item
	}
	any.engine.Patch.AnyOfPlayer_Position[any.ID] = any
}`
const _Kind_AnyOfItem_Player_ZoneItem_func string = `func (_any AnyOfItem_Player_ZoneItem) Kind() ElementKind {
	any := _any.anyOfItem_Player_ZoneItem.engine.anyOfItem_Player_ZoneItem(_any.anyOfItem_Player_ZoneItem.ID)
	return any.anyOfItem_Player_ZoneItem.ElementKind
}`
const _BeZoneItem_AnyOfItem_Player_ZoneItem_func string = `func (_any AnyOfItem_Player_ZoneItem) BeZoneItem() ZoneItem {
	any := _any.anyOfItem_Player_ZoneItem.engine.anyOfItem_Player_ZoneItem(_any.anyOfItem_Player_ZoneItem.ID)
	if any.anyOfItem_Player_ZoneItem.ElementKind == ElementKindZoneItem || any.anyOfItem_Player_ZoneItem.OperationKind == OperationKindDelete {
		return any.ZoneItem()
	}
	zoneItem := any.anyOfItem_Player_ZoneItem.engine.createZoneItem(any.anyOfItem_Player_ZoneItem.ParentElementPath, any.anyOfItem_Player_ZoneItem.FieldIdentifier)
	any.anyOfItem_Player_ZoneItem.beZoneItem(zoneItem.ID(), true)
	return zoneItem
}`
const beZoneItem_anyOfItem_Player_ZoneItemCore_func string = `func (_any anyOfItem_Player_ZoneItemCore) beZoneItem(zoneItemID ZoneItemID, deleteCurrentChild bool) {
	any := _any.engine.anyOfItem_Player_ZoneItem(_any.ID).anyOfItem_Player_ZoneItem
	any.engine.deleteAnyOfItem_Player_ZoneItem(any.ID, deleteCurrentChild)
	any = any.engine.createAnyOfItem_Player_ZoneItem(any.ParentID, int(zoneItemID), ElementKindZoneItem, any.ParentElementPath, any.FieldIdentifier).anyOfItem_Player_ZoneItem
	switch any.FieldIdentifier {
	}
	any.engine.Patch.AnyOfItem_Player_ZoneItem[any.ID] = any
}`
const _BePlayer_AnyOfItem_Player_ZoneItem_func string = `func (_any AnyOfItem_Player_ZoneItem) BePlayer() Player {
	any := _any.anyOfItem_Player_ZoneItem.engine.anyOfItem_Player_ZoneItem(_any.anyOfItem_Player_ZoneItem.ID)
	if any.anyOfItem_Player_ZoneItem.ElementKind == ElementKindPlayer || any.anyOfItem_Player_ZoneItem.OperationKind == OperationKindDelete {
		return any.Player()
	}
	player := any.anyOfItem_Player_ZoneItem.engine.createPlayer(any.anyOfItem_Player_ZoneItem.ParentElementPath, any.anyOfItem_Player_ZoneItem.FieldIdentifier)
	any.anyOfItem_Player_ZoneItem.bePlayer(player.ID(), true)
	return player
}`
const bePlayer_anyOfItem_Player_ZoneItemCore_func string = `func (_any anyOfItem_Player_ZoneItemCore) bePlayer(playerID PlayerID, deleteCurrentChild bool) {
	any := _any.engine.anyOfItem_Player_ZoneItem(_any.ID).anyOfItem_Player_ZoneItem
	any.engine.deleteAnyOfItem_Player_ZoneItem(any.ID, deleteCurrentChild)
	any = any.engine.createAnyOfItem_Player_ZoneItem(any.ParentID, int(playerID), ElementKindPlayer, any.ParentElementPath, any.FieldIdentifier).anyOfItem_Player_ZoneItem
	switch any.FieldIdentifier {
	}
	any.engine.Patch.AnyOfItem_Player_ZoneItem[any.ID] = any
}`
const _BeItem_AnyOfItem_Player_ZoneItem_func string = `func (_any AnyOfItem_Player_ZoneItem) BeItem() Item {
	any := _any.anyOfItem_Player_ZoneItem.engine.anyOfItem_Player_ZoneItem(_any.anyOfItem_Player_ZoneItem.ID)
	if any.anyOfItem_Player_ZoneItem.ElementKind == ElementKindItem || any.anyOfItem_Player_ZoneItem.OperationKind == OperationKindDelete {
		return any.Item()
	}
	item := any.anyOfItem_Player_ZoneItem.engine.createItem(any.anyOfItem_Player_ZoneItem.ParentElementPath, any.anyOfItem_Player_ZoneItem.FieldIdentifier)
	any.anyOfItem_Player_ZoneItem.beItem(item.ID(), true)
	return item
}`
const beItem_anyOfItem_Player_ZoneItemCore_func string = `func (_any anyOfItem_Player_ZoneItemCore) beItem(itemID ItemID, deleteCurrentChild bool) {
	any := _any.engine.anyOfItem_Player_ZoneItem(_any.ID).anyOfItem_Player_ZoneItem
	any.engine.deleteAnyOfItem_Player_ZoneItem(any.ID, deleteCurrentChild)
	any = any.engine.createAnyOfItem_Player_ZoneItem(any.ParentID, int(itemID), ElementKindItem, any.ParentElementPath, any.FieldIdentifier).anyOfItem_Player_ZoneItem
	switch any.FieldIdentifier {
	}
	any.engine.Patch.AnyOfItem_Player_ZoneItem[any.ID] = any
}`
const deleteChild_anyOfItem_Player_ZoneItemCore_func string = `func (_any anyOfItem_Player_ZoneItemCore) deleteChild() {
	any := _any.engine.anyOfItem_Player_ZoneItem(_any.ID).anyOfItem_Player_ZoneItem
	switch any.ElementKind {
	case ElementKindItem:
		any.engine.deleteItem(ItemID(any.ChildID))
	case ElementKindPlayer:
		any.engine.deletePlayer(PlayerID(any.ChildID))
	case ElementKindZoneItem:
		any.engine.deleteZoneItem(ZoneItemID(any.ChildID))
	}
}`
const assembleGearScorePath_Engine_func string = `func (engine *Engine) assembleGearScorePath(element *gearScore, p path, pIndex int, includedElements map[int]bool) {
	gearScoreData, ok := engine.Patch.GearScore[element.ID]
	if !ok {
		gearScoreData = engine.State.GearScore[element.ID]
	}
	if element.OperationKind != OperationKindUpdate && element.OperationKind != OperationKindDelete {
		element.OperationKind = gearScoreData.OperationKind
	}
	if pIndex+1 == len(p) {
		return
	}
	nextSeg := p[pIndex+1]
	switch nextSeg.Identifier {
	case gearScore_levelIdentifier:
		child := engine.intValue(gearScoreData.Level)
		element.OperationKind = child.OperationKind
		element.Level = &child.Value
	case gearScore_scoreIdentifier:
		child := engine.intValue(gearScoreData.Score)
		element.OperationKind = child.OperationKind
		element.Score = &child.Value
	}
	_ = gearScoreData
}`
const assemblePositionPath_Engine_func string = `func (engine *Engine) assemblePositionPath(element *position, p path, pIndex int, includedElements map[int]bool) {
	positionData, ok := engine.Patch.Position[element.ID]
	if !ok {
		positionData = engine.State.Position[element.ID]
	}
	if element.OperationKind != OperationKindUpdate && element.OperationKind != OperationKindDelete {
		element.OperationKind = positionData.OperationKind
	}
	if pIndex+1 == len(p) {
		return
	}
	nextSeg := p[pIndex+1]
	switch nextSeg.Identifier {
	case position_xIdentifier:
		child := engine.floatValue(positionData.X)
		element.OperationKind = child.OperationKind
		element.X = &child.Value
	case position_yIdentifier:
		child := engine.floatValue(positionData.Y)
		element.OperationKind = child.OperationKind
		element.Y = &child.Value
	}
	_ = positionData
}`
const assembleEquipmentSetPath_Engine_func string = `func (engine *Engine) assembleEquipmentSetPath(element *equipmentSet, p path, pIndex int, includedElements map[int]bool) {
	equipmentSetData, ok := engine.Patch.EquipmentSet[element.ID]
	if !ok {
		equipmentSetData = engine.State.EquipmentSet[element.ID]
	}
	if element.OperationKind != OperationKindUpdate && element.OperationKind != OperationKindDelete {
		element.OperationKind = equipmentSetData.OperationKind
	}
	if pIndex+1 == len(p) {
		return
	}
	nextSeg := p[pIndex+1]
	switch nextSeg.Identifier {
	case equipmentSet_equipmentIdentifier:
		ref := engine.equipmentSetEquipmentRef(EquipmentSetEquipmentRefID(nextSeg.RefID)).equipmentSetEquipmentRef
		referencedDataStatus := ReferencedDataUnchanged
		if _, ok := includedElements[int(ref.ReferencedElementID)]; ok {
			referencedDataStatus = ReferencedDataModified
		}
		referencedElement := engine.Item(ref.ReferencedElementID).item
		treeRef := elementReference{ID: int(ref.ID), OperationKind: ref.OperationKind, ElementID: int(ref.ReferencedElementID), ElementKind: ElementKindItem, ReferencedDataStatus: referencedDataStatus, ElementPath: referencedElement.JSONPath}
		if element.Equipment == nil {
			element.Equipment = make(map[ItemID]elementReference)
		}
		element.Equipment[referencedElement.ID] = treeRef
	case equipmentSet_nameIdentifier:
		child := engine.stringValue(equipmentSetData.Name)
		element.OperationKind = child.OperationKind
		element.Name = &child.Value
	}
	_ = equipmentSetData
}`
const assembleAttackEventPath_Engine_func string = `func (engine *Engine) assembleAttackEventPath(element *attackEvent, p path, pIndex int, includedElements map[int]bool) {
	attackEventData, ok := engine.Patch.AttackEvent[element.ID]
	if !ok {
		attackEventData = engine.State.AttackEvent[element.ID]
	}
	if element.OperationKind != OperationKindUpdate && element.OperationKind != OperationKindDelete {
		element.OperationKind = attackEventData.OperationKind
	}
	if pIndex+1 == len(p) {
		return
	}
	nextSeg := p[pIndex+1]
	switch nextSeg.Identifier {
	case attackEvent_targetIdentifier:
		ref := engine.attackEventTargetRef(AttackEventTargetRefID(nextSeg.RefID)).attackEventTargetRef
		if element.Target != nil && ref.OperationKind == OperationKindDelete {
			break
		}
		referencedDataStatus := ReferencedDataUnchanged
		if _, ok := includedElements[int(ref.ReferencedElementID)]; ok {
			referencedDataStatus = ReferencedDataModified
		}
		referencedElement := engine.Player(ref.ReferencedElementID).player
		treeRef := elementReference{ID: int(ref.ID), OperationKind: ref.OperationKind, ElementID: int(ref.ReferencedElementID), ElementKind: ElementKindPlayer, ReferencedDataStatus: referencedDataStatus, ElementPath: referencedElement.JSONPath}
		element.Target = &treeRef
	}
	_ = attackEventData
}`
const assembleItemPath_Engine_func string = `func (engine *Engine) assembleItemPath(element *item, p path, pIndex int, includedElements map[int]bool) {
	itemData, ok := engine.Patch.Item[element.ID]
	if !ok {
		itemData = engine.State.Item[element.ID]
	}
	if element.OperationKind != OperationKindUpdate && element.OperationKind != OperationKindDelete {
		element.OperationKind = itemData.OperationKind
	}
	if pIndex+1 == len(p) {
		return
	}
	nextSeg := p[pIndex+1]
	switch nextSeg.Identifier {
	case item_boundToIdentifier:
		ref := engine.itemBoundToRef(ItemBoundToRefID(nextSeg.RefID)).itemBoundToRef
		if element.BoundTo != nil && ref.OperationKind == OperationKindDelete {
			break
		}
		referencedDataStatus := ReferencedDataUnchanged
		if _, ok := includedElements[int(ref.ReferencedElementID)]; ok {
			referencedDataStatus = ReferencedDataModified
		}
		referencedElement := engine.Player(ref.ReferencedElementID).player
		treeRef := elementReference{ID: int(ref.ID), OperationKind: ref.OperationKind, ElementID: int(ref.ReferencedElementID), ElementKind: ElementKindPlayer, ReferencedDataStatus: referencedDataStatus, ElementPath: referencedElement.JSONPath}
		element.BoundTo = &treeRef
	case item_gearScoreIdentifier:
		child := element.GearScore
		if child == nil {
			child = &gearScore{ID: GearScoreID(nextSeg.ID)}
		}
		engine.assembleGearScorePath(child, p, pIndex+1, includedElements)
		element.GearScore = child
	case item_nameIdentifier:
		child := engine.stringValue(itemData.Name)
		element.OperationKind = child.OperationKind
		element.Name = &child.Value
	case item_originIdentifier:
		switch nextSeg.Kind {
		case ElementKindPlayer:
			child, ok := element.Origin.(*player)
			if !ok || child == nil {
				child = &player{ID: PlayerID(nextSeg.ID)}
			}
			engine.assemblePlayerPath(child, p, pIndex+1, includedElements)
			if child.OperationKind == OperationKindDelete && element.Origin != nil {
				break
			}
			element.Origin = child
		case ElementKindPosition:
			child, ok := element.Origin.(*position)
			if !ok || child == nil {
				child = &position{ID: PositionID(nextSeg.ID)}
			}
			engine.assemblePositionPath(child, p, pIndex+1, includedElements)
			if child.OperationKind == OperationKindDelete && element.Origin != nil {
				break
			}
			element.Origin = child
		}
	}
	_ = itemData
}`
const assembleZoneItemPath_Engine_func string = `func (engine *Engine) assembleZoneItemPath(element *zoneItem, p path, pIndex int, includedElements map[int]bool) {
	zoneItemData, ok := engine.Patch.ZoneItem[element.ID]
	if !ok {
		zoneItemData = engine.State.ZoneItem[element.ID]
	}
	if element.OperationKind != OperationKindUpdate && element.OperationKind != OperationKindDelete {
		element.OperationKind = zoneItemData.OperationKind
	}
	if pIndex+1 == len(p) {
		return
	}
	nextSeg := p[pIndex+1]
	switch nextSeg.Identifier {
	case zoneItem_itemIdentifier:
		child := element.Item
		if child == nil {
			child = &item{ID: ItemID(nextSeg.ID)}
		}
		engine.assembleItemPath(child, p, pIndex+1, includedElements)
		element.Item = child
	case zoneItem_positionIdentifier:
		child := element.Position
		if child == nil {
			child = &position{ID: PositionID(nextSeg.ID)}
		}
		engine.assemblePositionPath(child, p, pIndex+1, includedElements)
		element.Position = child
	}
	_ = zoneItemData
}`
const assemblePlayerPath_Engine_func string = `func (engine *Engine) assemblePlayerPath(element *player, p path, pIndex int, includedElements map[int]bool) {
	playerData, ok := engine.Patch.Player[element.ID]
	if !ok {
		playerData = engine.State.Player[element.ID]
	}
	if element.OperationKind != OperationKindUpdate && element.OperationKind != OperationKindDelete {
		element.OperationKind = playerData.OperationKind
	}
	if pIndex+1 == len(p) {
		return
	}
	nextSeg := p[pIndex+1]
	switch nextSeg.Identifier {
	case player_actionIdentifier:
		if element.Action == nil {
			element.Action = make(map[AttackEventID]attackEvent)
		}
		child, ok := element.Action[AttackEventID(nextSeg.ID)]
		if !ok {
			child = attackEvent{ID: AttackEventID(nextSeg.ID)}
		}
		engine.assembleAttackEventPath(&child, p, pIndex+1, includedElements)
		element.Action[child.ID] = child
	case player_equipmentSetsIdentifier:
		ref := engine.playerEquipmentSetRef(PlayerEquipmentSetRefID(nextSeg.RefID)).playerEquipmentSetRef
		referencedDataStatus := ReferencedDataUnchanged
		if _, ok := includedElements[int(ref.ReferencedElementID)]; ok {
			referencedDataStatus = ReferencedDataModified
		}
		referencedElement := engine.EquipmentSet(ref.ReferencedElementID).equipmentSet
		treeRef := elementReference{ID: int(ref.ID), OperationKind: ref.OperationKind, ElementID: int(ref.ReferencedElementID), ElementKind: ElementKindEquipmentSet, ReferencedDataStatus: referencedDataStatus, ElementPath: referencedElement.JSONPath}
		if element.EquipmentSets == nil {
			element.EquipmentSets = make(map[EquipmentSetID]elementReference)
		}
		element.EquipmentSets[referencedElement.ID] = treeRef
	case player_gearScoreIdentifier:
		child := element.GearScore
		if child == nil {
			child = &gearScore{ID: GearScoreID(nextSeg.ID)}
		}
		engine.assembleGearScorePath(child, p, pIndex+1, includedElements)
		element.GearScore = child
	case player_guildMembersIdentifier:
		ref := engine.playerGuildMemberRef(PlayerGuildMemberRefID(nextSeg.RefID)).playerGuildMemberRef
		referencedDataStatus := ReferencedDataUnchanged
		if _, ok := includedElements[int(ref.ReferencedElementID)]; ok {
			referencedDataStatus = ReferencedDataModified
		}
		referencedElement := engine.Player(ref.ReferencedElementID).player
		treeRef := elementReference{ID: int(ref.ID), OperationKind: ref.OperationKind, ElementID: int(ref.ReferencedElementID), ElementKind: ElementKindPlayer, ReferencedDataStatus: referencedDataStatus, ElementPath: referencedElement.JSONPath}
		if element.GuildMembers == nil {
			element.GuildMembers = make(map[PlayerID]elementReference)
		}
		element.GuildMembers[referencedElement.ID] = treeRef
	case player_itemsIdentifier:
		if element.Items == nil {
			element.Items = make(map[ItemID]item)
		}
		child, ok := element.Items[ItemID(nextSeg.ID)]
		if !ok {
			child = item{ID: ItemID(nextSeg.ID)}
		}
		engine.assembleItemPath(&child, p, pIndex+1, includedElements)
		element.Items[child.ID] = child
	case player_positionIdentifier:
		child := element.Position
		if child == nil {
			child = &position{ID: PositionID(nextSeg.ID)}
		}
		engine.assemblePositionPath(child, p, pIndex+1, includedElements)
		element.Position = child
	case player_targetIdentifier:
		ref := engine.playerTargetRef(PlayerTargetRefID(nextSeg.RefID)).playerTargetRef
		if element.Target != nil && ref.OperationKind == OperationKindDelete {
			break
		}
		referencedDataStatus := ReferencedDataUnchanged
		if _, ok := includedElements[ref.ChildID]; ok {
			referencedDataStatus = ReferencedDataModified
		}
		switch nextSeg.Kind {
		case ElementKindPlayer:
			referencedElement := engine.Player(PlayerID(ref.ChildID)).player
			treeRef := elementReference{ID: int(ref.ID), OperationKind: ref.OperationKind, ElementID: ref.ChildID, ElementKind: ElementKindPlayer, ReferencedDataStatus: referencedDataStatus, ElementPath: referencedElement.JSONPath}
			element.Target = &treeRef
		case ElementKindZoneItem:
			referencedElement := engine.ZoneItem(ZoneItemID(ref.ChildID)).zoneItem
			treeRef := elementReference{ID: int(ref.ID), OperationKind: ref.OperationKind, ElementID: ref.ChildID, ElementKind: ElementKindZoneItem, ReferencedDataStatus: referencedDataStatus, ElementPath: referencedElement.JSONPath}
			element.Target = &treeRef
		}
	case player_targetedByIdentifier:
		if element.TargetedBy == nil {
			element.TargetedBy = make(map[int]elementReference)
		}
		ref := engine.playerTargetedByRef(PlayerTargetedByRefID(nextSeg.RefID)).playerTargetedByRef
		referencedDataStatus := ReferencedDataUnchanged
		if _, ok := includedElements[ref.ChildID]; ok {
			referencedDataStatus = ReferencedDataModified
		}
		switch nextSeg.Kind {
		case ElementKindPlayer:
			referencedElement := engine.Player(PlayerID(ref.ChildID)).player
			treeRef := elementReference{ID: int(ref.ID), OperationKind: ref.OperationKind, ElementID: ref.ChildID, ElementKind: ElementKindPlayer, ReferencedDataStatus: referencedDataStatus, ElementPath: referencedElement.JSONPath}
			element.TargetedBy[ref.ChildID] = treeRef
		case ElementKindZoneItem:
			referencedElement := engine.ZoneItem(ZoneItemID(ref.ChildID)).zoneItem
			treeRef := elementReference{ID: int(ref.ID), OperationKind: ref.OperationKind, ElementID: ref.ChildID, ElementKind: ElementKindZoneItem, ReferencedDataStatus: referencedDataStatus, ElementPath: referencedElement.JSONPath}
			element.TargetedBy[ref.ChildID] = treeRef
		}
	}
	_ = playerData
}`
const assembleZonePath_Engine_func string = `func (engine *Engine) assembleZonePath(element *zone, p path, pIndex int, includedElements map[int]bool) {
	zoneData, ok := engine.Patch.Zone[element.ID]
	if !ok {
		zoneData = engine.State.Zone[element.ID]
	}
	if element.OperationKind != OperationKindUpdate && element.OperationKind != OperationKindDelete {
		element.OperationKind = zoneData.OperationKind
	}
	if pIndex+1 == len(p) {
		return
	}
	nextSeg := p[pIndex+1]
	switch nextSeg.Identifier {
	case zone_interactablesIdentifier:
		if element.Interactables == nil {
			element.Interactables = make(map[int]interface{})
		}
		switch nextSeg.Kind {
		case ElementKindItem:
			child, ok := element.Interactables[nextSeg.ID].(item)
			if !ok {
				child = item{ID: ItemID(nextSeg.ID)}
			}
			engine.assembleItemPath(&child, p, pIndex+1, includedElements)
			element.Interactables[nextSeg.ID] = child
		case ElementKindPlayer:
			child, ok := element.Interactables[nextSeg.ID].(player)
			if !ok {
				child = player{ID: PlayerID(nextSeg.ID)}
			}
			engine.assemblePlayerPath(&child, p, pIndex+1, includedElements)
			element.Interactables[nextSeg.ID] = child
		case ElementKindZoneItem:
			child, ok := element.Interactables[nextSeg.ID].(zoneItem)
			if !ok {
				child = zoneItem{ID: ZoneItemID(nextSeg.ID)}
			}
			engine.assembleZoneItemPath(&child, p, pIndex+1, includedElements)
			element.Interactables[nextSeg.ID] = child
		}
	case zone_itemsIdentifier:
		if element.Items == nil {
			element.Items = make(map[ZoneItemID]zoneItem)
		}
		child, ok := element.Items[ZoneItemID(nextSeg.ID)]
		if !ok {
			child = zoneItem{ID: ZoneItemID(nextSeg.ID)}
		}
		engine.assembleZoneItemPath(&child, p, pIndex+1, includedElements)
		element.Items[child.ID] = child
	case zone_playersIdentifier:
		if element.Players == nil {
			element.Players = make(map[PlayerID]player)
		}
		child, ok := element.Players[PlayerID(nextSeg.ID)]
		if !ok {
			child = player{ID: PlayerID(nextSeg.ID)}
		}
		engine.assemblePlayerPath(&child, p, pIndex+1, includedElements)
		element.Players[child.ID] = child
	case zone_tagsIdentifier:
		if element.Tags == nil {
			if stateZone, ok := engine.State.Zone[element.ID]; ok {
				element.Tags = make([]string, 0, len(stateZone.Tags))
				for _, valID := range stateZone.Tags {
					val := engine.stringValue(valID)
					element.Tags = append(element.Tags, val.Value)
				}
			} else {
				element.Tags = make([]string, 0, len(zoneData.Tags))
			}
		}
		child := engine.stringValue(StringValueID(nextSeg.ID))
		switch child.OperationKind {
		case OperationKindUnchanged:
			break
		case OperationKindDelete:
			var newValues []string
			for _, val := range element.Tags {
				if val != child.Value {
					newValues = append(newValues, val)
				}
			}
			element.Tags = newValues
		case OperationKindUpdate:
			element.Tags = append(element.Tags, child.Value)
		}
	}
	_ = zoneData
}`
const assemblePlanner_type string = `type assemblePlanner struct {
	updatedPaths		[]path
	updatedReferencePaths	map[int]path
	updatedElementPaths	map[int]path
	includedElements	map[int]bool
}`
const newAssemblePlanner_func string = `func newAssemblePlanner() *assemblePlanner {
	return &assemblePlanner{includedElements: make(map[int]bool), updatedElementPaths: make(map[int]path), updatedPaths: make([]path, 0), updatedReferencePaths: make(map[int]path)}
}`
const clear_assemblePlanner_func string = `func (a *assemblePlanner) clear() {
	a.updatedPaths = a.updatedPaths[:0]
	for key := range a.updatedElementPaths {
		delete(a.updatedElementPaths, key)
	}
	for key := range a.updatedReferencePaths {
		delete(a.updatedReferencePaths, key)
	}
	for key := range a.includedElements {
		delete(a.includedElements, key)
	}
}`
const plan_assemblePlanner_func string = `func (ap *assemblePlanner) plan(state, patch *State) {
	for _, boolValue := range patch.BoolValue {
		ap.updatedElementPaths[int(boolValue.ID)] = boolValue.Path
	}
	for _, floatValue := range patch.FloatValue {
		ap.updatedElementPaths[int(floatValue.ID)] = floatValue.Path
	}
	for _, intValue := range patch.IntValue {
		ap.updatedElementPaths[int(intValue.ID)] = intValue.Path
	}
	for _, stringValue := range patch.StringValue {
		ap.updatedElementPaths[int(stringValue.ID)] = stringValue.Path
	}
	for _, attackEvent := range patch.AttackEvent {
		ap.updatedElementPaths[int(attackEvent.ID)] = attackEvent.Path
	}
	for _, equipmentSet := range patch.EquipmentSet {
		ap.updatedElementPaths[int(equipmentSet.ID)] = equipmentSet.Path
	}
	for _, gearScore := range patch.GearScore {
		ap.updatedElementPaths[int(gearScore.ID)] = gearScore.Path
	}
	for _, item := range patch.Item {
		ap.updatedElementPaths[int(item.ID)] = item.Path
	}
	for _, player := range patch.Player {
		ap.updatedElementPaths[int(player.ID)] = player.Path
	}
	for _, position := range patch.Position {
		ap.updatedElementPaths[int(position.ID)] = position.Path
	}
	for _, zone := range patch.Zone {
		ap.updatedElementPaths[int(zone.ID)] = zone.Path
	}
	for _, zoneItem := range patch.ZoneItem {
		ap.updatedElementPaths[int(zoneItem.ID)] = zoneItem.Path
	}
	for _, attackEventTargetRef := range patch.AttackEventTargetRef {
		ap.updatedReferencePaths[int(attackEventTargetRef.ID)] = attackEventTargetRef.Path
	}
	for _, equipmentSetEquipmentRef := range patch.EquipmentSetEquipmentRef {
		ap.updatedReferencePaths[int(equipmentSetEquipmentRef.ID)] = equipmentSetEquipmentRef.Path
	}
	for _, itemBoundToRef := range patch.ItemBoundToRef {
		ap.updatedReferencePaths[int(itemBoundToRef.ID)] = itemBoundToRef.Path
	}
	for _, playerEquipmentSetRef := range patch.PlayerEquipmentSetRef {
		ap.updatedReferencePaths[int(playerEquipmentSetRef.ID)] = playerEquipmentSetRef.Path
	}
	for _, playerGuildMemberRef := range patch.PlayerGuildMemberRef {
		ap.updatedReferencePaths[int(playerGuildMemberRef.ID)] = playerGuildMemberRef.Path
	}
	for _, playerTargetRef := range patch.PlayerTargetRef {
		ap.updatedReferencePaths[int(playerTargetRef.ID)] = playerTargetRef.Path
	}
	for _, playerTargetedByRef := range patch.PlayerTargetedByRef {
		ap.updatedReferencePaths[int(playerTargetedByRef.ID)] = playerTargetedByRef.Path
	}
	previousLen := 0
	for {
		for _, p := range ap.updatedElementPaths {
			for _, seg := range p {
				ap.includedElements[seg.ID] = true
			}
		}
		for _, p := range ap.updatedReferencePaths {
			for _, seg := range p {
				if seg.RefID != 0 {
				} else {
					ap.includedElements[seg.ID] = true
				}
			}
		}
		if previousLen == len(ap.includedElements) {
			break
		}
		previousLen = len(ap.includedElements)
		for _, attackEventTargetRef := range patch.AttackEventTargetRef {
			if _, ok := ap.includedElements[int(attackEventTargetRef.ChildID)]; ok {
				ap.updatedReferencePaths[int(attackEventTargetRef.ID)] = attackEventTargetRef.Path
			}
		}
		for _, attackEventTargetRef := range state.AttackEventTargetRef {
			if _, ok := ap.updatedReferencePaths[int(attackEventTargetRef.ID)]; ok {
				continue
			}
			if _, ok := ap.includedElements[int(attackEventTargetRef.ChildID)]; ok {
				ap.updatedReferencePaths[int(attackEventTargetRef.ID)] = attackEventTargetRef.Path
			}
		}
		for _, equipmentSetEquipmentRef := range patch.EquipmentSetEquipmentRef {
			if _, ok := ap.includedElements[int(equipmentSetEquipmentRef.ChildID)]; ok {
				ap.updatedReferencePaths[int(equipmentSetEquipmentRef.ID)] = equipmentSetEquipmentRef.Path
			}
		}
		for _, equipmentSetEquipmentRef := range state.EquipmentSetEquipmentRef {
			if _, ok := ap.updatedReferencePaths[int(equipmentSetEquipmentRef.ID)]; ok {
				continue
			}
			if _, ok := ap.includedElements[int(equipmentSetEquipmentRef.ChildID)]; ok {
				ap.updatedReferencePaths[int(equipmentSetEquipmentRef.ID)] = equipmentSetEquipmentRef.Path
			}
		}
		for _, itemBoundToRef := range patch.ItemBoundToRef {
			if _, ok := ap.includedElements[int(itemBoundToRef.ChildID)]; ok {
				ap.updatedReferencePaths[int(itemBoundToRef.ID)] = itemBoundToRef.Path
			}
		}
		for _, itemBoundToRef := range state.ItemBoundToRef {
			if _, ok := ap.updatedReferencePaths[int(itemBoundToRef.ID)]; ok {
				continue
			}
			if _, ok := ap.includedElements[int(itemBoundToRef.ChildID)]; ok {
				ap.updatedReferencePaths[int(itemBoundToRef.ID)] = itemBoundToRef.Path
			}
		}
		for _, playerEquipmentSetRef := range patch.PlayerEquipmentSetRef {
			if _, ok := ap.includedElements[int(playerEquipmentSetRef.ChildID)]; ok {
				ap.updatedReferencePaths[int(playerEquipmentSetRef.ID)] = playerEquipmentSetRef.Path
			}
		}
		for _, playerEquipmentSetRef := range state.PlayerEquipmentSetRef {
			if _, ok := ap.updatedReferencePaths[int(playerEquipmentSetRef.ID)]; ok {
				continue
			}
			if _, ok := ap.includedElements[int(playerEquipmentSetRef.ChildID)]; ok {
				ap.updatedReferencePaths[int(playerEquipmentSetRef.ID)] = playerEquipmentSetRef.Path
			}
		}
		for _, playerGuildMemberRef := range patch.PlayerGuildMemberRef {
			if _, ok := ap.includedElements[int(playerGuildMemberRef.ChildID)]; ok {
				ap.updatedReferencePaths[int(playerGuildMemberRef.ID)] = playerGuildMemberRef.Path
			}
		}
		for _, playerGuildMemberRef := range state.PlayerGuildMemberRef {
			if _, ok := ap.updatedReferencePaths[int(playerGuildMemberRef.ID)]; ok {
				continue
			}
			if _, ok := ap.includedElements[int(playerGuildMemberRef.ChildID)]; ok {
				ap.updatedReferencePaths[int(playerGuildMemberRef.ID)] = playerGuildMemberRef.Path
			}
		}
		for _, playerTargetRef := range patch.PlayerTargetRef {
			if _, ok := ap.includedElements[int(playerTargetRef.ChildID)]; ok {
				ap.updatedReferencePaths[int(playerTargetRef.ID)] = playerTargetRef.Path
			}
		}
		for _, playerTargetRef := range state.PlayerTargetRef {
			if _, ok := ap.updatedReferencePaths[int(playerTargetRef.ID)]; ok {
				continue
			}
			if _, ok := ap.includedElements[int(playerTargetRef.ChildID)]; ok {
				ap.updatedReferencePaths[int(playerTargetRef.ID)] = playerTargetRef.Path
			}
		}
		for _, playerTargetedByRef := range patch.PlayerTargetedByRef {
			if _, ok := ap.includedElements[int(playerTargetedByRef.ChildID)]; ok {
				ap.updatedReferencePaths[int(playerTargetedByRef.ID)] = playerTargetedByRef.Path
			}
		}
		for _, playerTargetedByRef := range state.PlayerTargetedByRef {
			if _, ok := ap.updatedReferencePaths[int(playerTargetedByRef.ID)]; ok {
				continue
			}
			if _, ok := ap.includedElements[int(playerTargetedByRef.ChildID)]; ok {
				ap.updatedReferencePaths[int(playerTargetedByRef.ID)] = playerTargetedByRef.Path
			}
		}
	}
	for _, p := range ap.updatedElementPaths {
		ap.updatedPaths = append(ap.updatedPaths, p)
	}
	for _, p := range ap.updatedReferencePaths {
		ap.updatedPaths = append(ap.updatedPaths, p)
	}
}`
const fill_assemblePlanner_func string = `func (ap *assemblePlanner) fill(state *State) {
	for _, boolValue := range state.BoolValue {
		ap.updatedElementPaths[int(boolValue.ID)] = boolValue.Path
	}
	for _, floatValue := range state.FloatValue {
		ap.updatedElementPaths[int(floatValue.ID)] = floatValue.Path
	}
	for _, intValue := range state.IntValue {
		ap.updatedElementPaths[int(intValue.ID)] = intValue.Path
	}
	for _, stringValue := range state.StringValue {
		ap.updatedElementPaths[int(stringValue.ID)] = stringValue.Path
	}
	for _, attackEvent := range state.AttackEvent {
		ap.updatedElementPaths[int(attackEvent.ID)] = attackEvent.Path
	}
	for _, equipmentSet := range state.EquipmentSet {
		ap.updatedElementPaths[int(equipmentSet.ID)] = equipmentSet.Path
	}
	for _, gearScore := range state.GearScore {
		ap.updatedElementPaths[int(gearScore.ID)] = gearScore.Path
	}
	for _, item := range state.Item {
		ap.updatedElementPaths[int(item.ID)] = item.Path
	}
	for _, player := range state.Player {
		ap.updatedElementPaths[int(player.ID)] = player.Path
	}
	for _, position := range state.Position {
		ap.updatedElementPaths[int(position.ID)] = position.Path
	}
	for _, zone := range state.Zone {
		ap.updatedElementPaths[int(zone.ID)] = zone.Path
	}
	for _, zoneItem := range state.ZoneItem {
		ap.updatedElementPaths[int(zoneItem.ID)] = zoneItem.Path
	}
	for _, attackEventTargetRef := range state.AttackEventTargetRef {
		ap.updatedReferencePaths[int(attackEventTargetRef.ID)] = attackEventTargetRef.Path
	}
	for _, equipmentSetEquipmentRef := range state.EquipmentSetEquipmentRef {
		ap.updatedReferencePaths[int(equipmentSetEquipmentRef.ID)] = equipmentSetEquipmentRef.Path
	}
	for _, itemBoundToRef := range state.ItemBoundToRef {
		ap.updatedReferencePaths[int(itemBoundToRef.ID)] = itemBoundToRef.Path
	}
	for _, playerEquipmentSetRef := range state.PlayerEquipmentSetRef {
		ap.updatedReferencePaths[int(playerEquipmentSetRef.ID)] = playerEquipmentSetRef.Path
	}
	for _, playerGuildMemberRef := range state.PlayerGuildMemberRef {
		ap.updatedReferencePaths[int(playerGuildMemberRef.ID)] = playerGuildMemberRef.Path
	}
	for _, playerTargetRef := range state.PlayerTargetRef {
		ap.updatedReferencePaths[int(playerTargetRef.ID)] = playerTargetRef.Path
	}
	for _, playerTargetedByRef := range state.PlayerTargetedByRef {
		ap.updatedReferencePaths[int(playerTargetedByRef.ID)] = playerTargetedByRef.Path
	}
	for _, p := range ap.updatedElementPaths {
		ap.updatedPaths = append(ap.updatedPaths, p)
	}
	for _, p := range ap.updatedReferencePaths {
		ap.updatedPaths = append(ap.updatedPaths, p)
	}
}`
const _AssembleUpdateTree_Engine_func string = `func (engine *Engine) AssembleUpdateTree() {
	engine.planner.clear()
	engine.Tree.clear()
	engine.planner.plan(engine.State, engine.Patch)
	engine.assembleTree()
}`
const _AssembleFullTree_Engine_func string = `func (engine *Engine) AssembleFullTree() {
	engine.planner.clear()
	engine.Tree.clear()
	engine.planner.fill(engine.State)
	engine.assembleTree()
}`
const assembleTree_Engine_func string = `func (engine *Engine) assembleTree() {
	for _, p := range engine.planner.updatedPaths {
		switch p[0].Identifier {
		case attackEventIdentifier:
			child, ok := engine.Tree.AttackEvent[AttackEventID(p[0].ID)]
			if !ok {
				child = attackEvent{ID: AttackEventID(p[0].ID)}
			}
			engine.assembleAttackEventPath(&child, p, 0, engine.planner.includedElements)
			engine.Tree.AttackEvent[AttackEventID(p[0].ID)] = child
		case equipmentSetIdentifier:
			child, ok := engine.Tree.EquipmentSet[EquipmentSetID(p[0].ID)]
			if !ok {
				child = equipmentSet{ID: EquipmentSetID(p[0].ID)}
			}
			engine.assembleEquipmentSetPath(&child, p, 0, engine.planner.includedElements)
			engine.Tree.EquipmentSet[EquipmentSetID(p[0].ID)] = child
		case gearScoreIdentifier:
			child, ok := engine.Tree.GearScore[GearScoreID(p[0].ID)]
			if !ok {
				child = gearScore{ID: GearScoreID(p[0].ID)}
			}
			engine.assembleGearScorePath(&child, p, 0, engine.planner.includedElements)
			engine.Tree.GearScore[GearScoreID(p[0].ID)] = child
		case itemIdentifier:
			child, ok := engine.Tree.Item[ItemID(p[0].ID)]
			if !ok {
				child = item{ID: ItemID(p[0].ID)}
			}
			engine.assembleItemPath(&child, p, 0, engine.planner.includedElements)
			engine.Tree.Item[ItemID(p[0].ID)] = child
		case playerIdentifier:
			child, ok := engine.Tree.Player[PlayerID(p[0].ID)]
			if !ok {
				child = player{ID: PlayerID(p[0].ID)}
			}
			engine.assemblePlayerPath(&child, p, 0, engine.planner.includedElements)
			engine.Tree.Player[PlayerID(p[0].ID)] = child
		case positionIdentifier:
			child, ok := engine.Tree.Position[PositionID(p[0].ID)]
			if !ok {
				child = position{ID: PositionID(p[0].ID)}
			}
			engine.assemblePositionPath(&child, p, 0, engine.planner.includedElements)
			engine.Tree.Position[PositionID(p[0].ID)] = child
		case zoneIdentifier:
			child, ok := engine.Tree.Zone[ZoneID(p[0].ID)]
			if !ok {
				child = zone{ID: ZoneID(p[0].ID)}
			}
			engine.assembleZonePath(&child, p, 0, engine.planner.includedElements)
			engine.Tree.Zone[ZoneID(p[0].ID)] = child
		case zoneItemIdentifier:
			child, ok := engine.Tree.ZoneItem[ZoneItemID(p[0].ID)]
			if !ok {
				child = zoneItem{ID: ZoneItemID(p[0].ID)}
			}
			engine.assembleZoneItemPath(&child, p, 0, engine.planner.includedElements)
			engine.Tree.ZoneItem[ZoneItemID(p[0].ID)] = child
		}
	}
}`
const createBoolValue_Engine_func string = `func (engine *Engine) createBoolValue(p path, fieldIdentifier treeFieldIdentifier, value bool) boolValue {
	var element boolValue
	element.Value = value
	element.engine = engine
	element.ID = BoolValueID(engine.GenerateID())
	element.Path = p.extendAndCopy(fieldIdentifier, int(element.ID), ElementKindBoolValue, 0)
	element.JSONPath = element.Path.toJSONPath()
	element.OperationKind = OperationKindUpdate
	engine.Patch.BoolValue[element.ID] = element
	return element
}`
const createIntValue_Engine_func string = `func (engine *Engine) createIntValue(p path, fieldIdentifier treeFieldIdentifier, value int64) intValue {
	var element intValue
	element.Value = value
	element.engine = engine
	element.ID = IntValueID(engine.GenerateID())
	element.Path = p.extendAndCopy(fieldIdentifier, int(element.ID), ElementKindIntValue, 0)
	element.JSONPath = element.Path.toJSONPath()
	element.OperationKind = OperationKindUpdate
	engine.Patch.IntValue[element.ID] = element
	return element
}`
const createFloatValue_Engine_func string = `func (engine *Engine) createFloatValue(p path, fieldIdentifier treeFieldIdentifier, value float64) floatValue {
	var element floatValue
	element.Value = value
	element.engine = engine
	element.ID = FloatValueID(engine.GenerateID())
	element.Path = p.extendAndCopy(fieldIdentifier, int(element.ID), ElementKindFloatValue, 0)
	element.JSONPath = element.Path.toJSONPath()
	element.OperationKind = OperationKindUpdate
	engine.Patch.FloatValue[element.ID] = element
	return element
}`
const createStringValue_Engine_func string = `func (engine *Engine) createStringValue(p path, fieldIdentifier treeFieldIdentifier, value string) stringValue {
	var element stringValue
	element.Value = value
	element.engine = engine
	element.ID = StringValueID(engine.GenerateID())
	element.Path = p.extendAndCopy(fieldIdentifier, int(element.ID), ElementKindStringValue, 0)
	element.JSONPath = element.Path.toJSONPath()
	element.OperationKind = OperationKindUpdate
	engine.Patch.StringValue[element.ID] = element
	return element
}`
const _CreateEquipmentSet_Engine_func string = `func (engine *Engine) CreateEquipmentSet() EquipmentSet {
	return engine.createEquipmentSet(newPath(), equipmentSetIdentifier)
}`
const createEquipmentSet_Engine_func string = `func (engine *Engine) createEquipmentSet(p path, fieldIdentifier treeFieldIdentifier) EquipmentSet {
	var element equipmentSetCore
	element.engine = engine
	element.ID = EquipmentSetID(engine.GenerateID())
	element.Path = p.extendAndCopy(fieldIdentifier, int(element.ID), ElementKindEquipmentSet, 0)
	element.JSONPath = element.Path.toJSONPath()
	elementName := engine.createStringValue(element.Path, equipmentSet_nameIdentifier, "")
	element.Name = elementName.ID
	element.OperationKind = OperationKindUpdate
	element.HasParent = len(element.Path) > 1
	engine.Patch.EquipmentSet[element.ID] = element
	return EquipmentSet{equipmentSet: element}
}`
const _CreateGearScore_Engine_func string = `func (engine *Engine) CreateGearScore() GearScore {
	return engine.createGearScore(newPath(), gearScoreIdentifier)
}`
const createGearScore_Engine_func string = `func (engine *Engine) createGearScore(p path, fieldIdentifier treeFieldIdentifier) GearScore {
	var element gearScoreCore
	element.engine = engine
	element.ID = GearScoreID(engine.GenerateID())
	element.Path = p.extendAndCopy(fieldIdentifier, int(element.ID), ElementKindGearScore, 0)
	element.JSONPath = element.Path.toJSONPath()
	elementLevel := engine.createIntValue(element.Path, gearScore_levelIdentifier, 0)
	element.Level = elementLevel.ID
	elementScore := engine.createIntValue(element.Path, gearScore_scoreIdentifier, 0)
	element.Score = elementScore.ID
	element.OperationKind = OperationKindUpdate
	element.HasParent = len(element.Path) > 1
	engine.Patch.GearScore[element.ID] = element
	return GearScore{gearScore: element}
}`
const _CreatePosition_Engine_func string = `func (engine *Engine) CreatePosition() Position {
	return engine.createPosition(newPath(), positionIdentifier)
}`
const createPosition_Engine_func string = `func (engine *Engine) createPosition(p path, fieldIdentifier treeFieldIdentifier) Position {
	var element positionCore
	element.engine = engine
	element.ID = PositionID(engine.GenerateID())
	element.Path = p.extendAndCopy(fieldIdentifier, int(element.ID), ElementKindPosition, 0)
	element.JSONPath = element.Path.toJSONPath()
	elementX := engine.createFloatValue(element.Path, position_xIdentifier, 0.0)
	element.X = elementX.ID
	elementY := engine.createFloatValue(element.Path, position_yIdentifier, 0.0)
	element.Y = elementY.ID
	element.OperationKind = OperationKindUpdate
	element.HasParent = len(element.Path) > 1
	engine.Patch.Position[element.ID] = element
	return Position{position: element}
}`
const _CreateAttackEvent_Engine_func string = `func (engine *Engine) CreateAttackEvent() AttackEvent {
	return engine.createAttackEvent(newPath(), attackEventIdentifier)
}`
const createAttackEvent_Engine_func string = `func (engine *Engine) createAttackEvent(p path, fieldIdentifier treeFieldIdentifier) AttackEvent {
	var element attackEventCore
	element.engine = engine
	element.ID = AttackEventID(engine.GenerateID())
	element.Path = p.extendAndCopy(fieldIdentifier, int(element.ID), ElementKindAttackEvent, 0)
	element.JSONPath = element.Path.toJSONPath()
	element.OperationKind = OperationKindUpdate
	element.HasParent = len(element.Path) > 1
	engine.Patch.AttackEvent[element.ID] = element
	return AttackEvent{attackEvent: element}
}`
const _CreateItem_Engine_func string = `func (engine *Engine) CreateItem() Item {
	return engine.createItem(newPath(), itemIdentifier)
}`
const createItem_Engine_func string = `func (engine *Engine) createItem(p path, fieldIdentifier treeFieldIdentifier) Item {
	var element itemCore
	element.engine = engine
	element.ID = ItemID(engine.GenerateID())
	element.Path = p.extendAndCopy(fieldIdentifier, int(element.ID), ElementKindItem, 0)
	element.JSONPath = element.Path.toJSONPath()
	elementGearScore := engine.createGearScore(element.Path, item_gearScoreIdentifier)
	element.GearScore = elementGearScore.gearScore.ID
	elementName := engine.createStringValue(element.Path, item_nameIdentifier, "")
	element.Name = elementName.ID
	originElement := engine.createPlayer(element.Path, item_originIdentifier)
	elementOrigin := engine.createAnyOfPlayer_Position(int(element.ID), int(originElement.player.ID), ElementKindPlayer, element.Path, item_originIdentifier)
	element.Origin = elementOrigin.anyOfPlayer_Position.ID
	element.OperationKind = OperationKindUpdate
	element.HasParent = len(element.Path) > 1
	engine.Patch.Item[element.ID] = element
	return Item{item: element}
}`
const _CreateZoneItem_Engine_func string = `func (engine *Engine) CreateZoneItem() ZoneItem {
	return engine.createZoneItem(newPath(), zoneItemIdentifier)
}`
const createZoneItem_Engine_func string = `func (engine *Engine) createZoneItem(p path, fieldIdentifier treeFieldIdentifier) ZoneItem {
	var element zoneItemCore
	element.engine = engine
	element.ID = ZoneItemID(engine.GenerateID())
	element.Path = p.extendAndCopy(fieldIdentifier, int(element.ID), ElementKindZoneItem, 0)
	element.JSONPath = element.Path.toJSONPath()
	elementItem := engine.createItem(element.Path, zoneItem_itemIdentifier)
	element.Item = elementItem.item.ID
	elementPosition := engine.createPosition(element.Path, zoneItem_positionIdentifier)
	element.Position = elementPosition.position.ID
	element.OperationKind = OperationKindUpdate
	element.HasParent = len(element.Path) > 1
	engine.Patch.ZoneItem[element.ID] = element
	return ZoneItem{zoneItem: element}
}`
const _CreatePlayer_Engine_func string = `func (engine *Engine) CreatePlayer() Player {
	return engine.createPlayer(newPath(), playerIdentifier)
}`
const createPlayer_Engine_func string = `func (engine *Engine) createPlayer(p path, fieldIdentifier treeFieldIdentifier) Player {
	var element playerCore
	element.engine = engine
	element.ID = PlayerID(engine.GenerateID())
	element.Path = p.extendAndCopy(fieldIdentifier, int(element.ID), ElementKindPlayer, 0)
	element.JSONPath = element.Path.toJSONPath()
	elementGearScore := engine.createGearScore(element.Path, player_gearScoreIdentifier)
	element.GearScore = elementGearScore.gearScore.ID
	elementPosition := engine.createPosition(element.Path, player_positionIdentifier)
	element.Position = elementPosition.position.ID
	element.OperationKind = OperationKindUpdate
	element.HasParent = len(element.Path) > 1
	engine.Patch.Player[element.ID] = element
	return Player{player: element}
}`
const _CreateZone_Engine_func string = `func (engine *Engine) CreateZone() Zone {
	return engine.createZone(newPath(), zoneIdentifier)
}`
const createZone_Engine_func string = `func (engine *Engine) createZone(p path, fieldIdentifier treeFieldIdentifier) Zone {
	var element zoneCore
	element.engine = engine
	element.ID = ZoneID(engine.GenerateID())
	element.Path = p.extendAndCopy(fieldIdentifier, int(element.ID), ElementKindZone, 0)
	element.JSONPath = element.Path.toJSONPath()
	element.OperationKind = OperationKindUpdate
	element.HasParent = len(element.Path) > 1
	engine.Patch.Zone[element.ID] = element
	return Zone{zone: element}
}`
const createAttackEventTargetRef_Engine_func string = `func (engine *Engine) createAttackEventTargetRef(p path, fieldIdentifier treeFieldIdentifier, referencedElementID PlayerID, parentID AttackEventID, childID int) attackEventTargetRefCore {
	var element attackEventTargetRefCore
	element.engine = engine
	element.ReferencedElementID = referencedElementID
	element.ParentID = parentID
	element.ChildID = childID
	element.ID = AttackEventTargetRefID(engine.GenerateID())
	element.Path = p.extendAndCopy(fieldIdentifier, 0, ElementKindPlayer, int(element.ID))
	element.OperationKind = OperationKindUpdate
	engine.Patch.AttackEventTargetRef[element.ID] = element
	return element
}`
const createItemBoundToRef_Engine_func string = `func (engine *Engine) createItemBoundToRef(p path, fieldIdentifier treeFieldIdentifier, referencedElementID PlayerID, parentID ItemID, childID int) itemBoundToRefCore {
	var element itemBoundToRefCore
	element.engine = engine
	element.ReferencedElementID = referencedElementID
	element.ParentID = parentID
	element.ChildID = childID
	element.ID = ItemBoundToRefID(engine.GenerateID())
	element.Path = p.extendAndCopy(fieldIdentifier, 0, ElementKindPlayer, int(element.ID))
	element.OperationKind = OperationKindUpdate
	engine.Patch.ItemBoundToRef[element.ID] = element
	return element
}`
const createPlayerGuildMemberRef_Engine_func string = `func (engine *Engine) createPlayerGuildMemberRef(p path, fieldIdentifier treeFieldIdentifier, referencedElementID PlayerID, parentID PlayerID, childID int) playerGuildMemberRefCore {
	var element playerGuildMemberRefCore
	element.engine = engine
	element.ReferencedElementID = referencedElementID
	element.ParentID = parentID
	element.ChildID = childID
	element.ID = PlayerGuildMemberRefID(engine.GenerateID())
	element.Path = p.extendAndCopy(fieldIdentifier, 0, ElementKindPlayer, int(element.ID))
	element.OperationKind = OperationKindUpdate
	engine.Patch.PlayerGuildMemberRef[element.ID] = element
	return element
}`
const createEquipmentSetEquipmentRef_Engine_func string = `func (engine *Engine) createEquipmentSetEquipmentRef(p path, fieldIdentifier treeFieldIdentifier, referencedElementID ItemID, parentID EquipmentSetID, childID int) equipmentSetEquipmentRefCore {
	var element equipmentSetEquipmentRefCore
	element.engine = engine
	element.ReferencedElementID = referencedElementID
	element.ParentID = parentID
	element.ChildID = childID
	element.ID = EquipmentSetEquipmentRefID(engine.GenerateID())
	element.Path = p.extendAndCopy(fieldIdentifier, 0, ElementKindItem, int(element.ID))
	element.OperationKind = OperationKindUpdate
	engine.Patch.EquipmentSetEquipmentRef[element.ID] = element
	return element
}`
const createPlayerEquipmentSetRef_Engine_func string = `func (engine *Engine) createPlayerEquipmentSetRef(p path, fieldIdentifier treeFieldIdentifier, referencedElementID EquipmentSetID, parentID PlayerID, childID int) playerEquipmentSetRefCore {
	var element playerEquipmentSetRefCore
	element.engine = engine
	element.ReferencedElementID = referencedElementID
	element.ParentID = parentID
	element.ChildID = childID
	element.ID = PlayerEquipmentSetRefID(engine.GenerateID())
	element.Path = p.extendAndCopy(fieldIdentifier, 0, ElementKindEquipmentSet, int(element.ID))
	element.OperationKind = OperationKindUpdate
	engine.Patch.PlayerEquipmentSetRef[element.ID] = element
	return element
}`
const createPlayerTargetRef_Engine_func string = `func (engine *Engine) createPlayerTargetRef(p path, fieldIdentifier treeFieldIdentifier, referencedElementID AnyOfPlayer_ZoneItemID, parentID PlayerID, childKind ElementKind, childID int) playerTargetRefCore {
	var element playerTargetRefCore
	element.engine = engine
	element.ReferencedElementID = referencedElementID
	element.ParentID = parentID
	element.ChildID = childID
	element.ID = PlayerTargetRefID(engine.GenerateID())
	element.Path = p.extendAndCopy(fieldIdentifier, 0, childKind, int(element.ID))
	element.OperationKind = OperationKindUpdate
	engine.Patch.PlayerTargetRef[element.ID] = element
	return element
}`
const createPlayerTargetedByRef_Engine_func string = `func (engine *Engine) createPlayerTargetedByRef(p path, fieldIdentifier treeFieldIdentifier, referencedElementID AnyOfPlayer_ZoneItemID, parentID PlayerID, childKind ElementKind, childID int) playerTargetedByRefCore {
	var element playerTargetedByRefCore
	element.engine = engine
	element.ReferencedElementID = referencedElementID
	element.ParentID = parentID
	element.ChildID = childID
	element.ID = PlayerTargetedByRefID(engine.GenerateID())
	element.Path = p.extendAndCopy(fieldIdentifier, 0, childKind, int(element.ID))
	element.OperationKind = OperationKindUpdate
	engine.Patch.PlayerTargetedByRef[element.ID] = element
	return element
}`
const createAnyOfPlayer_ZoneItem_Engine_func string = `func (engine *Engine) createAnyOfPlayer_ZoneItem(parentID int, childID int, childKind ElementKind, p path, fieldIdentifier treeFieldIdentifier) AnyOfPlayer_ZoneItem {
	var element anyOfPlayer_ZoneItemCore
	element.engine = engine
	element.ID = AnyOfPlayer_ZoneItemID(engine.GenerateID())
	element.ParentID = parentID
	element.ChildID = childID
	element.ElementKind = childKind
	element.OperationKind = OperationKindUpdate
	element.ParentElementPath = p
	element.FieldIdentifier = fieldIdentifier
	engine.Patch.AnyOfPlayer_ZoneItem[element.ID] = element
	return AnyOfPlayer_ZoneItem{anyOfPlayer_ZoneItem: element}
}`
const createAnyOfPlayer_Position_Engine_func string = `func (engine *Engine) createAnyOfPlayer_Position(parentID int, childID int, childKind ElementKind, p path, fieldIdentifier treeFieldIdentifier) AnyOfPlayer_Position {
	var element anyOfPlayer_PositionCore
	element.engine = engine
	element.ID = AnyOfPlayer_PositionID(engine.GenerateID())
	element.ParentID = parentID
	element.ChildID = childID
	element.ElementKind = childKind
	element.OperationKind = OperationKindUpdate
	element.ParentElementPath = p
	element.FieldIdentifier = fieldIdentifier
	engine.Patch.AnyOfPlayer_Position[element.ID] = element
	return AnyOfPlayer_Position{anyOfPlayer_Position: element}
}`
const createAnyOfItem_Player_ZoneItem_Engine_func string = `func (engine *Engine) createAnyOfItem_Player_ZoneItem(parentID int, childID int, childKind ElementKind, p path, fieldIdentifier treeFieldIdentifier) AnyOfItem_Player_ZoneItem {
	var element anyOfItem_Player_ZoneItemCore
	element.engine = engine
	element.ID = AnyOfItem_Player_ZoneItemID(engine.GenerateID())
	element.ParentID = parentID
	element.ChildID = childID
	element.ElementKind = childKind
	element.OperationKind = OperationKindUpdate
	element.ParentElementPath = p
	element.FieldIdentifier = fieldIdentifier
	engine.Patch.AnyOfItem_Player_ZoneItem[element.ID] = element
	return AnyOfItem_Player_ZoneItem{anyOfItem_Player_ZoneItem: element}
}`
const _DeletePlayer_Engine_func string = `func (engine *Engine) DeletePlayer(playerID PlayerID) {
	player := engine.Player(playerID).player
	if player.HasParent {
		return
	}
	engine.deletePlayer(playerID)
}`
const deletePlayer_Engine_func string = `func (engine *Engine) deletePlayer(playerID PlayerID) {
	player := engine.Player(playerID).player
	if player.OperationKind == OperationKindDelete {
		return
	}
	engine.dereferenceAttackEventTargetRefs(playerID)
	engine.dereferenceItemBoundToRefs(playerID)
	engine.dereferencePlayerGuildMemberRefs(playerID)
	engine.dereferencePlayerTargetRefsPlayer(playerID)
	engine.dereferencePlayerTargetedByRefsPlayer(playerID)
	for _, actionID := range player.Action {
		engine.deleteAttackEvent(actionID)
	}
	for _, equipmentSetID := range player.EquipmentSets {
		engine.deletePlayerEquipmentSetRef(equipmentSetID)
	}
	engine.deleteGearScore(player.GearScore)
	for _, guildMemberID := range player.GuildMembers {
		engine.deletePlayerGuildMemberRef(guildMemberID)
	}
	for _, itemID := range player.Items {
		engine.deleteItem(itemID)
	}
	engine.deletePosition(player.Position)
	engine.deletePlayerTargetRef(player.Target)
	for _, targetedByID := range player.TargetedBy {
		engine.deletePlayerTargetedByRef(targetedByID)
	}
	if _, ok := engine.State.Player[playerID]; ok {
		player.OperationKind = OperationKindDelete
		engine.Patch.Player[player.ID] = player
	} else {
		delete(engine.Patch.Player, playerID)
	}
}`
const deleteBoolValue_Engine_func string = `func (engine *Engine) deleteBoolValue(boolValueID BoolValueID) {
	boolValue := engine.boolValue(boolValueID)
	if boolValue.OperationKind == OperationKindDelete {
		return
	}
	if _, ok := engine.State.BoolValue[boolValueID]; ok {
		boolValue.OperationKind = OperationKindDelete
		engine.Patch.BoolValue[boolValue.ID] = boolValue
	} else {
		delete(engine.Patch.BoolValue, boolValueID)
	}
}`
const deleteIntValue_Engine_func string = `func (engine *Engine) deleteIntValue(intValueID IntValueID) {
	intValue := engine.intValue(intValueID)
	if intValue.OperationKind == OperationKindDelete {
		return
	}
	if _, ok := engine.State.IntValue[intValueID]; ok {
		intValue.OperationKind = OperationKindDelete
		engine.Patch.IntValue[intValue.ID] = intValue
	} else {
		delete(engine.Patch.IntValue, intValueID)
	}
}`
const deleteFloatValue_Engine_func string = `func (engine *Engine) deleteFloatValue(floatValueID FloatValueID) {
	floatValue := engine.floatValue(floatValueID)
	if floatValue.OperationKind == OperationKindDelete {
		return
	}
	if _, ok := engine.State.FloatValue[floatValueID]; ok {
		floatValue.OperationKind = OperationKindDelete
		engine.Patch.FloatValue[floatValue.ID] = floatValue
	} else {
		delete(engine.Patch.FloatValue, floatValueID)
	}
}`
const deleteStringValue_Engine_func string = `func (engine *Engine) deleteStringValue(stringValueID StringValueID) {
	stringValue := engine.stringValue(stringValueID)
	if stringValue.OperationKind == OperationKindDelete {
		return
	}
	if _, ok := engine.State.StringValue[stringValueID]; ok {
		stringValue.OperationKind = OperationKindDelete
		engine.Patch.StringValue[stringValue.ID] = stringValue
	} else {
		delete(engine.Patch.StringValue, stringValueID)
	}
}`
const _DeleteGearScore_Engine_func string = `func (engine *Engine) DeleteGearScore(gearScoreID GearScoreID) {
	gearScore := engine.GearScore(gearScoreID).gearScore
	if gearScore.HasParent {
		return
	}
	engine.deleteGearScore(gearScoreID)
}`
const deleteGearScore_Engine_func string = `func (engine *Engine) deleteGearScore(gearScoreID GearScoreID) {
	gearScore := engine.GearScore(gearScoreID).gearScore
	if gearScore.OperationKind == OperationKindDelete {
		return
	}
	engine.deleteIntValue(gearScore.Level)
	engine.deleteIntValue(gearScore.Score)
	if _, ok := engine.State.GearScore[gearScoreID]; ok {
		gearScore.OperationKind = OperationKindDelete
		engine.Patch.GearScore[gearScore.ID] = gearScore
	} else {
		delete(engine.Patch.GearScore, gearScoreID)
	}
}`
const _DeletePosition_Engine_func string = `func (engine *Engine) DeletePosition(positionID PositionID) {
	position := engine.Position(positionID).position
	if position.HasParent {
		return
	}
	engine.deletePosition(positionID)
}`
const deletePosition_Engine_func string = `func (engine *Engine) deletePosition(positionID PositionID) {
	position := engine.Position(positionID).position
	if position.OperationKind == OperationKindDelete {
		return
	}
	engine.deleteFloatValue(position.X)
	engine.deleteFloatValue(position.Y)
	if _, ok := engine.State.Position[positionID]; ok {
		position.OperationKind = OperationKindDelete
		engine.Patch.Position[position.ID] = position
	} else {
		delete(engine.Patch.Position, positionID)
	}
}`
const _DeleteAttackEvent_Engine_func string = `func (engine *Engine) DeleteAttackEvent(attackEventID AttackEventID) {
	attackEvent := engine.AttackEvent(attackEventID).attackEvent
	if attackEvent.HasParent {
		return
	}
	engine.deleteAttackEvent(attackEventID)
}`
const deleteAttackEvent_Engine_func string = `func (engine *Engine) deleteAttackEvent(attackEventID AttackEventID) {
	attackEvent := engine.AttackEvent(attackEventID).attackEvent
	if attackEvent.OperationKind == OperationKindDelete {
		return
	}
	engine.deleteAttackEventTargetRef(attackEvent.Target)
	if _, ok := engine.State.AttackEvent[attackEventID]; ok {
		attackEvent.OperationKind = OperationKindDelete
		engine.Patch.AttackEvent[attackEvent.ID] = attackEvent
	} else {
		delete(engine.Patch.AttackEvent, attackEventID)
	}
}`
const _DeleteItem_Engine_func string = `func (engine *Engine) DeleteItem(itemID ItemID) {
	item := engine.Item(itemID).item
	if item.HasParent {
		return
	}
	engine.deleteItem(itemID)
}`
const deleteItem_Engine_func string = `func (engine *Engine) deleteItem(itemID ItemID) {
	item := engine.Item(itemID).item
	if item.OperationKind == OperationKindDelete {
		return
	}
	engine.dereferenceEquipmentSetEquipmentRefs(itemID)
	engine.deleteItemBoundToRef(item.BoundTo)
	engine.deleteGearScore(item.GearScore)
	engine.deleteStringValue(item.Name)
	engine.deleteAnyOfPlayer_Position(item.Origin, true)
	if _, ok := engine.State.Item[itemID]; ok {
		item.OperationKind = OperationKindDelete
		engine.Patch.Item[item.ID] = item
	} else {
		delete(engine.Patch.Item, itemID)
	}
}`
const _DeleteZoneItem_Engine_func string = `func (engine *Engine) DeleteZoneItem(zoneItemID ZoneItemID) {
	zoneItem := engine.ZoneItem(zoneItemID).zoneItem
	if zoneItem.HasParent {
		return
	}
	engine.deleteZoneItem(zoneItemID)
}`
const deleteZoneItem_Engine_func string = `func (engine *Engine) deleteZoneItem(zoneItemID ZoneItemID) {
	zoneItem := engine.ZoneItem(zoneItemID).zoneItem
	if zoneItem.OperationKind == OperationKindDelete {
		return
	}
	engine.dereferencePlayerTargetRefsZoneItem(zoneItemID)
	engine.dereferencePlayerTargetedByRefsZoneItem(zoneItemID)
	engine.deleteItem(zoneItem.Item)
	engine.deletePosition(zoneItem.Position)
	if _, ok := engine.State.ZoneItem[zoneItemID]; ok {
		zoneItem.OperationKind = OperationKindDelete
		engine.Patch.ZoneItem[zoneItem.ID] = zoneItem
	} else {
		delete(engine.Patch.ZoneItem, zoneItemID)
	}
}`
const _DeleteZone_Engine_func string = `func (engine *Engine) DeleteZone(zoneID ZoneID) {
	engine.deleteZone(zoneID)
}`
const deleteZone_Engine_func string = `func (engine *Engine) deleteZone(zoneID ZoneID) {
	zone := engine.Zone(zoneID).zone
	if zone.OperationKind == OperationKindDelete {
		return
	}
	for _, interactableID := range zone.Interactables {
		engine.deleteAnyOfItem_Player_ZoneItem(interactableID, true)
	}
	for _, itemID := range zone.Items {
		engine.deleteZoneItem(itemID)
	}
	for _, playerID := range zone.Players {
		engine.deletePlayer(playerID)
	}
	for _, tagID := range zone.Tags {
		engine.deleteStringValue(tagID)
	}
	if _, ok := engine.State.Zone[zoneID]; ok {
		zone.OperationKind = OperationKindDelete
		engine.Patch.Zone[zone.ID] = zone
	} else {
		delete(engine.Patch.Zone, zoneID)
	}
}`
const _DeleteEquipmentSet_Engine_func string = `func (engine *Engine) DeleteEquipmentSet(equipmentSetID EquipmentSetID) {
	engine.deleteEquipmentSet(equipmentSetID)
}`
const deleteEquipmentSet_Engine_func string = `func (engine *Engine) deleteEquipmentSet(equipmentSetID EquipmentSetID) {
	equipmentSet := engine.EquipmentSet(equipmentSetID).equipmentSet
	if equipmentSet.OperationKind == OperationKindDelete {
		return
	}
	engine.dereferencePlayerEquipmentSetRefs(equipmentSetID)
	for _, equipmentID := range equipmentSet.Equipment {
		engine.deleteEquipmentSetEquipmentRef(equipmentID)
	}
	engine.deleteStringValue(equipmentSet.Name)
	if _, ok := engine.State.EquipmentSet[equipmentSetID]; ok {
		equipmentSet.OperationKind = OperationKindDelete
		engine.Patch.EquipmentSet[equipmentSet.ID] = equipmentSet
	} else {
		delete(engine.Patch.EquipmentSet, equipmentSetID)
	}
}`
const deletePlayerGuildMemberRef_Engine_func string = `func (engine *Engine) deletePlayerGuildMemberRef(playerGuildMemberRefID PlayerGuildMemberRefID) {
	playerGuildMemberRef := engine.playerGuildMemberRef(playerGuildMemberRefID).playerGuildMemberRef
	if playerGuildMemberRef.OperationKind == OperationKindDelete {
		return
	}
	if _, ok := engine.State.PlayerGuildMemberRef[playerGuildMemberRefID]; ok {
		playerGuildMemberRef.OperationKind = OperationKindDelete
		engine.Patch.PlayerGuildMemberRef[playerGuildMemberRef.ID] = playerGuildMemberRef
	} else {
		delete(engine.Patch.PlayerGuildMemberRef, playerGuildMemberRefID)
	}
}`
const deletePlayerEquipmentSetRef_Engine_func string = `func (engine *Engine) deletePlayerEquipmentSetRef(playerEquipmentSetRefID PlayerEquipmentSetRefID) {
	playerEquipmentSetRef := engine.playerEquipmentSetRef(playerEquipmentSetRefID).playerEquipmentSetRef
	if playerEquipmentSetRef.OperationKind == OperationKindDelete {
		return
	}
	if _, ok := engine.State.PlayerEquipmentSetRef[playerEquipmentSetRefID]; ok {
		playerEquipmentSetRef.OperationKind = OperationKindDelete
		engine.Patch.PlayerEquipmentSetRef[playerEquipmentSetRef.ID] = playerEquipmentSetRef
	} else {
		delete(engine.Patch.PlayerEquipmentSetRef, playerEquipmentSetRefID)
	}
}`
const deleteItemBoundToRef_Engine_func string = `func (engine *Engine) deleteItemBoundToRef(itemBoundToRefID ItemBoundToRefID) {
	itemBoundToRef := engine.itemBoundToRef(itemBoundToRefID).itemBoundToRef
	if itemBoundToRef.OperationKind == OperationKindDelete {
		return
	}
	if _, ok := engine.State.ItemBoundToRef[itemBoundToRefID]; ok {
		itemBoundToRef.OperationKind = OperationKindDelete
		engine.Patch.ItemBoundToRef[itemBoundToRef.ID] = itemBoundToRef
	} else {
		delete(engine.Patch.ItemBoundToRef, itemBoundToRefID)
	}
}`
const deleteAttackEventTargetRef_Engine_func string = `func (engine *Engine) deleteAttackEventTargetRef(attackEventTargetRefID AttackEventTargetRefID) {
	attackEventTargetRef := engine.attackEventTargetRef(attackEventTargetRefID).attackEventTargetRef
	if attackEventTargetRef.OperationKind == OperationKindDelete {
		return
	}
	if _, ok := engine.State.AttackEventTargetRef[attackEventTargetRefID]; ok {
		attackEventTargetRef.OperationKind = OperationKindDelete
		engine.Patch.AttackEventTargetRef[attackEventTargetRef.ID] = attackEventTargetRef
	} else {
		delete(engine.Patch.AttackEventTargetRef, attackEventTargetRefID)
	}
}`
const deleteEquipmentSetEquipmentRef_Engine_func string = `func (engine *Engine) deleteEquipmentSetEquipmentRef(equipmentSetEquipmentRefID EquipmentSetEquipmentRefID) {
	equipmentSetEquipmentRef := engine.equipmentSetEquipmentRef(equipmentSetEquipmentRefID).equipmentSetEquipmentRef
	if equipmentSetEquipmentRef.OperationKind == OperationKindDelete {
		return
	}
	if _, ok := engine.State.EquipmentSetEquipmentRef[equipmentSetEquipmentRefID]; ok {
		equipmentSetEquipmentRef.OperationKind = OperationKindDelete
		engine.Patch.EquipmentSetEquipmentRef[equipmentSetEquipmentRef.ID] = equipmentSetEquipmentRef
	} else {
		delete(engine.Patch.EquipmentSetEquipmentRef, equipmentSetEquipmentRefID)
	}
}`
const deletePlayerTargetRef_Engine_func string = `func (engine *Engine) deletePlayerTargetRef(playerTargetRefID PlayerTargetRefID) {
	playerTargetRef := engine.playerTargetRef(playerTargetRefID).playerTargetRef
	if playerTargetRef.OperationKind == OperationKindDelete {
		return
	}
	engine.deleteAnyOfPlayer_ZoneItem(playerTargetRef.ReferencedElementID, false)
	if _, ok := engine.State.PlayerTargetRef[playerTargetRefID]; ok {
		playerTargetRef.OperationKind = OperationKindDelete
		engine.Patch.PlayerTargetRef[playerTargetRef.ID] = playerTargetRef
	} else {
		delete(engine.Patch.PlayerTargetRef, playerTargetRefID)
	}
}`
const deletePlayerTargetedByRef_Engine_func string = `func (engine *Engine) deletePlayerTargetedByRef(playerTargetedByRefID PlayerTargetedByRefID) {
	playerTargetedByRef := engine.playerTargetedByRef(playerTargetedByRefID).playerTargetedByRef
	if playerTargetedByRef.OperationKind == OperationKindDelete {
		return
	}
	engine.deleteAnyOfPlayer_ZoneItem(playerTargetedByRef.ReferencedElementID, false)
	if _, ok := engine.State.PlayerTargetedByRef[playerTargetedByRefID]; ok {
		playerTargetedByRef.OperationKind = OperationKindDelete
		engine.Patch.PlayerTargetedByRef[playerTargetedByRef.ID] = playerTargetedByRef
	} else {
		delete(engine.Patch.PlayerTargetedByRef, playerTargetedByRefID)
	}
}`
const deleteAnyOfPlayer_ZoneItem_Engine_func string = `func (engine *Engine) deleteAnyOfPlayer_ZoneItem(anyOfPlayer_ZoneItemID AnyOfPlayer_ZoneItemID, deleteChild bool) {
	anyOfPlayer_ZoneItem := engine.anyOfPlayer_ZoneItem(anyOfPlayer_ZoneItemID).anyOfPlayer_ZoneItem
	if anyOfPlayer_ZoneItem.OperationKind == OperationKindDelete {
		return
	}
	if deleteChild {
		anyOfPlayer_ZoneItem.deleteChild()
	}
	if _, ok := engine.State.AnyOfPlayer_ZoneItem[anyOfPlayer_ZoneItemID]; ok {
		anyOfPlayer_ZoneItem.OperationKind = OperationKindDelete
		engine.Patch.AnyOfPlayer_ZoneItem[anyOfPlayer_ZoneItem.ID] = anyOfPlayer_ZoneItem
	} else {
		delete(engine.Patch.AnyOfPlayer_ZoneItem, anyOfPlayer_ZoneItemID)
	}
}`
const deleteAnyOfPlayer_Position_Engine_func string = `func (engine *Engine) deleteAnyOfPlayer_Position(anyOfPlayer_PositionID AnyOfPlayer_PositionID, deleteChild bool) {
	anyOfPlayer_Position := engine.anyOfPlayer_Position(anyOfPlayer_PositionID).anyOfPlayer_Position
	if anyOfPlayer_Position.OperationKind == OperationKindDelete {
		return
	}
	if deleteChild {
		anyOfPlayer_Position.deleteChild()
	}
	if _, ok := engine.State.AnyOfPlayer_Position[anyOfPlayer_PositionID]; ok {
		anyOfPlayer_Position.OperationKind = OperationKindDelete
		engine.Patch.AnyOfPlayer_Position[anyOfPlayer_Position.ID] = anyOfPlayer_Position
	} else {
		delete(engine.Patch.AnyOfPlayer_Position, anyOfPlayer_PositionID)
	}
}`
const deleteAnyOfItem_Player_ZoneItem_Engine_func string = `func (engine *Engine) deleteAnyOfItem_Player_ZoneItem(anyOfItem_Player_ZoneItemID AnyOfItem_Player_ZoneItemID, deleteChild bool) {
	anyOfItem_Player_ZoneItem := engine.anyOfItem_Player_ZoneItem(anyOfItem_Player_ZoneItemID).anyOfItem_Player_ZoneItem
	if anyOfItem_Player_ZoneItem.OperationKind == OperationKindDelete {
		return
	}
	if deleteChild {
		anyOfItem_Player_ZoneItem.deleteChild()
	}
	if _, ok := engine.State.AnyOfItem_Player_ZoneItem[anyOfItem_Player_ZoneItemID]; ok {
		anyOfItem_Player_ZoneItem.OperationKind = OperationKindDelete
		engine.Patch.AnyOfItem_Player_ZoneItem[anyOfItem_Player_ZoneItem.ID] = anyOfItem_Player_ZoneItem
	} else {
		delete(engine.Patch.AnyOfItem_Player_ZoneItem, anyOfItem_Player_ZoneItemID)
	}
}`
const getters_generated_go_import string = `import (
	"sort"
)`
const boolValue_Engine_func string = `func (engine *Engine) boolValue(boolValueID BoolValueID) boolValue {
	patchingBoolValue, ok := engine.Patch.BoolValue[boolValueID]
	if ok {
		return patchingBoolValue
	}
	return engine.State.BoolValue[boolValueID]
}`
const intValue_Engine_func string = `func (engine *Engine) intValue(intValueID IntValueID) intValue {
	patchingIntValue, ok := engine.Patch.IntValue[intValueID]
	if ok {
		return patchingIntValue
	}
	return engine.State.IntValue[intValueID]
}`
const floatValue_Engine_func string = `func (engine *Engine) floatValue(floatValueID FloatValueID) floatValue {
	patchingFloatValue, ok := engine.Patch.FloatValue[floatValueID]
	if ok {
		return patchingFloatValue
	}
	return engine.State.FloatValue[floatValueID]
}`
const stringValue_Engine_func string = `func (engine *Engine) stringValue(stringValueID StringValueID) stringValue {
	patchingStringValue, ok := engine.Patch.StringValue[stringValueID]
	if ok {
		return patchingStringValue
	}
	return engine.State.StringValue[stringValueID]
}`
const _QueryPlayers_Engine_func string = `func (engine *Engine) QueryPlayers(matcher func(Player) bool) []Player {
	playerIDs := engine.allPlayerIDs()
	sort.Slice(playerIDs, func(i, j int) bool {
		return playerIDs[i] < playerIDs[j]
	})
	var players []Player
	for _, playerID := range playerIDs {
		player := engine.Player(playerID)
		if matcher(player) {
			players = append(players, player)
		}
	}
	playerIDSlicePool.Put(playerIDs)
	return players
}`
const _EveryPlayer_Engine_func string = `func (engine *Engine) EveryPlayer() []Player {
	playerIDs := engine.allPlayerIDs()
	sort.Slice(playerIDs, func(i, j int) bool {
		return playerIDs[i] < playerIDs[j]
	})
	var players []Player
	for _, playerID := range playerIDs {
		player := engine.Player(playerID)
		if player.player.HasParent {
			continue
		}
		players = append(players, player)
	}
	playerIDSlicePool.Put(playerIDs)
	return players
}`
const _Player_Engine_func string = `func (engine *Engine) Player(playerID PlayerID) Player {
	patchingPlayer, ok := engine.Patch.Player[playerID]
	if ok {
		return Player{player: patchingPlayer}
	}
	currentPlayer, ok := engine.State.Player[playerID]
	if ok {
		return Player{player: currentPlayer}
	}
	return Player{player: playerCore{OperationKind: OperationKindDelete, engine: engine}}
}`
const _ParentItem_Player_func string = `func (_player Player) ParentItem() Item {
	player := _player.player.engine.Player(_player.player.ID)
	if !player.player.HasParent {
		return Item{item: itemCore{OperationKind: OperationKindDelete, engine: player.player.engine}}
	}
	parentSeg := player.player.Path[len(player.player.Path)-2]
	return player.player.engine.Item(ItemID(parentSeg.ID))
}`
const _ParentZone_Player_func string = `func (_player Player) ParentZone() Zone {
	player := _player.player.engine.Player(_player.player.ID)
	if !player.player.HasParent {
		return Zone{zone: zoneCore{OperationKind: OperationKindDelete, engine: player.player.engine}}
	}
	parentSeg := player.player.Path[len(player.player.Path)-2]
	return player.player.engine.Zone(ZoneID(parentSeg.ID))
}`
const _ParentKind_Player_func string = `func (_player Player) ParentKind() (ElementKind, bool) {
	if !_player.player.HasParent {
		return "", false
	}
	return _player.player.Path[len(_player.player.Path)-2].Kind, true
}`
const _ID_Player_func string = `func (_player Player) ID() PlayerID {
	return _player.player.ID
}`
const _Exists_Player_func string = `func (_player Player) Exists() (Player, bool) {
	player := _player.player.engine.Player(_player.player.ID)
	return player, player.player.OperationKind != OperationKindDelete
}`
const _Path_Player_func string = `func (_player Player) Path() string {
	return _player.player.JSONPath
}`
const _Target_Player_func string = `func (_player Player) Target() PlayerTargetRef {
	player := _player.player.engine.Player(_player.player.ID)
	return player.player.engine.playerTargetRef(player.player.Target)
}`
const _TargetedBy_Player_func string = `func (_player Player) TargetedBy() []PlayerTargetedByRef {
	player := _player.player.engine.Player(_player.player.ID)
	var targetedBy []PlayerTargetedByRef
	sort.Slice(player.player.TargetedBy, func(i, j int) bool {
		return player.player.TargetedBy[i] < player.player.TargetedBy[j]
	})
	for _, refID := range player.player.TargetedBy {
		targetedBy = append(targetedBy, player.player.engine.playerTargetedByRef(refID))
	}
	return targetedBy
}`
const _Action_Player_func string = `func (_player Player) Action() []AttackEvent {
	player := _player.player.engine.Player(_player.player.ID)
	var action []AttackEvent
	sort.Slice(player.player.Action, func(i, j int) bool {
		return player.player.Action[i] < player.player.Action[j]
	})
	for _, attackEventID := range player.player.Action {
		action = append(action, player.player.engine.AttackEvent(attackEventID))
	}
	return action
}`
const _Items_Player_func string = `func (_player Player) Items() []Item {
	player := _player.player.engine.Player(_player.player.ID)
	var items []Item
	sort.Slice(player.player.Items, func(i, j int) bool {
		return player.player.Items[i] < player.player.Items[j]
	})
	for _, itemID := range player.player.Items {
		items = append(items, player.player.engine.Item(itemID))
	}
	return items
}`
const _GearScore_Player_func string = `func (_player Player) GearScore() GearScore {
	player := _player.player.engine.Player(_player.player.ID)
	return player.player.engine.GearScore(player.player.GearScore)
}`
const _GuildMembers_Player_func string = `func (_player Player) GuildMembers() []PlayerGuildMemberRef {
	player := _player.player.engine.Player(_player.player.ID)
	var guildMembers []PlayerGuildMemberRef
	sort.Slice(player.player.GuildMembers, func(i, j int) bool {
		return player.player.GuildMembers[i] < player.player.GuildMembers[j]
	})
	for _, refID := range player.player.GuildMembers {
		guildMembers = append(guildMembers, player.player.engine.playerGuildMemberRef(refID))
	}
	return guildMembers
}`
const _EquipmentSets_Player_func string = `func (_player Player) EquipmentSets() []PlayerEquipmentSetRef {
	player := _player.player.engine.Player(_player.player.ID)
	var equipmentSets []PlayerEquipmentSetRef
	sort.Slice(player.player.EquipmentSets, func(i, j int) bool {
		return player.player.EquipmentSets[i] < player.player.EquipmentSets[j]
	})
	for _, refID := range player.player.EquipmentSets {
		equipmentSets = append(equipmentSets, player.player.engine.playerEquipmentSetRef(refID))
	}
	return equipmentSets
}`
const _Position_Player_func string = `func (_player Player) Position() Position {
	player := _player.player.engine.Player(_player.player.ID)
	return player.player.engine.Position(player.player.Position)
}`
const _QueryGearScores_Engine_func string = `func (engine *Engine) QueryGearScores(matcher func(GearScore) bool) []GearScore {
	gearScoreIDs := engine.allGearScoreIDs()
	sort.Slice(gearScoreIDs, func(i, j int) bool {
		return gearScoreIDs[i] < gearScoreIDs[j]
	})
	var gearScores []GearScore
	for _, gearScoreID := range gearScoreIDs {
		gearScore := engine.GearScore(gearScoreID)
		if matcher(gearScore) {
			gearScores = append(gearScores, gearScore)
		}
	}
	gearScoreIDSlicePool.Put(gearScoreIDs)
	return gearScores
}`
const _EveryGearScore_Engine_func string = `func (engine *Engine) EveryGearScore() []GearScore {
	gearScoreIDs := engine.allGearScoreIDs()
	sort.Slice(gearScoreIDs, func(i, j int) bool {
		return gearScoreIDs[i] < gearScoreIDs[j]
	})
	var gearScores []GearScore
	for _, gearScoreID := range gearScoreIDs {
		gearScore := engine.GearScore(gearScoreID)
		if gearScore.gearScore.HasParent {
			continue
		}
		gearScores = append(gearScores, gearScore)
	}
	gearScoreIDSlicePool.Put(gearScoreIDs)
	return gearScores
}`
const _GearScore_Engine_func string = `func (engine *Engine) GearScore(gearScoreID GearScoreID) GearScore {
	patchingGearScore, ok := engine.Patch.GearScore[gearScoreID]
	if ok {
		return GearScore{gearScore: patchingGearScore}
	}
	currentGearScore, ok := engine.State.GearScore[gearScoreID]
	if ok {
		return GearScore{gearScore: currentGearScore}
	}
	return GearScore{gearScore: gearScoreCore{OperationKind: OperationKindDelete, engine: engine}}
}`
const _ParentItem_GearScore_func string = `func (_gearScore GearScore) ParentItem() Item {
	gearScore := _gearScore.gearScore.engine.GearScore(_gearScore.gearScore.ID)
	if !gearScore.gearScore.HasParent {
		return Item{item: itemCore{OperationKind: OperationKindDelete, engine: gearScore.gearScore.engine}}
	}
	parentSeg := gearScore.gearScore.Path[len(gearScore.gearScore.Path)-2]
	return gearScore.gearScore.engine.Item(ItemID(parentSeg.ID))
}`
const _ParentPlayer_GearScore_func string = `func (_gearScore GearScore) ParentPlayer() Player {
	gearScore := _gearScore.gearScore.engine.GearScore(_gearScore.gearScore.ID)
	if !gearScore.gearScore.HasParent {
		return Player{player: playerCore{OperationKind: OperationKindDelete, engine: gearScore.gearScore.engine}}
	}
	parentSeg := gearScore.gearScore.Path[len(gearScore.gearScore.Path)-2]
	return gearScore.gearScore.engine.Player(PlayerID(parentSeg.ID))
}`
const _ParentKind_GearScore_func string = `func (_gearScore GearScore) ParentKind() (ElementKind, bool) {
	if !_gearScore.gearScore.HasParent {
		return "", false
	}
	return _gearScore.gearScore.Path[len(_gearScore.gearScore.Path)-2].Kind, true
}`
const _ID_GearScore_func string = `func (_gearScore GearScore) ID() GearScoreID {
	return _gearScore.gearScore.ID
}`
const _Exists_GearScore_func string = `func (_gearScore GearScore) Exists() (GearScore, bool) {
	gearScore := _gearScore.gearScore.engine.GearScore(_gearScore.gearScore.ID)
	return gearScore, gearScore.gearScore.OperationKind != OperationKindDelete
}`
const _Path_GearScore_func string = `func (_gearScore GearScore) Path() string {
	return _gearScore.gearScore.JSONPath
}`
const _Level_GearScore_func string = `func (_gearScore GearScore) Level() int64 {
	gearScore := _gearScore.gearScore.engine.GearScore(_gearScore.gearScore.ID)
	return gearScore.gearScore.engine.intValue(gearScore.gearScore.Level).Value
}`
const _Score_GearScore_func string = `func (_gearScore GearScore) Score() int64 {
	gearScore := _gearScore.gearScore.engine.GearScore(_gearScore.gearScore.ID)
	return gearScore.gearScore.engine.intValue(gearScore.gearScore.Score).Value
}`
const _QueryItems_Engine_func string = `func (engine *Engine) QueryItems(matcher func(Item) bool) []Item {
	itemIDs := engine.allItemIDs()
	sort.Slice(itemIDs, func(i, j int) bool {
		return itemIDs[i] < itemIDs[j]
	})
	var items []Item
	for _, itemID := range itemIDs {
		item := engine.Item(itemID)
		if matcher(item) {
			items = append(items, item)
		}
	}
	itemIDSlicePool.Put(itemIDs)
	return items
}`
const _EveryItem_Engine_func string = `func (engine *Engine) EveryItem() []Item {
	itemIDs := engine.allItemIDs()
	sort.Slice(itemIDs, func(i, j int) bool {
		return itemIDs[i] < itemIDs[j]
	})
	var items []Item
	for _, itemID := range itemIDs {
		item := engine.Item(itemID)
		if item.item.HasParent {
			continue
		}
		items = append(items, item)
	}
	itemIDSlicePool.Put(itemIDs)
	return items
}`
const _Item_Engine_func string = `func (engine *Engine) Item(itemID ItemID) Item {
	patchingItem, ok := engine.Patch.Item[itemID]
	if ok {
		return Item{item: patchingItem}
	}
	currentItem, ok := engine.State.Item[itemID]
	if ok {
		return Item{item: currentItem}
	}
	return Item{item: itemCore{OperationKind: OperationKindDelete, engine: engine}}
}`
const _ParentPlayer_Item_func string = `func (_item Item) ParentPlayer() Player {
	item := _item.item.engine.Item(_item.item.ID)
	if !item.item.HasParent {
		return Player{player: playerCore{OperationKind: OperationKindDelete, engine: item.item.engine}}
	}
	parentSeg := item.item.Path[len(item.item.Path)-2]
	return item.item.engine.Player(PlayerID(parentSeg.ID))
}`
const _ParentZone_Item_func string = `func (_item Item) ParentZone() Zone {
	item := _item.item.engine.Item(_item.item.ID)
	if !item.item.HasParent {
		return Zone{zone: zoneCore{OperationKind: OperationKindDelete, engine: item.item.engine}}
	}
	parentSeg := item.item.Path[len(item.item.Path)-2]
	return item.item.engine.Zone(ZoneID(parentSeg.ID))
}`
const _ParentZoneItem_Item_func string = `func (_item Item) ParentZoneItem() ZoneItem {
	item := _item.item.engine.Item(_item.item.ID)
	if !item.item.HasParent {
		return ZoneItem{zoneItem: zoneItemCore{OperationKind: OperationKindDelete, engine: item.item.engine}}
	}
	parentSeg := item.item.Path[len(item.item.Path)-2]
	return item.item.engine.ZoneItem(ZoneItemID(parentSeg.ID))
}`
const _ParentKind_Item_func string = `func (_item Item) ParentKind() (ElementKind, bool) {
	if !_item.item.HasParent {
		return "", false
	}
	return _item.item.Path[len(_item.item.Path)-2].Kind, true
}`
const _ID_Item_func string = `func (_item Item) ID() ItemID {
	return _item.item.ID
}`
const _Exists_Item_func string = `func (_item Item) Exists() (Item, bool) {
	item := _item.item.engine.Item(_item.item.ID)
	return item, item.item.OperationKind != OperationKindDelete
}`
const _Path_Item_func string = `func (_item Item) Path() string {
	return _item.item.JSONPath
}`
const _Name_Item_func string = `func (_item Item) Name() string {
	item := _item.item.engine.Item(_item.item.ID)
	return item.item.engine.stringValue(item.item.Name).Value
}`
const _GearScore_Item_func string = `func (_item Item) GearScore() GearScore {
	item := _item.item.engine.Item(_item.item.ID)
	return item.item.engine.GearScore(item.item.GearScore)
}`
const _BoundTo_Item_func string = `func (_item Item) BoundTo() ItemBoundToRef {
	item := _item.item.engine.Item(_item.item.ID)
	return item.item.engine.itemBoundToRef(item.item.BoundTo)
}`
const _Origin_Item_func string = `func (_item Item) Origin() AnyOfPlayer_Position {
	item := _item.item.engine.Item(_item.item.ID)
	return item.item.engine.anyOfPlayer_Position(item.item.Origin)
}`
const _QueryAttackEvents_Engine_func string = `func (engine *Engine) QueryAttackEvents(matcher func(AttackEvent) bool) []AttackEvent {
	attackEventIDs := engine.allAttackEventIDs()
	sort.Slice(attackEventIDs, func(i, j int) bool {
		return attackEventIDs[i] < attackEventIDs[j]
	})
	var attackEvents []AttackEvent
	for _, attackEventID := range attackEventIDs {
		attackEvent := engine.AttackEvent(attackEventID)
		if matcher(attackEvent) {
			attackEvents = append(attackEvents, attackEvent)
		}
	}
	attackEventIDSlicePool.Put(attackEventIDs)
	return attackEvents
}`
const _EveryAttackEvent_Engine_func string = `func (engine *Engine) EveryAttackEvent() []AttackEvent {
	attackEventIDs := engine.allAttackEventIDs()
	sort.Slice(attackEventIDs, func(i, j int) bool {
		return attackEventIDs[i] < attackEventIDs[j]
	})
	var attackEvents []AttackEvent
	for _, attackEventID := range attackEventIDs {
		attackEvent := engine.AttackEvent(attackEventID)
		if attackEvent.attackEvent.HasParent {
			continue
		}
		attackEvents = append(attackEvents, attackEvent)
	}
	attackEventIDSlicePool.Put(attackEventIDs)
	return attackEvents
}`
const _AttackEvent_Engine_func string = `func (engine *Engine) AttackEvent(attackEventID AttackEventID) AttackEvent {
	patchingAttackEvent, ok := engine.Patch.AttackEvent[attackEventID]
	if ok {
		return AttackEvent{attackEvent: patchingAttackEvent}
	}
	currentAttackEvent, ok := engine.State.AttackEvent[attackEventID]
	if ok {
		return AttackEvent{attackEvent: currentAttackEvent}
	}
	return AttackEvent{attackEvent: attackEventCore{OperationKind: OperationKindDelete, engine: engine}}
}`
const _ParentPlayer_AttackEvent_func string = `func (_attackEvent AttackEvent) ParentPlayer() Player {
	attackEvent := _attackEvent.attackEvent.engine.AttackEvent(_attackEvent.attackEvent.ID)
	if !attackEvent.attackEvent.HasParent {
		return Player{player: playerCore{OperationKind: OperationKindDelete, engine: attackEvent.attackEvent.engine}}
	}
	parentSeg := attackEvent.attackEvent.Path[len(attackEvent.attackEvent.Path)-2]
	return attackEvent.attackEvent.engine.Player(PlayerID(parentSeg.ID))
}`
const _ParentKind_AttackEvent_func string = `func (_attackEvent AttackEvent) ParentKind() (ElementKind, bool) {
	if !_attackEvent.attackEvent.HasParent {
		return "", false
	}
	return _attackEvent.attackEvent.Path[len(_attackEvent.attackEvent.Path)-2].Kind, true
}`
const _ID_AttackEvent_func string = `func (_attackEvent AttackEvent) ID() AttackEventID {
	return _attackEvent.attackEvent.ID
}`
const _Exists_AttackEvent_func string = `func (_attackEvent AttackEvent) Exists() (AttackEvent, bool) {
	attackEvent := _attackEvent.attackEvent.engine.AttackEvent(_attackEvent.attackEvent.ID)
	return attackEvent, attackEvent.attackEvent.OperationKind != OperationKindDelete
}`
const _Path_AttackEvent_func string = `func (_attackEvent AttackEvent) Path() string {
	return _attackEvent.attackEvent.JSONPath
}`
const _Target_AttackEvent_func string = `func (_attackEvent AttackEvent) Target() AttackEventTargetRef {
	attackEvent := _attackEvent.attackEvent.engine.AttackEvent(_attackEvent.attackEvent.ID)
	return attackEvent.attackEvent.engine.attackEventTargetRef(attackEvent.attackEvent.Target)
}`
const _QueryPositions_Engine_func string = `func (engine *Engine) QueryPositions(matcher func(Position) bool) []Position {
	positionIDs := engine.allPositionIDs()
	sort.Slice(positionIDs, func(i, j int) bool {
		return positionIDs[i] < positionIDs[j]
	})
	var positions []Position
	for _, positionID := range positionIDs {
		position := engine.Position(positionID)
		if matcher(position) {
			positions = append(positions, position)
		}
	}
	positionIDSlicePool.Put(positionIDs)
	return positions
}`
const _EveryPosition_Engine_func string = `func (engine *Engine) EveryPosition() []Position {
	positionIDs := engine.allPositionIDs()
	sort.Slice(positionIDs, func(i, j int) bool {
		return positionIDs[i] < positionIDs[j]
	})
	var positions []Position
	for _, positionID := range positionIDs {
		position := engine.Position(positionID)
		if position.position.HasParent {
			continue
		}
		positions = append(positions, position)
	}
	positionIDSlicePool.Put(positionIDs)
	return positions
}`
const _Position_Engine_func string = `func (engine *Engine) Position(positionID PositionID) Position {
	patchingPosition, ok := engine.Patch.Position[positionID]
	if ok {
		return Position{position: patchingPosition}
	}
	currentPosition, ok := engine.State.Position[positionID]
	if ok {
		return Position{position: currentPosition}
	}
	return Position{position: positionCore{OperationKind: OperationKindDelete, engine: engine}}
}`
const _ParentItem_Position_func string = `func (_position Position) ParentItem() Item {
	position := _position.position.engine.Position(_position.position.ID)
	if !position.position.HasParent {
		return Item{item: itemCore{OperationKind: OperationKindDelete, engine: position.position.engine}}
	}
	parentSeg := position.position.Path[len(position.position.Path)-2]
	return position.position.engine.Item(ItemID(parentSeg.ID))
}`
const _ParentPlayer_Position_func string = `func (_position Position) ParentPlayer() Player {
	position := _position.position.engine.Position(_position.position.ID)
	if !position.position.HasParent {
		return Player{player: playerCore{OperationKind: OperationKindDelete, engine: position.position.engine}}
	}
	parentSeg := position.position.Path[len(position.position.Path)-2]
	return position.position.engine.Player(PlayerID(parentSeg.ID))
}`
const _ParentZoneItem_Position_func string = `func (_position Position) ParentZoneItem() ZoneItem {
	position := _position.position.engine.Position(_position.position.ID)
	if !position.position.HasParent {
		return ZoneItem{zoneItem: zoneItemCore{OperationKind: OperationKindDelete, engine: position.position.engine}}
	}
	parentSeg := position.position.Path[len(position.position.Path)-2]
	return position.position.engine.ZoneItem(ZoneItemID(parentSeg.ID))
}`
const _Exists_Position_func string = `func (_position Position) Exists() (Position, bool) {
	position := _position.position.engine.Position(_position.position.ID)
	return position, position.position.OperationKind != OperationKindDelete
}`
const _ParentKind_Position_func string = `func (_position Position) ParentKind() (ElementKind, bool) {
	if !_position.position.HasParent {
		return "", false
	}
	return _position.position.Path[len(_position.position.Path)-2].Kind, true
}`
const _ID_Position_func string = `func (_position Position) ID() PositionID {
	return _position.position.ID
}`
const _Path_Position_func string = `func (_position Position) Path() string {
	return _position.position.JSONPath
}`
const _X_Position_func string = `func (_position Position) X() float64 {
	position := _position.position.engine.Position(_position.position.ID)
	return position.position.engine.floatValue(position.position.X).Value
}`
const _Y_Position_func string = `func (_position Position) Y() float64 {
	position := _position.position.engine.Position(_position.position.ID)
	return position.position.engine.floatValue(position.position.Y).Value
}`
const _QueryZoneItems_Engine_func string = `func (engine *Engine) QueryZoneItems(matcher func(ZoneItem) bool) []ZoneItem {
	zoneItemIDs := engine.allZoneItemIDs()
	sort.Slice(zoneItemIDs, func(i, j int) bool {
		return zoneItemIDs[i] < zoneItemIDs[j]
	})
	var zoneItems []ZoneItem
	for _, zoneItemID := range zoneItemIDs {
		zoneItem := engine.ZoneItem(zoneItemID)
		if matcher(zoneItem) {
			zoneItems = append(zoneItems, zoneItem)
		}
	}
	zoneItemIDSlicePool.Put(zoneItemIDs)
	return zoneItems
}`
const _EveryZoneItem_Engine_func string = `func (engine *Engine) EveryZoneItem() []ZoneItem {
	zoneItemIDs := engine.allZoneItemIDs()
	sort.Slice(zoneItemIDs, func(i, j int) bool {
		return zoneItemIDs[i] < zoneItemIDs[j]
	})
	var zoneItems []ZoneItem
	for _, zoneItemID := range zoneItemIDs {
		zoneItem := engine.ZoneItem(zoneItemID)
		if zoneItem.zoneItem.HasParent {
			continue
		}
		zoneItems = append(zoneItems, zoneItem)
	}
	zoneItemIDSlicePool.Put(zoneItemIDs)
	return zoneItems
}`
const _ZoneItem_Engine_func string = `func (engine *Engine) ZoneItem(zoneItemID ZoneItemID) ZoneItem {
	patchingZoneItem, ok := engine.Patch.ZoneItem[zoneItemID]
	if ok {
		return ZoneItem{zoneItem: patchingZoneItem}
	}
	currentZoneItem, ok := engine.State.ZoneItem[zoneItemID]
	if ok {
		return ZoneItem{zoneItem: currentZoneItem}
	}
	return ZoneItem{zoneItem: zoneItemCore{OperationKind: OperationKindDelete, engine: engine}}
}`
const _ParentZone_ZoneItem_func string = `func (_zoneItem ZoneItem) ParentZone() Zone {
	zoneItem := _zoneItem.zoneItem.engine.ZoneItem(_zoneItem.zoneItem.ID)
	if !zoneItem.zoneItem.HasParent {
		return Zone{zone: zoneCore{OperationKind: OperationKindDelete, engine: zoneItem.zoneItem.engine}}
	}
	parentSeg := zoneItem.zoneItem.Path[len(zoneItem.zoneItem.Path)-2]
	return zoneItem.zoneItem.engine.Zone(ZoneID(parentSeg.ID))
}`
const _ParentKind_ZoneItem_func string = `func (_zoneItem ZoneItem) ParentKind() (ElementKind, bool) {
	if !_zoneItem.zoneItem.HasParent {
		return "", false
	}
	return _zoneItem.zoneItem.Path[len(_zoneItem.zoneItem.Path)-2].Kind, true
}`
const _ID_ZoneItem_func string = `func (_zoneItem ZoneItem) ID() ZoneItemID {
	return _zoneItem.zoneItem.ID
}`
const _Exists_ZoneItem_func string = `func (_zoneItem ZoneItem) Exists() (ZoneItem, bool) {
	zoneItem := _zoneItem.zoneItem.engine.ZoneItem(_zoneItem.zoneItem.ID)
	return zoneItem, zoneItem.zoneItem.OperationKind != OperationKindDelete
}`
const _Path_ZoneItem_func string = `func (_zoneItem ZoneItem) Path() string {
	return _zoneItem.zoneItem.JSONPath
}`
const _Position_ZoneItem_func string = `func (_zoneItem ZoneItem) Position() Position {
	zoneItem := _zoneItem.zoneItem.engine.ZoneItem(_zoneItem.zoneItem.ID)
	return zoneItem.zoneItem.engine.Position(zoneItem.zoneItem.Position)
}`
const _Item_ZoneItem_func string = `func (_zoneItem ZoneItem) Item() Item {
	zoneItem := _zoneItem.zoneItem.engine.ZoneItem(_zoneItem.zoneItem.ID)
	return zoneItem.zoneItem.engine.Item(zoneItem.zoneItem.Item)
}`
const _QueryZones_Engine_func string = `func (engine *Engine) QueryZones(matcher func(Zone) bool) []Zone {
	zoneIDs := engine.allZoneIDs()
	sort.Slice(zoneIDs, func(i, j int) bool {
		return zoneIDs[i] < zoneIDs[j]
	})
	var zones []Zone
	for _, zoneID := range zoneIDs {
		zone := engine.Zone(zoneID)
		if matcher(zone) {
			zones = append(zones, zone)
		}
	}
	zoneIDSlicePool.Put(zoneIDs)
	return zones
}`
const _EveryZone_Engine_func string = `func (engine *Engine) EveryZone() []Zone {
	zoneIDs := engine.allZoneIDs()
	sort.Slice(zoneIDs, func(i, j int) bool {
		return zoneIDs[i] < zoneIDs[j]
	})
	var zones []Zone
	for _, zoneID := range zoneIDs {
		zone := engine.Zone(zoneID)
		zones = append(zones, zone)
	}
	zoneIDSlicePool.Put(zoneIDs)
	return zones
}`
const _Zone_Engine_func string = `func (engine *Engine) Zone(zoneID ZoneID) Zone {
	patchingZone, ok := engine.Patch.Zone[zoneID]
	if ok {
		return Zone{zone: patchingZone}
	}
	currentZone, ok := engine.State.Zone[zoneID]
	if ok {
		return Zone{zone: currentZone}
	}
	return Zone{zone: zoneCore{OperationKind: OperationKindDelete, engine: engine}}
}`
const _ID_Zone_func string = `func (_zone Zone) ID() ZoneID {
	return _zone.zone.ID
}`
const _Exists_Zone_func string = `func (_zone Zone) Exists() (Zone, bool) {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	return zone, zone.zone.OperationKind != OperationKindDelete
}`
const _Path_Zone_func string = `func (_zone Zone) Path() string {
	return _zone.zone.JSONPath
}`
const _Players_Zone_func string = `func (_zone Zone) Players() []Player {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	var players []Player
	sort.Slice(zone.zone.Players, func(i, j int) bool {
		return zone.zone.Players[i] < zone.zone.Players[j]
	})
	for _, playerID := range zone.zone.Players {
		players = append(players, zone.zone.engine.Player(playerID))
	}
	return players
}`
const _Interactables_Zone_func string = `func (_zone Zone) Interactables() []AnyOfItem_Player_ZoneItemSliceElement {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	var interactables []AnyOfItem_Player_ZoneItemSliceElement
	sort.Slice(zone.zone.Interactables, func(i, j int) bool {
		return zone.zone.Interactables[i] < zone.zone.Interactables[j]
	})
	for _, anyOfItem_Player_ZoneItemID := range zone.zone.Interactables {
		interactables = append(interactables, zone.zone.engine.anyOfItem_Player_ZoneItem(anyOfItem_Player_ZoneItemID))
	}
	return interactables
}`
const _Items_Zone_func string = `func (_zone Zone) Items() []ZoneItem {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	var items []ZoneItem
	sort.Slice(zone.zone.Items, func(i, j int) bool {
		return zone.zone.Items[i] < zone.zone.Items[j]
	})
	for _, zoneItemID := range zone.zone.Items {
		items = append(items, zone.zone.engine.ZoneItem(zoneItemID))
	}
	return items
}`
const _Tags_Zone_func string = `func (_zone Zone) Tags() []string {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	var tags []string
	sort.Slice(zone.zone.Tags, func(i, j int) bool {
		return zone.zone.Tags[i] < zone.zone.Tags[j]
	})
	for _, stringValueID := range zone.zone.Tags {
		tags = append(tags, zone.zone.engine.stringValue(stringValueID).Value)
	}
	return tags
}`
const _ID_ItemBoundToRef_func string = `func (_itemBoundToRef ItemBoundToRef) ID() PlayerID {
	return _itemBoundToRef.itemBoundToRef.ReferencedElementID
}`
const itemBoundToRef_Engine_func string = `func (engine *Engine) itemBoundToRef(itemBoundToRefID ItemBoundToRefID) ItemBoundToRef {
	patchingItemBoundToRef, ok := engine.Patch.ItemBoundToRef[itemBoundToRefID]
	if ok {
		return ItemBoundToRef{itemBoundToRef: patchingItemBoundToRef}
	}
	currentItemBoundToRef, ok := engine.State.ItemBoundToRef[itemBoundToRefID]
	if ok {
		return ItemBoundToRef{itemBoundToRef: currentItemBoundToRef}
	}
	return ItemBoundToRef{itemBoundToRef: itemBoundToRefCore{OperationKind: OperationKindDelete, engine: engine}}
}`
const _ID_AttackEventTargetRef_func string = `func (_attackEventTargetRef AttackEventTargetRef) ID() PlayerID {
	return _attackEventTargetRef.attackEventTargetRef.ReferencedElementID
}`
const attackEventTargetRef_Engine_func string = `func (engine *Engine) attackEventTargetRef(attackEventTargetRefID AttackEventTargetRefID) AttackEventTargetRef {
	patchingAttackEventTargetRef, ok := engine.Patch.AttackEventTargetRef[attackEventTargetRefID]
	if ok {
		return AttackEventTargetRef{attackEventTargetRef: patchingAttackEventTargetRef}
	}
	currentAttackEventTargetRef, ok := engine.State.AttackEventTargetRef[attackEventTargetRefID]
	if ok {
		return AttackEventTargetRef{attackEventTargetRef: currentAttackEventTargetRef}
	}
	return AttackEventTargetRef{attackEventTargetRef: attackEventTargetRefCore{OperationKind: OperationKindDelete, engine: engine}}
}`
const _ID_PlayerGuildMemberRef_func string = `func (_playerGuildMemberRef PlayerGuildMemberRef) ID() PlayerID {
	return _playerGuildMemberRef.playerGuildMemberRef.ReferencedElementID
}`
const playerGuildMemberRef_Engine_func string = `func (engine *Engine) playerGuildMemberRef(playerGuildMemberRefID PlayerGuildMemberRefID) PlayerGuildMemberRef {
	patchingPlayerGuildMemberRef, ok := engine.Patch.PlayerGuildMemberRef[playerGuildMemberRefID]
	if ok {
		return PlayerGuildMemberRef{playerGuildMemberRef: patchingPlayerGuildMemberRef}
	}
	currentPlayerGuildMemberRef, ok := engine.State.PlayerGuildMemberRef[playerGuildMemberRefID]
	if ok {
		return PlayerGuildMemberRef{playerGuildMemberRef: currentPlayerGuildMemberRef}
	}
	return PlayerGuildMemberRef{playerGuildMemberRef: playerGuildMemberRefCore{OperationKind: OperationKindDelete, engine: engine}}
}`
const _ID_PlayerEquipmentSetRef_func string = `func (_playerEquipmentSetRef PlayerEquipmentSetRef) ID() EquipmentSetID {
	return _playerEquipmentSetRef.playerEquipmentSetRef.ReferencedElementID
}`
const _QueryEquipmentSets_Engine_func string = `func (engine *Engine) QueryEquipmentSets(matcher func(EquipmentSet) bool) []EquipmentSet {
	equipmentSetIDs := engine.allEquipmentSetIDs()
	sort.Slice(equipmentSetIDs, func(i, j int) bool {
		return equipmentSetIDs[i] < equipmentSetIDs[j]
	})
	var equipmentSets []EquipmentSet
	for _, equipmentSetID := range equipmentSetIDs {
		equipmentSet := engine.EquipmentSet(equipmentSetID)
		if matcher(equipmentSet) {
			equipmentSets = append(equipmentSets, equipmentSet)
		}
	}
	equipmentSetIDSlicePool.Put(equipmentSetIDs)
	return equipmentSets
}`
const _EveryEquipmentSet_Engine_func string = `func (engine *Engine) EveryEquipmentSet() []EquipmentSet {
	equipmentSetIDs := engine.allEquipmentSetIDs()
	sort.Slice(equipmentSetIDs, func(i, j int) bool {
		return equipmentSetIDs[i] < equipmentSetIDs[j]
	})
	var equipmentSets []EquipmentSet
	for _, equipmentSetID := range equipmentSetIDs {
		equipmentSet := engine.EquipmentSet(equipmentSetID)
		equipmentSets = append(equipmentSets, equipmentSet)
	}
	equipmentSetIDSlicePool.Put(equipmentSetIDs)
	return equipmentSets
}`
const _EquipmentSet_Engine_func string = `func (engine *Engine) EquipmentSet(equipmentSetID EquipmentSetID) EquipmentSet {
	patchingEquipmentSet, ok := engine.Patch.EquipmentSet[equipmentSetID]
	if ok {
		return EquipmentSet{equipmentSet: patchingEquipmentSet}
	}
	currentEquipmentSet, ok := engine.State.EquipmentSet[equipmentSetID]
	if ok {
		return EquipmentSet{equipmentSet: currentEquipmentSet}
	}
	return EquipmentSet{equipmentSet: equipmentSetCore{OperationKind: OperationKindDelete, engine: engine}}
}`
const _ID_EquipmentSet_func string = `func (_equipmentSet EquipmentSet) ID() EquipmentSetID {
	return _equipmentSet.equipmentSet.ID
}`
const _Exists_EquipmentSet_func string = `func (_equipmentSet EquipmentSet) Exists() (EquipmentSet, bool) {
	equipmentSet := _equipmentSet.equipmentSet.engine.EquipmentSet(_equipmentSet.equipmentSet.ID)
	return equipmentSet, equipmentSet.equipmentSet.OperationKind != OperationKindDelete
}`
const _Path_EquipmentSet_func string = `func (_equipmentSet EquipmentSet) Path() string {
	return _equipmentSet.equipmentSet.JSONPath
}`
const _Name_EquipmentSet_func string = `func (_equipmentSet EquipmentSet) Name() string {
	equipmentSet := _equipmentSet.equipmentSet.engine.EquipmentSet(_equipmentSet.equipmentSet.ID)
	return equipmentSet.equipmentSet.engine.stringValue(equipmentSet.equipmentSet.Name).Value
}`
const _Equipment_EquipmentSet_func string = `func (_equipmentSet EquipmentSet) Equipment() []EquipmentSetEquipmentRef {
	equipmentSet := _equipmentSet.equipmentSet.engine.EquipmentSet(_equipmentSet.equipmentSet.ID)
	var equipment []EquipmentSetEquipmentRef
	sort.Slice(equipmentSet.equipmentSet.Equipment, func(i, j int) bool {
		return equipmentSet.equipmentSet.Equipment[i] < equipmentSet.equipmentSet.Equipment[j]
	})
	for _, refID := range equipmentSet.equipmentSet.Equipment {
		equipment = append(equipment, equipmentSet.equipmentSet.engine.equipmentSetEquipmentRef(refID))
	}
	return equipment
}`
const playerEquipmentSetRef_Engine_func string = `func (engine *Engine) playerEquipmentSetRef(playerEquipmentSetRefID PlayerEquipmentSetRefID) PlayerEquipmentSetRef {
	patchingPlayerEquipmentSetRef, ok := engine.Patch.PlayerEquipmentSetRef[playerEquipmentSetRefID]
	if ok {
		return PlayerEquipmentSetRef{playerEquipmentSetRef: patchingPlayerEquipmentSetRef}
	}
	currentPlayerEquipmentSetRef, ok := engine.State.PlayerEquipmentSetRef[playerEquipmentSetRefID]
	if ok {
		return PlayerEquipmentSetRef{playerEquipmentSetRef: currentPlayerEquipmentSetRef}
	}
	return PlayerEquipmentSetRef{playerEquipmentSetRef: playerEquipmentSetRefCore{OperationKind: OperationKindDelete, engine: engine}}
}`
const _ID_EquipmentSetEquipmentRef_func string = `func (_equipmentSetEquipmentRef EquipmentSetEquipmentRef) ID() ItemID {
	return _equipmentSetEquipmentRef.equipmentSetEquipmentRef.ReferencedElementID
}`
const equipmentSetEquipmentRef_Engine_func string = `func (engine *Engine) equipmentSetEquipmentRef(equipmentSetEquipmentRefID EquipmentSetEquipmentRefID) EquipmentSetEquipmentRef {
	patchingEquipmentSetEquipmentRef, ok := engine.Patch.EquipmentSetEquipmentRef[equipmentSetEquipmentRefID]
	if ok {
		return EquipmentSetEquipmentRef{equipmentSetEquipmentRef: patchingEquipmentSetEquipmentRef}
	}
	currentEquipmentSetEquipmentRef, ok := engine.State.EquipmentSetEquipmentRef[equipmentSetEquipmentRefID]
	if ok {
		return EquipmentSetEquipmentRef{equipmentSetEquipmentRef: currentEquipmentSetEquipmentRef}
	}
	return EquipmentSetEquipmentRef{equipmentSetEquipmentRef: equipmentSetEquipmentRefCore{OperationKind: OperationKindDelete, engine: engine}}
}`
const _ID_PlayerTargetRef_func string = `func (_playerTargetRef PlayerTargetRef) ID() AnyOfPlayer_ZoneItemID {
	return _playerTargetRef.playerTargetRef.ReferencedElementID
}`
const playerTargetRef_Engine_func string = `func (engine *Engine) playerTargetRef(playerTargetRefID PlayerTargetRefID) PlayerTargetRef {
	patchingPlayerTargetRef, ok := engine.Patch.PlayerTargetRef[playerTargetRefID]
	if ok {
		return PlayerTargetRef{playerTargetRef: patchingPlayerTargetRef}
	}
	currentPlayerTargetRef, ok := engine.State.PlayerTargetRef[playerTargetRefID]
	if ok {
		return PlayerTargetRef{playerTargetRef: currentPlayerTargetRef}
	}
	return PlayerTargetRef{playerTargetRef: playerTargetRefCore{OperationKind: OperationKindDelete, engine: engine}}
}`
const _ID_PlayerTargetedByRef_func string = `func (_playerTargetedByRef PlayerTargetedByRef) ID() AnyOfPlayer_ZoneItemID {
	return _playerTargetedByRef.playerTargetedByRef.ReferencedElementID
}`
const playerTargetedByRef_Engine_func string = `func (engine *Engine) playerTargetedByRef(playerTargetedByRefID PlayerTargetedByRefID) PlayerTargetedByRef {
	patchingPlayerTargetedByRef, ok := engine.Patch.PlayerTargetedByRef[playerTargetedByRefID]
	if ok {
		return PlayerTargetedByRef{playerTargetedByRef: patchingPlayerTargetedByRef}
	}
	currentPlayerTargetedByRef, ok := engine.State.PlayerTargetedByRef[playerTargetedByRefID]
	if ok {
		return PlayerTargetedByRef{playerTargetedByRef: currentPlayerTargetedByRef}
	}
	return PlayerTargetedByRef{playerTargetedByRef: playerTargetedByRefCore{OperationKind: OperationKindDelete, engine: engine}}
}`
const _ID_AnyOfPlayer_Position_func string = `func (_anyOfPlayer_Position AnyOfPlayer_Position) ID() AnyOfPlayer_PositionID {
	return _anyOfPlayer_Position.anyOfPlayer_Position.ID
}`
const _Player_AnyOfPlayer_Position_func string = `func (_anyOfPlayer_Position AnyOfPlayer_Position) Player() Player {
	anyOfPlayer_Position := _anyOfPlayer_Position.anyOfPlayer_Position.engine.anyOfPlayer_Position(_anyOfPlayer_Position.anyOfPlayer_Position.ID)
	return anyOfPlayer_Position.anyOfPlayer_Position.engine.Player(PlayerID(anyOfPlayer_Position.anyOfPlayer_Position.ChildID))
}`
const _Position_AnyOfPlayer_Position_func string = `func (_anyOfPlayer_Position AnyOfPlayer_Position) Position() Position {
	anyOfPlayer_Position := _anyOfPlayer_Position.anyOfPlayer_Position.engine.anyOfPlayer_Position(_anyOfPlayer_Position.anyOfPlayer_Position.ID)
	return anyOfPlayer_Position.anyOfPlayer_Position.engine.Position(PositionID(anyOfPlayer_Position.anyOfPlayer_Position.ChildID))
}`
const anyOfPlayer_Position_Engine_func string = `func (engine *Engine) anyOfPlayer_Position(anyOfPlayer_PositionID AnyOfPlayer_PositionID) AnyOfPlayer_Position {
	patchingAnyOfPlayer_Position, ok := engine.Patch.AnyOfPlayer_Position[anyOfPlayer_PositionID]
	if ok {
		return AnyOfPlayer_Position{anyOfPlayer_Position: patchingAnyOfPlayer_Position}
	}
	currentAnyOfPlayer_Position, ok := engine.State.AnyOfPlayer_Position[anyOfPlayer_PositionID]
	if ok {
		return AnyOfPlayer_Position{anyOfPlayer_Position: currentAnyOfPlayer_Position}
	}
	return AnyOfPlayer_Position{anyOfPlayer_Position: anyOfPlayer_PositionCore{OperationKind: OperationKindDelete, engine: engine}}
}`
const _ID_AnyOfPlayer_ZoneItem_func string = `func (_anyOfPlayer_ZoneItem AnyOfPlayer_ZoneItem) ID() AnyOfPlayer_ZoneItemID {
	return _anyOfPlayer_ZoneItem.anyOfPlayer_ZoneItem.ID
}`
const _Player_AnyOfPlayer_ZoneItem_func string = `func (_anyOfPlayer_ZoneItem AnyOfPlayer_ZoneItem) Player() Player {
	anyOfPlayer_ZoneItem := _anyOfPlayer_ZoneItem.anyOfPlayer_ZoneItem.engine.anyOfPlayer_ZoneItem(_anyOfPlayer_ZoneItem.anyOfPlayer_ZoneItem.ID)
	return anyOfPlayer_ZoneItem.anyOfPlayer_ZoneItem.engine.Player(PlayerID(anyOfPlayer_ZoneItem.anyOfPlayer_ZoneItem.ChildID))
}`
const _ZoneItem_AnyOfPlayer_ZoneItem_func string = `func (_anyOfPlayer_ZoneItem AnyOfPlayer_ZoneItem) ZoneItem() ZoneItem {
	anyOfPlayer_ZoneItem := _anyOfPlayer_ZoneItem.anyOfPlayer_ZoneItem.engine.anyOfPlayer_ZoneItem(_anyOfPlayer_ZoneItem.anyOfPlayer_ZoneItem.ID)
	return anyOfPlayer_ZoneItem.anyOfPlayer_ZoneItem.engine.ZoneItem(ZoneItemID(anyOfPlayer_ZoneItem.anyOfPlayer_ZoneItem.ChildID))
}`
const anyOfPlayer_ZoneItem_Engine_func string = `func (engine *Engine) anyOfPlayer_ZoneItem(anyOfPlayer_ZoneItemID AnyOfPlayer_ZoneItemID) AnyOfPlayer_ZoneItem {
	patchingAnyOfPlayer_ZoneItem, ok := engine.Patch.AnyOfPlayer_ZoneItem[anyOfPlayer_ZoneItemID]
	if ok {
		return AnyOfPlayer_ZoneItem{anyOfPlayer_ZoneItem: patchingAnyOfPlayer_ZoneItem}
	}
	currentAnyOfPlayer_ZoneItem, ok := engine.State.AnyOfPlayer_ZoneItem[anyOfPlayer_ZoneItemID]
	if ok {
		return AnyOfPlayer_ZoneItem{anyOfPlayer_ZoneItem: currentAnyOfPlayer_ZoneItem}
	}
	return AnyOfPlayer_ZoneItem{anyOfPlayer_ZoneItem: anyOfPlayer_ZoneItemCore{OperationKind: OperationKindDelete, engine: engine}}
}`
const anyOfItem_Player_ZoneItem_Engine_func string = `func (engine *Engine) anyOfItem_Player_ZoneItem(anyOfItem_Player_ZoneItemID AnyOfItem_Player_ZoneItemID) AnyOfItem_Player_ZoneItem {
	patchingAnyOfItem_Player_ZoneItem, ok := engine.Patch.AnyOfItem_Player_ZoneItem[anyOfItem_Player_ZoneItemID]
	if ok {
		return AnyOfItem_Player_ZoneItem{anyOfItem_Player_ZoneItem: patchingAnyOfItem_Player_ZoneItem}
	}
	currentAnyOfItem_Player_ZoneItem, ok := engine.State.AnyOfItem_Player_ZoneItem[anyOfItem_Player_ZoneItemID]
	if ok {
		return AnyOfItem_Player_ZoneItem{anyOfItem_Player_ZoneItem: currentAnyOfItem_Player_ZoneItem}
	}
	return AnyOfItem_Player_ZoneItem{anyOfItem_Player_ZoneItem: anyOfItem_Player_ZoneItemCore{OperationKind: OperationKindDelete, engine: engine}}
}`
const _ID_AnyOfItem_Player_ZoneItem_func string = `func (_anyOfItem_Player_ZoneItem AnyOfItem_Player_ZoneItem) ID() AnyOfItem_Player_ZoneItemID {
	return _anyOfItem_Player_ZoneItem.anyOfItem_Player_ZoneItem.ID
}`
const _Player_AnyOfItem_Player_ZoneItem_func string = `func (_anyOfItem_Player_ZoneItem AnyOfItem_Player_ZoneItem) Player() Player {
	anyOfItem_Player_ZoneItem := _anyOfItem_Player_ZoneItem.anyOfItem_Player_ZoneItem.engine.anyOfItem_Player_ZoneItem(_anyOfItem_Player_ZoneItem.anyOfItem_Player_ZoneItem.ID)
	return anyOfItem_Player_ZoneItem.anyOfItem_Player_ZoneItem.engine.Player(PlayerID(anyOfItem_Player_ZoneItem.anyOfItem_Player_ZoneItem.ChildID))
}`
const _ZoneItem_AnyOfItem_Player_ZoneItem_func string = `func (_anyOfItem_Player_ZoneItem AnyOfItem_Player_ZoneItem) ZoneItem() ZoneItem {
	anyOfItem_Player_ZoneItem := _anyOfItem_Player_ZoneItem.anyOfItem_Player_ZoneItem.engine.anyOfItem_Player_ZoneItem(_anyOfItem_Player_ZoneItem.anyOfItem_Player_ZoneItem.ID)
	return anyOfItem_Player_ZoneItem.anyOfItem_Player_ZoneItem.engine.ZoneItem(ZoneItemID(anyOfItem_Player_ZoneItem.anyOfItem_Player_ZoneItem.ChildID))
}`
const _Item_AnyOfItem_Player_ZoneItem_func string = `func (_anyOfItem_Player_ZoneItem AnyOfItem_Player_ZoneItem) Item() Item {
	anyOfItem_Player_ZoneItem := _anyOfItem_Player_ZoneItem.anyOfItem_Player_ZoneItem.engine.anyOfItem_Player_ZoneItem(_anyOfItem_Player_ZoneItem.anyOfItem_Player_ZoneItem.ID)
	return anyOfItem_Player_ZoneItem.anyOfItem_Player_ZoneItem.engine.Item(ItemID(anyOfItem_Player_ZoneItem.anyOfItem_Player_ZoneItem.ChildID))
}`
const deduplicateZoneItemIDs_func string = `func deduplicateZoneItemIDs(a []ZoneItemID, b []ZoneItemID) []ZoneItemID {
	check := zoneItemCheckPool.Get().(map[ZoneItemID]bool)
	for k := range check {
		delete(check, k)
	}
	deduped := zoneItemIDSlicePool.Get().([]ZoneItemID)[:0]
	for _, val := range a {
		check[val] = true
	}
	for _, val := range b {
		check[val] = true
	}
	for val := range check {
		deduped = append(deduped, val)
	}
	zoneItemCheckPool.Put(check)
	return deduped
}`
const deduplicateZoneIDs_func string = `func deduplicateZoneIDs(a []ZoneID, b []ZoneID) []ZoneID {
	check := zoneCheckPool.Get().(map[ZoneID]bool)
	for k := range check {
		delete(check, k)
	}
	deduped := zoneIDSlicePool.Get().([]ZoneID)[:0]
	for _, val := range a {
		check[val] = true
	}
	for _, val := range b {
		check[val] = true
	}
	for val := range check {
		deduped = append(deduped, val)
	}
	zoneCheckPool.Put(check)
	return deduped
}`
const deduplicatePlayerIDs_func string = `func deduplicatePlayerIDs(a []PlayerID, b []PlayerID) []PlayerID {
	check := playerCheckPool.Get().(map[PlayerID]bool)
	for k := range check {
		delete(check, k)
	}
	deduped := playerIDSlicePool.Get().([]PlayerID)[:0]
	for _, val := range a {
		check[val] = true
	}
	for _, val := range b {
		check[val] = true
	}
	for val := range check {
		deduped = append(deduped, val)
	}
	playerCheckPool.Put(check)
	return deduped
}`
const deduplicatePositionIDs_func string = `func deduplicatePositionIDs(a []PositionID, b []PositionID) []PositionID {
	check := positionCheckPool.Get().(map[PositionID]bool)
	for k := range check {
		delete(check, k)
	}
	deduped := positionIDSlicePool.Get().([]PositionID)[:0]
	for _, val := range a {
		check[val] = true
	}
	for _, val := range b {
		check[val] = true
	}
	for val := range check {
		deduped = append(deduped, val)
	}
	positionCheckPool.Put(check)
	return deduped
}`
const deduplicateItemIDs_func string = `func deduplicateItemIDs(a []ItemID, b []ItemID) []ItemID {
	check := itemCheckPool.Get().(map[ItemID]bool)
	for k := range check {
		delete(check, k)
	}
	deduped := itemIDSlicePool.Get().([]ItemID)[:0]
	for _, val := range a {
		check[val] = true
	}
	for _, val := range b {
		check[val] = true
	}
	for val := range check {
		deduped = append(deduped, val)
	}
	itemCheckPool.Put(check)
	return deduped
}`
const deduplicateGearScoreIDs_func string = `func deduplicateGearScoreIDs(a []GearScoreID, b []GearScoreID) []GearScoreID {
	check := gearScoreCheckPool.Get().(map[GearScoreID]bool)
	for k := range check {
		delete(check, k)
	}
	deduped := gearScoreIDSlicePool.Get().([]GearScoreID)[:0]
	for _, val := range a {
		check[val] = true
	}
	for _, val := range b {
		check[val] = true
	}
	for val := range check {
		deduped = append(deduped, val)
	}
	gearScoreCheckPool.Put(check)
	return deduped
}`
const deduplicateAttackEventIDs_func string = `func deduplicateAttackEventIDs(a []AttackEventID, b []AttackEventID) []AttackEventID {
	check := attackEventCheckPool.Get().(map[AttackEventID]bool)
	for k := range check {
		delete(check, k)
	}
	deduped := attackEventIDSlicePool.Get().([]AttackEventID)[:0]
	for _, val := range a {
		check[val] = true
	}
	for _, val := range b {
		check[val] = true
	}
	for val := range check {
		deduped = append(deduped, val)
	}
	attackEventCheckPool.Put(check)
	return deduped
}`
const deduplicateEquipmentSetIDs_func string = `func deduplicateEquipmentSetIDs(a []EquipmentSetID, b []EquipmentSetID) []EquipmentSetID {
	check := equipmentSetCheckPool.Get().(map[EquipmentSetID]bool)
	for k := range check {
		delete(check, k)
	}
	deduped := equipmentSetIDSlicePool.Get().([]EquipmentSetID)[:0]
	for _, val := range a {
		check[val] = true
	}
	for _, val := range b {
		check[val] = true
	}
	for val := range check {
		deduped = append(deduped, val)
	}
	equipmentSetCheckPool.Put(check)
	return deduped
}`
const deduplicatePlayerTargetedByRefIDs_func string = `func deduplicatePlayerTargetedByRefIDs(a []PlayerTargetedByRefID, b []PlayerTargetedByRefID) []PlayerTargetedByRefID {
	check := playerTargetedByRefCheckPool.Get().(map[PlayerTargetedByRefID]bool)
	for k := range check {
		delete(check, k)
	}
	deduped := playerTargetedByRefIDSlicePool.Get().([]PlayerTargetedByRefID)[:0]
	for _, val := range a {
		check[val] = true
	}
	for _, val := range b {
		check[val] = true
	}
	for val := range check {
		deduped = append(deduped, val)
	}
	playerTargetedByRefCheckPool.Put(check)
	return deduped
}`
const deduplicatePlayerTargetRefIDs_func string = `func deduplicatePlayerTargetRefIDs(a []PlayerTargetRefID, b []PlayerTargetRefID) []PlayerTargetRefID {
	check := playerTargetRefCheckPool.Get().(map[PlayerTargetRefID]bool)
	for k := range check {
		delete(check, k)
	}
	deduped := playerTargetRefIDSlicePool.Get().([]PlayerTargetRefID)[:0]
	for _, val := range a {
		check[val] = true
	}
	for _, val := range b {
		check[val] = true
	}
	for val := range check {
		deduped = append(deduped, val)
	}
	playerTargetRefCheckPool.Put(check)
	return deduped
}`
const deduplicateAttackEventTargetRefIDs_func string = `func deduplicateAttackEventTargetRefIDs(a []AttackEventTargetRefID, b []AttackEventTargetRefID) []AttackEventTargetRefID {
	check := attackEventTargetRefCheckPool.Get().(map[AttackEventTargetRefID]bool)
	for k := range check {
		delete(check, k)
	}
	deduped := attackEventTargetRefIDSlicePool.Get().([]AttackEventTargetRefID)[:0]
	for _, val := range a {
		check[val] = true
	}
	for _, val := range b {
		check[val] = true
	}
	for val := range check {
		deduped = append(deduped, val)
	}
	attackEventTargetRefCheckPool.Put(check)
	return deduped
}`
const deduplicateItemBoundToRefIDs_func string = `func deduplicateItemBoundToRefIDs(a []ItemBoundToRefID, b []ItemBoundToRefID) []ItemBoundToRefID {
	check := itemBoundToRefCheckPool.Get().(map[ItemBoundToRefID]bool)
	for k := range check {
		delete(check, k)
	}
	deduped := itemBoundToRefIDSlicePool.Get().([]ItemBoundToRefID)[:0]
	for _, val := range a {
		check[val] = true
	}
	for _, val := range b {
		check[val] = true
	}
	for val := range check {
		deduped = append(deduped, val)
	}
	itemBoundToRefCheckPool.Put(check)
	return deduped
}`
const deduplicatePlayerGuildMemberRefIDs_func string = `func deduplicatePlayerGuildMemberRefIDs(a []PlayerGuildMemberRefID, b []PlayerGuildMemberRefID) []PlayerGuildMemberRefID {
	check := playerGuildMemberRefCheckPool.Get().(map[PlayerGuildMemberRefID]bool)
	for k := range check {
		delete(check, k)
	}
	deduped := playerGuildMemberRefIDSlicePool.Get().([]PlayerGuildMemberRefID)[:0]
	for _, val := range a {
		check[val] = true
	}
	for _, val := range b {
		check[val] = true
	}
	for val := range check {
		deduped = append(deduped, val)
	}
	playerGuildMemberRefCheckPool.Put(check)
	return deduped
}`
const deduplicatePlayerEquipmentSetRefIDs_func string = `func deduplicatePlayerEquipmentSetRefIDs(a []PlayerEquipmentSetRefID, b []PlayerEquipmentSetRefID) []PlayerEquipmentSetRefID {
	check := playerEquipmentSetRefCheckPool.Get().(map[PlayerEquipmentSetRefID]bool)
	for k := range check {
		delete(check, k)
	}
	deduped := playerEquipmentSetRefIDSlicePool.Get().([]PlayerEquipmentSetRefID)[:0]
	for _, val := range a {
		check[val] = true
	}
	for _, val := range b {
		check[val] = true
	}
	for val := range check {
		deduped = append(deduped, val)
	}
	playerEquipmentSetRefCheckPool.Put(check)
	return deduped
}`
const deduplicateEquipmentSetEquipmentRefIDs_func string = `func deduplicateEquipmentSetEquipmentRefIDs(a []EquipmentSetEquipmentRefID, b []EquipmentSetEquipmentRefID) []EquipmentSetEquipmentRefID {
	check := equipmentSetEquipmentRefCheckPool.Get().(map[EquipmentSetEquipmentRefID]bool)
	for k := range check {
		delete(check, k)
	}
	deduped := equipmentSetEquipmentRefIDSlicePool.Get().([]EquipmentSetEquipmentRefID)[:0]
	for _, val := range a {
		check[val] = true
	}
	for _, val := range b {
		check[val] = true
	}
	for val := range check {
		deduped = append(deduped, val)
	}
	equipmentSetEquipmentRefCheckPool.Put(check)
	return deduped
}`
const allAttackEventIDs_Engine_func string = `func (engine Engine) allAttackEventIDs() []AttackEventID {
	stateAttackEventIDs := attackEventIDSlicePool.Get().([]AttackEventID)[:0]
	for attackEventID := range engine.State.AttackEvent {
		stateAttackEventIDs = append(stateAttackEventIDs, attackEventID)
	}
	patchAttackEventIDs := attackEventIDSlicePool.Get().([]AttackEventID)[:0]
	for attackEventID := range engine.Patch.AttackEvent {
		patchAttackEventIDs = append(patchAttackEventIDs, attackEventID)
	}
	dedupedIDs := deduplicateAttackEventIDs(stateAttackEventIDs, patchAttackEventIDs)
	attackEventIDSlicePool.Put(stateAttackEventIDs)
	attackEventIDSlicePool.Put(patchAttackEventIDs)
	return dedupedIDs
}`
const allEquipmentSetIDs_Engine_func string = `func (engine Engine) allEquipmentSetIDs() []EquipmentSetID {
	stateEquipmentSetIDs := equipmentSetIDSlicePool.Get().([]EquipmentSetID)[:0]
	for equipmentSetID := range engine.State.EquipmentSet {
		stateEquipmentSetIDs = append(stateEquipmentSetIDs, equipmentSetID)
	}
	patchEquipmentSetIDs := equipmentSetIDSlicePool.Get().([]EquipmentSetID)[:0]
	for equipmentSetID := range engine.Patch.EquipmentSet {
		patchEquipmentSetIDs = append(patchEquipmentSetIDs, equipmentSetID)
	}
	dedupedIDs := deduplicateEquipmentSetIDs(stateEquipmentSetIDs, patchEquipmentSetIDs)
	equipmentSetIDSlicePool.Put(stateEquipmentSetIDs)
	equipmentSetIDSlicePool.Put(patchEquipmentSetIDs)
	return dedupedIDs
}`
const allGearScoreIDs_Engine_func string = `func (engine Engine) allGearScoreIDs() []GearScoreID {
	stateGearScoreIDs := gearScoreIDSlicePool.Get().([]GearScoreID)[:0]
	for gearScoreID := range engine.State.GearScore {
		stateGearScoreIDs = append(stateGearScoreIDs, gearScoreID)
	}
	patchGearScoreIDs := gearScoreIDSlicePool.Get().([]GearScoreID)[:0]
	for gearScoreID := range engine.Patch.GearScore {
		patchGearScoreIDs = append(patchGearScoreIDs, gearScoreID)
	}
	dedupedIDs := deduplicateGearScoreIDs(stateGearScoreIDs, patchGearScoreIDs)
	gearScoreIDSlicePool.Put(stateGearScoreIDs)
	gearScoreIDSlicePool.Put(patchGearScoreIDs)
	return dedupedIDs
}`
const allItemIDs_Engine_func string = `func (engine Engine) allItemIDs() []ItemID {
	stateItemIDs := itemIDSlicePool.Get().([]ItemID)[:0]
	for itemID := range engine.State.Item {
		stateItemIDs = append(stateItemIDs, itemID)
	}
	patchItemIDs := itemIDSlicePool.Get().([]ItemID)[:0]
	for itemID := range engine.Patch.Item {
		patchItemIDs = append(patchItemIDs, itemID)
	}
	dedupedIDs := deduplicateItemIDs(stateItemIDs, patchItemIDs)
	itemIDSlicePool.Put(stateItemIDs)
	itemIDSlicePool.Put(patchItemIDs)
	return dedupedIDs
}`
const allPositionIDs_Engine_func string = `func (engine Engine) allPositionIDs() []PositionID {
	statePositionIDs := positionIDSlicePool.Get().([]PositionID)[:0]
	for positionID := range engine.State.Position {
		statePositionIDs = append(statePositionIDs, positionID)
	}
	patchPositionIDs := positionIDSlicePool.Get().([]PositionID)[:0]
	for positionID := range engine.Patch.Position {
		patchPositionIDs = append(patchPositionIDs, positionID)
	}
	dedupedIDs := deduplicatePositionIDs(statePositionIDs, patchPositionIDs)
	positionIDSlicePool.Put(statePositionIDs)
	positionIDSlicePool.Put(patchPositionIDs)
	return dedupedIDs
}`
const allZoneIDs_Engine_func string = `func (engine Engine) allZoneIDs() []ZoneID {
	stateZoneIDs := zoneIDSlicePool.Get().([]ZoneID)[:0]
	for zoneID := range engine.State.Zone {
		stateZoneIDs = append(stateZoneIDs, zoneID)
	}
	patchZoneIDs := zoneIDSlicePool.Get().([]ZoneID)[:0]
	for zoneID := range engine.Patch.Zone {
		patchZoneIDs = append(patchZoneIDs, zoneID)
	}
	dedupedIDs := deduplicateZoneIDs(stateZoneIDs, patchZoneIDs)
	zoneIDSlicePool.Put(stateZoneIDs)
	zoneIDSlicePool.Put(patchZoneIDs)
	return dedupedIDs
}`
const allZoneItemIDs_Engine_func string = `func (engine Engine) allZoneItemIDs() []ZoneItemID {
	stateZoneItemIDs := zoneItemIDSlicePool.Get().([]ZoneItemID)[:0]
	for zoneItemID := range engine.State.ZoneItem {
		stateZoneItemIDs = append(stateZoneItemIDs, zoneItemID)
	}
	patchZoneItemIDs := zoneItemIDSlicePool.Get().([]ZoneItemID)[:0]
	for zoneItemID := range engine.Patch.ZoneItem {
		patchZoneItemIDs = append(patchZoneItemIDs, zoneItemID)
	}
	dedupedIDs := deduplicateZoneItemIDs(stateZoneItemIDs, patchZoneItemIDs)
	zoneItemIDSlicePool.Put(stateZoneItemIDs)
	zoneItemIDSlicePool.Put(patchZoneItemIDs)
	return dedupedIDs
}`
const allPlayerIDs_Engine_func string = `func (engine Engine) allPlayerIDs() []PlayerID {
	statePlayerIDs := playerIDSlicePool.Get().([]PlayerID)[:0]
	for playerID := range engine.State.Player {
		statePlayerIDs = append(statePlayerIDs, playerID)
	}
	patchPlayerIDs := playerIDSlicePool.Get().([]PlayerID)[:0]
	for playerID := range engine.Patch.Player {
		patchPlayerIDs = append(patchPlayerIDs, playerID)
	}
	dedupedIDs := deduplicatePlayerIDs(statePlayerIDs, patchPlayerIDs)
	playerIDSlicePool.Put(statePlayerIDs)
	playerIDSlicePool.Put(patchPlayerIDs)
	return dedupedIDs
}`
const allPlayerTargetedByRefIDs_Engine_func string = `func (engine Engine) allPlayerTargetedByRefIDs() []PlayerTargetedByRefID {
	statePlayerTargetedByRefIDs := playerTargetedByRefIDSlicePool.Get().([]PlayerTargetedByRefID)[:0]
	for playerTargetedByRefID := range engine.State.PlayerTargetedByRef {
		statePlayerTargetedByRefIDs = append(statePlayerTargetedByRefIDs, playerTargetedByRefID)
	}
	patchPlayerTargetedByRefIDs := playerTargetedByRefIDSlicePool.Get().([]PlayerTargetedByRefID)[:0]
	for playerTargetedByRefID := range engine.Patch.PlayerTargetedByRef {
		patchPlayerTargetedByRefIDs = append(patchPlayerTargetedByRefIDs, playerTargetedByRefID)
	}
	dedupedIDs := deduplicatePlayerTargetedByRefIDs(statePlayerTargetedByRefIDs, patchPlayerTargetedByRefIDs)
	playerTargetedByRefIDSlicePool.Put(statePlayerTargetedByRefIDs)
	playerTargetedByRefIDSlicePool.Put(patchPlayerTargetedByRefIDs)
	return dedupedIDs
}`
const allPlayerTargetRefIDs_Engine_func string = `func (engine Engine) allPlayerTargetRefIDs() []PlayerTargetRefID {
	statePlayerTargetRefIDs := playerTargetRefIDSlicePool.Get().([]PlayerTargetRefID)[:0]
	for playerTargetRefID := range engine.State.PlayerTargetRef {
		statePlayerTargetRefIDs = append(statePlayerTargetRefIDs, playerTargetRefID)
	}
	patchPlayerTargetRefIDs := playerTargetRefIDSlicePool.Get().([]PlayerTargetRefID)[:0]
	for playerTargetRefID := range engine.Patch.PlayerTargetRef {
		patchPlayerTargetRefIDs = append(patchPlayerTargetRefIDs, playerTargetRefID)
	}
	dedupedIDs := deduplicatePlayerTargetRefIDs(statePlayerTargetRefIDs, patchPlayerTargetRefIDs)
	playerTargetRefIDSlicePool.Put(statePlayerTargetRefIDs)
	playerTargetRefIDSlicePool.Put(patchPlayerTargetRefIDs)
	return dedupedIDs
}`
const allItemBoundToRefIDs_Engine_func string = `func (engine Engine) allItemBoundToRefIDs() []ItemBoundToRefID {
	stateItemBoundToRefIDs := itemBoundToRefIDSlicePool.Get().([]ItemBoundToRefID)[:0]
	for itemBoundToRefID := range engine.State.ItemBoundToRef {
		stateItemBoundToRefIDs = append(stateItemBoundToRefIDs, itemBoundToRefID)
	}
	patchItemBoundToRefIDs := itemBoundToRefIDSlicePool.Get().([]ItemBoundToRefID)[:0]
	for itemBoundToRefID := range engine.Patch.ItemBoundToRef {
		patchItemBoundToRefIDs = append(patchItemBoundToRefIDs, itemBoundToRefID)
	}
	dedupedIDs := deduplicateItemBoundToRefIDs(stateItemBoundToRefIDs, patchItemBoundToRefIDs)
	itemBoundToRefIDSlicePool.Put(stateItemBoundToRefIDs)
	itemBoundToRefIDSlicePool.Put(patchItemBoundToRefIDs)
	return dedupedIDs
}`
const allAttackEventTargetRefIDs_Engine_func string = `func (engine Engine) allAttackEventTargetRefIDs() []AttackEventTargetRefID {
	stateAttackEventTargetRefIDs := attackEventTargetRefIDSlicePool.Get().([]AttackEventTargetRefID)[:0]
	for attackEventTargetRefID := range engine.State.AttackEventTargetRef {
		stateAttackEventTargetRefIDs = append(stateAttackEventTargetRefIDs, attackEventTargetRefID)
	}
	patchAttackEventTargetRefIDs := attackEventTargetRefIDSlicePool.Get().([]AttackEventTargetRefID)[:0]
	for attackEventTargetRefID := range engine.Patch.AttackEventTargetRef {
		patchAttackEventTargetRefIDs = append(patchAttackEventTargetRefIDs, attackEventTargetRefID)
	}
	dedupedIDs := deduplicateAttackEventTargetRefIDs(stateAttackEventTargetRefIDs, patchAttackEventTargetRefIDs)
	attackEventTargetRefIDSlicePool.Put(stateAttackEventTargetRefIDs)
	attackEventTargetRefIDSlicePool.Put(patchAttackEventTargetRefIDs)
	return dedupedIDs
}`
const allPlayerGuildMemberRefIDs_Engine_func string = `func (engine Engine) allPlayerGuildMemberRefIDs() []PlayerGuildMemberRefID {
	statePlayerGuildMemberRefIDs := playerGuildMemberRefIDSlicePool.Get().([]PlayerGuildMemberRefID)[:0]
	for playerGuildMemberRefID := range engine.State.PlayerGuildMemberRef {
		statePlayerGuildMemberRefIDs = append(statePlayerGuildMemberRefIDs, playerGuildMemberRefID)
	}
	patchPlayerGuildMemberRefIDs := playerGuildMemberRefIDSlicePool.Get().([]PlayerGuildMemberRefID)[:0]
	for playerGuildMemberRefID := range engine.Patch.PlayerGuildMemberRef {
		patchPlayerGuildMemberRefIDs = append(patchPlayerGuildMemberRefIDs, playerGuildMemberRefID)
	}
	dedupedIDs := deduplicatePlayerGuildMemberRefIDs(statePlayerGuildMemberRefIDs, patchPlayerGuildMemberRefIDs)
	playerGuildMemberRefIDSlicePool.Put(statePlayerGuildMemberRefIDs)
	playerGuildMemberRefIDSlicePool.Put(patchPlayerGuildMemberRefIDs)
	return dedupedIDs
}`
const allPlayerEquipmentSetRefIDs_Engine_func string = `func (engine Engine) allPlayerEquipmentSetRefIDs() []PlayerEquipmentSetRefID {
	statePlayerEquipmentSetRefIDs := playerEquipmentSetRefIDSlicePool.Get().([]PlayerEquipmentSetRefID)[:0]
	for playerEquipmentSetRefID := range engine.State.PlayerEquipmentSetRef {
		statePlayerEquipmentSetRefIDs = append(statePlayerEquipmentSetRefIDs, playerEquipmentSetRefID)
	}
	patchPlayerEquipmentSetRefIDs := playerEquipmentSetRefIDSlicePool.Get().([]PlayerEquipmentSetRefID)[:0]
	for playerEquipmentSetRefID := range engine.Patch.PlayerEquipmentSetRef {
		patchPlayerEquipmentSetRefIDs = append(patchPlayerEquipmentSetRefIDs, playerEquipmentSetRefID)
	}
	dedupedIDs := deduplicatePlayerEquipmentSetRefIDs(statePlayerEquipmentSetRefIDs, patchPlayerEquipmentSetRefIDs)
	playerEquipmentSetRefIDSlicePool.Put(statePlayerEquipmentSetRefIDs)
	playerEquipmentSetRefIDSlicePool.Put(patchPlayerEquipmentSetRefIDs)
	return dedupedIDs
}`
const allEquipmentSetEquipmentRefIDs_Engine_func string = `func (engine Engine) allEquipmentSetEquipmentRefIDs() []EquipmentSetEquipmentRefID {
	stateEquipmentSetEquipmentRefIDs := equipmentSetEquipmentRefIDSlicePool.Get().([]EquipmentSetEquipmentRefID)[:0]
	for equipmentSetEquipmentRefID := range engine.State.EquipmentSetEquipmentRef {
		stateEquipmentSetEquipmentRefIDs = append(stateEquipmentSetEquipmentRefIDs, equipmentSetEquipmentRefID)
	}
	patchEquipmentSetEquipmentRefIDs := equipmentSetEquipmentRefIDSlicePool.Get().([]EquipmentSetEquipmentRefID)[:0]
	for equipmentSetEquipmentRefID := range engine.Patch.EquipmentSetEquipmentRef {
		patchEquipmentSetEquipmentRefIDs = append(patchEquipmentSetEquipmentRefIDs, equipmentSetEquipmentRefID)
	}
	dedupedIDs := deduplicateEquipmentSetEquipmentRefIDs(stateEquipmentSetEquipmentRefIDs, patchEquipmentSetEquipmentRefIDs)
	equipmentSetEquipmentRefIDSlicePool.Put(stateEquipmentSetEquipmentRefIDs)
	equipmentSetEquipmentRefIDSlicePool.Put(patchEquipmentSetEquipmentRefIDs)
	return dedupedIDs
}`
const path_generated_go_import string = `import (
	"fmt"
	"strconv"
)`
const treeFieldIdentifier_type string = `type treeFieldIdentifier int`
const attackEventIdentifier_type string = `const (
	attackEventIdentifier			treeFieldIdentifier	= 1
	equipmentSetIdentifier			treeFieldIdentifier	= 2
	gearScoreIdentifier			treeFieldIdentifier	= 3
	itemIdentifier				treeFieldIdentifier	= 4
	playerIdentifier			treeFieldIdentifier	= 5
	positionIdentifier			treeFieldIdentifier	= 6
	zoneIdentifier				treeFieldIdentifier	= 7
	zoneItemIdentifier			treeFieldIdentifier	= 8
	attackEvent_targetIdentifier		treeFieldIdentifier	= 9
	equipmentSet_equipmentIdentifier	treeFieldIdentifier	= 10
	equipmentSet_nameIdentifier		treeFieldIdentifier	= 11
	gearScore_levelIdentifier		treeFieldIdentifier	= 12
	gearScore_scoreIdentifier		treeFieldIdentifier	= 13
	item_boundToIdentifier			treeFieldIdentifier	= 14
	item_gearScoreIdentifier		treeFieldIdentifier	= 15
	item_nameIdentifier			treeFieldIdentifier	= 16
	item_originIdentifier			treeFieldIdentifier	= 17
	player_actionIdentifier			treeFieldIdentifier	= 18
	player_equipmentSetsIdentifier		treeFieldIdentifier	= 19
	player_gearScoreIdentifier		treeFieldIdentifier	= 20
	player_guildMembersIdentifier		treeFieldIdentifier	= 21
	player_itemsIdentifier			treeFieldIdentifier	= 22
	player_positionIdentifier		treeFieldIdentifier	= 23
	player_targetIdentifier			treeFieldIdentifier	= 24
	player_targetedByIdentifier		treeFieldIdentifier	= 25
	position_xIdentifier			treeFieldIdentifier	= 26
	position_yIdentifier			treeFieldIdentifier	= 27
	zone_interactablesIdentifier		treeFieldIdentifier	= 28
	zone_itemsIdentifier			treeFieldIdentifier	= 29
	zone_playersIdentifier			treeFieldIdentifier	= 30
	zone_tagsIdentifier			treeFieldIdentifier	= 31
	zoneItem_itemIdentifier			treeFieldIdentifier	= 32
	zoneItem_positionIdentifier		treeFieldIdentifier	= 33
)`
const toString_treeFieldIdentifier_func string = `func (t treeFieldIdentifier) toString() string {
	switch t {
	case attackEventIdentifier:
		return "attackEvent"
	case equipmentSetIdentifier:
		return "equipmentSet"
	case gearScoreIdentifier:
		return "gearScore"
	case itemIdentifier:
		return "item"
	case playerIdentifier:
		return "player"
	case positionIdentifier:
		return "position"
	case zoneIdentifier:
		return "zone"
	case zoneItemIdentifier:
		return "zoneItem"
	case attackEvent_targetIdentifier:
		return "target"
	case equipmentSet_equipmentIdentifier:
		return "equipment"
	case equipmentSet_nameIdentifier:
		return "name"
	case gearScore_levelIdentifier:
		return "level"
	case gearScore_scoreIdentifier:
		return "score"
	case item_boundToIdentifier:
		return "boundTo"
	case item_gearScoreIdentifier:
		return "gearScore"
	case item_nameIdentifier:
		return "name"
	case item_originIdentifier:
		return "origin"
	case player_actionIdentifier:
		return "action"
	case player_equipmentSetsIdentifier:
		return "equipmentSets"
	case player_gearScoreIdentifier:
		return "gearScore"
	case player_guildMembersIdentifier:
		return "guildMembers"
	case player_itemsIdentifier:
		return "items"
	case player_positionIdentifier:
		return "position"
	case player_targetIdentifier:
		return "target"
	case player_targetedByIdentifier:
		return "targetedBy"
	case position_xIdentifier:
		return "x"
	case position_yIdentifier:
		return "y"
	case zone_interactablesIdentifier:
		return "interactables"
	case zone_itemsIdentifier:
		return "items"
	case zone_playersIdentifier:
		return "players"
	case zone_tagsIdentifier:
		return "tags"
	case zoneItem_itemIdentifier:
		return "item"
	case zoneItem_positionIdentifier:
		return "position"
	default:
		panic(fmt.Sprintf("no string found for identifier <%d>", t))
	}
}`
const segment_type string = `type segment struct {
	ID		int			` + "`" + `json:"id"` + "`" + `
	Identifier	treeFieldIdentifier	` + "`" + `json:"identifier"` + "`" + `
	Kind		ElementKind		` + "`" + `json:"kind"` + "`" + `
	RefID		int			` + "`" + `json:"refID"` + "`" + `
}`
const path_type string = `type path []segment`
const newPath_func string = `func newPath() path {
	return make(path, 0)
}`
const extendAndCopy_path_func string = `func (p path) extendAndCopy(fieldIdentifier treeFieldIdentifier, id int, kind ElementKind, refID int) path {
	newPath := make(path, len(p), len(p)+1)
	copy(newPath, p)
	newPath = append(newPath, segment{id, fieldIdentifier, kind, refID})
	return newPath
}`
const toJSONPath_path_func string = `func (p path) toJSONPath() string {
	jsonPath := "$"
	for _, seg := range p {
		jsonPath += "." + seg.Identifier.toString()
		if isSliceFieldIdentifier(seg.Identifier) {
			jsonPath += "[" + strconv.Itoa(seg.ID) + "]"
		}
	}
	return jsonPath
}`
const isSliceFieldIdentifier_func string = `func isSliceFieldIdentifier(fieldIdentifier treeFieldIdentifier) bool {
	switch fieldIdentifier {
	case attackEventIdentifier:
		return true
	case equipmentSetIdentifier:
		return true
	case gearScoreIdentifier:
		return true
	case itemIdentifier:
		return true
	case playerIdentifier:
		return true
	case positionIdentifier:
		return true
	case zoneIdentifier:
		return true
	case zoneItemIdentifier:
		return true
	case equipmentSet_equipmentIdentifier:
		return true
	case player_actionIdentifier:
		return true
	case player_equipmentSetsIdentifier:
		return true
	case player_guildMembersIdentifier:
		return true
	case player_itemsIdentifier:
		return true
	case player_targetedByIdentifier:
		return true
	case zone_interactablesIdentifier:
		return true
	case zone_itemsIdentifier:
		return true
	case zone_playersIdentifier:
		return true
	}
	return false
}`
const pools_generated_go_import string = `import (
	"sync"
)`
const zoneItemCheckPool_type string = `var zoneItemCheckPool = sync.Pool{New: func() interface{} {
	return make(map[ZoneItemID]bool)
}}`
const zoneItemIDSlicePool_type string = `var zoneItemIDSlicePool = sync.Pool{New: func() interface{} {
	return make([]ZoneItemID, 0)
}}`
const zoneCheckPool_type string = `var zoneCheckPool = sync.Pool{New: func() interface{} {
	return make(map[ZoneID]bool)
}}`
const zoneIDSlicePool_type string = `var zoneIDSlicePool = sync.Pool{New: func() interface{} {
	return make([]ZoneID, 0)
}}`
const playerCheckPool_type string = `var playerCheckPool = sync.Pool{New: func() interface{} {
	return make(map[PlayerID]bool)
}}`
const playerIDSlicePool_type string = `var playerIDSlicePool = sync.Pool{New: func() interface{} {
	return make([]PlayerID, 0)
}}`
const positionCheckPool_type string = `var positionCheckPool = sync.Pool{New: func() interface{} {
	return make(map[PositionID]bool)
}}`
const positionIDSlicePool_type string = `var positionIDSlicePool = sync.Pool{New: func() interface{} {
	return make([]PositionID, 0)
}}`
const itemCheckPool_type string = `var itemCheckPool = sync.Pool{New: func() interface{} {
	return make(map[ItemID]bool)
}}`
const itemIDSlicePool_type string = `var itemIDSlicePool = sync.Pool{New: func() interface{} {
	return make([]ItemID, 0)
}}`
const gearScoreCheckPool_type string = `var gearScoreCheckPool = sync.Pool{New: func() interface{} {
	return make(map[GearScoreID]bool)
}}`
const gearScoreIDSlicePool_type string = `var gearScoreIDSlicePool = sync.Pool{New: func() interface{} {
	return make([]GearScoreID, 0)
}}`
const attackEventCheckPool_type string = `var attackEventCheckPool = sync.Pool{New: func() interface{} {
	return make(map[AttackEventID]bool)
}}`
const attackEventIDSlicePool_type string = `var attackEventIDSlicePool = sync.Pool{New: func() interface{} {
	return make([]AttackEventID, 0)
}}`
const equipmentSetCheckPool_type string = `var equipmentSetCheckPool = sync.Pool{New: func() interface{} {
	return make(map[EquipmentSetID]bool)
}}`
const equipmentSetIDSlicePool_type string = `var equipmentSetIDSlicePool = sync.Pool{New: func() interface{} {
	return make([]EquipmentSetID, 0)
}}`
const playerTargetedByRefCheckPool_type string = `var playerTargetedByRefCheckPool = sync.Pool{New: func() interface{} {
	return make(map[PlayerTargetedByRefID]bool)
}}`
const playerTargetedByRefIDSlicePool_type string = `var playerTargetedByRefIDSlicePool = sync.Pool{New: func() interface{} {
	return make([]PlayerTargetedByRefID, 0)
}}`
const playerTargetRefCheckPool_type string = `var playerTargetRefCheckPool = sync.Pool{New: func() interface{} {
	return make(map[PlayerTargetRefID]bool)
}}`
const playerTargetRefIDSlicePool_type string = `var playerTargetRefIDSlicePool = sync.Pool{New: func() interface{} {
	return make([]PlayerTargetRefID, 0)
}}`
const attackEventTargetRefCheckPool_type string = `var attackEventTargetRefCheckPool = sync.Pool{New: func() interface{} {
	return make(map[AttackEventTargetRefID]bool)
}}`
const attackEventTargetRefIDSlicePool_type string = `var attackEventTargetRefIDSlicePool = sync.Pool{New: func() interface{} {
	return make([]AttackEventTargetRefID, 0)
}}`
const itemBoundToRefCheckPool_type string = `var itemBoundToRefCheckPool = sync.Pool{New: func() interface{} {
	return make(map[ItemBoundToRefID]bool)
}}`
const itemBoundToRefIDSlicePool_type string = `var itemBoundToRefIDSlicePool = sync.Pool{New: func() interface{} {
	return make([]ItemBoundToRefID, 0)
}}`
const playerGuildMemberRefCheckPool_type string = `var playerGuildMemberRefCheckPool = sync.Pool{New: func() interface{} {
	return make(map[PlayerGuildMemberRefID]bool)
}}`
const playerGuildMemberRefIDSlicePool_type string = `var playerGuildMemberRefIDSlicePool = sync.Pool{New: func() interface{} {
	return make([]PlayerGuildMemberRefID, 0)
}}`
const playerEquipmentSetRefCheckPool_type string = `var playerEquipmentSetRefCheckPool = sync.Pool{New: func() interface{} {
	return make(map[PlayerEquipmentSetRefID]bool)
}}`
const playerEquipmentSetRefIDSlicePool_type string = `var playerEquipmentSetRefIDSlicePool = sync.Pool{New: func() interface{} {
	return make([]PlayerEquipmentSetRefID, 0)
}}`
const equipmentSetEquipmentRefCheckPool_type string = `var equipmentSetEquipmentRefCheckPool = sync.Pool{New: func() interface{} {
	return make(map[EquipmentSetEquipmentRefID]bool)
}}`
const equipmentSetEquipmentRefIDSlicePool_type string = `var equipmentSetEquipmentRefIDSlicePool = sync.Pool{New: func() interface{} {
	return make([]EquipmentSetEquipmentRefID, 0)
}}`
const _IsSet_ItemBoundToRef_func string = `func (_ref ItemBoundToRef) IsSet() (ItemBoundToRef, bool) {
	ref := _ref.itemBoundToRef.engine.itemBoundToRef(_ref.itemBoundToRef.ID)
	return ref, ref.itemBoundToRef.ID != 0
}`
const _Unset_ItemBoundToRef_func string = `func (_ref ItemBoundToRef) Unset() {
	ref := _ref.itemBoundToRef.engine.itemBoundToRef(_ref.itemBoundToRef.ID)
	if ref.itemBoundToRef.OperationKind == OperationKindDelete {
		return
	}
	ref.itemBoundToRef.engine.deleteItemBoundToRef(ref.itemBoundToRef.ID)
	parent := ref.itemBoundToRef.engine.Item(ref.itemBoundToRef.ParentID).item
	if parent.OperationKind == OperationKindDelete {
		return
	}
	parent.BoundTo = 0
	parent.OperationKind = OperationKindUpdate
	ref.itemBoundToRef.engine.Patch.Item[parent.ID] = parent
}`
const _Get_ItemBoundToRef_func string = `func (_ref ItemBoundToRef) Get() Player {
	ref := _ref.itemBoundToRef.engine.itemBoundToRef(_ref.itemBoundToRef.ID)
	return ref.itemBoundToRef.engine.Player(ref.itemBoundToRef.ReferencedElementID)
}`
const _IsSet_AttackEventTargetRef_func string = `func (_ref AttackEventTargetRef) IsSet() (AttackEventTargetRef, bool) {
	ref := _ref.attackEventTargetRef.engine.attackEventTargetRef(_ref.attackEventTargetRef.ID)
	return ref, ref.attackEventTargetRef.ID != 0
}`
const _Unset_AttackEventTargetRef_func string = `func (_ref AttackEventTargetRef) Unset() {
	ref := _ref.attackEventTargetRef.engine.attackEventTargetRef(_ref.attackEventTargetRef.ID)
	if ref.attackEventTargetRef.OperationKind == OperationKindDelete {
		return
	}
	ref.attackEventTargetRef.engine.deleteAttackEventTargetRef(ref.attackEventTargetRef.ID)
	parent := ref.attackEventTargetRef.engine.AttackEvent(ref.attackEventTargetRef.ParentID).attackEvent
	if parent.OperationKind == OperationKindDelete {
		return
	}
	parent.Target = 0
	parent.OperationKind = OperationKindUpdate
	ref.attackEventTargetRef.engine.Patch.AttackEvent[parent.ID] = parent
}`
const _Get_AttackEventTargetRef_func string = `func (_ref AttackEventTargetRef) Get() Player {
	ref := _ref.attackEventTargetRef.engine.attackEventTargetRef(_ref.attackEventTargetRef.ID)
	return ref.attackEventTargetRef.engine.Player(ref.attackEventTargetRef.ReferencedElementID)
}`
const _Get_PlayerGuildMemberRef_func string = `func (_ref PlayerGuildMemberRef) Get() Player {
	ref := _ref.playerGuildMemberRef.engine.playerGuildMemberRef(_ref.playerGuildMemberRef.ID)
	return ref.playerGuildMemberRef.engine.Player(ref.playerGuildMemberRef.ReferencedElementID)
}`
const _Get_PlayerEquipmentSetRef_func string = `func (_ref PlayerEquipmentSetRef) Get() EquipmentSet {
	ref := _ref.playerEquipmentSetRef.engine.playerEquipmentSetRef(_ref.playerEquipmentSetRef.ID)
	return ref.playerEquipmentSetRef.engine.EquipmentSet(ref.playerEquipmentSetRef.ReferencedElementID)
}`
const _Get_EquipmentSetEquipmentRef_func string = `func (_ref EquipmentSetEquipmentRef) Get() Item {
	ref := _ref.equipmentSetEquipmentRef.engine.equipmentSetEquipmentRef(_ref.equipmentSetEquipmentRef.ID)
	return ref.equipmentSetEquipmentRef.engine.Item(ref.equipmentSetEquipmentRef.ReferencedElementID)
}`
const _IsSet_PlayerTargetRef_func string = `func (_ref PlayerTargetRef) IsSet() (PlayerTargetRef, bool) {
	ref := _ref.playerTargetRef.engine.playerTargetRef(_ref.playerTargetRef.ID)
	return ref, ref.playerTargetRef.ID != 0
}`
const _Unset_PlayerTargetRef_func string = `func (_ref PlayerTargetRef) Unset() {
	ref := _ref.playerTargetRef.engine.playerTargetRef(_ref.playerTargetRef.ID)
	if ref.playerTargetRef.OperationKind == OperationKindDelete {
		return
	}
	ref.playerTargetRef.engine.deletePlayerTargetRef(ref.playerTargetRef.ID)
	parent := ref.playerTargetRef.engine.Player(ref.playerTargetRef.ParentID).player
	if parent.OperationKind == OperationKindDelete {
		return
	}
	parent.Target = 0
	parent.OperationKind = OperationKindUpdate
	ref.playerTargetRef.engine.Patch.Player[parent.ID] = parent
}`
const _Get_PlayerTargetRef_func string = `func (_ref PlayerTargetRef) Get() AnyOfPlayer_ZoneItemRef {
	ref := _ref.playerTargetRef.engine.playerTargetRef(_ref.playerTargetRef.ID)
	anyContainer := ref.playerTargetRef.engine.anyOfPlayer_ZoneItem(ref.playerTargetRef.ReferencedElementID)
	return anyContainer
}`
const _Get_PlayerTargetedByRef_func string = `func (_ref PlayerTargetedByRef) Get() AnyOfPlayer_ZoneItemRef {
	ref := _ref.playerTargetedByRef.engine.playerTargetedByRef(_ref.playerTargetedByRef.ID)
	anyContainer := ref.playerTargetedByRef.engine.anyOfPlayer_ZoneItem(ref.playerTargetedByRef.ReferencedElementID)
	return anyContainer
}`
const dereferenceAttackEventTargetRefs_Engine_func string = `func (engine *Engine) dereferenceAttackEventTargetRefs(playerID PlayerID) {
	allAttackEventTargetRefIDs := engine.allAttackEventTargetRefIDs()
	for _, refID := range allAttackEventTargetRefIDs {
		ref := engine.attackEventTargetRef(refID)
		if ref.attackEventTargetRef.ReferencedElementID == playerID {
			ref.Unset()
		}
	}
	attackEventTargetRefIDSlicePool.Put(allAttackEventTargetRefIDs)
}`
const dereferenceItemBoundToRefs_Engine_func string = `func (engine *Engine) dereferenceItemBoundToRefs(playerID PlayerID) {
	allItemBoundToRefIDs := engine.allItemBoundToRefIDs()
	for _, refID := range allItemBoundToRefIDs {
		ref := engine.itemBoundToRef(refID)
		if ref.itemBoundToRef.ReferencedElementID == playerID {
			ref.Unset()
		}
	}
	itemBoundToRefIDSlicePool.Put(allItemBoundToRefIDs)
}`
const dereferenceEquipmentSetEquipmentRefs_Engine_func string = `func (engine *Engine) dereferenceEquipmentSetEquipmentRefs(itemID ItemID) {
	allEquipmentSetEquipmentRefIDs := engine.allEquipmentSetEquipmentRefIDs()
	for _, refID := range allEquipmentSetEquipmentRefIDs {
		ref := engine.equipmentSetEquipmentRef(refID)
		if ref.equipmentSetEquipmentRef.ReferencedElementID == itemID {
			parent := engine.EquipmentSet(ref.equipmentSetEquipmentRef.ParentID)
			parent.RemoveEquipment(itemID)
		}
	}
	equipmentSetEquipmentRefIDSlicePool.Put(allEquipmentSetEquipmentRefIDs)
}`
const dereferencePlayerGuildMemberRefs_Engine_func string = `func (engine *Engine) dereferencePlayerGuildMemberRefs(playerID PlayerID) {
	allPlayerGuildMemberRefIDs := engine.allPlayerGuildMemberRefIDs()
	for _, refID := range allPlayerGuildMemberRefIDs {
		ref := engine.playerGuildMemberRef(refID)
		if ref.playerGuildMemberRef.ReferencedElementID == playerID {
			parent := engine.Player(ref.playerGuildMemberRef.ParentID)
			parent.RemoveGuildMember(playerID)
		}
	}
	playerGuildMemberRefIDSlicePool.Put(allPlayerGuildMemberRefIDs)
}`
const dereferencePlayerEquipmentSetRefs_Engine_func string = `func (engine *Engine) dereferencePlayerEquipmentSetRefs(equipmentSetID EquipmentSetID) {
	allPlayerEquipmentSetRefIDs := engine.allPlayerEquipmentSetRefIDs()
	for _, refID := range allPlayerEquipmentSetRefIDs {
		ref := engine.playerEquipmentSetRef(refID)
		if ref.playerEquipmentSetRef.ReferencedElementID == equipmentSetID {
			parent := engine.Player(ref.playerEquipmentSetRef.ParentID)
			parent.RemoveEquipmentSet(equipmentSetID)
		}
	}
	playerEquipmentSetRefIDSlicePool.Put(allPlayerEquipmentSetRefIDs)
}`
const dereferencePlayerTargetRefsPlayer_Engine_func string = `func (engine *Engine) dereferencePlayerTargetRefsPlayer(playerID PlayerID) {
	allPlayerTargetRefIDs := engine.allPlayerTargetRefIDs()
	for _, refID := range allPlayerTargetRefIDs {
		ref := engine.playerTargetRef(refID)
		anyContainer := ref.playerTargetRef.engine.anyOfPlayer_ZoneItem(ref.playerTargetRef.ReferencedElementID)
		if anyContainer.anyOfPlayer_ZoneItem.ElementKind != ElementKindPlayer {
			continue
		}
		if PlayerID(anyContainer.anyOfPlayer_ZoneItem.ChildID) == playerID {
			ref.Unset()
		}
	}
	playerTargetRefIDSlicePool.Put(allPlayerTargetRefIDs)
}`
const dereferencePlayerTargetRefsZoneItem_Engine_func string = `func (engine *Engine) dereferencePlayerTargetRefsZoneItem(zoneItemID ZoneItemID) {
	allPlayerTargetRefIDs := engine.allPlayerTargetRefIDs()
	for _, refID := range allPlayerTargetRefIDs {
		ref := engine.playerTargetRef(refID)
		anyContainer := ref.playerTargetRef.engine.anyOfPlayer_ZoneItem(ref.playerTargetRef.ReferencedElementID)
		if anyContainer.anyOfPlayer_ZoneItem.ElementKind != ElementKindZoneItem {
			continue
		}
		if ZoneItemID(anyContainer.anyOfPlayer_ZoneItem.ChildID) == zoneItemID {
			ref.Unset()
		}
	}
	playerTargetRefIDSlicePool.Put(allPlayerTargetRefIDs)
}`
const dereferencePlayerTargetedByRefsPlayer_Engine_func string = `func (engine *Engine) dereferencePlayerTargetedByRefsPlayer(playerID PlayerID) {
	allPlayerTargetedByRefIDs := engine.allPlayerTargetedByRefIDs()
	for _, refID := range allPlayerTargetedByRefIDs {
		ref := engine.playerTargetedByRef(refID)
		anyContainer := ref.playerTargetedByRef.engine.anyOfPlayer_ZoneItem(ref.playerTargetedByRef.ReferencedElementID)
		if anyContainer.anyOfPlayer_ZoneItem.ElementKind != ElementKindPlayer {
			continue
		}
		if PlayerID(anyContainer.anyOfPlayer_ZoneItem.ChildID) == playerID {
			parent := engine.Player(ref.playerTargetedByRef.ParentID)
			parent.RemoveTargetedByPlayer(playerID)
		}
	}
	playerTargetedByRefIDSlicePool.Put(allPlayerTargetedByRefIDs)
}`
const dereferencePlayerTargetedByRefsZoneItem_Engine_func string = `func (engine *Engine) dereferencePlayerTargetedByRefsZoneItem(zoneItemID ZoneItemID) {
	allPlayerTargetedByRefIDs := engine.allPlayerTargetedByRefIDs()
	for _, refID := range allPlayerTargetedByRefIDs {
		ref := engine.playerTargetedByRef(refID)
		anyContainer := ref.playerTargetedByRef.engine.anyOfPlayer_ZoneItem(ref.playerTargetedByRef.ReferencedElementID)
		if anyContainer.anyOfPlayer_ZoneItem.ElementKind != ElementKindZoneItem {
			continue
		}
		if ZoneItemID(anyContainer.anyOfPlayer_ZoneItem.ChildID) == zoneItemID {
			parent := engine.Player(ref.playerTargetedByRef.ParentID)
			parent.RemoveTargetedByZoneItem(zoneItemID)
		}
	}
	playerTargetedByRefIDSlicePool.Put(allPlayerTargetedByRefIDs)
}`
const _RemovePlayer_Zone_func string = `func (_zone Zone) RemovePlayer(playerToRemove PlayerID) Zone {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	if zone.zone.OperationKind == OperationKindDelete {
		return zone
	}
	if _, ok := zone.zone.engine.Patch.Zone[zone.zone.ID]; !ok {
		cp := make([]PlayerID, len(zone.zone.Players))
		copy(cp, zone.zone.Players)
		zone.zone.Players = cp
	}
	for i, playerID := range zone.zone.Players {
		if playerID != playerToRemove {
			continue
		}
		zone.zone.Players[i] = zone.zone.Players[len(zone.zone.Players)-1]
		zone.zone.Players[len(zone.zone.Players)-1] = 0
		zone.zone.Players = zone.zone.Players[:len(zone.zone.Players)-1]
		zone.zone.engine.deletePlayer(playerID)
		zone.zone.OperationKind = OperationKindUpdate
		zone.zone.engine.Patch.Zone[zone.zone.ID] = zone.zone
		break
	}
	return zone
}`
const _RemoveItem_Zone_func string = `func (_zone Zone) RemoveItem(itemToRemove ZoneItemID) Zone {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	if zone.zone.OperationKind == OperationKindDelete {
		return zone
	}
	if _, ok := zone.zone.engine.Patch.Zone[zone.zone.ID]; !ok {
		cp := make([]ZoneItemID, len(zone.zone.Items))
		copy(cp, zone.zone.Items)
		zone.zone.Items = cp
	}
	for i, zoneItemID := range zone.zone.Items {
		if zoneItemID != itemToRemove {
			continue
		}
		zone.zone.Items[i] = zone.zone.Items[len(zone.zone.Items)-1]
		zone.zone.Items[len(zone.zone.Items)-1] = 0
		zone.zone.Items = zone.zone.Items[:len(zone.zone.Items)-1]
		zone.zone.engine.deleteZoneItem(zoneItemID)
		zone.zone.OperationKind = OperationKindUpdate
		zone.zone.engine.Patch.Zone[zone.zone.ID] = zone.zone
		break
	}
	return zone
}`
const _RemoveInteractableItem_Zone_func string = `func (_zone Zone) RemoveInteractableItem(itemToRemove ItemID) Zone {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	if zone.zone.OperationKind == OperationKindDelete {
		return zone
	}
	if _, ok := zone.zone.engine.Patch.Zone[zone.zone.ID]; !ok {
		cp := make([]AnyOfItem_Player_ZoneItemID, len(zone.zone.Interactables))
		copy(cp, zone.zone.Interactables)
		zone.zone.Interactables = cp
	}
	for i, id := range zone.zone.Interactables {
		if childID := zone.zone.engine.anyOfItem_Player_ZoneItem(id).anyOfItem_Player_ZoneItem.ChildID; ItemID(childID) != itemToRemove {
			continue
		}
		zone.zone.Interactables[i] = zone.zone.Interactables[len(zone.zone.Interactables)-1]
		zone.zone.Interactables[len(zone.zone.Interactables)-1] = 0
		zone.zone.Interactables = zone.zone.Interactables[:len(zone.zone.Interactables)-1]
		zone.zone.engine.deleteAnyOfItem_Player_ZoneItem(id, true)
		zone.zone.OperationKind = OperationKindUpdate
		zone.zone.engine.Patch.Zone[zone.zone.ID] = zone.zone
		break
	}
	return zone
}`
const _RemoveInteractablePlayer_Zone_func string = `func (_zone Zone) RemoveInteractablePlayer(playerToRemove PlayerID) Zone {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	if zone.zone.OperationKind == OperationKindDelete {
		return zone
	}
	if _, ok := zone.zone.engine.Patch.Zone[zone.zone.ID]; !ok {
		cp := make([]AnyOfItem_Player_ZoneItemID, len(zone.zone.Interactables))
		copy(cp, zone.zone.Interactables)
		zone.zone.Interactables = cp
	}
	for i, id := range zone.zone.Interactables {
		if childID := zone.zone.engine.anyOfItem_Player_ZoneItem(id).anyOfItem_Player_ZoneItem.ChildID; PlayerID(childID) != playerToRemove {
			continue
		}
		zone.zone.Interactables[i] = zone.zone.Interactables[len(zone.zone.Interactables)-1]
		zone.zone.Interactables[len(zone.zone.Interactables)-1] = 0
		zone.zone.Interactables = zone.zone.Interactables[:len(zone.zone.Interactables)-1]
		zone.zone.engine.deleteAnyOfItem_Player_ZoneItem(id, true)
		zone.zone.OperationKind = OperationKindUpdate
		zone.zone.engine.Patch.Zone[zone.zone.ID] = zone.zone
		break
	}
	return zone
}`
const _RemoveInteractableZoneItem_Zone_func string = `func (_zone Zone) RemoveInteractableZoneItem(zoneItemToRemove ZoneItemID) Zone {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	if zone.zone.OperationKind == OperationKindDelete {
		return zone
	}
	if _, ok := zone.zone.engine.Patch.Zone[zone.zone.ID]; !ok {
		cp := make([]AnyOfItem_Player_ZoneItemID, len(zone.zone.Interactables))
		copy(cp, zone.zone.Interactables)
		zone.zone.Interactables = cp
	}
	for i, id := range zone.zone.Interactables {
		if childID := zone.zone.engine.anyOfItem_Player_ZoneItem(id).anyOfItem_Player_ZoneItem.ChildID; ZoneItemID(childID) != zoneItemToRemove {
			continue
		}
		zone.zone.Interactables[i] = zone.zone.Interactables[len(zone.zone.Interactables)-1]
		zone.zone.Interactables[len(zone.zone.Interactables)-1] = 0
		zone.zone.Interactables = zone.zone.Interactables[:len(zone.zone.Interactables)-1]
		zone.zone.engine.deleteAnyOfItem_Player_ZoneItem(id, true)
		zone.zone.OperationKind = OperationKindUpdate
		zone.zone.engine.Patch.Zone[zone.zone.ID] = zone.zone
		break
	}
	return zone
}`
const _RemoveAction_Player_func string = `func (_player Player) RemoveAction(actionToRemove AttackEventID) Player {
	player := _player.player.engine.Player(_player.player.ID)
	if player.player.OperationKind == OperationKindDelete {
		return player
	}
	if _, ok := player.player.engine.Patch.Player[player.player.ID]; !ok {
		cp := make([]AttackEventID, len(player.player.Action))
		copy(cp, player.player.Action)
		player.player.Action = cp
	}
	for i, attackEventID := range player.player.Action {
		if attackEventID != actionToRemove {
			continue
		}
		player.player.Action[i] = player.player.Action[len(player.player.Action)-1]
		player.player.Action[len(player.player.Action)-1] = 0
		player.player.Action = player.player.Action[:len(player.player.Action)-1]
		player.player.engine.deleteAttackEvent(attackEventID)
		player.player.OperationKind = OperationKindUpdate
		player.player.engine.Patch.Player[player.player.ID] = player.player
		break
	}
	return player
}`
const _RemoveItem_Player_func string = `func (_player Player) RemoveItem(itemToRemove ItemID) Player {
	player := _player.player.engine.Player(_player.player.ID)
	if player.player.OperationKind == OperationKindDelete {
		return player
	}
	if _, ok := player.player.engine.Patch.Player[player.player.ID]; !ok {
		cp := make([]ItemID, len(player.player.Items))
		copy(cp, player.player.Items)
		player.player.Items = cp
	}
	for i, itemID := range player.player.Items {
		if itemID != itemToRemove {
			continue
		}
		player.player.Items[i] = player.player.Items[len(player.player.Items)-1]
		player.player.Items[len(player.player.Items)-1] = 0
		player.player.Items = player.player.Items[:len(player.player.Items)-1]
		player.player.engine.deleteItem(itemID)
		player.player.OperationKind = OperationKindUpdate
		player.player.engine.Patch.Player[player.player.ID] = player.player
		break
	}
	return player
}`
const _RemoveEquipmentSet_Player_func string = `func (_player Player) RemoveEquipmentSet(equipmentSetToRemove EquipmentSetID) Player {
	player := _player.player.engine.Player(_player.player.ID)
	if player.player.OperationKind == OperationKindDelete {
		return player
	}
	if _, ok := player.player.engine.Patch.Player[player.player.ID]; !ok {
		cp := make([]PlayerEquipmentSetRefID, len(player.player.EquipmentSets))
		copy(cp, player.player.EquipmentSets)
		player.player.EquipmentSets = cp
	}
	for i, id := range player.player.EquipmentSets {
		if childID := player.player.engine.playerEquipmentSetRef(id).playerEquipmentSetRef.ChildID; EquipmentSetID(childID) != equipmentSetToRemove {
			continue
		}
		player.player.EquipmentSets[i] = player.player.EquipmentSets[len(player.player.EquipmentSets)-1]
		player.player.EquipmentSets[len(player.player.EquipmentSets)-1] = 0
		player.player.EquipmentSets = player.player.EquipmentSets[:len(player.player.EquipmentSets)-1]
		player.player.engine.deletePlayerEquipmentSetRef(id)
		player.player.OperationKind = OperationKindUpdate
		player.player.engine.Patch.Player[player.player.ID] = player.player
		break
	}
	return player
}`
const _RemoveGuildMember_Player_func string = `func (_player Player) RemoveGuildMember(guildMemberToRemove PlayerID) Player {
	player := _player.player.engine.Player(_player.player.ID)
	if player.player.OperationKind == OperationKindDelete {
		return player
	}
	if _, ok := player.player.engine.Patch.Player[player.player.ID]; !ok {
		cp := make([]PlayerGuildMemberRefID, len(player.player.GuildMembers))
		copy(cp, player.player.GuildMembers)
		player.player.GuildMembers = cp
	}
	for i, id := range player.player.GuildMembers {
		if childID := player.player.engine.playerGuildMemberRef(id).playerGuildMemberRef.ChildID; PlayerID(childID) != guildMemberToRemove {
			continue
		}
		player.player.GuildMembers[i] = player.player.GuildMembers[len(player.player.GuildMembers)-1]
		player.player.GuildMembers[len(player.player.GuildMembers)-1] = 0
		player.player.GuildMembers = player.player.GuildMembers[:len(player.player.GuildMembers)-1]
		player.player.engine.deletePlayerGuildMemberRef(id)
		player.player.OperationKind = OperationKindUpdate
		player.player.engine.Patch.Player[player.player.ID] = player.player
		break
	}
	return player
}`
const _RemoveTargetedByZoneItem_Player_func string = `func (_player Player) RemoveTargetedByZoneItem(zoneItemToRemove ZoneItemID) Player {
	player := _player.player.engine.Player(_player.player.ID)
	if player.player.OperationKind == OperationKindDelete {
		return player
	}
	if _, ok := player.player.engine.Patch.Player[player.player.ID]; !ok {
		cp := make([]PlayerTargetedByRefID, len(player.player.TargetedBy))
		copy(cp, player.player.TargetedBy)
		player.player.TargetedBy = cp
	}
	for i, id := range player.player.TargetedBy {
		if childID := player.player.engine.playerTargetedByRef(id).playerTargetedByRef.ChildID; ZoneItemID(childID) != zoneItemToRemove {
			continue
		}
		player.player.TargetedBy[i] = player.player.TargetedBy[len(player.player.TargetedBy)-1]
		player.player.TargetedBy[len(player.player.TargetedBy)-1] = 0
		player.player.TargetedBy = player.player.TargetedBy[:len(player.player.TargetedBy)-1]
		player.player.engine.deletePlayerTargetedByRef(id)
		player.player.OperationKind = OperationKindUpdate
		player.player.engine.Patch.Player[player.player.ID] = player.player
		break
	}
	return player
}`
const _RemoveTargetedByPlayer_Player_func string = `func (_player Player) RemoveTargetedByPlayer(playerToRemove PlayerID) Player {
	player := _player.player.engine.Player(_player.player.ID)
	if player.player.OperationKind == OperationKindDelete {
		return player
	}
	if _, ok := player.player.engine.Patch.Player[player.player.ID]; !ok {
		cp := make([]PlayerTargetedByRefID, len(player.player.TargetedBy))
		copy(cp, player.player.TargetedBy)
		player.player.TargetedBy = cp
	}
	for i, id := range player.player.TargetedBy {
		if childID := player.player.engine.playerTargetedByRef(id).playerTargetedByRef.ChildID; PlayerID(childID) != playerToRemove {
			continue
		}
		player.player.TargetedBy[i] = player.player.TargetedBy[len(player.player.TargetedBy)-1]
		player.player.TargetedBy[len(player.player.TargetedBy)-1] = 0
		player.player.TargetedBy = player.player.TargetedBy[:len(player.player.TargetedBy)-1]
		player.player.engine.deletePlayerTargetedByRef(id)
		player.player.OperationKind = OperationKindUpdate
		player.player.engine.Patch.Player[player.player.ID] = player.player
		break
	}
	return player
}`
const _RemoveTag_Zone_func string = `func (_zone Zone) RemoveTag(tagToRemove string) Zone {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	if zone.zone.OperationKind == OperationKindDelete {
		return zone
	}
	if _, ok := zone.zone.engine.Patch.Zone[zone.zone.ID]; !ok {
		cp := make([]StringValueID, len(zone.zone.Tags))
		copy(cp, zone.zone.Tags)
		zone.zone.Tags = cp
	}
	for i, valID := range zone.zone.Tags {
		if zone.zone.engine.stringValue(valID).Value != tagToRemove {
			continue
		}
		zone.zone.Tags[i] = zone.zone.Tags[len(zone.zone.Tags)-1]
		zone.zone.Tags[len(zone.zone.Tags)-1] = 0
		zone.zone.Tags = zone.zone.Tags[:len(zone.zone.Tags)-1]
		zone.zone.engine.deleteStringValue(valID)
		zone.zone.OperationKind = OperationKindUpdate
		zone.zone.engine.Patch.Zone[zone.zone.ID] = zone.zone
		break
	}
	return zone
}`
const _RemoveEquipment_EquipmentSet_func string = `func (_equipmentSet EquipmentSet) RemoveEquipment(equipmentToRemove ItemID) EquipmentSet {
	equipmentSet := _equipmentSet.equipmentSet.engine.EquipmentSet(_equipmentSet.equipmentSet.ID)
	if equipmentSet.equipmentSet.OperationKind == OperationKindDelete {
		return equipmentSet
	}
	if _, ok := equipmentSet.equipmentSet.engine.Patch.EquipmentSet[equipmentSet.equipmentSet.ID]; !ok {
		cp := make([]EquipmentSetEquipmentRefID, len(equipmentSet.equipmentSet.Equipment))
		copy(cp, equipmentSet.equipmentSet.Equipment)
		equipmentSet.equipmentSet.Equipment = cp
	}
	for i, id := range equipmentSet.equipmentSet.Equipment {
		if childID := equipmentSet.equipmentSet.engine.equipmentSetEquipmentRef(id).equipmentSetEquipmentRef.ChildID; ItemID(childID) != equipmentToRemove {
			continue
		}
		equipmentSet.equipmentSet.Equipment[i] = equipmentSet.equipmentSet.Equipment[len(equipmentSet.equipmentSet.Equipment)-1]
		equipmentSet.equipmentSet.Equipment[len(equipmentSet.equipmentSet.Equipment)-1] = 0
		equipmentSet.equipmentSet.Equipment = equipmentSet.equipmentSet.Equipment[:len(equipmentSet.equipmentSet.Equipment)-1]
		equipmentSet.equipmentSet.engine.deleteEquipmentSetEquipmentRef(id)
		equipmentSet.equipmentSet.OperationKind = OperationKindUpdate
		equipmentSet.equipmentSet.engine.Patch.EquipmentSet[equipmentSet.equipmentSet.ID] = equipmentSet.equipmentSet
		break
	}
	return equipmentSet
}`
const setBoolValue_Engine_func string = `func (engine *Engine) setBoolValue(id BoolValueID, val bool) {
	boolValue := engine.boolValue(id)
	if boolValue.OperationKind == OperationKindDelete {
		return
	}
	if boolValue.Value == val {
		return
	}
	boolValue.Value = val
	boolValue.OperationKind = OperationKindUpdate
	engine.Patch.BoolValue[id] = boolValue
}`
const setFloatValue_Engine_func string = `func (engine *Engine) setFloatValue(id FloatValueID, val float64) {
	floatValue := engine.floatValue(id)
	if floatValue.OperationKind == OperationKindDelete {
		return
	}
	if floatValue.Value == val {
		return
	}
	floatValue.Value = val
	floatValue.OperationKind = OperationKindUpdate
	engine.Patch.FloatValue[id] = floatValue
}`
const setIntValue_Engine_func string = `func (engine *Engine) setIntValue(id IntValueID, val int64) {
	intValue := engine.intValue(id)
	if intValue.OperationKind == OperationKindDelete {
		return
	}
	if intValue.Value == val {
		return
	}
	intValue.Value = val
	intValue.OperationKind = OperationKindUpdate
	engine.Patch.IntValue[id] = intValue
}`
const setStringValue_Engine_func string = `func (engine *Engine) setStringValue(id StringValueID, val string) {
	stringValue := engine.stringValue(id)
	if stringValue.OperationKind == OperationKindDelete {
		return
	}
	if stringValue.Value == val {
		return
	}
	stringValue.Value = val
	stringValue.OperationKind = OperationKindUpdate
	engine.Patch.StringValue[id] = stringValue
}`
const _SetLevel_GearScore_func string = `func (_gearScore GearScore) SetLevel(newLevel int64) GearScore {
	gearScore := _gearScore.gearScore.engine.GearScore(_gearScore.gearScore.ID)
	if gearScore.gearScore.OperationKind == OperationKindDelete {
		return gearScore
	}
	gearScore.gearScore.engine.setIntValue(gearScore.gearScore.Level, newLevel)
	return gearScore
}`
const _SetScore_GearScore_func string = `func (_gearScore GearScore) SetScore(newScore int64) GearScore {
	gearScore := _gearScore.gearScore.engine.GearScore(_gearScore.gearScore.ID)
	if gearScore.gearScore.OperationKind == OperationKindDelete {
		return gearScore
	}
	gearScore.gearScore.engine.setIntValue(gearScore.gearScore.Score, newScore)
	return gearScore
}`
const _SetX_Position_func string = `func (_position Position) SetX(newX float64) Position {
	position := _position.position.engine.Position(_position.position.ID)
	if position.position.OperationKind == OperationKindDelete {
		return position
	}
	position.position.engine.setFloatValue(position.position.X, newX)
	return position
}`
const _SetY_Position_func string = `func (_position Position) SetY(newY float64) Position {
	position := _position.position.engine.Position(_position.position.ID)
	if position.position.OperationKind == OperationKindDelete {
		return position
	}
	position.position.engine.setFloatValue(position.position.Y, newY)
	return position
}`
const _SetName_Item_func string = `func (_item Item) SetName(newName string) Item {
	item := _item.item.engine.Item(_item.item.ID)
	if item.item.OperationKind == OperationKindDelete {
		return item
	}
	item.item.engine.setStringValue(item.item.Name, newName)
	return item
}`
const _SetBoundTo_Item_func string = `func (_item Item) SetBoundTo(playerID PlayerID) Item {
	item := _item.item.engine.Item(_item.item.ID)
	if item.item.OperationKind == OperationKindDelete {
		return item
	}
	if item.item.engine.Player(playerID).player.OperationKind == OperationKindDelete {
		return item
	}
	if item.item.BoundTo != 0 {
		if childID := item.item.engine.itemBoundToRef(item.item.BoundTo).itemBoundToRef.ChildID; PlayerID(childID) == playerID {
			return item
		}
		item.item.engine.deleteItemBoundToRef(item.item.BoundTo)
	}
	ref := item.item.engine.createItemBoundToRef(item.item.Path, item_boundToIdentifier, playerID, item.item.ID, int(playerID))
	item.item.BoundTo = ref.ID
	item.item.OperationKind = OperationKindUpdate
	item.item.engine.Patch.Item[item.item.ID] = item.item
	return item
}`
const _SetTarget_AttackEvent_func string = `func (_attackEvent AttackEvent) SetTarget(playerID PlayerID) AttackEvent {
	attackEvent := _attackEvent.attackEvent.engine.AttackEvent(_attackEvent.attackEvent.ID)
	if attackEvent.attackEvent.OperationKind == OperationKindDelete {
		return attackEvent
	}
	if attackEvent.attackEvent.engine.Player(playerID).player.OperationKind == OperationKindDelete {
		return attackEvent
	}
	if attackEvent.attackEvent.Target != 0 {
		if childID := attackEvent.attackEvent.engine.attackEventTargetRef(attackEvent.attackEvent.Target).attackEventTargetRef.ChildID; PlayerID(childID) == playerID {
			return attackEvent
		}
		attackEvent.attackEvent.engine.deleteAttackEventTargetRef(attackEvent.attackEvent.Target)
	}
	ref := attackEvent.attackEvent.engine.createAttackEventTargetRef(attackEvent.attackEvent.Path, attackEvent_targetIdentifier, playerID, attackEvent.attackEvent.ID, int(playerID))
	attackEvent.attackEvent.Target = ref.ID
	attackEvent.attackEvent.OperationKind = OperationKindUpdate
	attackEvent.attackEvent.engine.Patch.AttackEvent[attackEvent.attackEvent.ID] = attackEvent.attackEvent
	return attackEvent
}`
const _SetName_EquipmentSet_func string = `func (_equipmentSet EquipmentSet) SetName(newName string) EquipmentSet {
	equipmentSet := _equipmentSet.equipmentSet.engine.EquipmentSet(_equipmentSet.equipmentSet.ID)
	if equipmentSet.equipmentSet.OperationKind == OperationKindDelete {
		return equipmentSet
	}
	equipmentSet.equipmentSet.engine.setStringValue(equipmentSet.equipmentSet.Name, newName)
	return equipmentSet
}`
const _SetTargetPlayer_Player_func string = `func (_player Player) SetTargetPlayer(playerID PlayerID) Player {
	player := _player.player.engine.Player(_player.player.ID)
	if player.player.OperationKind == OperationKindDelete {
		return player
	}
	if player.player.engine.Player(playerID).player.OperationKind == OperationKindDelete {
		return player
	}
	if player.player.Target != 0 {
		if childID := player.player.engine.playerTargetRef(player.player.Target).playerTargetRef.ChildID; PlayerID(childID) == playerID {
			return player
		}
		player.player.engine.deletePlayerTargetRef(player.player.Target)
	}
	anyContainer := player.player.engine.createAnyOfPlayer_ZoneItem(int(player.player.ID), int(playerID), ElementKindPlayer, player.player.Path, player_targetIdentifier)
	ref := player.player.engine.createPlayerTargetRef(player.player.Path, player_targetIdentifier, anyContainer.anyOfPlayer_ZoneItem.ID, player.player.ID, ElementKindPlayer, int(playerID))
	player.player.Target = ref.ID
	player.player.OperationKind = OperationKindUpdate
	player.player.engine.Patch.Player[player.player.ID] = player.player
	return player
}`
const _SetTargetZoneItem_Player_func string = `func (_player Player) SetTargetZoneItem(zoneItemID ZoneItemID) Player {
	player := _player.player.engine.Player(_player.player.ID)
	if player.player.OperationKind == OperationKindDelete {
		return player
	}
	if player.player.engine.ZoneItem(zoneItemID).zoneItem.OperationKind == OperationKindDelete {
		return player
	}
	if player.player.Target != 0 {
		if childID := player.player.engine.playerTargetRef(player.player.Target).playerTargetRef.ChildID; ZoneItemID(childID) == zoneItemID {
			return player
		}
		player.player.engine.deletePlayerTargetRef(player.player.Target)
	}
	anyContainer := player.player.engine.createAnyOfPlayer_ZoneItem(int(player.player.ID), int(zoneItemID), ElementKindZoneItem, player.player.Path, player_targetIdentifier)
	ref := player.player.engine.createPlayerTargetRef(player.player.Path, player_targetIdentifier, anyContainer.anyOfPlayer_ZoneItem.ID, player.player.ID, ElementKindZoneItem, int(zoneItemID))
	player.player.Target = ref.ID
	player.player.OperationKind = OperationKindUpdate
	player.player.engine.Patch.Player[player.player.ID] = player.player
	return player
}`
const _BoolValueID_type string = `type BoolValueID int`
const _FloatValueID_type string = `type FloatValueID int`
const _IntValueID_type string = `type IntValueID int`
const _StringValueID_type string = `type StringValueID int`
const _AttackEventID_type string = `type AttackEventID int`
const _EquipmentSetID_type string = `type EquipmentSetID int`
const _GearScoreID_type string = `type GearScoreID int`
const _ItemID_type string = `type ItemID int`
const _PlayerID_type string = `type PlayerID int`
const _PositionID_type string = `type PositionID int`
const _ZoneID_type string = `type ZoneID int`
const _ZoneItemID_type string = `type ZoneItemID int`
const _AttackEventTargetRefID_type string = `type AttackEventTargetRefID int`
const _PlayerGuildMemberRefID_type string = `type PlayerGuildMemberRefID int`
const _ItemBoundToRefID_type string = `type ItemBoundToRefID int`
const _EquipmentSetEquipmentRefID_type string = `type EquipmentSetEquipmentRefID int`
const _PlayerEquipmentSetRefID_type string = `type PlayerEquipmentSetRefID int`
const _AnyOfItem_Player_ZoneItemID_type string = `type AnyOfItem_Player_ZoneItemID int`
const _AnyOfPlayer_ZoneItemID_type string = `type AnyOfPlayer_ZoneItemID int`
const _AnyOfPlayer_PositionID_type string = `type AnyOfPlayer_PositionID int`
const _PlayerTargetRefID_type string = `type PlayerTargetRefID int`
const _PlayerTargetedByRefID_type string = `type PlayerTargetedByRefID int`
const _State_type string = `type State struct {
	BoolValue			map[BoolValueID]boolValue					` + "`" + `json:"boolValue"` + "`" + `
	FloatValue			map[FloatValueID]floatValue					` + "`" + `json:"floatValue"` + "`" + `
	IntValue			map[IntValueID]intValue						` + "`" + `json:"intValue"` + "`" + `
	StringValue			map[StringValueID]stringValue					` + "`" + `json:"stringValue"` + "`" + `
	AttackEvent			map[AttackEventID]attackEventCore				` + "`" + `json:"attackEvent"` + "`" + `
	EquipmentSet			map[EquipmentSetID]equipmentSetCore				` + "`" + `json:"equipmentSet"` + "`" + `
	GearScore			map[GearScoreID]gearScoreCore					` + "`" + `json:"gearScore"` + "`" + `
	Item				map[ItemID]itemCore						` + "`" + `json:"item"` + "`" + `
	Player				map[PlayerID]playerCore						` + "`" + `json:"player"` + "`" + `
	Position			map[PositionID]positionCore					` + "`" + `json:"position"` + "`" + `
	Zone				map[ZoneID]zoneCore						` + "`" + `json:"zone"` + "`" + `
	ZoneItem			map[ZoneItemID]zoneItemCore					` + "`" + `json:"zoneItem"` + "`" + `
	AttackEventTargetRef		map[AttackEventTargetRefID]attackEventTargetRefCore		` + "`" + `json:"attackEventTargetRef"` + "`" + `
	EquipmentSetEquipmentRef	map[EquipmentSetEquipmentRefID]equipmentSetEquipmentRefCore	` + "`" + `json:"equipmentSetEquipmentRef"` + "`" + `
	ItemBoundToRef			map[ItemBoundToRefID]itemBoundToRefCore				` + "`" + `json:"itemBoundToRef"` + "`" + `
	PlayerEquipmentSetRef		map[PlayerEquipmentSetRefID]playerEquipmentSetRefCore		` + "`" + `json:"playerEquipmentSetRef"` + "`" + `
	PlayerGuildMemberRef		map[PlayerGuildMemberRefID]playerGuildMemberRefCore		` + "`" + `json:"playerGuildMemberRef"` + "`" + `
	PlayerTargetRef			map[PlayerTargetRefID]playerTargetRefCore			` + "`" + `json:"playerTargetRef"` + "`" + `
	PlayerTargetedByRef		map[PlayerTargetedByRefID]playerTargetedByRefCore		` + "`" + `json:"playerTargetedByRef"` + "`" + `
	AnyOfPlayer_Position		map[AnyOfPlayer_PositionID]anyOfPlayer_PositionCore		` + "`" + `json:"anyOfPlayer_Position"` + "`" + `
	AnyOfPlayer_ZoneItem		map[AnyOfPlayer_ZoneItemID]anyOfPlayer_ZoneItemCore		` + "`" + `json:"anyOfPlayer_ZoneItem"` + "`" + `
	AnyOfItem_Player_ZoneItem	map[AnyOfItem_Player_ZoneItemID]anyOfItem_Player_ZoneItemCore	` + "`" + `json:"anyOfItem_Player_ZoneItem"` + "`" + `
}`
const newState_func string = `func newState() *State {
	return &State{BoolValue: make(map[BoolValueID]boolValue), FloatValue: make(map[FloatValueID]floatValue), IntValue: make(map[IntValueID]intValue), StringValue: make(map[StringValueID]stringValue), AttackEvent: make(map[AttackEventID]attackEventCore), EquipmentSet: make(map[EquipmentSetID]equipmentSetCore), GearScore: make(map[GearScoreID]gearScoreCore), Item: make(map[ItemID]itemCore), Player: make(map[PlayerID]playerCore), Position: make(map[PositionID]positionCore), Zone: make(map[ZoneID]zoneCore), ZoneItem: make(map[ZoneItemID]zoneItemCore), AttackEventTargetRef: make(map[AttackEventTargetRefID]attackEventTargetRefCore), EquipmentSetEquipmentRef: make(map[EquipmentSetEquipmentRefID]equipmentSetEquipmentRefCore), ItemBoundToRef: make(map[ItemBoundToRefID]itemBoundToRefCore), PlayerEquipmentSetRef: make(map[PlayerEquipmentSetRefID]playerEquipmentSetRefCore), PlayerGuildMemberRef: make(map[PlayerGuildMemberRefID]playerGuildMemberRefCore), PlayerTargetRef: make(map[PlayerTargetRefID]playerTargetRefCore), PlayerTargetedByRef: make(map[PlayerTargetedByRefID]playerTargetedByRefCore), AnyOfPlayer_Position: make(map[AnyOfPlayer_PositionID]anyOfPlayer_PositionCore), AnyOfPlayer_ZoneItem: make(map[AnyOfPlayer_ZoneItemID]anyOfPlayer_ZoneItemCore), AnyOfItem_Player_ZoneItem: make(map[AnyOfItem_Player_ZoneItemID]anyOfItem_Player_ZoneItemCore)}
}`
const boolValue_type string = `type boolValue struct {
	ID		BoolValueID	` + "`" + `json:"id"` + "`" + `
	Value		bool		` + "`" + `json:"value"` + "`" + `
	OperationKind	OperationKind	` + "`" + `json:"operationKind"` + "`" + `
	JSONPath	string		` + "`" + `json:"jsonPath"` + "`" + `
	Path		path		` + "`" + `json:"path"` + "`" + `
	engine		*Engine
}`
const intValue_type string = `type intValue struct {
	ID		IntValueID	` + "`" + `json:"id"` + "`" + `
	Value		int64		` + "`" + `json:"value"` + "`" + `
	OperationKind	OperationKind	` + "`" + `json:"operationKind"` + "`" + `
	JSONPath	string		` + "`" + `json:"jsonPath"` + "`" + `
	Path		path		` + "`" + `json:"path"` + "`" + `
	engine		*Engine
}`
const floatValue_type string = `type floatValue struct {
	ID		FloatValueID	` + "`" + `json:"id"` + "`" + `
	Value		float64		` + "`" + `json:"value"` + "`" + `
	OperationKind	OperationKind	` + "`" + `json:"operationKind"` + "`" + `
	JSONPath	string		` + "`" + `json:"jsonPath"` + "`" + `
	Path		path		` + "`" + `json:"path"` + "`" + `
	engine		*Engine
}`
const stringValue_type string = `type stringValue struct {
	ID		StringValueID	` + "`" + `json:"id"` + "`" + `
	Value		string		` + "`" + `json:"value"` + "`" + `
	OperationKind	OperationKind	` + "`" + `json:"operationKind"` + "`" + `
	JSONPath	string		` + "`" + `json:"jsonPath"` + "`" + `
	Path		path		` + "`" + `json:"path"` + "`" + `
	engine		*Engine
}`
const attackEventCore_type string = `type attackEventCore struct {
	ID		AttackEventID		` + "`" + `json:"id"` + "`" + `
	Target		AttackEventTargetRefID	` + "`" + `json:"target"` + "`" + `
	OperationKind	OperationKind		` + "`" + `json:"operationKind"` + "`" + `
	HasParent	bool			` + "`" + `json:"hasParent"` + "`" + `
	JSONPath	string			` + "`" + `json:"jsonPath"` + "`" + `
	Path		path			` + "`" + `json:"path"` + "`" + `
	engine		*Engine
}`
const _AttackEvent_type string = `type AttackEvent struct{ attackEvent attackEventCore }`
const zoneCore_type string = `type zoneCore struct {
	ID		ZoneID				` + "`" + `json:"id"` + "`" + `
	Interactables	[]AnyOfItem_Player_ZoneItemID	` + "`" + `json:"interactables"` + "`" + `
	Items		[]ZoneItemID			` + "`" + `json:"items"` + "`" + `
	Players		[]PlayerID			` + "`" + `json:"players"` + "`" + `
	Tags		[]StringValueID			` + "`" + `json:"tags"` + "`" + `
	OperationKind	OperationKind			` + "`" + `json:"operationKind"` + "`" + `
	HasParent	bool				` + "`" + `json:"hasParent"` + "`" + `
	JSONPath	string				` + "`" + `json:"jsonPath"` + "`" + `
	Path		path				` + "`" + `json:"path"` + "`" + `
	engine		*Engine
}`
const _Zone_type string = `type Zone struct{ zone zoneCore }`
const zoneItemCore_type string = `type zoneItemCore struct {
	ID		ZoneItemID	` + "`" + `json:"id"` + "`" + `
	Item		ItemID		` + "`" + `json:"item"` + "`" + `
	Position	PositionID	` + "`" + `json:"position"` + "`" + `
	OperationKind	OperationKind	` + "`" + `json:"operationKind"` + "`" + `
	HasParent	bool		` + "`" + `json:"hasParent"` + "`" + `
	JSONPath	string		` + "`" + `json:"jsonPath"` + "`" + `
	Path		path		` + "`" + `json:"path"` + "`" + `
	engine		*Engine
}`
const _ZoneItem_type string = `type ZoneItem struct{ zoneItem zoneItemCore }`
const itemCore_type string = `type itemCore struct {
	ID		ItemID			` + "`" + `json:"id"` + "`" + `
	BoundTo		ItemBoundToRefID	` + "`" + `json:"boundTo"` + "`" + `
	GearScore	GearScoreID		` + "`" + `json:"gearScore"` + "`" + `
	Name		StringValueID		` + "`" + `json:"name"` + "`" + `
	Origin		AnyOfPlayer_PositionID	` + "`" + `json:"origin"` + "`" + `
	OperationKind	OperationKind		` + "`" + `json:"operationKind"` + "`" + `
	HasParent	bool			` + "`" + `json:"hasParent"` + "`" + `
	JSONPath	string			` + "`" + `json:"jsonPath"` + "`" + `
	Path		path			` + "`" + `json:"path"` + "`" + `
	engine		*Engine
}`
const _Item_type string = `type Item struct{ item itemCore }`
const playerCore_type string = `type playerCore struct {
	ID		PlayerID			` + "`" + `json:"id"` + "`" + `
	Action		[]AttackEventID			` + "`" + `json:"action"` + "`" + `
	EquipmentSets	[]PlayerEquipmentSetRefID	` + "`" + `json:"equipmentSets"` + "`" + `
	GearScore	GearScoreID			` + "`" + `json:"gearScore"` + "`" + `
	GuildMembers	[]PlayerGuildMemberRefID	` + "`" + `json:"guildMembers"` + "`" + `
	Items		[]ItemID			` + "`" + `json:"items"` + "`" + `
	Position	PositionID			` + "`" + `json:"position"` + "`" + `
	Target		PlayerTargetRefID		` + "`" + `json:"target"` + "`" + `
	TargetedBy	[]PlayerTargetedByRefID		` + "`" + `json:"targetedBy"` + "`" + `
	OperationKind	OperationKind			` + "`" + `json:"operationKind"` + "`" + `
	HasParent	bool				` + "`" + `json:"hasParent"` + "`" + `
	JSONPath	string				` + "`" + `json:"jsonPath"` + "`" + `
	Path		path				` + "`" + `json:"path"` + "`" + `
	engine		*Engine
}`
const _Player_type string = `type Player struct{ player playerCore }`
const gearScoreCore_type string = `type gearScoreCore struct {
	ID		GearScoreID	` + "`" + `json:"id"` + "`" + `
	Level		IntValueID	` + "`" + `json:"level"` + "`" + `
	Score		IntValueID	` + "`" + `json:"score"` + "`" + `
	OperationKind	OperationKind	` + "`" + `json:"operationKind"` + "`" + `
	HasParent	bool		` + "`" + `json:"hasParent"` + "`" + `
	JSONPath	string		` + "`" + `json:"jsonPath"` + "`" + `
	Path		path		` + "`" + `json:"path"` + "`" + `
	engine		*Engine
}`
const _GearScore_type string = `type GearScore struct{ gearScore gearScoreCore }`
const positionCore_type string = `type positionCore struct {
	ID		PositionID	` + "`" + `json:"id"` + "`" + `
	X		FloatValueID	` + "`" + `json:"x"` + "`" + `
	Y		FloatValueID	` + "`" + `json:"y"` + "`" + `
	OperationKind	OperationKind	` + "`" + `json:"operationKind"` + "`" + `
	HasParent	bool		` + "`" + `json:"hasParent"` + "`" + `
	JSONPath	string		` + "`" + `json:"jsonPath"` + "`" + `
	Path		path		` + "`" + `json:"path"` + "`" + `
	engine		*Engine
}`
const _Position_type string = `type Position struct{ position positionCore }`
const equipmentSetCore_type string = `type equipmentSetCore struct {
	ID		EquipmentSetID			` + "`" + `json:"id"` + "`" + `
	Equipment	[]EquipmentSetEquipmentRefID	` + "`" + `json:"equipment"` + "`" + `
	Name		StringValueID			` + "`" + `json:"name"` + "`" + `
	OperationKind	OperationKind			` + "`" + `json:"operationKind"` + "`" + `
	HasParent	bool				` + "`" + `json:"hasParent"` + "`" + `
	JSONPath	string				` + "`" + `json:"jsonPath"` + "`" + `
	Path		path				` + "`" + `json:"path"` + "`" + `
	engine		*Engine
}`
const _EquipmentSet_type string = `type EquipmentSet struct{ equipmentSet equipmentSetCore }`
const itemBoundToRefCore_type string = `type itemBoundToRefCore struct {
	ID			ItemBoundToRefID	` + "`" + `json:"id"` + "`" + `
	ParentID		ItemID			` + "`" + `json:"parentID"` + "`" + `
	ChildID			int			` + "`" + `json:"childID"` + "`" + `
	ReferencedElementID	PlayerID		` + "`" + `json:"referencedElementID"` + "`" + `
	OperationKind		OperationKind		` + "`" + `json:"operationKind"` + "`" + `
	Path			path			` + "`" + `json:"path"` + "`" + `
	engine			*Engine
}`
const _ItemBoundToRef_type string = `type ItemBoundToRef struct{ itemBoundToRef itemBoundToRefCore }`
const attackEventTargetRefCore_type string = `type attackEventTargetRefCore struct {
	ID			AttackEventTargetRefID	` + "`" + `json:"id"` + "`" + `
	ParentID		AttackEventID		` + "`" + `json:"parentID"` + "`" + `
	ChildID			int			` + "`" + `json:"childID"` + "`" + `
	ReferencedElementID	PlayerID		` + "`" + `json:"referencedElementID"` + "`" + `
	OperationKind		OperationKind		` + "`" + `json:"operationKind"` + "`" + `
	Path			path			` + "`" + `json:"path"` + "`" + `
	engine			*Engine
}`
const _AttackEventTargetRef_type string = `type AttackEventTargetRef struct{ attackEventTargetRef attackEventTargetRefCore }`
const playerGuildMemberRefCore_type string = `type playerGuildMemberRefCore struct {
	ID			PlayerGuildMemberRefID	` + "`" + `json:"id"` + "`" + `
	ParentID		PlayerID		` + "`" + `json:"parentID"` + "`" + `
	ChildID			int			` + "`" + `json:"childID"` + "`" + `
	ReferencedElementID	PlayerID		` + "`" + `json:"referencedElementID"` + "`" + `
	OperationKind		OperationKind		` + "`" + `json:"operationKind"` + "`" + `
	Path			path			` + "`" + `json:"path"` + "`" + `
	engine			*Engine
}`
const _PlayerGuildMemberRef_type string = `type PlayerGuildMemberRef struct{ playerGuildMemberRef playerGuildMemberRefCore }`
const equipmentSetEquipmentRefCore_type string = `type equipmentSetEquipmentRefCore struct {
	ID			EquipmentSetEquipmentRefID	` + "`" + `json:"id"` + "`" + `
	ParentID		EquipmentSetID			` + "`" + `json:"parentID"` + "`" + `
	ChildID			int				` + "`" + `json:"childID"` + "`" + `
	ReferencedElementID	ItemID				` + "`" + `json:"referencedElementID"` + "`" + `
	OperationKind		OperationKind			` + "`" + `json:"operationKind"` + "`" + `
	Path			path				` + "`" + `json:"path"` + "`" + `
	engine			*Engine
}`
const _EquipmentSetEquipmentRef_type string = `type EquipmentSetEquipmentRef struct{ equipmentSetEquipmentRef equipmentSetEquipmentRefCore }`
const playerEquipmentSetRefCore_type string = `type playerEquipmentSetRefCore struct {
	ID			PlayerEquipmentSetRefID	` + "`" + `json:"id"` + "`" + `
	ParentID		PlayerID		` + "`" + `json:"parentID"` + "`" + `
	ChildID			int			` + "`" + `json:"childID"` + "`" + `
	ReferencedElementID	EquipmentSetID		` + "`" + `json:"referencedElementID"` + "`" + `
	OperationKind		OperationKind		` + "`" + `json:"operationKind"` + "`" + `
	Path			path			` + "`" + `json:"path"` + "`" + `
	engine			*Engine
}`
const _PlayerEquipmentSetRef_type string = `type PlayerEquipmentSetRef struct{ playerEquipmentSetRef playerEquipmentSetRefCore }`
const anyOfPlayer_PositionCore_type string = `type anyOfPlayer_PositionCore struct {
	ID			AnyOfPlayer_PositionID	` + "`" + `json:"id"` + "`" + `
	ElementKind		ElementKind		` + "`" + `json:"elementKind"` + "`" + `
	ParentID		int			` + "`" + `json:"parentID"` + "`" + `
	ChildID			int			` + "`" + `json:"childID"` + "`" + `
	ParentElementPath	path			` + "`" + `json:"parentElementPath"` + "`" + `
	FieldIdentifier		treeFieldIdentifier	` + "`" + `json:"fieldIdentifier"` + "`" + `
	OperationKind		OperationKind		` + "`" + `json:"operationKind"` + "`" + `
	engine			*Engine
}`
const _AnyOfPlayer_Position_type string = `type AnyOfPlayer_Position struct{ anyOfPlayer_Position anyOfPlayer_PositionCore }`
const anyOfPlayer_ZoneItemCore_type string = `type anyOfPlayer_ZoneItemCore struct {
	ID			AnyOfPlayer_ZoneItemID	` + "`" + `json:"id"` + "`" + `
	ElementKind		ElementKind		` + "`" + `json:"elementKind"` + "`" + `
	ParentID		int			` + "`" + `json:"parentID"` + "`" + `
	ChildID			int			` + "`" + `json:"childID"` + "`" + `
	ParentElementPath	path			` + "`" + `json:"parentElementPath"` + "`" + `
	FieldIdentifier		treeFieldIdentifier	` + "`" + `json:"fieldIdentifier"` + "`" + `
	OperationKind		OperationKind		` + "`" + `json:"operationKind"` + "`" + `
	engine			*Engine
}`
const _AnyOfPlayer_ZoneItem_type string = `type AnyOfPlayer_ZoneItem struct{ anyOfPlayer_ZoneItem anyOfPlayer_ZoneItemCore }`
const anyOfItem_Player_ZoneItemCore_type string = `type anyOfItem_Player_ZoneItemCore struct {
	ID			AnyOfItem_Player_ZoneItemID	` + "`" + `json:"id"` + "`" + `
	ElementKind		ElementKind			` + "`" + `json:"elementKind"` + "`" + `
	ParentID		int				` + "`" + `json:"parentID"` + "`" + `
	ChildID			int				` + "`" + `json:"childID"` + "`" + `
	ParentElementPath	path				` + "`" + `json:"parentElementPath"` + "`" + `
	FieldIdentifier		treeFieldIdentifier		` + "`" + `json:"fieldIdentifier"` + "`" + `
	OperationKind		OperationKind			` + "`" + `json:"operationKind"` + "`" + `
	engine			*Engine
}`
const _AnyOfItem_Player_ZoneItem_type string = `type AnyOfItem_Player_ZoneItem struct{ anyOfItem_Player_ZoneItem anyOfItem_Player_ZoneItemCore }`
const playerTargetRefCore_type string = `type playerTargetRefCore struct {
	ID			PlayerTargetRefID	` + "`" + `json:"id"` + "`" + `
	ParentID		PlayerID		` + "`" + `json:"parentID"` + "`" + `
	ChildID			int			` + "`" + `json:"childID"` + "`" + `
	ReferencedElementID	AnyOfPlayer_ZoneItemID	` + "`" + `json:"referencedElementID"` + "`" + `
	OperationKind		OperationKind		` + "`" + `json:"operationKind"` + "`" + `
	Path			path			` + "`" + `json:"path"` + "`" + `
	engine			*Engine
}`
const _PlayerTargetRef_type string = `type PlayerTargetRef struct{ playerTargetRef playerTargetRefCore }`
const playerTargetedByRefCore_type string = `type playerTargetedByRefCore struct {
	ID			PlayerTargetedByRefID	` + "`" + `json:"id"` + "`" + `
	ParentID		PlayerID		` + "`" + `json:"parentID"` + "`" + `
	ChildID			int			` + "`" + `json:"childID"` + "`" + `
	ReferencedElementID	AnyOfPlayer_ZoneItemID	` + "`" + `json:"referencedElementID"` + "`" + `
	OperationKind		OperationKind		` + "`" + `json:"operationKind"` + "`" + `
	Path			path			` + "`" + `json:"path"` + "`" + `
	engine			*Engine
}`
const _PlayerTargetedByRef_type string = `type PlayerTargetedByRef struct{ playerTargetedByRef playerTargetedByRefCore }`
const _IsEmpty_State_func string = `func (s State) IsEmpty() bool {
	if len(s.BoolValue) != 0 {
		return false
	}
	if len(s.FloatValue) != 0 {
		return false
	}
	if len(s.IntValue) != 0 {
		return false
	}
	if len(s.StringValue) != 0 {
		return false
	}
	if len(s.AttackEvent) != 0 {
		return false
	}
	if len(s.EquipmentSet) != 0 {
		return false
	}
	if len(s.GearScore) != 0 {
		return false
	}
	if len(s.Item) != 0 {
		return false
	}
	if len(s.Player) != 0 {
		return false
	}
	if len(s.Position) != 0 {
		return false
	}
	if len(s.Zone) != 0 {
		return false
	}
	if len(s.ZoneItem) != 0 {
		return false
	}
	if len(s.AttackEventTargetRef) != 0 {
		return false
	}
	if len(s.EquipmentSetEquipmentRef) != 0 {
		return false
	}
	if len(s.ItemBoundToRef) != 0 {
		return false
	}
	if len(s.PlayerEquipmentSetRef) != 0 {
		return false
	}
	if len(s.PlayerGuildMemberRef) != 0 {
		return false
	}
	if len(s.PlayerTargetRef) != 0 {
		return false
	}
	if len(s.PlayerTargetedByRef) != 0 {
		return false
	}
	if len(s.AnyOfPlayer_Position) != 0 {
		return false
	}
	if len(s.AnyOfPlayer_ZoneItem) != 0 {
		return false
	}
	if len(s.AnyOfItem_Player_ZoneItem) != 0 {
		return false
	}
	return true
}`
const _OperationKind_type string = `type OperationKind string`
const _OperationKindDelete_type string = `const (
	OperationKindDelete	OperationKind	= "DELETE"
	OperationKindUpdate	OperationKind	= "UPDATE"
	OperationKindUnchanged	OperationKind	= "UNCHANGED"
)`
const _Engine_type string = `type Engine struct {
	State			*State
	Patch			*State
	Tree			*Tree
	BroadcastingClientID	string
	ThisClientID		string
	planner			*assemblePlanner
	IDgen			int
}`
const _NewEngine_func string = `func NewEngine() *Engine {
	return &Engine{IDgen: 1, Patch: newState(), State: newState(), Tree: newTree(), planner: newAssemblePlanner()}
}`
const _GenerateID_Engine_func string = `func (engine *Engine) GenerateID() int {
	newID := engine.IDgen
	engine.IDgen = engine.IDgen + 1
	return newID
}`
const _UpdateState_Engine_func string = `func (engine *Engine) UpdateState() {
	for _, attackEvent := range engine.Patch.AttackEvent {
		engine.deleteAttackEvent(attackEvent.ID)
	}
	for _, boolValue := range engine.Patch.BoolValue {
		if boolValue.OperationKind == OperationKindDelete {
			delete(engine.State.BoolValue, boolValue.ID)
		} else {
			boolValue.OperationKind = OperationKindUnchanged
			engine.State.BoolValue[boolValue.ID] = boolValue
		}
	}
	for _, floatValue := range engine.Patch.FloatValue {
		if floatValue.OperationKind == OperationKindDelete {
			delete(engine.State.FloatValue, floatValue.ID)
		} else {
			floatValue.OperationKind = OperationKindUnchanged
			engine.State.FloatValue[floatValue.ID] = floatValue
		}
	}
	for _, intValue := range engine.Patch.IntValue {
		if intValue.OperationKind == OperationKindDelete {
			delete(engine.State.IntValue, intValue.ID)
		} else {
			intValue.OperationKind = OperationKindUnchanged
			engine.State.IntValue[intValue.ID] = intValue
		}
	}
	for _, stringValue := range engine.Patch.StringValue {
		if stringValue.OperationKind == OperationKindDelete {
			delete(engine.State.StringValue, stringValue.ID)
		} else {
			stringValue.OperationKind = OperationKindUnchanged
			engine.State.StringValue[stringValue.ID] = stringValue
		}
	}
	for _, attackEvent := range engine.Patch.AttackEvent {
		if attackEvent.OperationKind == OperationKindDelete {
			delete(engine.State.AttackEvent, attackEvent.ID)
		} else {
			attackEvent.OperationKind = OperationKindUnchanged
			engine.State.AttackEvent[attackEvent.ID] = attackEvent
		}
	}
	for _, equipmentSet := range engine.Patch.EquipmentSet {
		if equipmentSet.OperationKind == OperationKindDelete {
			delete(engine.State.EquipmentSet, equipmentSet.ID)
		} else {
			equipmentSet.OperationKind = OperationKindUnchanged
			engine.State.EquipmentSet[equipmentSet.ID] = equipmentSet
		}
	}
	for _, gearScore := range engine.Patch.GearScore {
		if gearScore.OperationKind == OperationKindDelete {
			delete(engine.State.GearScore, gearScore.ID)
		} else {
			gearScore.OperationKind = OperationKindUnchanged
			engine.State.GearScore[gearScore.ID] = gearScore
		}
	}
	for _, item := range engine.Patch.Item {
		if item.OperationKind == OperationKindDelete {
			delete(engine.State.Item, item.ID)
		} else {
			item.OperationKind = OperationKindUnchanged
			engine.State.Item[item.ID] = item
		}
	}
	for _, player := range engine.Patch.Player {
		if player.OperationKind == OperationKindDelete {
			delete(engine.State.Player, player.ID)
		} else {
			player.Action = player.Action[:0]
			player.OperationKind = OperationKindUnchanged
			engine.State.Player[player.ID] = player
		}
	}
	for _, position := range engine.Patch.Position {
		if position.OperationKind == OperationKindDelete {
			delete(engine.State.Position, position.ID)
		} else {
			position.OperationKind = OperationKindUnchanged
			engine.State.Position[position.ID] = position
		}
	}
	for _, zone := range engine.Patch.Zone {
		if zone.OperationKind == OperationKindDelete {
			delete(engine.State.Zone, zone.ID)
		} else {
			zone.OperationKind = OperationKindUnchanged
			engine.State.Zone[zone.ID] = zone
		}
	}
	for _, zoneItem := range engine.Patch.ZoneItem {
		if zoneItem.OperationKind == OperationKindDelete {
			delete(engine.State.ZoneItem, zoneItem.ID)
		} else {
			zoneItem.OperationKind = OperationKindUnchanged
			engine.State.ZoneItem[zoneItem.ID] = zoneItem
		}
	}
	for _, attackEventTargetRef := range engine.Patch.AttackEventTargetRef {
		if attackEventTargetRef.OperationKind == OperationKindDelete {
			delete(engine.State.AttackEventTargetRef, attackEventTargetRef.ID)
		} else {
			attackEventTargetRef.OperationKind = OperationKindUnchanged
			engine.State.AttackEventTargetRef[attackEventTargetRef.ID] = attackEventTargetRef
		}
	}
	for _, equipmentSetEquipmentRef := range engine.Patch.EquipmentSetEquipmentRef {
		if equipmentSetEquipmentRef.OperationKind == OperationKindDelete {
			delete(engine.State.EquipmentSetEquipmentRef, equipmentSetEquipmentRef.ID)
		} else {
			equipmentSetEquipmentRef.OperationKind = OperationKindUnchanged
			engine.State.EquipmentSetEquipmentRef[equipmentSetEquipmentRef.ID] = equipmentSetEquipmentRef
		}
	}
	for _, itemBoundToRef := range engine.Patch.ItemBoundToRef {
		if itemBoundToRef.OperationKind == OperationKindDelete {
			delete(engine.State.ItemBoundToRef, itemBoundToRef.ID)
		} else {
			itemBoundToRef.OperationKind = OperationKindUnchanged
			engine.State.ItemBoundToRef[itemBoundToRef.ID] = itemBoundToRef
		}
	}
	for _, playerEquipmentSetRef := range engine.Patch.PlayerEquipmentSetRef {
		if playerEquipmentSetRef.OperationKind == OperationKindDelete {
			delete(engine.State.PlayerEquipmentSetRef, playerEquipmentSetRef.ID)
		} else {
			playerEquipmentSetRef.OperationKind = OperationKindUnchanged
			engine.State.PlayerEquipmentSetRef[playerEquipmentSetRef.ID] = playerEquipmentSetRef
		}
	}
	for _, playerGuildMemberRef := range engine.Patch.PlayerGuildMemberRef {
		if playerGuildMemberRef.OperationKind == OperationKindDelete {
			delete(engine.State.PlayerGuildMemberRef, playerGuildMemberRef.ID)
		} else {
			playerGuildMemberRef.OperationKind = OperationKindUnchanged
			engine.State.PlayerGuildMemberRef[playerGuildMemberRef.ID] = playerGuildMemberRef
		}
	}
	for _, playerTargetRef := range engine.Patch.PlayerTargetRef {
		if playerTargetRef.OperationKind == OperationKindDelete {
			delete(engine.State.PlayerTargetRef, playerTargetRef.ID)
		} else {
			playerTargetRef.OperationKind = OperationKindUnchanged
			engine.State.PlayerTargetRef[playerTargetRef.ID] = playerTargetRef
		}
	}
	for _, playerTargetedByRef := range engine.Patch.PlayerTargetedByRef {
		if playerTargetedByRef.OperationKind == OperationKindDelete {
			delete(engine.State.PlayerTargetedByRef, playerTargetedByRef.ID)
		} else {
			playerTargetedByRef.OperationKind = OperationKindUnchanged
			engine.State.PlayerTargetedByRef[playerTargetedByRef.ID] = playerTargetedByRef
		}
	}
	for _, anyOfPlayer_Position := range engine.Patch.AnyOfPlayer_Position {
		if anyOfPlayer_Position.OperationKind == OperationKindDelete {
			delete(engine.State.AnyOfPlayer_Position, anyOfPlayer_Position.ID)
		} else {
			anyOfPlayer_Position.OperationKind = OperationKindUnchanged
			engine.State.AnyOfPlayer_Position[anyOfPlayer_Position.ID] = anyOfPlayer_Position
		}
	}
	for _, anyOfPlayer_ZoneItem := range engine.Patch.AnyOfPlayer_ZoneItem {
		if anyOfPlayer_ZoneItem.OperationKind == OperationKindDelete {
			delete(engine.State.AnyOfPlayer_ZoneItem, anyOfPlayer_ZoneItem.ID)
		} else {
			anyOfPlayer_ZoneItem.OperationKind = OperationKindUnchanged
			engine.State.AnyOfPlayer_ZoneItem[anyOfPlayer_ZoneItem.ID] = anyOfPlayer_ZoneItem
		}
	}
	for _, anyOfItem_Player_ZoneItem := range engine.Patch.AnyOfItem_Player_ZoneItem {
		if anyOfItem_Player_ZoneItem.OperationKind == OperationKindDelete {
			delete(engine.State.AnyOfItem_Player_ZoneItem, anyOfItem_Player_ZoneItem.ID)
		} else {
			anyOfItem_Player_ZoneItem.OperationKind = OperationKindUnchanged
			engine.State.AnyOfItem_Player_ZoneItem[anyOfItem_Player_ZoneItem.ID] = anyOfItem_Player_ZoneItem
		}
	}
	for key := range engine.Patch.BoolValue {
		delete(engine.Patch.BoolValue, key)
	}
	for key := range engine.Patch.FloatValue {
		delete(engine.Patch.FloatValue, key)
	}
	for key := range engine.Patch.IntValue {
		delete(engine.Patch.IntValue, key)
	}
	for key := range engine.Patch.StringValue {
		delete(engine.Patch.StringValue, key)
	}
	for key := range engine.Patch.AttackEvent {
		delete(engine.Patch.AttackEvent, key)
	}
	for key := range engine.Patch.EquipmentSet {
		delete(engine.Patch.EquipmentSet, key)
	}
	for key := range engine.Patch.GearScore {
		delete(engine.Patch.GearScore, key)
	}
	for key := range engine.Patch.Item {
		delete(engine.Patch.Item, key)
	}
	for key := range engine.Patch.Player {
		delete(engine.Patch.Player, key)
	}
	for key := range engine.Patch.Position {
		delete(engine.Patch.Position, key)
	}
	for key := range engine.Patch.Zone {
		delete(engine.Patch.Zone, key)
	}
	for key := range engine.Patch.ZoneItem {
		delete(engine.Patch.ZoneItem, key)
	}
	for key := range engine.Patch.AttackEventTargetRef {
		delete(engine.Patch.AttackEventTargetRef, key)
	}
	for key := range engine.Patch.EquipmentSetEquipmentRef {
		delete(engine.Patch.EquipmentSetEquipmentRef, key)
	}
	for key := range engine.Patch.ItemBoundToRef {
		delete(engine.Patch.ItemBoundToRef, key)
	}
	for key := range engine.Patch.PlayerEquipmentSetRef {
		delete(engine.Patch.PlayerEquipmentSetRef, key)
	}
	for key := range engine.Patch.PlayerGuildMemberRef {
		delete(engine.Patch.PlayerGuildMemberRef, key)
	}
	for key := range engine.Patch.PlayerTargetRef {
		delete(engine.Patch.PlayerTargetRef, key)
	}
	for key := range engine.Patch.PlayerTargetedByRef {
		delete(engine.Patch.PlayerTargetedByRef, key)
	}
	for key := range engine.Patch.AnyOfPlayer_Position {
		delete(engine.Patch.AnyOfPlayer_Position, key)
	}
	for key := range engine.Patch.AnyOfPlayer_ZoneItem {
		delete(engine.Patch.AnyOfPlayer_ZoneItem, key)
	}
	for key := range engine.Patch.AnyOfItem_Player_ZoneItem {
		delete(engine.Patch.AnyOfItem_Player_ZoneItem, key)
	}
}`
const _ImportPatch_Engine_func string = `func (engine *Engine) ImportPatch(patch *State) {
	for _, boolValue := range patch.BoolValue {
		engine.Patch.BoolValue[boolValue.ID] = boolValue
	}
	for _, floatValue := range patch.FloatValue {
		engine.Patch.FloatValue[floatValue.ID] = floatValue
	}
	for _, intValue := range patch.IntValue {
		engine.Patch.IntValue[intValue.ID] = intValue
	}
	for _, stringValue := range patch.StringValue {
		engine.Patch.StringValue[stringValue.ID] = stringValue
	}
	for _, attackEvent := range patch.AttackEvent {
		engine.Patch.AttackEvent[attackEvent.ID] = attackEvent
	}
	for _, equipmentSet := range patch.EquipmentSet {
		engine.Patch.EquipmentSet[equipmentSet.ID] = equipmentSet
	}
	for _, gearScore := range patch.GearScore {
		engine.Patch.GearScore[gearScore.ID] = gearScore
	}
	for _, item := range patch.Item {
		engine.Patch.Item[item.ID] = item
	}
	for _, player := range patch.Player {
		engine.Patch.Player[player.ID] = player
	}
	for _, position := range patch.Position {
		engine.Patch.Position[position.ID] = position
	}
	for _, zone := range patch.Zone {
		engine.Patch.Zone[zone.ID] = zone
	}
	for _, zoneItem := range patch.ZoneItem {
		engine.Patch.ZoneItem[zoneItem.ID] = zoneItem
	}
	for _, attackEventTargetRef := range patch.AttackEventTargetRef {
		engine.Patch.AttackEventTargetRef[attackEventTargetRef.ID] = attackEventTargetRef
	}
	for _, equipmentSetEquipmentRef := range patch.EquipmentSetEquipmentRef {
		engine.Patch.EquipmentSetEquipmentRef[equipmentSetEquipmentRef.ID] = equipmentSetEquipmentRef
	}
	for _, itemBoundToRef := range patch.ItemBoundToRef {
		engine.Patch.ItemBoundToRef[itemBoundToRef.ID] = itemBoundToRef
	}
	for _, playerEquipmentSetRef := range patch.PlayerEquipmentSetRef {
		engine.Patch.PlayerEquipmentSetRef[playerEquipmentSetRef.ID] = playerEquipmentSetRef
	}
	for _, playerGuildMemberRef := range patch.PlayerGuildMemberRef {
		engine.Patch.PlayerGuildMemberRef[playerGuildMemberRef.ID] = playerGuildMemberRef
	}
	for _, playerTargetRef := range patch.PlayerTargetRef {
		engine.Patch.PlayerTargetRef[playerTargetRef.ID] = playerTargetRef
	}
	for _, playerTargetedByRef := range patch.PlayerTargetedByRef {
		engine.Patch.PlayerTargetedByRef[playerTargetedByRef.ID] = playerTargetedByRef
	}
	for _, anyOfPlayer_Position := range patch.AnyOfPlayer_Position {
		engine.Patch.AnyOfPlayer_Position[anyOfPlayer_Position.ID] = anyOfPlayer_Position
	}
	for _, anyOfPlayer_ZoneItem := range patch.AnyOfPlayer_ZoneItem {
		engine.Patch.AnyOfPlayer_ZoneItem[anyOfPlayer_ZoneItem.ID] = anyOfPlayer_ZoneItem
	}
	for _, anyOfItem_Player_ZoneItem := range patch.AnyOfItem_Player_ZoneItem {
		engine.Patch.AnyOfItem_Player_ZoneItem[anyOfItem_Player_ZoneItem.ID] = anyOfItem_Player_ZoneItem
	}
}`
const _ReferencedDataStatus_type string = `type ReferencedDataStatus string`
const _ReferencedDataModified_type string = `const (
	ReferencedDataModified	ReferencedDataStatus	= "MODIFIED"
	ReferencedDataUnchanged	ReferencedDataStatus	= "UNCHANGED"
)`
const _ElementKind_type string = `type ElementKind string`
const _ElementKindBoolValue_type string = `const (
	ElementKindBoolValue	ElementKind	= "bool"
	ElementKindFloatValue	ElementKind	= "float64"
	ElementKindIntValue	ElementKind	= "int64"
	ElementKindStringValue	ElementKind	= "string"
	ElementKindAttackEvent	ElementKind	= "AttackEvent"
	ElementKindEquipmentSet	ElementKind	= "EquipmentSet"
	ElementKindGearScore	ElementKind	= "GearScore"
	ElementKindItem		ElementKind	= "Item"
	ElementKindPlayer	ElementKind	= "Player"
	ElementKindPosition	ElementKind	= "Position"
	ElementKindZone		ElementKind	= "Zone"
	ElementKindZoneItem	ElementKind	= "ZoneItem"
)`
const _Tree_type string = `type Tree struct {
	AttackEvent	map[AttackEventID]attackEvent	` + "`" + `json:"attackEvent"` + "`" + `
	EquipmentSet	map[EquipmentSetID]equipmentSet	` + "`" + `json:"equipmentSet"` + "`" + `
	GearScore	map[GearScoreID]gearScore	` + "`" + `json:"gearScore"` + "`" + `
	Item		map[ItemID]item			` + "`" + `json:"item"` + "`" + `
	Player		map[PlayerID]player		` + "`" + `json:"player"` + "`" + `
	Position	map[PositionID]position		` + "`" + `json:"position"` + "`" + `
	Zone		map[ZoneID]zone			` + "`" + `json:"zone"` + "`" + `
	ZoneItem	map[ZoneItemID]zoneItem		` + "`" + `json:"zoneItem"` + "`" + `
}`
const newTree_func string = `func newTree() *Tree {
	return &Tree{AttackEvent: make(map[AttackEventID]attackEvent), EquipmentSet: make(map[EquipmentSetID]equipmentSet), GearScore: make(map[GearScoreID]gearScore), Item: make(map[ItemID]item), Player: make(map[PlayerID]player), Position: make(map[PositionID]position), Zone: make(map[ZoneID]zone), ZoneItem: make(map[ZoneItemID]zoneItem)}
}`
const clear_Tree_func string = `func (t *Tree) clear() {
	for key := range t.AttackEvent {
		delete(t.AttackEvent, key)
	}
	for key := range t.EquipmentSet {
		delete(t.EquipmentSet, key)
	}
	for key := range t.GearScore {
		delete(t.GearScore, key)
	}
	for key := range t.Item {
		delete(t.Item, key)
	}
	for key := range t.Player {
		delete(t.Player, key)
	}
	for key := range t.Position {
		delete(t.Position, key)
	}
	for key := range t.Zone {
		delete(t.Zone, key)
	}
	for key := range t.ZoneItem {
		delete(t.ZoneItem, key)
	}
}`
const zoneItem_type string = `type zoneItem struct {
	ID		ZoneItemID	` + "`" + `json:"id"` + "`" + `
	Item		*item		` + "`" + `json:"item"` + "`" + `
	Position	*position	` + "`" + `json:"position"` + "`" + `
	OperationKind	OperationKind	` + "`" + `json:"operationKind"` + "`" + `
}`
const item_type string = `type item struct {
	ID		ItemID			` + "`" + `json:"id"` + "`" + `
	BoundTo		*elementReference	` + "`" + `json:"boundTo"` + "`" + `
	GearScore	*gearScore		` + "`" + `json:"gearScore"` + "`" + `
	Name		*string			` + "`" + `json:"name"` + "`" + `
	Origin		interface{}		` + "`" + `json:"origin"` + "`" + `
	OperationKind	OperationKind		` + "`" + `json:"operationKind"` + "`" + `
}`
const attackEvent_type string = `type attackEvent struct {
	ID		AttackEventID		` + "`" + `json:"id"` + "`" + `
	Target		*elementReference	` + "`" + `json:"target"` + "`" + `
	OperationKind	OperationKind		` + "`" + `json:"operationKind"` + "`" + `
}`
const equipmentSet_type string = `type equipmentSet struct {
	ID		EquipmentSetID			` + "`" + `json:"id"` + "`" + `
	Equipment	map[ItemID]elementReference	` + "`" + `json:"equipment"` + "`" + `
	Name		*string				` + "`" + `json:"name"` + "`" + `
	OperationKind	OperationKind			` + "`" + `json:"operationKind"` + "`" + `
}`
const position_type string = `type position struct {
	ID		PositionID	` + "`" + `json:"id"` + "`" + `
	X		*float64	` + "`" + `json:"x"` + "`" + `
	Y		*float64	` + "`" + `json:"y"` + "`" + `
	OperationKind	OperationKind	` + "`" + `json:"operationKind"` + "`" + `
}`
const gearScore_type string = `type gearScore struct {
	ID		GearScoreID	` + "`" + `json:"id"` + "`" + `
	Level		*int64		` + "`" + `json:"level"` + "`" + `
	Score		*int64		` + "`" + `json:"score"` + "`" + `
	OperationKind	OperationKind	` + "`" + `json:"operationKind"` + "`" + `
}`
const player_type string = `type player struct {
	ID		PlayerID				` + "`" + `json:"id"` + "`" + `
	Action		map[AttackEventID]attackEvent		` + "`" + `json:"action"` + "`" + `
	EquipmentSets	map[EquipmentSetID]elementReference	` + "`" + `json:"equipmentSets"` + "`" + `
	GearScore	*gearScore				` + "`" + `json:"gearScore"` + "`" + `
	GuildMembers	map[PlayerID]elementReference		` + "`" + `json:"guildMembers"` + "`" + `
	Items		map[ItemID]item				` + "`" + `json:"items"` + "`" + `
	Position	*position				` + "`" + `json:"position"` + "`" + `
	Target		*elementReference			` + "`" + `json:"target"` + "`" + `
	TargetedBy	map[int]elementReference		` + "`" + `json:"targetedBy"` + "`" + `
	OperationKind	OperationKind				` + "`" + `json:"operationKind"` + "`" + `
}`
const zone_type string = `type zone struct {
	ID		ZoneID			` + "`" + `json:"id"` + "`" + `
	Interactables	map[int]interface{}	` + "`" + `json:"interactables"` + "`" + `
	Items		map[ZoneItemID]zoneItem	` + "`" + `json:"items"` + "`" + `
	Players		map[PlayerID]player	` + "`" + `json:"players"` + "`" + `
	Tags		[]string		` + "`" + `json:"tags"` + "`" + `
	OperationKind	OperationKind		` + "`" + `json:"operationKind"` + "`" + `
}`
const elementReference_type string = `type elementReference struct {
	ID			int			` + "`" + `json:"id"` + "`" + `
	OperationKind		OperationKind		` + "`" + `json:"operationKind"` + "`" + `
	ElementID		int			` + "`" + `json:"elementID"` + "`" + `
	ElementKind		ElementKind		` + "`" + `json:"elementKind"` + "`" + `
	ReferencedDataStatus	ReferencedDataStatus	` + "`" + `json:"referencedDataStatus"` + "`" + `
	ElementPath		string			` + "`" + `json:"elementPath"` + "`" + `
}`

var decl_to_string_decl_collection = map[string]string{"_Action_Player_func": _Action_Player_func, "_AddAction_Player_func": _AddAction_Player_func, "_AddEquipmentSet_Player_func": _AddEquipmentSet_Player_func, "_AddEquipment_EquipmentSet_func": _AddEquipment_EquipmentSet_func, "_AddGuildMember_Player_func": _AddGuildMember_Player_func, "_AddInteractableItem_Zone_func": _AddInteractableItem_Zone_func, "_AddInteractablePlayer_Zone_func": _AddInteractablePlayer_Zone_func, "_AddInteractableZoneItem_Zone_func": _AddInteractableZoneItem_Zone_func, "_AddItem_Player_func": _AddItem_Player_func, "_AddItem_Zone_func": _AddItem_Zone_func, "_AddPlayer_Zone_func": _AddPlayer_Zone_func, "_AddTag_Zone_func": _AddTag_Zone_func, "_AddTargetedByPlayer_Player_func": _AddTargetedByPlayer_Player_func, "_AddTargetedByZoneItem_Player_func": _AddTargetedByZoneItem_Player_func, "_AnyOfItem_Player_ZoneItemID_type": _AnyOfItem_Player_ZoneItemID_type, "_AnyOfItem_Player_ZoneItemRef_type": _AnyOfItem_Player_ZoneItemRef_type, "_AnyOfItem_Player_ZoneItemSliceElement_type": _AnyOfItem_Player_ZoneItemSliceElement_type, "_AnyOfItem_Player_ZoneItem_type": _AnyOfItem_Player_ZoneItem_type, "_AnyOfPlayer_PositionID_type": _AnyOfPlayer_PositionID_type, "_AnyOfPlayer_PositionRef_type": _AnyOfPlayer_PositionRef_type, "_AnyOfPlayer_PositionSliceElement_type": _AnyOfPlayer_PositionSliceElement_type, "_AnyOfPlayer_Position_type": _AnyOfPlayer_Position_type, "_AnyOfPlayer_ZoneItemID_type": _AnyOfPlayer_ZoneItemID_type, "_AnyOfPlayer_ZoneItemRef_type": _AnyOfPlayer_ZoneItemRef_type, "_AnyOfPlayer_ZoneItemSliceElement_type": _AnyOfPlayer_ZoneItemSliceElement_type, "_AnyOfPlayer_ZoneItem_type": _AnyOfPlayer_ZoneItem_type, "_AssembleFullTree_Engine_func": _AssembleFullTree_Engine_func, "_AssembleUpdateTree_Engine_func": _AssembleUpdateTree_Engine_func, "_AttackEventID_type": _AttackEventID_type, "_AttackEventTargetRefID_type": _AttackEventTargetRefID_type, "_AttackEventTargetRef_type": _AttackEventTargetRef_type, "_AttackEvent_Engine_func": _AttackEvent_Engine_func, "_AttackEvent_type": _AttackEvent_type, "_BeItem_AnyOfItem_Player_ZoneItem_func": _BeItem_AnyOfItem_Player_ZoneItem_func, "_BePlayer_AnyOfItem_Player_ZoneItem_func": _BePlayer_AnyOfItem_Player_ZoneItem_func, "_BePlayer_AnyOfPlayer_Position_func": _BePlayer_AnyOfPlayer_Position_func, "_BePlayer_AnyOfPlayer_ZoneItem_func": _BePlayer_AnyOfPlayer_ZoneItem_func, "_BePosition_AnyOfPlayer_Position_func": _BePosition_AnyOfPlayer_Position_func, "_BeZoneItem_AnyOfItem_Player_ZoneItem_func": _BeZoneItem_AnyOfItem_Player_ZoneItem_func, "_BeZoneItem_AnyOfPlayer_ZoneItem_func": _BeZoneItem_AnyOfPlayer_ZoneItem_func, "_BoolValueID_type": _BoolValueID_type, "_BoundTo_Item_func": _BoundTo_Item_func, "_CreateAttackEvent_Engine_func": _CreateAttackEvent_Engine_func, "_CreateEquipmentSet_Engine_func": _CreateEquipmentSet_Engine_func, "_CreateGearScore_Engine_func": _CreateGearScore_Engine_func, "_CreateItem_Engine_func": _CreateItem_Engine_func, "_CreatePlayer_Engine_func": _CreatePlayer_Engine_func, "_CreatePosition_Engine_func": _CreatePosition_Engine_func, "_CreateZoneItem_Engine_func": _CreateZoneItem_Engine_func, "_CreateZone_Engine_func": _CreateZone_Engine_func, "_DeleteAttackEvent_Engine_func": _DeleteAttackEvent_Engine_func, "_DeleteEquipmentSet_Engine_func": _DeleteEquipmentSet_Engine_func, "_DeleteGearScore_Engine_func": _DeleteGearScore_Engine_func, "_DeleteItem_Engine_func": _DeleteItem_Engine_func, "_DeletePlayer_Engine_func": _DeletePlayer_Engine_func, "_DeletePosition_Engine_func": _DeletePosition_Engine_func, "_DeleteZoneItem_Engine_func": _DeleteZoneItem_Engine_func, "_DeleteZone_Engine_func": _DeleteZone_Engine_func, "_ElementKindBoolValue_type": _ElementKindBoolValue_type, "_ElementKind_type": _ElementKind_type, "_Engine_type": _Engine_type, "_EquipmentSetEquipmentRefID_type": _EquipmentSetEquipmentRefID_type, "_EquipmentSetEquipmentRef_type": _EquipmentSetEquipmentRef_type, "_EquipmentSetID_type": _EquipmentSetID_type, "_EquipmentSet_Engine_func": _EquipmentSet_Engine_func, "_EquipmentSet_type": _EquipmentSet_type, "_EquipmentSets_Player_func": _EquipmentSets_Player_func, "_Equipment_EquipmentSet_func": _Equipment_EquipmentSet_func, "_EveryAttackEvent_Engine_func": _EveryAttackEvent_Engine_func, "_EveryEquipmentSet_Engine_func": _EveryEquipmentSet_Engine_func, "_EveryGearScore_Engine_func": _EveryGearScore_Engine_func, "_EveryItem_Engine_func": _EveryItem_Engine_func, "_EveryPlayer_Engine_func": _EveryPlayer_Engine_func, "_EveryPosition_Engine_func": _EveryPosition_Engine_func, "_EveryZoneItem_Engine_func": _EveryZoneItem_Engine_func, "_EveryZone_Engine_func": _EveryZone_Engine_func, "_Exists_AttackEvent_func": _Exists_AttackEvent_func, "_Exists_EquipmentSet_func": _Exists_EquipmentSet_func, "_Exists_GearScore_func": _Exists_GearScore_func, "_Exists_Item_func": _Exists_Item_func, "_Exists_Player_func": _Exists_Player_func, "_Exists_Position_func": _Exists_Position_func, "_Exists_ZoneItem_func": _Exists_ZoneItem_func, "_Exists_Zone_func": _Exists_Zone_func, "_FloatValueID_type": _FloatValueID_type, "_GearScoreID_type": _GearScoreID_type, "_GearScore_Engine_func": _GearScore_Engine_func, "_GearScore_Item_func": _GearScore_Item_func, "_GearScore_Player_func": _GearScore_Player_func, "_GearScore_type": _GearScore_type, "_GenerateID_Engine_func": _GenerateID_Engine_func, "_Get_AttackEventTargetRef_func": _Get_AttackEventTargetRef_func, "_Get_EquipmentSetEquipmentRef_func": _Get_EquipmentSetEquipmentRef_func, "_Get_ItemBoundToRef_func": _Get_ItemBoundToRef_func, "_Get_PlayerEquipmentSetRef_func": _Get_PlayerEquipmentSetRef_func, "_Get_PlayerGuildMemberRef_func": _Get_PlayerGuildMemberRef_func, "_Get_PlayerTargetRef_func": _Get_PlayerTargetRef_func, "_Get_PlayerTargetedByRef_func": _Get_PlayerTargetedByRef_func, "_GuildMembers_Player_func": _GuildMembers_Player_func, "_ID_AnyOfItem_Player_ZoneItem_func": _ID_AnyOfItem_Player_ZoneItem_func, "_ID_AnyOfPlayer_Position_func": _ID_AnyOfPlayer_Position_func, "_ID_AnyOfPlayer_ZoneItem_func": _ID_AnyOfPlayer_ZoneItem_func, "_ID_AttackEventTargetRef_func": _ID_AttackEventTargetRef_func, "_ID_AttackEvent_func": _ID_AttackEvent_func, "_ID_EquipmentSetEquipmentRef_func": _ID_EquipmentSetEquipmentRef_func, "_ID_EquipmentSet_func": _ID_EquipmentSet_func, "_ID_GearScore_func": _ID_GearScore_func, "_ID_ItemBoundToRef_func": _ID_ItemBoundToRef_func, "_ID_Item_func": _ID_Item_func, "_ID_PlayerEquipmentSetRef_func": _ID_PlayerEquipmentSetRef_func, "_ID_PlayerGuildMemberRef_func": _ID_PlayerGuildMemberRef_func, "_ID_PlayerTargetRef_func": _ID_PlayerTargetRef_func, "_ID_PlayerTargetedByRef_func": _ID_PlayerTargetedByRef_func, "_ID_Player_func": _ID_Player_func, "_ID_Position_func": _ID_Position_func, "_ID_ZoneItem_func": _ID_ZoneItem_func, "_ID_Zone_func": _ID_Zone_func, "_ImportPatch_Engine_func": _ImportPatch_Engine_func, "_IntValueID_type": _IntValueID_type, "_Interactables_Zone_func": _Interactables_Zone_func, "_IsEmpty_State_func": _IsEmpty_State_func, "_IsSet_AttackEventTargetRef_func": _IsSet_AttackEventTargetRef_func, "_IsSet_ItemBoundToRef_func": _IsSet_ItemBoundToRef_func, "_IsSet_PlayerTargetRef_func": _IsSet_PlayerTargetRef_func, "_ItemBoundToRefID_type": _ItemBoundToRefID_type, "_ItemBoundToRef_type": _ItemBoundToRef_type, "_ItemID_type": _ItemID_type, "_Item_AnyOfItem_Player_ZoneItem_func": _Item_AnyOfItem_Player_ZoneItem_func, "_Item_Engine_func": _Item_Engine_func, "_Item_ZoneItem_func": _Item_ZoneItem_func, "_Item_type": _Item_type, "_Items_Player_func": _Items_Player_func, "_Items_Zone_func": _Items_Zone_func, "_Kind_AnyOfItem_Player_ZoneItem_func": _Kind_AnyOfItem_Player_ZoneItem_func, "_Kind_AnyOfPlayer_Position_func": _Kind_AnyOfPlayer_Position_func, "_Kind_AnyOfPlayer_ZoneItem_func": _Kind_AnyOfPlayer_ZoneItem_func, "_Level_GearScore_func": _Level_GearScore_func, "_Name_EquipmentSet_func": _Name_EquipmentSet_func, "_Name_Item_func": _Name_Item_func, "_NewEngine_func": _NewEngine_func, "_OperationKindDelete_type": _OperationKindDelete_type, "_OperationKind_type": _OperationKind_type, "_Origin_Item_func": _Origin_Item_func, "_ParentItem_GearScore_func": _ParentItem_GearScore_func, "_ParentItem_Player_func": _ParentItem_Player_func, "_ParentItem_Position_func": _ParentItem_Position_func, "_ParentKind_AttackEvent_func": _ParentKind_AttackEvent_func, "_ParentKind_GearScore_func": _ParentKind_GearScore_func, "_ParentKind_Item_func": _ParentKind_Item_func, "_ParentKind_Player_func": _ParentKind_Player_func, "_ParentKind_Position_func": _ParentKind_Position_func, "_ParentKind_ZoneItem_func": _ParentKind_ZoneItem_func, "_ParentPlayer_AttackEvent_func": _ParentPlayer_AttackEvent_func, "_ParentPlayer_GearScore_func": _ParentPlayer_GearScore_func, "_ParentPlayer_Item_func": _ParentPlayer_Item_func, "_ParentPlayer_Position_func": _ParentPlayer_Position_func, "_ParentZoneItem_Item_func": _ParentZoneItem_Item_func, "_ParentZoneItem_Position_func": _ParentZoneItem_Position_func, "_ParentZone_Item_func": _ParentZone_Item_func, "_ParentZone_Player_func": _ParentZone_Player_func, "_ParentZone_ZoneItem_func": _ParentZone_ZoneItem_func, "_Path_AttackEvent_func": _Path_AttackEvent_func, "_Path_EquipmentSet_func": _Path_EquipmentSet_func, "_Path_GearScore_func": _Path_GearScore_func, "_Path_Item_func": _Path_Item_func, "_Path_Player_func": _Path_Player_func, "_Path_Position_func": _Path_Position_func, "_Path_ZoneItem_func": _Path_ZoneItem_func, "_Path_Zone_func": _Path_Zone_func, "_PlayerEquipmentSetRefID_type": _PlayerEquipmentSetRefID_type, "_PlayerEquipmentSetRef_type": _PlayerEquipmentSetRef_type, "_PlayerGuildMemberRefID_type": _PlayerGuildMemberRefID_type, "_PlayerGuildMemberRef_type": _PlayerGuildMemberRef_type, "_PlayerID_type": _PlayerID_type, "_PlayerTargetRefID_type": _PlayerTargetRefID_type, "_PlayerTargetRef_type": _PlayerTargetRef_type, "_PlayerTargetedByRefID_type": _PlayerTargetedByRefID_type, "_PlayerTargetedByRef_type": _PlayerTargetedByRef_type, "_Player_AnyOfItem_Player_ZoneItem_func": _Player_AnyOfItem_Player_ZoneItem_func, "_Player_AnyOfPlayer_Position_func": _Player_AnyOfPlayer_Position_func, "_Player_AnyOfPlayer_ZoneItem_func": _Player_AnyOfPlayer_ZoneItem_func, "_Player_Engine_func": _Player_Engine_func, "_Player_type": _Player_type, "_Players_Zone_func": _Players_Zone_func, "_PositionID_type": _PositionID_type, "_Position_AnyOfPlayer_Position_func": _Position_AnyOfPlayer_Position_func, "_Position_Engine_func": _Position_Engine_func, "_Position_Player_func": _Position_Player_func, "_Position_ZoneItem_func": _Position_ZoneItem_func, "_Position_type": _Position_type, "_QueryAttackEvents_Engine_func": _QueryAttackEvents_Engine_func, "_QueryEquipmentSets_Engine_func": _QueryEquipmentSets_Engine_func, "_QueryGearScores_Engine_func": _QueryGearScores_Engine_func, "_QueryItems_Engine_func": _QueryItems_Engine_func, "_QueryPlayers_Engine_func": _QueryPlayers_Engine_func, "_QueryPositions_Engine_func": _QueryPositions_Engine_func, "_QueryZoneItems_Engine_func": _QueryZoneItems_Engine_func, "_QueryZones_Engine_func": _QueryZones_Engine_func, "_ReferencedDataModified_type": _ReferencedDataModified_type, "_ReferencedDataStatus_type": _ReferencedDataStatus_type, "_RemoveAction_Player_func": _RemoveAction_Player_func, "_RemoveEquipmentSet_Player_func": _RemoveEquipmentSet_Player_func, "_RemoveEquipment_EquipmentSet_func": _RemoveEquipment_EquipmentSet_func, "_RemoveGuildMember_Player_func": _RemoveGuildMember_Player_func, "_RemoveInteractableItem_Zone_func": _RemoveInteractableItem_Zone_func, "_RemoveInteractablePlayer_Zone_func": _RemoveInteractablePlayer_Zone_func, "_RemoveInteractableZoneItem_Zone_func": _RemoveInteractableZoneItem_Zone_func, "_RemoveItem_Player_func": _RemoveItem_Player_func, "_RemoveItem_Zone_func": _RemoveItem_Zone_func, "_RemovePlayer_Zone_func": _RemovePlayer_Zone_func, "_RemoveTag_Zone_func": _RemoveTag_Zone_func, "_RemoveTargetedByPlayer_Player_func": _RemoveTargetedByPlayer_Player_func, "_RemoveTargetedByZoneItem_Player_func": _RemoveTargetedByZoneItem_Player_func, "_Score_GearScore_func": _Score_GearScore_func, "_SetBoundTo_Item_func": _SetBoundTo_Item_func, "_SetLevel_GearScore_func": _SetLevel_GearScore_func, "_SetName_EquipmentSet_func": _SetName_EquipmentSet_func, "_SetName_Item_func": _SetName_Item_func, "_SetScore_GearScore_func": _SetScore_GearScore_func, "_SetTargetPlayer_Player_func": _SetTargetPlayer_Player_func, "_SetTargetZoneItem_Player_func": _SetTargetZoneItem_Player_func, "_SetTarget_AttackEvent_func": _SetTarget_AttackEvent_func, "_SetX_Position_func": _SetX_Position_func, "_SetY_Position_func": _SetY_Position_func, "_State_type": _State_type, "_StringValueID_type": _StringValueID_type, "_Tags_Zone_func": _Tags_Zone_func, "_Target_AttackEvent_func": _Target_AttackEvent_func, "_Target_Player_func": _Target_Player_func, "_TargetedBy_Player_func": _TargetedBy_Player_func, "_Tree_type": _Tree_type, "_Unset_AttackEventTargetRef_func": _Unset_AttackEventTargetRef_func, "_Unset_ItemBoundToRef_func": _Unset_ItemBoundToRef_func, "_Unset_PlayerTargetRef_func": _Unset_PlayerTargetRef_func, "_UpdateState_Engine_func": _UpdateState_Engine_func, "_X_Position_func": _X_Position_func, "_Y_Position_func": _Y_Position_func, "_ZoneID_type": _ZoneID_type, "_ZoneItemID_type": _ZoneItemID_type, "_ZoneItem_AnyOfItem_Player_ZoneItem_func": _ZoneItem_AnyOfItem_Player_ZoneItem_func, "_ZoneItem_AnyOfPlayer_ZoneItem_func": _ZoneItem_AnyOfPlayer_ZoneItem_func, "_ZoneItem_Engine_func": _ZoneItem_Engine_func, "_ZoneItem_type": _ZoneItem_type, "_Zone_Engine_func": _Zone_Engine_func, "_Zone_type": _Zone_type, "allAttackEventIDs_Engine_func": allAttackEventIDs_Engine_func, "allAttackEventTargetRefIDs_Engine_func": allAttackEventTargetRefIDs_Engine_func, "allEquipmentSetEquipmentRefIDs_Engine_func": allEquipmentSetEquipmentRefIDs_Engine_func, "allEquipmentSetIDs_Engine_func": allEquipmentSetIDs_Engine_func, "allGearScoreIDs_Engine_func": allGearScoreIDs_Engine_func, "allItemBoundToRefIDs_Engine_func": allItemBoundToRefIDs_Engine_func, "allItemIDs_Engine_func": allItemIDs_Engine_func, "allPlayerEquipmentSetRefIDs_Engine_func": allPlayerEquipmentSetRefIDs_Engine_func, "allPlayerGuildMemberRefIDs_Engine_func": allPlayerGuildMemberRefIDs_Engine_func, "allPlayerIDs_Engine_func": allPlayerIDs_Engine_func, "allPlayerTargetRefIDs_Engine_func": allPlayerTargetRefIDs_Engine_func, "allPlayerTargetedByRefIDs_Engine_func": allPlayerTargetedByRefIDs_Engine_func, "allPositionIDs_Engine_func": allPositionIDs_Engine_func, "allZoneIDs_Engine_func": allZoneIDs_Engine_func, "allZoneItemIDs_Engine_func": allZoneItemIDs_Engine_func, "anyOfItem_Player_ZoneItemCore_type": anyOfItem_Player_ZoneItemCore_type, "anyOfItem_Player_ZoneItem_Engine_func": anyOfItem_Player_ZoneItem_Engine_func, "anyOfPlayer_PositionCore_type": anyOfPlayer_PositionCore_type, "anyOfPlayer_Position_Engine_func": anyOfPlayer_Position_Engine_func, "anyOfPlayer_ZoneItemCore_type": anyOfPlayer_ZoneItemCore_type, "anyOfPlayer_ZoneItem_Engine_func": anyOfPlayer_ZoneItem_Engine_func, "assembleAttackEventPath_Engine_func": assembleAttackEventPath_Engine_func, "assembleEquipmentSetPath_Engine_func": assembleEquipmentSetPath_Engine_func, "assembleGearScorePath_Engine_func": assembleGearScorePath_Engine_func, "assembleItemPath_Engine_func": assembleItemPath_Engine_func, "assemblePlanner_type": assemblePlanner_type, "assemblePlayerPath_Engine_func": assemblePlayerPath_Engine_func, "assemblePositionPath_Engine_func": assemblePositionPath_Engine_func, "assembleTree_Engine_func": assembleTree_Engine_func, "assembleZoneItemPath_Engine_func": assembleZoneItemPath_Engine_func, "assembleZonePath_Engine_func": assembleZonePath_Engine_func, "attackEventCheckPool_type": attackEventCheckPool_type, "attackEventCore_type": attackEventCore_type, "attackEventIDSlicePool_type": attackEventIDSlicePool_type, "attackEventIdentifier_type": attackEventIdentifier_type, "attackEventTargetRefCheckPool_type": attackEventTargetRefCheckPool_type, "attackEventTargetRefCore_type": attackEventTargetRefCore_type, "attackEventTargetRefIDSlicePool_type": attackEventTargetRefIDSlicePool_type, "attackEventTargetRef_Engine_func": attackEventTargetRef_Engine_func, "attackEvent_type": attackEvent_type, "beItem_anyOfItem_Player_ZoneItemCore_func": beItem_anyOfItem_Player_ZoneItemCore_func, "bePlayer_anyOfItem_Player_ZoneItemCore_func": bePlayer_anyOfItem_Player_ZoneItemCore_func, "bePlayer_anyOfPlayer_PositionCore_func": bePlayer_anyOfPlayer_PositionCore_func, "bePlayer_anyOfPlayer_ZoneItemCore_func": bePlayer_anyOfPlayer_ZoneItemCore_func, "bePosition_anyOfPlayer_PositionCore_func": bePosition_anyOfPlayer_PositionCore_func, "beZoneItem_anyOfItem_Player_ZoneItemCore_func": beZoneItem_anyOfItem_Player_ZoneItemCore_func, "beZoneItem_anyOfPlayer_ZoneItemCore_func": beZoneItem_anyOfPlayer_ZoneItemCore_func, "boolValue_Engine_func": boolValue_Engine_func, "boolValue_type": boolValue_type, "clear_Tree_func": clear_Tree_func, "clear_assemblePlanner_func": clear_assemblePlanner_func, "createAnyOfItem_Player_ZoneItem_Engine_func": createAnyOfItem_Player_ZoneItem_Engine_func, "createAnyOfPlayer_Position_Engine_func": createAnyOfPlayer_Position_Engine_func, "createAnyOfPlayer_ZoneItem_Engine_func": createAnyOfPlayer_ZoneItem_Engine_func, "createAttackEventTargetRef_Engine_func": createAttackEventTargetRef_Engine_func, "createAttackEvent_Engine_func": createAttackEvent_Engine_func, "createBoolValue_Engine_func": createBoolValue_Engine_func, "createEquipmentSetEquipmentRef_Engine_func": createEquipmentSetEquipmentRef_Engine_func, "createEquipmentSet_Engine_func": createEquipmentSet_Engine_func, "createFloatValue_Engine_func": createFloatValue_Engine_func, "createGearScore_Engine_func": createGearScore_Engine_func, "createIntValue_Engine_func": createIntValue_Engine_func, "createItemBoundToRef_Engine_func": createItemBoundToRef_Engine_func, "createItem_Engine_func": createItem_Engine_func, "createPlayerEquipmentSetRef_Engine_func": createPlayerEquipmentSetRef_Engine_func, "createPlayerGuildMemberRef_Engine_func": createPlayerGuildMemberRef_Engine_func, "createPlayerTargetRef_Engine_func": createPlayerTargetRef_Engine_func, "createPlayerTargetedByRef_Engine_func": createPlayerTargetedByRef_Engine_func, "createPlayer_Engine_func": createPlayer_Engine_func, "createPosition_Engine_func": createPosition_Engine_func, "createStringValue_Engine_func": createStringValue_Engine_func, "createZoneItem_Engine_func": createZoneItem_Engine_func, "createZone_Engine_func": createZone_Engine_func, "deduplicateAttackEventIDs_func": deduplicateAttackEventIDs_func, "deduplicateAttackEventTargetRefIDs_func": deduplicateAttackEventTargetRefIDs_func, "deduplicateEquipmentSetEquipmentRefIDs_func": deduplicateEquipmentSetEquipmentRefIDs_func, "deduplicateEquipmentSetIDs_func": deduplicateEquipmentSetIDs_func, "deduplicateGearScoreIDs_func": deduplicateGearScoreIDs_func, "deduplicateItemBoundToRefIDs_func": deduplicateItemBoundToRefIDs_func, "deduplicateItemIDs_func": deduplicateItemIDs_func, "deduplicatePlayerEquipmentSetRefIDs_func": deduplicatePlayerEquipmentSetRefIDs_func, "deduplicatePlayerGuildMemberRefIDs_func": deduplicatePlayerGuildMemberRefIDs_func, "deduplicatePlayerIDs_func": deduplicatePlayerIDs_func, "deduplicatePlayerTargetRefIDs_func": deduplicatePlayerTargetRefIDs_func, "deduplicatePlayerTargetedByRefIDs_func": deduplicatePlayerTargetedByRefIDs_func, "deduplicatePositionIDs_func": deduplicatePositionIDs_func, "deduplicateZoneIDs_func": deduplicateZoneIDs_func, "deduplicateZoneItemIDs_func": deduplicateZoneItemIDs_func, "deleteAnyOfItem_Player_ZoneItem_Engine_func": deleteAnyOfItem_Player_ZoneItem_Engine_func, "deleteAnyOfPlayer_Position_Engine_func": deleteAnyOfPlayer_Position_Engine_func, "deleteAnyOfPlayer_ZoneItem_Engine_func": deleteAnyOfPlayer_ZoneItem_Engine_func, "deleteAttackEventTargetRef_Engine_func": deleteAttackEventTargetRef_Engine_func, "deleteAttackEvent_Engine_func": deleteAttackEvent_Engine_func, "deleteBoolValue_Engine_func": deleteBoolValue_Engine_func, "deleteChild_anyOfItem_Player_ZoneItemCore_func": deleteChild_anyOfItem_Player_ZoneItemCore_func, "deleteChild_anyOfPlayer_PositionCore_func": deleteChild_anyOfPlayer_PositionCore_func, "deleteChild_anyOfPlayer_ZoneItemCore_func": deleteChild_anyOfPlayer_ZoneItemCore_func, "deleteEquipmentSetEquipmentRef_Engine_func": deleteEquipmentSetEquipmentRef_Engine_func, "deleteEquipmentSet_Engine_func": deleteEquipmentSet_Engine_func, "deleteFloatValue_Engine_func": deleteFloatValue_Engine_func, "deleteGearScore_Engine_func": deleteGearScore_Engine_func, "deleteIntValue_Engine_func": deleteIntValue_Engine_func, "deleteItemBoundToRef_Engine_func": deleteItemBoundToRef_Engine_func, "deleteItem_Engine_func": deleteItem_Engine_func, "deletePlayerEquipmentSetRef_Engine_func": deletePlayerEquipmentSetRef_Engine_func, "deletePlayerGuildMemberRef_Engine_func": deletePlayerGuildMemberRef_Engine_func, "deletePlayerTargetRef_Engine_func": deletePlayerTargetRef_Engine_func, "deletePlayerTargetedByRef_Engine_func": deletePlayerTargetedByRef_Engine_func, "deletePlayer_Engine_func": deletePlayer_Engine_func, "deletePosition_Engine_func": deletePosition_Engine_func, "deleteStringValue_Engine_func": deleteStringValue_Engine_func, "deleteZoneItem_Engine_func": deleteZoneItem_Engine_func, "deleteZone_Engine_func": deleteZone_Engine_func, "dereferenceAttackEventTargetRefs_Engine_func": dereferenceAttackEventTargetRefs_Engine_func, "dereferenceEquipmentSetEquipmentRefs_Engine_func": dereferenceEquipmentSetEquipmentRefs_Engine_func, "dereferenceItemBoundToRefs_Engine_func": dereferenceItemBoundToRefs_Engine_func, "dereferencePlayerEquipmentSetRefs_Engine_func": dereferencePlayerEquipmentSetRefs_Engine_func, "dereferencePlayerGuildMemberRefs_Engine_func": dereferencePlayerGuildMemberRefs_Engine_func, "dereferencePlayerTargetRefsPlayer_Engine_func": dereferencePlayerTargetRefsPlayer_Engine_func, "dereferencePlayerTargetRefsZoneItem_Engine_func": dereferencePlayerTargetRefsZoneItem_Engine_func, "dereferencePlayerTargetedByRefsPlayer_Engine_func": dereferencePlayerTargetedByRefsPlayer_Engine_func, "dereferencePlayerTargetedByRefsZoneItem_Engine_func": dereferencePlayerTargetedByRefsZoneItem_Engine_func, "elementReference_type": elementReference_type, "equipmentSetCheckPool_type": equipmentSetCheckPool_type, "equipmentSetCore_type": equipmentSetCore_type, "equipmentSetEquipmentRefCheckPool_type": equipmentSetEquipmentRefCheckPool_type, "equipmentSetEquipmentRefCore_type": equipmentSetEquipmentRefCore_type, "equipmentSetEquipmentRefIDSlicePool_type": equipmentSetEquipmentRefIDSlicePool_type, "equipmentSetEquipmentRef_Engine_func": equipmentSetEquipmentRef_Engine_func, "equipmentSetIDSlicePool_type": equipmentSetIDSlicePool_type, "equipmentSet_type": equipmentSet_type, "extendAndCopy_path_func": extendAndCopy_path_func, "fill_assemblePlanner_func": fill_assemblePlanner_func, "floatValue_Engine_func": floatValue_Engine_func, "floatValue_type": floatValue_type, "gearScoreCheckPool_type": gearScoreCheckPool_type, "gearScoreCore_type": gearScoreCore_type, "gearScoreIDSlicePool_type": gearScoreIDSlicePool_type, "gearScore_type": gearScore_type, "intValue_Engine_func": intValue_Engine_func, "intValue_type": intValue_type, "isSliceFieldIdentifier_func": isSliceFieldIdentifier_func, "itemBoundToRefCheckPool_type": itemBoundToRefCheckPool_type, "itemBoundToRefCore_type": itemBoundToRefCore_type, "itemBoundToRefIDSlicePool_type": itemBoundToRefIDSlicePool_type, "itemBoundToRef_Engine_func": itemBoundToRef_Engine_func, "itemCheckPool_type": itemCheckPool_type, "itemCore_type": itemCore_type, "itemIDSlicePool_type": itemIDSlicePool_type, "item_type": item_type, "newAssemblePlanner_func": newAssemblePlanner_func, "newPath_func": newPath_func, "newState_func": newState_func, "newTree_func": newTree_func, "path_type": path_type, "plan_assemblePlanner_func": plan_assemblePlanner_func, "playerCheckPool_type": playerCheckPool_type, "playerCore_type": playerCore_type, "playerEquipmentSetRefCheckPool_type": playerEquipmentSetRefCheckPool_type, "playerEquipmentSetRefCore_type": playerEquipmentSetRefCore_type, "playerEquipmentSetRefIDSlicePool_type": playerEquipmentSetRefIDSlicePool_type, "playerEquipmentSetRef_Engine_func": playerEquipmentSetRef_Engine_func, "playerGuildMemberRefCheckPool_type": playerGuildMemberRefCheckPool_type, "playerGuildMemberRefCore_type": playerGuildMemberRefCore_type, "playerGuildMemberRefIDSlicePool_type": playerGuildMemberRefIDSlicePool_type, "playerGuildMemberRef_Engine_func": playerGuildMemberRef_Engine_func, "playerIDSlicePool_type": playerIDSlicePool_type, "playerTargetRefCheckPool_type": playerTargetRefCheckPool_type, "playerTargetRefCore_type": playerTargetRefCore_type, "playerTargetRefIDSlicePool_type": playerTargetRefIDSlicePool_type, "playerTargetRef_Engine_func": playerTargetRef_Engine_func, "playerTargetedByRefCheckPool_type": playerTargetedByRefCheckPool_type, "playerTargetedByRefCore_type": playerTargetedByRefCore_type, "playerTargetedByRefIDSlicePool_type": playerTargetedByRefIDSlicePool_type, "playerTargetedByRef_Engine_func": playerTargetedByRef_Engine_func, "player_type": player_type, "positionCheckPool_type": positionCheckPool_type, "positionCore_type": positionCore_type, "positionIDSlicePool_type": positionIDSlicePool_type, "position_type": position_type, "segment_type": segment_type, "setBoolValue_Engine_func": setBoolValue_Engine_func, "setFloatValue_Engine_func": setFloatValue_Engine_func, "setIntValue_Engine_func": setIntValue_Engine_func, "setStringValue_Engine_func": setStringValue_Engine_func, "stringValue_Engine_func": stringValue_Engine_func, "stringValue_type": stringValue_type, "toJSONPath_path_func": toJSONPath_path_func, "toString_treeFieldIdentifier_func": toString_treeFieldIdentifier_func, "treeFieldIdentifier_type": treeFieldIdentifier_type, "zoneCheckPool_type": zoneCheckPool_type, "zoneCore_type": zoneCore_type, "zoneIDSlicePool_type": zoneIDSlicePool_type, "zoneItemCheckPool_type": zoneItemCheckPool_type, "zoneItemCore_type": zoneItemCore_type, "zoneItemIDSlicePool_type": zoneItemIDSlicePool_type, "zoneItem_type": zoneItem_type, "zone_type": zone_type}
