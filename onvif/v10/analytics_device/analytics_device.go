package analytics_device

import (
	"github.com/zijiemeng/go-onvif/onvif/common"
)

var Namespace = "http://www.onvif.org/ver10/analyticsdevice/wsdl"

// NewAnalyticsDevicePort creates an initializes a AnalyticsDevicePort.
func NewAnalyticsDevicePort(endpoint string, cli common.Client) AnalyticsDevicePort {
	return &analyticsDevicePort{cli: cli, Endpoint: endpoint}
}

// AnalyticsDevicePort was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type AnalyticsDevicePort interface {
	OptCreateAnalyticsEngineControl(CreateAnalyticsEngineControl CreateAnalyticsEngineControl) (*CreateAnalyticsEngineControlResponse, *common.Fault)

	OptCreateAnalyticsEngineInputs(CreateAnalyticsEngineInputs CreateAnalyticsEngineInputs) (*CreateAnalyticsEngineInputsResponse, *common.Fault)

	OptDeleteAnalyticsEngineControl(DeleteAnalyticsEngineControl DeleteAnalyticsEngineControl) (*DeleteAnalyticsEngineControlResponse, *common.Fault)

	OptDeleteAnalyticsEngineInputs(DeleteAnalyticsEngineInputs DeleteAnalyticsEngineInputs) (*DeleteAnalyticsEngineInputsResponse, *common.Fault)

	OptGetAnalyticsDeviceStreamUri(GetAnalyticsDeviceStreamUri GetAnalyticsDeviceStreamUri) (*GetAnalyticsDeviceStreamUriResponse, *common.Fault)

	OptGetAnalyticsEngine(GetAnalyticsEngine GetAnalyticsEngine) (*GetAnalyticsEngineResponse, *common.Fault)

	OptGetAnalyticsEngineControl(GetAnalyticsEngineControl GetAnalyticsEngineControl) (*GetAnalyticsEngineControlResponse, *common.Fault)

	OptGetAnalyticsEngineControls(GetAnalyticsEngineControls GetAnalyticsEngineControls) (*GetAnalyticsEngineControlsResponse, *common.Fault)

	OptGetAnalyticsEngineInput(GetAnalyticsEngineInput GetAnalyticsEngineInput) (*GetAnalyticsEngineInputResponse, *common.Fault)

	OptGetAnalyticsEngineInputs(GetAnalyticsEngineInputs GetAnalyticsEngineInputs) (*GetAnalyticsEngineInputsResponse, *common.Fault)

	OptGetAnalyticsEngines(GetAnalyticsEngines GetAnalyticsEngines) (*GetAnalyticsEnginesResponse, *common.Fault)

	OptGetAnalyticsState(GetAnalyticsState GetAnalyticsState) (*GetAnalyticsStateResponse, *common.Fault)

	OptGetServiceCapabilities(GetServiceCapabilities GetServiceCapabilities) (*GetServiceCapabilitiesResponse, *common.Fault)

	OptGetVideoAnalyticsConfiguration(GetVideoAnalyticsConfiguration GetVideoAnalyticsConfiguration) (*GetVideoAnalyticsConfigurationResponse, *common.Fault)

	OptSetAnalyticsEngineControl(SetAnalyticsEngineControl SetAnalyticsEngineControl) (*SetAnalyticsEngineControlResponse, *common.Fault)

	OptSetAnalyticsEngineInput(SetAnalyticsEngineInput SetAnalyticsEngineInput) (*SetAnalyticsEngineInputResponse, *common.Fault)

	OptSetVideoAnalyticsConfiguration(SetVideoAnalyticsConfiguration SetVideoAnalyticsConfiguration) (*SetVideoAnalyticsConfigurationResponse, *common.Fault)
}
type Capabilities []interface{}

type CreateAnalyticsEngineControl struct {
	Configuration *common.AnalyticsEngineControl `xml:"Configuration,omitempty" json:"Configuration,omitempty" yaml:"Configuration,omitempty"`
}

type CreateAnalyticsEngineControlResponse struct {
	Configuration []common.AnalyticsEngineInput `xml:"Configuration" json:"Configuration" yaml:"Configuration"`
}

type CreateAnalyticsEngineInputs struct {
	Configuration    []common.AnalyticsEngineInput `xml:"Configuration" json:"Configuration" yaml:"Configuration"`
	ForcePersistence []bool                        `xml:"ForcePersistence" json:"ForcePersistence" yaml:"ForcePersistence"`
}

type CreateAnalyticsEngineInputsResponse struct {
	Configuration []common.AnalyticsEngineInput `xml:"Configuration" json:"Configuration" yaml:"Configuration"`
}

type DeleteAnalyticsEngineControl struct {
	ConfigurationToken *common.ReferenceToken `xml:"ConfigurationToken,omitempty" json:"ConfigurationToken,omitempty" yaml:"ConfigurationToken,omitempty"`
}

type DeleteAnalyticsEngineControlResponse struct {
}

type DeleteAnalyticsEngineInputs struct {
	ConfigurationToken []common.ReferenceToken `xml:"ConfigurationToken" json:"ConfigurationToken" yaml:"ConfigurationToken"`
}

type DeleteAnalyticsEngineInputsResponse struct {
}

type GetAnalyticsDeviceStreamUri struct {
	StreamSetup                 *common.StreamSetup    `xml:"StreamSetup,omitempty" json:"StreamSetup,omitempty" yaml:"StreamSetup,omitempty"`
	AnalyticsEngineControlToken *common.ReferenceToken `xml:"AnalyticsEngineControlToken,omitempty" json:"AnalyticsEngineControlToken,omitempty" yaml:"AnalyticsEngineControlToken,omitempty"`
}

type GetAnalyticsDeviceStreamUriResponse struct {
	Uri string `xml:"Uri,omitempty" json:"Uri,omitempty" yaml:"Uri,omitempty"`
}

type GetAnalyticsEngine struct {
	ConfigurationToken *common.ReferenceToken `xml:"ConfigurationToken,omitempty" json:"ConfigurationToken,omitempty" yaml:"ConfigurationToken,omitempty"`
}

type GetAnalyticsEngineControl struct {
	ConfigurationToken *common.ReferenceToken `xml:"ConfigurationToken,omitempty" json:"ConfigurationToken,omitempty" yaml:"ConfigurationToken,omitempty"`
}

type GetAnalyticsEngineControlResponse struct {
	Configuration *common.AnalyticsEngineControl `xml:"Configuration,omitempty" json:"Configuration,omitempty" yaml:"Configuration,omitempty"`
}

type GetAnalyticsEngineControls struct {
}

type GetAnalyticsEngineControlsResponse struct {
	AnalyticsEngineControls []common.AnalyticsEngineControl `xml:"AnalyticsEngineControls" json:"AnalyticsEngineControls" yaml:"AnalyticsEngineControls"`
}

type GetAnalyticsEngineInput struct {
	ConfigurationToken *common.ReferenceToken `xml:"ConfigurationToken,omitempty" json:"ConfigurationToken,omitempty" yaml:"ConfigurationToken,omitempty"`
}

type GetAnalyticsEngineInputResponse struct {
	Configuration *common.AnalyticsEngineInput `xml:"Configuration,omitempty" json:"Configuration,omitempty" yaml:"Configuration,omitempty"`
}

type GetAnalyticsEngineInputs struct {
}

type GetAnalyticsEngineInputsResponse struct {
	Configuration []common.AnalyticsEngineInput `xml:"Configuration" json:"Configuration" yaml:"Configuration"`
}

type GetAnalyticsEngineResponse struct {
	Configuration *common.AnalyticsEngine `xml:"Configuration,omitempty" json:"Configuration,omitempty" yaml:"Configuration,omitempty"`
}

type GetAnalyticsEngines struct {
}

type GetAnalyticsEnginesResponse struct {
	Configuration []common.AnalyticsEngine `xml:"Configuration" json:"Configuration" yaml:"Configuration"`
}

type GetAnalyticsState struct {
	AnalyticsEngineControlToken *common.ReferenceToken `xml:"AnalyticsEngineControlToken,omitempty" json:"AnalyticsEngineControlToken,omitempty" yaml:"AnalyticsEngineControlToken,omitempty"`
}

type GetAnalyticsStateResponse struct {
	State *common.AnalyticsStateInformation `xml:"State,omitempty" json:"State,omitempty" yaml:"State,omitempty"`
}

type GetServiceCapabilities struct {
}

type GetServiceCapabilitiesResponse struct {
	Capabilities Capabilities `xml:"Capabilities,omitempty" json:"Capabilities,omitempty" yaml:"Capabilities,omitempty"`
}

type GetVideoAnalyticsConfiguration struct {
	ConfigurationToken *common.ReferenceToken `xml:"ConfigurationToken,omitempty" json:"ConfigurationToken,omitempty" yaml:"ConfigurationToken,omitempty"`
}

type GetVideoAnalyticsConfigurationResponse struct {
	Configuration *common.VideoAnalyticsConfiguration `xml:"Configuration,omitempty" json:"Configuration,omitempty" yaml:"Configuration,omitempty"`
}

type SetAnalyticsEngineControl struct {
	Configuration    *common.AnalyticsEngineControl `xml:"Configuration,omitempty" json:"Configuration,omitempty" yaml:"Configuration,omitempty"`
	ForcePersistence bool                           `xml:"ForcePersistence,omitempty" json:"ForcePersistence,omitempty" yaml:"ForcePersistence,omitempty"`
}

type SetAnalyticsEngineControlResponse struct {
}

type SetAnalyticsEngineInput struct {
	Configuration    *common.AnalyticsEngineInput `xml:"Configuration,omitempty" json:"Configuration,omitempty" yaml:"Configuration,omitempty"`
	ForcePersistence bool                         `xml:"ForcePersistence,omitempty" json:"ForcePersistence,omitempty" yaml:"ForcePersistence,omitempty"`
}

type SetAnalyticsEngineInputResponse struct {
}

type SetVideoAnalyticsConfiguration struct {
	Configuration    *common.VideoAnalyticsConfiguration `xml:"Configuration,omitempty" json:"Configuration,omitempty" yaml:"Configuration,omitempty"`
	ForcePersistence bool                                `xml:"ForcePersistence,omitempty" json:"ForcePersistence,omitempty" yaml:"ForcePersistence,omitempty"`
}

type SetVideoAnalyticsConfigurationResponse struct {
}

// analyticsDevicePort implements the AnalyticsDevicePort interface.
type analyticsDevicePort struct {
	cli      common.Client
	Endpoint string
}

func (p *analyticsDevicePort) OptCreateAnalyticsEngineControl(args CreateAnalyticsEngineControl) (*CreateAnalyticsEngineControlResponse, *common.Fault) {
	req := struct {
		XMLName                      string `xml:"tad:CreateAnalyticsEngineControl"`
		CreateAnalyticsEngineControl CreateAnalyticsEngineControl
	}{
		CreateAnalyticsEngineControl: args,
	}

	resp := CreateAnalyticsEngineControlResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *analyticsDevicePort) OptCreateAnalyticsEngineInputs(args CreateAnalyticsEngineInputs) (*CreateAnalyticsEngineInputsResponse, *common.Fault) {
	req := struct {
		XMLName                     string `xml:"tad:CreateAnalyticsEngineInputs"`
		CreateAnalyticsEngineInputs CreateAnalyticsEngineInputs
	}{
		CreateAnalyticsEngineInputs: args,
	}

	resp := CreateAnalyticsEngineInputsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *analyticsDevicePort) OptDeleteAnalyticsEngineControl(args DeleteAnalyticsEngineControl) (*DeleteAnalyticsEngineControlResponse, *common.Fault) {
	req := struct {
		XMLName                      string `xml:"tad:DeleteAnalyticsEngineControl"`
		DeleteAnalyticsEngineControl DeleteAnalyticsEngineControl
	}{
		DeleteAnalyticsEngineControl: args,
	}

	resp := DeleteAnalyticsEngineControlResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *analyticsDevicePort) OptDeleteAnalyticsEngineInputs(args DeleteAnalyticsEngineInputs) (*DeleteAnalyticsEngineInputsResponse, *common.Fault) {
	req := struct {
		XMLName                     string `xml:"tad:DeleteAnalyticsEngineInputs"`
		DeleteAnalyticsEngineInputs DeleteAnalyticsEngineInputs
	}{
		DeleteAnalyticsEngineInputs: args,
	}

	resp := DeleteAnalyticsEngineInputsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *analyticsDevicePort) OptGetAnalyticsDeviceStreamUri(args GetAnalyticsDeviceStreamUri) (*GetAnalyticsDeviceStreamUriResponse, *common.Fault) {
	req := struct {
		XMLName                     string `xml:"tad:GetAnalyticsDeviceStreamUri"`
		GetAnalyticsDeviceStreamUri GetAnalyticsDeviceStreamUri
	}{
		GetAnalyticsDeviceStreamUri: args,
	}

	resp := GetAnalyticsDeviceStreamUriResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *analyticsDevicePort) OptGetAnalyticsEngine(args GetAnalyticsEngine) (*GetAnalyticsEngineResponse, *common.Fault) {
	req := struct {
		XMLName            string `xml:"tad:GetAnalyticsEngine"`
		GetAnalyticsEngine GetAnalyticsEngine
	}{
		GetAnalyticsEngine: args,
	}

	resp := GetAnalyticsEngineResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *analyticsDevicePort) OptGetAnalyticsEngineControl(args GetAnalyticsEngineControl) (*GetAnalyticsEngineControlResponse, *common.Fault) {
	req := struct {
		XMLName                   string `xml:"tad:GetAnalyticsEngineControl"`
		GetAnalyticsEngineControl GetAnalyticsEngineControl
	}{
		GetAnalyticsEngineControl: args,
	}

	resp := GetAnalyticsEngineControlResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *analyticsDevicePort) OptGetAnalyticsEngineControls(args GetAnalyticsEngineControls) (*GetAnalyticsEngineControlsResponse, *common.Fault) {
	req := struct {
		XMLName                    string `xml:"tad:GetAnalyticsEngineControls"`
		GetAnalyticsEngineControls GetAnalyticsEngineControls
	}{
		GetAnalyticsEngineControls: args,
	}

	resp := GetAnalyticsEngineControlsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *analyticsDevicePort) OptGetAnalyticsEngineInput(args GetAnalyticsEngineInput) (*GetAnalyticsEngineInputResponse, *common.Fault) {
	req := struct {
		XMLName                 string `xml:"tad:GetAnalyticsEngineInput"`
		GetAnalyticsEngineInput GetAnalyticsEngineInput
	}{
		GetAnalyticsEngineInput: args,
	}

	resp := GetAnalyticsEngineInputResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *analyticsDevicePort) OptGetAnalyticsEngineInputs(args GetAnalyticsEngineInputs) (*GetAnalyticsEngineInputsResponse, *common.Fault) {
	req := struct {
		XMLName                  string `xml:"tad:GetAnalyticsEngineInputs"`
		GetAnalyticsEngineInputs GetAnalyticsEngineInputs
	}{
		GetAnalyticsEngineInputs: args,
	}

	resp := GetAnalyticsEngineInputsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *analyticsDevicePort) OptGetAnalyticsEngines(args GetAnalyticsEngines) (*GetAnalyticsEnginesResponse, *common.Fault) {
	req := struct {
		XMLName             string `xml:"tad:GetAnalyticsEngines"`
		GetAnalyticsEngines GetAnalyticsEngines
	}{
		GetAnalyticsEngines: args,
	}

	resp := GetAnalyticsEnginesResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *analyticsDevicePort) OptGetAnalyticsState(args GetAnalyticsState) (*GetAnalyticsStateResponse, *common.Fault) {
	req := struct {
		XMLName           string `xml:"tad:GetAnalyticsState"`
		GetAnalyticsState GetAnalyticsState
	}{
		GetAnalyticsState: args,
	}

	resp := GetAnalyticsStateResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *analyticsDevicePort) OptGetServiceCapabilities(args GetServiceCapabilities) (*GetServiceCapabilitiesResponse, *common.Fault) {
	req := struct {
		XMLName                string `xml:"tad:GetServiceCapabilities"`
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

func (p *analyticsDevicePort) OptGetVideoAnalyticsConfiguration(args GetVideoAnalyticsConfiguration) (*GetVideoAnalyticsConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                        string `xml:"tad:GetVideoAnalyticsConfiguration"`
		GetVideoAnalyticsConfiguration GetVideoAnalyticsConfiguration
	}{
		GetVideoAnalyticsConfiguration: args,
	}

	resp := GetVideoAnalyticsConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *analyticsDevicePort) OptSetAnalyticsEngineControl(args SetAnalyticsEngineControl) (*SetAnalyticsEngineControlResponse, *common.Fault) {
	req := struct {
		XMLName                   string `xml:"tad:SetAnalyticsEngineControl"`
		SetAnalyticsEngineControl SetAnalyticsEngineControl
	}{
		SetAnalyticsEngineControl: args,
	}

	resp := SetAnalyticsEngineControlResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *analyticsDevicePort) OptSetAnalyticsEngineInput(args SetAnalyticsEngineInput) (*SetAnalyticsEngineInputResponse, *common.Fault) {
	req := struct {
		XMLName                 string `xml:"tad:SetAnalyticsEngineInput"`
		SetAnalyticsEngineInput SetAnalyticsEngineInput
	}{
		SetAnalyticsEngineInput: args,
	}

	resp := SetAnalyticsEngineInputResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *analyticsDevicePort) OptSetVideoAnalyticsConfiguration(args SetVideoAnalyticsConfiguration) (*SetVideoAnalyticsConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                        string `xml:"tad:SetVideoAnalyticsConfiguration"`
		SetVideoAnalyticsConfiguration SetVideoAnalyticsConfiguration
	}{
		SetVideoAnalyticsConfiguration: args,
	}

	resp := SetVideoAnalyticsConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}
