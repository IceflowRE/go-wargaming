// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wot"
	"strings"
)

// StrongholdClaninfo returns general information and the battle statistics of clans in the Stronghold mode. Please note that information about the number of battles fought as well as the number of defeats and victories is updated once every 24 hours.
//
// https://developers.wargaming.net/reference/all/wot/stronghold/claninfo
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
// clanId:
//     Clan ID. To get a clan ID, use the Clans method. Maximum limit: 10.
func (service *WotService) StrongholdClaninfo(ctx context.Context, realm Realm, clanId []int, options *wot.StrongholdClaninfoOptions) (*wot.StrongholdClaninfo, *GenericMeta, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, nil, err
	}

	reqParam := map[string]string{
		"clan_id": internal.SliceIntToString(clanId, ","),
	}

	if options != nil {
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
	}

	var data *wot.StrongholdClaninfo
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWot, realm, "/stronghold/claninfo/", reqParam, &data, &metaData)
	return data, metaData, err
}
