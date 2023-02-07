package analytics

import (
	"github.com/zijiemeng/go-onvif/onvif/common"
)

var Namespace = "http://www.onvif.org/ver20/analytics/wsdl"

// NewAnalyticsEnginePort creates an initializes a AnalyticsEnginePort.
func NewAnalyticsEnginePort(endpoint string, cli common.Client) AnalyticsEnginePort {
	return &analyticsEnginePort{cli: cli, Endpoint: endpoint}
}

// AnalyticsEnginePort was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type AnalyticsEnginePort interface {
	OptCreateAnalyticsModules(CreateAnalyticsModules CreateAnalyticsModules) (*CreateAnalyticsModulesResponse, *common.Fault)

	OptCreateRules(CreateRules CreateRules) (*CreateRulesResponse, *common.Fault)

	OptDeleteAnalyticsModules(DeleteAnalyticsModules DeleteAnalyticsModules) (*DeleteAnalyticsModulesResponse, *common.Fault)

	OptDeleteRules(DeleteRules DeleteRules) (*DeleteRulesResponse, *common.Fault)

	OptGetAnalyticsModuleOptions(GetAnalyticsModuleOptions GetAnalyticsModuleOptions) (*GetAnalyticsModuleOptionsResponse, *common.Fault)

	OptGetAnalyticsModules(GetAnalyticsModules GetAnalyticsModules) (*GetAnalyticsModulesResponse, *common.Fault)

	OptGetRuleOptions(GetRuleOptions GetRuleOptions) (*GetRuleOptionsResponse, *common.Fault)

	OptGetRules(GetRules GetRules) (*GetRulesResponse, *common.Fault)

	OptGetServiceCapabilities(GetServiceCapabilities GetServiceCapabilities) (*GetServiceCapabilitiesResponse, *common.Fault)

	OptGetSupportedAnalyticsModules(GetSupportedAnalyticsModules GetSupportedAnalyticsModules) (*GetSupportedAnalyticsModulesResponse, *common.Fault)

	OptGetSupportedMetadata(GetSupportedMetadata GetSupportedMetadata) (*GetSupportedMetadataResponse, *common.Fault)

	OptGetSupportedRules(GetSupportedRules GetSupportedRules) (*GetSupportedRulesResponse, *common.Fault)

	OptModifyAnalyticsModules(ModifyAnalyticsModules ModifyAnalyticsModules) (*ModifyAnalyticsModulesResponse, *common.Fault)

	OptModifyRules(ModifyRules ModifyRules) (*ModifyRulesResponse, *common.Fault)
}
type DateTime string

type Capabilities []interface{}

type ConfigOptions []interface{}

type CreateAnalyticsModules struct {
	ConfigurationToken *common.ReferenceToken `xml:"ConfigurationToken,omitempty" json:"ConfigurationToken,omitempty" yaml:"ConfigurationToken,omitempty"`
	AnalyticsModule    []*common.Config       `xml:"AnalyticsModule,omitempty" json:"AnalyticsModule,omitempty" yaml:"AnalyticsModule,omitempty"`
}

type CreateAnalyticsModulesResponse struct {
}

type CreateRules struct {
	ConfigurationToken *common.ReferenceToken `xml:"ConfigurationToken,omitempty" json:"ConfigurationToken,omitempty" yaml:"ConfigurationToken,omitempty"`
	Rule               []common.Config        `xml:"Rule" json:"Rule" yaml:"Rule"`
}

type CreateRulesResponse struct {
}

type DeleteAnalyticsModules struct {
	ConfigurationToken  *common.ReferenceToken `xml:"ConfigurationToken,omitempty" json:"ConfigurationToken,omitempty" yaml:"ConfigurationToken,omitempty"`
	AnalyticsModuleName []string               `xml:"AnalyticsModuleName,omitempty" json:"AnalyticsModuleName,omitempty" yaml:"AnalyticsModuleName,omitempty"`
}

type DeleteAnalyticsModulesResponse struct {
}

type DeleteRules struct {
	ConfigurationToken *common.ReferenceToken `xml:"ConfigurationToken,omitempty" json:"ConfigurationToken,omitempty" yaml:"ConfigurationToken,omitempty"`
	RuleName           []string               `xml:"RuleName" json:"RuleName" yaml:"RuleName"`
}

type DeleteRulesResponse struct {
}

type Frame struct {
	PTZStatus      *common.PTZStatus      `xml:"PTZStatus,omitempty" json:"PTZStatus,omitempty" yaml:"PTZStatus,omitempty"`
	Transformation *common.Transformation `xml:"Transformation,omitempty" json:"Transformation,omitempty" yaml:"Transformation,omitempty"`
	Object         []Object               `xml:"Object,omitempty" json:"Object,omitempty" yaml:"Object,omitempty"`
	ObjectTree     ObjectTree             `xml:"ObjectTree,omitempty" json:"ObjectTree,omitempty" yaml:"ObjectTree,omitempty"`
	Extension      FrameExtension         `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
	SceneImageRef  string                 `xml:"SceneImageRef,omitempty" json:"SceneImageRef,omitempty" yaml:"SceneImageRef,omitempty"`
	SceneImage     []byte                 `xml:"SceneImage,omitempty" json:"SceneImage,omitempty" yaml:"SceneImage,omitempty"`
	UtcTime        common.DateTime        `xml:"UtcTime,attr,omitempty" json:"UtcTime,attr,omitempty" yaml:"UtcTime,attr,omitempty"`
	Colorspace     string                 `xml:"Colorspace,attr,omitempty" json:"Colorspace,attr,omitempty" yaml:"Colorspace,attr,omitempty"`
	Source         string                 `xml:"Source,attr,omitempty" json:"Source,attr,omitempty" yaml:"Source,attr,omitempty"`
}

type FrameExtension struct {
	MotionInCells *common.MotionInCells   `xml:"MotionInCells,omitempty" json:"MotionInCells,omitempty" yaml:"MotionInCells,omitempty"`
	Extension     *common.FrameExtension2 `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type GetAnalyticsModuleOptions struct {
	Type               string                 `xml:"Type,omitempty" json:"Type,omitempty" yaml:"Type,omitempty"`
	ConfigurationToken *common.ReferenceToken `xml:"ConfigurationToken,omitempty" json:"ConfigurationToken,omitempty" yaml:"ConfigurationToken,omitempty"`
}

type GetAnalyticsModuleOptionsResponse struct {
	Options []ConfigOptions `xml:"Options,omitempty" json:"Options,omitempty" yaml:"Options,omitempty"`
}

type GetAnalyticsModules struct {
	ConfigurationToken *common.ReferenceToken `xml:"ConfigurationToken,omitempty" json:"ConfigurationToken,omitempty" yaml:"ConfigurationToken,omitempty"`
}

type GetAnalyticsModulesResponse struct {
	AnalyticsModule []*common.Config `xml:"AnalyticsModule,omitempty" json:"AnalyticsModule,omitempty" yaml:"AnalyticsModule,omitempty"`
}

type GetRuleOptions struct {
	RuleType           string                 `xml:"RuleType,omitempty" json:"RuleType,omitempty" yaml:"RuleType,omitempty"`
	ConfigurationToken *common.ReferenceToken `xml:"ConfigurationToken,omitempty" json:"ConfigurationToken,omitempty" yaml:"ConfigurationToken,omitempty"`
}

type GetRuleOptionsResponse struct {
	RuleOptions []ConfigOptions `xml:"RuleOptions,omitempty" json:"RuleOptions,omitempty" yaml:"RuleOptions,omitempty"`
}

type GetRules struct {
	ConfigurationToken *common.ReferenceToken `xml:"ConfigurationToken,omitempty" json:"ConfigurationToken,omitempty" yaml:"ConfigurationToken,omitempty"`
}

type GetRulesResponse struct {
	Rule []*common.Config `xml:"Rule,omitempty" json:"Rule,omitempty" yaml:"Rule,omitempty"`
}

type GetServiceCapabilities struct {
}

type GetServiceCapabilitiesResponse struct {
	Capabilities Capabilities `xml:"Capabilities,omitempty" json:"Capabilities,omitempty" yaml:"Capabilities,omitempty"`
}

type GetSupportedAnalyticsModules struct {
	ConfigurationToken *common.ReferenceToken `xml:"ConfigurationToken,omitempty" json:"ConfigurationToken,omitempty" yaml:"ConfigurationToken,omitempty"`
}

type GetSupportedAnalyticsModulesResponse struct {
	SupportedAnalyticsModules *common.SupportedAnalyticsModules `xml:"SupportedAnalyticsModules,omitempty" json:"SupportedAnalyticsModules,omitempty" yaml:"SupportedAnalyticsModules,omitempty"`
}

type GetSupportedMetadata struct {
	Type string `xml:"Type,omitempty" json:"Type,omitempty" yaml:"Type,omitempty"`
}

type GetSupportedMetadataResponse struct {
	AnalyticsModule []MetadataInfo `xml:"AnalyticsModule,omitempty" json:"AnalyticsModule,omitempty" yaml:"AnalyticsModule,omitempty"`
}

type GetSupportedRules struct {
	ConfigurationToken *common.ReferenceToken `xml:"ConfigurationToken,omitempty" json:"ConfigurationToken,omitempty" yaml:"ConfigurationToken,omitempty"`
}

type GetSupportedRulesResponse struct {
	SupportedRules *common.SupportedRules `xml:"SupportedRules,omitempty" json:"SupportedRules,omitempty" yaml:"SupportedRules,omitempty"`
}

type MetadataInfo struct {
	SampleFrame Frame  `xml:"SampleFrame,omitempty" json:"SampleFrame,omitempty" yaml:"SampleFrame,omitempty"`
	Type        string `xml:"Type,attr,omitempty" json:"Type,attr,omitempty" yaml:"Type,attr,omitempty"`
}

type ModifyAnalyticsModules struct {
	ConfigurationToken *common.ReferenceToken `xml:"ConfigurationToken,omitempty" json:"ConfigurationToken,omitempty" yaml:"ConfigurationToken,omitempty"`
	AnalyticsModule    []*common.Config       `xml:"AnalyticsModule,omitempty" json:"AnalyticsModule,omitempty" yaml:"AnalyticsModule,omitempty"`
}

type ModifyAnalyticsModulesResponse struct {
}

type ModifyRules struct {
	ConfigurationToken *common.ReferenceToken `xml:"ConfigurationToken,omitempty" json:"ConfigurationToken,omitempty" yaml:"ConfigurationToken,omitempty"`
	Rule               []common.Config        `xml:"Rule" json:"Rule" yaml:"Rule"`
}

type ModifyRulesResponse struct {
}

type Object struct {
	Parent        int64              `xml:"Parent,attr,omitempty" json:"Parent,attr,omitempty" yaml:"Parent,attr,omitempty"`
	Appearance    *common.Appearance `xml:"Appearance,omitempty" json:"Appearance,omitempty" yaml:"Appearance,omitempty"`
	Behaviour     *common.Behaviour  `xml:"Behaviour,omitempty" json:"Behaviour,omitempty" yaml:"Behaviour,omitempty"`
	Extension     ObjectExtension    `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
	TypeAttrXSI   string             `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string             `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *Object) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:Object"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://www.onvif.org/ver20/analytics/wsdl"
	}
}

type ObjectExtension []interface{}

type ObjectTree struct {
	Rename    []*common.Rename            `xml:"Rename,omitempty" json:"Rename,omitempty" yaml:"Rename,omitempty"`
	Split     []*common.Split             `xml:"Split,omitempty" json:"Split,omitempty" yaml:"Split,omitempty"`
	Merge     []*common.Merge             `xml:"Merge,omitempty" json:"Merge,omitempty" yaml:"Merge,omitempty"`
	Delete    []*common.ObjectId          `xml:"Delete,omitempty" json:"Delete,omitempty" yaml:"Delete,omitempty"`
	Extension *common.ObjectTreeExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

// analyticsEnginePort implements the AnalyticsEnginePort interface.
type analyticsEnginePort struct {
	cli      common.Client
	Endpoint string
}

func (p *analyticsEnginePort) OptCreateAnalyticsModules(args CreateAnalyticsModules) (*CreateAnalyticsModulesResponse, *common.Fault) {
	req := struct {
		XMLName                string `xml:"tan:CreateAnalyticsModules"`
		CreateAnalyticsModules CreateAnalyticsModules
	}{
		CreateAnalyticsModules: args,
	}

	resp := CreateAnalyticsModulesResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *analyticsEnginePort) OptCreateRules(args CreateRules) (*CreateRulesResponse, *common.Fault) {
	req := struct {
		XMLName     string `xml:"tan:CreateRules"`
		CreateRules CreateRules
	}{
		CreateRules: args,
	}

	resp := CreateRulesResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *analyticsEnginePort) OptDeleteAnalyticsModules(args DeleteAnalyticsModules) (*DeleteAnalyticsModulesResponse, *common.Fault) {
	req := struct {
		XMLName                string `xml:"tan:DeleteAnalyticsModules"`
		DeleteAnalyticsModules DeleteAnalyticsModules
	}{
		DeleteAnalyticsModules: args,
	}

	resp := DeleteAnalyticsModulesResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *analyticsEnginePort) OptDeleteRules(args DeleteRules) (*DeleteRulesResponse, *common.Fault) {
	req := struct {
		XMLName     string `xml:"tan:DeleteRules"`
		DeleteRules DeleteRules
	}{
		DeleteRules: args,
	}

	resp := DeleteRulesResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *analyticsEnginePort) OptGetAnalyticsModuleOptions(args GetAnalyticsModuleOptions) (*GetAnalyticsModuleOptionsResponse, *common.Fault) {
	req := struct {
		XMLName                   string `xml:"tan:GetAnalyticsModuleOptions"`
		GetAnalyticsModuleOptions GetAnalyticsModuleOptions
	}{
		GetAnalyticsModuleOptions: args,
	}

	resp := GetAnalyticsModuleOptionsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *analyticsEnginePort) OptGetAnalyticsModules(args GetAnalyticsModules) (*GetAnalyticsModulesResponse, *common.Fault) {
	req := struct {
		XMLName             string `xml:"tan:GetAnalyticsModules"`
		GetAnalyticsModules GetAnalyticsModules
	}{
		GetAnalyticsModules: args,
	}

	resp := GetAnalyticsModulesResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *analyticsEnginePort) OptGetRuleOptions(args GetRuleOptions) (*GetRuleOptionsResponse, *common.Fault) {
	req := struct {
		XMLName        string `xml:"tan:GetRuleOptions"`
		GetRuleOptions GetRuleOptions
	}{
		GetRuleOptions: args,
	}

	resp := GetRuleOptionsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *analyticsEnginePort) OptGetRules(args GetRules) (*GetRulesResponse, *common.Fault) {
	req := struct {
		XMLName  string `xml:"tan:GetRules"`
		GetRules GetRules
	}{
		GetRules: args,
	}

	resp := GetRulesResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *analyticsEnginePort) OptGetServiceCapabilities(args GetServiceCapabilities) (*GetServiceCapabilitiesResponse, *common.Fault) {
	req := struct {
		XMLName                string `xml:"tan:GetServiceCapabilities"`
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

func (p *analyticsEnginePort) OptGetSupportedAnalyticsModules(args GetSupportedAnalyticsModules) (*GetSupportedAnalyticsModulesResponse, *common.Fault) {
	req := struct {
		XMLName                      string `xml:"tan:GetSupportedAnalyticsModules"`
		GetSupportedAnalyticsModules GetSupportedAnalyticsModules
	}{
		GetSupportedAnalyticsModules: args,
	}

	resp := GetSupportedAnalyticsModulesResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *analyticsEnginePort) OptGetSupportedMetadata(args GetSupportedMetadata) (*GetSupportedMetadataResponse, *common.Fault) {
	req := struct {
		XMLName              string `xml:"tan:GetSupportedMetadata"`
		GetSupportedMetadata GetSupportedMetadata
	}{
		GetSupportedMetadata: args,
	}

	resp := GetSupportedMetadataResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *analyticsEnginePort) OptGetSupportedRules(args GetSupportedRules) (*GetSupportedRulesResponse, *common.Fault) {
	req := struct {
		XMLName           string `xml:"tan:GetSupportedRules"`
		GetSupportedRules GetSupportedRules
	}{
		GetSupportedRules: args,
	}

	resp := GetSupportedRulesResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *analyticsEnginePort) OptModifyAnalyticsModules(args ModifyAnalyticsModules) (*ModifyAnalyticsModulesResponse, *common.Fault) {
	req := struct {
		XMLName                string `xml:"tan:ModifyAnalyticsModules"`
		ModifyAnalyticsModules ModifyAnalyticsModules
	}{
		ModifyAnalyticsModules: args,
	}

	resp := ModifyAnalyticsModulesResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *analyticsEnginePort) OptModifyRules(args ModifyRules) (*ModifyRulesResponse, *common.Fault) {
	req := struct {
		XMLName     string `xml:"tan:ModifyRules"`
		ModifyRules ModifyRules
	}{
		ModifyRules: args,
	}

	resp := ModifyRulesResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}
