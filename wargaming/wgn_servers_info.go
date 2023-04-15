// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wgn"
	"strings"
)

// ServersInfo returns the number of online players on the servers.
//
// https://developers.wargaming.net/reference/all/wgn/servers/info
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
func (service *WgnService) ServersInfo(ctx context.Context, realm Realm, options *wgn.ServersInfoOptions) (*wgn.ServersInfo, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{}

	if options != nil {
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.Game != nil {
			reqParam["game"] = strings.Join(options.Game, ",")
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
	}

	var data *wgn.ServersInfo
	err := service.client.getRequest(ctx, sectionWgn, realm, "/servers/info/", reqParam, &data, nil)
	return data, err
}
