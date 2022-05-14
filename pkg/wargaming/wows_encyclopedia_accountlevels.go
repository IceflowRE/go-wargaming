package wargaming

import (
	"strings"
)

type WowsEncyclopediaAccountlevels struct {
	// URL to Service Record level icon
	Image string `json:"image,omitempty"`
	// Points to achieve the Service Record level
	Points int `json:"points,omitempty"`
	// Service Record level
	Tier int `json:"tier,omitempty"`
}

// WowsEncyclopediaAccountlevels Method returns information about Service Record levels.
//
// https://developers.wargaming.net/reference/all/wows/encyclopedia/accountlevels
//
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
func (client *Client) WowsEncyclopediaAccountlevels(realm Realm, fields []string) (*WowsEncyclopediaAccountlevels, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"fields": strings.Join(fields, ","),
	}

	var result *WowsEncyclopediaAccountlevels
	err := client.doGetDataRequest(realm, "/wows/encyclopedia/accountlevels/", reqParam, &result)
	return result, err
}
