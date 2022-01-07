package state

func (engine *Engine) assembleUpdateTree() {

	engine.planner.clear()
	engine.Tree.clear()

	engine.populateAssembler()

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

func (engine *Engine) populateAssembler() {
	// we want to find all nodes which have updated and collect their paths.
	// later we will loop over the paths we have collected, and "walk" them (assembleBranch)
	// in order to assemble the tree from top to bottom (leaf nodes to root nodes)
	for _, equipmentSet := range engine.Patch.EquipmentSet {
		engine.planner.updatedElementPaths[int(equipmentSet.ID)] = equipmentSet.path
	}
	for _, gearScore := range engine.Patch.GearScore {
		engine.planner.updatedElementPaths[int(gearScore.ID)] = gearScore.path
	}
	for _, item := range engine.Patch.Item {
		engine.planner.updatedElementPaths[int(item.ID)] = item.path
	}
	for _, player := range engine.Patch.Player {
		engine.planner.updatedElementPaths[int(player.ID)] = player.path
	}
	for _, position := range engine.Patch.Position {
		engine.planner.updatedElementPaths[int(position.ID)] = position.path
	}
	for _, zone := range engine.Patch.Zone {
		engine.planner.updatedElementPaths[int(zone.ID)] = zone.path
	}
	for _, zoneItem := range engine.Patch.ZoneItem {
		engine.planner.updatedElementPaths[int(zoneItem.ID)] = zoneItem.path
	}
	for _, equipmentSetEquipmentRef := range engine.Patch.EquipmentSetEquipmentRef {
		engine.planner.updatedReferencePaths[int(equipmentSetEquipmentRef.ID)] = equipmentSetEquipmentRef.path
	}
	for _, itemBoundToRef := range engine.Patch.ItemBoundToRef {
		engine.planner.updatedReferencePaths[int(itemBoundToRef.ID)] = itemBoundToRef.path
	}
	for _, playerEquipmentSetRef := range engine.Patch.PlayerEquipmentSetRef {
		engine.planner.updatedReferencePaths[int(playerEquipmentSetRef.ID)] = playerEquipmentSetRef.path
	}
	for _, playerGuildMemberRef := range engine.Patch.PlayerGuildMemberRef {
		engine.planner.updatedReferencePaths[int(playerGuildMemberRef.ID)] = playerGuildMemberRef.path
	}
	for _, playerTargetRef := range engine.Patch.PlayerTargetRef {
		engine.planner.updatedReferencePaths[int(playerTargetRef.ID)] = playerTargetRef.path
	}
	for _, playerTargetedByRef := range engine.Patch.PlayerTargetedByRef {
		engine.planner.updatedReferencePaths[int(playerTargetedByRef.ID)] = playerTargetedByRef.path
	}

	previousLen := 0
	// we'd be pretty much done collecting the required paths, but we also want to
	// build all paths ending with a reference which references an updated element.
	// (and want to populate includedElements map of course)
	// this needs to happen recursively, consider this example (-> = reference):
	// A -> B -> C -> ^D
	// Since "D" has updated (^), but no other element, we'd only include "C".
	// However, now that "C" is considered updated by reference, we also want
	// to include "B". This is why recursiveness is required.
	for {
		// here we populate out includedElements with all newly collected paths segments
		// so we can check if any of these  elements are referenced by any reference
		// in the loop below
		for _, p := range engine.planner.updatedElementPaths {
			for _, seg := range p {
				engine.planner.includedElements[seg.id] = true
			}
		}
		// add all elements of the updated reference paths to the includedElements
		for _, p := range engine.planner.updatedReferencePaths {
			for _, seg := range p {
				if seg.refID != 0 {
					engine.planner.includedElements[seg.refID] = true
				} else {
					engine.planner.includedElements[seg.id] = true
				}
			}
		}

		// we check if any new elements are involved, which could
		// mean that new paths containing references need to be looked at
		if previousLen == len(engine.planner.includedElements) {
			break
		}

		previousLen = len(engine.planner.includedElements)

		for _, equipmentSetEquipmentRef := range engine.Patch.EquipmentSetEquipmentRef {
			// if the reference references an element that has updated its path is collected
			// so that all segments can later be added to includedElements
			if _, ok := engine.planner.includedElements[int(equipmentSetEquipmentRef.ReferencedElementID)]; ok {
				engine.planner.updatedReferencePaths[int(equipmentSetEquipmentRef.ID)] = equipmentSetEquipmentRef.path
			}
		}
		// we also loop over all references in state because a reference which may not have updated
		// itself may still reference an element which has updated
		for _, equipmentSetEquipmentRef := range engine.State.EquipmentSetEquipmentRef {
			if _, ok := engine.planner.updatedReferencePaths[int(equipmentSetEquipmentRef.ID)]; ok {
				// we don't need to do the check if the reference is already included
				continue
			}
			if _, ok := engine.planner.includedElements[int(equipmentSetEquipmentRef.ReferencedElementID)]; ok {
				engine.planner.updatedReferencePaths[int(equipmentSetEquipmentRef.ID)] = equipmentSetEquipmentRef.path
			}
		}

		for _, itemBoundToRef := range engine.Patch.ItemBoundToRef {
			if _, ok := engine.planner.includedElements[int(itemBoundToRef.ReferencedElementID)]; ok {
				engine.planner.updatedReferencePaths[int(itemBoundToRef.ID)] = itemBoundToRef.path
			}
		}
		for _, itemBoundToRef := range engine.State.ItemBoundToRef {
			if _, ok := engine.planner.updatedReferencePaths[int(itemBoundToRef.ID)]; ok {
				continue
			}
			if _, ok := engine.planner.includedElements[int(itemBoundToRef.ReferencedElementID)]; ok {
				engine.planner.updatedReferencePaths[int(itemBoundToRef.ID)] = itemBoundToRef.path
			}
		}

		for _, playerEquipmentSetRef := range engine.Patch.PlayerEquipmentSetRef {
			if _, ok := engine.planner.includedElements[int(playerEquipmentSetRef.ReferencedElementID)]; ok {
				engine.planner.updatedReferencePaths[int(playerEquipmentSetRef.ID)] = playerEquipmentSetRef.path
			}
		}
		for _, playerEquipmentSetRef := range engine.State.PlayerEquipmentSetRef {
			if _, ok := engine.planner.updatedReferencePaths[int(playerEquipmentSetRef.ID)]; ok {
				continue
			}
			if _, ok := engine.planner.includedElements[int(playerEquipmentSetRef.ReferencedElementID)]; ok {
				engine.planner.updatedReferencePaths[int(playerEquipmentSetRef.ID)] = playerEquipmentSetRef.path
			}
		}

		for _, playerGuildMemberRef := range engine.Patch.PlayerGuildMemberRef {
			if _, ok := engine.planner.includedElements[int(playerGuildMemberRef.ReferencedElementID)]; ok {
				engine.planner.updatedReferencePaths[int(playerGuildMemberRef.ID)] = playerGuildMemberRef.path
			}
		}
		for _, playerGuildMemberRef := range engine.State.PlayerGuildMemberRef {
			if _, ok := engine.planner.updatedReferencePaths[int(playerGuildMemberRef.ID)]; ok {
				continue
			}
			if _, ok := engine.planner.includedElements[int(playerGuildMemberRef.ReferencedElementID)]; ok {
				engine.planner.updatedReferencePaths[int(playerGuildMemberRef.ID)] = playerGuildMemberRef.path
			}
		}

		for _, playerTargetRef := range engine.Patch.PlayerTargetRef {
			// if the reference exists in the patch, the anyContainer HAS to exist in patch as well
			// as both are always created and destroyed on unison
			anyContainer := engine.Patch.AnyOfPlayer_ZoneItem[playerTargetRef.ReferencedElementID]
			switch anyContainer.ElementKind {
			case ElementKindPlayer:
				if _, ok := engine.planner.includedElements[int(anyContainer.Player)]; ok {
					engine.planner.updatedReferencePaths[int(playerTargetRef.ID)] = playerTargetRef.path
				}
			case ElementKindZoneItem:
				if _, ok := engine.planner.includedElements[int(anyContainer.ZoneItem)]; ok {
					engine.planner.updatedReferencePaths[int(playerTargetRef.ID)] = playerTargetRef.path
				}
			}
		}
		for _, playerTargetRef := range engine.State.PlayerTargetRef {
			if _, ok := engine.planner.updatedReferencePaths[int(playerTargetRef.ID)]; ok {
				continue
			}
			// if the reference exists in the state, the anyContainer HAS to exist in state as well
			// as both are always created and destroyed on unison
			anyContainer := engine.State.AnyOfPlayer_ZoneItem[playerTargetRef.ReferencedElementID]
			switch anyContainer.ElementKind {
			case ElementKindPlayer:
				if _, ok := engine.planner.includedElements[int(anyContainer.Player)]; ok {
					engine.planner.updatedReferencePaths[int(playerTargetRef.ID)] = playerTargetRef.path
				}
			case ElementKindZoneItem:
				if _, ok := engine.planner.includedElements[int(anyContainer.ZoneItem)]; ok {
					engine.planner.updatedReferencePaths[int(playerTargetRef.ID)] = playerTargetRef.path
				}
			}
		}

		for _, playerTargetedByRef := range engine.Patch.PlayerTargetedByRef {
			anyContainer := engine.Patch.AnyOfPlayer_ZoneItem[playerTargetedByRef.ReferencedElementID]
			switch anyContainer.ElementKind {
			case ElementKindPlayer:
				if _, ok := engine.planner.includedElements[int(anyContainer.Player)]; ok {
					engine.planner.updatedReferencePaths[int(playerTargetedByRef.ID)] = playerTargetedByRef.path
				}
			case ElementKindZoneItem:
				if _, ok := engine.planner.includedElements[int(anyContainer.ZoneItem)]; ok {
					engine.planner.updatedReferencePaths[int(playerTargetedByRef.ID)] = playerTargetedByRef.path
				}
			}
		}
		for _, playerTargetedByRef := range engine.State.PlayerTargetedByRef {
			if _, ok := engine.planner.updatedReferencePaths[int(playerTargetedByRef.ID)]; ok {
				continue
			}
			anyContainer := engine.State.AnyOfPlayer_ZoneItem[playerTargetedByRef.ReferencedElementID]
			switch anyContainer.ElementKind {
			case ElementKindPlayer:
				if _, ok := engine.planner.includedElements[int(anyContainer.Player)]; ok {
					engine.planner.updatedReferencePaths[int(playerTargetedByRef.ID)] = playerTargetedByRef.path
				}
			case ElementKindZoneItem:
				if _, ok := engine.planner.includedElements[int(anyContainer.ZoneItem)]; ok {
					engine.planner.updatedReferencePaths[int(playerTargetedByRef.ID)] = playerTargetedByRef.path
				}
			}
		}
	}

	// TODO: might need this for full assembling
	// for _, path := range engine.assembler.updatedElementPaths {
	// 	if len(path) < 2 {
	// 		continue
	// 	}
	// 	for _, seg := range path[:len(path)-1] {
	// 		delete(engine.assembler.updatedElementPaths, seg.id)
	// 	}
	// }
	// for _, path := range engine.assembler.updatedReferencePaths {
	// 	if len(path) < 2 {
	// 		continue
	// 	}
	// 	for _, seg := range path[:len(path)-1] {
	// 		delete(engine.assembler.updatedReferencePaths, seg.id)
	// 	}
	// }

	// merge paths into one map, for convencience (they are recycled anyway)
	for id, path := range engine.planner.updatedElementPaths {
		engine.planner.updatedPaths[id] = path
	}
	for id, path := range engine.planner.updatedReferencePaths {
		engine.planner.updatedPaths[id] = path
	}

	// just to be a bit more organized
	for leafElementID, p := range engine.planner.updatedPaths {
		switch p[0].identifier {
		case equipmentSetIdentifier:
			engine.planner.equipmentSetPath[leafElementID] = p
		case gearScoreIdentifier:
			engine.planner.gearScorePath[leafElementID] = p
		case itemIdentifier:
			engine.planner.itemPath[leafElementID] = p
		case playerIdentifier:
			engine.planner.playerPath[leafElementID] = p
		case positionIdentifier:
			engine.planner.positionPath[leafElementID] = p
		case zoneIdentifier:
			engine.planner.zonePath[leafElementID] = p
		case zoneItemIdentifier:
			engine.planner.zoneItemPath[leafElementID] = p
		}
	}
}
