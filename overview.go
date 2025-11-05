package main

import "github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs/common"

type Tick struct {
	Tick_number int     `json:"tick_number"`
	Time_in_sec float32 `json:"time_in_sec"`

	Players map[uint64]Player_info `json:"players"`
}

// What are the poperties of a player?
type Player_info struct {
	steam_id  uint64                    `json:"steam_id"`
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
}

type Position_Player struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
	Z float32 `json:"z"`
}

// Will need to really think about this since
// I will be preprcessing the rounds into bins
// since I am processing the tick individually
type Round_info struct {
	round_number int
}

// More structs can be added here to hold additional information as needed
