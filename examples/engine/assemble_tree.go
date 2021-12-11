package state

import (
	"sync"
)

type assembleJob struct {
	id    int
	kind  ElementKind
	paths []path
}

func (engine *Engine) assembleUpdateTree() Tree {

	engine.clearAssembler()
	engine.populateAssembler()

	engine.clearTree()
	var wg sync.WaitGroup
	wg.Add(7)

	go engine.executeEquipmentSetsAssembling(&wg)
	go engine.executeGearScoresAssembling(&wg)
	go engine.executeItemsAssembling(&wg)
	go engine.executePlayersAssembling(&wg)
	go engine.executePositionsAssembling(&wg)
	go engine.executeZonesAssembling(&wg)
	go engine.executeZoneItemsAssembling(&wg)

	wg.Wait()

	return engine.Tree
}

func (engine *Engine) catchAssembledEquipmentSets(count int, done chan struct{}) {
	equipmentSets := make([]equipmentSet, 0, count)

	for i := 0; i < count; i++ {
		equipmentSet := <-engine.assembler.equipmentSetChan
		equipmentSets = append(equipmentSets, equipmentSet)
	}

	for _, equipmentSet := range equipmentSets {
		engine.Tree.EquipmentSet[equipmentSet.ID] = equipmentSet
	}

	done <- struct{}{}
}

func (engine *Engine) executeEquipmentSetsAssembling(wg *sync.WaitGroup) {
	jobs := make(map[int]assembleJob)

	for _, p := range engine.assembler.equipmentSetPath {
		neutralID := p[0].id
		id := EquipmentSetID(neutralID)

		_, ok := engine.Tree.EquipmentSet[id]
		if !ok {
			engine.Tree.EquipmentSet[id] = equipmentSet{ID: id}
		}

		job := jobs[neutralID]
		job.id = neutralID
		job.kind = ElementKindEquipmentSet
		job.paths = append(job.paths, p)
		jobs[neutralID] = job
	}

	done := make(chan struct{}, 1)

	go engine.catchAssembledEquipmentSets(len(jobs), done)

	for _, job := range jobs {
		engine.assembleJobChan <- job
	}

	<-done

	wg.Done()
}

func (engine *Engine) catchAssembledGearScores(count int, done chan struct{}) {
	gearScores := make([]gearScore, 0, count)

	for i := 0; i < count; i++ {
		gearScore := <-engine.assembler.gearScoreChan
		gearScores = append(gearScores, gearScore)
	}

	for _, gearScore := range gearScores {
		engine.Tree.GearScore[gearScore.ID] = gearScore
	}

	done <- struct{}{}
}

func (engine *Engine) executeGearScoresAssembling(wg *sync.WaitGroup) {
	jobs := make(map[int]assembleJob)

	for _, p := range engine.assembler.gearScorePath {
		neutralID := p[0].id
		id := GearScoreID(neutralID)

		_, ok := engine.Tree.GearScore[id]
		if !ok {
			engine.Tree.GearScore[id] = gearScore{ID: id}
		}

		job := jobs[neutralID]
		job.id = neutralID
		job.kind = ElementKindGearScore
		job.paths = append(job.paths, p)
		jobs[neutralID] = job
	}

	done := make(chan struct{}, 1)

	go engine.catchAssembledGearScores(len(jobs), done)

	for _, job := range jobs {
		engine.assembleJobChan <- job
	}

	<-done

	wg.Done()
}

func (engine *Engine) catchAssembledItems(count int, done chan struct{}) {
	items := make([]item, 0, count)

	for i := 0; i < count; i++ {
		item := <-engine.assembler.itemChan
		items = append(items, item)
	}

	for _, item := range items {
		engine.Tree.Item[item.ID] = item
	}

	done <- struct{}{}
}

func (engine *Engine) executeItemsAssembling(wg *sync.WaitGroup) {
	jobs := make(map[int]assembleJob)

	for _, p := range engine.assembler.itemPath {
		neutralID := p[0].id
		id := ItemID(neutralID)

		_, ok := engine.Tree.Item[id]
		if !ok {
			engine.Tree.Item[id] = item{ID: id}
		}

		job := jobs[neutralID]
		job.id = neutralID
		job.kind = ElementKindItem
		job.paths = append(job.paths, p)
		jobs[neutralID] = job
	}

	done := make(chan struct{}, 1)

	go engine.catchAssembledItems(len(jobs), done)

	for _, job := range jobs {
		engine.assembleJobChan <- job
	}

	<-done

	wg.Done()
}

func (engine *Engine) catchAssembledPlayers(count int, done chan struct{}) {
	players := make([]player, 0, count)

	for i := 0; i < count; i++ {
		player := <-engine.assembler.playerChan
		players = append(players, player)
	}

	for _, player := range players {
		engine.Tree.Player[player.ID] = player
	}

	done <- struct{}{}
}

func (engine *Engine) executePlayersAssembling(wg *sync.WaitGroup) {
	jobs := make(map[int]assembleJob)

	for _, p := range engine.assembler.playerPath {
		neutralID := p[0].id
		id := PlayerID(neutralID)

		_, ok := engine.Tree.Player[id]
		if !ok {
			engine.Tree.Player[id] = player{ID: id}
		}

		job := jobs[neutralID]
		job.id = neutralID
		job.kind = ElementKindPlayer
		job.paths = append(job.paths, p)
		jobs[neutralID] = job
	}

	done := make(chan struct{}, 1)

	go engine.catchAssembledPlayers(len(jobs), done)

	for _, job := range jobs {
		engine.assembleJobChan <- job
	}

	<-done

	wg.Done()
}

func (engine *Engine) catchAssembledPositions(count int, done chan struct{}) {
	positions := make([]position, 0, count)

	for i := 0; i < count; i++ {
		position := <-engine.assembler.positionChan
		positions = append(positions, position)
	}

	for _, position := range positions {
		engine.Tree.Position[position.ID] = position
	}

	done <- struct{}{}
}

func (engine *Engine) executePositionsAssembling(wg *sync.WaitGroup) {
	jobs := make(map[int]assembleJob)

	for _, p := range engine.assembler.positionPath {
		neutralID := p[0].id
		id := PositionID(neutralID)

		_, ok := engine.Tree.Position[id]
		if !ok {
			engine.Tree.Position[id] = position{ID: id}
		}

		job := jobs[neutralID]
		job.id = neutralID
		job.kind = ElementKindPosition
		job.paths = append(job.paths, p)
		jobs[neutralID] = job
	}

	done := make(chan struct{}, 1)

	go engine.catchAssembledPositions(len(jobs), done)

	for _, job := range jobs {
		engine.assembleJobChan <- job
	}

	<-done

	wg.Done()
}

func (engine *Engine) catchAssembledZones(count int, done chan struct{}) {
	zones := make([]zone, 0, count)

	for i := 0; i < count; i++ {
		zone := <-engine.assembler.zoneChan
		zones = append(zones, zone)
	}

	for _, zone := range zones {
		engine.Tree.Zone[zone.ID] = zone
	}

	done <- struct{}{}
}

func (engine *Engine) executeZonesAssembling(wg *sync.WaitGroup) {
	jobs := make(map[int]assembleJob)

	for _, p := range engine.assembler.zonePath {
		neutralID := p[0].id
		id := ZoneID(neutralID)

		_, ok := engine.Tree.Zone[id]
		if !ok {
			engine.Tree.Zone[id] = zone{ID: id}
		}

		job := jobs[neutralID]
		job.id = neutralID
		job.kind = ElementKindZone
		job.paths = append(job.paths, p)
		jobs[neutralID] = job
	}

	done := make(chan struct{}, 1)

	go engine.catchAssembledZones(len(jobs), done)

	for _, job := range jobs {
		engine.assembleJobChan <- job
	}

	<-done

	wg.Done()
}

func (engine *Engine) catchAssembledZoneItems(count int, done chan struct{}) {
	zoneItems := make([]zoneItem, 0, count)

	for i := 0; i < count; i++ {
		zoneItem := <-engine.assembler.zoneItemChan
		zoneItems = append(zoneItems, zoneItem)
	}

	for _, zoneItem := range zoneItems {
		engine.Tree.ZoneItem[zoneItem.ID] = zoneItem
	}

	done <- struct{}{}
}

func (engine *Engine) executeZoneItemsAssembling(wg *sync.WaitGroup) {
	jobs := make(map[int]assembleJob)

	for _, p := range engine.assembler.zoneItemPath {
		neutralID := p[0].id
		id := ZoneItemID(neutralID)

		_, ok := engine.Tree.ZoneItem[id]
		if !ok {
			engine.Tree.ZoneItem[id] = zoneItem{ID: id}
		}

		job := jobs[neutralID]
		job.id = neutralID
		job.kind = ElementKindZoneItem
		job.paths = append(job.paths, p)
		jobs[neutralID] = job
	}

	done := make(chan struct{}, 1)

	go engine.catchAssembledZoneItems(len(jobs), done)

	for _, job := range jobs {
		engine.assembleJobChan <- job
	}

	<-done

	wg.Done()
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
	// we want to find all lead nodes which have updated and collect their paths
	// later we will basically loop over the paths we have collected, and "walk" them
	// in order to assemble the tree from bottom to the top (leaf nodes)
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
	// we'd be pretty much done collecting the required paths, but we also want to
	// build all paths ending with a reference which references an updated element.
	// this needs to happen recursively, consider this example (-> = reference):
	// A -> B -> C -> ^D
	// Since "D" has updated (^), but no other element, we'd only include "C".
	// However, now that "C" is considered updated by reference, we also want
	// to include "B". This is why recursiveness is required.
	for {
		// here we populate out includedElements with out newly collected paths segments
		// so we can check if any of these  elements are referenced by any reference
		// in the loop below
		for _, p := range engine.assembler.updatedElementPaths {
			for _, seg := range p {
				engine.assembler.includedElements[seg.id] = true
			}
		}
		// add all elements of the updated reference paths to the includedElements
		for _, p := range engine.assembler.updatedReferencePaths {
			for _, seg := range p {
				// we know that the last segment of a reference path has a reference ID
				// if i == len(p)-1 {
				// 	engine.assembler.includedElements[seg.refID] = true
				// } else {
				// 	engine.assembler.includedElements[seg.id] = true
				// }
				if seg.refID != 0 {
					engine.assembler.includedElements[seg.refID] = true
				} else {
					engine.assembler.includedElements[seg.id] = true
				}
			}
		}

		// we check if ant new elements are involved, which could
		// mean that new paths containing references need to be looked at
		if prevousLength == len(engine.assembler.includedElements) {
			break
		}

		prevousLength = len(engine.assembler.includedElements)

		for _, equipmentSetEquipmentRef := range engine.Patch.EquipmentSetEquipmentRef {
			// if the reference references an element that has updated its path is collected
			// so that all segments can later be added to includedElements
			if _, ok := engine.assembler.includedElements[int(equipmentSetEquipmentRef.ReferencedElementID)]; ok {
				engine.assembler.updatedReferencePaths[int(equipmentSetEquipmentRef.ID)] = equipmentSetEquipmentRef.path
			}
		}
		for _, equipmentSetEquipmentRef := range engine.State.EquipmentSetEquipmentRef {
			// prioritize the ref from Patch
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

	// merge paths into one map, for convencience (they are recycled anyway)
	for id, path := range engine.assembler.updatedElementPaths {
		engine.assembler.updatedPaths[id] = path
	}
	for id, path := range engine.assembler.updatedReferencePaths {
		engine.assembler.updatedPaths[id] = path
	}

	// just to be a bit more organized
	for leafElementID, p := range engine.assembler.updatedPaths {
		switch p[0].identifier {
		case equipmentSetIdentifier:
			engine.assembler.equipmentSetPath[leafElementID] = p
		case gearScoreIdentifier:
			engine.assembler.gearScorePath[leafElementID] = p
		case itemIdentifier:
			engine.assembler.itemPath[leafElementID] = p
		case playerIdentifier:
			engine.assembler.playerPath[leafElementID] = p
		case positionIdentifier:
			engine.assembler.positionPath[leafElementID] = p
		case zoneIdentifier:
			engine.assembler.zonePath[leafElementID] = p
		case zoneItemIdentifier:
			engine.assembler.zoneItemPath[leafElementID] = p
		}
	}
}
