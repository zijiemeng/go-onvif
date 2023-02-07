package media

import (
	"reflect"

	"github.com/zijiemeng/go-onvif/onvif/common"
)

var Namespace = "http://www.onvif.org/ver20/media/wsdl"

// NewMedia2 creates an initializes a Media2.
func NewMedia2(endpoint string, cli common.Client) Media2 {
	return &media2{cli: cli, Endpoint: endpoint}
}

// Media2 was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type Media2 interface {
	OptAddConfiguration(AddConfiguration AddConfiguration) (*AddConfigurationResponse, *common.Fault)

	OptCreateMask(CreateMask CreateMask) (*CreateMaskResponse, *common.Fault)

	OptCreateOSD(CreateOSD CreateOSD) (*CreateOSDResponse, *common.Fault)

	OptCreateProfile(CreateProfile CreateProfile) (*CreateProfileResponse, *common.Fault)

	OptDeleteMask(DeleteMask DeleteMask) (*SetConfigurationResponse, *common.Fault)

	OptDeleteOSD(DeleteOSD DeleteOSD) (*SetConfigurationResponse, *common.Fault)

	OptDeleteProfile(DeleteProfile DeleteProfile) (*DeleteProfileResponse, *common.Fault)

	OptGetAnalyticsConfigurations(GetAnalyticsConfigurations GetConfiguration) (*GetAnalyticsConfigurationsResponse, *common.Fault)

	OptGetAudioDecoderConfigurationOptions(GetAudioDecoderConfigurationOptions GetConfiguration) (*GetAudioDecoderConfigurationOptionsResponse, *common.Fault)

	OptGetAudioDecoderConfigurations(GetAudioDecoderConfigurations GetConfiguration) (*GetAudioDecoderConfigurationsResponse, *common.Fault)

	OptGetAudioEncoderConfigurationOptions(GetAudioEncoderConfigurationOptions GetConfiguration) (*GetAudioEncoderConfigurationOptionsResponse, *common.Fault)

	OptGetAudioEncoderConfigurations(GetAudioEncoderConfigurations GetConfiguration) (*GetAudioEncoderConfigurationsResponse, *common.Fault)

	OptGetAudioOutputConfigurationOptions(GetAudioOutputConfigurationOptions GetConfiguration) (*GetAudioOutputConfigurationOptionsResponse, *common.Fault)

	OptGetAudioOutputConfigurations(GetAudioOutputConfigurations GetConfiguration) (*GetAudioOutputConfigurationsResponse, *common.Fault)

	OptGetAudioSourceConfigurationOptions(GetAudioSourceConfigurationOptions GetConfiguration) (*GetAudioSourceConfigurationOptionsResponse, *common.Fault)

	OptGetAudioSourceConfigurations(GetAudioSourceConfigurations GetConfiguration) (*GetAudioSourceConfigurationsResponse, *common.Fault)

	OptGetMaskOptions(GetMaskOptions GetMaskOptions) (*GetMaskOptionsResponse, *common.Fault)

	OptGetMasks(GetMasks GetMasks) (*GetMasksResponse, *common.Fault)

	OptGetMetadataConfigurationOptions(GetMetadataConfigurationOptions GetConfiguration) (*GetMetadataConfigurationOptionsResponse, *common.Fault)

	OptGetMetadataConfigurations(GetMetadataConfigurations GetConfiguration) (*GetMetadataConfigurationsResponse, *common.Fault)

	OptGetOSDOptions(GetOSDOptions GetOSDOptions) (*GetOSDOptionsResponse, *common.Fault)

	OptGetOSDs(GetOSDs GetOSDs) (*GetOSDsResponse, *common.Fault)

	OptGetProfiles(GetProfiles GetProfiles) (*GetProfilesResponse, *common.Fault)

	OptGetServiceCapabilities(GetServiceCapabilities GetServiceCapabilities) (*GetServiceCapabilitiesResponse, *common.Fault)

	OptGetSnapshotUri(GetSnapshotUri GetSnapshotUri) (*GetSnapshotUriResponse, *common.Fault)

	OptGetStreamUri(GetStreamUri GetStreamUri) (*GetStreamUriResponse, *common.Fault)

	OptGetVideoEncoderConfigurationOptions(GetVideoEncoderConfigurationOptions GetConfiguration) (*GetVideoEncoderConfigurationOptionsResponse, *common.Fault)

	OptGetVideoEncoderConfigurations(GetVideoEncoderConfigurations GetConfiguration) (*GetVideoEncoderConfigurationsResponse, *common.Fault)

	OptGetVideoEncoderInstances(GetVideoEncoderInstances GetVideoEncoderInstances) (*GetVideoEncoderInstancesResponse, *common.Fault)

	OptGetVideoSourceConfigurationOptions(GetVideoSourceConfigurationOptions GetConfiguration) (*GetVideoSourceConfigurationOptionsResponse, *common.Fault)

	OptGetVideoSourceConfigurations(GetVideoSourceConfigurations GetConfiguration) (*GetVideoSourceConfigurationsResponse, *common.Fault)

	OptGetVideoSourceModes(GetVideoSourceModes GetVideoSourceModes) (*GetVideoSourceModesResponse, *common.Fault)

	OptRemoveConfiguration(RemoveConfiguration RemoveConfiguration) (*RemoveConfigurationResponse, *common.Fault)

	OptSetAudioDecoderConfiguration(SetAudioDecoderConfiguration SetAudioDecoderConfiguration) (*SetConfigurationResponse, *common.Fault)

	OptSetAudioEncoderConfiguration(SetAudioEncoderConfiguration SetAudioEncoderConfiguration) (*SetConfigurationResponse, *common.Fault)

	OptSetAudioOutputConfiguration(SetAudioOutputConfiguration SetAudioOutputConfiguration) (*SetConfigurationResponse, *common.Fault)

	OptSetAudioSourceConfiguration(SetAudioSourceConfiguration SetAudioSourceConfiguration) (*SetConfigurationResponse, *common.Fault)

	OptSetMask(SetMask SetMask) (*SetConfigurationResponse, *common.Fault)

	OptSetMetadataConfiguration(SetMetadataConfiguration SetMetadataConfiguration) (*SetConfigurationResponse, *common.Fault)

	OptSetOSD(SetOSD SetOSD) (*SetConfigurationResponse, *common.Fault)

	OptSetSynchronizationPoint(SetSynchronizationPoint SetSynchronizationPoint) (*SetSynchronizationPointResponse, *common.Fault)

	OptSetVideoEncoderConfiguration(SetVideoEncoderConfiguration SetVideoEncoderConfiguration) (*SetConfigurationResponse, *common.Fault)

	OptSetVideoSourceConfiguration(SetVideoSourceConfiguration SetVideoSourceConfiguration) (*SetConfigurationResponse, *common.Fault)

	OptSetVideoSourceMode(SetVideoSourceMode SetVideoSourceMode) (*SetVideoSourceModeResponse, *common.Fault)

	OptStartMulticastStreaming(StartMulticastStreaming StartStopMulticastStreaming) (*SetConfigurationResponse, *common.Fault)

	OptStopMulticastStreaming(StopMulticastStreaming StartStopMulticastStreaming) (*SetConfigurationResponse, *common.Fault)
}
type ConfigurationEnumeration string

func (v ConfigurationEnumeration) Validate() bool {
	for _, vv := range []string{
		"All",
		"VideoSource",
		"VideoEncoder",
		"AudioSource",
		"AudioEncoder",
		"AudioOutput",
		"AudioDecoder",
		"Metadata",
		"Analytics",
		"PTZ",
		"Receiver",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type MaskType string

func (v MaskType) Validate() bool {
	for _, vv := range []string{
		"Color",
		"Pixelated",
		"Blurred",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type TransportProtocol string

func (v TransportProtocol) Validate() bool {
	for _, vv := range []string{
		"RtspUnicast",
		"RtspMulticast",
		"RtspsUnicast",
		"RtspsMulticast",
		"RTSP",
		"RtspOverHttp",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type AddConfiguration struct {
	ProfileToken  *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
	Name          *common.Name           `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Configuration []ConfigurationRef     `xml:"Configuration,omitempty" json:"Configuration,omitempty" yaml:"Configuration,omitempty"`
}

type AddConfigurationResponse struct {
}

type Capabilities2 struct {
	ProfileCapabilities   ProfileCapabilities   `xml:"ProfileCapabilities,omitempty" json:"ProfileCapabilities,omitempty" yaml:"ProfileCapabilities,omitempty"`
	StreamingCapabilities StreamingCapabilities `xml:"StreamingCapabilities,omitempty" json:"StreamingCapabilities,omitempty" yaml:"StreamingCapabilities,omitempty"`
	SnapshotUri           bool                  `xml:"SnapshotUri,attr,omitempty" json:"SnapshotUri,attr,omitempty" yaml:"SnapshotUri,attr,omitempty"`
	Rotation              bool                  `xml:"Rotation,attr,omitempty" json:"Rotation,attr,omitempty" yaml:"Rotation,attr,omitempty"`
	VideoSourceMode       bool                  `xml:"VideoSourceMode,attr,omitempty" json:"VideoSourceMode,attr,omitempty" yaml:"VideoSourceMode,attr,omitempty"`
	OSD                   bool                  `xml:"OSD,attr,omitempty" json:"OSD,attr,omitempty" yaml:"OSD,attr,omitempty"`
	TemporaryOSDText      bool                  `xml:"TemporaryOSDText,attr,omitempty" json:"TemporaryOSDText,attr,omitempty" yaml:"TemporaryOSDText,attr,omitempty"`
	Mask                  bool                  `xml:"Mask,attr,omitempty" json:"Mask,attr,omitempty" yaml:"Mask,attr,omitempty"`
	SourceMask            bool                  `xml:"SourceMask,attr,omitempty" json:"SourceMask,attr,omitempty" yaml:"SourceMask,attr,omitempty"`
}

type ConfigurationRef struct {
	Type  string                 `xml:"Type,omitempty" json:"Type,omitempty" yaml:"Type,omitempty"`
	Token *common.ReferenceToken `xml:"Token,omitempty" json:"Token,omitempty" yaml:"Token,omitempty"`
}

type ConfigurationSet struct {
	VideoSource  *common.VideoSourceConfiguration    `xml:"VideoSource,omitempty" json:"VideoSource,omitempty" yaml:"VideoSource,omitempty"`
	AudioSource  *common.AudioSourceConfiguration    `xml:"AudioSource,omitempty" json:"AudioSource,omitempty" yaml:"AudioSource,omitempty"`
	VideoEncoder *common.VideoEncoder2Configuration  `xml:"VideoEncoder,omitempty" json:"VideoEncoder,omitempty" yaml:"VideoEncoder,omitempty"`
	AudioEncoder *common.AudioEncoder2Configuration  `xml:"AudioEncoder,omitempty" json:"AudioEncoder,omitempty" yaml:"AudioEncoder,omitempty"`
	Analytics    *common.VideoAnalyticsConfiguration `xml:"Analytics,omitempty" json:"Analytics,omitempty" yaml:"Analytics,omitempty"`
	PTZ          *common.PTZConfiguration            `xml:"PTZ,omitempty" json:"PTZ,omitempty" yaml:"PTZ,omitempty"`
	Metadata     *common.MetadataConfiguration       `xml:"Metadata,omitempty" json:"Metadata,omitempty" yaml:"Metadata,omitempty"`
	AudioOutput  *common.AudioOutputConfiguration    `xml:"AudioOutput,omitempty" json:"AudioOutput,omitempty" yaml:"AudioOutput,omitempty"`
	AudioDecoder *common.AudioDecoderConfiguration   `xml:"AudioDecoder,omitempty" json:"AudioDecoder,omitempty" yaml:"AudioDecoder,omitempty"`
	Receiver     ReceiverConfiguration               `xml:"Receiver,omitempty" json:"Receiver,omitempty" yaml:"Receiver,omitempty"`
}

type CreateMask struct {
	Mask Mask `xml:"Mask,omitempty" json:"Mask,omitempty" yaml:"Mask,omitempty"`
}

type CreateMaskResponse struct {
	Token *common.ReferenceToken `xml:"Token,omitempty" json:"Token,omitempty" yaml:"Token,omitempty"`
}

type CreateOSD struct {
	OSD *common.OSDConfiguration `xml:"OSD,omitempty" json:"OSD,omitempty" yaml:"OSD,omitempty"`
}

type CreateOSDResponse struct {
	OSDToken *common.ReferenceToken `xml:"OSDToken,omitempty" json:"OSDToken,omitempty" yaml:"OSDToken,omitempty"`
}

type CreateProfile struct {
	Name          *common.Name       `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Configuration []ConfigurationRef `xml:"Configuration,omitempty" json:"Configuration,omitempty" yaml:"Configuration,omitempty"`
}

type CreateProfileResponse struct {
	Token *common.ReferenceToken `xml:"Token,omitempty" json:"Token,omitempty" yaml:"Token,omitempty"`
}

type DeleteMask struct {
	Token *common.ReferenceToken `xml:"Token,omitempty" json:"Token,omitempty" yaml:"Token,omitempty"`
}

type DeleteOSD struct {
	OSDToken *common.ReferenceToken `xml:"OSDToken,omitempty" json:"OSDToken,omitempty" yaml:"OSDToken,omitempty"`
}

type DeleteProfile struct {
	Token *common.ReferenceToken `xml:"Token,omitempty" json:"Token,omitempty" yaml:"Token,omitempty"`
}

type DeleteProfileResponse struct {
}

type EncoderInstance struct {
	Encoding string `xml:"Encoding,omitempty" json:"Encoding,omitempty" yaml:"Encoding,omitempty"`
	Number   int    `xml:"Number,omitempty" json:"Number,omitempty" yaml:"Number,omitempty"`
}

type EncoderInstanceInfo struct {
	Codec []EncoderInstance `xml:"Codec,omitempty" json:"Codec,omitempty" yaml:"Codec,omitempty"`
	Total int               `xml:"Total,omitempty" json:"Total,omitempty" yaml:"Total,omitempty"`
}

type GetAnalyticsConfigurationsResponse struct {
	Configurations []*common.VideoAnalyticsConfiguration `xml:"Configurations,omitempty" json:"Configurations,omitempty" yaml:"Configurations,omitempty"`
}

type GetAudioDecoderConfigurationOptionsResponse struct {
	Options []*common.AudioEncoder2ConfigurationOptions `xml:"Options,omitempty" json:"Options,omitempty" yaml:"Options,omitempty"`
}

type GetAudioDecoderConfigurationsResponse struct {
	Configurations []*common.AudioDecoderConfiguration `xml:"Configurations,omitempty" json:"Configurations,omitempty" yaml:"Configurations,omitempty"`
}

type GetAudioEncoderConfigurationOptionsResponse struct {
	Options []*common.AudioEncoder2ConfigurationOptions `xml:"Options,omitempty" json:"Options,omitempty" yaml:"Options,omitempty"`
}

type GetAudioEncoderConfigurationsResponse struct {
	Configurations []*common.AudioEncoder2Configuration `xml:"Configurations,omitempty" json:"Configurations,omitempty" yaml:"Configurations,omitempty"`
}

type GetAudioOutputConfigurationOptionsResponse struct {
	Options *common.AudioOutputConfigurationOptions `xml:"Options,omitempty" json:"Options,omitempty" yaml:"Options,omitempty"`
}

type GetAudioOutputConfigurationsResponse struct {
	Configurations []*common.AudioOutputConfiguration `xml:"Configurations,omitempty" json:"Configurations,omitempty" yaml:"Configurations,omitempty"`
}

type GetAudioSourceConfigurationOptionsResponse struct {
	Options *common.AudioSourceConfigurationOptions `xml:"Options,omitempty" json:"Options,omitempty" yaml:"Options,omitempty"`
}

type GetAudioSourceConfigurationsResponse struct {
	Configurations []*common.AudioSourceConfiguration `xml:"Configurations,omitempty" json:"Configurations,omitempty" yaml:"Configurations,omitempty"`
}

type GetConfiguration struct {
	ConfigurationToken *common.ReferenceToken `xml:"ConfigurationToken,omitempty" json:"ConfigurationToken,omitempty" yaml:"ConfigurationToken,omitempty"`
	ProfileToken       *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
}

type GetMaskOptions struct {
	ConfigurationToken *common.ReferenceToken `xml:"ConfigurationToken,omitempty" json:"ConfigurationToken,omitempty" yaml:"ConfigurationToken,omitempty"`
}

type GetMaskOptionsResponse struct {
	Options MaskOptions `xml:"Options,omitempty" json:"Options,omitempty" yaml:"Options,omitempty"`
}

type GetMasks struct {
	Token              *common.ReferenceToken `xml:"Token,omitempty" json:"Token,omitempty" yaml:"Token,omitempty"`
	ConfigurationToken *common.ReferenceToken `xml:"ConfigurationToken,omitempty" json:"ConfigurationToken,omitempty" yaml:"ConfigurationToken,omitempty"`
}

type GetMasksResponse struct {
	Masks []Mask `xml:"Masks,omitempty" json:"Masks,omitempty" yaml:"Masks,omitempty"`
}

type GetMetadataConfigurationOptionsResponse struct {
	Options *common.MetadataConfigurationOptions `xml:"Options,omitempty" json:"Options,omitempty" yaml:"Options,omitempty"`
}

type GetMetadataConfigurationsResponse struct {
	Configurations []*common.MetadataConfiguration `xml:"Configurations,omitempty" json:"Configurations,omitempty" yaml:"Configurations,omitempty"`
}

type GetOSDOptions struct {
	ConfigurationToken *common.ReferenceToken `xml:"ConfigurationToken,omitempty" json:"ConfigurationToken,omitempty" yaml:"ConfigurationToken,omitempty"`
}

type GetOSDOptionsResponse struct {
	OSDOptions *common.OSDConfigurationOptions `xml:"OSDOptions,omitempty" json:"OSDOptions,omitempty" yaml:"OSDOptions,omitempty"`
}

type GetOSDs struct {
	OSDToken           *common.ReferenceToken `xml:"OSDToken,omitempty" json:"OSDToken,omitempty" yaml:"OSDToken,omitempty"`
	ConfigurationToken *common.ReferenceToken `xml:"ConfigurationToken,omitempty" json:"ConfigurationToken,omitempty" yaml:"ConfigurationToken,omitempty"`
}

type GetOSDsResponse struct {
	OSDs []*common.OSDConfiguration `xml:"OSDs,omitempty" json:"OSDs,omitempty" yaml:"OSDs,omitempty"`
}

type GetProfiles struct {
	Token *common.ReferenceToken `xml:"Token,omitempty" json:"Token,omitempty" yaml:"Token,omitempty"`
	Type  []string               `xml:"Type,omitempty" json:"Type,omitempty" yaml:"Type,omitempty"`
}

type GetProfilesResponse struct {
	Profiles []MediaProfile `xml:"Profiles,omitempty" json:"Profiles,omitempty" yaml:"Profiles,omitempty"`
}

type GetServiceCapabilities struct {
}

type GetServiceCapabilitiesResponse struct {
	Capabilities Capabilities2 `xml:"Capabilities,omitempty" json:"Capabilities,omitempty" yaml:"Capabilities,omitempty"`
}

type GetSnapshotUri struct {
	ProfileToken *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
}

type GetSnapshotUriResponse struct {
	Uri string `xml:"Uri,omitempty" json:"Uri,omitempty" yaml:"Uri,omitempty"`
}

type GetStreamUri struct {
	Protocol     string                 `xml:"Protocol,omitempty" json:"Protocol,omitempty" yaml:"Protocol,omitempty"`
	ProfileToken *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
}

type GetStreamUriResponse struct {
	Uri string `xml:"Uri,omitempty" json:"Uri,omitempty" yaml:"Uri,omitempty"`
}

type GetVideoEncoderConfigurationOptionsResponse struct {
	Options []*common.VideoEncoder2ConfigurationOptions `xml:"Options,omitempty" json:"Options,omitempty" yaml:"Options,omitempty"`
}

type GetVideoEncoderConfigurationsResponse struct {
	Configurations []*common.VideoEncoder2Configuration `xml:"Configurations,omitempty" json:"Configurations,omitempty" yaml:"Configurations,omitempty"`
}

type GetVideoEncoderInstances struct {
	ConfigurationToken *common.ReferenceToken `xml:"ConfigurationToken,omitempty" json:"ConfigurationToken,omitempty" yaml:"ConfigurationToken,omitempty"`
}

type GetVideoEncoderInstancesResponse struct {
	Info EncoderInstanceInfo `xml:"Info,omitempty" json:"Info,omitempty" yaml:"Info,omitempty"`
}

type GetVideoSourceConfigurationOptionsResponse struct {
	Options *common.VideoSourceConfigurationOptions `xml:"Options,omitempty" json:"Options,omitempty" yaml:"Options,omitempty"`
}

type GetVideoSourceConfigurationsResponse struct {
	Configurations []*common.VideoSourceConfiguration `xml:"Configurations,omitempty" json:"Configurations,omitempty" yaml:"Configurations,omitempty"`
}

type GetVideoSourceModes struct {
	VideoSourceToken *common.ReferenceToken `xml:"VideoSourceToken,omitempty" json:"VideoSourceToken,omitempty" yaml:"VideoSourceToken,omitempty"`
}

type GetVideoSourceModesResponse struct {
	VideoSourceModes []VideoSourceMode `xml:"VideoSourceModes,omitempty" json:"VideoSourceModes,omitempty" yaml:"VideoSourceModes,omitempty"`
}

type Mask struct {
	ConfigurationToken *common.ReferenceToken `xml:"ConfigurationToken,omitempty" json:"ConfigurationToken,omitempty" yaml:"ConfigurationToken,omitempty"`
	Polygon            *common.Polygon        `xml:"Polygon,omitempty" json:"Polygon,omitempty" yaml:"Polygon,omitempty"`
	Type               string                 `xml:"Type,omitempty" json:"Type,omitempty" yaml:"Type,omitempty"`
	Color              *common.Color          `xml:"Color,omitempty" json:"Color,omitempty" yaml:"Color,omitempty"`
	Enabled            bool                   `xml:"Enabled,omitempty" json:"Enabled,omitempty" yaml:"Enabled,omitempty"`
	Token              common.ReferenceToken  `xml:"token,attr,omitempty" json:"token,attr,omitempty" yaml:"token,attr,omitempty"`
}

type MaskOptions struct {
	MaxMasks        int                  `xml:"MaxMasks,omitempty" json:"MaxMasks,omitempty" yaml:"MaxMasks,omitempty"`
	MaxPoints       int                  `xml:"MaxPoints,omitempty" json:"MaxPoints,omitempty" yaml:"MaxPoints,omitempty"`
	Types           []string             `xml:"Types,omitempty" json:"Types,omitempty" yaml:"Types,omitempty"`
	Color           *common.ColorOptions `xml:"Color,omitempty" json:"Color,omitempty" yaml:"Color,omitempty"`
	RectangleOnly   bool                 `xml:"RectangleOnly,attr,omitempty" json:"RectangleOnly,attr,omitempty" yaml:"RectangleOnly,attr,omitempty"`
	SingleColorOnly bool                 `xml:"SingleColorOnly,attr,omitempty" json:"SingleColorOnly,attr,omitempty" yaml:"SingleColorOnly,attr,omitempty"`
}

type MediaProfile struct {
	Name           *common.Name          `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Configurations ConfigurationSet      `xml:"Configurations,omitempty" json:"Configurations,omitempty" yaml:"Configurations,omitempty"`
	Token          common.ReferenceToken `xml:"token,attr,omitempty" json:"token,attr,omitempty" yaml:"token,attr,omitempty"`
	Fixed          bool                  `xml:"fixed,attr,omitempty" json:"fixed,attr,omitempty" yaml:"fixed,attr,omitempty"`
}

type ProfileCapabilities []interface{}

type ReceiverConfiguration struct {
	Token         common.ReferenceToken `xml:"token,attr,omitempty" json:"token,attr,omitempty" yaml:"token,attr,omitempty"`
	TypeAttrXSI   string                `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string                `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *ReceiverConfiguration) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:ReceiverConfiguration"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://www.onvif.org/ver20/media/wsdl"
	}
}

type RemoveConfiguration struct {
	ProfileToken  *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
	Configuration []ConfigurationRef     `xml:"Configuration,omitempty" json:"Configuration,omitempty" yaml:"Configuration,omitempty"`
}

type RemoveConfigurationResponse struct {
}

type SetAudioDecoderConfiguration struct {
	Configuration *common.AudioDecoderConfiguration `xml:"Configuration,omitempty" json:"Configuration,omitempty" yaml:"Configuration,omitempty"`
}

type SetAudioEncoderConfiguration struct {
	Configuration *common.AudioEncoder2Configuration `xml:"Configuration,omitempty" json:"Configuration,omitempty" yaml:"Configuration,omitempty"`
}

type SetAudioOutputConfiguration struct {
	Configuration *common.AudioOutputConfiguration `xml:"Configuration,omitempty" json:"Configuration,omitempty" yaml:"Configuration,omitempty"`
}

type SetAudioSourceConfiguration struct {
	Configuration *common.AudioSourceConfiguration `xml:"Configuration,omitempty" json:"Configuration,omitempty" yaml:"Configuration,omitempty"`
}

type SetConfigurationResponse struct {
}

type SetMask struct {
	Mask Mask `xml:"Mask,omitempty" json:"Mask,omitempty" yaml:"Mask,omitempty"`
}

type SetMetadataConfiguration struct {
	Configuration *common.MetadataConfiguration `xml:"Configuration,omitempty" json:"Configuration,omitempty" yaml:"Configuration,omitempty"`
}

type SetOSD struct {
	OSD *common.OSDConfiguration `xml:"OSD,omitempty" json:"OSD,omitempty" yaml:"OSD,omitempty"`
}

type SetSynchronizationPoint struct {
	ProfileToken *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
}

type SetSynchronizationPointResponse struct {
}

type SetVideoEncoderConfiguration struct {
	Configuration *common.VideoEncoder2Configuration `xml:"Configuration,omitempty" json:"Configuration,omitempty" yaml:"Configuration,omitempty"`
}

type SetVideoSourceConfiguration struct {
	Configuration *common.VideoSourceConfiguration `xml:"Configuration,omitempty" json:"Configuration,omitempty" yaml:"Configuration,omitempty"`
}

type SetVideoSourceMode struct {
	VideoSourceToken     *common.ReferenceToken `xml:"VideoSourceToken,omitempty" json:"VideoSourceToken,omitempty" yaml:"VideoSourceToken,omitempty"`
	VideoSourceModeToken *common.ReferenceToken `xml:"VideoSourceModeToken,omitempty" json:"VideoSourceModeToken,omitempty" yaml:"VideoSourceModeToken,omitempty"`
}

type SetVideoSourceModeResponse struct {
	Reboot bool `xml:"Reboot,omitempty" json:"Reboot,omitempty" yaml:"Reboot,omitempty"`
}

type StartStopMulticastStreaming struct {
	ProfileToken *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
}

type StreamingCapabilities []interface{}

type VideoSourceMode struct {
	MaxFramerate  float64                 `xml:"MaxFramerate,omitempty" json:"MaxFramerate,omitempty" yaml:"MaxFramerate,omitempty"`
	MaxResolution *common.VideoResolution `xml:"MaxResolution,omitempty" json:"MaxResolution,omitempty" yaml:"MaxResolution,omitempty"`
	Encodings     *common.StringList      `xml:"Encodings,omitempty" json:"Encodings,omitempty" yaml:"Encodings,omitempty"`
	Reboot        bool                    `xml:"Reboot,omitempty" json:"Reboot,omitempty" yaml:"Reboot,omitempty"`
	Description   *common.Description     `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	Token         common.ReferenceToken   `xml:"token,attr,omitempty" json:"token,attr,omitempty" yaml:"token,attr,omitempty"`
	Enabled       bool                    `xml:"Enabled,attr,omitempty" json:"Enabled,attr,omitempty" yaml:"Enabled,attr,omitempty"`
}

// media2 implements the Media2 interface.
type media2 struct {
	cli      common.Client
	Endpoint string
}

func (p *media2) OptAddConfiguration(args AddConfiguration) (*AddConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName          string `xml:"tr2:AddConfiguration"`
		AddConfiguration AddConfiguration
	}{
		AddConfiguration: args,
	}

	resp := AddConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media2) OptCreateMask(args CreateMask) (*CreateMaskResponse, *common.Fault) {
	req := struct {
		XMLName    string `xml:"tr2:CreateMask"`
		CreateMask CreateMask
	}{
		CreateMask: args,
	}

	resp := CreateMaskResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media2) OptCreateOSD(args CreateOSD) (*CreateOSDResponse, *common.Fault) {
	req := struct {
		XMLName   string `xml:"tr2:CreateOSD"`
		CreateOSD CreateOSD
	}{
		CreateOSD: args,
	}

	resp := CreateOSDResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media2) OptCreateProfile(args CreateProfile) (*CreateProfileResponse, *common.Fault) {
	req := struct {
		XMLName       string `xml:"tr2:CreateProfile"`
		CreateProfile CreateProfile
	}{
		CreateProfile: args,
	}

	resp := CreateProfileResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media2) OptDeleteMask(args DeleteMask) (*SetConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName    string `xml:"tr2:DeleteMask"`
		DeleteMask DeleteMask
	}{
		DeleteMask: args,
	}

	resp := SetConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media2) OptDeleteOSD(args DeleteOSD) (*SetConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName   string `xml:"tr2:DeleteOSD"`
		DeleteOSD DeleteOSD
	}{
		DeleteOSD: args,
	}

	resp := SetConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media2) OptDeleteProfile(args DeleteProfile) (*DeleteProfileResponse, *common.Fault) {
	req := struct {
		XMLName       string `xml:"tr2:DeleteProfile"`
		DeleteProfile DeleteProfile
	}{
		DeleteProfile: args,
	}

	resp := DeleteProfileResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media2) OptGetAnalyticsConfigurations(args GetConfiguration) (*GetAnalyticsConfigurationsResponse, *common.Fault) {
	req := struct {
		XMLName          string `xml:"tr2:GetAnalyticsConfigurations"`
		GetConfiguration GetConfiguration
	}{
		GetConfiguration: args,
	}

	resp := GetAnalyticsConfigurationsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media2) OptGetAudioDecoderConfigurationOptions(args GetConfiguration) (*GetAudioDecoderConfigurationOptionsResponse, *common.Fault) {
	req := struct {
		XMLName          string `xml:"tr2:GetAudioDecoderConfigurationOptions"`
		GetConfiguration GetConfiguration
	}{
		GetConfiguration: args,
	}

	resp := GetAudioDecoderConfigurationOptionsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media2) OptGetAudioDecoderConfigurations(args GetConfiguration) (*GetAudioDecoderConfigurationsResponse, *common.Fault) {
	req := struct {
		XMLName          string `xml:"tr2:GetAudioDecoderConfigurations"`
		GetConfiguration GetConfiguration
	}{
		GetConfiguration: args,
	}

	resp := GetAudioDecoderConfigurationsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media2) OptGetAudioEncoderConfigurationOptions(args GetConfiguration) (*GetAudioEncoderConfigurationOptionsResponse, *common.Fault) {
	req := struct {
		XMLName          string `xml:"tr2:GetAudioEncoderConfigurationOptions"`
		GetConfiguration GetConfiguration
	}{
		GetConfiguration: args,
	}

	resp := GetAudioEncoderConfigurationOptionsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media2) OptGetAudioEncoderConfigurations(args GetConfiguration) (*GetAudioEncoderConfigurationsResponse, *common.Fault) {
	req := struct {
		XMLName          string `xml:"tr2:GetAudioEncoderConfigurations"`
		GetConfiguration GetConfiguration
	}{
		GetConfiguration: args,
	}

	resp := GetAudioEncoderConfigurationsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media2) OptGetAudioOutputConfigurationOptions(args GetConfiguration) (*GetAudioOutputConfigurationOptionsResponse, *common.Fault) {
	req := struct {
		XMLName          string `xml:"tr2:GetAudioOutputConfigurationOptions"`
		GetConfiguration GetConfiguration
	}{
		GetConfiguration: args,
	}

	resp := GetAudioOutputConfigurationOptionsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media2) OptGetAudioOutputConfigurations(args GetConfiguration) (*GetAudioOutputConfigurationsResponse, *common.Fault) {
	req := struct {
		XMLName          string `xml:"tr2:GetAudioOutputConfigurations"`
		GetConfiguration GetConfiguration
	}{
		GetConfiguration: args,
	}

	resp := GetAudioOutputConfigurationsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media2) OptGetAudioSourceConfigurationOptions(args GetConfiguration) (*GetAudioSourceConfigurationOptionsResponse, *common.Fault) {
	req := struct {
		XMLName          string `xml:"tr2:GetAudioSourceConfigurationOptions"`
		GetConfiguration GetConfiguration
	}{
		GetConfiguration: args,
	}

	resp := GetAudioSourceConfigurationOptionsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media2) OptGetAudioSourceConfigurations(args GetConfiguration) (*GetAudioSourceConfigurationsResponse, *common.Fault) {
	req := struct {
		XMLName          string `xml:"tr2:GetAudioSourceConfigurations"`
		GetConfiguration GetConfiguration
	}{
		GetConfiguration: args,
	}

	resp := GetAudioSourceConfigurationsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media2) OptGetMaskOptions(args GetMaskOptions) (*GetMaskOptionsResponse, *common.Fault) {
	req := struct {
		XMLName        string `xml:"tr2:GetMaskOptions"`
		GetMaskOptions GetMaskOptions
	}{
		GetMaskOptions: args,
	}

	resp := GetMaskOptionsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media2) OptGetMasks(args GetMasks) (*GetMasksResponse, *common.Fault) {
	req := struct {
		XMLName  string `xml:"tr2:GetMasks"`
		GetMasks GetMasks
	}{
		GetMasks: args,
	}

	resp := GetMasksResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media2) OptGetMetadataConfigurationOptions(args GetConfiguration) (*GetMetadataConfigurationOptionsResponse, *common.Fault) {
	req := struct {
		XMLName          string `xml:"tr2:GetMetadataConfigurationOptions"`
		GetConfiguration GetConfiguration
	}{
		GetConfiguration: args,
	}

	resp := GetMetadataConfigurationOptionsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media2) OptGetMetadataConfigurations(args GetConfiguration) (*GetMetadataConfigurationsResponse, *common.Fault) {
	req := struct {
		XMLName          string `xml:"tr2:GetMetadataConfigurations"`
		GetConfiguration GetConfiguration
	}{
		GetConfiguration: args,
	}

	resp := GetMetadataConfigurationsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media2) OptGetOSDOptions(args GetOSDOptions) (*GetOSDOptionsResponse, *common.Fault) {
	req := struct {
		XMLName       string `xml:"tr2:GetOSDOptions"`
		GetOSDOptions GetOSDOptions
	}{
		GetOSDOptions: args,
	}

	resp := GetOSDOptionsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media2) OptGetOSDs(args GetOSDs) (*GetOSDsResponse, *common.Fault) {
	req := struct {
		XMLName string `xml:"tr2:GetOSDs"`
		GetOSDs GetOSDs
	}{
		GetOSDs: args,
	}

	resp := GetOSDsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media2) OptGetProfiles(args GetProfiles) (*GetProfilesResponse, *common.Fault) {
	req := struct {
		XMLName     string `xml:"tr2:GetProfiles"`
		GetProfiles GetProfiles
	}{
		GetProfiles: args,
	}

	resp := GetProfilesResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media2) OptGetServiceCapabilities(args GetServiceCapabilities) (*GetServiceCapabilitiesResponse, *common.Fault) {
	req := struct {
		XMLName                string `xml:"tr2:GetServiceCapabilities"`
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

func (p *media2) OptGetSnapshotUri(args GetSnapshotUri) (*GetSnapshotUriResponse, *common.Fault) {
	req := struct {
		XMLName        string `xml:"tr2:GetSnapshotUri"`
		GetSnapshotUri GetSnapshotUri
	}{
		GetSnapshotUri: args,
	}

	resp := GetSnapshotUriResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media2) OptGetStreamUri(args GetStreamUri) (*GetStreamUriResponse, *common.Fault) {
	req := struct {
		XMLName      string `xml:"tr2:GetStreamUri"`
		GetStreamUri GetStreamUri
	}{
		GetStreamUri: args,
	}

	resp := GetStreamUriResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media2) OptGetVideoEncoderConfigurationOptions(args GetConfiguration) (*GetVideoEncoderConfigurationOptionsResponse, *common.Fault) {
	req := struct {
		XMLName          string `xml:"tr2:GetVideoEncoderConfigurationOptions"`
		GetConfiguration GetConfiguration
	}{
		GetConfiguration: args,
	}

	resp := GetVideoEncoderConfigurationOptionsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media2) OptGetVideoEncoderConfigurations(args GetConfiguration) (*GetVideoEncoderConfigurationsResponse, *common.Fault) {
	req := struct {
		XMLName          string `xml:"tr2:GetVideoEncoderConfigurations"`
		GetConfiguration GetConfiguration
	}{
		GetConfiguration: args,
	}

	resp := GetVideoEncoderConfigurationsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media2) OptGetVideoEncoderInstances(args GetVideoEncoderInstances) (*GetVideoEncoderInstancesResponse, *common.Fault) {
	req := struct {
		XMLName                  string `xml:"tr2:GetVideoEncoderInstances"`
		GetVideoEncoderInstances GetVideoEncoderInstances
	}{
		GetVideoEncoderInstances: args,
	}

	resp := GetVideoEncoderInstancesResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media2) OptGetVideoSourceConfigurationOptions(args GetConfiguration) (*GetVideoSourceConfigurationOptionsResponse, *common.Fault) {
	req := struct {
		XMLName          string `xml:"tr2:GetVideoSourceConfigurationOptions"`
		GetConfiguration GetConfiguration
	}{
		GetConfiguration: args,
	}

	resp := GetVideoSourceConfigurationOptionsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media2) OptGetVideoSourceConfigurations(args GetConfiguration) (*GetVideoSourceConfigurationsResponse, *common.Fault) {
	req := struct {
		XMLName          string `xml:"tr2:GetVideoSourceConfigurations"`
		GetConfiguration GetConfiguration
	}{
		GetConfiguration: args,
	}

	resp := GetVideoSourceConfigurationsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media2) OptGetVideoSourceModes(args GetVideoSourceModes) (*GetVideoSourceModesResponse, *common.Fault) {
	req := struct {
		XMLName             string `xml:"tr2:GetVideoSourceModes"`
		GetVideoSourceModes GetVideoSourceModes
	}{
		GetVideoSourceModes: args,
	}

	resp := GetVideoSourceModesResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media2) OptRemoveConfiguration(args RemoveConfiguration) (*RemoveConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName             string `xml:"tr2:RemoveConfiguration"`
		RemoveConfiguration RemoveConfiguration
	}{
		RemoveConfiguration: args,
	}

	resp := RemoveConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media2) OptSetAudioDecoderConfiguration(args SetAudioDecoderConfiguration) (*SetConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                      string `xml:"tr2:SetAudioDecoderConfiguration"`
		SetAudioDecoderConfiguration SetAudioDecoderConfiguration
	}{
		SetAudioDecoderConfiguration: args,
	}

	resp := SetConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media2) OptSetAudioEncoderConfiguration(args SetAudioEncoderConfiguration) (*SetConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                      string `xml:"tr2:SetAudioEncoderConfiguration"`
		SetAudioEncoderConfiguration SetAudioEncoderConfiguration
	}{
		SetAudioEncoderConfiguration: args,
	}

	resp := SetConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media2) OptSetAudioOutputConfiguration(args SetAudioOutputConfiguration) (*SetConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                     string `xml:"tr2:SetAudioOutputConfiguration"`
		SetAudioOutputConfiguration SetAudioOutputConfiguration
	}{
		SetAudioOutputConfiguration: args,
	}

	resp := SetConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media2) OptSetAudioSourceConfiguration(args SetAudioSourceConfiguration) (*SetConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                     string `xml:"tr2:SetAudioSourceConfiguration"`
		SetAudioSourceConfiguration SetAudioSourceConfiguration
	}{
		SetAudioSourceConfiguration: args,
	}

	resp := SetConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media2) OptSetMask(args SetMask) (*SetConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName string `xml:"tr2:SetMask"`
		SetMask SetMask
	}{
		SetMask: args,
	}

	resp := SetConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media2) OptSetMetadataConfiguration(args SetMetadataConfiguration) (*SetConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                  string `xml:"tr2:SetMetadataConfiguration"`
		SetMetadataConfiguration SetMetadataConfiguration
	}{
		SetMetadataConfiguration: args,
	}

	resp := SetConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media2) OptSetOSD(args SetOSD) (*SetConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName string `xml:"tr2:SetOSD"`
		SetOSD  SetOSD
	}{
		SetOSD: args,
	}

	resp := SetConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media2) OptSetSynchronizationPoint(args SetSynchronizationPoint) (*SetSynchronizationPointResponse, *common.Fault) {
	req := struct {
		XMLName                 string `xml:"tr2:SetSynchronizationPoint"`
		SetSynchronizationPoint SetSynchronizationPoint
	}{
		SetSynchronizationPoint: args,
	}

	resp := SetSynchronizationPointResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media2) OptSetVideoEncoderConfiguration(args SetVideoEncoderConfiguration) (*SetConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                      string `xml:"tr2:SetVideoEncoderConfiguration"`
		SetVideoEncoderConfiguration SetVideoEncoderConfiguration
	}{
		SetVideoEncoderConfiguration: args,
	}

	resp := SetConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media2) OptSetVideoSourceConfiguration(args SetVideoSourceConfiguration) (*SetConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                     string `xml:"tr2:SetVideoSourceConfiguration"`
		SetVideoSourceConfiguration SetVideoSourceConfiguration
	}{
		SetVideoSourceConfiguration: args,
	}

	resp := SetConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media2) OptSetVideoSourceMode(args SetVideoSourceMode) (*SetVideoSourceModeResponse, *common.Fault) {
	req := struct {
		XMLName            string `xml:"tr2:SetVideoSourceMode"`
		SetVideoSourceMode SetVideoSourceMode
	}{
		SetVideoSourceMode: args,
	}

	resp := SetVideoSourceModeResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media2) OptStartMulticastStreaming(args StartStopMulticastStreaming) (*SetConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                     string `xml:"tr2:StartMulticastStreaming"`
		StartStopMulticastStreaming StartStopMulticastStreaming
	}{
		StartStopMulticastStreaming: args,
	}

	resp := SetConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media2) OptStopMulticastStreaming(args StartStopMulticastStreaming) (*SetConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                     string `xml:"tr2:StopMulticastStreaming"`
		StartStopMulticastStreaming StartStopMulticastStreaming
	}{
		StartStopMulticastStreaming: args,
	}

	resp := SetConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}
