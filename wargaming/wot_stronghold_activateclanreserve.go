package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/wargaming/wot"
	"strconv"
	"strings"
)

// StrongholdActivateclanreserve This method activates an available clan Reserve. A clan Reserve can be activated only by a clan member with the required permission.
//
// https://developers.wargaming.net/reference/all/wot/stronghold/activateclanreserve
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
// accessToken:
//     Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
// reserveLevel:
//     Level of clan Reserve to be activated
// reserveType:
//     Type of clan Reserve to be activated
func (service *wotService) StrongholdActivateclanreserve(ctx context.Context, realm Realm, accessToken string, reserveLevel int, reserveType string, options *wot.StrongholdActivateclanreserveOptions) (*wot.StrongholdActivateclanreserve, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"access_token":  accessToken,
		"reserve_level": strconv.Itoa(reserveLevel),
		"reserve_type":  reserveType,
	}
	if options != nil {
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
	}

	var data *wot.StrongholdActivateclanreserve
	err := service.client.postRequest(ctx, sectionWot, realm, "/stronghold/activateclanreserve/", reqParam, &data)
	return data, err
}
