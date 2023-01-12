package schedule

import (
	"code.byted.org/videoarch/go-onvif/onvif/common"
)

var Namespace = "http://www.onvif.org/ver10/schedule/wsdl"

// NewSchedulePort creates an initializes a SchedulePort.
func NewSchedulePort(endpoint string, cli common.Client) SchedulePort {
	return &schedulePort{cli: cli, Endpoint: endpoint}
}

// SchedulePort was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type SchedulePort interface {
	OptCreateSchedule(CreateSchedule CreateSchedule) (*CreateScheduleResponse, error)

	OptCreateSpecialDayGroup(CreateSpecialDayGroup CreateSpecialDayGroup) (*CreateSpecialDayGroupResponse, error)

	OptDeleteSchedule(DeleteSchedule DeleteSchedule) (*DeleteScheduleResponse, error)

	OptDeleteSpecialDayGroup(DeleteSpecialDayGroup DeleteSpecialDayGroup) (*DeleteSpecialDayGroupResponse, error)

	OptGetScheduleInfo(GetScheduleInfo GetScheduleInfo) (*GetScheduleInfoResponse, error)

	OptGetScheduleInfoList(GetScheduleInfoList GetScheduleInfoList) (*GetScheduleInfoListResponse, error)

	OptGetScheduleList(GetScheduleList GetScheduleList) (*GetScheduleListResponse, error)

	OptGetScheduleState(GetScheduleState GetScheduleState) (*GetScheduleStateResponse, error)

	OptGetSchedules(GetSchedules GetSchedules) (*GetSchedulesResponse, error)

	OptGetServiceCapabilities(GetServiceCapabilities GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error)

	OptGetSpecialDayGroupInfo(GetSpecialDayGroupInfo GetSpecialDayGroupInfo) (*GetSpecialDayGroupInfoResponse, error)

	OptGetSpecialDayGroupInfoList(GetSpecialDayGroupInfoList GetSpecialDayGroupInfoList) (*GetSpecialDayGroupInfoListResponse, error)

	OptGetSpecialDayGroupList(GetSpecialDayGroupList GetSpecialDayGroupList) (*GetSpecialDayGroupListResponse, error)

	OptGetSpecialDayGroups(GetSpecialDayGroups GetSpecialDayGroups) (*GetSpecialDayGroupsResponse, error)

	OptModifySchedule(ModifySchedule ModifySchedule) (*ModifyScheduleResponse, error)

	OptModifySpecialDayGroup(ModifySpecialDayGroup ModifySpecialDayGroup) (*ModifySpecialDayGroupResponse, error)

	OptSetSchedule(SetSchedule SetSchedule) (*SetScheduleResponse, error)

	OptSetSpecialDayGroup(SetSpecialDayGroup SetSpecialDayGroup) (*SetSpecialDayGroupResponse, error)
}
type Time string

type CreateSchedule struct {
	Schedule Schedule `xml:"Schedule" json:"Schedule" yaml:"Schedule"`
}

type CreateScheduleResponse struct {
	Token *common.ReferenceToken `xml:"Token,omitempty" json:"Token,omitempty" yaml:"Token,omitempty"`
}

type CreateSpecialDayGroup struct {
	SpecialDayGroup SpecialDayGroup `xml:"SpecialDayGroup" json:"SpecialDayGroup" yaml:"SpecialDayGroup"`
}

type CreateSpecialDayGroupResponse struct {
	Token common.ReferenceToken `xml:"Token" json:"Token" yaml:"Token"`
}

type DeleteSchedule struct {
	Token common.ReferenceToken `xml:"Token" json:"Token" yaml:"Token"`
}

type DeleteScheduleResponse struct {
}

type DeleteSpecialDayGroup struct {
	Token *common.ReferenceToken `xml:"Token,omitempty" json:"Token,omitempty" yaml:"Token,omitempty"`
}

type DeleteSpecialDayGroupResponse struct {
}

type GetScheduleInfo struct {
	Token []common.ReferenceToken `xml:"Token" json:"Token" yaml:"Token"`
}

type GetScheduleInfoList struct {
	Limit          int    `xml:"Limit,omitempty" json:"Limit,omitempty" yaml:"Limit,omitempty"`
	StartReference string `xml:"StartReference,omitempty" json:"StartReference,omitempty" yaml:"StartReference,omitempty"`
}

type GetScheduleInfoListResponse struct {
	NextStartReference string         `xml:"NextStartReference,omitempty" json:"NextStartReference,omitempty" yaml:"NextStartReference,omitempty"`
	ScheduleInfo       []ScheduleInfo `xml:"ScheduleInfo,omitempty" json:"ScheduleInfo,omitempty" yaml:"ScheduleInfo,omitempty"`
}

type GetScheduleInfoResponse struct {
	ScheduleInfo []ScheduleInfo `xml:"ScheduleInfo,omitempty" json:"ScheduleInfo,omitempty" yaml:"ScheduleInfo,omitempty"`
}

type GetScheduleList struct {
	Limit          int    `xml:"Limit,omitempty" json:"Limit,omitempty" yaml:"Limit,omitempty"`
	StartReference string `xml:"StartReference,omitempty" json:"StartReference,omitempty" yaml:"StartReference,omitempty"`
}

type GetScheduleListResponse struct {
	NextStartReference string     `xml:"NextStartReference,omitempty" json:"NextStartReference,omitempty" yaml:"NextStartReference,omitempty"`
	Schedule           []Schedule `xml:"Schedule,omitempty" json:"Schedule,omitempty" yaml:"Schedule,omitempty"`
}

type GetScheduleState struct {
	Token *common.ReferenceToken `xml:"Token,omitempty" json:"Token,omitempty" yaml:"Token,omitempty"`
}

type GetScheduleStateResponse struct {
	ScheduleState ScheduleState `xml:"ScheduleState,omitempty" json:"ScheduleState,omitempty" yaml:"ScheduleState,omitempty"`
}

type GetSchedules struct {
	Token []common.ReferenceToken `xml:"Token" json:"Token" yaml:"Token"`
}

type GetSchedulesResponse struct {
	Schedule []Schedule `xml:"Schedule,omitempty" json:"Schedule,omitempty" yaml:"Schedule,omitempty"`
}

type GetServiceCapabilities struct {
}

type GetServiceCapabilitiesResponse struct {
	Capabilities ServiceCapabilities `xml:"Capabilities,omitempty" json:"Capabilities,omitempty" yaml:"Capabilities,omitempty"`
}

type GetSpecialDayGroupInfo struct {
	Token []common.ReferenceToken `xml:"Token" json:"Token" yaml:"Token"`
}

type GetSpecialDayGroupInfoList struct {
	Limit          int    `xml:"Limit,omitempty" json:"Limit,omitempty" yaml:"Limit,omitempty"`
	StartReference string `xml:"StartReference,omitempty" json:"StartReference,omitempty" yaml:"StartReference,omitempty"`
}

type GetSpecialDayGroupInfoListResponse struct {
	NextStartReference  string                `xml:"NextStartReference,omitempty" json:"NextStartReference,omitempty" yaml:"NextStartReference,omitempty"`
	SpecialDayGroupInfo []SpecialDayGroupInfo `xml:"SpecialDayGroupInfo,omitempty" json:"SpecialDayGroupInfo,omitempty" yaml:"SpecialDayGroupInfo,omitempty"`
}

type GetSpecialDayGroupInfoResponse struct {
	SpecialDayGroupInfo []SpecialDayGroupInfo `xml:"SpecialDayGroupInfo,omitempty" json:"SpecialDayGroupInfo,omitempty" yaml:"SpecialDayGroupInfo,omitempty"`
}

type GetSpecialDayGroupList struct {
	Limit          int    `xml:"Limit,omitempty" json:"Limit,omitempty" yaml:"Limit,omitempty"`
	StartReference string `xml:"StartReference,omitempty" json:"StartReference,omitempty" yaml:"StartReference,omitempty"`
}

type GetSpecialDayGroupListResponse struct {
	NextStartReference string            `xml:"NextStartReference,omitempty" json:"NextStartReference,omitempty" yaml:"NextStartReference,omitempty"`
	SpecialDayGroup    []SpecialDayGroup `xml:"SpecialDayGroup,omitempty" json:"SpecialDayGroup,omitempty" yaml:"SpecialDayGroup,omitempty"`
}

type GetSpecialDayGroups struct {
	Token []common.ReferenceToken `xml:"Token" json:"Token" yaml:"Token"`
}

type GetSpecialDayGroupsResponse struct {
	SpecialDayGroup []SpecialDayGroup `xml:"SpecialDayGroup,omitempty" json:"SpecialDayGroup,omitempty" yaml:"SpecialDayGroup,omitempty"`
}

type ModifySchedule struct {
	Schedule Schedule `xml:"Schedule" json:"Schedule" yaml:"Schedule"`
}

type ModifyScheduleResponse struct {
}

type ModifySpecialDayGroup struct {
	SpecialDayGroup SpecialDayGroup `xml:"SpecialDayGroup" json:"SpecialDayGroup" yaml:"SpecialDayGroup"`
}

type ModifySpecialDayGroupResponse struct {
}

type Schedule struct {
	Name          *common.Name          `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Description   *common.Description   `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	Standard      string                `xml:"Standard,omitempty" json:"Standard,omitempty" yaml:"Standard,omitempty"`
	SpecialDays   []SpecialDaysSchedule `xml:"SpecialDays,omitempty" json:"SpecialDays,omitempty" yaml:"SpecialDays,omitempty"`
	Extension     ScheduleExtension     `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
	TypeAttrXSI   string                `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string                `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *Schedule) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:Schedule"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://www.onvif.org/ver10/schedule/wsdl"
	}
}

type ScheduleExtension []interface{}

type ScheduleInfo struct {
	Name          *common.Name        `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Description   *common.Description `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	TypeAttrXSI   string              `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string              `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *ScheduleInfo) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:ScheduleInfo"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://www.onvif.org/ver10/schedule/wsdl"
	}
}

type ScheduleState struct {
	Active     bool                   `xml:"Active,omitempty" json:"Active,omitempty" yaml:"Active,omitempty"`
	SpecialDay bool                   `xml:"SpecialDay,omitempty" json:"SpecialDay,omitempty" yaml:"SpecialDay,omitempty"`
	Extension  ScheduleStateExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type ScheduleStateExtension []interface{}

type ServiceCapabilities []interface{}

type SetSchedule struct {
	Schedule Schedule `xml:"Schedule" json:"Schedule" yaml:"Schedule"`
}

type SetScheduleResponse struct {
}

type SetSpecialDayGroup struct {
	SpecialDayGroup SpecialDayGroup `xml:"SpecialDayGroup" json:"SpecialDayGroup" yaml:"SpecialDayGroup"`
}

type SetSpecialDayGroupResponse struct {
}

type SpecialDayGroup struct {
	Name          *common.Name             `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Description   *common.Description      `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	Days          string                   `xml:"Days,omitempty" json:"Days,omitempty" yaml:"Days,omitempty"`
	Extension     SpecialDayGroupExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
	TypeAttrXSI   string                   `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string                   `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *SpecialDayGroup) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:SpecialDayGroup"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://www.onvif.org/ver10/schedule/wsdl"
	}
}

type SpecialDayGroupExtension []interface{}

type SpecialDayGroupInfo struct {
	Name          *common.Name        `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Description   *common.Description `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	TypeAttrXSI   string              `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string              `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *SpecialDayGroupInfo) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:SpecialDayGroupInfo"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://www.onvif.org/ver10/schedule/wsdl"
	}
}

type SpecialDaysSchedule struct {
	GroupToken *common.ReferenceToken       `xml:"GroupToken,omitempty" json:"GroupToken,omitempty" yaml:"GroupToken,omitempty"`
	TimeRange  []TimePeriod                 `xml:"TimeRange,omitempty" json:"TimeRange,omitempty" yaml:"TimeRange,omitempty"`
	Extension  SpecialDaysScheduleExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type SpecialDaysScheduleExtension []interface{}

type TimePeriod struct {
	From      *Time               `xml:"From,omitempty" json:"From,omitempty" yaml:"From,omitempty"`
	Until     *Time               `xml:"Until,omitempty" json:"Until,omitempty" yaml:"Until,omitempty"`
	Extension TimePeriodExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type TimePeriodExtension []interface{}

// schedulePort implements the SchedulePort interface.
type schedulePort struct {
	cli      common.Client
	Endpoint string
}

func (p *schedulePort) OptCreateSchedule(args CreateSchedule) (*CreateScheduleResponse, error) {
	req := struct {
		XMLName        string `xml:"tsc:CreateSchedule"`
		CreateSchedule CreateSchedule
	}{
		CreateSchedule: args,
	}

	resp := CreateScheduleResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *schedulePort) OptCreateSpecialDayGroup(args CreateSpecialDayGroup) (*CreateSpecialDayGroupResponse, error) {
	req := struct {
		XMLName               string `xml:"tsc:CreateSpecialDayGroup"`
		CreateSpecialDayGroup CreateSpecialDayGroup
	}{
		CreateSpecialDayGroup: args,
	}

	resp := CreateSpecialDayGroupResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *schedulePort) OptDeleteSchedule(args DeleteSchedule) (*DeleteScheduleResponse, error) {
	req := struct {
		XMLName        string `xml:"tsc:DeleteSchedule"`
		DeleteSchedule DeleteSchedule
	}{
		DeleteSchedule: args,
	}

	resp := DeleteScheduleResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *schedulePort) OptDeleteSpecialDayGroup(args DeleteSpecialDayGroup) (*DeleteSpecialDayGroupResponse, error) {
	req := struct {
		XMLName               string `xml:"tsc:DeleteSpecialDayGroup"`
		DeleteSpecialDayGroup DeleteSpecialDayGroup
	}{
		DeleteSpecialDayGroup: args,
	}

	resp := DeleteSpecialDayGroupResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *schedulePort) OptGetScheduleInfo(args GetScheduleInfo) (*GetScheduleInfoResponse, error) {
	req := struct {
		XMLName         string `xml:"tsc:GetScheduleInfo"`
		GetScheduleInfo GetScheduleInfo
	}{
		GetScheduleInfo: args,
	}

	resp := GetScheduleInfoResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *schedulePort) OptGetScheduleInfoList(args GetScheduleInfoList) (*GetScheduleInfoListResponse, error) {
	req := struct {
		XMLName             string `xml:"tsc:GetScheduleInfoList"`
		GetScheduleInfoList GetScheduleInfoList
	}{
		GetScheduleInfoList: args,
	}

	resp := GetScheduleInfoListResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *schedulePort) OptGetScheduleList(args GetScheduleList) (*GetScheduleListResponse, error) {
	req := struct {
		XMLName         string `xml:"tsc:GetScheduleList"`
		GetScheduleList GetScheduleList
	}{
		GetScheduleList: args,
	}

	resp := GetScheduleListResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *schedulePort) OptGetScheduleState(args GetScheduleState) (*GetScheduleStateResponse, error) {
	req := struct {
		XMLName          string `xml:"tsc:GetScheduleState"`
		GetScheduleState GetScheduleState
	}{
		GetScheduleState: args,
	}

	resp := GetScheduleStateResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *schedulePort) OptGetSchedules(args GetSchedules) (*GetSchedulesResponse, error) {
	req := struct {
		XMLName      string `xml:"tsc:GetSchedules"`
		GetSchedules GetSchedules
	}{
		GetSchedules: args,
	}

	resp := GetSchedulesResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *schedulePort) OptGetServiceCapabilities(args GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error) {
	req := struct {
		XMLName                string `xml:"tsc:GetServiceCapabilities"`
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

func (p *schedulePort) OptGetSpecialDayGroupInfo(args GetSpecialDayGroupInfo) (*GetSpecialDayGroupInfoResponse, error) {
	req := struct {
		XMLName                string `xml:"tsc:GetSpecialDayGroupInfo"`
		GetSpecialDayGroupInfo GetSpecialDayGroupInfo
	}{
		GetSpecialDayGroupInfo: args,
	}

	resp := GetSpecialDayGroupInfoResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *schedulePort) OptGetSpecialDayGroupInfoList(args GetSpecialDayGroupInfoList) (*GetSpecialDayGroupInfoListResponse, error) {
	req := struct {
		XMLName                    string `xml:"tsc:GetSpecialDayGroupInfoList"`
		GetSpecialDayGroupInfoList GetSpecialDayGroupInfoList
	}{
		GetSpecialDayGroupInfoList: args,
	}

	resp := GetSpecialDayGroupInfoListResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *schedulePort) OptGetSpecialDayGroupList(args GetSpecialDayGroupList) (*GetSpecialDayGroupListResponse, error) {
	req := struct {
		XMLName                string `xml:"tsc:GetSpecialDayGroupList"`
		GetSpecialDayGroupList GetSpecialDayGroupList
	}{
		GetSpecialDayGroupList: args,
	}

	resp := GetSpecialDayGroupListResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *schedulePort) OptGetSpecialDayGroups(args GetSpecialDayGroups) (*GetSpecialDayGroupsResponse, error) {
	req := struct {
		XMLName             string `xml:"tsc:GetSpecialDayGroups"`
		GetSpecialDayGroups GetSpecialDayGroups
	}{
		GetSpecialDayGroups: args,
	}

	resp := GetSpecialDayGroupsResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *schedulePort) OptModifySchedule(args ModifySchedule) (*ModifyScheduleResponse, error) {
	req := struct {
		XMLName        string `xml:"tsc:ModifySchedule"`
		ModifySchedule ModifySchedule
	}{
		ModifySchedule: args,
	}

	resp := ModifyScheduleResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *schedulePort) OptModifySpecialDayGroup(args ModifySpecialDayGroup) (*ModifySpecialDayGroupResponse, error) {
	req := struct {
		XMLName               string `xml:"tsc:ModifySpecialDayGroup"`
		ModifySpecialDayGroup ModifySpecialDayGroup
	}{
		ModifySpecialDayGroup: args,
	}

	resp := ModifySpecialDayGroupResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *schedulePort) OptSetSchedule(args SetSchedule) (*SetScheduleResponse, error) {
	req := struct {
		XMLName     string `xml:"tsc:SetSchedule"`
		SetSchedule SetSchedule
	}{
		SetSchedule: args,
	}

	resp := SetScheduleResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *schedulePort) OptSetSpecialDayGroup(args SetSpecialDayGroup) (*SetSpecialDayGroupResponse, error) {
	req := struct {
		XMLName            string `xml:"tsc:SetSpecialDayGroup"`
		SetSpecialDayGroup SetSpecialDayGroup
	}{
		SetSpecialDayGroup: args,
	}

	resp := SetSpecialDayGroupResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
