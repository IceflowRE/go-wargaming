// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wot"
	"strings"
)

// TanksMastery The method returns percentiles of the distribution of average damage or experience values for each piece of equipment
//
// https://developers.wargaming.net/reference/all/wot/tanks/mastery
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
// distribution:
//     Type of data. Valid values:
//
//     "damage" - Use damage distribution
//     "xp" - Use a distribution based on experience
// percentile:
//     A list of percentiles to be included in the response. Maximum limit: 10. Min value is 1. Maximum value: 100.
func (service *WotService) TanksMastery(ctx context.Context, realm Realm, distribution string, percentile []int, options *wot.TanksMasteryOptions) (*wot.TanksMastery, *GenericMeta, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, nil, err
	}

	reqParam := map[string]string{
		"distribution": distribution,
		"percentile":   internal.SliceIntToString(percentile, ","),
	}

	if options != nil {
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.TankId != nil {
			reqParam["tank_id"] = internal.SliceIntToString(options.TankId, ",")
		}
	}

	var data *wot.TanksMastery
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWot, realm, "/tanks/mastery/", reqParam, &data, &metaData)
	return data, metaData, err
}
