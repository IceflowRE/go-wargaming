package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wgn"
	"strconv"
	"strings"
)

// WgtvVideos Method returns list of videos filtered by the specified parameter.
//
// https://developers.wargaming.net/reference/all/wgn/wgtv/videos
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
func (service *WgnService) WgtvVideos(ctx context.Context, realm Realm, options *wgn.WgtvVideosOptions) ([]*wgn.WgtvVideos, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{}
	if options != nil {
		if options.VideoId != nil {
			reqParam["video_id"] = strings.Join(options.VideoId, ",")
		}
		if options.VehicleId != nil {
			reqParam["vehicle_id"] = internal.SliceIntToString(options.VehicleId, ",")
		}
		if options.Q != nil {
			reqParam["q"] = *options.Q
		}
		if options.ProjectId != nil {
			reqParam["project_id"] = internal.SliceIntToString(options.ProjectId, ",")
		}
		if options.ProgramId != nil {
			reqParam["program_id"] = internal.SliceIntToString(options.ProgramId, ",")
		}
		if options.PageNo != nil {
			reqParam["page_no"] = strconv.Itoa(*options.PageNo)
		}
		if options.Limit != nil {
			reqParam["limit"] = strconv.Itoa(*options.Limit)
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.Important != nil {
			reqParam["important"] = strconv.Itoa(*options.Important)
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.DateFrom != nil {
			reqParam["date_from"] = strconv.FormatInt(options.DateFrom.Unix(), 10)
		}
		if options.CategoryId != nil {
			reqParam["category_id"] = internal.SliceIntToString(options.CategoryId, ",")
		}
	}

	var data []*wgn.WgtvVideos
	err := service.client.getRequest(ctx, sectionWgn, realm, "/wgtv/videos/", reqParam, &data)
	return data, err
}
