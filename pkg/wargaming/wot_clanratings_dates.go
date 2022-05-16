package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wot"
	"strconv"
)

// WotClanratingsDates Method returns dates with available rating data.
//
// https://developers.wargaming.net/reference/all/wot/clanratings/dates
//
// limit:
//     Number of returned entries (fewer can be returned, but not more than 365). If the limit sent exceeds 365, a limit of 7 is applied (by default).
func (client *Client) WotClanratingsDates(realm Realm, limit int) (*wot.ClanratingsDates, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"limit": strconv.Itoa(limit),
	}

	var result *wot.ClanratingsDates
	err := client.doGetDataRequest(realm, "/wot/clanratings/dates/", reqParam, &result)
	return result, err
}
