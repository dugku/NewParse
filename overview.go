package main

import (
	"time"

	"github.com/golang/geo/r3"
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
	openRound    bool
	Players      map[uint64]PlayerStats
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
	DefuseStart    []BombDefuseStarted     `json:"defuse_started"`
	DefuseAbort    []BombDefuseAbort       `json:"bomb_defuse_abort"`
	BombaDropped   []BombDrop              `json:"bomb_dropped"`
	BombPickUp     []BombPickedUp          `json:"bomb_picked_up"`
	PlantBegin     []PlantBegin
	PlantAborted   []PlantAborted
	Planted        []Planted
	DecoyStarted   []DecoyStarted
	DecoyDone      []DecoyDone
	FireNadeStart  []FireNadeStart
	FireNadeEnd    []FireNadeEnd
	FlashBoom      []FlashBoom
	NadeBoom       []NadeBoom
}

type Weaps_fired struct {
	PlayerSteamIDFired uint64     `json:"player_fired_steam_id"`
	PlayerNameFired    string     `json:"player_name_fired"`
	WeaponFired        WeaponType `json:"weapons_fired"`
	WeaponFiredString  string     `json:"weapon_fired_string"`
	TickNumWeap        int        `json:"-"` //interal only.
	Position           Position
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
	TickNum           int        `json:"-"`                //for internal purposes only
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
	TickNum              int
}

type BombDefuseStarted struct {
	PlayerStartedSteamId uint64
	PlayerStartedName    string
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
	PlayerBeginPlantSteamId uint64
	PlayerBeginPlantName    string
	Position                Position
	TickNum                 int
}

type Planted struct {
	PlayerPlantedSteamId uint64
	PlayerPlantedName    string
	Position             Position
	TickNum              int
	Site                 Bombsite
}

type NadeEvent struct {
	NadeType      WeaponType
	Nade          *common.Equipment
	Position      r3.Vector
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

type RoundKill struct {
	TimeOfKill time.Duration
	Tick       int

	Weapon WeaponType

	VictimSteamID   uint64
	VictimName      string
	VictimX         float64
	VictimY         float64
	VictimViewX     float32
	VictimViewY     float32
	VictimIsFlashed bool

	KillerSteamID   uint64
	KillerName      string
	KillerHealth    int
	KillerX         float64
	KillerY         float64
	KillerViewX     float32
	KillerViewY     float32
	KillerIsFlashed bool

	AssistorSteamID uint64
	AssistorName    string

	IsOpeing        bool
	IsHeadshot      bool
	IsWallbang      bool
	IsNoScope       bool
	IsThroughtSmoke bool
	IsAssistFlash   bool
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
	Start_tick   int `json:"start_tick"`
	End_tick     int `json:"end_tick"`
	Round_number int `json:"round_number"`
	TimeStart    time.Duration
	TimeEnd      time.Duration

	TScore  int `json:"t_score"`
	CTScore int `json:"ct_score"`

	CTEcon         int `json:"ct_econ"` //Raw money
	TEcon          int `json:"t_econ"`
	CTEquipmentVal int `json:"ct_equipment_val"` //Value of equipment
	TEquipmentVal  int `json:"t_equipment_val"`
	TypeOfBuyCT    string
	TypeOfBuyT     string

	FirstKillCount int
	SuvivorsCT     []uint64
	SurvivorsT     []uint64

	BombPlanted      bool
	PlayerPlanted    string
	BombPlantedSite  string
	RoundEndedReason string

	PlayersAliveCT map[uint64]bool
	PlayerAliveT   map[uint64]bool
	OneVX          bool
	OneVXCount     int

	Kills map[int]RoundKill

	Ticks []Tick
	//Get more Later concept for now
}

type PlayerStats struct {
	Username string
	SteamId  uint64

	ImpactPreRound      float64
	Kills               int
	Deaths              int
	Assists             int
	HS                  int
	HeadshotPercent     float64
	ADR                 float64
	KAST                float64
	KDRatio             float64
	FirstKill           int
	FirstDeath          int
	FKDiff              int
	Round2k             int
	Round3k             int
	Round4k             int
	Round5k             int
	TotalDmg            int
	TradeKills          int
	TradeDeaths         int
	CTKills             int
	TKills              int
	EffectiveFlashes    int
	AvgFlashDuration    int
	WeaponKills         map[WeaponType]int
	WeaponKillsHeadshot map[WeaponType]int
	AvgDist             float64
	TotalDist           float64
	FlashesThrown       int
	TotalUtilDamage     int
	AvgKillRnd          float64
	AvgDeathsRnd        float64
	AvgAssistsRnd       float64
	AvgNadeDmg          float64
	AvgInferDmg         float64
	RoundsSurvived      int
	RoundTraded         int
	RoundContrid        []int
	InfernoDmg          int
	NadeDmg             int
	OpeningPercent      float64
	OpeningAttpPercent  float64
	OpeningRoundsWon    int
	OpeningWinPercent   float64
	OneVsOne            int
	OneVsTwo            int
	OneVsThree          int
	OneVsFour           int
	OneVsFive           int
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
