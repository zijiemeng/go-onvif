package common

import (
	"reflect"
)

type PositiveInteger uint

type RecognitionType string

// Validate validates RecognitionType.
func (v RecognitionType) Validate() bool {
	for _, vv := range []string{
		"pt:Card",
		"pt:PIN",
		"pt:Fingerprint",
		"pt:Face",
		"pt:Iris",
		"pt:Vein",
		"pt:Palm",
		"pt:Retina",
		"pt:LicensePlate",
		"pt:REX",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type Attribute struct {
	Name  string
	Value string
}

type DataEntity struct {
	Token ReferenceToken `xml:"token,attr,omitempty" json:"token,attr,omitempty" yaml:"token,attr,omitempty"`
}
