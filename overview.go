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

	IsFreezetime   bool                    `json:"is_freezetime"`
	IsWarmup       bool                    `json:"is_warmup"`
	IsMatchStarted bool                    `json:"is_match_started"`
	GamePhase      string                  `json:"game_phase"`
	Players        map[uint64]*Player_info `json:"players"`
	Nades          []Nades                 `json:"nades"`
	WeaponFired    []Weaps_fired           `json:"weapons_fired"`
	PlayersHurt    []PlayerHurt            `json:"players_hurt"`
}

type Weaps_fired struct {
	PlayerShot *common.Player
	Weap       *common.Equipment
}

type PlayerHurt struct {
	PlayerSteamID     uint64     `json:"player_steam_id"`
	AttackerSteamID   uint64     `json:"attacker_steam_id"`
	PlayerName        string     `json:"player_name"`
	AttackerName      string     `json:"attacker_name"`
	Health            int        `json:"health"`
	Armor             int        `json:"armor"`
	Weapon            WeaponType `json:"weapon"`
	WeaponString      string     `json:"weapon_string"`
	HealthDamage      int        `json:"health_dmg"`       // Raw health damage before overkill
	ArmorDamage       int        `json:"armor_dmg"`        // Raw armor damage before overkill
	HealthDamageTaken int        `json:"health_dmg_taken"` // Actual health damage (after overkill)
	ArmorDamageTaken  int        `json:"armor_dmg_taken"`  // Actual armor damage (after overkill)
	HitGroup          HitGroup   `json:"hit_group"`        // Body part that was hit
}

type Bullet_dmg struct {
	Attacker_Steam_Id  uint64  `json:"attacker_id_bullet"`
	Attacker_name      string  `json:"attacker_name_bulltet"`
	Victim_Steam_id    uint64  `json:"victim_id_bullet"`
	Victim_name        string  `json:"victim_name_bullet"`
	Dist               float32 `json:"distance"`
	DmgDirX            float32 `json:"damage_direction_x"`
	DmgDirY            float32 `json:"damage_direction_y"`
	DmgDirZ            float32 `json:"damage_direction_z"`
	NumberPenetrations int     `json:"number_penetrations"`
	IsNoScope          bool    `json:"is_no_scope_bullet"`
	IsAttInAir         bool    `json:"is_attacker_in_air"`
}

type BombDefuseAbort struct {
	PlayerAbortedSteamId uint64
	PayerAbortedName     string
	Position             Position
}

type BombDefuseStarted struct {
	PlayerStartedSteamId uint64
	PlayerStartedName    uint64
	Position             Position
	Kit                  bool
}

type BombDrop struct {
	PlayerDropSteamId uint64
	PlayerDropName    string
	Position          Position
}

type BombPickedUp struct {
	PlayerPickedUpSteamId uint64
	PlayerPickUpName      string
	Position              Position
}

type PlantAborted struct {
	PlayerAbortPlantSteamId uint64
	PlayerAbortPlantName    string
	Position                Position
}

type PlantBegin struct {
}

type Planted struct {
}

type NadeEvent struct {
	NadeType      WeaponType
	Nade          *common.Equipment
	Position      Position
	ThowerSteamId uint64
	ThowerName    string
	NadeEntityId  int
}

type DecoyDone struct {
	Event *NadeEvent
}

type DecoyStarted struct {
	Event *NadeEvent
}

type FireNadeStart struct {
	Event *NadeEvent
}

type FireNadeEnd struct {
	Event *NadeEvent
}

type FlashBoom struct {
	Event *NadeEvent
}

type NadeBoom struct {
	Event *NadeEvent
}

// What are the poperties of a player?
type Player_info struct {
	Name                  string                    `json:"name"`
	Inventory             map[int]*common.Equipment `json:"inventory"`
	Money                 int                       `json:"money"`
	Health                int                       `json:"health"`
	Armor                 int                       `json:"armor"`
	IsAlive               bool                      `json:"is_alive"`
	Team                  common.Team               `json:"team"`
	Entity_id             int                       `json:"entity_id"`
	Defusing              bool                      `json:"defusing"`
	HasHel                bool                      `json:"has_helmet"`
	Position              Position                  `json:"position_pl"`
	ActiveWeapon          *common.Equipment         `json:"active_weapon"`
	ActiveWeaponName      string                    `json:"active_weapon_name"`
	AmmoInMag             int                       `json:"ammo_in_mag"`
	AmmoInRes             int                       `json:"ammo_in_res" `
	FlashDurTime          time.Duration             `json:"flash_dur_time"`
	FlashDurTimeRemaining time.Duration             `json:"flash_dur_time_remaining"`
	HasKit                bool                      `json:"has_defuse_kit"`
	IsInAir               bool                      `json:"air_borne"`
	IsBlind               bool                      `json:"is_blinded"`
	InBombZone            bool                      `json:"in_bomb_zone"`
	InBuyZone             bool                      `json:"in_buy_zone"`
	IsScoped              bool                      `json:"is_scoped"`
	Standing              bool                      `json:"is_standing"`
	UnDuckingInProgress   bool                      `json:"un_ducking_in_prog"`
	Walking               bool                      `json:"is_walking"`
	LastPlace             string                    `json:"last_place_name"`
	ViewDirX              float32                   `json:"view_dir_x"`
	ViewDirY              float32                   `json:"view_dir_y"`
	Ducked                bool                      `json:"is_ducking"`
	FlashN                int8                      `json:"num_flashes"`
	NadeN                 int8                      `json:"num_nades"`
	SmokeN                int8                      `json:"num_smokes"`
	MollyN                int8                      `json:"num_mollies"`
	IncendiaryN           int8                      `json:"num_incindiary"`
	DecoyN                int8                      `json:"num_decoys"`
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
	Start_tick     int `json:"start_tick"`
	End_tick       int `json:"end_tick"`
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
