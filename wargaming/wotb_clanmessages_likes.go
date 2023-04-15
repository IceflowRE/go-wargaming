// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wotb"
	"strings"
)

// ClanmessagesLikes returns message likes on clan message board.
//
// https://developers.wargaming.net/reference/all/wotb/clanmessages/likes
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
// accessToken:
//     Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
// messageId:
//     Message ID. Maximum limit: 10.
func (service *WotbService) ClanmessagesLikes(ctx context.Context, realm Realm, accessToken string, messageId []int, options *wotb.ClanmessagesLikesOptions) ([]*wotb.ClanmessagesLikes, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"access_token": accessToken,
		"message_id":   internal.SliceIntToString(messageId, ","),
	}

	if options != nil {
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
	}

	var data []*wotb.ClanmessagesLikes
	err := service.client.getRequest(ctx, sectionWotb, realm, "/clanmessages/likes/", reqParam, &data, nil)
	return data, err
}
