package state

import (
	"math/rand"
	"testing"
)

func BenchmarkEngine(b *testing.B) {
	se := newEngine()
	zone := se.CreateZone()
	for i := 0; i < 10; i++ {
		player := zone.AddPlayer(se)
		for i := 0; i < 10; i++ {
			player.AddItem(se)
		}
		zone.AddItem(se)
		zone.AddTags(se, "string1", "string2")
	}

	for i := 0; i < b.N; i++ {
		players := zone.Players(se)
		zone.RemovePlayers(se, players[rand.Intn(len(players))].ID(se))
		for _, player := range players {
			playerGearScore := player.GearScore(se)
			playerGearScore.SetLevel(se, playerGearScore.Level(se)+1)
			items := player.Items(se)
			player.RemoveItems(se, items[rand.Intn(len(items))].ID(se))
			player.AddItem(se)
		}
		zone.AddPlayer(se).AddItem(se)
		items := zone.Items(se)
		zone.RemoveItems(se, items[rand.Intn(len(items))].ID(se))
		zoneItems := zone.Items(se)
		for _, zoneItem := range zoneItems {
			zoneItemItemGearScore := zoneItem.Item(se).GearScore(se)
			zoneItemItemGearScore.SetLevel(se, zoneItemItemGearScore.Level(se)+1)
		}
		zone.AddItem(se)
		se.UpdateState()
	}
}
