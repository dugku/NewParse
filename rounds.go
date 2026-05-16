package main

import (
	"fmt"

	"github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs"
	"github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs/common"
	"github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs/events"
)

func get_rounds_info(gs demoinfocs.GameState, tick *Tick) Tick {
	tick.IsFreezetime = gs.IsFreezetimePeriod()
	tick.IsWarmup = gs.IsWarmupPeriod()
	tick.IsMatchStarted = gs.IsMatchStarted()
	tick.GamePhase = gs.GamePhase().String()
	return *tick
}

func round_start_end(p demoinfocs.Parser, m *Match, c *int) {
	var (
		tScore    int
		ctScore   int
		round_num int
	)

	p.RegisterEventHandler(func(e events.RoundStart) {
		m.openRound = true

		gs := p.GameState()
		if gs == nil {
			return
		}

		round_num = gs.TotalRoundsPlayed()

		if t := gs.TeamCounterTerrorists(); t != nil {
			ctScore = t.Score()
		}
		if t := gs.TeamTerrorists(); t != nil {
			tScore = t.Score()
		}

		ctMoney := 0
		tMoney := 0
		for _, p := range gs.Participants().TeamMembers(common.TeamCounterTerrorists) {
			ctMoney += p.Money()
		}
		for _, p := range gs.Participants().TeamMembers(common.TeamTerrorists) {
			tMoney += p.Money()
		}

		fmt.Println(ctScore, tScore)

		m.CurrentRound.Start_tick = gs.IngameTick()
		m.CurrentRound.CTEcon = ctMoney
		m.CurrentRound.TEcon = tMoney
		m.CurrentRound.CTEquipmentVal = gs.TeamCounterTerrorists().CurrentEquipmentValue()
		m.CurrentRound.TEquipmentVal = gs.TeamTerrorists().CurrentEquipmentValue() // fixed: was CT twice
		m.CurrentRound.CTScore = ctScore
		m.CurrentRound.TScore = tScore
		m.CurrentRound.Round_number = round_num
	})

	p.RegisterEventHandler(func(e events.RoundEnd) {
		m.openRound = false

		gs := p.GameState()
		if gs != nil {
			m.CurrentRound.End_tick = gs.IngameTick()
		}

		// Capture who won and why — this is the label for the WPA model
		switch e.Winner {
		case common.TeamCounterTerrorists:
			m.CurrentRound.CTWin = 1
		case common.TeamTerrorists:
			m.CurrentRound.CTWin = 0
		default:
			m.CurrentRound.CTWin = -1 // unknown / draw
		}

		m.CurrentRound.RoundEndedReason = roundEndReasonString(e.Reason)

		m.StoreRoundInfo(round_num, *m.CurrentRound)
		*c += 1

		m.CurrentRound = &RoundInfo{
			Kills: make(map[int]RoundKill),
		}
	})
}
