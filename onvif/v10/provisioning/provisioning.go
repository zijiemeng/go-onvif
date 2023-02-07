package provisioning

import (
	"reflect"

	"github.com/zijiemeng/go-onvif/onvif/common"
)

var Namespace = "http://www.onvif.org/ver10/provisioning/wsdl"

// NewProvisioningService creates an initializes a ProvisioningService.
func NewProvisioningService(endpoint string, cli common.Client) ProvisioningService {
	return &provisioningService{cli: cli, Endpoint: endpoint}
}

// ProvisioningService was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type ProvisioningService interface {
	OptFocusMove(FocusMove FocusMove) (*FocusMoveResponse, *common.Fault)

	OptGetServiceCapabilities(GetServiceCapabilities GetServiceCapabilities) (*GetServiceCapabilitiesResponse, *common.Fault)

	OptGetUsage(GetUsage GetUsage) (*GetUsageResponse, *common.Fault)

	OptPanMove(PanMove PanMove) (*PanMoveResponse, *common.Fault)

	OptRollMove(RollMove RollMove) (*RollMoveResponse, *common.Fault)

	OptStop(Stop Stop) (*StopResponse, *common.Fault)

	OptTiltMove(TiltMove TiltMove) (*TiltMoveResponse, *common.Fault)

	OptZoomMove(ZoomMove ZoomMove) (*ZoomMoveResponse, *common.Fault)
}
type Duration string

type FocusDirection string

func (v FocusDirection) Validate() bool {
	for _, vv := range []string{
		"Near",
		"Far",
		"Auto",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type PanDirection string

func (v PanDirection) Validate() bool {
	for _, vv := range []string{
		"Left",
		"Right",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type RollDirection string

func (v RollDirection) Validate() bool {
	for _, vv := range []string{
		"Clockwise",
		"Counterclockwise",
		"Auto",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type TiltDirection string

func (v TiltDirection) Validate() bool {
	for _, vv := range []string{
		"Up",
		"Down",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type ZoomDirection string

func (v ZoomDirection) Validate() bool {
	for _, vv := range []string{
		"Wide",
		"Telephoto",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type Capabilities struct {
	DefaultTimeout *common.Duration     `xml:"DefaultTimeout,omitempty" json:"DefaultTimeout,omitempty" yaml:"DefaultTimeout,omitempty"`
	Source         []SourceCapabilities `xml:"Source,omitempty" json:"Source,omitempty" yaml:"Source,omitempty"`
}

type FocusMove struct {
	VideoSource *common.ReferenceToken `xml:"VideoSource,omitempty" json:"VideoSource,omitempty" yaml:"VideoSource,omitempty"`
	Direction   FocusDirection         `xml:"Direction,omitempty" json:"Direction,omitempty" yaml:"Direction,omitempty"`
	Timeout     *common.Duration       `xml:"Timeout,omitempty" json:"Timeout,omitempty" yaml:"Timeout,omitempty"`
}

type FocusMoveResponse struct {
}

type GetServiceCapabilities struct {
}

type GetServiceCapabilitiesResponse struct {
	Capabilities Capabilities `xml:"Capabilities,omitempty" json:"Capabilities,omitempty" yaml:"Capabilities,omitempty"`
}

type GetUsage struct {
	VideoSource *common.ReferenceToken `xml:"VideoSource,omitempty" json:"VideoSource,omitempty" yaml:"VideoSource,omitempty"`
}

type GetUsageResponse struct {
	Usage Usage `xml:"Usage,omitempty" json:"Usage,omitempty" yaml:"Usage,omitempty"`
}

type PanMove struct {
	VideoSource *common.ReferenceToken `xml:"VideoSource,omitempty" json:"VideoSource,omitempty" yaml:"VideoSource,omitempty"`
	Direction   PanDirection           `xml:"Direction,omitempty" json:"Direction,omitempty" yaml:"Direction,omitempty"`
	Timeout     *common.Duration       `xml:"Timeout,omitempty" json:"Timeout,omitempty" yaml:"Timeout,omitempty"`
}

type PanMoveResponse struct {
}

type RollMove struct {
	VideoSource *common.ReferenceToken `xml:"VideoSource,omitempty" json:"VideoSource,omitempty" yaml:"VideoSource,omitempty"`
	Direction   RollDirection          `xml:"Direction,omitempty" json:"Direction,omitempty" yaml:"Direction,omitempty"`
	Timeout     *common.Duration       `xml:"Timeout,omitempty" json:"Timeout,omitempty" yaml:"Timeout,omitempty"`
}

type RollMoveResponse struct {
}

type SourceCapabilities []interface{}

type Stop struct {
	VideoSource *common.ReferenceToken `xml:"VideoSource,omitempty" json:"VideoSource,omitempty" yaml:"VideoSource,omitempty"`
}

type StopResponse struct {
}

type TiltMove struct {
	VideoSource *common.ReferenceToken `xml:"VideoSource,omitempty" json:"VideoSource,omitempty" yaml:"VideoSource,omitempty"`
	Direction   TiltDirection          `xml:"Direction,omitempty" json:"Direction,omitempty" yaml:"Direction,omitempty"`
	Timeout     *common.Duration       `xml:"Timeout,omitempty" json:"Timeout,omitempty" yaml:"Timeout,omitempty"`
}

type TiltMoveResponse struct {
}

type Usage struct {
	Pan   uint64 `xml:"Pan,omitempty" json:"Pan,omitempty" yaml:"Pan,omitempty"`
	Tilt  uint64 `xml:"Tilt,omitempty" json:"Tilt,omitempty" yaml:"Tilt,omitempty"`
	Zoom  uint64 `xml:"Zoom,omitempty" json:"Zoom,omitempty" yaml:"Zoom,omitempty"`
	Roll  uint64 `xml:"Roll,omitempty" json:"Roll,omitempty" yaml:"Roll,omitempty"`
	Focus uint64 `xml:"Focus,omitempty" json:"Focus,omitempty" yaml:"Focus,omitempty"`
}

type ZoomMove struct {
	VideoSource *common.ReferenceToken `xml:"VideoSource,omitempty" json:"VideoSource,omitempty" yaml:"VideoSource,omitempty"`
	Direction   ZoomDirection          `xml:"Direction,omitempty" json:"Direction,omitempty" yaml:"Direction,omitempty"`
	Timeout     *common.Duration       `xml:"Timeout,omitempty" json:"Timeout,omitempty" yaml:"Timeout,omitempty"`
}

type ZoomMoveResponse struct {
}

// provisioningService implements the ProvisioningService interface.
type provisioningService struct {
	cli      common.Client
	Endpoint string
}

func (p *provisioningService) OptFocusMove(args FocusMove) (*FocusMoveResponse, *common.Fault) {
	req := struct {
		XMLName   string `xml:"tpv:FocusMove"`
		FocusMove FocusMove
	}{
		FocusMove: args,
	}

	resp := FocusMoveResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *provisioningService) OptGetServiceCapabilities(args GetServiceCapabilities) (*GetServiceCapabilitiesResponse, *common.Fault) {
	req := struct {
		XMLName                string `xml:"tpv:GetServiceCapabilities"`
		GetServiceCapabilities GetServiceCapabilities
	}{
		GetServiceCapabilities: args,
	}

	resp := GetServiceCapabilitiesResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *provisioningService) OptGetUsage(args GetUsage) (*GetUsageResponse, *common.Fault) {
	req := struct {
		XMLName  string `xml:"tpv:GetUsage"`
		GetUsage GetUsage
	}{
		GetUsage: args,
	}

	resp := GetUsageResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *provisioningService) OptPanMove(args PanMove) (*PanMoveResponse, *common.Fault) {
	req := struct {
		XMLName string `xml:"tpv:PanMove"`
		PanMove PanMove
	}{
		PanMove: args,
	}

	resp := PanMoveResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *provisioningService) OptRollMove(args RollMove) (*RollMoveResponse, *common.Fault) {
	req := struct {
		XMLName  string `xml:"tpv:RollMove"`
		RollMove RollMove
	}{
		RollMove: args,
	}

	resp := RollMoveResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *provisioningService) OptStop(args Stop) (*StopResponse, *common.Fault) {
	req := struct {
		XMLName string `xml:"tpv:Stop"`
		Stop    Stop
	}{
		Stop: args,
	}

	resp := StopResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *provisioningService) OptTiltMove(args TiltMove) (*TiltMoveResponse, *common.Fault) {
	req := struct {
		XMLName  string `xml:"tpv:TiltMove"`
		TiltMove TiltMove
	}{
		TiltMove: args,
	}

	resp := TiltMoveResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *provisioningService) OptZoomMove(args ZoomMove) (*ZoomMoveResponse, *common.Fault) {
	req := struct {
		XMLName  string `xml:"tpv:ZoomMove"`
		ZoomMove ZoomMove
	}{
		ZoomMove: args,
	}

	resp := ZoomMoveResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}
