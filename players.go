package main

import "github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs"

// take in game state to extract player info
func test_players(gs *demoinfocs.GameState) {
	players := gs.Participants().Players()
	for _, p := range players {
		log.Printf("Player: %s, SteamID: %d", p.Name, p.SteamID64
	}
}
