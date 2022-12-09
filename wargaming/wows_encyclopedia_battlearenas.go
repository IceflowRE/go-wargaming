// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v3/wargaming/wows"
	"strings"
)

// EncyclopediaBattlearenas returns the information about maps.
//
// https://developers.wargaming.net/reference/all/wows/encyclopedia/battlearenas
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
func (service *WowsService) EncyclopediaBattlearenas(ctx context.Context, realm Realm, options *wows.EncyclopediaBattlearenasOptions) (*wows.EncyclopediaBattlearenas, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
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

	var data *wows.EncyclopediaBattlearenas
	err := service.client.getRequest(ctx, sectionWows, realm, "/encyclopedia/battlearenas/", reqParam, &data, nil)
	return data, err
}
