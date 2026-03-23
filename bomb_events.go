package main

import (
	"fmt"

	"github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs"
	"github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs/events"
)

func bomb_handeler(p demoinfocs.Parser, m *Match, cTick *int) {
	p.RegisterEventHandler(func(e events.BombDefuseStart) {
		start := BombDefuseStarted{}
		fmt.Println("Defuse Started")

		if e.Player != nil {
			start.PlayerStartedSteamId = e.Player.SteamID64
			start.PlayerStartedName = e.Player.Name
			start.Position = Position{
				X: e.Player.Position().X,
				Y: e.Player.Position().Y,
				Z: e.Player.Position().Z,
			}
			start.Kit = e.HasKit
			start.TimeSec = int(p.CurrentTime().Seconds())
		} else {
			start.PlayerStartedName = "Unknown"
		}

		m.CurrentRound.DefuseStart = append(m.CurrentRound.DefuseStart, start)

	})

	p.RegisterEventHandler(func(e events.BombDefuseAborted) {
		abort := BombDefuseAbort{}
		fmt.Println("Defuse Abort")

		if e.Player != nil {
			abort.PlayerAbortedSteamId = e.Player.SteamID64
			abort.PayerAbortedName = e.Player.Name
			abort.Position = Position{
				X: e.Player.Position().X,
				Y: e.Player.Position().Y,
				Z: e.Player.Position().Z,
			}
			abort.TimeSec = int(p.CurrentTime().Seconds())
		} else {
			abort.PayerAbortedName = "Unknown"
		}

		abort.TickNum = p.GameState().IngameTick()

		m.CurrentRound.DefuseAbort = append(m.CurrentRound.DefuseAbort, abort)

	})

	p.RegisterEventHandler(func(e events.BombDropped) {
		dropped := BombDrop{}

		if e.Player != nil {
			dropped.PlayerDropSteamId = e.Player.SteamID64
			dropped.PlayerDropName = e.Player.Name
			dropped.Position = Position{
				X: e.Player.Position().X,
				Y: e.Player.Position().Y,
				Z: e.Player.Position().Z,
			}
			dropped.TimeSec = int(p.CurrentTime().Seconds())
		} else {
			dropped.PlayerDropName = "Uknown"
		}

		m.CurrentRound.BombaDropped = append(m.CurrentRound.BombaDropped, dropped)

	})

	p.RegisterEventHandler(func(e events.BombPickup) {
		picked := BombPickedUp{}

		if e.Player != nil {
			picked.PlayerPickedUpSteamId = e.Player.SteamID64
			picked.PlayerPickUpName = e.Player.Name
			picked.Position = Position{
				X: e.Player.Position().X,
				Y: e.Player.Position().Y,
				Z: e.Player.Position().Z,
			}
			picked.TimeSec = int(p.CurrentTime().Seconds())
		} else {
			picked.PlayerPickUpName = "Uknown"
		}

		m.CurrentRound.BombPickUp = append(m.CurrentRound.BombPickUp, picked)

	})

	p.RegisterEventHandler(func(e events.BombPlanted) {
		if m.CurrentRound == nil {
			return
		}

		planted := Planted{}
		fmt.Println("Bomb Planted")

		if e.Player != nil {
			planted.PlayerPlantedSteamId = e.Player.SteamID64
			planted.PlayerPlantedName = e.Player.Name
			planted.Position = Position{
				X: e.Player.Position().X,
				Y: e.Player.Position().Y,
				Z: e.Player.Position().Z,
			}
			planted.TimeSec = int(p.CurrentTime().Seconds())
		} else {
			planted.PlayerPlantedName = "Unknown"
		}

		gs := p.GameState()
		if gs != nil {
			planted.TickNum = gs.IngameTick()

			// Convert events.Bombsite to your Bombsite type
			switch e.Site {
			case events.BombsiteA:
				planted.Site = BombsiteA
			case events.BombsiteB:
				planted.Site = BombsiteB
			default:
				planted.Site = BombsiteUnknown
			}
		}

		m.CurrentRound.Planted = append(m.CurrentRound.Planted, planted)

	})

	p.RegisterEventHandler(func(e events.BombPlantAborted) {
		if m.CurrentRound == nil {
			return
		}

		plantAbort := PlantAborted{}
		fmt.Println("Bomb Plant Aborted")

		if e.Player != nil {
			plantAbort.PlayerAbortPlantSteamId = e.Player.SteamID64
			plantAbort.PlayerAbortPlantName = e.Player.Name
			plantAbort.Position = Position{
				X: e.Player.Position().X,
				Y: e.Player.Position().Y,
				Z: e.Player.Position().Z,
			}
			plantAbort.TimeSec = int(p.CurrentTime().Seconds())
		} else {
			plantAbort.PlayerAbortPlantName = "Unknown"
		}

		m.CurrentRound.PlantAborted = append(m.CurrentRound.PlantAborted, plantAbort)

	})

	p.RegisterEventHandler(func(e events.BombPlantBegin) {
		if m.CurrentRound == nil {
			return
		}

		plantBegin := PlantBegin{}
		fmt.Println("Bomb Plant Begin")

		if e.Player != nil {
			plantBegin.PlayerBeginPlantSteamId = e.Player.SteamID64
			plantBegin.PlayerBeginPlantName = e.Player.Name
			plantBegin.Position = Position{
				X: e.Player.Position().X,
				Y: e.Player.Position().Y,
				Z: e.Player.Position().Z,
			}
			plantBegin.TimeSec = int(p.CurrentTime().Seconds())
		} else {
			plantBegin.PlayerBeginPlantName = "Unknown"
		}

		gs := p.GameState()
		if gs != nil {
			plantBegin.TickNum = gs.IngameTick()
		}

		m.CurrentRound.PlantBegin = append(m.CurrentRound.PlantBegin, plantBegin)

	})

}
