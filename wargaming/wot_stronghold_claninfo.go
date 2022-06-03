package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wot"
	"strings"
)

// StrongholdClaninfo Method returns general information and the battle statistics of clans in the Stronghold mode. Please note that information about the number of battles fought as well as the number of defeats and victories is updated once every 24 hours.
//
// https://developers.wargaming.net/reference/all/wot/stronghold/claninfo
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
// clanId:
//     Clan ID. To get a clan ID, use the Clans method. Maximum limit: 10.
func (service *wotService) StrongholdClaninfo(ctx context.Context, realm Realm, clanId []int, options *wot.StrongholdClaninfoOptions) (*wot.StrongholdClaninfo, error) {
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
	}

	var data *wot.StrongholdClaninfo
	err := service.client.getRequest(ctx, sectionWot, realm, "/stronghold/claninfo/", reqParam, &data)
	return data, err
}
