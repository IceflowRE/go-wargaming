// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wgn"
	"strconv"
	"strings"
)

// WgtvVideos returns list of videos filtered by the specified parameter.
//
// https://developers.wargaming.net/reference/all/wgn/wgtv/videos
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
func (service *WgnService) WgtvVideos(ctx context.Context, realm Realm, options *wgn.WgtvVideosOptions) ([]*wgn.WgtvVideos, *GenericMeta, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, nil, err
	}

	reqParam := map[string]string{}

	if options != nil {
		if options.CategoryId != nil {
			reqParam["category_id"] = internal.SliceIntToString(options.CategoryId, ",")
		}
		if options.DateFrom != nil {
			reqParam["date_from"] = strconv.FormatInt(options.DateFrom.Unix(), 10)
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.Important != nil {
			reqParam["important"] = strconv.Itoa(*options.Important)
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.Limit != nil {
			reqParam["limit"] = strconv.Itoa(*options.Limit)
		}
		if options.PageNo != nil {
			reqParam["page_no"] = strconv.Itoa(*options.PageNo)
		}
		if options.ProgramId != nil {
			reqParam["program_id"] = internal.SliceIntToString(options.ProgramId, ",")
		}
		if options.ProjectId != nil {
			reqParam["project_id"] = internal.SliceIntToString(options.ProjectId, ",")
		}
		if options.Q != nil {
			reqParam["q"] = *options.Q
		}
		if options.VehicleId != nil {
			reqParam["vehicle_id"] = internal.SliceIntToString(options.VehicleId, ",")
		}
		if options.VideoId != nil {
			reqParam["video_id"] = strings.Join(options.VideoId, ",")
		}
	}

	var data []*wgn.WgtvVideos
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWgn, realm, "/wgtv/videos/", reqParam, &data, &metaData)
	return data, metaData, err
}
