package device

import (
	"reflect"

	"github.com/zijiemeng/go-onvif/onvif/common"
)

var Namespace = "http://www.onvif.org/ver10/device/wsdl"

// NewDevice creates an initializes a Device.
func NewDevice(endpoint string, cli common.Client) Device {
	return &device{cli: cli, Endpoint: endpoint}
}

// Device was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type Device interface {
	OptAddIPAddressFilter(AddIPAddressFilter AddIPAddressFilter) (*AddIPAddressFilterResponse, *common.Fault)

	OptAddScopes(AddScopes AddScopes) (*AddScopesResponse, *common.Fault)

	OptCreateCertificate(CreateCertificate CreateCertificate) (*CreateCertificateResponse, *common.Fault)

	OptCreateDot1XConfiguration(CreateDot1XConfiguration CreateDot1XConfiguration) (*CreateDot1XConfigurationResponse, *common.Fault)

	OptCreateStorageConfiguration(CreateStorageConfiguration CreateStorageConfiguration) (*CreateStorageConfigurationResponse, *common.Fault)

	OptCreateUsers(CreateUsers CreateUsers) (*CreateUsersResponse, *common.Fault)

	OptDeleteCertificates(DeleteCertificates DeleteCertificates) (*DeleteCertificatesResponse, *common.Fault)

	OptDeleteDot1XConfiguration(DeleteDot1XConfiguration DeleteDot1XConfiguration) (*DeleteDot1XConfigurationResponse, *common.Fault)

	OptDeleteGeoLocation(DeleteGeoLocation DeleteGeoLocation) (*DeleteGeoLocationResponse, *common.Fault)

	OptDeleteStorageConfiguration(DeleteStorageConfiguration DeleteStorageConfiguration) (*DeleteStorageConfigurationResponse, *common.Fault)

	OptDeleteUsers(DeleteUsers DeleteUsers) (*DeleteUsersResponse, *common.Fault)

	OptGetAccessPolicy(GetAccessPolicy GetAccessPolicy) (*GetAccessPolicyResponse, *common.Fault)

	OptGetAuthFailureWarningConfiguration(GetAuthFailureWarningConfiguration GetAuthFailureWarningConfiguration) (*GetAuthFailureWarningConfigurationResponse, *common.Fault)

	OptGetAuthFailureWarningOptions(GetAuthFailureWarningOptions GetAuthFailureWarningOptions) (*GetAuthFailureWarningOptionsResponse, *common.Fault)

	OptGetCACertificates(GetCACertificates GetCACertificates) (*GetCACertificatesResponse, *common.Fault)

	OptGetCapabilities(GetCapabilities GetCapabilities) (*GetCapabilitiesResponse, *common.Fault)

	OptGetCertificateInformation(GetCertificateInformation GetCertificateInformation) (*GetCertificateInformationResponse, *common.Fault)

	OptGetCertificates(GetCertificates GetCertificates) (*GetCertificatesResponse, *common.Fault)

	OptGetCertificatesStatus(GetCertificatesStatus GetCertificatesStatus) (*GetCertificatesStatusResponse, *common.Fault)

	OptGetClientCertificateMode(GetClientCertificateMode GetClientCertificateMode) (*GetClientCertificateModeResponse, *common.Fault)

	OptGetDNS(GetDNS GetDNS) (*GetDNSResponse, *common.Fault)

	OptGetDPAddresses(GetDPAddresses GetDPAddresses) (*GetDPAddressesResponse, *common.Fault)

	OptGetDeviceInformation(GetDeviceInformation GetDeviceInformation) (*GetDeviceInformationResponse, *common.Fault)

	OptGetDiscoveryMode(GetDiscoveryMode GetDiscoveryMode) (*GetDiscoveryModeResponse, *common.Fault)

	OptGetDot11Capabilities(GetDot11Capabilities GetDot11Capabilities) (*GetDot11CapabilitiesResponse, *common.Fault)

	OptGetDot11Status(GetDot11Status GetDot11Status) (*GetDot11StatusResponse, *common.Fault)

	OptGetDot1XConfiguration(GetDot1XConfiguration GetDot1XConfiguration) (*GetDot1XConfigurationResponse, *common.Fault)

	OptGetDot1XConfigurations(GetDot1XConfigurations GetDot1XConfigurations) (*GetDot1XConfigurationsResponse, *common.Fault)

	OptGetDynamicDNS(GetDynamicDNS GetDynamicDNS) (*GetDynamicDNSResponse, *common.Fault)

	OptGetEndpointReference(GetEndpointReference GetEndpointReference) (*GetEndpointReferenceResponse, *common.Fault)

	OptGetGeoLocation(GetGeoLocation GetGeoLocation) (*GetGeoLocationResponse, *common.Fault)

	OptGetHostname(GetHostname GetHostname) (*GetHostnameResponse, *common.Fault)

	OptGetIPAddressFilter(GetIPAddressFilter GetIPAddressFilter) (*GetIPAddressFilterResponse, *common.Fault)

	OptGetNTP(GetNTP GetNTP) (*GetNTPResponse, *common.Fault)

	OptGetNetworkDefaultGateway(GetNetworkDefaultGateway GetNetworkDefaultGateway) (*GetNetworkDefaultGatewayResponse, *common.Fault)

	OptGetNetworkInterfaces(GetNetworkInterfaces GetNetworkInterfaces) (*GetNetworkInterfacesResponse, *common.Fault)

	OptGetNetworkProtocols(GetNetworkProtocols GetNetworkProtocols) (*GetNetworkProtocolsResponse, *common.Fault)

	OptGetPasswordComplexityConfiguration(GetPasswordComplexityConfiguration GetPasswordComplexityConfiguration) (*GetPasswordComplexityConfigurationResponse, *common.Fault)

	OptGetPasswordComplexityOptions(GetPasswordComplexityOptions GetPasswordComplexityOptions) (*GetPasswordComplexityOptionsResponse, *common.Fault)

	OptGetPasswordHistoryConfiguration(GetPasswordHistoryConfiguration GetPasswordHistoryConfiguration) (*GetPasswordHistoryConfigurationResponse, *common.Fault)

	OptGetPkcs10Request(GetPkcs10Request GetPkcs10Request) (*GetPkcs10RequestResponse, *common.Fault)

	OptGetRelayOutputs(GetRelayOutputs GetRelayOutputs) (*GetRelayOutputsResponse, *common.Fault)

	OptGetRemoteDiscoveryMode(GetRemoteDiscoveryMode GetRemoteDiscoveryMode) (*GetRemoteDiscoveryModeResponse, *common.Fault)

	OptGetRemoteUser(GetRemoteUser GetRemoteUser) (*GetRemoteUserResponse, *common.Fault)

	OptGetScopes(GetScopes GetScopes) (*GetScopesResponse, *common.Fault)

	OptGetServiceCapabilities(GetServiceCapabilities GetServiceCapabilities) (*GetServiceCapabilitiesResponse, *common.Fault)

	OptGetServices(GetServices GetServices) (*GetServicesResponse, *common.Fault)

	OptGetStorageConfiguration(GetStorageConfiguration GetStorageConfiguration) (*GetStorageConfigurationResponse, *common.Fault)

	OptGetStorageConfigurations(GetStorageConfigurations GetStorageConfigurations) (*GetStorageConfigurationsResponse, *common.Fault)

	OptGetSystemBackup(GetSystemBackup GetSystemBackup) (*GetSystemBackupResponse, *common.Fault)

	OptGetSystemDateAndTime(GetSystemDateAndTime GetSystemDateAndTime) (*GetSystemDateAndTimeResponse, *common.Fault)

	OptGetSystemLog(GetSystemLog GetSystemLog) (*GetSystemLogResponse, *common.Fault)

	OptGetSystemSupportInformation(GetSystemSupportInformation GetSystemSupportInformation) (*GetSystemSupportInformationResponse, *common.Fault)

	OptGetSystemUris(GetSystemUris GetSystemUris) (*GetSystemUrisResponse, *common.Fault)

	OptGetUsers(GetUsers GetUsers) (*GetUsersResponse, *common.Fault)

	OptGetWsdlUrl(GetWsdlUrl GetWsdlUrl) (*GetWsdlUrlResponse, *common.Fault)

	OptGetZeroConfiguration(GetZeroConfiguration GetZeroConfiguration) (*GetZeroConfigurationResponse, *common.Fault)

	OptLoadCACertificates(LoadCACertificates LoadCACertificates) (*LoadCACertificatesResponse, *common.Fault)

	OptLoadCertificateWithPrivateKey(LoadCertificateWithPrivateKey LoadCertificateWithPrivateKey) (*LoadCertificateWithPrivateKeyResponse, *common.Fault)

	OptLoadCertificates(LoadCertificates LoadCertificates) (*LoadCertificatesResponse, *common.Fault)

	OptRemoveIPAddressFilter(RemoveIPAddressFilter RemoveIPAddressFilter) (*RemoveIPAddressFilterResponse, *common.Fault)

	OptRemoveScopes(RemoveScopes RemoveScopes) (*RemoveScopesResponse, *common.Fault)

	OptRestoreSystem(RestoreSystem RestoreSystem) (*RestoreSystemResponse, *common.Fault)

	OptScanAvailableDot11Networks(ScanAvailableDot11Networks ScanAvailableDot11Networks) (*ScanAvailableDot11NetworksResponse, *common.Fault)

	OptSendAuxiliaryCommand(SendAuxiliaryCommand SendAuxiliaryCommand) (*SendAuxiliaryCommandResponse, *common.Fault)

	OptSetAccessPolicy(SetAccessPolicy SetAccessPolicy) (*SetAccessPolicyResponse, *common.Fault)

	OptSetAuthFailureWarningConfiguration(SetAuthFailureWarningConfiguration SetAuthFailureWarningConfiguration) (*SetAuthFailureWarningConfigurationResponse, *common.Fault)

	OptSetCertificatesStatus(SetCertificatesStatus SetCertificatesStatus) (*SetCertificatesStatusResponse, *common.Fault)

	OptSetClientCertificateMode(SetClientCertificateMode SetClientCertificateMode) (*SetClientCertificateModeResponse, *common.Fault)

	OptSetDNS(SetDNS SetDNS) (*SetDNSResponse, *common.Fault)

	OptSetDPAddresses(SetDPAddresses SetDPAddresses) (*SetDPAddressesResponse, *common.Fault)

	OptSetDiscoveryMode(SetDiscoveryMode SetDiscoveryMode) (*SetDiscoveryModeResponse, *common.Fault)

	OptSetDot1XConfiguration(SetDot1XConfiguration SetDot1XConfiguration) (*SetDot1XConfigurationResponse, *common.Fault)

	OptSetDynamicDNS(SetDynamicDNS SetDynamicDNS) (*SetDynamicDNSResponse, *common.Fault)

	OptSetGeoLocation(SetGeoLocation SetGeoLocation) (*SetGeoLocationResponse, *common.Fault)

	OptSetHashingAlgorithm(SetHashingAlgorithm SetHashingAlgorithm) (*SetHashingAlgorithmResponse, *common.Fault)

	OptSetHostname(SetHostname SetHostname) (*SetHostnameResponse, *common.Fault)

	OptSetHostnameFromDHCP(SetHostnameFromDHCP SetHostnameFromDHCP) (*SetHostnameFromDHCPResponse, *common.Fault)

	OptSetIPAddressFilter(SetIPAddressFilter SetIPAddressFilter) (*SetIPAddressFilterResponse, *common.Fault)

	OptSetNTP(SetNTP SetNTP) (*SetNTPResponse, *common.Fault)

	OptSetNetworkDefaultGateway(SetNetworkDefaultGateway SetNetworkDefaultGateway) (*SetNetworkDefaultGatewayResponse, *common.Fault)

	OptSetNetworkInterfaces(SetNetworkInterfaces SetNetworkInterfaces) (*SetNetworkInterfacesResponse, *common.Fault)

	OptSetNetworkProtocols(SetNetworkProtocols SetNetworkProtocols) (*SetNetworkProtocolsResponse, *common.Fault)

	OptSetPasswordComplexityConfiguration(SetPasswordComplexityConfiguration SetPasswordComplexityConfiguration) (*SetPasswordComplexityConfigurationResponse, *common.Fault)

	OptSetPasswordHistoryConfiguration(SetPasswordHistoryConfiguration SetPasswordHistoryConfiguration) (*SetPasswordHistoryConfigurationResponse, *common.Fault)

	OptSetRelayOutputSettings(SetRelayOutputSettings SetRelayOutputSettings) (*SetRelayOutputSettingsResponse, *common.Fault)

	OptSetRelayOutputState(SetRelayOutputState SetRelayOutputState) (*SetRelayOutputStateResponse, *common.Fault)

	OptSetRemoteDiscoveryMode(SetRemoteDiscoveryMode SetRemoteDiscoveryMode) (*SetRemoteDiscoveryModeResponse, *common.Fault)

	OptSetRemoteUser(SetRemoteUser SetRemoteUser) (*SetRemoteUserResponse, *common.Fault)

	OptSetScopes(SetScopes SetScopes) (*SetScopesResponse, *common.Fault)

	OptSetStorageConfiguration(SetStorageConfiguration SetStorageConfiguration) (*SetStorageConfigurationResponse, *common.Fault)

	OptSetSystemDateAndTime(SetSystemDateAndTime SetSystemDateAndTime) (*SetSystemDateAndTimeResponse, *common.Fault)

	OptSetSystemFactoryDefault(SetSystemFactoryDefault SetSystemFactoryDefault) (*SetSystemFactoryDefaultResponse, *common.Fault)

	OptSetUser(SetUser SetUser) (*SetUserResponse, *common.Fault)

	OptSetZeroConfiguration(SetZeroConfiguration SetZeroConfiguration) (*SetZeroConfigurationResponse, *common.Fault)

	OptStartFirmwareUpgrade(StartFirmwareUpgrade StartFirmwareUpgrade) (*StartFirmwareUpgradeResponse, *common.Fault)

	OptStartSystemRestore(StartSystemRestore StartSystemRestore) (*StartSystemRestoreResponse, *common.Fault)

	OptSystemReboot(SystemReboot SystemReboot) (*SystemRebootResponse, *common.Fault)

	OptUpgradeSystemFirmware(UpgradeSystemFirmware UpgradeSystemFirmware) (*UpgradeSystemFirmwareResponse, *common.Fault)
}
type DateTime string

type Duration string

type AutoGeoModes string

func (v AutoGeoModes) Validate() bool {
	for _, vv := range []string{
		"Location",
		"Heading",
		"Leveling",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type StorageType string

func (v StorageType) Validate() bool {
	for _, vv := range []string{
		"NFS",
		"CIFS",
		"CDMI",
		"FTP",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type AddIPAddressFilter struct {
	IPAddressFilter *common.IPAddressFilter `xml:"IPAddressFilter,omitempty" json:"IPAddressFilter,omitempty" yaml:"IPAddressFilter,omitempty"`
}

type AddIPAddressFilterResponse struct {
}

type AddScopes struct {
	ScopeItem []string `xml:"ScopeItem" json:"ScopeItem" yaml:"ScopeItem"`
}

type AddScopesResponse struct {
}

type CreateCertificate struct {
	CertificateID  string           `xml:"CertificateID,omitempty" json:"CertificateID,omitempty" yaml:"CertificateID,omitempty"`
	Subject        string           `xml:"Subject,omitempty" json:"Subject,omitempty" yaml:"Subject,omitempty"`
	ValidNotBefore *common.DateTime `xml:"ValidNotBefore,omitempty" json:"ValidNotBefore,omitempty" yaml:"ValidNotBefore,omitempty"`
	ValidNotAfter  *common.DateTime `xml:"ValidNotAfter,omitempty" json:"ValidNotAfter,omitempty" yaml:"ValidNotAfter,omitempty"`
}

type CreateCertificateResponse struct {
	NvtCertificate *common.Certificate `xml:"NvtCertificate,omitempty" json:"NvtCertificate,omitempty" yaml:"NvtCertificate,omitempty"`
}

type CreateDot1XConfiguration struct {
	Dot1XConfiguration common.Dot1XConfiguration `xml:"Dot1XConfiguration" json:"Dot1XConfiguration" yaml:"Dot1XConfiguration"`
}

type CreateDot1XConfigurationResponse struct {
}

type CreateStorageConfiguration struct {
	StorageConfiguration StorageConfigurationData `xml:"StorageConfiguration,omitempty" json:"StorageConfiguration,omitempty" yaml:"StorageConfiguration,omitempty"`
}

type CreateStorageConfigurationResponse struct {
	Token *common.ReferenceToken `xml:"Token,omitempty" json:"Token,omitempty" yaml:"Token,omitempty"`
}

type CreateUsers struct {
	User []common.User `xml:"User" json:"User" yaml:"User"`
}

type CreateUsersResponse struct {
}

type DeleteCertificates struct {
	CertificateID []string `xml:"CertificateID" json:"CertificateID" yaml:"CertificateID"`
}

type DeleteCertificatesResponse struct {
}

type DeleteDot1XConfiguration struct {
	Dot1XConfigurationToken []*common.ReferenceToken `xml:"Dot1XConfigurationToken,omitempty" json:"Dot1XConfigurationToken,omitempty" yaml:"Dot1XConfigurationToken,omitempty"`
}

type DeleteDot1XConfigurationResponse struct {
}

type DeleteGeoLocation struct {
	Location []*common.LocationEntity `xml:"Location,omitempty" json:"Location,omitempty" yaml:"Location,omitempty"`
}

type DeleteGeoLocationResponse struct {
}

type DeleteStorageConfiguration struct {
	Token *common.ReferenceToken `xml:"Token,omitempty" json:"Token,omitempty" yaml:"Token,omitempty"`
}

type DeleteStorageConfigurationResponse struct {
}

type DeleteUsers struct {
	Username []string `xml:"Username" json:"Username" yaml:"Username"`
}

type DeleteUsersResponse struct {
}

type DeviceServiceCapabilities struct {
	Network  NetworkCapabilities  `xml:"Network,omitempty" json:"Network,omitempty" yaml:"Network,omitempty"`
	Security SecurityCapabilities `xml:"Security,omitempty" json:"Security,omitempty" yaml:"Security,omitempty"`
	System   SystemCapabilities   `xml:"System,omitempty" json:"System,omitempty" yaml:"System,omitempty"`
	Misc     MiscCapabilities     `xml:"Misc,omitempty" json:"Misc,omitempty" yaml:"Misc,omitempty"`
}

type GetAccessPolicy struct {
}

type GetAccessPolicyResponse struct {
	PolicyFile *common.BinaryData `xml:"PolicyFile,omitempty" json:"PolicyFile,omitempty" yaml:"PolicyFile,omitempty"`
}

type GetAuthFailureWarningConfiguration struct {
}

type GetAuthFailureWarningConfigurationResponse struct {
	Enabled         bool `xml:"Enabled,omitempty" json:"Enabled,omitempty" yaml:"Enabled,omitempty"`
	MonitorPeriod   int  `xml:"MonitorPeriod,omitempty" json:"MonitorPeriod,omitempty" yaml:"MonitorPeriod,omitempty"`
	MaxAuthFailures int  `xml:"MaxAuthFailures,omitempty" json:"MaxAuthFailures,omitempty" yaml:"MaxAuthFailures,omitempty"`
}

type GetAuthFailureWarningOptions struct {
}

type GetAuthFailureWarningOptionsResponse struct {
	MonitorPeriodRange *common.IntRange `xml:"MonitorPeriodRange,omitempty" json:"MonitorPeriodRange,omitempty" yaml:"MonitorPeriodRange,omitempty"`
	AuthFailureRange   *common.IntRange `xml:"AuthFailureRange,omitempty" json:"AuthFailureRange,omitempty" yaml:"AuthFailureRange,omitempty"`
}

type GetCACertificates struct {
}

type GetCACertificatesResponse struct {
	CACertificate []*common.Certificate `xml:"CACertificate,omitempty" json:"CACertificate,omitempty" yaml:"CACertificate,omitempty"`
}

type GetCapabilities struct {
	Category []common.CapabilityCategory `xml:"Category,omitempty" json:"Category,omitempty" yaml:"Category,omitempty"`
}

type GetCapabilitiesResponse struct {
	Capabilities *common.Capabilities `xml:"Capabilities,omitempty" json:"Capabilities,omitempty" yaml:"Capabilities,omitempty"`
}

type GetCertificateInformation struct {
	CertificateID string `xml:"CertificateID" json:"CertificateID" yaml:"CertificateID"`
}

type GetCertificateInformationResponse struct {
	CertificateInformation common.CertificateInformation `xml:"CertificateInformation" json:"CertificateInformation" yaml:"CertificateInformation"`
}

type GetCertificates struct {
}

type GetCertificatesResponse struct {
	NvtCertificate []*common.Certificate `xml:"NvtCertificate,omitempty" json:"NvtCertificate,omitempty" yaml:"NvtCertificate,omitempty"`
}

type GetCertificatesStatus struct {
}

type GetCertificatesStatusResponse struct {
	CertificateStatus []*common.CertificateStatus `xml:"CertificateStatus,omitempty" json:"CertificateStatus,omitempty" yaml:"CertificateStatus,omitempty"`
}

type GetClientCertificateMode struct {
}

type GetClientCertificateModeResponse struct {
	Enabled bool `xml:"Enabled,omitempty" json:"Enabled,omitempty" yaml:"Enabled,omitempty"`
}

type GetDNS struct {
}

type GetDNSResponse struct {
	DNSInformation *common.DNSInformation `xml:"DNSInformation,omitempty" json:"DNSInformation,omitempty" yaml:"DNSInformation,omitempty"`
}

type GetDPAddresses struct {
}

type GetDPAddressesResponse struct {
	DPAddress []*common.NetworkHost `xml:"DPAddress,omitempty" json:"DPAddress,omitempty" yaml:"DPAddress,omitempty"`
}

type GetDeviceInformation struct {
}

type GetDeviceInformationResponse struct {
	Manufacturer    string `xml:"Manufacturer,omitempty" json:"Manufacturer,omitempty" yaml:"Manufacturer,omitempty"`
	Model           string `xml:"Model,omitempty" json:"Model,omitempty" yaml:"Model,omitempty"`
	FirmwareVersion string `xml:"FirmwareVersion,omitempty" json:"FirmwareVersion,omitempty" yaml:"FirmwareVersion,omitempty"`
	SerialNumber    string `xml:"SerialNumber,omitempty" json:"SerialNumber,omitempty" yaml:"SerialNumber,omitempty"`
	HardwareId      string `xml:"HardwareId,omitempty" json:"HardwareId,omitempty" yaml:"HardwareId,omitempty"`
}

type GetDiscoveryMode struct {
}

type GetDiscoveryModeResponse struct {
	DiscoveryMode *common.DiscoveryMode `xml:"DiscoveryMode,omitempty" json:"DiscoveryMode,omitempty" yaml:"DiscoveryMode,omitempty"`
}

type GetDot11Capabilities []interface{}

type GetDot11CapabilitiesResponse struct {
	Capabilities *common.Dot11Capabilities `xml:"Capabilities,omitempty" json:"Capabilities,omitempty" yaml:"Capabilities,omitempty"`
}

type GetDot11Status struct {
	InterfaceToken *common.ReferenceToken `xml:"InterfaceToken,omitempty" json:"InterfaceToken,omitempty" yaml:"InterfaceToken,omitempty"`
}

type GetDot11StatusResponse struct {
	Status *common.Dot11Status `xml:"Status,omitempty" json:"Status,omitempty" yaml:"Status,omitempty"`
}

type GetDot1XConfiguration struct {
	Dot1XConfigurationToken common.ReferenceToken `xml:"Dot1XConfigurationToken" json:"Dot1XConfigurationToken" yaml:"Dot1XConfigurationToken"`
}

type GetDot1XConfigurationResponse struct {
	Dot1XConfiguration common.Dot1XConfiguration `xml:"Dot1XConfiguration" json:"Dot1XConfiguration" yaml:"Dot1XConfiguration"`
}

type GetDot1XConfigurations struct {
}

type GetDot1XConfigurationsResponse struct {
	Dot1XConfiguration []*common.Dot1XConfiguration `xml:"Dot1XConfiguration,omitempty" json:"Dot1XConfiguration,omitempty" yaml:"Dot1XConfiguration,omitempty"`
}

type GetDynamicDNS struct {
}

type GetDynamicDNSResponse struct {
	DynamicDNSInformation *common.DynamicDNSInformation `xml:"DynamicDNSInformation,omitempty" json:"DynamicDNSInformation,omitempty" yaml:"DynamicDNSInformation,omitempty"`
}

type GetEndpointReference struct {
}

type GetEndpointReferenceResponse struct {
	GUID string `xml:"GUID,omitempty" json:"GUID,omitempty" yaml:"GUID,omitempty"`
}

type GetGeoLocation struct {
}

type GetGeoLocationResponse struct {
	Location []*common.LocationEntity `xml:"Location,omitempty" json:"Location,omitempty" yaml:"Location,omitempty"`
}

type GetHostname struct {
}

type GetHostnameResponse struct {
	HostnameInformation *common.HostnameInformation `xml:"HostnameInformation,omitempty" json:"HostnameInformation,omitempty" yaml:"HostnameInformation,omitempty"`
}

type GetIPAddressFilter struct {
}

type GetIPAddressFilterResponse struct {
	IPAddressFilter *common.IPAddressFilter `xml:"IPAddressFilter,omitempty" json:"IPAddressFilter,omitempty" yaml:"IPAddressFilter,omitempty"`
}

type GetNTP struct {
}

type GetNTPResponse struct {
	NTPInformation *common.NTPInformation `xml:"NTPInformation,omitempty" json:"NTPInformation,omitempty" yaml:"NTPInformation,omitempty"`
}

type GetNetworkDefaultGateway struct {
}

type GetNetworkDefaultGatewayResponse struct {
	NetworkGateway *common.NetworkGateway `xml:"NetworkGateway,omitempty" json:"NetworkGateway,omitempty" yaml:"NetworkGateway,omitempty"`
}

type GetNetworkInterfaces struct {
}

type GetNetworkInterfacesResponse struct {
	NetworkInterfaces []common.NetworkInterface `xml:"NetworkInterfaces" json:"NetworkInterfaces" yaml:"NetworkInterfaces"`
}

type GetNetworkProtocols struct {
}

type GetNetworkProtocolsResponse struct {
	NetworkProtocols []*common.NetworkProtocol `xml:"NetworkProtocols,omitempty" json:"NetworkProtocols,omitempty" yaml:"NetworkProtocols,omitempty"`
}

type GetPasswordComplexityConfiguration struct {
}

type GetPasswordComplexityConfigurationResponse struct {
	MinLen                    int  `xml:"MinLen,omitempty" json:"MinLen,omitempty" yaml:"MinLen,omitempty"`
	Uppercase                 int  `xml:"Uppercase,omitempty" json:"Uppercase,omitempty" yaml:"Uppercase,omitempty"`
	Number                    int  `xml:"Number,omitempty" json:"Number,omitempty" yaml:"Number,omitempty"`
	SpecialChars              int  `xml:"SpecialChars,omitempty" json:"SpecialChars,omitempty" yaml:"SpecialChars,omitempty"`
	BlockUsernameOccurrence   bool `xml:"BlockUsernameOccurrence,omitempty" json:"BlockUsernameOccurrence,omitempty" yaml:"BlockUsernameOccurrence,omitempty"`
	PolicyConfigurationLocked bool `xml:"PolicyConfigurationLocked,omitempty" json:"PolicyConfigurationLocked,omitempty" yaml:"PolicyConfigurationLocked,omitempty"`
}

type GetPasswordComplexityOptions struct {
}

type GetPasswordComplexityOptionsResponse struct {
	MinLenRange                      *common.IntRange `xml:"MinLenRange,omitempty" json:"MinLenRange,omitempty" yaml:"MinLenRange,omitempty"`
	UppercaseRange                   *common.IntRange `xml:"UppercaseRange,omitempty" json:"UppercaseRange,omitempty" yaml:"UppercaseRange,omitempty"`
	NumberRange                      *common.IntRange `xml:"NumberRange,omitempty" json:"NumberRange,omitempty" yaml:"NumberRange,omitempty"`
	SpecialCharsRange                *common.IntRange `xml:"SpecialCharsRange,omitempty" json:"SpecialCharsRange,omitempty" yaml:"SpecialCharsRange,omitempty"`
	BlockUsernameOccurrenceSupported bool             `xml:"BlockUsernameOccurrenceSupported,omitempty" json:"BlockUsernameOccurrenceSupported,omitempty" yaml:"BlockUsernameOccurrenceSupported,omitempty"`
	PolicyConfigurationLockSupported bool             `xml:"PolicyConfigurationLockSupported,omitempty" json:"PolicyConfigurationLockSupported,omitempty" yaml:"PolicyConfigurationLockSupported,omitempty"`
}

type GetPasswordHistoryConfiguration struct {
}

type GetPasswordHistoryConfigurationResponse struct {
	Enabled bool `xml:"Enabled,omitempty" json:"Enabled,omitempty" yaml:"Enabled,omitempty"`
	Length  int  `xml:"Length,omitempty" json:"Length,omitempty" yaml:"Length,omitempty"`
}

type GetPkcs10Request struct {
	CertificateID string             `xml:"CertificateID,omitempty" json:"CertificateID,omitempty" yaml:"CertificateID,omitempty"`
	Subject       string             `xml:"Subject,omitempty" json:"Subject,omitempty" yaml:"Subject,omitempty"`
	Attributes    *common.BinaryData `xml:"Attributes,omitempty" json:"Attributes,omitempty" yaml:"Attributes,omitempty"`
}

type GetPkcs10RequestResponse struct {
	Pkcs10Request *common.BinaryData `xml:"Pkcs10Request,omitempty" json:"Pkcs10Request,omitempty" yaml:"Pkcs10Request,omitempty"`
}

type GetRelayOutputs struct {
}

type GetRelayOutputsResponse struct {
	RelayOutputs []*common.RelayOutput `xml:"RelayOutputs,omitempty" json:"RelayOutputs,omitempty" yaml:"RelayOutputs,omitempty"`
}

type GetRemoteDiscoveryMode struct {
}

type GetRemoteDiscoveryModeResponse struct {
	RemoteDiscoveryMode *common.DiscoveryMode `xml:"RemoteDiscoveryMode,omitempty" json:"RemoteDiscoveryMode,omitempty" yaml:"RemoteDiscoveryMode,omitempty"`
}

type GetRemoteUser struct {
}

type GetRemoteUserResponse struct {
	RemoteUser *common.RemoteUser `xml:"RemoteUser,omitempty" json:"RemoteUser,omitempty" yaml:"RemoteUser,omitempty"`
}

type GetScopes struct {
}

type GetScopesResponse struct {
	Scopes []common.Scope `xml:"Scopes" json:"Scopes" yaml:"Scopes"`
}

type GetServiceCapabilities struct {
}

type GetServiceCapabilitiesResponse struct {
	Capabilities DeviceServiceCapabilities `xml:"Capabilities,omitempty" json:"Capabilities,omitempty" yaml:"Capabilities,omitempty"`
}

type GetServices struct {
	IncludeCapability bool `xml:"IncludeCapability,omitempty" json:"IncludeCapability,omitempty" yaml:"IncludeCapability,omitempty"`
}

type GetServicesResponse struct {
	Service []Service `xml:"Service,omitempty" json:"Service,omitempty" yaml:"Service,omitempty"`
}

type GetStorageConfiguration struct {
	Token *common.ReferenceToken `xml:"Token,omitempty" json:"Token,omitempty" yaml:"Token,omitempty"`
}

type GetStorageConfigurationResponse struct {
	StorageConfiguration StorageConfiguration `xml:"StorageConfiguration,omitempty" json:"StorageConfiguration,omitempty" yaml:"StorageConfiguration,omitempty"`
}

type GetStorageConfigurations struct {
}

type GetStorageConfigurationsResponse struct {
	StorageConfigurations []StorageConfiguration `xml:"StorageConfigurations,omitempty" json:"StorageConfigurations,omitempty" yaml:"StorageConfigurations,omitempty"`
}

type GetSystemBackup struct {
}

type GetSystemBackupResponse struct {
	BackupFiles []common.BackupFile `xml:"BackupFiles" json:"BackupFiles" yaml:"BackupFiles"`
}

type GetSystemDateAndTime struct {
}

type GetSystemDateAndTimeResponse struct {
	SystemDateAndTime *common.SystemDateTime `xml:"SystemDateAndTime,omitempty" json:"SystemDateAndTime,omitempty" yaml:"SystemDateAndTime,omitempty"`
}

type GetSystemLog struct {
	LogType *common.SystemLogType `xml:"LogType,omitempty" json:"LogType,omitempty" yaml:"LogType,omitempty"`
}

type GetSystemLogResponse struct {
	SystemLog *common.SystemLog `xml:"SystemLog,omitempty" json:"SystemLog,omitempty" yaml:"SystemLog,omitempty"`
}

type GetSystemSupportInformation struct {
}

type GetSystemSupportInformationResponse struct {
	SupportInformation *common.SupportInformation `xml:"SupportInformation,omitempty" json:"SupportInformation,omitempty" yaml:"SupportInformation,omitempty"`
}

type GetSystemUris struct {
}

type GetSystemUrisResponse struct {
	SystemLogUris   *common.SystemLogUriList `xml:"SystemLogUris,omitempty" json:"SystemLogUris,omitempty" yaml:"SystemLogUris,omitempty"`
	SupportInfoUri  string                   `xml:"SupportInfoUri,omitempty" json:"SupportInfoUri,omitempty" yaml:"SupportInfoUri,omitempty"`
	SystemBackupUri string                   `xml:"SystemBackupUri,omitempty" json:"SystemBackupUri,omitempty" yaml:"SystemBackupUri,omitempty"`
	Extension       []string                 `xml:"Extension>Extension,omitempty" json:"Extension>Extension,omitempty" yaml:"Extension>Extension,omitempty"`
}

type GetUsers struct {
}

type GetUsersResponse struct {
	User []*common.User `xml:"User,omitempty" json:"User,omitempty" yaml:"User,omitempty"`
}

type GetWsdlUrl struct {
}

type GetWsdlUrlResponse struct {
	WsdlUrl string `xml:"WsdlUrl,omitempty" json:"WsdlUrl,omitempty" yaml:"WsdlUrl,omitempty"`
}

type GetZeroConfiguration struct {
}

type GetZeroConfigurationResponse struct {
	ZeroConfiguration *common.NetworkZeroConfiguration `xml:"ZeroConfiguration,omitempty" json:"ZeroConfiguration,omitempty" yaml:"ZeroConfiguration,omitempty"`
}

type LoadCACertificates struct {
	CACertificate []common.Certificate `xml:"CACertificate" json:"CACertificate" yaml:"CACertificate"`
}

type LoadCACertificatesResponse struct {
}

type LoadCertificateWithPrivateKey struct {
	CertificateWithPrivateKey []common.CertificateWithPrivateKey `xml:"CertificateWithPrivateKey" json:"CertificateWithPrivateKey" yaml:"CertificateWithPrivateKey"`
}

type LoadCertificateWithPrivateKeyResponse struct {
}

type LoadCertificates struct {
	NVTCertificate []common.Certificate `xml:"NVTCertificate" json:"NVTCertificate" yaml:"NVTCertificate"`
}

type LoadCertificatesResponse struct {
}

type MiscCapabilities struct {
	AuxiliaryCommands common.StringAttrList `xml:"AuxiliaryCommands,attr,omitempty" json:"AuxiliaryCommands,attr,omitempty" yaml:"AuxiliaryCommands,attr,omitempty"`
}

type NetworkCapabilities struct {
	IPFilter            bool `xml:"IPFilter,attr,omitempty" json:"IPFilter,attr,omitempty" yaml:"IPFilter,attr,omitempty"`
	ZeroConfiguration   bool `xml:"ZeroConfiguration,attr,omitempty" json:"ZeroConfiguration,attr,omitempty" yaml:"ZeroConfiguration,attr,omitempty"`
	IPVersion6          bool `xml:"IPVersion6,attr,omitempty" json:"IPVersion6,attr,omitempty" yaml:"IPVersion6,attr,omitempty"`
	DynDNS              bool `xml:"DynDNS,attr,omitempty" json:"DynDNS,attr,omitempty" yaml:"DynDNS,attr,omitempty"`
	Dot11Configuration  bool `xml:"Dot11Configuration,attr,omitempty" json:"Dot11Configuration,attr,omitempty" yaml:"Dot11Configuration,attr,omitempty"`
	Dot1XConfigurations int  `xml:"Dot1XConfigurations,attr,omitempty" json:"Dot1XConfigurations,attr,omitempty" yaml:"Dot1XConfigurations,attr,omitempty"`
	HostnameFromDHCP    bool `xml:"HostnameFromDHCP,attr,omitempty" json:"HostnameFromDHCP,attr,omitempty" yaml:"HostnameFromDHCP,attr,omitempty"`
	NTP                 int  `xml:"NTP,attr,omitempty" json:"NTP,attr,omitempty" yaml:"NTP,attr,omitempty"`
	DHCPv6              bool `xml:"DHCPv6,attr,omitempty" json:"DHCPv6,attr,omitempty" yaml:"DHCPv6,attr,omitempty"`
}

type RemoveIPAddressFilter struct {
	IPAddressFilter *common.IPAddressFilter `xml:"IPAddressFilter,omitempty" json:"IPAddressFilter,omitempty" yaml:"IPAddressFilter,omitempty"`
}

type RemoveIPAddressFilterResponse struct {
}

type RemoveScopes struct {
	ScopeItem []string `xml:"ScopeItem" json:"ScopeItem" yaml:"ScopeItem"`
}

type RemoveScopesResponse struct {
	ScopeItem []string `xml:"ScopeItem,omitempty" json:"ScopeItem,omitempty" yaml:"ScopeItem,omitempty"`
}

type RestoreSystem struct {
	BackupFiles []common.BackupFile `xml:"BackupFiles" json:"BackupFiles" yaml:"BackupFiles"`
}

type RestoreSystemResponse struct {
}

type ScanAvailableDot11Networks struct {
	InterfaceToken *common.ReferenceToken `xml:"InterfaceToken,omitempty" json:"InterfaceToken,omitempty" yaml:"InterfaceToken,omitempty"`
}

type ScanAvailableDot11NetworksResponse struct {
	Networks []*common.Dot11AvailableNetworks `xml:"Networks,omitempty" json:"Networks,omitempty" yaml:"Networks,omitempty"`
}

type SecurityCapabilities struct {
	TLS0                 bool              `xml:"TLS1.0,attr,omitempty" json:"TLS1.0,attr,omitempty" yaml:"TLS1.0,attr,omitempty"`
	TLS1                 bool              `xml:"TLS1.1,attr,omitempty" json:"TLS1.1,attr,omitempty" yaml:"TLS1.1,attr,omitempty"`
	TLS2                 bool              `xml:"TLS1.2,attr,omitempty" json:"TLS1.2,attr,omitempty" yaml:"TLS1.2,attr,omitempty"`
	OnboardKeyGeneration bool              `xml:"OnboardKeyGeneration,attr,omitempty" json:"OnboardKeyGeneration,attr,omitempty" yaml:"OnboardKeyGeneration,attr,omitempty"`
	AccessPolicyConfig   bool              `xml:"AccessPolicyConfig,attr,omitempty" json:"AccessPolicyConfig,attr,omitempty" yaml:"AccessPolicyConfig,attr,omitempty"`
	DefaultAccessPolicy  bool              `xml:"DefaultAccessPolicy,attr,omitempty" json:"DefaultAccessPolicy,attr,omitempty" yaml:"DefaultAccessPolicy,attr,omitempty"`
	Dot1X                bool              `xml:"Dot1X,attr,omitempty" json:"Dot1X,attr,omitempty" yaml:"Dot1X,attr,omitempty"`
	RemoteUserHandling   bool              `xml:"RemoteUserHandling,attr,omitempty" json:"RemoteUserHandling,attr,omitempty" yaml:"RemoteUserHandling,attr,omitempty"`
	X509Token            bool              `xml:"X.509Token,attr,omitempty" json:"X.509Token,attr,omitempty" yaml:"X.509Token,attr,omitempty"`
	SAMLToken            bool              `xml:"SAMLToken,attr,omitempty" json:"SAMLToken,attr,omitempty" yaml:"SAMLToken,attr,omitempty"`
	KerberosToken        bool              `xml:"KerberosToken,attr,omitempty" json:"KerberosToken,attr,omitempty" yaml:"KerberosToken,attr,omitempty"`
	UsernameToken        bool              `xml:"UsernameToken,attr,omitempty" json:"UsernameToken,attr,omitempty" yaml:"UsernameToken,attr,omitempty"`
	HttpDigest           bool              `xml:"HttpDigest,attr,omitempty" json:"HttpDigest,attr,omitempty" yaml:"HttpDigest,attr,omitempty"`
	RELToken             bool              `xml:"RELToken,attr,omitempty" json:"RELToken,attr,omitempty" yaml:"RELToken,attr,omitempty"`
	SupportedEAPMethods  common.IntList    `xml:"SupportedEAPMethods,attr,omitempty" json:"SupportedEAPMethods,attr,omitempty" yaml:"SupportedEAPMethods,attr,omitempty"`
	MaxUsers             int               `xml:"MaxUsers,attr,omitempty" json:"MaxUsers,attr,omitempty" yaml:"MaxUsers,attr,omitempty"`
	MaxUserNameLength    int               `xml:"MaxUserNameLength,attr,omitempty" json:"MaxUserNameLength,attr,omitempty" yaml:"MaxUserNameLength,attr,omitempty"`
	MaxPasswordLength    int               `xml:"MaxPasswordLength,attr,omitempty" json:"MaxPasswordLength,attr,omitempty" yaml:"MaxPasswordLength,attr,omitempty"`
	SecurityPolicies     common.StringList `xml:"SecurityPolicies,attr,omitempty" json:"SecurityPolicies,attr,omitempty" yaml:"SecurityPolicies,attr,omitempty"`
	MaxPasswordHistory   int               `xml:"MaxPasswordHistory,attr,omitempty" json:"MaxPasswordHistory,attr,omitempty" yaml:"MaxPasswordHistory,attr,omitempty"`
	HashingAlgorithms    common.StringList `xml:"HashingAlgorithms,attr,omitempty" json:"HashingAlgorithms,attr,omitempty" yaml:"HashingAlgorithms,attr,omitempty"`
}

type SendAuxiliaryCommand struct {
	AuxiliaryCommand common.AuxiliaryData `xml:"AuxiliaryCommand" json:"AuxiliaryCommand" yaml:"AuxiliaryCommand"`
}

type SendAuxiliaryCommandResponse struct {
	AuxiliaryCommandResponse *common.AuxiliaryData `xml:"AuxiliaryCommandResponse,omitempty" json:"AuxiliaryCommandResponse,omitempty" yaml:"AuxiliaryCommandResponse,omitempty"`
}

type Service struct {
	Namespace    string               `xml:"Namespace,omitempty" json:"Namespace,omitempty" yaml:"Namespace,omitempty"`
	XAddr        string               `xml:"XAddr,omitempty" json:"XAddr,omitempty" yaml:"XAddr,omitempty"`
	Capabilities string               `xml:"Capabilities,omitempty" json:"Capabilities,omitempty" yaml:"Capabilities,omitempty"`
	Version      *common.OnvifVersion `xml:"Version,omitempty" json:"Version,omitempty" yaml:"Version,omitempty"`
}

type SetAccessPolicy struct {
	PolicyFile *common.BinaryData `xml:"PolicyFile,omitempty" json:"PolicyFile,omitempty" yaml:"PolicyFile,omitempty"`
}

type SetAccessPolicyResponse struct {
}

type SetAuthFailureWarningConfiguration struct {
	Enabled         bool `xml:"Enabled,omitempty" json:"Enabled,omitempty" yaml:"Enabled,omitempty"`
	MonitorPeriod   int  `xml:"MonitorPeriod,omitempty" json:"MonitorPeriod,omitempty" yaml:"MonitorPeriod,omitempty"`
	MaxAuthFailures int  `xml:"MaxAuthFailures,omitempty" json:"MaxAuthFailures,omitempty" yaml:"MaxAuthFailures,omitempty"`
}

type SetAuthFailureWarningConfigurationResponse struct {
}

type SetCertificatesStatus struct {
	CertificateStatus []*common.CertificateStatus `xml:"CertificateStatus,omitempty" json:"CertificateStatus,omitempty" yaml:"CertificateStatus,omitempty"`
}

type SetCertificatesStatusResponse struct {
}

type SetClientCertificateMode struct {
	Enabled bool `xml:"Enabled,omitempty" json:"Enabled,omitempty" yaml:"Enabled,omitempty"`
}

type SetClientCertificateModeResponse struct {
}

type SetDNS struct {
	FromDHCP     bool                `xml:"FromDHCP,omitempty" json:"FromDHCP,omitempty" yaml:"FromDHCP,omitempty"`
	SearchDomain []string            `xml:"SearchDomain,omitempty" json:"SearchDomain,omitempty" yaml:"SearchDomain,omitempty"`
	DNSManual    []*common.IPAddress `xml:"DNSManual,omitempty" json:"DNSManual,omitempty" yaml:"DNSManual,omitempty"`
}

type SetDNSResponse struct {
}

type SetDPAddresses struct {
	DPAddress []*common.NetworkHost `xml:"DPAddress,omitempty" json:"DPAddress,omitempty" yaml:"DPAddress,omitempty"`
}

type SetDPAddressesResponse struct {
}

type SetDiscoveryMode struct {
	DiscoveryMode *common.DiscoveryMode `xml:"DiscoveryMode,omitempty" json:"DiscoveryMode,omitempty" yaml:"DiscoveryMode,omitempty"`
}

type SetDiscoveryModeResponse struct {
}

type SetDot1XConfiguration struct {
	Dot1XConfiguration common.Dot1XConfiguration `xml:"Dot1XConfiguration" json:"Dot1XConfiguration" yaml:"Dot1XConfiguration"`
}

type SetDot1XConfigurationResponse struct {
}

type SetDynamicDNS struct {
	Type *common.DynamicDNSType `xml:"Type,omitempty" json:"Type,omitempty" yaml:"Type,omitempty"`
	Name *common.DNSName        `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	TTL  *common.Duration       `xml:"TTL,omitempty" json:"TTL,omitempty" yaml:"TTL,omitempty"`
}

type SetDynamicDNSResponse struct {
}

type SetGeoLocation struct {
	Location []*common.LocationEntity `xml:"Location,omitempty" json:"Location,omitempty" yaml:"Location,omitempty"`
}

type SetGeoLocationResponse struct {
}

type SetHashingAlgorithm struct {
	Algorithm *common.StringList `xml:"Algorithm,omitempty" json:"Algorithm,omitempty" yaml:"Algorithm,omitempty"`
}

type SetHashingAlgorithmResponse struct {
}

type SetHostname struct {
	Name string `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
}

type SetHostnameFromDHCP struct {
	FromDHCP bool `xml:"FromDHCP,omitempty" json:"FromDHCP,omitempty" yaml:"FromDHCP,omitempty"`
}

type SetHostnameFromDHCPResponse struct {
	RebootNeeded bool `xml:"RebootNeeded,omitempty" json:"RebootNeeded,omitempty" yaml:"RebootNeeded,omitempty"`
}

type SetHostnameResponse struct {
}

type SetIPAddressFilter struct {
	IPAddressFilter *common.IPAddressFilter `xml:"IPAddressFilter,omitempty" json:"IPAddressFilter,omitempty" yaml:"IPAddressFilter,omitempty"`
}

type SetIPAddressFilterResponse struct {
}

type SetNTP struct {
	FromDHCP  bool                  `xml:"FromDHCP,omitempty" json:"FromDHCP,omitempty" yaml:"FromDHCP,omitempty"`
	NTPManual []*common.NetworkHost `xml:"NTPManual,omitempty" json:"NTPManual,omitempty" yaml:"NTPManual,omitempty"`
}

type SetNTPResponse struct {
}

type SetNetworkDefaultGateway struct {
	IPv4Address []*common.IPv4Address `xml:"IPv4Address,omitempty" json:"IPv4Address,omitempty" yaml:"IPv4Address,omitempty"`
	IPv6Address []*common.IPv6Address `xml:"IPv6Address,omitempty" json:"IPv6Address,omitempty" yaml:"IPv6Address,omitempty"`
}

type SetNetworkDefaultGatewayResponse struct {
}

type SetNetworkInterfaces struct {
	InterfaceToken   *common.ReferenceToken                   `xml:"InterfaceToken,omitempty" json:"InterfaceToken,omitempty" yaml:"InterfaceToken,omitempty"`
	NetworkInterface *common.NetworkInterfaceSetConfiguration `xml:"NetworkInterface,omitempty" json:"NetworkInterface,omitempty" yaml:"NetworkInterface,omitempty"`
}

type SetNetworkInterfacesResponse struct {
	RebootNeeded bool `xml:"RebootNeeded" json:"RebootNeeded" yaml:"RebootNeeded"`
}

type SetNetworkProtocols struct {
	NetworkProtocols []common.NetworkProtocol `xml:"NetworkProtocols" json:"NetworkProtocols" yaml:"NetworkProtocols"`
}

type SetNetworkProtocolsResponse struct {
}

type SetPasswordComplexityConfiguration struct {
	MinLen                    int  `xml:"MinLen,omitempty" json:"MinLen,omitempty" yaml:"MinLen,omitempty"`
	Uppercase                 int  `xml:"Uppercase,omitempty" json:"Uppercase,omitempty" yaml:"Uppercase,omitempty"`
	Number                    int  `xml:"Number,omitempty" json:"Number,omitempty" yaml:"Number,omitempty"`
	SpecialChars              int  `xml:"SpecialChars,omitempty" json:"SpecialChars,omitempty" yaml:"SpecialChars,omitempty"`
	BlockUsernameOccurrence   bool `xml:"BlockUsernameOccurrence,omitempty" json:"BlockUsernameOccurrence,omitempty" yaml:"BlockUsernameOccurrence,omitempty"`
	PolicyConfigurationLocked bool `xml:"PolicyConfigurationLocked,omitempty" json:"PolicyConfigurationLocked,omitempty" yaml:"PolicyConfigurationLocked,omitempty"`
}

type SetPasswordComplexityConfigurationResponse struct {
}

type SetPasswordHistoryConfiguration struct {
	Enabled bool `xml:"Enabled,omitempty" json:"Enabled,omitempty" yaml:"Enabled,omitempty"`
	Length  int  `xml:"Length,omitempty" json:"Length,omitempty" yaml:"Length,omitempty"`
}

type SetPasswordHistoryConfigurationResponse struct {
}

type SetRelayOutputSettings struct {
	RelayOutputToken *common.ReferenceToken      `xml:"RelayOutputToken,omitempty" json:"RelayOutputToken,omitempty" yaml:"RelayOutputToken,omitempty"`
	Properties       *common.RelayOutputSettings `xml:"Properties,omitempty" json:"Properties,omitempty" yaml:"Properties,omitempty"`
}

type SetRelayOutputSettingsResponse struct {
}

type SetRelayOutputState struct {
	RelayOutputToken common.ReferenceToken    `xml:"RelayOutputToken" json:"RelayOutputToken" yaml:"RelayOutputToken"`
	LogicalState     common.RelayLogicalState `xml:"LogicalState" json:"LogicalState" yaml:"LogicalState"`
}

type SetRelayOutputStateResponse struct {
}

type SetRemoteDiscoveryMode struct {
	RemoteDiscoveryMode *common.DiscoveryMode `xml:"RemoteDiscoveryMode,omitempty" json:"RemoteDiscoveryMode,omitempty" yaml:"RemoteDiscoveryMode,omitempty"`
}

type SetRemoteDiscoveryModeResponse struct {
}

type SetRemoteUser struct {
	RemoteUser *common.RemoteUser `xml:"RemoteUser,omitempty" json:"RemoteUser,omitempty" yaml:"RemoteUser,omitempty"`
}

type SetRemoteUserResponse struct {
}

type SetScopes struct {
	Scopes []string `xml:"Scopes" json:"Scopes" yaml:"Scopes"`
}

type SetScopesResponse struct {
}

type SetStorageConfiguration struct {
	StorageConfiguration StorageConfiguration `xml:"StorageConfiguration,omitempty" json:"StorageConfiguration,omitempty" yaml:"StorageConfiguration,omitempty"`
}

type SetStorageConfigurationResponse struct {
}

type SetSystemDateAndTime struct {
	DateTimeType    *common.SetDateTimeType `xml:"DateTimeType,omitempty" json:"DateTimeType,omitempty" yaml:"DateTimeType,omitempty"`
	DaylightSavings bool                    `xml:"DaylightSavings,omitempty" json:"DaylightSavings,omitempty" yaml:"DaylightSavings,omitempty"`
	TimeZone        *common.TimeZone        `xml:"TimeZone,omitempty" json:"TimeZone,omitempty" yaml:"TimeZone,omitempty"`
	UTCDateTime     *common.DateTime        `xml:"UTCDateTime,omitempty" json:"UTCDateTime,omitempty" yaml:"UTCDateTime,omitempty"`
}

type SetSystemDateAndTimeResponse struct {
}

type SetSystemFactoryDefault struct {
	FactoryDefault *common.FactoryDefaultType `xml:"FactoryDefault,omitempty" json:"FactoryDefault,omitempty" yaml:"FactoryDefault,omitempty"`
}

type SetSystemFactoryDefaultResponse struct {
}

type SetUser struct {
	User []common.User `xml:"User" json:"User" yaml:"User"`
}

type SetUserResponse struct {
}

type SetZeroConfiguration struct {
	InterfaceToken *common.ReferenceToken `xml:"InterfaceToken,omitempty" json:"InterfaceToken,omitempty" yaml:"InterfaceToken,omitempty"`
	Enabled        bool                   `xml:"Enabled,omitempty" json:"Enabled,omitempty" yaml:"Enabled,omitempty"`
}

type SetZeroConfigurationResponse struct {
}

type StartFirmwareUpgrade struct {
}

type StartFirmwareUpgradeResponse struct {
	UploadUri        string           `xml:"UploadUri,omitempty" json:"UploadUri,omitempty" yaml:"UploadUri,omitempty"`
	UploadDelay      *common.Duration `xml:"UploadDelay,omitempty" json:"UploadDelay,omitempty" yaml:"UploadDelay,omitempty"`
	ExpectedDownTime *common.Duration `xml:"ExpectedDownTime,omitempty" json:"ExpectedDownTime,omitempty" yaml:"ExpectedDownTime,omitempty"`
}

type StartSystemRestore struct {
}

type StartSystemRestoreResponse struct {
	UploadUri        string           `xml:"UploadUri,omitempty" json:"UploadUri,omitempty" yaml:"UploadUri,omitempty"`
	ExpectedDownTime *common.Duration `xml:"ExpectedDownTime,omitempty" json:"ExpectedDownTime,omitempty" yaml:"ExpectedDownTime,omitempty"`
}

type StorageConfiguration struct {
	Data          StorageConfigurationData `xml:"Data,omitempty" json:"Data,omitempty" yaml:"Data,omitempty"`
	TypeAttrXSI   string                   `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string                   `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *StorageConfiguration) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:StorageConfiguration"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://www.onvif.org/ver10/device/wsdl"
	}
}

type StorageConfigurationData struct {
	LocalPath  string         `xml:"LocalPath,omitempty" json:"LocalPath,omitempty" yaml:"LocalPath,omitempty"`
	StorageUri string         `xml:"StorageUri,omitempty" json:"StorageUri,omitempty" yaml:"StorageUri,omitempty"`
	User       UserCredential `xml:"User,omitempty" json:"User,omitempty" yaml:"User,omitempty"`
	Extension  []string       `xml:"Extension>Extension,omitempty" json:"Extension>Extension,omitempty" yaml:"Extension>Extension,omitempty"`
	Type       string         `xml:"type,attr,omitempty" json:"type,attr,omitempty" yaml:"type,attr,omitempty"`
}

type SystemCapabilities struct {
	DiscoveryResolve          bool                  `xml:"DiscoveryResolve,attr,omitempty" json:"DiscoveryResolve,attr,omitempty" yaml:"DiscoveryResolve,attr,omitempty"`
	DiscoveryBye              bool                  `xml:"DiscoveryBye,attr,omitempty" json:"DiscoveryBye,attr,omitempty" yaml:"DiscoveryBye,attr,omitempty"`
	RemoteDiscovery           bool                  `xml:"RemoteDiscovery,attr,omitempty" json:"RemoteDiscovery,attr,omitempty" yaml:"RemoteDiscovery,attr,omitempty"`
	SystemBackup              bool                  `xml:"SystemBackup,attr,omitempty" json:"SystemBackup,attr,omitempty" yaml:"SystemBackup,attr,omitempty"`
	SystemLogging             bool                  `xml:"SystemLogging,attr,omitempty" json:"SystemLogging,attr,omitempty" yaml:"SystemLogging,attr,omitempty"`
	FirmwareUpgrade           bool                  `xml:"FirmwareUpgrade,attr,omitempty" json:"FirmwareUpgrade,attr,omitempty" yaml:"FirmwareUpgrade,attr,omitempty"`
	HttpFirmwareUpgrade       bool                  `xml:"HttpFirmwareUpgrade,attr,omitempty" json:"HttpFirmwareUpgrade,attr,omitempty" yaml:"HttpFirmwareUpgrade,attr,omitempty"`
	HttpSystemBackup          bool                  `xml:"HttpSystemBackup,attr,omitempty" json:"HttpSystemBackup,attr,omitempty" yaml:"HttpSystemBackup,attr,omitempty"`
	HttpSystemLogging         bool                  `xml:"HttpSystemLogging,attr,omitempty" json:"HttpSystemLogging,attr,omitempty" yaml:"HttpSystemLogging,attr,omitempty"`
	HttpSupportInformation    bool                  `xml:"HttpSupportInformation,attr,omitempty" json:"HttpSupportInformation,attr,omitempty" yaml:"HttpSupportInformation,attr,omitempty"`
	StorageConfiguration      bool                  `xml:"StorageConfiguration,attr,omitempty" json:"StorageConfiguration,attr,omitempty" yaml:"StorageConfiguration,attr,omitempty"`
	MaxStorageConfigurations  int                   `xml:"MaxStorageConfigurations,attr,omitempty" json:"MaxStorageConfigurations,attr,omitempty" yaml:"MaxStorageConfigurations,attr,omitempty"`
	GeoLocationEntries        int                   `xml:"GeoLocationEntries,attr,omitempty" json:"GeoLocationEntries,attr,omitempty" yaml:"GeoLocationEntries,attr,omitempty"`
	AutoGeo                   common.StringAttrList `xml:"AutoGeo,attr,omitempty" json:"AutoGeo,attr,omitempty" yaml:"AutoGeo,attr,omitempty"`
	StorageTypesSupported     common.StringAttrList `xml:"StorageTypesSupported,attr,omitempty" json:"StorageTypesSupported,attr,omitempty" yaml:"StorageTypesSupported,attr,omitempty"`
	DiscoveryNotSupported     bool                  `xml:"DiscoveryNotSupported,attr,omitempty" json:"DiscoveryNotSupported,attr,omitempty" yaml:"DiscoveryNotSupported,attr,omitempty"`
	NetworkConfigNotSupported bool                  `xml:"NetworkConfigNotSupported,attr,omitempty" json:"NetworkConfigNotSupported,attr,omitempty" yaml:"NetworkConfigNotSupported,attr,omitempty"`
	UserConfigNotSupported    bool                  `xml:"UserConfigNotSupported,attr,omitempty" json:"UserConfigNotSupported,attr,omitempty" yaml:"UserConfigNotSupported,attr,omitempty"`
	Addons                    common.StringAttrList `xml:"Addons,attr,omitempty" json:"Addons,attr,omitempty" yaml:"Addons,attr,omitempty"`
}

type SystemReboot struct {
}

type SystemRebootResponse struct {
	Message string `xml:"Message,omitempty" json:"Message,omitempty" yaml:"Message,omitempty"`
}

type UpgradeSystemFirmware struct {
	Firmware *common.AttachmentData `xml:"Firmware,omitempty" json:"Firmware,omitempty" yaml:"Firmware,omitempty"`
}

type UpgradeSystemFirmwareResponse struct {
	Message string `xml:"Message,omitempty" json:"Message,omitempty" yaml:"Message,omitempty"`
}

type UserCredential struct {
	UserName  string   `xml:"UserName,omitempty" json:"UserName,omitempty" yaml:"UserName,omitempty"`
	Password  string   `xml:"Password,omitempty" json:"Password,omitempty" yaml:"Password,omitempty"`
	Extension []string `xml:"Extension>Extension,omitempty" json:"Extension>Extension,omitempty" yaml:"Extension>Extension,omitempty"`
}

// device implements the Device interface.
type device struct {
	cli      common.Client
	Endpoint string
}

func (p *device) OptAddIPAddressFilter(args AddIPAddressFilter) (*AddIPAddressFilterResponse, *common.Fault) {
	req := struct {
		XMLName            string `xml:"tds:AddIPAddressFilter"`
		AddIPAddressFilter AddIPAddressFilter
	}{
		AddIPAddressFilter: args,
	}

	resp := AddIPAddressFilterResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptAddScopes(args AddScopes) (*AddScopesResponse, *common.Fault) {
	req := struct {
		XMLName   string `xml:"tds:AddScopes"`
		AddScopes AddScopes
	}{
		AddScopes: args,
	}

	resp := AddScopesResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptCreateCertificate(args CreateCertificate) (*CreateCertificateResponse, *common.Fault) {
	req := struct {
		XMLName           string `xml:"tds:CreateCertificate"`
		CreateCertificate CreateCertificate
	}{
		CreateCertificate: args,
	}

	resp := CreateCertificateResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptCreateDot1XConfiguration(args CreateDot1XConfiguration) (*CreateDot1XConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                  string `xml:"tds:CreateDot1XConfiguration"`
		CreateDot1XConfiguration CreateDot1XConfiguration
	}{
		CreateDot1XConfiguration: args,
	}

	resp := CreateDot1XConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptCreateStorageConfiguration(args CreateStorageConfiguration) (*CreateStorageConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                    string `xml:"tds:CreateStorageConfiguration"`
		CreateStorageConfiguration CreateStorageConfiguration
	}{
		CreateStorageConfiguration: args,
	}

	resp := CreateStorageConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptCreateUsers(args CreateUsers) (*CreateUsersResponse, *common.Fault) {
	req := struct {
		XMLName     string `xml:"tds:CreateUsers"`
		CreateUsers CreateUsers
	}{
		CreateUsers: args,
	}

	resp := CreateUsersResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptDeleteCertificates(args DeleteCertificates) (*DeleteCertificatesResponse, *common.Fault) {
	req := struct {
		XMLName            string `xml:"tds:DeleteCertificates"`
		DeleteCertificates DeleteCertificates
	}{
		DeleteCertificates: args,
	}

	resp := DeleteCertificatesResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptDeleteDot1XConfiguration(args DeleteDot1XConfiguration) (*DeleteDot1XConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                  string `xml:"tds:DeleteDot1XConfiguration"`
		DeleteDot1XConfiguration DeleteDot1XConfiguration
	}{
		DeleteDot1XConfiguration: args,
	}

	resp := DeleteDot1XConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptDeleteGeoLocation(args DeleteGeoLocation) (*DeleteGeoLocationResponse, *common.Fault) {
	req := struct {
		XMLName           string `xml:"tds:DeleteGeoLocation"`
		DeleteGeoLocation DeleteGeoLocation
	}{
		DeleteGeoLocation: args,
	}

	resp := DeleteGeoLocationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptDeleteStorageConfiguration(args DeleteStorageConfiguration) (*DeleteStorageConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                    string `xml:"tds:DeleteStorageConfiguration"`
		DeleteStorageConfiguration DeleteStorageConfiguration
	}{
		DeleteStorageConfiguration: args,
	}

	resp := DeleteStorageConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptDeleteUsers(args DeleteUsers) (*DeleteUsersResponse, *common.Fault) {
	req := struct {
		XMLName     string `xml:"tds:DeleteUsers"`
		DeleteUsers DeleteUsers
	}{
		DeleteUsers: args,
	}

	resp := DeleteUsersResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptGetAccessPolicy(args GetAccessPolicy) (*GetAccessPolicyResponse, *common.Fault) {
	req := struct {
		XMLName         string `xml:"tds:GetAccessPolicy"`
		GetAccessPolicy GetAccessPolicy
	}{
		GetAccessPolicy: args,
	}

	resp := GetAccessPolicyResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptGetAuthFailureWarningConfiguration(args GetAuthFailureWarningConfiguration) (*GetAuthFailureWarningConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                            string `xml:"tds:GetAuthFailureWarningConfiguration"`
		GetAuthFailureWarningConfiguration GetAuthFailureWarningConfiguration
	}{
		GetAuthFailureWarningConfiguration: args,
	}

	resp := GetAuthFailureWarningConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptGetAuthFailureWarningOptions(args GetAuthFailureWarningOptions) (*GetAuthFailureWarningOptionsResponse, *common.Fault) {
	req := struct {
		XMLName                      string `xml:"tds:GetAuthFailureWarningOptions"`
		GetAuthFailureWarningOptions GetAuthFailureWarningOptions
	}{
		GetAuthFailureWarningOptions: args,
	}

	resp := GetAuthFailureWarningOptionsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptGetCACertificates(args GetCACertificates) (*GetCACertificatesResponse, *common.Fault) {
	req := struct {
		XMLName           string `xml:"tds:GetCACertificates"`
		GetCACertificates GetCACertificates
	}{
		GetCACertificates: args,
	}

	resp := GetCACertificatesResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptGetCapabilities(args GetCapabilities) (*GetCapabilitiesResponse, *common.Fault) {
	req := struct {
		XMLName         string `xml:"tds:GetCapabilities"`
		GetCapabilities GetCapabilities
	}{
		GetCapabilities: args,
	}

	resp := GetCapabilitiesResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptGetCertificateInformation(args GetCertificateInformation) (*GetCertificateInformationResponse, *common.Fault) {
	req := struct {
		XMLName                   string `xml:"tds:GetCertificateInformation"`
		GetCertificateInformation GetCertificateInformation
	}{
		GetCertificateInformation: args,
	}

	resp := GetCertificateInformationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptGetCertificates(args GetCertificates) (*GetCertificatesResponse, *common.Fault) {
	req := struct {
		XMLName         string `xml:"tds:GetCertificates"`
		GetCertificates GetCertificates
	}{
		GetCertificates: args,
	}

	resp := GetCertificatesResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptGetCertificatesStatus(args GetCertificatesStatus) (*GetCertificatesStatusResponse, *common.Fault) {
	req := struct {
		XMLName               string `xml:"tds:GetCertificatesStatus"`
		GetCertificatesStatus GetCertificatesStatus
	}{
		GetCertificatesStatus: args,
	}

	resp := GetCertificatesStatusResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptGetClientCertificateMode(args GetClientCertificateMode) (*GetClientCertificateModeResponse, *common.Fault) {
	req := struct {
		XMLName                  string `xml:"tds:GetClientCertificateMode"`
		GetClientCertificateMode GetClientCertificateMode
	}{
		GetClientCertificateMode: args,
	}

	resp := GetClientCertificateModeResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptGetDNS(args GetDNS) (*GetDNSResponse, *common.Fault) {
	req := struct {
		XMLName string `xml:"tds:GetDNS"`
		GetDNS  GetDNS
	}{
		GetDNS: args,
	}

	resp := GetDNSResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptGetDPAddresses(args GetDPAddresses) (*GetDPAddressesResponse, *common.Fault) {
	req := struct {
		XMLName        string `xml:"tds:GetDPAddresses"`
		GetDPAddresses GetDPAddresses
	}{
		GetDPAddresses: args,
	}

	resp := GetDPAddressesResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptGetDeviceInformation(args GetDeviceInformation) (*GetDeviceInformationResponse, *common.Fault) {
	req := struct {
		XMLName              string `xml:"tds:GetDeviceInformation"`
		GetDeviceInformation GetDeviceInformation
	}{
		GetDeviceInformation: args,
	}

	resp := GetDeviceInformationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptGetDiscoveryMode(args GetDiscoveryMode) (*GetDiscoveryModeResponse, *common.Fault) {
	req := struct {
		XMLName          string `xml:"tds:GetDiscoveryMode"`
		GetDiscoveryMode GetDiscoveryMode
	}{
		GetDiscoveryMode: args,
	}

	resp := GetDiscoveryModeResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptGetDot11Capabilities(args GetDot11Capabilities) (*GetDot11CapabilitiesResponse, *common.Fault) {
	req := struct {
		XMLName              string `xml:"tds:GetDot11Capabilities"`
		GetDot11Capabilities GetDot11Capabilities
	}{
		GetDot11Capabilities: args,
	}

	resp := GetDot11CapabilitiesResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptGetDot11Status(args GetDot11Status) (*GetDot11StatusResponse, *common.Fault) {
	req := struct {
		XMLName        string `xml:"tds:GetDot11Status"`
		GetDot11Status GetDot11Status
	}{
		GetDot11Status: args,
	}

	resp := GetDot11StatusResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptGetDot1XConfiguration(args GetDot1XConfiguration) (*GetDot1XConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName               string `xml:"tds:GetDot1XConfiguration"`
		GetDot1XConfiguration GetDot1XConfiguration
	}{
		GetDot1XConfiguration: args,
	}

	resp := GetDot1XConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptGetDot1XConfigurations(args GetDot1XConfigurations) (*GetDot1XConfigurationsResponse, *common.Fault) {
	req := struct {
		XMLName                string `xml:"tds:GetDot1XConfigurations"`
		GetDot1XConfigurations GetDot1XConfigurations
	}{
		GetDot1XConfigurations: args,
	}

	resp := GetDot1XConfigurationsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptGetDynamicDNS(args GetDynamicDNS) (*GetDynamicDNSResponse, *common.Fault) {
	req := struct {
		XMLName       string `xml:"tds:GetDynamicDNS"`
		GetDynamicDNS GetDynamicDNS
	}{
		GetDynamicDNS: args,
	}

	resp := GetDynamicDNSResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptGetEndpointReference(args GetEndpointReference) (*GetEndpointReferenceResponse, *common.Fault) {
	req := struct {
		XMLName              string `xml:"tds:GetEndpointReference"`
		GetEndpointReference GetEndpointReference
	}{
		GetEndpointReference: args,
	}

	resp := GetEndpointReferenceResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptGetGeoLocation(args GetGeoLocation) (*GetGeoLocationResponse, *common.Fault) {
	req := struct {
		XMLName        string `xml:"tds:GetGeoLocation"`
		GetGeoLocation GetGeoLocation
	}{
		GetGeoLocation: args,
	}

	resp := GetGeoLocationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptGetHostname(args GetHostname) (*GetHostnameResponse, *common.Fault) {
	req := struct {
		XMLName     string `xml:"tds:GetHostname"`
		GetHostname GetHostname
	}{
		GetHostname: args,
	}

	resp := GetHostnameResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptGetIPAddressFilter(args GetIPAddressFilter) (*GetIPAddressFilterResponse, *common.Fault) {
	req := struct {
		XMLName            string `xml:"tds:GetIPAddressFilter"`
		GetIPAddressFilter GetIPAddressFilter
	}{
		GetIPAddressFilter: args,
	}

	resp := GetIPAddressFilterResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptGetNTP(args GetNTP) (*GetNTPResponse, *common.Fault) {
	req := struct {
		XMLName string `xml:"tds:GetNTP"`
		GetNTP  GetNTP
	}{
		GetNTP: args,
	}

	resp := GetNTPResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptGetNetworkDefaultGateway(args GetNetworkDefaultGateway) (*GetNetworkDefaultGatewayResponse, *common.Fault) {
	req := struct {
		XMLName                  string `xml:"tds:GetNetworkDefaultGateway"`
		GetNetworkDefaultGateway GetNetworkDefaultGateway
	}{
		GetNetworkDefaultGateway: args,
	}

	resp := GetNetworkDefaultGatewayResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptGetNetworkInterfaces(args GetNetworkInterfaces) (*GetNetworkInterfacesResponse, *common.Fault) {
	req := struct {
		XMLName              string `xml:"tds:GetNetworkInterfaces"`
		GetNetworkInterfaces GetNetworkInterfaces
	}{
		GetNetworkInterfaces: args,
	}

	resp := GetNetworkInterfacesResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptGetNetworkProtocols(args GetNetworkProtocols) (*GetNetworkProtocolsResponse, *common.Fault) {
	req := struct {
		XMLName             string `xml:"tds:GetNetworkProtocols"`
		GetNetworkProtocols GetNetworkProtocols
	}{
		GetNetworkProtocols: args,
	}

	resp := GetNetworkProtocolsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptGetPasswordComplexityConfiguration(args GetPasswordComplexityConfiguration) (*GetPasswordComplexityConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                            string `xml:"tds:GetPasswordComplexityConfiguration"`
		GetPasswordComplexityConfiguration GetPasswordComplexityConfiguration
	}{
		GetPasswordComplexityConfiguration: args,
	}

	resp := GetPasswordComplexityConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptGetPasswordComplexityOptions(args GetPasswordComplexityOptions) (*GetPasswordComplexityOptionsResponse, *common.Fault) {
	req := struct {
		XMLName                      string `xml:"tds:GetPasswordComplexityOptions"`
		GetPasswordComplexityOptions GetPasswordComplexityOptions
	}{
		GetPasswordComplexityOptions: args,
	}

	resp := GetPasswordComplexityOptionsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptGetPasswordHistoryConfiguration(args GetPasswordHistoryConfiguration) (*GetPasswordHistoryConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                         string `xml:"tds:GetPasswordHistoryConfiguration"`
		GetPasswordHistoryConfiguration GetPasswordHistoryConfiguration
	}{
		GetPasswordHistoryConfiguration: args,
	}

	resp := GetPasswordHistoryConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptGetPkcs10Request(args GetPkcs10Request) (*GetPkcs10RequestResponse, *common.Fault) {
	req := struct {
		XMLName          string `xml:"tds:GetPkcs10Request"`
		GetPkcs10Request GetPkcs10Request
	}{
		GetPkcs10Request: args,
	}

	resp := GetPkcs10RequestResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptGetRelayOutputs(args GetRelayOutputs) (*GetRelayOutputsResponse, *common.Fault) {
	req := struct {
		XMLName         string `xml:"tds:GetRelayOutputs"`
		GetRelayOutputs GetRelayOutputs
	}{
		GetRelayOutputs: args,
	}

	resp := GetRelayOutputsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptGetRemoteDiscoveryMode(args GetRemoteDiscoveryMode) (*GetRemoteDiscoveryModeResponse, *common.Fault) {
	req := struct {
		XMLName                string `xml:"tds:GetRemoteDiscoveryMode"`
		GetRemoteDiscoveryMode GetRemoteDiscoveryMode
	}{
		GetRemoteDiscoveryMode: args,
	}

	resp := GetRemoteDiscoveryModeResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptGetRemoteUser(args GetRemoteUser) (*GetRemoteUserResponse, *common.Fault) {
	req := struct {
		XMLName       string `xml:"tds:GetRemoteUser"`
		GetRemoteUser GetRemoteUser
	}{
		GetRemoteUser: args,
	}

	resp := GetRemoteUserResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptGetScopes(args GetScopes) (*GetScopesResponse, *common.Fault) {
	req := struct {
		XMLName   string `xml:"tds:GetScopes"`
		GetScopes GetScopes
	}{
		GetScopes: args,
	}

	resp := GetScopesResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptGetServiceCapabilities(args GetServiceCapabilities) (*GetServiceCapabilitiesResponse, *common.Fault) {
	req := struct {
		XMLName                string `xml:"tds:GetServiceCapabilities"`
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

func (p *device) OptGetServices(args GetServices) (*GetServicesResponse, *common.Fault) {
	req := struct {
		XMLName     string `xml:"tds:GetServices"`
		GetServices GetServices
	}{
		GetServices: args,
	}

	resp := GetServicesResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptGetStorageConfiguration(args GetStorageConfiguration) (*GetStorageConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                 string `xml:"tds:GetStorageConfiguration"`
		GetStorageConfiguration GetStorageConfiguration
	}{
		GetStorageConfiguration: args,
	}

	resp := GetStorageConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptGetStorageConfigurations(args GetStorageConfigurations) (*GetStorageConfigurationsResponse, *common.Fault) {
	req := struct {
		XMLName                  string `xml:"tds:GetStorageConfigurations"`
		GetStorageConfigurations GetStorageConfigurations
	}{
		GetStorageConfigurations: args,
	}

	resp := GetStorageConfigurationsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptGetSystemBackup(args GetSystemBackup) (*GetSystemBackupResponse, *common.Fault) {
	req := struct {
		XMLName         string `xml:"tds:GetSystemBackup"`
		GetSystemBackup GetSystemBackup
	}{
		GetSystemBackup: args,
	}

	resp := GetSystemBackupResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptGetSystemDateAndTime(args GetSystemDateAndTime) (*GetSystemDateAndTimeResponse, *common.Fault) {
	req := struct {
		XMLName              string `xml:"tds:GetSystemDateAndTime"`
		GetSystemDateAndTime GetSystemDateAndTime
	}{
		GetSystemDateAndTime: args,
	}

	resp := GetSystemDateAndTimeResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptGetSystemLog(args GetSystemLog) (*GetSystemLogResponse, *common.Fault) {
	req := struct {
		XMLName      string `xml:"tds:GetSystemLog"`
		GetSystemLog GetSystemLog
	}{
		GetSystemLog: args,
	}

	resp := GetSystemLogResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptGetSystemSupportInformation(args GetSystemSupportInformation) (*GetSystemSupportInformationResponse, *common.Fault) {
	req := struct {
		XMLName                     string `xml:"tds:GetSystemSupportInformation"`
		GetSystemSupportInformation GetSystemSupportInformation
	}{
		GetSystemSupportInformation: args,
	}

	resp := GetSystemSupportInformationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptGetSystemUris(args GetSystemUris) (*GetSystemUrisResponse, *common.Fault) {
	req := struct {
		XMLName       string `xml:"tds:GetSystemUris"`
		GetSystemUris GetSystemUris
	}{
		GetSystemUris: args,
	}

	resp := GetSystemUrisResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptGetUsers(args GetUsers) (*GetUsersResponse, *common.Fault) {
	req := struct {
		XMLName  string `xml:"tds:GetUsers"`
		GetUsers GetUsers
	}{
		GetUsers: args,
	}

	resp := GetUsersResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptGetWsdlUrl(args GetWsdlUrl) (*GetWsdlUrlResponse, *common.Fault) {
	req := struct {
		XMLName    string `xml:"tds:GetWsdlUrl"`
		GetWsdlUrl GetWsdlUrl
	}{
		GetWsdlUrl: args,
	}

	resp := GetWsdlUrlResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptGetZeroConfiguration(args GetZeroConfiguration) (*GetZeroConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName              string `xml:"tds:GetZeroConfiguration"`
		GetZeroConfiguration GetZeroConfiguration
	}{
		GetZeroConfiguration: args,
	}

	resp := GetZeroConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptLoadCACertificates(args LoadCACertificates) (*LoadCACertificatesResponse, *common.Fault) {
	req := struct {
		XMLName            string `xml:"tds:LoadCACertificates"`
		LoadCACertificates LoadCACertificates
	}{
		LoadCACertificates: args,
	}

	resp := LoadCACertificatesResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptLoadCertificateWithPrivateKey(args LoadCertificateWithPrivateKey) (*LoadCertificateWithPrivateKeyResponse, *common.Fault) {
	req := struct {
		XMLName                       string `xml:"tds:LoadCertificateWithPrivateKey"`
		LoadCertificateWithPrivateKey LoadCertificateWithPrivateKey
	}{
		LoadCertificateWithPrivateKey: args,
	}

	resp := LoadCertificateWithPrivateKeyResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptLoadCertificates(args LoadCertificates) (*LoadCertificatesResponse, *common.Fault) {
	req := struct {
		XMLName          string `xml:"tds:LoadCertificates"`
		LoadCertificates LoadCertificates
	}{
		LoadCertificates: args,
	}

	resp := LoadCertificatesResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptRemoveIPAddressFilter(args RemoveIPAddressFilter) (*RemoveIPAddressFilterResponse, *common.Fault) {
	req := struct {
		XMLName               string `xml:"tds:RemoveIPAddressFilter"`
		RemoveIPAddressFilter RemoveIPAddressFilter
	}{
		RemoveIPAddressFilter: args,
	}

	resp := RemoveIPAddressFilterResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptRemoveScopes(args RemoveScopes) (*RemoveScopesResponse, *common.Fault) {
	req := struct {
		XMLName      string `xml:"tds:RemoveScopes"`
		RemoveScopes RemoveScopes
	}{
		RemoveScopes: args,
	}

	resp := RemoveScopesResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptRestoreSystem(args RestoreSystem) (*RestoreSystemResponse, *common.Fault) {
	req := struct {
		XMLName       string `xml:"tds:RestoreSystem"`
		RestoreSystem RestoreSystem
	}{
		RestoreSystem: args,
	}

	resp := RestoreSystemResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptScanAvailableDot11Networks(args ScanAvailableDot11Networks) (*ScanAvailableDot11NetworksResponse, *common.Fault) {
	req := struct {
		XMLName                    string `xml:"tds:ScanAvailableDot11Networks"`
		ScanAvailableDot11Networks ScanAvailableDot11Networks
	}{
		ScanAvailableDot11Networks: args,
	}

	resp := ScanAvailableDot11NetworksResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptSendAuxiliaryCommand(args SendAuxiliaryCommand) (*SendAuxiliaryCommandResponse, *common.Fault) {
	req := struct {
		XMLName              string `xml:"tds:SendAuxiliaryCommand"`
		SendAuxiliaryCommand SendAuxiliaryCommand
	}{
		SendAuxiliaryCommand: args,
	}

	resp := SendAuxiliaryCommandResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptSetAccessPolicy(args SetAccessPolicy) (*SetAccessPolicyResponse, *common.Fault) {
	req := struct {
		XMLName         string `xml:"tds:SetAccessPolicy"`
		SetAccessPolicy SetAccessPolicy
	}{
		SetAccessPolicy: args,
	}

	resp := SetAccessPolicyResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptSetAuthFailureWarningConfiguration(args SetAuthFailureWarningConfiguration) (*SetAuthFailureWarningConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                            string `xml:"tds:SetAuthFailureWarningConfiguration"`
		SetAuthFailureWarningConfiguration SetAuthFailureWarningConfiguration
	}{
		SetAuthFailureWarningConfiguration: args,
	}

	resp := SetAuthFailureWarningConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptSetCertificatesStatus(args SetCertificatesStatus) (*SetCertificatesStatusResponse, *common.Fault) {
	req := struct {
		XMLName               string `xml:"tds:SetCertificatesStatus"`
		SetCertificatesStatus SetCertificatesStatus
	}{
		SetCertificatesStatus: args,
	}

	resp := SetCertificatesStatusResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptSetClientCertificateMode(args SetClientCertificateMode) (*SetClientCertificateModeResponse, *common.Fault) {
	req := struct {
		XMLName                  string `xml:"tds:SetClientCertificateMode"`
		SetClientCertificateMode SetClientCertificateMode
	}{
		SetClientCertificateMode: args,
	}

	resp := SetClientCertificateModeResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptSetDNS(args SetDNS) (*SetDNSResponse, *common.Fault) {
	req := struct {
		XMLName string `xml:"tds:SetDNS"`
		SetDNS  SetDNS
	}{
		SetDNS: args,
	}

	resp := SetDNSResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptSetDPAddresses(args SetDPAddresses) (*SetDPAddressesResponse, *common.Fault) {
	req := struct {
		XMLName        string `xml:"tds:SetDPAddresses"`
		SetDPAddresses SetDPAddresses
	}{
		SetDPAddresses: args,
	}

	resp := SetDPAddressesResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptSetDiscoveryMode(args SetDiscoveryMode) (*SetDiscoveryModeResponse, *common.Fault) {
	req := struct {
		XMLName          string `xml:"tds:SetDiscoveryMode"`
		SetDiscoveryMode SetDiscoveryMode
	}{
		SetDiscoveryMode: args,
	}

	resp := SetDiscoveryModeResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptSetDot1XConfiguration(args SetDot1XConfiguration) (*SetDot1XConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName               string `xml:"tds:SetDot1XConfiguration"`
		SetDot1XConfiguration SetDot1XConfiguration
	}{
		SetDot1XConfiguration: args,
	}

	resp := SetDot1XConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptSetDynamicDNS(args SetDynamicDNS) (*SetDynamicDNSResponse, *common.Fault) {
	req := struct {
		XMLName       string `xml:"tds:SetDynamicDNS"`
		SetDynamicDNS SetDynamicDNS
	}{
		SetDynamicDNS: args,
	}

	resp := SetDynamicDNSResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptSetGeoLocation(args SetGeoLocation) (*SetGeoLocationResponse, *common.Fault) {
	req := struct {
		XMLName        string `xml:"tds:SetGeoLocation"`
		SetGeoLocation SetGeoLocation
	}{
		SetGeoLocation: args,
	}

	resp := SetGeoLocationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptSetHashingAlgorithm(args SetHashingAlgorithm) (*SetHashingAlgorithmResponse, *common.Fault) {
	req := struct {
		XMLName             string `xml:"tds:SetHashingAlgorithm"`
		SetHashingAlgorithm SetHashingAlgorithm
	}{
		SetHashingAlgorithm: args,
	}

	resp := SetHashingAlgorithmResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptSetHostname(args SetHostname) (*SetHostnameResponse, *common.Fault) {
	req := struct {
		XMLName     string `xml:"tds:SetHostname"`
		SetHostname SetHostname
	}{
		SetHostname: args,
	}

	resp := SetHostnameResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptSetHostnameFromDHCP(args SetHostnameFromDHCP) (*SetHostnameFromDHCPResponse, *common.Fault) {
	req := struct {
		XMLName             string `xml:"tds:SetHostnameFromDHCP"`
		SetHostnameFromDHCP SetHostnameFromDHCP
	}{
		SetHostnameFromDHCP: args,
	}

	resp := SetHostnameFromDHCPResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptSetIPAddressFilter(args SetIPAddressFilter) (*SetIPAddressFilterResponse, *common.Fault) {
	req := struct {
		XMLName            string `xml:"tds:SetIPAddressFilter"`
		SetIPAddressFilter SetIPAddressFilter
	}{
		SetIPAddressFilter: args,
	}

	resp := SetIPAddressFilterResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptSetNTP(args SetNTP) (*SetNTPResponse, *common.Fault) {
	req := struct {
		XMLName string `xml:"tds:SetNTP"`
		SetNTP  SetNTP
	}{
		SetNTP: args,
	}

	resp := SetNTPResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptSetNetworkDefaultGateway(args SetNetworkDefaultGateway) (*SetNetworkDefaultGatewayResponse, *common.Fault) {
	req := struct {
		XMLName                  string `xml:"tds:SetNetworkDefaultGateway"`
		SetNetworkDefaultGateway SetNetworkDefaultGateway
	}{
		SetNetworkDefaultGateway: args,
	}

	resp := SetNetworkDefaultGatewayResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptSetNetworkInterfaces(args SetNetworkInterfaces) (*SetNetworkInterfacesResponse, *common.Fault) {
	req := struct {
		XMLName              string `xml:"tds:SetNetworkInterfaces"`
		SetNetworkInterfaces SetNetworkInterfaces
	}{
		SetNetworkInterfaces: args,
	}

	resp := SetNetworkInterfacesResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptSetNetworkProtocols(args SetNetworkProtocols) (*SetNetworkProtocolsResponse, *common.Fault) {
	req := struct {
		XMLName             string `xml:"tds:SetNetworkProtocols"`
		SetNetworkProtocols SetNetworkProtocols
	}{
		SetNetworkProtocols: args,
	}

	resp := SetNetworkProtocolsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptSetPasswordComplexityConfiguration(args SetPasswordComplexityConfiguration) (*SetPasswordComplexityConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                            string `xml:"tds:SetPasswordComplexityConfiguration"`
		SetPasswordComplexityConfiguration SetPasswordComplexityConfiguration
	}{
		SetPasswordComplexityConfiguration: args,
	}

	resp := SetPasswordComplexityConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptSetPasswordHistoryConfiguration(args SetPasswordHistoryConfiguration) (*SetPasswordHistoryConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                         string `xml:"tds:SetPasswordHistoryConfiguration"`
		SetPasswordHistoryConfiguration SetPasswordHistoryConfiguration
	}{
		SetPasswordHistoryConfiguration: args,
	}

	resp := SetPasswordHistoryConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptSetRelayOutputSettings(args SetRelayOutputSettings) (*SetRelayOutputSettingsResponse, *common.Fault) {
	req := struct {
		XMLName                string `xml:"tds:SetRelayOutputSettings"`
		SetRelayOutputSettings SetRelayOutputSettings
	}{
		SetRelayOutputSettings: args,
	}

	resp := SetRelayOutputSettingsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptSetRelayOutputState(args SetRelayOutputState) (*SetRelayOutputStateResponse, *common.Fault) {
	req := struct {
		XMLName             string `xml:"tds:SetRelayOutputState"`
		SetRelayOutputState SetRelayOutputState
	}{
		SetRelayOutputState: args,
	}

	resp := SetRelayOutputStateResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptSetRemoteDiscoveryMode(args SetRemoteDiscoveryMode) (*SetRemoteDiscoveryModeResponse, *common.Fault) {
	req := struct {
		XMLName                string `xml:"tds:SetRemoteDiscoveryMode"`
		SetRemoteDiscoveryMode SetRemoteDiscoveryMode
	}{
		SetRemoteDiscoveryMode: args,
	}

	resp := SetRemoteDiscoveryModeResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptSetRemoteUser(args SetRemoteUser) (*SetRemoteUserResponse, *common.Fault) {
	req := struct {
		XMLName       string `xml:"tds:SetRemoteUser"`
		SetRemoteUser SetRemoteUser
	}{
		SetRemoteUser: args,
	}

	resp := SetRemoteUserResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptSetScopes(args SetScopes) (*SetScopesResponse, *common.Fault) {
	req := struct {
		XMLName   string `xml:"tds:SetScopes"`
		SetScopes SetScopes
	}{
		SetScopes: args,
	}

	resp := SetScopesResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptSetStorageConfiguration(args SetStorageConfiguration) (*SetStorageConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                 string `xml:"tds:SetStorageConfiguration"`
		SetStorageConfiguration SetStorageConfiguration
	}{
		SetStorageConfiguration: args,
	}

	resp := SetStorageConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptSetSystemDateAndTime(args SetSystemDateAndTime) (*SetSystemDateAndTimeResponse, *common.Fault) {
	req := struct {
		XMLName              string `xml:"tds:SetSystemDateAndTime"`
		SetSystemDateAndTime SetSystemDateAndTime
	}{
		SetSystemDateAndTime: args,
	}

	resp := SetSystemDateAndTimeResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptSetSystemFactoryDefault(args SetSystemFactoryDefault) (*SetSystemFactoryDefaultResponse, *common.Fault) {
	req := struct {
		XMLName                 string `xml:"tds:SetSystemFactoryDefault"`
		SetSystemFactoryDefault SetSystemFactoryDefault
	}{
		SetSystemFactoryDefault: args,
	}

	resp := SetSystemFactoryDefaultResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptSetUser(args SetUser) (*SetUserResponse, *common.Fault) {
	req := struct {
		XMLName string `xml:"tds:SetUser"`
		SetUser SetUser
	}{
		SetUser: args,
	}

	resp := SetUserResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptSetZeroConfiguration(args SetZeroConfiguration) (*SetZeroConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName              string `xml:"tds:SetZeroConfiguration"`
		SetZeroConfiguration SetZeroConfiguration
	}{
		SetZeroConfiguration: args,
	}

	resp := SetZeroConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptStartFirmwareUpgrade(args StartFirmwareUpgrade) (*StartFirmwareUpgradeResponse, *common.Fault) {
	req := struct {
		XMLName              string `xml:"tds:StartFirmwareUpgrade"`
		StartFirmwareUpgrade StartFirmwareUpgrade
	}{
		StartFirmwareUpgrade: args,
	}

	resp := StartFirmwareUpgradeResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptStartSystemRestore(args StartSystemRestore) (*StartSystemRestoreResponse, *common.Fault) {
	req := struct {
		XMLName            string `xml:"tds:StartSystemRestore"`
		StartSystemRestore StartSystemRestore
	}{
		StartSystemRestore: args,
	}

	resp := StartSystemRestoreResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptSystemReboot(args SystemReboot) (*SystemRebootResponse, *common.Fault) {
	req := struct {
		XMLName      string `xml:"tds:SystemReboot"`
		SystemReboot SystemReboot
	}{
		SystemReboot: args,
	}

	resp := SystemRebootResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *device) OptUpgradeSystemFirmware(args UpgradeSystemFirmware) (*UpgradeSystemFirmwareResponse, *common.Fault) {
	req := struct {
		XMLName               string `xml:"tds:UpgradeSystemFirmware"`
		UpgradeSystemFirmware UpgradeSystemFirmware
	}{
		UpgradeSystemFirmware: args,
	}

	resp := UpgradeSystemFirmwareResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}
