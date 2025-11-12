package main

import (
	"time"

	"github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs/common"
)

type MatchSink interface {
	StoreRoundInfo(round int, info RoundInfo)
	SeeFrame(tick Tick)
}

type RoundCap interface {
	CaptureFrame(ini RoundInfo, tick Tick)
}

type Match struct {
	Rounds       []RoundInfo
	CurrentRound *RoundInfo
}

type Tick struct {
	Tick_number int     `json:"tick_number"`
	Time_in_sec float32 `json:"time_in_sec"`

	IsFreezetime   bool                   `json:"is_freezetime"`
	IsWarmup       bool                   `json:"is_warmup"`
	IsMatchStarted bool                   `json:"is_match_started"`
	GamePhase      string                 `json:"game_phase"`
	Players        map[uint64]Player_info `json:"players"`
	Nades          []Nades                `json:"nades"`
}

// What are the poperties of a player?
type Player_info struct {
	Steam_id  uint64                    `json:"steam_id"`
	Name      string                    `json:"name"`
	Inventory map[int]*common.Equipment `json:"inventory"`
	Money     int                       `json:"money"`
	Health    int                       `json:"health"`
	Armor     int                       `json:"armor"`
	IsAlive   bool                      `json:"is_alive"`
	Team      common.Team               `json:"team"`
	Entity_id int                       `json:"entity_id"`
	Defusing  bool                      `json:"defusing"`
	HasHel    bool                      `json:"has_helmet"`
	Position  Position                  `json:"position_pl"`
}

type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

type Velocityy struct {
	VX float64
	VY float64
	VZ float64
}

// Will need to really think about this since
// I will be preprcessing the rounds into bins
// since I am processing the tick individually
type RoundInfo struct {
	Round_number   int `json:"round_number"`
	TScore         int `json:"t_score"`
	CTScore        int `json:"ct_score"`
	CTEcon         int `json:"ct_econ"` //Raw money
	TEcon          int `json:"t_econ"`
	CTEquipmentVal int `json:"ct_equipment_val"` //Value of equipment
	TEquipmentVal  int `json:"t_equipment_val"`

	Ticks []Tick
	//Get more Later concept for now
}

type Nades struct {
	EId       common.EquipmentClass
	Type      *common.Equipment
	Pos       Position
	TimeInSec time.Duration
	//Vel       Velocityy
}

type DmgDone struct {
}

// More structs can be added here to hold additional information as needed
