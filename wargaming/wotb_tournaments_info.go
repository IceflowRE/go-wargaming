// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wotb"
	"strings"
)

// TournamentsInfo returns detailed information on the specified tournament.
//
// https://developers.wargaming.net/reference/all/wotb/tournaments/info
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
// tournamentId:
//     Tournament ID that can be retrieved from the Tournaments list method. Maximum limit: 25.
func (service *WotbService) TournamentsInfo(ctx context.Context, realm Realm, tournamentId []int, options *wotb.TournamentsInfoOptions) (*wotb.TournamentsInfo, *GenericMeta, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, nil, err
	}

	reqParam := map[string]string{
		"tournament_id": internal.SliceIntToString(tournamentId, ","),
	}

	if options != nil {
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
	}

	var data *wotb.TournamentsInfo
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWotb, realm, "/tournaments/info/", reqParam, &data, &metaData)
	return data, metaData, err
}
