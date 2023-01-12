package replay

import (
	"code.byted.org/videoarch/go-onvif/onvif/common"
)

var Namespace = "http://www.onvif.org/ver10/replay/wsdl"

// NewReplayPort creates an initializes a ReplayPort.
func NewReplayPort(endpoint string, cli common.Client) ReplayPort {
	return &replayPort{cli: cli, Endpoint: endpoint}
}

// ReplayPort was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type ReplayPort interface {
	OptGetReplayConfiguration(GetReplayConfiguration GetReplayConfiguration) (*GetReplayConfigurationResponse, error)

	OptGetReplayUri(GetReplayUri GetReplayUri) (*GetReplayUriResponse, error)

	OptGetServiceCapabilities(GetServiceCapabilities GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error)

	OptSetReplayConfiguration(SetReplayConfiguration SetReplayConfiguration) (*SetReplayConfigurationResponse, error)
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

func (p *replayPort) OptGetReplayConfiguration(args GetReplayConfiguration) (*GetReplayConfigurationResponse, error) {
	req := struct {
		XMLName                string `xml:"trp:GetReplayConfiguration"`
		GetReplayConfiguration GetReplayConfiguration
	}{
		GetReplayConfiguration: args,
	}

	resp := GetReplayConfigurationResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *replayPort) OptGetReplayUri(args GetReplayUri) (*GetReplayUriResponse, error) {
	req := struct {
		XMLName      string `xml:"trp:GetReplayUri"`
		GetReplayUri GetReplayUri
	}{
		GetReplayUri: args,
	}

	resp := GetReplayUriResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *replayPort) OptGetServiceCapabilities(args GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error) {
	req := struct {
		XMLName                string `xml:"trp:GetServiceCapabilities"`
		GetServiceCapabilities GetServiceCapabilities
	}{
		GetServiceCapabilities: args,
	}

	resp := GetServiceCapabilitiesResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *replayPort) OptSetReplayConfiguration(args SetReplayConfiguration) (*SetReplayConfigurationResponse, error) {
	req := struct {
		XMLName                string `xml:"trp:SetReplayConfiguration"`
		SetReplayConfiguration SetReplayConfiguration
	}{
		SetReplayConfiguration: args,
	}

	resp := SetReplayConfigurationResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
