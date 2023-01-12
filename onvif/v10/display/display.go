package display

import (
	"code.byted.org/videoarch/go-onvif/onvif/common"
)

var Namespace = "http://www.onvif.org/ver10/display/wsdl"

// NewDisplayPort creates an initializes a DisplayPort.
func NewDisplayPort(endpoint string, cli common.Client) DisplayPort {
	return &displayPort{cli: cli, Endpoint: endpoint}
}

// DisplayPort was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type DisplayPort interface {
	OptCreatePaneConfiguration(CreatePaneConfiguration CreatePaneConfiguration) (*CreatePaneConfigurationResponse, error)

	OptDeletePaneConfiguration(DeletePaneConfiguration DeletePaneConfiguration) (*DeletePaneConfigurationResponse, error)

	OptGetDisplayOptions(GetDisplayOptions GetDisplayOptions) (*GetDisplayOptionsResponse, error)

	OptGetLayout(GetLayout GetLayout) (*GetLayoutResponse, error)

	OptGetPaneConfiguration(GetPaneConfiguration GetPaneConfiguration) (*GetPaneConfigurationResponse, error)

	OptGetPaneConfigurations(GetPaneConfigurations GetPaneConfigurations) (*GetPaneConfigurationsResponse, error)

	OptSetLayout(SetLayout SetLayout) (*SetLayoutResponse, error)

	OptSetPaneConfiguration(SetPaneConfiguration SetPaneConfiguration) (*SetPaneConfigurationResponse, error)

	OptSetPaneConfigurations(SetPaneConfigurations SetPaneConfigurations) (*SetPaneConfigurationsResponse, error)
}
type CreatePaneConfiguration struct {
	VideoOutput       *common.ReferenceToken    `xml:"VideoOutput,omitempty" json:"VideoOutput,omitempty" yaml:"VideoOutput,omitempty"`
	PaneConfiguration *common.PaneConfiguration `xml:"PaneConfiguration,omitempty" json:"PaneConfiguration,omitempty" yaml:"PaneConfiguration,omitempty"`
}

type CreatePaneConfigurationResponse []interface{}

type DeletePaneConfiguration struct {
	VideoOutput *common.ReferenceToken `xml:"VideoOutput,omitempty" json:"VideoOutput,omitempty" yaml:"VideoOutput,omitempty"`
	PaneToken   *common.ReferenceToken `xml:"PaneToken,omitempty" json:"PaneToken,omitempty" yaml:"PaneToken,omitempty"`
}

type DeletePaneConfigurationResponse []interface{}

type GetDisplayOptions struct {
	VideoOutput *common.ReferenceToken `xml:"VideoOutput,omitempty" json:"VideoOutput,omitempty" yaml:"VideoOutput,omitempty"`
}

type GetDisplayOptionsResponse struct {
	LayoutOptions      *common.LayoutOptions      `xml:"LayoutOptions,omitempty" json:"LayoutOptions,omitempty" yaml:"LayoutOptions,omitempty"`
	CodingCapabilities *common.CodingCapabilities `xml:"CodingCapabilities,omitempty" json:"CodingCapabilities,omitempty" yaml:"CodingCapabilities,omitempty"`
}

type GetLayout struct {
	VideoOutput *common.ReferenceToken `xml:"VideoOutput,omitempty" json:"VideoOutput,omitempty" yaml:"VideoOutput,omitempty"`
}

type GetLayoutResponse struct {
	Layout *common.Layout `xml:"Layout,omitempty" json:"Layout,omitempty" yaml:"Layout,omitempty"`
}

type GetPaneConfiguration struct {
	VideoOutput *common.ReferenceToken `xml:"VideoOutput,omitempty" json:"VideoOutput,omitempty" yaml:"VideoOutput,omitempty"`
	Pane        *common.ReferenceToken `xml:"Pane,omitempty" json:"Pane,omitempty" yaml:"Pane,omitempty"`
}

type GetPaneConfigurationResponse struct {
	PaneConfiguration *common.PaneConfiguration `xml:"PaneConfiguration,omitempty" json:"PaneConfiguration,omitempty" yaml:"PaneConfiguration,omitempty"`
}

type GetPaneConfigurations struct {
	VideoOutput *common.ReferenceToken `xml:"VideoOutput,omitempty" json:"VideoOutput,omitempty" yaml:"VideoOutput,omitempty"`
}

type GetPaneConfigurationsResponse struct {
	PaneConfiguration []*common.PaneConfiguration `xml:"PaneConfiguration,omitempty" json:"PaneConfiguration,omitempty" yaml:"PaneConfiguration,omitempty"`
}

type SetLayout struct {
	VideoOutput *common.ReferenceToken `xml:"VideoOutput,omitempty" json:"VideoOutput,omitempty" yaml:"VideoOutput,omitempty"`
	Layout      *common.Layout         `xml:"Layout,omitempty" json:"Layout,omitempty" yaml:"Layout,omitempty"`
}

type SetLayoutResponse []interface{}

type SetPaneConfiguration struct {
	VideoOutput       *common.ReferenceToken    `xml:"VideoOutput,omitempty" json:"VideoOutput,omitempty" yaml:"VideoOutput,omitempty"`
	PaneConfiguration *common.PaneConfiguration `xml:"PaneConfiguration,omitempty" json:"PaneConfiguration,omitempty" yaml:"PaneConfiguration,omitempty"`
}

type SetPaneConfigurationResponse []interface{}

type SetPaneConfigurations struct {
	VideoOutput       *common.ReferenceToken      `xml:"VideoOutput,omitempty" json:"VideoOutput,omitempty" yaml:"VideoOutput,omitempty"`
	PaneConfiguration []*common.PaneConfiguration `xml:"PaneConfiguration,omitempty" json:"PaneConfiguration,omitempty" yaml:"PaneConfiguration,omitempty"`
}

type SetPaneConfigurationsResponse []interface{}

// displayPort implements the DisplayPort interface.
type displayPort struct {
	cli      common.Client
	Endpoint string
}

func (p *displayPort) OptCreatePaneConfiguration(args CreatePaneConfiguration) (*CreatePaneConfigurationResponse, error) {
	req := struct {
		XMLName                 string `xml:"tls:CreatePaneConfiguration"`
		CreatePaneConfiguration CreatePaneConfiguration
	}{
		CreatePaneConfiguration: args,
	}

	resp := CreatePaneConfigurationResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *displayPort) OptDeletePaneConfiguration(args DeletePaneConfiguration) (*DeletePaneConfigurationResponse, error) {
	req := struct {
		XMLName                 string `xml:"tls:DeletePaneConfiguration"`
		DeletePaneConfiguration DeletePaneConfiguration
	}{
		DeletePaneConfiguration: args,
	}

	resp := DeletePaneConfigurationResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *displayPort) OptGetDisplayOptions(args GetDisplayOptions) (*GetDisplayOptionsResponse, error) {
	req := struct {
		XMLName           string `xml:"tls:GetDisplayOptions"`
		GetDisplayOptions GetDisplayOptions
	}{
		GetDisplayOptions: args,
	}

	resp := GetDisplayOptionsResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *displayPort) OptGetLayout(args GetLayout) (*GetLayoutResponse, error) {
	req := struct {
		XMLName   string `xml:"tls:GetLayout"`
		GetLayout GetLayout
	}{
		GetLayout: args,
	}

	resp := GetLayoutResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *displayPort) OptGetPaneConfiguration(args GetPaneConfiguration) (*GetPaneConfigurationResponse, error) {
	req := struct {
		XMLName              string `xml:"tls:GetPaneConfiguration"`
		GetPaneConfiguration GetPaneConfiguration
	}{
		GetPaneConfiguration: args,
	}

	resp := GetPaneConfigurationResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *displayPort) OptGetPaneConfigurations(args GetPaneConfigurations) (*GetPaneConfigurationsResponse, error) {
	req := struct {
		XMLName               string `xml:"tls:GetPaneConfigurations"`
		GetPaneConfigurations GetPaneConfigurations
	}{
		GetPaneConfigurations: args,
	}

	resp := GetPaneConfigurationsResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *displayPort) OptSetLayout(args SetLayout) (*SetLayoutResponse, error) {
	req := struct {
		XMLName   string `xml:"tls:SetLayout"`
		SetLayout SetLayout
	}{
		SetLayout: args,
	}

	resp := SetLayoutResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *displayPort) OptSetPaneConfiguration(args SetPaneConfiguration) (*SetPaneConfigurationResponse, error) {
	req := struct {
		XMLName              string `xml:"tls:SetPaneConfiguration"`
		SetPaneConfiguration SetPaneConfiguration
	}{
		SetPaneConfiguration: args,
	}

	resp := SetPaneConfigurationResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *displayPort) OptSetPaneConfigurations(args SetPaneConfigurations) (*SetPaneConfigurationsResponse, error) {
	req := struct {
		XMLName               string `xml:"tls:SetPaneConfigurations"`
		SetPaneConfigurations SetPaneConfigurations
	}{
		SetPaneConfigurations: args,
	}

	resp := SetPaneConfigurationsResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
