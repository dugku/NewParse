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
	var counter int
	var cur_tick int

	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("open demo %s: %w", path, err)
	}
	defer f.Close()

	config := demoinfocs.ParserConfig{
		IgnorePacketEntitiesPanic: true,
	}

	p := demoinfocs.NewParserWithConfig(f, config)

	// Fixed: r != nil so we only log on a real panic, not on normal exit
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered from panic in %s: %v (frame=%d, ingameTick=%d)",
				path, r, p.CurrentFrame(), p.GameState().IngameTick())
		}
	}()

	// Capture map name from server info
	p.RegisterNetMessageHandler(func(info *msg.CSVCMsg_ServerInfo) {
		m.MapName = info.GetMapName()
		log.Printf("ServerInfo — map=%s tickInterval=%.6f maxClients=%d",
			info.GetMapName(), info.GetTickInterval(), int(info.GetMaxClients()))
	})

	player_get(p, m)
	kill_logic(p, m)
	round_start_end(p, m, &counter)
	players_hurting(p, m, &cur_tick)
	weapons_firing(p, m, &cur_tick)
	bomb_handeler(p, m, &cur_tick)
	nades(m, p, &cur_tick)
	nade_handler(p, m, &cur_tick)

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
				Players:     make(map[uint64]*Player_info, 10),
			}
			if gs != nil {
				test_players(gs, &tick_current)
				if m != nil {
					m.SeeFrame(tick_current)
				}
			}
		}
	}

	return nil
}

func main() {
	demoPath := "parivision-vs-g2-m3-ancient.dem"

	m := Match{
		Rounds: make([]RoundInfo, 0),
		CurrentRound: &RoundInfo{
			Kills: make(map[int]RoundKill),
		},
		Players: map[uint64]PlayerStats{},
	}

	err := parser_start(demoPath, &m)
	if err != nil {
		log.Printf("Parser error: %v", err)
	}

	log.Printf("Parsed %d rounds", len(m.Rounds))

	if err := m.WriteMatch(demoPath); err != nil {
		log.Fatalf("Failed to write match JSON: %v", err)
	}
}
