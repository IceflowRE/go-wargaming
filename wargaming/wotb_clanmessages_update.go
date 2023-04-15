// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wotb"
	"strconv"
)

// ClanmessagesUpdate updates a message on clan message board. At least one changed parameter should be specified.
//
// https://developers.wargaming.net/reference/all/wotb/clanmessages/update
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
// accessToken:
//     Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
// messageId:
//     Message ID
func (service *WotbService) ClanmessagesUpdate(ctx context.Context, realm Realm, accessToken string, messageId int, options *wotb.ClanmessagesUpdateOptions) (*wotb.ClanmessagesUpdate, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"access_token": accessToken,
		"message_id":   strconv.Itoa(messageId),
	}

	if options != nil {
		if options.ExpiresAt != nil {
			reqParam["expires_at"] = strconv.FormatInt(options.ExpiresAt.Unix(), 10)
		}
		if options.Importance != nil {
			reqParam["importance"] = *options.Importance
		}
		if options.Text != nil {
			reqParam["text"] = *options.Text
		}
		if options.Title != nil {
			reqParam["title"] = *options.Title
		}
		if options.Type != nil {
			reqParam["type"] = *options.Type
		}
	}

	var data *wotb.ClanmessagesUpdate
	err := service.client.postRequest(ctx, sectionWotb, realm, "/clanmessages/update/", reqParam, &data, nil)
	return data, err
}
