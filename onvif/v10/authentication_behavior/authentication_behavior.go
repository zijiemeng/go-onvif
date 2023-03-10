package authentication_behavior

import (
	"github.com/zijiemeng/go-onvif/onvif/common"
)

var Namespace = "http://www.onvif.org/ver10/authenticationbehavior/wsdl"

// NewAuthenticationBehaviorPort creates an initializes a AuthenticationBehaviorPort.
func NewAuthenticationBehaviorPort(endpoint string, cli common.Client) AuthenticationBehaviorPort {
	return &authenticationBehaviorPort{cli: cli, Endpoint: endpoint}
}

// AuthenticationBehaviorPort was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type AuthenticationBehaviorPort interface {
	OptCreateAuthenticationProfile(CreateAuthenticationProfile CreateAuthenticationProfile) (*CreateAuthenticationProfileResponse, *common.Fault)

	OptCreateSecurityLevel(CreateSecurityLevel CreateSecurityLevel) (*CreateSecurityLevelResponse, *common.Fault)

	OptDeleteAuthenticationProfile(DeleteAuthenticationProfile DeleteAuthenticationProfile) (*DeleteAuthenticationProfileResponse, *common.Fault)

	OptDeleteSecurityLevel(DeleteSecurityLevel DeleteSecurityLevel) (*DeleteSecurityLevelResponse, *common.Fault)

	OptGetAuthenticationProfileInfo(GetAuthenticationProfileInfo GetAuthenticationProfileInfo) (*GetAuthenticationProfileInfoResponse, *common.Fault)

	OptGetAuthenticationProfileInfoList(GetAuthenticationProfileInfoList GetAuthenticationProfileInfoList) (*GetAuthenticationProfileInfoListResponse, *common.Fault)

	OptGetAuthenticationProfileList(GetAuthenticationProfileList GetAuthenticationProfileList) (*GetAuthenticationProfileListResponse, *common.Fault)

	OptGetAuthenticationProfiles(GetAuthenticationProfiles GetAuthenticationProfiles) (*GetAuthenticationProfilesResponse, *common.Fault)

	OptGetSecurityLevelInfo(GetSecurityLevelInfo GetSecurityLevelInfo) (*GetSecurityLevelInfoResponse, *common.Fault)

	OptGetSecurityLevelInfoList(GetSecurityLevelInfoList GetSecurityLevelInfoList) (*GetSecurityLevelInfoListResponse, *common.Fault)

	OptGetSecurityLevelList(GetSecurityLevelList GetSecurityLevelList) (*GetSecurityLevelListResponse, *common.Fault)

	OptGetSecurityLevels(GetSecurityLevels GetSecurityLevels) (*GetSecurityLevelsResponse, *common.Fault)

	OptGetServiceCapabilities(GetServiceCapabilities GetServiceCapabilities) (*GetServiceCapabilitiesResponse, *common.Fault)

	OptModifyAuthenticationProfile(ModifyAuthenticationProfile ModifyAuthenticationProfile) (*ModifyAuthenticationProfileResponse, *common.Fault)

	OptModifySecurityLevel(ModifySecurityLevel ModifySecurityLevel) (*ModifySecurityLevelResponse, *common.Fault)

	OptSetAuthenticationProfile(SetAuthenticationProfile SetAuthenticationProfile) (*SetAuthenticationProfileResponse, *common.Fault)

	OptSetSecurityLevel(SetSecurityLevel SetSecurityLevel) (*SetSecurityLevelResponse, *common.Fault)
}
type AuthenticationPolicy struct {
	ScheduleToken           *common.ReferenceToken        `xml:"ScheduleToken,omitempty" json:"ScheduleToken,omitempty" yaml:"ScheduleToken,omitempty"`
	SecurityLevelConstraint []SecurityLevelConstraint     `xml:"SecurityLevelConstraint" json:"SecurityLevelConstraint" yaml:"SecurityLevelConstraint"`
	Extension               AuthenticationPolicyExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type AuthenticationPolicyExtension []interface{}

type AuthenticationProfile struct {
	Name                      *common.Name                   `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Description               *common.Description            `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	DefaultSecurityLevelToken *common.ReferenceToken         `xml:"DefaultSecurityLevelToken,omitempty" json:"DefaultSecurityLevelToken,omitempty" yaml:"DefaultSecurityLevelToken,omitempty"`
	AuthenticationPolicy      []AuthenticationPolicy         `xml:"AuthenticationPolicy,omitempty" json:"AuthenticationPolicy,omitempty" yaml:"AuthenticationPolicy,omitempty"`
	Extension                 AuthenticationProfileExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
	TypeAttrXSI               string                         `xml:"xsi:type,attr,omitempty"`
	TypeNamespace             string                         `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *AuthenticationProfile) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:AuthenticationProfile"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://www.onvif.org/ver10/authenticationbehavior/wsdl"
	}
}

type AuthenticationProfileExtension []interface{}

type AuthenticationProfileInfo struct {
	Name          *common.Name        `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Description   *common.Description `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	TypeAttrXSI   string              `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string              `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *AuthenticationProfileInfo) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:AuthenticationProfileInfo"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://www.onvif.org/ver10/authenticationbehavior/wsdl"
	}
}

type CreateAuthenticationProfile struct {
	AuthenticationProfile AuthenticationProfile `xml:"AuthenticationProfile,omitempty" json:"AuthenticationProfile,omitempty" yaml:"AuthenticationProfile,omitempty"`
}

type CreateAuthenticationProfileResponse struct {
	Token *common.ReferenceToken `xml:"Token,omitempty" json:"Token,omitempty" yaml:"Token,omitempty"`
}

type CreateSecurityLevel struct {
	SecurityLevel SecurityLevel `xml:"SecurityLevel,omitempty" json:"SecurityLevel,omitempty" yaml:"SecurityLevel,omitempty"`
}

type CreateSecurityLevelResponse struct {
	Token *common.ReferenceToken `xml:"Token,omitempty" json:"Token,omitempty" yaml:"Token,omitempty"`
}

type DeleteAuthenticationProfile struct {
	Token *common.ReferenceToken `xml:"Token,omitempty" json:"Token,omitempty" yaml:"Token,omitempty"`
}

type DeleteAuthenticationProfileResponse struct {
}

type DeleteSecurityLevel struct {
	Token *common.ReferenceToken `xml:"Token,omitempty" json:"Token,omitempty" yaml:"Token,omitempty"`
}

type DeleteSecurityLevelResponse struct {
}

type GetAuthenticationProfileInfo struct {
	Token []common.ReferenceToken `xml:"Token" json:"Token" yaml:"Token"`
}

type GetAuthenticationProfileInfoList struct {
	Limit          int    `xml:"Limit,omitempty" json:"Limit,omitempty" yaml:"Limit,omitempty"`
	StartReference string `xml:"StartReference,omitempty" json:"StartReference,omitempty" yaml:"StartReference,omitempty"`
}

type GetAuthenticationProfileInfoListResponse struct {
	NextStartReference        string                      `xml:"NextStartReference,omitempty" json:"NextStartReference,omitempty" yaml:"NextStartReference,omitempty"`
	AuthenticationProfileInfo []AuthenticationProfileInfo `xml:"AuthenticationProfileInfo,omitempty" json:"AuthenticationProfileInfo,omitempty" yaml:"AuthenticationProfileInfo,omitempty"`
}

type GetAuthenticationProfileInfoResponse struct {
	AuthenticationProfileInfo []AuthenticationProfileInfo `xml:"AuthenticationProfileInfo,omitempty" json:"AuthenticationProfileInfo,omitempty" yaml:"AuthenticationProfileInfo,omitempty"`
}

type GetAuthenticationProfileList struct {
	Limit          int    `xml:"Limit,omitempty" json:"Limit,omitempty" yaml:"Limit,omitempty"`
	StartReference string `xml:"StartReference,omitempty" json:"StartReference,omitempty" yaml:"StartReference,omitempty"`
}

type GetAuthenticationProfileListResponse struct {
	NextStartReference    string                  `xml:"NextStartReference,omitempty" json:"NextStartReference,omitempty" yaml:"NextStartReference,omitempty"`
	AuthenticationProfile []AuthenticationProfile `xml:"AuthenticationProfile,omitempty" json:"AuthenticationProfile,omitempty" yaml:"AuthenticationProfile,omitempty"`
}

type GetAuthenticationProfiles struct {
	Token []common.ReferenceToken `xml:"Token" json:"Token" yaml:"Token"`
}

type GetAuthenticationProfilesResponse struct {
	AuthenticationProfile []AuthenticationProfile `xml:"AuthenticationProfile,omitempty" json:"AuthenticationProfile,omitempty" yaml:"AuthenticationProfile,omitempty"`
}

type GetSecurityLevelInfo struct {
	Token []common.ReferenceToken `xml:"Token" json:"Token" yaml:"Token"`
}

type GetSecurityLevelInfoList struct {
	Limit          int    `xml:"Limit,omitempty" json:"Limit,omitempty" yaml:"Limit,omitempty"`
	StartReference string `xml:"StartReference,omitempty" json:"StartReference,omitempty" yaml:"StartReference,omitempty"`
}

type GetSecurityLevelInfoListResponse struct {
	NextStartReference string              `xml:"NextStartReference,omitempty" json:"NextStartReference,omitempty" yaml:"NextStartReference,omitempty"`
	SecurityLevelInfo  []SecurityLevelInfo `xml:"SecurityLevelInfo,omitempty" json:"SecurityLevelInfo,omitempty" yaml:"SecurityLevelInfo,omitempty"`
}

type GetSecurityLevelInfoResponse struct {
	SecurityLevelInfo []SecurityLevelInfo `xml:"SecurityLevelInfo,omitempty" json:"SecurityLevelInfo,omitempty" yaml:"SecurityLevelInfo,omitempty"`
}

type GetSecurityLevelList struct {
	Limit          int    `xml:"Limit,omitempty" json:"Limit,omitempty" yaml:"Limit,omitempty"`
	StartReference string `xml:"StartReference,omitempty" json:"StartReference,omitempty" yaml:"StartReference,omitempty"`
}

type GetSecurityLevelListResponse struct {
	NextStartReference string          `xml:"NextStartReference,omitempty" json:"NextStartReference,omitempty" yaml:"NextStartReference,omitempty"`
	SecurityLevel      []SecurityLevel `xml:"SecurityLevel,omitempty" json:"SecurityLevel,omitempty" yaml:"SecurityLevel,omitempty"`
}

type GetSecurityLevels struct {
	Token []common.ReferenceToken `xml:"Token" json:"Token" yaml:"Token"`
}

type GetSecurityLevelsResponse struct {
	SecurityLevel []SecurityLevel `xml:"SecurityLevel,omitempty" json:"SecurityLevel,omitempty" yaml:"SecurityLevel,omitempty"`
}

type GetServiceCapabilities struct {
}

type GetServiceCapabilitiesResponse struct {
	Capabilities ServiceCapabilities `xml:"Capabilities,omitempty" json:"Capabilities,omitempty" yaml:"Capabilities,omitempty"`
}

type ModifyAuthenticationProfile struct {
	AuthenticationProfile AuthenticationProfile `xml:"AuthenticationProfile,omitempty" json:"AuthenticationProfile,omitempty" yaml:"AuthenticationProfile,omitempty"`
}

type ModifyAuthenticationProfileResponse struct {
}

type ModifySecurityLevel struct {
	SecurityLevel SecurityLevel `xml:"SecurityLevel,omitempty" json:"SecurityLevel,omitempty" yaml:"SecurityLevel,omitempty"`
}

type ModifySecurityLevelResponse struct {
}

type RecognitionGroup struct {
	RecognitionMethod []RecognitionMethod       `xml:"RecognitionMethod,omitempty" json:"RecognitionMethod,omitempty" yaml:"RecognitionMethod,omitempty"`
	Extension         RecognitionGroupExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type RecognitionGroupExtension []interface{}

type RecognitionMethod struct {
	RecognitionType string                     `xml:"RecognitionType,omitempty" json:"RecognitionType,omitempty" yaml:"RecognitionType,omitempty"`
	Order           int                        `xml:"Order,omitempty" json:"Order,omitempty" yaml:"Order,omitempty"`
	Extension       RecognitionMethodExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type RecognitionMethodExtension []interface{}

type SecurityLevel struct {
	Name             *common.Name           `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Priority         int                    `xml:"Priority,omitempty" json:"Priority,omitempty" yaml:"Priority,omitempty"`
	Description      *common.Description    `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	RecognitionGroup []RecognitionGroup     `xml:"RecognitionGroup,omitempty" json:"RecognitionGroup,omitempty" yaml:"RecognitionGroup,omitempty"`
	Extension        SecurityLevelExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
	TypeAttrXSI      string                 `xml:"xsi:type,attr,omitempty"`
	TypeNamespace    string                 `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *SecurityLevel) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:SecurityLevel"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://www.onvif.org/ver10/authenticationbehavior/wsdl"
	}
}

type SecurityLevelConstraint struct {
	ActiveRegularSchedule    bool                             `xml:"ActiveRegularSchedule,omitempty" json:"ActiveRegularSchedule,omitempty" yaml:"ActiveRegularSchedule,omitempty"`
	ActiveSpecialDaySchedule bool                             `xml:"ActiveSpecialDaySchedule,omitempty" json:"ActiveSpecialDaySchedule,omitempty" yaml:"ActiveSpecialDaySchedule,omitempty"`
	AuthenticationMode       *common.Name                     `xml:"AuthenticationMode,omitempty" json:"AuthenticationMode,omitempty" yaml:"AuthenticationMode,omitempty"`
	SecurityLevelToken       *common.ReferenceToken           `xml:"SecurityLevelToken,omitempty" json:"SecurityLevelToken,omitempty" yaml:"SecurityLevelToken,omitempty"`
	Extension                SecurityLevelConstraintExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type SecurityLevelConstraintExtension []interface{}

type SecurityLevelExtension []interface{}

type SecurityLevelInfo struct {
	Name          *common.Name        `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Priority      int                 `xml:"Priority,omitempty" json:"Priority,omitempty" yaml:"Priority,omitempty"`
	Description   *common.Description `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	TypeAttrXSI   string              `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string              `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *SecurityLevelInfo) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:SecurityLevelInfo"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://www.onvif.org/ver10/authenticationbehavior/wsdl"
	}
}

type ServiceCapabilities []interface{}

type SetAuthenticationProfile struct {
	AuthenticationProfile AuthenticationProfile `xml:"AuthenticationProfile,omitempty" json:"AuthenticationProfile,omitempty" yaml:"AuthenticationProfile,omitempty"`
}

type SetAuthenticationProfileResponse struct {
}

type SetSecurityLevel struct {
	SecurityLevel SecurityLevel `xml:"SecurityLevel,omitempty" json:"SecurityLevel,omitempty" yaml:"SecurityLevel,omitempty"`
}

type SetSecurityLevelResponse struct {
}

// authenticationBehaviorPort implements the AuthenticationBehaviorPort interface.
type authenticationBehaviorPort struct {
	cli      common.Client
	Endpoint string
}

func (p *authenticationBehaviorPort) OptCreateAuthenticationProfile(args CreateAuthenticationProfile) (*CreateAuthenticationProfileResponse, *common.Fault) {
	req := struct {
		XMLName                     string `xml:"tab:CreateAuthenticationProfile"`
		CreateAuthenticationProfile CreateAuthenticationProfile
	}{
		CreateAuthenticationProfile: args,
	}

	resp := CreateAuthenticationProfileResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *authenticationBehaviorPort) OptCreateSecurityLevel(args CreateSecurityLevel) (*CreateSecurityLevelResponse, *common.Fault) {
	req := struct {
		XMLName             string `xml:"tab:CreateSecurityLevel"`
		CreateSecurityLevel CreateSecurityLevel
	}{
		CreateSecurityLevel: args,
	}

	resp := CreateSecurityLevelResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *authenticationBehaviorPort) OptDeleteAuthenticationProfile(args DeleteAuthenticationProfile) (*DeleteAuthenticationProfileResponse, *common.Fault) {
	req := struct {
		XMLName                     string `xml:"tab:DeleteAuthenticationProfile"`
		DeleteAuthenticationProfile DeleteAuthenticationProfile
	}{
		DeleteAuthenticationProfile: args,
	}

	resp := DeleteAuthenticationProfileResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *authenticationBehaviorPort) OptDeleteSecurityLevel(args DeleteSecurityLevel) (*DeleteSecurityLevelResponse, *common.Fault) {
	req := struct {
		XMLName             string `xml:"tab:DeleteSecurityLevel"`
		DeleteSecurityLevel DeleteSecurityLevel
	}{
		DeleteSecurityLevel: args,
	}

	resp := DeleteSecurityLevelResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *authenticationBehaviorPort) OptGetAuthenticationProfileInfo(args GetAuthenticationProfileInfo) (*GetAuthenticationProfileInfoResponse, *common.Fault) {
	req := struct {
		XMLName                      string `xml:"tab:GetAuthenticationProfileInfo"`
		GetAuthenticationProfileInfo GetAuthenticationProfileInfo
	}{
		GetAuthenticationProfileInfo: args,
	}

	resp := GetAuthenticationProfileInfoResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *authenticationBehaviorPort) OptGetAuthenticationProfileInfoList(args GetAuthenticationProfileInfoList) (*GetAuthenticationProfileInfoListResponse, *common.Fault) {
	req := struct {
		XMLName                          string `xml:"tab:GetAuthenticationProfileInfoList"`
		GetAuthenticationProfileInfoList GetAuthenticationProfileInfoList
	}{
		GetAuthenticationProfileInfoList: args,
	}

	resp := GetAuthenticationProfileInfoListResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *authenticationBehaviorPort) OptGetAuthenticationProfileList(args GetAuthenticationProfileList) (*GetAuthenticationProfileListResponse, *common.Fault) {
	req := struct {
		XMLName                      string `xml:"tab:GetAuthenticationProfileList"`
		GetAuthenticationProfileList GetAuthenticationProfileList
	}{
		GetAuthenticationProfileList: args,
	}

	resp := GetAuthenticationProfileListResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *authenticationBehaviorPort) OptGetAuthenticationProfiles(args GetAuthenticationProfiles) (*GetAuthenticationProfilesResponse, *common.Fault) {
	req := struct {
		XMLName                   string `xml:"tab:GetAuthenticationProfiles"`
		GetAuthenticationProfiles GetAuthenticationProfiles
	}{
		GetAuthenticationProfiles: args,
	}

	resp := GetAuthenticationProfilesResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *authenticationBehaviorPort) OptGetSecurityLevelInfo(args GetSecurityLevelInfo) (*GetSecurityLevelInfoResponse, *common.Fault) {
	req := struct {
		XMLName              string `xml:"tab:GetSecurityLevelInfo"`
		GetSecurityLevelInfo GetSecurityLevelInfo
	}{
		GetSecurityLevelInfo: args,
	}

	resp := GetSecurityLevelInfoResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *authenticationBehaviorPort) OptGetSecurityLevelInfoList(args GetSecurityLevelInfoList) (*GetSecurityLevelInfoListResponse, *common.Fault) {
	req := struct {
		XMLName                  string `xml:"tab:GetSecurityLevelInfoList"`
		GetSecurityLevelInfoList GetSecurityLevelInfoList
	}{
		GetSecurityLevelInfoList: args,
	}

	resp := GetSecurityLevelInfoListResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *authenticationBehaviorPort) OptGetSecurityLevelList(args GetSecurityLevelList) (*GetSecurityLevelListResponse, *common.Fault) {
	req := struct {
		XMLName              string `xml:"tab:GetSecurityLevelList"`
		GetSecurityLevelList GetSecurityLevelList
	}{
		GetSecurityLevelList: args,
	}

	resp := GetSecurityLevelListResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *authenticationBehaviorPort) OptGetSecurityLevels(args GetSecurityLevels) (*GetSecurityLevelsResponse, *common.Fault) {
	req := struct {
		XMLName           string `xml:"tab:GetSecurityLevels"`
		GetSecurityLevels GetSecurityLevels
	}{
		GetSecurityLevels: args,
	}

	resp := GetSecurityLevelsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *authenticationBehaviorPort) OptGetServiceCapabilities(args GetServiceCapabilities) (*GetServiceCapabilitiesResponse, *common.Fault) {
	req := struct {
		XMLName                string `xml:"tab:GetServiceCapabilities"`
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

func (p *authenticationBehaviorPort) OptModifyAuthenticationProfile(args ModifyAuthenticationProfile) (*ModifyAuthenticationProfileResponse, *common.Fault) {
	req := struct {
		XMLName                     string `xml:"tab:ModifyAuthenticationProfile"`
		ModifyAuthenticationProfile ModifyAuthenticationProfile
	}{
		ModifyAuthenticationProfile: args,
	}

	resp := ModifyAuthenticationProfileResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *authenticationBehaviorPort) OptModifySecurityLevel(args ModifySecurityLevel) (*ModifySecurityLevelResponse, *common.Fault) {
	req := struct {
		XMLName             string `xml:"tab:ModifySecurityLevel"`
		ModifySecurityLevel ModifySecurityLevel
	}{
		ModifySecurityLevel: args,
	}

	resp := ModifySecurityLevelResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *authenticationBehaviorPort) OptSetAuthenticationProfile(args SetAuthenticationProfile) (*SetAuthenticationProfileResponse, *common.Fault) {
	req := struct {
		XMLName                  string `xml:"tab:SetAuthenticationProfile"`
		SetAuthenticationProfile SetAuthenticationProfile
	}{
		SetAuthenticationProfile: args,
	}

	resp := SetAuthenticationProfileResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *authenticationBehaviorPort) OptSetSecurityLevel(args SetSecurityLevel) (*SetSecurityLevelResponse, *common.Fault) {
	req := struct {
		XMLName          string `xml:"tab:SetSecurityLevel"`
		SetSecurityLevel SetSecurityLevel
	}{
		SetSecurityLevel: args,
	}

	resp := SetSecurityLevelResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}
