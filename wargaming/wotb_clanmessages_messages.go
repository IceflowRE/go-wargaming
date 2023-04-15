// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wotb"
	"strconv"
	"strings"
)

// ClanmessagesMessages returns messages of clan message board. If a single message requested by ID, all filters are ignored and total equals null.
//
// https://developers.wargaming.net/reference/all/wotb/clanmessages/messages
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
// accessToken:
//     Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
func (service *WotbService) ClanmessagesMessages(ctx context.Context, realm Realm, accessToken string, options *wotb.ClanmessagesMessagesOptions) (map[string][]*wotb.ClanmessagesMessages, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"access_token": accessToken,
	}

	if options != nil {
		if options.ExpiresAfter != nil {
			reqParam["expires_after"] = strconv.FormatInt(options.ExpiresAfter.Unix(), 10)
		}
		if options.ExpiresBefore != nil {
			reqParam["expires_before"] = strconv.FormatInt(options.ExpiresBefore.Unix(), 10)
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.Importance != nil {
			reqParam["importance"] = *options.Importance
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.Limit != nil {
			reqParam["limit"] = strconv.Itoa(*options.Limit)
		}
		if options.MessageId != nil {
			reqParam["message_id"] = strconv.Itoa(*options.MessageId)
		}
		if options.OrderBy != nil {
			reqParam["order_by"] = strings.Join(options.OrderBy, ",")
		}
		if options.PageNo != nil {
			reqParam["page_no"] = strconv.Itoa(*options.PageNo)
		}
		if options.Status != nil {
			reqParam["status"] = *options.Status
		}
		if options.Type != nil {
			reqParam["type"] = *options.Type
		}
	}

	var data map[string][]*wotb.ClanmessagesMessages
	err := service.client.getRequest(ctx, sectionWotb, realm, "/clanmessages/messages/", reqParam, &data, nil)
	return data, err
}
