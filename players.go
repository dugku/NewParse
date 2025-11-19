package main

import (
	"log"

	"github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs"
	"github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs/common"
	"github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs/events"
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
			continue //would log this but parsing would take forever if I did
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

func players_hurting(p demoinfocs.Parser, m *Match, curT *int) {
	p.RegisterEventHandler(func(e events.PlayerHurt) {

		gs := p.GameState()

		if !m.openRound || m.CurrentRound == nil {
			return
		}

		hurt := PlayerHurt{}
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Recovering in Player Hurt %v", r)
			}
		}()
		if e.Player != nil {
			hurt.PlayerSteamID = e.Player.SteamID64
			hurt.PlayerName = e.Player.Name
		} else {
			hurt.PlayerName = "Unknown"
			log.Printf("Warning: PlayerHurt event with nil player at tick %d", gs.IngameTick())
		}

		if e.Attacker != nil {
			hurt.AttackerSteamID = e.Attacker.SteamID64
			hurt.AttackerName = e.Attacker.Name
		} else {
			hurt.AttackerName = "World"
		}

		if e.Weapon != nil {
			hurt.Weapon = weaponTypeFromEquipment(e.Weapon)
			hurt.WeaponString = e.Weapon.String()
		} else {
			hurt.Weapon = WeaponUnknown
			hurt.WeaponString = "Unknown"
		}

		hurt.Health = e.Health
		hurt.Armor = e.Armor
		hurt.HealthDamage = e.HealthDamage
		hurt.ArmorDamage = e.ArmorDamage
		hurt.HealthDamageTaken = e.HealthDamageTaken
		hurt.ArmorDamageTaken = e.ArmorDamageTaken
		hurt.HitGroup = HitGroup(e.HitGroup)
		hurt.TickNum = *curT

		if len(m.CurrentRound.Ticks) > 0 {
			current := &m.CurrentRound.Ticks[len(m.CurrentRound.Ticks)-1]
			if current.Tick_number == *curT {
				current.PlayersHurt = append(current.PlayersHurt, hurt)
			}
		}

	})

}

func weapons_firing(p demoinfocs.Parser, m *Match, curT *int) {
	p.RegisterEventHandler(func(e events.WeaponFire) {
		if !m.openRound || m.CurrentRound == nil {
			return
		}

		fired := Weaps_fired{}

		if e.Weapon != nil {
			fired.WeaponFired = weaponTypeFromEquipment(e.Weapon)
			fired.WeaponFiredString = e.Weapon.String()
		} else {
			fired.WeaponFired = WeaponUnknown
			fired.WeaponFiredString = "Unknown"
		}

		if e.Shooter != nil {
			fired.PlayerSteamIDFired = e.Shooter.SteamID64
			fired.PlayerNameFired = e.Shooter.Name
			fired.Position.X = e.Shooter.Position().X
			fired.Position.Y = e.Shooter.Position().Y
			fired.Position.Z = e.Shooter.Position().Z
		} else {
			fired.PlayerNameFired = "World"
		}

		fired.TickNumWeap = *curT

		if len(m.CurrentRound.Ticks) > 0 {
			current := &m.CurrentRound.Ticks[len(m.CurrentRound.Ticks)-1]
			if current.Tick_number == *curT {
				current.WeaponFired = append(current.WeaponFired, fired)
			}
		}
	})
}

func kill_logic(p demoinfocs.Parser, m *Match) {
	current := &m.CurrentRound
	gs := p.GameState()

	if gs == nil {
		log.Println("GameState is nil somehow")
	}
	p.RegisterEventHandler(func(e events.Kill) {
		if e.Killer == nil || e.Victim == nil {
			log.Println("Killer or Victim is nil")
			return
		}

		if gs.IsWarmupPeriod() {
			return
		}

		if e.IsHeadshot && e.Weapon != nil {
			add_headshot(e.Killer, *e.Weapon, m)
		}

		var assistor string
		if e.Assister != nil {
			assistor = e.Assister.Name
		}
		open_kill := false
		if current != nil {
			count := len(m.CurrentRound.Kills) + 1

			if count == 1 {
				open_kill = true
			} else {
				open_kill = false
			}

			if _, exists := m.CurrentRound.Kills[count]; exists {
				return
			} else {
				m.CurrentRound.Kills[count] = RoundKill{
					TimeOfKill: p.CurrentTime(),
					Tick:       gs.IngameTick(),

					Weapon: weaponTypeFromEquipment(e.Weapon),

					VictimSteamID:   e.Victim.SteamID64,
					VictimName:      e.Victim.Name,
					VictimX:         e.Victim.Position().X,
					VictimY:         e.Victim.Position().Y,
					VictimViewX:     e.Victim.ViewDirectionX(),
					VictimViewY:     e.Victim.ViewDirectionY(),
					VictimIsFlashed: e.Victim.IsBlinded(),

					KillerSteamID:   e.Killer.SteamID64,
					KillerName:      e.Killer.Name,
					KillerHealth:    e.Killer.Health(),
					KillerX:         e.Killer.Position().X,
					KillerY:         e.Killer.Position().Y,
					KillerViewX:     e.Killer.ViewDirectionX(),
					KillerViewY:     e.Killer.ViewDirectionY(),
					KillerIsFlashed: e.Killer.IsBlinded(),

					AssistorSteamID: e.Assister.SteamID64,
					AssistorName:    assistor,

					IsOpeing:        open_kill,
					IsHeadshot:      e.IsHeadshot,
					IsWallbang:      e.IsWallBang(),
					IsNoScope:       e.NoScope,
					IsThroughtSmoke: e.ThroughSmoke,
					IsAssistFlash:   e.AssistedFlash,
				}
			}
			first_kill(m, e.Killer, e.Victim)
			round_contributed(m, e.Killer, e.Assister)
		}
	})
}

func first_kill(m *Match, k *common.Player, v *common.Player) {
	if m.CurrentRound.FirstKillCount == 1 {
		m.CurrentRound.FirstKillCount = 1

		player_k := k.SteamID64
		player_v := v.SteamID64

		player_stat_kill, exists := m.Players[player_k]
		if !exists {
			return
		}

		player_stat_kill.FirstKill++
		m.Players[player_k] = player_stat_kill

		player_stat_vict, exists2 := m.Players[player_v]
		if !exists2 {
			return
		}

		player_stat_vict.FirstDeath++
		m.Players[player_v] = player_stat_vict
	}
}

func round_contributed(m *Match, k *common.Player, a *common.Player) {
	kill_id := k.SteamID64

	player, exists := m.Players[kill_id]
	if !exists {
		return
	}

	player.RoundContrid = append(player.RoundContrid, m.CurrentRound.Round_number)
	m.Players[kill_id] = player

	if a != nil {
		if player_assist, exists2 := m.Players[a.SteamID64]; exists2 {
			player_assist.RoundContrid = append(player_assist.RoundContrid, m.CurrentRound.Round_number)
			m.Players[a.SteamID64] = player_assist
		}
	}
}

func update_weapon_kills() {

}

func one_verus_X(m *Match) {

}

func add_headshot(c *common.Player, w common.Equipment, m *Match) {
	player_id := c.SteamID64
	player, exists := m.Players[player_id]
	if !exists {
		return
	}

	player.HS++
	player.WeaponKillsHeadshot[weaponTypeFromEquipment(&w)]++

	m.Players[player_id] = player
}

func check_team(t *common.TeamState) {
	if t == nil {
		log.Println("One of the TeamStates is nil not good")
	}
}
