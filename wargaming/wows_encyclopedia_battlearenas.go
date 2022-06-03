package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wows"
	"strings"
)

// EncyclopediaBattlearenas Method returns the information about maps.
//
// https://developers.wargaming.net/reference/all/wows/encyclopedia/battlearenas
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
func (service *wowsService) EncyclopediaBattlearenas(ctx context.Context, realm Realm, options *wows.EncyclopediaBattlearenasOptions) (*wows.EncyclopediaBattlearenas, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{}
	if options != nil {
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
	}

	var data *wows.EncyclopediaBattlearenas
	err := service.client.getRequest(ctx, sectionWows, realm, "/encyclopedia/battlearenas/", reqParam, &data)
	return data, err
}
