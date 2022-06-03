package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wowp"
	"strings"
)

// RatingsDates Method returns dates with available rating data.
//
// https://developers.wargaming.net/reference/all/wowp/ratings/dates
//
// realm:
//     Valid realms: RealmEu, RealmNa, RealmRu
// type_:
//     Rating period. For valid values, check the Types of ratings method.
func (service *wowpService) RatingsDates(ctx context.Context, realm Realm, type_ string, options *wowp.RatingsDatesOptions) (*wowp.RatingsDates, error) {
	if err := validateRealm(realm, []Realm{RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"type": type_,
	}
	if options != nil {
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.AccountId != nil {
			reqParam["account_id"] = internal.SliceIntToString(options.AccountId, ",")
		}
	}

	var data *wowp.RatingsDates
	err := service.client.getRequest(ctx, sectionWowp, realm, "/ratings/dates/", reqParam, &data)
	return data, err
}
