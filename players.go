package main

import (
	"log"

	"github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs"
	"github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs/common"
)

// take in game state to extract player info and record player info into tick
func test_players(gs demoinfocs.GameState, tick *Tick) {

	//need to see if common.Team come backj nil

	TeamCounter := gs.Team(common.TeamCounterTerrorists)
	TeamTerror := gs.Team(common.TeamTerrorists)

	check_team(TeamCounter)
	check_team(TeamTerror)

	players_counter := gs.Participants().TeamMembers(common.TeamCounterTerrorists)
	for _, p := range players_counter {
		tick.Players[p.SteamID64] = Player_info{
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
			Position: Position{
				X: p.Position().X,
				Y: p.Position().Y,
				Z: p.Position().Z,
			},
		}

		for i, item := range p.Weapons() {
			tick.Players[p.SteamID64].Inventory[i] = item
		}
	}
	players_terror := gs.Participants().TeamMembers(common.TeamTerrorists)
	for _, p := range players_terror {
		tick.Players[p.SteamID64] = Player_info{
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
			Position: Position{
				X: p.Position().X,
				Y: p.Position().Y,
				Z: p.Position().Z,
			},
		}

		for i, item := range p.Weapons() {
			tick.Players[p.SteamID64].Inventory[i] = item
		}
	}
}

func check_team(t *common.TeamState) {
	if t == nil {
		log.Println("One of the TeamStates is nil not good")
	}
}
