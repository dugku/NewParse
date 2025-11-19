package main

import (
	"log"

	"github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs"
	"github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs/common"
	"github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs/events"
)

func player_setter(c *common.Player) PlayerStats {
	return PlayerStats{
		Username: c.Name,
		SteamId:  c.SteamID64,

		Kills:               0,
		Deaths:              0,
		Assists:             0,
		HS:                  0,
		KAST:                0,
		KDRatio:             0,
		FirstKill:           0,
		FirstDeath:          0,
		FKDiff:              0,
		Round2k:             0,
		Round3k:             0,
		Round4k:             0,
		Round5k:             0,
		TotalDmg:            0,
		TradeKills:          0,
		TradeDeaths:         0,
		CTKills:             0,
		TKills:              0,
		EffectiveFlashes:    0,
		AvgFlashDuration:    0,
		WeaponKills:         all_weapons(),
		WeaponKillsHeadshot: all_weapons(),
		AvgDist:             0,
		TotalDist:           0,
		AvgKillRnd:          0,
		AvgDeathsRnd:        0,
		AvgAssistsRnd:       0,
		AvgNadeDmg:          0,
		AvgInferDmg:         0,
		RoundsSurvived:      0,
		RoundTraded:         0,
		RoundContrid:        make([]int, 0),
		InfernoDmg:          0,
		NadeDmg:             0,
		OpeningPercent:      0,
		OpeningAttpPercent:  0,
		OpeningRoundsWon:    0,
		OneVsOne:            0,
		OneVsTwo:            0,
		OneVsThree:          0,
		OneVsFour:           0,
		OneVsFive:           0,
	}
}

func all_weapons() map[WeaponType]int { return make(map[WeaponType]int) }

func player_get(p demoinfocs.Parser, m *Match) {
	p.RegisterEventHandler(func(e events.RoundEnd) {
		gs := p.GameState()

		teamA := gs.TeamCounterTerrorists()
		teamB := gs.TeamTerrorists()

		if teamA == nil || teamB == nil {
			log.Panicln("Either one or both teams are nil")
			return
		}

		teamA_mem := teamA.Members()
		teamB_mem := teamB.Members()

		if teamA_mem == nil || teamB_mem == nil {
			log.Panicln("Either one or both team member list are nil")
			return
		}

		stat_setter(teamA_mem, m, gs)
		stat_setter(teamB_mem, m, gs)
	})
}

func stat_setter(c []*common.Player, m *Match, gs demoinfocs.GameState) {
	for i := range c {
		steam_id := c[i].SteamID64

		player, exists := m.Players[steam_id]
		if !exists {
			continue
		}

		player.Kills = c[i].Kills()
		player.Deaths = c[i].Deaths()
		player.Assists = c[i].Assists()
		player.TotalDmg = c[i].TotalDamage()
		player.ADR = calc_adr(gs, player.TotalDmg)
	}
}
