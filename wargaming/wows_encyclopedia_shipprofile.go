// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wows"
	"strconv"
	"strings"
)

// EncyclopediaShipprofile returns parameters of ships in all existing configurations.
//
// https://developers.wargaming.net/reference/all/wows/encyclopedia/shipprofile
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
// shipId:
//     Ship ID
func (service *WowsService) EncyclopediaShipprofile(ctx context.Context, realm Realm, shipId int, options *wows.EncyclopediaShipprofileOptions) (*wows.EncyclopediaShipprofile, *GenericMeta, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, nil, err
	}

	reqParam := map[string]string{
		"ship_id": strconv.Itoa(shipId),
	}

	if options != nil {
		if options.ArtilleryId != nil {
			reqParam["artillery_id"] = strconv.Itoa(*options.ArtilleryId)
		}
		if options.DiveBomberId != nil {
			reqParam["dive_bomber_id"] = strconv.Itoa(*options.DiveBomberId)
		}
		if options.EngineId != nil {
			reqParam["engine_id"] = strconv.Itoa(*options.EngineId)
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.FighterId != nil {
			reqParam["fighter_id"] = strconv.Itoa(*options.FighterId)
		}
		if options.FireControlId != nil {
			reqParam["fire_control_id"] = strconv.Itoa(*options.FireControlId)
		}
		if options.FlightControlId != nil {
			reqParam["flight_control_id"] = strconv.Itoa(*options.FlightControlId)
		}
		if options.HullId != nil {
			reqParam["hull_id"] = strconv.Itoa(*options.HullId)
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.TorpedoBomberId != nil {
			reqParam["torpedo_bomber_id"] = strconv.Itoa(*options.TorpedoBomberId)
		}
		if options.TorpedoesId != nil {
			reqParam["torpedoes_id"] = strconv.Itoa(*options.TorpedoesId)
		}
	}

	var data *wows.EncyclopediaShipprofile
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWows, realm, "/encyclopedia/shipprofile/", reqParam, &data, &metaData)
	return data, metaData, err
}
