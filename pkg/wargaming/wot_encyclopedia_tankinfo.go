package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"strings"
)

type WotEncyclopediaTankinfo struct {
	// Standard suspension traverse speed
	ChassisRotationSpeed int `json:"chassis_rotation_speed,omitempty"`
	// Standard turret view range
	CircularVisionRadius int `json:"circular_vision_radius,omitempty"`
	// URL to outline image of vehicle
	ContourImage string `json:"contour_image,omitempty"`
	// Standard engine power
	EnginePower int `json:"engine_power,omitempty"`
	// Maximum damage of standard gun
	GunDamageMax int `json:"gun_damage_max,omitempty"`
	// Minimum damage of standard gun
	GunDamageMin int `json:"gun_damage_min,omitempty"`
	// Ammunition of standard gun
	GunMaxAmmo int `json:"gun_max_ammo,omitempty"`
	// Standard gun name
	GunName string `json:"gun_name,omitempty"`
	// Maximum penetration of standard gun
	GunPiercingPowerMax int `json:"gun_piercing_power_max,omitempty"`
	// Minimum penetration of standard gun
	GunPiercingPowerMin int `json:"gun_piercing_power_min,omitempty"`
	// Rate of fire of standard gun
	GunRate float32 `json:"gun_rate,omitempty"`
	// URL to 160 x 100 px image of vehicle
	Image string `json:"image,omitempty"`
	// URL to 124 x 31 px image of vehicle
	ImageSmall string `json:"image_small,omitempty"`
	// Indicates if the vehicle was a gift
	IsGift bool `json:"is_gift,omitempty"`
	// Indicates if the vehicle is Premium vehicle
	IsPremium bool `json:"is_premium,omitempty"`
	// Tier
	Level int `json:"level,omitempty"`
	// Load limit
	LimitWeight float32 `json:"limit_weight,omitempty"`
	// Localized name field
	LocalizedName string `json:"localized_name,omitempty"`
	// Hit points
	MaxHealth int `json:"max_health,omitempty"`
	// Vehicle name
	Name string `json:"name,omitempty"`
	// Localized name field
	NameI18N string `json:"name_i18n,omitempty"`
	// Nation
	Nation string `json:"nation,omitempty"`
	// Localized nation field
	NationI18N string `json:"nation_i18n,omitempty"`
	// Purchase cost in credits
	PriceCredit int `json:"price_credit,omitempty"`
	// Purchase cost in gold
	PriceGold int `json:"price_gold,omitempty"`
	// Signal range of standard radio
	RadioDistance int `json:"radio_distance,omitempty"`
	// Localized short name of vehicle
	ShortNameI18N string `json:"short_name_i18n,omitempty"`
	// Speed limit
	SpeedLimit float32 `json:"speed_limit,omitempty"`
	// Vehicle ID
	TankId int `json:"tank_id,omitempty"`
	// Standard turret armor: sides
	TurretArmorBoard int `json:"turret_armor_board,omitempty"`
	// Standard turret armor: sides
	TurretArmorFedd int `json:"turret_armor_fedd,omitempty"`
	// Standard turret armor: front
	TurretArmorForehead int `json:"turret_armor_forehead,omitempty"`
	// Standard turret traverse speed
	TurretRotationSpeed int `json:"turret_rotation_speed,omitempty"`
	// Vehicle type
	Type_ string `json:"type,omitempty"`
	// Localized vehicle type
	TypeI18N string `json:"type_i18n,omitempty"`
	// Hull armor: sides
	VehicleArmorBoard int `json:"vehicle_armor_board,omitempty"`
	// Hull armor: rear
	VehicleArmorFedd int `json:"vehicle_armor_fedd,omitempty"`
	// Hull armor: front
	VehicleArmorForehead int `json:"vehicle_armor_forehead,omitempty"`
	// Parent vehicles in Tech Tree
	ParentTanks []int `json:"parent_tanks,omitempty"`
	// Cost of research in experience
	PriceXp int `json:"price_xp,omitempty"`
	// Weight
	Weight float32 `json:"weight,omitempty"`
	// Compatible suspensions
	Chassis struct {
		// Module ID
		ModuleId int `json:"module_id,omitempty"`
		// Indicates if the module is standard
		IsDefault bool `json:"is_default,omitempty"`
	} `json:"chassis,omitempty"`
	// Crew
	Crew struct {
		// Qualification of crew member
		Role string `json:"role,omitempty"`
		// Localized role field
		RoleI18N string `json:"role_i18n,omitempty"`
		// Additional qualifications of crew member
		AdditionalRoles []string `json:"additional_roles,omitempty"`
		// Additional qualifications of crew member
		AdditionalRolesI18N struct {
			// Qualification of crew member
			Role string `json:"role,omitempty"`
			// Localized role field
			RoleI18N string `json:"role_i18n,omitempty"`
		} `json:"additional_roles_i18n,omitempty"`
	} `json:"crew,omitempty"`
	// Compatible engines
	Engines struct {
		// Module ID
		ModuleId int `json:"module_id,omitempty"`
		// Indicates if the module is standard
		IsDefault bool `json:"is_default,omitempty"`
	} `json:"engines,omitempty"`
	// Compatible guns
	Guns struct {
		// Module ID
		ModuleId int `json:"module_id,omitempty"`
		// Indicates if the module is standard
		IsDefault bool `json:"is_default,omitempty"`
	} `json:"guns,omitempty"`
	// Compatible radios
	Radios struct {
		// Module ID
		ModuleId int `json:"module_id,omitempty"`
		// Indicates if the module is standard
		IsDefault bool `json:"is_default,omitempty"`
	} `json:"radios,omitempty"`
	// Compatible turrets
	Turrets struct {
		// Module ID
		ModuleId int `json:"module_id,omitempty"`
		// Indicates if the module is standard
		IsDefault bool `json:"is_default,omitempty"`
	} `json:"turrets,omitempty"`
}

// WotEncyclopediaTankinfo Deprecated: Attention! The method is deprecated.
// Method returns vehicle details from Tankopedia.
//
// https://developers.wargaming.net/reference/all/wot/encyclopedia/tankinfo
//
// tank_id:
//     Vehicle ID. Maximum limit: 1000.
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// language:
//     Localization language. Default is "ru". Valid values:
//     
//     "en" &mdash; English 
//     "ru" &mdash; Русский (by default)
//     "pl" &mdash; Polski 
//     "de" &mdash; Deutsch 
//     "fr" &mdash; Français 
//     "es" &mdash; Español 
//     "zh-cn" &mdash; 简体中文 
//     "zh-tw" &mdash; 繁體中文 
//     "tr" &mdash; Türkçe 
//     "cs" &mdash; Čeština 
//     "th" &mdash; ไทย 
//     "vi" &mdash; Tiếng Việt 
//     "ko" &mdash; 한국어
func (client *Client) WotEncyclopediaTankinfo(realm Realm, tankId []int, fields []string, language string) (*WotEncyclopediaTankinfo, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"tank_id": utils.SliceIntToString(tankId, ","),
		"fields": strings.Join(fields, ","),
		"language": language,
	}

	var result *WotEncyclopediaTankinfo
	err := client.doGetDataRequest(realm, "/wot/encyclopedia/tankinfo/", reqParam, &result)
	return result, err
}
