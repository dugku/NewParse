package main

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strings"
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

	if t.Tick_number%16 != 0 {
		return
	}

	m.CaptureFrame(t)
}

func (m *Match) StoreRoundInfo(round int, info RoundInfo) {
	m.Rounds = append(m.Rounds, info)
}

// MatchOutput is the top-level JSON structure written once per demo.
type MatchOutput struct {
	Demo   string      `json:"demo"`
	Map    string      `json:"map"`
	Rounds []RoundInfo `json:"rounds"`
}

// WriteMatch writes the entire match to a single JSON file once parsing is done.
// demoPath is the original .dem file path — used to derive the output filename.
func (m *Match) WriteMatch(demoPath string) error {
	// Derive a clean filename from the demo path:
	// e.g. "parivision-vs-g2-m2-dust2.dem" → "parivision-vs-g2-m2-dust2.json"
	base := filepath.Base(demoPath)
	name := strings.TrimSuffix(base, filepath.Ext(base))
	fileName := "./jsonFolder/" + name + ".json"

	out := MatchOutput{
		Demo:   name,
		Map:    m.MapName,
		Rounds: m.Rounds,
	}

	data, err := json.Marshal(out) // No indent — saves significant space vs MarshalIndent
	if err != nil {
		return err
	}

	err = os.WriteFile(fileName, data, 0644)
	if err != nil {
		return err
	}

	log.Printf("Wrote %d rounds to %s (%.2f MB)", len(m.Rounds), fileName, float64(len(data))/(1024*1024))
	return nil
}

func writeFile(filename string, data []byte) error {
	return os.WriteFile(filename, data, 0644)
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
