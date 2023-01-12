package action_engine

import (
	"reflect"

	"code.byted.org/videoarch/go-onvif/onvif/common"
)

var Namespace = "http://www.onvif.org/ver10/actionengine/wsdl"

// NewActionEnginePort creates an initializes a ActionEnginePort.
func NewActionEnginePort(endpoint string, cli common.Client) ActionEnginePort {
	return &actionEnginePort{cli: cli, Endpoint: endpoint}
}

// ActionEnginePort was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type ActionEnginePort interface {
	OptCreateActionTriggers(CreateActionTriggers CreateActionTriggers) (*CreateActionTriggersResponse, error)

	OptCreateActions(CreateActions CreateActions) (*CreateActionsResponse, error)

	OptDeleteActionTriggers(DeleteActionTriggers DeleteActionTriggers) (*DeleteActionTriggersResponse, error)

	OptDeleteActions(DeleteActions DeleteActions) (*DeleteActionsResponse, error)

	OptGetActionTriggers(GetActionTriggers GetActionTriggers) (*GetActionTriggersResponse, error)

	OptGetActions(GetActions GetActions) (*GetActionsResponse, error)

	OptGetServiceCapabilities(GetServiceCapabilities GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error)

	OptGetSupportedActions(GetSupportedActions GetSupportedActions) (*GetSupportedActionsResponse, error)

	OptModifyActionTriggers(ModifyActionTriggers ModifyActionTriggers) (*ModifyActionTriggersResponse, error)

	OptModifyActions(ModifyActions ModifyActions) (*ModifyActionsResponse, error)
}
type Duration string

type AddressFormatType string

func (v AddressFormatType) Validate() bool {
	for _, vv := range []string{
		"hostname",
		"ipv4",
		"ipv6",
		"Extended",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type EMailAuthenticationMode string

func (v EMailAuthenticationMode) Validate() bool {
	for _, vv := range []string{
		"none",
		"SMTP",
		"POPSMTP",
		"Extended",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type FileSuffixType string

func (v FileSuffixType) Validate() bool {
	for _, vv := range []string{
		"none",
		"sequence",
		"dateTime",
		"Extended",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type HttpAuthenticationMethodType string

func (v HttpAuthenticationMethodType) Validate() bool {
	for _, vv := range []string{
		"none",
		"MD5Digest",
		"Extended",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type HttpProtocolType string

func (v HttpProtocolType) Validate() bool {
	for _, vv := range []string{
		"http",
		"https",
		"Extended",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type Action struct {
	Configuration ActionConfiguration   `xml:"Configuration,omitempty" json:"Configuration,omitempty" yaml:"Configuration,omitempty"`
	Token         common.ReferenceToken `xml:"Token,attr,omitempty" json:"Token,attr,omitempty" yaml:"Token,attr,omitempty"`
}

type ActionConfigDescription struct {
	ParameterDescription *common.ItemListDescription `xml:"ParameterDescription,omitempty" json:"ParameterDescription,omitempty" yaml:"ParameterDescription,omitempty"`
	Name                 string                      `xml:"Name,attr,omitempty" json:"Name,attr,omitempty" yaml:"Name,attr,omitempty"`
}

type ActionConfiguration struct {
	Parameters *common.ItemList `xml:"Parameters,omitempty" json:"Parameters,omitempty" yaml:"Parameters,omitempty"`
	Name       string           `xml:"Name,attr,omitempty" json:"Name,attr,omitempty" yaml:"Name,attr,omitempty"`
	Type       string           `xml:"Type,attr,omitempty" json:"Type,attr,omitempty" yaml:"Type,attr,omitempty"`
}

type ActionEngineCapabilities struct {
	ActionCapabilities []ActionTypeLimits                `xml:"ActionCapabilities,omitempty" json:"ActionCapabilities,omitempty" yaml:"ActionCapabilities,omitempty"`
	Extension          ActionEngineCapabilitiesExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
	MaximumTriggers    uint64                            `xml:"MaximumTriggers,attr,omitempty" json:"MaximumTriggers,attr,omitempty" yaml:"MaximumTriggers,attr,omitempty"`
	MaximumActions     uint64                            `xml:"MaximumActions,attr,omitempty" json:"MaximumActions,attr,omitempty" yaml:"MaximumActions,attr,omitempty"`
}

type ActionEngineCapabilitiesExtension []interface{}

type ActionTrigger struct {
	Configuration ActionTriggerConfiguration `xml:"Configuration,omitempty" json:"Configuration,omitempty" yaml:"Configuration,omitempty"`
	Token         common.ReferenceToken      `xml:"Token,attr,omitempty" json:"Token,attr,omitempty" yaml:"Token,attr,omitempty"`
}

type ActionTriggerConfiguration struct {
	TopicExpression   *common.TopicExpressionType         `xml:"TopicExpression,omitempty" json:"TopicExpression,omitempty" yaml:"TopicExpression,omitempty"`
	ContentExpression *common.QueryExpressionType         `xml:"ContentExpression,omitempty" json:"ContentExpression,omitempty" yaml:"ContentExpression,omitempty"`
	ActionToken       []*common.ReferenceToken            `xml:"ActionToken,omitempty" json:"ActionToken,omitempty" yaml:"ActionToken,omitempty"`
	Extension         ActionTriggerConfigurationExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type ActionTriggerConfigurationExtension []interface{}

type ActionTypeLimits []interface{}

type AuthenticationConfig struct {
	User UserCredentials         `xml:"User,omitempty" json:"User,omitempty" yaml:"User,omitempty"`
	Mode EMailAuthenticationMode `xml:"mode,attr,omitempty" json:"mode,attr,omitempty" yaml:"mode,attr,omitempty"`
}

type CreateActionTriggers struct {
	ActionTrigger []ActionTriggerConfiguration `xml:"ActionTrigger,omitempty" json:"ActionTrigger,omitempty" yaml:"ActionTrigger,omitempty"`
}

type CreateActionTriggersResponse struct {
	ActionTrigger []ActionTrigger `xml:"ActionTrigger,omitempty" json:"ActionTrigger,omitempty" yaml:"ActionTrigger,omitempty"`
}

type CreateActions struct {
	Action []ActionConfiguration `xml:"Action,omitempty" json:"Action,omitempty" yaml:"Action,omitempty"`
}

type CreateActionsResponse struct {
	Action []Action `xml:"Action,omitempty" json:"Action,omitempty" yaml:"Action,omitempty"`
}

type DeleteActionTriggers struct {
	Token []*common.ReferenceToken `xml:"Token,omitempty" json:"Token,omitempty" yaml:"Token,omitempty"`
}

type DeleteActionTriggersResponse struct {
}

type DeleteActions struct {
	Token []*common.ReferenceToken `xml:"Token,omitempty" json:"Token,omitempty" yaml:"Token,omitempty"`
}

type DeleteActionsResponse struct {
}

type EMailAttachmentConfiguration struct {
	FileName  string                                `xml:"FileName,omitempty" json:"FileName,omitempty" yaml:"FileName,omitempty"`
	DoSuffix  FileSuffixType                        `xml:"doSuffix,omitempty" json:"doSuffix,omitempty" yaml:"doSuffix,omitempty"`
	Extension EMailAttachmentConfigurationExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type EMailAttachmentConfigurationExtension []interface{}

type EMailBodyTextConfiguration []interface{}

type EMailReceiverConfiguration struct {
	TO        []string                            `xml:"TO" json:"TO" yaml:"TO"`
	CC        []string                            `xml:"CC,omitempty" json:"CC,omitempty" yaml:"CC,omitempty"`
	Extension EMailReceiverConfigurationExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type EMailReceiverConfigurationExtension []interface{}

type EMailServerConfiguration struct {
	SMTPConfig           SMTPConfig           `xml:"SMTPConfig,omitempty" json:"SMTPConfig,omitempty" yaml:"SMTPConfig,omitempty"`
	POPConfig            POPConfig            `xml:"POPConfig,omitempty" json:"POPConfig,omitempty" yaml:"POPConfig,omitempty"`
	AuthenticationConfig AuthenticationConfig `xml:"AuthenticationConfig,omitempty" json:"AuthenticationConfig,omitempty" yaml:"AuthenticationConfig,omitempty"`
}

type FtpAuthenticationConfiguration struct {
	User      UserCredentials                         `xml:"User,omitempty" json:"User,omitempty" yaml:"User,omitempty"`
	Extension FtpAuthenticationConfigurationExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type FtpAuthenticationConfigurationExtension []interface{}

type FtpContent struct {
	FtpContentConfig FtpContentConfiguration `xml:"FtpContentConfig,omitempty" json:"FtpContentConfig,omitempty" yaml:"FtpContentConfig,omitempty"`
}

type FtpContentConfiguration []interface{}

type FtpContentConfigurationUploadFile struct {
	SourceFileName      string `xml:"sourceFileName,omitempty" json:"sourceFileName,omitempty" yaml:"sourceFileName,omitempty"`
	DestinationFileName string `xml:"destinationFileName,omitempty" json:"destinationFileName,omitempty" yaml:"destinationFileName,omitempty"`
}

type FtpContentConfigurationUploadImages struct {
	HowLong        *common.Duration          `xml:"HowLong,omitempty" json:"HowLong,omitempty" yaml:"HowLong,omitempty"`
	SampleInterval *common.Duration          `xml:"SampleInterval,omitempty" json:"SampleInterval,omitempty" yaml:"SampleInterval,omitempty"`
	FileName       FtpFileNameConfigurations `xml:"FileName,omitempty" json:"FileName,omitempty" yaml:"FileName,omitempty"`
}

type FtpDestinationConfiguration struct {
	HostAddress       FtpHostAddress                       `xml:"HostAddress,omitempty" json:"HostAddress,omitempty" yaml:"HostAddress,omitempty"`
	UploadPath        string                               `xml:"UploadPath,omitempty" json:"UploadPath,omitempty" yaml:"UploadPath,omitempty"`
	FtpAuthentication FtpAuthenticationConfiguration       `xml:"FtpAuthentication,omitempty" json:"FtpAuthentication,omitempty" yaml:"FtpAuthentication,omitempty"`
	Extension         FtpDestinationConfigurationExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type FtpDestinationConfigurationExtension []interface{}

type FtpFileNameConfigurations []interface{}

type FtpHostAddress struct {
	Value      string            `xml:"Value,omitempty" json:"Value,omitempty" yaml:"Value,omitempty"`
	FormatType AddressFormatType `xml:"formatType,attr,omitempty" json:"formatType,attr,omitempty" yaml:"formatType,attr,omitempty"`
	PortNo     int64             `xml:"portNo,attr,omitempty" json:"portNo,attr,omitempty" yaml:"portNo,attr,omitempty"`
}

type FtpHostConfigurations struct {
	FtpDestination []FtpDestinationConfiguration  `xml:"FtpDestination" json:"FtpDestination" yaml:"FtpDestination"`
	Extension      FtpHostConfigurationsExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type FtpHostConfigurationsExtension []interface{}

type GetActionTriggers struct {
}

type GetActionTriggersResponse struct {
	ActionTrigger []ActionTrigger `xml:"ActionTrigger,omitempty" json:"ActionTrigger,omitempty" yaml:"ActionTrigger,omitempty"`
}

type GetActions struct {
}

type GetActionsResponse struct {
	Action []Action `xml:"Action,omitempty" json:"Action,omitempty" yaml:"Action,omitempty"`
}

type GetServiceCapabilities struct {
}

type GetServiceCapabilitiesResponse struct {
	Capabilities ActionEngineCapabilities `xml:"Capabilities,omitempty" json:"Capabilities,omitempty" yaml:"Capabilities,omitempty"`
}

type GetSupportedActions struct {
}

type GetSupportedActionsResponse struct {
	SupportedActions SupportedActions `xml:"SupportedActions,omitempty" json:"SupportedActions,omitempty" yaml:"SupportedActions,omitempty"`
}

type HostAddress struct {
	Value      string            `xml:"Value,omitempty" json:"Value,omitempty" yaml:"Value,omitempty"`
	FormatType AddressFormatType `xml:"formatType,attr,omitempty" json:"formatType,attr,omitempty" yaml:"formatType,attr,omitempty"`
}

type HttpAuthenticationConfiguration struct {
	User      UserCredentials                          `xml:"User,omitempty" json:"User,omitempty" yaml:"User,omitempty"`
	Extension HttpAuthenticationConfigurationExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
	Method    HttpAuthenticationMethodType             `xml:"method,attr,omitempty" json:"method,attr,omitempty" yaml:"method,attr,omitempty"`
}

type HttpAuthenticationConfigurationExtension []interface{}

type HttpDestinationConfiguration struct {
	HostAddress        HttpHostAddress                       `xml:"HostAddress,omitempty" json:"HostAddress,omitempty" yaml:"HostAddress,omitempty"`
	HttpAuthentication HttpAuthenticationConfiguration       `xml:"HttpAuthentication,omitempty" json:"HttpAuthentication,omitempty" yaml:"HttpAuthentication,omitempty"`
	Extension          HttpDestinationConfigurationExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
	Uri                string                                `xml:"uri,attr,omitempty" json:"uri,attr,omitempty" yaml:"uri,attr,omitempty"`
	Protocol           HttpProtocolType                      `xml:"protocol,attr,omitempty" json:"protocol,attr,omitempty" yaml:"protocol,attr,omitempty"`
}

type HttpDestinationConfigurationExtension []interface{}

type HttpHostAddress struct {
	Value      string            `xml:"Value,omitempty" json:"Value,omitempty" yaml:"Value,omitempty"`
	FormatType AddressFormatType `xml:"formatType,attr,omitempty" json:"formatType,attr,omitempty" yaml:"formatType,attr,omitempty"`
	PortNo     int64             `xml:"portNo,attr,omitempty" json:"portNo,attr,omitempty" yaml:"portNo,attr,omitempty"`
}

type HttpHostConfigurations struct {
	HttpDestination []HttpDestinationConfiguration  `xml:"HttpDestination" json:"HttpDestination" yaml:"HttpDestination"`
	Extension       HttpHostConfigurationsExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type HttpHostConfigurationsExtension []interface{}

type MediaSource struct {
	ProfileToken *common.ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty" yaml:"ProfileToken,omitempty"`
}

type ModifyActionTriggers struct {
	ActionTrigger []ActionTrigger `xml:"ActionTrigger,omitempty" json:"ActionTrigger,omitempty" yaml:"ActionTrigger,omitempty"`
}

type ModifyActionTriggersResponse struct {
}

type ModifyActions struct {
	Action []Action `xml:"Action,omitempty" json:"Action,omitempty" yaml:"Action,omitempty"`
}

type ModifyActionsResponse struct {
}

type POPConfig struct {
	HostAddress HostAddress `xml:"HostAddress,omitempty" json:"HostAddress,omitempty" yaml:"HostAddress,omitempty"`
}

type PostBodyConfiguration []interface{}

type PostContentConfiguration struct {
	MediaReference MediaSource           `xml:"MediaReference,omitempty" json:"MediaReference,omitempty" yaml:"MediaReference,omitempty"`
	PostBody       PostBodyConfiguration `xml:"PostBody,omitempty" json:"PostBody,omitempty" yaml:"PostBody,omitempty"`
}

type RecordingActionConfiguration struct {
	RecordConfig TriggeredRecordingConfiguration `xml:"RecordConfig,omitempty" json:"RecordConfig,omitempty" yaml:"RecordConfig,omitempty"`
}

type SMSMessage struct {
	Text string `xml:"Text,omitempty" json:"Text,omitempty" yaml:"Text,omitempty"`
}

type SMSProviderConfiguration struct {
	ProviderURL string          `xml:"ProviderURL,omitempty" json:"ProviderURL,omitempty" yaml:"ProviderURL,omitempty"`
	User        UserCredentials `xml:"User,omitempty" json:"User,omitempty" yaml:"User,omitempty"`
}

type SMSSenderConfiguration struct {
	EMail string `xml:"EMail,omitempty" json:"EMail,omitempty" yaml:"EMail,omitempty"`
}

type SMTPConfig struct {
	HostAddress HostAddress `xml:"HostAddress,omitempty" json:"HostAddress,omitempty" yaml:"HostAddress,omitempty"`
	PortNo      uint64      `xml:"portNo,attr,omitempty" json:"portNo,attr,omitempty" yaml:"portNo,attr,omitempty"`
}

type SupportedActions struct {
	ActionContentSchemaLocation []string                  `xml:"ActionContentSchemaLocation,omitempty" json:"ActionContentSchemaLocation,omitempty" yaml:"ActionContentSchemaLocation,omitempty"`
	ActionDescription           []ActionConfigDescription `xml:"ActionDescription,omitempty" json:"ActionDescription,omitempty" yaml:"ActionDescription,omitempty"`
	Extension                   SupportedActionsExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type SupportedActionsExtension []interface{}

type TriggeredRecordingConfiguration struct {
	PreRecordDuration  *common.Duration `xml:"PreRecordDuration,omitempty" json:"PreRecordDuration,omitempty" yaml:"PreRecordDuration,omitempty"`
	PostRecordDuration *common.Duration `xml:"PostRecordDuration,omitempty" json:"PostRecordDuration,omitempty" yaml:"PostRecordDuration,omitempty"`
	RecordDuration     *common.Duration `xml:"RecordDuration,omitempty" json:"RecordDuration,omitempty" yaml:"RecordDuration,omitempty"`
	RecordFrameRate    uint64           `xml:"RecordFrameRate,omitempty" json:"RecordFrameRate,omitempty" yaml:"RecordFrameRate,omitempty"`
	DoRecordAudio      bool             `xml:"DoRecordAudio,omitempty" json:"DoRecordAudio,omitempty" yaml:"DoRecordAudio,omitempty"`
}

type UserCredentials struct {
	Username  string                   `xml:"username,omitempty" json:"username,omitempty" yaml:"username,omitempty"`
	Password  []byte                   `xml:"password,omitempty" json:"password,omitempty" yaml:"password,omitempty"`
	Extension UserCredentialsExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type UserCredentialsExtension []interface{}

type Onvif_action struct {
	ActionDescription []ActionConfigDescription `xml:"ActionDescription,omitempty" json:"ActionDescription,omitempty" yaml:"ActionDescription,omitempty"`
}

// actionEnginePort implements the ActionEnginePort interface.
type actionEnginePort struct {
	cli      common.Client
	Endpoint string
}

func (p *actionEnginePort) OptCreateActionTriggers(args CreateActionTriggers) (*CreateActionTriggersResponse, error) {
	req := struct {
		XMLName              string `xml:"tae:CreateActionTriggers"`
		CreateActionTriggers CreateActionTriggers
	}{
		CreateActionTriggers: args,
	}

	resp := CreateActionTriggersResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *actionEnginePort) OptCreateActions(args CreateActions) (*CreateActionsResponse, error) {
	req := struct {
		XMLName       string `xml:"tae:CreateActions"`
		CreateActions CreateActions
	}{
		CreateActions: args,
	}

	resp := CreateActionsResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *actionEnginePort) OptDeleteActionTriggers(args DeleteActionTriggers) (*DeleteActionTriggersResponse, error) {
	req := struct {
		XMLName              string `xml:"tae:DeleteActionTriggers"`
		DeleteActionTriggers DeleteActionTriggers
	}{
		DeleteActionTriggers: args,
	}

	resp := DeleteActionTriggersResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *actionEnginePort) OptDeleteActions(args DeleteActions) (*DeleteActionsResponse, error) {
	req := struct {
		XMLName       string `xml:"tae:DeleteActions"`
		DeleteActions DeleteActions
	}{
		DeleteActions: args,
	}

	resp := DeleteActionsResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *actionEnginePort) OptGetActionTriggers(args GetActionTriggers) (*GetActionTriggersResponse, error) {
	req := struct {
		XMLName           string `xml:"tae:GetActionTriggers"`
		GetActionTriggers GetActionTriggers
	}{
		GetActionTriggers: args,
	}

	resp := GetActionTriggersResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *actionEnginePort) OptGetActions(args GetActions) (*GetActionsResponse, error) {
	req := struct {
		XMLName    string `xml:"tae:GetActions"`
		GetActions GetActions
	}{
		GetActions: args,
	}

	resp := GetActionsResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *actionEnginePort) OptGetServiceCapabilities(args GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error) {
	req := struct {
		XMLName                string `xml:"tae:GetServiceCapabilities"`
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

func (p *actionEnginePort) OptGetSupportedActions(args GetSupportedActions) (*GetSupportedActionsResponse, error) {
	req := struct {
		XMLName             string `xml:"tae:GetSupportedActions"`
		GetSupportedActions GetSupportedActions
	}{
		GetSupportedActions: args,
	}

	resp := GetSupportedActionsResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *actionEnginePort) OptModifyActionTriggers(args ModifyActionTriggers) (*ModifyActionTriggersResponse, error) {
	req := struct {
		XMLName              string `xml:"tae:ModifyActionTriggers"`
		ModifyActionTriggers ModifyActionTriggers
	}{
		ModifyActionTriggers: args,
	}

	resp := ModifyActionTriggersResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *actionEnginePort) OptModifyActions(args ModifyActions) (*ModifyActionsResponse, error) {
	req := struct {
		XMLName       string `xml:"tae:ModifyActions"`
		ModifyActions ModifyActions
	}{
		ModifyActions: args,
	}

	resp := ModifyActionsResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
