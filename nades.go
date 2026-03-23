package main

import (
	"fmt"
	"log"

	"github.com/golang/geo/r3"
	"github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs"
	"github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs/common"
	"github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs/events"
)

func nades(m *Match, p demoinfocs.Parser, cutTick *int) {
	gs := p.GameState()
	gernade_entities := gs.GrenadeProjectiles()
	check_nades(gernade_entities)

	if m.CurrentRound == nil {
		return
	}

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

			m.CurrentRound.Nades = append(m.CurrentRound.Nades, n)
		}
	}()
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

			m.CurrentRound.DecoyDone = append(m.CurrentRound.DecoyDone, decoyDone)

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

			m.CurrentRound.DecoyStarted = append(m.CurrentRound.DecoyStarted, decoyStarted)

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

			m.CurrentRound.FireNadeStart = append(m.CurrentRound.FireNadeStart, fireStart)

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

			fmt.Println("HERE IN NADES FIREEE THINGNNNGNNGN")
			m.CurrentRound.FireNadeEnd = append(m.CurrentRound.FireNadeEnd, fireEnd)
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

			m.CurrentRound.FlashBoom = append(m.CurrentRound.FlashBoom, flashBoom)

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

			m.CurrentRound.NadeBoom = append(m.CurrentRound.NadeBoom, nadeBoom)

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
		TimeInSec:    int(p.CurrentTime().Seconds()),
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

type FlatGameFrame struct {
	// --- Global Context ---
	RoundNum      int     `json:"round_num"`
	TickNum       int     `json:"tick_num"`
	TimeRemaining float32 `json:"time_remaining"`

	// --- Match State ---
	IsBombPlanted int // 0=False, 1=True
	BombSite      int // 0=None, 1=A, 2=B (If planted)
	CTScore       int
	TScore        int

	// --- Team Economies (Aggregated features are very strong for ML) ---
	CTTotalMoney     int
	TTotalMoney      int
	CTEquipmentValue int
	TEquipmentValue  int
	CTAliveCount     int
	TAliveCount      int

	// --- The Players (ALWAYS size 5) ---
	// We use arrays, not slices, to enforce fixed size.
	CTPlayers [5]FlatPlayer `json:"ct_players"`
	TPlayers  [5]FlatPlayer `json:"t_players"`
}

// FlatPlayer contains strictly numeric/boolean data for ML.
type FlatPlayer struct {
	IsActive  int // 1 if this slot is occupied by a real player, 0 if bot/empty
	IsAlive   int // 1=Alive, 0=Dead
	Health    int
	Armor     int
	HasHelmet int // 0 or 1
	HasDefuse int // 0 or 1
	HasBomb   int // 0 or 1

	// Position
	X, Y, Z      float32
	ViewX, ViewY float32 // Where are they looking? (Yaw/Pitch)

	// Status
	IsBlinded  int     // 1 if fully blind
	FlashDur   float32 // Exact duration remaining
	IsScoped   int
	IsAirborne int

	// Economy / Equipment
	Money      int
	WeaponID   int // Integer mapping (e.g., 1=Pistol, 2=Rifle, 3=AWP)
	Ammo       int
	TotalNades int // Simple count of util
}
