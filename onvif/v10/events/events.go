package events

import (
	"reflect"

	"code.byted.org/videoarch/go-onvif/onvif/common"
)

var Namespace = "http://docs.oasis-open.org/wsrf/rw-2"

func NewPort(endpoint string, cli common.Client) Port {
	return &pausableSubscriptionManager{cli: cli, Endpoint: endpoint}
}

type Port interface {
	OptAddEventBroker(AddEventBroker AddEventBroker) (*AddEventBrokerResponse, error)

	OptCreatePullPoint(CreatePullPoint CreatePullPoint) (*CreatePullPointResponse, error)

	OptCreatePullPointSubscription(CreatePullPointSubscription CreatePullPointSubscription) (*CreatePullPointSubscriptionResponse, error)

	OptDeleteEventBroker(DeleteEventBroker DeleteEventBroker) (*DeleteEventBrokerResponse, error)

	OptDestroyPullPoint(DestroyPullPoint DestroyPullPoint) (*DestroyPullPointResponse, error)

	OptGetCurrentMessage(GetCurrentMessage GetCurrentMessage) (*GetCurrentMessageResponse, error)

	OptGetEventBrokers(GetEventBrokers GetEventBrokers) (*GetEventBrokersResponse, error)

	OptGetEventProperties(GetEventProperties GetEventProperties) (*GetEventPropertiesResponse, error)

	OptGetMessages(GetMessages GetMessages) (*GetMessagesResponse, error)

	OptGetServiceCapabilities(GetServiceCapabilities GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error)

	OptNotify(Notify Notify) error

	OptPauseSubscription(PauseSubscription PauseSubscription) (*PauseSubscriptionResponse, error)

	OptPullMessages(PullMessages PullMessages) (*PullMessagesResponse, error)

	OptRenew(Renew Renew) (*RenewResponse, error)

	OptResumeSubscription(ResumeSubscription ResumeSubscription) (*ResumeSubscriptionResponse, error)

	OptSeek(Seek Seek) (*SeekResponse, error)

	OptSetSynchronizationPoint(SetSynchronizationPoint SetSynchronizationPoint) (*SetSynchronizationPointResponse, error)

	OptSubscribe(Subscribe Subscribe) (*SubscribeResponse, error)

	OptUnsubscribe(Unsubscribe Unsubscribe) (*UnsubscribeResponse, error)
}
type DateTime string

type Duration string

type AbsoluteOrRelativeTimeType interface{}

type ConcreteTopicExpression string

type ConnectionStatus string

func (v ConnectionStatus) Validate() bool {
	for _, vv := range []string{
		"Offline",
		"Connecting",
		"Connected",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type EventBrokerProtocol string

func (v EventBrokerProtocol) Validate() bool {
	for _, vv := range []string{
		"mqtt",
		"mqtts",
		"ws",
		"wss",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type FaultCodesOpenEnumType interface{}

type FaultCodesType string

func (v FaultCodesType) Validate() bool {
	for _, vv := range []string{
		"tns:InvalidAddressingHeader",
		"tns:InvalidAddress",
		"tns:InvalidEPR",
		"tns:InvalidCardinality",
		"tns:MissingAddressInEPR",
		"tns:DuplicateMessageID",
		"tns:ActionMismatch",
		"tns:MessageAddressingHeaderRequired",
		"tns:DestinationUnreachable",
		"tns:ActionNotSupported",
		"tns:EndpointUnavailable",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type FullTopicExpression string

type RelationshipType string

func (v RelationshipType) Validate() bool {
	for _, vv := range []string{
		"http://www.w3.org/2005/08/addressing/reply",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type RelationshipTypeOpenEnum interface{}

type SimpleTopicExpression string

type AddEventBroker struct {
	EventBroker EventBrokerConfig `xml:"EventBroker" json:"EventBroker" yaml:"EventBroker"`
}

type AddEventBrokerResponse struct {
}

type AttributedAnyType []interface{}

type AttributedQNameType struct {
	Content string `xml:"Content,omitempty" json:"Content,omitempty" yaml:"Content,omitempty"`
}

type AttributedURIType struct {
	Content string `xml:"Content,omitempty" json:"Content,omitempty" yaml:"Content,omitempty"`
}

type AttributedUnsignedLongType struct {
	Content *uint64 `xml:"Content,omitempty" json:"Content,omitempty" yaml:"Content,omitempty"`
}

type BaseFaultType struct {
	Timestamp   common.DateTime       `xml:"Timestamp" json:"Timestamp" yaml:"Timestamp"`
	Originator  EndpointReferenceType `xml:"Originator,omitempty" json:"Originator,omitempty" yaml:"Originator,omitempty"`
	ErrorCode   string                `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	Description []string              `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	FaultCause  string                `xml:"FaultCause" json:"FaultCause" yaml:"FaultCause"`
}

type Capabilities []interface{}

type CreatePullPoint []interface{}

type CreatePullPointResponse struct {
	PullPoint EndpointReferenceType `xml:"PullPoint,omitempty" json:"PullPoint,omitempty" yaml:"PullPoint,omitempty"`
}

type CreatePullPointSubscription struct {
	Filter                 FilterType                 `xml:"Filter,omitempty" json:"Filter,omitempty" yaml:"Filter,omitempty"`
	InitialTerminationTime AbsoluteOrRelativeTimeType `xml:"InitialTerminationTime,omitempty" json:"InitialTerminationTime,omitempty" yaml:"InitialTerminationTime,omitempty"`
	SubscriptionPolicy     []string                   `xml:"SubscriptionPolicy>SubscriptionPolicy,omitempty" json:"SubscriptionPolicy>SubscriptionPolicy,omitempty" yaml:"SubscriptionPolicy>SubscriptionPolicy,omitempty"`
}

type CreatePullPointSubscriptionResponse struct {
	SubscriptionReference EndpointReferenceType `xml:"SubscriptionReference,omitempty" json:"SubscriptionReference,omitempty" yaml:"SubscriptionReference,omitempty"`
	CurrentTime           *common.DateTime      `xml:"CurrentTime,omitempty" json:"CurrentTime,omitempty" yaml:"CurrentTime,omitempty"`
	TerminationTime       *common.DateTime      `xml:"TerminationTime,omitempty" json:"TerminationTime,omitempty" yaml:"TerminationTime,omitempty"`
}

type DeleteEventBroker struct {
	Address string `xml:"Address" json:"Address" yaml:"Address"`
}

type DeleteEventBrokerResponse struct {
}

type DestroyPullPoint []interface{}

type DestroyPullPointResponse []interface{}

type Documentation []interface{}

type EndpointReferenceType struct {
	Address             AttributedURIType       `xml:"Address,omitempty" json:"Address,omitempty" yaml:"Address,omitempty"`
	ReferenceParameters ReferenceParametersType `xml:"ReferenceParameters,omitempty" json:"ReferenceParameters,omitempty" yaml:"ReferenceParameters,omitempty"`
	Metadata            MetadataType            `xml:"Metadata,omitempty" json:"Metadata,omitempty" yaml:"Metadata,omitempty"`
}

type EventBrokerConfig struct {
	Address                    string     `xml:"Address" json:"Address" yaml:"Address"`
	TopicPrefix                string     `xml:"TopicPrefix" json:"TopicPrefix" yaml:"TopicPrefix"`
	UserName                   string     `xml:"UserName,omitempty" json:"UserName,omitempty" yaml:"UserName,omitempty"`
	Password                   string     `xml:"Password,omitempty" json:"Password,omitempty" yaml:"Password,omitempty"`
	CertificateID              string     `xml:"CertificateID,omitempty" json:"CertificateID,omitempty" yaml:"CertificateID,omitempty"`
	PublishFilter              FilterType `xml:"PublishFilter,omitempty" json:"PublishFilter,omitempty" yaml:"PublishFilter,omitempty"`
	QoS                        int        `xml:"QoS,omitempty" json:"QoS,omitempty" yaml:"QoS,omitempty"`
	Status                     string     `xml:"Status,omitempty" json:"Status,omitempty" yaml:"Status,omitempty"`
	CertPathValidationPolicyID string     `xml:"CertPathValidationPolicyID,omitempty" json:"CertPathValidationPolicyID,omitempty" yaml:"CertPathValidationPolicyID,omitempty"`
	MetadataFilter             FilterType `xml:"MetadataFilter,omitempty" json:"MetadataFilter,omitempty" yaml:"MetadataFilter,omitempty"`
}

type ExtensibleDocumented interface{}

type FilterType []interface{}

type GetCurrentMessage struct {
	Topic TopicExpressionType `xml:"Topic,omitempty" json:"Topic,omitempty" yaml:"Topic,omitempty"`
}

type GetCurrentMessageResponse []interface{}

type GetEventBrokers struct {
	Address string `xml:"Address,omitempty" json:"Address,omitempty" yaml:"Address,omitempty"`
}

type GetEventBrokersResponse struct {
	EventBroker []EventBrokerConfig `xml:"EventBroker,omitempty" json:"EventBroker,omitempty" yaml:"EventBroker,omitempty"`
}

type GetEventProperties struct {
}

type GetEventPropertiesResponse struct {
	TopicNamespaceLocation          []string     `xml:"TopicNamespaceLocation" json:"TopicNamespaceLocation" yaml:"TopicNamespaceLocation"`
	FixedTopicSet                   bool         `xml:"FixedTopicSet,omitempty" json:"FixedTopicSet,omitempty" yaml:"FixedTopicSet,omitempty"`
	TopicSet                        TopicSetType `xml:"TopicSet,omitempty" json:"TopicSet,omitempty" yaml:"TopicSet,omitempty"`
	TopicExpressionDialect          string       `xml:"TopicExpressionDialect,omitempty" json:"TopicExpressionDialect,omitempty" yaml:"TopicExpressionDialect,omitempty"`
	MessageContentFilterDialect     []string     `xml:"MessageContentFilterDialect" json:"MessageContentFilterDialect" yaml:"MessageContentFilterDialect"`
	ProducerPropertiesFilterDialect []string     `xml:"ProducerPropertiesFilterDialect,omitempty" json:"ProducerPropertiesFilterDialect,omitempty" yaml:"ProducerPropertiesFilterDialect,omitempty"`
	MessageContentSchemaLocation    []string     `xml:"MessageContentSchemaLocation" json:"MessageContentSchemaLocation" yaml:"MessageContentSchemaLocation"`
}

type GetMessages struct {
	MaximumNumber uint `xml:"MaximumNumber,omitempty" json:"MaximumNumber,omitempty" yaml:"MaximumNumber,omitempty"`
}

type GetMessagesResponse struct {
	NotificationMessage NotificationMessageHolderType `xml:"NotificationMessage,omitempty" json:"NotificationMessage,omitempty" yaml:"NotificationMessage,omitempty"`
}

type GetServiceCapabilities struct {
}

type GetServiceCapabilitiesResponse struct {
	Capabilities Capabilities `xml:"Capabilities,omitempty" json:"Capabilities,omitempty" yaml:"Capabilities,omitempty"`
}

type InvalidFilterFaultType struct {
	Timestamp     common.DateTime       `xml:"Timestamp" json:"Timestamp" yaml:"Timestamp"`
	Originator    EndpointReferenceType `xml:"Originator,omitempty" json:"Originator,omitempty" yaml:"Originator,omitempty"`
	ErrorCode     string                `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	Description   []string              `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	FaultCause    string                `xml:"FaultCause" json:"FaultCause" yaml:"FaultCause"`
	UnknownFilter []string              `xml:"UnknownFilter" json:"UnknownFilter" yaml:"UnknownFilter"`
	TypeAttrXSI   string                `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string                `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *InvalidFilterFaultType) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:InvalidFilterFaultType"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://docs.oasis-open.org/wsn/b-2"
	}
}

type InvalidMessageContentExpressionFaultType struct {
	Timestamp     common.DateTime       `xml:"Timestamp" json:"Timestamp" yaml:"Timestamp"`
	Originator    EndpointReferenceType `xml:"Originator,omitempty" json:"Originator,omitempty" yaml:"Originator,omitempty"`
	ErrorCode     string                `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	Description   []string              `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	FaultCause    string                `xml:"FaultCause" json:"FaultCause" yaml:"FaultCause"`
	TypeAttrXSI   string                `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string                `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *InvalidMessageContentExpressionFaultType) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:InvalidMessageContentExpressionFaultType"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://docs.oasis-open.org/wsn/b-2"
	}
}

type InvalidProducerPropertiesExpressionFaultType struct {
	Timestamp     common.DateTime       `xml:"Timestamp" json:"Timestamp" yaml:"Timestamp"`
	Originator    EndpointReferenceType `xml:"Originator,omitempty" json:"Originator,omitempty" yaml:"Originator,omitempty"`
	ErrorCode     string                `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	Description   []string              `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	FaultCause    string                `xml:"FaultCause" json:"FaultCause" yaml:"FaultCause"`
	TypeAttrXSI   string                `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string                `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *InvalidProducerPropertiesExpressionFaultType) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:InvalidProducerPropertiesExpressionFaultType"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://docs.oasis-open.org/wsn/b-2"
	}
}

type InvalidTopicExpressionFaultType struct {
	Timestamp     common.DateTime       `xml:"Timestamp" json:"Timestamp" yaml:"Timestamp"`
	Originator    EndpointReferenceType `xml:"Originator,omitempty" json:"Originator,omitempty" yaml:"Originator,omitempty"`
	ErrorCode     string                `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	Description   []string              `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	FaultCause    string                `xml:"FaultCause" json:"FaultCause" yaml:"FaultCause"`
	TypeAttrXSI   string                `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string                `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *InvalidTopicExpressionFaultType) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:InvalidTopicExpressionFaultType"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://docs.oasis-open.org/wsn/b-2"
	}
}

type MetadataType []interface{}

type MultipleTopicsSpecifiedFaultType struct {
	Timestamp     common.DateTime       `xml:"Timestamp" json:"Timestamp" yaml:"Timestamp"`
	Originator    EndpointReferenceType `xml:"Originator,omitempty" json:"Originator,omitempty" yaml:"Originator,omitempty"`
	ErrorCode     string                `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	Description   []string              `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	FaultCause    string                `xml:"FaultCause" json:"FaultCause" yaml:"FaultCause"`
	TypeAttrXSI   string                `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string                `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *MultipleTopicsSpecifiedFaultType) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:MultipleTopicsSpecifiedFaultType"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://docs.oasis-open.org/wsn/b-2"
	}
}

type NoCurrentMessageOnTopicFaultType struct {
	Timestamp     common.DateTime       `xml:"Timestamp" json:"Timestamp" yaml:"Timestamp"`
	Originator    EndpointReferenceType `xml:"Originator,omitempty" json:"Originator,omitempty" yaml:"Originator,omitempty"`
	ErrorCode     string                `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	Description   []string              `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	FaultCause    string                `xml:"FaultCause" json:"FaultCause" yaml:"FaultCause"`
	TypeAttrXSI   string                `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string                `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *NoCurrentMessageOnTopicFaultType) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:NoCurrentMessageOnTopicFaultType"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://docs.oasis-open.org/wsn/b-2"
	}
}

type NotificationMessageHolderType struct {
	SubscriptionReference EndpointReferenceType `xml:"SubscriptionReference,omitempty" json:"SubscriptionReference,omitempty" yaml:"SubscriptionReference,omitempty"`
	Topic                 TopicExpressionType   `xml:"Topic,omitempty" json:"Topic,omitempty" yaml:"Topic,omitempty"`
	ProducerReference     EndpointReferenceType `xml:"ProducerReference,omitempty" json:"ProducerReference,omitempty" yaml:"ProducerReference,omitempty"`
	Message               string                `xml:"Message" json:"Message" yaml:"Message"`
}

type NotificationProducerRP struct {
	TopicExpression        TopicExpressionType `xml:"TopicExpression,omitempty" json:"TopicExpression,omitempty" yaml:"TopicExpression,omitempty"`
	FixedTopicSet          bool                `xml:"FixedTopicSet,omitempty" json:"FixedTopicSet,omitempty" yaml:"FixedTopicSet,omitempty"`
	TopicExpressionDialect string              `xml:"TopicExpressionDialect,omitempty" json:"TopicExpressionDialect,omitempty" yaml:"TopicExpressionDialect,omitempty"`
	TopicSet               TopicSetType        `xml:"TopicSet,omitempty" json:"TopicSet,omitempty" yaml:"TopicSet,omitempty"`
}

type Notify struct {
	NotificationMessage NotificationMessageHolderType `xml:"NotificationMessage,omitempty" json:"NotificationMessage,omitempty" yaml:"NotificationMessage,omitempty"`
}

type NotifyMessageNotSupportedFaultType struct {
	Timestamp     common.DateTime       `xml:"Timestamp" json:"Timestamp" yaml:"Timestamp"`
	Originator    EndpointReferenceType `xml:"Originator,omitempty" json:"Originator,omitempty" yaml:"Originator,omitempty"`
	ErrorCode     string                `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	Description   []string              `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	FaultCause    string                `xml:"FaultCause" json:"FaultCause" yaml:"FaultCause"`
	TypeAttrXSI   string                `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string                `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *NotifyMessageNotSupportedFaultType) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:NotifyMessageNotSupportedFaultType"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://docs.oasis-open.org/wsn/b-2"
	}
}

type PauseFailedFaultType struct {
	Timestamp     common.DateTime       `xml:"Timestamp" json:"Timestamp" yaml:"Timestamp"`
	Originator    EndpointReferenceType `xml:"Originator,omitempty" json:"Originator,omitempty" yaml:"Originator,omitempty"`
	ErrorCode     string                `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	Description   []string              `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	FaultCause    string                `xml:"FaultCause" json:"FaultCause" yaml:"FaultCause"`
	TypeAttrXSI   string                `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string                `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *PauseFailedFaultType) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:PauseFailedFaultType"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://docs.oasis-open.org/wsn/b-2"
	}
}

type PauseSubscription []interface{}

type PauseSubscriptionResponse []interface{}

type ProblemActionType struct {
	Action     AttributedURIType `xml:"Action,omitempty" json:"Action,omitempty" yaml:"Action,omitempty"`
	SoapAction string            `xml:"SoapAction,omitempty" json:"SoapAction,omitempty" yaml:"SoapAction,omitempty"`
}

type PullMessages struct {
	Timeout      *common.Duration `xml:"Timeout,omitempty" json:"Timeout,omitempty" yaml:"Timeout,omitempty"`
	MessageLimit int              `xml:"MessageLimit,omitempty" json:"MessageLimit,omitempty" yaml:"MessageLimit,omitempty"`
}

type PullMessagesFaultResponse struct {
	MaxTimeout      *common.Duration `xml:"MaxTimeout,omitempty" json:"MaxTimeout,omitempty" yaml:"MaxTimeout,omitempty"`
	MaxMessageLimit int              `xml:"MaxMessageLimit,omitempty" json:"MaxMessageLimit,omitempty" yaml:"MaxMessageLimit,omitempty"`
}

type PullMessagesResponse struct {
	CurrentTime         *common.DateTime              `xml:"CurrentTime,omitempty" json:"CurrentTime,omitempty" yaml:"CurrentTime,omitempty"`
	TerminationTime     *common.DateTime              `xml:"TerminationTime,omitempty" json:"TerminationTime,omitempty" yaml:"TerminationTime,omitempty"`
	NotificationMessage NotificationMessageHolderType `xml:"NotificationMessage,omitempty" json:"NotificationMessage,omitempty" yaml:"NotificationMessage,omitempty"`
}

type QueryExpressionType []interface{}

type ReferenceParametersType []interface{}

type RelatesToType struct {
	Content          string                   `xml:"Content,omitempty" json:"Content,omitempty" yaml:"Content,omitempty"`
	RelationshipType RelationshipTypeOpenEnum `xml:"RelationshipType,attr,omitempty" json:"RelationshipType,attr,omitempty" yaml:"RelationshipType,attr,omitempty"`
}

type Renew struct {
	TerminationTime AbsoluteOrRelativeTimeType `xml:"TerminationTime,omitempty" json:"TerminationTime,omitempty" yaml:"TerminationTime,omitempty"`
}

type RenewResponse struct {
	TerminationTime *common.DateTime `xml:"TerminationTime,omitempty" json:"TerminationTime,omitempty" yaml:"TerminationTime,omitempty"`
	CurrentTime     *common.DateTime `xml:"CurrentTime,omitempty" json:"CurrentTime,omitempty" yaml:"CurrentTime,omitempty"`
}

type ResourceUnavailableFaultType struct {
	Timestamp     common.DateTime       `xml:"Timestamp" json:"Timestamp" yaml:"Timestamp"`
	Originator    EndpointReferenceType `xml:"Originator,omitempty" json:"Originator,omitempty" yaml:"Originator,omitempty"`
	ErrorCode     string                `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	Description   []string              `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	FaultCause    string                `xml:"FaultCause" json:"FaultCause" yaml:"FaultCause"`
	TypeAttrXSI   string                `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string                `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *ResourceUnavailableFaultType) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:ResourceUnavailableFaultType"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://docs.oasis-open.org/wsrf/r-2"
	}
}

type ResourceUnknownFaultType struct {
	Timestamp     common.DateTime       `xml:"Timestamp" json:"Timestamp" yaml:"Timestamp"`
	Originator    EndpointReferenceType `xml:"Originator,omitempty" json:"Originator,omitempty" yaml:"Originator,omitempty"`
	ErrorCode     string                `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	Description   []string              `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	FaultCause    string                `xml:"FaultCause" json:"FaultCause" yaml:"FaultCause"`
	TypeAttrXSI   string                `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string                `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *ResourceUnknownFaultType) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:ResourceUnknownFaultType"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://docs.oasis-open.org/wsrf/r-2"
	}
}

type ResumeFailedFaultType struct {
	Timestamp     common.DateTime       `xml:"Timestamp" json:"Timestamp" yaml:"Timestamp"`
	Originator    EndpointReferenceType `xml:"Originator,omitempty" json:"Originator,omitempty" yaml:"Originator,omitempty"`
	ErrorCode     string                `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	Description   []string              `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	FaultCause    string                `xml:"FaultCause" json:"FaultCause" yaml:"FaultCause"`
	TypeAttrXSI   string                `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string                `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *ResumeFailedFaultType) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:ResumeFailedFaultType"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://docs.oasis-open.org/wsn/b-2"
	}
}

type ResumeSubscription []interface{}

type ResumeSubscriptionResponse []interface{}

type Seek struct {
	UtcTime *common.DateTime `xml:"UtcTime,omitempty" json:"UtcTime,omitempty" yaml:"UtcTime,omitempty"`
	Reverse bool             `xml:"Reverse,omitempty" json:"Reverse,omitempty" yaml:"Reverse,omitempty"`
}

type SeekResponse struct {
}

type SetSynchronizationPoint struct {
}

type SetSynchronizationPointResponse struct {
}

type Subscribe struct {
	ConsumerReference      EndpointReferenceType      `xml:"ConsumerReference" json:"ConsumerReference" yaml:"ConsumerReference"`
	Filter                 FilterType                 `xml:"Filter,omitempty" json:"Filter,omitempty" yaml:"Filter,omitempty"`
	InitialTerminationTime AbsoluteOrRelativeTimeType `xml:"InitialTerminationTime,omitempty" json:"InitialTerminationTime,omitempty" yaml:"InitialTerminationTime,omitempty"`
	SubscriptionPolicy     []string                   `xml:"SubscriptionPolicy>SubscriptionPolicy,omitempty" json:"SubscriptionPolicy>SubscriptionPolicy,omitempty" yaml:"SubscriptionPolicy>SubscriptionPolicy,omitempty"`
}

type SubscribeCreationFailedFaultType struct {
	Timestamp     common.DateTime       `xml:"Timestamp" json:"Timestamp" yaml:"Timestamp"`
	Originator    EndpointReferenceType `xml:"Originator,omitempty" json:"Originator,omitempty" yaml:"Originator,omitempty"`
	ErrorCode     string                `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	Description   []string              `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	FaultCause    string                `xml:"FaultCause" json:"FaultCause" yaml:"FaultCause"`
	TypeAttrXSI   string                `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string                `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *SubscribeCreationFailedFaultType) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:SubscribeCreationFailedFaultType"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://docs.oasis-open.org/wsn/b-2"
	}
}

type SubscribeResponse struct {
	SubscriptionReference EndpointReferenceType `xml:"SubscriptionReference" json:"SubscriptionReference" yaml:"SubscriptionReference"`
	CurrentTime           *common.DateTime      `xml:"CurrentTime,omitempty" json:"CurrentTime,omitempty" yaml:"CurrentTime,omitempty"`
	TerminationTime       *common.DateTime      `xml:"TerminationTime,omitempty" json:"TerminationTime,omitempty" yaml:"TerminationTime,omitempty"`
}

type SubscriptionManagerRP struct {
	ConsumerReference  EndpointReferenceType  `xml:"ConsumerReference,omitempty" json:"ConsumerReference,omitempty" yaml:"ConsumerReference,omitempty"`
	Filter             FilterType             `xml:"Filter,omitempty" json:"Filter,omitempty" yaml:"Filter,omitempty"`
	SubscriptionPolicy SubscriptionPolicyType `xml:"SubscriptionPolicy,omitempty" json:"SubscriptionPolicy,omitempty" yaml:"SubscriptionPolicy,omitempty"`
	CreationTime       *common.DateTime       `xml:"CreationTime,omitempty" json:"CreationTime,omitempty" yaml:"CreationTime,omitempty"`
}

type SubscriptionPolicyType []interface{}

type TopicExpressionDialectUnknownFaultType struct {
	Timestamp     common.DateTime       `xml:"Timestamp" json:"Timestamp" yaml:"Timestamp"`
	Originator    EndpointReferenceType `xml:"Originator,omitempty" json:"Originator,omitempty" yaml:"Originator,omitempty"`
	ErrorCode     string                `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	Description   []string              `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	FaultCause    string                `xml:"FaultCause" json:"FaultCause" yaml:"FaultCause"`
	TypeAttrXSI   string                `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string                `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *TopicExpressionDialectUnknownFaultType) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:TopicExpressionDialectUnknownFaultType"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://docs.oasis-open.org/wsn/b-2"
	}
}

type TopicExpressionType []interface{}

type TopicNamespaceType struct {
	Documentation   Documentation `xml:"documentation,omitempty" json:"documentation,omitempty" yaml:"documentation,omitempty"`
	Name            common.NCName `xml:"name,attr,omitempty" json:"name,attr,omitempty" yaml:"name,attr,omitempty"`
	TargetNamespace string        `xml:"targetNamespace,attr,omitempty" json:"targetNamespace,attr,omitempty" yaml:"targetNamespace,attr,omitempty"`
	Final           bool          `xml:"final,attr,omitempty" json:"final,attr,omitempty" yaml:"final,attr,omitempty"`
	Topic           []string      `xml:"Topic,omitempty" json:"Topic,omitempty" yaml:"Topic,omitempty"`
	TypeAttrXSI     string        `xml:"xsi:type,attr,omitempty"`
	TypeNamespace   string        `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *TopicNamespaceType) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:TopicNamespaceType"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://docs.oasis-open.org/wsn/t-1"
	}
}

type TopicNotSupportedFaultType struct {
	Timestamp     common.DateTime       `xml:"Timestamp" json:"Timestamp" yaml:"Timestamp"`
	Originator    EndpointReferenceType `xml:"Originator,omitempty" json:"Originator,omitempty" yaml:"Originator,omitempty"`
	ErrorCode     string                `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	Description   []string              `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	FaultCause    string                `xml:"FaultCause" json:"FaultCause" yaml:"FaultCause"`
	TypeAttrXSI   string                `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string                `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *TopicNotSupportedFaultType) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:TopicNotSupportedFaultType"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://docs.oasis-open.org/wsn/b-2"
	}
}

type TopicSetType struct {
	Documentation Documentation `xml:"documentation,omitempty" json:"documentation,omitempty" yaml:"documentation,omitempty"`
	TypeAttrXSI   string        `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string        `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *TopicSetType) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:TopicSetType"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://docs.oasis-open.org/wsn/t-1"
	}
}

type TopicType struct {
	Documentation  Documentation       `xml:"documentation,omitempty" json:"documentation,omitempty" yaml:"documentation,omitempty"`
	Name           common.NCName       `xml:"name,attr,omitempty" json:"name,attr,omitempty" yaml:"name,attr,omitempty"`
	MessageTypes   string              `xml:"messageTypes,attr,omitempty" json:"messageTypes,attr,omitempty" yaml:"messageTypes,attr,omitempty"`
	Final          bool                `xml:"final,attr,omitempty" json:"final,attr,omitempty" yaml:"final,attr,omitempty"`
	MessagePattern QueryExpressionType `xml:"MessagePattern,omitempty" json:"MessagePattern,omitempty" yaml:"MessagePattern,omitempty"`
	Topic          []TopicType         `xml:"Topic,omitempty" json:"Topic,omitempty" yaml:"Topic,omitempty"`
	TypeAttrXSI    string              `xml:"xsi:type,attr,omitempty"`
	TypeNamespace  string              `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *TopicType) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:TopicType"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://docs.oasis-open.org/wsn/t-1"
	}
}

type UnableToCreatePullPointFaultType struct {
	Timestamp     common.DateTime       `xml:"Timestamp" json:"Timestamp" yaml:"Timestamp"`
	Originator    EndpointReferenceType `xml:"Originator,omitempty" json:"Originator,omitempty" yaml:"Originator,omitempty"`
	ErrorCode     string                `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	Description   []string              `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	FaultCause    string                `xml:"FaultCause" json:"FaultCause" yaml:"FaultCause"`
	TypeAttrXSI   string                `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string                `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *UnableToCreatePullPointFaultType) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:UnableToCreatePullPointFaultType"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://docs.oasis-open.org/wsn/b-2"
	}
}

type UnableToDestroyPullPointFaultType struct {
	Timestamp     common.DateTime       `xml:"Timestamp" json:"Timestamp" yaml:"Timestamp"`
	Originator    EndpointReferenceType `xml:"Originator,omitempty" json:"Originator,omitempty" yaml:"Originator,omitempty"`
	ErrorCode     string                `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	Description   []string              `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	FaultCause    string                `xml:"FaultCause" json:"FaultCause" yaml:"FaultCause"`
	TypeAttrXSI   string                `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string                `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *UnableToDestroyPullPointFaultType) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:UnableToDestroyPullPointFaultType"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://docs.oasis-open.org/wsn/b-2"
	}
}

type UnableToDestroySubscriptionFaultType struct {
	Timestamp     common.DateTime       `xml:"Timestamp" json:"Timestamp" yaml:"Timestamp"`
	Originator    EndpointReferenceType `xml:"Originator,omitempty" json:"Originator,omitempty" yaml:"Originator,omitempty"`
	ErrorCode     string                `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	Description   []string              `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	FaultCause    string                `xml:"FaultCause" json:"FaultCause" yaml:"FaultCause"`
	TypeAttrXSI   string                `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string                `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *UnableToDestroySubscriptionFaultType) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:UnableToDestroySubscriptionFaultType"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://docs.oasis-open.org/wsn/b-2"
	}
}

type UnableToGetMessagesFaultType struct {
	Timestamp     common.DateTime       `xml:"Timestamp" json:"Timestamp" yaml:"Timestamp"`
	Originator    EndpointReferenceType `xml:"Originator,omitempty" json:"Originator,omitempty" yaml:"Originator,omitempty"`
	ErrorCode     string                `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	Description   []string              `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	FaultCause    string                `xml:"FaultCause" json:"FaultCause" yaml:"FaultCause"`
	TypeAttrXSI   string                `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string                `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *UnableToGetMessagesFaultType) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:UnableToGetMessagesFaultType"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://docs.oasis-open.org/wsn/b-2"
	}
}

type UnacceptableInitialTerminationTimeFaultType struct {
	Timestamp     common.DateTime       `xml:"Timestamp" json:"Timestamp" yaml:"Timestamp"`
	Originator    EndpointReferenceType `xml:"Originator,omitempty" json:"Originator,omitempty" yaml:"Originator,omitempty"`
	ErrorCode     string                `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	Description   []string              `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	FaultCause    string                `xml:"FaultCause" json:"FaultCause" yaml:"FaultCause"`
	MinimumTime   *common.DateTime      `xml:"MinimumTime,omitempty" json:"MinimumTime,omitempty" yaml:"MinimumTime,omitempty"`
	MaximumTime   *common.DateTime      `xml:"MaximumTime,omitempty" json:"MaximumTime,omitempty" yaml:"MaximumTime,omitempty"`
	TypeAttrXSI   string                `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string                `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *UnacceptableInitialTerminationTimeFaultType) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:UnacceptableInitialTerminationTimeFaultType"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://docs.oasis-open.org/wsn/b-2"
	}
}

type UnacceptableTerminationTimeFaultType struct {
	Timestamp     common.DateTime       `xml:"Timestamp" json:"Timestamp" yaml:"Timestamp"`
	Originator    EndpointReferenceType `xml:"Originator,omitempty" json:"Originator,omitempty" yaml:"Originator,omitempty"`
	ErrorCode     string                `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	Description   []string              `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	FaultCause    string                `xml:"FaultCause" json:"FaultCause" yaml:"FaultCause"`
	MinimumTime   *common.DateTime      `xml:"MinimumTime,omitempty" json:"MinimumTime,omitempty" yaml:"MinimumTime,omitempty"`
	MaximumTime   *common.DateTime      `xml:"MaximumTime,omitempty" json:"MaximumTime,omitempty" yaml:"MaximumTime,omitempty"`
	TypeAttrXSI   string                `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string                `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *UnacceptableTerminationTimeFaultType) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:UnacceptableTerminationTimeFaultType"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://docs.oasis-open.org/wsn/b-2"
	}
}

type UnrecognizedPolicyRequestFaultType struct {
	Timestamp          common.DateTime       `xml:"Timestamp" json:"Timestamp" yaml:"Timestamp"`
	Originator         EndpointReferenceType `xml:"Originator,omitempty" json:"Originator,omitempty" yaml:"Originator,omitempty"`
	ErrorCode          string                `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	Description        []string              `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	FaultCause         string                `xml:"FaultCause" json:"FaultCause" yaml:"FaultCause"`
	UnrecognizedPolicy []string              `xml:"UnrecognizedPolicy,omitempty" json:"UnrecognizedPolicy,omitempty" yaml:"UnrecognizedPolicy,omitempty"`
	TypeAttrXSI        string                `xml:"xsi:type,attr,omitempty"`
	TypeNamespace      string                `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *UnrecognizedPolicyRequestFaultType) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:UnrecognizedPolicyRequestFaultType"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://docs.oasis-open.org/wsn/b-2"
	}
}

type Unsubscribe []interface{}

type UnsubscribeResponse []interface{}

type UnsupportedPolicyRequestFaultType struct {
	Timestamp         common.DateTime       `xml:"Timestamp" json:"Timestamp" yaml:"Timestamp"`
	Originator        EndpointReferenceType `xml:"Originator,omitempty" json:"Originator,omitempty" yaml:"Originator,omitempty"`
	ErrorCode         string                `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	Description       []string              `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	FaultCause        string                `xml:"FaultCause" json:"FaultCause" yaml:"FaultCause"`
	UnsupportedPolicy []string              `xml:"UnsupportedPolicy,omitempty" json:"UnsupportedPolicy,omitempty" yaml:"UnsupportedPolicy,omitempty"`
	TypeAttrXSI       string                `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string                `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *UnsupportedPolicyRequestFaultType) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:UnsupportedPolicyRequestFaultType"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://docs.oasis-open.org/wsn/b-2"
	}
}

type UseRaw struct {
}

// pausableSubscriptionManager implements the PausableSubscriptionManager interface.
type pausableSubscriptionManager struct {
	cli      common.Client
	Endpoint string
}

func (p *pausableSubscriptionManager) OptAddEventBroker(args AddEventBroker) (*AddEventBrokerResponse, error) {
	req := struct {
		XMLName        string `xml:"tev:AddEventBroker"`
		AddEventBroker AddEventBroker
	}{
		AddEventBroker: args,
	}

	resp := AddEventBrokerResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *pausableSubscriptionManager) OptCreatePullPoint(args CreatePullPoint) (*CreatePullPointResponse, error) {
	req := struct {
		XMLName         string `xml:"wsntw:CreatePullPoint"`
		CreatePullPoint CreatePullPoint
	}{
		CreatePullPoint: args,
	}

	resp := CreatePullPointResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *pausableSubscriptionManager) OptCreatePullPointSubscription(args CreatePullPointSubscription) (*CreatePullPointSubscriptionResponse, error) {
	req := struct {
		XMLName                     string `xml:"tev:CreatePullPointSubscription"`
		CreatePullPointSubscription CreatePullPointSubscription
	}{
		CreatePullPointSubscription: args,
	}

	resp := CreatePullPointSubscriptionResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *pausableSubscriptionManager) OptDeleteEventBroker(args DeleteEventBroker) (*DeleteEventBrokerResponse, error) {
	req := struct {
		XMLName           string `xml:"tev:DeleteEventBroker"`
		DeleteEventBroker DeleteEventBroker
	}{
		DeleteEventBroker: args,
	}

	resp := DeleteEventBrokerResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *pausableSubscriptionManager) OptDestroyPullPoint(args DestroyPullPoint) (*DestroyPullPointResponse, error) {
	req := struct {
		XMLName          string `xml:"wsntw:DestroyPullPoint"`
		DestroyPullPoint DestroyPullPoint
	}{
		DestroyPullPoint: args,
	}

	resp := DestroyPullPointResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *pausableSubscriptionManager) OptGetCurrentMessage(args GetCurrentMessage) (*GetCurrentMessageResponse, error) {
	req := struct {
		XMLName           string `xml:"wsntw:GetCurrentMessage"`
		GetCurrentMessage GetCurrentMessage
	}{
		GetCurrentMessage: args,
	}

	resp := GetCurrentMessageResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *pausableSubscriptionManager) OptGetEventBrokers(args GetEventBrokers) (*GetEventBrokersResponse, error) {
	req := struct {
		XMLName         string `xml:"tev:GetEventBrokers"`
		GetEventBrokers GetEventBrokers
	}{
		GetEventBrokers: args,
	}

	resp := GetEventBrokersResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *pausableSubscriptionManager) OptGetEventProperties(args GetEventProperties) (*GetEventPropertiesResponse, error) {
	req := struct {
		XMLName            string `xml:"tev:GetEventProperties"`
		GetEventProperties GetEventProperties
	}{
		GetEventProperties: args,
	}

	resp := GetEventPropertiesResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *pausableSubscriptionManager) OptGetMessages(args GetMessages) (*GetMessagesResponse, error) {
	req := struct {
		XMLName     string `xml:"wsntw:GetMessages"`
		GetMessages GetMessages
	}{
		GetMessages: args,
	}

	resp := GetMessagesResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *pausableSubscriptionManager) OptGetServiceCapabilities(args GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error) {
	req := struct {
		XMLName                string `xml:"tev:GetServiceCapabilities"`
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

func (p *pausableSubscriptionManager) OptNotify(args Notify) error {
	req := struct {
		XMLName string `xml:"wsntw:Notify"`
		Notify  Notify
	}{
		Notify: args,
	}

	resp := map[string]interface{}{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return err
	}
	return nil
}

func (p *pausableSubscriptionManager) OptPauseSubscription(args PauseSubscription) (*PauseSubscriptionResponse, error) {
	req := struct {
		XMLName           string `xml:"wsntw:PauseSubscription"`
		PauseSubscription PauseSubscription
	}{
		PauseSubscription: args,
	}

	resp := PauseSubscriptionResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *pausableSubscriptionManager) OptPullMessages(args PullMessages) (*PullMessagesResponse, error) {
	req := struct {
		XMLName      string `xml:"tev:PullMessages"`
		PullMessages PullMessages
	}{
		PullMessages: args,
	}

	resp := PullMessagesResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *pausableSubscriptionManager) OptRenew(args Renew) (*RenewResponse, error) {
	req := struct {
		XMLName string `xml:"wsntw:Renew"`
		Renew   Renew
	}{
		Renew: args,
	}

	resp := RenewResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *pausableSubscriptionManager) OptResumeSubscription(args ResumeSubscription) (*ResumeSubscriptionResponse, error) {
	req := struct {
		XMLName            string `xml:"wsntw:ResumeSubscription"`
		ResumeSubscription ResumeSubscription
	}{
		ResumeSubscription: args,
	}

	resp := ResumeSubscriptionResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *pausableSubscriptionManager) OptSeek(args Seek) (*SeekResponse, error) {
	req := struct {
		XMLName string `xml:"tev:Seek"`
		Seek    Seek
	}{
		Seek: args,
	}

	resp := SeekResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *pausableSubscriptionManager) OptSetSynchronizationPoint(args SetSynchronizationPoint) (*SetSynchronizationPointResponse, error) {
	req := struct {
		XMLName                 string `xml:"tev:SetSynchronizationPoint"`
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

func (p *pausableSubscriptionManager) OptSubscribe(args Subscribe) (*SubscribeResponse, error) {
	req := struct {
		XMLName   string `xml:"wsntw:Subscribe"`
		Subscribe Subscribe
	}{
		Subscribe: args,
	}

	resp := SubscribeResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p *pausableSubscriptionManager) OptUnsubscribe(args Unsubscribe) (*UnsubscribeResponse, error) {
	req := struct {
		XMLName     string `xml:"wsntw:Unsubscribe"`
		Unsubscribe Unsubscribe
	}{
		Unsubscribe: args,
	}

	resp := UnsubscribeResponse{}

	if err := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
