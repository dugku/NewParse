package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs"
	"github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs/msg"
)

func parser_start(path string, m *Match) error {
	//initalize tick struck
	var counter int
	var cur_tick int
	//This is going to use the demoinfo-cs library
	f, _ := os.Open(path)
	defer f.Close()

	config := demoinfocs.ParserConfig{
		IgnorePacketEntitiesPanic: true,
	}

	p := demoinfocs.NewParserWithConfig(f, config)

	defer func() {
		if r := recover(); r == nil {
			log.Printf("Recovered from panic in %s: %v (frame=%d, ingameTick=%d)",
				path, r, p.CurrentFrame(), p.GameState().IngameTick())
			fmt.Printf("r: %v\n", r)
		}
	}()

	p.RegisterNetMessageHandler(func(m *msg.CSVCMsg_ServerInfo) {
		log.Printf("ServerInfo — map=%s tickInterval=%.6f maxClients=%d",
			m.GetMapName(), m.GetTickInterval(), int(m.GetMaxClients()))
	})

	//need a way to store lots of ticks huge slie..?
	//ticks_data := make([]Tick, 0)

	player_get(p, m)
	kill_logic(p, m)
	round_start_end(p, m, &counter)
	players_hurting(p, m, &cur_tick)
	weapons_firing(p, m, &cur_tick)
	bomb_handeler(p, m, &cur_tick)
	nades(m, p, &cur_tick)
	nade_handler(p, m, &cur_tick)
	//nade_handler(p, m, &cur_tick)
	for {
		more, err := p.ParseNextFrame()
		if err != nil && !errors.Is(err, io.EOF) && !errors.Is(err, demoinfocs.ErrUnexpectedEndOfDemo) {
			if strings.Contains(err.Error(), "packet entities") {
				log.Printf("Skipping problematic frame: %v", err)
				continue
			}
			return fmt.Errorf("parse %s: %w", path, err)
		}

		if !more || errors.Is(err, io.EOF) {
			break
		}
		gs := p.GameState()
		if m.openRound {
			cur_tick = gs.IngameTick()
			tick_current := Tick{
				Tick_number: cur_tick,
				Time_in_sec: 0,

				Players: make(map[uint64]*Player_info, 10),
			}
			if gs != nil {

				test_players(gs, &tick_current)
				if m == nil {
					log.Println("Sink is nil not good")
				} else {
					m.SeeFrame(tick_current)
				}

			}
		}
	}

	return nil
}

func frame(m *Match) {
	fmt.Println("Here")
	rounds := m.Rounds

	fmt.Println(rounds)
}

func main() {
	m := Match{
		Rounds: make([]RoundInfo, 0),

		// FIX: Remove the '&'. Initialize as a struct value.
		CurrentRound: &RoundInfo{
			Kills: make(map[int]RoundKill),
		},

		Players: map[uint64]PlayerStats{},
	}

	err := parser_start("vitality-vs-the-mongolz-m1-mirage.dem", &m)
	fmt.Println(err)

	// Assuming frame takes a pointer
	frame(&m)
}
