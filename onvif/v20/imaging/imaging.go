package imaging

import (
	"reflect"

	"code.byted.org/videoarch/go-onvif/onvif/common"
)

var Namespace = "http://www.onvif.org/ver20/imaging/wsdl"

// NewImagingPort creates an initializes a ImagingPort.
func NewImagingPort(endpoint string, cli common.Client) ImagingPort {
	return &imagingPort{cli: cli, Endpoint: endpoint}
}

// ImagingPort was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type ImagingPort interface {
	OptGetCurrentPreset(GetCurrentPreset GetCurrentPreset) (*GetCurrentPresetResponse, *common.Fault)

	OptGetImagingSettings(GetImagingSettings GetImagingSettings) (*GetImagingSettingsResponse, *common.Fault)

	OptGetMoveOptions(GetMoveOptions GetMoveOptions) (*GetMoveOptionsResponse, *common.Fault)

	OptGetOptions(GetOptions GetOptions) (*GetOptionsResponse, *common.Fault)

	OptGetPresets(GetPresets GetPresets) (*GetPresetsResponse, *common.Fault)

	OptGetServiceCapabilities(GetServiceCapabilities GetServiceCapabilities) (*GetServiceCapabilitiesResponse, *common.Fault)

	OptGetStatus(GetStatus GetStatus) (*GetStatusResponse, *common.Fault)

	OptMove(Move Move) (*MoveResponse, *common.Fault)

	OptSetCurrentPreset(SetCurrentPreset SetCurrentPreset) (*SetCurrentPresetResponse, *common.Fault)

	OptSetImagingSettings(SetImagingSettings SetImagingSettings) (*SetImagingSettingsResponse, *common.Fault)

	OptStop(Stop Stop) (*StopResponse, *common.Fault)
}
type ImagingPresetType string

func (v ImagingPresetType) Validate() bool {
	for _, vv := range []string{
		"Custom",
		"ClearWeather",
		"Cloudy",
		"Fog",
		"Rain",
		"Snowing",
		"Snow",
		"WDR",
		"Shade",
		"Night",
		"Indoor",
		"Fluorescent",
		"Incandescent",
		"Sodium(Natrium)",
		"Sunrise(Horizon)",
		"Sunset(Rear)",
		"ExtremeHot",
		"ExtremeCold",
		"Underwater",
		"CloseUp",
		"Motion",
		"FlickerFree50",
		"FlickerFree60",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type Capabilities []interface{}

type GetCurrentPreset struct {
	VideoSourceToken *common.ReferenceToken `xml:"VideoSourceToken,omitempty" json:"VideoSourceToken,omitempty" yaml:"VideoSourceToken,omitempty"`
}

type GetCurrentPresetResponse struct {
	Preset ImagingPreset `xml:"Preset,omitempty" json:"Preset,omitempty" yaml:"Preset,omitempty"`
}

type GetImagingSettings struct {
	VideoSourceToken *common.ReferenceToken `xml:"VideoSourceToken,omitempty" json:"VideoSourceToken,omitempty" yaml:"VideoSourceToken,omitempty"`
}

type GetImagingSettingsResponse struct {
	ImagingSettings *common.ImagingSettings20 `xml:"ImagingSettings,omitempty" json:"ImagingSettings,omitempty" yaml:"ImagingSettings,omitempty"`
}

type GetMoveOptions struct {
	VideoSourceToken *common.ReferenceToken `xml:"VideoSourceToken,omitempty" json:"VideoSourceToken,omitempty" yaml:"VideoSourceToken,omitempty"`
}

type GetMoveOptionsResponse struct {
	MoveOptions *common.MoveOptions20 `xml:"MoveOptions,omitempty" json:"MoveOptions,omitempty" yaml:"MoveOptions,omitempty"`
}

type GetOptions struct {
	VideoSourceToken *common.ReferenceToken `xml:"VideoSourceToken,omitempty" json:"VideoSourceToken,omitempty" yaml:"VideoSourceToken,omitempty"`
}

type GetOptionsResponse struct {
	ImagingOptions *common.ImagingOptions20 `xml:"ImagingOptions,omitempty" json:"ImagingOptions,omitempty" yaml:"ImagingOptions,omitempty"`
}

type GetPresets struct {
	VideoSourceToken *common.ReferenceToken `xml:"VideoSourceToken,omitempty" json:"VideoSourceToken,omitempty" yaml:"VideoSourceToken,omitempty"`
}

type GetPresetsResponse struct {
	Preset []ImagingPreset `xml:"Preset,omitempty" json:"Preset,omitempty" yaml:"Preset,omitempty"`
}

type GetServiceCapabilities struct {
}

type GetServiceCapabilitiesResponse struct {
	Capabilities Capabilities `xml:"Capabilities,omitempty" json:"Capabilities,omitempty" yaml:"Capabilities,omitempty"`
}

type GetStatus struct {
	VideoSourceToken *common.ReferenceToken `xml:"VideoSourceToken,omitempty" json:"VideoSourceToken,omitempty" yaml:"VideoSourceToken,omitempty"`
}

type GetStatusResponse struct {
	Status *common.ImagingStatus20 `xml:"Status,omitempty" json:"Status,omitempty" yaml:"Status,omitempty"`
}

type ImagingPreset struct {
	Name  *common.Name          `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Token common.ReferenceToken `xml:"token,attr,omitempty" json:"token,attr,omitempty" yaml:"token,attr,omitempty"`
	Type  string                `xml:"type,attr,omitempty" json:"type,attr,omitempty" yaml:"type,attr,omitempty"`
}

type Move struct {
	VideoSourceToken *common.ReferenceToken `xml:"VideoSourceToken,omitempty" json:"VideoSourceToken,omitempty" yaml:"VideoSourceToken,omitempty"`
	Focus            *common.FocusMove      `xml:"Focus,omitempty" json:"Focus,omitempty" yaml:"Focus,omitempty"`
}

type MoveResponse struct {
}

type SetCurrentPreset struct {
	VideoSourceToken *common.ReferenceToken `xml:"VideoSourceToken,omitempty" json:"VideoSourceToken,omitempty" yaml:"VideoSourceToken,omitempty"`
	PresetToken      *common.ReferenceToken `xml:"PresetToken,omitempty" json:"PresetToken,omitempty" yaml:"PresetToken,omitempty"`
}

type SetCurrentPresetResponse struct {
}

type SetImagingSettings struct {
	VideoSourceToken *common.ReferenceToken    `xml:"VideoSourceToken,omitempty" json:"VideoSourceToken,omitempty" yaml:"VideoSourceToken,omitempty"`
	ImagingSettings  *common.ImagingSettings20 `xml:"ImagingSettings,omitempty" json:"ImagingSettings,omitempty" yaml:"ImagingSettings,omitempty"`
	ForcePersistence bool                      `xml:"ForcePersistence,omitempty" json:"ForcePersistence,omitempty" yaml:"ForcePersistence,omitempty"`
}

type SetImagingSettingsResponse struct {
}

type Stop struct {
	VideoSourceToken *common.ReferenceToken `xml:"VideoSourceToken,omitempty" json:"VideoSourceToken,omitempty" yaml:"VideoSourceToken,omitempty"`
}

type StopResponse struct {
}

// imagingPort implements the ImagingPort interface.
type imagingPort struct {
	cli      common.Client
	Endpoint string
}

func (p *imagingPort) OptGetCurrentPreset(args GetCurrentPreset) (*GetCurrentPresetResponse, *common.Fault) {
	req := struct {
		XMLName          string `xml:"timg:GetCurrentPreset"`
		GetCurrentPreset GetCurrentPreset
	}{
		GetCurrentPreset: args,
	}

	resp := GetCurrentPresetResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *imagingPort) OptGetImagingSettings(args GetImagingSettings) (*GetImagingSettingsResponse, *common.Fault) {
	req := struct {
		XMLName            string `xml:"timg:GetImagingSettings"`
		GetImagingSettings GetImagingSettings
	}{
		GetImagingSettings: args,
	}

	resp := GetImagingSettingsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *imagingPort) OptGetMoveOptions(args GetMoveOptions) (*GetMoveOptionsResponse, *common.Fault) {
	req := struct {
		XMLName        string `xml:"timg:GetMoveOptions"`
		GetMoveOptions GetMoveOptions
	}{
		GetMoveOptions: args,
	}

	resp := GetMoveOptionsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *imagingPort) OptGetOptions(args GetOptions) (*GetOptionsResponse, *common.Fault) {
	req := struct {
		XMLName    string `xml:"timg:GetOptions"`
		GetOptions GetOptions
	}{
		GetOptions: args,
	}

	resp := GetOptionsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *imagingPort) OptGetPresets(args GetPresets) (*GetPresetsResponse, *common.Fault) {
	req := struct {
		XMLName    string `xml:"timg:GetPresets"`
		GetPresets GetPresets
	}{
		GetPresets: args,
	}

	resp := GetPresetsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *imagingPort) OptGetServiceCapabilities(args GetServiceCapabilities) (*GetServiceCapabilitiesResponse, *common.Fault) {
	req := struct {
		XMLName                string `xml:"timg:GetServiceCapabilities"`
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

func (p *imagingPort) OptGetStatus(args GetStatus) (*GetStatusResponse, *common.Fault) {
	req := struct {
		XMLName   string `xml:"timg:GetStatus"`
		GetStatus GetStatus
	}{
		GetStatus: args,
	}

	resp := GetStatusResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *imagingPort) OptMove(args Move) (*MoveResponse, *common.Fault) {
	req := struct {
		XMLName string `xml:"timg:Move"`
		Move    Move
	}{
		Move: args,
	}

	resp := MoveResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *imagingPort) OptSetCurrentPreset(args SetCurrentPreset) (*SetCurrentPresetResponse, *common.Fault) {
	req := struct {
		XMLName          string `xml:"timg:SetCurrentPreset"`
		SetCurrentPreset SetCurrentPreset
	}{
		SetCurrentPreset: args,
	}

	resp := SetCurrentPresetResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *imagingPort) OptSetImagingSettings(args SetImagingSettings) (*SetImagingSettingsResponse, *common.Fault) {
	req := struct {
		XMLName            string `xml:"timg:SetImagingSettings"`
		SetImagingSettings SetImagingSettings
	}{
		SetImagingSettings: args,
	}

	resp := SetImagingSettingsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *imagingPort) OptStop(args Stop) (*StopResponse, *common.Fault) {
	req := struct {
		XMLName string `xml:"timg:Stop"`
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
