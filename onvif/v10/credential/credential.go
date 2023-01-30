package credential

import (
	"code.byted.org/videoarch/go-onvif/onvif/common"
)

var Namespace = "http://www.onvif.org/ver10/credential/wsdl"

// NewCredentialPort creates an initializes a CredentialPort.
func NewCredentialPort(endpoint string, cli common.Client) CredentialPort {
	return &credentialPort{cli: cli, Endpoint: endpoint}
}

// CredentialPort was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type CredentialPort interface {
	OptAddToBlacklist(AddToBlacklist AddToBlacklist) (*AddToBlacklistResponse, *common.Fault)

	OptAddToWhitelist(AddToWhitelist AddToWhitelist) (*AddToWhitelistResponse, *common.Fault)

	OptCreateCredential(CreateCredential CreateCredential) (*CreateCredentialResponse, *common.Fault)

	OptDeleteBlacklist(DeleteBlacklist DeleteBlacklist) (*DeleteBlacklistResponse, *common.Fault)

	OptDeleteCredential(DeleteCredential DeleteCredential) (*DeleteCredentialResponse, *common.Fault)

	OptDeleteCredentialAccessProfiles(DeleteCredentialAccessProfiles DeleteCredentialAccessProfiles) (*DeleteCredentialAccessProfilesResponse, *common.Fault)

	OptDeleteCredentialIdentifier(DeleteCredentialIdentifier DeleteCredentialIdentifier) (*DeleteCredentialIdentifierResponse, *common.Fault)

	OptDeleteWhitelist(DeleteWhitelist DeleteWhitelist) (*DeleteWhitelistResponse, *common.Fault)

	OptDisableCredential(DisableCredential DisableCredential) (*DisableCredentialResponse, *common.Fault)

	OptEnableCredential(EnableCredential EnableCredential) (*EnableCredentialResponse, *common.Fault)

	OptGetBlacklist(GetBlacklist GetBlacklist) (*GetBlacklistResponse, *common.Fault)

	OptGetCredentialAccessProfiles(GetCredentialAccessProfiles GetCredentialAccessProfiles) (*GetCredentialAccessProfilesResponse, *common.Fault)

	OptGetCredentialIdentifiers(GetCredentialIdentifiers GetCredentialIdentifiers) (*GetCredentialIdentifiersResponse, *common.Fault)

	OptGetCredentialInfo(GetCredentialInfo GetCredentialInfo) (*GetCredentialInfoResponse, *common.Fault)

	OptGetCredentialInfoList(GetCredentialInfoList GetCredentialInfoList) (*GetCredentialInfoListResponse, *common.Fault)

	OptGetCredentialList(GetCredentialList GetCredentialList) (*GetCredentialListResponse, *common.Fault)

	OptGetCredentialState(GetCredentialState GetCredentialState) (*GetCredentialStateResponse, *common.Fault)

	OptGetCredentials(GetCredentials GetCredentials) (*GetCredentialsResponse, *common.Fault)

	OptGetServiceCapabilities(GetServiceCapabilities GetServiceCapabilities) (*GetServiceCapabilitiesResponse, *common.Fault)

	OptGetSupportedFormatTypes(GetSupportedFormatTypes GetSupportedFormatTypes) (*GetSupportedFormatTypesResponse, *common.Fault)

	OptGetWhitelist(GetWhitelist GetWhitelist) (*GetWhitelistResponse, *common.Fault)

	OptModifyCredential(ModifyCredential ModifyCredential) (*ModifyCredentialResponse, *common.Fault)

	OptRemoveFromBlacklist(RemoveFromBlacklist RemoveFromBlacklist) (*RemoveFromBlacklistResponse, *common.Fault)

	OptRemoveFromWhitelist(RemoveFromWhitelist RemoveFromWhitelist) (*RemoveFromWhitelistResponse, *common.Fault)

	OptResetAntipassbackViolation(ResetAntipassbackViolation ResetAntipassbackViolation) (*ResetAntipassbackViolationResponse, *common.Fault)

	OptSetCredential(SetCredential SetCredential) (*SetCredentialResponse, *common.Fault)

	OptSetCredentialAccessProfiles(SetCredentialAccessProfiles SetCredentialAccessProfiles) (*SetCredentialAccessProfilesResponse, *common.Fault)

	OptSetCredentialIdentifier(SetCredentialIdentifier SetCredentialIdentifier) (*SetCredentialIdentifierResponse, *common.Fault)
}
type DateTime string

type Duration string

type AddToBlacklist struct {
	Identifier []CredentialIdentifierItem `xml:"Identifier,omitempty" json:"Identifier,omitempty" yaml:"Identifier,omitempty"`
}

type AddToBlacklistResponse struct {
}

type AddToWhitelist struct {
	Identifier []CredentialIdentifierItem `xml:"Identifier,omitempty" json:"Identifier,omitempty" yaml:"Identifier,omitempty"`
}

type AddToWhitelistResponse struct {
}

type AntipassbackState struct {
	AntipassbackViolated bool `xml:"AntipassbackViolated,omitempty" json:"AntipassbackViolated,omitempty" yaml:"AntipassbackViolated,omitempty"`
}

type CreateCredential struct {
	Credential Credential      `xml:"Credential,omitempty" json:"Credential,omitempty" yaml:"Credential,omitempty"`
	State      CredentialState `xml:"State,omitempty" json:"State,omitempty" yaml:"State,omitempty"`
}

type CreateCredentialResponse struct {
	Token *common.ReferenceToken `xml:"Token,omitempty" json:"Token,omitempty" yaml:"Token,omitempty"`
}

type Credential struct {
	Description               *common.Description       `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	CredentialHolderReference string                    `xml:"CredentialHolderReference,omitempty" json:"CredentialHolderReference,omitempty" yaml:"CredentialHolderReference,omitempty"`
	ValidFrom                 *common.DateTime          `xml:"ValidFrom,omitempty" json:"ValidFrom,omitempty" yaml:"ValidFrom,omitempty"`
	ValidTo                   *common.DateTime          `xml:"ValidTo,omitempty" json:"ValidTo,omitempty" yaml:"ValidTo,omitempty"`
	CredentialIdentifier      []CredentialIdentifier    `xml:"CredentialIdentifier" json:"CredentialIdentifier" yaml:"CredentialIdentifier"`
	CredentialAccessProfile   []CredentialAccessProfile `xml:"CredentialAccessProfile,omitempty" json:"CredentialAccessProfile,omitempty" yaml:"CredentialAccessProfile,omitempty"`
	ExtendedGrantTime         bool                      `xml:"ExtendedGrantTime,omitempty" json:"ExtendedGrantTime,omitempty" yaml:"ExtendedGrantTime,omitempty"`
	Attribute                 []*common.Attribute       `xml:"Attribute,omitempty" json:"Attribute,omitempty" yaml:"Attribute,omitempty"`
	Extension                 CredentialExtension       `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
	TypeAttrXSI               string                    `xml:"xsi:type,attr,omitempty"`
	TypeNamespace             string                    `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *Credential) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:Credential"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://www.onvif.org/ver10/credential/wsdl"
	}
}

type CredentialAccessProfile struct {
	AccessProfileToken *common.ReferenceToken `xml:"AccessProfileToken,omitempty" json:"AccessProfileToken,omitempty" yaml:"AccessProfileToken,omitempty"`
	ValidFrom          *common.DateTime       `xml:"ValidFrom,omitempty" json:"ValidFrom,omitempty" yaml:"ValidFrom,omitempty"`
	ValidTo            *common.DateTime       `xml:"ValidTo,omitempty" json:"ValidTo,omitempty" yaml:"ValidTo,omitempty"`
}

type CredentialData struct {
	Credential      Credential              `xml:"Credential,omitempty" json:"Credential,omitempty" yaml:"Credential,omitempty"`
	CredentialState CredentialState         `xml:"CredentialState,omitempty" json:"CredentialState,omitempty" yaml:"CredentialState,omitempty"`
	Extension       CredentialDataExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type CredentialDataExtension []interface{}

type CredentialExtension []interface{}

type CredentialIdentifier struct {
	Type                       CredentialIdentifierType `xml:"Type,omitempty" json:"Type,omitempty" yaml:"Type,omitempty"`
	ExemptedFromAuthentication bool                     `xml:"ExemptedFromAuthentication,omitempty" json:"ExemptedFromAuthentication,omitempty" yaml:"ExemptedFromAuthentication,omitempty"`
	Value                      []byte                   `xml:"Value,omitempty" json:"Value,omitempty" yaml:"Value,omitempty"`
}

type CredentialIdentifierFormatTypeInfo struct {
	FormatType  string                                      `xml:"FormatType,omitempty" json:"FormatType,omitempty" yaml:"FormatType,omitempty"`
	Description *common.Description                         `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	Extension   CredentialIdentifierFormatTypeInfoExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type CredentialIdentifierFormatTypeInfoExtension []interface{}

type CredentialIdentifierItem struct {
	Type  CredentialIdentifierType `xml:"Type,omitempty" json:"Type,omitempty" yaml:"Type,omitempty"`
	Value []byte                   `xml:"Value,omitempty" json:"Value,omitempty" yaml:"Value,omitempty"`
}

type CredentialIdentifierType struct {
	Name       *common.Name `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	FormatType string       `xml:"FormatType,omitempty" json:"FormatType,omitempty" yaml:"FormatType,omitempty"`
}

type CredentialInfo struct {
	Description               *common.Description `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	CredentialHolderReference string              `xml:"CredentialHolderReference,omitempty" json:"CredentialHolderReference,omitempty" yaml:"CredentialHolderReference,omitempty"`
	ValidFrom                 *common.DateTime    `xml:"ValidFrom,omitempty" json:"ValidFrom,omitempty" yaml:"ValidFrom,omitempty"`
	ValidTo                   *common.DateTime    `xml:"ValidTo,omitempty" json:"ValidTo,omitempty" yaml:"ValidTo,omitempty"`
	TypeAttrXSI               string              `xml:"xsi:type,attr,omitempty"`
	TypeNamespace             string              `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *CredentialInfo) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:CredentialInfo"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://www.onvif.org/ver10/credential/wsdl"
	}
}

type CredentialState struct {
	Enabled           bool                     `xml:"Enabled,omitempty" json:"Enabled,omitempty" yaml:"Enabled,omitempty"`
	Reason            *common.Name             `xml:"Reason,omitempty" json:"Reason,omitempty" yaml:"Reason,omitempty"`
	AntipassbackState AntipassbackState        `xml:"AntipassbackState,omitempty" json:"AntipassbackState,omitempty" yaml:"AntipassbackState,omitempty"`
	Extension         CredentialStateExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type CredentialStateExtension []interface{}

type DeleteBlacklist struct {
}

type DeleteBlacklistResponse struct {
}

type DeleteCredential struct {
	Token *common.ReferenceToken `xml:"Token,omitempty" json:"Token,omitempty" yaml:"Token,omitempty"`
}

type DeleteCredentialAccessProfiles struct {
	CredentialToken    *common.ReferenceToken  `xml:"CredentialToken,omitempty" json:"CredentialToken,omitempty" yaml:"CredentialToken,omitempty"`
	AccessProfileToken []common.ReferenceToken `xml:"AccessProfileToken" json:"AccessProfileToken" yaml:"AccessProfileToken"`
}

type DeleteCredentialAccessProfilesResponse struct {
}

type DeleteCredentialIdentifier struct {
	CredentialToken              *common.ReferenceToken `xml:"CredentialToken,omitempty" json:"CredentialToken,omitempty" yaml:"CredentialToken,omitempty"`
	CredentialIdentifierTypeName *common.Name           `xml:"CredentialIdentifierTypeName,omitempty" json:"CredentialIdentifierTypeName,omitempty" yaml:"CredentialIdentifierTypeName,omitempty"`
}

type DeleteCredentialIdentifierResponse struct {
}

type DeleteCredentialResponse struct {
}

type DeleteWhitelist struct {
}

type DeleteWhitelistResponse struct {
}

type DisableCredential struct {
	Token  *common.ReferenceToken `xml:"Token,omitempty" json:"Token,omitempty" yaml:"Token,omitempty"`
	Reason *common.Name           `xml:"Reason,omitempty" json:"Reason,omitempty" yaml:"Reason,omitempty"`
}

type DisableCredentialResponse struct {
}

type EnableCredential struct {
	Token  *common.ReferenceToken `xml:"Token,omitempty" json:"Token,omitempty" yaml:"Token,omitempty"`
	Reason *common.Name           `xml:"Reason,omitempty" json:"Reason,omitempty" yaml:"Reason,omitempty"`
}

type EnableCredentialResponse struct {
}

type FaultResponse struct {
	Token     *common.ReferenceToken `xml:"Token,omitempty" json:"Token,omitempty" yaml:"Token,omitempty"`
	Fault     string                 `xml:"Fault,omitempty" json:"Fault,omitempty" yaml:"Fault,omitempty"`
	Extension FaultResponseExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type FaultResponseExtension []interface{}

type GetBlacklist struct {
	Limit          int    `xml:"Limit,omitempty" json:"Limit,omitempty" yaml:"Limit,omitempty"`
	StartReference string `xml:"StartReference,omitempty" json:"StartReference,omitempty" yaml:"StartReference,omitempty"`
	IdentifierType string `xml:"IdentifierType,omitempty" json:"IdentifierType,omitempty" yaml:"IdentifierType,omitempty"`
	FormatType     string `xml:"FormatType,omitempty" json:"FormatType,omitempty" yaml:"FormatType,omitempty"`
	Value          []byte `xml:"Value,omitempty" json:"Value,omitempty" yaml:"Value,omitempty"`
}

type GetBlacklistResponse struct {
	NextStartReference string                     `xml:"NextStartReference,omitempty" json:"NextStartReference,omitempty" yaml:"NextStartReference,omitempty"`
	Identifier         []CredentialIdentifierItem `xml:"Identifier,omitempty" json:"Identifier,omitempty" yaml:"Identifier,omitempty"`
}

type GetCredentialAccessProfiles struct {
	CredentialToken *common.ReferenceToken `xml:"CredentialToken,omitempty" json:"CredentialToken,omitempty" yaml:"CredentialToken,omitempty"`
}

type GetCredentialAccessProfilesResponse struct {
	CredentialAccessProfile []CredentialAccessProfile `xml:"CredentialAccessProfile,omitempty" json:"CredentialAccessProfile,omitempty" yaml:"CredentialAccessProfile,omitempty"`
}

type GetCredentialIdentifiers struct {
	CredentialToken *common.ReferenceToken `xml:"CredentialToken,omitempty" json:"CredentialToken,omitempty" yaml:"CredentialToken,omitempty"`
}

type GetCredentialIdentifiersResponse struct {
	CredentialIdentifier []CredentialIdentifier `xml:"CredentialIdentifier,omitempty" json:"CredentialIdentifier,omitempty" yaml:"CredentialIdentifier,omitempty"`
}

type GetCredentialInfo struct {
	Token []common.ReferenceToken `xml:"Token" json:"Token" yaml:"Token"`
}

type GetCredentialInfoList struct {
	Limit          int    `xml:"Limit,omitempty" json:"Limit,omitempty" yaml:"Limit,omitempty"`
	StartReference string `xml:"StartReference,omitempty" json:"StartReference,omitempty" yaml:"StartReference,omitempty"`
}

type GetCredentialInfoListResponse struct {
	NextStartReference string           `xml:"NextStartReference,omitempty" json:"NextStartReference,omitempty" yaml:"NextStartReference,omitempty"`
	CredentialInfo     []CredentialInfo `xml:"CredentialInfo,omitempty" json:"CredentialInfo,omitempty" yaml:"CredentialInfo,omitempty"`
}

type GetCredentialInfoResponse struct {
	CredentialInfo []CredentialInfo `xml:"CredentialInfo,omitempty" json:"CredentialInfo,omitempty" yaml:"CredentialInfo,omitempty"`
}

type GetCredentialList struct {
	Limit          int    `xml:"Limit,omitempty" json:"Limit,omitempty" yaml:"Limit,omitempty"`
	StartReference string `xml:"StartReference,omitempty" json:"StartReference,omitempty" yaml:"StartReference,omitempty"`
}

type GetCredentialListResponse struct {
	NextStartReference string       `xml:"NextStartReference,omitempty" json:"NextStartReference,omitempty" yaml:"NextStartReference,omitempty"`
	Credential         []Credential `xml:"Credential,omitempty" json:"Credential,omitempty" yaml:"Credential,omitempty"`
}

type GetCredentialState struct {
	Token *common.ReferenceToken `xml:"Token,omitempty" json:"Token,omitempty" yaml:"Token,omitempty"`
}

type GetCredentialStateResponse struct {
	State CredentialState `xml:"State,omitempty" json:"State,omitempty" yaml:"State,omitempty"`
}

type GetCredentials struct {
	Token []common.ReferenceToken `xml:"Token" json:"Token" yaml:"Token"`
}

type GetCredentialsResponse struct {
	Credential []Credential `xml:"Credential,omitempty" json:"Credential,omitempty" yaml:"Credential,omitempty"`
}

type GetServiceCapabilities struct {
}

type GetServiceCapabilitiesResponse struct {
	Capabilities ServiceCapabilities `xml:"Capabilities,omitempty" json:"Capabilities,omitempty" yaml:"Capabilities,omitempty"`
}

type GetSupportedFormatTypes struct {
	CredentialIdentifierTypeName string `xml:"CredentialIdentifierTypeName,omitempty" json:"CredentialIdentifierTypeName,omitempty" yaml:"CredentialIdentifierTypeName,omitempty"`
}

type GetSupportedFormatTypesResponse struct {
	FormatTypeInfo []CredentialIdentifierFormatTypeInfo `xml:"FormatTypeInfo" json:"FormatTypeInfo" yaml:"FormatTypeInfo"`
}

type GetWhitelist struct {
	Limit          int    `xml:"Limit,omitempty" json:"Limit,omitempty" yaml:"Limit,omitempty"`
	StartReference string `xml:"StartReference,omitempty" json:"StartReference,omitempty" yaml:"StartReference,omitempty"`
	IdentifierType string `xml:"IdentifierType,omitempty" json:"IdentifierType,omitempty" yaml:"IdentifierType,omitempty"`
	FormatType     string `xml:"FormatType,omitempty" json:"FormatType,omitempty" yaml:"FormatType,omitempty"`
	Value          []byte `xml:"Value,omitempty" json:"Value,omitempty" yaml:"Value,omitempty"`
}

type GetWhitelistResponse struct {
	NextStartReference string                     `xml:"NextStartReference,omitempty" json:"NextStartReference,omitempty" yaml:"NextStartReference,omitempty"`
	Identifier         []CredentialIdentifierItem `xml:"Identifier,omitempty" json:"Identifier,omitempty" yaml:"Identifier,omitempty"`
}

type ModifyCredential struct {
	Credential Credential `xml:"Credential,omitempty" json:"Credential,omitempty" yaml:"Credential,omitempty"`
}

type ModifyCredentialResponse struct {
}

type RemoveFromBlacklist struct {
	Identifier []CredentialIdentifierItem `xml:"Identifier,omitempty" json:"Identifier,omitempty" yaml:"Identifier,omitempty"`
}

type RemoveFromBlacklistResponse struct {
}

type RemoveFromWhitelist struct {
	Identifier []CredentialIdentifierItem `xml:"Identifier,omitempty" json:"Identifier,omitempty" yaml:"Identifier,omitempty"`
}

type RemoveFromWhitelistResponse struct {
}

type ResetAntipassbackViolation struct {
	CredentialToken *common.ReferenceToken `xml:"CredentialToken,omitempty" json:"CredentialToken,omitempty" yaml:"CredentialToken,omitempty"`
}

type ResetAntipassbackViolationResponse struct {
}

type ServiceCapabilities struct {
	SupportedIdentifierType                  []common.Name                `xml:"SupportedIdentifierType" json:"SupportedIdentifierType" yaml:"SupportedIdentifierType"`
	Extension                                ServiceCapabilitiesExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
	MaxLimit                                 uint64                       `xml:"MaxLimit,attr,omitempty" json:"MaxLimit,attr,omitempty" yaml:"MaxLimit,attr,omitempty"`
	CredentialValiditySupported              bool                         `xml:"CredentialValiditySupported,attr,omitempty" json:"CredentialValiditySupported,attr,omitempty" yaml:"CredentialValiditySupported,attr,omitempty"`
	CredentialAccessProfileValiditySupported bool                         `xml:"CredentialAccessProfileValiditySupported,attr,omitempty" json:"CredentialAccessProfileValiditySupported,attr,omitempty" yaml:"CredentialAccessProfileValiditySupported,attr,omitempty"`
	ValiditySupportsTimeValue                bool                         `xml:"ValiditySupportsTimeValue,attr,omitempty" json:"ValiditySupportsTimeValue,attr,omitempty" yaml:"ValiditySupportsTimeValue,attr,omitempty"`
	MaxCredentials                           uint                         `xml:"MaxCredentials,attr,omitempty" json:"MaxCredentials,attr,omitempty" yaml:"MaxCredentials,attr,omitempty"`
	MaxAccessProfilesPerCredential           uint64                       `xml:"MaxAccessProfilesPerCredential,attr,omitempty" json:"MaxAccessProfilesPerCredential,attr,omitempty" yaml:"MaxAccessProfilesPerCredential,attr,omitempty"`
	ResetAntipassbackSupported               bool                         `xml:"ResetAntipassbackSupported,attr,omitempty" json:"ResetAntipassbackSupported,attr,omitempty" yaml:"ResetAntipassbackSupported,attr,omitempty"`
	ClientSuppliedTokenSupported             bool                         `xml:"ClientSuppliedTokenSupported,attr,omitempty" json:"ClientSuppliedTokenSupported,attr,omitempty" yaml:"ClientSuppliedTokenSupported,attr,omitempty"`
	DefaultCredentialSuspensionDuration      common.Duration              `xml:"DefaultCredentialSuspensionDuration,attr,omitempty" json:"DefaultCredentialSuspensionDuration,attr,omitempty" yaml:"DefaultCredentialSuspensionDuration,attr,omitempty"`
	MaxWhitelistedItems                      uint                         `xml:"MaxWhitelistedItems,attr,omitempty" json:"MaxWhitelistedItems,attr,omitempty" yaml:"MaxWhitelistedItems,attr,omitempty"`
	MaxBlacklistedItems                      uint                         `xml:"MaxBlacklistedItems,attr,omitempty" json:"MaxBlacklistedItems,attr,omitempty" yaml:"MaxBlacklistedItems,attr,omitempty"`
}

type ServiceCapabilitiesExtension struct {
	SupportedExemptionType []*common.Name `xml:"SupportedExemptionType,omitempty" json:"SupportedExemptionType,omitempty" yaml:"SupportedExemptionType,omitempty"`
}

type SetCredential struct {
	CredentialData CredentialData `xml:"CredentialData,omitempty" json:"CredentialData,omitempty" yaml:"CredentialData,omitempty"`
}

type SetCredentialAccessProfiles struct {
	CredentialToken         *common.ReferenceToken    `xml:"CredentialToken,omitempty" json:"CredentialToken,omitempty" yaml:"CredentialToken,omitempty"`
	CredentialAccessProfile []CredentialAccessProfile `xml:"CredentialAccessProfile" json:"CredentialAccessProfile" yaml:"CredentialAccessProfile"`
}

type SetCredentialAccessProfilesResponse struct {
}

type SetCredentialIdentifier struct {
	CredentialToken      *common.ReferenceToken `xml:"CredentialToken,omitempty" json:"CredentialToken,omitempty" yaml:"CredentialToken,omitempty"`
	CredentialIdentifier CredentialIdentifier   `xml:"CredentialIdentifier,omitempty" json:"CredentialIdentifier,omitempty" yaml:"CredentialIdentifier,omitempty"`
}

type SetCredentialIdentifierResponse struct {
}

type SetCredentialResponse struct {
}

// credentialPort implements the CredentialPort interface.
type credentialPort struct {
	cli      common.Client
	Endpoint string
}

func (p *credentialPort) OptAddToBlacklist(args AddToBlacklist) (*AddToBlacklistResponse, *common.Fault) {
	req := struct {
		XMLName        string `xml:"tcr:AddToBlacklist"`
		AddToBlacklist AddToBlacklist
	}{
		AddToBlacklist: args,
	}

	resp := AddToBlacklistResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *credentialPort) OptAddToWhitelist(args AddToWhitelist) (*AddToWhitelistResponse, *common.Fault) {
	req := struct {
		XMLName        string `xml:"tcr:AddToWhitelist"`
		AddToWhitelist AddToWhitelist
	}{
		AddToWhitelist: args,
	}

	resp := AddToWhitelistResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *credentialPort) OptCreateCredential(args CreateCredential) (*CreateCredentialResponse, *common.Fault) {
	req := struct {
		XMLName          string `xml:"tcr:CreateCredential"`
		CreateCredential CreateCredential
	}{
		CreateCredential: args,
	}

	resp := CreateCredentialResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *credentialPort) OptDeleteBlacklist(args DeleteBlacklist) (*DeleteBlacklistResponse, *common.Fault) {
	req := struct {
		XMLName         string `xml:"tcr:DeleteBlacklist"`
		DeleteBlacklist DeleteBlacklist
	}{
		DeleteBlacklist: args,
	}

	resp := DeleteBlacklistResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *credentialPort) OptDeleteCredential(args DeleteCredential) (*DeleteCredentialResponse, *common.Fault) {
	req := struct {
		XMLName          string `xml:"tcr:DeleteCredential"`
		DeleteCredential DeleteCredential
	}{
		DeleteCredential: args,
	}

	resp := DeleteCredentialResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *credentialPort) OptDeleteCredentialAccessProfiles(args DeleteCredentialAccessProfiles) (*DeleteCredentialAccessProfilesResponse, *common.Fault) {
	req := struct {
		XMLName                        string `xml:"tcr:DeleteCredentialAccessProfiles"`
		DeleteCredentialAccessProfiles DeleteCredentialAccessProfiles
	}{
		DeleteCredentialAccessProfiles: args,
	}

	resp := DeleteCredentialAccessProfilesResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *credentialPort) OptDeleteCredentialIdentifier(args DeleteCredentialIdentifier) (*DeleteCredentialIdentifierResponse, *common.Fault) {
	req := struct {
		XMLName                    string `xml:"tcr:DeleteCredentialIdentifier"`
		DeleteCredentialIdentifier DeleteCredentialIdentifier
	}{
		DeleteCredentialIdentifier: args,
	}

	resp := DeleteCredentialIdentifierResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *credentialPort) OptDeleteWhitelist(args DeleteWhitelist) (*DeleteWhitelistResponse, *common.Fault) {
	req := struct {
		XMLName         string `xml:"tcr:DeleteWhitelist"`
		DeleteWhitelist DeleteWhitelist
	}{
		DeleteWhitelist: args,
	}

	resp := DeleteWhitelistResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *credentialPort) OptDisableCredential(args DisableCredential) (*DisableCredentialResponse, *common.Fault) {
	req := struct {
		XMLName           string `xml:"tcr:DisableCredential"`
		DisableCredential DisableCredential
	}{
		DisableCredential: args,
	}

	resp := DisableCredentialResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *credentialPort) OptEnableCredential(args EnableCredential) (*EnableCredentialResponse, *common.Fault) {
	req := struct {
		XMLName          string `xml:"tcr:EnableCredential"`
		EnableCredential EnableCredential
	}{
		EnableCredential: args,
	}

	resp := EnableCredentialResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *credentialPort) OptGetBlacklist(args GetBlacklist) (*GetBlacklistResponse, *common.Fault) {
	req := struct {
		XMLName      string `xml:"tcr:GetBlacklist"`
		GetBlacklist GetBlacklist
	}{
		GetBlacklist: args,
	}

	resp := GetBlacklistResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *credentialPort) OptGetCredentialAccessProfiles(args GetCredentialAccessProfiles) (*GetCredentialAccessProfilesResponse, *common.Fault) {
	req := struct {
		XMLName                     string `xml:"tcr:GetCredentialAccessProfiles"`
		GetCredentialAccessProfiles GetCredentialAccessProfiles
	}{
		GetCredentialAccessProfiles: args,
	}

	resp := GetCredentialAccessProfilesResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *credentialPort) OptGetCredentialIdentifiers(args GetCredentialIdentifiers) (*GetCredentialIdentifiersResponse, *common.Fault) {
	req := struct {
		XMLName                  string `xml:"tcr:GetCredentialIdentifiers"`
		GetCredentialIdentifiers GetCredentialIdentifiers
	}{
		GetCredentialIdentifiers: args,
	}

	resp := GetCredentialIdentifiersResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *credentialPort) OptGetCredentialInfo(args GetCredentialInfo) (*GetCredentialInfoResponse, *common.Fault) {
	req := struct {
		XMLName           string `xml:"tcr:GetCredentialInfo"`
		GetCredentialInfo GetCredentialInfo
	}{
		GetCredentialInfo: args,
	}

	resp := GetCredentialInfoResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *credentialPort) OptGetCredentialInfoList(args GetCredentialInfoList) (*GetCredentialInfoListResponse, *common.Fault) {
	req := struct {
		XMLName               string `xml:"tcr:GetCredentialInfoList"`
		GetCredentialInfoList GetCredentialInfoList
	}{
		GetCredentialInfoList: args,
	}

	resp := GetCredentialInfoListResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *credentialPort) OptGetCredentialList(args GetCredentialList) (*GetCredentialListResponse, *common.Fault) {
	req := struct {
		XMLName           string `xml:"tcr:GetCredentialList"`
		GetCredentialList GetCredentialList
	}{
		GetCredentialList: args,
	}

	resp := GetCredentialListResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *credentialPort) OptGetCredentialState(args GetCredentialState) (*GetCredentialStateResponse, *common.Fault) {
	req := struct {
		XMLName            string `xml:"tcr:GetCredentialState"`
		GetCredentialState GetCredentialState
	}{
		GetCredentialState: args,
	}

	resp := GetCredentialStateResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *credentialPort) OptGetCredentials(args GetCredentials) (*GetCredentialsResponse, *common.Fault) {
	req := struct {
		XMLName        string `xml:"tcr:GetCredentials"`
		GetCredentials GetCredentials
	}{
		GetCredentials: args,
	}

	resp := GetCredentialsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *credentialPort) OptGetServiceCapabilities(args GetServiceCapabilities) (*GetServiceCapabilitiesResponse, *common.Fault) {
	req := struct {
		XMLName                string `xml:"tcr:GetServiceCapabilities"`
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

func (p *credentialPort) OptGetSupportedFormatTypes(args GetSupportedFormatTypes) (*GetSupportedFormatTypesResponse, *common.Fault) {
	req := struct {
		XMLName                 string `xml:"tcr:GetSupportedFormatTypes"`
		GetSupportedFormatTypes GetSupportedFormatTypes
	}{
		GetSupportedFormatTypes: args,
	}

	resp := GetSupportedFormatTypesResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *credentialPort) OptGetWhitelist(args GetWhitelist) (*GetWhitelistResponse, *common.Fault) {
	req := struct {
		XMLName      string `xml:"tcr:GetWhitelist"`
		GetWhitelist GetWhitelist
	}{
		GetWhitelist: args,
	}

	resp := GetWhitelistResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *credentialPort) OptModifyCredential(args ModifyCredential) (*ModifyCredentialResponse, *common.Fault) {
	req := struct {
		XMLName          string `xml:"tcr:ModifyCredential"`
		ModifyCredential ModifyCredential
	}{
		ModifyCredential: args,
	}

	resp := ModifyCredentialResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *credentialPort) OptRemoveFromBlacklist(args RemoveFromBlacklist) (*RemoveFromBlacklistResponse, *common.Fault) {
	req := struct {
		XMLName             string `xml:"tcr:RemoveFromBlacklist"`
		RemoveFromBlacklist RemoveFromBlacklist
	}{
		RemoveFromBlacklist: args,
	}

	resp := RemoveFromBlacklistResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *credentialPort) OptRemoveFromWhitelist(args RemoveFromWhitelist) (*RemoveFromWhitelistResponse, *common.Fault) {
	req := struct {
		XMLName             string `xml:"tcr:RemoveFromWhitelist"`
		RemoveFromWhitelist RemoveFromWhitelist
	}{
		RemoveFromWhitelist: args,
	}

	resp := RemoveFromWhitelistResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *credentialPort) OptResetAntipassbackViolation(args ResetAntipassbackViolation) (*ResetAntipassbackViolationResponse, *common.Fault) {
	req := struct {
		XMLName                    string `xml:"tcr:ResetAntipassbackViolation"`
		ResetAntipassbackViolation ResetAntipassbackViolation
	}{
		ResetAntipassbackViolation: args,
	}

	resp := ResetAntipassbackViolationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *credentialPort) OptSetCredential(args SetCredential) (*SetCredentialResponse, *common.Fault) {
	req := struct {
		XMLName       string `xml:"tcr:SetCredential"`
		SetCredential SetCredential
	}{
		SetCredential: args,
	}

	resp := SetCredentialResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *credentialPort) OptSetCredentialAccessProfiles(args SetCredentialAccessProfiles) (*SetCredentialAccessProfilesResponse, *common.Fault) {
	req := struct {
		XMLName                     string `xml:"tcr:SetCredentialAccessProfiles"`
		SetCredentialAccessProfiles SetCredentialAccessProfiles
	}{
		SetCredentialAccessProfiles: args,
	}

	resp := SetCredentialAccessProfilesResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *credentialPort) OptSetCredentialIdentifier(args SetCredentialIdentifier) (*SetCredentialIdentifierResponse, *common.Fault) {
	req := struct {
		XMLName                 string `xml:"tcr:SetCredentialIdentifier"`
		SetCredentialIdentifier SetCredentialIdentifier
	}{
		SetCredentialIdentifier: args,
	}

	resp := SetCredentialIdentifierResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}
