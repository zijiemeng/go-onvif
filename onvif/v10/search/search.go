package search

import (
	"code.byted.org/videoarch/go-onvif/onvif/common"
)

var Namespace = "http://www.onvif.org/ver10/search/wsdl"

// NewSearchPort creates an initializes a SearchPort.
func NewSearchPort(endpoint string, cli common.Client) SearchPort {
	return &searchPort{cli: cli, Endpoint: endpoint}
}

// SearchPort was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type SearchPort interface {
	OptEndSearch(EndSearch EndSearch) (*EndSearchResponse, error)

	OptFindEvents(FindEvents FindEvents) (*FindEventsResponse, error)

	OptFindMetadata(FindMetadata FindMetadata) (*FindMetadataResponse, error)

	OptFindPTZPosition(FindPTZPosition FindPTZPosition) (*FindPTZPositionResponse, error)

	OptFindRecordings(FindRecordings FindRecordings) (*FindRecordingsResponse, error)

	OptGetEventSearchResults(GetEventSearchResults GetEventSearchResults) (*GetEventSearchResultsResponse, error)

	OptGetMediaAttributes(GetMediaAttributes GetMediaAttributes) (*GetMediaAttributesResponse, error)

	OptGetMetadataSearchResults(GetMetadataSearchResults GetMetadataSearchResults) (*GetMetadataSearchResultsResponse, error)

	OptGetPTZPositionSearchResults(GetPTZPositionSearchResults GetPTZPositionSearchResults) (*GetPTZPositionSearchResultsResponse, error)

	OptGetRecordingInformation(GetRecordingInformation GetRecordingInformation) (*GetRecordingInformationResponse, error)

	OptGetRecordingSearchResults(GetRecordingSearchResults GetRecordingSearchResults) (*GetRecordingSearchResultsResponse, error)

	OptGetRecordingSummary(GetRecordingSummary GetRecordingSummary) (*GetRecordingSummaryResponse, error)

	OptGetSearchState(GetSearchState GetSearchState) (*GetSearchStateResponse, error)

	OptGetServiceCapabilities(GetServiceCapabilities GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error)
}
type DateTime string

type Duration string

type Capabilities []interface{}

type EndSearch struct {
	SearchToken *common.JobToken `xml:"SearchToken,omitempty" json:"SearchToken,omitempty" yaml:"SearchToken,omitempty"`
}

type EndSearchResponse struct {
	Endpoint *common.DateTime `xml:"Endpoint,omitempty" json:"Endpoint,omitempty" yaml:"Endpoint,omitempty"`
}

type FindEvents struct {
	StartPoint        *common.DateTime    `xml:"StartPoint,omitempty" json:"StartPoint,omitempty" yaml:"StartPoint,omitempty"`
	EndPoint          *common.DateTime    `xml:"EndPoint,omitempty" json:"EndPoint,omitempty" yaml:"EndPoint,omitempty"`
	Scope             *common.SearchScope `xml:"Scope,omitempty" json:"Scope,omitempty" yaml:"Scope,omitempty"`
	SearchFilter      *common.EventFilter `xml:"SearchFilter,omitempty" json:"SearchFilter,omitempty" yaml:"SearchFilter,omitempty"`
	IncludeStartState bool                `xml:"IncludeStartState,omitempty" json:"IncludeStartState,omitempty" yaml:"IncludeStartState,omitempty"`
	MaxMatches        int                 `xml:"MaxMatches,omitempty" json:"MaxMatches,omitempty" yaml:"MaxMatches,omitempty"`
	KeepAliveTime     *common.Duration    `xml:"KeepAliveTime,omitempty" json:"KeepAliveTime,omitempty" yaml:"KeepAliveTime,omitempty"`
}

type FindEventsResponse struct {
	SearchToken *common.JobToken `xml:"SearchToken,omitempty" json:"SearchToken,omitempty" yaml:"SearchToken,omitempty"`
}

type FindMetadata struct {
	StartPoint     *common.DateTime       `xml:"StartPoint,omitempty" json:"StartPoint,omitempty" yaml:"StartPoint,omitempty"`
	EndPoint       *common.DateTime       `xml:"EndPoint,omitempty" json:"EndPoint,omitempty" yaml:"EndPoint,omitempty"`
	Scope          *common.SearchScope    `xml:"Scope,omitempty" json:"Scope,omitempty" yaml:"Scope,omitempty"`
	MetadataFilter *common.MetadataFilter `xml:"MetadataFilter,omitempty" json:"MetadataFilter,omitempty" yaml:"MetadataFilter,omitempty"`
	MaxMatches     int                    `xml:"MaxMatches,omitempty" json:"MaxMatches,omitempty" yaml:"MaxMatches,omitempty"`
	KeepAliveTime  *common.Duration       `xml:"KeepAliveTime,omitempty" json:"KeepAliveTime,omitempty" yaml:"KeepAliveTime,omitempty"`
}

type FindMetadataResponse struct {
	SearchToken *common.JobToken `xml:"SearchToken,omitempty" json:"SearchToken,omitempty" yaml:"SearchToken,omitempty"`
}

type FindPTZPosition struct {
	StartPoint    *common.DateTime          `xml:"StartPoint,omitempty" json:"StartPoint,omitempty" yaml:"StartPoint,omitempty"`
	EndPoint      *common.DateTime          `xml:"EndPoint,omitempty" json:"EndPoint,omitempty" yaml:"EndPoint,omitempty"`
	Scope         *common.SearchScope       `xml:"Scope,omitempty" json:"Scope,omitempty" yaml:"Scope,omitempty"`
	SearchFilter  *common.PTZPositionFilter `xml:"SearchFilter,omitempty" json:"SearchFilter,omitempty" yaml:"SearchFilter,omitempty"`
	MaxMatches    int                       `xml:"MaxMatches,omitempty" json:"MaxMatches,omitempty" yaml:"MaxMatches,omitempty"`
	KeepAliveTime *common.Duration          `xml:"KeepAliveTime,omitempty" json:"KeepAliveTime,omitempty" yaml:"KeepAliveTime,omitempty"`
}

type FindPTZPositionResponse struct {
	SearchToken *common.JobToken `xml:"SearchToken,omitempty" json:"SearchToken,omitempty" yaml:"SearchToken,omitempty"`
}

type FindRecordings struct {
	Scope         *common.SearchScope `xml:"Scope,omitempty" json:"Scope,omitempty" yaml:"Scope,omitempty"`
	MaxMatches    int                 `xml:"MaxMatches,omitempty" json:"MaxMatches,omitempty" yaml:"MaxMatches,omitempty"`
	KeepAliveTime *common.Duration    `xml:"KeepAliveTime,omitempty" json:"KeepAliveTime,omitempty" yaml:"KeepAliveTime,omitempty"`
}

type FindRecordingsResponse struct {
	SearchToken *common.JobToken `xml:"SearchToken,omitempty" json:"SearchToken,omitempty" yaml:"SearchToken,omitempty"`
}

type GetEventSearchResults struct {
	SearchToken *common.JobToken `xml:"SearchToken,omitempty" json:"SearchToken,omitempty" yaml:"SearchToken,omitempty"`
	MinResults  int              `xml:"MinResults,omitempty" json:"MinResults,omitempty" yaml:"MinResults,omitempty"`
	MaxResults  int              `xml:"MaxResults,omitempty" json:"MaxResults,omitempty" yaml:"MaxResults,omitempty"`
	WaitTime    *common.Duration `xml:"WaitTime,omitempty" json:"WaitTime,omitempty" yaml:"WaitTime,omitempty"`
}

type GetEventSearchResultsResponse struct {
	ResultList *common.FindEventResultList `xml:"ResultList,omitempty" json:"ResultList,omitempty" yaml:"ResultList,omitempty"`
}

type GetMediaAttributes struct {
	RecordingTokens []*common.RecordingReference `xml:"RecordingTokens,omitempty" json:"RecordingTokens,omitempty" yaml:"RecordingTokens,omitempty"`
	Time            *common.DateTime             `xml:"Time,omitempty" json:"Time,omitempty" yaml:"Time,omitempty"`
}

type GetMediaAttributesResponse struct {
	MediaAttributes []*common.MediaAttributes `xml:"MediaAttributes,omitempty" json:"MediaAttributes,omitempty" yaml:"MediaAttributes,omitempty"`
}

type GetMetadataSearchResults struct {
	SearchToken *common.JobToken `xml:"SearchToken,omitempty" json:"SearchToken,omitempty" yaml:"SearchToken,omitempty"`
	MinResults  int              `xml:"MinResults,omitempty" json:"MinResults,omitempty" yaml:"MinResults,omitempty"`
	MaxResults  int              `xml:"MaxResults,omitempty" json:"MaxResults,omitempty" yaml:"MaxResults,omitempty"`
	WaitTime    *common.Duration `xml:"WaitTime,omitempty" json:"WaitTime,omitempty" yaml:"WaitTime,omitempty"`
}

type GetMetadataSearchResultsResponse struct {
	ResultList *common.FindMetadataResultList `xml:"ResultList,omitempty" json:"ResultList,omitempty" yaml:"ResultList,omitempty"`
}

type GetPTZPositionSearchResults struct {
	SearchToken *common.JobToken `xml:"SearchToken,omitempty" json:"SearchToken,omitempty" yaml:"SearchToken,omitempty"`
	MinResults  int              `xml:"MinResults,omitempty" json:"MinResults,omitempty" yaml:"MinResults,omitempty"`
	MaxResults  int              `xml:"MaxResults,omitempty" json:"MaxResults,omitempty" yaml:"MaxResults,omitempty"`
	WaitTime    *common.Duration `xml:"WaitTime,omitempty" json:"WaitTime,omitempty" yaml:"WaitTime,omitempty"`
}

type GetPTZPositionSearchResultsResponse struct {
	ResultList *common.FindPTZPositionResultList `xml:"ResultList,omitempty" json:"ResultList,omitempty" yaml:"ResultList,omitempty"`
}

type GetRecordingInformation struct {
	RecordingToken *common.RecordingReference `xml:"RecordingToken,omitempty" json:"RecordingToken,omitempty" yaml:"RecordingToken,omitempty"`
}

type GetRecordingInformationResponse struct {
	RecordingInformation *common.RecordingInformation `xml:"RecordingInformation,omitempty" json:"RecordingInformation,omitempty" yaml:"RecordingInformation,omitempty"`
}

type GetRecordingSearchResults struct {
	SearchToken *common.JobToken `xml:"SearchToken,omitempty" json:"SearchToken,omitempty" yaml:"SearchToken,omitempty"`
	MinResults  int              `xml:"MinResults,omitempty" json:"MinResults,omitempty" yaml:"MinResults,omitempty"`
	MaxResults  int              `xml:"MaxResults,omitempty" json:"MaxResults,omitempty" yaml:"MaxResults,omitempty"`
	WaitTime    *common.Duration `xml:"WaitTime,omitempty" json:"WaitTime,omitempty" yaml:"WaitTime,omitempty"`
}

type GetRecordingSearchResultsResponse struct {
	ResultList *common.FindRecordingResultList `xml:"ResultList,omitempty" json:"ResultList,omitempty" yaml:"ResultList,omitempty"`
}

type GetRecordingSummary struct {
}

type GetRecordingSummaryResponse struct {
	Summary *common.RecordingSummary `xml:"Summary,omitempty" json:"Summary,omitempty" yaml:"Summary,omitempty"`
}

type GetSearchState struct {
	SearchToken *common.JobToken `xml:"SearchToken,omitempty" json:"SearchToken,omitempty" yaml:"SearchToken,omitempty"`
}

type GetSearchStateResponse struct {
	State *common.SearchState `xml:"State,omitempty" json:"State,omitempty" yaml:"State,omitempty"`
}

type GetServiceCapabilities struct {
}

type GetServiceCapabilitiesResponse struct {
	Capabilities Capabilities `xml:"Capabilities,omitempty" json:"Capabilities,omitempty" yaml:"Capabilities,omitempty"`
}

// searchPort implements the SearchPort interface.
type searchPort struct {
	cli      common.Client
	Endpoint string
}

func (p *searchPort) OptEndSearch(args EndSearch) (*EndSearchResponse, error) {
	req := struct {
		XMLName   string `xml:"tse:EndSearch"`
		EndSearch EndSearch
	}{
		EndSearch: args,
	}

	resp := EndSearchResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *searchPort) OptFindEvents(args FindEvents) (*FindEventsResponse, error) {
	req := struct {
		XMLName    string `xml:"tse:FindEvents"`
		FindEvents FindEvents
	}{
		FindEvents: args,
	}

	resp := FindEventsResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *searchPort) OptFindMetadata(args FindMetadata) (*FindMetadataResponse, error) {
	req := struct {
		XMLName      string `xml:"tse:FindMetadata"`
		FindMetadata FindMetadata
	}{
		FindMetadata: args,
	}

	resp := FindMetadataResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *searchPort) OptFindPTZPosition(args FindPTZPosition) (*FindPTZPositionResponse, error) {
	req := struct {
		XMLName         string `xml:"tse:FindPTZPosition"`
		FindPTZPosition FindPTZPosition
	}{
		FindPTZPosition: args,
	}

	resp := FindPTZPositionResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *searchPort) OptFindRecordings(args FindRecordings) (*FindRecordingsResponse, error) {
	req := struct {
		XMLName        string `xml:"tse:FindRecordings"`
		FindRecordings FindRecordings
	}{
		FindRecordings: args,
	}

	resp := FindRecordingsResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *searchPort) OptGetEventSearchResults(args GetEventSearchResults) (*GetEventSearchResultsResponse, error) {
	req := struct {
		XMLName               string `xml:"tse:GetEventSearchResults"`
		GetEventSearchResults GetEventSearchResults
	}{
		GetEventSearchResults: args,
	}

	resp := GetEventSearchResultsResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *searchPort) OptGetMediaAttributes(args GetMediaAttributes) (*GetMediaAttributesResponse, error) {
	req := struct {
		XMLName            string `xml:"tse:GetMediaAttributes"`
		GetMediaAttributes GetMediaAttributes
	}{
		GetMediaAttributes: args,
	}

	resp := GetMediaAttributesResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *searchPort) OptGetMetadataSearchResults(args GetMetadataSearchResults) (*GetMetadataSearchResultsResponse, error) {
	req := struct {
		XMLName                  string `xml:"tse:GetMetadataSearchResults"`
		GetMetadataSearchResults GetMetadataSearchResults
	}{
		GetMetadataSearchResults: args,
	}

	resp := GetMetadataSearchResultsResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *searchPort) OptGetPTZPositionSearchResults(args GetPTZPositionSearchResults) (*GetPTZPositionSearchResultsResponse, error) {
	req := struct {
		XMLName                     string `xml:"tse:GetPTZPositionSearchResults"`
		GetPTZPositionSearchResults GetPTZPositionSearchResults
	}{
		GetPTZPositionSearchResults: args,
	}

	resp := GetPTZPositionSearchResultsResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *searchPort) OptGetRecordingInformation(args GetRecordingInformation) (*GetRecordingInformationResponse, error) {
	req := struct {
		XMLName                 string `xml:"tse:GetRecordingInformation"`
		GetRecordingInformation GetRecordingInformation
	}{
		GetRecordingInformation: args,
	}

	resp := GetRecordingInformationResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *searchPort) OptGetRecordingSearchResults(args GetRecordingSearchResults) (*GetRecordingSearchResultsResponse, error) {
	req := struct {
		XMLName                   string `xml:"tse:GetRecordingSearchResults"`
		GetRecordingSearchResults GetRecordingSearchResults
	}{
		GetRecordingSearchResults: args,
	}

	resp := GetRecordingSearchResultsResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *searchPort) OptGetRecordingSummary(args GetRecordingSummary) (*GetRecordingSummaryResponse, error) {
	req := struct {
		XMLName             string `xml:"tse:GetRecordingSummary"`
		GetRecordingSummary GetRecordingSummary
	}{
		GetRecordingSummary: args,
	}

	resp := GetRecordingSummaryResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *searchPort) OptGetSearchState(args GetSearchState) (*GetSearchStateResponse, error) {
	req := struct {
		XMLName        string `xml:"tse:GetSearchState"`
		GetSearchState GetSearchState
	}{
		GetSearchState: args,
	}

	resp := GetSearchStateResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *searchPort) OptGetServiceCapabilities(args GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error) {
	req := struct {
		XMLName                string `xml:"tse:GetServiceCapabilities"`
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
