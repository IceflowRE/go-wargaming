package wows

type EncyclopediaModules struct {
	// Module type
	Type_ string `json:"type,omitempty"`
	// Tag
	Tag string `json:"tag,omitempty"`
	// Module parameters, values related to the module type
	Profile struct {
		// Torpedo tubes
		Torpedoes struct {
			// Torpedo Speed (knots)
			TorpedoSpeed int `json:"torpedo_speed,omitempty"`
			// Reload Time (sec)
			ShotSpeed float32 `json:"shot_speed,omitempty"`
			// Maximum Damage
			MaxDamage int `json:"max_damage,omitempty"`
			// Firing range
			Distance float32 `json:"distance,omitempty"`
		} `json:"torpedoes,omitempty"`
		// Torpedo Bombers
		TorpedoBomber struct {
			// Torpedo name
			TorpedoName string `json:"torpedo_name,omitempty"`
			// Top Speed (knots)
			TorpedoMaxSpeed int `json:"torpedo_max_speed,omitempty"`
			// Maximum torpedo damage
			TorpedoDamage int `json:"torpedo_damage,omitempty"`
			// Survivability
			MaxHealth int `json:"max_health,omitempty"`
			// Maximum Bomb Damage
			MaxDamage int `json:"max_damage,omitempty"`
			// Firing range
			Distance float32 `json:"distance,omitempty"`
			// Cruise Speed (knots)
			CruiseSpeed int `json:"cruise_speed,omitempty"`
		} `json:"torpedo_bomber,omitempty"`
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
			// Hit points
			Health int `json:"health,omitempty"`
			// Secondary gun turrets
			AtbaBarrels int `json:"atba_barrels,omitempty"`
			// Number of main turrets
			ArtilleryBarrels int `json:"artillery_barrels,omitempty"`
			// AA Mounts
			AntiAircraftBarrels int `json:"anti_aircraft_barrels,omitempty"`
		} `json:"hull,omitempty"`
		// Flight Control
		FlightControl struct {
			// Number of torpedo bomber squadrons
			TorpedoSquadrons int `json:"torpedo_squadrons,omitempty"`
			// Number of fighter squadrons
			FighterSquadrons int `json:"fighter_squadrons,omitempty"`
			// Number of bomber squadrons
			BomberSquadrons int `json:"bomber_squadrons,omitempty"`
		} `json:"flight_control,omitempty"`
		// Gun Fire Control System
		FireControl struct {
			// Firing Range extension (%)
			DistanceIncrease int `json:"distance_increase,omitempty"`
			// Firing range
			Distance float32 `json:"distance,omitempty"`
		} `json:"fire_control,omitempty"`
		// Fighters
		Fighter struct {
			// Survivability
			MaxHealth int `json:"max_health,omitempty"`
			// Ammunition
			MaxAmmo int `json:"max_ammo,omitempty"`
			// Cruise Speed (knots)
			CruiseSpeed int `json:"cruise_speed,omitempty"`
			// Average damage caused per second
			AvgDamage int `json:"avg_damage,omitempty"`
		} `json:"fighter,omitempty"`
		// Engine
		Engine struct {
			// Top Speed (knots)
			MaxSpeed float32 `json:"max_speed,omitempty"`
		} `json:"engine,omitempty"`
		// Dive bombers
		DiveBomber struct {
			// Survivability
			MaxHealth int `json:"max_health,omitempty"`
			// Maximum Bomb Damage
			MaxDamage int `json:"max_damage,omitempty"`
			// Cruise Speed (knots)
			CruiseSpeed int `json:"cruise_speed,omitempty"`
			// Chance of Fire on target caused by bomb (%)
			BombBurnProbability float32 `json:"bomb_burn_probability,omitempty"`
			// Accuracy
			Accuracy struct {
				// Minimum value
				Min float32 `json:"min,omitempty"`
				// Maximum value
				Max float32 `json:"max,omitempty"`
			} `json:"accuracy,omitempty"`
		} `json:"dive_bomber,omitempty"`
		// Main battery
		Artillery struct {
			// 180 Degree Turn Time (sec)
			RotationTime float32 `json:"rotation_time,omitempty"`
			// Maximum Damage caused by High Explosive Shells
			MaxDamageHe int `json:"max_damage_HE,omitempty"`
			// Maximum Damage caused by Armor Piercing Shells
			MaxDamageAp int `json:"max_damage_AP,omitempty"`
			// Rate of fire (rounds / min)
			GunRate float32 `json:"gun_rate,omitempty"`
		} `json:"artillery,omitempty"`
	} `json:"profile,omitempty"`
	// Cost in credits
	PriceCredit int `json:"price_credit,omitempty"`
	// Module name
	Name string `json:"name,omitempty"`
	// Module string ID
	ModuleIdStr string `json:"module_id_str,omitempty"`
	// Module ID
	ModuleId int `json:"module_id,omitempty"`
	// Image link
	Image string `json:"image,omitempty"`
} 
