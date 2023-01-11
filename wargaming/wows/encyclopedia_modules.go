// Auto generated file!

package wows

type EncyclopediaModulesOptions struct {
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use "-" in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string `json:"fields,omitempty"`
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
	// Number of returned entries (fewer can be returned, but not more than 100). If the limit sent exceeds 100, a limit of 100 is applied (by default).
	Limit *int `json:"limit,omitempty"`
	// Module ID. Maximum limit: 100.
	ModuleId []int `json:"module_id,omitempty"`
	// Result page number. Default is 1. Min value is 1.
	PageNo *int `json:"page_no,omitempty"`
	// Module type. Valid values:
	//
	// "Artillery" - Main battery
	// "Torpedoes" - Torpedo tubes
	// "Suo" - Gun Fire Control System
	// "FlightControl" - Flight Control
	// "Hull" - Hull
	// "Engine" - Engine
	// "Fighter" - Fighters
	// "TorpedoBomber" - Torpedo Bombers
	// "DiveBomber" - Dive bombers
	Type *string `json:"type,omitempty"`
}

type EncyclopediaModules struct {
	// Image link
	Image *string `json:"image,omitempty"`
	// Module ID
	ModuleId *int `json:"module_id,omitempty"`
	// Module string ID
	ModuleIdStr *string `json:"module_id_str,omitempty"`
	// Module name
	Name *string `json:"name,omitempty"`
	// Cost in credits
	PriceCredit *int `json:"price_credit,omitempty"`
	// Module parameters, values related to the module type
	Profile *struct {
		// Main battery
		Artillery *struct {
			// Rate of fire (rounds / min)
			GunRate *float32 `json:"gun_rate,omitempty"`
			// Maximum Damage caused by Armor Piercing Shells
			MaxDamageAP *int `json:"max_damage_AP,omitempty"`
			// Maximum Damage caused by High Explosive Shells
			MaxDamageHE *int `json:"max_damage_HE,omitempty"`
			// 180 Degree Turn Time (sec)
			RotationTime *float32 `json:"rotation_time,omitempty"`
		} `json:"artillery,omitempty"`
		// Dive bombers
		DiveBomber *struct {
			// Accuracy
			Accuracy *struct {
				// Maximum value
				Max *float32 `json:"max,omitempty"`
				// Minimum value
				Min *float32 `json:"min,omitempty"`
			} `json:"accuracy,omitempty"`
			// Chance of Fire on target caused by bomb (%)
			BombBurnProbability *float32 `json:"bomb_burn_probability,omitempty"`
			// Cruise Speed (knots)
			CruiseSpeed *int `json:"cruise_speed,omitempty"`
			// Maximum Bomb Damage
			MaxDamage *int `json:"max_damage,omitempty"`
			// Survivability
			MaxHealth *int `json:"max_health,omitempty"`
		} `json:"dive_bomber,omitempty"`
		// Engine
		Engine *struct {
			// Top Speed (knots)
			MaxSpeed *float32 `json:"max_speed,omitempty"`
		} `json:"engine,omitempty"`
		// Fighters
		Fighter *struct {
			// Average damage caused per second
			AvgDamage *int `json:"avg_damage,omitempty"`
			// Cruise Speed (knots)
			CruiseSpeed *int `json:"cruise_speed,omitempty"`
			// Ammunition
			MaxAmmo *int `json:"max_ammo,omitempty"`
			// Survivability
			MaxHealth *int `json:"max_health,omitempty"`
		} `json:"fighter,omitempty"`
		// Gun Fire Control System
		FireControl *struct {
			// Firing range
			Distance *float32 `json:"distance,omitempty"`
			// Firing Range extension (%)
			DistanceIncrease *int `json:"distance_increase,omitempty"`
		} `json:"fire_control,omitempty"`
		// Flight Control
		FlightControl *struct {
			// Number of bomber squadrons
			BomberSquadrons *int `json:"bomber_squadrons,omitempty"`
			// Number of fighter squadrons
			FighterSquadrons *int `json:"fighter_squadrons,omitempty"`
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
		// Torpedo Bombers
		TorpedoBomber *struct {
			// Cruise Speed (knots)
			CruiseSpeed *int `json:"cruise_speed,omitempty"`
			// Firing range
			Distance *float32 `json:"distance,omitempty"`
			// Maximum Bomb Damage
			MaxDamage *int `json:"max_damage,omitempty"`
			// Survivability
			MaxHealth *int `json:"max_health,omitempty"`
			// Maximum torpedo damage
			TorpedoDamage *int `json:"torpedo_damage,omitempty"`
			// Top Speed (knots)
			TorpedoMaxSpeed *int `json:"torpedo_max_speed,omitempty"`
			// Torpedo name
			TorpedoName *string `json:"torpedo_name,omitempty"`
		} `json:"torpedo_bomber,omitempty"`
		// Torpedo tubes
		Torpedoes *struct {
			// Firing range
			Distance *float32 `json:"distance,omitempty"`
			// Maximum Damage
			MaxDamage *int `json:"max_damage,omitempty"`
			// Reload Time (sec)
			ShotSpeed *float32 `json:"shot_speed,omitempty"`
			// Torpedo Speed (knots)
			TorpedoSpeed *int `json:"torpedo_speed,omitempty"`
		} `json:"torpedoes,omitempty"`
	} `json:"profile,omitempty"`
	// Tag
	Tag *string `json:"tag,omitempty"`
	// Module type
	Type *string `json:"type,omitempty"`
}
