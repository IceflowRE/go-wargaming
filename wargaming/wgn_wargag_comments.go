package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wgn"
	"strconv"
	"strings"
)

// WargagComments Method returns comments to the content.
//
// https://developers.wargaming.net/reference/all/wgn/wargag/comments
//
// realm:
//     Valid realms: RealmRu
// contentId:
//     Publication ID
func (service *wgnService) WargagComments(ctx context.Context, realm Realm, contentId int, options *wgn.WargagCommentsOptions) (*wgn.WargagComments, error) {
	if err := validateRealm(realm, []Realm{RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"content_id": strconv.Itoa(contentId),
	}
	if options != nil {
		if options.PageNo != nil {
			reqParam["page_no"] = strconv.Itoa(*options.PageNo)
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
	}

	var data *wgn.WargagComments
	err := service.client.getRequest(ctx, sectionWgn, realm, "/wargag/comments/", reqParam, &data)
	return data, err
}
