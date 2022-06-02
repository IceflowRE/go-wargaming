package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/wargaming/wotb"
	"strings"
)

// TournamentsInfo Method returns detailed information on the specified tournament.
//
// https://developers.wargaming.net/reference/all/wotb/tournaments/info
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
// tournamentId:
//     Tournament ID that can be retrieved from the Tournaments list method. Maximum limit: 25.
func (service *wotbService) TournamentsInfo(ctx context.Context, realm Realm, tournamentId []int, options *wotb.TournamentsInfoOptions) (map[string]*wotb.TournamentsInfo, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"tournament_id": internal.SliceIntToString(tournamentId, ","),
	}
	if options != nil {
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
	}

	var data map[string]*wotb.TournamentsInfo
	err := service.client.getRequest(ctx, sectionWotb, realm, "/tournaments/info/", reqParam, &data)
	return data, err
}
