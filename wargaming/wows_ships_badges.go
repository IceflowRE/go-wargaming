// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wows"
	"strconv"
	"strings"
)

// ShipsBadges returns badges for player's ships.
//
// https://developers.wargaming.net/reference/all/wows/ships/badges
//
// realm:
//
//	Valid realms: RealmAsia, RealmEu, RealmNa
//
// accountId:
//
//	Player account ID
func (service *WowsService) ShipsBadges(ctx context.Context, realm Realm, accountId int, options *wows.ShipsBadgesOptions) (*wows.ShipsBadges, error) {
	if !containsRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}) {
		return nil, InvalidRealm{realm}
	}

	reqParam := map[string]string{
		"account_id": strconv.Itoa(accountId),
	}

	if options != nil {
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.ShipId != nil {
			reqParam["ship_id"] = internal.SliceIntToString(options.ShipId, ",")
		}
	}

	var data *wows.ShipsBadges
	err := service.client.getRequest(ctx, sectionWows, realm, "/ships/badges/", reqParam, &data, nil)
	return data, err
}
