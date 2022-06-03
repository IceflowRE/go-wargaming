package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wows"
	"strconv"
	"strings"
)

// EncyclopediaShipprofile Method returns parameters of ships in all existing configurations.
//
// https://developers.wargaming.net/reference/all/wows/encyclopedia/shipprofile
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
// shipId:
//     Ship ID
func (service *wowsService) EncyclopediaShipprofile(ctx context.Context, realm Realm, shipId int, options *wows.EncyclopediaShipprofileOptions) (*wows.EncyclopediaShipprofile, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"ship_id": strconv.Itoa(shipId),
	}
	if options != nil {
		if options.TorpedoesId != nil {
			reqParam["torpedoes_id"] = strconv.Itoa(*options.TorpedoesId)
		}
		if options.TorpedoBomberId != nil {
			reqParam["torpedo_bomber_id"] = strconv.Itoa(*options.TorpedoBomberId)
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.HullId != nil {
			reqParam["hull_id"] = strconv.Itoa(*options.HullId)
		}
		if options.FlightControlId != nil {
			reqParam["flight_control_id"] = strconv.Itoa(*options.FlightControlId)
		}
		if options.FireControlId != nil {
			reqParam["fire_control_id"] = strconv.Itoa(*options.FireControlId)
		}
		if options.FighterId != nil {
			reqParam["fighter_id"] = strconv.Itoa(*options.FighterId)
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.EngineId != nil {
			reqParam["engine_id"] = strconv.Itoa(*options.EngineId)
		}
		if options.DiveBomberId != nil {
			reqParam["dive_bomber_id"] = strconv.Itoa(*options.DiveBomberId)
		}
		if options.ArtilleryId != nil {
			reqParam["artillery_id"] = strconv.Itoa(*options.ArtilleryId)
		}
	}

	var data *wows.EncyclopediaShipprofile
	err := service.client.getRequest(ctx, sectionWows, realm, "/encyclopedia/shipprofile/", reqParam, &data)
	return data, err
}
