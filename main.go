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

func parser_start(path string, sink RoundSink) error {

	//initalize tick struck
	var round_open bool
	//This is going to use the demoinfo-cs library
	f, _ := os.Open(path)
	defer f.Close()

	p := demoinfocs.NewParser(f)

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

	round_start_end(p, &round_open, sink)
	for {
		more, err := p.ParseNextFrame()
		if err != nil && !errors.Is(err, io.EOF) && !errors.Is(err, demoinfocs.ErrUnexpectedEndOfDemo) {
			return fmt.Errorf("parse %s: %w", path, err)
		}

		if !more || errors.Is(err, io.EOF) {
			fmt.Print("Here")
			break
		}
		if round_open {
			tick_current := Tick{
				Tick_number: p.GameState().IngameTick(),
				Time_in_sec: 0,

				Players: make(map[uint64]Player_info),
			}
			gs := p.GameState()

			if gs != nil {

				test_players(gs, &tick_current)

			}
		}
	}

	return nil

}

func main() {
	var m RoundSink
	parser_start("spirit-vs-the-mongolz-m1-dust2.dem", m)
}
