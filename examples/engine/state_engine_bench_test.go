package state

import (
	"math/rand"
	"testing"
)

const benchTestNumberOfZones = 3
const benchTestNumberOfPlayers = 20
const benchTestNumberOfPlayerItems = 12
const benchTestNumberOfEquipmentSetItems = 10
const benchTestNumberOfEquipmentSets = 10
const benchTestNumberOfZoneTags = 16
const benchTestNumberOfInteractables = 24
const benchTestNumberOfZoneItems = 14

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
			randomPlayer := zone.Players()[rand.Intn(benchTestNumberOfPlayers)]
			randomPlayer.AddTargetedByZoneItem(zoneItem.ID())
		}
	}

	// make players target something
	for i, player := range zone.Players() {
		if i%2 == 0 {
			// make player target random player
			player.SetTargetPlayer(zone.Players()[rand.Intn(benchTestNumberOfPlayers)].ID())
		} else {
			// make player target random zoneItem
			player.SetTargetZoneItem(zone.Items()[rand.Intn(benchTestNumberOfZoneItems)].ID())
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
		randomEquipmentSet := engine.EveryEquipmentSet()[rand.Intn(benchTestNumberOfEquipmentSets)]
		// bind all items in this equipment set to this player
		for _, itemRef := range randomEquipmentSet.Equipment() {
			itemRef.Get().SetBoundTo(player.ID())
		}
		// add equipment set to this player
		player.AddEquipmentSet(randomEquipmentSet.ID())
	}
}

func benchTestModifyPlayerPosition(engine *Engine) {
	zone := engine.EveryZone()[rand.Intn(benchTestNumberOfZones)]
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

func benchTestAddNewPlayersAsGuildMembers(engine *Engine, zone zone) {
	for i := 0; i < int(benchTestNumberOfPlayers/3); i++ {
		// get the player with getter method
		player := engine.Player(zone.Players()[rand.Intn(benchTestNumberOfPlayers)].ID())
		newPlayer := zone.AddPlayer()
		// add each other as guild member
		player.AddGuildMember(newPlayer.ID())
		newPlayer.AddGuildMember(player.ID())
	}
}

func benchTestRemovePlayers(engine *Engine, zone zone) {
	for i := 0; i < int(benchTestNumberOfPlayers/3); i++ {
		// get the player with getter method
		player := engine.Player(zone.Players()[rand.Intn(benchTestNumberOfPlayers)].ID())
		zone.RemovePlayers(player.ID())
	}
}

func benchTestModifyItemGearScore(engine *Engine) {
	zone := engine.EveryZone()[rand.Intn(benchTestNumberOfZones)]
	for i, _zoneItem := range zone.Items() {
		if i%5 == 0 {
			continue
		}
		// get the zoneItem with getter method
		zoneItem := engine.ZoneItem(_zoneItem.ID())
		zoneItem.Item().GearScore().SetLevel(zoneItem.Item().GearScore().Level() + 1)
	}
}

func benchTestAddInteractables(engine *Engine, zone zone) {
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

func benchTestRemoveInteractables(engine *Engine, zone zone) {
	for i := 0; i < int(benchTestNumberOfInteractables/9)*3; i++ {
		// get the interactable with getter method
		randomInteractable := zone.Interactables()[rand.Intn(benchTestNumberOfInteractables)]
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

func BenchmarkAssembleTreeForceInclude(b *testing.B) {
	engine := newEngine()
	for i := 0; i < benchTestNumberOfZones; i++ {
		setUpRealisticZoneForBenchmarkExample(engine)
	}
	engine.UpdateState()

	randomZone1 := engine.EveryZone()[rand.Intn(benchTestNumberOfZones)]
	benchTestAddInteractables(engine, randomZone1)
	benchTestRemoveInteractables(engine, randomZone1)
	randomZone2 := engine.EveryZone()[rand.Intn(benchTestNumberOfZones)]
	benchTestAddNewPlayersAsGuildMembers(engine, randomZone2)
	benchTestRemovePlayers(engine, randomZone2)
	benchTestModifyPlayerPosition(engine)
	benchTestModifyItemGearScore(engine)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = engine.assembleTree(true)
	}
}

func BenchmarkAssembleTree(b *testing.B) {
	engine := newEngine()
	for i := 0; i < benchTestNumberOfZones; i++ {
		setUpRealisticZoneForBenchmarkExample(engine)
	}
	engine.UpdateState()

	randomZone1 := engine.EveryZone()[rand.Intn(benchTestNumberOfZones)]
	benchTestAddInteractables(engine, randomZone1)
	benchTestRemoveInteractables(engine, randomZone1)
	randomZone2 := engine.EveryZone()[rand.Intn(benchTestNumberOfZones)]
	benchTestAddNewPlayersAsGuildMembers(engine, randomZone2)
	benchTestRemovePlayers(engine, randomZone2)
	benchTestModifyPlayerPosition(engine)
	benchTestModifyItemGearScore(engine)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = engine.assembleTree(false)
	}
}

func BenchmarkEngine(b *testing.B) {
	engine := newEngine()
	for i := 0; i < benchTestNumberOfZones; i++ {
		setUpRealisticZoneForBenchmarkExample(engine)
	}
	engine.UpdateState()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		randomZone1 := engine.EveryZone()[rand.Intn(benchTestNumberOfZones)]
		benchTestAddInteractables(engine, randomZone1)
		benchTestRemoveInteractables(engine, randomZone1)
		randomZone2 := engine.EveryZone()[rand.Intn(benchTestNumberOfZones)]
		benchTestAddNewPlayersAsGuildMembers(engine, randomZone2)
		benchTestRemovePlayers(engine, randomZone2)
		benchTestModifyPlayerPosition(engine)
		benchTestModifyItemGearScore(engine)
		_ = engine.assembleTree(false)
		engine.UpdateState()
	}
}

func BenchmarkUpdateState(b *testing.B) {
	engine := newEngine()
	for i := 0; i < benchTestNumberOfZones; i++ {
		setUpRealisticZoneForBenchmarkExample(engine)
	}
	engine.UpdateState()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		randomZone1 := engine.EveryZone()[rand.Intn(benchTestNumberOfZones)]
		benchTestAddInteractables(engine, randomZone1)
		benchTestRemoveInteractables(engine, randomZone1)
		randomZone2 := engine.EveryZone()[rand.Intn(benchTestNumberOfZones)]
		benchTestAddNewPlayersAsGuildMembers(engine, randomZone2)
		benchTestRemovePlayers(engine, randomZone2)
		benchTestModifyPlayerPosition(engine)
		benchTestModifyItemGearScore(engine)
		engine.UpdateState()
	}
}

func BenchmarkElementModificaton(b *testing.B) {
	engine := newEngine()
	for i := 0; i < benchTestNumberOfZones; i++ {
		setUpRealisticZoneForBenchmarkExample(engine)
	}
	engine.UpdateState()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		randomZone1 := engine.EveryZone()[rand.Intn(benchTestNumberOfZones)]
		benchTestAddInteractables(engine, randomZone1)
		benchTestRemoveInteractables(engine, randomZone1)
		randomZone2 := engine.EveryZone()[rand.Intn(benchTestNumberOfZones)]
		benchTestAddNewPlayersAsGuildMembers(engine, randomZone2)
		benchTestRemovePlayers(engine, randomZone2)
		benchTestModifyPlayerPosition(engine)
		benchTestModifyItemGearScore(engine)
	}
}
