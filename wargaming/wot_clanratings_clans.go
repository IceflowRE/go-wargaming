package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wot"
	"strconv"
	"strings"
)

// ClanratingsClans returns clan ratings by specified IDs.
//
// https://developers.wargaming.net/reference/all/wot/clanratings/clans
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
// clanId:
//     Clan IDs. Maximum limit: 100.
func (service *WotService) ClanratingsClans(ctx context.Context, realm Realm, clanId []int, options *wot.ClanratingsClansOptions) (*wot.ClanratingsClans, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"clan_id": internal.SliceIntToString(clanId, ","),
	}
	if options != nil {
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.Date != nil {
			reqParam["date"] = strconv.FormatInt(options.Date.Unix(), 10)
		}
	}

	var data *wot.ClanratingsClans
	err := service.client.getRequest(ctx, sectionWot, realm, "/clanratings/clans/", reqParam, &data)
	return data, err
}
