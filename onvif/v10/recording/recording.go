package recording

import (
	"github.com/zijiemeng/go-onvif/onvif/common"
)

var Namespace = "http://www.onvif.org/ver10/recording/wsdl"

// NewRecordingPort creates an initializes a RecordingPort.
func NewRecordingPort(endpoint string, cli common.Client) RecordingPort {
	return &recordingPort{cli: cli, Endpoint: endpoint}
}

// RecordingPort was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type RecordingPort interface {
	OptCreateRecording(CreateRecording CreateRecording) (*CreateRecordingResponse, *common.Fault)

	OptCreateRecordingJob(CreateRecordingJob CreateRecordingJob) (*CreateRecordingJobResponse, *common.Fault)

	OptCreateTrack(CreateTrack CreateTrack) (*CreateTrackResponse, *common.Fault)

	OptDeleteRecording(DeleteRecording DeleteRecording) (*DeleteRecordingResponse, *common.Fault)

	OptDeleteRecordingJob(DeleteRecordingJob DeleteRecordingJob) (*DeleteRecordingJobResponse, *common.Fault)

	OptDeleteTrack(DeleteTrack DeleteTrack) (*DeleteTrackResponse, *common.Fault)

	OptExportRecordedData(ExportRecordedData ExportRecordedData) (*ExportRecordedDataResponse, *common.Fault)

	OptGetExportRecordedDataState(GetExportRecordedDataState GetExportRecordedDataState) (*GetExportRecordedDataStateResponse, *common.Fault)

	OptGetRecordingConfiguration(GetRecordingConfiguration GetRecordingConfiguration) (*GetRecordingConfigurationResponse, *common.Fault)

	OptGetRecordingJobConfiguration(GetRecordingJobConfiguration GetRecordingJobConfiguration) (*GetRecordingJobConfigurationResponse, *common.Fault)

	OptGetRecordingJobState(GetRecordingJobState GetRecordingJobState) (*GetRecordingJobStateResponse, *common.Fault)

	OptGetRecordingJobs(GetRecordingJobs GetRecordingJobs) (*GetRecordingJobsResponse, *common.Fault)

	OptGetRecordingOptions(GetRecordingOptions GetRecordingOptions) (*GetRecordingOptionsResponse, *common.Fault)

	OptGetRecordings(GetRecordings GetRecordings) (*GetRecordingsResponse, *common.Fault)

	OptGetServiceCapabilities(GetServiceCapabilities GetServiceCapabilities) (*GetServiceCapabilitiesResponse, *common.Fault)

	OptGetTrackConfiguration(GetTrackConfiguration GetTrackConfiguration) (*GetTrackConfigurationResponse, *common.Fault)

	OptSetRecordingConfiguration(SetRecordingConfiguration SetRecordingConfiguration) (*SetRecordingConfigurationResponse, *common.Fault)

	OptSetRecordingJobConfiguration(SetRecordingJobConfiguration SetRecordingJobConfiguration) (*SetRecordingJobConfigurationResponse, *common.Fault)

	OptSetRecordingJobMode(SetRecordingJobMode SetRecordingJobMode) (*SetRecordingJobModeResponse, *common.Fault)

	OptSetTrackConfiguration(SetTrackConfiguration SetTrackConfiguration) (*SetTrackConfigurationResponse, *common.Fault)

	OptStopExportRecordedData(StopExportRecordedData StopExportRecordedData) (*StopExportRecordedDataResponse, *common.Fault)
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

func (p *recordingPort) OptCreateRecording(args CreateRecording) (*CreateRecordingResponse, *common.Fault) {
	req := struct {
		XMLName         string `xml:"trc:CreateRecording"`
		CreateRecording CreateRecording
	}{
		CreateRecording: args,
	}

	resp := CreateRecordingResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *recordingPort) OptCreateRecordingJob(args CreateRecordingJob) (*CreateRecordingJobResponse, *common.Fault) {
	req := struct {
		XMLName            string `xml:"trc:CreateRecordingJob"`
		CreateRecordingJob CreateRecordingJob
	}{
		CreateRecordingJob: args,
	}

	resp := CreateRecordingJobResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *recordingPort) OptCreateTrack(args CreateTrack) (*CreateTrackResponse, *common.Fault) {
	req := struct {
		XMLName     string `xml:"trc:CreateTrack"`
		CreateTrack CreateTrack
	}{
		CreateTrack: args,
	}

	resp := CreateTrackResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *recordingPort) OptDeleteRecording(args DeleteRecording) (*DeleteRecordingResponse, *common.Fault) {
	req := struct {
		XMLName         string `xml:"trc:DeleteRecording"`
		DeleteRecording DeleteRecording
	}{
		DeleteRecording: args,
	}

	resp := DeleteRecordingResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *recordingPort) OptDeleteRecordingJob(args DeleteRecordingJob) (*DeleteRecordingJobResponse, *common.Fault) {
	req := struct {
		XMLName            string `xml:"trc:DeleteRecordingJob"`
		DeleteRecordingJob DeleteRecordingJob
	}{
		DeleteRecordingJob: args,
	}

	resp := DeleteRecordingJobResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *recordingPort) OptDeleteTrack(args DeleteTrack) (*DeleteTrackResponse, *common.Fault) {
	req := struct {
		XMLName     string `xml:"trc:DeleteTrack"`
		DeleteTrack DeleteTrack
	}{
		DeleteTrack: args,
	}

	resp := DeleteTrackResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *recordingPort) OptExportRecordedData(args ExportRecordedData) (*ExportRecordedDataResponse, *common.Fault) {
	req := struct {
		XMLName            string `xml:"trc:ExportRecordedData"`
		ExportRecordedData ExportRecordedData
	}{
		ExportRecordedData: args,
	}

	resp := ExportRecordedDataResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *recordingPort) OptGetExportRecordedDataState(args GetExportRecordedDataState) (*GetExportRecordedDataStateResponse, *common.Fault) {
	req := struct {
		XMLName                    string `xml:"trc:GetExportRecordedDataState"`
		GetExportRecordedDataState GetExportRecordedDataState
	}{
		GetExportRecordedDataState: args,
	}

	resp := GetExportRecordedDataStateResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *recordingPort) OptGetRecordingConfiguration(args GetRecordingConfiguration) (*GetRecordingConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                   string `xml:"trc:GetRecordingConfiguration"`
		GetRecordingConfiguration GetRecordingConfiguration
	}{
		GetRecordingConfiguration: args,
	}

	resp := GetRecordingConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *recordingPort) OptGetRecordingJobConfiguration(args GetRecordingJobConfiguration) (*GetRecordingJobConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                      string `xml:"trc:GetRecordingJobConfiguration"`
		GetRecordingJobConfiguration GetRecordingJobConfiguration
	}{
		GetRecordingJobConfiguration: args,
	}

	resp := GetRecordingJobConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *recordingPort) OptGetRecordingJobState(args GetRecordingJobState) (*GetRecordingJobStateResponse, *common.Fault) {
	req := struct {
		XMLName              string `xml:"trc:GetRecordingJobState"`
		GetRecordingJobState GetRecordingJobState
	}{
		GetRecordingJobState: args,
	}

	resp := GetRecordingJobStateResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *recordingPort) OptGetRecordingJobs(args GetRecordingJobs) (*GetRecordingJobsResponse, *common.Fault) {
	req := struct {
		XMLName          string `xml:"trc:GetRecordingJobs"`
		GetRecordingJobs GetRecordingJobs
	}{
		GetRecordingJobs: args,
	}

	resp := GetRecordingJobsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *recordingPort) OptGetRecordingOptions(args GetRecordingOptions) (*GetRecordingOptionsResponse, *common.Fault) {
	req := struct {
		XMLName             string `xml:"trc:GetRecordingOptions"`
		GetRecordingOptions GetRecordingOptions
	}{
		GetRecordingOptions: args,
	}

	resp := GetRecordingOptionsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *recordingPort) OptGetRecordings(args GetRecordings) (*GetRecordingsResponse, *common.Fault) {
	req := struct {
		XMLName       string `xml:"trc:GetRecordings"`
		GetRecordings GetRecordings
	}{
		GetRecordings: args,
	}

	resp := GetRecordingsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *recordingPort) OptGetServiceCapabilities(args GetServiceCapabilities) (*GetServiceCapabilitiesResponse, *common.Fault) {
	req := struct {
		XMLName                string `xml:"trc:GetServiceCapabilities"`
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

func (p *recordingPort) OptGetTrackConfiguration(args GetTrackConfiguration) (*GetTrackConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName               string `xml:"trc:GetTrackConfiguration"`
		GetTrackConfiguration GetTrackConfiguration
	}{
		GetTrackConfiguration: args,
	}

	resp := GetTrackConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *recordingPort) OptSetRecordingConfiguration(args SetRecordingConfiguration) (*SetRecordingConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                   string `xml:"trc:SetRecordingConfiguration"`
		SetRecordingConfiguration SetRecordingConfiguration
	}{
		SetRecordingConfiguration: args,
	}

	resp := SetRecordingConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *recordingPort) OptSetRecordingJobConfiguration(args SetRecordingJobConfiguration) (*SetRecordingJobConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                      string `xml:"trc:SetRecordingJobConfiguration"`
		SetRecordingJobConfiguration SetRecordingJobConfiguration
	}{
		SetRecordingJobConfiguration: args,
	}

	resp := SetRecordingJobConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *recordingPort) OptSetRecordingJobMode(args SetRecordingJobMode) (*SetRecordingJobModeResponse, *common.Fault) {
	req := struct {
		XMLName             string `xml:"trc:SetRecordingJobMode"`
		SetRecordingJobMode SetRecordingJobMode
	}{
		SetRecordingJobMode: args,
	}

	resp := SetRecordingJobModeResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *recordingPort) OptSetTrackConfiguration(args SetTrackConfiguration) (*SetTrackConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName               string `xml:"trc:SetTrackConfiguration"`
		SetTrackConfiguration SetTrackConfiguration
	}{
		SetTrackConfiguration: args,
	}

	resp := SetTrackConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *recordingPort) OptStopExportRecordedData(args StopExportRecordedData) (*StopExportRecordedDataResponse, *common.Fault) {
	req := struct {
		XMLName                string `xml:"trc:StopExportRecordedData"`
		StopExportRecordedData StopExportRecordedData
	}{
		StopExportRecordedData: args,
	}

	resp := StopExportRecordedDataResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}
