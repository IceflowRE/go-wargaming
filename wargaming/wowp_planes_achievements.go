package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wowp"
	"strconv"
	"strings"
)

// PlanesAchievements returns achievements on player's aircraft.
//
// https://developers.wargaming.net/reference/all/wowp/planes/achievements
//
// realm:
//     Valid realms: RealmEu, RealmNa, RealmRu
// accountId:
//     Player account ID
func (service *WowpService) PlanesAchievements(ctx context.Context, realm Realm, accountId int, options *wowp.PlanesAchievementsOptions) (*wowp.PlanesAchievements, error) {
	if err := validateRealm(realm, []Realm{RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"account_id": strconv.Itoa(accountId),
	}
	if options != nil {
		if options.PlaneId != nil {
			reqParam["plane_id"] = internal.SliceIntToString(options.PlaneId, ",")
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
	}

	var data *wowp.PlanesAchievements
	err := service.client.getRequest(ctx, sectionWowp, realm, "/planes/achievements/", reqParam, &data)
	return data, err
}
