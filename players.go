package main

import (
	"github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs"
	"github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs/common"
)

// take in game state to extract player info and record player info into tick
func test_players(gs demoinfocs.GameState, tick *Tick) *Tick {
	players := gs.Participants().TeamMembers(common.TeamCounterTerrorists)
	for _, p := range players {
		tick.Players[p.SteamID64] = Player_info{
			Steam_id:  p.SteamID64,
			Name:      p.Name,
			Inventory: make(map[int]*common.Equipment),
			Money:     p.Money(),
			Health:    p.Health(),
			Armor:     p.Armor(),
			IsAlive:   p.IsAlive(),
			Team:      p.Team,
			Entity_id: p.EntityID,
			Defusing:  p.IsDefusing,
			HasHel:    p.HasHelmet(),
			Position: Position_Player{
				X: p.Position().X,
				Y: p.Position().Y,
				Z: p.Position().Z,
			},
		}

		for i, item := range p.Weapons() {
			tick.Players[p.SteamID64].Inventory[i] = item
		}
	}
	return tick
}
