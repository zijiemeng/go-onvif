package onvif

import (
	"fmt"
	"github.com/zijiemeng/go-onvif/onvif/common"
	"github.com/zijiemeng/go-onvif/onvif/v10/access_control"
	"github.com/zijiemeng/go-onvif/onvif/v10/access_rules"
	"github.com/zijiemeng/go-onvif/onvif/v10/action_engine"
	"github.com/zijiemeng/go-onvif/onvif/v10/analytics_device"
	"github.com/zijiemeng/go-onvif/onvif/v10/app_mgmt"
	"github.com/zijiemeng/go-onvif/onvif/v10/authentication_behavior"
	"github.com/zijiemeng/go-onvif/onvif/v10/credential"
	"github.com/zijiemeng/go-onvif/onvif/v10/device"
	"github.com/zijiemeng/go-onvif/onvif/v10/device_io"
	"github.com/zijiemeng/go-onvif/onvif/v10/display"
	"github.com/zijiemeng/go-onvif/onvif/v10/door_control"
	"github.com/zijiemeng/go-onvif/onvif/v10/events"
	"github.com/zijiemeng/go-onvif/onvif/v10/federated_search"
	"github.com/zijiemeng/go-onvif/onvif/v10/media"
	"github.com/zijiemeng/go-onvif/onvif/v10/provisioning"
	"github.com/zijiemeng/go-onvif/onvif/v10/receiver"
	"github.com/zijiemeng/go-onvif/onvif/v10/recording"
	"github.com/zijiemeng/go-onvif/onvif/v10/replay"
	"github.com/zijiemeng/go-onvif/onvif/v10/schedule"
	"github.com/zijiemeng/go-onvif/onvif/v10/search"
	"github.com/zijiemeng/go-onvif/onvif/v10/security"
	"github.com/zijiemeng/go-onvif/onvif/v10/thermal"
	"github.com/zijiemeng/go-onvif/onvif/v10/uplink"
	"github.com/zijiemeng/go-onvif/onvif/v20/analytics"
	"github.com/zijiemeng/go-onvif/onvif/v20/imaging"
	media2 "github.com/zijiemeng/go-onvif/onvif/v20/media"
	"github.com/zijiemeng/go-onvif/onvif/v20/ptz"
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
	for k, _ := range endpoints {
		switch strings.ToLower(k) {
		case "accesscontrol":
			mgr.AccessControl = &EptAccessControl{
				access_control.NewPACSPort(k, client),
			}
		case "accessrules":
			mgr.AccessRules = &EptAccessRules{
				access_rules.NewAccessRulesPort(k, client),
			}
		case "actionengine":
			mgr.ActionEngine = &EptActionEngine{
				action_engine.NewActionEnginePort(k, client),
			}
		case "analyticsdevice":
			mgr.AnalyticsDevice = &EptAnalyticsDevice{
				analytics_device.NewAnalyticsDevicePort(k, client),
			}
		case "appmgmt":
			mgr.AppMgmt = &EptAppMgmt{
				app_mgmt.NewAppManagement(k, client),
			}
		case "authenticationbehavior":
			mgr.AuthenticationBehavior = &EptAuthenticationBehavior{
				authentication_behavior.NewAuthenticationBehaviorPort(k, client),
			}
		case "credential":
			mgr.Credential = &EptCredential{
				credential.NewCredentialPort(k, client),
			}
		case "device":
			mgr.Device = &EptDevice{
				device.NewDevice(k, client),
			}
		case "deviceio":
			mgr.DeviceIO = &EptDeviceIO{
				device_io.NewDeviceIOPort(k, client),
			}
		case "display":
			mgr.Display = &EptDisplay{
				display.NewDisplayPort(k, client),
			}
		case "doorcontrol":
			mgr.DoorControl = &EptDoorControl{
				door_control.NewDoorControlPort(k, client),
			}
		case "events":
			mgr.Event = &EptEvent{
				events.NewPort(k, client),
			}
		case "federatedsearch":
			mgr.FederatedSearch = &EptFederateSearch{
				federated_search.NewFederatedSearchPort(k, client),
			}
		case "media":
			mgr.Media = &EptMedia{
				Media:  media.NewMedia(k, client),
				Media2: media2.NewMedia2(k, client),
			}
		case "provisioning":
			mgr.Provisioning = &EptProvision{
				provisioning.NewProvisioningService(k, client),
			}
		case "receiver":
			mgr.Receiver = &EptReceiver{
				receiver.NewReceiverPort(k, client),
			}
		case "recording":
			mgr.Recording = &EptRecord{
				recording.NewRecordingPort(k, client),
			}
		case "replay":
			mgr.Replay = &EptReplay{
				replay.NewReplayPort(k, client),
			}
		case "schedule":
			mgr.Schedule = &EptSchedule{
				schedule.NewSchedulePort(k, client),
			}
		case "search":
			mgr.Search = &EptSearch{
				search.NewSearchPort(k, client),
			}
		case "security":
			mgr.Security = &EptSecurity{
				security.NewDot1X(k, client),
			}
		case "thermal":
			mgr.Thermal = &EptThermal{
				thermal.NewThermalPort(k, client),
			}
		case "uplink":
			mgr.Uplink = &EptUplink{
				uplink.NewUplinkPort(k, client),
			}
		case "analytics":
			mgr.Analytics = &EptAnalytics{
				analytics.NewAnalyticsEnginePort(k, client),
			}
		case "imaging":
			mgr.Imaging = &EptImaging{
				imaging.NewImagingPort(k, client),
			}
		case "ptz":
			mgr.PTZ = &EptPTZ{
				ptz.NewPTZ(k, client),
			}
		default:
			fmt.Printf("[Endpoint] Unsupported endpoint: [%s : %s]\n", k, k)
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
