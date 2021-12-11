package state

import (
	"testing"
)

const benchTestNumberOfZones = 3

var benchTestNumberOfZonesCounter = 0

const benchTestNumberOfPlayers = 20

var benchTestNumberOfPlayersCounter = 0

const benchTestNumberOfPlayerItems = 12

var benchTestNumberOfPlayerItemsCounter = 0

const benchTestNumberOfEquipmentSetItems = 10

var benchTestNumberOfEquipmentSetItemsCounter = 0

const benchTestNumberOfEquipmentSets = 10

var benchTestNumberOfEquipmentSetsCounter = 0

const benchTestNumberOfZoneTags = 16

var benchTestNumberOfZoneTagsCounter = 0

const benchTestNumberOfInteractables = 24

var benchTestNumberOfInteractablesCounter = 0

const benchTestNumberOfZoneItems = 14

var benchTestNumberOfZoneItemsCounter = 0

func nextNum(_const int, _var *int) int {
	if _const == *_var {
		*_var = 0
		return 0
	}
	x := *_var
	*_var++
	return x
}

func setUpRealisticZoneForBenchmarkExample(engine *Engine) {
	zone := engine.CreateZone()

	// create interactables
	for i := 0; i < benchTestNumberOfInteractables; i++ {
		if i%3 == 0 {
			zone.AddInteractableItem()
		} else if i%3 == 1 {
			zone.AddInteractablePlayer()
		} else if i%3 == 2 {
			zone.AddInteractableZoneItem()
		}
	}

	// create zone players
	for i := 0; i < benchTestNumberOfPlayers; i++ {
		player := zone.AddPlayer()
		for j := 0; j < benchTestNumberOfPlayerItems; j++ {
			player.AddItem()
		}
	}

	// create ZoneItems
	for i := 0; i < benchTestNumberOfZoneItems; i++ {
		zoneItem := zone.AddItem()
		// target random players
		for j := 0; j < int(benchTestNumberOfZoneItems/2); j++ {
			randomPlayer := zone.Players()[nextNum(benchTestNumberOfPlayers, &benchTestNumberOfPlayersCounter)]
			randomPlayer.AddTargetedByZoneItem(zoneItem.ID())
		}
	}

	// make players target something
	for i, player := range zone.Players() {
		if i%2 == 0 {
			// make player target random player
			player.SetTargetPlayer(zone.Players()[nextNum(benchTestNumberOfPlayers, &benchTestNumberOfPlayersCounter)].ID())
		} else {
			// make player target random zoneItem
			player.SetTargetZoneItem(zone.Items()[nextNum(benchTestNumberOfZoneItems, &benchTestNumberOfZoneItemsCounter)].ID())
		}
	}

	// add tags to zone
	for i := 0; i < benchTestNumberOfZoneTags; i++ {
		zone.AddTags("foo")
	}

	// create equipmentSets
	for i := 0; i < benchTestNumberOfEquipmentSets; i++ {
		equipmentSet := engine.CreateEquipmentSet()
		for j := 0; j < benchTestNumberOfEquipmentSetItems; j++ {
			// create items for equipmentSet
			item := engine.CreateItem()
			equipmentSet.AddEquipment(item.ID())
		}
	}

	zonePlayers := zone.Players()
	for _, player := range zonePlayers {
		for j, _player := range zonePlayers {
			if j%2 == 0 {
				// add guild members to zone players
				player.AddGuildMember(_player.ID())
			}
		}
		// pick random equipment set
		randomEquipmentSet := engine.EveryEquipmentSet()[nextNum(benchTestNumberOfEquipmentSets, &benchTestNumberOfEquipmentSetsCounter)]
		// bind all items in this equipment set to this player
		for _, itemRef := range randomEquipmentSet.Equipment() {
			itemRef.Get().SetBoundTo(player.ID())
		}
		// add equipment set to this player
		player.AddEquipmentSet(randomEquipmentSet.ID())
	}
}

func benchTestModifyPlayerPosition(engine *Engine) {
	zone := engine.EveryZone()[nextNum(benchTestNumberOfZones, &benchTestNumberOfZonesCounter)]
	for i, _player := range zone.Players() {
		if i%3 == 0 {
			continue
		}
		// get the player with getter method
		player := engine.Player(_player.ID())
		// set new position
		player.Position().SetX(player.Position().X() + 1)
	}
}

func benchTestAddNewPlayersAsGuildMembers(engine *Engine, zone Zone) {
	for i := 0; i < int(benchTestNumberOfPlayers/3); i++ {
		// get the player with getter method
		player := engine.Player(zone.Players()[nextNum(benchTestNumberOfPlayers, &benchTestNumberOfPlayersCounter)].ID())
		newPlayer := zone.AddPlayer()
		// add each other as guild member
		player.AddGuildMember(newPlayer.ID())
		newPlayer.AddGuildMember(player.ID())
	}
}

func benchTestRemovePlayers(engine *Engine, zone Zone) {
	for i := 0; i < int(benchTestNumberOfPlayers/3); i++ {
		// get the player with getter method
		player := engine.Player(zone.Players()[nextNum(benchTestNumberOfPlayers, &benchTestNumberOfPlayersCounter)].ID())
		zone.RemovePlayers(player.ID())
	}
}

func benchTestModifyItemGearScore(engine *Engine) {
	zone := engine.EveryZone()[nextNum(benchTestNumberOfZones, &benchTestNumberOfZonesCounter)]
	for i, _zoneItem := range zone.Items() {
		if i%5 == 0 {
			continue
		}
		// get the zoneItem with getter method
		zoneItem := engine.ZoneItem(_zoneItem.ID())
		zoneItem.Item().GearScore().SetLevel(zoneItem.Item().GearScore().Level() + 1)
	}
}

func benchTestAddInteractables(engine *Engine, zone Zone) {
	for i := 0; i < int(benchTestNumberOfInteractables/9)*3; i++ {
		if i%3 == 0 {
			zone.AddInteractableItem()
		} else if i%3 == 1 {
			zone.AddInteractablePlayer()
		} else if i%3 == 2 {
			zone.AddInteractableZoneItem()
		}
	}
}

func benchTestRemoveInteractables(engine *Engine, zone Zone) {
	for i := 0; i < int(benchTestNumberOfInteractables/9)*3; i++ {
		// get the interactable with getter method
		randomInteractable := zone.Interactables()[nextNum(benchTestNumberOfInteractables, &benchTestNumberOfInteractablesCounter)]
		switch randomInteractable.Kind() {
		case ElementKindPlayer:
			zone.RemoveInteractablesPlayer(randomInteractable.Player().ID())
		case ElementKindItem:
			zone.RemoveInteractablesItem(randomInteractable.Item().ID())
		case ElementKindZoneItem:
			zone.RemoveInteractablesZoneItem(randomInteractable.ZoneItem().ID())
		}
	}
}

// func BenchmarkAssembleTreeForceInclude(b *testing.B) {
// 	engine := newEngine()
// 	for i := 0; i < benchTestNumberOfZones; i++ {
// 		setUpRealisticZoneForBenchmarkExample(engine)
// 	}
// 	engine.UpdateState()

// 	randomZone1 := engine.EveryZone()[nextNum(benchTestNumberOfZones, &benchTestNumberOfZonesCounter)]
// 	benchTestAddInteractables(engine, randomZone1)
// 	benchTestRemoveInteractables(engine, randomZone1)
// 	randomZone2 := engine.EveryZone()[nextNum(benchTestNumberOfZones, &benchTestNumberOfZonesCounter)]
// 	benchTestAddNewPlayersAsGuildMembers(engine, randomZone2)
// 	benchTestRemovePlayers(engine, randomZone2)
// 	benchTestModifyPlayerPosition(engine)
// 	benchTestModifyItemGearScore(engine)

// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		_ = engine.assembleTree(true)
// 	}
// }

func BenchmarkAssembleTree(b *testing.B) {
	engine := newEngine()
	for i := 0; i < benchTestNumberOfZones; i++ {
		setUpRealisticZoneForBenchmarkExample(engine)
	}
	engine.UpdateState()

	randomZone1 := engine.EveryZone()[nextNum(benchTestNumberOfZones, &benchTestNumberOfZonesCounter)]
	benchTestAddInteractables(engine, randomZone1)
	benchTestRemoveInteractables(engine, randomZone1)
	randomZone2 := engine.EveryZone()[nextNum(benchTestNumberOfZones, &benchTestNumberOfZonesCounter)]
	benchTestAddNewPlayersAsGuildMembers(engine, randomZone2)
	benchTestRemovePlayers(engine, randomZone2)
	benchTestModifyPlayerPosition(engine)
	benchTestModifyItemGearScore(engine)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = engine.assembleUpdateTree()
	}
}

// func BenchmarkEngine(b *testing.B) {
// 	engine := newEngine()
// 	for i := 0; i < benchTestNumberOfZones; i++ {
// 		setUpRealisticZoneForBenchmarkExample(engine)
// 	}
// 	engine.UpdateState()

// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		randomZone1 := engine.EveryZone()[nextNum(benchTestNumberOfZones, &benchTestNumberOfZonesCounter)]
// 		benchTestAddInteractables(engine, randomZone1)
// 		benchTestRemoveInteractables(engine, randomZone1)
// 		randomZone2 := engine.EveryZone()[nextNum(benchTestNumberOfZones, &benchTestNumberOfZonesCounter)]
// 		benchTestAddNewPlayersAsGuildMembers(engine, randomZone2)
// 		benchTestRemovePlayers(engine, randomZone2)
// 		benchTestModifyPlayerPosition(engine)
// 		benchTestModifyItemGearScore(engine)
// 		_ = engine.assembleUpdateTree()
// 		engine.UpdateState()
// 	}
// }

// func BenchmarkUpdateState(b *testing.B) {
// 	engine := newEngine()
// 	for i := 0; i < benchTestNumberOfZones; i++ {
// 		setUpRealisticZoneForBenchmarkExample(engine)
// 	}
// 	engine.UpdateState()

// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		randomZone1 := engine.EveryZone()[nextNum(benchTestNumberOfZones, &benchTestNumberOfZonesCounter)]
// 		benchTestAddInteractables(engine, randomZone1)
// 		benchTestRemoveInteractables(engine, randomZone1)
// 		randomZone2 := engine.EveryZone()[nextNum(benchTestNumberOfZones, &benchTestNumberOfZonesCounter)]
// 		benchTestAddNewPlayersAsGuildMembers(engine, randomZone2)
// 		benchTestRemovePlayers(engine, randomZone2)
// 		benchTestModifyPlayerPosition(engine)
// 		benchTestModifyItemGearScore(engine)
// 		engine.UpdateState()
// 	}
// }
