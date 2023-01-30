package ptz

import (
	"code.byted.org/videoarch/go-onvif/onvif/common"
)

var Namespace = "http://www.onvif.org/ver20/ptz/wsdl"

// NewPTZ creates an initializes a PTZ.
func NewPTZ(endpoint string, cli common.Client) PTZ {
	return &ptz{cli: cli, Endpoint: endpoint}
}

// PTZ was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type PTZ interface {
	OptAbsoluteMove(AbsoluteMove AbsoluteMove) (*AbsoluteMoveResponse, *common.Fault)

	OptContinuousMove(ContinuousMove ContinuousMove) (*ContinuousMoveResponse, *common.Fault)

	OptCreatePresetTour(CreatePresetTour CreatePresetTour) (*CreatePresetTourResponse, *common.Fault)

	OptGeoMove(GeoMove GeoMove) (*GeoMoveResponse, *common.Fault)

	OptGetCompatibleConfigurations(GetCompatibleConfigurations GetCompatibleConfigurations) (*GetCompatibleConfigurationsResponse, *common.Fault)

	OptGetConfiguration(GetConfiguration GetConfiguration) (*GetConfigurationResponse, *common.Fault)

	OptGetConfigurationOptions(GetConfigurationOptions GetConfigurationOptions) (*GetConfigurationOptionsResponse, *common.Fault)

	OptGetConfigurations(GetConfigurations GetConfigurations) (*GetConfigurationsResponse, *common.Fault)

	OptGetNode(GetNode GetNode) (*GetNodeResponse, *common.Fault)

	OptGetNodes(GetNodes GetNodes) (*GetNodesResponse, *common.Fault)

	OptGetPresetTour(GetPresetTour GetPresetTour) (*GetPresetTourResponse, *common.Fault)

	OptGetPresetTourOptions(GetPresetTourOptions GetPresetTourOptions) (*GetPresetTourOptionsResponse, *common.Fault)

	OptGetPresetTours(GetPresetTours GetPresetTours) (*GetPresetToursResponse, *common.Fault)

	OptGetPresets(GetPresets GetPresets) (*GetPresetsResponse, *common.Fault)

	OptGetServiceCapabilities(GetServiceCapabilities GetServiceCapabilities) (*GetServiceCapabilitiesResponse, *common.Fault)

	OptGetStatus(GetStatus GetStatus) (*GetStatusResponse, *common.Fault)

	OptGotoHomePosition(GotoHomePosition GotoHomePosition) (*GotoHomePositionResponse, *common.Fault)

	OptGotoPreset(GotoPreset GotoPreset) (*GotoPresetResponse, *common.Fault)

	OptModifyPresetTour(ModifyPresetTour ModifyPresetTour) (*ModifyPresetTourResponse, *common.Fault)

	OptMoveAndStartTracking(MoveAndStartTracking MoveAndStartTracking) (*MoveAndStartTrackingResponse, *common.Fault)

	OptOperatePresetTour(OperatePresetTour OperatePresetTour) (*OperatePresetTourResponse, *common.Fault)

	OptRelativeMove(RelativeMove RelativeMove) (*RelativeMoveResponse, *common.Fault)

	OptRemovePreset(RemovePreset RemovePreset) (*RemovePresetResponse, *common.Fault)

	OptRemovePresetTour(RemovePresetTour RemovePresetTour) (*RemovePresetTourResponse, *common.Fault)

	OptSendAuxiliaryCommand(SendAuxiliaryCommand SendAuxiliaryCommand) (*SendAuxiliaryCommandResponse, *common.Fault)

	OptSetConfiguration(SetConfiguration SetConfiguration) (*SetConfigurationResponse, *common.Fault)

	OptSetHomePosition(SetHomePosition SetHomePosition) (*SetHomePositionResponse, *common.Fault)

	OptSetPreset(SetPreset SetPreset) (*SetPresetResponse, *common.Fault)

	OptStop(Stop Stop) (*StopResponse, *common.Fault)
}
type Duration string

type AbsoluteMove struct {
	ProfileToken *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
	Position     *common.PTZVector      `xml:"Position,omitempty" json:"Position,omitempty" yaml:"Position,omitempty"`
	Speed        *common.PTZSpeed       `xml:"Speed,omitempty" json:"Speed,omitempty" yaml:"Speed,omitempty"`
}

type AbsoluteMoveResponse struct {
}

type Capabilities []interface{}

type ContinuousMove struct {
	ProfileToken *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
	Velocity     *common.PTZSpeed       `xml:"Velocity,omitempty" json:"Velocity,omitempty" yaml:"Velocity,omitempty"`
	Timeout      *common.Duration       `xml:"Timeout,omitempty" json:"Timeout,omitempty" yaml:"Timeout,omitempty"`
}

type ContinuousMoveResponse struct {
}

type CreatePresetTour struct {
	ProfileToken *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
}

type CreatePresetTourResponse struct {
	PresetTourToken *common.ReferenceToken `xml:"PresetTourToken,omitempty" json:"PresetTourToken,omitempty" yaml:"PresetTourToken,omitempty"`
}

type GeoMove struct {
	ProfileToken *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
	Target       *common.GeoLocation    `xml:"Target,omitempty" json:"Target,omitempty" yaml:"Target,omitempty"`
	Speed        *common.PTZSpeed       `xml:"Speed,omitempty" json:"Speed,omitempty" yaml:"Speed,omitempty"`
	AreaHeight   float64                `xml:"AreaHeight,omitempty" json:"AreaHeight,omitempty" yaml:"AreaHeight,omitempty"`
	AreaWidth    float64                `xml:"AreaWidth,omitempty" json:"AreaWidth,omitempty" yaml:"AreaWidth,omitempty"`
}

type GeoMoveResponse struct {
}

type GetCompatibleConfigurations struct {
	ProfileToken *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
}

type GetCompatibleConfigurationsResponse struct {
	PTZConfiguration []*common.PTZConfiguration `xml:"PTZConfiguration,omitempty" json:"PTZConfiguration,omitempty" yaml:"PTZConfiguration,omitempty"`
}

type GetConfiguration struct {
	PTZConfigurationToken *common.ReferenceToken `xml:"PTZConfigurationToken,omitempty" json:"PTZConfigurationToken,omitempty" yaml:"PTZConfigurationToken,omitempty"`
}

type GetConfigurationOptions struct {
	ConfigurationToken *common.ReferenceToken `xml:"ConfigurationToken,omitempty" json:"ConfigurationToken,omitempty" yaml:"ConfigurationToken,omitempty"`
}

type GetConfigurationOptionsResponse struct {
	PTZConfigurationOptions *common.PTZConfigurationOptions `xml:"PTZConfigurationOptions,omitempty" json:"PTZConfigurationOptions,omitempty" yaml:"PTZConfigurationOptions,omitempty"`
}

type GetConfigurationResponse struct {
	PTZConfiguration *common.PTZConfiguration `xml:"PTZConfiguration,omitempty" json:"PTZConfiguration,omitempty" yaml:"PTZConfiguration,omitempty"`
}

type GetConfigurations struct {
}

type GetConfigurationsResponse struct {
	PTZConfiguration []*common.PTZConfiguration `xml:"PTZConfiguration,omitempty" json:"PTZConfiguration,omitempty" yaml:"PTZConfiguration,omitempty"`
}

type GetNode struct {
	NodeToken *common.ReferenceToken `xml:"NodeToken,omitempty" json:"NodeToken,omitempty" yaml:"NodeToken,omitempty"`
}

type GetNodeResponse struct {
	PTZNode *common.PTZNode `xml:"PTZNode,omitempty" json:"PTZNode,omitempty" yaml:"PTZNode,omitempty"`
}

type GetNodes struct {
}

type GetNodesResponse struct {
	PTZNode []*common.PTZNode `xml:"PTZNode,omitempty" json:"PTZNode,omitempty" yaml:"PTZNode,omitempty"`
}

type GetPresetTour struct {
	ProfileToken    *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
	PresetTourToken *common.ReferenceToken `xml:"PresetTourToken,omitempty" json:"PresetTourToken,omitempty" yaml:"PresetTourToken,omitempty"`
}

type GetPresetTourOptions struct {
	ProfileToken    *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
	PresetTourToken *common.ReferenceToken `xml:"PresetTourToken,omitempty" json:"PresetTourToken,omitempty" yaml:"PresetTourToken,omitempty"`
}

type GetPresetTourOptionsResponse struct {
	Options *common.PTZPresetTourOptions `xml:"Options,omitempty" json:"Options,omitempty" yaml:"Options,omitempty"`
}

type GetPresetTourResponse struct {
	PresetTour *common.PresetTour `xml:"PresetTour,omitempty" json:"PresetTour,omitempty" yaml:"PresetTour,omitempty"`
}

type GetPresetTours struct {
	ProfileToken *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
}

type GetPresetToursResponse struct {
	PresetTour []*common.PresetTour `xml:"PresetTour,omitempty" json:"PresetTour,omitempty" yaml:"PresetTour,omitempty"`
}

type GetPresets struct {
	ProfileToken *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
}

type GetPresetsResponse struct {
	Preset []*common.PTZPreset `xml:"Preset,omitempty" json:"Preset,omitempty" yaml:"Preset,omitempty"`
}

type GetServiceCapabilities struct {
}

type GetServiceCapabilitiesResponse struct {
	Capabilities Capabilities `xml:"Capabilities,omitempty" json:"Capabilities,omitempty" yaml:"Capabilities,omitempty"`
}

type GetStatus struct {
	ProfileToken *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
}

type GetStatusResponse struct {
	PTZStatus *common.PTZStatus `xml:"PTZStatus,omitempty" json:"PTZStatus,omitempty" yaml:"PTZStatus,omitempty"`
}

type GotoHomePosition struct {
	ProfileToken *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
	Speed        *common.PTZSpeed       `xml:"Speed,omitempty" json:"Speed,omitempty" yaml:"Speed,omitempty"`
}

type GotoHomePositionResponse struct {
}

type GotoPreset struct {
	ProfileToken *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
	PresetToken  *common.ReferenceToken `xml:"PresetToken,omitempty" json:"PresetToken,omitempty" yaml:"PresetToken,omitempty"`
	Speed        *common.PTZSpeed       `xml:"Speed,omitempty" json:"Speed,omitempty" yaml:"Speed,omitempty"`
}

type GotoPresetResponse struct {
}

type ModifyPresetTour struct {
	ProfileToken *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
	PresetTour   *common.PresetTour     `xml:"PresetTour,omitempty" json:"PresetTour,omitempty" yaml:"PresetTour,omitempty"`
}

type ModifyPresetTourResponse struct {
}

type MoveAndStartTracking struct {
	ProfileToken   *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
	PresetToken    *common.ReferenceToken `xml:"PresetToken,omitempty" json:"PresetToken,omitempty" yaml:"PresetToken,omitempty"`
	GeoLocation    *common.GeoLocation    `xml:"GeoLocation,omitempty" json:"GeoLocation,omitempty" yaml:"GeoLocation,omitempty"`
	TargetPosition *common.PTZVector      `xml:"TargetPosition,omitempty" json:"TargetPosition,omitempty" yaml:"TargetPosition,omitempty"`
	Speed          *common.PTZSpeed       `xml:"Speed,omitempty" json:"Speed,omitempty" yaml:"Speed,omitempty"`
	ObjectID       int64                  `xml:"ObjectID,omitempty" json:"ObjectID,omitempty" yaml:"ObjectID,omitempty"`
}

type MoveAndStartTrackingResponse struct {
}

type OperatePresetTour struct {
	ProfileToken    *common.ReferenceToken         `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
	PresetTourToken *common.ReferenceToken         `xml:"PresetTourToken,omitempty" json:"PresetTourToken,omitempty" yaml:"PresetTourToken,omitempty"`
	Operation       *common.PTZPresetTourOperation `xml:"Operation,omitempty" json:"Operation,omitempty" yaml:"Operation,omitempty"`
}

type OperatePresetTourResponse struct {
}

type RelativeMove struct {
	ProfileToken *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
	Translation  *common.PTZVector      `xml:"Translation,omitempty" json:"Translation,omitempty" yaml:"Translation,omitempty"`
	Speed        *common.PTZSpeed       `xml:"Speed,omitempty" json:"Speed,omitempty" yaml:"Speed,omitempty"`
}

type RelativeMoveResponse struct {
}

type RemovePreset struct {
	ProfileToken *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
	PresetToken  *common.ReferenceToken `xml:"PresetToken,omitempty" json:"PresetToken,omitempty" yaml:"PresetToken,omitempty"`
}

type RemovePresetResponse struct {
}

type RemovePresetTour struct {
	ProfileToken    *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
	PresetTourToken *common.ReferenceToken `xml:"PresetTourToken,omitempty" json:"PresetTourToken,omitempty" yaml:"PresetTourToken,omitempty"`
}

type RemovePresetTourResponse struct {
}

type SendAuxiliaryCommand struct {
	ProfileToken  *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
	AuxiliaryData *common.AuxiliaryData  `xml:"AuxiliaryData,omitempty" json:"AuxiliaryData,omitempty" yaml:"AuxiliaryData,omitempty"`
}

type SendAuxiliaryCommandResponse struct {
	AuxiliaryResponse *common.AuxiliaryData `xml:"AuxiliaryResponse,omitempty" json:"AuxiliaryResponse,omitempty" yaml:"AuxiliaryResponse,omitempty"`
}

type SetConfiguration struct {
	PTZConfiguration *common.PTZConfiguration `xml:"PTZConfiguration,omitempty" json:"PTZConfiguration,omitempty" yaml:"PTZConfiguration,omitempty"`
	ForcePersistence bool                     `xml:"ForcePersistence,omitempty" json:"ForcePersistence,omitempty" yaml:"ForcePersistence,omitempty"`
}

type SetConfigurationResponse struct {
}

type SetHomePosition struct {
	ProfileToken *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
}

type SetHomePositionResponse struct {
}

type SetPreset struct {
	ProfileToken *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
	PresetName   string                 `xml:"PresetName,omitempty" json:"PresetName,omitempty" yaml:"PresetName,omitempty"`
	PresetToken  *common.ReferenceToken `xml:"PresetToken,omitempty" json:"PresetToken,omitempty" yaml:"PresetToken,omitempty"`
}

type SetPresetResponse struct {
	PresetToken *common.ReferenceToken `xml:"PresetToken,omitempty" json:"PresetToken,omitempty" yaml:"PresetToken,omitempty"`
}

type Stop struct {
	ProfileToken *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
	PanTilt      bool                   `xml:"PanTilt,omitempty" json:"PanTilt,omitempty" yaml:"PanTilt,omitempty"`
	Zoom         bool                   `xml:"Zoom,omitempty" json:"Zoom,omitempty" yaml:"Zoom,omitempty"`
}

type StopResponse struct {
}

// ptz implements the PTZ interface.
type ptz struct {
	cli      common.Client
	Endpoint string
}

func (p *ptz) OptAbsoluteMove(args AbsoluteMove) (*AbsoluteMoveResponse, *common.Fault) {
	req := struct {
		XMLName      string `xml:"tptz:AbsoluteMove"`
		AbsoluteMove AbsoluteMove
	}{
		AbsoluteMove: args,
	}

	resp := AbsoluteMoveResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *ptz) OptContinuousMove(args ContinuousMove) (*ContinuousMoveResponse, *common.Fault) {
	req := struct {
		XMLName        string `xml:"tptz:ContinuousMove"`
		ContinuousMove ContinuousMove
	}{
		ContinuousMove: args,
	}

	resp := ContinuousMoveResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *ptz) OptCreatePresetTour(args CreatePresetTour) (*CreatePresetTourResponse, *common.Fault) {
	req := struct {
		XMLName          string `xml:"tptz:CreatePresetTour"`
		CreatePresetTour CreatePresetTour
	}{
		CreatePresetTour: args,
	}

	resp := CreatePresetTourResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *ptz) OptGeoMove(args GeoMove) (*GeoMoveResponse, *common.Fault) {
	req := struct {
		XMLName string `xml:"tptz:GeoMove"`
		GeoMove GeoMove
	}{
		GeoMove: args,
	}

	resp := GeoMoveResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *ptz) OptGetCompatibleConfigurations(args GetCompatibleConfigurations) (*GetCompatibleConfigurationsResponse, *common.Fault) {
	req := struct {
		XMLName                     string `xml:"tptz:GetCompatibleConfigurations"`
		GetCompatibleConfigurations GetCompatibleConfigurations
	}{
		GetCompatibleConfigurations: args,
	}

	resp := GetCompatibleConfigurationsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *ptz) OptGetConfiguration(args GetConfiguration) (*GetConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName          string `xml:"tptz:GetConfiguration"`
		GetConfiguration GetConfiguration
	}{
		GetConfiguration: args,
	}

	resp := GetConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *ptz) OptGetConfigurationOptions(args GetConfigurationOptions) (*GetConfigurationOptionsResponse, *common.Fault) {
	req := struct {
		XMLName                 string `xml:"tptz:GetConfigurationOptions"`
		GetConfigurationOptions GetConfigurationOptions
	}{
		GetConfigurationOptions: args,
	}

	resp := GetConfigurationOptionsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *ptz) OptGetConfigurations(args GetConfigurations) (*GetConfigurationsResponse, *common.Fault) {
	req := struct {
		XMLName           string `xml:"tptz:GetConfigurations"`
		GetConfigurations GetConfigurations
	}{
		GetConfigurations: args,
	}

	resp := GetConfigurationsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *ptz) OptGetNode(args GetNode) (*GetNodeResponse, *common.Fault) {
	req := struct {
		XMLName string `xml:"tptz:GetNode"`
		GetNode GetNode
	}{
		GetNode: args,
	}

	resp := GetNodeResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *ptz) OptGetNodes(args GetNodes) (*GetNodesResponse, *common.Fault) {
	req := struct {
		XMLName  string `xml:"tptz:GetNodes"`
		GetNodes GetNodes
	}{
		GetNodes: args,
	}

	resp := GetNodesResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *ptz) OptGetPresetTour(args GetPresetTour) (*GetPresetTourResponse, *common.Fault) {
	req := struct {
		XMLName       string `xml:"tptz:GetPresetTour"`
		GetPresetTour GetPresetTour
	}{
		GetPresetTour: args,
	}

	resp := GetPresetTourResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *ptz) OptGetPresetTourOptions(args GetPresetTourOptions) (*GetPresetTourOptionsResponse, *common.Fault) {
	req := struct {
		XMLName              string `xml:"tptz:GetPresetTourOptions"`
		GetPresetTourOptions GetPresetTourOptions
	}{
		GetPresetTourOptions: args,
	}

	resp := GetPresetTourOptionsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *ptz) OptGetPresetTours(args GetPresetTours) (*GetPresetToursResponse, *common.Fault) {
	req := struct {
		XMLName        string `xml:"tptz:GetPresetTours"`
		GetPresetTours GetPresetTours
	}{
		GetPresetTours: args,
	}

	resp := GetPresetToursResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *ptz) OptGetPresets(args GetPresets) (*GetPresetsResponse, *common.Fault) {
	req := struct {
		XMLName    string `xml:"tptz:GetPresets"`
		GetPresets GetPresets
	}{
		GetPresets: args,
	}

	resp := GetPresetsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *ptz) OptGetServiceCapabilities(args GetServiceCapabilities) (*GetServiceCapabilitiesResponse, *common.Fault) {
	req := struct {
		XMLName                string `xml:"tptz:GetServiceCapabilities"`
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

func (p *ptz) OptGetStatus(args GetStatus) (*GetStatusResponse, *common.Fault) {
	req := struct {
		XMLName   string `xml:"tptz:GetStatus"`
		GetStatus GetStatus
	}{
		GetStatus: args,
	}

	resp := GetStatusResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *ptz) OptGotoHomePosition(args GotoHomePosition) (*GotoHomePositionResponse, *common.Fault) {
	req := struct {
		XMLName          string `xml:"tptz:GotoHomePosition"`
		GotoHomePosition GotoHomePosition
	}{
		GotoHomePosition: args,
	}

	resp := GotoHomePositionResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *ptz) OptGotoPreset(args GotoPreset) (*GotoPresetResponse, *common.Fault) {
	req := struct {
		XMLName    string `xml:"tptz:GotoPreset"`
		GotoPreset GotoPreset
	}{
		GotoPreset: args,
	}

	resp := GotoPresetResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *ptz) OptModifyPresetTour(args ModifyPresetTour) (*ModifyPresetTourResponse, *common.Fault) {
	req := struct {
		XMLName          string `xml:"tptz:ModifyPresetTour"`
		ModifyPresetTour ModifyPresetTour
	}{
		ModifyPresetTour: args,
	}

	resp := ModifyPresetTourResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *ptz) OptMoveAndStartTracking(args MoveAndStartTracking) (*MoveAndStartTrackingResponse, *common.Fault) {
	req := struct {
		XMLName              string `xml:"tptz:MoveAndStartTracking"`
		MoveAndStartTracking MoveAndStartTracking
	}{
		MoveAndStartTracking: args,
	}

	resp := MoveAndStartTrackingResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *ptz) OptOperatePresetTour(args OperatePresetTour) (*OperatePresetTourResponse, *common.Fault) {
	req := struct {
		XMLName           string `xml:"tptz:OperatePresetTour"`
		OperatePresetTour OperatePresetTour
	}{
		OperatePresetTour: args,
	}

	resp := OperatePresetTourResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *ptz) OptRelativeMove(args RelativeMove) (*RelativeMoveResponse, *common.Fault) {
	req := struct {
		XMLName      string `xml:"tptz:RelativeMove"`
		RelativeMove RelativeMove
	}{
		RelativeMove: args,
	}

	resp := RelativeMoveResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *ptz) OptRemovePreset(args RemovePreset) (*RemovePresetResponse, *common.Fault) {
	req := struct {
		XMLName      string `xml:"tptz:RemovePreset"`
		RemovePreset RemovePreset
	}{
		RemovePreset: args,
	}

	resp := RemovePresetResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *ptz) OptRemovePresetTour(args RemovePresetTour) (*RemovePresetTourResponse, *common.Fault) {
	req := struct {
		XMLName          string `xml:"tptz:RemovePresetTour"`
		RemovePresetTour RemovePresetTour
	}{
		RemovePresetTour: args,
	}

	resp := RemovePresetTourResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *ptz) OptSendAuxiliaryCommand(args SendAuxiliaryCommand) (*SendAuxiliaryCommandResponse, *common.Fault) {
	req := struct {
		XMLName              string `xml:"tptz:SendAuxiliaryCommand"`
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

func (p *ptz) OptSetConfiguration(args SetConfiguration) (*SetConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName          string `xml:"tptz:SetConfiguration"`
		SetConfiguration SetConfiguration
	}{
		SetConfiguration: args,
	}

	resp := SetConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *ptz) OptSetHomePosition(args SetHomePosition) (*SetHomePositionResponse, *common.Fault) {
	req := struct {
		XMLName         string `xml:"tptz:SetHomePosition"`
		SetHomePosition SetHomePosition
	}{
		SetHomePosition: args,
	}

	resp := SetHomePositionResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *ptz) OptSetPreset(args SetPreset) (*SetPresetResponse, *common.Fault) {
	req := struct {
		XMLName   string `xml:"tptz:SetPreset"`
		SetPreset SetPreset
	}{
		SetPreset: args,
	}

	resp := SetPresetResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *ptz) OptStop(args Stop) (*StopResponse, *common.Fault) {
	req := struct {
		XMLName string `xml:"tptz:Stop"`
		Stop    Stop
	}{
		Stop: args,
	}

	resp := StopResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}
