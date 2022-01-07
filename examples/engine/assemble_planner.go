package state

type assemblePlanner struct {
	updatedPaths          map[int]path
	updatedReferencePaths map[int]path
	updatedElementPaths   map[int]path
	includedElements      map[int]bool
	equipmentSetPath      map[int]path
	gearScorePath         map[int]path
	itemPath              map[int]path
	playerPath            map[int]path
	positionPath          map[int]path
	zonePath              map[int]path
	zoneItemPath          map[int]path
}

func newAssemblePlanner() *assemblePlanner {
	return &assemblePlanner{
		updatedPaths:          make(map[int]path),
		updatedElementPaths:   make(map[int]path),
		updatedReferencePaths: make(map[int]path),
		includedElements:      make(map[int]bool),
		equipmentSetPath:      make(map[int]path),
		gearScorePath:         make(map[int]path),
		itemPath:              make(map[int]path),
		playerPath:            make(map[int]path),
		positionPath:          make(map[int]path),
		zonePath:              make(map[int]path),
		zoneItemPath:          make(map[int]path),
	}
}

func (a *assemblePlanner) clear() {
	for key := range a.updatedPaths {
		delete(a.updatedPaths, key)
	}
	for key := range a.updatedElementPaths {
		delete(a.updatedElementPaths, key)
	}
	for key := range a.updatedReferencePaths {
		delete(a.updatedReferencePaths, key)
	}
	for key := range a.includedElements {
		delete(a.includedElements, key)
	}
	for key := range a.equipmentSetPath {
		delete(a.equipmentSetPath, key)
	}
	for key := range a.gearScorePath {
		delete(a.gearScorePath, key)
	}
	for key := range a.itemPath {
		delete(a.itemPath, key)
	}
	for key := range a.playerPath {
		delete(a.playerPath, key)
	}
	for key := range a.positionPath {
		delete(a.positionPath, key)
	}
	for key := range a.zonePath {
		delete(a.zonePath, key)
	}
	for key := range a.zoneItemPath {
		delete(a.zoneItemPath, key)
	}
}

func (ap *assemblePlanner) plan(state, patch *State) {
	// we want to find all nodes which have updated and collect their paths.
	// later we will loop over the paths we have collected, and "walk" them (assembleBranch)
	// in order to assemble the tree from top to bottom (leaf nodes to root nodes)
	for _, equipmentSet := range patch.EquipmentSet {
		ap.updatedElementPaths[int(equipmentSet.ID)] = equipmentSet.path
	}
	for _, gearScore := range patch.GearScore {
		ap.updatedElementPaths[int(gearScore.ID)] = gearScore.path
	}
	for _, item := range patch.Item {
		ap.updatedElementPaths[int(item.ID)] = item.path
	}
	for _, player := range patch.Player {
		ap.updatedElementPaths[int(player.ID)] = player.path
	}
	for _, position := range patch.Position {
		ap.updatedElementPaths[int(position.ID)] = position.path
	}
	for _, zone := range patch.Zone {
		ap.updatedElementPaths[int(zone.ID)] = zone.path
	}
	for _, zoneItem := range patch.ZoneItem {
		ap.updatedElementPaths[int(zoneItem.ID)] = zoneItem.path
	}
	for _, equipmentSetEquipmentRef := range patch.EquipmentSetEquipmentRef {
		ap.updatedReferencePaths[int(equipmentSetEquipmentRef.ID)] = equipmentSetEquipmentRef.path
	}
	for _, itemBoundToRef := range patch.ItemBoundToRef {
		ap.updatedReferencePaths[int(itemBoundToRef.ID)] = itemBoundToRef.path
	}
	for _, playerEquipmentSetRef := range patch.PlayerEquipmentSetRef {
		ap.updatedReferencePaths[int(playerEquipmentSetRef.ID)] = playerEquipmentSetRef.path
	}
	for _, playerGuildMemberRef := range patch.PlayerGuildMemberRef {
		ap.updatedReferencePaths[int(playerGuildMemberRef.ID)] = playerGuildMemberRef.path
	}
	for _, playerTargetRef := range patch.PlayerTargetRef {
		ap.updatedReferencePaths[int(playerTargetRef.ID)] = playerTargetRef.path
	}
	for _, playerTargetedByRef := range patch.PlayerTargetedByRef {
		ap.updatedReferencePaths[int(playerTargetedByRef.ID)] = playerTargetedByRef.path
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
		for _, p := range ap.updatedElementPaths {
			for _, seg := range p {
				ap.includedElements[seg.id] = true
			}
		}
		// add all elements of the updated reference paths to the includedElements
		for _, p := range ap.updatedReferencePaths {
			for _, seg := range p {
				if seg.refID != 0 {
					ap.includedElements[seg.refID] = true
				} else {
					ap.includedElements[seg.id] = true
				}
			}
		}

		// we check if any new elements are involved, which could
		// mean that new paths containing references need to be looked at
		if previousLen == len(ap.includedElements) {
			break
		}

		previousLen = len(ap.includedElements)

		for _, equipmentSetEquipmentRef := range patch.EquipmentSetEquipmentRef {
			// if the reference references an element that has updated its path is collected
			// so that all segments can later be added to includedElements
			if _, ok := ap.includedElements[int(equipmentSetEquipmentRef.ReferencedElementID)]; ok {
				ap.updatedReferencePaths[int(equipmentSetEquipmentRef.ID)] = equipmentSetEquipmentRef.path
			}
		}
		// we also loop over all references in state because a reference which may not have updated
		// itself may still reference an element which has updated
		for _, equipmentSetEquipmentRef := range state.EquipmentSetEquipmentRef {
			if _, ok := ap.updatedReferencePaths[int(equipmentSetEquipmentRef.ID)]; ok {
				// we don't need to do the check if the reference is already included
				continue
			}
			if _, ok := ap.includedElements[int(equipmentSetEquipmentRef.ReferencedElementID)]; ok {
				ap.updatedReferencePaths[int(equipmentSetEquipmentRef.ID)] = equipmentSetEquipmentRef.path
			}
		}

		for _, itemBoundToRef := range patch.ItemBoundToRef {
			if _, ok := ap.includedElements[int(itemBoundToRef.ReferencedElementID)]; ok {
				ap.updatedReferencePaths[int(itemBoundToRef.ID)] = itemBoundToRef.path
			}
		}
		for _, itemBoundToRef := range state.ItemBoundToRef {
			if _, ok := ap.updatedReferencePaths[int(itemBoundToRef.ID)]; ok {
				continue
			}
			if _, ok := ap.includedElements[int(itemBoundToRef.ReferencedElementID)]; ok {
				ap.updatedReferencePaths[int(itemBoundToRef.ID)] = itemBoundToRef.path
			}
		}

		for _, playerEquipmentSetRef := range patch.PlayerEquipmentSetRef {
			if _, ok := ap.includedElements[int(playerEquipmentSetRef.ReferencedElementID)]; ok {
				ap.updatedReferencePaths[int(playerEquipmentSetRef.ID)] = playerEquipmentSetRef.path
			}
		}
		for _, playerEquipmentSetRef := range state.PlayerEquipmentSetRef {
			if _, ok := ap.updatedReferencePaths[int(playerEquipmentSetRef.ID)]; ok {
				continue
			}
			if _, ok := ap.includedElements[int(playerEquipmentSetRef.ReferencedElementID)]; ok {
				ap.updatedReferencePaths[int(playerEquipmentSetRef.ID)] = playerEquipmentSetRef.path
			}
		}

		for _, playerGuildMemberRef := range patch.PlayerGuildMemberRef {
			if _, ok := ap.includedElements[int(playerGuildMemberRef.ReferencedElementID)]; ok {
				ap.updatedReferencePaths[int(playerGuildMemberRef.ID)] = playerGuildMemberRef.path
			}
		}
		for _, playerGuildMemberRef := range state.PlayerGuildMemberRef {
			if _, ok := ap.updatedReferencePaths[int(playerGuildMemberRef.ID)]; ok {
				continue
			}
			if _, ok := ap.includedElements[int(playerGuildMemberRef.ReferencedElementID)]; ok {
				ap.updatedReferencePaths[int(playerGuildMemberRef.ID)] = playerGuildMemberRef.path
			}
		}

		for _, playerTargetRef := range patch.PlayerTargetRef {
			// if the reference exists in the patch, the anyContainer HAS to exist in patch as well
			// as both are always created and destroyed on unison
			anyContainer := patch.AnyOfPlayer_ZoneItem[playerTargetRef.ReferencedElementID]
			switch anyContainer.ElementKind {
			case ElementKindPlayer:
				if _, ok := ap.includedElements[int(anyContainer.Player)]; ok {
					ap.updatedReferencePaths[int(playerTargetRef.ID)] = playerTargetRef.path
				}
			case ElementKindZoneItem:
				if _, ok := ap.includedElements[int(anyContainer.ZoneItem)]; ok {
					ap.updatedReferencePaths[int(playerTargetRef.ID)] = playerTargetRef.path
				}
			}
		}
		for _, playerTargetRef := range state.PlayerTargetRef {
			if _, ok := ap.updatedReferencePaths[int(playerTargetRef.ID)]; ok {
				continue
			}
			// if the reference exists in the state, the anyContainer HAS to exist in state as well
			// as both are always created and destroyed on unison
			anyContainer := state.AnyOfPlayer_ZoneItem[playerTargetRef.ReferencedElementID]
			switch anyContainer.ElementKind {
			case ElementKindPlayer:
				if _, ok := ap.includedElements[int(anyContainer.Player)]; ok {
					ap.updatedReferencePaths[int(playerTargetRef.ID)] = playerTargetRef.path
				}
			case ElementKindZoneItem:
				if _, ok := ap.includedElements[int(anyContainer.ZoneItem)]; ok {
					ap.updatedReferencePaths[int(playerTargetRef.ID)] = playerTargetRef.path
				}
			}
		}

		for _, playerTargetedByRef := range patch.PlayerTargetedByRef {
			anyContainer := patch.AnyOfPlayer_ZoneItem[playerTargetedByRef.ReferencedElementID]
			switch anyContainer.ElementKind {
			case ElementKindPlayer:
				if _, ok := ap.includedElements[int(anyContainer.Player)]; ok {
					ap.updatedReferencePaths[int(playerTargetedByRef.ID)] = playerTargetedByRef.path
				}
			case ElementKindZoneItem:
				if _, ok := ap.includedElements[int(anyContainer.ZoneItem)]; ok {
					ap.updatedReferencePaths[int(playerTargetedByRef.ID)] = playerTargetedByRef.path
				}
			}
		}
		for _, playerTargetedByRef := range state.PlayerTargetedByRef {
			if _, ok := ap.updatedReferencePaths[int(playerTargetedByRef.ID)]; ok {
				continue
			}
			anyContainer := state.AnyOfPlayer_ZoneItem[playerTargetedByRef.ReferencedElementID]
			switch anyContainer.ElementKind {
			case ElementKindPlayer:
				if _, ok := ap.includedElements[int(anyContainer.Player)]; ok {
					ap.updatedReferencePaths[int(playerTargetedByRef.ID)] = playerTargetedByRef.path
				}
			case ElementKindZoneItem:
				if _, ok := ap.includedElements[int(anyContainer.ZoneItem)]; ok {
					ap.updatedReferencePaths[int(playerTargetedByRef.ID)] = playerTargetedByRef.path
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
	for id, path := range ap.updatedElementPaths {
		ap.updatedPaths[id] = path
	}
	for id, path := range ap.updatedReferencePaths {
		ap.updatedPaths[id] = path
	}

	// just to be a bit more organized
	for leafElementID, p := range ap.updatedPaths {
		switch p[0].identifier {
		case equipmentSetIdentifier:
			ap.equipmentSetPath[leafElementID] = p
		case gearScoreIdentifier:
			ap.gearScorePath[leafElementID] = p
		case itemIdentifier:
			ap.itemPath[leafElementID] = p
		case playerIdentifier:
			ap.playerPath[leafElementID] = p
		case positionIdentifier:
			ap.positionPath[leafElementID] = p
		case zoneIdentifier:
			ap.zonePath[leafElementID] = p
		case zoneItemIdentifier:
			ap.zoneItemPath[leafElementID] = p
		}
	}
}
