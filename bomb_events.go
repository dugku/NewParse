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
		} else {
			start.PlayerStartedName = "Unknown"
		}

		if len(m.CurrentRound.Ticks) > 0 {
			current := &m.CurrentRound.Ticks[len(m.CurrentRound.Ticks)-1]
			if current.Tick_number == *cTick {
				current.DefuseStart = append(current.DefuseStart, start)
			} else {
				// Create a new tick if needed, or add to the most recent one
				current.DefuseStart = append(current.DefuseStart, start)
			}
		}
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
		} else {
			abort.PayerAbortedName = "Unknown"
		}

		abort.TickNum = p.GameState().IngameTick()

		if len(m.CurrentRound.Ticks) > 0 {
			current := &m.CurrentRound.Ticks[len(m.CurrentRound.Ticks)-1]
			if current.Tick_number == *cTick {
				current.DefuseAbort = append(current.DefuseAbort, abort)
			} else {
				current.DefuseAbort = append(current.DefuseAbort, abort)
			}
		}
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
		} else {
			dropped.PlayerDropName = "Uknown"
		}

		if len(m.CurrentRound.Ticks) > 0 {
			current := &m.CurrentRound.Ticks[len(m.CurrentRound.Ticks)-1]
			if current.Tick_number == *cTick {
				current.BombaDropped = append(current.BombaDropped, dropped)
			}
		}
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
		} else {
			picked.PlayerPickUpName = "Uknown"
		}

		if len(m.CurrentRound.Ticks) > 0 {
			current := &m.CurrentRound.Ticks[len(m.CurrentRound.Ticks)-1]
			if current.Tick_number == *cTick {
				current.BombPickUp = append(current.BombPickUp, picked)
			}
		}

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

		if len(m.CurrentRound.Ticks) > 0 {
			current := &m.CurrentRound.Ticks[len(m.CurrentRound.Ticks)-1]
			if current.Tick_number == *cTick {
				current.Planted = append(current.Planted, planted)
			} else {
				current.Planted = append(current.Planted, planted)
			}
		}
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
		} else {
			plantAbort.PlayerAbortPlantName = "Unknown"
		}

		if len(m.CurrentRound.Ticks) > 0 {
			current := &m.CurrentRound.Ticks[len(m.CurrentRound.Ticks)-1]
			if current.Tick_number == *cTick {
				current.PlantAborted = append(current.PlantAborted, plantAbort)
			} else {
				current.PlantAborted = append(current.PlantAborted, plantAbort)
			}
		}
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
		} else {
			plantBegin.PlayerBeginPlantName = "Unknown"
		}

		gs := p.GameState()
		if gs != nil {
			plantBegin.TickNum = gs.IngameTick()
		}

		if len(m.CurrentRound.Ticks) > 0 {
			current := &m.CurrentRound.Ticks[len(m.CurrentRound.Ticks)-1]
			if current.Tick_number == *cTick {
				current.PlantBegin = append(current.PlantBegin, plantBegin)
			} else {
				current.PlantBegin = append(current.PlantBegin, plantBegin)
			}
		}
	})

}
