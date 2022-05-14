package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"strconv"
	"strings"
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
	if client.localCheck {
		if search == "" || len(search) > 24 {
			return nil, NewInvalidParameter("search", search, "cannot be empty or longer than 24 chars")
		}
		if err := CheckResponseFields(fields, []string{"account_id", "nickname"}); err != nil {
			return nil, err
		}
		if len(game) > 10 || !utils.ContainsAll(validGames, strings.Split(game, ",")) {
			return nil, NewInvalidParameter("game", game, "")
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
		"game":     game,
		"language": language,
		"limit":    strconv.Itoa(limit),
		"type":     typ,
	}
	var result []*WgnAccountList
	err := client.sendGetRequest(realm, "wgn/account/list/", reqParam, &result)
	return result, err
}
