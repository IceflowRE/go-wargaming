// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wowp"
	"strconv"
	"strings"
)

// PlanesAchievements returns achievements on player's aircraft.
//
// https://developers.wargaming.net/reference/all/wowp/planes/achievements
//
// realm:
//     Valid realms: RealmEu, RealmNa
// accountId:
//     Player account ID
func (service *WowpService) PlanesAchievements(ctx context.Context, realm Realm, accountId int, options *wowp.PlanesAchievementsOptions) (*wowp.PlanesAchievements, *GenericMeta, error) {
	if err := validateRealm(realm, []Realm{RealmEu, RealmNa}); err != nil {
		return nil, nil, err
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
		if options.PlaneId != nil {
			reqParam["plane_id"] = internal.SliceIntToString(options.PlaneId, ",")
		}
	}

	var data *wowp.PlanesAchievements
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWowp, realm, "/planes/achievements/", reqParam, &data, &metaData)
	return data, metaData, err
}
