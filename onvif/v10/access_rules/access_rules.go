package access_rules

import (
	"code.byted.org/videoarch/go-onvif/onvif/common"
)

var Namespace = "http://www.onvif.org/ver10/accessrules/wsdl"

// NewAccessRulesPort creates an initializes a AccessRulesPort.
func NewAccessRulesPort(endpoint string, cli common.Client) AccessRulesPort {
	return &accessRulesPort{cli: cli, Endpoint: endpoint}
}

// AccessRulesPort was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type AccessRulesPort interface {
	OptCreateAccessProfile(CreateAccessProfile CreateAccessProfile) (*CreateAccessProfileResponse, *common.Fault)

	OptDeleteAccessProfile(DeleteAccessProfile DeleteAccessProfile) (*DeleteAccessProfileResponse, *common.Fault)

	OptGetAccessProfileInfo(GetAccessProfileInfo GetAccessProfileInfo) (*GetAccessProfileInfoResponse, *common.Fault)

	OptGetAccessProfileInfoList(GetAccessProfileInfoList GetAccessProfileInfoList) (*GetAccessProfileInfoListResponse, *common.Fault)

	OptGetAccessProfileList(GetAccessProfileList GetAccessProfileList) (*GetAccessProfileListResponse, *common.Fault)

	OptGetAccessProfiles(GetAccessProfiles GetAccessProfiles) (*GetAccessProfilesResponse, *common.Fault)

	OptGetServiceCapabilities(GetServiceCapabilities GetServiceCapabilities) (*GetServiceCapabilitiesResponse, *common.Fault)

	OptModifyAccessProfile(ModifyAccessProfile ModifyAccessProfile) (*ModifyAccessProfileResponse, *common.Fault)

	OptSetAccessProfile(SetAccessProfile SetAccessProfile) (*SetAccessProfileResponse, *common.Fault)
}
type AccessPolicy struct {
	ScheduleToken *common.ReferenceToken `xml:"ScheduleToken,omitempty" json:"ScheduleToken,omitempty" yaml:"ScheduleToken,omitempty"`
	Entity        *common.ReferenceToken `xml:"Entity,omitempty" json:"Entity,omitempty" yaml:"Entity,omitempty"`
	EntityType    string                 `xml:"EntityType,omitempty" json:"EntityType,omitempty" yaml:"EntityType,omitempty"`
	Extension     AccessPolicyExtension  `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type AccessPolicyExtension []interface{}

type AccessProfile struct {
	Name          *common.Name           `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Description   *common.Description    `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	AccessPolicy  []AccessPolicy         `xml:"AccessPolicy,omitempty" json:"AccessPolicy,omitempty" yaml:"AccessPolicy,omitempty"`
	Extension     AccessProfileExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
	TypeAttrXSI   string                 `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string                 `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *AccessProfile) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:AccessProfile"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://www.onvif.org/ver10/accessrules/wsdl"
	}
}

type AccessProfileExtension []interface{}

type AccessProfileInfo struct {
	Name          *common.Name        `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Description   *common.Description `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	TypeAttrXSI   string              `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string              `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *AccessProfileInfo) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:AccessProfileInfo"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://www.onvif.org/ver10/accessrules/wsdl"
	}
}

type CreateAccessProfile struct {
	AccessProfile AccessProfile `xml:"AccessProfile,omitempty" json:"AccessProfile,omitempty" yaml:"AccessProfile,omitempty"`
}

type CreateAccessProfileResponse struct {
	Token *common.ReferenceToken `xml:"Token,omitempty" json:"Token,omitempty" yaml:"Token,omitempty"`
}

type DeleteAccessProfile struct {
	Token *common.ReferenceToken `xml:"Token,omitempty" json:"Token,omitempty" yaml:"Token,omitempty"`
}

type DeleteAccessProfileResponse struct {
}

type GetAccessProfileInfo struct {
	Token []common.ReferenceToken `xml:"Token" json:"Token" yaml:"Token"`
}

type GetAccessProfileInfoList struct {
	Limit          int    `xml:"Limit,omitempty" json:"Limit,omitempty" yaml:"Limit,omitempty"`
	StartReference string `xml:"StartReference,omitempty" json:"StartReference,omitempty" yaml:"StartReference,omitempty"`
}

type GetAccessProfileInfoListResponse struct {
	NextStartReference string              `xml:"NextStartReference,omitempty" json:"NextStartReference,omitempty" yaml:"NextStartReference,omitempty"`
	AccessProfileInfo  []AccessProfileInfo `xml:"AccessProfileInfo,omitempty" json:"AccessProfileInfo,omitempty" yaml:"AccessProfileInfo,omitempty"`
}

type GetAccessProfileInfoResponse struct {
	AccessProfileInfo []AccessProfileInfo `xml:"AccessProfileInfo,omitempty" json:"AccessProfileInfo,omitempty" yaml:"AccessProfileInfo,omitempty"`
}

type GetAccessProfileList struct {
	Limit          int    `xml:"Limit,omitempty" json:"Limit,omitempty" yaml:"Limit,omitempty"`
	StartReference string `xml:"StartReference,omitempty" json:"StartReference,omitempty" yaml:"StartReference,omitempty"`
}

type GetAccessProfileListResponse struct {
	NextStartReference string          `xml:"NextStartReference,omitempty" json:"NextStartReference,omitempty" yaml:"NextStartReference,omitempty"`
	AccessProfile      []AccessProfile `xml:"AccessProfile,omitempty" json:"AccessProfile,omitempty" yaml:"AccessProfile,omitempty"`
}

type GetAccessProfiles struct {
	Token []common.ReferenceToken `xml:"Token" json:"Token" yaml:"Token"`
}

type GetAccessProfilesResponse struct {
	AccessProfile []AccessProfile `xml:"AccessProfile,omitempty" json:"AccessProfile,omitempty" yaml:"AccessProfile,omitempty"`
}

type GetServiceCapabilities struct {
}

type GetServiceCapabilitiesResponse struct {
	Capabilities ServiceCapabilities `xml:"Capabilities,omitempty" json:"Capabilities,omitempty" yaml:"Capabilities,omitempty"`
}

type ModifyAccessProfile struct {
	AccessProfile AccessProfile `xml:"AccessProfile,omitempty" json:"AccessProfile,omitempty" yaml:"AccessProfile,omitempty"`
}

type ModifyAccessProfileResponse struct {
}

type ServiceCapabilities []interface{}

type SetAccessProfile struct {
	AccessProfile AccessProfile `xml:"AccessProfile,omitempty" json:"AccessProfile,omitempty" yaml:"AccessProfile,omitempty"`
}

type SetAccessProfileResponse struct {
}

// accessRulesPort implements the AccessRulesPort interface.
type accessRulesPort struct {
	cli      common.Client
	Endpoint string
}

func (p *accessRulesPort) OptCreateAccessProfile(args CreateAccessProfile) (*CreateAccessProfileResponse, *common.Fault) {
	req := struct {
		XMLName             string `xml:"tar:CreateAccessProfile"`
		CreateAccessProfile CreateAccessProfile
	}{
		CreateAccessProfile: args,
	}

	resp := CreateAccessProfileResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *accessRulesPort) OptDeleteAccessProfile(args DeleteAccessProfile) (*DeleteAccessProfileResponse, *common.Fault) {
	req := struct {
		XMLName             string `xml:"tar:DeleteAccessProfile"`
		DeleteAccessProfile DeleteAccessProfile
	}{
		DeleteAccessProfile: args,
	}

	resp := DeleteAccessProfileResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *accessRulesPort) OptGetAccessProfileInfo(args GetAccessProfileInfo) (*GetAccessProfileInfoResponse, *common.Fault) {
	req := struct {
		XMLName              string `xml:"tar:GetAccessProfileInfo"`
		GetAccessProfileInfo GetAccessProfileInfo
	}{
		GetAccessProfileInfo: args,
	}

	resp := GetAccessProfileInfoResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *accessRulesPort) OptGetAccessProfileInfoList(args GetAccessProfileInfoList) (*GetAccessProfileInfoListResponse, *common.Fault) {
	req := struct {
		XMLName                  string `xml:"tar:GetAccessProfileInfoList"`
		GetAccessProfileInfoList GetAccessProfileInfoList
	}{
		GetAccessProfileInfoList: args,
	}

	resp := GetAccessProfileInfoListResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *accessRulesPort) OptGetAccessProfileList(args GetAccessProfileList) (*GetAccessProfileListResponse, *common.Fault) {
	req := struct {
		XMLName              string `xml:"tar:GetAccessProfileList"`
		GetAccessProfileList GetAccessProfileList
	}{
		GetAccessProfileList: args,
	}

	resp := GetAccessProfileListResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *accessRulesPort) OptGetAccessProfiles(args GetAccessProfiles) (*GetAccessProfilesResponse, *common.Fault) {
	req := struct {
		XMLName           string `xml:"tar:GetAccessProfiles"`
		GetAccessProfiles GetAccessProfiles
	}{
		GetAccessProfiles: args,
	}

	resp := GetAccessProfilesResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *accessRulesPort) OptGetServiceCapabilities(args GetServiceCapabilities) (*GetServiceCapabilitiesResponse, *common.Fault) {
	req := struct {
		XMLName                string `xml:"tar:GetServiceCapabilities"`
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

func (p *accessRulesPort) OptModifyAccessProfile(args ModifyAccessProfile) (*ModifyAccessProfileResponse, *common.Fault) {
	req := struct {
		XMLName             string `xml:"tar:ModifyAccessProfile"`
		ModifyAccessProfile ModifyAccessProfile
	}{
		ModifyAccessProfile: args,
	}

	resp := ModifyAccessProfileResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *accessRulesPort) OptSetAccessProfile(args SetAccessProfile) (*SetAccessProfileResponse, *common.Fault) {
	req := struct {
		XMLName          string `xml:"tar:SetAccessProfile"`
		SetAccessProfile SetAccessProfile
	}{
		SetAccessProfile: args,
	}

	resp := SetAccessProfileResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}
