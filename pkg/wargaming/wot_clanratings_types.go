package wargaming

type WotClanratingsTypes struct {
	// Rating categories
	RankFields []string `json:"rank_fields,omitempty"`
	// Rating period
	Type_ string `json:"type,omitempty"`
}

// WotClanratingsTypes Method returns details on ratings types and categories.
//
// https://developers.wargaming.net/reference/all/wot/clanratings/types
//

func (client *Client) WotClanratingsTypes(realm Realm) (*WotClanratingsTypes, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{

	}

	var result *WotClanratingsTypes
	err := client.doGetDataRequest(realm, "/wot/clanratings/types/", reqParam, &result)
	return result, err
}
