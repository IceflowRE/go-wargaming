package wows

type EncyclopediaShips struct {
	// List of compatible Modifications
	Upgrades []int `json:"upgrades,omitempty"`
	// Type of ship
	Type_ string `json:"type,omitempty"`
	// Tier
	Tier int `json:"tier,omitempty"`
	// Ship string ID
	ShipIdStr string `json:"ship_id_str,omitempty"`
	// Ship ID
	ShipId int `json:"ship_id,omitempty"`
	// Cost in gold
	PriceGold int `json:"price_gold,omitempty"`
	// Cost in credits
	PriceCredit int `json:"price_credit,omitempty"`
	// List of ships available for research in form of pairs
	NextShips map[string]string `json:"next_ships,omitempty"`
	// Nation
	Nation string `json:"nation,omitempty"`
	// Ship name
	Name string `json:"name,omitempty"`
	// Module research information
	ModulesTree struct {
		// Module type
		Type_ string `json:"type,omitempty"`
		// Research cost
		PriceXp int `json:"price_xp,omitempty"`
		// Cost in credits
		PriceCredit int `json:"price_credit,omitempty"`
		// List of vehicle IDs available after research of the module
		NextShips []int `json:"next_ships,omitempty"`
		// List of module IDs available after research of the module
		NextModules []int `json:"next_modules,omitempty"`
		// Module name
		Name string `json:"name,omitempty"`
		// Module string ID
		ModuleIdStr string `json:"module_id_str,omitempty"`
		// Module ID
		ModuleId int `json:"module_id,omitempty"`
		// Indicates if the module is basic
		IsDefault bool `json:"is_default,omitempty"`
	} `json:"modules_tree,omitempty"`
	// List of compatible modules
	Modules struct {
		// Torpedo tubes. If the module is absent on the ship, field value is null.
		Torpedoes []int `json:"torpedoes,omitempty"`
		// Torpedo bombers. If the module is absent on the ship, field value is null.
		TorpedoBomber []int `json:"torpedo_bomber,omitempty"`
		// Hull
		Hull []int `json:"hull,omitempty"`
		// Flight Control. If the module is absent on the ship, field value is null.
		FlightControl []int `json:"flight_control,omitempty"`
		// Gun Fire Control System. If the module is absent on the ship, field value is null.
		FireControl []int `json:"fire_control,omitempty"`
		// Fighters. If the module is absent on the ship, field value is null.
		Fighter []int `json:"fighter,omitempty"`
		// Engines
		Engine []int `json:"engine,omitempty"`
		// Dive bombers. If the module is absent on the ship, field value is null.
		DiveBomber []int `json:"dive_bomber,omitempty"`
		// Main battery. If the module is absent on the ship, field value is null.
		Artillery []int `json:"artillery,omitempty"`
	} `json:"modules,omitempty"`
	// Number of slots for upgrades
	ModSlots int `json:"mod_slots,omitempty"`
	// Indicates if the ship is on a special offer
	IsSpecial bool `json:"is_special,omitempty"`
	// Indicates if the ship is Premium ship
	IsPremium bool `json:"is_premium,omitempty"`
	// Image of a ship
	Images struct {
		// URL to 214 x 126 px image of ship
		Small string `json:"small,omitempty"`
		// URL to 435 x 256 px image of ship
		Medium string `json:"medium,omitempty"`
		// URL to 870 x 512 px image of ship
		Large string `json:"large,omitempty"`
		// URL to 186 x 48 px outline image of ship
		Contour string `json:"contour,omitempty"`
	} `json:"images,omitempty"`
	// Indicates that ship characteristics are used for illustration, and may be changed.
	HasDemoProfile bool `json:"has_demo_profile,omitempty"`
	// Ship description
	Description string `json:"description,omitempty"`
	// Parameters of basic configuration
	DefaultProfile struct {
		// Armament effectiveness of basic configuration
		Weaponry struct {
			// Torpedoes (%)
			Torpedoes int `json:"torpedoes,omitempty"`
			// Artillery (%)
			Artillery int `json:"artillery,omitempty"`
			// AA Guns (%)
			AntiAircraft int `json:"anti_aircraft,omitempty"`
			// Aircraft (%)
			Aircraft int `json:"aircraft,omitempty"`
		} `json:"weaponry,omitempty"`
		// Torpedo tubes. If the module is absent on the ship, field value is null.
		Torpedoes struct {
			// Torpedo range (km)
			VisibilityDist float32 `json:"visibility_dist,omitempty"`
			// Torpedo tubes string ID
			TorpedoesIdStr string `json:"torpedoes_id_str,omitempty"`
			// Torpedo tubes' ID
			TorpedoesId int `json:"torpedoes_id,omitempty"`
			// Torpedo Speed (knots)
			TorpedoSpeed int `json:"torpedo_speed,omitempty"`
			// Torpedo
			TorpedoName string `json:"torpedo_name,omitempty"`
			// Slots for Torpedo tubes
			Slots struct {
				// Name
				Name string `json:"name,omitempty"`
				// Torpedo tubes
				Guns int `json:"guns,omitempty"`
				// Caliber
				Caliber int `json:"caliber,omitempty"`
				// Torpedo tubes per slot
				Barrels int `json:"barrels,omitempty"`
			} `json:"slots,omitempty"`
			// 180 Degree Turn Time (sec)
			RotationTime float32 `json:"rotation_time,omitempty"`
			// Reload Time (sec)
			ReloadTime int `json:"reload_time,omitempty"`
			// Maximum Damage
			MaxDamage int `json:"max_damage,omitempty"`
			// Firing range
			Distance float32 `json:"distance,omitempty"`
		} `json:"torpedoes,omitempty"`
		// Torpedo bombers. If the module is absent on the ship, field value is null.
		TorpedoBomber struct {
			// Torpedo name
			TorpedoName string `json:"torpedo_name,omitempty"`
			// Top Speed (knots)
			TorpedoMaxSpeed int `json:"torpedo_max_speed,omitempty"`
			// Firing range
			TorpedoDistance float32 `json:"torpedo_distance,omitempty"`
			// Maximum torpedo damage
			TorpedoDamage int `json:"torpedo_damage,omitempty"`
			// Torpedo bombers string ID
			TorpedoBomberIdStr string `json:"torpedo_bomber_id_str,omitempty"`
			// Torpedo bombers' ID
			TorpedoBomberId int `json:"torpedo_bomber_id,omitempty"`
			// Number of squadrons
			Squadrons int `json:"squadrons,omitempty"`
			// Time of preparation for takeoff (sec)
			PrepareTime int `json:"prepare_time,omitempty"`
			// Torpedo bomber tier
			PlaneLevel int `json:"plane_level,omitempty"`
			// Name of squadron
			Name string `json:"name,omitempty"`
			// Survivability
			MaxHealth int `json:"max_health,omitempty"`
			// Maximum Bomb Damage
			MaxDamage int `json:"max_damage,omitempty"`
			// Average damage per rear gunner of a torpedo bomber per second
			GunnerDamage int `json:"gunner_damage,omitempty"`
			// Cruise Speed (knots)
			CruiseSpeed int `json:"cruise_speed,omitempty"`
			// Number of aircraft in a squadron
			CountInSquadron struct {
				// Minimum value
				Min int `json:"min,omitempty"`
				// Maximum value
				Max int `json:"max,omitempty"`
			} `json:"count_in_squadron,omitempty"`
		} `json:"torpedo_bomber,omitempty"`
		// Maneuverability of basic configuration
		Mobility struct {
			// Turning Circle Radius (m)
			TurningRadius int `json:"turning_radius,omitempty"`
			// Maneuverability (%)
			Total int `json:"total,omitempty"`
			// Rudder Shift Time (sec)
			RudderTime float32 `json:"rudder_time,omitempty"`
			// Top Speed (knots)
			MaxSpeed float32 `json:"max_speed,omitempty"`
		} `json:"mobility,omitempty"`
		// Hull
		Hull struct {
			// Torpedo tubes
			TorpedoesBarrels int `json:"torpedoes_barrels,omitempty"`
			// Armor (mm)
			Range struct {
				// Minimum value
				Min int `json:"min,omitempty"`
				// Maximum value
				Max int `json:"max,omitempty"`
			} `json:"range,omitempty"`
			// Hangar capacity
			PlanesAmount int `json:"planes_amount,omitempty"`
			// Hull string ID
			HullIdStr string `json:"hull_id_str,omitempty"`
			// Hull ID
			HullId int `json:"hull_id,omitempty"`
			// Hit points
			Health int `json:"health,omitempty"`
			// Secondary gun turrets
			AtbaBarrels int `json:"atba_barrels,omitempty"`
			// Number of main turrets
			ArtilleryBarrels int `json:"artillery_barrels,omitempty"`
			// AA Mounts
			AntiAircraftBarrels int `json:"anti_aircraft_barrels,omitempty"`
		} `json:"hull,omitempty"`
		// Flight Control. If the module is absent on the ship, field value is null.
		FlightControl struct {
			// Number of torpedo bomber squadrons
			TorpedoSquadrons int `json:"torpedo_squadrons,omitempty"`
			// Flight Control string ID
			FlightControlIdStr string `json:"flight_control_id_str,omitempty"`
			// Flight Control ID
			FlightControlId int `json:"flight_control_id,omitempty"`
			// Number of fighter squadrons
			FighterSquadrons int `json:"fighter_squadrons,omitempty"`
			// Number of bomber squadrons
			BomberSquadrons int `json:"bomber_squadrons,omitempty"`
		} `json:"flight_control,omitempty"`
		// Gun Fire Control System. If the module is absent on the ship, field value is null.
		FireControl struct {
			// String ID of Gun Fire Control System
			FireControlIdStr string `json:"fire_control_id_str,omitempty"`
			// ID of Gun Fire Control System
			FireControlId int `json:"fire_control_id,omitempty"`
			// Firing Range extension (%)
			DistanceIncrease int `json:"distance_increase,omitempty"`
			// Firing range
			Distance float32 `json:"distance,omitempty"`
		} `json:"fire_control,omitempty"`
		// Fighters. If the module is absent on the ship, field value is null.
		Fighters struct {
			// Number of squadrons
			Squadrons int `json:"squadrons,omitempty"`
			// Time of preparation for takeoff (sec)
			PrepareTime int `json:"prepare_time,omitempty"`
			// Fighter tier
			PlaneLevel int `json:"plane_level,omitempty"`
			// Name of squadron
			Name string `json:"name,omitempty"`
			// Survivability
			MaxHealth int `json:"max_health,omitempty"`
			// Ammunition
			MaxAmmo int `json:"max_ammo,omitempty"`
			// Average damage per gunner of a fighter per second
			GunnerDamage int `json:"gunner_damage,omitempty"`
			// Fighters string ID
			FightersIdStr string `json:"fighters_id_str,omitempty"`
			// Fighters' ID
			FightersId int `json:"fighters_id,omitempty"`
			// Cruise Speed (knots)
			CruiseSpeed int `json:"cruise_speed,omitempty"`
			// Number of aircraft in a squadron
			CountInSquadron struct {
				// Minimum value
				Min int `json:"min,omitempty"`
				// Maximum value
				Max int `json:"max,omitempty"`
			} `json:"count_in_squadron,omitempty"`
			// Average damage caused per second
			AvgDamage int `json:"avg_damage,omitempty"`
		} `json:"fighters,omitempty"`
		// Engine
		Engine struct {
			// Top Speed (knots)
			MaxSpeed float32 `json:"max_speed,omitempty"`
			// Engine string ID
			EngineIdStr string `json:"engine_id_str,omitempty"`
			// Engine ID
			EngineId int `json:"engine_id,omitempty"`
		} `json:"engine,omitempty"`
		// Dive bombers. If the module is absent on the ship, field value is null.
		DiveBomber struct {
			// Number of squadrons
			Squadrons int `json:"squadrons,omitempty"`
			// Time of preparation for takeoff (sec)
			PrepareTime int `json:"prepare_time,omitempty"`
			// Dive bomber tier
			PlaneLevel int `json:"plane_level,omitempty"`
			// Name of squadron
			Name string `json:"name,omitempty"`
			// Survivability
			MaxHealth int `json:"max_health,omitempty"`
			// Maximum Bomb Damage
			MaxDamage int `json:"max_damage,omitempty"`
			// Average damage per rear gunner of a dive bomber per second
			GunnerDamage int `json:"gunner_damage,omitempty"`
			// Dive bombers string ID
			DiveBomberIdStr string `json:"dive_bomber_id_str,omitempty"`
			// Dive bombers' ID
			DiveBomberId int `json:"dive_bomber_id,omitempty"`
			// Cruise Speed (knots)
			CruiseSpeed int `json:"cruise_speed,omitempty"`
			// Number of aircraft in a squadron
			CountInSquadron struct {
				// Minimum value
				Min int `json:"min,omitempty"`
				// Maximum value
				Max int `json:"max,omitempty"`
			} `json:"count_in_squadron,omitempty"`
			// Bomb name
			BombName string `json:"bomb_name,omitempty"`
			// Maximum bomb damage
			BombDamage int `json:"bomb_damage,omitempty"`
			// Chance of Fire on target caused by bomb (%)
			BombBurnProbability float32 `json:"bomb_burn_probability,omitempty"`
			// Bomb weight
			BombBulletMass int `json:"bomb_bullet_mass,omitempty"`
			// Accuracy
			Accuracy struct {
				// Minimum value
				Min float32 `json:"min,omitempty"`
				// Maximum value
				Max float32 `json:"max,omitempty"`
			} `json:"accuracy,omitempty"`
		} `json:"dive_bomber,omitempty"`
		// Concealment of basic configuration
		Concealment struct {
			// Concealment (%)
			Total int `json:"total,omitempty"`
			// Surface Detectability Range
			DetectDistanceByShip float32 `json:"detect_distance_by_ship,omitempty"`
			// Air Detectability Range
			DetectDistanceByPlane float32 `json:"detect_distance_by_plane,omitempty"`
		} `json:"concealment,omitempty"`
		// Minimum battle tier for a random battle
		BattleLevelRangeMin int `json:"battle_level_range_min,omitempty"`
		// Maximum battle tier for a random battle
		BattleLevelRangeMax int `json:"battle_level_range_max,omitempty"`
		// Secondary armament. If the module is absent on the ship, field value is null.
		Atbas struct {
			// Main gun slots
			Slots struct {
				// Shell type
				Type_ string `json:"type,omitempty"`
				// Reload time (s)
				ShotDelay float32 `json:"shot_delay,omitempty"`
				// Shell name
				Name string `json:"name,omitempty"`
				// Rate of fire (rounds / min)
				GunRate float32 `json:"gun_rate,omitempty"`
				// Maximum Damage
				Damage int `json:"damage,omitempty"`
				// Chance of Fire on target caused by shell (%)
				BurnProbability float32 `json:"burn_probability,omitempty"`
				// Shell speed
				BulletSpeed int `json:"bullet_speed,omitempty"`
				// Shell weight
				BulletMass int `json:"bullet_mass,omitempty"`
			} `json:"slots,omitempty"`
			// Firing range
			Distance float32 `json:"distance,omitempty"`
		} `json:"atbas,omitempty"`
		// Main battery. If the module is absent on the ship, field value is null.
		Artillery struct {
			// Main gun slots
			Slots struct {
				// Name
				Name string `json:"name,omitempty"`
				// Number of main turrets
				Guns int `json:"guns,omitempty"`
				// Number of barrels per slot
				Barrels int `json:"barrels,omitempty"`
			} `json:"slots,omitempty"`
			// Main battery reload time (s)
			ShotDelay float32 `json:"shot_delay,omitempty"`
			// Shells
			Shells struct {
				// Shell type
				Type_ string `json:"type,omitempty"`
				// Shell name
				Name string `json:"name,omitempty"`
				// Maximum Damage
				Damage int `json:"damage,omitempty"`
				// Chance of Fire on target caused by shell (%)
				BurnProbability float32 `json:"burn_probability,omitempty"`
				// Shell speed
				BulletSpeed int `json:"bullet_speed,omitempty"`
				// Shell weight
				BulletMass int `json:"bullet_mass,omitempty"`
			} `json:"shells,omitempty"`
			// 180 Degree Turn Time (sec)
			RotationTime float32 `json:"rotation_time,omitempty"`
			// Maximum dispersion (m)
			MaxDispersion int `json:"max_dispersion,omitempty"`
			// Rate of fire (rounds / min)
			GunRate float32 `json:"gun_rate,omitempty"`
			// Firing range
			Distance float32 `json:"distance,omitempty"`
			// Gun string ID
			ArtilleryIdStr string `json:"artillery_id_str,omitempty"`
			// Gun ID
			ArtilleryId int `json:"artillery_id,omitempty"`
		} `json:"artillery,omitempty"`
		// Survivability of basic configuration
		Armour struct {
			// Armor protection (%)
			Total int `json:"total,omitempty"`
			// Armor
			Range struct {
				// Minimum value
				Min int `json:"min,omitempty"`
				// Maximum value
				Max int `json:"max,omitempty"`
			} `json:"range,omitempty"`
			// Hit points
			Health int `json:"health,omitempty"`
			// Torpedo Protection. Flooding Risk Reduction (%)
			FloodProb int `json:"flood_prob,omitempty"`
			// Torpedo Protection. Damage Reduction (%)
			FloodDamage int `json:"flood_damage,omitempty"`
			// Forward and After Ends
			Extremities struct {
				// Minimum value
				Min int `json:"min,omitempty"`
				// Maximum value
				Max int `json:"max,omitempty"`
			} `json:"extremities,omitempty"`
			// Armored Deck
			Deck struct {
				// Minimum value
				Min int `json:"min,omitempty"`
				// Maximum value
				Max int `json:"max,omitempty"`
			} `json:"deck,omitempty"`
			// Citadel
			Citadel struct {
				// Minimum value
				Min int `json:"min,omitempty"`
				// Maximum value
				Max int `json:"max,omitempty"`
			} `json:"citadel,omitempty"`
			// Gun Casemate
			Casemate struct {
				// Minimum value
				Min int `json:"min,omitempty"`
				// Maximum value
				Max int `json:"max,omitempty"`
			} `json:"casemate,omitempty"`
		} `json:"armour,omitempty"`
		// Anti-aircraft guns. If the module is absent on the ship, field value is null.
		AntiAircraft struct {
			// Gun slots
			Slots struct {
				// Gun name
				Name string `json:"name,omitempty"`
				// Number of guns
				Guns int `json:"guns,omitempty"`
				// Firing range (km)
				Distance float32 `json:"distance,omitempty"`
				// Caliber
				Caliber int `json:"caliber,omitempty"`
				// Average damage per second
				AvgDamage int `json:"avg_damage,omitempty"`
			} `json:"slots,omitempty"`
			// Anti-aircraft effectiveness
			Defense int `json:"defense,omitempty"`
		} `json:"anti_aircraft,omitempty"`
	} `json:"default_profile,omitempty"`
} 
