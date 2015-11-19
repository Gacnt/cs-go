# cs-go
An addon to interface with CSGO from Gamecoordinator using [go-steam](https://github.com/Philipp15b/go-steam)

in the `*steam.LoggedOnEvent` simply call:

```
cs = csgo.NewCSGO(client)
cs.SetPlaying(true) // Must be set to activate counter-strike
cs.ShakeHands() // This will shake hands with the gamecoordinator
```

After those 3 are called, simply wait for the `*csgo.GCReadyEvent`. Once that is called you are ready to go and can use any methods. e.g.

```
cs.GetPlayerProfile(123451613613)
```
Will return a struct layed out like:
```
{
  "request_id": null,
  "account_profiles": [
    {
      "account_id": 137013074,
      "ongoingmatch": null,
      "global_stats": null,
      "penalty_seconds": null,
      "penalty_reason": null,
      "vac_banned": null,
      "ranking": {
        "account_id": 137013074,
        "rank_id": 11,
        "wins": 192,
        "rank_change": null
      },
      "commendation": {
        "cmd_friendly": 3,
        "cmd_teaching": 3,
        "cmd_leader": 3
      },
      "medals": {
        "medal_team": 0,
        "medal_combat": 0,
        "medal_weapon": 0,
        "medal_global": 0,
        "medal_arms": 0,
        "display_items_defidx": [],
        "featured_display_item_defidx": null
      },
      "my_current_event": null,
      "my_current_event_teams": [],
      "my_current_team": null,
      "my_current_event_stages": [],
      "survey_vote": null,
      "activity": null,
      "player_level": 3,
      "player_cur_xp": 327684342,
      "player_xp_bonus_flags": null
    }
  ]
}
```

currently working on expanding and adding more methods.

