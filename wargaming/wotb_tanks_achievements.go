package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wotb"
	"strconv"
	"strings"
)

// TanksAchievements Method returns list of player's achievements on all vehicles.
//
// https://developers.wargaming.net/reference/all/wotb/tanks/achievements
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
// accountId:
//     Player account ID
func (service *wotbService) TanksAchievements(ctx context.Context, realm Realm, accountId int, options *wotb.TanksAchievementsOptions) (*wotb.TanksAchievements, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"account_id": strconv.Itoa(accountId),
	}
	if options != nil {
		if options.TankId != nil {
			reqParam["tank_id"] = internal.SliceIntToString(options.TankId, ",")
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.InGarage != nil {
			reqParam["in_garage"] = *options.InGarage
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.AccessToken != nil {
			reqParam["access_token"] = *options.AccessToken
		}
	}

	var data *wotb.TanksAchievements
	err := service.client.getRequest(ctx, sectionWotb, realm, "/tanks/achievements/", reqParam, &data)
	return data, err
}
