package device_io

import (
	"code.byted.org/videoarch/go-onvif/onvif/v10/device"
	"reflect"

	"code.byted.org/videoarch/go-onvif/onvif/common"
)

var Namespace = "http://www.onvif.org/ver10/deviceIO/wsdl"

// NewDeviceIOPort creates an initializes a DeviceIOPort.
func NewDeviceIOPort(endpoint string, cli common.Client) DeviceIOPort {
	return &deviceIOPort{cli: cli, Endpoint: endpoint}
}

// DeviceIOPort was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type DeviceIOPort interface {
	OptGetAudioOutputConfiguration(GetAudioOutputConfiguration GetAudioOutputConfiguration) (*GetAudioOutputConfigurationResponse, *common.Fault)

	OptGetAudioOutputConfigurationOptions(GetAudioOutputConfigurationOptions GetAudioOutputConfigurationOptions) (*GetAudioOutputConfigurationOptionsResponse, *common.Fault)

	OptGetAudioOutputs(GetAudioOutputs Get) (*GetResponse, *common.Fault)

	OptGetAudioSourceConfiguration(GetAudioSourceConfiguration GetAudioSourceConfiguration) (*GetAudioSourceConfigurationResponse, *common.Fault)

	OptGetAudioSourceConfigurationOptions(GetAudioSourceConfigurationOptions GetAudioSourceConfigurationOptions) (*GetAudioSourceConfigurationOptionsResponse, *common.Fault)

	OptGetAudioSources(GetAudioSources Get) (*GetResponse, *common.Fault)

	OptGetDigitalInputConfigurationOptions(GetDigitalInputConfigurationOptions GetDigitalInputConfigurationOptions) (*GetDigitalInputConfigurationOptionsResponse, *common.Fault)

	OptGetDigitalInputs(GetDigitalInputs GetDigitalInputs) (*GetDigitalInputsResponse, *common.Fault)

	OptGetRelayOutputOptions(GetRelayOutputOptions GetRelayOutputOptions) (*GetRelayOutputOptionsResponse, *common.Fault)

	OptGetRelayOutputs(GetRelayOutputs device.GetRelayOutputs) (*device.GetRelayOutputsResponse, *common.Fault)

	OptGetSerialPortConfiguration(GetSerialPortConfiguration GetSerialPortConfiguration) (*GetSerialPortConfigurationResponse, *common.Fault)

	OptGetSerialPortConfigurationOptions(GetSerialPortConfigurationOptions GetSerialPortConfigurationOptions) (*GetSerialPortConfigurationOptionsResponse, *common.Fault)

	OptGetSerialPorts(GetSerialPorts GetSerialPorts) (*GetSerialPortsResponse, *common.Fault)

	OptGetServiceCapabilities(GetServiceCapabilities GetServiceCapabilities) (*GetServiceCapabilitiesResponse, *common.Fault)

	OptGetVideoOutputConfiguration(GetVideoOutputConfiguration GetVideoOutputConfiguration) (*GetVideoOutputConfigurationResponse, *common.Fault)

	OptGetVideoOutputConfigurationOptions(GetVideoOutputConfigurationOptions GetVideoOutputConfigurationOptions) (*GetVideoOutputConfigurationOptionsResponse, *common.Fault)

	OptGetVideoOutputs(GetVideoOutputs GetVideoOutputs) (*GetVideoOutputsResponse, *common.Fault)

	OptGetVideoSourceConfiguration(GetVideoSourceConfiguration GetVideoSourceConfiguration) (*GetVideoSourceConfigurationResponse, *common.Fault)

	OptGetVideoSourceConfigurationOptions(GetVideoSourceConfigurationOptions GetVideoSourceConfigurationOptions) (*GetVideoSourceConfigurationOptionsResponse, *common.Fault)

	OptGetVideoSources(GetVideoSources Get) (*GetResponse, *common.Fault)

	OptSendReceiveSerialCommand(SendReceiveSerialCommand SendReceiveSerialCommand) (*SendReceiveSerialCommandResponse, *common.Fault)

	OptSetAudioOutputConfiguration(SetAudioOutputConfiguration SetAudioOutputConfiguration) (*SetAudioOutputConfigurationResponse, *common.Fault)

	OptSetAudioSourceConfiguration(SetAudioSourceConfiguration SetAudioSourceConfiguration) (*SetAudioSourceConfigurationResponse, *common.Fault)

	OptSetDigitalInputConfigurations(SetDigitalInputConfigurations SetDigitalInputConfigurations) (*SetDigitalInputConfigurationsResponse, *common.Fault)

	OptSetRelayOutputSettings(SetRelayOutputSettings SetRelayOutputSettings) (*SetRelayOutputSettingsResponse, *common.Fault)

	OptSetRelayOutputState(SetRelayOutputState device.SetRelayOutputState) (*device.SetRelayOutputStateResponse, *common.Fault)

	OptSetSerialPortConfiguration(SetSerialPortConfiguration SetSerialPortConfiguration) (*SetSerialPortConfigurationResponse, *common.Fault)

	OptSetVideoOutputConfiguration(SetVideoOutputConfiguration SetVideoOutputConfiguration) (*SetVideoOutputConfigurationResponse, *common.Fault)

	OptSetVideoSourceConfiguration(SetVideoSourceConfiguration SetVideoSourceConfiguration) (*SetVideoSourceConfigurationResponse, *common.Fault)
}
type Duration string

type ParityBit string

func (v ParityBit) Validate() bool {
	for _, vv := range []string{
		"None",
		"Even",
		"Odd",
		"Mark",
		"Space",
		"Extended",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type SerialPortType string

func (v SerialPortType) Validate() bool {
	for _, vv := range []string{
		"RS232",
		"RS422HalfDuplex",
		"RS422FullDuplex",
		"RS485HalfDuplex",
		"RS485FullDuplex",
		"Generic",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type Capabilities []interface{}

type DigitalInputConfigurationOptions struct {
	IdleState []*common.DigitalIdleState `xml:"IdleState,omitempty" json:"IdleState,omitempty" yaml:"IdleState,omitempty"`
}

type Get struct {
}

type GetAudioOutputConfiguration struct {
	AudioOutputToken *common.ReferenceToken `xml:"AudioOutputToken,omitempty" json:"AudioOutputToken,omitempty" yaml:"AudioOutputToken,omitempty"`
}

type GetAudioOutputConfigurationOptions struct {
	AudioOutputToken *common.ReferenceToken `xml:"AudioOutputToken,omitempty" json:"AudioOutputToken,omitempty" yaml:"AudioOutputToken,omitempty"`
}

type GetAudioOutputConfigurationOptionsResponse struct {
	AudioOutputOptions *common.AudioOutputConfigurationOptions `xml:"AudioOutputOptions,omitempty" json:"AudioOutputOptions,omitempty" yaml:"AudioOutputOptions,omitempty"`
}

type GetAudioOutputConfigurationResponse struct {
	AudioOutputConfiguration *common.AudioOutputConfiguration `xml:"AudioOutputConfiguration,omitempty" json:"AudioOutputConfiguration,omitempty" yaml:"AudioOutputConfiguration,omitempty"`
}

type GetAudioSourceConfiguration struct {
	AudioSourceToken *common.ReferenceToken `xml:"AudioSourceToken,omitempty" json:"AudioSourceToken,omitempty" yaml:"AudioSourceToken,omitempty"`
}

type GetAudioSourceConfigurationOptions struct {
	AudioSourceToken *common.ReferenceToken `xml:"AudioSourceToken,omitempty" json:"AudioSourceToken,omitempty" yaml:"AudioSourceToken,omitempty"`
}

type GetAudioSourceConfigurationOptionsResponse struct {
	AudioSourceOptions *common.AudioSourceConfigurationOptions `xml:"AudioSourceOptions,omitempty" json:"AudioSourceOptions,omitempty" yaml:"AudioSourceOptions,omitempty"`
}

type GetAudioSourceConfigurationResponse struct {
	AudioSourceConfiguration *common.AudioSourceConfiguration `xml:"AudioSourceConfiguration,omitempty" json:"AudioSourceConfiguration,omitempty" yaml:"AudioSourceConfiguration,omitempty"`
}

type GetDigitalInputConfigurationOptions struct {
	Token *common.ReferenceToken `xml:"Token,omitempty" json:"Token,omitempty" yaml:"Token,omitempty"`
}

type GetDigitalInputConfigurationOptionsResponse struct {
	DigitalInputOptions DigitalInputConfigurationOptions `xml:"DigitalInputOptions,omitempty" json:"DigitalInputOptions,omitempty" yaml:"DigitalInputOptions,omitempty"`
}

type GetDigitalInputs struct {
}

type GetDigitalInputsResponse struct {
	DigitalInputs []*common.DigitalInput `xml:"DigitalInputs,omitempty" json:"DigitalInputs,omitempty" yaml:"DigitalInputs,omitempty"`
}

type GetRelayOutputOptions struct {
	RelayOutputToken *common.ReferenceToken `xml:"RelayOutputToken,omitempty" json:"RelayOutputToken,omitempty" yaml:"RelayOutputToken,omitempty"`
}

type GetRelayOutputOptionsResponse struct {
	RelayOutputOptions []RelayOutputOptions `xml:"RelayOutputOptions,omitempty" json:"RelayOutputOptions,omitempty" yaml:"RelayOutputOptions,omitempty"`
}

type GetResponse struct {
	Token []*common.ReferenceToken `xml:"Token,omitempty" json:"Token,omitempty" yaml:"Token,omitempty"`
}

type GetSerialPortConfiguration struct {
	SerialPortToken *common.ReferenceToken `xml:"SerialPortToken,omitempty" json:"SerialPortToken,omitempty" yaml:"SerialPortToken,omitempty"`
}

type GetSerialPortConfigurationOptions struct {
	SerialPortToken *common.ReferenceToken `xml:"SerialPortToken,omitempty" json:"SerialPortToken,omitempty" yaml:"SerialPortToken,omitempty"`
}

type GetSerialPortConfigurationOptionsResponse struct {
	SerialPortOptions SerialPortConfigurationOptions `xml:"SerialPortOptions,omitempty" json:"SerialPortOptions,omitempty" yaml:"SerialPortOptions,omitempty"`
}

type GetSerialPortConfigurationResponse struct {
	SerialPortConfiguration SerialPortConfiguration `xml:"SerialPortConfiguration,omitempty" json:"SerialPortConfiguration,omitempty" yaml:"SerialPortConfiguration,omitempty"`
}

type GetSerialPorts struct {
}

type GetSerialPortsResponse struct {
	SerialPort []SerialPort `xml:"SerialPort,omitempty" json:"SerialPort,omitempty" yaml:"SerialPort,omitempty"`
}

type GetServiceCapabilities struct {
}

type GetServiceCapabilitiesResponse struct {
	Capabilities Capabilities `xml:"Capabilities,omitempty" json:"Capabilities,omitempty" yaml:"Capabilities,omitempty"`
}

type GetVideoOutputConfiguration struct {
	VideoOutputToken *common.ReferenceToken `xml:"VideoOutputToken,omitempty" json:"VideoOutputToken,omitempty" yaml:"VideoOutputToken,omitempty"`
}

type GetVideoOutputConfigurationOptions struct {
	VideoOutputToken *common.ReferenceToken `xml:"VideoOutputToken,omitempty" json:"VideoOutputToken,omitempty" yaml:"VideoOutputToken,omitempty"`
}

type GetVideoOutputConfigurationOptionsResponse struct {
	VideoOutputConfigurationOptions *common.VideoOutputConfigurationOptions `xml:"VideoOutputConfigurationOptions,omitempty" json:"VideoOutputConfigurationOptions,omitempty" yaml:"VideoOutputConfigurationOptions,omitempty"`
}

type GetVideoOutputConfigurationResponse struct {
	VideoOutputConfiguration *common.VideoOutputConfiguration `xml:"VideoOutputConfiguration,omitempty" json:"VideoOutputConfiguration,omitempty" yaml:"VideoOutputConfiguration,omitempty"`
}

type GetVideoOutputs struct {
}

type GetVideoOutputsResponse struct {
	VideoOutputs []*common.VideoOutput `xml:"VideoOutputs,omitempty" json:"VideoOutputs,omitempty" yaml:"VideoOutputs,omitempty"`
}

type GetVideoSourceConfiguration struct {
	VideoSourceToken *common.ReferenceToken `xml:"VideoSourceToken,omitempty" json:"VideoSourceToken,omitempty" yaml:"VideoSourceToken,omitempty"`
}

type GetVideoSourceConfigurationOptions struct {
	VideoSourceToken *common.ReferenceToken `xml:"VideoSourceToken,omitempty" json:"VideoSourceToken,omitempty" yaml:"VideoSourceToken,omitempty"`
}

type GetVideoSourceConfigurationOptionsResponse struct {
	VideoSourceConfigurationOptions *common.VideoSourceConfigurationOptions `xml:"VideoSourceConfigurationOptions,omitempty" json:"VideoSourceConfigurationOptions,omitempty" yaml:"VideoSourceConfigurationOptions,omitempty"`
}

type GetVideoSourceConfigurationResponse struct {
	VideoSourceConfiguration *common.VideoSourceConfiguration `xml:"VideoSourceConfiguration,omitempty" json:"VideoSourceConfiguration,omitempty" yaml:"VideoSourceConfiguration,omitempty"`
}

type ParityBitList struct {
	Items []ParityBit `xml:"Items,omitempty" json:"Items,omitempty" yaml:"Items,omitempty"`
}

type RelayOutputOptions struct {
	Mode       []*common.RelayMode         `xml:"Mode,omitempty" json:"Mode,omitempty" yaml:"Mode,omitempty"`
	DelayTimes *common.FloatList           `xml:"DelayTimes,omitempty" json:"DelayTimes,omitempty" yaml:"DelayTimes,omitempty"`
	Discrete   bool                        `xml:"Discrete,omitempty" json:"Discrete,omitempty" yaml:"Discrete,omitempty"`
	Extension  RelayOutputOptionsExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
	Token      common.ReferenceToken       `xml:"token,attr,omitempty" json:"token,attr,omitempty" yaml:"token,attr,omitempty"`
}

type RelayOutputOptionsExtension []interface{}

type SendReceiveSerialCommand struct {
	Token      *common.ReferenceToken `xml:"Token,omitempty" json:"Token,omitempty" yaml:"Token,omitempty"`
	SerialData SerialData             `xml:"SerialData,omitempty" json:"SerialData,omitempty" yaml:"SerialData,omitempty"`
	TimeOut    *common.Duration       `xml:"TimeOut,omitempty" json:"TimeOut,omitempty" yaml:"TimeOut,omitempty"`
	DataLength int64                  `xml:"DataLength,omitempty" json:"DataLength,omitempty" yaml:"DataLength,omitempty"`
	Delimiter  string                 `xml:"Delimiter,omitempty" json:"Delimiter,omitempty" yaml:"Delimiter,omitempty"`
}

type SendReceiveSerialCommandResponse struct {
	SerialData SerialData `xml:"SerialData,omitempty" json:"SerialData,omitempty" yaml:"SerialData,omitempty"`
}

type SerialData struct {
	Binary []byte `xml:"Binary,omitempty" json:"Binary,omitempty" yaml:"Binary,omitempty"`
	String string `xml:"String,omitempty" json:"String,omitempty" yaml:"String,omitempty"`
}

type SerialPort struct {
	TypeAttrXSI   string `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *SerialPort) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:SerialPort"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://www.onvif.org/ver10/deviceIO/wsdl"
	}
}

type SerialPortConfiguration struct {
	BaudRate        int                   `xml:"BaudRate,omitempty" json:"BaudRate,omitempty" yaml:"BaudRate,omitempty"`
	ParityBit       ParityBit             `xml:"ParityBit,omitempty" json:"ParityBit,omitempty" yaml:"ParityBit,omitempty"`
	CharacterLength int                   `xml:"CharacterLength,omitempty" json:"CharacterLength,omitempty" yaml:"CharacterLength,omitempty"`
	StopBit         float64               `xml:"StopBit,omitempty" json:"StopBit,omitempty" yaml:"StopBit,omitempty"`
	Token           common.ReferenceToken `xml:"token,attr,omitempty" json:"token,attr,omitempty" yaml:"token,attr,omitempty"`
	Type            SerialPortType        `xml:"type,attr,omitempty" json:"type,attr,omitempty" yaml:"type,attr,omitempty"`
}

type SerialPortConfigurationOptions struct {
	BaudRateList        *common.IntItems      `xml:"BaudRateList,omitempty" json:"BaudRateList,omitempty" yaml:"BaudRateList,omitempty"`
	ParityBitList       ParityBitList         `xml:"ParityBitList,omitempty" json:"ParityBitList,omitempty" yaml:"ParityBitList,omitempty"`
	CharacterLengthList *common.IntItems      `xml:"CharacterLengthList,omitempty" json:"CharacterLengthList,omitempty" yaml:"CharacterLengthList,omitempty"`
	StopBitList         *common.FloatItems    `xml:"StopBitList,omitempty" json:"StopBitList,omitempty" yaml:"StopBitList,omitempty"`
	Token               common.ReferenceToken `xml:"token,attr,omitempty" json:"token,attr,omitempty" yaml:"token,attr,omitempty"`
}

type SetAudioOutputConfiguration struct {
	Configuration    *common.AudioOutputConfiguration `xml:"Configuration,omitempty" json:"Configuration,omitempty" yaml:"Configuration,omitempty"`
	ForcePersistence bool                             `xml:"ForcePersistence,omitempty" json:"ForcePersistence,omitempty" yaml:"ForcePersistence,omitempty"`
}

type SetAudioOutputConfigurationResponse []interface{}

type SetAudioSourceConfiguration struct {
	Configuration    *common.AudioSourceConfiguration `xml:"Configuration,omitempty" json:"Configuration,omitempty" yaml:"Configuration,omitempty"`
	ForcePersistence bool                             `xml:"ForcePersistence,omitempty" json:"ForcePersistence,omitempty" yaml:"ForcePersistence,omitempty"`
}

type SetAudioSourceConfigurationResponse []interface{}

type SetDigitalInputConfigurations struct {
	DigitalInputs []*common.DigitalInput `xml:"DigitalInputs,omitempty" json:"DigitalInputs,omitempty" yaml:"DigitalInputs,omitempty"`
}

type SetDigitalInputConfigurationsResponse struct {
}

type SetRelayOutputSettings struct {
	RelayOutput *common.RelayOutput `xml:"RelayOutput,omitempty" json:"RelayOutput,omitempty" yaml:"RelayOutput,omitempty"`
}

type SetRelayOutputSettingsResponse struct {
}

type SetSerialPortConfiguration struct {
	SerialPortConfiguration SerialPortConfiguration `xml:"SerialPortConfiguration,omitempty" json:"SerialPortConfiguration,omitempty" yaml:"SerialPortConfiguration,omitempty"`
	ForcePersistance        bool                    `xml:"ForcePersistance,omitempty" json:"ForcePersistance,omitempty" yaml:"ForcePersistance,omitempty"`
}

type SetSerialPortConfigurationResponse struct {
}

type SetVideoOutputConfiguration struct {
	Configuration    *common.VideoOutputConfiguration `xml:"Configuration,omitempty" json:"Configuration,omitempty" yaml:"Configuration,omitempty"`
	ForcePersistence bool                             `xml:"ForcePersistence,omitempty" json:"ForcePersistence,omitempty" yaml:"ForcePersistence,omitempty"`
}

type SetVideoOutputConfigurationResponse []interface{}

type SetVideoSourceConfiguration struct {
	Configuration    *common.VideoSourceConfiguration `xml:"Configuration,omitempty" json:"Configuration,omitempty" yaml:"Configuration,omitempty"`
	ForcePersistence bool                             `xml:"ForcePersistence,omitempty" json:"ForcePersistence,omitempty" yaml:"ForcePersistence,omitempty"`
}

type SetVideoSourceConfigurationResponse []interface{}

// deviceIOPort implements the DeviceIOPort interface.
type deviceIOPort struct {
	cli      common.Client
	Endpoint string
}

func (p *deviceIOPort) OptGetAudioOutputConfiguration(args GetAudioOutputConfiguration) (*GetAudioOutputConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                     string `xml:"tmd:GetAudioOutputConfiguration"`
		GetAudioOutputConfiguration GetAudioOutputConfiguration
	}{
		GetAudioOutputConfiguration: args,
	}

	resp := GetAudioOutputConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *deviceIOPort) OptGetAudioOutputConfigurationOptions(args GetAudioOutputConfigurationOptions) (*GetAudioOutputConfigurationOptionsResponse, *common.Fault) {
	req := struct {
		XMLName                            string `xml:"tmd:GetAudioOutputConfigurationOptions"`
		GetAudioOutputConfigurationOptions GetAudioOutputConfigurationOptions
	}{
		GetAudioOutputConfigurationOptions: args,
	}

	resp := GetAudioOutputConfigurationOptionsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *deviceIOPort) OptGetAudioOutputs(args Get) (*GetResponse, *common.Fault) {
	req := struct {
		XMLName string `xml:"tmd:GetAudioOutputs"`
		Get     Get
	}{
		Get: args,
	}

	resp := GetResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *deviceIOPort) OptGetAudioSourceConfiguration(args GetAudioSourceConfiguration) (*GetAudioSourceConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                     string `xml:"tmd:GetAudioSourceConfiguration"`
		GetAudioSourceConfiguration GetAudioSourceConfiguration
	}{
		GetAudioSourceConfiguration: args,
	}

	resp := GetAudioSourceConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *deviceIOPort) OptGetAudioSourceConfigurationOptions(args GetAudioSourceConfigurationOptions) (*GetAudioSourceConfigurationOptionsResponse, *common.Fault) {
	req := struct {
		XMLName                            string `xml:"tmd:GetAudioSourceConfigurationOptions"`
		GetAudioSourceConfigurationOptions GetAudioSourceConfigurationOptions
	}{
		GetAudioSourceConfigurationOptions: args,
	}

	resp := GetAudioSourceConfigurationOptionsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *deviceIOPort) OptGetAudioSources(args Get) (*GetResponse, *common.Fault) {
	req := struct {
		XMLName string `xml:"tmd:GetAudioSources"`
		Get     Get
	}{
		Get: args,
	}

	resp := GetResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *deviceIOPort) OptGetDigitalInputConfigurationOptions(args GetDigitalInputConfigurationOptions) (*GetDigitalInputConfigurationOptionsResponse, *common.Fault) {
	req := struct {
		XMLName                             string `xml:"tmd:GetDigitalInputConfigurationOptions"`
		GetDigitalInputConfigurationOptions GetDigitalInputConfigurationOptions
	}{
		GetDigitalInputConfigurationOptions: args,
	}

	resp := GetDigitalInputConfigurationOptionsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *deviceIOPort) OptGetDigitalInputs(args GetDigitalInputs) (*GetDigitalInputsResponse, *common.Fault) {
	req := struct {
		XMLName          string `xml:"tmd:GetDigitalInputs"`
		GetDigitalInputs GetDigitalInputs
	}{
		GetDigitalInputs: args,
	}

	resp := GetDigitalInputsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *deviceIOPort) OptGetRelayOutputOptions(args GetRelayOutputOptions) (*GetRelayOutputOptionsResponse, *common.Fault) {
	req := struct {
		XMLName               string `xml:"tmd:GetRelayOutputOptions"`
		GetRelayOutputOptions GetRelayOutputOptions
	}{
		GetRelayOutputOptions: args,
	}

	resp := GetRelayOutputOptionsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *deviceIOPort) OptGetRelayOutputs(args device.GetRelayOutputs) (*device.GetRelayOutputsResponse, *common.Fault) {
	req := struct {
		XMLName         string `xml:"tmd:GetRelayOutputs"`
		GetRelayOutputs device.GetRelayOutputs
	}{
		GetRelayOutputs: args,
	}

	resp := device.GetRelayOutputsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *deviceIOPort) OptGetSerialPortConfiguration(args GetSerialPortConfiguration) (*GetSerialPortConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                    string `xml:"tmd:GetSerialPortConfiguration"`
		GetSerialPortConfiguration GetSerialPortConfiguration
	}{
		GetSerialPortConfiguration: args,
	}

	resp := GetSerialPortConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *deviceIOPort) OptGetSerialPortConfigurationOptions(args GetSerialPortConfigurationOptions) (*GetSerialPortConfigurationOptionsResponse, *common.Fault) {
	req := struct {
		XMLName                           string `xml:"tmd:GetSerialPortConfigurationOptions"`
		GetSerialPortConfigurationOptions GetSerialPortConfigurationOptions
	}{
		GetSerialPortConfigurationOptions: args,
	}

	resp := GetSerialPortConfigurationOptionsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *deviceIOPort) OptGetSerialPorts(args GetSerialPorts) (*GetSerialPortsResponse, *common.Fault) {
	req := struct {
		XMLName        string `xml:"tmd:GetSerialPorts"`
		GetSerialPorts GetSerialPorts
	}{
		GetSerialPorts: args,
	}

	resp := GetSerialPortsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *deviceIOPort) OptGetServiceCapabilities(args GetServiceCapabilities) (*GetServiceCapabilitiesResponse, *common.Fault) {
	req := struct {
		XMLName                string `xml:"tmd:GetServiceCapabilities"`
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

func (p *deviceIOPort) OptGetVideoOutputConfiguration(args GetVideoOutputConfiguration) (*GetVideoOutputConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                     string `xml:"tmd:GetVideoOutputConfiguration"`
		GetVideoOutputConfiguration GetVideoOutputConfiguration
	}{
		GetVideoOutputConfiguration: args,
	}

	resp := GetVideoOutputConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *deviceIOPort) OptGetVideoOutputConfigurationOptions(args GetVideoOutputConfigurationOptions) (*GetVideoOutputConfigurationOptionsResponse, *common.Fault) {
	req := struct {
		XMLName                            string `xml:"tmd:GetVideoOutputConfigurationOptions"`
		GetVideoOutputConfigurationOptions GetVideoOutputConfigurationOptions
	}{
		GetVideoOutputConfigurationOptions: args,
	}

	resp := GetVideoOutputConfigurationOptionsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *deviceIOPort) OptGetVideoOutputs(args GetVideoOutputs) (*GetVideoOutputsResponse, *common.Fault) {
	req := struct {
		XMLName         string `xml:"tmd:GetVideoOutputs"`
		GetVideoOutputs GetVideoOutputs
	}{
		GetVideoOutputs: args,
	}

	resp := GetVideoOutputsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *deviceIOPort) OptGetVideoSourceConfiguration(args GetVideoSourceConfiguration) (*GetVideoSourceConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                     string `xml:"tmd:GetVideoSourceConfiguration"`
		GetVideoSourceConfiguration GetVideoSourceConfiguration
	}{
		GetVideoSourceConfiguration: args,
	}

	resp := GetVideoSourceConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *deviceIOPort) OptGetVideoSourceConfigurationOptions(args GetVideoSourceConfigurationOptions) (*GetVideoSourceConfigurationOptionsResponse, *common.Fault) {
	req := struct {
		XMLName                            string `xml:"tmd:GetVideoSourceConfigurationOptions"`
		GetVideoSourceConfigurationOptions GetVideoSourceConfigurationOptions
	}{
		GetVideoSourceConfigurationOptions: args,
	}

	resp := GetVideoSourceConfigurationOptionsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *deviceIOPort) OptGetVideoSources(args Get) (*GetResponse, *common.Fault) {
	req := struct {
		XMLName string `xml:"tmd:GetVideoSources"`
		Get     Get
	}{
		Get: args,
	}

	resp := GetResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *deviceIOPort) OptSendReceiveSerialCommand(args SendReceiveSerialCommand) (*SendReceiveSerialCommandResponse, *common.Fault) {
	req := struct {
		XMLName                  string `xml:"tmd:SendReceiveSerialCommand"`
		SendReceiveSerialCommand SendReceiveSerialCommand
	}{
		SendReceiveSerialCommand: args,
	}

	resp := SendReceiveSerialCommandResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *deviceIOPort) OptSetAudioOutputConfiguration(args SetAudioOutputConfiguration) (*SetAudioOutputConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                     string `xml:"tmd:SetAudioOutputConfiguration"`
		SetAudioOutputConfiguration SetAudioOutputConfiguration
	}{
		SetAudioOutputConfiguration: args,
	}

	resp := SetAudioOutputConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *deviceIOPort) OptSetAudioSourceConfiguration(args SetAudioSourceConfiguration) (*SetAudioSourceConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                     string `xml:"tmd:SetAudioSourceConfiguration"`
		SetAudioSourceConfiguration SetAudioSourceConfiguration
	}{
		SetAudioSourceConfiguration: args,
	}

	resp := SetAudioSourceConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *deviceIOPort) OptSetDigitalInputConfigurations(args SetDigitalInputConfigurations) (*SetDigitalInputConfigurationsResponse, *common.Fault) {
	req := struct {
		XMLName                       string `xml:"tmd:SetDigitalInputConfigurations"`
		SetDigitalInputConfigurations SetDigitalInputConfigurations
	}{
		SetDigitalInputConfigurations: args,
	}

	resp := SetDigitalInputConfigurationsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *deviceIOPort) OptSetRelayOutputSettings(args SetRelayOutputSettings) (*SetRelayOutputSettingsResponse, *common.Fault) {
	req := struct {
		XMLName                string `xml:"tmd:SetRelayOutputSettings"`
		SetRelayOutputSettings SetRelayOutputSettings
	}{
		SetRelayOutputSettings: args,
	}

	resp := SetRelayOutputSettingsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *deviceIOPort) OptSetRelayOutputState(args device.SetRelayOutputState) (*device.SetRelayOutputStateResponse, *common.Fault) {
	req := struct {
		XMLName             string `xml:"tmd:SetRelayOutputState"`
		SetRelayOutputState device.SetRelayOutputState
	}{
		SetRelayOutputState: args,
	}

	resp := device.SetRelayOutputStateResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *deviceIOPort) OptSetSerialPortConfiguration(args SetSerialPortConfiguration) (*SetSerialPortConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                    string `xml:"tmd:SetSerialPortConfiguration"`
		SetSerialPortConfiguration SetSerialPortConfiguration
	}{
		SetSerialPortConfiguration: args,
	}

	resp := SetSerialPortConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *deviceIOPort) OptSetVideoOutputConfiguration(args SetVideoOutputConfiguration) (*SetVideoOutputConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                     string `xml:"tmd:SetVideoOutputConfiguration"`
		SetVideoOutputConfiguration SetVideoOutputConfiguration
	}{
		SetVideoOutputConfiguration: args,
	}

	resp := SetVideoOutputConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *deviceIOPort) OptSetVideoSourceConfiguration(args SetVideoSourceConfiguration) (*SetVideoSourceConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                     string `xml:"tmd:SetVideoSourceConfiguration"`
		SetVideoSourceConfiguration SetVideoSourceConfiguration
	}{
		SetVideoSourceConfiguration: args,
	}

	resp := SetVideoSourceConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}
