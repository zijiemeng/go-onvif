package federated_search

import (
	"code.byted.org/videoarch/go-onvif/onvif/common"
)

var Namespace = "http://www.onvif.org/ver10/federatedsearch/wsdl"

// NewFederatedSearchPort creates an initializes a FederatedSearchPort.
func NewFederatedSearchPort(endpoint string, cli common.Client) FederatedSearchPort {
	return &federatedSearchPort{cli: cli, Endpoint: endpoint}
}

// FederatedSearchPort was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type FederatedSearchPort interface {
	OptGetSearchResults(GetSearchResults GetSearchResults) (*GetSearchResultsResponse, error)

	OptGetServiceCapabilities(GetServiceCapabilities GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error)

	OptGetServiceFeatures(GetServiceFeatures GetServiceFeatures) (*GetServiceFeaturesResponse, error)

	OptRegisterDatabase(RegisterDatabase RegisterDatabase) (*RegisterDatabaseResponse, error)

	OptSearch(Search Search) (*SearchResponse, error)
}
type DateTime string

type Duration string

type SimpleTermType string

type MediaDurationType string

type MimeType string

type StartTimePointType string

type XPathType string

type ZeroToOneType float64

type IDREF common.NCName

type IDREFS []IDREF

type AND struct {
	PreferenceValue ZeroToOneType           `xml:"preferenceValue,attr,omitempty" json:"preferenceValue,attr,omitempty" yaml:"preferenceValue,attr,omitempty"`
	ThresholdValue  ZeroToOneType           `xml:"thresholdValue,attr,omitempty" json:"thresholdValue,attr,omitempty" yaml:"thresholdValue,attr,omitempty"`
	ScoringFunction SimpleTermType          `xml:"scoringFunction,attr,omitempty" json:"scoringFunction,attr,omitempty" yaml:"scoringFunction,attr,omitempty"`
	Condition       []BooleanExpressionType `xml:"Condition" json:"Condition" yaml:"Condition"`
	TypeAttrXSI     string                  `xml:"xsi:type,attr,omitempty"`
	TypeNamespace   string                  `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *AND) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:AND"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "urn:mpeg:mpqf:schema:2008"
	}
}

type AVG struct {
	AggregateID   string    `xml:"aggregateID,attr,omitempty" json:"aggregateID,attr,omitempty" yaml:"aggregateID,attr,omitempty"`
	Field         FieldType `xml:"Field,omitempty" json:"Field,omitempty" yaml:"Field,omitempty"`
	TypeAttrXSI   string    `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string    `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *AVG) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:AVG"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "urn:mpeg:mpqf:schema:2008"
	}
}

type Abs struct {
	DoubleValue          float64                  `xml:"DoubleValue,omitempty" json:"DoubleValue,omitempty" yaml:"DoubleValue,omitempty"`
	LongValue            int64                    `xml:"LongValue,omitempty" json:"LongValue,omitempty" yaml:"LongValue,omitempty"`
	ArithmeticField      FieldType                `xml:"ArithmeticField,omitempty" json:"ArithmeticField,omitempty" yaml:"ArithmeticField,omitempty"`
	ArithmeticExpression ArithmeticExpressionType `xml:"ArithmeticExpression,omitempty" json:"ArithmeticExpression,omitempty" yaml:"ArithmeticExpression,omitempty"`
	TypeAttrXSI          string                   `xml:"xsi:type,attr,omitempty"`
	TypeNamespace        string                   `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *Abs) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:Abs"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "urn:mpeg:mpqf:schema:2008"
	}
}

type AbstractSortByType interface{}

type Add struct {
	DoubleValue          float64                  `xml:"DoubleValue,omitempty" json:"DoubleValue,omitempty" yaml:"DoubleValue,omitempty"`
	LongValue            int64                    `xml:"LongValue,omitempty" json:"LongValue,omitempty" yaml:"LongValue,omitempty"`
	ArithmeticField      FieldType                `xml:"ArithmeticField,omitempty" json:"ArithmeticField,omitempty" yaml:"ArithmeticField,omitempty"`
	ArithmeticExpression ArithmeticExpressionType `xml:"ArithmeticExpression,omitempty" json:"ArithmeticExpression,omitempty" yaml:"ArithmeticExpression,omitempty"`
	TypeAttrXSI          string                   `xml:"xsi:type,attr,omitempty"`
	TypeNamespace        string                   `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *Add) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:Add"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "urn:mpeg:mpqf:schema:2008"
	}
}

type AggregateExpressionType interface{}

type ArithmeticExpressionType interface{}

type AvailableCapabilityType struct {
	SupportedQFProfile         CapabilityTermType   `xml:"SupportedQFProfile,omitempty" json:"SupportedQFProfile,omitempty" yaml:"SupportedQFProfile,omitempty"`
	SupportedMetadata          []string             `xml:"SupportedMetadata,omitempty" json:"SupportedMetadata,omitempty" yaml:"SupportedMetadata,omitempty"`
	SupportedExampleMediaTypes string               `xml:"SupportedExampleMediaTypes,omitempty" json:"SupportedExampleMediaTypes,omitempty" yaml:"SupportedExampleMediaTypes,omitempty"`
	SupportedResultMediaTypes  string               `xml:"SupportedResultMediaTypes,omitempty" json:"SupportedResultMediaTypes,omitempty" yaml:"SupportedResultMediaTypes,omitempty"`
	SupportedQueryTypes        []CapabilityTermType `xml:"SupportedQueryTypes,omitempty" json:"SupportedQueryTypes,omitempty" yaml:"SupportedQueryTypes,omitempty"`
	SupportedExpressions       []CapabilityTermType `xml:"SupportedExpressions,omitempty" json:"SupportedExpressions,omitempty" yaml:"SupportedExpressions,omitempty"`
	UsageConditions            []string             `xml:"UsageConditions,omitempty" json:"UsageConditions,omitempty" yaml:"UsageConditions,omitempty"`
	ServiceID                  string               `xml:"serviceID,attr,omitempty" json:"serviceID,attr,omitempty" yaml:"serviceID,attr,omitempty"`
	TypeAttrXSI                string               `xml:"xsi:type,attr,omitempty"`
	TypeNamespace              string               `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *AvailableCapabilityType) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:AvailableCapabilityType"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "urn:mpeg:mpqf:schema:2008"
	}
}

type BooleanExpressionType interface{}

type Capabilities []interface{}

type CapabilityTermType struct {
	Name          string         `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Description   string         `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	Term          []TermType     `xml:"Term,omitempty" json:"Term,omitempty" yaml:"Term,omitempty"`
	Href          SimpleTermType `xml:"href,attr,omitempty" json:"href,attr,omitempty" yaml:"href,attr,omitempty"`
	UsageRefList  IDREFS         `xml:"usageRefList,attr,omitempty" json:"usageRefList,attr,omitempty" yaml:"usageRefList,attr,omitempty"`
	TypeAttrXSI   string         `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string         `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *CapabilityTermType) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:CapabilityTermType"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "urn:mpeg:mpqf:schema:2008"
	}
}

type CapabilityType struct {
	SupportedQFProfile         CapabilityTermType   `xml:"SupportedQFProfile,omitempty" json:"SupportedQFProfile,omitempty" yaml:"SupportedQFProfile,omitempty"`
	SupportedMetadata          []string             `xml:"SupportedMetadata,omitempty" json:"SupportedMetadata,omitempty" yaml:"SupportedMetadata,omitempty"`
	SupportedExampleMediaTypes string               `xml:"SupportedExampleMediaTypes,omitempty" json:"SupportedExampleMediaTypes,omitempty" yaml:"SupportedExampleMediaTypes,omitempty"`
	SupportedResultMediaTypes  string               `xml:"SupportedResultMediaTypes,omitempty" json:"SupportedResultMediaTypes,omitempty" yaml:"SupportedResultMediaTypes,omitempty"`
	SupportedQueryTypes        []CapabilityTermType `xml:"SupportedQueryTypes,omitempty" json:"SupportedQueryTypes,omitempty" yaml:"SupportedQueryTypes,omitempty"`
	SupportedExpressions       []CapabilityTermType `xml:"SupportedExpressions,omitempty" json:"SupportedExpressions,omitempty" yaml:"SupportedExpressions,omitempty"`
	UsageConditions            []string             `xml:"UsageConditions,omitempty" json:"UsageConditions,omitempty" yaml:"UsageConditions,omitempty"`
}

type Ceiling struct {
	DoubleValue          float64                  `xml:"DoubleValue,omitempty" json:"DoubleValue,omitempty" yaml:"DoubleValue,omitempty"`
	LongValue            int64                    `xml:"LongValue,omitempty" json:"LongValue,omitempty" yaml:"LongValue,omitempty"`
	ArithmeticField      FieldType                `xml:"ArithmeticField,omitempty" json:"ArithmeticField,omitempty" yaml:"ArithmeticField,omitempty"`
	ArithmeticExpression ArithmeticExpressionType `xml:"ArithmeticExpression,omitempty" json:"ArithmeticExpression,omitempty" yaml:"ArithmeticExpression,omitempty"`
	TypeAttrXSI          string                   `xml:"xsi:type,attr,omitempty"`
	TypeNamespace        string                   `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *Ceiling) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:Ceiling"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "urn:mpeg:mpqf:schema:2008"
	}
}

type ClassificationScheme struct {
	Term   []TermType `xml:"Term,omitempty" json:"Term,omitempty" yaml:"Term,omitempty"`
	Uri    string     `xml:"uri,attr,omitempty" json:"uri,attr,omitempty" yaml:"uri,attr,omitempty"`
	Domain string     `xml:"domain,attr,omitempty" json:"domain,attr,omitempty" yaml:"domain,attr,omitempty"`
}

type ComparisonExpressionType interface{}

type ComplementOf struct {
	PreferenceValue ZeroToOneType `xml:"preferenceValue,attr,omitempty" json:"preferenceValue,attr,omitempty" yaml:"preferenceValue,attr,omitempty"`
	ThresholdValue  ZeroToOneType `xml:"thresholdValue,attr,omitempty" json:"thresholdValue,attr,omitempty" yaml:"thresholdValue,attr,omitempty"`
	AnchorDistance  uint          `xml:"anchorDistance,attr,omitempty" json:"anchorDistance,attr,omitempty" yaml:"anchorDistance,attr,omitempty"`
	Anchor          bool          `xml:"anchor,attr,omitempty" json:"anchor,attr,omitempty" yaml:"anchor,attr,omitempty"`
	Var             string        `xml:"var,attr,omitempty" json:"var,attr,omitempty" yaml:"var,attr,omitempty"`
	Class           string        `xml:"class,attr,omitempty" json:"class,attr,omitempty" yaml:"class,attr,omitempty"`
	TypeAttrXSI     string        `xml:"xsi:type,attr,omitempty"`
	TypeNamespace   string        `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *ComplementOf) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:ComplementOf"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "urn:mpeg:mpqf:schema:2008"
	}
}

type Contains struct {
	PreferenceValue     ZeroToOneType        `xml:"preferenceValue,attr,omitempty" json:"preferenceValue,attr,omitempty" yaml:"preferenceValue,attr,omitempty"`
	ThresholdValue      ZeroToOneType        `xml:"thresholdValue,attr,omitempty" json:"thresholdValue,attr,omitempty" yaml:"thresholdValue,attr,omitempty"`
	StringValue         string               `xml:"StringValue,omitempty" json:"StringValue,omitempty" yaml:"StringValue,omitempty"`
	StringField         FieldType            `xml:"StringField,omitempty" json:"StringField,omitempty" yaml:"StringField,omitempty"`
	StringExpression    StringExpressionType `xml:"StringExpression,omitempty" json:"StringExpression,omitempty" yaml:"StringExpression,omitempty"`
	SemanticStringField SemanticFieldType    `xml:"SemanticStringField,omitempty" json:"SemanticStringField,omitempty" yaml:"SemanticStringField,omitempty"`
	TypeAttrXSI         string               `xml:"xsi:type,attr,omitempty"`
	TypeNamespace       string               `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *Contains) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:Contains"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "urn:mpeg:mpqf:schema:2008"
	}
}

type Count struct {
	AggregateID   string    `xml:"aggregateID,attr,omitempty" json:"aggregateID,attr,omitempty" yaml:"aggregateID,attr,omitempty"`
	Field         FieldType `xml:"Field,omitempty" json:"Field,omitempty" yaml:"Field,omitempty"`
	TypeAttrXSI   string    `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string    `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *Count) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:Count"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "urn:mpeg:mpqf:schema:2008"
	}
}

type DeclaredFieldType struct {
	Content  *string `xml:"Content,omitempty" json:"Content,omitempty" yaml:"Content,omitempty"`
	Id       string  `xml:"id,attr,omitempty" json:"id,attr,omitempty" yaml:"id,attr,omitempty"`
	TypeName string  `xml:"typeName,attr,omitempty" json:"typeName,attr,omitempty" yaml:"typeName,attr,omitempty"`
}

type DescriptionResourceType struct {
	ResourceID     string `xml:"resourceID,attr,omitempty" json:"resourceID,attr,omitempty" yaml:"resourceID,attr,omitempty"`
	AnyDescription string `xml:"AnyDescription,omitempty" json:"AnyDescription,omitempty" yaml:"AnyDescription,omitempty"`
	TypeAttrXSI    string `xml:"xsi:type,attr,omitempty"`
	TypeNamespace  string `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *DescriptionResourceType) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:DescriptionResourceType"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "urn:mpeg:mpqf:schema:2008"
	}
}

type DisjointWith struct {
	PreferenceValue ZeroToOneType `xml:"preferenceValue,attr,omitempty" json:"preferenceValue,attr,omitempty" yaml:"preferenceValue,attr,omitempty"`
	ThresholdValue  ZeroToOneType `xml:"thresholdValue,attr,omitempty" json:"thresholdValue,attr,omitempty" yaml:"thresholdValue,attr,omitempty"`
	AnchorDistance  uint          `xml:"anchorDistance,attr,omitempty" json:"anchorDistance,attr,omitempty" yaml:"anchorDistance,attr,omitempty"`
	Anchor          bool          `xml:"anchor,attr,omitempty" json:"anchor,attr,omitempty" yaml:"anchor,attr,omitempty"`
	Var             string        `xml:"var,attr,omitempty" json:"var,attr,omitempty" yaml:"var,attr,omitempty"`
	Class           string        `xml:"class,attr,omitempty" json:"class,attr,omitempty" yaml:"class,attr,omitempty"`
	TypeAttrXSI     string        `xml:"xsi:type,attr,omitempty"`
	TypeNamespace   string        `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *DisjointWith) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:DisjointWith"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "urn:mpeg:mpqf:schema:2008"
	}
}

type Divide struct {
	DoubleValue          float64                  `xml:"DoubleValue,omitempty" json:"DoubleValue,omitempty" yaml:"DoubleValue,omitempty"`
	LongValue            int64                    `xml:"LongValue,omitempty" json:"LongValue,omitempty" yaml:"LongValue,omitempty"`
	ArithmeticField      FieldType                `xml:"ArithmeticField,omitempty" json:"ArithmeticField,omitempty" yaml:"ArithmeticField,omitempty"`
	ArithmeticExpression ArithmeticExpressionType `xml:"ArithmeticExpression,omitempty" json:"ArithmeticExpression,omitempty" yaml:"ArithmeticExpression,omitempty"`
	TypeAttrXSI          string                   `xml:"xsi:type,attr,omitempty"`
	TypeNamespace        string                   `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *Divide) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:Divide"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "urn:mpeg:mpqf:schema:2008"
	}
}

type Equal struct {
	PreferenceValue ZeroToOneType `xml:"preferenceValue,attr,omitempty" json:"preferenceValue,attr,omitempty" yaml:"preferenceValue,attr,omitempty"`
	ThresholdValue  ZeroToOneType `xml:"thresholdValue,attr,omitempty" json:"thresholdValue,attr,omitempty" yaml:"thresholdValue,attr,omitempty"`
	TypeAttrXSI     string        `xml:"xsi:type,attr,omitempty"`
	TypeNamespace   string        `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *Equal) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:Equal"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "urn:mpeg:mpqf:schema:2008"
	}
}

type EquivalentClass struct {
	PreferenceValue ZeroToOneType `xml:"preferenceValue,attr,omitempty" json:"preferenceValue,attr,omitempty" yaml:"preferenceValue,attr,omitempty"`
	ThresholdValue  ZeroToOneType `xml:"thresholdValue,attr,omitempty" json:"thresholdValue,attr,omitempty" yaml:"thresholdValue,attr,omitempty"`
	AnchorDistance  uint          `xml:"anchorDistance,attr,omitempty" json:"anchorDistance,attr,omitempty" yaml:"anchorDistance,attr,omitempty"`
	Anchor          bool          `xml:"anchor,attr,omitempty" json:"anchor,attr,omitempty" yaml:"anchor,attr,omitempty"`
	Var             string        `xml:"var,attr,omitempty" json:"var,attr,omitempty" yaml:"var,attr,omitempty"`
	Class           string        `xml:"class,attr,omitempty" json:"class,attr,omitempty" yaml:"class,attr,omitempty"`
	TypeAttrXSI     string        `xml:"xsi:type,attr,omitempty"`
	TypeNamespace   string        `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *EquivalentClass) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:EquivalentClass"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "urn:mpeg:mpqf:schema:2008"
	}
}

type FieldType struct {
	Content            *string `xml:"Content,omitempty" json:"Content,omitempty" yaml:"Content,omitempty"`
	TypeName           string  `xml:"typeName,attr,omitempty" json:"typeName,attr,omitempty" yaml:"typeName,attr,omitempty"`
	FragmentResultName string  `xml:"fragmentResultName,attr,omitempty" json:"fragmentResultName,attr,omitempty" yaml:"fragmentResultName,attr,omitempty"`
	FromREF            IDREF   `xml:"fromREF,attr,omitempty" json:"fromREF,attr,omitempty" yaml:"fromREF,attr,omitempty"`
	FieldREF           IDREF   `xml:"fieldREF,attr,omitempty" json:"fieldREF,attr,omitempty" yaml:"fieldREF,attr,omitempty"`
	ResultMode         string  `xml:"resultMode,attr,omitempty" json:"resultMode,attr,omitempty" yaml:"resultMode,attr,omitempty"`
}

type Floor struct {
	DoubleValue          float64                  `xml:"DoubleValue,omitempty" json:"DoubleValue,omitempty" yaml:"DoubleValue,omitempty"`
	LongValue            int64                    `xml:"LongValue,omitempty" json:"LongValue,omitempty" yaml:"LongValue,omitempty"`
	ArithmeticField      FieldType                `xml:"ArithmeticField,omitempty" json:"ArithmeticField,omitempty" yaml:"ArithmeticField,omitempty"`
	ArithmeticExpression ArithmeticExpressionType `xml:"ArithmeticExpression,omitempty" json:"ArithmeticExpression,omitempty" yaml:"ArithmeticExpression,omitempty"`
	TypeAttrXSI          string                   `xml:"xsi:type,attr,omitempty"`
	TypeNamespace        string                   `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *Floor) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:Floor"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "urn:mpeg:mpqf:schema:2008"
	}
}

type GetSearchResults struct {
	Results MpegQueryType `xml:"Results,omitempty" json:"Results,omitempty" yaml:"Results,omitempty"`
}

type GetSearchResultsResponse struct {
	ResultItem []MpegQueryType `xml:"ResultItem,omitempty" json:"ResultItem,omitempty" yaml:"ResultItem,omitempty"`
}

type GetServiceCapabilities struct {
}

type GetServiceCapabilitiesResponse struct {
	Capabilities Capabilities `xml:"Capabilities,omitempty" json:"Capabilities,omitempty" yaml:"Capabilities,omitempty"`
}

type GetServiceFeatures struct {
	InputCapabilities MpegQueryType `xml:"InputCapabilities,omitempty" json:"InputCapabilities,omitempty" yaml:"InputCapabilities,omitempty"`
}

type GetServiceFeaturesResponse struct {
	OutputCapabilities MpegQueryType `xml:"OutputCapabilities,omitempty" json:"OutputCapabilities,omitempty" yaml:"OutputCapabilities,omitempty"`
}

type GreaterThan struct {
	PreferenceValue ZeroToOneType `xml:"preferenceValue,attr,omitempty" json:"preferenceValue,attr,omitempty" yaml:"preferenceValue,attr,omitempty"`
	ThresholdValue  ZeroToOneType `xml:"thresholdValue,attr,omitempty" json:"thresholdValue,attr,omitempty" yaml:"thresholdValue,attr,omitempty"`
	TypeAttrXSI     string        `xml:"xsi:type,attr,omitempty"`
	TypeNamespace   string        `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *GreaterThan) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:GreaterThan"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "urn:mpeg:mpqf:schema:2008"
	}
}

type GreaterThanEqual struct {
	PreferenceValue ZeroToOneType `xml:"preferenceValue,attr,omitempty" json:"preferenceValue,attr,omitempty" yaml:"preferenceValue,attr,omitempty"`
	ThresholdValue  ZeroToOneType `xml:"thresholdValue,attr,omitempty" json:"thresholdValue,attr,omitempty" yaml:"thresholdValue,attr,omitempty"`
	TypeAttrXSI     string        `xml:"xsi:type,attr,omitempty"`
	TypeNamespace   string        `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *GreaterThanEqual) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:GreaterThanEqual"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "urn:mpeg:mpqf:schema:2008"
	}
}

type InformationType struct {
	Code        uint64 `xml:"Code,omitempty" json:"Code,omitempty" yaml:"Code,omitempty"`
	Description string `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
}

type InlineMediaType struct {
	MediaData16 []byte `xml:"MediaData16,omitempty" json:"MediaData16,omitempty" yaml:"MediaData16,omitempty"`
	MediaData64 []byte `xml:"MediaData64,omitempty" json:"MediaData64,omitempty" yaml:"MediaData64,omitempty"`
	Type        string `xml:"type,attr,omitempty" json:"type,attr,omitempty" yaml:"type,attr,omitempty"`
}

type InputManagementType struct {
	DesiredCapability CapabilityType `xml:"DesiredCapability,omitempty" json:"DesiredCapability,omitempty" yaml:"DesiredCapability,omitempty"`
	ServiceID         []string       `xml:"ServiceID,omitempty" json:"ServiceID,omitempty" yaml:"ServiceID,omitempty"`
}

type InputQueryType struct {
	QFDeclaration     QFDeclarationType     `xml:"QFDeclaration,omitempty" json:"QFDeclaration,omitempty" yaml:"QFDeclaration,omitempty"`
	OutputDescription OutputDescriptionType `xml:"OutputDescription,omitempty" json:"OutputDescription,omitempty" yaml:"OutputDescription,omitempty"`
	QueryCondition    QueryConditionType    `xml:"QueryCondition,omitempty" json:"QueryCondition,omitempty" yaml:"QueryCondition,omitempty"`
	ServiceSelection  ServiceSelectionType  `xml:"ServiceSelection,omitempty" json:"ServiceSelection,omitempty" yaml:"ServiceSelection,omitempty"`
	PreviousAnswerID  string                `xml:"previousAnswerID,attr,omitempty" json:"previousAnswerID,attr,omitempty" yaml:"previousAnswerID,attr,omitempty"`
	ImmediateResponse bool                  `xml:"immediateResponse,attr,omitempty" json:"immediateResponse,attr,omitempty" yaml:"immediateResponse,attr,omitempty"`
	StreamingResponse bool                  `xml:"streamingResponse,attr,omitempty" json:"streamingResponse,attr,omitempty" yaml:"streamingResponse,attr,omitempty"`
	Timeout           common.Duration       `xml:"timeout,attr,omitempty" json:"timeout,attr,omitempty" yaml:"timeout,attr,omitempty"`
}

type IntegerMatrixType struct {
	Content []uint64 `xml:"Content,omitempty" json:"Content,omitempty" yaml:"Content,omitempty"`
	Dim     uint64   `xml:"dim,attr,omitempty" json:"dim,attr,omitempty" yaml:"dim,attr,omitempty"`
}

type IntersectionOf struct {
	PreferenceValue ZeroToOneType `xml:"preferenceValue,attr,omitempty" json:"preferenceValue,attr,omitempty" yaml:"preferenceValue,attr,omitempty"`
	ThresholdValue  ZeroToOneType `xml:"thresholdValue,attr,omitempty" json:"thresholdValue,attr,omitempty" yaml:"thresholdValue,attr,omitempty"`
	AnchorDistance  uint          `xml:"anchorDistance,attr,omitempty" json:"anchorDistance,attr,omitempty" yaml:"anchorDistance,attr,omitempty"`
	Anchor          bool          `xml:"anchor,attr,omitempty" json:"anchor,attr,omitempty" yaml:"anchor,attr,omitempty"`
	Var             string        `xml:"var,attr,omitempty" json:"var,attr,omitempty" yaml:"var,attr,omitempty"`
	Class           string        `xml:"class,attr,omitempty" json:"class,attr,omitempty" yaml:"class,attr,omitempty"`
	TypeAttrXSI     string        `xml:"xsi:type,attr,omitempty"`
	TypeNamespace   string        `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *IntersectionOf) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:IntersectionOf"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "urn:mpeg:mpqf:schema:2008"
	}
}

type InverseOf struct {
	PreferenceValue ZeroToOneType `xml:"preferenceValue,attr,omitempty" json:"preferenceValue,attr,omitempty" yaml:"preferenceValue,attr,omitempty"`
	ThresholdValue  ZeroToOneType `xml:"thresholdValue,attr,omitempty" json:"thresholdValue,attr,omitempty" yaml:"thresholdValue,attr,omitempty"`
	AnchorDistance  uint          `xml:"anchorDistance,attr,omitempty" json:"anchorDistance,attr,omitempty" yaml:"anchorDistance,attr,omitempty"`
	Anchor          bool          `xml:"anchor,attr,omitempty" json:"anchor,attr,omitempty" yaml:"anchor,attr,omitempty"`
	Var             string        `xml:"var,attr,omitempty" json:"var,attr,omitempty" yaml:"var,attr,omitempty"`
	Class           string        `xml:"class,attr,omitempty" json:"class,attr,omitempty" yaml:"class,attr,omitempty"`
	TypeAttrXSI     string        `xml:"xsi:type,attr,omitempty"`
	TypeNamespace   string        `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *InverseOf) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:InverseOf"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "urn:mpeg:mpqf:schema:2008"
	}
}

type JoinType struct {
	From          []string              `xml:"From,omitempty" json:"From,omitempty" yaml:"From,omitempty"`
	JoinCondition BooleanExpressionType `xml:"JoinCondition,omitempty" json:"JoinCondition,omitempty" yaml:"JoinCondition,omitempty"`
}

type LessThan struct {
	PreferenceValue ZeroToOneType `xml:"preferenceValue,attr,omitempty" json:"preferenceValue,attr,omitempty" yaml:"preferenceValue,attr,omitempty"`
	ThresholdValue  ZeroToOneType `xml:"thresholdValue,attr,omitempty" json:"thresholdValue,attr,omitempty" yaml:"thresholdValue,attr,omitempty"`
	TypeAttrXSI     string        `xml:"xsi:type,attr,omitempty"`
	TypeNamespace   string        `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *LessThan) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:LessThan"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "urn:mpeg:mpqf:schema:2008"
	}
}

type LessThanEqual struct {
	PreferenceValue ZeroToOneType `xml:"preferenceValue,attr,omitempty" json:"preferenceValue,attr,omitempty" yaml:"preferenceValue,attr,omitempty"`
	ThresholdValue  ZeroToOneType `xml:"thresholdValue,attr,omitempty" json:"thresholdValue,attr,omitempty" yaml:"thresholdValue,attr,omitempty"`
	TypeAttrXSI     string        `xml:"xsi:type,attr,omitempty"`
	TypeNamespace   string        `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *LessThanEqual) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:LessThanEqual"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "urn:mpeg:mpqf:schema:2008"
	}
}

type LowerCase struct {
	StringValue      string               `xml:"StringValue,omitempty" json:"StringValue,omitempty" yaml:"StringValue,omitempty"`
	StringField      FieldType            `xml:"StringField,omitempty" json:"StringField,omitempty" yaml:"StringField,omitempty"`
	StringExpression StringExpressionType `xml:"StringExpression,omitempty" json:"StringExpression,omitempty" yaml:"StringExpression,omitempty"`
	TypeAttrXSI      string               `xml:"xsi:type,attr,omitempty"`
	TypeNamespace    string               `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *LowerCase) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:LowerCase"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "urn:mpeg:mpqf:schema:2008"
	}
}

type MAX struct {
	AggregateID   string    `xml:"aggregateID,attr,omitempty" json:"aggregateID,attr,omitempty" yaml:"aggregateID,attr,omitempty"`
	Field         FieldType `xml:"Field,omitempty" json:"Field,omitempty" yaml:"Field,omitempty"`
	TypeAttrXSI   string    `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string    `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *MAX) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:MAX"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "urn:mpeg:mpqf:schema:2008"
	}
}

type MIN struct {
	AggregateID   string    `xml:"aggregateID,attr,omitempty" json:"aggregateID,attr,omitempty" yaml:"aggregateID,attr,omitempty"`
	Field         FieldType `xml:"Field,omitempty" json:"Field,omitempty" yaml:"Field,omitempty"`
	TypeAttrXSI   string    `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string    `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *MIN) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:MIN"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "urn:mpeg:mpqf:schema:2008"
	}
}

type MediaLocatorType struct {
	MediaUri    string          `xml:"MediaUri,omitempty" json:"MediaUri,omitempty" yaml:"MediaUri,omitempty"`
	InlineMedia InlineMediaType `xml:"InlineMedia,omitempty" json:"InlineMedia,omitempty" yaml:"InlineMedia,omitempty"`
}

type MediaResourceType struct {
	ResourceID    string           `xml:"resourceID,attr,omitempty" json:"resourceID,attr,omitempty" yaml:"resourceID,attr,omitempty"`
	MediaResource MediaLocatorType `xml:"MediaResource,omitempty" json:"MediaResource,omitempty" yaml:"MediaResource,omitempty"`
	TypeAttrXSI   string           `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string           `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *MediaResourceType) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:MediaResourceType"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "urn:mpeg:mpqf:schema:2008"
	}
}

type Modulus struct {
	DoubleValue          float64                  `xml:"DoubleValue,omitempty" json:"DoubleValue,omitempty" yaml:"DoubleValue,omitempty"`
	LongValue            int64                    `xml:"LongValue,omitempty" json:"LongValue,omitempty" yaml:"LongValue,omitempty"`
	ArithmeticField      FieldType                `xml:"ArithmeticField,omitempty" json:"ArithmeticField,omitempty" yaml:"ArithmeticField,omitempty"`
	ArithmeticExpression ArithmeticExpressionType `xml:"ArithmeticExpression,omitempty" json:"ArithmeticExpression,omitempty" yaml:"ArithmeticExpression,omitempty"`
	TypeAttrXSI          string                   `xml:"xsi:type,attr,omitempty"`
	TypeNamespace        string                   `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *Modulus) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:Modulus"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "urn:mpeg:mpqf:schema:2008"
	}
}

type MpegQueryType struct {
	Query      string `xml:"Query,omitempty" json:"Query,omitempty" yaml:"Query,omitempty"`
	Management string `xml:"Management,omitempty" json:"Management,omitempty" yaml:"Management,omitempty"`
	MpqfID     string `xml:"mpqfID,attr,omitempty" json:"mpqfID,attr,omitempty" yaml:"mpqfID,attr,omitempty"`
}

type Multiply struct {
	DoubleValue          float64                  `xml:"DoubleValue,omitempty" json:"DoubleValue,omitempty" yaml:"DoubleValue,omitempty"`
	LongValue            int64                    `xml:"LongValue,omitempty" json:"LongValue,omitempty" yaml:"LongValue,omitempty"`
	ArithmeticField      FieldType                `xml:"ArithmeticField,omitempty" json:"ArithmeticField,omitempty" yaml:"ArithmeticField,omitempty"`
	ArithmeticExpression ArithmeticExpressionType `xml:"ArithmeticExpression,omitempty" json:"ArithmeticExpression,omitempty" yaml:"ArithmeticExpression,omitempty"`
	TypeAttrXSI          string                   `xml:"xsi:type,attr,omitempty"`
	TypeNamespace        string                   `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *Multiply) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:Multiply"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "urn:mpeg:mpqf:schema:2008"
	}
}

type NOT struct {
	PreferenceValue ZeroToOneType         `xml:"preferenceValue,attr,omitempty" json:"preferenceValue,attr,omitempty" yaml:"preferenceValue,attr,omitempty"`
	ThresholdValue  ZeroToOneType         `xml:"thresholdValue,attr,omitempty" json:"thresholdValue,attr,omitempty" yaml:"thresholdValue,attr,omitempty"`
	Condition       BooleanExpressionType `xml:"Condition,omitempty" json:"Condition,omitempty" yaml:"Condition,omitempty"`
	TypeAttrXSI     string                `xml:"xsi:type,attr,omitempty"`
	TypeNamespace   string                `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *NOT) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:NOT"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "urn:mpeg:mpqf:schema:2008"
	}
}

type NotEqual struct {
	PreferenceValue ZeroToOneType `xml:"preferenceValue,attr,omitempty" json:"preferenceValue,attr,omitempty" yaml:"preferenceValue,attr,omitempty"`
	ThresholdValue  ZeroToOneType `xml:"thresholdValue,attr,omitempty" json:"thresholdValue,attr,omitempty" yaml:"thresholdValue,attr,omitempty"`
	TypeAttrXSI     string        `xml:"xsi:type,attr,omitempty"`
	TypeNamespace   string        `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *NotEqual) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:NotEqual"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "urn:mpeg:mpqf:schema:2008"
	}
}

type OR struct {
	PreferenceValue ZeroToOneType           `xml:"preferenceValue,attr,omitempty" json:"preferenceValue,attr,omitempty" yaml:"preferenceValue,attr,omitempty"`
	ThresholdValue  ZeroToOneType           `xml:"thresholdValue,attr,omitempty" json:"thresholdValue,attr,omitempty" yaml:"thresholdValue,attr,omitempty"`
	ScoringFunction SimpleTermType          `xml:"scoringFunction,attr,omitempty" json:"scoringFunction,attr,omitempty" yaml:"scoringFunction,attr,omitempty"`
	Condition       []BooleanExpressionType `xml:"Condition" json:"Condition" yaml:"Condition"`
	TypeAttrXSI     string                  `xml:"xsi:type,attr,omitempty"`
	TypeNamespace   string                  `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *OR) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:OR"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "urn:mpeg:mpqf:schema:2008"
	}
}

type OutputDescriptionType struct {
	ReqField         []FieldType          `xml:"ReqField,omitempty" json:"ReqField,omitempty" yaml:"ReqField,omitempty"`
	ReqAggregateID   []*IDREF             `xml:"ReqAggregateID,omitempty" json:"ReqAggregateID,omitempty" yaml:"ReqAggregateID,omitempty"`
	ReqSemanticField []string             `xml:"ReqSemanticField,omitempty" json:"ReqSemanticField,omitempty" yaml:"ReqSemanticField,omitempty"`
	GroupBy          string               `xml:"GroupBy,omitempty" json:"GroupBy,omitempty" yaml:"GroupBy,omitempty"`
	SortBy           []AbstractSortByType `xml:"SortBy,omitempty" json:"SortBy,omitempty" yaml:"SortBy,omitempty"`
	MaxPageEntries   uint64               `xml:"maxPageEntries,attr,omitempty" json:"maxPageEntries,attr,omitempty" yaml:"maxPageEntries,attr,omitempty"`
	MaxItemCount     uint64               `xml:"maxItemCount,attr,omitempty" json:"maxItemCount,attr,omitempty" yaml:"maxItemCount,attr,omitempty"`
	FreeTextUse      bool                 `xml:"freeTextUse,attr,omitempty" json:"freeTextUse,attr,omitempty" yaml:"freeTextUse,attr,omitempty"`
	ThumbnailUse     bool                 `xml:"thumbnailUse,attr,omitempty" json:"thumbnailUse,attr,omitempty" yaml:"thumbnailUse,attr,omitempty"`
	MediaResourceUse bool                 `xml:"mediaResourceUse,attr,omitempty" json:"mediaResourceUse,attr,omitempty" yaml:"mediaResourceUse,attr,omitempty"`
	OutputNameSpace  string               `xml:"outputNameSpace,attr,omitempty" json:"outputNameSpace,attr,omitempty" yaml:"outputNameSpace,attr,omitempty"`
	Distinct         bool                 `xml:"distinct,attr,omitempty" json:"distinct,attr,omitempty" yaml:"distinct,attr,omitempty"`
}

type OutputManagementType struct {
	AvailableCapability []AvailableCapabilityType `xml:"AvailableCapability,omitempty" json:"AvailableCapability,omitempty" yaml:"AvailableCapability,omitempty"`
	SystemMessage       SystemMessageType         `xml:"SystemMessage,omitempty" json:"SystemMessage,omitempty" yaml:"SystemMessage,omitempty"`
}

type OutputQueryType struct {
	GlobalComment  string            `xml:"GlobalComment,omitempty" json:"GlobalComment,omitempty" yaml:"GlobalComment,omitempty"`
	ResultItem     []ResultItemType  `xml:"ResultItem,omitempty" json:"ResultItem,omitempty" yaml:"ResultItem,omitempty"`
	SystemMessage  SystemMessageType `xml:"SystemMessage,omitempty" json:"SystemMessage,omitempty" yaml:"SystemMessage,omitempty"`
	CurrPage       uint64            `xml:"currPage,attr,omitempty" json:"currPage,attr,omitempty" yaml:"currPage,attr,omitempty"`
	TotalPages     uint64            `xml:"totalPages,attr,omitempty" json:"totalPages,attr,omitempty" yaml:"totalPages,attr,omitempty"`
	ExpirationDate common.DateTime   `xml:"expirationDate,attr,omitempty" json:"expirationDate,attr,omitempty" yaml:"expirationDate,attr,omitempty"`
}

type QFDeclarationType struct {
	DeclaredField []DeclaredFieldType `xml:"DeclaredField,omitempty" json:"DeclaredField,omitempty" yaml:"DeclaredField,omitempty"`
	Resource      []ResourceType      `xml:"Resource,omitempty" json:"Resource,omitempty" yaml:"Resource,omitempty"`
	Prefix        []string            `xml:"Prefix,omitempty" json:"Prefix,omitempty" yaml:"Prefix,omitempty"`
}

type QueryByDescription struct {
	PreferenceValue        ZeroToOneType           `xml:"preferenceValue,attr,omitempty" json:"preferenceValue,attr,omitempty" yaml:"preferenceValue,attr,omitempty"`
	ThresholdValue         ZeroToOneType           `xml:"thresholdValue,attr,omitempty" json:"thresholdValue,attr,omitempty" yaml:"thresholdValue,attr,omitempty"`
	MatchType              string                  `xml:"matchType,attr,omitempty" json:"matchType,attr,omitempty" yaml:"matchType,attr,omitempty"`
	TargetMediaPath        string                  `xml:" TargetMediaPath,attr,omitempty" json:" TargetMediaPath,attr,omitempty" yaml:" TargetMediaPath,attr,omitempty"`
	DescriptionResource    DescriptionResourceType `xml:"DescriptionResource,omitempty" json:"DescriptionResource,omitempty" yaml:"DescriptionResource,omitempty"`
	DescriptionResourceREF *IDREF                  `xml:"DescriptionResourceREF,omitempty" json:"DescriptionResourceREF,omitempty" yaml:"DescriptionResourceREF,omitempty"`
	TypeAttrXSI            string                  `xml:"xsi:type,attr,omitempty"`
	TypeNamespace          string                  `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *QueryByDescription) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:QueryByDescription"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "urn:mpeg:mpqf:schema:2008"
	}
}

type QueryByFeatureRange struct {
	PreferenceValue ZeroToOneType `xml:"preferenceValue,attr,omitempty" json:"preferenceValue,attr,omitempty" yaml:"preferenceValue,attr,omitempty"`
	ThresholdValue  ZeroToOneType `xml:"thresholdValue,attr,omitempty" json:"thresholdValue,attr,omitempty" yaml:"thresholdValue,attr,omitempty"`
	TargetMediaPath string        `xml:" TargetMediaPath,attr,omitempty" json:" TargetMediaPath,attr,omitempty" yaml:" TargetMediaPath,attr,omitempty"`
	Range           string        `xml:"Range,omitempty" json:"Range,omitempty" yaml:"Range,omitempty"`
	Distance        string        `xml:"Distance,omitempty" json:"Distance,omitempty" yaml:"Distance,omitempty"`
	TypeAttrXSI     string        `xml:"xsi:type,attr,omitempty"`
	TypeNamespace   string        `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *QueryByFeatureRange) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:QueryByFeatureRange"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "urn:mpeg:mpqf:schema:2008"
	}
}

type QueryByFreeText struct {
	PreferenceValue ZeroToOneType `xml:"preferenceValue,attr,omitempty" json:"preferenceValue,attr,omitempty" yaml:"preferenceValue,attr,omitempty"`
	ThresholdValue  ZeroToOneType `xml:"thresholdValue,attr,omitempty" json:"thresholdValue,attr,omitempty" yaml:"thresholdValue,attr,omitempty"`
	FreeText        string        `xml:"FreeText,omitempty" json:"FreeText,omitempty" yaml:"FreeText,omitempty"`
	RegExp          string        `xml:"RegExp,omitempty" json:"RegExp,omitempty" yaml:"RegExp,omitempty"`
	SearchField     []FieldType   `xml:"SearchField,omitempty" json:"SearchField,omitempty" yaml:"SearchField,omitempty"`
	IgnoreField     []FieldType   `xml:"IgnoreField,omitempty" json:"IgnoreField,omitempty" yaml:"IgnoreField,omitempty"`
	TypeAttrXSI     string        `xml:"xsi:type,attr,omitempty"`
	TypeNamespace   string        `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *QueryByFreeText) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:QueryByFreeText"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "urn:mpeg:mpqf:schema:2008"
	}
}

type QueryByMedia struct {
	PreferenceValue  ZeroToOneType     `xml:"preferenceValue,attr,omitempty" json:"preferenceValue,attr,omitempty" yaml:"preferenceValue,attr,omitempty"`
	ThresholdValue   ZeroToOneType     `xml:"thresholdValue,attr,omitempty" json:"thresholdValue,attr,omitempty" yaml:"thresholdValue,attr,omitempty"`
	MatchType        string            `xml:"matchType,attr,omitempty" json:"matchType,attr,omitempty" yaml:"matchType,attr,omitempty"`
	TargetMediaPath  string            `xml:" TargetMediaPath,attr,omitempty" json:" TargetMediaPath,attr,omitempty" yaml:" TargetMediaPath,attr,omitempty"`
	MediaResource    MediaResourceType `xml:"MediaResource,omitempty" json:"MediaResource,omitempty" yaml:"MediaResource,omitempty"`
	MediaResourceREF *IDREF            `xml:"MediaResourceREF,omitempty" json:"MediaResourceREF,omitempty" yaml:"MediaResourceREF,omitempty"`
	TypeAttrXSI      string            `xml:"xsi:type,attr,omitempty"`
	TypeNamespace    string            `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *QueryByMedia) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:QueryByMedia"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "urn:mpeg:mpqf:schema:2008"
	}
}

type QueryByROI struct {
	PreferenceValue          ZeroToOneType        `xml:"preferenceValue,attr,omitempty" json:"preferenceValue,attr,omitempty" yaml:"preferenceValue,attr,omitempty"`
	ThresholdValue           ZeroToOneType        `xml:"thresholdValue,attr,omitempty" json:"thresholdValue,attr,omitempty" yaml:"thresholdValue,attr,omitempty"`
	MatchType                string               `xml:"matchType,attr,omitempty" json:"matchType,attr,omitempty" yaml:"matchType,attr,omitempty"`
	TargetMediaPath          string               `xml:" TargetMediaPath,attr,omitempty" json:" TargetMediaPath,attr,omitempty" yaml:" TargetMediaPath,attr,omitempty"`
	MediaResource            MediaResourceType    `xml:"MediaResource,omitempty" json:"MediaResource,omitempty" yaml:"MediaResource,omitempty"`
	MediaResourceREF         *IDREF               `xml:"MediaResourceREF,omitempty" json:"MediaResourceREF,omitempty" yaml:"MediaResourceREF,omitempty"`
	TemporalRegionOfInterest []TemporalRegionType `xml:"TemporalRegionOfInterest,omitempty" json:"TemporalRegionOfInterest,omitempty" yaml:"TemporalRegionOfInterest,omitempty"`
	SpatialRegionOfInterest  []IntegerMatrixType  `xml:"SpatialRegionOfInterest,omitempty" json:"SpatialRegionOfInterest,omitempty" yaml:"SpatialRegionOfInterest,omitempty"`
	TypeAttrXSI              string               `xml:"xsi:type,attr,omitempty"`
	TypeNamespace            string               `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *QueryByROI) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:QueryByROI"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "urn:mpeg:mpqf:schema:2008"
	}
}

type QueryByRelevanceFeedback struct {
	PreferenceValue ZeroToOneType `xml:"preferenceValue,attr,omitempty" json:"preferenceValue,attr,omitempty" yaml:"preferenceValue,attr,omitempty"`
	ThresholdValue  ZeroToOneType `xml:"thresholdValue,attr,omitempty" json:"thresholdValue,attr,omitempty" yaml:"thresholdValue,attr,omitempty"`
	AnswerID        string        `xml:"answerID,attr,omitempty" json:"answerID,attr,omitempty" yaml:"answerID,attr,omitempty"`
	ResultItem      []uint64      `xml:"ResultItem,omitempty" json:"ResultItem,omitempty" yaml:"ResultItem,omitempty"`
	TypeAttrXSI     string        `xml:"xsi:type,attr,omitempty"`
	TypeNamespace   string        `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *QueryByRelevanceFeedback) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:QueryByRelevanceFeedback"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "urn:mpeg:mpqf:schema:2008"
	}
}

type QueryBySPARQL struct {
	PreferenceValue ZeroToOneType `xml:"preferenceValue,attr,omitempty" json:"preferenceValue,attr,omitempty" yaml:"preferenceValue,attr,omitempty"`
	ThresholdValue  ZeroToOneType `xml:"thresholdValue,attr,omitempty" json:"thresholdValue,attr,omitempty" yaml:"thresholdValue,attr,omitempty"`
	SPARQL          string        `xml:"SPARQL,omitempty" json:"SPARQL,omitempty" yaml:"SPARQL,omitempty"`
	TypeAttrXSI     string        `xml:"xsi:type,attr,omitempty"`
	TypeNamespace   string        `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *QueryBySPARQL) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:QueryBySPARQL"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "urn:mpeg:mpqf:schema:2008"
	}
}

type QueryByXQuery struct {
	PreferenceValue ZeroToOneType `xml:"preferenceValue,attr,omitempty" json:"preferenceValue,attr,omitempty" yaml:"preferenceValue,attr,omitempty"`
	ThresholdValue  ZeroToOneType `xml:"thresholdValue,attr,omitempty" json:"thresholdValue,attr,omitempty" yaml:"thresholdValue,attr,omitempty"`
	XQuery          string        `xml:"XQuery,omitempty" json:"XQuery,omitempty" yaml:"XQuery,omitempty"`
	TypeAttrXSI     string        `xml:"xsi:type,attr,omitempty"`
	TypeNamespace   string        `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *QueryByXQuery) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:QueryByXQuery"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "urn:mpeg:mpqf:schema:2008"
	}
}

type QueryConditionType struct {
	EvaluationPath  *string               `xml:"EvaluationPath,omitempty" json:"EvaluationPath,omitempty" yaml:"EvaluationPath,omitempty"`
	TargetMediaType []*string             `xml:"TargetMediaType,omitempty" json:"TargetMediaType,omitempty" yaml:"TargetMediaType,omitempty"`
	Join            JoinType              `xml:"Join,omitempty" json:"Join,omitempty" yaml:"Join,omitempty"`
	Condition       BooleanExpressionType `xml:"Condition,omitempty" json:"Condition,omitempty" yaml:"Condition,omitempty"`
}

type QueryType interface{}

type RegisterDatabase struct {
	SupportedQFProfile                   CapabilityTermType        `xml:"SupportedQFProfile,omitempty" json:"SupportedQFProfile,omitempty" yaml:"SupportedQFProfile,omitempty"`
	SupportedMetadata                    []string                  `xml:"SupportedMetadata,omitempty" json:"SupportedMetadata,omitempty" yaml:"SupportedMetadata,omitempty"`
	SupportedExampleMediaTypes           string                    `xml:"SupportedExampleMediaTypes,omitempty" json:"SupportedExampleMediaTypes,omitempty" yaml:"SupportedExampleMediaTypes,omitempty"`
	SupportedResultMediaTypes            string                    `xml:"SupportedResultMediaTypes,omitempty" json:"SupportedResultMediaTypes,omitempty" yaml:"SupportedResultMediaTypes,omitempty"`
	SupportedQueryTypes                  []CapabilityTermType      `xml:"SupportedQueryTypes,omitempty" json:"SupportedQueryTypes,omitempty" yaml:"SupportedQueryTypes,omitempty"`
	SupportedExpressions                 []CapabilityTermType      `xml:"SupportedExpressions,omitempty" json:"SupportedExpressions,omitempty" yaml:"SupportedExpressions,omitempty"`
	UsageConditions                      []string                  `xml:"UsageConditions,omitempty" json:"UsageConditions,omitempty" yaml:"UsageConditions,omitempty"`
	ServiceID                            string                    `xml:"serviceID,attr,omitempty" json:"serviceID,attr,omitempty" yaml:"serviceID,attr,omitempty"`
	RecordingSearchInterfaceRegistration bool                      `xml:"RecordingSearchInterfaceRegistration,attr,omitempty" json:"RecordingSearchInterfaceRegistration,attr,omitempty" yaml:"RecordingSearchInterfaceRegistration,attr,omitempty"`
	Extension                            RegisterDatabaseExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
	TypeAttrXSI                          string                    `xml:"xsi:type,attr,omitempty"`
	TypeNamespace                        string                    `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *RegisterDatabase) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:RegisterDatabase"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://www.onvif.org/ver10/federatedsearch/wsdl"
	}
}

type RegisterDatabaseExtension []interface{}

type RegisterDatabaseResponse []interface{}

type RelationType struct {
	RelationType   string `xml:"relationType,attr,omitempty" json:"relationType,attr,omitempty" yaml:"relationType,attr,omitempty"`
	SourceResource IDREF  `xml:"sourceResource,attr,omitempty" json:"sourceResource,attr,omitempty" yaml:"sourceResource,attr,omitempty"`
	TargetResource IDREF  `xml:"targetResource,attr,omitempty" json:"targetResource,attr,omitempty" yaml:"targetResource,attr,omitempty"`
}

type ResourceType interface{}

type ResultItemBaseType interface{}

type ResultItemType struct {
	RecordNumber      uint64        `xml:"recordNumber,attr,omitempty" json:"recordNumber,attr,omitempty" yaml:"recordNumber,attr,omitempty"`
	Rank              uint64        `xml:"rank,attr,omitempty" json:"rank,attr,omitempty" yaml:"rank,attr,omitempty"`
	Confidence        ZeroToOneType `xml:"confidence,attr,omitempty" json:"confidence,attr,omitempty" yaml:"confidence,attr,omitempty"`
	OriginID          string        `xml:"originID,attr,omitempty" json:"originID,attr,omitempty" yaml:"originID,attr,omitempty"`
	Comment           []string      `xml:"Comment,omitempty" json:"Comment,omitempty" yaml:"Comment,omitempty"`
	TextResult        []string      `xml:"TextResult,omitempty" json:"TextResult,omitempty" yaml:"TextResult,omitempty"`
	Thumbnail         []string      `xml:"Thumbnail,omitempty" json:"Thumbnail,omitempty" yaml:"Thumbnail,omitempty"`
	MediaResource     []string      `xml:"MediaResource,omitempty" json:"MediaResource,omitempty" yaml:"MediaResource,omitempty"`
	Description       []string      `xml:"Description>Description,omitempty" json:"Description>Description,omitempty" yaml:"Description>Description,omitempty"`
	AggregationResult []string      `xml:"AggregationResult,omitempty" json:"AggregationResult,omitempty" yaml:"AggregationResult,omitempty"`
	FragmentResult    []string      `xml:"FragmentResult,omitempty" json:"FragmentResult,omitempty" yaml:"FragmentResult,omitempty"`
	TypeAttrXSI       string        `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string        `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *ResultItemType) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:ResultItemType"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "urn:mpeg:mpqf:schema:2008"
	}
}

type Round struct {
	DoubleValue          float64                  `xml:"DoubleValue,omitempty" json:"DoubleValue,omitempty" yaml:"DoubleValue,omitempty"`
	LongValue            int64                    `xml:"LongValue,omitempty" json:"LongValue,omitempty" yaml:"LongValue,omitempty"`
	ArithmeticField      FieldType                `xml:"ArithmeticField,omitempty" json:"ArithmeticField,omitempty" yaml:"ArithmeticField,omitempty"`
	ArithmeticExpression ArithmeticExpressionType `xml:"ArithmeticExpression,omitempty" json:"ArithmeticExpression,omitempty" yaml:"ArithmeticExpression,omitempty"`
	TypeAttrXSI          string                   `xml:"xsi:type,attr,omitempty"`
	TypeNamespace        string                   `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *Round) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:Round"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "urn:mpeg:mpqf:schema:2008"
	}
}

type SUM struct {
	AggregateID   string    `xml:"aggregateID,attr,omitempty" json:"aggregateID,attr,omitempty" yaml:"aggregateID,attr,omitempty"`
	Field         FieldType `xml:"Field,omitempty" json:"Field,omitempty" yaml:"Field,omitempty"`
	TypeAttrXSI   string    `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string    `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *SUM) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:SUM"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "urn:mpeg:mpqf:schema:2008"
	}
}

type Search struct {
	InputQuery MpegQueryType `xml:"InputQuery,omitempty" json:"InputQuery,omitempty" yaml:"InputQuery,omitempty"`
}

type SearchResponse struct {
	OutputQuery MpegQueryType `xml:"OutputQuery,omitempty" json:"OutputQuery,omitempty" yaml:"OutputQuery,omitempty"`
}

type SemanticExpressionType interface{}

type SemanticFieldType struct {
	Var string `xml:"Var,omitempty" json:"Var,omitempty" yaml:"Var,omitempty"`
}

type SemanticRelation struct {
	PreferenceValue ZeroToOneType `xml:"preferenceValue,attr,omitempty" json:"preferenceValue,attr,omitempty" yaml:"preferenceValue,attr,omitempty"`
	ThresholdValue  ZeroToOneType `xml:"thresholdValue,attr,omitempty" json:"thresholdValue,attr,omitempty" yaml:"thresholdValue,attr,omitempty"`
	AnchorDistance  uint          `xml:"anchorDistance,attr,omitempty" json:"anchorDistance,attr,omitempty" yaml:"anchorDistance,attr,omitempty"`
	Anchor          bool          `xml:"anchor,attr,omitempty" json:"anchor,attr,omitempty" yaml:"anchor,attr,omitempty"`
	Subject         string        `xml:"Subject,omitempty" json:"Subject,omitempty" yaml:"Subject,omitempty"`
	Property        string        `xml:"Property,omitempty" json:"Property,omitempty" yaml:"Property,omitempty"`
	Object          string        `xml:"Object,omitempty" json:"Object,omitempty" yaml:"Object,omitempty"`
	TypeAttrXSI     string        `xml:"xsi:type,attr,omitempty"`
	TypeNamespace   string        `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *SemanticRelation) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:SemanticRelation"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "urn:mpeg:mpqf:schema:2008"
	}
}

type ServiceSelectionType struct {
	ServiceID []string `xml:"ServiceID,omitempty" json:"ServiceID,omitempty" yaml:"ServiceID,omitempty"`
}

type SortByAggregateType struct {
	Order         string                  `xml:"order,attr,omitempty" json:"order,attr,omitempty" yaml:"order,attr,omitempty"`
	Aggregate     AggregateExpressionType `xml:"Aggregate,omitempty" json:"Aggregate,omitempty" yaml:"Aggregate,omitempty"`
	AggregateID   *IDREF                  `xml:"AggregateID,omitempty" json:"AggregateID,omitempty" yaml:"AggregateID,omitempty"`
	TypeAttrXSI   string                  `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string                  `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *SortByAggregateType) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:SortByAggregateType"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "urn:mpeg:mpqf:schema:2008"
	}
}

type SortByFieldType struct {
	Order         string            `xml:"order,attr,omitempty" json:"order,attr,omitempty" yaml:"order,attr,omitempty"`
	Field         FieldType         `xml:"Field,omitempty" json:"Field,omitempty" yaml:"Field,omitempty"`
	SemanticField SemanticFieldType `xml:"SemanticField,omitempty" json:"SemanticField,omitempty" yaml:"SemanticField,omitempty"`
	TypeAttrXSI   string            `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string            `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *SortByFieldType) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:SortByFieldType"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "urn:mpeg:mpqf:schema:2008"
	}
}

type SpatialQuery struct {
	PreferenceValue ZeroToOneType `xml:"preferenceValue,attr,omitempty" json:"preferenceValue,attr,omitempty" yaml:"preferenceValue,attr,omitempty"`
	ThresholdValue  ZeroToOneType `xml:"thresholdValue,attr,omitempty" json:"thresholdValue,attr,omitempty" yaml:"thresholdValue,attr,omitempty"`
	TargetMediaPath string        `xml:" TargetMediaPath,attr,omitempty" json:" TargetMediaPath,attr,omitempty" yaml:" TargetMediaPath,attr,omitempty"`
	SpatialRelation RelationType  `xml:"SpatialRelation,omitempty" json:"SpatialRelation,omitempty" yaml:"SpatialRelation,omitempty"`
	TypeAttrXSI     string        `xml:"xsi:type,attr,omitempty"`
	TypeNamespace   string        `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *SpatialQuery) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:SpatialQuery"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "urn:mpeg:mpqf:schema:2008"
	}
}

type StdDev struct {
	AggregateID   string    `xml:"aggregateID,attr,omitempty" json:"aggregateID,attr,omitempty" yaml:"aggregateID,attr,omitempty"`
	Field         FieldType `xml:"Field,omitempty" json:"Field,omitempty" yaml:"Field,omitempty"`
	TypeAttrXSI   string    `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string    `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *StdDev) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:StdDev"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "urn:mpeg:mpqf:schema:2008"
	}
}

type StringExpressionType interface{}

type SubClassOf struct {
	PreferenceValue ZeroToOneType `xml:"preferenceValue,attr,omitempty" json:"preferenceValue,attr,omitempty" yaml:"preferenceValue,attr,omitempty"`
	ThresholdValue  ZeroToOneType `xml:"thresholdValue,attr,omitempty" json:"thresholdValue,attr,omitempty" yaml:"thresholdValue,attr,omitempty"`
	AnchorDistance  uint          `xml:"anchorDistance,attr,omitempty" json:"anchorDistance,attr,omitempty" yaml:"anchorDistance,attr,omitempty"`
	Anchor          bool          `xml:"anchor,attr,omitempty" json:"anchor,attr,omitempty" yaml:"anchor,attr,omitempty"`
	Var             string        `xml:"var,attr,omitempty" json:"var,attr,omitempty" yaml:"var,attr,omitempty"`
	Class           string        `xml:"class,attr,omitempty" json:"class,attr,omitempty" yaml:"class,attr,omitempty"`
	TypeAttrXSI     string        `xml:"xsi:type,attr,omitempty"`
	TypeNamespace   string        `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *SubClassOf) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:SubClassOf"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "urn:mpeg:mpqf:schema:2008"
	}
}

type Subtract struct {
	DoubleValue          float64                  `xml:"DoubleValue,omitempty" json:"DoubleValue,omitempty" yaml:"DoubleValue,omitempty"`
	LongValue            int64                    `xml:"LongValue,omitempty" json:"LongValue,omitempty" yaml:"LongValue,omitempty"`
	ArithmeticField      FieldType                `xml:"ArithmeticField,omitempty" json:"ArithmeticField,omitempty" yaml:"ArithmeticField,omitempty"`
	ArithmeticExpression ArithmeticExpressionType `xml:"ArithmeticExpression,omitempty" json:"ArithmeticExpression,omitempty" yaml:"ArithmeticExpression,omitempty"`
	TypeAttrXSI          string                   `xml:"xsi:type,attr,omitempty"`
	TypeNamespace        string                   `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *Subtract) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:Subtract"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "urn:mpeg:mpqf:schema:2008"
	}
}

type SystemMessageType struct {
	Status    []InformationType `xml:"Status,omitempty" json:"Status,omitempty" yaml:"Status,omitempty"`
	Warning   []InformationType `xml:"Warning,omitempty" json:"Warning,omitempty" yaml:"Warning,omitempty"`
	Exception []InformationType `xml:"Exception,omitempty" json:"Exception,omitempty" yaml:"Exception,omitempty"`
}

type TemporalQuery struct {
	PreferenceValue  ZeroToOneType `xml:"preferenceValue,attr,omitempty" json:"preferenceValue,attr,omitempty" yaml:"preferenceValue,attr,omitempty"`
	ThresholdValue   ZeroToOneType `xml:"thresholdValue,attr,omitempty" json:"thresholdValue,attr,omitempty" yaml:"thresholdValue,attr,omitempty"`
	TargetMediaPath  string        `xml:" TargetMediaPath,attr,omitempty" json:" TargetMediaPath,attr,omitempty" yaml:" TargetMediaPath,attr,omitempty"`
	TemporalRelation RelationType  `xml:"TemporalRelation,omitempty" json:"TemporalRelation,omitempty" yaml:"TemporalRelation,omitempty"`
	TypeAttrXSI      string        `xml:"xsi:type,attr,omitempty"`
	TypeNamespace    string        `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *TemporalQuery) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:TemporalQuery"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "urn:mpeg:mpqf:schema:2008"
	}
}

type TemporalRegionType struct {
	StartTime StartTimePointType `xml:"StartTime,omitempty" json:"StartTime,omitempty" yaml:"StartTime,omitempty"`
	Duration  MediaDurationType  `xml:"Duration,omitempty" json:"Duration,omitempty" yaml:"Duration,omitempty"`
}

type TermType struct {
	Name        string         `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Description string         `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	Term        []TermType     `xml:"Term,omitempty" json:"Term,omitempty" yaml:"Term,omitempty"`
	Href        SimpleTermType `xml:"href,attr,omitempty" json:"href,attr,omitempty" yaml:"href,attr,omitempty"`
}

type TypeOf struct {
	PreferenceValue ZeroToOneType `xml:"preferenceValue,attr,omitempty" json:"preferenceValue,attr,omitempty" yaml:"preferenceValue,attr,omitempty"`
	ThresholdValue  ZeroToOneType `xml:"thresholdValue,attr,omitempty" json:"thresholdValue,attr,omitempty" yaml:"thresholdValue,attr,omitempty"`
	AnchorDistance  uint          `xml:"anchorDistance,attr,omitempty" json:"anchorDistance,attr,omitempty" yaml:"anchorDistance,attr,omitempty"`
	Anchor          bool          `xml:"anchor,attr,omitempty" json:"anchor,attr,omitempty" yaml:"anchor,attr,omitempty"`
	Var             string        `xml:"var,attr,omitempty" json:"var,attr,omitempty" yaml:"var,attr,omitempty"`
	Class           string        `xml:"class,attr,omitempty" json:"class,attr,omitempty" yaml:"class,attr,omitempty"`
	TypeAttrXSI     string        `xml:"xsi:type,attr,omitempty"`
	TypeNamespace   string        `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *TypeOf) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:TypeOf"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "urn:mpeg:mpqf:schema:2008"
	}
}

type UnionOf struct {
	PreferenceValue ZeroToOneType `xml:"preferenceValue,attr,omitempty" json:"preferenceValue,attr,omitempty" yaml:"preferenceValue,attr,omitempty"`
	ThresholdValue  ZeroToOneType `xml:"thresholdValue,attr,omitempty" json:"thresholdValue,attr,omitempty" yaml:"thresholdValue,attr,omitempty"`
	AnchorDistance  uint          `xml:"anchorDistance,attr,omitempty" json:"anchorDistance,attr,omitempty" yaml:"anchorDistance,attr,omitempty"`
	Anchor          bool          `xml:"anchor,attr,omitempty" json:"anchor,attr,omitempty" yaml:"anchor,attr,omitempty"`
	Var             string        `xml:"var,attr,omitempty" json:"var,attr,omitempty" yaml:"var,attr,omitempty"`
	Class           string        `xml:"class,attr,omitempty" json:"class,attr,omitempty" yaml:"class,attr,omitempty"`
	TypeAttrXSI     string        `xml:"xsi:type,attr,omitempty"`
	TypeNamespace   string        `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *UnionOf) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:UnionOf"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "urn:mpeg:mpqf:schema:2008"
	}
}

type UpperCase struct {
	StringValue      string               `xml:"StringValue,omitempty" json:"StringValue,omitempty" yaml:"StringValue,omitempty"`
	StringField      FieldType            `xml:"StringField,omitempty" json:"StringField,omitempty" yaml:"StringField,omitempty"`
	StringExpression StringExpressionType `xml:"StringExpression,omitempty" json:"StringExpression,omitempty" yaml:"StringExpression,omitempty"`
	TypeAttrXSI      string               `xml:"xsi:type,attr,omitempty"`
	TypeNamespace    string               `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *UpperCase) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:UpperCase"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "urn:mpeg:mpqf:schema:2008"
	}
}

type Variance struct {
	AggregateID   string    `xml:"aggregateID,attr,omitempty" json:"aggregateID,attr,omitempty" yaml:"aggregateID,attr,omitempty"`
	Field         FieldType `xml:"Field,omitempty" json:"Field,omitempty" yaml:"Field,omitempty"`
	TypeAttrXSI   string    `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string    `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *Variance) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:Variance"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "urn:mpeg:mpqf:schema:2008"
	}
}

type XOR struct {
	PreferenceValue ZeroToOneType           `xml:"preferenceValue,attr,omitempty" json:"preferenceValue,attr,omitempty" yaml:"preferenceValue,attr,omitempty"`
	ThresholdValue  ZeroToOneType           `xml:"thresholdValue,attr,omitempty" json:"thresholdValue,attr,omitempty" yaml:"thresholdValue,attr,omitempty"`
	ScoringFunction SimpleTermType          `xml:"scoringFunction,attr,omitempty" json:"scoringFunction,attr,omitempty" yaml:"scoringFunction,attr,omitempty"`
	Condition       []BooleanExpressionType `xml:"Condition" json:"Condition" yaml:"Condition"`
	TypeAttrXSI     string                  `xml:"xsi:type,attr,omitempty"`
	TypeNamespace   string                  `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *XOR) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:XOR"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "urn:mpeg:mpqf:schema:2008"
	}
}

// federatedSearchPort implements the FederatedSearchPort interface.
type federatedSearchPort struct {
	cli      common.Client
	Endpoint string
}

func (p *federatedSearchPort) OptGetSearchResults(args GetSearchResults) (*GetSearchResultsResponse, error) {
	req := struct {
		XMLName          string `xml:"tfs:GetSearchResults"`
		GetSearchResults GetSearchResults
	}{
		GetSearchResults: args,
	}

	resp := GetSearchResultsResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *federatedSearchPort) OptGetServiceCapabilities(args GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error) {
	req := struct {
		XMLName                string `xml:"tfs:GetServiceCapabilities"`
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

func (p *federatedSearchPort) OptGetServiceFeatures(args GetServiceFeatures) (*GetServiceFeaturesResponse, error) {
	req := struct {
		XMLName            string `xml:"tfs:GetServiceFeatures"`
		GetServiceFeatures GetServiceFeatures
	}{
		GetServiceFeatures: args,
	}

	resp := GetServiceFeaturesResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *federatedSearchPort) OptRegisterDatabase(args RegisterDatabase) (*RegisterDatabaseResponse, error) {
	req := struct {
		XMLName          string `xml:"tfs:RegisterDatabase"`
		RegisterDatabase RegisterDatabase
	}{
		RegisterDatabase: args,
	}

	resp := RegisterDatabaseResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *federatedSearchPort) OptSearch(args Search) (*SearchResponse, error) {
	req := struct {
		XMLName string `xml:"tfs:Search"`
		Search  Search
	}{
		Search: args,
	}

	resp := SearchResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
