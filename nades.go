package main

import (
	"log"

	"github.com/golang/geo/r3"
	"github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs"
	"github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs/common"
	"github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs/events"
)

func nades(gs demoinfocs.GameState, tick *Tick, p demoinfocs.Parser) *Tick {
	gernade_entities := gs.GrenadeProjectiles()
	check_nades(gernade_entities)

	func() {

		defer func() {
			if r := recover(); r != nil {
				log.Printf("Recover of a panic in Nades %v\n", r)
			}
		}()

		for _, g := range gernade_entities {

			if g == nil {
				log.Println("g is nil..?")
			}

			n := Nades{
				EId: common.EquipmentClass(g.UniqueID()),
				Pos: Position{
					X: g.Position().X,
					Y: g.Position().Y,
					Z: g.Position().Z,
				},
				Type:      g.WeaponInstance,
				TimeInSec: p.CurrentTime(),
				//Vel: Velocityy{
				//	VX: g.Velocity().X,
				//	VY: g.Velocity().Y, for some reason this was a problem nice
				//	VZ: g.Velocity().Z,
				//},
			}

			tick.Nades = append(tick.Nades, n)
		}
	}()

	return tick
}

/*
Need to redo the below function since it's doesn't do what it's suppose to do,
don't know how to fix it yet but need to do some of the other stuff first.
*/

func nade_handler(p demoinfocs.Parser, m *Match, curTick *int) {
	p.RegisterEventHandler(func(e events.DecoyExpired) {
		if m.CurrentRound == nil {
			return
		}

		if e.Grenade == nil {
			return
		}

		nadeEvent := createNadeEvent(e.Grenade, p)
		if nadeEvent != nil {
			decoyDone := DecoyDone{Event: nadeEvent}

			if len(m.CurrentRound.Ticks) > 0 {
				current := &m.CurrentRound.Ticks[len(m.CurrentRound.Ticks)-1]
				if current.Tick_number == *curTick {
					current.DecoyDone = append(current.DecoyDone, decoyDone)
				}
			}
		}
	})

	p.RegisterEventHandler(func(e events.DecoyStart) {
		if m.CurrentRound == nil {
			return
		}

		if e.Grenade == nil {
			return
		}

		nadeEvent := createNadeEvent(e.Grenade, p)
		if nadeEvent != nil {
			decoyStarted := DecoyStarted{Event: nadeEvent}

			if len(m.CurrentRound.Ticks) > 0 {
				current := &m.CurrentRound.Ticks[len(m.CurrentRound.Ticks)-1]
				if current.Tick_number == *curTick {
					current.DecoyStarted = append(current.DecoyStarted, decoyStarted)
				}
			}
		}
	})

	p.RegisterEventHandler(func(e events.FireGrenadeStart) {
		if m.CurrentRound == nil {
			return
		}

		if e.Grenade == nil {
			return
		}

		nadeEvent := createNadeEvent(e.Grenade, p)
		if nadeEvent != nil {
			fireStart := FireNadeStart{Event: nadeEvent}

			if len(m.CurrentRound.Ticks) > 0 {
				current := &m.CurrentRound.Ticks[len(m.CurrentRound.Ticks)-1]
				if current.Tick_number == *curTick {
					current.FireNadeStart = append(current.FireNadeStart, fireStart)
				}
			}
		}
	})

	p.RegisterEventHandler(func(e events.FireGrenadeExpired) {
		if m.CurrentRound == nil {
			return
		}

		if e.Grenade == nil {
			return
		}

		nadeEvent := createNadeEvent(e.Grenade, p)
		if nadeEvent != nil {
			fireEnd := FireNadeEnd{Event: nadeEvent}

			if len(m.CurrentRound.Ticks) > 0 {
				current := &m.CurrentRound.Ticks[len(m.CurrentRound.Ticks)-1]
				if current.Tick_number == *curTick {
					current.FireNadeEnd = append(current.FireNadeEnd, fireEnd)
				}
			}
		}
	})

	p.RegisterEventHandler(func(e events.FlashExplode) {
		if m.CurrentRound == nil {
			return
		}

		if e.Grenade == nil {
			return
		}

		nadeEvent := createNadeEvent(e.Grenade, p)
		if nadeEvent != nil {
			flashBoom := FlashBoom{
				Event: nadeEvent,
			}

			if len(m.CurrentRound.Ticks) > 0 {
				current := &m.CurrentRound.Ticks[len(m.CurrentRound.Ticks)-1]
				if current.Tick_number == *curTick {
					current.FlashBoom = append(current.FlashBoom, flashBoom)
				}
			}
		}
	})

	p.RegisterEventHandler(func(e events.HeExplode) {
		if m.CurrentRound == nil {
			return
		}

		if e.Grenade == nil {
			return
		}

		nadeEvent := createNadeEvent(e.Grenade, p)
		if nadeEvent != nil {
			nadeBoom := NadeBoom{
				Event: nadeEvent,
			}

			if len(m.CurrentRound.Ticks) > 0 {
				current := &m.CurrentRound.Ticks[len(m.CurrentRound.Ticks)-1]
				if current.Tick_number == *curTick {
					current.NadeBoom = append(current.NadeBoom, nadeBoom)
				}
			}
		}
	})
}

// Helper function to create NadeEvent from grenade
func createNadeEvent(grenade *common.Equipment, p demoinfocs.Parser) *NadeEvent {
	if grenade == nil || grenade.Entity == nil {
		return nil
	}

	nadeEvent := &NadeEvent{
		Position: r3.Vector{
			X: grenade.Entity.Position().X,
			Y: grenade.Entity.Position().Y,
			Z: grenade.Entity.Position().Z,
		},
		NadeEntityId: grenade.Entity.ID(),
	}

	// Get thrower information
	if grenade.Owner != nil {
		nadeEvent.ThowerSteamId = grenade.Owner.SteamID64
		nadeEvent.ThowerName = grenade.Owner.Name
	}

	// Get weapon type
	if grenade != nil {
		nadeEvent.NadeType = weaponTypeFromEquipment(grenade)
	}

	return nadeEvent
}

func check_nades(g map[int]*common.GrenadeProjectile) {
	for _, i := range g {
		if i == nil {
			log.Println("This Gernade is empty")
		}
	}
}
