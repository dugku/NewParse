import json
import glob
import os
import pandas as pd

TEAM_CT = 3
TEAM_T  = 2
WEAPON_RIFLE  = 5
WEAPON_SNIPER = 6
WEAPON_PISTOL = 1
WEAPON_KNIFE  = 8
TICKS_PER_SEC = 64.0

SITE_A_NAMES = {"BombsiteA", "Ramp", "Short", "Long", "ALong", "AShort"}
SITE_B_NAMES = {"BombsiteB", "BDoors", "BAppartments", "Tunnels", "BShort"}

def get_team_stats(players, team):
    alive = total_hp = total_money = rifle_count = awp_alive = 0
    pistol_only = smokes = mollies = flashes = nades = 0
    any_blinded = any_scoped = any_defusing = has_kit = 0
    in_a = in_b = bomb_carrier = 0

    for pid, p in players.items():
        if p.get("team") != team or not p.get("is_alive", False):
            continue
        alive       += 1
        total_hp    += p.get("health", 0)
        total_money += p.get("money", 0)

        w = p.get("active_weapon", 0)
        if w == WEAPON_RIFLE:  rifle_count += 1
        if w == WEAPON_SNIPER: awp_alive   += 1
        if w in (WEAPON_PISTOL, WEAPON_KNIFE, 0): pistol_only += 1

        smokes  += p.get("num_smokes", 0)
        mollies += p.get("num_mollies", 0) + p.get("num_incindiary", 0)
        flashes += p.get("num_flashes", 0)
        nades   += p.get("num_nades", 0)

        if p.get("is_blinded", False):  any_blinded  = 1
        if p.get("is_scoped",  False):  any_scoped   = 1
        if p.get("defusing",   False):  any_defusing = 1
        if p.get("has_defuse_kit", False): has_kit   = 1

        place = p.get("last_place_name", "")
        if any(s in place for s in SITE_A_NAMES):   in_a += 1
        elif any(s in place for s in SITE_B_NAMES): in_b += 1

        if team == TEAM_T and p.get("in_bomb_zone", False):
            bomb_carrier = 1

    return dict(alive=alive, total_hp=total_hp, money=total_money,
                rifle_count=rifle_count, awp_alive=awp_alive, pistol_only=pistol_only,
                smokes=smokes, mollies=mollies, flashes=flashes, nades=nades,
                any_blinded=any_blinded, any_scoped=any_scoped,
                any_defusing=any_defusing, has_kit=has_kit,
                in_a=in_a, in_b=in_b, bomb_carrier=bomb_carrier)

def derive_time_remaining(tick_num, start_tick, end_tick):
    duration = max(end_tick - start_tick, 1)
    elapsed  = max(tick_num - start_tick, 0)
    return round(max(duration - elapsed, 0) / TICKS_PER_SEC, 2)

def bomb_state(round_data, tick_num):
    for event in (round_data.get("Planted") or []):
        plant_tick = event.get("TickNum", 0)
        if plant_tick <= tick_num:
            return 1, event.get("Site", 0), round((tick_num - plant_tick) / TICKS_PER_SEC, 2)
    return 0, 0, 0.0

def process_file(filepath):
    tick_rows, round_rows = [], []
    with open(filepath) as f:
        data = json.load(f)

    demo     = data.get("demo", os.path.basename(filepath))
    map_name = data.get("map", "")

    for r in data.get("rounds", []):
        ct_win = r.get("ct_win", -1)
        if ct_win == -1:
            continue

        round_num  = r.get("round_number", -1)
        ct_score   = r.get("ct_score", 0)
        t_score    = r.get("t_score", 0)
        ct_econ    = r.get("ct_econ", 0)
        t_econ     = r.get("t_econ", 0)
        ct_equip   = r.get("ct_equipment_val", 0)
        t_equip    = r.get("t_equipment_val", 0)
        start_tick = r.get("start_tick", 0)
        end_tick   = r.get("end_tick", 0)
        end_reason = r.get("round_ended_reason", "")

        round_rows.append(dict(demo=demo, map=map_name, round_number=round_num,
            ct_score=ct_score, t_score=t_score, ct_econ=ct_econ, t_econ=t_econ,
            ct_equip_val=ct_equip, t_equip_val=t_equip, start_tick=start_tick,
            end_tick=end_tick, round_end_reason=end_reason, ct_win=ct_win))

        for tick in (r.get("Ticks") or []):
            if tick.get("is_freezetime", False):
                continue
            tick_num = tick.get("tick_number", 0)
            players  = tick.get("players") or {}

            ct = get_team_stats(players, TEAM_CT)
            t  = get_team_stats(players, TEAM_T)
            if ct["alive"] == 0 and t["alive"] == 0:
                continue

            bp, bs, tsp = bomb_state(r, tick_num)

            tick_rows.append({
                "demo": demo, "map": map_name,
                "round_number": round_num, "tick_number": tick_num,
                "time_remaining": derive_time_remaining(tick_num, start_tick, end_tick),
                "ct_score": ct_score, "t_score": t_score,
                # CT
                "ct_alive": ct["alive"], "ct_total_hp": ct["total_hp"],
                "ct_money": ct["money"], "ct_equip_val": ct_equip,
                "ct_rifle_count": ct["rifle_count"], "ct_awp_alive": ct["awp_alive"],
                "ct_pistol_only": ct["pistol_only"],
                "ct_smokes": ct["smokes"], "ct_mollies": ct["mollies"],
                "ct_flashes": ct["flashes"], "ct_nades": ct["nades"],
                "ct_any_blinded": ct["any_blinded"], "ct_any_scoped": ct["any_scoped"],
                "ct_any_defusing": ct["any_defusing"], "ct_has_kit": ct["has_kit"],
                "ct_in_bombsite_a": ct["in_a"], "ct_in_bombsite_b": ct["in_b"],
                # T
                "t_alive": t["alive"], "t_total_hp": t["total_hp"],
                "t_money": t["money"], "t_equip_val": t_equip,
                "t_rifle_count": t["rifle_count"], "t_awp_alive": t["awp_alive"],
                "t_pistol_only": t["pistol_only"],
                "t_smokes": t["smokes"], "t_mollies": t["mollies"],
                "t_flashes": t["flashes"], "t_nades": t["nades"],
                "t_any_blinded": t["any_blinded"], "t_any_scoped": t["any_scoped"],
                "t_any_planting": t["any_defusing"],
                "t_in_bombsite_a": t["in_a"], "t_in_bombsite_b": t["in_b"],
                "t_bomb_carrier_alive": t["bomb_carrier"],
                # Bomb
                "bomb_planted": bp, "bomb_site": bs, "time_since_plant": tsp,
                # Label
                "ct_win": ct_win,
            })

    return tick_rows, round_rows

def main():
    files = glob.glob("../jsonFolder/*.json")
    if not files:
        print("No JSON files found")
        return

    all_ticks, all_rounds = [], []
    for fp in files:
        print(f"Processing {fp}...")
        t, r = process_file(fp)
        all_ticks.extend(t)
        all_rounds.extend(r)
        print(f"  → {len(r)} rounds, {len(t)} ticks")

    tick_df  = pd.DataFrame(all_ticks)
    round_df = pd.DataFrame(all_rounds)
    tick_df.to_csv("ticks.csv", index=False)
    round_df.to_csv("rounds.csv", index=False)

    print(f"\nDone.")
    print(f"ticks.csv  → {len(tick_df)} rows, {len(tick_df.columns)} features")
    print(f"rounds.csv → {len(round_df)} rows")
    print(f"\nFeatures: {list(tick_df.columns)}")

if __name__ == "__main__":
    main()