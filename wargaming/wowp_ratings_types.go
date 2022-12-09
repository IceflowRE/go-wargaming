// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v3/wargaming/wowp"
	"strings"
)

// RatingsTypes returns dictionary of rating periods and ratings details.
//
// https://developers.wargaming.net/reference/all/wowp/ratings/types
//
// realm:
//     Valid realms: RealmEu, RealmNa
func (service *WowpService) RatingsTypes(ctx context.Context, realm Realm, options *wowp.RatingsTypesOptions) (*wowp.RatingsTypes, error) {
	if err := validateRealm(realm, []Realm{RealmEu, RealmNa}); err != nil {
		return nil, err
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
	err := service.client.getRequest(ctx, sectionWowp, realm, "/ratings/types/", reqParam, &data, nil)
	return data, err
}
