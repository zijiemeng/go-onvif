package onvif

import (
	"code.byted.org/videoarch/go-onvif/onvif/common"
	"code.byted.org/videoarch/go-onvif/onvif/v10/access_control"
	"code.byted.org/videoarch/go-onvif/onvif/v10/access_rules"
	"code.byted.org/videoarch/go-onvif/onvif/v10/action_engine"
	"code.byted.org/videoarch/go-onvif/onvif/v10/analytics_device"
	"code.byted.org/videoarch/go-onvif/onvif/v10/app_mgmt"
	"code.byted.org/videoarch/go-onvif/onvif/v10/authentication_behavior"
	"code.byted.org/videoarch/go-onvif/onvif/v10/credential"
	"code.byted.org/videoarch/go-onvif/onvif/v10/device"
	"code.byted.org/videoarch/go-onvif/onvif/v10/device_io"
	"code.byted.org/videoarch/go-onvif/onvif/v10/display"
	"code.byted.org/videoarch/go-onvif/onvif/v10/door_control"
	"code.byted.org/videoarch/go-onvif/onvif/v10/events"
	"code.byted.org/videoarch/go-onvif/onvif/v10/federated_search"
	"code.byted.org/videoarch/go-onvif/onvif/v10/media"
	"code.byted.org/videoarch/go-onvif/onvif/v10/provisioning"
	"code.byted.org/videoarch/go-onvif/onvif/v10/receiver"
	"code.byted.org/videoarch/go-onvif/onvif/v10/recording"
	"code.byted.org/videoarch/go-onvif/onvif/v10/replay"
	"code.byted.org/videoarch/go-onvif/onvif/v10/schedule"
	"code.byted.org/videoarch/go-onvif/onvif/v10/search"
	"code.byted.org/videoarch/go-onvif/onvif/v10/security"
	"code.byted.org/videoarch/go-onvif/onvif/v10/thermal"
	"code.byted.org/videoarch/go-onvif/onvif/v10/uplink"
	"code.byted.org/videoarch/go-onvif/onvif/v20/analytics"
	"code.byted.org/videoarch/go-onvif/onvif/v20/imaging"
	media2 "code.byted.org/videoarch/go-onvif/onvif/v20/media"
	"code.byted.org/videoarch/go-onvif/onvif/v20/ptz"
	"fmt"
	"strings"
)

type Client interface {
	common.Client
}

type EndpointManager struct {
	endpoints map[string]string

	AccessControl          *EptAccessControl
	AccessRules            *EptAccessRules
	ActionEngine           *EptActionEngine
	AnalyticsDevice        *EptAnalyticsDevice
	AppMgmt                *EptAppMgmt
	AuthenticationBehavior *EptAuthenticationBehavior
	Credential             *EptCredential
	Device                 *EptDevice
	DeviceIO               *EptDeviceIO
	Display                *EptDisplay
	DoorControl            *EptDoorControl
	Event                  *EptEvent
	FederatedSearch        *EptFederateSearch
	Media                  *EptMedia
	Provisioning           *EptProvision
	Receiver               *EptReceiver
	Recording              *EptRecord
	Replay                 *EptReplay
	Schedule               *EptSchedule
	Search                 *EptSearch
	Security               *EptSecurity
	Thermal                *EptThermal
	Uplink                 *EptUplink
	Analytics              *EptAnalytics
	Imaging                *EptImaging
	PTZ                    *EptPTZ
}

func NewEndpointManager(endpoints map[string]string, client Client) *EndpointManager {
	mgr := &EndpointManager{
		endpoints: endpoints,
	}
	for k, v := range endpoints {
		switch strings.ToLower(k) {
		case "accesscontrol":
			mgr.AccessControl = &EptAccessControl{
				access_control.NewPACSPort(v, client),
			}
		case "accessrules":
			mgr.AccessRules = &EptAccessRules{
				access_rules.NewAccessRulesPort(v, client),
			}
		case "actionengine":
			mgr.ActionEngine = &EptActionEngine{
				action_engine.NewActionEnginePort(v, client),
			}
		case "analyticsdevice":
			mgr.AnalyticsDevice = &EptAnalyticsDevice{
				analytics_device.NewAnalyticsDevicePort(v, client),
			}
		case "appmgmt":
			mgr.AppMgmt = &EptAppMgmt{
				app_mgmt.NewAppManagement(v, client),
			}
		case "authenticationbehavior":
			mgr.AuthenticationBehavior = &EptAuthenticationBehavior{
				authentication_behavior.NewAuthenticationBehaviorPort(v, client),
			}
		case "credential":
			mgr.Credential = &EptCredential{
				credential.NewCredentialPort(v, client),
			}
		case "device":
			mgr.Device = &EptDevice{
				device.NewDevice(v, client),
			}
		case "deviceio":
			mgr.DeviceIO = &EptDeviceIO{
				device_io.NewDeviceIOPort(v, client),
			}
		case "display":
			mgr.Display = &EptDisplay{
				display.NewDisplayPort(v, client),
			}
		case "doorcontrol":
			mgr.DoorControl = &EptDoorControl{
				door_control.NewDoorControlPort(v, client),
			}
		case "events":
			mgr.Event = &EptEvent{
				events.NewPort(v, client),
			}
		case "federatedsearch":
			mgr.FederatedSearch = &EptFederateSearch{
				federated_search.NewFederatedSearchPort(v, client),
			}
		case "media":
			mgr.Media = &EptMedia{
				Media:  media.NewMedia(v, client),
				Media2: media2.NewMedia2(v, client),
			}
		case "provisioning":
			mgr.Provisioning = &EptProvision{
				provisioning.NewProvisioningService(v, client),
			}
		case "receiver":
			mgr.Receiver = &EptReceiver{
				receiver.NewReceiverPort(v, client),
			}
		case "recording":
			mgr.Recording = &EptRecord{
				recording.NewRecordingPort(v, client),
			}
		case "replay":
			mgr.Replay = &EptReplay{
				replay.NewReplayPort(v, client),
			}
		case "schedule":
			mgr.Schedule = &EptSchedule{
				schedule.NewSchedulePort(v, client),
			}
		case "search":
			mgr.Search = &EptSearch{
				search.NewSearchPort(v, client),
			}
		case "security":
			mgr.Security = &EptSecurity{
				security.NewDot1X(v, client),
			}
		case "thermal":
			mgr.Thermal = &EptThermal{
				thermal.NewThermalPort(v, client),
			}
		case "uplink":
			mgr.Uplink = &EptUplink{
				uplink.NewUplinkPort(v, client),
			}
		case "analytics":
			mgr.Analytics = &EptAnalytics{
				analytics.NewAnalyticsEnginePort(v, client),
			}
		case "imaging":
			mgr.Imaging = &EptImaging{
				imaging.NewImagingPort(v, client),
			}
		case "ptz":
			mgr.PTZ = &EptPTZ{
				ptz.NewPTZ(v, client),
			}
		default:
			fmt.Printf("[Endpoint] Unsupported endpoint: [%s : %s]\n", k, v)
		}
	}
	return mgr
}

type EptAccessControl struct {
	access_control.PACSPort
}

type EptAccessRules struct {
	access_rules.AccessRulesPort
}

type EptActionEngine struct {
	action_engine.ActionEnginePort
}

type EptAnalyticsDevice struct {
	analytics_device.AnalyticsDevicePort
}

type EptAppMgmt struct {
	app_mgmt.AppManagement
}

type EptAuthenticationBehavior struct {
	authentication_behavior.AuthenticationBehaviorPort
}

type EptCredential struct {
	credential.CredentialPort
}

type EptDevice struct {
	device.Device
}

type EptDeviceIO struct {
	device_io.DeviceIOPort
}

type EptDisplay struct {
	display.DisplayPort
}

type EptDoorControl struct {
	door_control.DoorControlPort
}

type EptEvent struct {
	events.Port
}

type EptFederateSearch struct {
	federated_search.FederatedSearchPort
}

type EptMedia struct {
	media.Media
	media2.Media2
}

type EptProvision struct {
	provisioning.ProvisioningService
}

type EptReceiver struct {
	receiver.ReceiverPort
}

type EptRecord struct {
	recording.RecordingPort
}

type EptReplay struct {
	replay.ReplayPort
}

type EptSchedule struct {
	schedule.SchedulePort
}

type EptSearch struct {
	search.SearchPort
}

type EptSecurity struct {
	security.Dot1X
}

type EptThermal struct {
	thermal.ThermalPort
}

type EptUplink struct {
	uplink.UplinkPort
}

type EptAnalytics struct {
	analytics.AnalyticsEnginePort
}

type EptImaging struct {
	imaging.ImagingPort
}

type EptPTZ struct {
	ptz.PTZ
}
