package main

import (
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
	m.CurrentRound = &RoundInfo{}

	p.RegisterEventHandler(func(e events.RoundStart) {
		m.openRound = true
		//this is messy wtf will refactor later if I don't forget since it will be messier later
		gs := p.GameState()
		if gs != nil {
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
			m.CurrentRound.Start_tick = gs.IngameTick()
			m.CurrentRound.CTEcon = ctMoney
			m.CurrentRound.TEcon = tMoney
			m.CurrentRound.CTEquipmentVal = gs.TeamCounterTerrorists().CurrentEquipmentValue()
			m.CurrentRound.TEquipmentVal = gs.TeamCounterTerrorists().CurrentEquipmentValue()
			m.CurrentRound.CTScore = ctScore
			m.CurrentRound.TScore = tScore
			m.CurrentRound.Round_number = round_num

		}

	})

	p.RegisterEventHandler(func(e events.RoundEnd) {
		m.openRound = false
		gs := p.GameState()
		if gs != nil {
			m.CurrentRound.End_tick = gs.IngameTick()

		}
		m.StoreRoundInfo(round_num, *m.CurrentRound)
		*c += 1
	})

}
