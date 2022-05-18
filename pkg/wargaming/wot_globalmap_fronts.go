package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wot"
	"strconv"
	"strings"
)

// WotGlobalmapFronts Method returns information about the Global Map Fronts.
//
// https://developers.wargaming.net/reference/all/wot/globalmap/fronts
//
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// front_id:
//     List of Front IDs, to specify what fronts need to be returned. Maximum limit: 100.
// language:
//     Language. Default is "ru". Valid values:
//     
//     "ru" &mdash; Russian (by default)
// limit:
//     Number of returned entries (fewer can be returned, but not more than 100). If the limit sent exceeds 100, a limit of 100 is applied (by default).
// page_no:
//     Page number. Default is 1. Min value is 1.
func (client *Client) WotGlobalmapFronts(realm Realm, fields []string, frontId []string, language string, limit int, pageNo int) (*wot.GlobalmapFronts, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"fields": strings.Join(fields, ","),
		"front_id": strings.Join(frontId, ","),
		"language": language,
		"limit": strconv.Itoa(limit),
		"page_no": strconv.Itoa(pageNo),
	}

	var result *wot.GlobalmapFronts
	err := client.doGetDataRequest(realm, "/wot/globalmap/fronts/", reqParam, &result)
	return result, err
}
