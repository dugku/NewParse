package main

import (
	"log"

	"github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs"
	"github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs/common"
)

// take in game state to extract player info
func test_players(gs demoinfocs.GameState) {
	players := gs.Participants().TeamMembers(common.TeamCounterTerrorists)
	for _, p := range players {
		log.Printf("Player: %s, SteamID: %d", p.Name, p.SteamID64)
	}
}
