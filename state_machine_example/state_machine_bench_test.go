package statemachine

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
		players := zone.GetPlayers(sm)
		zone.RemovePlayers(sm, players[rand.Intn(len(players))].GetID(sm))
		for _, player := range players {
			playerGearScore := player.GetGearScore(sm)
			playerGearScore.SetLevel(sm, playerGearScore.GetLevel(sm)+1)
			items := player.GetItems(sm)
			player.RemoveItems(sm, items[rand.Intn(len(items))].GetID(sm))
			player.AddItem(sm)
		}
		zone.AddPlayer(sm).AddItem(sm)
		items := zone.GetItems(sm)
		zone.RemoveItems(sm, items[rand.Intn(len(items))].GetID(sm))
		zoneItems := zone.GetItems(sm)
		for _, zoneItem := range zoneItems {
			zoneItemItemGearScore := zoneItem.GetItem(sm).GetGearScore(sm)
			zoneItemItemGearScore.SetLevel(sm, zoneItemItemGearScore.GetLevel(sm)+1)
		}
		zone.AddItem(sm)
		sm.UpdateState()
	}
}
