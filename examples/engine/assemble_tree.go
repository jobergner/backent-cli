package state

func (engine *Engine) assembleUpdateTree() {
	engine.planner.clear()
	engine.Tree.clear()

	engine.planner.plan(engine.State, engine.Patch)

	engine.assembleTree()
}

func (engine *Engine) assembleFullTree() {
	engine.planner.clear()
	engine.Tree.clear()

	engine.planner.fill(engine.State)

	engine.assembleTree()
}

func (engine *Engine) assembleTree() {
	for _, p := range engine.planner.updatedPaths {
		switch p[0].identifier {
		case equipmentSetIdentifier:
			child, ok := engine.Tree.EquipmentSet[EquipmentSetID(p[0].id)]
			if !ok {
				child = equipmentSet{ID: EquipmentSetID(p[0].id)}
			}
			engine.assembleEquipmentSetPath(&child, p, 0, engine.planner.includedElements)
			engine.Tree.EquipmentSet[EquipmentSetID(p[0].id)] = child
		case gearScoreIdentifier:
			child, ok := engine.Tree.GearScore[GearScoreID(p[0].id)]
			if !ok {
				child = gearScore{ID: GearScoreID(p[0].id)}
			}
			engine.assembleGearScorePath(&child, p, 0, engine.planner.includedElements)
			engine.Tree.GearScore[GearScoreID(p[0].id)] = child
		case itemIdentifier:
			child, ok := engine.Tree.Item[ItemID(p[0].id)]
			if !ok {
				child = item{ID: ItemID(p[0].id)}
			}
			engine.assembleItemPath(&child, p, 0, engine.planner.includedElements)
			engine.Tree.Item[ItemID(p[0].id)] = child
		case playerIdentifier:
			child, ok := engine.Tree.Player[PlayerID(p[0].id)]
			if !ok {
				child = player{ID: PlayerID(p[0].id)}
			}
			engine.assemblePlayerPath(&child, p, 0, engine.planner.includedElements)
			engine.Tree.Player[PlayerID(p[0].id)] = child
		case positionIdentifier:
			child, ok := engine.Tree.Position[PositionID(p[0].id)]
			if !ok {
				child = position{ID: PositionID(p[0].id)}
			}
			engine.assemblePositionPath(&child, p, 0, engine.planner.includedElements)
			engine.Tree.Position[PositionID(p[0].id)] = child
		case zoneIdentifier:
			child, ok := engine.Tree.Zone[ZoneID(p[0].id)]
			if !ok {
				child = zone{ID: ZoneID(p[0].id)}
			}
			engine.assembleZonePath(&child, p, 0, engine.planner.includedElements)
			engine.Tree.Zone[ZoneID(p[0].id)] = child
		case zoneItemIdentifier:
			child, ok := engine.Tree.ZoneItem[ZoneItemID(p[0].id)]
			if !ok {
				child = zoneItem{ID: ZoneItemID(p[0].id)}
			}
			engine.assembleZoneItemPath(&child, p, 0, engine.planner.includedElements)
			engine.Tree.ZoneItem[ZoneItemID(p[0].id)] = child
		}
	}
}
