package onvif

import (
	"code.byted.org/videoarch/go-onvif/onvif/common"
	"code.byted.org/videoarch/go-onvif/onvif/v10/device"
	"code.byted.org/videoarch/go-onvif/onvif/v10/device_io"
	"code.byted.org/videoarch/go-onvif/onvif/v10/media"
	media2 "code.byted.org/videoarch/go-onvif/onvif/v20/media"
	"fmt"
	"strings"
)

type Client interface {
	common.Client
}

type EndpointManager struct {
	endpoints map[string]string
	Device    *EptDevice
	Media     *EptMedia
}

func NewEndpointManager(endpoints map[string]string, client Client) *EndpointManager {
	mgr := &EndpointManager{
		endpoints: endpoints,
	}
	for k, v := range endpoints {
		switch strings.ToLower(k) {
		case "device":
			mgr.Device = &EptDevice{
				device.NewDevice(v, client), device_io.NewDeviceIOPort(v, client),
			}
		case "media":
			mgr.Media = &EptMedia{
				Media:  media.NewMedia(v, client),
				Media2: media2.NewMedia2(v, client),
			}
		default:
			fmt.Printf("unsupported endpoint: [%s : %s]", k, v)
		}
	}
	return mgr
}

type EptDevice struct {
	device.Device
	device_io.DeviceIOPort
}

type EptMedia struct {
	media.Media
	media2.Media2
}
