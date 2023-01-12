package thermal

import (
	"reflect"

	"code.byted.org/videoarch/go-onvif/onvif/common"
)

var Namespace = "http://www.onvif.org/ver10/thermal/wsdl"

// NewThermalPort creates an initializes a ThermalPort.
func NewThermalPort(endpoint string, cli common.Client) ThermalPort {
	return &thermalPort{cli: cli, Endpoint: endpoint}
}

// ThermalPort was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type ThermalPort interface {
	OptGetConfiguration(GetConfiguration GetConfiguration) (*GetConfigurationResponse, error)

	OptGetConfigurationOptions(GetConfigurationOptions GetConfigurationOptions) (*GetConfigurationOptionsResponse, error)

	OptGetConfigurations(GetConfigurations GetConfigurations) (*GetConfigurationsResponse, error)

	OptGetRadiometryConfiguration(GetRadiometryConfiguration GetRadiometryConfiguration) (*GetRadiometryConfigurationResponse, error)

	OptGetRadiometryConfigurationOptions(GetRadiometryConfigurationOptions GetRadiometryConfigurationOptions) (*GetRadiometryConfigurationOptionsResponse, error)

	OptGetServiceCapabilities(GetServiceCapabilities GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error)

	OptSetConfiguration(SetConfiguration SetConfiguration) (*SetConfigurationResponse, error)

	OptSetRadiometryConfiguration(SetRadiometryConfiguration SetRadiometryConfiguration) (*SetRadiometryConfigurationResponse, error)
}
type ColorPaletteType string

func (v ColorPaletteType) Validate() bool {
	for _, vv := range []string{
		"Custom",
		"Grayscale",
		"BlackHot",
		"WhiteHot",
		"Sepia",
		"Red",
		"Iron",
		"Rain",
		"Rainbow",
		"Isotherm",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type Polarity string

func (v Polarity) Validate() bool {
	for _, vv := range []string{
		"WhiteHot",
		"BlackHot",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type Capabilities []interface{}

type ColorPalette struct {
	Name  *common.Name          `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Token common.ReferenceToken `xml:"token,attr,omitempty" json:"token,attr,omitempty" yaml:"token,attr,omitempty"`
	Type  string                `xml:"Type,attr,omitempty" json:"Type,attr,omitempty" yaml:"Type,attr,omitempty"`
}

type Configuration struct {
	ColorPalette ColorPalette `xml:"ColorPalette,omitempty" json:"ColorPalette,omitempty" yaml:"ColorPalette,omitempty"`
	Polarity     Polarity     `xml:"Polarity,omitempty" json:"Polarity,omitempty" yaml:"Polarity,omitempty"`
	NUCTable     NUCTable     `xml:"NUCTable,omitempty" json:"NUCTable,omitempty" yaml:"NUCTable,omitempty"`
	Cooler       Cooler       `xml:"Cooler,omitempty" json:"Cooler,omitempty" yaml:"Cooler,omitempty"`
}

type ConfigurationOptions struct {
	ColorPalette  []ColorPalette `xml:"ColorPalette,omitempty" json:"ColorPalette,omitempty" yaml:"ColorPalette,omitempty"`
	NUCTable      []NUCTable     `xml:"NUCTable,omitempty" json:"NUCTable,omitempty" yaml:"NUCTable,omitempty"`
	CoolerOptions CoolerOptions  `xml:"CoolerOptions,omitempty" json:"CoolerOptions,omitempty" yaml:"CoolerOptions,omitempty"`
}

type Configurations struct {
	Configuration Configuration         `xml:"Configuration,omitempty" json:"Configuration,omitempty" yaml:"Configuration,omitempty"`
	Token         common.ReferenceToken `xml:"token,attr,omitempty" json:"token,attr,omitempty" yaml:"token,attr,omitempty"`
}

type Cooler struct {
	Enabled bool    `xml:"Enabled,omitempty" json:"Enabled,omitempty" yaml:"Enabled,omitempty"`
	RunTime float64 `xml:"RunTime,omitempty" json:"RunTime,omitempty" yaml:"RunTime,omitempty"`
}

type CoolerOptions struct {
	Enabled bool `xml:"Enabled,omitempty" json:"Enabled,omitempty" yaml:"Enabled,omitempty"`
}

type GetConfiguration struct {
	VideoSourceToken *common.ReferenceToken `xml:"VideoSourceToken,omitempty" json:"VideoSourceToken,omitempty" yaml:"VideoSourceToken,omitempty"`
}

type GetConfigurationOptions struct {
	VideoSourceToken *common.ReferenceToken `xml:"VideoSourceToken,omitempty" json:"VideoSourceToken,omitempty" yaml:"VideoSourceToken,omitempty"`
}

type GetConfigurationOptionsResponse struct {
	ConfigurationOptions ConfigurationOptions `xml:"ConfigurationOptions,omitempty" json:"ConfigurationOptions,omitempty" yaml:"ConfigurationOptions,omitempty"`
}

type GetConfigurationResponse struct {
	Configuration Configuration `xml:"Configuration,omitempty" json:"Configuration,omitempty" yaml:"Configuration,omitempty"`
}

type GetConfigurations struct {
}

type GetConfigurationsResponse struct {
	Configurations []Configurations `xml:"Configurations,omitempty" json:"Configurations,omitempty" yaml:"Configurations,omitempty"`
}

type GetRadiometryConfiguration struct {
	VideoSourceToken *common.ReferenceToken `xml:"VideoSourceToken,omitempty" json:"VideoSourceToken,omitempty" yaml:"VideoSourceToken,omitempty"`
}

type GetRadiometryConfigurationOptions struct {
	VideoSourceToken *common.ReferenceToken `xml:"VideoSourceToken,omitempty" json:"VideoSourceToken,omitempty" yaml:"VideoSourceToken,omitempty"`
}

type GetRadiometryConfigurationOptionsResponse struct {
	ConfigurationOptions RadiometryConfigurationOptions `xml:"ConfigurationOptions,omitempty" json:"ConfigurationOptions,omitempty" yaml:"ConfigurationOptions,omitempty"`
}

type GetRadiometryConfigurationResponse struct {
	Configuration RadiometryConfiguration `xml:"Configuration,omitempty" json:"Configuration,omitempty" yaml:"Configuration,omitempty"`
}

type GetServiceCapabilities struct {
}

type GetServiceCapabilitiesResponse struct {
	Capabilities Capabilities `xml:"Capabilities,omitempty" json:"Capabilities,omitempty" yaml:"Capabilities,omitempty"`
}

type NUCTable struct {
	Name            *common.Name          `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Token           common.ReferenceToken `xml:"token,attr,omitempty" json:"token,attr,omitempty" yaml:"token,attr,omitempty"`
	LowTemperature  float64               `xml:"LowTemperature,attr,omitempty" json:"LowTemperature,attr,omitempty" yaml:"LowTemperature,attr,omitempty"`
	HighTemperature float64               `xml:"HighTemperature,attr,omitempty" json:"HighTemperature,attr,omitempty" yaml:"HighTemperature,attr,omitempty"`
}

type RadiometryConfiguration struct {
	RadiometryGlobalParameters RadiometryGlobalParameters `xml:"RadiometryGlobalParameters,omitempty" json:"RadiometryGlobalParameters,omitempty" yaml:"RadiometryGlobalParameters,omitempty"`
}

type RadiometryConfigurationOptions struct {
	RadiometryGlobalParameterOptions RadiometryGlobalParameterOptions `xml:"RadiometryGlobalParameterOptions,omitempty" json:"RadiometryGlobalParameterOptions,omitempty" yaml:"RadiometryGlobalParameterOptions,omitempty"`
}

type RadiometryGlobalParameterOptions struct {
	ReflectedAmbientTemperature *common.FloatRange `xml:"ReflectedAmbientTemperature,omitempty" json:"ReflectedAmbientTemperature,omitempty" yaml:"ReflectedAmbientTemperature,omitempty"`
	Emissivity                  *common.FloatRange `xml:"Emissivity,omitempty" json:"Emissivity,omitempty" yaml:"Emissivity,omitempty"`
	DistanceToObject            *common.FloatRange `xml:"DistanceToObject,omitempty" json:"DistanceToObject,omitempty" yaml:"DistanceToObject,omitempty"`
	RelativeHumidity            *common.FloatRange `xml:"RelativeHumidity,omitempty" json:"RelativeHumidity,omitempty" yaml:"RelativeHumidity,omitempty"`
	AtmosphericTemperature      *common.FloatRange `xml:"AtmosphericTemperature,omitempty" json:"AtmosphericTemperature,omitempty" yaml:"AtmosphericTemperature,omitempty"`
	AtmosphericTransmittance    *common.FloatRange `xml:"AtmosphericTransmittance,omitempty" json:"AtmosphericTransmittance,omitempty" yaml:"AtmosphericTransmittance,omitempty"`
	ExtOpticsTemperature        *common.FloatRange `xml:"ExtOpticsTemperature,omitempty" json:"ExtOpticsTemperature,omitempty" yaml:"ExtOpticsTemperature,omitempty"`
	ExtOpticsTransmittance      *common.FloatRange `xml:"ExtOpticsTransmittance,omitempty" json:"ExtOpticsTransmittance,omitempty" yaml:"ExtOpticsTransmittance,omitempty"`
}

type RadiometryGlobalParameters struct {
	ReflectedAmbientTemperature float64 `xml:"ReflectedAmbientTemperature,omitempty" json:"ReflectedAmbientTemperature,omitempty" yaml:"ReflectedAmbientTemperature,omitempty"`
	Emissivity                  float64 `xml:"Emissivity,omitempty" json:"Emissivity,omitempty" yaml:"Emissivity,omitempty"`
	DistanceToObject            float64 `xml:"DistanceToObject,omitempty" json:"DistanceToObject,omitempty" yaml:"DistanceToObject,omitempty"`
	RelativeHumidity            float64 `xml:"RelativeHumidity,omitempty" json:"RelativeHumidity,omitempty" yaml:"RelativeHumidity,omitempty"`
	AtmosphericTemperature      float64 `xml:"AtmosphericTemperature,omitempty" json:"AtmosphericTemperature,omitempty" yaml:"AtmosphericTemperature,omitempty"`
	AtmosphericTransmittance    float64 `xml:"AtmosphericTransmittance,omitempty" json:"AtmosphericTransmittance,omitempty" yaml:"AtmosphericTransmittance,omitempty"`
	ExtOpticsTemperature        float64 `xml:"ExtOpticsTemperature,omitempty" json:"ExtOpticsTemperature,omitempty" yaml:"ExtOpticsTemperature,omitempty"`
	ExtOpticsTransmittance      float64 `xml:"ExtOpticsTransmittance,omitempty" json:"ExtOpticsTransmittance,omitempty" yaml:"ExtOpticsTransmittance,omitempty"`
}

type SetConfiguration struct {
	VideoSourceToken *common.ReferenceToken `xml:"VideoSourceToken,omitempty" json:"VideoSourceToken,omitempty" yaml:"VideoSourceToken,omitempty"`
	Configuration    Configuration          `xml:"Configuration,omitempty" json:"Configuration,omitempty" yaml:"Configuration,omitempty"`
}

type SetConfigurationResponse struct {
}

type SetRadiometryConfiguration struct {
	VideoSourceToken *common.ReferenceToken  `xml:"VideoSourceToken,omitempty" json:"VideoSourceToken,omitempty" yaml:"VideoSourceToken,omitempty"`
	Configuration    RadiometryConfiguration `xml:"Configuration,omitempty" json:"Configuration,omitempty" yaml:"Configuration,omitempty"`
}

type SetRadiometryConfigurationResponse struct {
}

// thermalPort implements the ThermalPort interface.
type thermalPort struct {
	cli      common.Client
	Endpoint string
}

func (p *thermalPort) OptGetConfiguration(args GetConfiguration) (*GetConfigurationResponse, error) {
	req := struct {
		XMLName          string `xml:"tth:GetConfiguration"`
		GetConfiguration GetConfiguration
	}{
		GetConfiguration: args,
	}

	resp := GetConfigurationResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *thermalPort) OptGetConfigurationOptions(args GetConfigurationOptions) (*GetConfigurationOptionsResponse, error) {
	req := struct {
		XMLName                 string `xml:"tth:GetConfigurationOptions"`
		GetConfigurationOptions GetConfigurationOptions
	}{
		GetConfigurationOptions: args,
	}

	resp := GetConfigurationOptionsResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *thermalPort) OptGetConfigurations(args GetConfigurations) (*GetConfigurationsResponse, error) {
	req := struct {
		XMLName           string `xml:"tth:GetConfigurations"`
		GetConfigurations GetConfigurations
	}{
		GetConfigurations: args,
	}

	resp := GetConfigurationsResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *thermalPort) OptGetRadiometryConfiguration(args GetRadiometryConfiguration) (*GetRadiometryConfigurationResponse, error) {
	req := struct {
		XMLName                    string `xml:"tth:GetRadiometryConfiguration"`
		GetRadiometryConfiguration GetRadiometryConfiguration
	}{
		GetRadiometryConfiguration: args,
	}

	resp := GetRadiometryConfigurationResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *thermalPort) OptGetRadiometryConfigurationOptions(args GetRadiometryConfigurationOptions) (*GetRadiometryConfigurationOptionsResponse, error) {
	req := struct {
		XMLName                           string `xml:"tth:GetRadiometryConfigurationOptions"`
		GetRadiometryConfigurationOptions GetRadiometryConfigurationOptions
	}{
		GetRadiometryConfigurationOptions: args,
	}

	resp := GetRadiometryConfigurationOptionsResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *thermalPort) OptGetServiceCapabilities(args GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error) {
	req := struct {
		XMLName                string `xml:"tth:GetServiceCapabilities"`
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

func (p *thermalPort) OptSetConfiguration(args SetConfiguration) (*SetConfigurationResponse, error) {
	req := struct {
		XMLName          string `xml:"tth:SetConfiguration"`
		SetConfiguration SetConfiguration
	}{
		SetConfiguration: args,
	}

	resp := SetConfigurationResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *thermalPort) OptSetRadiometryConfiguration(args SetRadiometryConfiguration) (*SetRadiometryConfigurationResponse, error) {
	req := struct {
		XMLName                    string `xml:"tth:SetRadiometryConfiguration"`
		SetRadiometryConfiguration SetRadiometryConfiguration
	}{
		SetRadiometryConfiguration: args,
	}

	resp := SetRadiometryConfigurationResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
