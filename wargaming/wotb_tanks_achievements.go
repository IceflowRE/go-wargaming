// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wotb"
	"strconv"
	"strings"
)

// TanksAchievements returns list of player's achievements on all vehicles.
//
// https://developers.wargaming.net/reference/all/wotb/tanks/achievements
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
// accountId:
//     Player account ID
func (service *WotbService) TanksAchievements(ctx context.Context, realm Realm, accountId int, options *wotb.TanksAchievementsOptions) (*wotb.TanksAchievements, *GenericMeta, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, nil, err
	}

	reqParam := map[string]string{
		"account_id": strconv.Itoa(accountId),
	}

	if options != nil {
		if options.AccessToken != nil {
			reqParam["access_token"] = *options.AccessToken
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.InGarage != nil {
			reqParam["in_garage"] = *options.InGarage
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.TankId != nil {
			reqParam["tank_id"] = internal.SliceIntToString(options.TankId, ",")
		}
	}

	var data *wotb.TanksAchievements
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWotb, realm, "/tanks/achievements/", reqParam, &data, &metaData)
	return data, metaData, err
}
