package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs"
	"github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs/msg"
)

func parser_start(path string, m *Match) error {
	//initalize tick struck
	var round_open bool
	var counter int
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
		}
	}()

	p.RegisterNetMessageHandler(func(m *msg.CSVCMsg_ServerInfo) {
		log.Printf("ServerInfo — map=%s tickInterval=%.6f maxClients=%d",
			m.GetMapName(), m.GetTickInterval(), int(m.GetMaxClients()))
	})

	//need a way to store lots of ticks huge slie..?
	//ticks_data := make([]Tick, 0)

	round_start_end(p, &round_open, m, &counter)
	for {
		more, err := p.ParseNextFrame()
		if err != nil && !errors.Is(err, io.EOF) && !errors.Is(err, demoinfocs.ErrUnexpectedEndOfDemo) {
			return fmt.Errorf("parse %s: %w", path, err)
		}

		if !more || errors.Is(err, io.EOF) {
			break
		}
		if round_open {
			tick_current := Tick{
				Tick_number: p.GameState().IngameTick(),
				Time_in_sec: 0,

				Players:     make(map[uint64]*Player_info, 10),
				Nades:       make([]Nades, 0),
				PlayersHurt: make([]PlayerHurt, 0),
				WeaponFired: make([]Weaps_fired, 0),
			}
			gs := p.GameState()
			if gs != nil {
				test_players(gs, &tick_current)
				nades(gs, &tick_current, p)
				players_hurting(p, &tick_current)
				weapons_firing(p, &tick_current)
				if m == nil {
					log.Println("Sink is nil not good")
				} else {
					m.SeeFrame(tick_current)
				}

			}
		}
		if counter == 2 {
			break
		}
	}

	return nil
}

func frame(m *Match) {
	fmt.Println("Here")
	rounds := m.Rounds

	for _, i := range rounds {
		for _, j := range i.Ticks {
			fmt.Println(j.PlayersHurt)
		}

	}
}

func main() {
	var m Match
	m = Match{
		Rounds:       make([]RoundInfo, 0),
		CurrentRound: &RoundInfo{},
	}
	parser_start("spirit-vs-the-mongolz-m1-dust2.dem", &m)

	frame(&m)

}
