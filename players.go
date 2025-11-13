package main

import (
	"fmt"
	"log"

	"github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs"
	"github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs/common"
)

type NadeType int

const (
	Decoy      NadeType = 501
	Molotov             = 502
	Incendiary          = 503
	Flash               = 504
	Smoke               = 505
	HE                  = 506
)

// take in game state to extract player info and record player info into tick
func test_players(gs demoinfocs.GameState, tick *Tick) {

	//need to see if common.Team come backj nil

	TeamCounter := gs.Team(common.TeamCounterTerrorists)
	TeamTerror := gs.Team(common.TeamTerrorists)

	check_team(TeamCounter)
	check_team(TeamTerror)

	players_counter := gs.Participants().TeamMembers(common.TeamCounterTerrorists)
	players_terror := gs.Participants().TeamMembers(common.TeamTerrorists)

	set_player(players_counter, tick, gs)
	set_player(players_terror, tick, gs)
}

func set_player(players []*common.Player, tick *Tick, gs demoinfocs.GameState) {
	for _, p := range players {
		if !p.IsAlive() || p == nil {
			fmt.Println("Either Player is Dead or is nil")
			continue
		}

		tick.Players[p.SteamID64] = &Player_info{
			Name:      p.Name,
			Inventory: make(map[int]*common.Equipment),
			Money:     p.Money(),
			Health:    p.Health(),
			Armor:     p.Armor(),
			IsAlive:   p.IsAlive(),
			Team:      p.Team,
			Entity_id: p.EntityID,
			Defusing:  p.IsDefusing,
			HasHel:    p.HasHelmet(),
			Position: Position{
				X: p.Position().X,
				Y: p.Position().Y,
				Z: p.Position().Z,
			},
			ActiveWeapon:          p.ActiveWeapon(),
			ActiveWeaponName:      p.ActiveWeapon().OriginalString,
			AmmoInMag:             p.ActiveWeapon().AmmoInMagazine(),
			AmmoInRes:             p.ActiveWeapon().AmmoReserve(),
			FlashDurTime:          p.FlashDurationTime(),
			FlashDurTimeRemaining: p.FlashDurationTimeRemaining(),
			HasKit:                p.HasDefuseKit(),
			IsInAir:               p.IsAirborne(),
			IsBlind:               p.IsBlinded(),
			InBombZone:            p.IsInBombZone(),
			InBuyZone:             p.IsInBuyZone(),
			IsScoped:              p.IsScoped(),
			Standing:              p.IsStanding(),
			UnDuckingInProgress:   p.IsUnDuckingInProgress(),
			Walking:               p.IsWalking(),
			LastPlace:             p.LastPlaceName(),
			ViewDirX:              p.ViewDirectionX(),
			ViewDirY:              p.ViewDirectionY(),
			FlashN:                0,
			NadeN:                 0,
			SmokeN:                0,
			MollyN:                0,
			DecoyN:                0,
		}

		for i, item := range p.Weapons() {
			tick.Players[p.SteamID64].Inventory[i] = item
		}

		for _, item := range p.Weapons() {
			if item.Type == common.EquipmentType(Decoy) {
				tick.Players[p.SteamID64].DecoyN += 1
			}
			if item.Type == common.EquipmentType(Molotov) {
				tick.Players[p.SteamID64].MollyN += 1
			}
			if item.Type == common.EquipmentType(Incendiary) {
				tick.Players[p.SteamID64].IncendiaryN += 1
			}
			if item.Type == common.EquipmentType(Flash) {
				tick.Players[p.SteamID64].FlashN += 1
			}
			if item.Type == common.EquipmentType(Smoke) {
				tick.Players[p.SteamID64].SmokeN += 1
			}
			if item.Type == common.EquipmentType(HE) {
				tick.Players[p.SteamID64].NadeN += 1
			}
		}
	}
}

func check_team(t *common.TeamState) {
	if t == nil {
		log.Println("One of the TeamStates is nil not good")
	}
}
