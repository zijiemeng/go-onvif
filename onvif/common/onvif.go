package common

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// Duration in WSDL format.
type Duration string

type AFModes string

// Validate validates AFModes.
func (v AFModes) Validate() bool {
	for _, vv := range []string{
		"OnceAfterMove",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// AbsoluteOrRelativeTimeType is a union of: DateTime, Duration
type AbsoluteOrRelativeTimeType interface{}

type AudioClassType string

// Validate validates AudioClassType.
func (v AudioClassType) Validate() bool {
	for _, vv := range []string{
		"gun_shot",
		"scream",
		"glass_breaking",
		"tire_screech",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type AudioEncoding string

// Validate validates AudioEncoding.
func (v AudioEncoding) Validate() bool {
	for _, vv := range []string{
		"G711",
		"G726",
		"AAC",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type AudioEncodingMimeNames string

// Validate validates AudioEncodingMimeNames.
func (v AudioEncodingMimeNames) Validate() bool {
	for _, vv := range []string{
		"PCMU",
		"G726",
		"MP4A-LATM",
		"mpeg4-generic",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type AutoFocusMode string

// Validate validates AutoFocusMode.
func (v AutoFocusMode) Validate() bool {
	for _, vv := range []string{
		"AUTO",
		"MANUAL",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type AuxiliaryData string

type BacklightCompensationMode string

// Validate validates BacklightCompensationMode.
func (v BacklightCompensationMode) Validate() bool {
	for _, vv := range []string{
		"OFF",
		"ON",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type CapabilityCategory string

// Validate validates CapabilityCategory.
func (v CapabilityCategory) Validate() bool {
	for _, vv := range []string{
		"All",
		"Analytics",
		"Device",
		"Events",
		"Imaging",
		"Media",
		"PTZ",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type ConcreteTopicExpression string

type DNSName string

type DefoggingMode string

// Validate validates DefoggingMode.
func (v DefoggingMode) Validate() bool {
	for _, vv := range []string{
		"OFF",
		"ON",
		"AUTO",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type Description string

type DigitalIdleState string

// Validate validates DigitalIdleState.
func (v DigitalIdleState) Validate() bool {
	for _, vv := range []string{
		"closed",
		"open",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type Direction string

// Validate validates Direction.
func (v Direction) Validate() bool {
	for _, vv := range []string{
		"Left",
		"Right",
		"Any",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type DiscoveryMode string

// Validate validates DiscoveryMode.
func (v DiscoveryMode) Validate() bool {
	for _, vv := range []string{
		"Discoverable",
		"NonDiscoverable",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type Domain string

type Dot11AuthAndMangementSuite string

// Validate validates Dot11AuthAndMangementSuite.
func (v Dot11AuthAndMangementSuite) Validate() bool {
	for _, vv := range []string{
		"None",
		"Dot1X",
		"PSK",
		"Extended",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type Dot11Cipher string

// Validate validates Dot11Cipher.
func (v Dot11Cipher) Validate() bool {
	for _, vv := range []string{
		"CCMP",
		"TKIP",
		"Any",
		"Extended",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type Dot11PSK HexBinary

type Dot11PSKPassphrase string

type Dot11SSIDType HexBinary

type Dot11SecurityMode string

// Validate validates Dot11SecurityMode.
func (v Dot11SecurityMode) Validate() bool {
	for _, vv := range []string{
		"None",
		"WEP",
		"PSK",
		"Dot1X",
		"Extended",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type Dot11SignalStrength string

// Validate validates Dot11SignalStrength.
func (v Dot11SignalStrength) Validate() bool {
	for _, vv := range []string{
		"None",
		"Very Bad",
		"Bad",
		"Good",
		"Very Good",
		"Extended",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type Dot11StationMode string

// Validate validates Dot11StationMode.
func (v Dot11StationMode) Validate() bool {
	for _, vv := range []string{
		"Ad-hoc",
		"Infrastructure",
		"Extended",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type Duplex string

// Validate validates Duplex.
func (v Duplex) Validate() bool {
	for _, vv := range []string{
		"Full",
		"Half",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type DynamicDNSType string

// Validate validates DynamicDNSType.
func (v DynamicDNSType) Validate() bool {
	for _, vv := range []string{
		"NoUpdate",
		"ClientUpdates",
		"ServerUpdates",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type EFlipMode string

// Validate validates EFlipMode.
func (v EFlipMode) Validate() bool {
	for _, vv := range []string{
		"OFF",
		"ON",
		"Extended",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type Enabled string

// Validate validates Enabled.
func (v Enabled) Validate() bool {
	for _, vv := range []string{
		"ENABLED",
		"DISABLED",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type ExposureMode string

// Validate validates ExposureMode.
func (v ExposureMode) Validate() bool {
	for _, vv := range []string{
		"AUTO",
		"MANUAL",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type ExposurePriority string

// Validate validates ExposurePriority.
func (v ExposurePriority) Validate() bool {
	for _, vv := range []string{
		"LowNoise",
		"FrameRate",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type FactoryDefaultType string

// Validate validates FactoryDefaultType.
func (v FactoryDefaultType) Validate() bool {
	for _, vv := range []string{
		"Hard",
		"Soft",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// FaultCodesOpenEnumType is a union of: FaultCodesType, string
type FaultCodesOpenEnumType interface{}

type FaultCodesType string

// Validate validates FaultCodesType.
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

type H264Profile string

// Validate validates H264Profile.
func (v H264Profile) Validate() bool {
	for _, vv := range []string{
		"Baseline",
		"Main",
		"Extended",
		"High",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type HwAddress string

type IANAIfTypes int

type IPAddressFilterType string

// Validate validates IPAddressFilterType.
func (v IPAddressFilterType) Validate() bool {
	for _, vv := range []string{
		"Allow",
		"Deny",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type IPType string

// Validate validates IPType.
func (v IPType) Validate() bool {
	for _, vv := range []string{
		"IPv4",
		"IPv6",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type IPv4Address string

type IPv6Address string

type IPv6DHCPConfiguration string

// Validate validates IPv6DHCPConfiguration.
func (v IPv6DHCPConfiguration) Validate() bool {
	for _, vv := range []string{
		"Auto",
		"Stateful",
		"Stateless",
		"Off",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type ImageSendingType string

// Validate validates ImageSendingType.
func (v ImageSendingType) Validate() bool {
	for _, vv := range []string{
		"Embedded",
		"LocalStorage",
		"RemoteStorage",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type ImageStabilizationMode string

// Validate validates ImageStabilizationMode.
func (v ImageStabilizationMode) Validate() bool {
	for _, vv := range []string{
		"OFF",
		"ON",
		"AUTO",
		"Extended",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type IrCutFilterAutoBoundaryType string

// Validate validates IrCutFilterAutoBoundaryType.
func (v IrCutFilterAutoBoundaryType) Validate() bool {
	for _, vv := range []string{
		"Common",
		"ToOn",
		"ToOff",
		"Extended",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type IrCutFilterMode string

// Validate validates IrCutFilterMode.
func (v IrCutFilterMode) Validate() bool {
	for _, vv := range []string{
		"ON",
		"OFF",
		"AUTO",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type JobToken ReferenceToken

type MetadataCompressionType string

// Validate validates MetadataCompressionType.
func (v MetadataCompressionType) Validate() bool {
	for _, vv := range []string{
		"None",
		"GZIP",
		"EXI",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type ModeOfOperation string

// Validate validates ModeOfOperation.
func (v ModeOfOperation) Validate() bool {
	for _, vv := range []string{
		"Idle",
		"Active",
		"Unknown",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type MoveAndTrackMethod string

// Validate validates MoveAndTrackMethod.
func (v MoveAndTrackMethod) Validate() bool {
	for _, vv := range []string{
		"PresetToken",
		"GeoLocation",
		"PTZVector",
		"ObjectID",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type Mpeg4Profile string

// Validate validates Mpeg4Profile.
func (v Mpeg4Profile) Validate() bool {
	for _, vv := range []string{
		"SP",
		"ASP",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type Name string

type NetworkHostType string

// Validate validates NetworkHostType.
func (v NetworkHostType) Validate() bool {
	for _, vv := range []string{
		"IPv4",
		"IPv6",
		"DNS",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type NetworkInterfaceConfigPriority int64

type NetworkProtocolType string

// Validate validates NetworkProtocolType.
func (v NetworkProtocolType) Validate() bool {
	for _, vv := range []string{
		"HTTP",
		"HTTPS",
		"RTSP",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type OSDType string

// Validate validates OSDType.
func (v OSDType) Validate() bool {
	for _, vv := range []string{
		"Text",
		"Image",
		"Extended",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type PTZPresetTourDirection string

// Validate validates PTZPresetTourDirection.
func (v PTZPresetTourDirection) Validate() bool {
	for _, vv := range []string{
		"Forward",
		"Backward",
		"Extended",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type PTZPresetTourOperation string

// Validate validates PTZPresetTourOperation.
func (v PTZPresetTourOperation) Validate() bool {
	for _, vv := range []string{
		"Start",
		"Stop",
		"Pause",
		"Extended",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type PTZPresetTourState string

// Validate validates PTZPresetTourState.
func (v PTZPresetTourState) Validate() bool {
	for _, vv := range []string{
		"Idle",
		"Touring",
		"Paused",
		"Extended",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type PropertyOperation string

// Validate validates PropertyOperation.
func (v PropertyOperation) Validate() bool {
	for _, vv := range []string{
		"Initialized",
		"Deleted",
		"Changed",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type ReceiverMode string

// Validate validates ReceiverMode.
func (v ReceiverMode) Validate() bool {
	for _, vv := range []string{
		"AutoConnect",
		"AlwaysConnect",
		"NeverConnect",
		"Unknown",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type ReceiverReference ReferenceToken

type ReceiverState string

// Validate validates ReceiverState.
func (v ReceiverState) Validate() bool {
	for _, vv := range []string{
		"NotConnected",
		"Connecting",
		"Connected",
		"Unknown",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type RecordingJobMode string

type RecordingJobReference ReferenceToken

type RecordingJobState string

type RecordingReference ReferenceToken

type RecordingStatus string

// Validate validates RecordingStatus.
func (v RecordingStatus) Validate() bool {
	for _, vv := range []string{
		"Initiated",
		"Recording",
		"Stopped",
		"Removing",
		"Removed",
		"Unknown",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type RelationshipType string

// Validate validates RelationshipType.
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

// RelationshipTypeOpenEnum is a union of: RelationshipType, string
type RelationshipTypeOpenEnum interface{}

type RelayIdleState string

// Validate validates RelayIdleState.
func (v RelayIdleState) Validate() bool {
	for _, vv := range []string{
		"closed",
		"open",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type RelayLogicalState string

// Validate validates RelayLogicalState.
func (v RelayLogicalState) Validate() bool {
	for _, vv := range []string{
		"active",
		"inactive",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type RelayMode string

// Validate validates RelayMode.
func (v RelayMode) Validate() bool {
	for _, vv := range []string{
		"Monostable",
		"Bistable",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type ReverseMode string

// Validate validates ReverseMode.
func (v ReverseMode) Validate() bool {
	for _, vv := range []string{
		"OFF",
		"ON",
		"AUTO",
		"Extended",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type RotateMode string

// Validate validates RotateMode.
func (v RotateMode) Validate() bool {
	for _, vv := range []string{
		"OFF",
		"ON",
		"AUTO",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type SceneOrientationMode string

// Validate validates SceneOrientationMode.
func (v SceneOrientationMode) Validate() bool {
	for _, vv := range []string{
		"MANUAL",
		"AUTO",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type SceneOrientationOption string

// Validate validates SceneOrientationOption.
func (v SceneOrientationOption) Validate() bool {
	for _, vv := range []string{
		"Below",
		"Horizon",
		"Above",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type ScopeDefinition string

// Validate validates ScopeDefinition.
func (v ScopeDefinition) Validate() bool {
	for _, vv := range []string{
		"Fixed",
		"Configurable",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type SearchState string

// Validate validates SearchState.
func (v SearchState) Validate() bool {
	for _, vv := range []string{
		"Queued",
		"Searching",
		"Completed",
		"Unknown",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type SetDateTimeType string

// Validate validates SetDateTimeType.
func (v SetDateTimeType) Validate() bool {
	for _, vv := range []string{
		"Manual",
		"NTP",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type SimpleTopicExpression string

type StreamType string

// Validate validates StreamType.
func (v StreamType) Validate() bool {
	for _, vv := range []string{
		"RTP-Unicast",
		"RTP-Multicast",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type SystemLogType string

// Validate validates SystemLogType.
func (v SystemLogType) Validate() bool {
	for _, vv := range []string{
		"System",
		"Access",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type ToneCompensationMode string

// Validate validates ToneCompensationMode.
func (v ToneCompensationMode) Validate() bool {
	for _, vv := range []string{
		"OFF",
		"ON",
		"AUTO",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type TrackReference ReferenceToken

type TrackType string

// Validate validates TrackType.
func (v TrackType) Validate() bool {
	for _, vv := range []string{
		"Video",
		"Audio",
		"Metadata",
		"Extended",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type TransportProtocol string

// Validate validates TransportProtocol.
func (v TransportProtocol) Validate() bool {
	for _, vv := range []string{
		"UDP",
		"TCP",
		"RTSP",
		"HTTP",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type UserLevel string

// Validate validates UserLevel.
func (v UserLevel) Validate() bool {
	for _, vv := range []string{
		"Administrator",
		"Operator",
		"User",
		"Anonymous",
		"Extended",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type VideoEncoding string

// Validate validates VideoEncoding.
func (v VideoEncoding) Validate() bool {
	for _, vv := range []string{
		"JPEG",
		"MPEG4",
		"H264",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type VideoEncodingMimeNames string

// Validate validates VideoEncodingMimeNames.
func (v VideoEncodingMimeNames) Validate() bool {
	for _, vv := range []string{
		"JPEG",
		"MPV4-ES",
		"H264",
		"H265",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type VideoEncodingProfiles string

// Validate validates VideoEncodingProfiles.
func (v VideoEncodingProfiles) Validate() bool {
	for _, vv := range []string{
		"Simple",
		"AdvancedSimple",
		"Baseline",
		"Main",
		"Main10",
		"Extended",
		"High",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type ViewModes string

// Validate validates ViewModes.
func (v ViewModes) Validate() bool {
	for _, vv := range []string{
		"tt:Fisheye",
		"tt:360Panorama",
		"tt:180Panorama",
		"tt:Quad",
		"tt:Original",
		"tt:LeftHalf",
		"tt:RightHalf",
		"tt:Dewarp",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type WhiteBalanceMode string

// Validate validates WhiteBalanceMode.
func (v WhiteBalanceMode) Validate() bool {
	for _, vv := range []string{
		"AUTO",
		"MANUAL",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type WideDynamicMode string

// Validate validates WideDynamicMode.
func (v WideDynamicMode) Validate() bool {
	for _, vv := range []string{
		"OFF",
		"ON",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type XPathExpression string

type FaultcodeEnum string

// Validate validates FaultcodeEnum.
func (v FaultcodeEnum) Validate() bool {
	for _, vv := range []string{
		"tns:DataEncodingUnknown",
		"tns:MustUnderstand",
		"tns:Receiver",
		"tns:Sender",
		"tns:VersionMismatch",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type AACDecOptions struct {
	Bitrate         *IntItems `xml:"Bitrate,omitempty" json:"Bitrate,omitempty" yaml:"Bitrate,omitempty"`
	SampleRateRange *IntItems `xml:"SampleRateRange,omitempty" json:"SampleRateRange,omitempty" yaml:"SampleRateRange,omitempty"`
}

type AbsoluteFocus struct {
	Position *float64 `xml:"Position,omitempty" json:"Position,omitempty" yaml:"Position,omitempty"`
	Speed    *float64 `xml:"Speed,omitempty" json:"Speed,omitempty" yaml:"Speed,omitempty"`
}

type AbsoluteFocusOptions struct {
	Position *FloatRange `xml:"Position,omitempty" json:"Position,omitempty" yaml:"Position,omitempty"`
	Speed    *FloatRange `xml:"Speed,omitempty" json:"Speed,omitempty" yaml:"Speed,omitempty"`
}

// Action Engine Event Payload data structure contains the
//                     information about the ONVIF command invocations.
// Since this event could be                         generated
// by other or proprietary actions, the command invocation specific
//                         fields are defined as optional and additional
// extension mechanism is                         provided for
// future or additional action definitions.
type ActionEngineEventPayload struct {
	RequestInfo  *Envelope                          `xml:"RequestInfo,omitempty" json:"RequestInfo,omitempty" yaml:"RequestInfo,omitempty"`
	ResponseInfo *Envelope                          `xml:"ResponseInfo,omitempty" json:"ResponseInfo,omitempty" yaml:"ResponseInfo,omitempty"`
	Fault        *Fault                             `xml:"Fault,omitempty" json:"Fault,omitempty" yaml:"Fault,omitempty"`
	Extension    *ActionEngineEventPayloadExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type ActionEngineEventPayloadExtension []interface{}

type ActiveConnection struct {
	CurrentBitrate *float64 `xml:"CurrentBitrate,omitempty" json:"CurrentBitrate,omitempty" yaml:"CurrentBitrate,omitempty"`
	CurrentFps     *float64 `xml:"CurrentFps,omitempty" json:"CurrentFps,omitempty" yaml:"CurrentFps,omitempty"`
}

type AnalyticsCapabilities struct {
	XAddr                  *string `xml:"XAddr,omitempty" json:"XAddr,omitempty" yaml:"XAddr,omitempty"`
	RuleSupport            *bool   `xml:"RuleSupport,omitempty" json:"RuleSupport,omitempty" yaml:"RuleSupport,omitempty"`
	AnalyticsModuleSupport *bool   `xml:"AnalyticsModuleSupport,omitempty" json:"AnalyticsModuleSupport,omitempty" yaml:"AnalyticsModuleSupport,omitempty"`
}

type AnalyticsDeviceCapabilities struct {
	XAddr       *string                   `xml:"XAddr,omitempty" json:"XAddr,omitempty" yaml:"XAddr,omitempty"`
	RuleSupport *bool                     `xml:"RuleSupport,omitempty" json:"RuleSupport,omitempty" yaml:"RuleSupport,omitempty"`
	Extension   *AnalyticsDeviceExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type AnalyticsDeviceEngineConfiguration struct {
	EngineConfiguration []*EngineConfiguration                       `xml:"EngineConfiguration,omitempty" json:"EngineConfiguration,omitempty" yaml:"EngineConfiguration,omitempty"`
	Extension           *AnalyticsDeviceEngineConfigurationExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type AnalyticsDeviceEngineConfigurationExtension []interface{}

type AnalyticsDeviceExtension []interface{}

type AnalyticsEngine struct {
	Name                         *Name                               `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	UseCount                     *int                                `xml:"UseCount,omitempty" json:"UseCount,omitempty" yaml:"UseCount,omitempty"`
	Token                        ReferenceToken                      `xml:"token,attr,omitempty" json:"token,attr,omitempty" yaml:"token,attr,omitempty"`
	AnalyticsEngineConfiguration *AnalyticsDeviceEngineConfiguration `xml:"AnalyticsEngineConfiguration,omitempty" json:"AnalyticsEngineConfiguration,omitempty" yaml:"AnalyticsEngineConfiguration,omitempty"`
	TypeAttrXSI                  string                              `xml:"xsi:type,attr,omitempty"`
	TypeNamespace                string                              `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *AnalyticsEngine) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:AnalyticsEngine"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://www.onvif.org/ver10/schema"
	}
}

type AnalyticsEngineConfiguration struct {
	AnalyticsModule []*Config                              `xml:"AnalyticsModule,omitempty" json:"AnalyticsModule,omitempty" yaml:"AnalyticsModule,omitempty"`
	Extension       *AnalyticsEngineConfigurationExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type AnalyticsEngineConfigurationExtension []interface{}

type AnalyticsEngineControl struct {
	Name              *Name                   `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	UseCount          *int                    `xml:"UseCount,omitempty" json:"UseCount,omitempty" yaml:"UseCount,omitempty"`
	Token             ReferenceToken          `xml:"token,attr,omitempty" json:"token,attr,omitempty" yaml:"token,attr,omitempty"`
	EngineToken       ReferenceToken          `xml:"EngineToken,omitempty" json:"EngineToken,omitempty" yaml:"EngineToken,omitempty"`
	EngineConfigToken ReferenceToken          `xml:"EngineConfigToken,omitempty" json:"EngineConfigToken,omitempty" yaml:"EngineConfigToken,omitempty"`
	InputToken        []ReferenceToken        `xml:"InputToken,omitempty" json:"InputToken,omitempty" yaml:"InputToken,omitempty"`
	ReceiverToken     []ReferenceToken        `xml:"ReceiverToken,omitempty" json:"ReceiverToken,omitempty" yaml:"ReceiverToken,omitempty"`
	Multicast         *MulticastConfiguration `xml:"Multicast,omitempty" json:"Multicast,omitempty" yaml:"Multicast,omitempty"`
	Subscription      *Config                 `xml:"Subscription,omitempty" json:"Subscription,omitempty" yaml:"Subscription,omitempty"`
	Mode              *ModeOfOperation        `xml:"Mode,omitempty" json:"Mode,omitempty" yaml:"Mode,omitempty"`
	TypeAttrXSI       string                  `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string                  `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *AnalyticsEngineControl) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:AnalyticsEngineControl"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://www.onvif.org/ver10/schema"
	}
}

type AnalyticsEngineInput struct {
	Name                 *Name                      `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	UseCount             *int                       `xml:"UseCount,omitempty" json:"UseCount,omitempty" yaml:"UseCount,omitempty"`
	Token                ReferenceToken             `xml:"token,attr,omitempty" json:"token,attr,omitempty" yaml:"token,attr,omitempty"`
	SourceIdentification *SourceIdentification      `xml:"SourceIdentification,omitempty" json:"SourceIdentification,omitempty" yaml:"SourceIdentification,omitempty"`
	VideoInput           *VideoEncoderConfiguration `xml:"VideoInput,omitempty" json:"VideoInput,omitempty" yaml:"VideoInput,omitempty"`
	MetadataInput        *MetadataInput             `xml:"MetadataInput,omitempty" json:"MetadataInput,omitempty" yaml:"MetadataInput,omitempty"`
	TypeAttrXSI          string                     `xml:"xsi:type,attr,omitempty"`
	TypeNamespace        string                     `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *AnalyticsEngineInput) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:AnalyticsEngineInput"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://www.onvif.org/ver10/schema"
	}
}

type AnalyticsEngineInputInfo struct {
	InputInfo *Config                            `xml:"InputInfo,omitempty" json:"InputInfo,omitempty" yaml:"InputInfo,omitempty"`
	Extension *AnalyticsEngineInputInfoExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type AnalyticsEngineInputInfoExtension []interface{}

type AnalyticsState struct {
	Error *string `xml:"Error,omitempty" json:"Error,omitempty" yaml:"Error,omitempty"`
	State *string `xml:"State,omitempty" json:"State,omitempty" yaml:"State,omitempty"`
}

type AnalyticsStateInformation struct {
	AnalyticsEngineControlToken ReferenceToken  `xml:"AnalyticsEngineControlToken,omitempty" json:"AnalyticsEngineControlToken,omitempty" yaml:"AnalyticsEngineControlToken,omitempty"`
	State                       *AnalyticsState `xml:"State,omitempty" json:"State,omitempty" yaml:"State,omitempty"`
}

type AnyHolder []interface{}

type ArrayOfFileProgress struct {
	FileProgress []*FileProgress               `xml:"FileProgress,omitempty" json:"FileProgress,omitempty" yaml:"FileProgress,omitempty"`
	Extension    *ArrayOfFileProgressExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type ArrayOfFileProgressExtension []interface{}

type AttachmentData struct {
	Include     *Include `xml:"Include,omitempty" json:"Include,omitempty" yaml:"Include,omitempty"`
	ContentType string   `xml:"contentType,attr,omitempty" json:"contentType,attr,omitempty" yaml:"contentType,attr,omitempty"`
}

type AttributedAnyType []interface{}

type AttributedQNameType struct {
	Content *string `xml:"Content,omitempty" json:"Content,omitempty" yaml:"Content,omitempty"`
}

type AttributedURIType struct {
	Content *string `xml:"Content,omitempty" json:"Content,omitempty" yaml:"Content,omitempty"`
}

type AttributedUnsignedLongType struct {
	Content uint64 `xml:"Content,omitempty" json:"Content,omitempty" yaml:"Content,omitempty"`
}

type AudioAttributes struct {
	Bitrate    *int    `xml:"Bitrate,omitempty" json:"Bitrate,omitempty" yaml:"Bitrate,omitempty"`
	Encoding   *string `xml:"Encoding,omitempty" json:"Encoding,omitempty" yaml:"Encoding,omitempty"`
	Samplerate *int    `xml:"Samplerate,omitempty" json:"Samplerate,omitempty" yaml:"Samplerate,omitempty"`
}

type AudioClassCandidate struct {
	Type       *AudioClassType `xml:"Type,omitempty" json:"Type,omitempty" yaml:"Type,omitempty"`
	Likelihood *float64        `xml:"Likelihood,omitempty" json:"Likelihood,omitempty" yaml:"Likelihood,omitempty"`
}

type AudioClassDescriptor struct {
	ClassCandidate []*AudioClassCandidate         `xml:"ClassCandidate,omitempty" json:"ClassCandidate,omitempty" yaml:"ClassCandidate,omitempty"`
	Extension      *AudioClassDescriptorExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type AudioClassDescriptorExtension []interface{}

// The Audio Decoder Configuration does not contain any that
//                       parameter to configure the decoding .A
// decoder shall decode every data it                         receives
// (according to its capabilities).
type AudioDecoderConfiguration struct {
	Name          *Name          `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	UseCount      *int           `xml:"UseCount,omitempty" json:"UseCount,omitempty" yaml:"UseCount,omitempty"`
	Token         ReferenceToken `xml:"token,attr,omitempty" json:"token,attr,omitempty" yaml:"token,attr,omitempty"`
	TypeAttrXSI   string         `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string         `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *AudioDecoderConfiguration) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:AudioDecoderConfiguration"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://www.onvif.org/ver10/schema"
	}
}

type AudioDecoderConfigurationOptions struct {
	AACDecOptions  *AACDecOptions                             `xml:"AACDecOptions,omitempty" json:"AACDecOptions,omitempty" yaml:"AACDecOptions,omitempty"`
	G711DecOptions *G711DecOptions                            `xml:"G711DecOptions,omitempty" json:"G711DecOptions,omitempty" yaml:"G711DecOptions,omitempty"`
	G726DecOptions *G726DecOptions                            `xml:"G726DecOptions,omitempty" json:"G726DecOptions,omitempty" yaml:"G726DecOptions,omitempty"`
	Extension      *AudioDecoderConfigurationOptionsExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type AudioDecoderConfigurationOptionsExtension []interface{}

type AudioEncoder2Configuration struct {
	Name          *Name                   `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	UseCount      *int                    `xml:"UseCount,omitempty" json:"UseCount,omitempty" yaml:"UseCount,omitempty"`
	Token         ReferenceToken          `xml:"token,attr,omitempty" json:"token,attr,omitempty" yaml:"token,attr,omitempty"`
	Encoding      *string                 `xml:"Encoding,omitempty" json:"Encoding,omitempty" yaml:"Encoding,omitempty"`
	Multicast     *MulticastConfiguration `xml:"Multicast,omitempty" json:"Multicast,omitempty" yaml:"Multicast,omitempty"`
	Bitrate       *int                    `xml:"Bitrate,omitempty" json:"Bitrate,omitempty" yaml:"Bitrate,omitempty"`
	SampleRate    *int                    `xml:"SampleRate,omitempty" json:"SampleRate,omitempty" yaml:"SampleRate,omitempty"`
	TypeAttrXSI   string                  `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string                  `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *AudioEncoder2Configuration) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:AudioEncoder2Configuration"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://www.onvif.org/ver10/schema"
	}
}

type AudioEncoder2ConfigurationOptions struct {
	Encoding       *string   `xml:"Encoding,omitempty" json:"Encoding,omitempty" yaml:"Encoding,omitempty"`
	BitrateList    *IntItems `xml:"BitrateList,omitempty" json:"BitrateList,omitempty" yaml:"BitrateList,omitempty"`
	SampleRateList *IntItems `xml:"SampleRateList,omitempty" json:"SampleRateList,omitempty" yaml:"SampleRateList,omitempty"`
}

type AudioEncoderConfiguration struct {
	Name           *Name                   `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	UseCount       *int                    `xml:"UseCount,omitempty" json:"UseCount,omitempty" yaml:"UseCount,omitempty"`
	Token          ReferenceToken          `xml:"token,attr,omitempty" json:"token,attr,omitempty" yaml:"token,attr,omitempty"`
	Encoding       *AudioEncoding          `xml:"Encoding,omitempty" json:"Encoding,omitempty" yaml:"Encoding,omitempty"`
	Bitrate        *int                    `xml:"Bitrate,omitempty" json:"Bitrate,omitempty" yaml:"Bitrate,omitempty"`
	SampleRate     *int                    `xml:"SampleRate,omitempty" json:"SampleRate,omitempty" yaml:"SampleRate,omitempty"`
	Multicast      *MulticastConfiguration `xml:"Multicast,omitempty" json:"Multicast,omitempty" yaml:"Multicast,omitempty"`
	SessionTimeout *Duration               `xml:"SessionTimeout,omitempty" json:"SessionTimeout,omitempty" yaml:"SessionTimeout,omitempty"`
	TypeAttrXSI    string                  `xml:"xsi:type,attr,omitempty"`
	TypeNamespace  string                  `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *AudioEncoderConfiguration) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:AudioEncoderConfiguration"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://www.onvif.org/ver10/schema"
	}
}

type AudioEncoderConfigurationOption struct {
	Encoding       *AudioEncoding `xml:"Encoding,omitempty" json:"Encoding,omitempty" yaml:"Encoding,omitempty"`
	BitrateList    *IntItems      `xml:"BitrateList,omitempty" json:"BitrateList,omitempty" yaml:"BitrateList,omitempty"`
	SampleRateList *IntItems      `xml:"SampleRateList,omitempty" json:"SampleRateList,omitempty" yaml:"SampleRateList,omitempty"`
}

type AudioEncoderConfigurationOptions struct {
	Options []*AudioEncoderConfigurationOption `xml:"Options,omitempty" json:"Options,omitempty" yaml:"Options,omitempty"`
}

// Representation of a physical audio outputs.
type AudioOutput struct {
	Token         ReferenceToken `xml:"token,attr,omitempty" json:"token,attr,omitempty" yaml:"token,attr,omitempty"`
	TypeAttrXSI   string         `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string         `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *AudioOutput) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:AudioOutput"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://www.onvif.org/ver10/schema"
	}
}

type AudioOutputConfiguration struct {
	Name          *Name          `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	UseCount      *int           `xml:"UseCount,omitempty" json:"UseCount,omitempty" yaml:"UseCount,omitempty"`
	Token         ReferenceToken `xml:"token,attr,omitempty" json:"token,attr,omitempty" yaml:"token,attr,omitempty"`
	OutputToken   ReferenceToken `xml:"OutputToken,omitempty" json:"OutputToken,omitempty" yaml:"OutputToken,omitempty"`
	SendPrimacy   *string        `xml:"SendPrimacy,omitempty" json:"SendPrimacy,omitempty" yaml:"SendPrimacy,omitempty"`
	OutputLevel   *int           `xml:"OutputLevel,omitempty" json:"OutputLevel,omitempty" yaml:"OutputLevel,omitempty"`
	TypeAttrXSI   string         `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string         `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *AudioOutputConfiguration) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:AudioOutputConfiguration"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://www.onvif.org/ver10/schema"
	}
}

type AudioOutputConfigurationOptions struct {
	OutputTokensAvailable []ReferenceToken `xml:"OutputTokensAvailable,omitempty" json:"OutputTokensAvailable,omitempty" yaml:"OutputTokensAvailable,omitempty"`
	SendPrimacyOptions    []*string        `xml:"SendPrimacyOptions,omitempty" json:"SendPrimacyOptions,omitempty" yaml:"SendPrimacyOptions,omitempty"`
	OutputLevelRange      *IntRange        `xml:"OutputLevelRange,omitempty" json:"OutputLevelRange,omitempty" yaml:"OutputLevelRange,omitempty"`
}

// Representation of a physical audio input.
type AudioSource struct {
	Token         ReferenceToken `xml:"token,attr,omitempty" json:"token,attr,omitempty" yaml:"token,attr,omitempty"`
	Channels      *int           `xml:"Channels,omitempty" json:"Channels,omitempty" yaml:"Channels,omitempty"`
	TypeAttrXSI   string         `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string         `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *AudioSource) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:AudioSource"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://www.onvif.org/ver10/schema"
	}
}

type AudioSourceConfiguration struct {
	Name          *Name          `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	UseCount      *int           `xml:"UseCount,omitempty" json:"UseCount,omitempty" yaml:"UseCount,omitempty"`
	Token         ReferenceToken `xml:"token,attr,omitempty" json:"token,attr,omitempty" yaml:"token,attr,omitempty"`
	SourceToken   ReferenceToken `xml:"SourceToken,omitempty" json:"SourceToken,omitempty" yaml:"SourceToken,omitempty"`
	TypeAttrXSI   string         `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string         `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *AudioSourceConfiguration) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:AudioSourceConfiguration"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://www.onvif.org/ver10/schema"
	}
}

type AudioSourceConfigurationOptions struct {
	InputTokensAvailable []ReferenceToken             `xml:"InputTokensAvailable,omitempty" json:"InputTokensAvailable,omitempty" yaml:"InputTokensAvailable,omitempty"`
	Extension            *AudioSourceOptionsExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type AudioSourceOptionsExtension []interface{}

type BacklightCompensation struct {
	Mode  *BacklightCompensationMode `xml:"Mode,omitempty" json:"Mode,omitempty" yaml:"Mode,omitempty"`
	Level *float64                   `xml:"Level,omitempty" json:"Level,omitempty" yaml:"Level,omitempty"`
}

// Type describing whether BLC mode is enabled or disabled
//                     (on/off).
type BacklightCompensation20 struct {
	Mode  *BacklightCompensationMode `xml:"Mode,omitempty" json:"Mode,omitempty" yaml:"Mode,omitempty"`
	Level *float64                   `xml:"Level,omitempty" json:"Level,omitempty" yaml:"Level,omitempty"`
}

type BacklightCompensationOptions struct {
	Mode  []*WideDynamicMode `xml:"Mode,omitempty" json:"Mode,omitempty" yaml:"Mode,omitempty"`
	Level *FloatRange        `xml:"Level,omitempty" json:"Level,omitempty" yaml:"Level,omitempty"`
}

type BacklightCompensationOptions20 struct {
	Mode  []*BacklightCompensationMode `xml:"Mode,omitempty" json:"Mode,omitempty" yaml:"Mode,omitempty"`
	Level *FloatRange                  `xml:"Level,omitempty" json:"Level,omitempty" yaml:"Level,omitempty"`
}

type BackupFile struct {
	Name *string         `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Data *AttachmentData `xml:"Data,omitempty" json:"Data,omitempty" yaml:"Data,omitempty"`
}

type BaseFaultType struct {
	Timestamp   DateTime               `xml:"Timestamp" json:"Timestamp" yaml:"Timestamp"`
	Originator  *EndpointReferenceType `xml:"Originator,omitempty" json:"Originator,omitempty" yaml:"Originator,omitempty"`
	ErrorCode   *string                `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	Description []*string              `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	FaultCause  interface{}            `xml:"FaultCause" json:"FaultCause" yaml:"FaultCause"`
}

type BinaryData struct {
	Data        *Base64Binary `xml:"Data,omitempty" json:"Data,omitempty" yaml:"Data,omitempty"`
	ContentType string        `xml:"contentType,attr,omitempty" json:"contentType,attr,omitempty" yaml:"contentType,attr,omitempty"`
}

type Body []interface{}

type Capabilities struct {
	Analytics *AnalyticsCapabilities `xml:"Analytics,omitempty" json:"Analytics,omitempty" yaml:"Analytics,omitempty"`
	Device    *DeviceCapabilities    `xml:"Device,omitempty" json:"Device,omitempty" yaml:"Device,omitempty"`
	Events    *EventCapabilities     `xml:"Events,omitempty" json:"Events,omitempty" yaml:"Events,omitempty"`
	Imaging   *ImagingCapabilities   `xml:"Imaging,omitempty" json:"Imaging,omitempty" yaml:"Imaging,omitempty"`
	Media     *MediaCapabilities     `xml:"Media,omitempty" json:"Media,omitempty" yaml:"Media,omitempty"`
	PTZ       *PTZCapabilities       `xml:"PTZ,omitempty" json:"PTZ,omitempty" yaml:"PTZ,omitempty"`
	Extension *CapabilitiesExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type CapabilitiesExtension struct {
	DeviceIO        *DeviceIOCapabilities        `xml:"DeviceIO,omitempty" json:"DeviceIO,omitempty" yaml:"DeviceIO,omitempty"`
	Display         *DisplayCapabilities         `xml:"Display,omitempty" json:"Display,omitempty" yaml:"Display,omitempty"`
	Recording       *RecordingCapabilities       `xml:"Recording,omitempty" json:"Recording,omitempty" yaml:"Recording,omitempty"`
	Search          *SearchCapabilities          `xml:"Search,omitempty" json:"Search,omitempty" yaml:"Search,omitempty"`
	Replay          *ReplayCapabilities          `xml:"Replay,omitempty" json:"Replay,omitempty" yaml:"Replay,omitempty"`
	Receiver        *ReceiverCapabilities        `xml:"Receiver,omitempty" json:"Receiver,omitempty" yaml:"Receiver,omitempty"`
	AnalyticsDevice *AnalyticsDeviceCapabilities `xml:"AnalyticsDevice,omitempty" json:"AnalyticsDevice,omitempty" yaml:"AnalyticsDevice,omitempty"`
	Extensions      *CapabilitiesExtension2      `xml:"Extensions,omitempty" json:"Extensions,omitempty" yaml:"Extensions,omitempty"`
}

type CapabilitiesExtension2 []interface{}

type CellLayout struct {
	Transformation *Transformation `xml:"Transformation,omitempty" json:"Transformation,omitempty" yaml:"Transformation,omitempty"`
	Columns        int64           `xml:"Columns,attr,omitempty" json:"Columns,attr,omitempty" yaml:"Columns,attr,omitempty"`
	Rows           int64           `xml:"Rows,attr,omitempty" json:"Rows,attr,omitempty" yaml:"Rows,attr,omitempty"`
}

type Certificate struct {
	CertificateID *string     `xml:"CertificateID,omitempty" json:"CertificateID,omitempty" yaml:"CertificateID,omitempty"`
	Certificate   *BinaryData `xml:"Certificate,omitempty" json:"Certificate,omitempty" yaml:"Certificate,omitempty"`
}

type CertificateGenerationParameters struct {
	CertificateID  *string                                   `xml:"CertificateID,omitempty" json:"CertificateID,omitempty" yaml:"CertificateID,omitempty"`
	Subject        *string                                   `xml:"Subject,omitempty" json:"Subject,omitempty" yaml:"Subject,omitempty"`
	ValidNotBefore *string                                   `xml:"ValidNotBefore,omitempty" json:"ValidNotBefore,omitempty" yaml:"ValidNotBefore,omitempty"`
	ValidNotAfter  *string                                   `xml:"ValidNotAfter,omitempty" json:"ValidNotAfter,omitempty" yaml:"ValidNotAfter,omitempty"`
	Extension      *CertificateGenerationParametersExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type CertificateGenerationParametersExtension []interface{}

type CertificateInformation struct {
	CertificateID      *string                          `xml:"CertificateID,omitempty" json:"CertificateID,omitempty" yaml:"CertificateID,omitempty"`
	IssuerDN           *string                          `xml:"IssuerDN,omitempty" json:"IssuerDN,omitempty" yaml:"IssuerDN,omitempty"`
	SubjectDN          *string                          `xml:"SubjectDN,omitempty" json:"SubjectDN,omitempty" yaml:"SubjectDN,omitempty"`
	KeyUsage           *CertificateUsage                `xml:"KeyUsage,omitempty" json:"KeyUsage,omitempty" yaml:"KeyUsage,omitempty"`
	ExtendedKeyUsage   *CertificateUsage                `xml:"ExtendedKeyUsage,omitempty" json:"ExtendedKeyUsage,omitempty" yaml:"ExtendedKeyUsage,omitempty"`
	KeyLength          *int                             `xml:"KeyLength,omitempty" json:"KeyLength,omitempty" yaml:"KeyLength,omitempty"`
	Version            *string                          `xml:"Version,omitempty" json:"Version,omitempty" yaml:"Version,omitempty"`
	SerialNum          *string                          `xml:"SerialNum,omitempty" json:"SerialNum,omitempty" yaml:"SerialNum,omitempty"`
	SignatureAlgorithm *string                          `xml:"SignatureAlgorithm,omitempty" json:"SignatureAlgorithm,omitempty" yaml:"SignatureAlgorithm,omitempty"`
	Validity           *DateTimeRange                   `xml:"Validity,omitempty" json:"Validity,omitempty" yaml:"Validity,omitempty"`
	Extension          *CertificateInformationExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type CertificateInformationExtension []interface{}

type CertificateStatus struct {
	CertificateID *string `xml:"CertificateID,omitempty" json:"CertificateID,omitempty" yaml:"CertificateID,omitempty"`
	Status        *bool   `xml:"Status,omitempty" json:"Status,omitempty" yaml:"Status,omitempty"`
}

type CertificateUsage struct {
	Content  *string `xml:"Content,omitempty" json:"Content,omitempty" yaml:"Content,omitempty"`
	Critical bool    `xml:"Critical,attr,omitempty" json:"Critical,attr,omitempty" yaml:"Critical,attr,omitempty"`
}

type CertificateWithPrivateKey struct {
	CertificateID *string     `xml:"CertificateID,omitempty" json:"CertificateID,omitempty" yaml:"CertificateID,omitempty"`
	Certificate   *BinaryData `xml:"Certificate,omitempty" json:"Certificate,omitempty" yaml:"Certificate,omitempty"`
	PrivateKey    *BinaryData `xml:"PrivateKey,omitempty" json:"PrivateKey,omitempty" yaml:"PrivateKey,omitempty"`
}

// This type contains the Audio and Video coding capabilities of
//                         a display service.
type CodingCapabilities struct {
	AudioEncodingCapabilities *AudioEncoderConfigurationOptions `xml:"AudioEncodingCapabilities,omitempty" json:"AudioEncodingCapabilities,omitempty" yaml:"AudioEncodingCapabilities,omitempty"`
	AudioDecodingCapabilities *AudioDecoderConfigurationOptions `xml:"AudioDecodingCapabilities,omitempty" json:"AudioDecodingCapabilities,omitempty" yaml:"AudioDecodingCapabilities,omitempty"`
	VideoDecodingCapabilities *VideoDecoderConfigurationOptions `xml:"VideoDecodingCapabilities,omitempty" json:"VideoDecodingCapabilities,omitempty" yaml:"VideoDecodingCapabilities,omitempty"`
}

// Describe the colors supported. Either list each color or
//                      define the range of color values.
type ColorOptions struct {
	ColorList       []*Color           `xml:"ColorList,omitempty" json:"ColorList,omitempty" yaml:"ColorList,omitempty"`
	ColorspaceRange []*ColorspaceRange `xml:"ColorspaceRange,omitempty" json:"ColorspaceRange,omitempty" yaml:"ColorspaceRange,omitempty"`
}

type ColorspaceRange struct {
	X          *FloatRange `xml:"X,omitempty" json:"X,omitempty" yaml:"X,omitempty"`
	Y          *FloatRange `xml:"Y,omitempty" json:"Y,omitempty" yaml:"Y,omitempty"`
	Z          *FloatRange `xml:"Z,omitempty" json:"Z,omitempty" yaml:"Z,omitempty"`
	Colorspace *string     `xml:"Colorspace,omitempty" json:"Colorspace,omitempty" yaml:"Colorspace,omitempty"`
}

type Config struct {
	Parameters *ItemList `xml:"Parameters,omitempty" json:"Parameters,omitempty" yaml:"Parameters,omitempty"`
	Name       string    `xml:"Name,attr,omitempty" json:"Name,attr,omitempty" yaml:"Name,attr,omitempty"`
	Type       string    `xml:"Type,attr,omitempty" json:"Type,attr,omitempty" yaml:"Type,attr,omitempty"`
}

type Frame struct {
	PTZStatus      *PTZStatus      `xml:"PTZStatus,omitempty" json:"PTZStatus,omitempty" yaml:"PTZStatus,omitempty"`
	Transformation *Transformation `xml:"Transformation,omitempty" json:"Transformation,omitempty" yaml:"Transformation,omitempty"`
	Object         []*Object       `xml:"Object,omitempty" json:"Object,omitempty" yaml:"Object,omitempty"`
	ObjectTree     *ObjectTree     `xml:"ObjectTree,omitempty" json:"ObjectTree,omitempty" yaml:"ObjectTree,omitempty"`
	Extension      *FrameExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
	SceneImageRef  *string         `xml:"SceneImageRef,omitempty" json:"SceneImageRef,omitempty" yaml:"SceneImageRef,omitempty"`
	SceneImage     *[]byte         `xml:"SceneImage,omitempty" json:"SceneImage,omitempty" yaml:"SceneImage,omitempty"`
	UtcTime        DateTime        `xml:"UtcTime,attr,omitempty" json:"UtcTime,attr,omitempty" yaml:"UtcTime,attr,omitempty"`
	Colorspace     string          `xml:"Colorspace,attr,omitempty" json:"Colorspace,attr,omitempty" yaml:"Colorspace,attr,omitempty"`
	Source         string          `xml:"Source,attr,omitempty" json:"Source,attr,omitempty" yaml:"Source,attr,omitempty"`
}

type FrameExtension struct {
	MotionInCells *MotionInCells   `xml:"MotionInCells,omitempty" json:"MotionInCells,omitempty" yaml:"MotionInCells,omitempty"`
	Extension     *FrameExtension2 `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type ConfigDescription struct {
	Parameters   *ItemListDescription        `xml:"Parameters,omitempty" json:"Parameters,omitempty" yaml:"Parameters,omitempty"`
	Messages     []*string                   `xml:"Messages,omitempty" json:"Messages,omitempty" yaml:"Messages,omitempty"`
	Extension    *ConfigDescriptionExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
	Name         string                      `xml:"Name,attr,omitempty" json:"Name,attr,omitempty" yaml:"Name,attr,omitempty"`
	Fixed        bool                        `xml:"fixed,attr,omitempty" json:"fixed,attr,omitempty" yaml:"fixed,attr,omitempty"`
	MaxInstances int64                       `xml:"maxInstances,attr,omitempty" json:"maxInstances,attr,omitempty" yaml:"maxInstances,attr,omitempty"`
}

type ConfigDescriptionExtension []interface{}

// Base type defining the common properties of a configuration.
type ConfigurationEntity struct {
	Name     *Name          `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	UseCount *int           `xml:"UseCount,omitempty" json:"UseCount,omitempty" yaml:"UseCount,omitempty"`
	Token    ReferenceToken `xml:"token,attr,omitempty" json:"token,attr,omitempty" yaml:"token,attr,omitempty"`
}

type ContinuousFocus struct {
	Speed *float64 `xml:"Speed,omitempty" json:"Speed,omitempty" yaml:"Speed,omitempty"`
}

type ContinuousFocusOptions struct {
	Speed *FloatRange `xml:"Speed,omitempty" json:"Speed,omitempty" yaml:"Speed,omitempty"`
}

type CreatePullPoint []interface{}

type CreatePullPointResponse struct {
	PullPoint *EndpointReferenceType `xml:"PullPoint,omitempty" json:"PullPoint,omitempty" yaml:"PullPoint,omitempty"`
}

type DNSInformation struct {
	FromDHCP     *bool                    `xml:"FromDHCP,omitempty" json:"FromDHCP,omitempty" yaml:"FromDHCP,omitempty"`
	SearchDomain []*string                `xml:"SearchDomain,omitempty" json:"SearchDomain,omitempty" yaml:"SearchDomain,omitempty"`
	DNSFromDHCP  []*IPAddress             `xml:"DNSFromDHCP,omitempty" json:"DNSFromDHCP,omitempty" yaml:"DNSFromDHCP,omitempty"`
	DNSManual    []*IPAddress             `xml:"DNSManual,omitempty" json:"DNSManual,omitempty" yaml:"DNSManual,omitempty"`
	Extension    *DNSInformationExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type DNSInformationExtension []interface{}

type DateStruct struct {
	Year  *int `xml:"Year,omitempty" json:"Year,omitempty" yaml:"Year,omitempty"`
	Month *int `xml:"Month,omitempty" json:"Month,omitempty" yaml:"Month,omitempty"`
	Day   *int `xml:"Day,omitempty" json:"Day,omitempty" yaml:"Day,omitempty"`
}

type DateTimeRange struct {
	From  *DateTime `xml:"From,omitempty" json:"From,omitempty" yaml:"From,omitempty"`
	Until *DateTime `xml:"Until,omitempty" json:"Until,omitempty" yaml:"Until,omitempty"`
}

type DateTimeStruct struct {
	Time *TimeStruct `xml:"Time,omitempty" json:"Time,omitempty" yaml:"Time,omitempty"`
	Date *DateStruct `xml:"Date,omitempty" json:"Date,omitempty" yaml:"Date,omitempty"`
}

type Defogging struct {
	Mode      *string             `xml:"Mode,omitempty" json:"Mode,omitempty" yaml:"Mode,omitempty"`
	Level     *float64            `xml:"Level,omitempty" json:"Level,omitempty" yaml:"Level,omitempty"`
	Extension *DefoggingExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type DefoggingExtension []interface{}

type DefoggingOptions struct {
	Mode  []*string `xml:"Mode,omitempty" json:"Mode,omitempty" yaml:"Mode,omitempty"`
	Level *bool     `xml:"Level,omitempty" json:"Level,omitempty" yaml:"Level,omitempty"`
}

type DestroyPullPoint []interface{}

type DestroyPullPointResponse []interface{}

type DeviceCapabilities struct {
	XAddr     *string                      `xml:"XAddr,omitempty" json:"XAddr,omitempty" yaml:"XAddr,omitempty"`
	Network   *NetworkCapabilities         `xml:"Network,omitempty" json:"Network,omitempty" yaml:"Network,omitempty"`
	System    *SystemCapabilities          `xml:"System,omitempty" json:"System,omitempty" yaml:"System,omitempty"`
	IO        *IOCapabilities              `xml:"IO,omitempty" json:"IO,omitempty" yaml:"IO,omitempty"`
	Security  *SecurityCapabilities        `xml:"Security,omitempty" json:"Security,omitempty" yaml:"Security,omitempty"`
	Extension *DeviceCapabilitiesExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type DeviceCapabilitiesExtension []interface{}

// Base class for physical entities like inputs and outputs.
type DeviceEntity struct {
	Token ReferenceToken `xml:"token,attr,omitempty" json:"token,attr,omitempty" yaml:"token,attr,omitempty"`
}

type DeviceIOCapabilities struct {
	XAddr        *string `xml:"XAddr,omitempty" json:"XAddr,omitempty" yaml:"XAddr,omitempty"`
	VideoSources *int    `xml:"VideoSources,omitempty" json:"VideoSources,omitempty" yaml:"VideoSources,omitempty"`
	VideoOutputs *int    `xml:"VideoOutputs,omitempty" json:"VideoOutputs,omitempty" yaml:"VideoOutputs,omitempty"`
	AudioSources *int    `xml:"AudioSources,omitempty" json:"AudioSources,omitempty" yaml:"AudioSources,omitempty"`
	AudioOutputs *int    `xml:"AudioOutputs,omitempty" json:"AudioOutputs,omitempty" yaml:"AudioOutputs,omitempty"`
	RelayOutputs *int    `xml:"RelayOutputs,omitempty" json:"RelayOutputs,omitempty" yaml:"RelayOutputs,omitempty"`
}

type DigitalInput struct {
	Token         ReferenceToken   `xml:"token,attr,omitempty" json:"token,attr,omitempty" yaml:"token,attr,omitempty"`
	IdleState     DigitalIdleState `xml:"IdleState,attr,omitempty" json:"IdleState,attr,omitempty" yaml:"IdleState,attr,omitempty"`
	TypeAttrXSI   string           `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string           `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *DigitalInput) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:DigitalInput"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://www.onvif.org/ver10/schema"
	}
}

type DisplayCapabilities struct {
	XAddr       *string `xml:"XAddr,omitempty" json:"XAddr,omitempty" yaml:"XAddr,omitempty"`
	FixedLayout *bool   `xml:"FixedLayout,omitempty" json:"FixedLayout,omitempty" yaml:"FixedLayout,omitempty"`
}

type Documentation []interface{}

type Dot11AvailableNetworks struct {
	SSID                  *Dot11SSIDType                   `xml:"SSID,omitempty" json:"SSID,omitempty" yaml:"SSID,omitempty"`
	BSSID                 *string                          `xml:"BSSID,omitempty" json:"BSSID,omitempty" yaml:"BSSID,omitempty"`
	AuthAndMangementSuite []*Dot11AuthAndMangementSuite    `xml:"AuthAndMangementSuite,omitempty" json:"AuthAndMangementSuite,omitempty" yaml:"AuthAndMangementSuite,omitempty"`
	PairCipher            []*Dot11Cipher                   `xml:"PairCipher,omitempty" json:"PairCipher,omitempty" yaml:"PairCipher,omitempty"`
	GroupCipher           []*Dot11Cipher                   `xml:"GroupCipher,omitempty" json:"GroupCipher,omitempty" yaml:"GroupCipher,omitempty"`
	SignalStrength        *Dot11SignalStrength             `xml:"SignalStrength,omitempty" json:"SignalStrength,omitempty" yaml:"SignalStrength,omitempty"`
	Extension             *Dot11AvailableNetworksExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type Dot11AvailableNetworksExtension []interface{}

type Dot11Capabilities struct {
	TKIP                  *bool `xml:"TKIP,omitempty" json:"TKIP,omitempty" yaml:"TKIP,omitempty"`
	ScanAvailableNetworks *bool `xml:"ScanAvailableNetworks,omitempty" json:"ScanAvailableNetworks,omitempty" yaml:"ScanAvailableNetworks,omitempty"`
	MultipleConfiguration *bool `xml:"MultipleConfiguration,omitempty" json:"MultipleConfiguration,omitempty" yaml:"MultipleConfiguration,omitempty"`
	AdHocStationMode      *bool `xml:"AdHocStationMode,omitempty" json:"AdHocStationMode,omitempty" yaml:"AdHocStationMode,omitempty"`
	WEP                   *bool `xml:"WEP,omitempty" json:"WEP,omitempty" yaml:"WEP,omitempty"`
}

type Dot11Configuration struct {
	SSID     *Dot11SSIDType                  `xml:"SSID,omitempty" json:"SSID,omitempty" yaml:"SSID,omitempty"`
	Mode     *Dot11StationMode               `xml:"Mode,omitempty" json:"Mode,omitempty" yaml:"Mode,omitempty"`
	Alias    *Name                           `xml:"Alias,omitempty" json:"Alias,omitempty" yaml:"Alias,omitempty"`
	Priority *NetworkInterfaceConfigPriority `xml:"Priority,omitempty" json:"Priority,omitempty" yaml:"Priority,omitempty"`
	Security *Dot11SecurityConfiguration     `xml:"Security,omitempty" json:"Security,omitempty" yaml:"Security,omitempty"`
}

type Dot11PSKSet struct {
	Key        *Dot11PSK             `xml:"Key,omitempty" json:"Key,omitempty" yaml:"Key,omitempty"`
	Passphrase *Dot11PSKPassphrase   `xml:"Passphrase,omitempty" json:"Passphrase,omitempty" yaml:"Passphrase,omitempty"`
	Extension  *Dot11PSKSetExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type Dot11PSKSetExtension []interface{}

type Dot11SecurityConfiguration struct {
	Mode      *Dot11SecurityMode                   `xml:"Mode,omitempty" json:"Mode,omitempty" yaml:"Mode,omitempty"`
	Algorithm *Dot11Cipher                         `xml:"Algorithm,omitempty" json:"Algorithm,omitempty" yaml:"Algorithm,omitempty"`
	PSK       *Dot11PSKSet                         `xml:"PSK,omitempty" json:"PSK,omitempty" yaml:"PSK,omitempty"`
	Dot1X     ReferenceToken                       `xml:"Dot1X,omitempty" json:"Dot1X,omitempty" yaml:"Dot1X,omitempty"`
	Extension *Dot11SecurityConfigurationExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type Dot11SecurityConfigurationExtension []interface{}

type Dot11Status struct {
	SSID              *Dot11SSIDType       `xml:"SSID,omitempty" json:"SSID,omitempty" yaml:"SSID,omitempty"`
	BSSID             *string              `xml:"BSSID,omitempty" json:"BSSID,omitempty" yaml:"BSSID,omitempty"`
	PairCipher        *Dot11Cipher         `xml:"PairCipher,omitempty" json:"PairCipher,omitempty" yaml:"PairCipher,omitempty"`
	GroupCipher       *Dot11Cipher         `xml:"GroupCipher,omitempty" json:"GroupCipher,omitempty" yaml:"GroupCipher,omitempty"`
	SignalStrength    *Dot11SignalStrength `xml:"SignalStrength,omitempty" json:"SignalStrength,omitempty" yaml:"SignalStrength,omitempty"`
	ActiveConfigAlias ReferenceToken       `xml:"ActiveConfigAlias,omitempty" json:"ActiveConfigAlias,omitempty" yaml:"ActiveConfigAlias,omitempty"`
}

type Dot1XConfiguration struct {
	Dot1XConfigurationToken ReferenceToken               `xml:"Dot1XConfigurationToken,omitempty" json:"Dot1XConfigurationToken,omitempty" yaml:"Dot1XConfigurationToken,omitempty"`
	Identity                *string                      `xml:"Identity,omitempty" json:"Identity,omitempty" yaml:"Identity,omitempty"`
	AnonymousID             *string                      `xml:"AnonymousID,omitempty" json:"AnonymousID,omitempty" yaml:"AnonymousID,omitempty"`
	EAPMethod               *int                         `xml:"EAPMethod,omitempty" json:"EAPMethod,omitempty" yaml:"EAPMethod,omitempty"`
	CACertificateID         []*string                    `xml:"CACertificateID,omitempty" json:"CACertificateID,omitempty" yaml:"CACertificateID,omitempty"`
	EAPMethodConfiguration  *EAPMethodConfiguration      `xml:"EAPMethodConfiguration,omitempty" json:"EAPMethodConfiguration,omitempty" yaml:"EAPMethodConfiguration,omitempty"`
	Extension               *Dot1XConfigurationExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type Dot1XConfigurationExtension []interface{}

type Dot3Configuration []interface{}

// Range of duration greater equal Min duration and less equal
//                         Max duration.
type DurationRange struct {
	Min *Duration `xml:"Min,omitempty" json:"Min,omitempty" yaml:"Min,omitempty"`
	Max *Duration `xml:"Max,omitempty" json:"Max,omitempty" yaml:"Max,omitempty"`
}

type DynamicDNSInformation struct {
	Type      *DynamicDNSType                 `xml:"Type,omitempty" json:"Type,omitempty" yaml:"Type,omitempty"`
	Name      *DNSName                        `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	TTL       *Duration                       `xml:"TTL,omitempty" json:"TTL,omitempty" yaml:"TTL,omitempty"`
	Extension *DynamicDNSInformationExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type DynamicDNSInformationExtension []interface{}

type EAPMethodConfiguration struct {
	TLSConfiguration *TLSConfiguration   `xml:"TLSConfiguration,omitempty" json:"TLSConfiguration,omitempty" yaml:"TLSConfiguration,omitempty"`
	Password         *string             `xml:"Password,omitempty" json:"Password,omitempty" yaml:"Password,omitempty"`
	Extension        *EapMethodExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type EFlip struct {
	Mode *EFlipMode `xml:"Mode,omitempty" json:"Mode,omitempty" yaml:"Mode,omitempty"`
}

type EFlipOptions struct {
	Mode      []*EFlipMode           `xml:"Mode,omitempty" json:"Mode,omitempty" yaml:"Mode,omitempty"`
	Extension *EFlipOptionsExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type EFlipOptionsExtension []interface{}

type EapMethodExtension []interface{}

type EndpointReferenceType struct {
	Address             *AttributedURIType       `xml:"Address,omitempty" json:"Address,omitempty" yaml:"Address,omitempty"`
	ReferenceParameters *ReferenceParametersType `xml:"ReferenceParameters,omitempty" json:"ReferenceParameters,omitempty" yaml:"ReferenceParameters,omitempty"`
	Metadata            *MetadataType            `xml:"Metadata,omitempty" json:"Metadata,omitempty" yaml:"Metadata,omitempty"`
}

type EngineConfiguration struct {
	VideoAnalyticsConfiguration *VideoAnalyticsConfiguration `xml:"VideoAnalyticsConfiguration,omitempty" json:"VideoAnalyticsConfiguration,omitempty" yaml:"VideoAnalyticsConfiguration,omitempty"`
	AnalyticsEngineInputInfo    *AnalyticsEngineInputInfo    `xml:"AnalyticsEngineInputInfo,omitempty" json:"AnalyticsEngineInputInfo,omitempty" yaml:"AnalyticsEngineInputInfo,omitempty"`
}

type Envelope struct {
	Header *Header `xml:"Header,omitempty" json:"Header,omitempty" yaml:"Header,omitempty"`
	Body   *Body   `xml:"Body,omitempty" json:"Body,omitempty" yaml:"Body,omitempty"`
}

type EventCapabilities struct {
	XAddr                                         *string `xml:"XAddr,omitempty" json:"XAddr,omitempty" yaml:"XAddr,omitempty"`
	WSSubscriptionPolicySupport                   *bool   `xml:"WSSubscriptionPolicySupport,omitempty" json:"WSSubscriptionPolicySupport,omitempty" yaml:"WSSubscriptionPolicySupport,omitempty"`
	WSPullPointSupport                            *bool   `xml:"WSPullPointSupport,omitempty" json:"WSPullPointSupport,omitempty" yaml:"WSPullPointSupport,omitempty"`
	WSPausableSubscriptionManagerInterfaceSupport *bool   `xml:"WSPausableSubscriptionManagerInterfaceSupport,omitempty" json:"WSPausableSubscriptionManagerInterfaceSupport,omitempty" yaml:"WSPausableSubscriptionManagerInterfaceSupport,omitempty"`
}

type EventFilter struct {
	TypeAttrXSI   string `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *EventFilter) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:EventFilter"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://www.onvif.org/ver10/schema"
	}
}

// Subcription handling in the same way as base notification
//                       subscription.
type EventSubscription struct {
	Filter             *FilterType    `xml:"Filter,omitempty" json:"Filter,omitempty" yaml:"Filter,omitempty"`
	SubscriptionPolicy []*interface{} `xml:"SubscriptionPolicy>SubscriptionPolicy,omitempty" json:"SubscriptionPolicy>SubscriptionPolicy,omitempty" yaml:"SubscriptionPolicy>SubscriptionPolicy,omitempty"`
}

type Exposure struct {
	Mode            *ExposureMode     `xml:"Mode,omitempty" json:"Mode,omitempty" yaml:"Mode,omitempty"`
	Priority        *ExposurePriority `xml:"Priority,omitempty" json:"Priority,omitempty" yaml:"Priority,omitempty"`
	Window          *Rectangle        `xml:"Window,omitempty" json:"Window,omitempty" yaml:"Window,omitempty"`
	MinExposureTime *float64          `xml:"MinExposureTime,omitempty" json:"MinExposureTime,omitempty" yaml:"MinExposureTime,omitempty"`
	MaxExposureTime *float64          `xml:"MaxExposureTime,omitempty" json:"MaxExposureTime,omitempty" yaml:"MaxExposureTime,omitempty"`
	MinGain         *float64          `xml:"MinGain,omitempty" json:"MinGain,omitempty" yaml:"MinGain,omitempty"`
	MaxGain         *float64          `xml:"MaxGain,omitempty" json:"MaxGain,omitempty" yaml:"MaxGain,omitempty"`
	MinIris         *float64          `xml:"MinIris,omitempty" json:"MinIris,omitempty" yaml:"MinIris,omitempty"`
	MaxIris         *float64          `xml:"MaxIris,omitempty" json:"MaxIris,omitempty" yaml:"MaxIris,omitempty"`
	ExposureTime    *float64          `xml:"ExposureTime,omitempty" json:"ExposureTime,omitempty" yaml:"ExposureTime,omitempty"`
	Gain            *float64          `xml:"Gain,omitempty" json:"Gain,omitempty" yaml:"Gain,omitempty"`
	Iris            *float64          `xml:"Iris,omitempty" json:"Iris,omitempty" yaml:"Iris,omitempty"`
}

// Type describing the exposure settings.
type Exposure20 struct {
	Mode            *ExposureMode     `xml:"Mode,omitempty" json:"Mode,omitempty" yaml:"Mode,omitempty"`
	Priority        *ExposurePriority `xml:"Priority,omitempty" json:"Priority,omitempty" yaml:"Priority,omitempty"`
	Window          *Rectangle        `xml:"Window,omitempty" json:"Window,omitempty" yaml:"Window,omitempty"`
	MinExposureTime *float64          `xml:"MinExposureTime,omitempty" json:"MinExposureTime,omitempty" yaml:"MinExposureTime,omitempty"`
	MaxExposureTime *float64          `xml:"MaxExposureTime,omitempty" json:"MaxExposureTime,omitempty" yaml:"MaxExposureTime,omitempty"`
	MinGain         *float64          `xml:"MinGain,omitempty" json:"MinGain,omitempty" yaml:"MinGain,omitempty"`
	MaxGain         *float64          `xml:"MaxGain,omitempty" json:"MaxGain,omitempty" yaml:"MaxGain,omitempty"`
	MinIris         *float64          `xml:"MinIris,omitempty" json:"MinIris,omitempty" yaml:"MinIris,omitempty"`
	MaxIris         *float64          `xml:"MaxIris,omitempty" json:"MaxIris,omitempty" yaml:"MaxIris,omitempty"`
	ExposureTime    *float64          `xml:"ExposureTime,omitempty" json:"ExposureTime,omitempty" yaml:"ExposureTime,omitempty"`
	Gain            *float64          `xml:"Gain,omitempty" json:"Gain,omitempty" yaml:"Gain,omitempty"`
	Iris            *float64          `xml:"Iris,omitempty" json:"Iris,omitempty" yaml:"Iris,omitempty"`
}

type ExposureOptions struct {
	Mode            []*ExposureMode     `xml:"Mode,omitempty" json:"Mode,omitempty" yaml:"Mode,omitempty"`
	Priority        []*ExposurePriority `xml:"Priority,omitempty" json:"Priority,omitempty" yaml:"Priority,omitempty"`
	MinExposureTime *FloatRange         `xml:"MinExposureTime,omitempty" json:"MinExposureTime,omitempty" yaml:"MinExposureTime,omitempty"`
	MaxExposureTime *FloatRange         `xml:"MaxExposureTime,omitempty" json:"MaxExposureTime,omitempty" yaml:"MaxExposureTime,omitempty"`
	MinGain         *FloatRange         `xml:"MinGain,omitempty" json:"MinGain,omitempty" yaml:"MinGain,omitempty"`
	MaxGain         *FloatRange         `xml:"MaxGain,omitempty" json:"MaxGain,omitempty" yaml:"MaxGain,omitempty"`
	MinIris         *FloatRange         `xml:"MinIris,omitempty" json:"MinIris,omitempty" yaml:"MinIris,omitempty"`
	MaxIris         *FloatRange         `xml:"MaxIris,omitempty" json:"MaxIris,omitempty" yaml:"MaxIris,omitempty"`
	ExposureTime    *FloatRange         `xml:"ExposureTime,omitempty" json:"ExposureTime,omitempty" yaml:"ExposureTime,omitempty"`
	Gain            *FloatRange         `xml:"Gain,omitempty" json:"Gain,omitempty" yaml:"Gain,omitempty"`
	Iris            *FloatRange         `xml:"Iris,omitempty" json:"Iris,omitempty" yaml:"Iris,omitempty"`
}

type ExposureOptions20 struct {
	Mode            []*ExposureMode     `xml:"Mode,omitempty" json:"Mode,omitempty" yaml:"Mode,omitempty"`
	Priority        []*ExposurePriority `xml:"Priority,omitempty" json:"Priority,omitempty" yaml:"Priority,omitempty"`
	MinExposureTime *FloatRange         `xml:"MinExposureTime,omitempty" json:"MinExposureTime,omitempty" yaml:"MinExposureTime,omitempty"`
	MaxExposureTime *FloatRange         `xml:"MaxExposureTime,omitempty" json:"MaxExposureTime,omitempty" yaml:"MaxExposureTime,omitempty"`
	MinGain         *FloatRange         `xml:"MinGain,omitempty" json:"MinGain,omitempty" yaml:"MinGain,omitempty"`
	MaxGain         *FloatRange         `xml:"MaxGain,omitempty" json:"MaxGain,omitempty" yaml:"MaxGain,omitempty"`
	MinIris         *FloatRange         `xml:"MinIris,omitempty" json:"MinIris,omitempty" yaml:"MinIris,omitempty"`
	MaxIris         *FloatRange         `xml:"MaxIris,omitempty" json:"MaxIris,omitempty" yaml:"MaxIris,omitempty"`
	ExposureTime    *FloatRange         `xml:"ExposureTime,omitempty" json:"ExposureTime,omitempty" yaml:"ExposureTime,omitempty"`
	Gain            *FloatRange         `xml:"Gain,omitempty" json:"Gain,omitempty" yaml:"Gain,omitempty"`
	Iris            *FloatRange         `xml:"Iris,omitempty" json:"Iris,omitempty" yaml:"Iris,omitempty"`
}

type ExtensibleDocumented interface{}

// 	    Fault reporting structure
type Fault struct {
	error
	ErrStr string       `xml:"-" json:"ErrStr,omitempty" yaml:"ErrStr,omitempty"`
	Action string       `xml:"-" json:"Action,omitempty" yaml:"Action,omitempty"`
	Code   *Faultcode   `xml:"Code,omitempty" json:"Code,omitempty" yaml:"Code,omitempty"`
	Reason *Faultreason `xml:"Reason,omitempty" json:"Reason,omitempty" yaml:"Reason,omitempty"`
	Node   *string      `xml:"Node,omitempty" json:"Node,omitempty" yaml:"Node,omitempty"`
	Role   *string      `xml:"Role,omitempty" json:"Role,omitempty" yaml:"Role,omitempty"`
	Detail *Detail      `xml:"Detail,omitempty" json:"Detail,omitempty" yaml:"Detail,omitempty"`
}

func (fault *Fault) Error() string {
	b, _ := json.Marshal(fault)
	errStr := ""
	if fault.error != nil {
		errStr = fault.error.Error()
	} else {
		errStr = "nil"
	}
	return fmt.Sprintf("error: %s, info: %s", errStr, string(b))
}

func (fault *Fault) SetErr(err error) {
	fault.error = err
}

func NewFault(err error) *Fault {
	fault := &Fault{
		error: err,
	}
	if err != nil {
		fault.ErrStr = err.Error()
	}
	return fault
}

type FileProgress struct {
	FileName *string  `xml:"FileName,omitempty" json:"FileName,omitempty" yaml:"FileName,omitempty"`
	Progress *float64 `xml:"Progress,omitempty" json:"Progress,omitempty" yaml:"Progress,omitempty"`
}

type FilterType []interface{}

type FindEventResult struct {
	RecordingToken  *RecordingReference            `xml:"RecordingToken,omitempty" json:"RecordingToken,omitempty" yaml:"RecordingToken,omitempty"`
	TrackToken      *TrackReference                `xml:"TrackToken,omitempty" json:"TrackToken,omitempty" yaml:"TrackToken,omitempty"`
	Time            *DateTime                      `xml:"Time,omitempty" json:"Time,omitempty" yaml:"Time,omitempty"`
	Event           *NotificationMessageHolderType `xml:"Event,omitempty" json:"Event,omitempty" yaml:"Event,omitempty"`
	StartStateEvent *bool                          `xml:"StartStateEvent,omitempty" json:"StartStateEvent,omitempty" yaml:"StartStateEvent,omitempty"`
}

type FindEventResultList struct {
	SearchState *SearchState       `xml:"SearchState,omitempty" json:"SearchState,omitempty" yaml:"SearchState,omitempty"`
	Result      []*FindEventResult `xml:"Result,omitempty" json:"Result,omitempty" yaml:"Result,omitempty"`
}

type FindMetadataResult struct {
	RecordingToken *RecordingReference `xml:"RecordingToken,omitempty" json:"RecordingToken,omitempty" yaml:"RecordingToken,omitempty"`
	TrackToken     *TrackReference     `xml:"TrackToken,omitempty" json:"TrackToken,omitempty" yaml:"TrackToken,omitempty"`
	Time           *DateTime           `xml:"Time,omitempty" json:"Time,omitempty" yaml:"Time,omitempty"`
}

type FindMetadataResultList struct {
	SearchState *SearchState          `xml:"SearchState,omitempty" json:"SearchState,omitempty" yaml:"SearchState,omitempty"`
	Result      []*FindMetadataResult `xml:"Result,omitempty" json:"Result,omitempty" yaml:"Result,omitempty"`
}

type FindPTZPositionResult struct {
	RecordingToken *RecordingReference `xml:"RecordingToken,omitempty" json:"RecordingToken,omitempty" yaml:"RecordingToken,omitempty"`
	TrackToken     *TrackReference     `xml:"TrackToken,omitempty" json:"TrackToken,omitempty" yaml:"TrackToken,omitempty"`
	Time           *DateTime           `xml:"Time,omitempty" json:"Time,omitempty" yaml:"Time,omitempty"`
	Position       *PTZVector          `xml:"Position,omitempty" json:"Position,omitempty" yaml:"Position,omitempty"`
}

type FindPTZPositionResultList struct {
	SearchState *SearchState             `xml:"SearchState,omitempty" json:"SearchState,omitempty" yaml:"SearchState,omitempty"`
	Result      []*FindPTZPositionResult `xml:"Result,omitempty" json:"Result,omitempty" yaml:"Result,omitempty"`
}

type FindRecordingResultList struct {
	SearchState          *SearchState            `xml:"SearchState,omitempty" json:"SearchState,omitempty" yaml:"SearchState,omitempty"`
	RecordingInformation []*RecordingInformation `xml:"RecordingInformation,omitempty" json:"RecordingInformation,omitempty" yaml:"RecordingInformation,omitempty"`
}

type FloatItems struct {
	Items []*float64 `xml:"Items,omitempty" json:"Items,omitempty" yaml:"Items,omitempty"`
}

// Range of values greater equal Min value and less equal Max
//                        value.
type FloatRange struct {
	Min *float64 `xml:"Min,omitempty" json:"Min,omitempty" yaml:"Min,omitempty"`
	Max *float64 `xml:"Max,omitempty" json:"Max,omitempty" yaml:"Max,omitempty"`
}

type FocusConfiguration struct {
	AutoFocusMode *AutoFocusMode `xml:"AutoFocusMode,omitempty" json:"AutoFocusMode,omitempty" yaml:"AutoFocusMode,omitempty"`
	DefaultSpeed  *float64       `xml:"DefaultSpeed,omitempty" json:"DefaultSpeed,omitempty" yaml:"DefaultSpeed,omitempty"`
	NearLimit     *float64       `xml:"NearLimit,omitempty" json:"NearLimit,omitempty" yaml:"NearLimit,omitempty"`
	FarLimit      *float64       `xml:"FarLimit,omitempty" json:"FarLimit,omitempty" yaml:"FarLimit,omitempty"`
}

type FocusConfiguration20 struct {
	AutoFocusMode *AutoFocusMode                 `xml:"AutoFocusMode,omitempty" json:"AutoFocusMode,omitempty" yaml:"AutoFocusMode,omitempty"`
	DefaultSpeed  *float64                       `xml:"DefaultSpeed,omitempty" json:"DefaultSpeed,omitempty" yaml:"DefaultSpeed,omitempty"`
	NearLimit     *float64                       `xml:"NearLimit,omitempty" json:"NearLimit,omitempty" yaml:"NearLimit,omitempty"`
	FarLimit      *float64                       `xml:"FarLimit,omitempty" json:"FarLimit,omitempty" yaml:"FarLimit,omitempty"`
	Extension     *FocusConfiguration20Extension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
	AFMode        StringAttrList                 `xml:"AFMode,attr,omitempty" json:"AFMode,attr,omitempty" yaml:"AFMode,attr,omitempty"`
}

type FocusConfiguration20Extension []interface{}

type FocusMove struct {
	Absolute   *AbsoluteFocus   `xml:"Absolute,omitempty" json:"Absolute,omitempty" yaml:"Absolute,omitempty"`
	Relative   *RelativeFocus   `xml:"Relative,omitempty" json:"Relative,omitempty" yaml:"Relative,omitempty"`
	Continuous *ContinuousFocus `xml:"Continuous,omitempty" json:"Continuous,omitempty" yaml:"Continuous,omitempty"`
}

type FocusOptions struct {
	AutoFocusModes []*AutoFocusMode `xml:"AutoFocusModes,omitempty" json:"AutoFocusModes,omitempty" yaml:"AutoFocusModes,omitempty"`
	DefaultSpeed   *FloatRange      `xml:"DefaultSpeed,omitempty" json:"DefaultSpeed,omitempty" yaml:"DefaultSpeed,omitempty"`
	NearLimit      *FloatRange      `xml:"NearLimit,omitempty" json:"NearLimit,omitempty" yaml:"NearLimit,omitempty"`
	FarLimit       *FloatRange      `xml:"FarLimit,omitempty" json:"FarLimit,omitempty" yaml:"FarLimit,omitempty"`
}

type FocusOptions20 struct {
	AutoFocusModes []*AutoFocusMode         `xml:"AutoFocusModes,omitempty" json:"AutoFocusModes,omitempty" yaml:"AutoFocusModes,omitempty"`
	DefaultSpeed   *FloatRange              `xml:"DefaultSpeed,omitempty" json:"DefaultSpeed,omitempty" yaml:"DefaultSpeed,omitempty"`
	NearLimit      *FloatRange              `xml:"NearLimit,omitempty" json:"NearLimit,omitempty" yaml:"NearLimit,omitempty"`
	FarLimit       *FloatRange              `xml:"FarLimit,omitempty" json:"FarLimit,omitempty" yaml:"FarLimit,omitempty"`
	Extension      *FocusOptions20Extension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type FocusOptions20Extension struct {
	AFModes *StringAttrList `xml:"AFModes,omitempty" json:"AFModes,omitempty" yaml:"AFModes,omitempty"`
}

type FocusStatus struct {
	Position   *float64    `xml:"Position,omitempty" json:"Position,omitempty" yaml:"Position,omitempty"`
	MoveStatus *MoveStatus `xml:"MoveStatus,omitempty" json:"MoveStatus,omitempty" yaml:"MoveStatus,omitempty"`
	Error      *string     `xml:"Error,omitempty" json:"Error,omitempty" yaml:"Error,omitempty"`
}

type FocusStatus20 struct {
	Position   *float64                `xml:"Position,omitempty" json:"Position,omitempty" yaml:"Position,omitempty"`
	MoveStatus *MoveStatus             `xml:"MoveStatus,omitempty" json:"MoveStatus,omitempty" yaml:"MoveStatus,omitempty"`
	Error      *string                 `xml:"Error,omitempty" json:"Error,omitempty" yaml:"Error,omitempty"`
	Extension  *FocusStatus20Extension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type FocusStatus20Extension []interface{}

type G711DecOptions struct {
	Bitrate         *IntItems `xml:"Bitrate,omitempty" json:"Bitrate,omitempty" yaml:"Bitrate,omitempty"`
	SampleRateRange *IntItems `xml:"SampleRateRange,omitempty" json:"SampleRateRange,omitempty" yaml:"SampleRateRange,omitempty"`
}

type G726DecOptions struct {
	Bitrate         *IntItems `xml:"Bitrate,omitempty" json:"Bitrate,omitempty" yaml:"Bitrate,omitempty"`
	SampleRateRange *IntItems `xml:"SampleRateRange,omitempty" json:"SampleRateRange,omitempty" yaml:"SampleRateRange,omitempty"`
}

type GenericEapPwdConfigurationExtension []interface{}

type GetCurrentMessage struct {
	Topic *TopicExpressionType `xml:"Topic,omitempty" json:"Topic,omitempty" yaml:"Topic,omitempty"`
}

type GetCurrentMessageResponse []interface{}

type GetMessages struct {
	MaximumNumber *uint `xml:"MaximumNumber,omitempty" json:"MaximumNumber,omitempty" yaml:"MaximumNumber,omitempty"`
}

type GetMessagesResponse struct {
	NotificationMessage *NotificationMessageHolderType `xml:"NotificationMessage,omitempty" json:"NotificationMessage,omitempty" yaml:"NotificationMessage,omitempty"`
}

type GetRecordingJobsResponseItem struct {
	JobToken         *RecordingJobReference     `xml:"JobToken,omitempty" json:"JobToken,omitempty" yaml:"JobToken,omitempty"`
	JobConfiguration *RecordingJobConfiguration `xml:"JobConfiguration,omitempty" json:"JobConfiguration,omitempty" yaml:"JobConfiguration,omitempty"`
}

type GetRecordingsResponseItem struct {
	RecordingToken *RecordingReference     `xml:"RecordingToken,omitempty" json:"RecordingToken,omitempty" yaml:"RecordingToken,omitempty"`
	Configuration  *RecordingConfiguration `xml:"Configuration,omitempty" json:"Configuration,omitempty" yaml:"Configuration,omitempty"`
	Tracks         *GetTracksResponseList  `xml:"Tracks,omitempty" json:"Tracks,omitempty" yaml:"Tracks,omitempty"`
}

type GetTracksResponseItem struct {
	TrackToken    *TrackReference     `xml:"TrackToken,omitempty" json:"TrackToken,omitempty" yaml:"TrackToken,omitempty"`
	Configuration *TrackConfiguration `xml:"Configuration,omitempty" json:"Configuration,omitempty" yaml:"Configuration,omitempty"`
}

type GetTracksResponseList struct {
	Track []*GetTracksResponseItem `xml:"Track,omitempty" json:"Track,omitempty" yaml:"Track,omitempty"`
}

type H264Configuration struct {
	GovLength   *int         `xml:"GovLength,omitempty" json:"GovLength,omitempty" yaml:"GovLength,omitempty"`
	H264Profile *H264Profile `xml:"H264Profile,omitempty" json:"H264Profile,omitempty" yaml:"H264Profile,omitempty"`
}

type H264DecOptions struct {
	ResolutionsAvailable  []*VideoResolution `xml:"ResolutionsAvailable,omitempty" json:"ResolutionsAvailable,omitempty" yaml:"ResolutionsAvailable,omitempty"`
	SupportedH264Profiles []*H264Profile     `xml:"SupportedH264Profiles,omitempty" json:"SupportedH264Profiles,omitempty" yaml:"SupportedH264Profiles,omitempty"`
	SupportedInputBitrate *IntRange          `xml:"SupportedInputBitrate,omitempty" json:"SupportedInputBitrate,omitempty" yaml:"SupportedInputBitrate,omitempty"`
	SupportedFrameRate    *IntRange          `xml:"SupportedFrameRate,omitempty" json:"SupportedFrameRate,omitempty" yaml:"SupportedFrameRate,omitempty"`
}

type H264Options struct {
	ResolutionsAvailable  []*VideoResolution `xml:"ResolutionsAvailable,omitempty" json:"ResolutionsAvailable,omitempty" yaml:"ResolutionsAvailable,omitempty"`
	GovLengthRange        *IntRange          `xml:"GovLengthRange,omitempty" json:"GovLengthRange,omitempty" yaml:"GovLengthRange,omitempty"`
	FrameRateRange        *IntRange          `xml:"FrameRateRange,omitempty" json:"FrameRateRange,omitempty" yaml:"FrameRateRange,omitempty"`
	EncodingIntervalRange *IntRange          `xml:"EncodingIntervalRange,omitempty" json:"EncodingIntervalRange,omitempty" yaml:"EncodingIntervalRange,omitempty"`
	H264ProfilesSupported []*H264Profile     `xml:"H264ProfilesSupported,omitempty" json:"H264ProfilesSupported,omitempty" yaml:"H264ProfilesSupported,omitempty"`
}

type H264Options2 struct {
	ResolutionsAvailable  []*VideoResolution `xml:"ResolutionsAvailable,omitempty" json:"ResolutionsAvailable,omitempty" yaml:"ResolutionsAvailable,omitempty"`
	GovLengthRange        *IntRange          `xml:"GovLengthRange,omitempty" json:"GovLengthRange,omitempty" yaml:"GovLengthRange,omitempty"`
	FrameRateRange        *IntRange          `xml:"FrameRateRange,omitempty" json:"FrameRateRange,omitempty" yaml:"FrameRateRange,omitempty"`
	EncodingIntervalRange *IntRange          `xml:"EncodingIntervalRange,omitempty" json:"EncodingIntervalRange,omitempty" yaml:"EncodingIntervalRange,omitempty"`
	H264ProfilesSupported []*H264Profile     `xml:"H264ProfilesSupported,omitempty" json:"H264ProfilesSupported,omitempty" yaml:"H264ProfilesSupported,omitempty"`
	BitrateRange          *IntRange          `xml:"BitrateRange,omitempty" json:"BitrateRange,omitempty" yaml:"BitrateRange,omitempty"`
	TypeAttrXSI           string             `xml:"xsi:type,attr,omitempty"`
	TypeNamespace         string             `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *H264Options2) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:H264Options2"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://www.onvif.org/ver10/schema"
	}
}

// 	  Elements replacing the wildcard MUST be namespace qualified,
// but can be in the targetNamespace
type Header []interface{}

type HostnameInformation struct {
	FromDHCP  *bool                         `xml:"FromDHCP,omitempty" json:"FromDHCP,omitempty" yaml:"FromDHCP,omitempty"`
	Name      *string                       `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Extension *HostnameInformationExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type HostnameInformationExtension []interface{}

type IOCapabilities struct {
	InputConnectors *int                     `xml:"InputConnectors,omitempty" json:"InputConnectors,omitempty" yaml:"InputConnectors,omitempty"`
	RelayOutputs    *int                     `xml:"RelayOutputs,omitempty" json:"RelayOutputs,omitempty" yaml:"RelayOutputs,omitempty"`
	Extension       *IOCapabilitiesExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type IOCapabilitiesExtension struct {
	Auxiliary         *bool                     `xml:"Auxiliary,omitempty" json:"Auxiliary,omitempty" yaml:"Auxiliary,omitempty"`
	AuxiliaryCommands []*AuxiliaryData          `xml:"AuxiliaryCommands,omitempty" json:"AuxiliaryCommands,omitempty" yaml:"AuxiliaryCommands,omitempty"`
	Extension         *IOCapabilitiesExtension2 `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type IOCapabilitiesExtension2 []interface{}

type IPAddress struct {
	Type        *IPType      `xml:"Type,omitempty" json:"Type,omitempty" yaml:"Type,omitempty"`
	IPv4Address *IPv4Address `xml:"IPv4Address,omitempty" json:"IPv4Address,omitempty" yaml:"IPv4Address,omitempty"`
	IPv6Address *IPv6Address `xml:"IPv6Address,omitempty" json:"IPv6Address,omitempty" yaml:"IPv6Address,omitempty"`
}

type IPAddressFilter struct {
	Type        *IPAddressFilterType      `xml:"Type,omitempty" json:"Type,omitempty" yaml:"Type,omitempty"`
	IPv4Address []*PrefixedIPv4Address    `xml:"IPv4Address,omitempty" json:"IPv4Address,omitempty" yaml:"IPv4Address,omitempty"`
	IPv6Address []*PrefixedIPv6Address    `xml:"IPv6Address,omitempty" json:"IPv6Address,omitempty" yaml:"IPv6Address,omitempty"`
	Extension   *IPAddressFilterExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type IPAddressFilterExtension []interface{}

type IPv4Configuration struct {
	Manual    []*PrefixedIPv4Address `xml:"Manual,omitempty" json:"Manual,omitempty" yaml:"Manual,omitempty"`
	LinkLocal *PrefixedIPv4Address   `xml:"LinkLocal,omitempty" json:"LinkLocal,omitempty" yaml:"LinkLocal,omitempty"`
	FromDHCP  *PrefixedIPv4Address   `xml:"FromDHCP,omitempty" json:"FromDHCP,omitempty" yaml:"FromDHCP,omitempty"`
	DHCP      *bool                  `xml:"DHCP,omitempty" json:"DHCP,omitempty" yaml:"DHCP,omitempty"`
}

type IPv4NetworkInterface struct {
	Enabled *bool              `xml:"Enabled,omitempty" json:"Enabled,omitempty" yaml:"Enabled,omitempty"`
	Config  *IPv4Configuration `xml:"Config,omitempty" json:"Config,omitempty" yaml:"Config,omitempty"`
}

type IPv4NetworkInterfaceSetConfiguration struct {
	Enabled *bool                  `xml:"Enabled,omitempty" json:"Enabled,omitempty" yaml:"Enabled,omitempty"`
	Manual  []*PrefixedIPv4Address `xml:"Manual,omitempty" json:"Manual,omitempty" yaml:"Manual,omitempty"`
	DHCP    *bool                  `xml:"DHCP,omitempty" json:"DHCP,omitempty" yaml:"DHCP,omitempty"`
}

type IPv6Configuration struct {
	AcceptRouterAdvert *bool                       `xml:"AcceptRouterAdvert,omitempty" json:"AcceptRouterAdvert,omitempty" yaml:"AcceptRouterAdvert,omitempty"`
	DHCP               *IPv6DHCPConfiguration      `xml:"DHCP,omitempty" json:"DHCP,omitempty" yaml:"DHCP,omitempty"`
	Manual             []*PrefixedIPv6Address      `xml:"Manual,omitempty" json:"Manual,omitempty" yaml:"Manual,omitempty"`
	LinkLocal          []*PrefixedIPv6Address      `xml:"LinkLocal,omitempty" json:"LinkLocal,omitempty" yaml:"LinkLocal,omitempty"`
	FromDHCP           []*PrefixedIPv6Address      `xml:"FromDHCP,omitempty" json:"FromDHCP,omitempty" yaml:"FromDHCP,omitempty"`
	FromRA             []*PrefixedIPv6Address      `xml:"FromRA,omitempty" json:"FromRA,omitempty" yaml:"FromRA,omitempty"`
	Extension          *IPv6ConfigurationExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type IPv6ConfigurationExtension []interface{}

type IPv6NetworkInterface struct {
	Enabled *bool              `xml:"Enabled,omitempty" json:"Enabled,omitempty" yaml:"Enabled,omitempty"`
	Config  *IPv6Configuration `xml:"Config,omitempty" json:"Config,omitempty" yaml:"Config,omitempty"`
}

type IPv6NetworkInterfaceSetConfiguration struct {
	Enabled            *bool                  `xml:"Enabled,omitempty" json:"Enabled,omitempty" yaml:"Enabled,omitempty"`
	AcceptRouterAdvert *bool                  `xml:"AcceptRouterAdvert,omitempty" json:"AcceptRouterAdvert,omitempty" yaml:"AcceptRouterAdvert,omitempty"`
	Manual             []*PrefixedIPv6Address `xml:"Manual,omitempty" json:"Manual,omitempty" yaml:"Manual,omitempty"`
	DHCP               *IPv6DHCPConfiguration `xml:"DHCP,omitempty" json:"DHCP,omitempty" yaml:"DHCP,omitempty"`
}

type ImageStabilization struct {
	Mode      *ImageStabilizationMode      `xml:"Mode,omitempty" json:"Mode,omitempty" yaml:"Mode,omitempty"`
	Level     *float64                     `xml:"Level,omitempty" json:"Level,omitempty" yaml:"Level,omitempty"`
	Extension *ImageStabilizationExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type ImageStabilizationExtension []interface{}

type ImageStabilizationOptions struct {
	Mode      []*ImageStabilizationMode           `xml:"Mode,omitempty" json:"Mode,omitempty" yaml:"Mode,omitempty"`
	Level     *FloatRange                         `xml:"Level,omitempty" json:"Level,omitempty" yaml:"Level,omitempty"`
	Extension *ImageStabilizationOptionsExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type ImageStabilizationOptionsExtension []interface{}

type ImagingCapabilities struct {
	XAddr *string `xml:"XAddr,omitempty" json:"XAddr,omitempty" yaml:"XAddr,omitempty"`
}

type ImagingOptions struct {
	BacklightCompensation *BacklightCompensationOptions `xml:"BacklightCompensation,omitempty" json:"BacklightCompensation,omitempty" yaml:"BacklightCompensation,omitempty"`
	Brightness            *FloatRange                   `xml:"Brightness,omitempty" json:"Brightness,omitempty" yaml:"Brightness,omitempty"`
	ColorSaturation       *FloatRange                   `xml:"ColorSaturation,omitempty" json:"ColorSaturation,omitempty" yaml:"ColorSaturation,omitempty"`
	Contrast              *FloatRange                   `xml:"Contrast,omitempty" json:"Contrast,omitempty" yaml:"Contrast,omitempty"`
	Exposure              *ExposureOptions              `xml:"Exposure,omitempty" json:"Exposure,omitempty" yaml:"Exposure,omitempty"`
	Focus                 *FocusOptions                 `xml:"Focus,omitempty" json:"Focus,omitempty" yaml:"Focus,omitempty"`
	IrCutFilterModes      []*IrCutFilterMode            `xml:"IrCutFilterModes,omitempty" json:"IrCutFilterModes,omitempty" yaml:"IrCutFilterModes,omitempty"`
	Sharpness             *FloatRange                   `xml:"Sharpness,omitempty" json:"Sharpness,omitempty" yaml:"Sharpness,omitempty"`
	WideDynamicRange      *WideDynamicRangeOptions      `xml:"WideDynamicRange,omitempty" json:"WideDynamicRange,omitempty" yaml:"WideDynamicRange,omitempty"`
	WhiteBalance          *WhiteBalanceOptions          `xml:"WhiteBalance,omitempty" json:"WhiteBalance,omitempty" yaml:"WhiteBalance,omitempty"`
}

type ImagingOptions20 struct {
	BacklightCompensation *BacklightCompensationOptions20 `xml:"BacklightCompensation,omitempty" json:"BacklightCompensation,omitempty" yaml:"BacklightCompensation,omitempty"`
	Brightness            *FloatRange                     `xml:"Brightness,omitempty" json:"Brightness,omitempty" yaml:"Brightness,omitempty"`
	ColorSaturation       *FloatRange                     `xml:"ColorSaturation,omitempty" json:"ColorSaturation,omitempty" yaml:"ColorSaturation,omitempty"`
	Contrast              *FloatRange                     `xml:"Contrast,omitempty" json:"Contrast,omitempty" yaml:"Contrast,omitempty"`
	Exposure              *ExposureOptions20              `xml:"Exposure,omitempty" json:"Exposure,omitempty" yaml:"Exposure,omitempty"`
	Focus                 *FocusOptions20                 `xml:"Focus,omitempty" json:"Focus,omitempty" yaml:"Focus,omitempty"`
	IrCutFilterModes      []*IrCutFilterMode              `xml:"IrCutFilterModes,omitempty" json:"IrCutFilterModes,omitempty" yaml:"IrCutFilterModes,omitempty"`
	Sharpness             *FloatRange                     `xml:"Sharpness,omitempty" json:"Sharpness,omitempty" yaml:"Sharpness,omitempty"`
	WideDynamicRange      *WideDynamicRangeOptions20      `xml:"WideDynamicRange,omitempty" json:"WideDynamicRange,omitempty" yaml:"WideDynamicRange,omitempty"`
	WhiteBalance          *WhiteBalanceOptions20          `xml:"WhiteBalance,omitempty" json:"WhiteBalance,omitempty" yaml:"WhiteBalance,omitempty"`
	Extension             *ImagingOptions20Extension      `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type ImagingOptions20Extension struct {
	ImageStabilization *ImageStabilizationOptions  `xml:"ImageStabilization,omitempty" json:"ImageStabilization,omitempty" yaml:"ImageStabilization,omitempty"`
	Extension          *ImagingOptions20Extension2 `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type ImagingOptions20Extension2 struct {
	IrCutFilterAutoAdjustment *IrCutFilterAutoAdjustmentOptions `xml:"IrCutFilterAutoAdjustment,omitempty" json:"IrCutFilterAutoAdjustment,omitempty" yaml:"IrCutFilterAutoAdjustment,omitempty"`
	Extension                 *ImagingOptions20Extension3       `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type ImagingOptions20Extension3 struct {
	ToneCompensationOptions *ToneCompensationOptions    `xml:"ToneCompensationOptions,omitempty" json:"ToneCompensationOptions,omitempty" yaml:"ToneCompensationOptions,omitempty"`
	DefoggingOptions        *DefoggingOptions           `xml:"DefoggingOptions,omitempty" json:"DefoggingOptions,omitempty" yaml:"DefoggingOptions,omitempty"`
	NoiseReductionOptions   *NoiseReductionOptions      `xml:"NoiseReductionOptions,omitempty" json:"NoiseReductionOptions,omitempty" yaml:"NoiseReductionOptions,omitempty"`
	Extension               *ImagingOptions20Extension4 `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type ImagingOptions20Extension4 []interface{}

type ImagingSettings struct {
	BacklightCompensation *BacklightCompensation    `xml:"BacklightCompensation,omitempty" json:"BacklightCompensation,omitempty" yaml:"BacklightCompensation,omitempty"`
	Brightness            *float64                  `xml:"Brightness,omitempty" json:"Brightness,omitempty" yaml:"Brightness,omitempty"`
	ColorSaturation       *float64                  `xml:"ColorSaturation,omitempty" json:"ColorSaturation,omitempty" yaml:"ColorSaturation,omitempty"`
	Contrast              *float64                  `xml:"Contrast,omitempty" json:"Contrast,omitempty" yaml:"Contrast,omitempty"`
	Exposure              *Exposure                 `xml:"Exposure,omitempty" json:"Exposure,omitempty" yaml:"Exposure,omitempty"`
	Focus                 *FocusConfiguration       `xml:"Focus,omitempty" json:"Focus,omitempty" yaml:"Focus,omitempty"`
	IrCutFilter           *IrCutFilterMode          `xml:"IrCutFilter,omitempty" json:"IrCutFilter,omitempty" yaml:"IrCutFilter,omitempty"`
	Sharpness             *float64                  `xml:"Sharpness,omitempty" json:"Sharpness,omitempty" yaml:"Sharpness,omitempty"`
	WideDynamicRange      *WideDynamicRange         `xml:"WideDynamicRange,omitempty" json:"WideDynamicRange,omitempty" yaml:"WideDynamicRange,omitempty"`
	WhiteBalance          *WhiteBalance             `xml:"WhiteBalance,omitempty" json:"WhiteBalance,omitempty" yaml:"WhiteBalance,omitempty"`
	Extension             *ImagingSettingsExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

// Type describing the ImagingSettings of a VideoSource. The
//                       supported options and ranges can be obtained
// via the GetOptions command.
type ImagingSettings20 struct {
	BacklightCompensation *BacklightCompensation20    `xml:"BacklightCompensation,omitempty" json:"BacklightCompensation,omitempty" yaml:"BacklightCompensation,omitempty"`
	Brightness            *float64                    `xml:"Brightness,omitempty" json:"Brightness,omitempty" yaml:"Brightness,omitempty"`
	ColorSaturation       *float64                    `xml:"ColorSaturation,omitempty" json:"ColorSaturation,omitempty" yaml:"ColorSaturation,omitempty"`
	Contrast              *float64                    `xml:"Contrast,omitempty" json:"Contrast,omitempty" yaml:"Contrast,omitempty"`
	Exposure              *Exposure20                 `xml:"Exposure,omitempty" json:"Exposure,omitempty" yaml:"Exposure,omitempty"`
	Focus                 *FocusConfiguration20       `xml:"Focus,omitempty" json:"Focus,omitempty" yaml:"Focus,omitempty"`
	IrCutFilter           *IrCutFilterMode            `xml:"IrCutFilter,omitempty" json:"IrCutFilter,omitempty" yaml:"IrCutFilter,omitempty"`
	Sharpness             *float64                    `xml:"Sharpness,omitempty" json:"Sharpness,omitempty" yaml:"Sharpness,omitempty"`
	WideDynamicRange      *WideDynamicRange20         `xml:"WideDynamicRange,omitempty" json:"WideDynamicRange,omitempty" yaml:"WideDynamicRange,omitempty"`
	WhiteBalance          *WhiteBalance20             `xml:"WhiteBalance,omitempty" json:"WhiteBalance,omitempty" yaml:"WhiteBalance,omitempty"`
	Extension             *ImagingSettingsExtension20 `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type ImagingSettingsExtension []interface{}

type ImagingSettingsExtension20 struct {
	ImageStabilization *ImageStabilization          `xml:"ImageStabilization,omitempty" json:"ImageStabilization,omitempty" yaml:"ImageStabilization,omitempty"`
	Extension          *ImagingSettingsExtension202 `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type ImagingSettingsExtension202 struct {
	IrCutFilterAutoAdjustment []*IrCutFilterAutoAdjustment `xml:"IrCutFilterAutoAdjustment,omitempty" json:"IrCutFilterAutoAdjustment,omitempty" yaml:"IrCutFilterAutoAdjustment,omitempty"`
	Extension                 *ImagingSettingsExtension203 `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type ImagingSettingsExtension203 struct {
	ToneCompensation *ToneCompensation            `xml:"ToneCompensation,omitempty" json:"ToneCompensation,omitempty" yaml:"ToneCompensation,omitempty"`
	Defogging        *Defogging                   `xml:"Defogging,omitempty" json:"Defogging,omitempty" yaml:"Defogging,omitempty"`
	NoiseReduction   *NoiseReduction              `xml:"NoiseReduction,omitempty" json:"NoiseReduction,omitempty" yaml:"NoiseReduction,omitempty"`
	Extension        *ImagingSettingsExtension204 `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type ImagingSettingsExtension204 []interface{}

type ImagingStatus struct {
	FocusStatus *FocusStatus `xml:"FocusStatus,omitempty" json:"FocusStatus,omitempty" yaml:"FocusStatus,omitempty"`
}

type ImagingStatus20 struct {
	FocusStatus20 *FocusStatus20            `xml:"FocusStatus20,omitempty" json:"FocusStatus20,omitempty" yaml:"FocusStatus20,omitempty"`
	Extension     *ImagingStatus20Extension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type ImagingStatus20Extension []interface{}

type Include []interface{}

// List of values.
type IntItems struct {
	Items []*int `xml:"Items,omitempty" json:"Items,omitempty" yaml:"Items,omitempty"`
}

// Rectangle defined by lower left corner position and size.
//                       Units are pixel.
type IntRectangle struct {
	X      int `xml:"x,attr,omitempty" json:"x,attr,omitempty" yaml:"x,attr,omitempty"`
	Y      int `xml:"y,attr,omitempty" json:"y,attr,omitempty" yaml:"y,attr,omitempty"`
	Width  int `xml:"width,attr,omitempty" json:"width,attr,omitempty" yaml:"width,attr,omitempty"`
	Height int `xml:"height,attr,omitempty" json:"height,attr,omitempty" yaml:"height,attr,omitempty"`
}

// Range of a rectangle. The rectangle itself is defined by lower
//                         left corner position and size. Units
// are pixel.
type IntRectangleRange struct {
	XRange      *IntRange `xml:"XRange,omitempty" json:"XRange,omitempty" yaml:"XRange,omitempty"`
	YRange      *IntRange `xml:"YRange,omitempty" json:"YRange,omitempty" yaml:"YRange,omitempty"`
	WidthRange  *IntRange `xml:"WidthRange,omitempty" json:"WidthRange,omitempty" yaml:"WidthRange,omitempty"`
	HeightRange *IntRange `xml:"HeightRange,omitempty" json:"HeightRange,omitempty" yaml:"HeightRange,omitempty"`
}

type InvalidFilterFaultType struct {
	Timestamp     DateTime               `xml:"Timestamp" json:"Timestamp" yaml:"Timestamp"`
	Originator    *EndpointReferenceType `xml:"Originator,omitempty" json:"Originator,omitempty" yaml:"Originator,omitempty"`
	ErrorCode     *string                `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	Description   []*string              `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	FaultCause    interface{}            `xml:"FaultCause" json:"FaultCause" yaml:"FaultCause"`
	UnknownFilter []string               `xml:"UnknownFilter" json:"UnknownFilter" yaml:"UnknownFilter"`
	TypeAttrXSI   string                 `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string                 `xml:"xmlns:objtype,attr,omitempty"`

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
	Timestamp     DateTime               `xml:"Timestamp" json:"Timestamp" yaml:"Timestamp"`
	Originator    *EndpointReferenceType `xml:"Originator,omitempty" json:"Originator,omitempty" yaml:"Originator,omitempty"`
	ErrorCode     *string                `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	Description   []*string              `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	FaultCause    interface{}            `xml:"FaultCause" json:"FaultCause" yaml:"FaultCause"`
	TypeAttrXSI   string                 `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string                 `xml:"xmlns:objtype,attr,omitempty"`

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
	Timestamp     DateTime               `xml:"Timestamp" json:"Timestamp" yaml:"Timestamp"`
	Originator    *EndpointReferenceType `xml:"Originator,omitempty" json:"Originator,omitempty" yaml:"Originator,omitempty"`
	ErrorCode     *string                `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	Description   []*string              `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	FaultCause    interface{}            `xml:"FaultCause" json:"FaultCause" yaml:"FaultCause"`
	TypeAttrXSI   string                 `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string                 `xml:"xmlns:objtype,attr,omitempty"`

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
	Timestamp     DateTime               `xml:"Timestamp" json:"Timestamp" yaml:"Timestamp"`
	Originator    *EndpointReferenceType `xml:"Originator,omitempty" json:"Originator,omitempty" yaml:"Originator,omitempty"`
	ErrorCode     *string                `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	Description   []*string              `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	FaultCause    interface{}            `xml:"FaultCause" json:"FaultCause" yaml:"FaultCause"`
	TypeAttrXSI   string                 `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string                 `xml:"xmlns:objtype,attr,omitempty"`

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

type IrCutFilterAutoAdjustment struct {
	BoundaryType   *string                             `xml:"BoundaryType,omitempty" json:"BoundaryType,omitempty" yaml:"BoundaryType,omitempty"`
	BoundaryOffset *float64                            `xml:"BoundaryOffset,omitempty" json:"BoundaryOffset,omitempty" yaml:"BoundaryOffset,omitempty"`
	ResponseTime   *Duration                           `xml:"ResponseTime,omitempty" json:"ResponseTime,omitempty" yaml:"ResponseTime,omitempty"`
	Extension      *IrCutFilterAutoAdjustmentExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type IrCutFilterAutoAdjustmentExtension []interface{}

type IrCutFilterAutoAdjustmentOptions struct {
	BoundaryType      []*string                                  `xml:"BoundaryType,omitempty" json:"BoundaryType,omitempty" yaml:"BoundaryType,omitempty"`
	BoundaryOffset    *bool                                      `xml:"BoundaryOffset,omitempty" json:"BoundaryOffset,omitempty" yaml:"BoundaryOffset,omitempty"`
	ResponseTimeRange *DurationRange                             `xml:"ResponseTimeRange,omitempty" json:"ResponseTimeRange,omitempty" yaml:"ResponseTimeRange,omitempty"`
	Extension         *IrCutFilterAutoAdjustmentOptionsExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type IrCutFilterAutoAdjustmentOptionsExtension []interface{}

type ItemList struct {
	SimpleItem  []*string          `xml:"SimpleItem,omitempty" json:"SimpleItem,omitempty" yaml:"SimpleItem,omitempty"`
	ElementItem *interface{}       `xml:"ElementItem,omitempty" json:"ElementItem,omitempty" yaml:"ElementItem,omitempty"`
	Extension   *ItemListExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

// Describes a list of items. Each item in the list shall have
// a                         unique name. The list is designed
// as linear structure without optional or
//     unbounded elements. Use ElementItems only when complex structures
// are                         inevitable.
type ItemListDescription struct {
	SimpleItemDescription  []*string                     `xml:"SimpleItemDescription,omitempty" json:"SimpleItemDescription,omitempty" yaml:"SimpleItemDescription,omitempty"`
	ElementItemDescription []*string                     `xml:"ElementItemDescription,omitempty" json:"ElementItemDescription,omitempty" yaml:"ElementItemDescription,omitempty"`
	Extension              *ItemListDescriptionExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type ItemListDescriptionExtension []interface{}

type ItemListExtension []interface{}

type JpegDecOptions struct {
	ResolutionsAvailable  []*VideoResolution `xml:"ResolutionsAvailable,omitempty" json:"ResolutionsAvailable,omitempty" yaml:"ResolutionsAvailable,omitempty"`
	SupportedInputBitrate *IntRange          `xml:"SupportedInputBitrate,omitempty" json:"SupportedInputBitrate,omitempty" yaml:"SupportedInputBitrate,omitempty"`
	SupportedFrameRate    *IntRange          `xml:"SupportedFrameRate,omitempty" json:"SupportedFrameRate,omitempty" yaml:"SupportedFrameRate,omitempty"`
}

type JpegOptions struct {
	ResolutionsAvailable  []*VideoResolution `xml:"ResolutionsAvailable,omitempty" json:"ResolutionsAvailable,omitempty" yaml:"ResolutionsAvailable,omitempty"`
	FrameRateRange        *IntRange          `xml:"FrameRateRange,omitempty" json:"FrameRateRange,omitempty" yaml:"FrameRateRange,omitempty"`
	EncodingIntervalRange *IntRange          `xml:"EncodingIntervalRange,omitempty" json:"EncodingIntervalRange,omitempty" yaml:"EncodingIntervalRange,omitempty"`
}

type JpegOptions2 struct {
	ResolutionsAvailable  []*VideoResolution `xml:"ResolutionsAvailable,omitempty" json:"ResolutionsAvailable,omitempty" yaml:"ResolutionsAvailable,omitempty"`
	FrameRateRange        *IntRange          `xml:"FrameRateRange,omitempty" json:"FrameRateRange,omitempty" yaml:"FrameRateRange,omitempty"`
	EncodingIntervalRange *IntRange          `xml:"EncodingIntervalRange,omitempty" json:"EncodingIntervalRange,omitempty" yaml:"EncodingIntervalRange,omitempty"`
	BitrateRange          *IntRange          `xml:"BitrateRange,omitempty" json:"BitrateRange,omitempty" yaml:"BitrateRange,omitempty"`
	TypeAttrXSI           string             `xml:"xsi:type,attr,omitempty"`
	TypeNamespace         string             `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *JpegOptions2) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:JpegOptions2"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://www.onvif.org/ver10/schema"
	}
}

// A layout describes a set of Video windows that are displayed
//                         simultaniously on a display.
type Layout struct {
	PaneLayout []*PaneLayout    `xml:"PaneLayout,omitempty" json:"PaneLayout,omitempty" yaml:"PaneLayout,omitempty"`
	Extension  *LayoutExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type LayoutExtension []interface{}

// The options supported for a display layout.
type LayoutOptions struct {
	PaneLayoutOptions []*PaneLayoutOptions    `xml:"PaneLayoutOptions,omitempty" json:"PaneLayoutOptions,omitempty" yaml:"PaneLayoutOptions,omitempty"`
	Extension         *LayoutOptionsExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type LayoutOptionsExtension []interface{}

type LensDescription struct {
	Offset      *LensOffset       `xml:"Offset,omitempty" json:"Offset,omitempty" yaml:"Offset,omitempty"`
	Projection  []*LensProjection `xml:"Projection,omitempty" json:"Projection,omitempty" yaml:"Projection,omitempty"`
	XFactor     *float64          `xml:"XFactor,omitempty" json:"XFactor,omitempty" yaml:"XFactor,omitempty"`
	FocalLength float64           `xml:"FocalLength,attr,omitempty" json:"FocalLength,attr,omitempty" yaml:"FocalLength,attr,omitempty"`
}

type LensOffset struct {
	X float64 `xml:"x,attr,omitempty" json:"x,attr,omitempty" yaml:"x,attr,omitempty"`
	Y float64 `xml:"y,attr,omitempty" json:"y,attr,omitempty" yaml:"y,attr,omitempty"`
}

type LensProjection struct {
	Angle         *float64 `xml:"Angle,omitempty" json:"Angle,omitempty" yaml:"Angle,omitempty"`
	Radius        *float64 `xml:"Radius,omitempty" json:"Radius,omitempty" yaml:"Radius,omitempty"`
	Transmittance *float64 `xml:"Transmittance,omitempty" json:"Transmittance,omitempty" yaml:"Transmittance,omitempty"`
}

type MaximumNumberOfOSDs struct {
	Total       int `xml:"Total,attr,omitempty" json:"Total,attr,omitempty" yaml:"Total,attr,omitempty"`
	Image       int `xml:"Image,attr,omitempty" json:"Image,attr,omitempty" yaml:"Image,attr,omitempty"`
	PlainText   int `xml:"PlainText,attr,omitempty" json:"PlainText,attr,omitempty" yaml:"PlainText,attr,omitempty"`
	Date        int `xml:"Date,attr,omitempty" json:"Date,attr,omitempty" yaml:"Date,attr,omitempty"`
	Time        int `xml:"Time,attr,omitempty" json:"Time,attr,omitempty" yaml:"Time,attr,omitempty"`
	DateAndTime int `xml:"DateAndTime,attr,omitempty" json:"DateAndTime,attr,omitempty" yaml:"DateAndTime,attr,omitempty"`
}

// A set of media attributes valid for a recording at a point in
//                         time or for a time interval.
type MediaAttributes struct {
	RecordingToken  *RecordingReference `xml:"RecordingToken,omitempty" json:"RecordingToken,omitempty" yaml:"RecordingToken,omitempty"`
	TrackAttributes []*TrackAttributes  `xml:"TrackAttributes,omitempty" json:"TrackAttributes,omitempty" yaml:"TrackAttributes,omitempty"`
	From            *DateTime           `xml:"From,omitempty" json:"From,omitempty" yaml:"From,omitempty"`
	Until           *DateTime           `xml:"Until,omitempty" json:"Until,omitempty" yaml:"Until,omitempty"`
}

type MediaCapabilities struct {
	XAddr                 *string                        `xml:"XAddr,omitempty" json:"XAddr,omitempty" yaml:"XAddr,omitempty"`
	StreamingCapabilities *RealTimeStreamingCapabilities `xml:"StreamingCapabilities,omitempty" json:"StreamingCapabilities,omitempty" yaml:"StreamingCapabilities,omitempty"`
	Extension             *MediaCapabilitiesExtension    `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type MediaCapabilitiesExtension struct {
	ProfileCapabilities *ProfileCapabilities `xml:"ProfileCapabilities,omitempty" json:"ProfileCapabilities,omitempty" yaml:"ProfileCapabilities,omitempty"`
}

type MediaUri struct {
	Uri                 *string   `xml:"Uri,omitempty" json:"Uri,omitempty" yaml:"Uri,omitempty"`
	InvalidAfterConnect *bool     `xml:"InvalidAfterConnect,omitempty" json:"InvalidAfterConnect,omitempty" yaml:"InvalidAfterConnect,omitempty"`
	InvalidAfterReboot  *bool     `xml:"InvalidAfterReboot,omitempty" json:"InvalidAfterReboot,omitempty" yaml:"InvalidAfterReboot,omitempty"`
	Timeout             *Duration `xml:"Timeout,omitempty" json:"Timeout,omitempty" yaml:"Timeout,omitempty"`
}

type Message struct {
	Source            *ItemList         `xml:"Source,omitempty" json:"Source,omitempty" yaml:"Source,omitempty"`
	Key               *ItemList         `xml:"Key,omitempty" json:"Key,omitempty" yaml:"Key,omitempty"`
	Data              *ItemList         `xml:"Data,omitempty" json:"Data,omitempty" yaml:"Data,omitempty"`
	Extension         *MessageExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
	UtcTime           DateTime          `xml:"UtcTime,attr,omitempty" json:"UtcTime,attr,omitempty" yaml:"UtcTime,attr,omitempty"`
	PropertyOperation PropertyOperation `xml:"PropertyOperation,attr,omitempty" json:"PropertyOperation,attr,omitempty" yaml:"PropertyOperation,attr,omitempty"`
}

type MessageDescription struct {
	Source     *ItemListDescription         `xml:"Source,omitempty" json:"Source,omitempty" yaml:"Source,omitempty"`
	Key        *ItemListDescription         `xml:"Key,omitempty" json:"Key,omitempty" yaml:"Key,omitempty"`
	Data       *ItemListDescription         `xml:"Data,omitempty" json:"Data,omitempty" yaml:"Data,omitempty"`
	Extension  *MessageDescriptionExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
	IsProperty bool                         `xml:"IsProperty,attr,omitempty" json:"IsProperty,attr,omitempty" yaml:"IsProperty,attr,omitempty"`
}

type MessageDescriptionExtension []interface{}

type MessageExtension []interface{}

type MetadataAttributes struct {
	CanContainPTZ           *bool          `xml:"CanContainPTZ,omitempty" json:"CanContainPTZ,omitempty" yaml:"CanContainPTZ,omitempty"`
	CanContainAnalytics     *bool          `xml:"CanContainAnalytics,omitempty" json:"CanContainAnalytics,omitempty" yaml:"CanContainAnalytics,omitempty"`
	CanContainNotifications *bool          `xml:"CanContainNotifications,omitempty" json:"CanContainNotifications,omitempty" yaml:"CanContainNotifications,omitempty"`
	PtzSpaces               StringAttrList `xml:"PtzSpaces,attr,omitempty" json:"PtzSpaces,attr,omitempty" yaml:"PtzSpaces,attr,omitempty"`
}

type MetadataConfiguration struct {
	Name                         *Name                           `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	UseCount                     *int                            `xml:"UseCount,omitempty" json:"UseCount,omitempty" yaml:"UseCount,omitempty"`
	Token                        ReferenceToken                  `xml:"token,attr,omitempty" json:"token,attr,omitempty" yaml:"token,attr,omitempty"`
	CompressionType              string                          `xml:"CompressionType,attr,omitempty" json:"CompressionType,attr,omitempty" yaml:"CompressionType,attr,omitempty"`
	GeoLocation                  bool                            `xml:"GeoLocation,attr,omitempty" json:"GeoLocation,attr,omitempty" yaml:"GeoLocation,attr,omitempty"`
	ShapePolygon                 bool                            `xml:"ShapePolygon,attr,omitempty" json:"ShapePolygon,attr,omitempty" yaml:"ShapePolygon,attr,omitempty"`
	PTZStatus                    *PTZFilter                      `xml:"PTZStatus,omitempty" json:"PTZStatus,omitempty" yaml:"PTZStatus,omitempty"`
	Events                       *EventSubscription              `xml:"Events,omitempty" json:"Events,omitempty" yaml:"Events,omitempty"`
	Analytics                    *bool                           `xml:"Analytics,omitempty" json:"Analytics,omitempty" yaml:"Analytics,omitempty"`
	Multicast                    *MulticastConfiguration         `xml:"Multicast,omitempty" json:"Multicast,omitempty" yaml:"Multicast,omitempty"`
	SessionTimeout               *Duration                       `xml:"SessionTimeout,omitempty" json:"SessionTimeout,omitempty" yaml:"SessionTimeout,omitempty"`
	AnalyticsEngineConfiguration *AnalyticsEngineConfiguration   `xml:"AnalyticsEngineConfiguration,omitempty" json:"AnalyticsEngineConfiguration,omitempty" yaml:"AnalyticsEngineConfiguration,omitempty"`
	Extension                    *MetadataConfigurationExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
	TypeAttrXSI                  string                          `xml:"xsi:type,attr,omitempty"`
	TypeNamespace                string                          `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *MetadataConfiguration) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:MetadataConfiguration"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://www.onvif.org/ver10/schema"
	}
}

type MetadataConfigurationExtension []interface{}

type MetadataConfigurationOptions struct {
	PTZStatusFilterOptions *PTZStatusFilterOptions                `xml:"PTZStatusFilterOptions,omitempty" json:"PTZStatusFilterOptions,omitempty" yaml:"PTZStatusFilterOptions,omitempty"`
	Extension              *MetadataConfigurationOptionsExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
	GeoLocation            bool                                   `xml:"GeoLocation,attr,omitempty" json:"GeoLocation,attr,omitempty" yaml:"GeoLocation,attr,omitempty"`
	MaxContentFilterSize   int                                    `xml:"MaxContentFilterSize,attr,omitempty" json:"MaxContentFilterSize,attr,omitempty" yaml:"MaxContentFilterSize,attr,omitempty"`
}

type MetadataConfigurationOptionsExtension struct {
	CompressionType []*string                               `xml:"CompressionType,omitempty" json:"CompressionType,omitempty" yaml:"CompressionType,omitempty"`
	Extension       *MetadataConfigurationOptionsExtension2 `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type MetadataConfigurationOptionsExtension2 []interface{}

type MetadataFilter struct {
	MetadataStreamFilter *XPathExpression `xml:"MetadataStreamFilter,omitempty" json:"MetadataStreamFilter,omitempty" yaml:"MetadataStreamFilter,omitempty"`
}

type MetadataInput struct {
	MetadataConfig []*Config               `xml:"MetadataConfig,omitempty" json:"MetadataConfig,omitempty" yaml:"MetadataConfig,omitempty"`
	Extension      *MetadataInputExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type MetadataInputExtension []interface{}

type MetadataType []interface{}

type MotionExpression struct {
	Expression *string `xml:"Expression,omitempty" json:"Expression,omitempty" yaml:"Expression,omitempty"`
	Type       string  `xml:"Type,attr,omitempty" json:"Type,attr,omitempty" yaml:"Type,attr,omitempty"`
}

type MotionExpressionConfiguration struct {
	MotionExpression *MotionExpression `xml:"MotionExpression,omitempty" json:"MotionExpression,omitempty" yaml:"MotionExpression,omitempty"`
}

type MoveOptions struct {
	Absolute   *AbsoluteFocusOptions   `xml:"Absolute,omitempty" json:"Absolute,omitempty" yaml:"Absolute,omitempty"`
	Relative   *RelativeFocusOptions   `xml:"Relative,omitempty" json:"Relative,omitempty" yaml:"Relative,omitempty"`
	Continuous *ContinuousFocusOptions `xml:"Continuous,omitempty" json:"Continuous,omitempty" yaml:"Continuous,omitempty"`
}

type MoveOptions20 struct {
	Absolute   *AbsoluteFocusOptions   `xml:"Absolute,omitempty" json:"Absolute,omitempty" yaml:"Absolute,omitempty"`
	Relative   *RelativeFocusOptions20 `xml:"Relative,omitempty" json:"Relative,omitempty" yaml:"Relative,omitempty"`
	Continuous *ContinuousFocusOptions `xml:"Continuous,omitempty" json:"Continuous,omitempty" yaml:"Continuous,omitempty"`
}

type Mpeg4Configuration struct {
	GovLength    *int          `xml:"GovLength,omitempty" json:"GovLength,omitempty" yaml:"GovLength,omitempty"`
	Mpeg4Profile *Mpeg4Profile `xml:"Mpeg4Profile,omitempty" json:"Mpeg4Profile,omitempty" yaml:"Mpeg4Profile,omitempty"`
}

type Mpeg4DecOptions struct {
	ResolutionsAvailable   []*VideoResolution `xml:"ResolutionsAvailable,omitempty" json:"ResolutionsAvailable,omitempty" yaml:"ResolutionsAvailable,omitempty"`
	SupportedMpeg4Profiles []*Mpeg4Profile    `xml:"SupportedMpeg4Profiles,omitempty" json:"SupportedMpeg4Profiles,omitempty" yaml:"SupportedMpeg4Profiles,omitempty"`
	SupportedInputBitrate  *IntRange          `xml:"SupportedInputBitrate,omitempty" json:"SupportedInputBitrate,omitempty" yaml:"SupportedInputBitrate,omitempty"`
	SupportedFrameRate     *IntRange          `xml:"SupportedFrameRate,omitempty" json:"SupportedFrameRate,omitempty" yaml:"SupportedFrameRate,omitempty"`
}

type Mpeg4Options struct {
	ResolutionsAvailable   []*VideoResolution `xml:"ResolutionsAvailable,omitempty" json:"ResolutionsAvailable,omitempty" yaml:"ResolutionsAvailable,omitempty"`
	GovLengthRange         *IntRange          `xml:"GovLengthRange,omitempty" json:"GovLengthRange,omitempty" yaml:"GovLengthRange,omitempty"`
	FrameRateRange         *IntRange          `xml:"FrameRateRange,omitempty" json:"FrameRateRange,omitempty" yaml:"FrameRateRange,omitempty"`
	EncodingIntervalRange  *IntRange          `xml:"EncodingIntervalRange,omitempty" json:"EncodingIntervalRange,omitempty" yaml:"EncodingIntervalRange,omitempty"`
	Mpeg4ProfilesSupported []*Mpeg4Profile    `xml:"Mpeg4ProfilesSupported,omitempty" json:"Mpeg4ProfilesSupported,omitempty" yaml:"Mpeg4ProfilesSupported,omitempty"`
}

type Mpeg4Options2 struct {
	ResolutionsAvailable   []*VideoResolution `xml:"ResolutionsAvailable,omitempty" json:"ResolutionsAvailable,omitempty" yaml:"ResolutionsAvailable,omitempty"`
	GovLengthRange         *IntRange          `xml:"GovLengthRange,omitempty" json:"GovLengthRange,omitempty" yaml:"GovLengthRange,omitempty"`
	FrameRateRange         *IntRange          `xml:"FrameRateRange,omitempty" json:"FrameRateRange,omitempty" yaml:"FrameRateRange,omitempty"`
	EncodingIntervalRange  *IntRange          `xml:"EncodingIntervalRange,omitempty" json:"EncodingIntervalRange,omitempty" yaml:"EncodingIntervalRange,omitempty"`
	Mpeg4ProfilesSupported []*Mpeg4Profile    `xml:"Mpeg4ProfilesSupported,omitempty" json:"Mpeg4ProfilesSupported,omitempty" yaml:"Mpeg4ProfilesSupported,omitempty"`
	BitrateRange           *IntRange          `xml:"BitrateRange,omitempty" json:"BitrateRange,omitempty" yaml:"BitrateRange,omitempty"`
	TypeAttrXSI            string             `xml:"xsi:type,attr,omitempty"`
	TypeNamespace          string             `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *Mpeg4Options2) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:Mpeg4Options2"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://www.onvif.org/ver10/schema"
	}
}

type MulticastConfiguration struct {
	Address   *IPAddress `xml:"Address,omitempty" json:"Address,omitempty" yaml:"Address,omitempty"`
	Port      *int       `xml:"Port,omitempty" json:"Port,omitempty" yaml:"Port,omitempty"`
	TTL       *int       `xml:"TTL,omitempty" json:"TTL,omitempty" yaml:"TTL,omitempty"`
	AutoStart *bool      `xml:"AutoStart,omitempty" json:"AutoStart,omitempty" yaml:"AutoStart,omitempty"`
}

type MultipleTopicsSpecifiedFaultType struct {
	Timestamp     DateTime               `xml:"Timestamp" json:"Timestamp" yaml:"Timestamp"`
	Originator    *EndpointReferenceType `xml:"Originator,omitempty" json:"Originator,omitempty" yaml:"Originator,omitempty"`
	ErrorCode     *string                `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	Description   []*string              `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	FaultCause    interface{}            `xml:"FaultCause" json:"FaultCause" yaml:"FaultCause"`
	TypeAttrXSI   string                 `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string                 `xml:"xmlns:objtype,attr,omitempty"`

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

type NTPInformation struct {
	FromDHCP    *bool                    `xml:"FromDHCP,omitempty" json:"FromDHCP,omitempty" yaml:"FromDHCP,omitempty"`
	NTPFromDHCP []*NetworkHost           `xml:"NTPFromDHCP,omitempty" json:"NTPFromDHCP,omitempty" yaml:"NTPFromDHCP,omitempty"`
	NTPManual   []*NetworkHost           `xml:"NTPManual,omitempty" json:"NTPManual,omitempty" yaml:"NTPManual,omitempty"`
	Extension   *NTPInformationExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type NTPInformationExtension []interface{}

type NetworkCapabilities struct {
	IPFilter          *bool                         `xml:"IPFilter,omitempty" json:"IPFilter,omitempty" yaml:"IPFilter,omitempty"`
	ZeroConfiguration *bool                         `xml:"ZeroConfiguration,omitempty" json:"ZeroConfiguration,omitempty" yaml:"ZeroConfiguration,omitempty"`
	IPVersion6        *bool                         `xml:"IPVersion6,omitempty" json:"IPVersion6,omitempty" yaml:"IPVersion6,omitempty"`
	DynDNS            *bool                         `xml:"DynDNS,omitempty" json:"DynDNS,omitempty" yaml:"DynDNS,omitempty"`
	Extension         *NetworkCapabilitiesExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type NetworkCapabilitiesExtension struct {
	Dot11Configuration *bool                          `xml:"Dot11Configuration,omitempty" json:"Dot11Configuration,omitempty" yaml:"Dot11Configuration,omitempty"`
	Extension          *NetworkCapabilitiesExtension2 `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type NetworkCapabilitiesExtension2 []interface{}

type NetworkGateway struct {
	IPv4Address []*IPv4Address `xml:"IPv4Address,omitempty" json:"IPv4Address,omitempty" yaml:"IPv4Address,omitempty"`
	IPv6Address []*IPv6Address `xml:"IPv6Address,omitempty" json:"IPv6Address,omitempty" yaml:"IPv6Address,omitempty"`
}

type NetworkHost struct {
	Type        *NetworkHostType      `xml:"Type,omitempty" json:"Type,omitempty" yaml:"Type,omitempty"`
	IPv4Address *IPv4Address          `xml:"IPv4Address,omitempty" json:"IPv4Address,omitempty" yaml:"IPv4Address,omitempty"`
	IPv6Address *IPv6Address          `xml:"IPv6Address,omitempty" json:"IPv6Address,omitempty" yaml:"IPv6Address,omitempty"`
	DNSname     *DNSName              `xml:"DNSname,omitempty" json:"DNSname,omitempty" yaml:"DNSname,omitempty"`
	Extension   *NetworkHostExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type NetworkHostExtension []interface{}

type NetworkInterface struct {
	Token         ReferenceToken             `xml:"token,attr,omitempty" json:"token,attr,omitempty" yaml:"token,attr,omitempty"`
	Enabled       *bool                      `xml:"Enabled,omitempty" json:"Enabled,omitempty" yaml:"Enabled,omitempty"`
	Info          *NetworkInterfaceInfo      `xml:"Info,omitempty" json:"Info,omitempty" yaml:"Info,omitempty"`
	Link          *NetworkInterfaceLink      `xml:"Link,omitempty" json:"Link,omitempty" yaml:"Link,omitempty"`
	IPv4          *IPv4NetworkInterface      `xml:"IPv4,omitempty" json:"IPv4,omitempty" yaml:"IPv4,omitempty"`
	IPv6          *IPv6NetworkInterface      `xml:"IPv6,omitempty" json:"IPv6,omitempty" yaml:"IPv6,omitempty"`
	Extension     *NetworkInterfaceExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
	TypeAttrXSI   string                     `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string                     `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *NetworkInterface) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:NetworkInterface"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://www.onvif.org/ver10/schema"
	}
}

type NetworkInterfaceConnectionSetting struct {
	AutoNegotiation *bool   `xml:"AutoNegotiation,omitempty" json:"AutoNegotiation,omitempty" yaml:"AutoNegotiation,omitempty"`
	Speed           *int    `xml:"Speed,omitempty" json:"Speed,omitempty" yaml:"Speed,omitempty"`
	Duplex          *Duplex `xml:"Duplex,omitempty" json:"Duplex,omitempty" yaml:"Duplex,omitempty"`
}

type NetworkInterfaceExtension struct {
	InterfaceType *IANAIfTypes                `xml:"InterfaceType,omitempty" json:"InterfaceType,omitempty" yaml:"InterfaceType,omitempty"`
	Dot3          []*Dot3Configuration        `xml:"Dot3,omitempty" json:"Dot3,omitempty" yaml:"Dot3,omitempty"`
	Dot11         []*Dot11Configuration       `xml:"Dot11,omitempty" json:"Dot11,omitempty" yaml:"Dot11,omitempty"`
	Extension     *NetworkInterfaceExtension2 `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type NetworkInterfaceExtension2 []interface{}

type NetworkInterfaceInfo struct {
	Name      *string    `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	HwAddress *HwAddress `xml:"HwAddress,omitempty" json:"HwAddress,omitempty" yaml:"HwAddress,omitempty"`
	MTU       *int       `xml:"MTU,omitempty" json:"MTU,omitempty" yaml:"MTU,omitempty"`
}

type NetworkInterfaceLink struct {
	AdminSettings *NetworkInterfaceConnectionSetting `xml:"AdminSettings,omitempty" json:"AdminSettings,omitempty" yaml:"AdminSettings,omitempty"`
	OperSettings  *NetworkInterfaceConnectionSetting `xml:"OperSettings,omitempty" json:"OperSettings,omitempty" yaml:"OperSettings,omitempty"`
	InterfaceType *IANAIfTypes                       `xml:"InterfaceType,omitempty" json:"InterfaceType,omitempty" yaml:"InterfaceType,omitempty"`
}

type NetworkInterfaceSetConfiguration struct {
	Enabled   *bool                                      `xml:"Enabled,omitempty" json:"Enabled,omitempty" yaml:"Enabled,omitempty"`
	Link      *NetworkInterfaceConnectionSetting         `xml:"Link,omitempty" json:"Link,omitempty" yaml:"Link,omitempty"`
	MTU       *int                                       `xml:"MTU,omitempty" json:"MTU,omitempty" yaml:"MTU,omitempty"`
	IPv4      *IPv4NetworkInterfaceSetConfiguration      `xml:"IPv4,omitempty" json:"IPv4,omitempty" yaml:"IPv4,omitempty"`
	IPv6      *IPv6NetworkInterfaceSetConfiguration      `xml:"IPv6,omitempty" json:"IPv6,omitempty" yaml:"IPv6,omitempty"`
	Extension *NetworkInterfaceSetConfigurationExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type NetworkInterfaceSetConfigurationExtension struct {
	Dot3      []*Dot3Configuration                        `xml:"Dot3,omitempty" json:"Dot3,omitempty" yaml:"Dot3,omitempty"`
	Dot11     []*Dot11Configuration                       `xml:"Dot11,omitempty" json:"Dot11,omitempty" yaml:"Dot11,omitempty"`
	Extension *NetworkInterfaceSetConfigurationExtension2 `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type NetworkInterfaceSetConfigurationExtension2 []interface{}

type NetworkProtocol struct {
	Name      *NetworkProtocolType      `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Enabled   *bool                     `xml:"Enabled,omitempty" json:"Enabled,omitempty" yaml:"Enabled,omitempty"`
	Port      []*int                    `xml:"Port,omitempty" json:"Port,omitempty" yaml:"Port,omitempty"`
	Extension *NetworkProtocolExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type NetworkProtocolExtension []interface{}

type NetworkZeroConfiguration struct {
	InterfaceToken ReferenceToken                     `xml:"InterfaceToken,omitempty" json:"InterfaceToken,omitempty" yaml:"InterfaceToken,omitempty"`
	Enabled        *bool                              `xml:"Enabled,omitempty" json:"Enabled,omitempty" yaml:"Enabled,omitempty"`
	Addresses      []*IPv4Address                     `xml:"Addresses,omitempty" json:"Addresses,omitempty" yaml:"Addresses,omitempty"`
	Extension      *NetworkZeroConfigurationExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type NetworkZeroConfigurationExtension struct {
	Additional []*NetworkZeroConfiguration         `xml:"Additional,omitempty" json:"Additional,omitempty" yaml:"Additional,omitempty"`
	Extension  *NetworkZeroConfigurationExtension2 `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type NetworkZeroConfigurationExtension2 []interface{}

type NoCurrentMessageOnTopicFaultType struct {
	Timestamp     DateTime               `xml:"Timestamp" json:"Timestamp" yaml:"Timestamp"`
	Originator    *EndpointReferenceType `xml:"Originator,omitempty" json:"Originator,omitempty" yaml:"Originator,omitempty"`
	ErrorCode     *string                `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	Description   []*string              `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	FaultCause    interface{}            `xml:"FaultCause" json:"FaultCause" yaml:"FaultCause"`
	TypeAttrXSI   string                 `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string                 `xml:"xmlns:objtype,attr,omitempty"`

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

type NoiseReduction struct {
	Level *float64 `xml:"Level,omitempty" json:"Level,omitempty" yaml:"Level,omitempty"`
}

type NoiseReductionOptions struct {
	Level *bool `xml:"Level,omitempty" json:"Level,omitempty" yaml:"Level,omitempty"`
}

type NotUnderstoodType struct {
	Qname string `xml:"qname,attr,omitempty" json:"qname,attr,omitempty" yaml:"qname,attr,omitempty"`
}

type NotificationMessageHolderType struct {
	SubscriptionReference *EndpointReferenceType `xml:"SubscriptionReference,omitempty" json:"SubscriptionReference,omitempty" yaml:"SubscriptionReference,omitempty"`
	Topic                 *TopicExpressionType   `xml:"Topic,omitempty" json:"Topic,omitempty" yaml:"Topic,omitempty"`
	ProducerReference     *EndpointReferenceType `xml:"ProducerReference,omitempty" json:"ProducerReference,omitempty" yaml:"ProducerReference,omitempty"`
	Message               interface{}            `xml:"Message" json:"Message" yaml:"Message"`
}

type NotificationProducerRP struct {
	TopicExpression        *TopicExpressionType `xml:"TopicExpression,omitempty" json:"TopicExpression,omitempty" yaml:"TopicExpression,omitempty"`
	FixedTopicSet          *bool                `xml:"FixedTopicSet,omitempty" json:"FixedTopicSet,omitempty" yaml:"FixedTopicSet,omitempty"`
	TopicExpressionDialect *string              `xml:"TopicExpressionDialect,omitempty" json:"TopicExpressionDialect,omitempty" yaml:"TopicExpressionDialect,omitempty"`
	TopicSet               *TopicSetType        `xml:"TopicSet,omitempty" json:"TopicSet,omitempty" yaml:"TopicSet,omitempty"`
}

type Notify struct {
	NotificationMessage *NotificationMessageHolderType `xml:"NotificationMessage,omitempty" json:"NotificationMessage,omitempty" yaml:"NotificationMessage,omitempty"`
}

type NotifyMessageNotSupportedFaultType struct {
	Timestamp     DateTime               `xml:"Timestamp" json:"Timestamp" yaml:"Timestamp"`
	Originator    *EndpointReferenceType `xml:"Originator,omitempty" json:"Originator,omitempty" yaml:"Originator,omitempty"`
	ErrorCode     *string                `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	Description   []*string              `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	FaultCause    interface{}            `xml:"FaultCause" json:"FaultCause" yaml:"FaultCause"`
	TypeAttrXSI   string                 `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string                 `xml:"xmlns:objtype,attr,omitempty"`

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

// The value range of "Transparent" could be defined by vendors
//                         only should follow this rule: the minimum
// value means non-transparent and                         the
// maximum value maens fully transparent.
type OSDColor struct {
	Color       *Color `xml:"Color,omitempty" json:"Color,omitempty" yaml:"Color,omitempty"`
	Transparent int    `xml:"Transparent,attr,omitempty" json:"Transparent,attr,omitempty" yaml:"Transparent,attr,omitempty"`
}

// Describe the option of the color and its transparency.
type OSDColorOptions struct {
	Color       *ColorOptions             `xml:"Color,omitempty" json:"Color,omitempty" yaml:"Color,omitempty"`
	Transparent *IntRange                 `xml:"Transparent,omitempty" json:"Transparent,omitempty" yaml:"Transparent,omitempty"`
	Extension   *OSDColorOptionsExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type OSDColorOptionsExtension []interface{}

type OSDConfiguration struct {
	Token                         ReferenceToken             `xml:"token,attr,omitempty" json:"token,attr,omitempty" yaml:"token,attr,omitempty"`
	VideoSourceConfigurationToken *OSDReference              `xml:"VideoSourceConfigurationToken,omitempty" json:"VideoSourceConfigurationToken,omitempty" yaml:"VideoSourceConfigurationToken,omitempty"`
	Type                          *OSDType                   `xml:"Type,omitempty" json:"Type,omitempty" yaml:"Type,omitempty"`
	Position                      *OSDPosConfiguration       `xml:"Position,omitempty" json:"Position,omitempty" yaml:"Position,omitempty"`
	TextString                    *OSDTextConfiguration      `xml:"TextString,omitempty" json:"TextString,omitempty" yaml:"TextString,omitempty"`
	Image                         *OSDImgConfiguration       `xml:"Image,omitempty" json:"Image,omitempty" yaml:"Image,omitempty"`
	Extension                     *OSDConfigurationExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
	TypeAttrXSI                   string                     `xml:"xsi:type,attr,omitempty"`
	TypeNamespace                 string                     `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *OSDConfiguration) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:OSDConfiguration"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://www.onvif.org/ver10/schema"
	}
}

type OSDConfigurationExtension []interface{}

type OSDConfigurationOptions struct {
	MaximumNumberOfOSDs *MaximumNumberOfOSDs              `xml:"MaximumNumberOfOSDs,omitempty" json:"MaximumNumberOfOSDs,omitempty" yaml:"MaximumNumberOfOSDs,omitempty"`
	Type                []*OSDType                        `xml:"Type,omitempty" json:"Type,omitempty" yaml:"Type,omitempty"`
	PositionOption      []*string                         `xml:"PositionOption,omitempty" json:"PositionOption,omitempty" yaml:"PositionOption,omitempty"`
	TextOption          *OSDTextOptions                   `xml:"TextOption,omitempty" json:"TextOption,omitempty" yaml:"TextOption,omitempty"`
	ImageOption         *OSDImgOptions                    `xml:"ImageOption,omitempty" json:"ImageOption,omitempty" yaml:"ImageOption,omitempty"`
	Extension           *OSDConfigurationOptionsExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type OSDConfigurationOptionsExtension []interface{}

type OSDImgConfiguration struct {
	ImgPath   *string                       `xml:"ImgPath,omitempty" json:"ImgPath,omitempty" yaml:"ImgPath,omitempty"`
	Extension *OSDImgConfigurationExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type OSDImgConfigurationExtension []interface{}

type OSDImgOptions struct {
	ImagePath        []*string               `xml:"ImagePath,omitempty" json:"ImagePath,omitempty" yaml:"ImagePath,omitempty"`
	Extension        *OSDImgOptionsExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
	FormatsSupported StringAttrList          `xml:"FormatsSupported,attr,omitempty" json:"FormatsSupported,attr,omitempty" yaml:"FormatsSupported,attr,omitempty"`
	MaxSize          int                     `xml:"MaxSize,attr,omitempty" json:"MaxSize,attr,omitempty" yaml:"MaxSize,attr,omitempty"`
	MaxWidth         int                     `xml:"MaxWidth,attr,omitempty" json:"MaxWidth,attr,omitempty" yaml:"MaxWidth,attr,omitempty"`
	MaxHeight        int                     `xml:"MaxHeight,attr,omitempty" json:"MaxHeight,attr,omitempty" yaml:"MaxHeight,attr,omitempty"`
}

type OSDImgOptionsExtension []interface{}

type OSDPosConfiguration struct {
	Type      *string                       `xml:"Type,omitempty" json:"Type,omitempty" yaml:"Type,omitempty"`
	Pos       *Vector                       `xml:"Pos,omitempty" json:"Pos,omitempty" yaml:"Pos,omitempty"`
	Extension *OSDPosConfigurationExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type OSDPosConfigurationExtension []interface{}

type OSDReference struct {
	Content ReferenceToken `xml:"Content,omitempty" json:"Content,omitempty" yaml:"Content,omitempty"`
}

type OSDTextConfiguration struct {
	Type             *string                        `xml:"Type,omitempty" json:"Type,omitempty" yaml:"Type,omitempty"`
	DateFormat       *string                        `xml:"DateFormat,omitempty" json:"DateFormat,omitempty" yaml:"DateFormat,omitempty"`
	TimeFormat       *string                        `xml:"TimeFormat,omitempty" json:"TimeFormat,omitempty" yaml:"TimeFormat,omitempty"`
	FontSize         *int                           `xml:"FontSize,omitempty" json:"FontSize,omitempty" yaml:"FontSize,omitempty"`
	FontColor        *OSDColor                      `xml:"FontColor,omitempty" json:"FontColor,omitempty" yaml:"FontColor,omitempty"`
	BackgroundColor  *OSDColor                      `xml:"BackgroundColor,omitempty" json:"BackgroundColor,omitempty" yaml:"BackgroundColor,omitempty"`
	PlainText        *string                        `xml:"PlainText,omitempty" json:"PlainText,omitempty" yaml:"PlainText,omitempty"`
	Extension        *OSDTextConfigurationExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
	IsPersistentText bool                           `xml:"IsPersistentText,attr,omitempty" json:"IsPersistentText,attr,omitempty" yaml:"IsPersistentText,attr,omitempty"`
}

type OSDTextConfigurationExtension []interface{}

type OSDTextOptions struct {
	Type            []*string                `xml:"Type,omitempty" json:"Type,omitempty" yaml:"Type,omitempty"`
	FontSizeRange   *IntRange                `xml:"FontSizeRange,omitempty" json:"FontSizeRange,omitempty" yaml:"FontSizeRange,omitempty"`
	DateFormat      []*string                `xml:"DateFormat,omitempty" json:"DateFormat,omitempty" yaml:"DateFormat,omitempty"`
	TimeFormat      []*string                `xml:"TimeFormat,omitempty" json:"TimeFormat,omitempty" yaml:"TimeFormat,omitempty"`
	FontColor       *OSDColorOptions         `xml:"FontColor,omitempty" json:"FontColor,omitempty" yaml:"FontColor,omitempty"`
	BackgroundColor *OSDColorOptions         `xml:"BackgroundColor,omitempty" json:"BackgroundColor,omitempty" yaml:"BackgroundColor,omitempty"`
	Extension       *OSDTextOptionsExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type OSDTextOptionsExtension []interface{}

type OnvifVersion struct {
	Major *int `xml:"Major,omitempty" json:"Major,omitempty" yaml:"Major,omitempty"`
	Minor *int `xml:"Minor,omitempty" json:"Minor,omitempty" yaml:"Minor,omitempty"`
}

type PTControlDirection struct {
	EFlip     *EFlip                       `xml:"EFlip,omitempty" json:"EFlip,omitempty" yaml:"EFlip,omitempty"`
	Reverse   *Reverse                     `xml:"Reverse,omitempty" json:"Reverse,omitempty" yaml:"Reverse,omitempty"`
	Extension *PTControlDirectionExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type PTControlDirectionExtension []interface{}

type PTControlDirectionOptions struct {
	EFlip     *EFlipOptions                       `xml:"EFlip,omitempty" json:"EFlip,omitempty" yaml:"EFlip,omitempty"`
	Reverse   *ReverseOptions                     `xml:"Reverse,omitempty" json:"Reverse,omitempty" yaml:"Reverse,omitempty"`
	Extension *PTControlDirectionOptionsExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type PTControlDirectionOptionsExtension []interface{}

type PTZCapabilities struct {
	XAddr *string `xml:"XAddr,omitempty" json:"XAddr,omitempty" yaml:"XAddr,omitempty"`
}

type PTZConfiguration struct {
	Name                                   *Name                      `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	UseCount                               *int                       `xml:"UseCount,omitempty" json:"UseCount,omitempty" yaml:"UseCount,omitempty"`
	Token                                  ReferenceToken             `xml:"token,attr,omitempty" json:"token,attr,omitempty" yaml:"token,attr,omitempty"`
	MoveRamp                               int                        `xml:"MoveRamp,attr,omitempty" json:"MoveRamp,attr,omitempty" yaml:"MoveRamp,attr,omitempty"`
	PresetRamp                             int                        `xml:"PresetRamp,attr,omitempty" json:"PresetRamp,attr,omitempty" yaml:"PresetRamp,attr,omitempty"`
	PresetTourRamp                         int                        `xml:"PresetTourRamp,attr,omitempty" json:"PresetTourRamp,attr,omitempty" yaml:"PresetTourRamp,attr,omitempty"`
	NodeToken                              ReferenceToken             `xml:"NodeToken,omitempty" json:"NodeToken,omitempty" yaml:"NodeToken,omitempty"`
	DefaultAbsolutePantTiltPositionSpace   *string                    `xml:"DefaultAbsolutePantTiltPositionSpace,omitempty" json:"DefaultAbsolutePantTiltPositionSpace,omitempty" yaml:"DefaultAbsolutePantTiltPositionSpace,omitempty"`
	DefaultAbsoluteZoomPositionSpace       *string                    `xml:"DefaultAbsoluteZoomPositionSpace,omitempty" json:"DefaultAbsoluteZoomPositionSpace,omitempty" yaml:"DefaultAbsoluteZoomPositionSpace,omitempty"`
	DefaultRelativePanTiltTranslationSpace *string                    `xml:"DefaultRelativePanTiltTranslationSpace,omitempty" json:"DefaultRelativePanTiltTranslationSpace,omitempty" yaml:"DefaultRelativePanTiltTranslationSpace,omitempty"`
	DefaultRelativeZoomTranslationSpace    *string                    `xml:"DefaultRelativeZoomTranslationSpace,omitempty" json:"DefaultRelativeZoomTranslationSpace,omitempty" yaml:"DefaultRelativeZoomTranslationSpace,omitempty"`
	DefaultContinuousPanTiltVelocitySpace  *string                    `xml:"DefaultContinuousPanTiltVelocitySpace,omitempty" json:"DefaultContinuousPanTiltVelocitySpace,omitempty" yaml:"DefaultContinuousPanTiltVelocitySpace,omitempty"`
	DefaultContinuousZoomVelocitySpace     *string                    `xml:"DefaultContinuousZoomVelocitySpace,omitempty" json:"DefaultContinuousZoomVelocitySpace,omitempty" yaml:"DefaultContinuousZoomVelocitySpace,omitempty"`
	DefaultPTZSpeed                        *PTZSpeed                  `xml:"DefaultPTZSpeed,omitempty" json:"DefaultPTZSpeed,omitempty" yaml:"DefaultPTZSpeed,omitempty"`
	DefaultPTZTimeout                      *Duration                  `xml:"DefaultPTZTimeout,omitempty" json:"DefaultPTZTimeout,omitempty" yaml:"DefaultPTZTimeout,omitempty"`
	PanTiltLimits                          *PanTiltLimits             `xml:"PanTiltLimits,omitempty" json:"PanTiltLimits,omitempty" yaml:"PanTiltLimits,omitempty"`
	ZoomLimits                             *ZoomLimits                `xml:"ZoomLimits,omitempty" json:"ZoomLimits,omitempty" yaml:"ZoomLimits,omitempty"`
	Extension                              *PTZConfigurationExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
	TypeAttrXSI                            string                     `xml:"xsi:type,attr,omitempty"`
	TypeNamespace                          string                     `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *PTZConfiguration) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:PTZConfiguration"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://www.onvif.org/ver10/schema"
	}
}

type PTZConfigurationExtension struct {
	PTControlDirection *PTControlDirection         `xml:"PTControlDirection,omitempty" json:"PTControlDirection,omitempty" yaml:"PTControlDirection,omitempty"`
	Extension          *PTZConfigurationExtension2 `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type PTZConfigurationExtension2 []interface{}

type PTZConfigurationOptions struct {
	Spaces             *PTZSpaces                 `xml:"Spaces,omitempty" json:"Spaces,omitempty" yaml:"Spaces,omitempty"`
	PTZTimeout         *DurationRange             `xml:"PTZTimeout,omitempty" json:"PTZTimeout,omitempty" yaml:"PTZTimeout,omitempty"`
	PTControlDirection *PTControlDirectionOptions `xml:"PTControlDirection,omitempty" json:"PTControlDirection,omitempty" yaml:"PTControlDirection,omitempty"`
	Extension          *PTZConfigurationOptions2  `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
	PTZRamps           *IntList                   `xml:"PTZRamps,attr,omitempty" json:"PTZRamps,attr,omitempty" yaml:"PTZRamps,attr,omitempty"`
}

type PTZConfigurationOptions2 []interface{}

type PTZFilter struct {
	Status   *bool `xml:"Status,omitempty" json:"Status,omitempty" yaml:"Status,omitempty"`
	Position *bool `xml:"Position,omitempty" json:"Position,omitempty" yaml:"Position,omitempty"`
}

type PTZNode struct {
	Token                  ReferenceToken    `xml:"token,attr,omitempty" json:"token,attr,omitempty" yaml:"token,attr,omitempty"`
	FixedHomePosition      bool              `xml:"FixedHomePosition,attr,omitempty" json:"FixedHomePosition,attr,omitempty" yaml:"FixedHomePosition,attr,omitempty"`
	GeoMove                bool              `xml:"GeoMove,attr,omitempty" json:"GeoMove,attr,omitempty" yaml:"GeoMove,attr,omitempty"`
	Name                   *Name             `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	SupportedPTZSpaces     *PTZSpaces        `xml:"SupportedPTZSpaces,omitempty" json:"SupportedPTZSpaces,omitempty" yaml:"SupportedPTZSpaces,omitempty"`
	MaximumNumberOfPresets *int              `xml:"MaximumNumberOfPresets,omitempty" json:"MaximumNumberOfPresets,omitempty" yaml:"MaximumNumberOfPresets,omitempty"`
	HomeSupported          *bool             `xml:"HomeSupported,omitempty" json:"HomeSupported,omitempty" yaml:"HomeSupported,omitempty"`
	AuxiliaryCommands      []*AuxiliaryData  `xml:"AuxiliaryCommands,omitempty" json:"AuxiliaryCommands,omitempty" yaml:"AuxiliaryCommands,omitempty"`
	Extension              *PTZNodeExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
	TypeAttrXSI            string            `xml:"xsi:type,attr,omitempty"`
	TypeNamespace          string            `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *PTZNode) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:PTZNode"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://www.onvif.org/ver10/schema"
	}
}

type PTZNodeExtension struct {
	SupportedPresetTour *PTZPresetTourSupported `xml:"SupportedPresetTour,omitempty" json:"SupportedPresetTour,omitempty" yaml:"SupportedPresetTour,omitempty"`
	Extension           *PTZNodeExtension2      `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type PTZNodeExtension2 []interface{}

type PTZPositionFilter struct {
	MinPosition *PTZVector `xml:"MinPosition,omitempty" json:"MinPosition,omitempty" yaml:"MinPosition,omitempty"`
	MaxPosition *PTZVector `xml:"MaxPosition,omitempty" json:"MaxPosition,omitempty" yaml:"MaxPosition,omitempty"`
	EnterOrExit *bool      `xml:"EnterOrExit,omitempty" json:"EnterOrExit,omitempty" yaml:"EnterOrExit,omitempty"`
}

type PTZPreset struct {
	Name        *Name          `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	PTZPosition *PTZVector     `xml:"PTZPosition,omitempty" json:"PTZPosition,omitempty" yaml:"PTZPosition,omitempty"`
	Token       ReferenceToken `xml:"token,attr,omitempty" json:"token,attr,omitempty" yaml:"token,attr,omitempty"`
}

type PTZPresetTourExtension []interface{}

type PTZPresetTourOptions struct {
	AutoStart         *bool                                  `xml:"AutoStart,omitempty" json:"AutoStart,omitempty" yaml:"AutoStart,omitempty"`
	StartingCondition *PTZPresetTourStartingConditionOptions `xml:"StartingCondition,omitempty" json:"StartingCondition,omitempty" yaml:"StartingCondition,omitempty"`
	TourSpot          *PTZPresetTourSpotOptions              `xml:"TourSpot,omitempty" json:"TourSpot,omitempty" yaml:"TourSpot,omitempty"`
}

type PTZPresetTourPresetDetail []interface{}

type PTZPresetTourPresetDetailOptions struct {
	PresetToken          []ReferenceToken                           `xml:"PresetToken,omitempty" json:"PresetToken,omitempty" yaml:"PresetToken,omitempty"`
	Home                 *bool                                      `xml:"Home,omitempty" json:"Home,omitempty" yaml:"Home,omitempty"`
	PanTiltPositionSpace *Space2DDescription                        `xml:"PanTiltPositionSpace,omitempty" json:"PanTiltPositionSpace,omitempty" yaml:"PanTiltPositionSpace,omitempty"`
	ZoomPositionSpace    *Space1DDescription                        `xml:"ZoomPositionSpace,omitempty" json:"ZoomPositionSpace,omitempty" yaml:"ZoomPositionSpace,omitempty"`
	Extension            *PTZPresetTourPresetDetailOptionsExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type PTZPresetTourPresetDetailOptionsExtension []interface{}

type PTZPresetTourSpot struct {
	PresetDetail *PTZPresetTourPresetDetail  `xml:"PresetDetail,omitempty" json:"PresetDetail,omitempty" yaml:"PresetDetail,omitempty"`
	Speed        *PTZSpeed                   `xml:"Speed,omitempty" json:"Speed,omitempty" yaml:"Speed,omitempty"`
	StayTime     *Duration                   `xml:"StayTime,omitempty" json:"StayTime,omitempty" yaml:"StayTime,omitempty"`
	Extension    *PTZPresetTourSpotExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type PTZPresetTourSpotExtension []interface{}

type PTZPresetTourSpotOptions struct {
	PresetDetail *PTZPresetTourPresetDetailOptions `xml:"PresetDetail,omitempty" json:"PresetDetail,omitempty" yaml:"PresetDetail,omitempty"`
	StayTime     *DurationRange                    `xml:"StayTime,omitempty" json:"StayTime,omitempty" yaml:"StayTime,omitempty"`
}

type PTZPresetTourStartingCondition struct {
	RecurringTime     *int                                     `xml:"RecurringTime,omitempty" json:"RecurringTime,omitempty" yaml:"RecurringTime,omitempty"`
	RecurringDuration *Duration                                `xml:"RecurringDuration,omitempty" json:"RecurringDuration,omitempty" yaml:"RecurringDuration,omitempty"`
	Direction         *PTZPresetTourDirection                  `xml:"Direction,omitempty" json:"Direction,omitempty" yaml:"Direction,omitempty"`
	Extension         *PTZPresetTourStartingConditionExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
	RandomPresetOrder bool                                     `xml:"RandomPresetOrder,attr,omitempty" json:"RandomPresetOrder,attr,omitempty" yaml:"RandomPresetOrder,attr,omitempty"`
}

type PTZPresetTourStartingConditionExtension []interface{}

type PTZPresetTourStartingConditionOptions struct {
	RecurringTime     *IntRange                                       `xml:"RecurringTime,omitempty" json:"RecurringTime,omitempty" yaml:"RecurringTime,omitempty"`
	RecurringDuration *DurationRange                                  `xml:"RecurringDuration,omitempty" json:"RecurringDuration,omitempty" yaml:"RecurringDuration,omitempty"`
	Direction         []*PTZPresetTourDirection                       `xml:"Direction,omitempty" json:"Direction,omitempty" yaml:"Direction,omitempty"`
	Extension         *PTZPresetTourStartingConditionOptionsExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type PTZPresetTourStartingConditionOptionsExtension []interface{}

type PTZPresetTourStatus struct {
	State           *PTZPresetTourState           `xml:"State,omitempty" json:"State,omitempty" yaml:"State,omitempty"`
	CurrentTourSpot *PTZPresetTourSpot            `xml:"CurrentTourSpot,omitempty" json:"CurrentTourSpot,omitempty" yaml:"CurrentTourSpot,omitempty"`
	Extension       *PTZPresetTourStatusExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type PTZPresetTourStatusExtension []interface{}

type PTZPresetTourSupported struct {
	MaximumNumberOfPresetTours *int                             `xml:"MaximumNumberOfPresetTours,omitempty" json:"MaximumNumberOfPresetTours,omitempty" yaml:"MaximumNumberOfPresetTours,omitempty"`
	PTZPresetTourOperation     []*PTZPresetTourOperation        `xml:"PTZPresetTourOperation,omitempty" json:"PTZPresetTourOperation,omitempty" yaml:"PTZPresetTourOperation,omitempty"`
	Extension                  *PTZPresetTourSupportedExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type PTZPresetTourSupportedExtension []interface{}

type PTZPresetTourTypeExtension []interface{}

type PTZSpaces struct {
	AbsolutePanTiltPositionSpace    []*Space2DDescription `xml:"AbsolutePanTiltPositionSpace,omitempty" json:"AbsolutePanTiltPositionSpace,omitempty" yaml:"AbsolutePanTiltPositionSpace,omitempty"`
	AbsoluteZoomPositionSpace       []*Space1DDescription `xml:"AbsoluteZoomPositionSpace,omitempty" json:"AbsoluteZoomPositionSpace,omitempty" yaml:"AbsoluteZoomPositionSpace,omitempty"`
	RelativePanTiltTranslationSpace []*Space2DDescription `xml:"RelativePanTiltTranslationSpace,omitempty" json:"RelativePanTiltTranslationSpace,omitempty" yaml:"RelativePanTiltTranslationSpace,omitempty"`
	RelativeZoomTranslationSpace    []*Space1DDescription `xml:"RelativeZoomTranslationSpace,omitempty" json:"RelativeZoomTranslationSpace,omitempty" yaml:"RelativeZoomTranslationSpace,omitempty"`
	ContinuousPanTiltVelocitySpace  []*Space2DDescription `xml:"ContinuousPanTiltVelocitySpace,omitempty" json:"ContinuousPanTiltVelocitySpace,omitempty" yaml:"ContinuousPanTiltVelocitySpace,omitempty"`
	ContinuousZoomVelocitySpace     []*Space1DDescription `xml:"ContinuousZoomVelocitySpace,omitempty" json:"ContinuousZoomVelocitySpace,omitempty" yaml:"ContinuousZoomVelocitySpace,omitempty"`
	PanTiltSpeedSpace               []*Space1DDescription `xml:"PanTiltSpeedSpace,omitempty" json:"PanTiltSpeedSpace,omitempty" yaml:"PanTiltSpeedSpace,omitempty"`
	ZoomSpeedSpace                  []*Space1DDescription `xml:"ZoomSpeedSpace,omitempty" json:"ZoomSpeedSpace,omitempty" yaml:"ZoomSpeedSpace,omitempty"`
	Extension                       *PTZSpacesExtension   `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type PTZSpacesExtension []interface{}

type PTZSpeed struct {
	PanTilt *Vector2D `xml:"PanTilt,omitempty" json:"PanTilt,omitempty" yaml:"PanTilt,omitempty"`
	Zoom    *Vector1D `xml:"Zoom,omitempty" json:"Zoom,omitempty" yaml:"Zoom,omitempty"`
}

type PTZStatusFilterOptions struct {
	PanTiltStatusSupported   *bool                            `xml:"PanTiltStatusSupported,omitempty" json:"PanTiltStatusSupported,omitempty" yaml:"PanTiltStatusSupported,omitempty"`
	ZoomStatusSupported      *bool                            `xml:"ZoomStatusSupported,omitempty" json:"ZoomStatusSupported,omitempty" yaml:"ZoomStatusSupported,omitempty"`
	PanTiltPositionSupported *bool                            `xml:"PanTiltPositionSupported,omitempty" json:"PanTiltPositionSupported,omitempty" yaml:"PanTiltPositionSupported,omitempty"`
	ZoomPositionSupported    *bool                            `xml:"ZoomPositionSupported,omitempty" json:"ZoomPositionSupported,omitempty" yaml:"ZoomPositionSupported,omitempty"`
	Extension                *PTZStatusFilterOptionsExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type PTZStatusFilterOptionsExtension []interface{}

type PanTiltLimits struct {
	Range *Space2DDescription `xml:"Range,omitempty" json:"Range,omitempty" yaml:"Range,omitempty"`
}

// Configuration of the streaming and coding settings of a Video
//                         window.
type PaneConfiguration struct {
	PaneName                  *string                    `xml:"PaneName,omitempty" json:"PaneName,omitempty" yaml:"PaneName,omitempty"`
	AudioOutputToken          ReferenceToken             `xml:"AudioOutputToken,omitempty" json:"AudioOutputToken,omitempty" yaml:"AudioOutputToken,omitempty"`
	AudioSourceToken          ReferenceToken             `xml:"AudioSourceToken,omitempty" json:"AudioSourceToken,omitempty" yaml:"AudioSourceToken,omitempty"`
	AudioEncoderConfiguration *AudioEncoderConfiguration `xml:"AudioEncoderConfiguration,omitempty" json:"AudioEncoderConfiguration,omitempty" yaml:"AudioEncoderConfiguration,omitempty"`
	ReceiverToken             ReferenceToken             `xml:"ReceiverToken,omitempty" json:"ReceiverToken,omitempty" yaml:"ReceiverToken,omitempty"`
	Token                     ReferenceToken             `xml:"Token,omitempty" json:"Token,omitempty" yaml:"Token,omitempty"`
}

// A pane layout describes one Video window of a display. It
//                       links a pane configuration to a region
// of the screen.
type PaneLayout struct {
	Pane ReferenceToken `xml:"Pane,omitempty" json:"Pane,omitempty" yaml:"Pane,omitempty"`
	Area *Rectangle     `xml:"Area,omitempty" json:"Area,omitempty" yaml:"Area,omitempty"`
}

// Description of a pane layout describing a complete display
//                        layout.
type PaneLayoutOptions struct {
	Area      []*Rectangle         `xml:"Area,omitempty" json:"Area,omitempty" yaml:"Area,omitempty"`
	Extension *PaneOptionExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type PaneOptionExtension []interface{}

type PauseFailedFaultType struct {
	Timestamp     DateTime               `xml:"Timestamp" json:"Timestamp" yaml:"Timestamp"`
	Originator    *EndpointReferenceType `xml:"Originator,omitempty" json:"Originator,omitempty" yaml:"Originator,omitempty"`
	ErrorCode     *string                `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	Description   []*string              `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	FaultCause    interface{}            `xml:"FaultCause" json:"FaultCause" yaml:"FaultCause"`
	TypeAttrXSI   string                 `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string                 `xml:"xmlns:objtype,attr,omitempty"`

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

type PolygonOptions struct {
	RectangleOnly *bool     `xml:"RectangleOnly,omitempty" json:"RectangleOnly,omitempty" yaml:"RectangleOnly,omitempty"`
	VertexLimits  *IntRange `xml:"VertexLimits,omitempty" json:"VertexLimits,omitempty" yaml:"VertexLimits,omitempty"`
}

type Polyline struct {
	Point []*Vector `xml:"Point" json:"Point" yaml:"Point"`
}

type PolylineArray struct {
	Segment   []*Polyline             `xml:"Segment,omitempty" json:"Segment,omitempty" yaml:"Segment,omitempty"`
	Extension *PolylineArrayExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type PolylineArrayConfiguration struct {
	PolylineArray *PolylineArray `xml:"PolylineArray,omitempty" json:"PolylineArray,omitempty" yaml:"PolylineArray,omitempty"`
}

type PolylineArrayExtension []interface{}

type PrefixedIPv4Address struct {
	Address      *IPv4Address `xml:"Address,omitempty" json:"Address,omitempty" yaml:"Address,omitempty"`
	PrefixLength *int         `xml:"PrefixLength,omitempty" json:"PrefixLength,omitempty" yaml:"PrefixLength,omitempty"`
}

type PrefixedIPv6Address struct {
	Address      *IPv6Address `xml:"Address,omitempty" json:"Address,omitempty" yaml:"Address,omitempty"`
	PrefixLength *int         `xml:"PrefixLength,omitempty" json:"PrefixLength,omitempty" yaml:"PrefixLength,omitempty"`
}

type PresetTour struct {
	Name              *Name                           `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Status            *PTZPresetTourStatus            `xml:"Status,omitempty" json:"Status,omitempty" yaml:"Status,omitempty"`
	AutoStart         *bool                           `xml:"AutoStart,omitempty" json:"AutoStart,omitempty" yaml:"AutoStart,omitempty"`
	StartingCondition *PTZPresetTourStartingCondition `xml:"StartingCondition,omitempty" json:"StartingCondition,omitempty" yaml:"StartingCondition,omitempty"`
	TourSpot          []*PTZPresetTourSpot            `xml:"TourSpot,omitempty" json:"TourSpot,omitempty" yaml:"TourSpot,omitempty"`
	Extension         *PTZPresetTourExtension         `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
	Token             ReferenceToken                  `xml:"token,attr,omitempty" json:"token,attr,omitempty" yaml:"token,attr,omitempty"`
}

type ProblemActionType struct {
	Action     *AttributedURIType `xml:"Action,omitempty" json:"Action,omitempty" yaml:"Action,omitempty"`
	SoapAction *string            `xml:"SoapAction,omitempty" json:"SoapAction,omitempty" yaml:"SoapAction,omitempty"`
}

// A media profile consists of a set of media configurations.
//                        Media profiles are used by a client to
// configure properties of a media                         stream
// from an NVT. An NVT shall provide at least one media profile
// at                         boot. An NVT should provide “ready
// to use” profiles for the most common
//    media configurations that the device offers. A profile consists
// of a                         set of interconnected configuration
// entities. Configurations are provided by
//      the NVT and can be either static or created dynamically
// by the NVT. For                         example, the dynamic
// configurations can be created by the NVT depending on
//                   current available encoding resources.
type Profile struct {
	Name                        *Name                        `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	VideoSourceConfiguration    *VideoSourceConfiguration    `xml:"VideoSourceConfiguration,omitempty" json:"VideoSourceConfiguration,omitempty" yaml:"VideoSourceConfiguration,omitempty"`
	AudioSourceConfiguration    *AudioSourceConfiguration    `xml:"AudioSourceConfiguration,omitempty" json:"AudioSourceConfiguration,omitempty" yaml:"AudioSourceConfiguration,omitempty"`
	VideoEncoderConfiguration   *VideoEncoderConfiguration   `xml:"VideoEncoderConfiguration,omitempty" json:"VideoEncoderConfiguration,omitempty" yaml:"VideoEncoderConfiguration,omitempty"`
	AudioEncoderConfiguration   *AudioEncoderConfiguration   `xml:"AudioEncoderConfiguration,omitempty" json:"AudioEncoderConfiguration,omitempty" yaml:"AudioEncoderConfiguration,omitempty"`
	VideoAnalyticsConfiguration *VideoAnalyticsConfiguration `xml:"VideoAnalyticsConfiguration,omitempty" json:"VideoAnalyticsConfiguration,omitempty" yaml:"VideoAnalyticsConfiguration,omitempty"`
	PTZConfiguration            *PTZConfiguration            `xml:"PTZConfiguration,omitempty" json:"PTZConfiguration,omitempty" yaml:"PTZConfiguration,omitempty"`
	MetadataConfiguration       *MetadataConfiguration       `xml:"MetadataConfiguration,omitempty" json:"MetadataConfiguration,omitempty" yaml:"MetadataConfiguration,omitempty"`
	Extension                   *ProfileExtension            `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
	Token                       ReferenceToken               `xml:"token,attr,omitempty" json:"token,attr,omitempty" yaml:"token,attr,omitempty"`
	Fixed                       bool                         `xml:"fixed,attr,omitempty" json:"fixed,attr,omitempty" yaml:"fixed,attr,omitempty"`
}

type ProfileCapabilities struct {
	MaximumNumberOfProfiles *int `xml:"MaximumNumberOfProfiles,omitempty" json:"MaximumNumberOfProfiles,omitempty" yaml:"MaximumNumberOfProfiles,omitempty"`
}

type ProfileExtension struct {
	AudioOutputConfiguration  *AudioOutputConfiguration  `xml:"AudioOutputConfiguration,omitempty" json:"AudioOutputConfiguration,omitempty" yaml:"AudioOutputConfiguration,omitempty"`
	AudioDecoderConfiguration *AudioDecoderConfiguration `xml:"AudioDecoderConfiguration,omitempty" json:"AudioDecoderConfiguration,omitempty" yaml:"AudioDecoderConfiguration,omitempty"`
	Extension                 *ProfileExtension2         `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type ProfileExtension2 []interface{}

type ProfileStatus struct {
	ActiveConnections []*ActiveConnection     `xml:"ActiveConnections,omitempty" json:"ActiveConnections,omitempty" yaml:"ActiveConnections,omitempty"`
	Extension         *ProfileStatusExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type ProfileStatusExtension []interface{}

type QueryExpressionType []interface{}

type RealTimeStreamingCapabilities struct {
	RTPMulticast *bool                                   `xml:"RTPMulticast,omitempty" json:"RTPMulticast,omitempty" yaml:"RTPMulticast,omitempty"`
	RTP_TCP      *bool                                   `xml:"RTP_TCP,omitempty" json:"RTP_TCP,omitempty" yaml:"RTP_TCP,omitempty"`
	RTP_RTSP_TCP *bool                                   `xml:"RTP_RTSP_TCP,omitempty" json:"RTP_RTSP_TCP,omitempty" yaml:"RTP_RTSP_TCP,omitempty"`
	Extension    *RealTimeStreamingCapabilitiesExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type RealTimeStreamingCapabilitiesExtension []interface{}

// Description of a receiver, including its token and
//                configuration.
type Receiver struct {
	Token         ReferenceToken         `xml:"Token,omitempty" json:"Token,omitempty" yaml:"Token,omitempty"`
	Configuration *ReceiverConfiguration `xml:"Configuration,omitempty" json:"Configuration,omitempty" yaml:"Configuration,omitempty"`
}

type ReceiverCapabilities struct {
	XAddr                *string `xml:"XAddr,omitempty" json:"XAddr,omitempty" yaml:"XAddr,omitempty"`
	RTP_Multicast        *bool   `xml:"RTP_Multicast,omitempty" json:"RTP_Multicast,omitempty" yaml:"RTP_Multicast,omitempty"`
	RTP_TCP              *bool   `xml:"RTP_TCP,omitempty" json:"RTP_TCP,omitempty" yaml:"RTP_TCP,omitempty"`
	RTP_RTSP_TCP         *bool   `xml:"RTP_RTSP_TCP,omitempty" json:"RTP_RTSP_TCP,omitempty" yaml:"RTP_RTSP_TCP,omitempty"`
	SupportedReceivers   *int    `xml:"SupportedReceivers,omitempty" json:"SupportedReceivers,omitempty" yaml:"SupportedReceivers,omitempty"`
	MaximumRTSPURILength *int    `xml:"MaximumRTSPURILength,omitempty" json:"MaximumRTSPURILength,omitempty" yaml:"MaximumRTSPURILength,omitempty"`
}

// Describes the configuration of a receiver.
type ReceiverConfiguration struct {
	Mode        *ReceiverMode `xml:"Mode,omitempty" json:"Mode,omitempty" yaml:"Mode,omitempty"`
	MediaUri    *string       `xml:"MediaUri,omitempty" json:"MediaUri,omitempty" yaml:"MediaUri,omitempty"`
	StreamSetup *StreamSetup  `xml:"StreamSetup,omitempty" json:"StreamSetup,omitempty" yaml:"StreamSetup,omitempty"`
}

// Contains information about a receiver's current state.
type ReceiverStateInformation struct {
	State       *ReceiverState `xml:"State,omitempty" json:"State,omitempty" yaml:"State,omitempty"`
	AutoCreated *bool          `xml:"AutoCreated,omitempty" json:"AutoCreated,omitempty" yaml:"AutoCreated,omitempty"`
}

type RecordingCapabilities struct {
	XAddr              *string `xml:"XAddr,omitempty" json:"XAddr,omitempty" yaml:"XAddr,omitempty"`
	ReceiverSource     *bool   `xml:"ReceiverSource,omitempty" json:"ReceiverSource,omitempty" yaml:"ReceiverSource,omitempty"`
	MediaProfileSource *bool   `xml:"MediaProfileSource,omitempty" json:"MediaProfileSource,omitempty" yaml:"MediaProfileSource,omitempty"`
	DynamicRecordings  *bool   `xml:"DynamicRecordings,omitempty" json:"DynamicRecordings,omitempty" yaml:"DynamicRecordings,omitempty"`
	DynamicTracks      *bool   `xml:"DynamicTracks,omitempty" json:"DynamicTracks,omitempty" yaml:"DynamicTracks,omitempty"`
	MaxStringLength    *int    `xml:"MaxStringLength,omitempty" json:"MaxStringLength,omitempty" yaml:"MaxStringLength,omitempty"`
}

type RecordingConfiguration struct {
	Source               *RecordingSourceInformation `xml:"Source,omitempty" json:"Source,omitempty" yaml:"Source,omitempty"`
	Content              *Description                `xml:"Content,omitempty" json:"Content,omitempty" yaml:"Content,omitempty"`
	MaximumRetentionTime *Duration                   `xml:"MaximumRetentionTime,omitempty" json:"MaximumRetentionTime,omitempty" yaml:"MaximumRetentionTime,omitempty"`
}

type RecordingEventFilter struct {
	Filter []*interface{} `xml:"Filter>Filter,omitempty" json:"Filter>Filter,omitempty" yaml:"Filter>Filter,omitempty"`
	Before *Duration      `xml:"Before,omitempty" json:"Before,omitempty" yaml:"Before,omitempty"`
	After  *Duration      `xml:"After,omitempty" json:"After,omitempty" yaml:"After,omitempty"`
}

type RecordingInformation struct {
	RecordingToken    *RecordingReference         `xml:"RecordingToken,omitempty" json:"RecordingToken,omitempty" yaml:"RecordingToken,omitempty"`
	Source            *RecordingSourceInformation `xml:"Source,omitempty" json:"Source,omitempty" yaml:"Source,omitempty"`
	EarliestRecording *DateTime                   `xml:"EarliestRecording,omitempty" json:"EarliestRecording,omitempty" yaml:"EarliestRecording,omitempty"`
	LatestRecording   *DateTime                   `xml:"LatestRecording,omitempty" json:"LatestRecording,omitempty" yaml:"LatestRecording,omitempty"`
	Content           *Description                `xml:"Content,omitempty" json:"Content,omitempty" yaml:"Content,omitempty"`
	Track             []*TrackInformation         `xml:"Track,omitempty" json:"Track,omitempty" yaml:"Track,omitempty"`
	RecordingStatus   *RecordingStatus            `xml:"RecordingStatus,omitempty" json:"RecordingStatus,omitempty" yaml:"RecordingStatus,omitempty"`
}

type RecordingJobConfiguration struct {
	RecordingToken *RecordingReference                 `xml:"RecordingToken,omitempty" json:"RecordingToken,omitempty" yaml:"RecordingToken,omitempty"`
	Mode           *RecordingJobMode                   `xml:"Mode,omitempty" json:"Mode,omitempty" yaml:"Mode,omitempty"`
	Priority       *int                                `xml:"Priority,omitempty" json:"Priority,omitempty" yaml:"Priority,omitempty"`
	Source         []*RecordingJobSource               `xml:"Source,omitempty" json:"Source,omitempty" yaml:"Source,omitempty"`
	Extension      *RecordingJobConfigurationExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
	EventFilter    *RecordingEventFilter               `xml:"EventFilter,omitempty" json:"EventFilter,omitempty" yaml:"EventFilter,omitempty"`
	ScheduleToken  string                              `xml:"ScheduleToken,attr,omitempty" json:"ScheduleToken,attr,omitempty" yaml:"ScheduleToken,attr,omitempty"`
}

type RecordingJobConfigurationExtension []interface{}

type RecordingJobSource struct {
	SourceToken        *SourceReference             `xml:"SourceToken,omitempty" json:"SourceToken,omitempty" yaml:"SourceToken,omitempty"`
	AutoCreateReceiver *bool                        `xml:"AutoCreateReceiver,omitempty" json:"AutoCreateReceiver,omitempty" yaml:"AutoCreateReceiver,omitempty"`
	Tracks             []*RecordingJobTrack         `xml:"Tracks,omitempty" json:"Tracks,omitempty" yaml:"Tracks,omitempty"`
	Extension          *RecordingJobSourceExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type RecordingJobSourceExtension []interface{}

type RecordingJobStateInformation struct {
	RecordingToken *RecordingReference                    `xml:"RecordingToken,omitempty" json:"RecordingToken,omitempty" yaml:"RecordingToken,omitempty"`
	State          *RecordingJobState                     `xml:"State,omitempty" json:"State,omitempty" yaml:"State,omitempty"`
	Sources        []*RecordingJobStateSource             `xml:"Sources,omitempty" json:"Sources,omitempty" yaml:"Sources,omitempty"`
	Extension      *RecordingJobStateInformationExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type RecordingJobStateInformationExtension []interface{}

type RecordingJobStateSource struct {
	SourceToken *SourceReference         `xml:"SourceToken,omitempty" json:"SourceToken,omitempty" yaml:"SourceToken,omitempty"`
	State       *RecordingJobState       `xml:"State,omitempty" json:"State,omitempty" yaml:"State,omitempty"`
	Tracks      *RecordingJobStateTracks `xml:"Tracks,omitempty" json:"Tracks,omitempty" yaml:"Tracks,omitempty"`
}

type RecordingJobStateTrack struct {
	SourceTag   *string            `xml:"SourceTag,omitempty" json:"SourceTag,omitempty" yaml:"SourceTag,omitempty"`
	Destination *TrackReference    `xml:"Destination,omitempty" json:"Destination,omitempty" yaml:"Destination,omitempty"`
	Error       *string            `xml:"Error,omitempty" json:"Error,omitempty" yaml:"Error,omitempty"`
	State       *RecordingJobState `xml:"State,omitempty" json:"State,omitempty" yaml:"State,omitempty"`
}

type RecordingJobStateTracks struct {
	Track []*RecordingJobStateTrack `xml:"Track,omitempty" json:"Track,omitempty" yaml:"Track,omitempty"`
}

type RecordingJobTrack struct {
	SourceTag   *string         `xml:"SourceTag,omitempty" json:"SourceTag,omitempty" yaml:"SourceTag,omitempty"`
	Destination *TrackReference `xml:"Destination,omitempty" json:"Destination,omitempty" yaml:"Destination,omitempty"`
}

// A set of informative desciptions of a data source. The Search
//                         searvice allows a client to filter on
// recordings based on information in                         this
// structure.
type RecordingSourceInformation struct {
	SourceId    *string      `xml:"SourceId,omitempty" json:"SourceId,omitempty" yaml:"SourceId,omitempty"`
	Name        *Name        `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Location    *Description `xml:"Location,omitempty" json:"Location,omitempty" yaml:"Location,omitempty"`
	Description *Description `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	Address     *string      `xml:"Address,omitempty" json:"Address,omitempty" yaml:"Address,omitempty"`
}

type RecordingSummary struct {
	DataFrom         *DateTime `xml:"DataFrom,omitempty" json:"DataFrom,omitempty" yaml:"DataFrom,omitempty"`
	DataUntil        *DateTime `xml:"DataUntil,omitempty" json:"DataUntil,omitempty" yaml:"DataUntil,omitempty"`
	NumberRecordings *int      `xml:"NumberRecordings,omitempty" json:"NumberRecordings,omitempty" yaml:"NumberRecordings,omitempty"`
}

type ReferenceParametersType []interface{}

type RelatesToType struct {
	Content          *string                  `xml:"Content,omitempty" json:"Content,omitempty" yaml:"Content,omitempty"`
	RelationshipType RelationshipTypeOpenEnum `xml:"RelationshipType,attr,omitempty" json:"RelationshipType,attr,omitempty" yaml:"RelationshipType,attr,omitempty"`
}

type RelativeFocus struct {
	Distance *float64 `xml:"Distance,omitempty" json:"Distance,omitempty" yaml:"Distance,omitempty"`
	Speed    *float64 `xml:"Speed,omitempty" json:"Speed,omitempty" yaml:"Speed,omitempty"`
}

type RelativeFocusOptions struct {
	Distance *FloatRange `xml:"Distance,omitempty" json:"Distance,omitempty" yaml:"Distance,omitempty"`
	Speed    *FloatRange `xml:"Speed,omitempty" json:"Speed,omitempty" yaml:"Speed,omitempty"`
}

type RelativeFocusOptions20 struct {
	Distance *FloatRange `xml:"Distance,omitempty" json:"Distance,omitempty" yaml:"Distance,omitempty"`
	Speed    *FloatRange `xml:"Speed,omitempty" json:"Speed,omitempty" yaml:"Speed,omitempty"`
}

type RelayOutput struct {
	Token         ReferenceToken       `xml:"token,attr,omitempty" json:"token,attr,omitempty" yaml:"token,attr,omitempty"`
	Properties    *RelayOutputSettings `xml:"Properties,omitempty" json:"Properties,omitempty" yaml:"Properties,omitempty"`
	TypeAttrXSI   string               `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string               `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *RelayOutput) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:RelayOutput"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://www.onvif.org/ver10/schema"
	}
}

type RelayOutputSettings struct {
	Mode      *RelayMode      `xml:"Mode,omitempty" json:"Mode,omitempty" yaml:"Mode,omitempty"`
	DelayTime *Duration       `xml:"DelayTime,omitempty" json:"DelayTime,omitempty" yaml:"DelayTime,omitempty"`
	IdleState *RelayIdleState `xml:"IdleState,omitempty" json:"IdleState,omitempty" yaml:"IdleState,omitempty"`
}

type RemoteUser struct {
	Username           *string `xml:"Username,omitempty" json:"Username,omitempty" yaml:"Username,omitempty"`
	Password           *string `xml:"Password,omitempty" json:"Password,omitempty" yaml:"Password,omitempty"`
	UseDerivedPassword *bool   `xml:"UseDerivedPassword,omitempty" json:"UseDerivedPassword,omitempty" yaml:"UseDerivedPassword,omitempty"`
}

type Renew struct {
	TerminationTime *AbsoluteOrRelativeTimeType `xml:"TerminationTime,omitempty" json:"TerminationTime,omitempty" yaml:"TerminationTime,omitempty"`
}

type RenewResponse struct {
	TerminationTime *DateTime `xml:"TerminationTime,omitempty" json:"TerminationTime,omitempty" yaml:"TerminationTime,omitempty"`
	CurrentTime     *DateTime `xml:"CurrentTime,omitempty" json:"CurrentTime,omitempty" yaml:"CurrentTime,omitempty"`
}

type ReplayCapabilities struct {
	XAddr *string `xml:"XAddr,omitempty" json:"XAddr,omitempty" yaml:"XAddr,omitempty"`
}

// Configuration parameters for the replay service.
type ReplayConfiguration struct {
	SessionTimeout *Duration `xml:"SessionTimeout,omitempty" json:"SessionTimeout,omitempty" yaml:"SessionTimeout,omitempty"`
}

type ResumeFailedFaultType struct {
	Timestamp     DateTime               `xml:"Timestamp" json:"Timestamp" yaml:"Timestamp"`
	Originator    *EndpointReferenceType `xml:"Originator,omitempty" json:"Originator,omitempty" yaml:"Originator,omitempty"`
	ErrorCode     *string                `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	Description   []*string              `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	FaultCause    interface{}            `xml:"FaultCause" json:"FaultCause" yaml:"FaultCause"`
	TypeAttrXSI   string                 `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string                 `xml:"xmlns:objtype,attr,omitempty"`

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

type Reverse struct {
	Mode *ReverseMode `xml:"Mode,omitempty" json:"Mode,omitempty" yaml:"Mode,omitempty"`
}

type ReverseOptions struct {
	Mode      []*ReverseMode           `xml:"Mode,omitempty" json:"Mode,omitempty" yaml:"Mode,omitempty"`
	Extension *ReverseOptionsExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type ReverseOptionsExtension []interface{}

type Rotate struct {
	Mode      *RotateMode      `xml:"Mode,omitempty" json:"Mode,omitempty" yaml:"Mode,omitempty"`
	Degree    *int             `xml:"Degree,omitempty" json:"Degree,omitempty" yaml:"Degree,omitempty"`
	Extension *RotateExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type RotateExtension []interface{}

type RotateOptions struct {
	Mode       []*RotateMode           `xml:"Mode,omitempty" json:"Mode,omitempty" yaml:"Mode,omitempty"`
	DegreeList *IntItems               `xml:"DegreeList,omitempty" json:"DegreeList,omitempty" yaml:"DegreeList,omitempty"`
	Extension  *RotateOptionsExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
	Reboot     bool                    `xml:"Reboot,attr,omitempty" json:"Reboot,attr,omitempty" yaml:"Reboot,attr,omitempty"`
}

type RotateOptionsExtension []interface{}

type RuleEngineConfiguration struct {
	Rule      []*Config                         `xml:"Rule,omitempty" json:"Rule,omitempty" yaml:"Rule,omitempty"`
	Extension *RuleEngineConfigurationExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type RuleEngineConfigurationExtension []interface{}

type SceneOrientation struct {
	Mode        *SceneOrientationMode `xml:"Mode,omitempty" json:"Mode,omitempty" yaml:"Mode,omitempty"`
	Orientation *string               `xml:"Orientation,omitempty" json:"Orientation,omitempty" yaml:"Orientation,omitempty"`
}

type Scope struct {
	ScopeDef  *ScopeDefinition `xml:"ScopeDef,omitempty" json:"ScopeDef,omitempty" yaml:"ScopeDef,omitempty"`
	ScopeItem *string          `xml:"ScopeItem,omitempty" json:"ScopeItem,omitempty" yaml:"ScopeItem,omitempty"`
}

type SearchCapabilities struct {
	XAddr          *string `xml:"XAddr,omitempty" json:"XAddr,omitempty" yaml:"XAddr,omitempty"`
	MetadataSearch *bool   `xml:"MetadataSearch,omitempty" json:"MetadataSearch,omitempty" yaml:"MetadataSearch,omitempty"`
}

// A structure for defining a limited scope when searching in
//                        recorded data.
type SearchScope struct {
	IncludedSources            []*SourceReference    `xml:"IncludedSources,omitempty" json:"IncludedSources,omitempty" yaml:"IncludedSources,omitempty"`
	IncludedRecordings         []*RecordingReference `xml:"IncludedRecordings,omitempty" json:"IncludedRecordings,omitempty" yaml:"IncludedRecordings,omitempty"`
	RecordingInformationFilter *XPathExpression      `xml:"RecordingInformationFilter,omitempty" json:"RecordingInformationFilter,omitempty" yaml:"RecordingInformationFilter,omitempty"`
	Extension                  *SearchScopeExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type SearchScopeExtension []interface{}

type SecurityCapabilities struct {
	TLS1                 *bool                          `xml:"TLS1.1,omitempty" json:"TLS1.1,omitempty" yaml:"TLS1.1,omitempty"`
	TLS2                 *bool                          `xml:"TLS1.2,omitempty" json:"TLS1.2,omitempty" yaml:"TLS1.2,omitempty"`
	OnboardKeyGeneration *bool                          `xml:"OnboardKeyGeneration,omitempty" json:"OnboardKeyGeneration,omitempty" yaml:"OnboardKeyGeneration,omitempty"`
	AccessPolicyConfig   *bool                          `xml:"AccessPolicyConfig,omitempty" json:"AccessPolicyConfig,omitempty" yaml:"AccessPolicyConfig,omitempty"`
	X509Token            *bool                          `xml:"X.509Token,omitempty" json:"X.509Token,omitempty" yaml:"X.509Token,omitempty"`
	SAMLToken            *bool                          `xml:"SAMLToken,omitempty" json:"SAMLToken,omitempty" yaml:"SAMLToken,omitempty"`
	KerberosToken        *bool                          `xml:"KerberosToken,omitempty" json:"KerberosToken,omitempty" yaml:"KerberosToken,omitempty"`
	RELToken             *bool                          `xml:"RELToken,omitempty" json:"RELToken,omitempty" yaml:"RELToken,omitempty"`
	Extension            *SecurityCapabilitiesExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type SecurityCapabilitiesExtension struct {
	TLS0      *bool                           `xml:"TLS1.0,omitempty" json:"TLS1.0,omitempty" yaml:"TLS1.0,omitempty"`
	Extension *SecurityCapabilitiesExtension2 `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type SecurityCapabilitiesExtension2 struct {
	Dot1X              *bool  `xml:"Dot1X,omitempty" json:"Dot1X,omitempty" yaml:"Dot1X,omitempty"`
	SupportedEAPMethod []*int `xml:"SupportedEAPMethod,omitempty" json:"SupportedEAPMethod,omitempty" yaml:"SupportedEAPMethod,omitempty"`
	RemoteUserHandling *bool  `xml:"RemoteUserHandling,omitempty" json:"RemoteUserHandling,omitempty" yaml:"RemoteUserHandling,omitempty"`
}

type SourceIdentification struct {
	Name      *string                        `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Token     []ReferenceToken               `xml:"Token,omitempty" json:"Token,omitempty" yaml:"Token,omitempty"`
	Extension *SourceIdentificationExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type SourceIdentificationExtension []interface{}

type SourceReference struct {
	Token ReferenceToken `xml:"Token,omitempty" json:"Token,omitempty" yaml:"Token,omitempty"`
	Type  string         `xml:"Type,attr,omitempty" json:"Type,attr,omitempty" yaml:"Type,attr,omitempty"`
}

type Space1DDescription struct {
	URI    *string     `xml:"URI,omitempty" json:"URI,omitempty" yaml:"URI,omitempty"`
	XRange *FloatRange `xml:"XRange,omitempty" json:"XRange,omitempty" yaml:"XRange,omitempty"`
}

type Space2DDescription struct {
	URI    *string     `xml:"URI,omitempty" json:"URI,omitempty" yaml:"URI,omitempty"`
	XRange *FloatRange `xml:"XRange,omitempty" json:"XRange,omitempty" yaml:"XRange,omitempty"`
	YRange *FloatRange `xml:"YRange,omitempty" json:"YRange,omitempty" yaml:"YRange,omitempty"`
}

type StorageReferencePath struct {
	StorageToken ReferenceToken                 `xml:"StorageToken,omitempty" json:"StorageToken,omitempty" yaml:"StorageToken,omitempty"`
	RelativePath *string                        `xml:"RelativePath,omitempty" json:"RelativePath,omitempty" yaml:"RelativePath,omitempty"`
	Extension    *StorageReferencePathExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type StorageReferencePathExtension []interface{}

type StreamSetup struct {
	Stream    *StreamType `xml:"Stream,omitempty" json:"Stream,omitempty" yaml:"Stream,omitempty"`
	Transport *Transport  `xml:"Transport,omitempty" json:"Transport,omitempty" yaml:"Transport,omitempty"`
}

type StringItems struct {
	Item []*string `xml:"Item,omitempty" json:"Item,omitempty" yaml:"Item,omitempty"`
}

type Subscribe struct {
	ConsumerReference      *EndpointReferenceType      `xml:"ConsumerReference" json:"ConsumerReference" yaml:"ConsumerReference"`
	Filter                 *FilterType                 `xml:"Filter,omitempty" json:"Filter,omitempty" yaml:"Filter,omitempty"`
	InitialTerminationTime *AbsoluteOrRelativeTimeType `xml:"InitialTerminationTime,omitempty" json:"InitialTerminationTime,omitempty" yaml:"InitialTerminationTime,omitempty"`
	SubscriptionPolicy     []*interface{}              `xml:"SubscriptionPolicy>SubscriptionPolicy,omitempty" json:"SubscriptionPolicy>SubscriptionPolicy,omitempty" yaml:"SubscriptionPolicy>SubscriptionPolicy,omitempty"`
}

type SubscribeCreationFailedFaultType struct {
	Timestamp     DateTime               `xml:"Timestamp" json:"Timestamp" yaml:"Timestamp"`
	Originator    *EndpointReferenceType `xml:"Originator,omitempty" json:"Originator,omitempty" yaml:"Originator,omitempty"`
	ErrorCode     *string                `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	Description   []*string              `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	FaultCause    interface{}            `xml:"FaultCause" json:"FaultCause" yaml:"FaultCause"`
	TypeAttrXSI   string                 `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string                 `xml:"xmlns:objtype,attr,omitempty"`

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
	SubscriptionReference *EndpointReferenceType `xml:"SubscriptionReference" json:"SubscriptionReference" yaml:"SubscriptionReference"`
	CurrentTime           *DateTime              `xml:"CurrentTime,omitempty" json:"CurrentTime,omitempty" yaml:"CurrentTime,omitempty"`
	TerminationTime       *DateTime              `xml:"TerminationTime,omitempty" json:"TerminationTime,omitempty" yaml:"TerminationTime,omitempty"`
}

type SubscriptionManagerRP struct {
	ConsumerReference  *EndpointReferenceType  `xml:"ConsumerReference,omitempty" json:"ConsumerReference,omitempty" yaml:"ConsumerReference,omitempty"`
	Filter             *FilterType             `xml:"Filter,omitempty" json:"Filter,omitempty" yaml:"Filter,omitempty"`
	SubscriptionPolicy *SubscriptionPolicyType `xml:"SubscriptionPolicy,omitempty" json:"SubscriptionPolicy,omitempty" yaml:"SubscriptionPolicy,omitempty"`
	CreationTime       *DateTime               `xml:"CreationTime,omitempty" json:"CreationTime,omitempty" yaml:"CreationTime,omitempty"`
}

type SubscriptionPolicyType []interface{}

type SupportInformation struct {
	Binary *AttachmentData `xml:"Binary,omitempty" json:"Binary,omitempty" yaml:"Binary,omitempty"`
	String *string         `xml:"String,omitempty" json:"String,omitempty" yaml:"String,omitempty"`
}

type SupportedAnalyticsModules struct {
	AnalyticsModuleContentSchemaLocation []*string                           `xml:"AnalyticsModuleContentSchemaLocation,omitempty" json:"AnalyticsModuleContentSchemaLocation,omitempty" yaml:"AnalyticsModuleContentSchemaLocation,omitempty"`
	AnalyticsModuleDescription           []*ConfigDescription                `xml:"AnalyticsModuleDescription,omitempty" json:"AnalyticsModuleDescription,omitempty" yaml:"AnalyticsModuleDescription,omitempty"`
	Extension                            *SupportedAnalyticsModulesExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
	Limit                                int                                 `xml:"Limit,attr,omitempty" json:"Limit,attr,omitempty" yaml:"Limit,attr,omitempty"`
}

type SupportedAnalyticsModulesExtension []interface{}

type SupportedEnvType struct {
	Qname string `xml:"qname,attr,omitempty" json:"qname,attr,omitempty" yaml:"qname,attr,omitempty"`
}

type SupportedRules struct {
	RuleContentSchemaLocation []*string                `xml:"RuleContentSchemaLocation,omitempty" json:"RuleContentSchemaLocation,omitempty" yaml:"RuleContentSchemaLocation,omitempty"`
	RuleDescription           []*ConfigDescription     `xml:"RuleDescription,omitempty" json:"RuleDescription,omitempty" yaml:"RuleDescription,omitempty"`
	Extension                 *SupportedRulesExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
	Limit                     int                      `xml:"Limit,attr,omitempty" json:"Limit,attr,omitempty" yaml:"Limit,attr,omitempty"`
}

type SupportedRulesExtension []interface{}

type SystemCapabilities struct {
	DiscoveryResolve  *bool                        `xml:"DiscoveryResolve,omitempty" json:"DiscoveryResolve,omitempty" yaml:"DiscoveryResolve,omitempty"`
	DiscoveryBye      *bool                        `xml:"DiscoveryBye,omitempty" json:"DiscoveryBye,omitempty" yaml:"DiscoveryBye,omitempty"`
	RemoteDiscovery   *bool                        `xml:"RemoteDiscovery,omitempty" json:"RemoteDiscovery,omitempty" yaml:"RemoteDiscovery,omitempty"`
	SystemBackup      *bool                        `xml:"SystemBackup,omitempty" json:"SystemBackup,omitempty" yaml:"SystemBackup,omitempty"`
	SystemLogging     *bool                        `xml:"SystemLogging,omitempty" json:"SystemLogging,omitempty" yaml:"SystemLogging,omitempty"`
	FirmwareUpgrade   *bool                        `xml:"FirmwareUpgrade,omitempty" json:"FirmwareUpgrade,omitempty" yaml:"FirmwareUpgrade,omitempty"`
	SupportedVersions []*OnvifVersion              `xml:"SupportedVersions,omitempty" json:"SupportedVersions,omitempty" yaml:"SupportedVersions,omitempty"`
	Extension         *SystemCapabilitiesExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type SystemCapabilitiesExtension struct {
	HttpFirmwareUpgrade    *bool                         `xml:"HttpFirmwareUpgrade,omitempty" json:"HttpFirmwareUpgrade,omitempty" yaml:"HttpFirmwareUpgrade,omitempty"`
	HttpSystemBackup       *bool                         `xml:"HttpSystemBackup,omitempty" json:"HttpSystemBackup,omitempty" yaml:"HttpSystemBackup,omitempty"`
	HttpSystemLogging      *bool                         `xml:"HttpSystemLogging,omitempty" json:"HttpSystemLogging,omitempty" yaml:"HttpSystemLogging,omitempty"`
	HttpSupportInformation *bool                         `xml:"HttpSupportInformation,omitempty" json:"HttpSupportInformation,omitempty" yaml:"HttpSupportInformation,omitempty"`
	Extension              *SystemCapabilitiesExtension2 `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type SystemCapabilitiesExtension2 []interface{}

// General date time inforamtion returned by the
//           GetSystemDateTime method.
type SystemDateTime struct {
	DateTimeType    *SetDateTimeType         `xml:"DateTimeType,omitempty" json:"DateTimeType,omitempty" yaml:"DateTimeType,omitempty"`
	DaylightSavings *bool                    `xml:"DaylightSavings,omitempty" json:"DaylightSavings,omitempty" yaml:"DaylightSavings,omitempty"`
	TimeZone        *TimeZone                `xml:"TimeZone,omitempty" json:"TimeZone,omitempty" yaml:"TimeZone,omitempty"`
	UTCDateTime     *DateTimeStruct          `xml:"UTCDateTime,omitempty" json:"UTCDateTime,omitempty" yaml:"UTCDateTime,omitempty"`
	LocalDateTime   *DateTimeStruct          `xml:"LocalDateTime,omitempty" json:"LocalDateTime,omitempty" yaml:"LocalDateTime,omitempty"`
	Extension       *SystemDateTimeExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type SystemDateTimeExtension []interface{}

type SystemLog struct {
	Binary *AttachmentData `xml:"Binary,omitempty" json:"Binary,omitempty" yaml:"Binary,omitempty"`
	String *string         `xml:"String,omitempty" json:"String,omitempty" yaml:"String,omitempty"`
}

type SystemLogUri struct {
	Type *SystemLogType `xml:"Type,omitempty" json:"Type,omitempty" yaml:"Type,omitempty"`
	Uri  *string        `xml:"Uri,omitempty" json:"Uri,omitempty" yaml:"Uri,omitempty"`
}

type SystemLogUriList struct {
	SystemLog []*SystemLogUri `xml:"SystemLog,omitempty" json:"SystemLog,omitempty" yaml:"SystemLog,omitempty"`
}

type TLSConfiguration struct {
	CertificateID *string `xml:"CertificateID,omitempty" json:"CertificateID,omitempty" yaml:"CertificateID,omitempty"`
}

type TimeStruct struct {
	Hour   *int `xml:"Hour,omitempty" json:"Hour,omitempty" yaml:"Hour,omitempty"`
	Minute *int `xml:"Minute,omitempty" json:"Minute,omitempty" yaml:"Minute,omitempty"`
	Second *int `xml:"Second,omitempty" json:"Second,omitempty" yaml:"Second,omitempty"`
}

// The TZ format is specified by POSIX, please refer to POSIX
//                        1003.1 section 8.3 Example: Europe, Paris
//                         TZ=CET-1CEST,M3.5.0/2,M10.5.0/3 CET
// = designation for standard time                         when
// daylight saving is not in force -1 = offset in hours = negative
//                         so 1 hour east of Greenwich meridian
// CEST = designation when daylight                         saving
// is in force ("Central European Summer Time") , = no offset
//                        number between code and comma, so default
// to one hour ahead for daylight                         saving
// M3.5.0 = when daylight saving starts = the last Sunday in March
//                         (the "5th" week means the last in the
// month) /2, = the local time when                         the
// switch occurs = 2 a.m. in this case M10.5.0 = when daylight
// saving                         ends = the last Sunday in October.
// /3, = the local time when the                         switch
// occurs = 3 a.m. in this case
type TimeZone struct {
	TZ *string `xml:"TZ,omitempty" json:"TZ,omitempty" yaml:"TZ,omitempty"`
}

type ToneCompensation struct {
	Mode      *string                    `xml:"Mode,omitempty" json:"Mode,omitempty" yaml:"Mode,omitempty"`
	Level     *float64                   `xml:"Level,omitempty" json:"Level,omitempty" yaml:"Level,omitempty"`
	Extension *ToneCompensationExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type ToneCompensationExtension []interface{}

type ToneCompensationOptions struct {
	Mode  []*string `xml:"Mode,omitempty" json:"Mode,omitempty" yaml:"Mode,omitempty"`
	Level *bool     `xml:"Level,omitempty" json:"Level,omitempty" yaml:"Level,omitempty"`
}

type TopicExpressionDialectUnknownFaultType struct {
	Timestamp     DateTime               `xml:"Timestamp" json:"Timestamp" yaml:"Timestamp"`
	Originator    *EndpointReferenceType `xml:"Originator,omitempty" json:"Originator,omitempty" yaml:"Originator,omitempty"`
	ErrorCode     *string                `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	Description   []*string              `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	FaultCause    interface{}            `xml:"FaultCause" json:"FaultCause" yaml:"FaultCause"`
	TypeAttrXSI   string                 `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string                 `xml:"xmlns:objtype,attr,omitempty"`

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
	Documentation   *Documentation `xml:"documentation,omitempty" json:"documentation,omitempty" yaml:"documentation,omitempty"`
	Name            *NCName        `xml:"name,attr,omitempty" json:"name,attr,omitempty" yaml:"name,attr,omitempty"`
	TargetNamespace string         `xml:"targetNamespace,attr,omitempty" json:"targetNamespace,attr,omitempty" yaml:"targetNamespace,attr,omitempty"`
	Final           bool           `xml:"final,attr,omitempty" json:"final,attr,omitempty" yaml:"final,attr,omitempty"`
	Topic           []*string      `xml:"Topic,omitempty" json:"Topic,omitempty" yaml:"Topic,omitempty"`
	TypeAttrXSI     string         `xml:"xsi:type,attr,omitempty"`
	TypeNamespace   string         `xml:"xmlns:objtype,attr,omitempty"`

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
	Timestamp     DateTime               `xml:"Timestamp" json:"Timestamp" yaml:"Timestamp"`
	Originator    *EndpointReferenceType `xml:"Originator,omitempty" json:"Originator,omitempty" yaml:"Originator,omitempty"`
	ErrorCode     *string                `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	Description   []*string              `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	FaultCause    interface{}            `xml:"FaultCause" json:"FaultCause" yaml:"FaultCause"`
	TypeAttrXSI   string                 `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string                 `xml:"xmlns:objtype,attr,omitempty"`

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
	Documentation *Documentation `xml:"documentation,omitempty" json:"documentation,omitempty" yaml:"documentation,omitempty"`
	TypeAttrXSI   string         `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string         `xml:"xmlns:objtype,attr,omitempty"`

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
	Documentation  *Documentation       `xml:"documentation,omitempty" json:"documentation,omitempty" yaml:"documentation,omitempty"`
	Name           *NCName              `xml:"name,attr,omitempty" json:"name,attr,omitempty" yaml:"name,attr,omitempty"`
	MessageTypes   string               `xml:"messageTypes,attr,omitempty" json:"messageTypes,attr,omitempty" yaml:"messageTypes,attr,omitempty"`
	Final          bool                 `xml:"final,attr,omitempty" json:"final,attr,omitempty" yaml:"final,attr,omitempty"`
	MessagePattern *QueryExpressionType `xml:"MessagePattern,omitempty" json:"MessagePattern,omitempty" yaml:"MessagePattern,omitempty"`
	Topic          []*TopicType         `xml:"Topic,omitempty" json:"Topic,omitempty" yaml:"Topic,omitempty"`
	TypeAttrXSI    string               `xml:"xsi:type,attr,omitempty"`
	TypeNamespace  string               `xml:"xmlns:objtype,attr,omitempty"`

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

type TrackAttributes struct {
	TrackInformation   *TrackInformation         `xml:"TrackInformation,omitempty" json:"TrackInformation,omitempty" yaml:"TrackInformation,omitempty"`
	VideoAttributes    *VideoAttributes          `xml:"VideoAttributes,omitempty" json:"VideoAttributes,omitempty" yaml:"VideoAttributes,omitempty"`
	AudioAttributes    *AudioAttributes          `xml:"AudioAttributes,omitempty" json:"AudioAttributes,omitempty" yaml:"AudioAttributes,omitempty"`
	MetadataAttributes *MetadataAttributes       `xml:"MetadataAttributes,omitempty" json:"MetadataAttributes,omitempty" yaml:"MetadataAttributes,omitempty"`
	Extension          *TrackAttributesExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type TrackAttributesExtension []interface{}

type TrackConfiguration struct {
	TrackType   *TrackType   `xml:"TrackType,omitempty" json:"TrackType,omitempty" yaml:"TrackType,omitempty"`
	Description *Description `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
}

type TrackInformation struct {
	TrackToken  *TrackReference `xml:"TrackToken,omitempty" json:"TrackToken,omitempty" yaml:"TrackToken,omitempty"`
	TrackType   *TrackType      `xml:"TrackType,omitempty" json:"TrackType,omitempty" yaml:"TrackType,omitempty"`
	Description *Description    `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	DataFrom    *DateTime       `xml:"DataFrom,omitempty" json:"DataFrom,omitempty" yaml:"DataFrom,omitempty"`
	DataTo      *DateTime       `xml:"DataTo,omitempty" json:"DataTo,omitempty" yaml:"DataTo,omitempty"`
}

type Transport struct {
	Protocol *TransportProtocol `xml:"Protocol,omitempty" json:"Protocol,omitempty" yaml:"Protocol,omitempty"`
	Tunnel   *Transport         `xml:"Tunnel,omitempty" json:"Tunnel,omitempty" yaml:"Tunnel,omitempty"`
}

type UnableToCreatePullPointFaultType struct {
	Timestamp     DateTime               `xml:"Timestamp" json:"Timestamp" yaml:"Timestamp"`
	Originator    *EndpointReferenceType `xml:"Originator,omitempty" json:"Originator,omitempty" yaml:"Originator,omitempty"`
	ErrorCode     *string                `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	Description   []*string              `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	FaultCause    interface{}            `xml:"FaultCause" json:"FaultCause" yaml:"FaultCause"`
	TypeAttrXSI   string                 `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string                 `xml:"xmlns:objtype,attr,omitempty"`

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
	Timestamp     DateTime               `xml:"Timestamp" json:"Timestamp" yaml:"Timestamp"`
	Originator    *EndpointReferenceType `xml:"Originator,omitempty" json:"Originator,omitempty" yaml:"Originator,omitempty"`
	ErrorCode     *string                `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	Description   []*string              `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	FaultCause    interface{}            `xml:"FaultCause" json:"FaultCause" yaml:"FaultCause"`
	TypeAttrXSI   string                 `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string                 `xml:"xmlns:objtype,attr,omitempty"`

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
	Timestamp     DateTime               `xml:"Timestamp" json:"Timestamp" yaml:"Timestamp"`
	Originator    *EndpointReferenceType `xml:"Originator,omitempty" json:"Originator,omitempty" yaml:"Originator,omitempty"`
	ErrorCode     *string                `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	Description   []*string              `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	FaultCause    interface{}            `xml:"FaultCause" json:"FaultCause" yaml:"FaultCause"`
	TypeAttrXSI   string                 `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string                 `xml:"xmlns:objtype,attr,omitempty"`

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
	Timestamp     DateTime               `xml:"Timestamp" json:"Timestamp" yaml:"Timestamp"`
	Originator    *EndpointReferenceType `xml:"Originator,omitempty" json:"Originator,omitempty" yaml:"Originator,omitempty"`
	ErrorCode     *string                `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	Description   []*string              `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	FaultCause    interface{}            `xml:"FaultCause" json:"FaultCause" yaml:"FaultCause"`
	TypeAttrXSI   string                 `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string                 `xml:"xmlns:objtype,attr,omitempty"`

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
	Timestamp     DateTime               `xml:"Timestamp" json:"Timestamp" yaml:"Timestamp"`
	Originator    *EndpointReferenceType `xml:"Originator,omitempty" json:"Originator,omitempty" yaml:"Originator,omitempty"`
	ErrorCode     *string                `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	Description   []*string              `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	FaultCause    interface{}            `xml:"FaultCause" json:"FaultCause" yaml:"FaultCause"`
	MinimumTime   *DateTime              `xml:"MinimumTime,omitempty" json:"MinimumTime,omitempty" yaml:"MinimumTime,omitempty"`
	MaximumTime   *DateTime              `xml:"MaximumTime,omitempty" json:"MaximumTime,omitempty" yaml:"MaximumTime,omitempty"`
	TypeAttrXSI   string                 `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string                 `xml:"xmlns:objtype,attr,omitempty"`

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
	Timestamp     DateTime               `xml:"Timestamp" json:"Timestamp" yaml:"Timestamp"`
	Originator    *EndpointReferenceType `xml:"Originator,omitempty" json:"Originator,omitempty" yaml:"Originator,omitempty"`
	ErrorCode     *string                `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	Description   []*string              `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	FaultCause    interface{}            `xml:"FaultCause" json:"FaultCause" yaml:"FaultCause"`
	MinimumTime   *DateTime              `xml:"MinimumTime,omitempty" json:"MinimumTime,omitempty" yaml:"MinimumTime,omitempty"`
	MaximumTime   *DateTime              `xml:"MaximumTime,omitempty" json:"MaximumTime,omitempty" yaml:"MaximumTime,omitempty"`
	TypeAttrXSI   string                 `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string                 `xml:"xmlns:objtype,attr,omitempty"`

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
	Timestamp          DateTime               `xml:"Timestamp" json:"Timestamp" yaml:"Timestamp"`
	Originator         *EndpointReferenceType `xml:"Originator,omitempty" json:"Originator,omitempty" yaml:"Originator,omitempty"`
	ErrorCode          *string                `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	Description        []*string              `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	FaultCause         interface{}            `xml:"FaultCause" json:"FaultCause" yaml:"FaultCause"`
	UnrecognizedPolicy []*string              `xml:"UnrecognizedPolicy,omitempty" json:"UnrecognizedPolicy,omitempty" yaml:"UnrecognizedPolicy,omitempty"`
	TypeAttrXSI        string                 `xml:"xsi:type,attr,omitempty"`
	TypeNamespace      string                 `xml:"xmlns:objtype,attr,omitempty"`

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
	Timestamp         DateTime               `xml:"Timestamp" json:"Timestamp" yaml:"Timestamp"`
	Originator        *EndpointReferenceType `xml:"Originator,omitempty" json:"Originator,omitempty" yaml:"Originator,omitempty"`
	ErrorCode         *string                `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	Description       []*string              `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	FaultCause        interface{}            `xml:"FaultCause" json:"FaultCause" yaml:"FaultCause"`
	UnsupportedPolicy []*string              `xml:"UnsupportedPolicy,omitempty" json:"UnsupportedPolicy,omitempty" yaml:"UnsupportedPolicy,omitempty"`
	TypeAttrXSI       string                 `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string                 `xml:"xmlns:objtype,attr,omitempty"`

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

type UpgradeType struct {
	SupportedEnvelope []*SupportedEnvType `xml:"SupportedEnvelope" json:"SupportedEnvelope" yaml:"SupportedEnvelope"`
}

type UseRaw struct {
}

type User struct {
	Username  *string        `xml:"Username,omitempty" json:"Username,omitempty" yaml:"Username,omitempty"`
	Password  *string        `xml:"Password,omitempty" json:"Password,omitempty" yaml:"Password,omitempty"`
	UserLevel *UserLevel     `xml:"UserLevel,omitempty" json:"UserLevel,omitempty" yaml:"UserLevel,omitempty"`
	Extension *UserExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type UserExtension []interface{}

type VideoAnalyticsConfiguration struct {
	Name                         *Name                         `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	UseCount                     *int                          `xml:"UseCount,omitempty" json:"UseCount,omitempty" yaml:"UseCount,omitempty"`
	Token                        ReferenceToken                `xml:"token,attr,omitempty" json:"token,attr,omitempty" yaml:"token,attr,omitempty"`
	AnalyticsEngineConfiguration *AnalyticsEngineConfiguration `xml:"AnalyticsEngineConfiguration,omitempty" json:"AnalyticsEngineConfiguration,omitempty" yaml:"AnalyticsEngineConfiguration,omitempty"`
	RuleEngineConfiguration      *RuleEngineConfiguration      `xml:"RuleEngineConfiguration,omitempty" json:"RuleEngineConfiguration,omitempty" yaml:"RuleEngineConfiguration,omitempty"`
	TypeAttrXSI                  string                        `xml:"xsi:type,attr,omitempty"`
	TypeNamespace                string                        `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *VideoAnalyticsConfiguration) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:VideoAnalyticsConfiguration"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://www.onvif.org/ver10/schema"
	}
}

type VideoAttributes struct {
	Bitrate   *int     `xml:"Bitrate,omitempty" json:"Bitrate,omitempty" yaml:"Bitrate,omitempty"`
	Width     *int     `xml:"Width,omitempty" json:"Width,omitempty" yaml:"Width,omitempty"`
	Height    *int     `xml:"Height,omitempty" json:"Height,omitempty" yaml:"Height,omitempty"`
	Encoding  *string  `xml:"Encoding,omitempty" json:"Encoding,omitempty" yaml:"Encoding,omitempty"`
	Framerate *float64 `xml:"Framerate,omitempty" json:"Framerate,omitempty" yaml:"Framerate,omitempty"`
}

type VideoDecoderConfigurationOptions struct {
	JpegDecOptions  *JpegDecOptions                            `xml:"JpegDecOptions,omitempty" json:"JpegDecOptions,omitempty" yaml:"JpegDecOptions,omitempty"`
	H264DecOptions  *H264DecOptions                            `xml:"H264DecOptions,omitempty" json:"H264DecOptions,omitempty" yaml:"H264DecOptions,omitempty"`
	Mpeg4DecOptions *Mpeg4DecOptions                           `xml:"Mpeg4DecOptions,omitempty" json:"Mpeg4DecOptions,omitempty" yaml:"Mpeg4DecOptions,omitempty"`
	Extension       *VideoDecoderConfigurationOptionsExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type VideoDecoderConfigurationOptionsExtension []interface{}

type VideoEncoder2Configuration struct {
	Name                *Name                   `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	UseCount            *int                    `xml:"UseCount,omitempty" json:"UseCount,omitempty" yaml:"UseCount,omitempty"`
	Token               ReferenceToken          `xml:"token,attr,omitempty" json:"token,attr,omitempty" yaml:"token,attr,omitempty"`
	GovLength           int                     `xml:"GovLength,attr,omitempty" json:"GovLength,attr,omitempty" yaml:"GovLength,attr,omitempty"`
	Profile             string                  `xml:"Profile,attr,omitempty" json:"Profile,attr,omitempty" yaml:"Profile,attr,omitempty"`
	GuaranteedFrameRate bool                    `xml:"GuaranteedFrameRate,attr,omitempty" json:"GuaranteedFrameRate,attr,omitempty" yaml:"GuaranteedFrameRate,attr,omitempty"`
	Encoding            *string                 `xml:"Encoding,omitempty" json:"Encoding,omitempty" yaml:"Encoding,omitempty"`
	Resolution          *VideoResolution2       `xml:"Resolution,omitempty" json:"Resolution,omitempty" yaml:"Resolution,omitempty"`
	RateControl         *VideoRateControl2      `xml:"RateControl,omitempty" json:"RateControl,omitempty" yaml:"RateControl,omitempty"`
	Multicast           *MulticastConfiguration `xml:"Multicast,omitempty" json:"Multicast,omitempty" yaml:"Multicast,omitempty"`
	Quality             *float64                `xml:"Quality,omitempty" json:"Quality,omitempty" yaml:"Quality,omitempty"`
	TypeAttrXSI         string                  `xml:"xsi:type,attr,omitempty"`
	TypeNamespace       string                  `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *VideoEncoder2Configuration) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:VideoEncoder2Configuration"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://www.onvif.org/ver10/schema"
	}
}

type VideoEncoder2ConfigurationOptions struct {
	Encoding                     *string             `xml:"Encoding,omitempty" json:"Encoding,omitempty" yaml:"Encoding,omitempty"`
	QualityRange                 *FloatRange         `xml:"QualityRange,omitempty" json:"QualityRange,omitempty" yaml:"QualityRange,omitempty"`
	ResolutionsAvailable         []*VideoResolution2 `xml:"ResolutionsAvailable,omitempty" json:"ResolutionsAvailable,omitempty" yaml:"ResolutionsAvailable,omitempty"`
	BitrateRange                 *IntRange           `xml:"BitrateRange,omitempty" json:"BitrateRange,omitempty" yaml:"BitrateRange,omitempty"`
	GovLengthRange               IntList             `xml:"GovLengthRange,attr,omitempty" json:"GovLengthRange,attr,omitempty" yaml:"GovLengthRange,attr,omitempty"`
	FrameRatesSupported          FloatList           `xml:"FrameRatesSupported,attr,omitempty" json:"FrameRatesSupported,attr,omitempty" yaml:"FrameRatesSupported,attr,omitempty"`
	ProfilesSupported            StringAttrList      `xml:"ProfilesSupported,attr,omitempty" json:"ProfilesSupported,attr,omitempty" yaml:"ProfilesSupported,attr,omitempty"`
	ConstantBitRateSupported     bool                `xml:"ConstantBitRateSupported,attr,omitempty" json:"ConstantBitRateSupported,attr,omitempty" yaml:"ConstantBitRateSupported,attr,omitempty"`
	GuaranteedFrameRateSupported bool                `xml:"GuaranteedFrameRateSupported,attr,omitempty" json:"GuaranteedFrameRateSupported,attr,omitempty" yaml:"GuaranteedFrameRateSupported,attr,omitempty"`
}

type VideoEncoderConfiguration struct {
	Name                *Name                   `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	UseCount            *int                    `xml:"UseCount,omitempty" json:"UseCount,omitempty" yaml:"UseCount,omitempty"`
	Token               ReferenceToken          `xml:"token,attr,omitempty" json:"token,attr,omitempty" yaml:"token,attr,omitempty"`
	GuaranteedFrameRate bool                    `xml:"GuaranteedFrameRate,attr,omitempty" json:"GuaranteedFrameRate,attr,omitempty" yaml:"GuaranteedFrameRate,attr,omitempty"`
	Encoding            *VideoEncoding          `xml:"Encoding,omitempty" json:"Encoding,omitempty" yaml:"Encoding,omitempty"`
	Resolution          *VideoResolution        `xml:"Resolution,omitempty" json:"Resolution,omitempty" yaml:"Resolution,omitempty"`
	Quality             *float64                `xml:"Quality,omitempty" json:"Quality,omitempty" yaml:"Quality,omitempty"`
	RateControl         *VideoRateControl       `xml:"RateControl,omitempty" json:"RateControl,omitempty" yaml:"RateControl,omitempty"`
	MPEG4               *Mpeg4Configuration     `xml:"MPEG4,omitempty" json:"MPEG4,omitempty" yaml:"MPEG4,omitempty"`
	H264                *H264Configuration      `xml:"H264,omitempty" json:"H264,omitempty" yaml:"H264,omitempty"`
	Multicast           *MulticastConfiguration `xml:"Multicast,omitempty" json:"Multicast,omitempty" yaml:"Multicast,omitempty"`
	SessionTimeout      *Duration               `xml:"SessionTimeout,omitempty" json:"SessionTimeout,omitempty" yaml:"SessionTimeout,omitempty"`
	TypeAttrXSI         string                  `xml:"xsi:type,attr,omitempty"`
	TypeNamespace       string                  `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *VideoEncoderConfiguration) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:VideoEncoderConfiguration"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://www.onvif.org/ver10/schema"
	}
}

type VideoEncoderConfigurationOptions struct {
	QualityRange                 *IntRange                     `xml:"QualityRange,omitempty" json:"QualityRange,omitempty" yaml:"QualityRange,omitempty"`
	JPEG                         *JpegOptions                  `xml:"JPEG,omitempty" json:"JPEG,omitempty" yaml:"JPEG,omitempty"`
	MPEG4                        *Mpeg4Options                 `xml:"MPEG4,omitempty" json:"MPEG4,omitempty" yaml:"MPEG4,omitempty"`
	H264                         *H264Options                  `xml:"H264,omitempty" json:"H264,omitempty" yaml:"H264,omitempty"`
	Extension                    *VideoEncoderOptionsExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
	GuaranteedFrameRateSupported bool                          `xml:"GuaranteedFrameRateSupported,attr,omitempty" json:"GuaranteedFrameRateSupported,attr,omitempty" yaml:"GuaranteedFrameRateSupported,attr,omitempty"`
}

type VideoEncoderOptionsExtension struct {
	JPEG      *JpegOptions2                  `xml:"JPEG,omitempty" json:"JPEG,omitempty" yaml:"JPEG,omitempty"`
	MPEG4     *Mpeg4Options2                 `xml:"MPEG4,omitempty" json:"MPEG4,omitempty" yaml:"MPEG4,omitempty"`
	H264      *H264Options2                  `xml:"H264,omitempty" json:"H264,omitempty" yaml:"H264,omitempty"`
	Extension *VideoEncoderOptionsExtension2 `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type VideoEncoderOptionsExtension2 []interface{}

// Representation of a physical video outputs.
type VideoOutput struct {
	Token         ReferenceToken        `xml:"token,attr,omitempty" json:"token,attr,omitempty" yaml:"token,attr,omitempty"`
	Layout        *Layout               `xml:"Layout,omitempty" json:"Layout,omitempty" yaml:"Layout,omitempty"`
	Resolution    *VideoResolution      `xml:"Resolution,omitempty" json:"Resolution,omitempty" yaml:"Resolution,omitempty"`
	RefreshRate   *float64              `xml:"RefreshRate,omitempty" json:"RefreshRate,omitempty" yaml:"RefreshRate,omitempty"`
	AspectRatio   *float64              `xml:"AspectRatio,omitempty" json:"AspectRatio,omitempty" yaml:"AspectRatio,omitempty"`
	Extension     *VideoOutputExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
	TypeAttrXSI   string                `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string                `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *VideoOutput) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:VideoOutput"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://www.onvif.org/ver10/schema"
	}
}

type VideoOutputConfiguration struct {
	Name          *Name          `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	UseCount      *int           `xml:"UseCount,omitempty" json:"UseCount,omitempty" yaml:"UseCount,omitempty"`
	Token         ReferenceToken `xml:"token,attr,omitempty" json:"token,attr,omitempty" yaml:"token,attr,omitempty"`
	OutputToken   ReferenceToken `xml:"OutputToken,omitempty" json:"OutputToken,omitempty" yaml:"OutputToken,omitempty"`
	TypeAttrXSI   string         `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string         `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *VideoOutputConfiguration) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:VideoOutputConfiguration"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://www.onvif.org/ver10/schema"
	}
}

type VideoOutputConfigurationOptions []interface{}

type VideoOutputExtension []interface{}

type VideoRateControl struct {
	FrameRateLimit   *int `xml:"FrameRateLimit,omitempty" json:"FrameRateLimit,omitempty" yaml:"FrameRateLimit,omitempty"`
	EncodingInterval *int `xml:"EncodingInterval,omitempty" json:"EncodingInterval,omitempty" yaml:"EncodingInterval,omitempty"`
	BitrateLimit     *int `xml:"BitrateLimit,omitempty" json:"BitrateLimit,omitempty" yaml:"BitrateLimit,omitempty"`
}

type VideoRateControl2 struct {
	FrameRateLimit  *float64 `xml:"FrameRateLimit,omitempty" json:"FrameRateLimit,omitempty" yaml:"FrameRateLimit,omitempty"`
	BitrateLimit    *int     `xml:"BitrateLimit,omitempty" json:"BitrateLimit,omitempty" yaml:"BitrateLimit,omitempty"`
	ConstantBitRate bool     `xml:"ConstantBitRate,attr,omitempty" json:"ConstantBitRate,attr,omitempty" yaml:"ConstantBitRate,attr,omitempty"`
}

type VideoResolution struct {
	Width  *int `xml:"Width,omitempty" json:"Width,omitempty" yaml:"Width,omitempty"`
	Height *int `xml:"Height,omitempty" json:"Height,omitempty" yaml:"Height,omitempty"`
}

type VideoResolution2 struct {
	Width  *int `xml:"Width,omitempty" json:"Width,omitempty" yaml:"Width,omitempty"`
	Height *int `xml:"Height,omitempty" json:"Height,omitempty" yaml:"Height,omitempty"`
}

// Representation of a physical video input.
type VideoSource struct {
	Token         ReferenceToken        `xml:"token,attr,omitempty" json:"token,attr,omitempty" yaml:"token,attr,omitempty"`
	Framerate     *float64              `xml:"Framerate,omitempty" json:"Framerate,omitempty" yaml:"Framerate,omitempty"`
	Resolution    *VideoResolution      `xml:"Resolution,omitempty" json:"Resolution,omitempty" yaml:"Resolution,omitempty"`
	Imaging       *ImagingSettings      `xml:"Imaging,omitempty" json:"Imaging,omitempty" yaml:"Imaging,omitempty"`
	Extension     *VideoSourceExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
	TypeAttrXSI   string                `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string                `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *VideoSource) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:VideoSource"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://www.onvif.org/ver10/schema"
	}
}

type VideoSourceConfiguration struct {
	Name          *Name                              `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	UseCount      *int                               `xml:"UseCount,omitempty" json:"UseCount,omitempty" yaml:"UseCount,omitempty"`
	Token         ReferenceToken                     `xml:"token,attr,omitempty" json:"token,attr,omitempty" yaml:"token,attr,omitempty"`
	ViewMode      string                             `xml:"ViewMode,attr,omitempty" json:"ViewMode,attr,omitempty" yaml:"ViewMode,attr,omitempty"`
	SourceToken   ReferenceToken                     `xml:"SourceToken,omitempty" json:"SourceToken,omitempty" yaml:"SourceToken,omitempty"`
	Bounds        *IntRectangle                      `xml:"Bounds,omitempty" json:"Bounds,omitempty" yaml:"Bounds,omitempty"`
	Extension     *VideoSourceConfigurationExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
	TypeAttrXSI   string                             `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string                             `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

func (t *VideoSourceConfiguration) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:VideoSourceConfiguration"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://www.onvif.org/ver10/schema"
	}
}

type VideoSourceConfigurationExtension struct {
	Rotate    *Rotate                             `xml:"Rotate,omitempty" json:"Rotate,omitempty" yaml:"Rotate,omitempty"`
	Extension *VideoSourceConfigurationExtension2 `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type VideoSourceConfigurationExtension2 struct {
	LensDescription  []*LensDescription `xml:"LensDescription,omitempty" json:"LensDescription,omitempty" yaml:"LensDescription,omitempty"`
	SceneOrientation *SceneOrientation  `xml:"SceneOrientation,omitempty" json:"SceneOrientation,omitempty" yaml:"SceneOrientation,omitempty"`
}

type VideoSourceConfigurationOptions struct {
	BoundsRange                *IntRectangleRange                        `xml:"BoundsRange,omitempty" json:"BoundsRange,omitempty" yaml:"BoundsRange,omitempty"`
	VideoSourceTokensAvailable []ReferenceToken                          `xml:"VideoSourceTokensAvailable,omitempty" json:"VideoSourceTokensAvailable,omitempty" yaml:"VideoSourceTokensAvailable,omitempty"`
	Extension                  *VideoSourceConfigurationOptionsExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
	MaximumNumberOfProfiles    int                                       `xml:"MaximumNumberOfProfiles,attr,omitempty" json:"MaximumNumberOfProfiles,attr,omitempty" yaml:"MaximumNumberOfProfiles,attr,omitempty"`
}

type VideoSourceConfigurationOptionsExtension struct {
	Rotate    *RotateOptions                             `xml:"Rotate,omitempty" json:"Rotate,omitempty" yaml:"Rotate,omitempty"`
	Extension *VideoSourceConfigurationOptionsExtension2 `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type VideoSourceConfigurationOptionsExtension2 struct {
	SceneOrientationMode []*SceneOrientationMode `xml:"SceneOrientationMode,omitempty" json:"SceneOrientationMode,omitempty" yaml:"SceneOrientationMode,omitempty"`
}

type VideoSourceExtension struct {
	Imaging   *ImagingSettings20     `xml:"Imaging,omitempty" json:"Imaging,omitempty" yaml:"Imaging,omitempty"`
	Extension *VideoSourceExtension2 `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type VideoSourceExtension2 []interface{}

type WhiteBalance struct {
	Mode   *WhiteBalanceMode `xml:"Mode,omitempty" json:"Mode,omitempty" yaml:"Mode,omitempty"`
	CrGain *float64          `xml:"CrGain,omitempty" json:"CrGain,omitempty" yaml:"CrGain,omitempty"`
	CbGain *float64          `xml:"CbGain,omitempty" json:"CbGain,omitempty" yaml:"CbGain,omitempty"`
}

type WhiteBalance20 struct {
	Mode      *WhiteBalanceMode        `xml:"Mode,omitempty" json:"Mode,omitempty" yaml:"Mode,omitempty"`
	CrGain    *float64                 `xml:"CrGain,omitempty" json:"CrGain,omitempty" yaml:"CrGain,omitempty"`
	CbGain    *float64                 `xml:"CbGain,omitempty" json:"CbGain,omitempty" yaml:"CbGain,omitempty"`
	Extension *WhiteBalance20Extension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type WhiteBalance20Extension []interface{}

type WhiteBalanceOptions struct {
	Mode   []*WhiteBalanceMode `xml:"Mode,omitempty" json:"Mode,omitempty" yaml:"Mode,omitempty"`
	YrGain *FloatRange         `xml:"YrGain,omitempty" json:"YrGain,omitempty" yaml:"YrGain,omitempty"`
	YbGain *FloatRange         `xml:"YbGain,omitempty" json:"YbGain,omitempty" yaml:"YbGain,omitempty"`
}

type WhiteBalanceOptions20 struct {
	Mode      []*WhiteBalanceMode             `xml:"Mode,omitempty" json:"Mode,omitempty" yaml:"Mode,omitempty"`
	YrGain    *FloatRange                     `xml:"YrGain,omitempty" json:"YrGain,omitempty" yaml:"YrGain,omitempty"`
	YbGain    *FloatRange                     `xml:"YbGain,omitempty" json:"YbGain,omitempty" yaml:"YbGain,omitempty"`
	Extension *WhiteBalanceOptions20Extension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type WhiteBalanceOptions20Extension []interface{}

type WideDynamicRange struct {
	Mode  *WideDynamicMode `xml:"Mode,omitempty" json:"Mode,omitempty" yaml:"Mode,omitempty"`
	Level *float64         `xml:"Level,omitempty" json:"Level,omitempty" yaml:"Level,omitempty"`
}

// Type describing whether WDR mode is enabled or disabled
//                     (on/off).
type WideDynamicRange20 struct {
	Mode  *WideDynamicMode `xml:"Mode,omitempty" json:"Mode,omitempty" yaml:"Mode,omitempty"`
	Level *float64         `xml:"Level,omitempty" json:"Level,omitempty" yaml:"Level,omitempty"`
}

type WideDynamicRangeOptions struct {
	Mode  []*WideDynamicMode `xml:"Mode,omitempty" json:"Mode,omitempty" yaml:"Mode,omitempty"`
	Level *FloatRange        `xml:"Level,omitempty" json:"Level,omitempty" yaml:"Level,omitempty"`
}

type WideDynamicRangeOptions20 struct {
	Mode  []*WideDynamicMode `xml:"Mode,omitempty" json:"Mode,omitempty" yaml:"Mode,omitempty"`
	Level *FloatRange        `xml:"Level,omitempty" json:"Level,omitempty" yaml:"Level,omitempty"`
}

type ZoomLimits struct {
	Range *Space1DDescription `xml:"Range,omitempty" json:"Range,omitempty" yaml:"Range,omitempty"`
}

type Base64Binary struct {
	ContentType string `xml:"contentType,attr,omitempty" json:"contentType,attr,omitempty" yaml:"contentType,attr,omitempty"`
}

type Detail []interface{}

type Faultcode struct {
	Value   *FaultcodeEnum `xml:"Value,omitempty" json:"Value,omitempty" yaml:"Value,omitempty"`
	Subcode *Subcode       `xml:"Subcode,omitempty" json:"Subcode,omitempty" yaml:"Subcode,omitempty"`
}

type Faultreason struct {
	Text []*Reasontext `xml:"Text" json:"Text" yaml:"Text"`
}

type HexBinary struct {
	ContentType string `xml:"contentType,attr,omitempty" json:"contentType,attr,omitempty" yaml:"contentType,attr,omitempty"`
}

type Reasontext struct {
	Content *string `xml:"Content,omitempty" json:"Content,omitempty" yaml:"Content,omitempty"`
	Lang    string  `xml:"lang,attr,omitempty" json:"lang,attr,omitempty" yaml:"lang,attr,omitempty"`
}

type Subcode struct {
	Value   *string  `xml:"Value,omitempty" json:"Value,omitempty" yaml:"Value,omitempty"`
	Subcode *Subcode `xml:"Subcode,omitempty" json:"Subcode,omitempty" yaml:"Subcode,omitempty"`
}
