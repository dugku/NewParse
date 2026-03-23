package main

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
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
	//wirte current round
	m.WriteRound(info)
	m.Rounds = append(m.Rounds, info)
}

func (m *Match) WriteRound(round RoundInfo) {
	//write to json
	json_data, err := json.MarshalIndent(round, "", " ")
	if err != nil {
		log.Println("Error marshalling to JSON:", err)
		return
	}

	roundNum := strconv.Itoa(round.Round_number)
	file_name := "round_" + roundNum + ".json"
	err = writeFile(file_name, json_data)
	if err != nil {
		log.Println("Error writing to file:", err)
		return
	}
}

func writeFile(filename string, data []byte) error {
	err := os.WriteFile(filename, data, 0644)
	return err
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
