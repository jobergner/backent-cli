package state

import "sync"

// 1. get all basic elements and references out of patch, put their paths in updatedReferencePaths
// 2. go through all paths and put ids in includedElements map, save len(includedElements)
// 3. get all references out of STATE, check if they reference element in includedElements, if TRUE put reference path into updatedByReferecePaths
// 4. if len(updatedByReferecePaths) != 0: go through all updatedByReferecePaths and put ids in includedElements map, ELSE continue with 6.
// 5. back to step 3.

// TODO PROBLEM?? if a reference is Set and then Unset and Set again, does that mess with pathuilding, as referenceID might be 0 in patch or state
// TODO what happens if you call SetPlayer? will the path be built with a player-update or a position-delete??

func (engine *Engine) assembleEquipmentSets(wg *sync.WaitGroup) {
	for id, p := range engine.assembler.equipmentSetPath {
		child, ok := engine.Tree.EquipmentSet[id]
		if !ok {
			child = equipmentSet{ID: id}
		}
		engine.assembleEquipmentSetPath(&child, p, 0, engine.assembler.includedElements)
		engine.Tree.EquipmentSet[id] = child
	}
	wg.Done()
}

func (engine *Engine) assembleGearScores(wg *sync.WaitGroup) {
	for id, p := range engine.assembler.gearScorePath {
		child, ok := engine.Tree.GearScore[id]
		if !ok {
			child = gearScore{ID: id}
		}
		engine.assembleGearScorePath(&child, p, 0, engine.assembler.includedElements)
		engine.Tree.GearScore[id] = child
	}
	wg.Done()
}

func (engine *Engine) assembleItems(wg *sync.WaitGroup) {
	for id, p := range engine.assembler.itemPath {
		child, ok := engine.Tree.Item[id]
		if !ok {
			child = item{ID: id}
		}
		engine.assembleItemPath(&child, p, 0, engine.assembler.includedElements)
		engine.Tree.Item[id] = child
	}
	wg.Done()
}

func (engine *Engine) assemblePlayers(wg *sync.WaitGroup) {
	for id, p := range engine.assembler.playerPath {
		child, ok := engine.Tree.Player[id]
		if !ok {
			child = player{ID: id}
		}
		engine.assemblePlayerPath(&child, p, 0, engine.assembler.includedElements)
		engine.Tree.Player[id] = child
	}
	wg.Done()
}

func (engine *Engine) assemblePositions(wg *sync.WaitGroup) {
	for id, p := range engine.assembler.positionPath {
		child, ok := engine.Tree.Position[id]
		if !ok {
			child = position{ID: id}
		}
		engine.assemblePositionPath(&child, p, 0, engine.assembler.includedElements)
		engine.Tree.Position[id] = child
	}
	wg.Done()
}

func (engine *Engine) assembleZones(wg *sync.WaitGroup) {
	for id, p := range engine.assembler.zonePath {
		child, ok := engine.Tree.Zone[id]
		if !ok {
			child = zone{ID: id}
		}
		engine.assembleZonePath(&child, p, 0, engine.assembler.includedElements)
		engine.Tree.Zone[id] = child
	}
	wg.Done()
}

func (engine *Engine) assembleZoneItems(wg *sync.WaitGroup) {
	for id, p := range engine.assembler.zoneItemPath {
		child, ok := engine.Tree.ZoneItem[id]
		if !ok {
			child = zoneItem{ID: id}
		}
		engine.assembleZoneItemPath(&child, p, 0, engine.assembler.includedElements)
		engine.Tree.ZoneItem[id] = child
	}
	wg.Done()
}

func (engine *Engine) assembleUpdateTree() Tree {

	engine.clearAssembler()
	engine.clearTree()

	engine.populateAssembler()

	var wg sync.WaitGroup

	wg.Add(7)

	go engine.assembleEquipmentSets(&wg)
	go engine.assembleGearScores(&wg)
	go engine.assembleItems(&wg)
	go engine.assemblePlayers(&wg)
	go engine.assemblePositions(&wg)
	go engine.assembleZones(&wg)
	go engine.assembleZoneItems(&wg)

	wg.Wait()

	return engine.Tree
}

func (engine *Engine) clearAssembler() {
	for key := range engine.assembler.updatedPaths {
		delete(engine.assembler.updatedPaths, key)
	}
	for key := range engine.assembler.updatedElementPaths {
		delete(engine.assembler.updatedElementPaths, key)
	}
	for key := range engine.assembler.updatedReferencePaths {
		delete(engine.assembler.updatedReferencePaths, key)
	}
	for key := range engine.assembler.includedElements {
		delete(engine.assembler.includedElements, key)
	}

	for key := range engine.assembler.equipmentSetPath {
		delete(engine.assembler.equipmentSetPath, key)
	}
	for key := range engine.assembler.gearScorePath {
		delete(engine.assembler.gearScorePath, key)
	}
	for key := range engine.assembler.itemPath {
		delete(engine.assembler.itemPath, key)
	}
	for key := range engine.assembler.playerPath {
		delete(engine.assembler.playerPath, key)
	}
	for key := range engine.assembler.positionPath {
		delete(engine.assembler.positionPath, key)
	}
	for key := range engine.assembler.zoneItemPath {
		delete(engine.assembler.zoneItemPath, key)
	}
	for key := range engine.assembler.zonePath {
		delete(engine.assembler.zonePath, key)
	}
}

func (engine *Engine) clearTree() {
	for key := range engine.Tree.EquipmentSet {
		delete(engine.Tree.EquipmentSet, key)
	}
	for key := range engine.Tree.GearScore {
		delete(engine.Tree.GearScore, key)
	}
	for key := range engine.Tree.Item {
		delete(engine.Tree.Item, key)
	}
	for key := range engine.Tree.Player {
		delete(engine.Tree.Player, key)
	}
	for key := range engine.Tree.Position {
		delete(engine.Tree.Position, key)
	}
	for key := range engine.Tree.Zone {
		delete(engine.Tree.Zone, key)
	}
	for key := range engine.Tree.ZoneItem {
		delete(engine.Tree.ZoneItem, key)
	}
}

func (engine *Engine) populateAssembler() {
	// TODO possibly big performance boost
	for _, equipmentSet := range engine.Patch.EquipmentSet {
		engine.assembler.updatedElementPaths[int(equipmentSet.ID)] = equipmentSet.path
	}
	for _, gearScore := range engine.Patch.GearScore {
		engine.assembler.updatedElementPaths[int(gearScore.ID)] = gearScore.path
	}
	for _, item := range engine.Patch.Item {
		engine.assembler.updatedElementPaths[int(item.ID)] = item.path
	}
	for _, player := range engine.Patch.Player {
		engine.assembler.updatedElementPaths[int(player.ID)] = player.path
	}
	for _, position := range engine.Patch.Position {
		engine.assembler.updatedElementPaths[int(position.ID)] = position.path
	}
	for _, zone := range engine.Patch.Zone {
		engine.assembler.updatedElementPaths[int(zone.ID)] = zone.path
	}
	for _, zoneItem := range engine.Patch.ZoneItem {
		engine.assembler.updatedElementPaths[int(zoneItem.ID)] = zoneItem.path
	}
	for _, equipmentSetEquipmentRef := range engine.Patch.EquipmentSetEquipmentRef {
		engine.assembler.updatedReferencePaths[int(equipmentSetEquipmentRef.ID)] = equipmentSetEquipmentRef.path
	}
	for _, itemBoundToRef := range engine.Patch.ItemBoundToRef {
		engine.assembler.updatedReferencePaths[int(itemBoundToRef.ID)] = itemBoundToRef.path
	}
	for _, playerEquipmentSetRef := range engine.Patch.PlayerEquipmentSetRef {
		engine.assembler.updatedReferencePaths[int(playerEquipmentSetRef.ID)] = playerEquipmentSetRef.path
	}
	for _, playerGuildMemberRef := range engine.Patch.PlayerGuildMemberRef {
		engine.assembler.updatedReferencePaths[int(playerGuildMemberRef.ID)] = playerGuildMemberRef.path
	}
	for _, playerTargetRef := range engine.Patch.PlayerTargetRef {
		engine.assembler.updatedReferencePaths[int(playerTargetRef.ID)] = playerTargetRef.path
	}
	for _, playerTargetedByRef := range engine.Patch.PlayerTargetedByRef {
		engine.assembler.updatedReferencePaths[int(playerTargetedByRef.ID)] = playerTargetedByRef.path
	}

	prevousLength := 0
	for {
		for _, p := range engine.assembler.updatedElementPaths {
			for _, seg := range p {
				engine.assembler.includedElements[seg.id] = true
			}
		}
		for _, p := range engine.assembler.updatedReferencePaths {
			for _, seg := range p {
				if seg.refID != 0 {
					engine.assembler.includedElements[seg.refID] = true
				} else {
					engine.assembler.includedElements[seg.id] = true
				}
			}
		}

		if prevousLength == len(engine.assembler.includedElements) {
			break
		}

		prevousLength = len(engine.assembler.includedElements)

		for _, equipmentSetEquipmentRef := range engine.Patch.EquipmentSetEquipmentRef {
			if _, ok := engine.assembler.includedElements[int(equipmentSetEquipmentRef.ReferencedElementID)]; ok {
				engine.assembler.updatedReferencePaths[int(equipmentSetEquipmentRef.ID)] = equipmentSetEquipmentRef.path
			}
		}
		for _, equipmentSetEquipmentRef := range engine.State.EquipmentSetEquipmentRef {
			if _, ok := engine.assembler.updatedReferencePaths[int(equipmentSetEquipmentRef.ID)]; ok {
				continue
			}
			if _, ok := engine.assembler.includedElements[int(equipmentSetEquipmentRef.ReferencedElementID)]; ok {
				engine.assembler.updatedReferencePaths[int(equipmentSetEquipmentRef.ID)] = equipmentSetEquipmentRef.path
			}
		}

		for _, itemBoundToRef := range engine.Patch.ItemBoundToRef {
			if _, ok := engine.assembler.includedElements[int(itemBoundToRef.ReferencedElementID)]; ok {
				engine.assembler.updatedReferencePaths[int(itemBoundToRef.ID)] = itemBoundToRef.path
			}
		}
		for _, itemBoundToRef := range engine.State.ItemBoundToRef {
			if _, ok := engine.assembler.updatedReferencePaths[int(itemBoundToRef.ID)]; ok {
				continue
			}
			if _, ok := engine.assembler.includedElements[int(itemBoundToRef.ReferencedElementID)]; ok {
				engine.assembler.updatedReferencePaths[int(itemBoundToRef.ID)] = itemBoundToRef.path
			}
		}

		for _, playerEquipmentSetRef := range engine.Patch.PlayerEquipmentSetRef {
			if _, ok := engine.assembler.includedElements[int(playerEquipmentSetRef.ReferencedElementID)]; ok {
				engine.assembler.updatedReferencePaths[int(playerEquipmentSetRef.ID)] = playerEquipmentSetRef.path
			}
		}
		for _, playerEquipmentSetRef := range engine.State.PlayerEquipmentSetRef {
			if _, ok := engine.assembler.updatedReferencePaths[int(playerEquipmentSetRef.ID)]; ok {
				continue
			}
			if _, ok := engine.assembler.includedElements[int(playerEquipmentSetRef.ReferencedElementID)]; ok {
				engine.assembler.updatedReferencePaths[int(playerEquipmentSetRef.ID)] = playerEquipmentSetRef.path
			}
		}

		for _, playerGuildMemberRef := range engine.Patch.PlayerGuildMemberRef {
			if _, ok := engine.assembler.includedElements[int(playerGuildMemberRef.ReferencedElementID)]; ok {
				engine.assembler.updatedReferencePaths[int(playerGuildMemberRef.ID)] = playerGuildMemberRef.path
			}
		}
		for _, playerGuildMemberRef := range engine.State.PlayerGuildMemberRef {
			if _, ok := engine.assembler.updatedReferencePaths[int(playerGuildMemberRef.ID)]; ok {
				continue
			}
			if _, ok := engine.assembler.includedElements[int(playerGuildMemberRef.ReferencedElementID)]; ok {
				engine.assembler.updatedReferencePaths[int(playerGuildMemberRef.ID)] = playerGuildMemberRef.path
			}
		}

		for _, playerTargetRef := range engine.Patch.PlayerTargetRef {
			anyContainer := engine.anyOfPlayer_ZoneItem(playerTargetRef.ReferencedElementID)
			switch anyContainer.anyOfPlayer_ZoneItem.ElementKind {
			case ElementKindPlayer:
				if _, ok := engine.assembler.includedElements[int(anyContainer.anyOfPlayer_ZoneItem.Player)]; ok {
					engine.assembler.updatedReferencePaths[int(playerTargetRef.ID)] = playerTargetRef.path
				}
			case ElementKindZoneItem:
				if _, ok := engine.assembler.includedElements[int(anyContainer.anyOfPlayer_ZoneItem.ZoneItem)]; ok {
					engine.assembler.updatedReferencePaths[int(playerTargetRef.ID)] = playerTargetRef.path
				}
			}
		}
		for _, playerTargetRef := range engine.State.PlayerTargetRef {
			if _, ok := engine.assembler.updatedReferencePaths[int(playerTargetRef.ID)]; ok {
				continue
			}
			anyContainer := engine.anyOfPlayer_ZoneItem(playerTargetRef.ReferencedElementID)
			switch anyContainer.anyOfPlayer_ZoneItem.ElementKind {
			case ElementKindPlayer:
				if _, ok := engine.assembler.includedElements[int(anyContainer.anyOfPlayer_ZoneItem.Player)]; ok {
					engine.assembler.updatedReferencePaths[int(playerTargetRef.ID)] = playerTargetRef.path
				}
			case ElementKindZoneItem:
				if _, ok := engine.assembler.includedElements[int(anyContainer.anyOfPlayer_ZoneItem.Player)]; ok {
					engine.assembler.updatedReferencePaths[int(playerTargetRef.ID)] = playerTargetRef.path
				}
			}
		}

		for _, playerTargetedByRef := range engine.Patch.PlayerTargetedByRef {
			anyContainer := engine.anyOfPlayer_ZoneItem(playerTargetedByRef.ReferencedElementID)
			switch anyContainer.anyOfPlayer_ZoneItem.ElementKind {
			case ElementKindPlayer:
				if _, ok := engine.assembler.includedElements[int(anyContainer.anyOfPlayer_ZoneItem.Player)]; ok {
					engine.assembler.updatedReferencePaths[int(playerTargetedByRef.ID)] = playerTargetedByRef.path
				}
			case ElementKindZoneItem:
				if _, ok := engine.assembler.includedElements[int(anyContainer.anyOfPlayer_ZoneItem.Player)]; ok {
					engine.assembler.updatedReferencePaths[int(playerTargetedByRef.ID)] = playerTargetedByRef.path
				}
			}
		}
		for _, playerTargetedByRef := range engine.State.PlayerTargetedByRef {
			if _, ok := engine.assembler.updatedReferencePaths[int(playerTargetedByRef.ID)]; ok {
				continue
			}
			anyContainer := engine.anyOfPlayer_ZoneItem(playerTargetedByRef.ReferencedElementID)
			switch anyContainer.anyOfPlayer_ZoneItem.ElementKind {
			case ElementKindPlayer:
				if _, ok := engine.assembler.includedElements[int(anyContainer.anyOfPlayer_ZoneItem.Player)]; ok {
					engine.assembler.updatedReferencePaths[int(playerTargetedByRef.ID)] = playerTargetedByRef.path
				}
			case ElementKindZoneItem:
				if _, ok := engine.assembler.includedElements[int(anyContainer.anyOfPlayer_ZoneItem.Player)]; ok {
					engine.assembler.updatedReferencePaths[int(playerTargetedByRef.ID)] = playerTargetedByRef.path
				}
			}
		}
	}

	for id, path := range engine.assembler.updatedElementPaths {
		engine.assembler.updatedPaths[id] = path
	}
	for id, path := range engine.assembler.updatedReferencePaths {
		engine.assembler.updatedPaths[id] = path
	}

	for _, p := range engine.assembler.updatedPaths {
		switch p[0].identifier {
		case equipmentSetIdentifier:
			engine.assembler.equipmentSetPath[EquipmentSetID(p[0].id)] = p
		case gearScoreIdentifier:
			engine.assembler.gearScorePath[GearScoreID(p[0].id)] = p
		case itemIdentifier:
			engine.assembler.itemPath[ItemID(p[0].id)] = p
		case playerIdentifier:
			engine.assembler.playerPath[PlayerID(p[0].id)] = p
		case positionIdentifier:
			engine.assembler.positionPath[PositionID(p[0].id)] = p
		case zoneIdentifier:
			engine.assembler.zonePath[ZoneID(p[0].id)] = p
		case zoneItemIdentifier:
			engine.assembler.zoneItemPath[ZoneItemID(p[0].id)] = p
		}
	}
}
