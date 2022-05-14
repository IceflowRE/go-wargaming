package wargaming

type WgnAccountInfo struct {
	AccountId int      `json:"account_id,omitempty"`
	CreatedAt UnixTime `json:"created_at,omitempty"`
	Games     []string `json:"games,omitempty"`
	Nickname  string   `json:"nickname,omitempty"`
	Private   struct {
		FreeXp           int      `json:"free_xp,omitempty"`
		Gold             int      `json:"gold,omitempty"`
		PremiumExpiresAt UnixTime `json:"premium_expires_at,omitempty"`
	} `json:"private,omitempty"`
}

func (client *Client) WgnAccountInfo(realm Realm, accountId string, accessToken string, fields string, language string) (map[string]*WgnAccountInfo, error) {
	if err := ValidateRealm(realm, []Realm{RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"account_id":   accountId,
		"access_token": accessToken,
		"fields":       fields,
		"language":     language,
	}
	var result map[string]*WgnAccountInfo
	err := client.sendGetRequest(realm, "/wgn/account/info/", reqParam, &result)
	return result, err
}
