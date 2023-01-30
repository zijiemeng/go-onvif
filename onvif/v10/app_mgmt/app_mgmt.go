package app_mgmt

import (
	"reflect"

	"code.byted.org/videoarch/go-onvif/onvif/common"
)

var Namespace = "http://www.onvif.org/ver10/appmgmt/wsdl"

// NewAppManagement creates an initializes a AppManagement.
func NewAppManagement(endpoint string, cli common.Client) AppManagement {
	return &appManagement{cli: cli, Endpoint: endpoint}
}

// AppManagement was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type AppManagement interface {
	OptActivate(Activate Activate) (*ActivateResponse, *common.Fault)

	OptDeactivate(Deactivate Deactivate) (*DeactivateResponse, *common.Fault)

	OptGetAppsInfo(GetAppsInfo GetAppsInfo) (*GetAppsInfoResponse, *common.Fault)

	OptGetDeviceId(GetDeviceId GetDeviceId) (*GetDeviceIdResponse, *common.Fault)

	OptGetInstalledApps(GetInstalledApps GetInstalledApps) (*GetInstalledAppsResponse, *common.Fault)

	OptGetServiceCapabilities(GetServiceCapabilities GetServiceCapabilities) (*GetServiceCapabilitiesResponse, *common.Fault)

	OptInstallLicense(InstallLicense InstallLicense) (*InstallLicenseResponse, *common.Fault)

	OptUninstall(Uninstall Uninstall) (*UninstallResponse, *common.Fault)
}
type DateTime string

type AppState string

func (v AppState) Validate() bool {
	for _, vv := range []string{
		"Active",
		"Inactive",
		"Installing",
		"Uninstalling",
		"Removed",
		"InstallationFailed",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type Activate struct {
	AppID string `xml:"AppID,omitempty" json:"AppID,omitempty" yaml:"AppID,omitempty"`
}

type ActivateResponse struct {
}

type AppInfo struct {
	AppID                string           `xml:"AppID,omitempty" json:"AppID,omitempty" yaml:"AppID,omitempty"`
	Name                 string           `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Version              string           `xml:"Version,omitempty" json:"Version,omitempty" yaml:"Version,omitempty"`
	Licenses             []LicenseInfo    `xml:"Licenses,omitempty" json:"Licenses,omitempty" yaml:"Licenses,omitempty"`
	Privileges           []string         `xml:"Privileges,omitempty" json:"Privileges,omitempty" yaml:"Privileges,omitempty"`
	InstallationDate     *common.DateTime `xml:"InstallationDate,omitempty" json:"InstallationDate,omitempty" yaml:"InstallationDate,omitempty"`
	LastUpdate           *common.DateTime `xml:"LastUpdate,omitempty" json:"LastUpdate,omitempty" yaml:"LastUpdate,omitempty"`
	State                AppState         `xml:"State,omitempty" json:"State,omitempty" yaml:"State,omitempty"`
	Status               string           `xml:"Status,omitempty" json:"Status,omitempty" yaml:"Status,omitempty"`
	Autostart            bool             `xml:"Autostart,omitempty" json:"Autostart,omitempty" yaml:"Autostart,omitempty"`
	Website              string           `xml:"Website,omitempty" json:"Website,omitempty" yaml:"Website,omitempty"`
	OpenSource           string           `xml:"OpenSource,omitempty" json:"OpenSource,omitempty" yaml:"OpenSource,omitempty"`
	Configuration        string           `xml:"Configuration,omitempty" json:"Configuration,omitempty" yaml:"Configuration,omitempty"`
	InterfaceDescription []string         `xml:"InterfaceDescription,omitempty" json:"InterfaceDescription,omitempty" yaml:"InterfaceDescription,omitempty"`
}

type Capabilities []interface{}

type Deactivate struct {
	AppID string `xml:"AppID,omitempty" json:"AppID,omitempty" yaml:"AppID,omitempty"`
}

type DeactivateResponse struct {
}

type GetAppsInfo struct {
	AppID string `xml:"AppID,omitempty" json:"AppID,omitempty" yaml:"AppID,omitempty"`
}

type GetAppsInfoResponse struct {
	Info []AppInfo `xml:"Info,omitempty" json:"Info,omitempty" yaml:"Info,omitempty"`
}

type GetDeviceId struct {
}

type GetDeviceIdResponse struct {
	DeviceId string `xml:"DeviceId,omitempty" json:"DeviceId,omitempty" yaml:"DeviceId,omitempty"`
}

type GetInstalledApps struct {
}

type GetInstalledAppsResponse struct {
	App []string `xml:"App>App,omitempty" json:"App>App,omitempty" yaml:"App>App,omitempty"`
}

type GetServiceCapabilities struct {
}

type GetServiceCapabilitiesResponse struct {
	Capabilities Capabilities `xml:"Capabilities,omitempty" json:"Capabilities,omitempty" yaml:"Capabilities,omitempty"`
}

type InstallLicense struct {
	AppID   string `xml:"AppID,omitempty" json:"AppID,omitempty" yaml:"AppID,omitempty"`
	License string `xml:"License,omitempty" json:"License,omitempty" yaml:"License,omitempty"`
}

type InstallLicenseResponse struct {
}

type LicenseInfo struct {
	Name       string           `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	ValidFrom  *common.DateTime `xml:"ValidFrom,omitempty" json:"ValidFrom,omitempty" yaml:"ValidFrom,omitempty"`
	ValidUntil *common.DateTime `xml:"ValidUntil,omitempty" json:"ValidUntil,omitempty" yaml:"ValidUntil,omitempty"`
}

type Uninstall struct {
	AppID string `xml:"AppID,omitempty" json:"AppID,omitempty" yaml:"AppID,omitempty"`
}

type UninstallResponse struct {
}

// appManagement implements the AppManagement interface.
type appManagement struct {
	cli      common.Client
	Endpoint string
}

func (p *appManagement) OptActivate(args Activate) (*ActivateResponse, *common.Fault) {
	req := struct {
		XMLName  string `xml:"ans:Activate"`
		Activate Activate
	}{
		Activate: args,
	}

	resp := ActivateResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *appManagement) OptDeactivate(args Deactivate) (*DeactivateResponse, *common.Fault) {
	req := struct {
		XMLName    string `xml:"ans:Deactivate"`
		Deactivate Deactivate
	}{
		Deactivate: args,
	}

	resp := DeactivateResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *appManagement) OptGetAppsInfo(args GetAppsInfo) (*GetAppsInfoResponse, *common.Fault) {
	req := struct {
		XMLName     string `xml:"ans:GetAppsInfo"`
		GetAppsInfo GetAppsInfo
	}{
		GetAppsInfo: args,
	}

	resp := GetAppsInfoResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *appManagement) OptGetDeviceId(args GetDeviceId) (*GetDeviceIdResponse, *common.Fault) {
	req := struct {
		XMLName     string `xml:"ans:GetDeviceId"`
		GetDeviceId GetDeviceId
	}{
		GetDeviceId: args,
	}

	resp := GetDeviceIdResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *appManagement) OptGetInstalledApps(args GetInstalledApps) (*GetInstalledAppsResponse, *common.Fault) {
	req := struct {
		XMLName          string `xml:"ans:GetInstalledApps"`
		GetInstalledApps GetInstalledApps
	}{
		GetInstalledApps: args,
	}

	resp := GetInstalledAppsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *appManagement) OptGetServiceCapabilities(args GetServiceCapabilities) (*GetServiceCapabilitiesResponse, *common.Fault) {
	req := struct {
		XMLName                string `xml:"ans:GetServiceCapabilities"`
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

func (p *appManagement) OptInstallLicense(args InstallLicense) (*InstallLicenseResponse, *common.Fault) {
	req := struct {
		XMLName        string `xml:"ans:InstallLicense"`
		InstallLicense InstallLicense
	}{
		InstallLicense: args,
	}

	resp := InstallLicenseResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *appManagement) OptUninstall(args Uninstall) (*UninstallResponse, *common.Fault) {
	req := struct {
		XMLName   string `xml:"ans:Uninstall"`
		Uninstall Uninstall
	}{
		Uninstall: args,
	}

	resp := UninstallResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}
