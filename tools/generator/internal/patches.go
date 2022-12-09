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
	case "wot_encyclopedia_vehicles":
		ep.DataType.F("DefaultProfile").F("Ammo").F("Stun").F("Duration").TypeStr = "[]int"
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
	case "wgn_wgtv_tags":
		ep.DataType.F("Categories").TypeStr = "[]*struct"
		ep.DataType.F("Projects").TypeStr = "[]*struct"
		ep.DataType.F("Programs").TypeStr = "[]*struct"
	case "wot_clans_messageboard", "wotb_clanmessages_messages":
		ep.DataType.TypeStr = "map[string][]" + ep.DataType.TypeStr
	case "wotb_tournaments_teams":
		ep.DataType.F("Players").TypeStr = "[]*struct"
	case "wows_encyclopedia_ships":
		ep.DataType.TypeStr = "map[int]" + ep.DataType.TypeStr
	case "wows_ships_stats":
		ep.DataType.TypeStr = "map[int][]" + ep.DataType.TypeStr
	}

	if contains([]string{
		"wows_encyclopedia_ships",
	}, ep.Id) {
		ep.MetaType = &goType{
			Name:    ep.Name + "Meta",
			TypeStr: "struct",
			Fields: []*goType{
				{
					Name:    "Count",
					TypeStr: "int",
					JsonTag: "count",
					Imports: make(map[string]struct{}),
				},
				{
					Name:    "Limit",
					TypeStr: "int",
					JsonTag: "limit",
					Imports: make(map[string]struct{}),
				},
				{
					Name:    "Page",
					TypeStr: "int",
					JsonTag: "page",
					Imports: make(map[string]struct{}),
				},
				{
					Name:    "PageTotal",
					TypeStr: "int",
					JsonTag: "page_total",
					Imports: make(map[string]struct{}),
				},
				{
					Name:    "Total",
					TypeStr: "int",
					JsonTag: "total",
					Imports: make(map[string]struct{}),
				},
			},
		}
	}

	if contains([]string{
		"wows_account_achievements", "wows_seasons_accountinfo", "wows_seasons_shipstats", "wows_ships_stats",
	}, ep.Id) {
		ep.MetaType = &goType{
			Name:    ep.Name + "Meta",
			TypeStr: "struct",
			Fields: []*goType{
				{
					Name:        "Hidden",
					TypeStr:     "[]int",
					Description: "Hidden profiles.",
					JsonTag:     "hidden",
					Imports:     make(map[string]struct{}),
				},
			},
		}
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
		ep.DataType.TypeStr = "[]" + ep.DataType.TypeStr
	}
}
