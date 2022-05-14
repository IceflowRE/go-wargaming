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
	if client.localCheck {
		if search == "" || len(search) > 24 {
			return nil, NewInvalidParameter("search", search, "cannot be empty or longer than 24 chars")
		}
		if err := CheckResponseFields(fields, []string{"account_id", "nickname"}); err != nil {
			return nil, err
		}
		if err := CheckLanguage(language); err != nil {
			return nil, err
		}
		if err := CheckLimit(limit); err != nil {
			return nil, err
		}
		if err := CheckType(typ); err != nil {
			return nil, err
		}
	}

	reqParam := map[string]string{
		"search":   search,
		"fields":   fields,
		"language": language,
		"limit":    strconv.Itoa(limit),
		"type":     typ,
	}
	var result []*WotAccountList
	err := client.sendGetRequest(realm, "wot/account/list/", reqParam, &result)
	return result, err
}
