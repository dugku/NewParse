package main

import (
	"errors"
	"io"
	"log"
	"os"

	"github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs"
	"github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs/msg"
)

func parser_start(path string) error {
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

	for {
		more, err := p.ParseNextFrame()
		if err != nil {
			log.Fatalf("Error parsing demo %s: %v", path, err)
			if !errors.Is(err, demoinfocs.ErrUnexpectedEndOfDemo) && !errors.Is(err, io.EOF) {
				return err
			}
		}
		if !more || errors.Is(err, io.EOF) {
			break
		}

		gs := p.GameState()
		if gs == nil {
			continue
		}

		//this is where the fun starts :)
		test_players(&gs)
	}

	return nil

}

func main() {
	parser_start("furia-vs-mouz-m2-overpass.dem")
}
