package state

import (
	"math/rand"
	"testing"
)

func BenchmarkStateMachine(b *testing.B) {
	sm := newStateMachine()
	zone := sm.CreateZone()
	for i := 0; i < 10; i++ {
		player := zone.AddPlayer(sm)
		for i := 0; i < 10; i++ {
			player.AddItem(sm)
		}
		zone.AddItem(sm)
		zone.AddTags(sm, "string1", "string2")
	}

	for i := 0; i < b.N; i++ {
		players := zone.Players(sm)
		zone.RemovePlayers(sm, players[rand.Intn(len(players))].ID(sm))
		for _, player := range players {
			playerGearScore := player.GearScore(sm)
			playerGearScore.SetLevel(sm, playerGearScore.Level(sm)+1)
			items := player.Items(sm)
			player.RemoveItems(sm, items[rand.Intn(len(items))].ID(sm))
			player.AddItem(sm)
		}
		zone.AddPlayer(sm).AddItem(sm)
		items := zone.Items(sm)
		zone.RemoveItems(sm, items[rand.Intn(len(items))].ID(sm))
		zoneItems := zone.Items(sm)
		for _, zoneItem := range zoneItems {
			zoneItemItemGearScore := zoneItem.Item(sm).GearScore(sm)
			zoneItemItemGearScore.SetLevel(sm, zoneItemItemGearScore.Level(sm)+1)
		}
		zone.AddItem(sm)
		sm.UpdateState()
	}
}
