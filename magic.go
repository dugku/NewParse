package main

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
	WeaponRifle
	WeaponSMG
	WeaponHeavy
	WeaponSniper
	WeaponGrenade
	WeaponKnife
	WeaponWorld
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
	return [...]string{
		"Unknown",
		"Pistol",
		"Rifle",
		"SMG",
		"Heavy",
		"Sniper",
		"Grenade",
		"Knife",
		"World",
	}[wt]
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
