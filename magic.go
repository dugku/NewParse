package main

import "github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs/common"

type HitGroup int
type WeaponType int
type Bombsite int

const (
	HitGroupGeneric HitGroup = iota
	HitGroupHead
	HitGroupChest
	HitGroupStomach
	HitGroupLeftArm
	HitGroupRightArm
	HitGroupLeftLeg
	HitGroupRightLeg
	HitGroupGear
)

const (
	WeaponUnknown WeaponType = iota
	WeaponPistol
	WeaponSMG
	WeaponShotgun
	WeaponMachineGun
	WeaponRifle
	WeaponSniper
	WeaponGrenade
	WeaponKnife
	WeaponZeus
	WeaponEquipment
	WeaponWorld
	WeaponMax
)

const (
	BombsiteUnknown Bombsite = iota
	BombsiteA
	BombsiteB
)

func (hg HitGroup) String() string {
	return [...]string{
		"Generic",
		"Head",
		"Chest",
		"Stomach",
		"Left Arm",
		"Right Arm",
		"Left Leg",
		"Right Leg",
		"World",
	}[hg]
}

// String returns the human-readable name of the weapon type
func (wt WeaponType) String() string {
	names := []string{
		"Unknown",
		"Pistol",
		"SMG",
		"Shotgun",
		"MachineGun",
		"Rifle",
		"Sniper",
		"Grenade",
		"Knife",
		"Zeus",
		"Equipment",
		"World",
	}

	if int(wt) < 0 || int(wt) >= len(names) {
		return "Unknown"
	}
	return names[wt]
}

func (b Bombsite) String() string {
	switch b {
	case BombsiteA:
		return "A"
	case BombsiteB:
		return "B"
	default:
		return "Unknown"
	}
}

func weaponTypeFromEquipment(eq *common.Equipment) WeaponType {
	if eq == nil {
		return WeaponUnknown
	}

	switch eq.Type {
	// Pistols
	case common.EqP2000, common.EqGlock, common.EqP250, common.EqDeagle,
		common.EqFiveSeven, common.EqDualBerettas, common.EqTec9, common.EqCZ,
		common.EqUSP, common.EqRevolver:
		return WeaponPistol

	// SMGs
	case common.EqMP7, common.EqMP9, common.EqBizon, common.EqMac10,
		common.EqUMP, common.EqP90, common.EqMP5:
		return WeaponSMG

	// Shotguns
	case common.EqSawedOff, common.EqNova, common.EqMag7, common.EqXM1014:
		return WeaponShotgun

	// Machine Guns
	case common.EqM249, common.EqNegev:
		return WeaponMachineGun

	// Rifles
	case common.EqGalil, common.EqFamas, common.EqAK47, common.EqM4A4,
		common.EqM4A1, common.EqSG553, common.EqAUG:
		return WeaponRifle

	// Snipers
	case common.EqSSG08, common.EqAWP, common.EqScar20, common.EqG3SG1:
		return WeaponSniper

	// Grenades
	case common.EqDecoy, common.EqMolotov, common.EqIncendiary, common.EqFlash,
		common.EqSmoke, common.EqHE:
		return WeaponGrenade

	// Knives
	case common.EqKnife, common.EqAxe, common.EqHammer, common.EqWrench,
		common.EqFists:
		return WeaponKnife

	// Zeus
	case common.EqZeus:
		return WeaponZeus

	// World damage (fall damage, bomb, etc.)
	case common.EqWorld:
		return WeaponWorld

	// Equipment
	case common.EqKevlar, common.EqHelmet, common.EqBomb,
		common.EqDefuseKit, common.EqZoneRepulsor, common.EqShield,
		common.EqHeavyAssaultSuit, common.EqNightVision, common.EqHealthShot,
		common.EqTacticalAwarenessGrenade, common.EqBreachCharge,
		common.EqTablet, common.EqSnowball, common.EqBumpMine:
		return WeaponEquipment

	default:
		return WeaponUnknown
	}
}
