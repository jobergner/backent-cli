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
	for _, elementPath := range engine.planner.updatedPaths {
		switch elementPath[0].identifier {
		case equipmentSetIdentifier:
			child, ok := engine.Tree.EquipmentSet[EquipmentSetID(elementPath[0].id)]
			if !ok {
				child = equipmentSet{ID: EquipmentSetID(elementPath[0].id)}
			}
			engine.assembleEquipmentSetPath(&child, elementPath, 0, engine.planner.includedElements)
			engine.Tree.EquipmentSet[EquipmentSetID(elementPath[0].id)] = child
		case gearScoreIdentifier:
			child, ok := engine.Tree.GearScore[GearScoreID(elementPath[0].id)]
			if !ok {
				child = gearScore{ID: GearScoreID(elementPath[0].id)}
			}
			engine.assembleGearScorePath(&child, elementPath, 0, engine.planner.includedElements)
			engine.Tree.GearScore[GearScoreID(elementPath[0].id)] = child
		case itemIdentifier:
			child, ok := engine.Tree.Item[ItemID(elementPath[0].id)]
			if !ok {
				child = item{ID: ItemID(elementPath[0].id)}
			}
			engine.assembleItemPath(&child, elementPath, 0, engine.planner.includedElements)
			engine.Tree.Item[ItemID(elementPath[0].id)] = child
		case playerIdentifier:
			child, ok := engine.Tree.Player[PlayerID(elementPath[0].id)]
			if !ok {
				child = player{ID: PlayerID(elementPath[0].id)}
			}
			engine.assemblePlayerPath(&child, elementPath, 0, engine.planner.includedElements)
			engine.Tree.Player[PlayerID(elementPath[0].id)] = child
		case positionIdentifier:
			child, ok := engine.Tree.Position[PositionID(elementPath[0].id)]
			if !ok {
				child = position{ID: PositionID(elementPath[0].id)}
			}
			engine.assemblePositionPath(&child, elementPath, 0, engine.planner.includedElements)
			engine.Tree.Position[PositionID(elementPath[0].id)] = child
		case zoneIdentifier:
			child, ok := engine.Tree.Zone[ZoneID(elementPath[0].id)]
			if !ok {
				child = zone{ID: ZoneID(elementPath[0].id)}
			}
			engine.assembleZonePath(&child, elementPath, 0, engine.planner.includedElements)
			engine.Tree.Zone[ZoneID(elementPath[0].id)] = child
		case zoneItemIdentifier:
			child, ok := engine.Tree.ZoneItem[ZoneItemID(elementPath[0].id)]
			if !ok {
				child = zoneItem{ID: ZoneItemID(elementPath[0].id)}
			}
			engine.assembleZoneItemPath(&child, elementPath, 0, engine.planner.includedElements)
			engine.Tree.ZoneItem[ZoneItemID(elementPath[0].id)] = child
		}
	}
}
