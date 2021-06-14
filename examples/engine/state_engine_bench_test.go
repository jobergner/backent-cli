package state

import (
	"math/rand"
	"testing"
)

func BenchmarkEngine(b *testing.B) {
	se := newEngine()
	zone := se.CreateZone()
	for i := 0; i < 10; i++ {
		player := zone.AddPlayer()
		for i := 0; i < 10; i++ {
			player.AddItem()
		}
		zone.AddItem()
		zone.AddTags("string1", "string2")
	}

	for i := 0; i < b.N; i++ {
		players := zone.Players()
		zone.RemovePlayers(players[rand.Intn(len(players))].ID())
		for _, player := range zone.Players() {
			playerGearScore := player.GearScore()
			playerGearScore.SetLevel(playerGearScore.Level() + 1)
			items := player.Items()
			player.RemoveItems(items[rand.Intn(len(items))].ID())
			player.AddItem()
		}
		zone.AddPlayer().AddItem()
		items := zone.Items()
		zone.RemoveItems(items[rand.Intn(len(items))].ID())
		zoneItems := zone.Items()
		for _, zoneItem := range zoneItems {
			zoneItemItemGearScore := zoneItem.Item().GearScore()
			zoneItemItemGearScore.SetLevel(zoneItemItemGearScore.Level() + 1)
		}
		zone.AddItem()
		se.UpdateState()
	}
}
