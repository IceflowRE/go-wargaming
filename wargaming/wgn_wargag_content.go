package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/wargaming/wgn"
	"strconv"
	"strings"
)

// WargagContent Method returns information about content.
//
// https://developers.wargaming.net/reference/all/wgn/wargag/content
//
// realm:
//     Valid realms: RealmRu
func (service *wgnService) WargagContent(ctx context.Context, realm Realm, options *wgn.WargagContentOptions) (*wgn.WargagContent, error) {
	if err := validateRealm(realm, []Realm{RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{}
	if options != nil {
		if options.Type_ != nil {
			reqParam["type"] = *options.Type_
		}
		if options.TagId != nil {
			reqParam["tag_id"] = strconv.Itoa(*options.TagId)
		}
		if options.RatingThreshold != nil {
			reqParam["rating_threshold"] = strconv.Itoa(*options.RatingThreshold)
		}
		if options.PageNo != nil {
			reqParam["page_no"] = strconv.Itoa(*options.PageNo)
		}
		if options.OrderBy != nil {
			reqParam["order_by"] = *options.OrderBy
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.Date != nil {
			reqParam["date"] = strconv.FormatInt(options.Date.Unix(), 10)
		}
		if options.ContentId != nil {
			reqParam["content_id"] = strconv.Itoa(*options.ContentId)
		}
		if options.CategoryId != nil {
			reqParam["category_id"] = strconv.Itoa(*options.CategoryId)
		}
		if options.AccountId != nil {
			reqParam["account_id"] = strconv.Itoa(*options.AccountId)
		}
		if options.AccessToken != nil {
			reqParam["access_token"] = *options.AccessToken
		}
	}

	var data *wgn.WargagContent
	err := service.client.getRequest(ctx, sectionWgn, realm, "/wargag/content/", reqParam, &data)
	return data, err
}
