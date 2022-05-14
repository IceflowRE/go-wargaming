package wargaming

type WowpEncyclopediaInfo struct {
	// Date when Encyclopedia data were updated
	UpdatedAt UnixTime `json:"updated_at,omitempty"`
}

// WowpEncyclopediaInfo Method returns information about Encyclopedia.
//
// https://developers.wargaming.net/reference/all/wowp/encyclopedia/info
//

func (client *Client) WowpEncyclopediaInfo(realm Realm) (*WowpEncyclopediaInfo, error) {
	if err := ValidateRealm(realm, []Realm{RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{

	}

	var result *WowpEncyclopediaInfo
	err := client.doGetDataRequest(realm, "/wowp/encyclopedia/info/", reqParam, &result)
	return result, err
}
