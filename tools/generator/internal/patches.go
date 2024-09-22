package internal

import "slices"

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
	patchTimeTypes(ep.DataType)
	patchTimeTypes(ep.OptionsType)
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
		ep.DataType.F("Statistics").F("Frags").TypeStr = "map[string]int"
		ep.DataType.F("Private").F("Boosters").TypeStr = "map[string]*struct"
	case "wot_account_tanks":
		ep.DataType.TypeStr = "map[string][]" + ep.DataType.TypeStr
	case "wot_encyclopedia_vehicles":
		ep.DataType.F("Crew").TypeStr = "[]" + ep.DataType.F("Crew").TypeStr
		ep.DataType.F("DefaultProfile").F("Ammo").TypeStr = "[]*struct"
		ep.DataType.F("DefaultProfile").F("Ammo").F("Stun").F("Duration").TypeStr = "[]int"
		ep.DataType.F("NextTanks").TypeStr = "map[string]int"
		ep.DataType.F("PricesXp").TypeStr = "map[string]int"
		ep.DataType.TypeStr = "map[string]" + ep.DataType.TypeStr
	case "wot_globalmap_events":
		ep.DataType.F("Fronts").TypeStr = "[]*struct"
	case "wot_globalmap_eventaccountinfo":
		// account_id OR clan_id is required
		for idx, param := range ep.Parameters {
			if param.Name != "accountId" {
				continue
			}
			param.TypeStr = toPointerType(param.TypeStr)
			param.Name = camelLowerToCamel(param.Name)
			ep.OptionsType.Fields = append(ep.OptionsType.Fields, param)
			ep.Parameters = append(ep.Parameters[:idx], ep.Parameters[idx+1:]...)
			break
		}
	case "wot_globalmap_fronts":
		ep.DataType.F("AvailableExtensions").TypeStr = "[]*struct"
	case "wot_globalmap_seasons":
		ep.DataType.F("Fronts").TypeStr = "[]*struct"
	case "wot_tanks_mastery":
		ep.DataType.F("Distribution").TypeStr = "map[string]map[string]int"
	case "wotb_tournaments_stages":
		ep.DataType.F("Groups").TypeStr = "[]*struct"
	case "wotb_encyclopedia_modules":
		ep.DataType.F("Engines").TypeStr = "[]*struct"
		ep.DataType.F("Guns").TypeStr = "[]*struct"
		ep.DataType.F("Guns").F("Shells").TypeStr = "[]*struct"
		ep.DataType.F("Suspensions").TypeStr = "[]*struct"
		ep.DataType.F("Turrets").TypeStr = "[]*struct"
	case "wot_clans_messageboard":
		ep.DataType.TypeStr = "map[string][]" + ep.DataType.TypeStr
	case "wotb_tournaments_teams":
		ep.DataType.F("Players").TypeStr = "[]*struct"
	case "wows_encyclopedia_ships":
		ep.DataType.TypeStr = "map[int]" + ep.DataType.TypeStr
	case "wows_ships_stats":
		ep.DataType.TypeStr = "map[int][]" + ep.DataType.TypeStr
	case "wows_clans_info":
		ep.DataType.TypeStr = "map[int]" + ep.DataType.TypeStr
	case "wows_clans_accountinfo":
		ep.DataType.TypeStr = "map[int]" + ep.DataType.TypeStr
	case "wows_account_info":
		ep.DataType.TypeStr = "map[int]" + ep.DataType.TypeStr
	}

	if slices.Contains([]string{
		"wgn_account_list",
		"wot_account_list", "wot_clans_list",
		"wot_clanratings_neighbors", "wot_clanratings_top", "wot_globalmap_clanbattles",
		"wot_globalmap_eventaccountratingneighbors", "wot_globalmap_eventaccountratings", "wot_globalmap_eventrating",
		"wot_globalmap_eventratingneighbors", "wot_globalmap_events", "wot_globalmap_fronts", "wot_globalmap_provinces",
		"wot_globalmap_seasonrating", "wot_globalmap_seasonratingneighbors", "wot_globalmap_seasons",
		"wot_stronghold_clanreserves",
		"wotb_account_list", "wotb_clans_list", "wotb_tournaments_list",
		"wotb_tournaments_teams", "wotb_tournaments_stages", "wotb_tournaments_standings", "wotb_tournaments_tables",

		"wotx_account_list", "wotx_clans_list", "wotx_account_xuidinfo", "wotx_account_psninfo",
		"wowp_account_list", "wowp_clans_list", "wowp_ratings_neighbors",
		"wowp_ratings_top",
		"wows_account_list", "wows_clans_list",
	}, ep.Id) {
		ep.DataType.TypeStr = "[]" + ep.DataType.TypeStr
	}
	patchMeta(ep)
}

func patchMeta(ep *endpoint) {
	if slices.Contains([]string{
		"wot_account_list", "wot_account_info", "wot_account_tanks", "wot_account_achievements", "wot_stronghold_claninfo",
		"wot_globalmap_fronts", "wot_globalmap_provinces", "wot_globalmap_claninfo", "wot_globalmap_clanprovinces",
		"wot_globalmap_clanbattles", "wot_globalmap_seasonclaninfo", "wot_globalmap_seasonaccountinfo",
		"wot_globalmap_seasonratingneighbors", "wot_globalmap_eventclaninfo", "wot_globalmap_eventaccountinfo",
		"wot_globalmap_eventaccountinfo", "wot_encyclopedia_vehicleprofile", "wot_encyclopedia_achievements",
		"wot_encyclopedia_info", "wot_encyclopedia_arenas", "wot_encyclopedia_personalmissions", "wot_encyclopedia_boosters",
		"wot_encyclopedia_vehicleprofiles", "wot_encyclopedia_badges", "wot_encyclopedia_crewroles",
		"wot_encyclopedia_crewskills", "wot_clanratings_types", "wot_clanratings_dates", "wot_clanratings_clans",
		"wot_clanratings_neighbors", "wot_clanratings_top", "wot_tanks_stats", "wot_tanks_achievements",
		"wot_tanks_mastery", "wot_clans_info", "wot_clans_accountinfo", "wot_clans_memberhistory",
		"wotb_account_list", "wotb_account_info", "wotb_account_achievements", "wotb_account_tankstats",
		"wotb_encyclopedia_vehicles", "wotb_encyclopedia_vehicleprofile", "wotb_encyclopedia_provisions",
		"wotb_encyclopedia_achievements", "wotb_encyclopedia_vehicleprofiles", "wotb_clans_info",
		"wotb_clans_accountinfo", "wotb_tanks_stats", "wotb_tanks_achievements", "wotb_tournaments_info",
		"wotx_account_list", "wotx_account_info", "wotx_account_achievements", "wotx_clans_info",
		"wotx_clans_accountinfo", "wotx_encyclopedia_crewroles", "wotx_encyclopedia_vehicleupgrades",
		"wotx_encyclopedia_achievements", "wotx_encyclopedia_arenas", "wotx_encyclopedia_vehicleprofile",
		"wotx_tanks_stats", "wotx_tanks_achievements", "wowp_account_list", "wowp_account_info2",
		"wowp_account_achievements", "wowp_encyclopedia_planes", "wowp_encyclopedia_planeinfo",
		"wowp_encyclopedia_planemodules", "wowp_encyclopedia_planeupgrades", "wowp_encyclopedia_planespecification",
		"wowp_encyclopedia_achievements", "wowp_encyclopedia_info", "wowp_ratings_types", "wowp_ratings_accounts",
		"wowp_ratings_neighbors", "wowp_ratings_top", "wowp_ratings_dates", "wowp_planes_stats",
		"wowp_planes_achievements", "wowp_clans_info", "wowp_clans_accountinfo", "wows_account_list",
		"wows_encyclopedia_info", "wows_encyclopedia_achievements", "wows_encyclopedia_shipprofile",
		"wows_encyclopedia_accountlevels", "wows_encyclopedia_crews", "wows_encyclopedia_crewranks",
		"wows_encyclopedia_battletypes", "wows_encyclopedia_collections", "wows_encyclopedia_collectioncards",
		"wows_encyclopedia_battlearenas", "wows_clans_info", "wows_clans_accountinfo", "wows_clans_season",
		"wgn_account_list", "wgn_account_info",

		"wows_account_info", "wows_account_achievements", "wows_ships_stats",
		"wows_seasons_shipstats", "wows_seasons_accountinfo",

		"wot_clans_list", "wotb_clans_list", "wotx_clans_list",

		"wot_globalmap_seasons", "wot_globalmap_seasonrating", "wot_globalmap_events", "wot_globalmap_info",
		"wotb_tournaments_teams", "wotb_tournaments_stages", "wotx_encyclopedia_modules",

		"wotb_tournaments_list", "wowp_clans_list", "wows_clans_list",

		"wot_encyclopedia_vehicles", "wot_encyclopedia_provisions", "wot_encyclopedia_modules",
		"wotx_encyclopedia_vehicles", "wows_encyclopedia_ships", "wows_encyclopedia_modules",
		"wows_encyclopedia_consumables",
	}, ep.Id) {
		ep.MetaType = &goType{
			Name:    "",
			TypeStr: "*GenericMeta",
		}
	}
}
