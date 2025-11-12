package main

import (
	"log"
)

func (s *Match) CaptureFrame(t Tick) {
	if s == nil {
		log.Fatalln("This is nil")
		return
	}
	s.CurrentRound.Ticks = append(s.CurrentRound.Ticks, t)
}

func (m *Match) SeeFrame(t Tick) {
	if m.CurrentRound == nil {
		log.Panicln("See frame Stops")
		return
	}
	m.CaptureFrame(t)
}

func (m *Match) StoreRoundInfo(round int, info RoundInfo) {
	m.Rounds = append(m.Rounds, info)
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
