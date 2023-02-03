package replay

import (
	"github.com/zijiemeng/go-onvif/onvif/common"
)

var Namespace = "http://www.onvif.org/ver10/replay/wsdl"

// NewReplayPort creates an initializes a ReplayPort.
func NewReplayPort(endpoint string, cli common.Client) ReplayPort {
	return &replayPort{cli: cli, Endpoint: endpoint}
}

// ReplayPort was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type ReplayPort interface {
	OptGetReplayConfiguration(GetReplayConfiguration GetReplayConfiguration) (*GetReplayConfigurationResponse, *common.Fault)

	OptGetReplayUri(GetReplayUri GetReplayUri) (*GetReplayUriResponse, *common.Fault)

	OptGetServiceCapabilities(GetServiceCapabilities GetServiceCapabilities) (*GetServiceCapabilitiesResponse, *common.Fault)

	OptSetReplayConfiguration(SetReplayConfiguration SetReplayConfiguration) (*SetReplayConfigurationResponse, *common.Fault)
}
type Capabilities []interface{}

type GetReplayConfiguration struct {
}

type GetReplayConfigurationResponse struct {
	Configuration *common.ReplayConfiguration `xml:"Configuration,omitempty" json:"Configuration,omitempty" yaml:"Configuration,omitempty"`
}

type GetReplayUri struct {
	StreamSetup    *common.StreamSetup    `xml:"StreamSetup,omitempty" json:"StreamSetup,omitempty" yaml:"StreamSetup,omitempty"`
	RecordingToken *common.ReferenceToken `xml:"RecordingToken,omitempty" json:"RecordingToken,omitempty" yaml:"RecordingToken,omitempty"`
}

type GetReplayUriResponse struct {
	Uri string `xml:"Uri,omitempty" json:"Uri,omitempty" yaml:"Uri,omitempty"`
}

type GetServiceCapabilities struct {
}

type GetServiceCapabilitiesResponse struct {
	Capabilities Capabilities `xml:"Capabilities,omitempty" json:"Capabilities,omitempty" yaml:"Capabilities,omitempty"`
}

type SetReplayConfiguration struct {
	Configuration *common.ReplayConfiguration `xml:"Configuration,omitempty" json:"Configuration,omitempty" yaml:"Configuration,omitempty"`
}

type SetReplayConfigurationResponse struct {
}

// replayPort implements the ReplayPort interface.
type replayPort struct {
	cli      common.Client
	Endpoint string
}

func (p *replayPort) OptGetReplayConfiguration(args GetReplayConfiguration) (*GetReplayConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                string `xml:"trp:GetReplayConfiguration"`
		GetReplayConfiguration GetReplayConfiguration
	}{
		GetReplayConfiguration: args,
	}

	resp := GetReplayConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *replayPort) OptGetReplayUri(args GetReplayUri) (*GetReplayUriResponse, *common.Fault) {
	req := struct {
		XMLName      string `xml:"trp:GetReplayUri"`
		GetReplayUri GetReplayUri
	}{
		GetReplayUri: args,
	}

	resp := GetReplayUriResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *replayPort) OptGetServiceCapabilities(args GetServiceCapabilities) (*GetServiceCapabilitiesResponse, *common.Fault) {
	req := struct {
		XMLName                string `xml:"trp:GetServiceCapabilities"`
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

func (p *replayPort) OptSetReplayConfiguration(args SetReplayConfiguration) (*SetReplayConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                string `xml:"trp:SetReplayConfiguration"`
		SetReplayConfiguration SetReplayConfiguration
	}{
		SetReplayConfiguration: args,
	}

	resp := SetReplayConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}
