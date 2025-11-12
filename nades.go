package main

import (
	"log"

	"github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs"
	"github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs/common"
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

func check_nades(g map[int]*common.GrenadeProjectile) {
	for _, i := range g {
		if i == nil {
			log.Println("This Gernade is empty")
		}
	}
}
