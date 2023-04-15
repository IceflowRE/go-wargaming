// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wowp"
	"strings"
)

// RatingsDates returns dates with available rating data.
//
// https://developers.wargaming.net/reference/all/wowp/ratings/dates
//
// realm:
//     Valid realms: RealmEu, RealmNa
// typ:
//     Rating period. For valid values, check the Types of ratings method.
func (service *WowpService) RatingsDates(ctx context.Context, realm Realm, typ string, options *wowp.RatingsDatesOptions) (*wowp.RatingsDates, *GenericMeta, error) {
	if err := validateRealm(realm, []Realm{RealmEu, RealmNa}); err != nil {
		return nil, nil, err
	}

	reqParam := map[string]string{
		"type": typ,
	}

	if options != nil {
		if options.AccountId != nil {
			reqParam["account_id"] = internal.SliceIntToString(options.AccountId, ",")
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
	}

	var data *wowp.RatingsDates
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWowp, realm, "/ratings/dates/", reqParam, &data, &metaData)
	return data, metaData, err
}
