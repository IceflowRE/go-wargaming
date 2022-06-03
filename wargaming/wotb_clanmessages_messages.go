package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wotb"
	"strconv"
	"strings"
)

// ClanmessagesMessages Method returns messages of clan message board. If a single message requested by ID, all filters are ignored and total equals null.
//
// https://developers.wargaming.net/reference/all/wotb/clanmessages/messages
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
// accessToken:
//     Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
func (service *WotbService) ClanmessagesMessages(ctx context.Context, realm Realm, accessToken string, options *wotb.ClanmessagesMessagesOptions) (map[string][]*wotb.ClanmessagesMessages, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"access_token": accessToken,
	}
	if options != nil {
		if options.Type_ != nil {
			reqParam["type"] = *options.Type_
		}
		if options.Status != nil {
			reqParam["status"] = *options.Status
		}
		if options.PageNo != nil {
			reqParam["page_no"] = strconv.Itoa(*options.PageNo)
		}
		if options.OrderBy != nil {
			reqParam["order_by"] = strings.Join(options.OrderBy, ",")
		}
		if options.MessageId != nil {
			reqParam["message_id"] = strconv.Itoa(*options.MessageId)
		}
		if options.Limit != nil {
			reqParam["limit"] = strconv.Itoa(*options.Limit)
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.Importance != nil {
			reqParam["importance"] = *options.Importance
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.ExpiresBefore != nil {
			reqParam["expires_before"] = strconv.FormatInt(options.ExpiresBefore.Unix(), 10)
		}
		if options.ExpiresAfter != nil {
			reqParam["expires_after"] = strconv.FormatInt(options.ExpiresAfter.Unix(), 10)
		}
	}

	var data map[string][]*wotb.ClanmessagesMessages
	err := service.client.getRequest(ctx, sectionWotb, realm, "/clanmessages/messages/", reqParam, &data)
	return data, err
}
