package access_control

import (
	"reflect"

	"code.byted.org/videoarch/go-onvif/onvif/common"
)

var Namespace = "http://www.onvif.org/ver10/accesscontrol/wsdl"

// NewPACSPort creates an initializes a PACSPort.
func NewPACSPort(endpoint string, cli common.Client) PACSPort {
	return &pACSPort{cli: cli, Endpoint: endpoint}
}

// PACSPort was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type PACSPort interface {
	OptCreateAccessPoint(CreateAccessPoint CreateAccessPoint) (*CreateAccessPointResponse, error)

	OptCreateArea(CreateArea CreateArea) (*CreateAreaResponse, error)

	OptDeleteAccessPoint(DeleteAccessPoint DeleteAccessPoint) (*DeleteAccessPointResponse, error)

	OptDeleteAccessPointAuthenticationProfile(DeleteAccessPointAuthenticationProfile DeleteAccessPointAuthenticationProfile) (*DeleteAccessPointAuthenticationProfileResponse, error)

	OptDeleteArea(DeleteArea DeleteArea) (*DeleteAreaResponse, error)

	OptDisableAccessPoint(DisableAccessPoint DisableAccessPoint) (*DisableAccessPointResponse, error)

	OptEnableAccessPoint(EnableAccessPoint EnableAccessPoint) (*EnableAccessPointResponse, error)

	OptExternalAuthorization(ExternalAuthorization ExternalAuthorization) (*ExternalAuthorizationResponse, error)

	OptFeedback(Feedback Feedback) (*FeedbackResponse, error)

	OptGetAccessPointInfo(GetAccessPointInfo GetAccessPointInfo) (*GetAccessPointInfoResponse, error)

	OptGetAccessPointInfoList(GetAccessPointInfoList GetAccessPointInfoList) (*GetAccessPointInfoListResponse, error)

	OptGetAccessPointList(GetAccessPointList GetAccessPointList) (*GetAccessPointListResponse, error)

	OptGetAccessPointState(GetAccessPointState GetAccessPointState) (*GetAccessPointStateResponse, error)

	OptGetAccessPoints(GetAccessPoints GetAccessPoints) (*GetAccessPointsResponse, error)

	OptGetAreaInfo(GetAreaInfo GetAreaInfo) (*GetAreaInfoResponse, error)

	OptGetAreaInfoList(GetAreaInfoList GetAreaInfoList) (*GetAreaInfoListResponse, error)

	OptGetAreaList(GetAreaList GetAreaList) (*GetAreaListResponse, error)

	OptGetAreas(GetAreas GetAreas) (*GetAreasResponse, error)

	OptGetServiceCapabilities(GetServiceCapabilities GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error)

	OptModifyAccessPoint(ModifyAccessPoint ModifyAccessPoint) (*ModifyAccessPointResponse, error)

	OptModifyArea(ModifyArea ModifyArea) (*ModifyAreaResponse, error)

	OptSetAccessPoint(SetAccessPoint SetAccessPoint) (*SetAccessPointResponse, error)

	OptSetAccessPointAuthenticationProfile(SetAccessPointAuthenticationProfile SetAccessPointAuthenticationProfile) (*SetAccessPointAuthenticationProfileResponse, error)

	OptSetArea(SetArea SetArea) (*SetAreaResponse, error)
}
type Decision string

func (v Decision) Validate() bool {
	for _, vv := range []string{
		"Granted",
		"Denied",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type DenyReason string

func (v DenyReason) Validate() bool {
	for _, vv := range []string{
		"CredentialNotEnabled",
		"CredentialNotActive",
		"CredentialExpired",
		"InvalidPIN",
		"NotPermittedAtThisTime",
		"Unauthorized",
		"Other",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type FeedbackType string

func (v FeedbackType) Validate() bool {
	for _, vv := range []string{
		"pt:Disabled",
		"pt:Idle",
		"pt:DoorLocked",
		"pt:DoorUnlocked",
		"pt:DoorOpenTooLong",
		"pt:DoorPreAlarmWarning",
		"pt:RequireIdentifier",
		"pt:TextMessage",
		"pt:Processing",
		"pt:RetryIdentifier",
		"pt:AccessGranted",
		"pt:AccessDenied",
		"pt:Ok",
		"pt:Fault",
		"pt:Warning",
		"pt:Alarm",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type AccessPoint struct {
	Name                       *common.Name            `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Description                *common.Description     `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	AreaFrom                   *common.ReferenceToken  `xml:"AreaFrom,omitempty" json:"AreaFrom,omitempty" yaml:"AreaFrom,omitempty"`
	AreaTo                     *common.ReferenceToken  `xml:"AreaTo,omitempty" json:"AreaTo,omitempty" yaml:"AreaTo,omitempty"`
	EntityType                 string                  `xml:"EntityType,omitempty" json:"EntityType,omitempty" yaml:"EntityType,omitempty"`
	Entity                     *common.ReferenceToken  `xml:"Entity,omitempty" json:"Entity,omitempty" yaml:"Entity,omitempty"`
	Capabilities               AccessPointCapabilities `xml:"Capabilities,omitempty" json:"Capabilities,omitempty" yaml:"Capabilities,omitempty"`
	AuthenticationProfileToken *common.ReferenceToken  `xml:"AuthenticationProfileToken,omitempty" json:"AuthenticationProfileToken,omitempty" yaml:"AuthenticationProfileToken,omitempty"`
	Extension                  AccessPointExtension    `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
	TypeAttrXSI                string                  `xml:"xsi:type,attr,omitempty"`
	TypeNamespace              string                  `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *AccessPoint) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:AccessPoint"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://www.onvif.org/ver10/accesscontrol/wsdl"
	}
}

type AccessPointCapabilities struct {
	SupportedSecurityLevels   []*common.ReferenceToken         `xml:"SupportedSecurityLevels,omitempty" json:"SupportedSecurityLevels,omitempty" yaml:"SupportedSecurityLevels,omitempty"`
	Extension                 SupportedSecurityLevelsExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
	DisableAccessPoint        bool                             `xml:"DisableAccessPoint,attr,omitempty" json:"DisableAccessPoint,attr,omitempty" yaml:"DisableAccessPoint,attr,omitempty"`
	Duress                    bool                             `xml:"Duress,attr,omitempty" json:"Duress,attr,omitempty" yaml:"Duress,attr,omitempty"`
	AnonymousAccess           bool                             `xml:"AnonymousAccess,attr,omitempty" json:"AnonymousAccess,attr,omitempty" yaml:"AnonymousAccess,attr,omitempty"`
	AccessTaken               bool                             `xml:"AccessTaken,attr,omitempty" json:"AccessTaken,attr,omitempty" yaml:"AccessTaken,attr,omitempty"`
	ExternalAuthorization     bool                             `xml:"ExternalAuthorization,attr,omitempty" json:"ExternalAuthorization,attr,omitempty" yaml:"ExternalAuthorization,attr,omitempty"`
	SupportedRecognitionTypes common.StringList                `xml:"SupportedRecognitionTypes,attr,omitempty" json:"SupportedRecognitionTypes,attr,omitempty" yaml:"SupportedRecognitionTypes,attr,omitempty"`
	IdentifierAccess          bool                             `xml:"IdentifierAccess,attr,omitempty" json:"IdentifierAccess,attr,omitempty" yaml:"IdentifierAccess,attr,omitempty"`
	SupportedFeedbackTypes    common.StringList                `xml:"SupportedFeedbackTypes,attr,omitempty" json:"SupportedFeedbackTypes,attr,omitempty" yaml:"SupportedFeedbackTypes,attr,omitempty"`
}

type AccessPointExtension []interface{}

type AccessPointInfo struct {
	Name          *common.Name            `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Description   *common.Description     `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	AreaFrom      *common.ReferenceToken  `xml:"AreaFrom,omitempty" json:"AreaFrom,omitempty" yaml:"AreaFrom,omitempty"`
	AreaTo        *common.ReferenceToken  `xml:"AreaTo,omitempty" json:"AreaTo,omitempty" yaml:"AreaTo,omitempty"`
	EntityType    string                  `xml:"EntityType,omitempty" json:"EntityType,omitempty" yaml:"EntityType,omitempty"`
	Entity        *common.ReferenceToken  `xml:"Entity,omitempty" json:"Entity,omitempty" yaml:"Entity,omitempty"`
	Capabilities  AccessPointCapabilities `xml:"Capabilities,omitempty" json:"Capabilities,omitempty" yaml:"Capabilities,omitempty"`
	TypeAttrXSI   string                  `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string                  `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *AccessPointInfo) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:AccessPointInfo"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://www.onvif.org/ver10/accesscontrol/wsdl"
	}
}

type AccessPointInfoBase struct {
	Name          *common.Name           `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Description   *common.Description    `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	AreaFrom      *common.ReferenceToken `xml:"AreaFrom,omitempty" json:"AreaFrom,omitempty" yaml:"AreaFrom,omitempty"`
	AreaTo        *common.ReferenceToken `xml:"AreaTo,omitempty" json:"AreaTo,omitempty" yaml:"AreaTo,omitempty"`
	EntityType    string                 `xml:"EntityType,omitempty" json:"EntityType,omitempty" yaml:"EntityType,omitempty"`
	Entity        *common.ReferenceToken `xml:"Entity,omitempty" json:"Entity,omitempty" yaml:"Entity,omitempty"`
	TypeAttrXSI   string                 `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string                 `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *AccessPointInfoBase) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:AccessPointInfoBase"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://www.onvif.org/ver10/accesscontrol/wsdl"
	}
}

type AccessPointState struct {
	Enabled bool `xml:"Enabled,omitempty" json:"Enabled,omitempty" yaml:"Enabled,omitempty"`
}

type Area struct {
	Name          *common.Name        `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Description   *common.Description `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	Extension     AreaExtension       `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
	TypeAttrXSI   string              `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string              `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *Area) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:Area"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://www.onvif.org/ver10/accesscontrol/wsdl"
	}
}

type AreaExtension []interface{}

type AreaInfo struct {
	Name          *common.Name        `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Description   *common.Description `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	TypeAttrXSI   string              `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string              `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *AreaInfo) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:AreaInfo"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://www.onvif.org/ver10/accesscontrol/wsdl"
	}
}

type AreaInfoBase struct {
	Name          *common.Name        `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Description   *common.Description `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	TypeAttrXSI   string              `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string              `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *AreaInfoBase) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:AreaInfoBase"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://www.onvif.org/ver10/accesscontrol/wsdl"
	}
}

type CreateAccessPoint struct {
	AccessPoint AccessPoint `xml:"AccessPoint,omitempty" json:"AccessPoint,omitempty" yaml:"AccessPoint,omitempty"`
}

type CreateAccessPointResponse struct {
	Token *common.ReferenceToken `xml:"Token,omitempty" json:"Token,omitempty" yaml:"Token,omitempty"`
}

type CreateArea struct {
	Area Area `xml:"Area,omitempty" json:"Area,omitempty" yaml:"Area,omitempty"`
}

type CreateAreaResponse struct {
	Token *common.ReferenceToken `xml:"Token,omitempty" json:"Token,omitempty" yaml:"Token,omitempty"`
}

type DeleteAccessPoint struct {
	Token *common.ReferenceToken `xml:"Token,omitempty" json:"Token,omitempty" yaml:"Token,omitempty"`
}

type DeleteAccessPointAuthenticationProfile struct {
	Token *common.ReferenceToken `xml:"Token,omitempty" json:"Token,omitempty" yaml:"Token,omitempty"`
}

type DeleteAccessPointAuthenticationProfileResponse struct {
}

type DeleteAccessPointResponse struct {
}

type DeleteArea struct {
	Token *common.ReferenceToken `xml:"Token,omitempty" json:"Token,omitempty" yaml:"Token,omitempty"`
}

type DeleteAreaResponse struct {
}

type DisableAccessPoint struct {
	Token *common.ReferenceToken `xml:"Token,omitempty" json:"Token,omitempty" yaml:"Token,omitempty"`
}

type DisableAccessPointResponse struct {
}

type EnableAccessPoint struct {
	Token *common.ReferenceToken `xml:"Token,omitempty" json:"Token,omitempty" yaml:"Token,omitempty"`
}

type EnableAccessPointResponse struct {
}

type ExternalAuthorization struct {
	AccessPointToken *common.ReferenceToken `xml:"AccessPointToken,omitempty" json:"AccessPointToken,omitempty" yaml:"AccessPointToken,omitempty"`
	CredentialToken  *common.ReferenceToken `xml:"CredentialToken,omitempty" json:"CredentialToken,omitempty" yaml:"CredentialToken,omitempty"`
	Reason           string                 `xml:"Reason,omitempty" json:"Reason,omitempty" yaml:"Reason,omitempty"`
	Decision         Decision               `xml:"Decision,omitempty" json:"Decision,omitempty" yaml:"Decision,omitempty"`
}

type ExternalAuthorizationResponse struct {
}

type Feedback struct {
	AccessPointToken *common.ReferenceToken `xml:"AccessPointToken,omitempty" json:"AccessPointToken,omitempty" yaml:"AccessPointToken,omitempty"`
	FeedbackType     string                 `xml:"FeedbackType,omitempty" json:"FeedbackType,omitempty" yaml:"FeedbackType,omitempty"`
	RecognitionType  []string               `xml:"RecognitionType,omitempty" json:"RecognitionType,omitempty" yaml:"RecognitionType,omitempty"`
	TextMessage      string                 `xml:"TextMessage,omitempty" json:"TextMessage,omitempty" yaml:"TextMessage,omitempty"`
}

type FeedbackResponse []interface{}

type GetAccessPointInfo struct {
	Token []*common.ReferenceToken `xml:"Token,omitempty" json:"Token,omitempty" yaml:"Token,omitempty"`
}

type GetAccessPointInfoList struct {
	Limit          int    `xml:"Limit,omitempty" json:"Limit,omitempty" yaml:"Limit,omitempty"`
	StartReference string `xml:"StartReference,omitempty" json:"StartReference,omitempty" yaml:"StartReference,omitempty"`
}

type GetAccessPointInfoListResponse struct {
	NextStartReference string            `xml:"NextStartReference,omitempty" json:"NextStartReference,omitempty" yaml:"NextStartReference,omitempty"`
	AccessPointInfo    []AccessPointInfo `xml:"AccessPointInfo,omitempty" json:"AccessPointInfo,omitempty" yaml:"AccessPointInfo,omitempty"`
}

type GetAccessPointInfoResponse struct {
	AccessPointInfo []AccessPointInfo `xml:"AccessPointInfo,omitempty" json:"AccessPointInfo,omitempty" yaml:"AccessPointInfo,omitempty"`
}

type GetAccessPointList struct {
	Limit          int    `xml:"Limit,omitempty" json:"Limit,omitempty" yaml:"Limit,omitempty"`
	StartReference string `xml:"StartReference,omitempty" json:"StartReference,omitempty" yaml:"StartReference,omitempty"`
}

type GetAccessPointListResponse struct {
	NextStartReference string        `xml:"NextStartReference,omitempty" json:"NextStartReference,omitempty" yaml:"NextStartReference,omitempty"`
	AccessPoint        []AccessPoint `xml:"AccessPoint,omitempty" json:"AccessPoint,omitempty" yaml:"AccessPoint,omitempty"`
}

type GetAccessPointState struct {
	Token *common.ReferenceToken `xml:"Token,omitempty" json:"Token,omitempty" yaml:"Token,omitempty"`
}

type GetAccessPointStateResponse struct {
	AccessPointState AccessPointState `xml:"AccessPointState,omitempty" json:"AccessPointState,omitempty" yaml:"AccessPointState,omitempty"`
}

type GetAccessPoints struct {
	Token []*common.ReferenceToken `xml:"Token,omitempty" json:"Token,omitempty" yaml:"Token,omitempty"`
}

type GetAccessPointsResponse struct {
	AccessPoint []AccessPoint `xml:"AccessPoint,omitempty" json:"AccessPoint,omitempty" yaml:"AccessPoint,omitempty"`
}

type GetAreaInfo struct {
	Token []*common.ReferenceToken `xml:"Token,omitempty" json:"Token,omitempty" yaml:"Token,omitempty"`
}

type GetAreaInfoList struct {
	Limit          int    `xml:"Limit,omitempty" json:"Limit,omitempty" yaml:"Limit,omitempty"`
	StartReference string `xml:"StartReference,omitempty" json:"StartReference,omitempty" yaml:"StartReference,omitempty"`
}

type GetAreaInfoListResponse struct {
	NextStartReference string     `xml:"NextStartReference,omitempty" json:"NextStartReference,omitempty" yaml:"NextStartReference,omitempty"`
	AreaInfo           []AreaInfo `xml:"AreaInfo,omitempty" json:"AreaInfo,omitempty" yaml:"AreaInfo,omitempty"`
}

type GetAreaInfoResponse struct {
	AreaInfo []AreaInfo `xml:"AreaInfo,omitempty" json:"AreaInfo,omitempty" yaml:"AreaInfo,omitempty"`
}

type GetAreaList struct {
	Limit          int    `xml:"Limit,omitempty" json:"Limit,omitempty" yaml:"Limit,omitempty"`
	StartReference string `xml:"StartReference,omitempty" json:"StartReference,omitempty" yaml:"StartReference,omitempty"`
}

type GetAreaListResponse struct {
	NextStartReference string `xml:"NextStartReference,omitempty" json:"NextStartReference,omitempty" yaml:"NextStartReference,omitempty"`
	Area               []Area `xml:"Area,omitempty" json:"Area,omitempty" yaml:"Area,omitempty"`
}

type GetAreas struct {
	Token []*common.ReferenceToken `xml:"Token,omitempty" json:"Token,omitempty" yaml:"Token,omitempty"`
}

type GetAreasResponse struct {
	Area []Area `xml:"Area,omitempty" json:"Area,omitempty" yaml:"Area,omitempty"`
}

type GetServiceCapabilities struct {
}

type GetServiceCapabilitiesResponse struct {
	Capabilities ServiceCapabilities `xml:"Capabilities,omitempty" json:"Capabilities,omitempty" yaml:"Capabilities,omitempty"`
}

type ModifyAccessPoint struct {
	AccessPoint AccessPoint `xml:"AccessPoint,omitempty" json:"AccessPoint,omitempty" yaml:"AccessPoint,omitempty"`
}

type ModifyAccessPointResponse struct {
}

type ModifyArea struct {
	Area Area `xml:"Area,omitempty" json:"Area,omitempty" yaml:"Area,omitempty"`
}

type ModifyAreaResponse struct {
}

type ServiceCapabilities []interface{}

type SetAccessPoint struct {
	AccessPoint AccessPoint `xml:"AccessPoint,omitempty" json:"AccessPoint,omitempty" yaml:"AccessPoint,omitempty"`
}

type SetAccessPointAuthenticationProfile struct {
	Token                      *common.ReferenceToken `xml:"Token,omitempty" json:"Token,omitempty" yaml:"Token,omitempty"`
	AuthenticationProfileToken *common.ReferenceToken `xml:"AuthenticationProfileToken,omitempty" json:"AuthenticationProfileToken,omitempty" yaml:"AuthenticationProfileToken,omitempty"`
}

type SetAccessPointAuthenticationProfileResponse struct {
}

type SetAccessPointResponse struct {
}

type SetArea struct {
	Area Area `xml:"Area,omitempty" json:"Area,omitempty" yaml:"Area,omitempty"`
}

type SetAreaResponse struct {
}

type SupportedSecurityLevelsExtension []interface{}

// pACSPort implements the PACSPort interface.
type pACSPort struct {
	cli      common.Client
	Endpoint string
}

func (p *pACSPort) OptCreateAccessPoint(args CreateAccessPoint) (*CreateAccessPointResponse, error) {
	req := struct {
		XMLName           string `xml:"tac:CreateAccessPoint"`
		CreateAccessPoint CreateAccessPoint
	}{
		CreateAccessPoint: args,
	}

	resp := CreateAccessPointResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *pACSPort) OptCreateArea(args CreateArea) (*CreateAreaResponse, error) {
	req := struct {
		XMLName    string `xml:"tac:CreateArea"`
		CreateArea CreateArea
	}{
		CreateArea: args,
	}

	resp := CreateAreaResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *pACSPort) OptDeleteAccessPoint(args DeleteAccessPoint) (*DeleteAccessPointResponse, error) {
	req := struct {
		XMLName           string `xml:"tac:DeleteAccessPoint"`
		DeleteAccessPoint DeleteAccessPoint
	}{
		DeleteAccessPoint: args,
	}

	resp := DeleteAccessPointResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *pACSPort) OptDeleteAccessPointAuthenticationProfile(args DeleteAccessPointAuthenticationProfile) (*DeleteAccessPointAuthenticationProfileResponse, error) {
	req := struct {
		XMLName                                string `xml:"tac:DeleteAccessPointAuthenticationProfile"`
		DeleteAccessPointAuthenticationProfile DeleteAccessPointAuthenticationProfile
	}{
		DeleteAccessPointAuthenticationProfile: args,
	}

	resp := DeleteAccessPointAuthenticationProfileResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *pACSPort) OptDeleteArea(args DeleteArea) (*DeleteAreaResponse, error) {
	req := struct {
		XMLName    string `xml:"tac:DeleteArea"`
		DeleteArea DeleteArea
	}{
		DeleteArea: args,
	}

	resp := DeleteAreaResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *pACSPort) OptDisableAccessPoint(args DisableAccessPoint) (*DisableAccessPointResponse, error) {
	req := struct {
		XMLName            string `xml:"tac:DisableAccessPoint"`
		DisableAccessPoint DisableAccessPoint
	}{
		DisableAccessPoint: args,
	}

	resp := DisableAccessPointResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *pACSPort) OptEnableAccessPoint(args EnableAccessPoint) (*EnableAccessPointResponse, error) {
	req := struct {
		XMLName           string `xml:"tac:EnableAccessPoint"`
		EnableAccessPoint EnableAccessPoint
	}{
		EnableAccessPoint: args,
	}

	resp := EnableAccessPointResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *pACSPort) OptExternalAuthorization(args ExternalAuthorization) (*ExternalAuthorizationResponse, error) {
	req := struct {
		XMLName               string `xml:"tac:ExternalAuthorization"`
		ExternalAuthorization ExternalAuthorization
	}{
		ExternalAuthorization: args,
	}

	resp := ExternalAuthorizationResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *pACSPort) OptFeedback(args Feedback) (*FeedbackResponse, error) {
	req := struct {
		XMLName  string `xml:"tac:Feedback"`
		Feedback Feedback
	}{
		Feedback: args,
	}

	resp := FeedbackResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *pACSPort) OptGetAccessPointInfo(args GetAccessPointInfo) (*GetAccessPointInfoResponse, error) {
	req := struct {
		XMLName            string `xml:"tac:GetAccessPointInfo"`
		GetAccessPointInfo GetAccessPointInfo
	}{
		GetAccessPointInfo: args,
	}

	resp := GetAccessPointInfoResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *pACSPort) OptGetAccessPointInfoList(args GetAccessPointInfoList) (*GetAccessPointInfoListResponse, error) {
	req := struct {
		XMLName                string `xml:"tac:GetAccessPointInfoList"`
		GetAccessPointInfoList GetAccessPointInfoList
	}{
		GetAccessPointInfoList: args,
	}

	resp := GetAccessPointInfoListResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *pACSPort) OptGetAccessPointList(args GetAccessPointList) (*GetAccessPointListResponse, error) {
	req := struct {
		XMLName            string `xml:"tac:GetAccessPointList"`
		GetAccessPointList GetAccessPointList
	}{
		GetAccessPointList: args,
	}

	resp := GetAccessPointListResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *pACSPort) OptGetAccessPointState(args GetAccessPointState) (*GetAccessPointStateResponse, error) {
	req := struct {
		XMLName             string `xml:"tac:GetAccessPointState"`
		GetAccessPointState GetAccessPointState
	}{
		GetAccessPointState: args,
	}

	resp := GetAccessPointStateResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *pACSPort) OptGetAccessPoints(args GetAccessPoints) (*GetAccessPointsResponse, error) {
	req := struct {
		XMLName         string `xml:"tac:GetAccessPoints"`
		GetAccessPoints GetAccessPoints
	}{
		GetAccessPoints: args,
	}

	resp := GetAccessPointsResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *pACSPort) OptGetAreaInfo(args GetAreaInfo) (*GetAreaInfoResponse, error) {
	req := struct {
		XMLName     string `xml:"tac:GetAreaInfo"`
		GetAreaInfo GetAreaInfo
	}{
		GetAreaInfo: args,
	}

	resp := GetAreaInfoResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *pACSPort) OptGetAreaInfoList(args GetAreaInfoList) (*GetAreaInfoListResponse, error) {
	req := struct {
		XMLName         string `xml:"tac:GetAreaInfoList"`
		GetAreaInfoList GetAreaInfoList
	}{
		GetAreaInfoList: args,
	}

	resp := GetAreaInfoListResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *pACSPort) OptGetAreaList(args GetAreaList) (*GetAreaListResponse, error) {
	req := struct {
		XMLName     string `xml:"tac:GetAreaList"`
		GetAreaList GetAreaList
	}{
		GetAreaList: args,
	}

	resp := GetAreaListResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *pACSPort) OptGetAreas(args GetAreas) (*GetAreasResponse, error) {
	req := struct {
		XMLName  string `xml:"tac:GetAreas"`
		GetAreas GetAreas
	}{
		GetAreas: args,
	}

	resp := GetAreasResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *pACSPort) OptGetServiceCapabilities(args GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error) {
	req := struct {
		XMLName                string `xml:"tac:GetServiceCapabilities"`
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

func (p *pACSPort) OptModifyAccessPoint(args ModifyAccessPoint) (*ModifyAccessPointResponse, error) {
	req := struct {
		XMLName           string `xml:"tac:ModifyAccessPoint"`
		ModifyAccessPoint ModifyAccessPoint
	}{
		ModifyAccessPoint: args,
	}

	resp := ModifyAccessPointResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *pACSPort) OptModifyArea(args ModifyArea) (*ModifyAreaResponse, error) {
	req := struct {
		XMLName    string `xml:"tac:ModifyArea"`
		ModifyArea ModifyArea
	}{
		ModifyArea: args,
	}

	resp := ModifyAreaResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *pACSPort) OptSetAccessPoint(args SetAccessPoint) (*SetAccessPointResponse, error) {
	req := struct {
		XMLName        string `xml:"tac:SetAccessPoint"`
		SetAccessPoint SetAccessPoint
	}{
		SetAccessPoint: args,
	}

	resp := SetAccessPointResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *pACSPort) OptSetAccessPointAuthenticationProfile(args SetAccessPointAuthenticationProfile) (*SetAccessPointAuthenticationProfileResponse, error) {
	req := struct {
		XMLName                             string `xml:"tac:SetAccessPointAuthenticationProfile"`
		SetAccessPointAuthenticationProfile SetAccessPointAuthenticationProfile
	}{
		SetAccessPointAuthenticationProfile: args,
	}

	resp := SetAccessPointAuthenticationProfileResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *pACSPort) OptSetArea(args SetArea) (*SetAreaResponse, error) {
	req := struct {
		XMLName string `xml:"tac:SetArea"`
		SetArea SetArea
	}{
		SetArea: args,
	}

	resp := SetAreaResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
