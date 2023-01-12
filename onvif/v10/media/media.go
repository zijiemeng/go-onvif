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
	OptAddAudioDecoderConfiguration(AddAudioDecoderConfiguration AddAudioDecoderConfiguration) (*AddAudioDecoderConfigurationResponse, error)

	OptAddAudioEncoderConfiguration(AddAudioEncoderConfiguration AddAudioEncoderConfiguration) (*AddAudioEncoderConfigurationResponse, error)

	OptAddAudioOutputConfiguration(AddAudioOutputConfiguration AddAudioOutputConfiguration) (*AddAudioOutputConfigurationResponse, error)

	OptAddAudioSourceConfiguration(AddAudioSourceConfiguration AddAudioSourceConfiguration) (*AddAudioSourceConfigurationResponse, error)

	OptAddMetadataConfiguration(AddMetadataConfiguration AddMetadataConfiguration) (*AddMetadataConfigurationResponse, error)

	OptAddPTZConfiguration(AddPTZConfiguration AddPTZConfiguration) (*AddPTZConfigurationResponse, error)

	OptAddVideoAnalyticsConfiguration(AddVideoAnalyticsConfiguration AddVideoAnalyticsConfiguration) (*AddVideoAnalyticsConfigurationResponse, error)

	OptAddVideoEncoderConfiguration(AddVideoEncoderConfiguration AddVideoEncoderConfiguration) (*AddVideoEncoderConfigurationResponse, error)

	OptAddVideoSourceConfiguration(AddVideoSourceConfiguration AddVideoSourceConfiguration) (*AddVideoSourceConfigurationResponse, error)

	OptCreateOSD(CreateOSD CreateOSD) (*CreateOSDResponse, error)

	OptCreateProfile(CreateProfile CreateProfile) (*CreateProfileResponse, error)

	OptDeleteOSD(DeleteOSD DeleteOSD) (*DeleteOSDResponse, error)

	OptDeleteProfile(DeleteProfile DeleteProfile) (*DeleteProfileResponse, error)

	OptGetAudioDecoderConfiguration(GetAudioDecoderConfiguration GetAudioDecoderConfiguration) (*GetAudioDecoderConfigurationResponse, error)

	OptGetAudioDecoderConfigurationOptions(GetAudioDecoderConfigurationOptions GetAudioDecoderConfigurationOptions) (*GetAudioDecoderConfigurationOptionsResponse, error)

	OptGetAudioDecoderConfigurations(GetAudioDecoderConfigurations GetAudioDecoderConfigurations) (*GetAudioDecoderConfigurationsResponse, error)

	OptGetAudioEncoderConfiguration(GetAudioEncoderConfiguration GetAudioEncoderConfiguration) (*GetAudioEncoderConfigurationResponse, error)

	OptGetAudioEncoderConfigurationOptions(GetAudioEncoderConfigurationOptions GetAudioEncoderConfigurationOptions) (*GetAudioEncoderConfigurationOptionsResponse, error)

	OptGetAudioEncoderConfigurations(GetAudioEncoderConfigurations GetAudioEncoderConfigurations) (*GetAudioEncoderConfigurationsResponse, error)

	OptGetAudioOutputConfiguration(GetAudioOutputConfiguration GetAudioOutputConfiguration) (*GetAudioOutputConfigurationResponse, error)

	OptGetAudioOutputConfigurationOptions(GetAudioOutputConfigurationOptions GetAudioOutputConfigurationOptions) (*GetAudioOutputConfigurationOptionsResponse, error)

	OptGetAudioOutputConfigurations(GetAudioOutputConfigurations GetAudioOutputConfigurations) (*GetAudioOutputConfigurationsResponse, error)

	OptGetAudioOutputs(GetAudioOutputs GetAudioOutputs) (*GetAudioOutputsResponse, error)

	OptGetAudioSourceConfiguration(GetAudioSourceConfiguration GetAudioSourceConfiguration) (*GetAudioSourceConfigurationResponse, error)

	OptGetAudioSourceConfigurationOptions(GetAudioSourceConfigurationOptions GetAudioSourceConfigurationOptions) (*GetAudioSourceConfigurationOptionsResponse, error)

	OptGetAudioSourceConfigurations(GetAudioSourceConfigurations GetAudioSourceConfigurations) (*GetAudioSourceConfigurationsResponse, error)

	OptGetAudioSources(GetAudioSources GetAudioSources) (*GetAudioSourcesResponse, error)

	OptGetCompatibleAudioDecoderConfigurations(GetCompatibleAudioDecoderConfigurations GetCompatibleAudioDecoderConfigurations) (*GetCompatibleAudioDecoderConfigurationsResponse, error)

	OptGetCompatibleAudioEncoderConfigurations(GetCompatibleAudioEncoderConfigurations GetCompatibleAudioEncoderConfigurations) (*GetCompatibleAudioEncoderConfigurationsResponse, error)

	OptGetCompatibleAudioOutputConfigurations(GetCompatibleAudioOutputConfigurations GetCompatibleAudioOutputConfigurations) (*GetCompatibleAudioOutputConfigurationsResponse, error)

	OptGetCompatibleAudioSourceConfigurations(GetCompatibleAudioSourceConfigurations GetCompatibleAudioSourceConfigurations) (*GetCompatibleAudioSourceConfigurationsResponse, error)

	OptGetCompatibleMetadataConfigurations(GetCompatibleMetadataConfigurations GetCompatibleMetadataConfigurations) (*GetCompatibleMetadataConfigurationsResponse, error)

	OptGetCompatibleVideoAnalyticsConfigurations(GetCompatibleVideoAnalyticsConfigurations GetCompatibleVideoAnalyticsConfigurations) (*GetCompatibleVideoAnalyticsConfigurationsResponse, error)

	OptGetCompatibleVideoEncoderConfigurations(GetCompatibleVideoEncoderConfigurations GetCompatibleVideoEncoderConfigurations) (*GetCompatibleVideoEncoderConfigurationsResponse, error)

	OptGetCompatibleVideoSourceConfigurations(GetCompatibleVideoSourceConfigurations GetCompatibleVideoSourceConfigurations) (*GetCompatibleVideoSourceConfigurationsResponse, error)

	OptGetGuaranteedNumberOfVideoEncoderInstances(GetGuaranteedNumberOfVideoEncoderInstances GetGuaranteedNumberOfVideoEncoderInstances) (*GetGuaranteedNumberOfVideoEncoderInstancesResponse, error)

	OptGetMetadataConfiguration(GetMetadataConfiguration GetMetadataConfiguration) (*GetMetadataConfigurationResponse, error)

	OptGetMetadataConfigurationOptions(GetMetadataConfigurationOptions GetMetadataConfigurationOptions) (*GetMetadataConfigurationOptionsResponse, error)

	OptGetMetadataConfigurations(GetMetadataConfigurations GetMetadataConfigurations) (*GetMetadataConfigurationsResponse, error)

	OptGetOSD(GetOSD GetOSD) (*GetOSDResponse, error)

	OptGetOSDOptions(GetOSDOptions GetOSDOptions) (*GetOSDOptionsResponse, error)

	OptGetOSDs(GetOSDs GetOSDs) (*GetOSDsResponse, error)

	OptGetProfile(GetProfile GetProfile) (*GetProfileResponse, error)

	OptGetProfiles(GetProfiles GetProfiles) (*GetProfilesResponse, error)

	OptGetServiceCapabilities(GetServiceCapabilities GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error)

	OptGetSnapshotUri(GetSnapshotUri GetSnapshotUri) (*GetSnapshotUriResponse, error)

	OptGetStreamUri(GetStreamUri GetStreamUri) (*GetStreamUriResponse, error)

	OptGetVideoAnalyticsConfiguration(GetVideoAnalyticsConfiguration GetVideoAnalyticsConfiguration) (*GetVideoAnalyticsConfigurationResponse, error)

	OptGetVideoAnalyticsConfigurations(GetVideoAnalyticsConfigurations GetVideoAnalyticsConfigurations) (*GetVideoAnalyticsConfigurationsResponse, error)

	OptGetVideoEncoderConfiguration(GetVideoEncoderConfiguration GetVideoEncoderConfiguration) (*GetVideoEncoderConfigurationResponse, error)

	OptGetVideoEncoderConfigurationOptions(GetVideoEncoderConfigurationOptions GetVideoEncoderConfigurationOptions) (*GetVideoEncoderConfigurationOptionsResponse, error)

	OptGetVideoEncoderConfigurations(GetVideoEncoderConfigurations GetVideoEncoderConfigurations) (*GetVideoEncoderConfigurationsResponse, error)

	OptGetVideoSourceConfiguration(GetVideoSourceConfiguration GetVideoSourceConfiguration) (*GetVideoSourceConfigurationResponse, error)

	OptGetVideoSourceConfigurationOptions(GetVideoSourceConfigurationOptions GetVideoSourceConfigurationOptions) (*GetVideoSourceConfigurationOptionsResponse, error)

	OptGetVideoSourceConfigurations(GetVideoSourceConfigurations GetVideoSourceConfigurations) (*GetVideoSourceConfigurationsResponse, error)

	OptGetVideoSourceModes(GetVideoSourceModes GetVideoSourceModes) (*GetVideoSourceModesResponse, error)

	OptGetVideoSources(GetVideoSources GetVideoSources) (*GetVideoSourcesResponse, error)

	OptRemoveAudioDecoderConfiguration(RemoveAudioDecoderConfiguration RemoveAudioDecoderConfiguration) (*RemoveAudioDecoderConfigurationResponse, error)

	OptRemoveAudioEncoderConfiguration(RemoveAudioEncoderConfiguration RemoveAudioEncoderConfiguration) (*RemoveAudioEncoderConfigurationResponse, error)

	OptRemoveAudioOutputConfiguration(RemoveAudioOutputConfiguration RemoveAudioOutputConfiguration) (*RemoveAudioOutputConfigurationResponse, error)

	OptRemoveAudioSourceConfiguration(RemoveAudioSourceConfiguration RemoveAudioSourceConfiguration) (*RemoveAudioSourceConfigurationResponse, error)

	OptRemoveMetadataConfiguration(RemoveMetadataConfiguration RemoveMetadataConfiguration) (*RemoveMetadataConfigurationResponse, error)

	OptRemovePTZConfiguration(RemovePTZConfiguration RemovePTZConfiguration) (*RemovePTZConfigurationResponse, error)

	OptRemoveVideoAnalyticsConfiguration(RemoveVideoAnalyticsConfiguration RemoveVideoAnalyticsConfiguration) (*RemoveVideoAnalyticsConfigurationResponse, error)

	OptRemoveVideoEncoderConfiguration(RemoveVideoEncoderConfiguration RemoveVideoEncoderConfiguration) (*RemoveVideoEncoderConfigurationResponse, error)

	OptRemoveVideoSourceConfiguration(RemoveVideoSourceConfiguration RemoveVideoSourceConfiguration) (*RemoveVideoSourceConfigurationResponse, error)

	OptSetAudioDecoderConfiguration(SetAudioDecoderConfiguration SetAudioDecoderConfiguration) (*SetAudioDecoderConfigurationResponse, error)

	OptSetAudioEncoderConfiguration(SetAudioEncoderConfiguration SetAudioEncoderConfiguration) (*SetAudioEncoderConfigurationResponse, error)

	OptSetAudioOutputConfiguration(SetAudioOutputConfiguration SetAudioOutputConfiguration) (*SetAudioOutputConfigurationResponse, error)

	OptSetAudioSourceConfiguration(SetAudioSourceConfiguration SetAudioSourceConfiguration) (*SetAudioSourceConfigurationResponse, error)

	OptSetMetadataConfiguration(SetMetadataConfiguration SetMetadataConfiguration) (*SetMetadataConfigurationResponse, error)

	OptSetOSD(SetOSD SetOSD) (*SetOSDResponse, error)

	OptSetSynchronizationPoint(SetSynchronizationPoint SetSynchronizationPoint) (*SetSynchronizationPointResponse, error)

	OptSetVideoAnalyticsConfiguration(SetVideoAnalyticsConfiguration SetVideoAnalyticsConfiguration) (*SetVideoAnalyticsConfigurationResponse, error)

	OptSetVideoEncoderConfiguration(SetVideoEncoderConfiguration SetVideoEncoderConfiguration) (*SetVideoEncoderConfigurationResponse, error)

	OptSetVideoSourceConfiguration(SetVideoSourceConfiguration SetVideoSourceConfiguration) (*SetVideoSourceConfigurationResponse, error)

	OptSetVideoSourceMode(SetVideoSourceMode SetVideoSourceMode) (*SetVideoSourceModeResponse, error)

	OptStartMulticastStreaming(StartMulticastStreaming StartMulticastStreaming) (*StartMulticastStreamingResponse, error)

	OptStopMulticastStreaming(StopMulticastStreaming StopMulticastStreaming) (*StopMulticastStreamingResponse, error)
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

func (p *media) OptAddAudioDecoderConfiguration(args AddAudioDecoderConfiguration) (*AddAudioDecoderConfigurationResponse, error) {
	req := struct {
		XMLName                      string `xml:"trt:AddAudioDecoderConfiguration"`
		AddAudioDecoderConfiguration AddAudioDecoderConfiguration
	}{
		AddAudioDecoderConfiguration: args,
	}

	resp := AddAudioDecoderConfigurationResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptAddAudioEncoderConfiguration(args AddAudioEncoderConfiguration) (*AddAudioEncoderConfigurationResponse, error) {
	req := struct {
		XMLName                      string `xml:"trt:AddAudioEncoderConfiguration"`
		AddAudioEncoderConfiguration AddAudioEncoderConfiguration
	}{
		AddAudioEncoderConfiguration: args,
	}

	resp := AddAudioEncoderConfigurationResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptAddAudioOutputConfiguration(args AddAudioOutputConfiguration) (*AddAudioOutputConfigurationResponse, error) {
	req := struct {
		XMLName                     string `xml:"trt:AddAudioOutputConfiguration"`
		AddAudioOutputConfiguration AddAudioOutputConfiguration
	}{
		AddAudioOutputConfiguration: args,
	}

	resp := AddAudioOutputConfigurationResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptAddAudioSourceConfiguration(args AddAudioSourceConfiguration) (*AddAudioSourceConfigurationResponse, error) {
	req := struct {
		XMLName                     string `xml:"trt:AddAudioSourceConfiguration"`
		AddAudioSourceConfiguration AddAudioSourceConfiguration
	}{
		AddAudioSourceConfiguration: args,
	}

	resp := AddAudioSourceConfigurationResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptAddMetadataConfiguration(args AddMetadataConfiguration) (*AddMetadataConfigurationResponse, error) {
	req := struct {
		XMLName                  string `xml:"trt:AddMetadataConfiguration"`
		AddMetadataConfiguration AddMetadataConfiguration
	}{
		AddMetadataConfiguration: args,
	}

	resp := AddMetadataConfigurationResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptAddPTZConfiguration(args AddPTZConfiguration) (*AddPTZConfigurationResponse, error) {
	req := struct {
		XMLName             string `xml:"trt:AddPTZConfiguration"`
		AddPTZConfiguration AddPTZConfiguration
	}{
		AddPTZConfiguration: args,
	}

	resp := AddPTZConfigurationResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptAddVideoAnalyticsConfiguration(args AddVideoAnalyticsConfiguration) (*AddVideoAnalyticsConfigurationResponse, error) {
	req := struct {
		XMLName                        string `xml:"trt:AddVideoAnalyticsConfiguration"`
		AddVideoAnalyticsConfiguration AddVideoAnalyticsConfiguration
	}{
		AddVideoAnalyticsConfiguration: args,
	}

	resp := AddVideoAnalyticsConfigurationResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptAddVideoEncoderConfiguration(args AddVideoEncoderConfiguration) (*AddVideoEncoderConfigurationResponse, error) {
	req := struct {
		XMLName                      string `xml:"trt:AddVideoEncoderConfiguration"`
		AddVideoEncoderConfiguration AddVideoEncoderConfiguration
	}{
		AddVideoEncoderConfiguration: args,
	}

	resp := AddVideoEncoderConfigurationResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptAddVideoSourceConfiguration(args AddVideoSourceConfiguration) (*AddVideoSourceConfigurationResponse, error) {
	req := struct {
		XMLName                     string `xml:"trt:AddVideoSourceConfiguration"`
		AddVideoSourceConfiguration AddVideoSourceConfiguration
	}{
		AddVideoSourceConfiguration: args,
	}

	resp := AddVideoSourceConfigurationResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptCreateOSD(args CreateOSD) (*CreateOSDResponse, error) {
	req := struct {
		XMLName   string `xml:"trt:CreateOSD"`
		CreateOSD CreateOSD
	}{
		CreateOSD: args,
	}

	resp := CreateOSDResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptCreateProfile(args CreateProfile) (*CreateProfileResponse, error) {
	req := struct {
		XMLName       string `xml:"trt:CreateProfile"`
		CreateProfile CreateProfile
	}{
		CreateProfile: args,
	}

	resp := CreateProfileResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptDeleteOSD(args DeleteOSD) (*DeleteOSDResponse, error) {
	req := struct {
		XMLName   string `xml:"trt:DeleteOSD"`
		DeleteOSD DeleteOSD
	}{
		DeleteOSD: args,
	}

	resp := DeleteOSDResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptDeleteProfile(args DeleteProfile) (*DeleteProfileResponse, error) {
	req := struct {
		XMLName       string `xml:"trt:DeleteProfile"`
		DeleteProfile DeleteProfile
	}{
		DeleteProfile: args,
	}

	resp := DeleteProfileResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptGetAudioDecoderConfiguration(args GetAudioDecoderConfiguration) (*GetAudioDecoderConfigurationResponse, error) {
	req := struct {
		XMLName                      string `xml:"trt:GetAudioDecoderConfiguration"`
		GetAudioDecoderConfiguration GetAudioDecoderConfiguration
	}{
		GetAudioDecoderConfiguration: args,
	}

	resp := GetAudioDecoderConfigurationResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptGetAudioDecoderConfigurationOptions(args GetAudioDecoderConfigurationOptions) (*GetAudioDecoderConfigurationOptionsResponse, error) {
	req := struct {
		XMLName                             string `xml:"trt:GetAudioDecoderConfigurationOptions"`
		GetAudioDecoderConfigurationOptions GetAudioDecoderConfigurationOptions
	}{
		GetAudioDecoderConfigurationOptions: args,
	}

	resp := GetAudioDecoderConfigurationOptionsResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptGetAudioDecoderConfigurations(args GetAudioDecoderConfigurations) (*GetAudioDecoderConfigurationsResponse, error) {
	req := struct {
		XMLName                       string `xml:"trt:GetAudioDecoderConfigurations"`
		GetAudioDecoderConfigurations GetAudioDecoderConfigurations
	}{
		GetAudioDecoderConfigurations: args,
	}

	resp := GetAudioDecoderConfigurationsResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptGetAudioEncoderConfiguration(args GetAudioEncoderConfiguration) (*GetAudioEncoderConfigurationResponse, error) {
	req := struct {
		XMLName                      string `xml:"trt:GetAudioEncoderConfiguration"`
		GetAudioEncoderConfiguration GetAudioEncoderConfiguration
	}{
		GetAudioEncoderConfiguration: args,
	}

	resp := GetAudioEncoderConfigurationResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptGetAudioEncoderConfigurationOptions(args GetAudioEncoderConfigurationOptions) (*GetAudioEncoderConfigurationOptionsResponse, error) {
	req := struct {
		XMLName                             string `xml:"trt:GetAudioEncoderConfigurationOptions"`
		GetAudioEncoderConfigurationOptions GetAudioEncoderConfigurationOptions
	}{
		GetAudioEncoderConfigurationOptions: args,
	}

	resp := GetAudioEncoderConfigurationOptionsResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptGetAudioEncoderConfigurations(args GetAudioEncoderConfigurations) (*GetAudioEncoderConfigurationsResponse, error) {
	req := struct {
		XMLName                       string `xml:"trt:GetAudioEncoderConfigurations"`
		GetAudioEncoderConfigurations GetAudioEncoderConfigurations
	}{
		GetAudioEncoderConfigurations: args,
	}

	resp := GetAudioEncoderConfigurationsResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptGetAudioOutputConfiguration(args GetAudioOutputConfiguration) (*GetAudioOutputConfigurationResponse, error) {
	req := struct {
		XMLName                     string `xml:"trt:GetAudioOutputConfiguration"`
		GetAudioOutputConfiguration GetAudioOutputConfiguration
	}{
		GetAudioOutputConfiguration: args,
	}

	resp := GetAudioOutputConfigurationResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptGetAudioOutputConfigurationOptions(args GetAudioOutputConfigurationOptions) (*GetAudioOutputConfigurationOptionsResponse, error) {
	req := struct {
		XMLName                            string `xml:"trt:GetAudioOutputConfigurationOptions"`
		GetAudioOutputConfigurationOptions GetAudioOutputConfigurationOptions
	}{
		GetAudioOutputConfigurationOptions: args,
	}

	resp := GetAudioOutputConfigurationOptionsResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptGetAudioOutputConfigurations(args GetAudioOutputConfigurations) (*GetAudioOutputConfigurationsResponse, error) {
	req := struct {
		XMLName                      string `xml:"trt:GetAudioOutputConfigurations"`
		GetAudioOutputConfigurations GetAudioOutputConfigurations
	}{
		GetAudioOutputConfigurations: args,
	}

	resp := GetAudioOutputConfigurationsResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptGetAudioOutputs(args GetAudioOutputs) (*GetAudioOutputsResponse, error) {
	req := struct {
		XMLName         string `xml:"trt:GetAudioOutputs"`
		GetAudioOutputs GetAudioOutputs
	}{
		GetAudioOutputs: args,
	}

	resp := GetAudioOutputsResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptGetAudioSourceConfiguration(args GetAudioSourceConfiguration) (*GetAudioSourceConfigurationResponse, error) {
	req := struct {
		XMLName                     string `xml:"trt:GetAudioSourceConfiguration"`
		GetAudioSourceConfiguration GetAudioSourceConfiguration
	}{
		GetAudioSourceConfiguration: args,
	}

	resp := GetAudioSourceConfigurationResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptGetAudioSourceConfigurationOptions(args GetAudioSourceConfigurationOptions) (*GetAudioSourceConfigurationOptionsResponse, error) {
	req := struct {
		XMLName                            string `xml:"trt:GetAudioSourceConfigurationOptions"`
		GetAudioSourceConfigurationOptions GetAudioSourceConfigurationOptions
	}{
		GetAudioSourceConfigurationOptions: args,
	}

	resp := GetAudioSourceConfigurationOptionsResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptGetAudioSourceConfigurations(args GetAudioSourceConfigurations) (*GetAudioSourceConfigurationsResponse, error) {
	req := struct {
		XMLName                      string `xml:"trt:GetAudioSourceConfigurations"`
		GetAudioSourceConfigurations GetAudioSourceConfigurations
	}{
		GetAudioSourceConfigurations: args,
	}

	resp := GetAudioSourceConfigurationsResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptGetAudioSources(args GetAudioSources) (*GetAudioSourcesResponse, error) {
	req := struct {
		XMLName         string `xml:"trt:GetAudioSources"`
		GetAudioSources GetAudioSources
	}{
		GetAudioSources: args,
	}

	resp := GetAudioSourcesResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptGetCompatibleAudioDecoderConfigurations(args GetCompatibleAudioDecoderConfigurations) (*GetCompatibleAudioDecoderConfigurationsResponse, error) {
	req := struct {
		XMLName                                 string `xml:"trt:GetCompatibleAudioDecoderConfigurations"`
		GetCompatibleAudioDecoderConfigurations GetCompatibleAudioDecoderConfigurations
	}{
		GetCompatibleAudioDecoderConfigurations: args,
	}

	resp := GetCompatibleAudioDecoderConfigurationsResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptGetCompatibleAudioEncoderConfigurations(args GetCompatibleAudioEncoderConfigurations) (*GetCompatibleAudioEncoderConfigurationsResponse, error) {
	req := struct {
		XMLName                                 string `xml:"trt:GetCompatibleAudioEncoderConfigurations"`
		GetCompatibleAudioEncoderConfigurations GetCompatibleAudioEncoderConfigurations
	}{
		GetCompatibleAudioEncoderConfigurations: args,
	}

	resp := GetCompatibleAudioEncoderConfigurationsResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptGetCompatibleAudioOutputConfigurations(args GetCompatibleAudioOutputConfigurations) (*GetCompatibleAudioOutputConfigurationsResponse, error) {
	req := struct {
		XMLName                                string `xml:"trt:GetCompatibleAudioOutputConfigurations"`
		GetCompatibleAudioOutputConfigurations GetCompatibleAudioOutputConfigurations
	}{
		GetCompatibleAudioOutputConfigurations: args,
	}

	resp := GetCompatibleAudioOutputConfigurationsResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptGetCompatibleAudioSourceConfigurations(args GetCompatibleAudioSourceConfigurations) (*GetCompatibleAudioSourceConfigurationsResponse, error) {
	req := struct {
		XMLName                                string `xml:"trt:GetCompatibleAudioSourceConfigurations"`
		GetCompatibleAudioSourceConfigurations GetCompatibleAudioSourceConfigurations
	}{
		GetCompatibleAudioSourceConfigurations: args,
	}

	resp := GetCompatibleAudioSourceConfigurationsResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptGetCompatibleMetadataConfigurations(args GetCompatibleMetadataConfigurations) (*GetCompatibleMetadataConfigurationsResponse, error) {
	req := struct {
		XMLName                             string `xml:"trt:GetCompatibleMetadataConfigurations"`
		GetCompatibleMetadataConfigurations GetCompatibleMetadataConfigurations
	}{
		GetCompatibleMetadataConfigurations: args,
	}

	resp := GetCompatibleMetadataConfigurationsResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptGetCompatibleVideoAnalyticsConfigurations(args GetCompatibleVideoAnalyticsConfigurations) (*GetCompatibleVideoAnalyticsConfigurationsResponse, error) {
	req := struct {
		XMLName                                   string `xml:"trt:GetCompatibleVideoAnalyticsConfigurations"`
		GetCompatibleVideoAnalyticsConfigurations GetCompatibleVideoAnalyticsConfigurations
	}{
		GetCompatibleVideoAnalyticsConfigurations: args,
	}

	resp := GetCompatibleVideoAnalyticsConfigurationsResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptGetCompatibleVideoEncoderConfigurations(args GetCompatibleVideoEncoderConfigurations) (*GetCompatibleVideoEncoderConfigurationsResponse, error) {
	req := struct {
		XMLName                                 string `xml:"trt:GetCompatibleVideoEncoderConfigurations"`
		GetCompatibleVideoEncoderConfigurations GetCompatibleVideoEncoderConfigurations
	}{
		GetCompatibleVideoEncoderConfigurations: args,
	}

	resp := GetCompatibleVideoEncoderConfigurationsResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptGetCompatibleVideoSourceConfigurations(args GetCompatibleVideoSourceConfigurations) (*GetCompatibleVideoSourceConfigurationsResponse, error) {
	req := struct {
		XMLName                                string `xml:"trt:GetCompatibleVideoSourceConfigurations"`
		GetCompatibleVideoSourceConfigurations GetCompatibleVideoSourceConfigurations
	}{
		GetCompatibleVideoSourceConfigurations: args,
	}

	resp := GetCompatibleVideoSourceConfigurationsResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptGetGuaranteedNumberOfVideoEncoderInstances(args GetGuaranteedNumberOfVideoEncoderInstances) (*GetGuaranteedNumberOfVideoEncoderInstancesResponse, error) {
	req := struct {
		XMLName                                    string `xml:"trt:GetGuaranteedNumberOfVideoEncoderInstances"`
		GetGuaranteedNumberOfVideoEncoderInstances GetGuaranteedNumberOfVideoEncoderInstances
	}{
		GetGuaranteedNumberOfVideoEncoderInstances: args,
	}

	resp := GetGuaranteedNumberOfVideoEncoderInstancesResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptGetMetadataConfiguration(args GetMetadataConfiguration) (*GetMetadataConfigurationResponse, error) {
	req := struct {
		XMLName                  string `xml:"trt:GetMetadataConfiguration"`
		GetMetadataConfiguration GetMetadataConfiguration
	}{
		GetMetadataConfiguration: args,
	}

	resp := GetMetadataConfigurationResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptGetMetadataConfigurationOptions(args GetMetadataConfigurationOptions) (*GetMetadataConfigurationOptionsResponse, error) {
	req := struct {
		XMLName                         string `xml:"trt:GetMetadataConfigurationOptions"`
		GetMetadataConfigurationOptions GetMetadataConfigurationOptions
	}{
		GetMetadataConfigurationOptions: args,
	}

	resp := GetMetadataConfigurationOptionsResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptGetMetadataConfigurations(args GetMetadataConfigurations) (*GetMetadataConfigurationsResponse, error) {
	req := struct {
		XMLName                   string `xml:"trt:GetMetadataConfigurations"`
		GetMetadataConfigurations GetMetadataConfigurations
	}{
		GetMetadataConfigurations: args,
	}

	resp := GetMetadataConfigurationsResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptGetOSD(args GetOSD) (*GetOSDResponse, error) {
	req := struct {
		XMLName string `xml:"trt:GetOSD"`
		GetOSD  GetOSD
	}{
		GetOSD: args,
	}

	resp := GetOSDResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptGetOSDOptions(args GetOSDOptions) (*GetOSDOptionsResponse, error) {
	req := struct {
		XMLName       string `xml:"trt:GetOSDOptions"`
		GetOSDOptions GetOSDOptions
	}{
		GetOSDOptions: args,
	}

	resp := GetOSDOptionsResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptGetOSDs(args GetOSDs) (*GetOSDsResponse, error) {
	req := struct {
		XMLName string `xml:"trt:GetOSDs"`
		GetOSDs GetOSDs
	}{
		GetOSDs: args,
	}

	resp := GetOSDsResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptGetProfile(args GetProfile) (*GetProfileResponse, error) {
	req := struct {
		XMLName    string `xml:"trt:GetProfile"`
		GetProfile GetProfile
	}{
		GetProfile: args,
	}

	resp := GetProfileResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptGetProfiles(args GetProfiles) (*GetProfilesResponse, error) {
	req := struct {
		XMLName     string `xml:"trt:GetProfiles"`
		GetProfiles GetProfiles
	}{
		GetProfiles: args,
	}

	resp := GetProfilesResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptGetServiceCapabilities(args GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error) {
	req := struct {
		XMLName                string `xml:"trt:GetServiceCapabilities"`
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

func (p *media) OptGetSnapshotUri(args GetSnapshotUri) (*GetSnapshotUriResponse, error) {
	req := struct {
		XMLName        string `xml:"trt:GetSnapshotUri"`
		GetSnapshotUri GetSnapshotUri
	}{
		GetSnapshotUri: args,
	}

	resp := GetSnapshotUriResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptGetStreamUri(args GetStreamUri) (*GetStreamUriResponse, error) {
	req := struct {
		XMLName      string `xml:"trt:GetStreamUri"`
		GetStreamUri GetStreamUri
	}{
		GetStreamUri: args,
	}

	resp := GetStreamUriResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptGetVideoAnalyticsConfiguration(args GetVideoAnalyticsConfiguration) (*GetVideoAnalyticsConfigurationResponse, error) {
	req := struct {
		XMLName                        string `xml:"trt:GetVideoAnalyticsConfiguration"`
		GetVideoAnalyticsConfiguration GetVideoAnalyticsConfiguration
	}{
		GetVideoAnalyticsConfiguration: args,
	}

	resp := GetVideoAnalyticsConfigurationResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptGetVideoAnalyticsConfigurations(args GetVideoAnalyticsConfigurations) (*GetVideoAnalyticsConfigurationsResponse, error) {
	req := struct {
		XMLName                         string `xml:"trt:GetVideoAnalyticsConfigurations"`
		GetVideoAnalyticsConfigurations GetVideoAnalyticsConfigurations
	}{
		GetVideoAnalyticsConfigurations: args,
	}

	resp := GetVideoAnalyticsConfigurationsResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptGetVideoEncoderConfiguration(args GetVideoEncoderConfiguration) (*GetVideoEncoderConfigurationResponse, error) {
	req := struct {
		XMLName                      string `xml:"trt:GetVideoEncoderConfiguration"`
		GetVideoEncoderConfiguration GetVideoEncoderConfiguration
	}{
		GetVideoEncoderConfiguration: args,
	}

	resp := GetVideoEncoderConfigurationResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptGetVideoEncoderConfigurationOptions(args GetVideoEncoderConfigurationOptions) (*GetVideoEncoderConfigurationOptionsResponse, error) {
	req := struct {
		XMLName                             string `xml:"trt:GetVideoEncoderConfigurationOptions"`
		GetVideoEncoderConfigurationOptions GetVideoEncoderConfigurationOptions
	}{
		GetVideoEncoderConfigurationOptions: args,
	}

	resp := GetVideoEncoderConfigurationOptionsResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptGetVideoEncoderConfigurations(args GetVideoEncoderConfigurations) (*GetVideoEncoderConfigurationsResponse, error) {
	req := struct {
		XMLName                       string `xml:"trt:GetVideoEncoderConfigurations"`
		GetVideoEncoderConfigurations GetVideoEncoderConfigurations
	}{
		GetVideoEncoderConfigurations: args,
	}

	resp := GetVideoEncoderConfigurationsResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptGetVideoSourceConfiguration(args GetVideoSourceConfiguration) (*GetVideoSourceConfigurationResponse, error) {
	req := struct {
		XMLName                     string `xml:"trt:GetVideoSourceConfiguration"`
		GetVideoSourceConfiguration GetVideoSourceConfiguration
	}{
		GetVideoSourceConfiguration: args,
	}

	resp := GetVideoSourceConfigurationResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptGetVideoSourceConfigurationOptions(args GetVideoSourceConfigurationOptions) (*GetVideoSourceConfigurationOptionsResponse, error) {
	req := struct {
		XMLName                            string `xml:"trt:GetVideoSourceConfigurationOptions"`
		GetVideoSourceConfigurationOptions GetVideoSourceConfigurationOptions
	}{
		GetVideoSourceConfigurationOptions: args,
	}

	resp := GetVideoSourceConfigurationOptionsResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptGetVideoSourceConfigurations(args GetVideoSourceConfigurations) (*GetVideoSourceConfigurationsResponse, error) {
	req := struct {
		XMLName                      string `xml:"trt:GetVideoSourceConfigurations"`
		GetVideoSourceConfigurations GetVideoSourceConfigurations
	}{
		GetVideoSourceConfigurations: args,
	}

	resp := GetVideoSourceConfigurationsResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptGetVideoSourceModes(args GetVideoSourceModes) (*GetVideoSourceModesResponse, error) {
	req := struct {
		XMLName             string `xml:"trt:GetVideoSourceModes"`
		GetVideoSourceModes GetVideoSourceModes
	}{
		GetVideoSourceModes: args,
	}

	resp := GetVideoSourceModesResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptGetVideoSources(args GetVideoSources) (*GetVideoSourcesResponse, error) {
	req := struct {
		XMLName         string `xml:"trt:GetVideoSources"`
		GetVideoSources GetVideoSources
	}{
		GetVideoSources: args,
	}

	resp := GetVideoSourcesResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptRemoveAudioDecoderConfiguration(args RemoveAudioDecoderConfiguration) (*RemoveAudioDecoderConfigurationResponse, error) {
	req := struct {
		XMLName                         string `xml:"trt:RemoveAudioDecoderConfiguration"`
		RemoveAudioDecoderConfiguration RemoveAudioDecoderConfiguration
	}{
		RemoveAudioDecoderConfiguration: args,
	}

	resp := RemoveAudioDecoderConfigurationResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptRemoveAudioEncoderConfiguration(args RemoveAudioEncoderConfiguration) (*RemoveAudioEncoderConfigurationResponse, error) {
	req := struct {
		XMLName                         string `xml:"trt:RemoveAudioEncoderConfiguration"`
		RemoveAudioEncoderConfiguration RemoveAudioEncoderConfiguration
	}{
		RemoveAudioEncoderConfiguration: args,
	}

	resp := RemoveAudioEncoderConfigurationResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptRemoveAudioOutputConfiguration(args RemoveAudioOutputConfiguration) (*RemoveAudioOutputConfigurationResponse, error) {
	req := struct {
		XMLName                        string `xml:"trt:RemoveAudioOutputConfiguration"`
		RemoveAudioOutputConfiguration RemoveAudioOutputConfiguration
	}{
		RemoveAudioOutputConfiguration: args,
	}

	resp := RemoveAudioOutputConfigurationResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptRemoveAudioSourceConfiguration(args RemoveAudioSourceConfiguration) (*RemoveAudioSourceConfigurationResponse, error) {
	req := struct {
		XMLName                        string `xml:"trt:RemoveAudioSourceConfiguration"`
		RemoveAudioSourceConfiguration RemoveAudioSourceConfiguration
	}{
		RemoveAudioSourceConfiguration: args,
	}

	resp := RemoveAudioSourceConfigurationResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptRemoveMetadataConfiguration(args RemoveMetadataConfiguration) (*RemoveMetadataConfigurationResponse, error) {
	req := struct {
		XMLName                     string `xml:"trt:RemoveMetadataConfiguration"`
		RemoveMetadataConfiguration RemoveMetadataConfiguration
	}{
		RemoveMetadataConfiguration: args,
	}

	resp := RemoveMetadataConfigurationResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptRemovePTZConfiguration(args RemovePTZConfiguration) (*RemovePTZConfigurationResponse, error) {
	req := struct {
		XMLName                string `xml:"trt:RemovePTZConfiguration"`
		RemovePTZConfiguration RemovePTZConfiguration
	}{
		RemovePTZConfiguration: args,
	}

	resp := RemovePTZConfigurationResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptRemoveVideoAnalyticsConfiguration(args RemoveVideoAnalyticsConfiguration) (*RemoveVideoAnalyticsConfigurationResponse, error) {
	req := struct {
		XMLName                           string `xml:"trt:RemoveVideoAnalyticsConfiguration"`
		RemoveVideoAnalyticsConfiguration RemoveVideoAnalyticsConfiguration
	}{
		RemoveVideoAnalyticsConfiguration: args,
	}

	resp := RemoveVideoAnalyticsConfigurationResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptRemoveVideoEncoderConfiguration(args RemoveVideoEncoderConfiguration) (*RemoveVideoEncoderConfigurationResponse, error) {
	req := struct {
		XMLName                         string `xml:"trt:RemoveVideoEncoderConfiguration"`
		RemoveVideoEncoderConfiguration RemoveVideoEncoderConfiguration
	}{
		RemoveVideoEncoderConfiguration: args,
	}

	resp := RemoveVideoEncoderConfigurationResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptRemoveVideoSourceConfiguration(args RemoveVideoSourceConfiguration) (*RemoveVideoSourceConfigurationResponse, error) {
	req := struct {
		XMLName                        string `xml:"trt:RemoveVideoSourceConfiguration"`
		RemoveVideoSourceConfiguration RemoveVideoSourceConfiguration
	}{
		RemoveVideoSourceConfiguration: args,
	}

	resp := RemoveVideoSourceConfigurationResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptSetAudioDecoderConfiguration(args SetAudioDecoderConfiguration) (*SetAudioDecoderConfigurationResponse, error) {
	req := struct {
		XMLName                      string `xml:"trt:SetAudioDecoderConfiguration"`
		SetAudioDecoderConfiguration SetAudioDecoderConfiguration
	}{
		SetAudioDecoderConfiguration: args,
	}

	resp := SetAudioDecoderConfigurationResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptSetAudioEncoderConfiguration(args SetAudioEncoderConfiguration) (*SetAudioEncoderConfigurationResponse, error) {
	req := struct {
		XMLName                      string `xml:"trt:SetAudioEncoderConfiguration"`
		SetAudioEncoderConfiguration SetAudioEncoderConfiguration
	}{
		SetAudioEncoderConfiguration: args,
	}

	resp := SetAudioEncoderConfigurationResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptSetAudioOutputConfiguration(args SetAudioOutputConfiguration) (*SetAudioOutputConfigurationResponse, error) {
	req := struct {
		XMLName                     string `xml:"trt:SetAudioOutputConfiguration"`
		SetAudioOutputConfiguration SetAudioOutputConfiguration
	}{
		SetAudioOutputConfiguration: args,
	}

	resp := SetAudioOutputConfigurationResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptSetAudioSourceConfiguration(args SetAudioSourceConfiguration) (*SetAudioSourceConfigurationResponse, error) {
	req := struct {
		XMLName                     string `xml:"trt:SetAudioSourceConfiguration"`
		SetAudioSourceConfiguration SetAudioSourceConfiguration
	}{
		SetAudioSourceConfiguration: args,
	}

	resp := SetAudioSourceConfigurationResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptSetMetadataConfiguration(args SetMetadataConfiguration) (*SetMetadataConfigurationResponse, error) {
	req := struct {
		XMLName                  string `xml:"trt:SetMetadataConfiguration"`
		SetMetadataConfiguration SetMetadataConfiguration
	}{
		SetMetadataConfiguration: args,
	}

	resp := SetMetadataConfigurationResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptSetOSD(args SetOSD) (*SetOSDResponse, error) {
	req := struct {
		XMLName string `xml:"trt:SetOSD"`
		SetOSD  SetOSD
	}{
		SetOSD: args,
	}

	resp := SetOSDResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptSetSynchronizationPoint(args SetSynchronizationPoint) (*SetSynchronizationPointResponse, error) {
	req := struct {
		XMLName                 string `xml:"trt:SetSynchronizationPoint"`
		SetSynchronizationPoint SetSynchronizationPoint
	}{
		SetSynchronizationPoint: args,
	}

	resp := SetSynchronizationPointResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptSetVideoAnalyticsConfiguration(args SetVideoAnalyticsConfiguration) (*SetVideoAnalyticsConfigurationResponse, error) {
	req := struct {
		XMLName                        string `xml:"trt:SetVideoAnalyticsConfiguration"`
		SetVideoAnalyticsConfiguration SetVideoAnalyticsConfiguration
	}{
		SetVideoAnalyticsConfiguration: args,
	}

	resp := SetVideoAnalyticsConfigurationResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptSetVideoEncoderConfiguration(args SetVideoEncoderConfiguration) (*SetVideoEncoderConfigurationResponse, error) {
	req := struct {
		XMLName                      string `xml:"trt:SetVideoEncoderConfiguration"`
		SetVideoEncoderConfiguration SetVideoEncoderConfiguration
	}{
		SetVideoEncoderConfiguration: args,
	}

	resp := SetVideoEncoderConfigurationResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptSetVideoSourceConfiguration(args SetVideoSourceConfiguration) (*SetVideoSourceConfigurationResponse, error) {
	req := struct {
		XMLName                     string `xml:"trt:SetVideoSourceConfiguration"`
		SetVideoSourceConfiguration SetVideoSourceConfiguration
	}{
		SetVideoSourceConfiguration: args,
	}

	resp := SetVideoSourceConfigurationResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptSetVideoSourceMode(args SetVideoSourceMode) (*SetVideoSourceModeResponse, error) {
	req := struct {
		XMLName            string `xml:"trt:SetVideoSourceMode"`
		SetVideoSourceMode SetVideoSourceMode
	}{
		SetVideoSourceMode: args,
	}

	resp := SetVideoSourceModeResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptStartMulticastStreaming(args StartMulticastStreaming) (*StartMulticastStreamingResponse, error) {
	req := struct {
		XMLName                 string `xml:"trt:StartMulticastStreaming"`
		StartMulticastStreaming StartMulticastStreaming
	}{
		StartMulticastStreaming: args,
	}

	resp := StartMulticastStreamingResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *media) OptStopMulticastStreaming(args StopMulticastStreaming) (*StopMulticastStreamingResponse, error) {
	req := struct {
		XMLName                string `xml:"trt:StopMulticastStreaming"`
		StopMulticastStreaming StopMulticastStreaming
	}{
		StopMulticastStreaming: args,
	}

	resp := StopMulticastStreamingResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
