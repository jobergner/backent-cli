package state

import (
	"math/rand"
	"testing"
)

const benchTestNumberOfZones = 10
const benchTestNumberOfPlayers = 100
const benchTestNumberOfPlayerItems = 20
const benchTestNumberOfEquipmentSetItems = 10
const benchTestNumberOfEquipmentSets = 30
const benchTestNumberOfZoneTags = 40
const benchTestNumberOfInteractables = 80
const benchTestNumberOfZoneItems = 80

func setUpRealisticZoneForBenchmarkExample(engine *Engine) {
	zone := engine.CreateZone()

	// create interactables
	for i := 0; i < benchTestNumberOfInteractables; i++ {
		if n := rand.Intn(3); n == 0 {
			zone.AddInteractableItem()
		} else if n == 1 {
			zone.AddInteractablePlayer()
		} else if n == 2 {
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
	for _, player := range zone.Players() {
		if n := rand.Intn(2); n == 0 {
			// make player target random player
			player.SetTargetPlayer(zone.Players()[rand.Intn(benchTestNumberOfPlayers)].ID())
		} else if n == 1 {
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
		for _, _player := range zonePlayers {
			if rand.Intn(2) == 1 {
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

func BenchmarkEngine(b *testing.B) {
	engine := newEngine()
	for i := 0; i < benchTestNumberOfZones; i++ {
		// clutter engine with multiple zones for real life conditions
		setUpRealisticZoneForBenchmarkExample(engine)
	}

	for i := 0; i < b.N; i++ {
	}
}
