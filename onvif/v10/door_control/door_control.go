package door_control

import (
	"reflect"

	"code.byted.org/videoarch/go-onvif/onvif/common"
)

var Namespace = "http://www.onvif.org/ver10/doorcontrol/wsdl"

// NewDoorControlPort creates an initializes a DoorControlPort.
func NewDoorControlPort(endpoint string, cli common.Client) DoorControlPort {
	return &doorControlPort{cli: cli, Endpoint: endpoint}
}

// DoorControlPort was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type DoorControlPort interface {
	OptAccessDoor(AccessDoor AccessDoor) (*AccessDoorResponse, *common.Fault)

	OptBlockDoor(BlockDoor BlockDoor) (*BlockDoorResponse, *common.Fault)

	OptCreateDoor(CreateDoor CreateDoor) (*CreateDoorResponse, *common.Fault)

	OptDeleteDoor(DeleteDoor DeleteDoor) (*DeleteDoorResponse, *common.Fault)

	OptDoubleLockDoor(DoubleLockDoor DoubleLockDoor) (*DoubleLockDoorResponse, *common.Fault)

	OptGetDoorInfo(GetDoorInfo GetDoorInfo) (*GetDoorInfoResponse, *common.Fault)

	OptGetDoorInfoList(GetDoorInfoList GetDoorInfoList) (*GetDoorInfoListResponse, *common.Fault)

	OptGetDoorList(GetDoorList GetDoorList) (*GetDoorListResponse, *common.Fault)

	OptGetDoorState(GetDoorState GetDoorState) (*GetDoorStateResponse, *common.Fault)

	OptGetDoors(GetDoors GetDoors) (*GetDoorsResponse, *common.Fault)

	OptGetServiceCapabilities(GetServiceCapabilities GetServiceCapabilities) (*GetServiceCapabilitiesResponse, *common.Fault)

	OptLockDoor(LockDoor LockDoor) (*LockDoorResponse, *common.Fault)

	OptLockDownDoor(LockDownDoor LockDownDoor) (*LockDownDoorResponse, *common.Fault)

	OptLockDownReleaseDoor(LockDownReleaseDoor LockDownReleaseDoor) (*LockDownReleaseDoorResponse, *common.Fault)

	OptLockOpenDoor(LockOpenDoor LockOpenDoor) (*LockOpenDoorResponse, *common.Fault)

	OptLockOpenReleaseDoor(LockOpenReleaseDoor LockOpenReleaseDoor) (*LockOpenReleaseDoorResponse, *common.Fault)

	OptModifyDoor(ModifyDoor ModifyDoor) (*ModifyDoorResponse, *common.Fault)

	OptSetDoor(SetDoor SetDoor) (*SetDoorResponse, *common.Fault)

	OptUnlockDoor(UnlockDoor UnlockDoor) (*UnlockDoorResponse, *common.Fault)
}
type Duration string

type DoorAlarmState string

func (v DoorAlarmState) Validate() bool {
	for _, vv := range []string{
		"Normal",
		"DoorForcedOpen",
		"DoorOpenTooLong",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type DoorFaultState string

func (v DoorFaultState) Validate() bool {
	for _, vv := range []string{
		"Unknown",
		"NotInFault",
		"FaultDetected",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type DoorMode string

func (v DoorMode) Validate() bool {
	for _, vv := range []string{
		"Unknown",
		"Locked",
		"Unlocked",
		"Accessed",
		"Blocked",
		"LockedDown",
		"LockedOpen",
		"DoubleLocked",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type DoorPhysicalState string

func (v DoorPhysicalState) Validate() bool {
	for _, vv := range []string{
		"Unknown",
		"Open",
		"Closed",
		"Fault",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type DoorTamperState string

func (v DoorTamperState) Validate() bool {
	for _, vv := range []string{
		"Unknown",
		"NotInTamper",
		"TamperDetected",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type LockPhysicalState string

func (v LockPhysicalState) Validate() bool {
	for _, vv := range []string{
		"Unknown",
		"Locked",
		"Unlocked",
		"Fault",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type AccessDoor struct {
	Token           *common.ReferenceToken `xml:"Token,omitempty" json:"Token,omitempty" yaml:"Token,omitempty"`
	UseExtendedTime bool                   `xml:"UseExtendedTime,omitempty" json:"UseExtendedTime,omitempty" yaml:"UseExtendedTime,omitempty"`
	AccessTime      *common.Duration       `xml:"AccessTime,omitempty" json:"AccessTime,omitempty" yaml:"AccessTime,omitempty"`
	OpenTooLongTime *common.Duration       `xml:"OpenTooLongTime,omitempty" json:"OpenTooLongTime,omitempty" yaml:"OpenTooLongTime,omitempty"`
	PreAlarmTime    *common.Duration       `xml:"PreAlarmTime,omitempty" json:"PreAlarmTime,omitempty" yaml:"PreAlarmTime,omitempty"`
	Extension       AccessDoorExtension    `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type AccessDoorExtension []interface{}

type AccessDoorResponse struct {
}

type BlockDoor struct {
	Token *common.ReferenceToken `xml:"Token,omitempty" json:"Token,omitempty" yaml:"Token,omitempty"`
}

type BlockDoorResponse struct {
}

type CreateDoor struct {
	Door Door `xml:"Door,omitempty" json:"Door,omitempty" yaml:"Door,omitempty"`
}

type CreateDoorResponse struct {
	Token *common.ReferenceToken `xml:"Token,omitempty" json:"Token,omitempty" yaml:"Token,omitempty"`
}

type DeleteDoor struct {
	Token *common.ReferenceToken `xml:"Token,omitempty" json:"Token,omitempty" yaml:"Token,omitempty"`
}

type DeleteDoorResponse struct {
}

type Door struct {
	Name          *common.Name        `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Description   *common.Description `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	Capabilities  DoorCapabilities    `xml:"Capabilities,omitempty" json:"Capabilities,omitempty" yaml:"Capabilities,omitempty"`
	DoorType      *common.Name        `xml:"DoorType,omitempty" json:"DoorType,omitempty" yaml:"DoorType,omitempty"`
	Timings       Timings             `xml:"Timings,omitempty" json:"Timings,omitempty" yaml:"Timings,omitempty"`
	Extension     DoorExtension       `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
	TypeAttrXSI   string              `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string              `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *Door) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:Door"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://www.onvif.org/ver10/doorcontrol/wsdl"
	}
}

type DoorCapabilities []interface{}

type DoorExtension []interface{}

type DoorFault struct {
	Reason string         `xml:"Reason,omitempty" json:"Reason,omitempty" yaml:"Reason,omitempty"`
	State  DoorFaultState `xml:"State,omitempty" json:"State,omitempty" yaml:"State,omitempty"`
}

type DoorInfo struct {
	Name          *common.Name        `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Description   *common.Description `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	Capabilities  DoorCapabilities    `xml:"Capabilities,omitempty" json:"Capabilities,omitempty" yaml:"Capabilities,omitempty"`
	TypeAttrXSI   string              `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string              `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *DoorInfo) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:DoorInfo"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://www.onvif.org/ver10/doorcontrol/wsdl"
	}
}

type DoorInfoBase struct {
	Name          *common.Name        `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Description   *common.Description `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	TypeAttrXSI   string              `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string              `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *DoorInfoBase) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:DoorInfoBase"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://www.onvif.org/ver10/doorcontrol/wsdl"
	}
}

type DoorState struct {
	DoorPhysicalState       DoorPhysicalState `xml:"DoorPhysicalState,omitempty" json:"DoorPhysicalState,omitempty" yaml:"DoorPhysicalState,omitempty"`
	LockPhysicalState       LockPhysicalState `xml:"LockPhysicalState,omitempty" json:"LockPhysicalState,omitempty" yaml:"LockPhysicalState,omitempty"`
	DoubleLockPhysicalState LockPhysicalState `xml:"DoubleLockPhysicalState,omitempty" json:"DoubleLockPhysicalState,omitempty" yaml:"DoubleLockPhysicalState,omitempty"`
	Alarm                   DoorAlarmState    `xml:"Alarm,omitempty" json:"Alarm,omitempty" yaml:"Alarm,omitempty"`
	Tamper                  DoorTamper        `xml:"Tamper,omitempty" json:"Tamper,omitempty" yaml:"Tamper,omitempty"`
	Fault                   DoorFault         `xml:"Fault,omitempty" json:"Fault,omitempty" yaml:"Fault,omitempty"`
	DoorMode                DoorMode          `xml:"DoorMode,omitempty" json:"DoorMode,omitempty" yaml:"DoorMode,omitempty"`
}

type DoorTamper struct {
	Reason string          `xml:"Reason,omitempty" json:"Reason,omitempty" yaml:"Reason,omitempty"`
	State  DoorTamperState `xml:"State,omitempty" json:"State,omitempty" yaml:"State,omitempty"`
}

type DoubleLockDoor struct {
	Token *common.ReferenceToken `xml:"Token,omitempty" json:"Token,omitempty" yaml:"Token,omitempty"`
}

type DoubleLockDoorResponse struct {
}

type GetDoorInfo struct {
	Token []*common.ReferenceToken `xml:"Token,omitempty" json:"Token,omitempty" yaml:"Token,omitempty"`
}

type GetDoorInfoList struct {
	Limit          int    `xml:"Limit,omitempty" json:"Limit,omitempty" yaml:"Limit,omitempty"`
	StartReference string `xml:"StartReference,omitempty" json:"StartReference,omitempty" yaml:"StartReference,omitempty"`
}

type GetDoorInfoListResponse struct {
	NextStartReference string     `xml:"NextStartReference,omitempty" json:"NextStartReference,omitempty" yaml:"NextStartReference,omitempty"`
	DoorInfo           []DoorInfo `xml:"DoorInfo,omitempty" json:"DoorInfo,omitempty" yaml:"DoorInfo,omitempty"`
}

type GetDoorInfoResponse struct {
	DoorInfo []DoorInfo `xml:"DoorInfo,omitempty" json:"DoorInfo,omitempty" yaml:"DoorInfo,omitempty"`
}

type GetDoorList struct {
	Limit          int    `xml:"Limit,omitempty" json:"Limit,omitempty" yaml:"Limit,omitempty"`
	StartReference string `xml:"StartReference,omitempty" json:"StartReference,omitempty" yaml:"StartReference,omitempty"`
}

type GetDoorListResponse struct {
	NextStartReference string `xml:"NextStartReference,omitempty" json:"NextStartReference,omitempty" yaml:"NextStartReference,omitempty"`
	Door               []Door `xml:"Door,omitempty" json:"Door,omitempty" yaml:"Door,omitempty"`
}

type GetDoorState struct {
	Token *common.ReferenceToken `xml:"Token,omitempty" json:"Token,omitempty" yaml:"Token,omitempty"`
}

type GetDoorStateResponse struct {
	DoorState DoorState `xml:"DoorState,omitempty" json:"DoorState,omitempty" yaml:"DoorState,omitempty"`
}

type GetDoors struct {
	Token []*common.ReferenceToken `xml:"Token,omitempty" json:"Token,omitempty" yaml:"Token,omitempty"`
}

type GetDoorsResponse struct {
	Door []Door `xml:"Door,omitempty" json:"Door,omitempty" yaml:"Door,omitempty"`
}

type GetServiceCapabilities struct {
}

type GetServiceCapabilitiesResponse struct {
	Capabilities ServiceCapabilities `xml:"Capabilities,omitempty" json:"Capabilities,omitempty" yaml:"Capabilities,omitempty"`
}

type LockDoor struct {
	Token *common.ReferenceToken `xml:"Token,omitempty" json:"Token,omitempty" yaml:"Token,omitempty"`
}

type LockDoorResponse struct {
}

type LockDownDoor struct {
	Token *common.ReferenceToken `xml:"Token,omitempty" json:"Token,omitempty" yaml:"Token,omitempty"`
}

type LockDownDoorResponse struct {
}

type LockDownReleaseDoor struct {
	Token *common.ReferenceToken `xml:"Token,omitempty" json:"Token,omitempty" yaml:"Token,omitempty"`
}

type LockDownReleaseDoorResponse struct {
}

type LockOpenDoor struct {
	Token *common.ReferenceToken `xml:"Token,omitempty" json:"Token,omitempty" yaml:"Token,omitempty"`
}

type LockOpenDoorResponse struct {
}

type LockOpenReleaseDoor struct {
	Token *common.ReferenceToken `xml:"Token,omitempty" json:"Token,omitempty" yaml:"Token,omitempty"`
}

type LockOpenReleaseDoorResponse struct {
}

type ModifyDoor struct {
	Door Door `xml:"Door,omitempty" json:"Door,omitempty" yaml:"Door,omitempty"`
}

type ModifyDoorResponse struct {
}

type ServiceCapabilities []interface{}

type SetDoor struct {
	Door Door `xml:"Door,omitempty" json:"Door,omitempty" yaml:"Door,omitempty"`
}

type SetDoorResponse struct {
}

type Timings struct {
	ReleaseTime           *common.Duration `xml:"ReleaseTime,omitempty" json:"ReleaseTime,omitempty" yaml:"ReleaseTime,omitempty"`
	OpenTime              *common.Duration `xml:"OpenTime,omitempty" json:"OpenTime,omitempty" yaml:"OpenTime,omitempty"`
	ExtendedReleaseTime   *common.Duration `xml:"ExtendedReleaseTime,omitempty" json:"ExtendedReleaseTime,omitempty" yaml:"ExtendedReleaseTime,omitempty"`
	DelayTimeBeforeRelock *common.Duration `xml:"DelayTimeBeforeRelock,omitempty" json:"DelayTimeBeforeRelock,omitempty" yaml:"DelayTimeBeforeRelock,omitempty"`
	ExtendedOpenTime      *common.Duration `xml:"ExtendedOpenTime,omitempty" json:"ExtendedOpenTime,omitempty" yaml:"ExtendedOpenTime,omitempty"`
	PreAlarmTime          *common.Duration `xml:"PreAlarmTime,omitempty" json:"PreAlarmTime,omitempty" yaml:"PreAlarmTime,omitempty"`
	Extension             TimingsExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type TimingsExtension []interface{}

type UnlockDoor struct {
	Token *common.ReferenceToken `xml:"Token,omitempty" json:"Token,omitempty" yaml:"Token,omitempty"`
}

type UnlockDoorResponse struct {
}

// doorControlPort implements the DoorControlPort interface.
type doorControlPort struct {
	cli      common.Client
	Endpoint string
}

func (p *doorControlPort) OptAccessDoor(args AccessDoor) (*AccessDoorResponse, *common.Fault) {
	req := struct {
		XMLName    string `xml:"tdc:AccessDoor"`
		AccessDoor AccessDoor
	}{
		AccessDoor: args,
	}

	resp := AccessDoorResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *doorControlPort) OptBlockDoor(args BlockDoor) (*BlockDoorResponse, *common.Fault) {
	req := struct {
		XMLName   string `xml:"tdc:BlockDoor"`
		BlockDoor BlockDoor
	}{
		BlockDoor: args,
	}

	resp := BlockDoorResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *doorControlPort) OptCreateDoor(args CreateDoor) (*CreateDoorResponse, *common.Fault) {
	req := struct {
		XMLName    string `xml:"tdc:CreateDoor"`
		CreateDoor CreateDoor
	}{
		CreateDoor: args,
	}

	resp := CreateDoorResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *doorControlPort) OptDeleteDoor(args DeleteDoor) (*DeleteDoorResponse, *common.Fault) {
	req := struct {
		XMLName    string `xml:"tdc:DeleteDoor"`
		DeleteDoor DeleteDoor
	}{
		DeleteDoor: args,
	}

	resp := DeleteDoorResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *doorControlPort) OptDoubleLockDoor(args DoubleLockDoor) (*DoubleLockDoorResponse, *common.Fault) {
	req := struct {
		XMLName        string `xml:"tdc:DoubleLockDoor"`
		DoubleLockDoor DoubleLockDoor
	}{
		DoubleLockDoor: args,
	}

	resp := DoubleLockDoorResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *doorControlPort) OptGetDoorInfo(args GetDoorInfo) (*GetDoorInfoResponse, *common.Fault) {
	req := struct {
		XMLName     string `xml:"tdc:GetDoorInfo"`
		GetDoorInfo GetDoorInfo
	}{
		GetDoorInfo: args,
	}

	resp := GetDoorInfoResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *doorControlPort) OptGetDoorInfoList(args GetDoorInfoList) (*GetDoorInfoListResponse, *common.Fault) {
	req := struct {
		XMLName         string `xml:"tdc:GetDoorInfoList"`
		GetDoorInfoList GetDoorInfoList
	}{
		GetDoorInfoList: args,
	}

	resp := GetDoorInfoListResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *doorControlPort) OptGetDoorList(args GetDoorList) (*GetDoorListResponse, *common.Fault) {
	req := struct {
		XMLName     string `xml:"tdc:GetDoorList"`
		GetDoorList GetDoorList
	}{
		GetDoorList: args,
	}

	resp := GetDoorListResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *doorControlPort) OptGetDoorState(args GetDoorState) (*GetDoorStateResponse, *common.Fault) {
	req := struct {
		XMLName      string `xml:"tdc:GetDoorState"`
		GetDoorState GetDoorState
	}{
		GetDoorState: args,
	}

	resp := GetDoorStateResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *doorControlPort) OptGetDoors(args GetDoors) (*GetDoorsResponse, *common.Fault) {
	req := struct {
		XMLName  string `xml:"tdc:GetDoors"`
		GetDoors GetDoors
	}{
		GetDoors: args,
	}

	resp := GetDoorsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *doorControlPort) OptGetServiceCapabilities(args GetServiceCapabilities) (*GetServiceCapabilitiesResponse, *common.Fault) {
	req := struct {
		XMLName                string `xml:"tdc:GetServiceCapabilities"`
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

func (p *doorControlPort) OptLockDoor(args LockDoor) (*LockDoorResponse, *common.Fault) {
	req := struct {
		XMLName  string `xml:"tdc:LockDoor"`
		LockDoor LockDoor
	}{
		LockDoor: args,
	}

	resp := LockDoorResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *doorControlPort) OptLockDownDoor(args LockDownDoor) (*LockDownDoorResponse, *common.Fault) {
	req := struct {
		XMLName      string `xml:"tdc:LockDownDoor"`
		LockDownDoor LockDownDoor
	}{
		LockDownDoor: args,
	}

	resp := LockDownDoorResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *doorControlPort) OptLockDownReleaseDoor(args LockDownReleaseDoor) (*LockDownReleaseDoorResponse, *common.Fault) {
	req := struct {
		XMLName             string `xml:"tdc:LockDownReleaseDoor"`
		LockDownReleaseDoor LockDownReleaseDoor
	}{
		LockDownReleaseDoor: args,
	}

	resp := LockDownReleaseDoorResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *doorControlPort) OptLockOpenDoor(args LockOpenDoor) (*LockOpenDoorResponse, *common.Fault) {
	req := struct {
		XMLName      string `xml:"tdc:LockOpenDoor"`
		LockOpenDoor LockOpenDoor
	}{
		LockOpenDoor: args,
	}

	resp := LockOpenDoorResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *doorControlPort) OptLockOpenReleaseDoor(args LockOpenReleaseDoor) (*LockOpenReleaseDoorResponse, *common.Fault) {
	req := struct {
		XMLName             string `xml:"tdc:LockOpenReleaseDoor"`
		LockOpenReleaseDoor LockOpenReleaseDoor
	}{
		LockOpenReleaseDoor: args,
	}

	resp := LockOpenReleaseDoorResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *doorControlPort) OptModifyDoor(args ModifyDoor) (*ModifyDoorResponse, *common.Fault) {
	req := struct {
		XMLName    string `xml:"tdc:ModifyDoor"`
		ModifyDoor ModifyDoor
	}{
		ModifyDoor: args,
	}

	resp := ModifyDoorResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *doorControlPort) OptSetDoor(args SetDoor) (*SetDoorResponse, *common.Fault) {
	req := struct {
		XMLName string `xml:"tdc:SetDoor"`
		SetDoor SetDoor
	}{
		SetDoor: args,
	}

	resp := SetDoorResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *doorControlPort) OptUnlockDoor(args UnlockDoor) (*UnlockDoorResponse, *common.Fault) {
	req := struct {
		XMLName    string `xml:"tdc:UnlockDoor"`
		UnlockDoor UnlockDoor
	}{
		UnlockDoor: args,
	}

	resp := UnlockDoorResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}
