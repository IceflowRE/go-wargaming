package wargaming

import (
	"strconv"
)

type WgnAccountList struct {
	AccountId int      `json:"account_id,omitempty"`
	CreatedAt UnixTime `json:"created_at,omitempty"`
	Games     []string `json:"games,omitempty"`
	Nickname  string   `json:"nickname,omitempty"`
}

func (client *Client) WgnAccountList(realm Realm, search string, fields string, game string, language string, limit int, typ string) ([]*WgnAccountList, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"search":   search,
		"fields":   fields,
		"game":     game,
		"language": language,
		"limit":    strconv.Itoa(limit),
		"type":     typ,
	}
	var result []*WgnAccountList
	err := client.sendGetRequest(realm, "/wgn/account/list/", reqParam, &result)
	return result, err
}
