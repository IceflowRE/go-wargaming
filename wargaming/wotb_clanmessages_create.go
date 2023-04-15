// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wgnTime"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wotb"
	"strconv"
)

// ClanmessagesCreate creates a message on clan message board.
//
// https://developers.wargaming.net/reference/all/wotb/clanmessages/create
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
// accessToken:
//     Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
// expiresAt:
//     Date when message will become irrelevant. Date in UNIX timestamp or ISO 8601 format. E.g.: 1376542800 or 2013-08-15T00:00:00. Date must be in the future.
// importance:
//     Message importance. Valid values:
//
//     "important" - Important messages
//     "standard" - Standard messages
// text:
//     Message text. Max. length: 1000. Maximum length: 1000.
// title:
//     Message title. Max. length: 100. Maximum length: 100.
// typ:
//     Message type. Valid values:
//
//     "general" - General messages
//     "training" - Training messages
//     "meeting" - Meeting messages
//     "battle" - Battle messages
func (service *WotbService) ClanmessagesCreate(ctx context.Context, realm Realm, accessToken string, expiresAt wgnTime.UnixTime, importance string, text string, title string, typ string) (*wotb.ClanmessagesCreate, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"access_token": accessToken,
		"expires_at":   strconv.FormatInt(expiresAt.Unix(), 10),
		"importance":   importance,
		"text":         text,
		"title":        title,
		"type":         typ,
	}

	var data *wotb.ClanmessagesCreate
	err := service.client.postRequest(ctx, sectionWotb, realm, "/clanmessages/create/", reqParam, &data, nil)
	return data, err
}
