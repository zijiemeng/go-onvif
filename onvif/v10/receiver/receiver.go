package receiver

import (
	"code.byted.org/videoarch/go-onvif/onvif/common"
)

var Namespace = "http://www.onvif.org/ver10/receiver/wsdl"

// NewReceiverPort creates an initializes a ReceiverPort.
func NewReceiverPort(endpoint string, cli common.Client) ReceiverPort {
	return &receiverPort{cli: cli, Endpoint: endpoint}
}

// ReceiverPort was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type ReceiverPort interface {
	OptConfigureReceiver(ConfigureReceiver ConfigureReceiver) (*ConfigureReceiverResponse, error)

	OptCreateReceiver(CreateReceiver CreateReceiver) (*CreateReceiverResponse, error)

	OptDeleteReceiver(DeleteReceiver DeleteReceiver) (*DeleteReceiverResponse, error)

	OptGetReceiver(GetReceiver GetReceiver) (*GetReceiverResponse, error)

	OptGetReceiverState(GetReceiverState GetReceiverState) (*GetReceiverStateResponse, error)

	OptGetReceivers(GetReceivers GetReceivers) (*GetReceiversResponse, error)

	OptGetServiceCapabilities(GetServiceCapabilities GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error)

	OptSetReceiverMode(SetReceiverMode SetReceiverMode) (*SetReceiverModeResponse, error)
}
type Capabilities []interface{}

type ConfigureReceiver struct {
	ReceiverToken *common.ReferenceToken        `xml:"ReceiverToken,omitempty" json:"ReceiverToken,omitempty" yaml:"ReceiverToken,omitempty"`
	Configuration *common.ReceiverConfiguration `xml:"Configuration,omitempty" json:"Configuration,omitempty" yaml:"Configuration,omitempty"`
}

type ConfigureReceiverResponse struct {
}

type CreateReceiver struct {
	Configuration *common.ReceiverConfiguration `xml:"Configuration,omitempty" json:"Configuration,omitempty" yaml:"Configuration,omitempty"`
}

type CreateReceiverResponse struct {
	Receiver *common.Receiver `xml:"Receiver,omitempty" json:"Receiver,omitempty" yaml:"Receiver,omitempty"`
}

type DeleteReceiver struct {
	ReceiverToken *common.ReferenceToken `xml:"ReceiverToken,omitempty" json:"ReceiverToken,omitempty" yaml:"ReceiverToken,omitempty"`
}

type DeleteReceiverResponse struct {
}

type GetReceiver struct {
	ReceiverToken *common.ReferenceToken `xml:"ReceiverToken,omitempty" json:"ReceiverToken,omitempty" yaml:"ReceiverToken,omitempty"`
}

type GetReceiverResponse struct {
	Receiver *common.Receiver `xml:"Receiver,omitempty" json:"Receiver,omitempty" yaml:"Receiver,omitempty"`
}

type GetReceiverState struct {
	ReceiverToken *common.ReferenceToken `xml:"ReceiverToken,omitempty" json:"ReceiverToken,omitempty" yaml:"ReceiverToken,omitempty"`
}

type GetReceiverStateResponse struct {
	ReceiverState *common.ReceiverStateInformation `xml:"ReceiverState,omitempty" json:"ReceiverState,omitempty" yaml:"ReceiverState,omitempty"`
}

type GetReceivers struct {
}

type GetReceiversResponse struct {
	Receivers []*common.Receiver `xml:"Receivers,omitempty" json:"Receivers,omitempty" yaml:"Receivers,omitempty"`
}

type GetServiceCapabilities struct {
}

type GetServiceCapabilitiesResponse struct {
	Capabilities Capabilities `xml:"Capabilities,omitempty" json:"Capabilities,omitempty" yaml:"Capabilities,omitempty"`
}

type SetReceiverMode struct {
	ReceiverToken *common.ReferenceToken `xml:"ReceiverToken,omitempty" json:"ReceiverToken,omitempty" yaml:"ReceiverToken,omitempty"`
	Mode          *common.ReceiverMode   `xml:"Mode,omitempty" json:"Mode,omitempty" yaml:"Mode,omitempty"`
}

type SetReceiverModeResponse struct {
}

// receiverPort implements the ReceiverPort interface.
type receiverPort struct {
	cli      common.Client
	Endpoint string
}

func (p *receiverPort) OptConfigureReceiver(args ConfigureReceiver) (*ConfigureReceiverResponse, error) {
	req := struct {
		XMLName           string `xml:"trv:ConfigureReceiver"`
		ConfigureReceiver ConfigureReceiver
	}{
		ConfigureReceiver: args,
	}

	resp := ConfigureReceiverResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *receiverPort) OptCreateReceiver(args CreateReceiver) (*CreateReceiverResponse, error) {
	req := struct {
		XMLName        string `xml:"trv:CreateReceiver"`
		CreateReceiver CreateReceiver
	}{
		CreateReceiver: args,
	}

	resp := CreateReceiverResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *receiverPort) OptDeleteReceiver(args DeleteReceiver) (*DeleteReceiverResponse, error) {
	req := struct {
		XMLName        string `xml:"trv:DeleteReceiver"`
		DeleteReceiver DeleteReceiver
	}{
		DeleteReceiver: args,
	}

	resp := DeleteReceiverResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *receiverPort) OptGetReceiver(args GetReceiver) (*GetReceiverResponse, error) {
	req := struct {
		XMLName     string `xml:"trv:GetReceiver"`
		GetReceiver GetReceiver
	}{
		GetReceiver: args,
	}

	resp := GetReceiverResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *receiverPort) OptGetReceiverState(args GetReceiverState) (*GetReceiverStateResponse, error) {
	req := struct {
		XMLName          string `xml:"trv:GetReceiverState"`
		GetReceiverState GetReceiverState
	}{
		GetReceiverState: args,
	}

	resp := GetReceiverStateResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *receiverPort) OptGetReceivers(args GetReceivers) (*GetReceiversResponse, error) {
	req := struct {
		XMLName      string `xml:"trv:GetReceivers"`
		GetReceivers GetReceivers
	}{
		GetReceivers: args,
	}

	resp := GetReceiversResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *receiverPort) OptGetServiceCapabilities(args GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error) {
	req := struct {
		XMLName                string `xml:"trv:GetServiceCapabilities"`
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

func (p *receiverPort) OptSetReceiverMode(args SetReceiverMode) (*SetReceiverModeResponse, error) {
	req := struct {
		XMLName         string `xml:"trv:SetReceiverMode"`
		SetReceiverMode SetReceiverMode
	}{
		SetReceiverMode: args,
	}

	resp := SetReceiverModeResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
