package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wot"
	"strconv"
	"strings"
)

// WotGlobalmapSeasons Method returns information about seasons.
//
// https://developers.wargaming.net/reference/all/wot/globalmap/seasons
//
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// language:
//     Language. Default is "ru". Valid values:
//     
//     "ru" - Russian (by default)
// limit:
//     Page limit. Default is 5. Min value is 1. Maximum value: 20.
// pageNo:
//     Page number. Default is 1. Min value is 1.
// seasonId:
//     Season ID
// status:
//     Response with seasons filtered by status. Valid values:
//     
//     "PLANNED" - Upcoming season 
//     "ACTIVE" - Current season 
//     "FINISHED" - Season is over
func (client *Client) WotGlobalmapSeasons(realm Realm, fields []string, language string, limit int, pageNo int, seasonId string, status string) (*wot.GlobalmapSeasons, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"fields": strings.Join(fields, ","),
		"language": language,
		"limit": strconv.Itoa(limit),
		"page_no": strconv.Itoa(pageNo),
		"season_id": seasonId,
		"status": status,
	}

	var result *wot.GlobalmapSeasons
	err := client.doGetDataRequest(realm, "/wot/globalmap/seasons/", reqParam, &result)
	return result, err
}