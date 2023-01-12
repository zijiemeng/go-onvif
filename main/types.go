package biz

import (
	"code.byted.org/videoarch/go-onvif/media"
)

type ReflectTypes struct {
	GetVideoSources struct {
		media.GetVideoSources         `req:""`
		media.GetVideoSourcesResponse `resp:""`
	} `method:"GetVideoSources"`
	GetStreamUri struct {
		media.GetStreamUri         `req:""`
		media.GetStreamUriResponse `resp:""`
	} `method:"GetStreamUri"`
	GetProfiles struct {
		media.GetProfiles         `req:""`
		media.GetProfilesResponse `resp:""`
	} `method:"GetProfiles"`
	GetProfile struct {
		media.GetProfile         `req:""`
		media.GetProfileResponse `req:""`
	}
}
