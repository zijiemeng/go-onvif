package uplink

import (
	"reflect"

	"code.byted.org/videoarch/go-onvif/onvif/common"
)

var Namespace = "http://www.onvif.org/ver10/uplink/wsdl"

// NewUplinkPort creates an initializes a UplinkPort.
func NewUplinkPort(endpoint string, cli common.Client) UplinkPort {
	return &uplinkPort{cli: cli, Endpoint: endpoint}
}

// UplinkPort was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type UplinkPort interface {
	OptDeleteUplink(DeleteUplink DeleteUplink) (*DeleteUplinkResponse, error)

	OptGetServiceCapabilities(GetServiceCapabilities GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error)

	OptGetUplinks(GetUplinks GetUplinks) (*GetUplinksResponse, error)

	OptSetUplink(SetUplink SetUplink) (*SetUplinkResponse, error)
}
type ConnectionStatus string

func (v ConnectionStatus) Validate() bool {
	for _, vv := range []string{
		"Offline",
		"Connecting",
		"Connected",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type Capabilities []interface{}

type Configuration struct {
	RemoteAddress              string `xml:"RemoteAddress,omitempty" json:"RemoteAddress,omitempty" yaml:"RemoteAddress,omitempty"`
	CertificateID              string `xml:"CertificateID,omitempty" json:"CertificateID,omitempty" yaml:"CertificateID,omitempty"`
	UserLevel                  string `xml:"UserLevel,omitempty" json:"UserLevel,omitempty" yaml:"UserLevel,omitempty"`
	Status                     string `xml:"Status,omitempty" json:"Status,omitempty" yaml:"Status,omitempty"`
	CertPathValidationPolicyID string `xml:"CertPathValidationPolicyID,omitempty" json:"CertPathValidationPolicyID,omitempty" yaml:"CertPathValidationPolicyID,omitempty"`
}

type DeleteUplink struct {
	RemoteAddress string `xml:"RemoteAddress,omitempty" json:"RemoteAddress,omitempty" yaml:"RemoteAddress,omitempty"`
}

type DeleteUplinkResponse struct {
}

type GetServiceCapabilities struct {
}

type GetServiceCapabilitiesResponse struct {
	Capabilities Capabilities `xml:"Capabilities,omitempty" json:"Capabilities,omitempty" yaml:"Capabilities,omitempty"`
}

type GetUplinks struct {
}

type GetUplinksResponse struct {
	Configuration []Configuration `xml:"Configuration,omitempty" json:"Configuration,omitempty" yaml:"Configuration,omitempty"`
}

type SetUplink struct {
	Configuration Configuration `xml:"Configuration,omitempty" json:"Configuration,omitempty" yaml:"Configuration,omitempty"`
}

type SetUplinkResponse struct {
}

// uplinkPort implements the UplinkPort interface.
type uplinkPort struct {
	cli      common.Client
	Endpoint string
}

func (p *uplinkPort) OptDeleteUplink(args DeleteUplink) (*DeleteUplinkResponse, error) {
	req := struct {
		XMLName      string `xml:"tup:DeleteUplink"`
		DeleteUplink DeleteUplink
	}{
		DeleteUplink: args,
	}

	resp := DeleteUplinkResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *uplinkPort) OptGetServiceCapabilities(args GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error) {
	req := struct {
		XMLName                string `xml:"tup:GetServiceCapabilities"`
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

func (p *uplinkPort) OptGetUplinks(args GetUplinks) (*GetUplinksResponse, error) {
	req := struct {
		XMLName    string `xml:"tup:GetUplinks"`
		GetUplinks GetUplinks
	}{
		GetUplinks: args,
	}

	resp := GetUplinksResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *uplinkPort) OptSetUplink(args SetUplink) (*SetUplinkResponse, error) {
	req := struct {
		XMLName   string `xml:"tup:SetUplink"`
		SetUplink SetUplink
	}{
		SetUplink: args,
	}

	resp := SetUplinkResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
