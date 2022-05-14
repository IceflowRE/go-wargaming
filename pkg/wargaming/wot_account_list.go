package wargaming

import (
	"strconv"
)

type WotAccountList struct {
	AccountId int    `json:"account_id,omitempty"`
	Nickname  string `json:"nickname,omitempty"`
}

func (client *Client) WotAccountList(realm Realm, search string, fields string, language string, limit int, typ string) ([]*WotAccountList, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"search":   search,
		"fields":   fields,
		"language": language,
		"limit":    strconv.Itoa(limit),
		"type":     typ,
	}
	var result []*WotAccountList
	err := client.sendGetRequest(realm, "/wot/account/list/", reqParam, &result)
	return result, err
}
