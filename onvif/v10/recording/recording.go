package recording

import (
	"code.byted.org/videoarch/go-onvif/onvif/common"
)

var Namespace = "http://www.onvif.org/ver10/recording/wsdl"

// NewRecordingPort creates an initializes a RecordingPort.
func NewRecordingPort(endpoint string, cli common.Client) RecordingPort {
	return &recordingPort{cli: cli, Endpoint: endpoint}
}

// RecordingPort was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type RecordingPort interface {
	OptCreateRecording(CreateRecording CreateRecording) (*CreateRecordingResponse, error)

	OptCreateRecordingJob(CreateRecordingJob CreateRecordingJob) (*CreateRecordingJobResponse, error)

	OptCreateTrack(CreateTrack CreateTrack) (*CreateTrackResponse, error)

	OptDeleteRecording(DeleteRecording DeleteRecording) (*DeleteRecordingResponse, error)

	OptDeleteRecordingJob(DeleteRecordingJob DeleteRecordingJob) (*DeleteRecordingJobResponse, error)

	OptDeleteTrack(DeleteTrack DeleteTrack) (*DeleteTrackResponse, error)

	OptExportRecordedData(ExportRecordedData ExportRecordedData) (*ExportRecordedDataResponse, error)

	OptGetExportRecordedDataState(GetExportRecordedDataState GetExportRecordedDataState) (*GetExportRecordedDataStateResponse, error)

	OptGetRecordingConfiguration(GetRecordingConfiguration GetRecordingConfiguration) (*GetRecordingConfigurationResponse, error)

	OptGetRecordingJobConfiguration(GetRecordingJobConfiguration GetRecordingJobConfiguration) (*GetRecordingJobConfigurationResponse, error)

	OptGetRecordingJobState(GetRecordingJobState GetRecordingJobState) (*GetRecordingJobStateResponse, error)

	OptGetRecordingJobs(GetRecordingJobs GetRecordingJobs) (*GetRecordingJobsResponse, error)

	OptGetRecordingOptions(GetRecordingOptions GetRecordingOptions) (*GetRecordingOptionsResponse, error)

	OptGetRecordings(GetRecordings GetRecordings) (*GetRecordingsResponse, error)

	OptGetServiceCapabilities(GetServiceCapabilities GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error)

	OptGetTrackConfiguration(GetTrackConfiguration GetTrackConfiguration) (*GetTrackConfigurationResponse, error)

	OptSetRecordingConfiguration(SetRecordingConfiguration SetRecordingConfiguration) (*SetRecordingConfigurationResponse, error)

	OptSetRecordingJobConfiguration(SetRecordingJobConfiguration SetRecordingJobConfiguration) (*SetRecordingJobConfigurationResponse, error)

	OptSetRecordingJobMode(SetRecordingJobMode SetRecordingJobMode) (*SetRecordingJobModeResponse, error)

	OptSetTrackConfiguration(SetTrackConfiguration SetTrackConfiguration) (*SetTrackConfigurationResponse, error)

	OptStopExportRecordedData(StopExportRecordedData StopExportRecordedData) (*StopExportRecordedDataResponse, error)
}
type DateTime string

type Capabilities []interface{}

type CreateRecording struct {
	RecordingConfiguration *common.RecordingConfiguration `xml:"RecordingConfiguration,omitempty" json:"RecordingConfiguration,omitempty" yaml:"RecordingConfiguration,omitempty"`
}

type CreateRecordingJob struct {
	JobConfiguration *common.RecordingJobConfiguration `xml:"JobConfiguration,omitempty" json:"JobConfiguration,omitempty" yaml:"JobConfiguration,omitempty"`
}

type CreateRecordingJobResponse struct {
	JobToken         *common.RecordingJobReference     `xml:"JobToken,omitempty" json:"JobToken,omitempty" yaml:"JobToken,omitempty"`
	JobConfiguration *common.RecordingJobConfiguration `xml:"JobConfiguration,omitempty" json:"JobConfiguration,omitempty" yaml:"JobConfiguration,omitempty"`
}

type CreateRecordingResponse struct {
	RecordingToken *common.RecordingReference `xml:"RecordingToken,omitempty" json:"RecordingToken,omitempty" yaml:"RecordingToken,omitempty"`
}

type CreateTrack struct {
	RecordingToken     *common.RecordingReference `xml:"RecordingToken,omitempty" json:"RecordingToken,omitempty" yaml:"RecordingToken,omitempty"`
	TrackConfiguration *common.TrackConfiguration `xml:"TrackConfiguration,omitempty" json:"TrackConfiguration,omitempty" yaml:"TrackConfiguration,omitempty"`
}

type CreateTrackResponse struct {
	TrackToken *common.TrackReference `xml:"TrackToken,omitempty" json:"TrackToken,omitempty" yaml:"TrackToken,omitempty"`
}

type DeleteRecording struct {
	RecordingToken *common.RecordingReference `xml:"RecordingToken,omitempty" json:"RecordingToken,omitempty" yaml:"RecordingToken,omitempty"`
}

type DeleteRecordingJob struct {
	JobToken *common.RecordingJobReference `xml:"JobToken,omitempty" json:"JobToken,omitempty" yaml:"JobToken,omitempty"`
}

type DeleteRecordingJobResponse struct {
}

type DeleteRecordingResponse struct {
}

type DeleteTrack struct {
	RecordingToken *common.RecordingReference `xml:"RecordingToken,omitempty" json:"RecordingToken,omitempty" yaml:"RecordingToken,omitempty"`
	TrackToken     *common.TrackReference     `xml:"TrackToken,omitempty" json:"TrackToken,omitempty" yaml:"TrackToken,omitempty"`
}

type DeleteTrackResponse struct {
}

type ExportRecordedData struct {
	StartPoint         *common.DateTime             `xml:"StartPoint,omitempty" json:"StartPoint,omitempty" yaml:"StartPoint,omitempty"`
	EndPoint           *common.DateTime             `xml:"EndPoint,omitempty" json:"EndPoint,omitempty" yaml:"EndPoint,omitempty"`
	SearchScope        *common.SearchScope          `xml:"SearchScope,omitempty" json:"SearchScope,omitempty" yaml:"SearchScope,omitempty"`
	FileFormat         string                       `xml:"FileFormat,omitempty" json:"FileFormat,omitempty" yaml:"FileFormat,omitempty"`
	StorageDestination *common.StorageReferencePath `xml:"StorageDestination,omitempty" json:"StorageDestination,omitempty" yaml:"StorageDestination,omitempty"`
}

type ExportRecordedDataResponse struct {
	OperationToken *common.ReferenceToken `xml:"OperationToken,omitempty" json:"OperationToken,omitempty" yaml:"OperationToken,omitempty"`
	FileNames      []string               `xml:"FileNames,omitempty" json:"FileNames,omitempty" yaml:"FileNames,omitempty"`
	Extension      string                 `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type GetExportRecordedDataState struct {
	OperationToken *common.ReferenceToken `xml:"OperationToken,omitempty" json:"OperationToken,omitempty" yaml:"OperationToken,omitempty"`
}

type GetExportRecordedDataStateResponse struct {
	Progress           float64                     `xml:"Progress,omitempty" json:"Progress,omitempty" yaml:"Progress,omitempty"`
	FileProgressStatus *common.ArrayOfFileProgress `xml:"FileProgressStatus,omitempty" json:"FileProgressStatus,omitempty" yaml:"FileProgressStatus,omitempty"`
}

type GetRecordingConfiguration struct {
	RecordingToken *common.RecordingReference `xml:"RecordingToken,omitempty" json:"RecordingToken,omitempty" yaml:"RecordingToken,omitempty"`
}

type GetRecordingConfigurationResponse struct {
	RecordingConfiguration *common.RecordingConfiguration `xml:"RecordingConfiguration,omitempty" json:"RecordingConfiguration,omitempty" yaml:"RecordingConfiguration,omitempty"`
}

type GetRecordingJobConfiguration struct {
	JobToken *common.RecordingJobReference `xml:"JobToken,omitempty" json:"JobToken,omitempty" yaml:"JobToken,omitempty"`
}

type GetRecordingJobConfigurationResponse struct {
	JobConfiguration *common.RecordingJobConfiguration `xml:"JobConfiguration,omitempty" json:"JobConfiguration,omitempty" yaml:"JobConfiguration,omitempty"`
}

type GetRecordingJobState struct {
	JobToken *common.RecordingJobReference `xml:"JobToken,omitempty" json:"JobToken,omitempty" yaml:"JobToken,omitempty"`
}

type GetRecordingJobStateResponse struct {
	State *common.RecordingJobStateInformation `xml:"State,omitempty" json:"State,omitempty" yaml:"State,omitempty"`
}

type GetRecordingJobs struct {
}

type GetRecordingJobsResponse struct {
	JobItem []*common.GetRecordingJobsResponseItem `xml:"JobItem,omitempty" json:"JobItem,omitempty" yaml:"JobItem,omitempty"`
}

type GetRecordingOptions struct {
	RecordingToken *common.RecordingReference `xml:"RecordingToken,omitempty" json:"RecordingToken,omitempty" yaml:"RecordingToken,omitempty"`
}

type GetRecordingOptionsResponse struct {
	Options RecordingOptions `xml:"Options,omitempty" json:"Options,omitempty" yaml:"Options,omitempty"`
}

type GetRecordings struct {
}

type GetRecordingsResponse struct {
	RecordingItem []*common.GetRecordingsResponseItem `xml:"RecordingItem,omitempty" json:"RecordingItem,omitempty" yaml:"RecordingItem,omitempty"`
}

type GetServiceCapabilities struct {
}

type GetServiceCapabilitiesResponse struct {
	Capabilities Capabilities `xml:"Capabilities,omitempty" json:"Capabilities,omitempty" yaml:"Capabilities,omitempty"`
}

type GetTrackConfiguration struct {
	RecordingToken *common.RecordingReference `xml:"RecordingToken,omitempty" json:"RecordingToken,omitempty" yaml:"RecordingToken,omitempty"`
	TrackToken     *common.TrackReference     `xml:"TrackToken,omitempty" json:"TrackToken,omitempty" yaml:"TrackToken,omitempty"`
}

type GetTrackConfigurationResponse struct {
	TrackConfiguration *common.TrackConfiguration `xml:"TrackConfiguration,omitempty" json:"TrackConfiguration,omitempty" yaml:"TrackConfiguration,omitempty"`
}

type JobOptions struct {
	Spare             int                   `xml:"Spare,attr,omitempty" json:"Spare,attr,omitempty" yaml:"Spare,attr,omitempty"`
	CompatibleSources common.StringAttrList `xml:"CompatibleSources,attr,omitempty" json:"CompatibleSources,attr,omitempty" yaml:"CompatibleSources,attr,omitempty"`
}

type RecordingOptions struct {
	Job   JobOptions   `xml:"Job,omitempty" json:"Job,omitempty" yaml:"Job,omitempty"`
	Track TrackOptions `xml:"Track,omitempty" json:"Track,omitempty" yaml:"Track,omitempty"`
}

type SetRecordingConfiguration struct {
	RecordingToken         *common.RecordingReference     `xml:"RecordingToken,omitempty" json:"RecordingToken,omitempty" yaml:"RecordingToken,omitempty"`
	RecordingConfiguration *common.RecordingConfiguration `xml:"RecordingConfiguration,omitempty" json:"RecordingConfiguration,omitempty" yaml:"RecordingConfiguration,omitempty"`
}

type SetRecordingConfigurationResponse struct {
}

type SetRecordingJobConfiguration struct {
	JobToken         *common.RecordingJobReference     `xml:"JobToken,omitempty" json:"JobToken,omitempty" yaml:"JobToken,omitempty"`
	JobConfiguration *common.RecordingJobConfiguration `xml:"JobConfiguration,omitempty" json:"JobConfiguration,omitempty" yaml:"JobConfiguration,omitempty"`
}

type SetRecordingJobConfigurationResponse struct {
	JobConfiguration *common.RecordingJobConfiguration `xml:"JobConfiguration,omitempty" json:"JobConfiguration,omitempty" yaml:"JobConfiguration,omitempty"`
}

type SetRecordingJobMode struct {
	JobToken *common.RecordingJobReference `xml:"JobToken,omitempty" json:"JobToken,omitempty" yaml:"JobToken,omitempty"`
	Mode     *common.RecordingJobMode      `xml:"Mode,omitempty" json:"Mode,omitempty" yaml:"Mode,omitempty"`
}

type SetRecordingJobModeResponse struct {
}

type SetTrackConfiguration struct {
	RecordingToken     *common.RecordingReference `xml:"RecordingToken,omitempty" json:"RecordingToken,omitempty" yaml:"RecordingToken,omitempty"`
	TrackToken         *common.TrackReference     `xml:"TrackToken,omitempty" json:"TrackToken,omitempty" yaml:"TrackToken,omitempty"`
	TrackConfiguration *common.TrackConfiguration `xml:"TrackConfiguration,omitempty" json:"TrackConfiguration,omitempty" yaml:"TrackConfiguration,omitempty"`
}

type SetTrackConfigurationResponse struct {
}

type StopExportRecordedData struct {
	OperationToken *common.ReferenceToken `xml:"OperationToken,omitempty" json:"OperationToken,omitempty" yaml:"OperationToken,omitempty"`
}

type StopExportRecordedDataResponse struct {
	Progress           float64                     `xml:"Progress,omitempty" json:"Progress,omitempty" yaml:"Progress,omitempty"`
	FileProgressStatus *common.ArrayOfFileProgress `xml:"FileProgressStatus,omitempty" json:"FileProgressStatus,omitempty" yaml:"FileProgressStatus,omitempty"`
}

type TrackOptions struct {
	SpareTotal    int `xml:"SpareTotal,attr,omitempty" json:"SpareTotal,attr,omitempty" yaml:"SpareTotal,attr,omitempty"`
	SpareVideo    int `xml:"SpareVideo,attr,omitempty" json:"SpareVideo,attr,omitempty" yaml:"SpareVideo,attr,omitempty"`
	SpareAudio    int `xml:"SpareAudio,attr,omitempty" json:"SpareAudio,attr,omitempty" yaml:"SpareAudio,attr,omitempty"`
	SpareMetadata int `xml:"SpareMetadata,attr,omitempty" json:"SpareMetadata,attr,omitempty" yaml:"SpareMetadata,attr,omitempty"`
}

// recordingPort implements the RecordingPort interface.
type recordingPort struct {
	cli      common.Client
	Endpoint string
}

func (p *recordingPort) OptCreateRecording(args CreateRecording) (*CreateRecordingResponse, error) {
	req := struct {
		XMLName         string `xml:"trc:CreateRecording"`
		CreateRecording CreateRecording
	}{
		CreateRecording: args,
	}

	resp := CreateRecordingResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *recordingPort) OptCreateRecordingJob(args CreateRecordingJob) (*CreateRecordingJobResponse, error) {
	req := struct {
		XMLName            string `xml:"trc:CreateRecordingJob"`
		CreateRecordingJob CreateRecordingJob
	}{
		CreateRecordingJob: args,
	}

	resp := CreateRecordingJobResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *recordingPort) OptCreateTrack(args CreateTrack) (*CreateTrackResponse, error) {
	req := struct {
		XMLName     string `xml:"trc:CreateTrack"`
		CreateTrack CreateTrack
	}{
		CreateTrack: args,
	}

	resp := CreateTrackResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *recordingPort) OptDeleteRecording(args DeleteRecording) (*DeleteRecordingResponse, error) {
	req := struct {
		XMLName         string `xml:"trc:DeleteRecording"`
		DeleteRecording DeleteRecording
	}{
		DeleteRecording: args,
	}

	resp := DeleteRecordingResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *recordingPort) OptDeleteRecordingJob(args DeleteRecordingJob) (*DeleteRecordingJobResponse, error) {
	req := struct {
		XMLName            string `xml:"trc:DeleteRecordingJob"`
		DeleteRecordingJob DeleteRecordingJob
	}{
		DeleteRecordingJob: args,
	}

	resp := DeleteRecordingJobResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *recordingPort) OptDeleteTrack(args DeleteTrack) (*DeleteTrackResponse, error) {
	req := struct {
		XMLName     string `xml:"trc:DeleteTrack"`
		DeleteTrack DeleteTrack
	}{
		DeleteTrack: args,
	}

	resp := DeleteTrackResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *recordingPort) OptExportRecordedData(args ExportRecordedData) (*ExportRecordedDataResponse, error) {
	req := struct {
		XMLName            string `xml:"trc:ExportRecordedData"`
		ExportRecordedData ExportRecordedData
	}{
		ExportRecordedData: args,
	}

	resp := ExportRecordedDataResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *recordingPort) OptGetExportRecordedDataState(args GetExportRecordedDataState) (*GetExportRecordedDataStateResponse, error) {
	req := struct {
		XMLName                    string `xml:"trc:GetExportRecordedDataState"`
		GetExportRecordedDataState GetExportRecordedDataState
	}{
		GetExportRecordedDataState: args,
	}

	resp := GetExportRecordedDataStateResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *recordingPort) OptGetRecordingConfiguration(args GetRecordingConfiguration) (*GetRecordingConfigurationResponse, error) {
	req := struct {
		XMLName                   string `xml:"trc:GetRecordingConfiguration"`
		GetRecordingConfiguration GetRecordingConfiguration
	}{
		GetRecordingConfiguration: args,
	}

	resp := GetRecordingConfigurationResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *recordingPort) OptGetRecordingJobConfiguration(args GetRecordingJobConfiguration) (*GetRecordingJobConfigurationResponse, error) {
	req := struct {
		XMLName                      string `xml:"trc:GetRecordingJobConfiguration"`
		GetRecordingJobConfiguration GetRecordingJobConfiguration
	}{
		GetRecordingJobConfiguration: args,
	}

	resp := GetRecordingJobConfigurationResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *recordingPort) OptGetRecordingJobState(args GetRecordingJobState) (*GetRecordingJobStateResponse, error) {
	req := struct {
		XMLName              string `xml:"trc:GetRecordingJobState"`
		GetRecordingJobState GetRecordingJobState
	}{
		GetRecordingJobState: args,
	}

	resp := GetRecordingJobStateResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *recordingPort) OptGetRecordingJobs(args GetRecordingJobs) (*GetRecordingJobsResponse, error) {
	req := struct {
		XMLName          string `xml:"trc:GetRecordingJobs"`
		GetRecordingJobs GetRecordingJobs
	}{
		GetRecordingJobs: args,
	}

	resp := GetRecordingJobsResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *recordingPort) OptGetRecordingOptions(args GetRecordingOptions) (*GetRecordingOptionsResponse, error) {
	req := struct {
		XMLName             string `xml:"trc:GetRecordingOptions"`
		GetRecordingOptions GetRecordingOptions
	}{
		GetRecordingOptions: args,
	}

	resp := GetRecordingOptionsResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *recordingPort) OptGetRecordings(args GetRecordings) (*GetRecordingsResponse, error) {
	req := struct {
		XMLName       string `xml:"trc:GetRecordings"`
		GetRecordings GetRecordings
	}{
		GetRecordings: args,
	}

	resp := GetRecordingsResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *recordingPort) OptGetServiceCapabilities(args GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error) {
	req := struct {
		XMLName                string `xml:"trc:GetServiceCapabilities"`
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

func (p *recordingPort) OptGetTrackConfiguration(args GetTrackConfiguration) (*GetTrackConfigurationResponse, error) {
	req := struct {
		XMLName               string `xml:"trc:GetTrackConfiguration"`
		GetTrackConfiguration GetTrackConfiguration
	}{
		GetTrackConfiguration: args,
	}

	resp := GetTrackConfigurationResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *recordingPort) OptSetRecordingConfiguration(args SetRecordingConfiguration) (*SetRecordingConfigurationResponse, error) {
	req := struct {
		XMLName                   string `xml:"trc:SetRecordingConfiguration"`
		SetRecordingConfiguration SetRecordingConfiguration
	}{
		SetRecordingConfiguration: args,
	}

	resp := SetRecordingConfigurationResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *recordingPort) OptSetRecordingJobConfiguration(args SetRecordingJobConfiguration) (*SetRecordingJobConfigurationResponse, error) {
	req := struct {
		XMLName                      string `xml:"trc:SetRecordingJobConfiguration"`
		SetRecordingJobConfiguration SetRecordingJobConfiguration
	}{
		SetRecordingJobConfiguration: args,
	}

	resp := SetRecordingJobConfigurationResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *recordingPort) OptSetRecordingJobMode(args SetRecordingJobMode) (*SetRecordingJobModeResponse, error) {
	req := struct {
		XMLName             string `xml:"trc:SetRecordingJobMode"`
		SetRecordingJobMode SetRecordingJobMode
	}{
		SetRecordingJobMode: args,
	}

	resp := SetRecordingJobModeResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *recordingPort) OptSetTrackConfiguration(args SetTrackConfiguration) (*SetTrackConfigurationResponse, error) {
	req := struct {
		XMLName               string `xml:"trc:SetTrackConfiguration"`
		SetTrackConfiguration SetTrackConfiguration
	}{
		SetTrackConfiguration: args,
	}

	resp := SetTrackConfigurationResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *recordingPort) OptStopExportRecordedData(args StopExportRecordedData) (*StopExportRecordedDataResponse, error) {
	req := struct {
		XMLName                string `xml:"trc:StopExportRecordedData"`
		StopExportRecordedData StopExportRecordedData
	}{
		StopExportRecordedData: args,
	}

	resp := StopExportRecordedDataResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
