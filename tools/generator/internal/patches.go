package internal

func patchTimeTypes(t *goType) {
	if t == nil {
		return
	}
	for _, field := range t.Fields {
		if field.IsStruct() {
			patchTimeTypes(field)
		} else if field.Name == "ExpiresAt" && (field.TypeStr == "int" || field.TypeStr == "*int") {
			field.TypeStr = "*wgnTime.UnixTime"
			field.AddImport(wgModuleName + "/wargaming/wgnTime")
		}
	}
}

func patchEndpoint(ep *endpoint) {
	patchTimeTypes(ep.ReturnType)
	patchTimeTypes(ep.Options)
	for _, param := range ep.Parameters {
		if param.Name == "type" {
			param.Name = "typ"
		}
		if param.Name == "expiresAt" && (param.TypeStr == "int" || param.TypeStr == "*int") {
			param.TypeStr = "*wgnTime.UnixTime"
			param.AddImport(wgModuleName + "/wargaming/wgnTime")
		}
	}

	switch ep.Id {
	case "wot_account_info":
		ep.ReturnType.F("Statistics").F("Frags").TypeStr = "map[string]int"
		ep.ReturnType.F("Private").F("Boosters").TypeStr = "map[string]*struct"
	case "wot_encyclopedia_vehicles":
		ep.ReturnType.F("DefaultProfile").F("Ammo").F("Stun").F("Duration").TypeStr = "[]int"
	case "wot_globalmap_events":
		ep.ReturnType.F("Fronts").TypeStr = "[]*struct"
	case "wot_globalmap_eventaccountinfo":
		// account_id OR clan_id is required
		for idx, param := range ep.Parameters {
			if param.Name != "accountId" {
				continue
			}
			param.TypeStr = toPointerType(param.TypeStr)
			param.Name = camelLowerToCamel(param.Name)
			ep.Options.Fields = append(ep.Options.Fields, param)
			ep.Parameters = append(ep.Parameters[:idx], ep.Parameters[idx+1:]...)
			break
		}
	case "wot_globalmap_fronts":
		ep.ReturnType.F("AvailableExtensions").TypeStr = "[]*struct"
	case "wot_globalmap_seasons":
		ep.ReturnType.F("Fronts").TypeStr = "[]*struct"
	case "wot_tanks_mastery":
		ep.ReturnType.F("Distribution").TypeStr = "map[string]map[string]int"
	case "wotb_tournaments_stages":
		ep.ReturnType.F("Groups").TypeStr = "[]*struct"
	case "wotb_encyclopedia_modules":
		ep.ReturnType.F("Engines").TypeStr = "[]*struct"
		ep.ReturnType.F("Guns").TypeStr = "[]*struct"
		ep.ReturnType.F("Guns").F("Shells").TypeStr = "[]*struct"
		ep.ReturnType.F("Suspensions").TypeStr = "[]*struct"
		ep.ReturnType.F("Turrets").TypeStr = "[]*struct"
	case "wgn_wgtv_tags":
		ep.ReturnType.F("Categories").TypeStr = "[]*struct"
		ep.ReturnType.F("Projects").TypeStr = "[]*struct"
		ep.ReturnType.F("Programs").TypeStr = "[]*struct"
	case "wot_clans_messageboard", "wotb_clanmessages_messages":
		ep.ReturnType.TypeStr = "map[string][]" + ep.ReturnType.TypeStr
	case "wotb_tournaments_teams":
		ep.ReturnType.F("Players").TypeStr = "[]*struct"
	case "wows_encyclopedia_ships":
		ep.ReturnType.TypeStr = "map[int]" + ep.ReturnType.TypeStr
	case "wows_ships_stats":
		ep.ReturnType.TypeStr = "map[int][]" + ep.ReturnType.TypeStr
	}
	if contains([]string{
		"wgn_account_list", "wgn_wgtv_videos",
		"wot_account_list", "wot_clans_list",
		"wot_clanratings_neighbors", "wot_clanratings_top", "wot_globalmap_clanbattles",
		"wot_globalmap_eventaccountratingneighbors", "wot_globalmap_eventaccountratings", "wot_globalmap_eventrating",
		"wot_globalmap_eventratingneighbors", "wot_globalmap_events", "wot_globalmap_fronts", "wot_globalmap_provinces",
		"wot_globalmap_seasonrating", "wot_globalmap_seasonratingneighbors", "wot_globalmap_seasons",
		"wot_stronghold_clanreserves",
		"wotb_account_list", "wotb_clans_list", "wotb_clanmessages_likes", "wotb_tournaments_list",
		"wotb_tournaments_teams", "wotb_tournaments_stages", "wotb_tournaments_standings", "wotb_tournaments_tables",

		"wotx_account_list", "wotx_clans_list", "wotx_account_xuidinfo", "wotx_account_psninfo",
		"wowp_account_list", "wowp_clans_list", "wowp_ratings_neighbors",
		"wowp_ratings_top",
		"wows_account_list", "wows_clans_list",
	}, ep.Id) {
		ep.ReturnType.TypeStr = "[]" + ep.ReturnType.TypeStr
	}
}
