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

func round_start_end(p demoinfocs.Parser, open_round *bool, m *Match) {
	var (
		tScore    int
		ctScore   int
		round_num int
	)
	var info RoundInfo

	p.RegisterEventHandler(func(e events.RoundStart) {
		*open_round = true

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

			info.CTEcon = ctMoney
			info.TEcon = tMoney
			info.CTEquipmentVal = gs.TeamCounterTerrorists().CurrentEquipmentValue()
			info.TEquipmentVal = gs.TeamCounterTerrorists().CurrentEquipmentValue()
			info.CTScore = ctScore
			info.TScore = tScore
			info.Round_number = round_num
		}

	})

	p.RegisterEventHandler(func(e events.RoundEnd) {
		fmt.Println("Here in end")
		*open_round = false
		m.StoreRoundInfo(round_num, *m.CurrentRound)
	})

}
