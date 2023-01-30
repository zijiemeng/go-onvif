package media

import (
	"code.byted.org/videoarch/go-onvif/onvif/common"
)

var Namespace = "http://www.onvif.org/ver10/media/wsdl"

// NewMedia creates an initializes a Media.
func NewMedia(endpoint string, cli common.Client) Media {
	return &media{cli: cli, Endpoint: endpoint}
}

// Media was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type Media interface {
	OptAddAudioDecoderConfiguration(AddAudioDecoderConfiguration AddAudioDecoderConfiguration) (*AddAudioDecoderConfigurationResponse, *common.Fault)

	OptAddAudioEncoderConfiguration(AddAudioEncoderConfiguration AddAudioEncoderConfiguration) (*AddAudioEncoderConfigurationResponse, *common.Fault)

	OptAddAudioOutputConfiguration(AddAudioOutputConfiguration AddAudioOutputConfiguration) (*AddAudioOutputConfigurationResponse, *common.Fault)

	OptAddAudioSourceConfiguration(AddAudioSourceConfiguration AddAudioSourceConfiguration) (*AddAudioSourceConfigurationResponse, *common.Fault)

	OptAddMetadataConfiguration(AddMetadataConfiguration AddMetadataConfiguration) (*AddMetadataConfigurationResponse, *common.Fault)

	OptAddPTZConfiguration(AddPTZConfiguration AddPTZConfiguration) (*AddPTZConfigurationResponse, *common.Fault)

	OptAddVideoAnalyticsConfiguration(AddVideoAnalyticsConfiguration AddVideoAnalyticsConfiguration) (*AddVideoAnalyticsConfigurationResponse, *common.Fault)

	OptAddVideoEncoderConfiguration(AddVideoEncoderConfiguration AddVideoEncoderConfiguration) (*AddVideoEncoderConfigurationResponse, *common.Fault)

	OptAddVideoSourceConfiguration(AddVideoSourceConfiguration AddVideoSourceConfiguration) (*AddVideoSourceConfigurationResponse, *common.Fault)

	OptCreateOSD(CreateOSD CreateOSD) (*CreateOSDResponse, *common.Fault)

	OptCreateProfile(CreateProfile CreateProfile) (*CreateProfileResponse, *common.Fault)

	OptDeleteOSD(DeleteOSD DeleteOSD) (*DeleteOSDResponse, *common.Fault)

	OptDeleteProfile(DeleteProfile DeleteProfile) (*DeleteProfileResponse, *common.Fault)

	OptGetAudioDecoderConfiguration(GetAudioDecoderConfiguration GetAudioDecoderConfiguration) (*GetAudioDecoderConfigurationResponse, *common.Fault)

	OptGetAudioDecoderConfigurationOptions(GetAudioDecoderConfigurationOptions GetAudioDecoderConfigurationOptions) (*GetAudioDecoderConfigurationOptionsResponse, *common.Fault)

	OptGetAudioDecoderConfigurations(GetAudioDecoderConfigurations GetAudioDecoderConfigurations) (*GetAudioDecoderConfigurationsResponse, *common.Fault)

	OptGetAudioEncoderConfiguration(GetAudioEncoderConfiguration GetAudioEncoderConfiguration) (*GetAudioEncoderConfigurationResponse, *common.Fault)

	OptGetAudioEncoderConfigurationOptions(GetAudioEncoderConfigurationOptions GetAudioEncoderConfigurationOptions) (*GetAudioEncoderConfigurationOptionsResponse, *common.Fault)

	OptGetAudioEncoderConfigurations(GetAudioEncoderConfigurations GetAudioEncoderConfigurations) (*GetAudioEncoderConfigurationsResponse, *common.Fault)

	OptGetAudioOutputConfiguration(GetAudioOutputConfiguration GetAudioOutputConfiguration) (*GetAudioOutputConfigurationResponse, *common.Fault)

	OptGetAudioOutputConfigurationOptions(GetAudioOutputConfigurationOptions GetAudioOutputConfigurationOptions) (*GetAudioOutputConfigurationOptionsResponse, *common.Fault)

	OptGetAudioOutputConfigurations(GetAudioOutputConfigurations GetAudioOutputConfigurations) (*GetAudioOutputConfigurationsResponse, *common.Fault)

	OptGetAudioOutputs(GetAudioOutputs GetAudioOutputs) (*GetAudioOutputsResponse, *common.Fault)

	OptGetAudioSourceConfiguration(GetAudioSourceConfiguration GetAudioSourceConfiguration) (*GetAudioSourceConfigurationResponse, *common.Fault)

	OptGetAudioSourceConfigurationOptions(GetAudioSourceConfigurationOptions GetAudioSourceConfigurationOptions) (*GetAudioSourceConfigurationOptionsResponse, *common.Fault)

	OptGetAudioSourceConfigurations(GetAudioSourceConfigurations GetAudioSourceConfigurations) (*GetAudioSourceConfigurationsResponse, *common.Fault)

	OptGetAudioSources(GetAudioSources GetAudioSources) (*GetAudioSourcesResponse, *common.Fault)

	OptGetCompatibleAudioDecoderConfigurations(GetCompatibleAudioDecoderConfigurations GetCompatibleAudioDecoderConfigurations) (*GetCompatibleAudioDecoderConfigurationsResponse, *common.Fault)

	OptGetCompatibleAudioEncoderConfigurations(GetCompatibleAudioEncoderConfigurations GetCompatibleAudioEncoderConfigurations) (*GetCompatibleAudioEncoderConfigurationsResponse, *common.Fault)

	OptGetCompatibleAudioOutputConfigurations(GetCompatibleAudioOutputConfigurations GetCompatibleAudioOutputConfigurations) (*GetCompatibleAudioOutputConfigurationsResponse, *common.Fault)

	OptGetCompatibleAudioSourceConfigurations(GetCompatibleAudioSourceConfigurations GetCompatibleAudioSourceConfigurations) (*GetCompatibleAudioSourceConfigurationsResponse, *common.Fault)

	OptGetCompatibleMetadataConfigurations(GetCompatibleMetadataConfigurations GetCompatibleMetadataConfigurations) (*GetCompatibleMetadataConfigurationsResponse, *common.Fault)

	OptGetCompatibleVideoAnalyticsConfigurations(GetCompatibleVideoAnalyticsConfigurations GetCompatibleVideoAnalyticsConfigurations) (*GetCompatibleVideoAnalyticsConfigurationsResponse, *common.Fault)

	OptGetCompatibleVideoEncoderConfigurations(GetCompatibleVideoEncoderConfigurations GetCompatibleVideoEncoderConfigurations) (*GetCompatibleVideoEncoderConfigurationsResponse, *common.Fault)

	OptGetCompatibleVideoSourceConfigurations(GetCompatibleVideoSourceConfigurations GetCompatibleVideoSourceConfigurations) (*GetCompatibleVideoSourceConfigurationsResponse, *common.Fault)

	OptGetGuaranteedNumberOfVideoEncoderInstances(GetGuaranteedNumberOfVideoEncoderInstances GetGuaranteedNumberOfVideoEncoderInstances) (*GetGuaranteedNumberOfVideoEncoderInstancesResponse, *common.Fault)

	OptGetMetadataConfiguration(GetMetadataConfiguration GetMetadataConfiguration) (*GetMetadataConfigurationResponse, *common.Fault)

	OptGetMetadataConfigurationOptions(GetMetadataConfigurationOptions GetMetadataConfigurationOptions) (*GetMetadataConfigurationOptionsResponse, *common.Fault)

	OptGetMetadataConfigurations(GetMetadataConfigurations GetMetadataConfigurations) (*GetMetadataConfigurationsResponse, *common.Fault)

	OptGetOSD(GetOSD GetOSD) (*GetOSDResponse, *common.Fault)

	OptGetOSDOptions(GetOSDOptions GetOSDOptions) (*GetOSDOptionsResponse, *common.Fault)

	OptGetOSDs(GetOSDs GetOSDs) (*GetOSDsResponse, *common.Fault)

	OptGetProfile(GetProfile GetProfile) (*GetProfileResponse, *common.Fault)

	OptGetProfiles(GetProfiles GetProfiles) (*GetProfilesResponse, *common.Fault)

	OptGetServiceCapabilities(GetServiceCapabilities GetServiceCapabilities) (*GetServiceCapabilitiesResponse, *common.Fault)

	OptGetSnapshotUri(GetSnapshotUri GetSnapshotUri) (*GetSnapshotUriResponse, *common.Fault)

	OptGetStreamUri(GetStreamUri GetStreamUri) (*GetStreamUriResponse, *common.Fault)

	OptGetVideoAnalyticsConfiguration(GetVideoAnalyticsConfiguration GetVideoAnalyticsConfiguration) (*GetVideoAnalyticsConfigurationResponse, *common.Fault)

	OptGetVideoAnalyticsConfigurations(GetVideoAnalyticsConfigurations GetVideoAnalyticsConfigurations) (*GetVideoAnalyticsConfigurationsResponse, *common.Fault)

	OptGetVideoEncoderConfiguration(GetVideoEncoderConfiguration GetVideoEncoderConfiguration) (*GetVideoEncoderConfigurationResponse, *common.Fault)

	OptGetVideoEncoderConfigurationOptions(GetVideoEncoderConfigurationOptions GetVideoEncoderConfigurationOptions) (*GetVideoEncoderConfigurationOptionsResponse, *common.Fault)

	OptGetVideoEncoderConfigurations(GetVideoEncoderConfigurations GetVideoEncoderConfigurations) (*GetVideoEncoderConfigurationsResponse, *common.Fault)

	OptGetVideoSourceConfiguration(GetVideoSourceConfiguration GetVideoSourceConfiguration) (*GetVideoSourceConfigurationResponse, *common.Fault)

	OptGetVideoSourceConfigurationOptions(GetVideoSourceConfigurationOptions GetVideoSourceConfigurationOptions) (*GetVideoSourceConfigurationOptionsResponse, *common.Fault)

	OptGetVideoSourceConfigurations(GetVideoSourceConfigurations GetVideoSourceConfigurations) (*GetVideoSourceConfigurationsResponse, *common.Fault)

	OptGetVideoSourceModes(GetVideoSourceModes GetVideoSourceModes) (*GetVideoSourceModesResponse, *common.Fault)

	OptGetVideoSources(GetVideoSources GetVideoSources) (*GetVideoSourcesResponse, *common.Fault)

	OptRemoveAudioDecoderConfiguration(RemoveAudioDecoderConfiguration RemoveAudioDecoderConfiguration) (*RemoveAudioDecoderConfigurationResponse, *common.Fault)

	OptRemoveAudioEncoderConfiguration(RemoveAudioEncoderConfiguration RemoveAudioEncoderConfiguration) (*RemoveAudioEncoderConfigurationResponse, *common.Fault)

	OptRemoveAudioOutputConfiguration(RemoveAudioOutputConfiguration RemoveAudioOutputConfiguration) (*RemoveAudioOutputConfigurationResponse, *common.Fault)

	OptRemoveAudioSourceConfiguration(RemoveAudioSourceConfiguration RemoveAudioSourceConfiguration) (*RemoveAudioSourceConfigurationResponse, *common.Fault)

	OptRemoveMetadataConfiguration(RemoveMetadataConfiguration RemoveMetadataConfiguration) (*RemoveMetadataConfigurationResponse, *common.Fault)

	OptRemovePTZConfiguration(RemovePTZConfiguration RemovePTZConfiguration) (*RemovePTZConfigurationResponse, *common.Fault)

	OptRemoveVideoAnalyticsConfiguration(RemoveVideoAnalyticsConfiguration RemoveVideoAnalyticsConfiguration) (*RemoveVideoAnalyticsConfigurationResponse, *common.Fault)

	OptRemoveVideoEncoderConfiguration(RemoveVideoEncoderConfiguration RemoveVideoEncoderConfiguration) (*RemoveVideoEncoderConfigurationResponse, *common.Fault)

	OptRemoveVideoSourceConfiguration(RemoveVideoSourceConfiguration RemoveVideoSourceConfiguration) (*RemoveVideoSourceConfigurationResponse, *common.Fault)

	OptSetAudioDecoderConfiguration(SetAudioDecoderConfiguration SetAudioDecoderConfiguration) (*SetAudioDecoderConfigurationResponse, *common.Fault)

	OptSetAudioEncoderConfiguration(SetAudioEncoderConfiguration SetAudioEncoderConfiguration) (*SetAudioEncoderConfigurationResponse, *common.Fault)

	OptSetAudioOutputConfiguration(SetAudioOutputConfiguration SetAudioOutputConfiguration) (*SetAudioOutputConfigurationResponse, *common.Fault)

	OptSetAudioSourceConfiguration(SetAudioSourceConfiguration SetAudioSourceConfiguration) (*SetAudioSourceConfigurationResponse, *common.Fault)

	OptSetMetadataConfiguration(SetMetadataConfiguration SetMetadataConfiguration) (*SetMetadataConfigurationResponse, *common.Fault)

	OptSetOSD(SetOSD SetOSD) (*SetOSDResponse, *common.Fault)

	OptSetSynchronizationPoint(SetSynchronizationPoint SetSynchronizationPoint) (*SetSynchronizationPointResponse, *common.Fault)

	OptSetVideoAnalyticsConfiguration(SetVideoAnalyticsConfiguration SetVideoAnalyticsConfiguration) (*SetVideoAnalyticsConfigurationResponse, *common.Fault)

	OptSetVideoEncoderConfiguration(SetVideoEncoderConfiguration SetVideoEncoderConfiguration) (*SetVideoEncoderConfigurationResponse, *common.Fault)

	OptSetVideoSourceConfiguration(SetVideoSourceConfiguration SetVideoSourceConfiguration) (*SetVideoSourceConfigurationResponse, *common.Fault)

	OptSetVideoSourceMode(SetVideoSourceMode SetVideoSourceMode) (*SetVideoSourceModeResponse, *common.Fault)

	OptStartMulticastStreaming(StartMulticastStreaming StartMulticastStreaming) (*StartMulticastStreamingResponse, *common.Fault)

	OptStopMulticastStreaming(StopMulticastStreaming StopMulticastStreaming) (*StopMulticastStreamingResponse, *common.Fault)
}
type AddAudioDecoderConfiguration struct {
	ProfileToken       *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
	ConfigurationToken *common.ReferenceToken `xml:"ConfigurationToken,omitempty" json:"ConfigurationToken,omitempty" yaml:"ConfigurationToken,omitempty"`
}

type AddAudioDecoderConfigurationResponse struct {
}

type AddAudioEncoderConfiguration struct {
	ProfileToken       *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
	ConfigurationToken *common.ReferenceToken `xml:"ConfigurationToken,omitempty" json:"ConfigurationToken,omitempty" yaml:"ConfigurationToken,omitempty"`
}

type AddAudioEncoderConfigurationResponse struct {
}

type AddAudioOutputConfiguration struct {
	ProfileToken       *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
	ConfigurationToken *common.ReferenceToken `xml:"ConfigurationToken,omitempty" json:"ConfigurationToken,omitempty" yaml:"ConfigurationToken,omitempty"`
}

type AddAudioOutputConfigurationResponse struct {
}

type AddAudioSourceConfiguration struct {
	ProfileToken       *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
	ConfigurationToken *common.ReferenceToken `xml:"ConfigurationToken,omitempty" json:"ConfigurationToken,omitempty" yaml:"ConfigurationToken,omitempty"`
}

type AddAudioSourceConfigurationResponse struct {
}

type AddMetadataConfiguration struct {
	ProfileToken       *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
	ConfigurationToken *common.ReferenceToken `xml:"ConfigurationToken,omitempty" json:"ConfigurationToken,omitempty" yaml:"ConfigurationToken,omitempty"`
}

type AddMetadataConfigurationResponse struct {
}

type AddPTZConfiguration struct {
	ProfileToken       *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
	ConfigurationToken *common.ReferenceToken `xml:"ConfigurationToken,omitempty" json:"ConfigurationToken,omitempty" yaml:"ConfigurationToken,omitempty"`
}

type AddPTZConfigurationResponse struct {
}

type AddVideoAnalyticsConfiguration struct {
	ProfileToken       *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
	ConfigurationToken *common.ReferenceToken `xml:"ConfigurationToken,omitempty" json:"ConfigurationToken,omitempty" yaml:"ConfigurationToken,omitempty"`
}

type AddVideoAnalyticsConfigurationResponse struct {
}

type AddVideoEncoderConfiguration struct {
	ProfileToken       *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
	ConfigurationToken *common.ReferenceToken `xml:"ConfigurationToken,omitempty" json:"ConfigurationToken,omitempty" yaml:"ConfigurationToken,omitempty"`
}

type AddVideoEncoderConfigurationResponse struct {
}

type AddVideoSourceConfiguration struct {
	ProfileToken       *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
	ConfigurationToken *common.ReferenceToken `xml:"ConfigurationToken,omitempty" json:"ConfigurationToken,omitempty" yaml:"ConfigurationToken,omitempty"`
}

type AddVideoSourceConfigurationResponse struct {
}

type Capabilities struct {
	ProfileCapabilities   ProfileCapabilities   `xml:"ProfileCapabilities,omitempty" json:"ProfileCapabilities,omitempty" yaml:"ProfileCapabilities,omitempty"`
	StreamingCapabilities StreamingCapabilities `xml:"StreamingCapabilities,omitempty" json:"StreamingCapabilities,omitempty" yaml:"StreamingCapabilities,omitempty"`
	SnapshotUri           bool                  `xml:"SnapshotUri,attr,omitempty" json:"SnapshotUri,attr,omitempty" yaml:"SnapshotUri,attr,omitempty"`
	Rotation              bool                  `xml:"Rotation,attr,omitempty" json:"Rotation,attr,omitempty" yaml:"Rotation,attr,omitempty"`
	VideoSourceMode       bool                  `xml:"VideoSourceMode,attr,omitempty" json:"VideoSourceMode,attr,omitempty" yaml:"VideoSourceMode,attr,omitempty"`
	OSD                   bool                  `xml:"OSD,attr,omitempty" json:"OSD,attr,omitempty" yaml:"OSD,attr,omitempty"`
	TemporaryOSDText      bool                  `xml:"TemporaryOSDText,attr,omitempty" json:"TemporaryOSDText,attr,omitempty" yaml:"TemporaryOSDText,attr,omitempty"`
	EXICompression        bool                  `xml:"EXICompression,attr,omitempty" json:"EXICompression,attr,omitempty" yaml:"EXICompression,attr,omitempty"`
}

type CreateOSD struct {
	OSD *common.OSDConfiguration `xml:"OSD,omitempty" json:"OSD,omitempty" yaml:"OSD,omitempty"`
}

type CreateOSDResponse struct {
	OSDToken *common.ReferenceToken `xml:"OSDToken,omitempty" json:"OSDToken,omitempty" yaml:"OSDToken,omitempty"`
}

type CreateProfile struct {
	Name  *common.Name           `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Token *common.ReferenceToken `xml:"Token,omitempty" json:"Token,omitempty" yaml:"Token,omitempty"`
}

type CreateProfileResponse struct {
	Profile *common.Profile `xml:"Profile,omitempty" json:"Profile,omitempty" yaml:"Profile,omitempty"`
}

type DeleteOSD struct {
	OSDToken *common.ReferenceToken `xml:"OSDToken,omitempty" json:"OSDToken,omitempty" yaml:"OSDToken,omitempty"`
}

type DeleteOSDResponse []interface{}

type DeleteProfile struct {
	ProfileToken *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
}

type DeleteProfileResponse struct {
}

type GetAudioDecoderConfiguration struct {
	ConfigurationToken *common.ReferenceToken `xml:"ConfigurationToken,omitempty" json:"ConfigurationToken,omitempty" yaml:"ConfigurationToken,omitempty"`
}

type GetAudioDecoderConfigurationOptions struct {
	ConfigurationToken *common.ReferenceToken `xml:"ConfigurationToken,omitempty" json:"ConfigurationToken,omitempty" yaml:"ConfigurationToken,omitempty"`
	ProfileToken       *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
}

type GetAudioDecoderConfigurationOptionsResponse struct {
	Options *common.AudioDecoderConfigurationOptions `xml:"Options,omitempty" json:"Options,omitempty" yaml:"Options,omitempty"`
}

type GetAudioDecoderConfigurationResponse struct {
	Configuration *common.AudioDecoderConfiguration `xml:"Configuration,omitempty" json:"Configuration,omitempty" yaml:"Configuration,omitempty"`
}

type GetAudioDecoderConfigurations struct {
}

type GetAudioDecoderConfigurationsResponse struct {
	Configurations []*common.AudioDecoderConfiguration `xml:"Configurations,omitempty" json:"Configurations,omitempty" yaml:"Configurations,omitempty"`
}

type GetAudioEncoderConfiguration struct {
	ConfigurationToken *common.ReferenceToken `xml:"ConfigurationToken,omitempty" json:"ConfigurationToken,omitempty" yaml:"ConfigurationToken,omitempty"`
}

type GetAudioEncoderConfigurationOptions struct {
	ConfigurationToken *common.ReferenceToken `xml:"ConfigurationToken,omitempty" json:"ConfigurationToken,omitempty" yaml:"ConfigurationToken,omitempty"`
	ProfileToken       *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
}

type GetAudioEncoderConfigurationOptionsResponse struct {
	Options *common.AudioEncoderConfigurationOptions `xml:"Options,omitempty" json:"Options,omitempty" yaml:"Options,omitempty"`
}

type GetAudioEncoderConfigurationResponse struct {
	Configuration *common.AudioEncoderConfiguration `xml:"Configuration,omitempty" json:"Configuration,omitempty" yaml:"Configuration,omitempty"`
}

type GetAudioEncoderConfigurations struct {
}

type GetAudioEncoderConfigurationsResponse struct {
	Configurations []*common.AudioEncoderConfiguration `xml:"Configurations,omitempty" json:"Configurations,omitempty" yaml:"Configurations,omitempty"`
}

type GetAudioOutputConfiguration struct {
	ConfigurationToken *common.ReferenceToken `xml:"ConfigurationToken,omitempty" json:"ConfigurationToken,omitempty" yaml:"ConfigurationToken,omitempty"`
}

type GetAudioOutputConfigurationOptions struct {
	ConfigurationToken *common.ReferenceToken `xml:"ConfigurationToken,omitempty" json:"ConfigurationToken,omitempty" yaml:"ConfigurationToken,omitempty"`
	ProfileToken       *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
}

type GetAudioOutputConfigurationOptionsResponse struct {
	Options *common.AudioOutputConfigurationOptions `xml:"Options,omitempty" json:"Options,omitempty" yaml:"Options,omitempty"`
}

type GetAudioOutputConfigurationResponse struct {
	Configuration *common.AudioOutputConfiguration `xml:"Configuration,omitempty" json:"Configuration,omitempty" yaml:"Configuration,omitempty"`
}

type GetAudioOutputConfigurations struct {
}

type GetAudioOutputConfigurationsResponse struct {
	Configurations []*common.AudioOutputConfiguration `xml:"Configurations,omitempty" json:"Configurations,omitempty" yaml:"Configurations,omitempty"`
}

type GetAudioOutputs struct {
}

type GetAudioOutputsResponse struct {
	AudioOutputs []*common.AudioOutput `xml:"AudioOutputs,omitempty" json:"AudioOutputs,omitempty" yaml:"AudioOutputs,omitempty"`
}

type GetAudioSourceConfiguration struct {
	ConfigurationToken *common.ReferenceToken `xml:"ConfigurationToken,omitempty" json:"ConfigurationToken,omitempty" yaml:"ConfigurationToken,omitempty"`
}

type GetAudioSourceConfigurationOptions struct {
	ConfigurationToken *common.ReferenceToken `xml:"ConfigurationToken,omitempty" json:"ConfigurationToken,omitempty" yaml:"ConfigurationToken,omitempty"`
	ProfileToken       *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
}

type GetAudioSourceConfigurationOptionsResponse struct {
	Options *common.AudioSourceConfigurationOptions `xml:"Options,omitempty" json:"Options,omitempty" yaml:"Options,omitempty"`
}

type GetAudioSourceConfigurationResponse struct {
	Configuration *common.AudioSourceConfiguration `xml:"Configuration,omitempty" json:"Configuration,omitempty" yaml:"Configuration,omitempty"`
}

type GetAudioSourceConfigurations struct {
}

type GetAudioSourceConfigurationsResponse struct {
	Configurations []*common.AudioSourceConfiguration `xml:"Configurations,omitempty" json:"Configurations,omitempty" yaml:"Configurations,omitempty"`
}

type GetAudioSources struct {
}

type GetAudioSourcesResponse struct {
	AudioSources []*common.AudioSource `xml:"AudioSources,omitempty" json:"AudioSources,omitempty" yaml:"AudioSources,omitempty"`
}

type GetCompatibleAudioDecoderConfigurations struct {
	ProfileToken *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
}

type GetCompatibleAudioDecoderConfigurationsResponse struct {
	Configurations []*common.AudioDecoderConfiguration `xml:"Configurations,omitempty" json:"Configurations,omitempty" yaml:"Configurations,omitempty"`
}

type GetCompatibleAudioEncoderConfigurations struct {
	ProfileToken *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
}

type GetCompatibleAudioEncoderConfigurationsResponse struct {
	Configurations []*common.AudioEncoderConfiguration `xml:"Configurations,omitempty" json:"Configurations,omitempty" yaml:"Configurations,omitempty"`
}

type GetCompatibleAudioOutputConfigurations struct {
	ProfileToken *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
}

type GetCompatibleAudioOutputConfigurationsResponse struct {
	Configurations []*common.AudioOutputConfiguration `xml:"Configurations,omitempty" json:"Configurations,omitempty" yaml:"Configurations,omitempty"`
}

type GetCompatibleAudioSourceConfigurations struct {
	ProfileToken *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
}

type GetCompatibleAudioSourceConfigurationsResponse struct {
	Configurations []*common.AudioSourceConfiguration `xml:"Configurations,omitempty" json:"Configurations,omitempty" yaml:"Configurations,omitempty"`
}

type GetCompatibleMetadataConfigurations struct {
	ProfileToken *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
}

type GetCompatibleMetadataConfigurationsResponse struct {
	Configurations []*common.MetadataConfiguration `xml:"Configurations,omitempty" json:"Configurations,omitempty" yaml:"Configurations,omitempty"`
}

type GetCompatibleVideoAnalyticsConfigurations struct {
	ProfileToken *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
}

type GetCompatibleVideoAnalyticsConfigurationsResponse struct {
	Configurations []*common.VideoAnalyticsConfiguration `xml:"Configurations,omitempty" json:"Configurations,omitempty" yaml:"Configurations,omitempty"`
}

type GetCompatibleVideoEncoderConfigurations struct {
	ProfileToken *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
}

type GetCompatibleVideoEncoderConfigurationsResponse struct {
	Configurations []*common.VideoEncoderConfiguration `xml:"Configurations,omitempty" json:"Configurations,omitempty" yaml:"Configurations,omitempty"`
}

type GetCompatibleVideoSourceConfigurations struct {
	ProfileToken *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
}

type GetCompatibleVideoSourceConfigurationsResponse struct {
	Configurations []*common.VideoSourceConfiguration `xml:"Configurations,omitempty" json:"Configurations,omitempty" yaml:"Configurations,omitempty"`
}

type GetGuaranteedNumberOfVideoEncoderInstances struct {
	ConfigurationToken *common.ReferenceToken `xml:"ConfigurationToken,omitempty" json:"ConfigurationToken,omitempty" yaml:"ConfigurationToken,omitempty"`
}

type GetGuaranteedNumberOfVideoEncoderInstancesResponse struct {
	TotalNumber int `xml:"TotalNumber,omitempty" json:"TotalNumber,omitempty" yaml:"TotalNumber,omitempty"`
	JPEG        int `xml:"JPEG,omitempty" json:"JPEG,omitempty" yaml:"JPEG,omitempty"`
	H264        int `xml:"H264,omitempty" json:"H264,omitempty" yaml:"H264,omitempty"`
	MPEG4       int `xml:"MPEG4,omitempty" json:"MPEG4,omitempty" yaml:"MPEG4,omitempty"`
}

type GetMetadataConfiguration struct {
	ConfigurationToken *common.ReferenceToken `xml:"ConfigurationToken,omitempty" json:"ConfigurationToken,omitempty" yaml:"ConfigurationToken,omitempty"`
}

type GetMetadataConfigurationOptions struct {
	ConfigurationToken *common.ReferenceToken `xml:"ConfigurationToken,omitempty" json:"ConfigurationToken,omitempty" yaml:"ConfigurationToken,omitempty"`
	ProfileToken       *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
}

type GetMetadataConfigurationOptionsResponse struct {
	Options *common.MetadataConfigurationOptions `xml:"Options,omitempty" json:"Options,omitempty" yaml:"Options,omitempty"`
}

type GetMetadataConfigurationResponse struct {
	Configuration *common.MetadataConfiguration `xml:"Configuration,omitempty" json:"Configuration,omitempty" yaml:"Configuration,omitempty"`
}

type GetMetadataConfigurations struct {
}

type GetMetadataConfigurationsResponse struct {
	Configurations []*common.MetadataConfiguration `xml:"Configurations,omitempty" json:"Configurations,omitempty" yaml:"Configurations,omitempty"`
}

type GetOSD struct {
	OSDToken *common.ReferenceToken `xml:"OSDToken,omitempty" json:"OSDToken,omitempty" yaml:"OSDToken,omitempty"`
}

type GetOSDOptions struct {
	ConfigurationToken *common.ReferenceToken `xml:"ConfigurationToken,omitempty" json:"ConfigurationToken,omitempty" yaml:"ConfigurationToken,omitempty"`
}

type GetOSDOptionsResponse struct {
	OSDOptions *common.OSDConfigurationOptions `xml:"OSDOptions,omitempty" json:"OSDOptions,omitempty" yaml:"OSDOptions,omitempty"`
}

type GetOSDResponse struct {
	OSD *common.OSDConfiguration `xml:"OSD,omitempty" json:"OSD,omitempty" yaml:"OSD,omitempty"`
}

type GetOSDs struct {
	ConfigurationToken *common.ReferenceToken `xml:"ConfigurationToken,omitempty" json:"ConfigurationToken,omitempty" yaml:"ConfigurationToken,omitempty"`
}

type GetOSDsResponse struct {
	OSDs []*common.OSDConfiguration `xml:"OSDs,omitempty" json:"OSDs,omitempty" yaml:"OSDs,omitempty"`
}

type GetProfile struct {
	ProfileToken *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
}

type GetProfileResponse struct {
	Profile *common.Profile `xml:"Profile,omitempty" json:"Profile,omitempty" yaml:"Profile,omitempty"`
}

type GetProfiles struct {
}

type GetProfilesResponse struct {
	Profiles []*common.Profile `xml:"Profiles,omitempty" json:"Profiles,omitempty" yaml:"Profiles,omitempty"`
}

type GetServiceCapabilities struct {
}

type GetServiceCapabilitiesResponse struct {
	Capabilities Capabilities `xml:"Capabilities,omitempty" json:"Capabilities,omitempty" yaml:"Capabilities,omitempty"`
}

type GetSnapshotUri struct {
	ProfileToken *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
}

type GetSnapshotUriResponse struct {
	MediaUri *common.MediaUri `xml:"MediaUri,omitempty" json:"MediaUri,omitempty" yaml:"MediaUri,omitempty"`
}

type GetStreamUri struct {
	StreamSetup  *common.StreamSetup    `xml:"StreamSetup,omitempty" json:"StreamSetup,omitempty" yaml:"StreamSetup,omitempty"`
	ProfileToken *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
}

type GetStreamUriResponse struct {
	MediaUri *common.MediaUri `xml:"MediaUri,omitempty" json:"MediaUri,omitempty" yaml:"MediaUri,omitempty"`
}

type GetVideoAnalyticsConfiguration struct {
	ConfigurationToken *common.ReferenceToken `xml:"ConfigurationToken,omitempty" json:"ConfigurationToken,omitempty" yaml:"ConfigurationToken,omitempty"`
}

type GetVideoAnalyticsConfigurationResponse struct {
	Configuration *common.VideoAnalyticsConfiguration `xml:"Configuration,omitempty" json:"Configuration,omitempty" yaml:"Configuration,omitempty"`
}

type GetVideoAnalyticsConfigurations struct {
}

type GetVideoAnalyticsConfigurationsResponse struct {
	Configurations []*common.VideoAnalyticsConfiguration `xml:"Configurations,omitempty" json:"Configurations,omitempty" yaml:"Configurations,omitempty"`
}

type GetVideoEncoderConfiguration struct {
	ConfigurationToken *common.ReferenceToken `xml:"ConfigurationToken,omitempty" json:"ConfigurationToken,omitempty" yaml:"ConfigurationToken,omitempty"`
}

type GetVideoEncoderConfigurationOptions struct {
	ConfigurationToken *common.ReferenceToken `xml:"ConfigurationToken,omitempty" json:"ConfigurationToken,omitempty" yaml:"ConfigurationToken,omitempty"`
	ProfileToken       *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
}

type GetVideoEncoderConfigurationOptionsResponse struct {
	Options *common.VideoEncoderConfigurationOptions `xml:"Options,omitempty" json:"Options,omitempty" yaml:"Options,omitempty"`
}

type GetVideoEncoderConfigurationResponse struct {
	Configuration *common.VideoEncoderConfiguration `xml:"Configuration,omitempty" json:"Configuration,omitempty" yaml:"Configuration,omitempty"`
}

type GetVideoEncoderConfigurations struct {
}

type GetVideoEncoderConfigurationsResponse struct {
	Configurations []*common.VideoEncoderConfiguration `xml:"Configurations,omitempty" json:"Configurations,omitempty" yaml:"Configurations,omitempty"`
}

type GetVideoSourceConfiguration struct {
	ConfigurationToken *common.ReferenceToken `xml:"ConfigurationToken,omitempty" json:"ConfigurationToken,omitempty" yaml:"ConfigurationToken,omitempty"`
}

type GetVideoSourceConfigurationOptions struct {
	ConfigurationToken *common.ReferenceToken `xml:"ConfigurationToken,omitempty" json:"ConfigurationToken,omitempty" yaml:"ConfigurationToken,omitempty"`
	ProfileToken       *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
}

type GetVideoSourceConfigurationOptionsResponse struct {
	Options *common.VideoSourceConfigurationOptions `xml:"Options,omitempty" json:"Options,omitempty" yaml:"Options,omitempty"`
}

type GetVideoSourceConfigurationResponse struct {
	Configuration *common.VideoSourceConfiguration `xml:"Configuration,omitempty" json:"Configuration,omitempty" yaml:"Configuration,omitempty"`
}

type GetVideoSourceConfigurations struct {
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

type GetVideoSources struct {
}

type GetVideoSourcesResponse struct {
	VideoSources []*common.VideoSource `xml:"VideoSources,omitempty" json:"VideoSources,omitempty" yaml:"VideoSources,omitempty"`
}

type ProfileCapabilities []interface{}

type RemoveAudioDecoderConfiguration struct {
	ProfileToken *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
}

type RemoveAudioDecoderConfigurationResponse struct {
}

type RemoveAudioEncoderConfiguration struct {
	ProfileToken *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
}

type RemoveAudioEncoderConfigurationResponse struct {
}

type RemoveAudioOutputConfiguration struct {
	ProfileToken *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
}

type RemoveAudioOutputConfigurationResponse struct {
}

type RemoveAudioSourceConfiguration struct {
	ProfileToken *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
}

type RemoveAudioSourceConfigurationResponse struct {
}

type RemoveMetadataConfiguration struct {
	ProfileToken *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
}

type RemoveMetadataConfigurationResponse struct {
}

type RemovePTZConfiguration struct {
	ProfileToken *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
}

type RemovePTZConfigurationResponse struct {
}

type RemoveVideoAnalyticsConfiguration struct {
	ProfileToken *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
}

type RemoveVideoAnalyticsConfigurationResponse struct {
}

type RemoveVideoEncoderConfiguration struct {
	ProfileToken *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
}

type RemoveVideoEncoderConfigurationResponse struct {
}

type RemoveVideoSourceConfiguration struct {
	ProfileToken *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
}

type RemoveVideoSourceConfigurationResponse struct {
}

type SetAudioDecoderConfiguration struct {
	Configuration    *common.AudioDecoderConfiguration `xml:"Configuration,omitempty" json:"Configuration,omitempty" yaml:"Configuration,omitempty"`
	ForcePersistence bool                              `xml:"ForcePersistence,omitempty" json:"ForcePersistence,omitempty" yaml:"ForcePersistence,omitempty"`
}

type SetAudioDecoderConfigurationResponse struct {
}

type SetAudioEncoderConfiguration struct {
	Configuration    *common.AudioEncoderConfiguration `xml:"Configuration,omitempty" json:"Configuration,omitempty" yaml:"Configuration,omitempty"`
	ForcePersistence bool                              `xml:"ForcePersistence,omitempty" json:"ForcePersistence,omitempty" yaml:"ForcePersistence,omitempty"`
}

type SetAudioEncoderConfigurationResponse struct {
}

type SetAudioOutputConfiguration struct {
	Configuration    *common.AudioOutputConfiguration `xml:"Configuration,omitempty" json:"Configuration,omitempty" yaml:"Configuration,omitempty"`
	ForcePersistence bool                             `xml:"ForcePersistence,omitempty" json:"ForcePersistence,omitempty" yaml:"ForcePersistence,omitempty"`
}

type SetAudioOutputConfigurationResponse struct {
}

type SetAudioSourceConfiguration struct {
	Configuration    *common.AudioSourceConfiguration `xml:"Configuration,omitempty" json:"Configuration,omitempty" yaml:"Configuration,omitempty"`
	ForcePersistence bool                             `xml:"ForcePersistence,omitempty" json:"ForcePersistence,omitempty" yaml:"ForcePersistence,omitempty"`
}

type SetAudioSourceConfigurationResponse struct {
}

type SetMetadataConfiguration struct {
	Configuration    *common.MetadataConfiguration `xml:"Configuration,omitempty" json:"Configuration,omitempty" yaml:"Configuration,omitempty"`
	ForcePersistence bool                          `xml:"ForcePersistence,omitempty" json:"ForcePersistence,omitempty" yaml:"ForcePersistence,omitempty"`
}

type SetMetadataConfigurationResponse struct {
}

type SetOSD struct {
	OSD *common.OSDConfiguration `xml:"OSD,omitempty" json:"OSD,omitempty" yaml:"OSD,omitempty"`
}

type SetOSDResponse []interface{}

type SetSynchronizationPoint struct {
	ProfileToken *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
}

type SetSynchronizationPointResponse struct {
}

type SetVideoAnalyticsConfiguration struct {
	Configuration    *common.VideoAnalyticsConfiguration `xml:"Configuration,omitempty" json:"Configuration,omitempty" yaml:"Configuration,omitempty"`
	ForcePersistence bool                                `xml:"ForcePersistence,omitempty" json:"ForcePersistence,omitempty" yaml:"ForcePersistence,omitempty"`
}

type SetVideoAnalyticsConfigurationResponse struct {
}

type SetVideoEncoderConfiguration struct {
	Configuration    *common.VideoEncoderConfiguration `xml:"Configuration,omitempty" json:"Configuration,omitempty" yaml:"Configuration,omitempty"`
	ForcePersistence bool                              `xml:"ForcePersistence,omitempty" json:"ForcePersistence,omitempty" yaml:"ForcePersistence,omitempty"`
}

type SetVideoEncoderConfigurationResponse struct {
}

type SetVideoSourceConfiguration struct {
	Configuration    *common.VideoSourceConfiguration `xml:"Configuration,omitempty" json:"Configuration,omitempty" yaml:"Configuration,omitempty"`
	ForcePersistence bool                             `xml:"ForcePersistence,omitempty" json:"ForcePersistence,omitempty" yaml:"ForcePersistence,omitempty"`
}

type SetVideoSourceConfigurationResponse struct {
}

type SetVideoSourceMode struct {
	VideoSourceToken     *common.ReferenceToken `xml:"VideoSourceToken,omitempty" json:"VideoSourceToken,omitempty" yaml:"VideoSourceToken,omitempty"`
	VideoSourceModeToken *common.ReferenceToken `xml:"VideoSourceModeToken,omitempty" json:"VideoSourceModeToken,omitempty" yaml:"VideoSourceModeToken,omitempty"`
}

type SetVideoSourceModeResponse struct {
	Reboot bool `xml:"Reboot,omitempty" json:"Reboot,omitempty" yaml:"Reboot,omitempty"`
}

type StartMulticastStreaming struct {
	ProfileToken *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
}

type StartMulticastStreamingResponse struct {
}

type StopMulticastStreaming struct {
	ProfileToken *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
}

type StopMulticastStreamingResponse struct {
}

type StreamingCapabilities []interface{}

type VideoSourceMode struct {
	MaxFramerate  float64                  `xml:"MaxFramerate,omitempty" json:"MaxFramerate,omitempty" yaml:"MaxFramerate,omitempty"`
	MaxResolution *common.VideoResolution  `xml:"MaxResolution,omitempty" json:"MaxResolution,omitempty" yaml:"MaxResolution,omitempty"`
	Encodings     *common.StringList       `xml:"Encodings,omitempty" json:"Encodings,omitempty" yaml:"Encodings,omitempty"`
	Reboot        bool                     `xml:"Reboot,omitempty" json:"Reboot,omitempty" yaml:"Reboot,omitempty"`
	Description   *common.Description      `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	Extension     VideoSourceModeExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
	Token         common.ReferenceToken    `xml:"token,attr,omitempty" json:"token,attr,omitempty" yaml:"token,attr,omitempty"`
	Enabled       bool                     `xml:"Enabled,attr,omitempty" json:"Enabled,attr,omitempty" yaml:"Enabled,attr,omitempty"`
}

type VideoSourceModeExtension []interface{}

// media implements the Media interface.
type media struct {
	cli      common.Client
	Endpoint string
}

func (p *media) OptAddAudioDecoderConfiguration(args AddAudioDecoderConfiguration) (*AddAudioDecoderConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                      string `xml:"trt:AddAudioDecoderConfiguration"`
		AddAudioDecoderConfiguration AddAudioDecoderConfiguration
	}{
		AddAudioDecoderConfiguration: args,
	}

	resp := AddAudioDecoderConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media) OptAddAudioEncoderConfiguration(args AddAudioEncoderConfiguration) (*AddAudioEncoderConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                      string `xml:"trt:AddAudioEncoderConfiguration"`
		AddAudioEncoderConfiguration AddAudioEncoderConfiguration
	}{
		AddAudioEncoderConfiguration: args,
	}

	resp := AddAudioEncoderConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media) OptAddAudioOutputConfiguration(args AddAudioOutputConfiguration) (*AddAudioOutputConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                     string `xml:"trt:AddAudioOutputConfiguration"`
		AddAudioOutputConfiguration AddAudioOutputConfiguration
	}{
		AddAudioOutputConfiguration: args,
	}

	resp := AddAudioOutputConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media) OptAddAudioSourceConfiguration(args AddAudioSourceConfiguration) (*AddAudioSourceConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                     string `xml:"trt:AddAudioSourceConfiguration"`
		AddAudioSourceConfiguration AddAudioSourceConfiguration
	}{
		AddAudioSourceConfiguration: args,
	}

	resp := AddAudioSourceConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media) OptAddMetadataConfiguration(args AddMetadataConfiguration) (*AddMetadataConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                  string `xml:"trt:AddMetadataConfiguration"`
		AddMetadataConfiguration AddMetadataConfiguration
	}{
		AddMetadataConfiguration: args,
	}

	resp := AddMetadataConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media) OptAddPTZConfiguration(args AddPTZConfiguration) (*AddPTZConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName             string `xml:"trt:AddPTZConfiguration"`
		AddPTZConfiguration AddPTZConfiguration
	}{
		AddPTZConfiguration: args,
	}

	resp := AddPTZConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media) OptAddVideoAnalyticsConfiguration(args AddVideoAnalyticsConfiguration) (*AddVideoAnalyticsConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                        string `xml:"trt:AddVideoAnalyticsConfiguration"`
		AddVideoAnalyticsConfiguration AddVideoAnalyticsConfiguration
	}{
		AddVideoAnalyticsConfiguration: args,
	}

	resp := AddVideoAnalyticsConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media) OptAddVideoEncoderConfiguration(args AddVideoEncoderConfiguration) (*AddVideoEncoderConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                      string `xml:"trt:AddVideoEncoderConfiguration"`
		AddVideoEncoderConfiguration AddVideoEncoderConfiguration
	}{
		AddVideoEncoderConfiguration: args,
	}

	resp := AddVideoEncoderConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media) OptAddVideoSourceConfiguration(args AddVideoSourceConfiguration) (*AddVideoSourceConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                     string `xml:"trt:AddVideoSourceConfiguration"`
		AddVideoSourceConfiguration AddVideoSourceConfiguration
	}{
		AddVideoSourceConfiguration: args,
	}

	resp := AddVideoSourceConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media) OptCreateOSD(args CreateOSD) (*CreateOSDResponse, *common.Fault) {
	req := struct {
		XMLName   string `xml:"trt:CreateOSD"`
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

func (p *media) OptCreateProfile(args CreateProfile) (*CreateProfileResponse, *common.Fault) {
	req := struct {
		XMLName       string `xml:"trt:CreateProfile"`
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

func (p *media) OptDeleteOSD(args DeleteOSD) (*DeleteOSDResponse, *common.Fault) {
	req := struct {
		XMLName   string `xml:"trt:DeleteOSD"`
		DeleteOSD DeleteOSD
	}{
		DeleteOSD: args,
	}

	resp := DeleteOSDResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media) OptDeleteProfile(args DeleteProfile) (*DeleteProfileResponse, *common.Fault) {
	req := struct {
		XMLName       string `xml:"trt:DeleteProfile"`
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

func (p *media) OptGetAudioDecoderConfiguration(args GetAudioDecoderConfiguration) (*GetAudioDecoderConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                      string `xml:"trt:GetAudioDecoderConfiguration"`
		GetAudioDecoderConfiguration GetAudioDecoderConfiguration
	}{
		GetAudioDecoderConfiguration: args,
	}

	resp := GetAudioDecoderConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media) OptGetAudioDecoderConfigurationOptions(args GetAudioDecoderConfigurationOptions) (*GetAudioDecoderConfigurationOptionsResponse, *common.Fault) {
	req := struct {
		XMLName                             string `xml:"trt:GetAudioDecoderConfigurationOptions"`
		GetAudioDecoderConfigurationOptions GetAudioDecoderConfigurationOptions
	}{
		GetAudioDecoderConfigurationOptions: args,
	}

	resp := GetAudioDecoderConfigurationOptionsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media) OptGetAudioDecoderConfigurations(args GetAudioDecoderConfigurations) (*GetAudioDecoderConfigurationsResponse, *common.Fault) {
	req := struct {
		XMLName                       string `xml:"trt:GetAudioDecoderConfigurations"`
		GetAudioDecoderConfigurations GetAudioDecoderConfigurations
	}{
		GetAudioDecoderConfigurations: args,
	}

	resp := GetAudioDecoderConfigurationsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media) OptGetAudioEncoderConfiguration(args GetAudioEncoderConfiguration) (*GetAudioEncoderConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                      string `xml:"trt:GetAudioEncoderConfiguration"`
		GetAudioEncoderConfiguration GetAudioEncoderConfiguration
	}{
		GetAudioEncoderConfiguration: args,
	}

	resp := GetAudioEncoderConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media) OptGetAudioEncoderConfigurationOptions(args GetAudioEncoderConfigurationOptions) (*GetAudioEncoderConfigurationOptionsResponse, *common.Fault) {
	req := struct {
		XMLName                             string `xml:"trt:GetAudioEncoderConfigurationOptions"`
		GetAudioEncoderConfigurationOptions GetAudioEncoderConfigurationOptions
	}{
		GetAudioEncoderConfigurationOptions: args,
	}

	resp := GetAudioEncoderConfigurationOptionsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media) OptGetAudioEncoderConfigurations(args GetAudioEncoderConfigurations) (*GetAudioEncoderConfigurationsResponse, *common.Fault) {
	req := struct {
		XMLName                       string `xml:"trt:GetAudioEncoderConfigurations"`
		GetAudioEncoderConfigurations GetAudioEncoderConfigurations
	}{
		GetAudioEncoderConfigurations: args,
	}

	resp := GetAudioEncoderConfigurationsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media) OptGetAudioOutputConfiguration(args GetAudioOutputConfiguration) (*GetAudioOutputConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                     string `xml:"trt:GetAudioOutputConfiguration"`
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

func (p *media) OptGetAudioOutputConfigurationOptions(args GetAudioOutputConfigurationOptions) (*GetAudioOutputConfigurationOptionsResponse, *common.Fault) {
	req := struct {
		XMLName                            string `xml:"trt:GetAudioOutputConfigurationOptions"`
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

func (p *media) OptGetAudioOutputConfigurations(args GetAudioOutputConfigurations) (*GetAudioOutputConfigurationsResponse, *common.Fault) {
	req := struct {
		XMLName                      string `xml:"trt:GetAudioOutputConfigurations"`
		GetAudioOutputConfigurations GetAudioOutputConfigurations
	}{
		GetAudioOutputConfigurations: args,
	}

	resp := GetAudioOutputConfigurationsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media) OptGetAudioOutputs(args GetAudioOutputs) (*GetAudioOutputsResponse, *common.Fault) {
	req := struct {
		XMLName         string `xml:"trt:GetAudioOutputs"`
		GetAudioOutputs GetAudioOutputs
	}{
		GetAudioOutputs: args,
	}

	resp := GetAudioOutputsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media) OptGetAudioSourceConfiguration(args GetAudioSourceConfiguration) (*GetAudioSourceConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                     string `xml:"trt:GetAudioSourceConfiguration"`
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

func (p *media) OptGetAudioSourceConfigurationOptions(args GetAudioSourceConfigurationOptions) (*GetAudioSourceConfigurationOptionsResponse, *common.Fault) {
	req := struct {
		XMLName                            string `xml:"trt:GetAudioSourceConfigurationOptions"`
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

func (p *media) OptGetAudioSourceConfigurations(args GetAudioSourceConfigurations) (*GetAudioSourceConfigurationsResponse, *common.Fault) {
	req := struct {
		XMLName                      string `xml:"trt:GetAudioSourceConfigurations"`
		GetAudioSourceConfigurations GetAudioSourceConfigurations
	}{
		GetAudioSourceConfigurations: args,
	}

	resp := GetAudioSourceConfigurationsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media) OptGetAudioSources(args GetAudioSources) (*GetAudioSourcesResponse, *common.Fault) {
	req := struct {
		XMLName         string `xml:"trt:GetAudioSources"`
		GetAudioSources GetAudioSources
	}{
		GetAudioSources: args,
	}

	resp := GetAudioSourcesResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media) OptGetCompatibleAudioDecoderConfigurations(args GetCompatibleAudioDecoderConfigurations) (*GetCompatibleAudioDecoderConfigurationsResponse, *common.Fault) {
	req := struct {
		XMLName                                 string `xml:"trt:GetCompatibleAudioDecoderConfigurations"`
		GetCompatibleAudioDecoderConfigurations GetCompatibleAudioDecoderConfigurations
	}{
		GetCompatibleAudioDecoderConfigurations: args,
	}

	resp := GetCompatibleAudioDecoderConfigurationsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media) OptGetCompatibleAudioEncoderConfigurations(args GetCompatibleAudioEncoderConfigurations) (*GetCompatibleAudioEncoderConfigurationsResponse, *common.Fault) {
	req := struct {
		XMLName                                 string `xml:"trt:GetCompatibleAudioEncoderConfigurations"`
		GetCompatibleAudioEncoderConfigurations GetCompatibleAudioEncoderConfigurations
	}{
		GetCompatibleAudioEncoderConfigurations: args,
	}

	resp := GetCompatibleAudioEncoderConfigurationsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media) OptGetCompatibleAudioOutputConfigurations(args GetCompatibleAudioOutputConfigurations) (*GetCompatibleAudioOutputConfigurationsResponse, *common.Fault) {
	req := struct {
		XMLName                                string `xml:"trt:GetCompatibleAudioOutputConfigurations"`
		GetCompatibleAudioOutputConfigurations GetCompatibleAudioOutputConfigurations
	}{
		GetCompatibleAudioOutputConfigurations: args,
	}

	resp := GetCompatibleAudioOutputConfigurationsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media) OptGetCompatibleAudioSourceConfigurations(args GetCompatibleAudioSourceConfigurations) (*GetCompatibleAudioSourceConfigurationsResponse, *common.Fault) {
	req := struct {
		XMLName                                string `xml:"trt:GetCompatibleAudioSourceConfigurations"`
		GetCompatibleAudioSourceConfigurations GetCompatibleAudioSourceConfigurations
	}{
		GetCompatibleAudioSourceConfigurations: args,
	}

	resp := GetCompatibleAudioSourceConfigurationsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media) OptGetCompatibleMetadataConfigurations(args GetCompatibleMetadataConfigurations) (*GetCompatibleMetadataConfigurationsResponse, *common.Fault) {
	req := struct {
		XMLName                             string `xml:"trt:GetCompatibleMetadataConfigurations"`
		GetCompatibleMetadataConfigurations GetCompatibleMetadataConfigurations
	}{
		GetCompatibleMetadataConfigurations: args,
	}

	resp := GetCompatibleMetadataConfigurationsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media) OptGetCompatibleVideoAnalyticsConfigurations(args GetCompatibleVideoAnalyticsConfigurations) (*GetCompatibleVideoAnalyticsConfigurationsResponse, *common.Fault) {
	req := struct {
		XMLName                                   string `xml:"trt:GetCompatibleVideoAnalyticsConfigurations"`
		GetCompatibleVideoAnalyticsConfigurations GetCompatibleVideoAnalyticsConfigurations
	}{
		GetCompatibleVideoAnalyticsConfigurations: args,
	}

	resp := GetCompatibleVideoAnalyticsConfigurationsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media) OptGetCompatibleVideoEncoderConfigurations(args GetCompatibleVideoEncoderConfigurations) (*GetCompatibleVideoEncoderConfigurationsResponse, *common.Fault) {
	req := struct {
		XMLName                                 string `xml:"trt:GetCompatibleVideoEncoderConfigurations"`
		GetCompatibleVideoEncoderConfigurations GetCompatibleVideoEncoderConfigurations
	}{
		GetCompatibleVideoEncoderConfigurations: args,
	}

	resp := GetCompatibleVideoEncoderConfigurationsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media) OptGetCompatibleVideoSourceConfigurations(args GetCompatibleVideoSourceConfigurations) (*GetCompatibleVideoSourceConfigurationsResponse, *common.Fault) {
	req := struct {
		XMLName                                string `xml:"trt:GetCompatibleVideoSourceConfigurations"`
		GetCompatibleVideoSourceConfigurations GetCompatibleVideoSourceConfigurations
	}{
		GetCompatibleVideoSourceConfigurations: args,
	}

	resp := GetCompatibleVideoSourceConfigurationsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media) OptGetGuaranteedNumberOfVideoEncoderInstances(args GetGuaranteedNumberOfVideoEncoderInstances) (*GetGuaranteedNumberOfVideoEncoderInstancesResponse, *common.Fault) {
	req := struct {
		XMLName                                    string `xml:"trt:GetGuaranteedNumberOfVideoEncoderInstances"`
		GetGuaranteedNumberOfVideoEncoderInstances GetGuaranteedNumberOfVideoEncoderInstances
	}{
		GetGuaranteedNumberOfVideoEncoderInstances: args,
	}

	resp := GetGuaranteedNumberOfVideoEncoderInstancesResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media) OptGetMetadataConfiguration(args GetMetadataConfiguration) (*GetMetadataConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                  string `xml:"trt:GetMetadataConfiguration"`
		GetMetadataConfiguration GetMetadataConfiguration
	}{
		GetMetadataConfiguration: args,
	}

	resp := GetMetadataConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media) OptGetMetadataConfigurationOptions(args GetMetadataConfigurationOptions) (*GetMetadataConfigurationOptionsResponse, *common.Fault) {
	req := struct {
		XMLName                         string `xml:"trt:GetMetadataConfigurationOptions"`
		GetMetadataConfigurationOptions GetMetadataConfigurationOptions
	}{
		GetMetadataConfigurationOptions: args,
	}

	resp := GetMetadataConfigurationOptionsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media) OptGetMetadataConfigurations(args GetMetadataConfigurations) (*GetMetadataConfigurationsResponse, *common.Fault) {
	req := struct {
		XMLName                   string `xml:"trt:GetMetadataConfigurations"`
		GetMetadataConfigurations GetMetadataConfigurations
	}{
		GetMetadataConfigurations: args,
	}

	resp := GetMetadataConfigurationsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media) OptGetOSD(args GetOSD) (*GetOSDResponse, *common.Fault) {
	req := struct {
		XMLName string `xml:"trt:GetOSD"`
		GetOSD  GetOSD
	}{
		GetOSD: args,
	}

	resp := GetOSDResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media) OptGetOSDOptions(args GetOSDOptions) (*GetOSDOptionsResponse, *common.Fault) {
	req := struct {
		XMLName       string `xml:"trt:GetOSDOptions"`
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

func (p *media) OptGetOSDs(args GetOSDs) (*GetOSDsResponse, *common.Fault) {
	req := struct {
		XMLName string `xml:"trt:GetOSDs"`
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

func (p *media) OptGetProfile(args GetProfile) (*GetProfileResponse, *common.Fault) {
	req := struct {
		XMLName    string `xml:"trt:GetProfile"`
		GetProfile GetProfile
	}{
		GetProfile: args,
	}

	resp := GetProfileResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media) OptGetProfiles(args GetProfiles) (*GetProfilesResponse, *common.Fault) {
	req := struct {
		XMLName     string `xml:"trt:GetProfiles"`
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

func (p *media) OptGetServiceCapabilities(args GetServiceCapabilities) (*GetServiceCapabilitiesResponse, *common.Fault) {
	req := struct {
		XMLName                string `xml:"trt:GetServiceCapabilities"`
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

func (p *media) OptGetSnapshotUri(args GetSnapshotUri) (*GetSnapshotUriResponse, *common.Fault) {
	req := struct {
		XMLName        string `xml:"trt:GetSnapshotUri"`
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

func (p *media) OptGetStreamUri(args GetStreamUri) (*GetStreamUriResponse, *common.Fault) {
	req := struct {
		XMLName      string `xml:"trt:GetStreamUri"`
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

func (p *media) OptGetVideoAnalyticsConfiguration(args GetVideoAnalyticsConfiguration) (*GetVideoAnalyticsConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                        string `xml:"trt:GetVideoAnalyticsConfiguration"`
		GetVideoAnalyticsConfiguration GetVideoAnalyticsConfiguration
	}{
		GetVideoAnalyticsConfiguration: args,
	}

	resp := GetVideoAnalyticsConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media) OptGetVideoAnalyticsConfigurations(args GetVideoAnalyticsConfigurations) (*GetVideoAnalyticsConfigurationsResponse, *common.Fault) {
	req := struct {
		XMLName                         string `xml:"trt:GetVideoAnalyticsConfigurations"`
		GetVideoAnalyticsConfigurations GetVideoAnalyticsConfigurations
	}{
		GetVideoAnalyticsConfigurations: args,
	}

	resp := GetVideoAnalyticsConfigurationsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media) OptGetVideoEncoderConfiguration(args GetVideoEncoderConfiguration) (*GetVideoEncoderConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                      string `xml:"trt:GetVideoEncoderConfiguration"`
		GetVideoEncoderConfiguration GetVideoEncoderConfiguration
	}{
		GetVideoEncoderConfiguration: args,
	}

	resp := GetVideoEncoderConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media) OptGetVideoEncoderConfigurationOptions(args GetVideoEncoderConfigurationOptions) (*GetVideoEncoderConfigurationOptionsResponse, *common.Fault) {
	req := struct {
		XMLName                             string `xml:"trt:GetVideoEncoderConfigurationOptions"`
		GetVideoEncoderConfigurationOptions GetVideoEncoderConfigurationOptions
	}{
		GetVideoEncoderConfigurationOptions: args,
	}

	resp := GetVideoEncoderConfigurationOptionsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media) OptGetVideoEncoderConfigurations(args GetVideoEncoderConfigurations) (*GetVideoEncoderConfigurationsResponse, *common.Fault) {
	req := struct {
		XMLName                       string `xml:"trt:GetVideoEncoderConfigurations"`
		GetVideoEncoderConfigurations GetVideoEncoderConfigurations
	}{
		GetVideoEncoderConfigurations: args,
	}

	resp := GetVideoEncoderConfigurationsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media) OptGetVideoSourceConfiguration(args GetVideoSourceConfiguration) (*GetVideoSourceConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                     string `xml:"trt:GetVideoSourceConfiguration"`
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

func (p *media) OptGetVideoSourceConfigurationOptions(args GetVideoSourceConfigurationOptions) (*GetVideoSourceConfigurationOptionsResponse, *common.Fault) {
	req := struct {
		XMLName                            string `xml:"trt:GetVideoSourceConfigurationOptions"`
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

func (p *media) OptGetVideoSourceConfigurations(args GetVideoSourceConfigurations) (*GetVideoSourceConfigurationsResponse, *common.Fault) {
	req := struct {
		XMLName                      string `xml:"trt:GetVideoSourceConfigurations"`
		GetVideoSourceConfigurations GetVideoSourceConfigurations
	}{
		GetVideoSourceConfigurations: args,
	}

	resp := GetVideoSourceConfigurationsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media) OptGetVideoSourceModes(args GetVideoSourceModes) (*GetVideoSourceModesResponse, *common.Fault) {
	req := struct {
		XMLName             string `xml:"trt:GetVideoSourceModes"`
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

func (p *media) OptGetVideoSources(args GetVideoSources) (*GetVideoSourcesResponse, *common.Fault) {
	req := struct {
		XMLName         string `xml:"trt:GetVideoSources"`
		GetVideoSources GetVideoSources
	}{
		GetVideoSources: args,
	}

	resp := GetVideoSourcesResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media) OptRemoveAudioDecoderConfiguration(args RemoveAudioDecoderConfiguration) (*RemoveAudioDecoderConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                         string `xml:"trt:RemoveAudioDecoderConfiguration"`
		RemoveAudioDecoderConfiguration RemoveAudioDecoderConfiguration
	}{
		RemoveAudioDecoderConfiguration: args,
	}

	resp := RemoveAudioDecoderConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media) OptRemoveAudioEncoderConfiguration(args RemoveAudioEncoderConfiguration) (*RemoveAudioEncoderConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                         string `xml:"trt:RemoveAudioEncoderConfiguration"`
		RemoveAudioEncoderConfiguration RemoveAudioEncoderConfiguration
	}{
		RemoveAudioEncoderConfiguration: args,
	}

	resp := RemoveAudioEncoderConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media) OptRemoveAudioOutputConfiguration(args RemoveAudioOutputConfiguration) (*RemoveAudioOutputConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                        string `xml:"trt:RemoveAudioOutputConfiguration"`
		RemoveAudioOutputConfiguration RemoveAudioOutputConfiguration
	}{
		RemoveAudioOutputConfiguration: args,
	}

	resp := RemoveAudioOutputConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media) OptRemoveAudioSourceConfiguration(args RemoveAudioSourceConfiguration) (*RemoveAudioSourceConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                        string `xml:"trt:RemoveAudioSourceConfiguration"`
		RemoveAudioSourceConfiguration RemoveAudioSourceConfiguration
	}{
		RemoveAudioSourceConfiguration: args,
	}

	resp := RemoveAudioSourceConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media) OptRemoveMetadataConfiguration(args RemoveMetadataConfiguration) (*RemoveMetadataConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                     string `xml:"trt:RemoveMetadataConfiguration"`
		RemoveMetadataConfiguration RemoveMetadataConfiguration
	}{
		RemoveMetadataConfiguration: args,
	}

	resp := RemoveMetadataConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media) OptRemovePTZConfiguration(args RemovePTZConfiguration) (*RemovePTZConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                string `xml:"trt:RemovePTZConfiguration"`
		RemovePTZConfiguration RemovePTZConfiguration
	}{
		RemovePTZConfiguration: args,
	}

	resp := RemovePTZConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media) OptRemoveVideoAnalyticsConfiguration(args RemoveVideoAnalyticsConfiguration) (*RemoveVideoAnalyticsConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                           string `xml:"trt:RemoveVideoAnalyticsConfiguration"`
		RemoveVideoAnalyticsConfiguration RemoveVideoAnalyticsConfiguration
	}{
		RemoveVideoAnalyticsConfiguration: args,
	}

	resp := RemoveVideoAnalyticsConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media) OptRemoveVideoEncoderConfiguration(args RemoveVideoEncoderConfiguration) (*RemoveVideoEncoderConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                         string `xml:"trt:RemoveVideoEncoderConfiguration"`
		RemoveVideoEncoderConfiguration RemoveVideoEncoderConfiguration
	}{
		RemoveVideoEncoderConfiguration: args,
	}

	resp := RemoveVideoEncoderConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media) OptRemoveVideoSourceConfiguration(args RemoveVideoSourceConfiguration) (*RemoveVideoSourceConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                        string `xml:"trt:RemoveVideoSourceConfiguration"`
		RemoveVideoSourceConfiguration RemoveVideoSourceConfiguration
	}{
		RemoveVideoSourceConfiguration: args,
	}

	resp := RemoveVideoSourceConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media) OptSetAudioDecoderConfiguration(args SetAudioDecoderConfiguration) (*SetAudioDecoderConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                      string `xml:"trt:SetAudioDecoderConfiguration"`
		SetAudioDecoderConfiguration SetAudioDecoderConfiguration
	}{
		SetAudioDecoderConfiguration: args,
	}

	resp := SetAudioDecoderConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media) OptSetAudioEncoderConfiguration(args SetAudioEncoderConfiguration) (*SetAudioEncoderConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                      string `xml:"trt:SetAudioEncoderConfiguration"`
		SetAudioEncoderConfiguration SetAudioEncoderConfiguration
	}{
		SetAudioEncoderConfiguration: args,
	}

	resp := SetAudioEncoderConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media) OptSetAudioOutputConfiguration(args SetAudioOutputConfiguration) (*SetAudioOutputConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                     string `xml:"trt:SetAudioOutputConfiguration"`
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

func (p *media) OptSetAudioSourceConfiguration(args SetAudioSourceConfiguration) (*SetAudioSourceConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                     string `xml:"trt:SetAudioSourceConfiguration"`
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

func (p *media) OptSetMetadataConfiguration(args SetMetadataConfiguration) (*SetMetadataConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                  string `xml:"trt:SetMetadataConfiguration"`
		SetMetadataConfiguration SetMetadataConfiguration
	}{
		SetMetadataConfiguration: args,
	}

	resp := SetMetadataConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media) OptSetOSD(args SetOSD) (*SetOSDResponse, *common.Fault) {
	req := struct {
		XMLName string `xml:"trt:SetOSD"`
		SetOSD  SetOSD
	}{
		SetOSD: args,
	}

	resp := SetOSDResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media) OptSetSynchronizationPoint(args SetSynchronizationPoint) (*SetSynchronizationPointResponse, *common.Fault) {
	req := struct {
		XMLName                 string `xml:"trt:SetSynchronizationPoint"`
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

func (p *media) OptSetVideoAnalyticsConfiguration(args SetVideoAnalyticsConfiguration) (*SetVideoAnalyticsConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                        string `xml:"trt:SetVideoAnalyticsConfiguration"`
		SetVideoAnalyticsConfiguration SetVideoAnalyticsConfiguration
	}{
		SetVideoAnalyticsConfiguration: args,
	}

	resp := SetVideoAnalyticsConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media) OptSetVideoEncoderConfiguration(args SetVideoEncoderConfiguration) (*SetVideoEncoderConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                      string `xml:"trt:SetVideoEncoderConfiguration"`
		SetVideoEncoderConfiguration SetVideoEncoderConfiguration
	}{
		SetVideoEncoderConfiguration: args,
	}

	resp := SetVideoEncoderConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media) OptSetVideoSourceConfiguration(args SetVideoSourceConfiguration) (*SetVideoSourceConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                     string `xml:"trt:SetVideoSourceConfiguration"`
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

func (p *media) OptSetVideoSourceMode(args SetVideoSourceMode) (*SetVideoSourceModeResponse, *common.Fault) {
	req := struct {
		XMLName            string `xml:"trt:SetVideoSourceMode"`
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

func (p *media) OptStartMulticastStreaming(args StartMulticastStreaming) (*StartMulticastStreamingResponse, *common.Fault) {
	req := struct {
		XMLName                 string `xml:"trt:StartMulticastStreaming"`
		StartMulticastStreaming StartMulticastStreaming
	}{
		StartMulticastStreaming: args,
	}

	resp := StartMulticastStreamingResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *media) OptStopMulticastStreaming(args StopMulticastStreaming) (*StopMulticastStreamingResponse, *common.Fault) {
	req := struct {
		XMLName                string `xml:"trt:StopMulticastStreaming"`
		StopMulticastStreaming StopMulticastStreaming
	}{
		StopMulticastStreaming: args,
	}

	resp := StopMulticastStreamingResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}
