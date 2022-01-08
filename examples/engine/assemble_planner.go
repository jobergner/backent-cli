package state

type assemblePlanner struct {
	updatedPaths          map[int]path
	updatedReferencePaths map[int]path
	updatedElementPaths   map[int]path
	includedElements      map[int]bool
}

func newAssemblePlanner() *assemblePlanner {
	return &assemblePlanner{
		updatedPaths:          make(map[int]path),
		updatedElementPaths:   make(map[int]path),
		updatedReferencePaths: make(map[int]path),
		includedElements:      make(map[int]bool),
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

	// merge paths into one map, for convencience (they are recycled anyway)
	for id, p := range ap.updatedElementPaths {
		ap.updatedPaths[id] = p
	}
	for id, p := range ap.updatedReferencePaths {
		ap.updatedPaths[id] = p
	}
}

func (ap *assemblePlanner) fill(state *State) {
	for _, equipmentSet := range state.EquipmentSet {
		ap.updatedElementPaths[int(equipmentSet.ID)] = equipmentSet.path
	}
	for _, gearScore := range state.GearScore {
		ap.updatedElementPaths[int(gearScore.ID)] = gearScore.path
	}
	for _, item := range state.Item {
		ap.updatedElementPaths[int(item.ID)] = item.path
	}
	for _, player := range state.Player {
		ap.updatedElementPaths[int(player.ID)] = player.path
	}
	for _, position := range state.Position {
		ap.updatedElementPaths[int(position.ID)] = position.path
	}
	for _, zone := range state.Zone {
		ap.updatedElementPaths[int(zone.ID)] = zone.path
	}
	for _, zoneItem := range state.ZoneItem {
		ap.updatedElementPaths[int(zoneItem.ID)] = zoneItem.path
	}
	for _, equipmentSetEquipmentRef := range state.EquipmentSetEquipmentRef {
		ap.updatedReferencePaths[int(equipmentSetEquipmentRef.ID)] = equipmentSetEquipmentRef.path
	}
	for _, itemBoundToRef := range state.ItemBoundToRef {
		ap.updatedReferencePaths[int(itemBoundToRef.ID)] = itemBoundToRef.path
	}
	for _, playerEquipmentSetRef := range state.PlayerEquipmentSetRef {
		ap.updatedReferencePaths[int(playerEquipmentSetRef.ID)] = playerEquipmentSetRef.path
	}
	for _, playerGuildMemberRef := range state.PlayerGuildMemberRef {
		ap.updatedReferencePaths[int(playerGuildMemberRef.ID)] = playerGuildMemberRef.path
	}
	for _, playerTargetRef := range state.PlayerTargetRef {
		ap.updatedReferencePaths[int(playerTargetRef.ID)] = playerTargetRef.path
	}
	for _, playerTargetedByRef := range state.PlayerTargetedByRef {
		ap.updatedReferencePaths[int(playerTargetedByRef.ID)] = playerTargetedByRef.path
	}

	// merge paths into one map, for convencience (they are recycled anyway)
	for id, p := range ap.updatedElementPaths {
		ap.updatedPaths[id] = p
	}
	for id, p := range ap.updatedReferencePaths {
		ap.updatedPaths[id] = p
	}
}
