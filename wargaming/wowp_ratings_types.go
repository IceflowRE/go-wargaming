// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wowp"
	"strings"
)

// RatingsTypes returns dictionary of rating periods and ratings details.
//
// https://developers.wargaming.net/reference/all/wowp/ratings/types
//
// realm:
//     Valid realms: RealmEu, RealmNa
func (service *WowpService) RatingsTypes(ctx context.Context, realm Realm, options *wowp.RatingsTypesOptions) (*wowp.RatingsTypes, *GenericMeta, error) {
	if err := validateRealm(realm, []Realm{RealmEu, RealmNa}); err != nil {
		return nil, nil, err
	}

	reqParam := map[string]string{}

	if options != nil {
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
	}

	var data *wowp.RatingsTypes
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWowp, realm, "/ratings/types/", reqParam, &data, &metaData)
	return data, metaData, err
}
