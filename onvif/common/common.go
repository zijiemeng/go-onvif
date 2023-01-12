package common

import (
	"reflect"
)

// DateTime in WSDL format.
type DateTime string

type NCName string

type Entity string

type MoveStatus string

// Validate validates Entity.
func (v Entity) Validate() bool {
	for _, vv := range []string{
		"Device",
		"VideoSource",
		"AudioSource",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// Validate validates MoveStatus.
func (v MoveStatus) Validate() bool {
	for _, vv := range []string{
		"IDLE",
		"MOVING",
		"UNKNOWN",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type ReferenceToken string

type Color struct {
	X          float64 `xml:"X,attr,omitempty" json:"X,attr,omitempty" yaml:"X,attr,omitempty"`
	Y          float64 `xml:"Y,attr,omitempty" json:"Y,attr,omitempty" yaml:"Y,attr,omitempty"`
	Z          float64 `xml:"Z,attr,omitempty" json:"Z,attr,omitempty" yaml:"Z,attr,omitempty"`
	Colorspace string  `xml:"Colorspace,attr,omitempty" json:"Colorspace,attr,omitempty" yaml:"Colorspace,attr,omitempty"`
	Likelihood float64 `xml:"Likelihood,attr,omitempty" json:"Likelihood,attr,omitempty" yaml:"Likelihood,attr,omitempty"`
}

type ColorCovariance struct {
	XX         float64 `xml:"XX,attr,omitempty" json:"XX,attr,omitempty" yaml:"XX,attr,omitempty"`
	YY         float64 `xml:"YY,attr,omitempty" json:"YY,attr,omitempty" yaml:"YY,attr,omitempty"`
	ZZ         float64 `xml:"ZZ,attr,omitempty" json:"ZZ,attr,omitempty" yaml:"ZZ,attr,omitempty"`
	XY         float64 `xml:"XY,attr,omitempty" json:"XY,attr,omitempty" yaml:"XY,attr,omitempty"`
	XZ         float64 `xml:"XZ,attr,omitempty" json:"XZ,attr,omitempty" yaml:"XZ,attr,omitempty"`
	YZ         float64 `xml:"YZ,attr,omitempty" json:"YZ,attr,omitempty" yaml:"YZ,attr,omitempty"`
	Colorspace string  `xml:"Colorspace,attr,omitempty" json:"Colorspace,attr,omitempty" yaml:"Colorspace,attr,omitempty"`
}

type ColorDescriptor struct {
	ColorCluster []*interface{} `xml:"ColorCluster>ColorCluster,omitempty" json:"ColorCluster>ColorCluster,omitempty" yaml:"ColorCluster>ColorCluster,omitempty"`
	Extension    *interface{}   `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type GeoLocation []interface{}

type GeoOrientation []interface{}

// IntRange : Range of values greater equal Min value and less equal Max value.
type IntRange struct {
	Min *int `xml:"Min,omitempty" json:"Min,omitempty" yaml:"Min,omitempty"`
	Max *int `xml:"Max,omitempty" json:"Max,omitempty" yaml:"Max,omitempty"`
}

type LocalLocation []interface{}

type LocalOrientation []interface{}

type LocationEntity struct {
	GeoLocation      *GeoLocation      `xml:"GeoLocation,omitempty" json:"GeoLocation,omitempty" yaml:"GeoLocation,omitempty"`
	GeoOrientation   *GeoOrientation   `xml:"GeoOrientation,omitempty" json:"GeoOrientation,omitempty" yaml:"GeoOrientation,omitempty"`
	LocalLocation    *LocalLocation    `xml:"LocalLocation,omitempty" json:"LocalLocation,omitempty" yaml:"LocalLocation,omitempty"`
	LocalOrientation *LocalOrientation `xml:"LocalOrientation,omitempty" json:"LocalOrientation,omitempty" yaml:"LocalOrientation,omitempty"`
	Entity           string            `xml:"Entity,attr,omitempty" json:"Entity,attr,omitempty" yaml:"Entity,attr,omitempty"`
	Token            ReferenceToken    `xml:"Token,attr,omitempty" json:"Token,attr,omitempty" yaml:"Token,attr,omitempty"`
	Fixed            bool              `xml:"Fixed,attr,omitempty" json:"Fixed,attr,omitempty" yaml:"Fixed,attr,omitempty"`
	GeoSource        string            `xml:"GeoSource,attr,omitempty" json:"GeoSource,attr,omitempty" yaml:"GeoSource,attr,omitempty"`
	AutoGeo          bool              `xml:"AutoGeo,attr,omitempty" json:"AutoGeo,attr,omitempty" yaml:"AutoGeo,attr,omitempty"`
}

type PTZMoveStatus struct {
	PanTilt *MoveStatus `xml:"PanTilt,omitempty" json:"PanTilt,omitempty" yaml:"PanTilt,omitempty"`
	Zoom    *MoveStatus `xml:"Zoom,omitempty" json:"Zoom,omitempty" yaml:"Zoom,omitempty"`
}

type PTZStatus struct {
	Position   *PTZVector     `xml:"Position,omitempty" json:"Position,omitempty" yaml:"Position,omitempty"`
	MoveStatus *PTZMoveStatus `xml:"MoveStatus,omitempty" json:"MoveStatus,omitempty" yaml:"MoveStatus,omitempty"`
	Error      *string        `xml:"Error,omitempty" json:"Error,omitempty" yaml:"Error,omitempty"`
	UtcTime    *DateTime      `xml:"UtcTime,omitempty" json:"UtcTime,omitempty" yaml:"UtcTime,omitempty"`
}

type PTZVector struct {
	PanTilt *Vector2D `xml:"PanTilt,omitempty" json:"PanTilt,omitempty" yaml:"PanTilt,omitempty"`
	Zoom    *Vector1D `xml:"Zoom,omitempty" json:"Zoom,omitempty" yaml:"Zoom,omitempty"`
}

type Polygon struct {
	Point []*Vector `xml:"Point" json:"Point" yaml:"Point"`
}

type Rectangle struct {
	Bottom float64 `xml:"bottom,attr,omitempty" json:"bottom,attr,omitempty" yaml:"bottom,attr,omitempty"`
	Top    float64 `xml:"top,attr,omitempty" json:"top,attr,omitempty" yaml:"top,attr,omitempty"`
	Right  float64 `xml:"right,attr,omitempty" json:"right,attr,omitempty" yaml:"right,attr,omitempty"`
	Left   float64 `xml:"left,attr,omitempty" json:"left,attr,omitempty" yaml:"left,attr,omitempty"`
}

type Transformation struct {
	Translate *Vector                  `xml:"Translate,omitempty" json:"Translate,omitempty" yaml:"Translate,omitempty"`
	Scale     *Vector                  `xml:"Scale,omitempty" json:"Scale,omitempty" yaml:"Scale,omitempty"`
	Extension *TransformationExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type TransformationExtension []interface{}

type Vector struct {
	X float64 `xml:"x,attr,omitempty" json:"x,attr,omitempty" yaml:"x,attr,omitempty"`
	Y float64 `xml:"y,attr,omitempty" json:"y,attr,omitempty" yaml:"y,attr,omitempty"`
}

type Vector1D struct {
	X     float64 `xml:"x,attr,omitempty" json:"x,attr,omitempty" yaml:"x,attr,omitempty"`
	Space string  `xml:"space,attr,omitempty" json:"space,attr,omitempty" yaml:"space,attr,omitempty"`
}

type Vector2D struct {
	X     float64 `xml:"x,attr,omitempty" json:"x,attr,omitempty" yaml:"x,attr,omitempty"`
	Y     float64 `xml:"y,attr,omitempty" json:"y,attr,omitempty" yaml:"y,attr,omitempty"`
	Space string  `xml:"space,attr,omitempty" json:"space,attr,omitempty" yaml:"space,attr,omitempty"`
}

type StringAttrList struct {
	AttrList []string
}

type IntList struct {
	Items []int
}

type FloatList struct {
	Items []float64
}

type StringList struct {
	Items []StringList
}
