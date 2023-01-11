// Auto generated file!

package wows

type EncyclopediaShipprofileOptions struct {
	// Main Battery ID. If the module is not indicated, module of basic configuration is used.
	ArtilleryId *int `json:"artillery_id,omitempty"`
	// Dive bombers' ID. If the module is not indicated, module of basic configuration is used.
	DiveBomberId *int `json:"dive_bomber_id,omitempty"`
	// Engine ID. If the module is not indicated, module of basic configuration is used.
	EngineId *int `json:"engine_id,omitempty"`
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use "-" in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string `json:"fields,omitempty"`
	// Fighters' ID. If the module is not indicated, module of basic configuration is used.
	FighterId *int `json:"fighter_id,omitempty"`
	// ID of Gun Fire Control System. If the module is not indicated, module of basic configuration is used.
	FireControlId *int `json:"fire_control_id,omitempty"`
	// ID of Flight Control System. If the module is not indicated, module of basic configuration is used.
	FlightControlId *int `json:"flight_control_id,omitempty"`
	// Hull ID. If the module is not indicated, module of basic configuration is used.
	HullId *int `json:"hull_id,omitempty"`
	// Localization language. Default is "en". Valid values:
	//
	// "cs" - Čeština
	// "de" - Deutsch
	// "en" - English (by default)
	// "es" - Español
	// "fr" - Français
	// "ja" - 日本語
	// "pl" - Polski
	// "ru" - Русский
	// "th" - ไทย
	// "zh-tw" - 繁體中文
	// "tr" - Türkçe
	// "zh-cn" - 中文
	// "pt-br" - Português do Brasil
	// "es-mx" - Español (México)
	Language *string `json:"language,omitempty"`
	// Torpedo bombers' ID. If the module is not indicated, module of basic configuration is used.
	TorpedoBomberId *int `json:"torpedo_bomber_id,omitempty"`
	// Torpedo tubes' ID. If the module is not indicated, module of basic configuration is used.
	TorpedoesId *int `json:"torpedoes_id,omitempty"`
}

type EncyclopediaShipprofile struct {
	// Anti-aircraft guns. If the module is absent on the ship, field value is null.
	AntiAircraft *struct {
		// Anti-aircraft effectiveness
		Defense *int `json:"defense,omitempty"`
		// Gun slots
		Slots *struct {
			// Average damage per second
			AvgDamage *int `json:"avg_damage,omitempty"`
			// Caliber
			Caliber *int `json:"caliber,omitempty"`
			// Firing range (km)
			Distance *float32 `json:"distance,omitempty"`
			// Number of guns
			Guns *int `json:"guns,omitempty"`
			// Gun name
			Name *string `json:"name,omitempty"`
		} `json:"slots,omitempty"`
	} `json:"anti_aircraft,omitempty"`
	// Survivability of basic configuration
	Armour *struct {
		// Gun Casemate
		Casemate *struct {
			// Maximum value
			Max *int `json:"max,omitempty"`
			// Minimum value
			Min *int `json:"min,omitempty"`
		} `json:"casemate,omitempty"`
		// Citadel
		Citadel *struct {
			// Maximum value
			Max *int `json:"max,omitempty"`
			// Minimum value
			Min *int `json:"min,omitempty"`
		} `json:"citadel,omitempty"`
		// Armored Deck
		Deck *struct {
			// Maximum value
			Max *int `json:"max,omitempty"`
			// Minimum value
			Min *int `json:"min,omitempty"`
		} `json:"deck,omitempty"`
		// Forward and After Ends
		Extremities *struct {
			// Maximum value
			Max *int `json:"max,omitempty"`
			// Minimum value
			Min *int `json:"min,omitempty"`
		} `json:"extremities,omitempty"`
		// Torpedo Protection. Damage Reduction (%)
		FloodDamage *int `json:"flood_damage,omitempty"`
		// Torpedo Protection. Flooding Risk Reduction (%)
		FloodProb *int `json:"flood_prob,omitempty"`
		// Hit points
		Health *int `json:"health,omitempty"`
		// Armor
		Range *struct {
			// Maximum value
			Max *int `json:"max,omitempty"`
			// Minimum value
			Min *int `json:"min,omitempty"`
		} `json:"range,omitempty"`
		// Armor protection (%)
		Total *int `json:"total,omitempty"`
	} `json:"armour,omitempty"`
	// Main battery. If the module is absent on the ship, field value is null.
	Artillery *struct {
		// Gun ID
		ArtilleryId *int `json:"artillery_id,omitempty"`
		// Gun string ID
		ArtilleryIdStr *string `json:"artillery_id_str,omitempty"`
		// Firing range
		Distance *float32 `json:"distance,omitempty"`
		// Rate of fire (rounds / min)
		GunRate *float32 `json:"gun_rate,omitempty"`
		// Maximum dispersion (m)
		MaxDispersion *int `json:"max_dispersion,omitempty"`
		// 180 Degree Turn Time (sec)
		RotationTime *float32 `json:"rotation_time,omitempty"`
		// Shells
		Shells *struct {
			// Shell weight
			BulletMass *int `json:"bullet_mass,omitempty"`
			// Shell speed
			BulletSpeed *int `json:"bullet_speed,omitempty"`
			// Chance of Fire on target caused by shell (%)
			BurnProbability *float32 `json:"burn_probability,omitempty"`
			// Maximum Damage
			Damage *int `json:"damage,omitempty"`
			// Shell name
			Name *string `json:"name,omitempty"`
			// Shell type
			Type *string `json:"type,omitempty"`
		} `json:"shells,omitempty"`
		// Main battery reload time (s)
		ShotDelay *float32 `json:"shot_delay,omitempty"`
		// Main gun slots
		Slots *struct {
			// Number of barrels per slot
			Barrels *int `json:"barrels,omitempty"`
			// Number of main turrets
			Guns *int `json:"guns,omitempty"`
			// Name
			Name *string `json:"name,omitempty"`
		} `json:"slots,omitempty"`
	} `json:"artillery,omitempty"`
	// Secondary armament. If the module is absent on the ship, field value is null.
	Atbas *struct {
		// Firing range
		Distance *float32 `json:"distance,omitempty"`
		// Main gun slots
		Slots *struct {
			// Shell weight
			BulletMass *int `json:"bullet_mass,omitempty"`
			// Shell speed
			BulletSpeed *int `json:"bullet_speed,omitempty"`
			// Chance of Fire on target caused by shell (%)
			BurnProbability *float32 `json:"burn_probability,omitempty"`
			// Maximum Damage
			Damage *int `json:"damage,omitempty"`
			// Rate of fire (rounds / min)
			GunRate *float32 `json:"gun_rate,omitempty"`
			// Shell name
			Name *string `json:"name,omitempty"`
			// Reload time (s)
			ShotDelay *float32 `json:"shot_delay,omitempty"`
			// Shell type
			Type *string `json:"type,omitempty"`
		} `json:"slots,omitempty"`
	} `json:"atbas,omitempty"`
	// Maximum battle tier for a random battle
	BattleLevelRangeMax *int `json:"battle_level_range_max,omitempty"`
	// Minimum battle tier for a random battle
	BattleLevelRangeMin *int `json:"battle_level_range_min,omitempty"`
	// Concealment of basic configuration
	Concealment *struct {
		// Air Detectability Range
		DetectDistanceByPlane *float32 `json:"detect_distance_by_plane,omitempty"`
		// Surface Detectability Range
		DetectDistanceByShip *float32 `json:"detect_distance_by_ship,omitempty"`
		// Concealment (%)
		Total *int `json:"total,omitempty"`
	} `json:"concealment,omitempty"`
	// Dive bombers. If the module is absent on the ship, field value is null.
	DiveBomber *struct {
		// Accuracy
		Accuracy *struct {
			// Maximum value
			Max *float32 `json:"max,omitempty"`
			// Minimum value
			Min *float32 `json:"min,omitempty"`
		} `json:"accuracy,omitempty"`
		// Bomb weight
		BombBulletMass *int `json:"bomb_bullet_mass,omitempty"`
		// Chance of Fire on target caused by bomb (%)
		BombBurnProbability *float32 `json:"bomb_burn_probability,omitempty"`
		// Maximum bomb damage
		BombDamage *int `json:"bomb_damage,omitempty"`
		// Bomb name
		BombName *string `json:"bomb_name,omitempty"`
		// Number of aircraft in a squadron
		CountInSquadron *struct {
			// Maximum value
			Max *int `json:"max,omitempty"`
			// Minimum value
			Min *int `json:"min,omitempty"`
		} `json:"count_in_squadron,omitempty"`
		// Cruise Speed (knots)
		CruiseSpeed *int `json:"cruise_speed,omitempty"`
		// Dive bombers' ID
		DiveBomberId *int `json:"dive_bomber_id,omitempty"`
		// Dive bombers string ID
		DiveBomberIdStr *string `json:"dive_bomber_id_str,omitempty"`
		// Average damage per rear gunner of a dive bomber per second
		GunnerDamage *int `json:"gunner_damage,omitempty"`
		// Maximum Bomb Damage
		MaxDamage *int `json:"max_damage,omitempty"`
		// Survivability
		MaxHealth *int `json:"max_health,omitempty"`
		// Name of squadron
		Name *string `json:"name,omitempty"`
		// Dive bomber tier
		PlaneLevel *int `json:"plane_level,omitempty"`
		// Time of preparation for takeoff (sec)
		PrepareTime *int `json:"prepare_time,omitempty"`
		// Number of squadrons
		Squadrons *int `json:"squadrons,omitempty"`
	} `json:"dive_bomber,omitempty"`
	// Engine
	Engine *struct {
		// Engine ID
		EngineId *int `json:"engine_id,omitempty"`
		// Engine string ID
		EngineIdStr *string `json:"engine_id_str,omitempty"`
		// Top Speed (knots)
		MaxSpeed *float32 `json:"max_speed,omitempty"`
	} `json:"engine,omitempty"`
	// Fighters. If the module is absent on the ship, field value is null.
	Fighters *struct {
		// Average damage caused per second
		AvgDamage *int `json:"avg_damage,omitempty"`
		// Number of aircraft in a squadron
		CountInSquadron *struct {
			// Maximum value
			Max *int `json:"max,omitempty"`
			// Minimum value
			Min *int `json:"min,omitempty"`
		} `json:"count_in_squadron,omitempty"`
		// Cruise Speed (knots)
		CruiseSpeed *int `json:"cruise_speed,omitempty"`
		// Fighters' ID
		FightersId *int `json:"fighters_id,omitempty"`
		// Fighters string ID
		FightersIdStr *string `json:"fighters_id_str,omitempty"`
		// Average damage per gunner of a fighter per second
		GunnerDamage *int `json:"gunner_damage,omitempty"`
		// Ammunition
		MaxAmmo *int `json:"max_ammo,omitempty"`
		// Survivability
		MaxHealth *int `json:"max_health,omitempty"`
		// Name of squadron
		Name *string `json:"name,omitempty"`
		// Fighter tier
		PlaneLevel *int `json:"plane_level,omitempty"`
		// Time of preparation for takeoff (sec)
		PrepareTime *int `json:"prepare_time,omitempty"`
		// Number of squadrons
		Squadrons *int `json:"squadrons,omitempty"`
	} `json:"fighters,omitempty"`
	// Gun Fire Control System. If the module is absent on the ship, field value is null.
	FireControl *struct {
		// Firing range
		Distance *float32 `json:"distance,omitempty"`
		// Firing Range extension (%)
		DistanceIncrease *int `json:"distance_increase,omitempty"`
		// ID of Gun Fire Control System
		FireControlId *int `json:"fire_control_id,omitempty"`
		// String ID of Gun Fire Control System
		FireControlIdStr *string `json:"fire_control_id_str,omitempty"`
	} `json:"fire_control,omitempty"`
	// Flight Control. If the module is absent on the ship, field value is null.
	FlightControl *struct {
		// Number of bomber squadrons
		BomberSquadrons *int `json:"bomber_squadrons,omitempty"`
		// Number of fighter squadrons
		FighterSquadrons *int `json:"fighter_squadrons,omitempty"`
		// Flight Control ID
		FlightControlId *int `json:"flight_control_id,omitempty"`
		// Flight Control string ID
		FlightControlIdStr *string `json:"flight_control_id_str,omitempty"`
		// Number of torpedo bomber squadrons
		TorpedoSquadrons *int `json:"torpedo_squadrons,omitempty"`
	} `json:"flight_control,omitempty"`
	// Hull
	Hull *struct {
		// AA Mounts
		AntiAircraftBarrels *int `json:"anti_aircraft_barrels,omitempty"`
		// Number of main turrets
		ArtilleryBarrels *int `json:"artillery_barrels,omitempty"`
		// Secondary gun turrets
		AtbaBarrels *int `json:"atba_barrels,omitempty"`
		// Hit points
		Health *int `json:"health,omitempty"`
		// Hull ID
		HullId *int `json:"hull_id,omitempty"`
		// Hull string ID
		HullIdStr *string `json:"hull_id_str,omitempty"`
		// Hangar capacity
		PlanesAmount *int `json:"planes_amount,omitempty"`
		// Armor (mm)
		Range *struct {
			// Maximum value
			Max *int `json:"max,omitempty"`
			// Minimum value
			Min *int `json:"min,omitempty"`
		} `json:"range,omitempty"`
		// Torpedo tubes
		TorpedoesBarrels *int `json:"torpedoes_barrels,omitempty"`
	} `json:"hull,omitempty"`
	// Maneuverability of basic configuration
	Mobility *struct {
		// Top Speed (knots)
		MaxSpeed *float32 `json:"max_speed,omitempty"`
		// Rudder Shift Time (sec)
		RudderTime *float32 `json:"rudder_time,omitempty"`
		// Maneuverability (%)
		Total *int `json:"total,omitempty"`
		// Turning Circle Radius (m)
		TurningRadius *int `json:"turning_radius,omitempty"`
	} `json:"mobility,omitempty"`
	// Ship ID
	ShipId *int `json:"ship_id,omitempty"`
	// Torpedo bombers. If the module is absent on the ship, field value is null.
	TorpedoBomber *struct {
		// Number of aircraft in a squadron
		CountInSquadron *struct {
			// Maximum value
			Max *int `json:"max,omitempty"`
			// Minimum value
			Min *int `json:"min,omitempty"`
		} `json:"count_in_squadron,omitempty"`
		// Cruise Speed (knots)
		CruiseSpeed *int `json:"cruise_speed,omitempty"`
		// Average damage per rear gunner of a torpedo bomber per second
		GunnerDamage *int `json:"gunner_damage,omitempty"`
		// Maximum Bomb Damage
		MaxDamage *int `json:"max_damage,omitempty"`
		// Survivability
		MaxHealth *int `json:"max_health,omitempty"`
		// Name of squadron
		Name *string `json:"name,omitempty"`
		// Torpedo bomber tier
		PlaneLevel *int `json:"plane_level,omitempty"`
		// Time of preparation for takeoff (sec)
		PrepareTime *int `json:"prepare_time,omitempty"`
		// Number of squadrons
		Squadrons *int `json:"squadrons,omitempty"`
		// Torpedo bombers' ID
		TorpedoBomberId *int `json:"torpedo_bomber_id,omitempty"`
		// Torpedo bombers string ID
		TorpedoBomberIdStr *string `json:"torpedo_bomber_id_str,omitempty"`
		// Maximum torpedo damage
		TorpedoDamage *int `json:"torpedo_damage,omitempty"`
		// Firing range
		TorpedoDistance *float32 `json:"torpedo_distance,omitempty"`
		// Top Speed (knots)
		TorpedoMaxSpeed *int `json:"torpedo_max_speed,omitempty"`
		// Torpedo name
		TorpedoName *string `json:"torpedo_name,omitempty"`
	} `json:"torpedo_bomber,omitempty"`
	// Torpedo tubes. If the module is absent on the ship, field value is null.
	Torpedoes *struct {
		// Firing range
		Distance *float32 `json:"distance,omitempty"`
		// Maximum Damage
		MaxDamage *int `json:"max_damage,omitempty"`
		// Reload Time (sec)
		ReloadTime *int `json:"reload_time,omitempty"`
		// 180 Degree Turn Time (sec)
		RotationTime *float32 `json:"rotation_time,omitempty"`
		// Slots for Torpedo tubes
		Slots *struct {
			// Torpedo tubes per slot
			Barrels *int `json:"barrels,omitempty"`
			// Caliber
			Caliber *int `json:"caliber,omitempty"`
			// Torpedo tubes
			Guns *int `json:"guns,omitempty"`
			// Name
			Name *string `json:"name,omitempty"`
		} `json:"slots,omitempty"`
		// Torpedo
		TorpedoName *string `json:"torpedo_name,omitempty"`
		// Torpedo Speed (knots)
		TorpedoSpeed *int `json:"torpedo_speed,omitempty"`
		// Torpedo tubes' ID
		TorpedoesId *int `json:"torpedoes_id,omitempty"`
		// Torpedo tubes string ID
		TorpedoesIdStr *string `json:"torpedoes_id_str,omitempty"`
		// Torpedo range (km)
		VisibilityDist *float32 `json:"visibility_dist,omitempty"`
	} `json:"torpedoes,omitempty"`
	// Armament effectiveness of basic configuration
	Weaponry *struct {
		// Aircraft (%)
		Aircraft *int `json:"aircraft,omitempty"`
		// AA Guns (%)
		AntiAircraft *int `json:"anti_aircraft,omitempty"`
		// Artillery (%)
		Artillery *int `json:"artillery,omitempty"`
		// Torpedoes (%)
		Torpedoes *int `json:"torpedoes,omitempty"`
	} `json:"weaponry,omitempty"`
}
