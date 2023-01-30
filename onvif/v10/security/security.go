package security

import (
	"reflect"

	"code.byted.org/videoarch/go-onvif/onvif/common"
)

var Namespace = "http://www.onvif.org/ver10/advancedsecurity/wsdl"

// NewDot1X creates an initializes a Dot1X.
func NewDot1X(endpoint string, cli common.Client) Dot1X {
	return &dot1X{cli: cli, Endpoint: endpoint}
}

// Dot1X was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type Dot1X interface {
	OptAddCertPathValidationPolicyAssignment(AddCertPathValidationPolicyAssignment AddCertPathValidationPolicyAssignment) (*AddCertPathValidationPolicyAssignmentResponse, *common.Fault)

	OptAddDot1XConfiguration(AddDot1XConfiguration AddDot1XConfiguration) (*AddDot1XConfigurationResponse, *common.Fault)

	OptAddServerCertificateAssignment(AddServerCertificateAssignment AddServerCertificateAssignment) (*AddServerCertificateAssignmentResponse, *common.Fault)

	OptCreateCertPathValidationPolicy(CreateCertPathValidationPolicy CreateCertPathValidationPolicy) (*CreateCertPathValidationPolicyResponse, *common.Fault)

	OptCreateCertificationPath(CreateCertificationPath CreateCertificationPath) (*CreateCertificationPathResponse, *common.Fault)

	OptCreatePKCS10CSR(CreatePKCS10CSR CreatePKCS10CSR) (*CreatePKCS10CSRResponse, *common.Fault)

	OptCreateRSAKeyPair(CreateRSAKeyPair CreateRSAKeyPair) (*CreateRSAKeyPairResponse, *common.Fault)

	OptCreateSelfSignedCertificate(CreateSelfSignedCertificate CreateSelfSignedCertificate) (*CreateSelfSignedCertificateResponse, *common.Fault)

	OptDeleteCRL(DeleteCRL DeleteCRL) (*DeleteCRLResponse, *common.Fault)

	OptDeleteCertPathValidationPolicy(DeleteCertPathValidationPolicy DeleteCertPathValidationPolicy) (*DeleteCertPathValidationPolicyResponse, *common.Fault)

	OptDeleteCertificate(DeleteCertificate DeleteCertificate) (*DeleteCertificateResponse, *common.Fault)

	OptDeleteCertificationPath(DeleteCertificationPath DeleteCertificationPath) (*DeleteCertificationPathResponse, *common.Fault)

	OptDeleteDot1XConfiguration(DeleteDot1XConfiguration DeleteDot1XConfiguration) (*DeleteDot1XConfigurationResponse, *common.Fault)

	OptDeleteKey(DeleteKey DeleteKey) (*DeleteKeyResponse, *common.Fault)

	OptDeleteNetworkInterfaceDot1XConfiguration(DeleteNetworkInterfaceDot1XConfiguration DeleteNetworkInterfaceDot1XConfiguration) (*DeleteNetworkInterfaceDot1XConfigurationResponse, *common.Fault)

	OptDeletePassphrase(DeletePassphrase DeletePassphrase) (*DeletePassphraseResponse, *common.Fault)

	OptGetAllCRLs(GetAllCRLs GetAllCRLs) (*GetAllCRLsResponse, *common.Fault)

	OptGetAllCertPathValidationPolicies(GetAllCertPathValidationPolicies GetAllCertPathValidationPolicies) (*GetAllCertPathValidationPoliciesResponse, *common.Fault)

	OptGetAllCertificates(GetAllCertificates GetAllCertificates) (*GetAllCertificatesResponse, *common.Fault)

	OptGetAllCertificationPaths(GetAllCertificationPaths GetAllCertificationPaths) (*GetAllCertificationPathsResponse, *common.Fault)

	OptGetAllDot1XConfigurations(GetAllDot1XConfigurations GetAllDot1XConfigurations) (*GetAllDot1XConfigurationsResponse, *common.Fault)

	OptGetAllKeys(GetAllKeys GetAllKeys) (*GetAllKeysResponse, *common.Fault)

	OptGetAllPassphrases(GetAllPassphrases GetAllPassphrases) (*GetAllPassphrasesResponse, *common.Fault)

	OptGetAssignedCertPathValidationPolicies(GetAssignedCertPathValidationPolicies GetAssignedCertPathValidationPolicies) (*GetAssignedCertPathValidationPoliciesResponse, *common.Fault)

	OptGetAssignedServerCertificates(GetAssignedServerCertificates GetAssignedServerCertificates) (*GetAssignedServerCertificatesResponse, *common.Fault)

	OptGetCRL(GetCRL GetCRL) (*GetCRLResponse, *common.Fault)

	OptGetCertPathValidationPolicy(GetCertPathValidationPolicy GetCertPathValidationPolicy) (*GetCertPathValidationPolicyResponse, *common.Fault)

	OptGetCertificate(GetCertificate GetCertificate) (*GetCertificateResponse, *common.Fault)

	OptGetCertificationPath(GetCertificationPath GetCertificationPath) (*GetCertificationPathResponse, *common.Fault)

	OptGetClientAuthenticationRequired(GetClientAuthenticationRequired GetClientAuthenticationRequired) (*GetClientAuthenticationRequiredResponse, *common.Fault)

	OptGetCnMapsToUser(GetCnMapsToUser GetCnMapsToUser) (*GetCnMapsToUserResponse, *common.Fault)

	OptGetDot1XConfiguration(GetDot1XConfiguration GetDot1XConfiguration) (*GetDot1XConfigurationResponse, *common.Fault)

	OptGetEnabledTLSVersions(GetEnabledTLSVersions GetEnabledTLSVersions) (*GetEnabledTLSVersionsResponse, *common.Fault)

	OptGetKeyStatus(GetKeyStatus GetKeyStatus) (*GetKeyStatusResponse, *common.Fault)

	OptGetNetworkInterfaceDot1XConfiguration(GetNetworkInterfaceDot1XConfiguration GetNetworkInterfaceDot1XConfiguration) (*GetNetworkInterfaceDot1XConfigurationResponse, *common.Fault)

	OptGetPrivateKeyStatus(GetPrivateKeyStatus GetPrivateKeyStatus) (*GetPrivateKeyStatusResponse, *common.Fault)

	OptGetServiceCapabilities(GetServiceCapabilities GetServiceCapabilities) (*GetServiceCapabilitiesResponse, *common.Fault)

	OptRemoveCertPathValidationPolicyAssignment(RemoveCertPathValidationPolicyAssignment RemoveCertPathValidationPolicyAssignment) (*RemoveCertPathValidationPolicyAssignmentResponse, *common.Fault)

	OptRemoveServerCertificateAssignment(RemoveServerCertificateAssignment RemoveServerCertificateAssignment) (*RemoveServerCertificateAssignmentResponse, *common.Fault)

	OptReplaceCertPathValidationPolicyAssignment(ReplaceCertPathValidationPolicyAssignment ReplaceCertPathValidationPolicyAssignment) (*ReplaceCertPathValidationPolicyAssignmentResponse, *common.Fault)

	OptReplaceServerCertificateAssignment(ReplaceServerCertificateAssignment ReplaceServerCertificateAssignment) (*ReplaceServerCertificateAssignmentResponse, *common.Fault)

	OptSetClientAuthenticationRequired(SetClientAuthenticationRequired SetClientAuthenticationRequired) (*SetClientAuthenticationRequiredResponse, *common.Fault)

	OptSetCnMapsToUser(SetCnMapsToUser SetCnMapsToUser) (*SetCnMapsToUserResponse, *common.Fault)

	OptSetEnabledTLSVersions(SetEnabledTLSVersions SetEnabledTLSVersions) (*SetEnabledTLSVersionsResponse, *common.Fault)

	OptSetNetworkInterfaceDot1XConfiguration(SetNetworkInterfaceDot1XConfiguration SetNetworkInterfaceDot1XConfiguration) (*SetNetworkInterfaceDot1XConfigurationResponse, *common.Fault)

	OptUploadCRL(UploadCRL UploadCRL) (*UploadCRLResponse, *common.Fault)

	OptUploadCertificate(UploadCertificate UploadCertificate) (*UploadCertificateResponse, *common.Fault)

	OptUploadCertificateWithPrivateKeyInPKCS12(UploadCertificateWithPrivateKeyInPKCS12 UploadCertificateWithPrivateKeyInPKCS12) (*UploadCertificateWithPrivateKeyInPKCS12Response, *common.Fault)

	OptUploadKeyPairInPKCS8(UploadKeyPairInPKCS8 UploadKeyPairInPKCS8) (*UploadKeyPairInPKCS8Response, *common.Fault)

	OptUploadPassphrase(UploadPassphrase UploadPassphrase) (*UploadPassphraseResponse, *common.Fault)
}
type DateTime string

type Duration string

type Base64DERencodedASN1Value []byte

type NCName string

type CRLID NCName

type CertPathValidationPolicyID NCName

type CertificateID NCName

type CertificationPathID NCName

type DNAttributeType string

type DNAttributeValue string

type Dot1XID NCName

type DotDecimalOID string

type KeyID NCName

type KeyStatus string

type TLSVersions string

func (v KeyStatus) Validate() bool {
	for _, vv := range []string{
		"ok",
		"generating",
		"corrupt",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

type PassphraseID NCName

type AddCertPathValidationPolicyAssignment struct {
	CertPathValidationPolicyID CertPathValidationPolicyID `xml:"CertPathValidationPolicyID,omitempty" json:"CertPathValidationPolicyID,omitempty" yaml:"CertPathValidationPolicyID,omitempty"`
}

type AddCertPathValidationPolicyAssignmentResponse struct {
}

type AddDot1XConfiguration struct {
	Dot1XConfiguration Dot1XConfiguration `xml:"Dot1XConfiguration,omitempty" json:"Dot1XConfiguration,omitempty" yaml:"Dot1XConfiguration,omitempty"`
}

type AddDot1XConfigurationResponse struct {
	Dot1XID Dot1XID `xml:"Dot1XID,omitempty" json:"Dot1XID,omitempty" yaml:"Dot1XID,omitempty"`
}

type AddServerCertificateAssignment struct {
	CertificationPathID CertificationPathID `xml:"CertificationPathID,omitempty" json:"CertificationPathID,omitempty" yaml:"CertificationPathID,omitempty"`
}

type AddServerCertificateAssignmentResponse struct {
}

type AlgorithmIdentifier struct {
	Algorithm     DotDecimalOID             `xml:"algorithm,omitempty" json:"algorithm,omitempty" yaml:"algorithm,omitempty"`
	Parameters    Base64DERencodedASN1Value `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
	AnyParameters []string                  `xml:"anyParameters>anyParameters,omitempty" json:"anyParameters>anyParameters,omitempty" yaml:"anyParameters>anyParameters,omitempty"`
}

type BasicRequestAttribute struct {
	OID   DotDecimalOID             `xml:"OID,omitempty" json:"OID,omitempty" yaml:"OID,omitempty"`
	Value Base64DERencodedASN1Value `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
}

type CRL struct {
	CRLID      CRLID                     `xml:"CRLID,omitempty" json:"CRLID,omitempty" yaml:"CRLID,omitempty"`
	Alias      string                    `xml:"Alias,omitempty" json:"Alias,omitempty" yaml:"Alias,omitempty"`
	CRLContent Base64DERencodedASN1Value `xml:"CRLContent,omitempty" json:"CRLContent,omitempty" yaml:"CRLContent,omitempty"`
}

type CSRAttribute struct {
	X509v3Extension       X509v3Extension       `xml:"X509v3Extension,omitempty" json:"X509v3Extension,omitempty" yaml:"X509v3Extension,omitempty"`
	BasicRequestAttribute BasicRequestAttribute `xml:"BasicRequestAttribute,omitempty" json:"BasicRequestAttribute,omitempty" yaml:"BasicRequestAttribute,omitempty"`
	AnyAttribute          []string              `xml:"anyAttribute>anyAttribute,omitempty" json:"anyAttribute>anyAttribute,omitempty" yaml:"anyAttribute>anyAttribute,omitempty"`
}

type Capabilities struct {
	KeystoreCapabilities  KeystoreCapabilities  `xml:"KeystoreCapabilities,omitempty" json:"KeystoreCapabilities,omitempty" yaml:"KeystoreCapabilities,omitempty"`
	TLSServerCapabilities TLSServerCapabilities `xml:"TLSServerCapabilities,omitempty" json:"TLSServerCapabilities,omitempty" yaml:"TLSServerCapabilities,omitempty"`
	Dot1XCapabilities     Dot1XCapabilities     `xml:"Dot1XCapabilities,omitempty" json:"Dot1XCapabilities,omitempty" yaml:"Dot1XCapabilities,omitempty"`
}

type CertPathValidationParameters struct {
	RequireTLSWWWClientAuthExtendedKeyUsage bool     `xml:"RequireTLSWWWClientAuthExtendedKeyUsage,omitempty" json:"RequireTLSWWWClientAuthExtendedKeyUsage,omitempty" yaml:"RequireTLSWWWClientAuthExtendedKeyUsage,omitempty"`
	UseDeltaCRLs                            bool     `xml:"UseDeltaCRLs,omitempty" json:"UseDeltaCRLs,omitempty" yaml:"UseDeltaCRLs,omitempty"`
	AnyParameters                           []string `xml:"anyParameters>anyParameters,omitempty" json:"anyParameters>anyParameters,omitempty" yaml:"anyParameters>anyParameters,omitempty"`
}

type CertPathValidationPolicy struct {
	CertPathValidationPolicyID CertPathValidationPolicyID   `xml:"CertPathValidationPolicyID,omitempty" json:"CertPathValidationPolicyID,omitempty" yaml:"CertPathValidationPolicyID,omitempty"`
	Alias                      string                       `xml:"Alias,omitempty" json:"Alias,omitempty" yaml:"Alias,omitempty"`
	Parameters                 CertPathValidationParameters `xml:"Parameters,omitempty" json:"Parameters,omitempty" yaml:"Parameters,omitempty"`
	TrustAnchor                []TrustAnchor                `xml:"TrustAnchor,omitempty" json:"TrustAnchor,omitempty" yaml:"TrustAnchor,omitempty"`
	AnyParameters              []string                     `xml:"anyParameters>anyParameters,omitempty" json:"anyParameters>anyParameters,omitempty" yaml:"anyParameters>anyParameters,omitempty"`
}

type CertificateIDs struct {
	CertificateID []CertificateID `xml:"CertificateID,omitempty" json:"CertificateID,omitempty" yaml:"CertificateID,omitempty"`
}

type CertificationPath struct {
	CertificateID []CertificateID `xml:"CertificateID,omitempty" json:"CertificateID,omitempty" yaml:"CertificateID,omitempty"`
	Alias         string          `xml:"Alias,omitempty" json:"Alias,omitempty" yaml:"Alias,omitempty"`
	AnyElement    []string        `xml:"anyElement>anyElement,omitempty" json:"anyElement>anyElement,omitempty" yaml:"anyElement>anyElement,omitempty"`
}

type CreateCertPathValidationPolicy struct {
	Alias         string                       `xml:"Alias,omitempty" json:"Alias,omitempty" yaml:"Alias,omitempty"`
	Parameters    CertPathValidationParameters `xml:"Parameters,omitempty" json:"Parameters,omitempty" yaml:"Parameters,omitempty"`
	TrustAnchor   []TrustAnchor                `xml:"TrustAnchor,omitempty" json:"TrustAnchor,omitempty" yaml:"TrustAnchor,omitempty"`
	AnyParameters []string                     `xml:"anyParameters>anyParameters,omitempty" json:"anyParameters>anyParameters,omitempty" yaml:"anyParameters>anyParameters,omitempty"`
}

type CreateCertPathValidationPolicyResponse struct {
	CertPathValidationPolicyID CertPathValidationPolicyID `xml:"CertPathValidationPolicyID,omitempty" json:"CertPathValidationPolicyID,omitempty" yaml:"CertPathValidationPolicyID,omitempty"`
}

type CreateCertificationPath struct {
	CertificateIDs CertificateIDs `xml:"CertificateIDs,omitempty" json:"CertificateIDs,omitempty" yaml:"CertificateIDs,omitempty"`
	Alias          string         `xml:"Alias,omitempty" json:"Alias,omitempty" yaml:"Alias,omitempty"`
}

type CreateCertificationPathResponse struct {
	CertificationPathID CertificationPathID `xml:"CertificationPathID,omitempty" json:"CertificationPathID,omitempty" yaml:"CertificationPathID,omitempty"`
}

type CreatePKCS10CSR struct {
	Subject            DistinguishedName   `xml:"Subject,omitempty" json:"Subject,omitempty" yaml:"Subject,omitempty"`
	KeyID              KeyID               `xml:"KeyID,omitempty" json:"KeyID,omitempty" yaml:"KeyID,omitempty"`
	CSRAttribute       []CSRAttribute      `xml:"CSRAttribute,omitempty" json:"CSRAttribute,omitempty" yaml:"CSRAttribute,omitempty"`
	SignatureAlgorithm AlgorithmIdentifier `xml:"SignatureAlgorithm,omitempty" json:"SignatureAlgorithm,omitempty" yaml:"SignatureAlgorithm,omitempty"`
}

type CreatePKCS10CSRResponse struct {
	PKCS10CSR Base64DERencodedASN1Value `xml:"PKCS10CSR,omitempty" json:"PKCS10CSR,omitempty" yaml:"PKCS10CSR,omitempty"`
}

type CreateRSAKeyPair struct {
	KeyLength uint   `xml:"KeyLength,omitempty" json:"KeyLength,omitempty" yaml:"KeyLength,omitempty"`
	Alias     string `xml:"Alias,omitempty" json:"Alias,omitempty" yaml:"Alias,omitempty"`
}

type CreateRSAKeyPairResponse struct {
	KeyID                 KeyID            `xml:"KeyID,omitempty" json:"KeyID,omitempty" yaml:"KeyID,omitempty"`
	EstimatedCreationTime *common.Duration `xml:"EstimatedCreationTime,omitempty" json:"EstimatedCreationTime,omitempty" yaml:"EstimatedCreationTime,omitempty"`
}

type CreateSelfSignedCertificate struct {
	X509Version        uint64              `xml:"X509Version,omitempty" json:"X509Version,omitempty" yaml:"X509Version,omitempty"`
	Subject            DistinguishedName   `xml:"Subject,omitempty" json:"Subject,omitempty" yaml:"Subject,omitempty"`
	KeyID              KeyID               `xml:"KeyID,omitempty" json:"KeyID,omitempty" yaml:"KeyID,omitempty"`
	Alias              string              `xml:"Alias,omitempty" json:"Alias,omitempty" yaml:"Alias,omitempty"`
	NotValidBefore     *common.DateTime    `xml:"notValidBefore,omitempty" json:"notValidBefore,omitempty" yaml:"notValidBefore,omitempty"`
	NotValidAfter      *common.DateTime    `xml:"notValidAfter,omitempty" json:"notValidAfter,omitempty" yaml:"notValidAfter,omitempty"`
	SignatureAlgorithm AlgorithmIdentifier `xml:"SignatureAlgorithm,omitempty" json:"SignatureAlgorithm,omitempty" yaml:"SignatureAlgorithm,omitempty"`
	Extension          []X509v3Extension   `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
}

type CreateSelfSignedCertificateResponse struct {
	CertificateID CertificateID `xml:"CertificateID,omitempty" json:"CertificateID,omitempty" yaml:"CertificateID,omitempty"`
}

type DNAttributeTypeAndValue struct {
	Type  DNAttributeType  `xml:"Type,omitempty" json:"Type,omitempty" yaml:"Type,omitempty"`
	Value DNAttributeValue `xml:"Value,omitempty" json:"Value,omitempty" yaml:"Value,omitempty"`
}

type DeleteCRL struct {
	CrlID CRLID `xml:"CrlID,omitempty" json:"CrlID,omitempty" yaml:"CrlID,omitempty"`
}

type DeleteCRLResponse struct {
}

type DeleteCertPathValidationPolicy struct {
	CertPathValidationPolicyID CertPathValidationPolicyID `xml:"CertPathValidationPolicyID,omitempty" json:"CertPathValidationPolicyID,omitempty" yaml:"CertPathValidationPolicyID,omitempty"`
}

type DeleteCertPathValidationPolicyResponse struct {
}

type DeleteCertificate struct {
	CertificateID CertificateID `xml:"CertificateID,omitempty" json:"CertificateID,omitempty" yaml:"CertificateID,omitempty"`
}

type DeleteCertificateResponse struct {
}

type DeleteCertificationPath struct {
	CertificationPathID CertificationPathID `xml:"CertificationPathID,omitempty" json:"CertificationPathID,omitempty" yaml:"CertificationPathID,omitempty"`
}

type DeleteCertificationPathResponse struct {
}

type DeleteDot1XConfiguration struct {
	Dot1XID Dot1XID `xml:"Dot1XID,omitempty" json:"Dot1XID,omitempty" yaml:"Dot1XID,omitempty"`
}

type DeleteDot1XConfigurationResponse struct {
}

type DeleteKey struct {
	KeyID KeyID `xml:"KeyID,omitempty" json:"KeyID,omitempty" yaml:"KeyID,omitempty"`
}

type DeleteKeyResponse struct {
}

type DeleteNetworkInterfaceDot1XConfiguration struct {
	Token string `xml:"token,omitempty" json:"token,omitempty" yaml:"token,omitempty"`
}

type DeleteNetworkInterfaceDot1XConfigurationResponse struct {
	RebootNeeded bool `xml:"RebootNeeded,omitempty" json:"RebootNeeded,omitempty" yaml:"RebootNeeded,omitempty"`
}

type DeletePassphrase struct {
	PassphraseID PassphraseID `xml:"PassphraseID,omitempty" json:"PassphraseID,omitempty" yaml:"PassphraseID,omitempty"`
}

type DeletePassphraseResponse struct {
}

type DistinguishedName struct {
	Country                    []DNAttributeValue        `xml:"Country,omitempty" json:"Country,omitempty" yaml:"Country,omitempty"`
	Organization               []DNAttributeValue        `xml:"Organization,omitempty" json:"Organization,omitempty" yaml:"Organization,omitempty"`
	OrganizationalUnit         []DNAttributeValue        `xml:"OrganizationalUnit,omitempty" json:"OrganizationalUnit,omitempty" yaml:"OrganizationalUnit,omitempty"`
	DistinguishedNameQualifier []DNAttributeValue        `xml:"DistinguishedNameQualifier,omitempty" json:"DistinguishedNameQualifier,omitempty" yaml:"DistinguishedNameQualifier,omitempty"`
	StateOrProvinceName        []DNAttributeValue        `xml:"StateOrProvinceName,omitempty" json:"StateOrProvinceName,omitempty" yaml:"StateOrProvinceName,omitempty"`
	CommonName                 []DNAttributeValue        `xml:"CommonName,omitempty" json:"CommonName,omitempty" yaml:"CommonName,omitempty"`
	SerialNumber               []DNAttributeValue        `xml:"SerialNumber,omitempty" json:"SerialNumber,omitempty" yaml:"SerialNumber,omitempty"`
	Locality                   []DNAttributeValue        `xml:"Locality,omitempty" json:"Locality,omitempty" yaml:"Locality,omitempty"`
	Title                      []DNAttributeValue        `xml:"Title,omitempty" json:"Title,omitempty" yaml:"Title,omitempty"`
	Surname                    []DNAttributeValue        `xml:"Surname,omitempty" json:"Surname,omitempty" yaml:"Surname,omitempty"`
	GivenName                  []DNAttributeValue        `xml:"GivenName,omitempty" json:"GivenName,omitempty" yaml:"GivenName,omitempty"`
	Initials                   []DNAttributeValue        `xml:"Initials,omitempty" json:"Initials,omitempty" yaml:"Initials,omitempty"`
	Pseudonym                  []DNAttributeValue        `xml:"Pseudonym,omitempty" json:"Pseudonym,omitempty" yaml:"Pseudonym,omitempty"`
	GenerationQualifier        []DNAttributeValue        `xml:"GenerationQualifier,omitempty" json:"GenerationQualifier,omitempty" yaml:"GenerationQualifier,omitempty"`
	GenericAttribute           []DNAttributeTypeAndValue `xml:"GenericAttribute,omitempty" json:"GenericAttribute,omitempty" yaml:"GenericAttribute,omitempty"`
	MultiValuedRDN             []MultiValuedRDN          `xml:"MultiValuedRDN,omitempty" json:"MultiValuedRDN,omitempty" yaml:"MultiValuedRDN,omitempty"`
	AnyAttribute               []DNAttributeValue        `xml:"anyAttribute>DomainComponent,omitempty" json:"anyAttribute>DomainComponent,omitempty" yaml:"anyAttribute>DomainComponent,omitempty"`
}

type Dot1XCapabilities []interface{}

type Dot1XConfiguration struct {
	Dot1XID Dot1XID    `xml:"Dot1XID,omitempty" json:"Dot1XID,omitempty" yaml:"Dot1XID,omitempty"`
	Alias   string     `xml:"Alias,omitempty" json:"Alias,omitempty" yaml:"Alias,omitempty"`
	Outer   Dot1XStage `xml:"Outer,omitempty" json:"Outer,omitempty" yaml:"Outer,omitempty"`
}

type Dot1XStage struct {
	Identity            string              `xml:"Identity,omitempty" json:"Identity,omitempty" yaml:"Identity,omitempty"`
	CertificationPathID CertificationPathID `xml:"CertificationPathID,omitempty" json:"CertificationPathID,omitempty" yaml:"CertificationPathID,omitempty"`
	PassphraseID        PassphraseID        `xml:"PassphraseID,omitempty" json:"PassphraseID,omitempty" yaml:"PassphraseID,omitempty"`
	Inner               *Dot1XStage         `xml:"Inner,omitempty" json:"Inner,omitempty" yaml:"Inner,omitempty"`
	Extension           Dot1XStageExtension `xml:"Extension,omitempty" json:"Extension,omitempty" yaml:"Extension,omitempty"`
	Method              string              `xml:"Method,attr,omitempty" json:"Method,attr,omitempty" yaml:"Method,attr,omitempty"`
}

type Dot1XStageExtension []interface{}

type GetAllCRLs struct {
}

type GetAllCRLsResponse struct {
	Crl []CRL `xml:"Crl,omitempty" json:"Crl,omitempty" yaml:"Crl,omitempty"`
}

type GetAllCertPathValidationPolicies struct {
}

type GetAllCertPathValidationPoliciesResponse struct {
	CertPathValidationPolicy []CertPathValidationPolicy `xml:"CertPathValidationPolicy,omitempty" json:"CertPathValidationPolicy,omitempty" yaml:"CertPathValidationPolicy,omitempty"`
}

type GetAllCertificates struct {
}

type GetAllCertificatesResponse struct {
	Certificate []X509Certificate `xml:"Certificate,omitempty" json:"Certificate,omitempty" yaml:"Certificate,omitempty"`
}

type GetAllCertificationPaths struct {
}

type GetAllCertificationPathsResponse struct {
	CertificationPathID []CertificationPathID `xml:"CertificationPathID,omitempty" json:"CertificationPathID,omitempty" yaml:"CertificationPathID,omitempty"`
}

type GetAllDot1XConfigurations struct {
}

type GetAllDot1XConfigurationsResponse struct {
	Configuration []Dot1XConfiguration `xml:"Configuration,omitempty" json:"Configuration,omitempty" yaml:"Configuration,omitempty"`
}

type GetAllKeys struct {
}

type GetAllKeysResponse struct {
	KeyAttribute []KeyAttribute `xml:"KeyAttribute,omitempty" json:"KeyAttribute,omitempty" yaml:"KeyAttribute,omitempty"`
}

type GetAllPassphrases struct {
}

type GetAllPassphrasesResponse struct {
	PassphraseAttribute []PassphraseAttribute `xml:"PassphraseAttribute,omitempty" json:"PassphraseAttribute,omitempty" yaml:"PassphraseAttribute,omitempty"`
}

type GetAssignedCertPathValidationPolicies struct {
}

type GetAssignedCertPathValidationPoliciesResponse struct {
	CertPathValidationPolicyID []CertPathValidationPolicyID `xml:"CertPathValidationPolicyID,omitempty" json:"CertPathValidationPolicyID,omitempty" yaml:"CertPathValidationPolicyID,omitempty"`
}

type GetAssignedServerCertificates struct {
}

type GetAssignedServerCertificatesResponse struct {
	CertificationPathID []CertificationPathID `xml:"CertificationPathID,omitempty" json:"CertificationPathID,omitempty" yaml:"CertificationPathID,omitempty"`
}

type GetCRL struct {
	CrlID CRLID `xml:"CrlID,omitempty" json:"CrlID,omitempty" yaml:"CrlID,omitempty"`
}

type GetCRLResponse struct {
	Crl CRL `xml:"Crl,omitempty" json:"Crl,omitempty" yaml:"Crl,omitempty"`
}

type GetCertPathValidationPolicy struct {
	CertPathValidationPolicyID CertPathValidationPolicyID `xml:"CertPathValidationPolicyID,omitempty" json:"CertPathValidationPolicyID,omitempty" yaml:"CertPathValidationPolicyID,omitempty"`
}

type GetCertPathValidationPolicyResponse struct {
	CertPathValidationPolicy CertPathValidationPolicy `xml:"CertPathValidationPolicy,omitempty" json:"CertPathValidationPolicy,omitempty" yaml:"CertPathValidationPolicy,omitempty"`
}

type GetCertificate struct {
	CertificateID CertificateID `xml:"CertificateID,omitempty" json:"CertificateID,omitempty" yaml:"CertificateID,omitempty"`
}

type GetCertificateResponse struct {
	Certificate X509Certificate `xml:"Certificate,omitempty" json:"Certificate,omitempty" yaml:"Certificate,omitempty"`
}

type GetCertificationPath struct {
	CertificationPathID CertificationPathID `xml:"CertificationPathID,omitempty" json:"CertificationPathID,omitempty" yaml:"CertificationPathID,omitempty"`
}

type GetCertificationPathResponse struct {
	CertificationPath CertificationPath `xml:"CertificationPath,omitempty" json:"CertificationPath,omitempty" yaml:"CertificationPath,omitempty"`
}

type GetClientAuthenticationRequired struct {
}

type GetClientAuthenticationRequiredResponse struct {
	ClientAuthenticationRequired bool `xml:"clientAuthenticationRequired,omitempty" json:"clientAuthenticationRequired,omitempty" yaml:"clientAuthenticationRequired,omitempty"`
}

type GetCnMapsToUser struct {
}

type GetCnMapsToUserResponse struct {
	CnMapsToUser bool `xml:"cnMapsToUser,omitempty" json:"cnMapsToUser,omitempty" yaml:"cnMapsToUser,omitempty"`
}

type GetDot1XConfiguration struct {
	Dot1XID Dot1XID `xml:"Dot1XID,omitempty" json:"Dot1XID,omitempty" yaml:"Dot1XID,omitempty"`
}

type GetDot1XConfigurationResponse struct {
	Dot1XConfiguration Dot1XConfiguration `xml:"Dot1XConfiguration,omitempty" json:"Dot1XConfiguration,omitempty" yaml:"Dot1XConfiguration,omitempty"`
}

type GetEnabledTLSVersions struct {
}

type GetEnabledTLSVersionsResponse struct {
	Versions TLSVersions `xml:"Versions,omitempty" json:"Versions,omitempty" yaml:"Versions,omitempty"`
}

type GetKeyStatus struct {
	KeyID KeyID `xml:"KeyID,omitempty" json:"KeyID,omitempty" yaml:"KeyID,omitempty"`
}

type GetKeyStatusResponse struct {
	KeyStatus string `xml:"KeyStatus,omitempty" json:"KeyStatus,omitempty" yaml:"KeyStatus,omitempty"`
}

type GetNetworkInterfaceDot1XConfiguration struct {
	Token string `xml:"token,omitempty" json:"token,omitempty" yaml:"token,omitempty"`
}

type GetNetworkInterfaceDot1XConfigurationResponse struct {
	Dot1XID Dot1XID `xml:"Dot1XID,omitempty" json:"Dot1XID,omitempty" yaml:"Dot1XID,omitempty"`
}

type GetPrivateKeyStatus struct {
	KeyID KeyID `xml:"KeyID,omitempty" json:"KeyID,omitempty" yaml:"KeyID,omitempty"`
}

type GetPrivateKeyStatusResponse struct {
	HasPrivateKey bool `xml:"hasPrivateKey,omitempty" json:"hasPrivateKey,omitempty" yaml:"hasPrivateKey,omitempty"`
}

type GetServiceCapabilities struct {
}

type GetServiceCapabilitiesResponse struct {
	Capabilities Capabilities `xml:"Capabilities,omitempty" json:"Capabilities,omitempty" yaml:"Capabilities,omitempty"`
}

type KeyAttribute struct {
	KeyID               KeyID    `xml:"KeyID,omitempty" json:"KeyID,omitempty" yaml:"KeyID,omitempty"`
	Alias               string   `xml:"Alias,omitempty" json:"Alias,omitempty" yaml:"Alias,omitempty"`
	HasPrivateKey       bool     `xml:"hasPrivateKey,omitempty" json:"hasPrivateKey,omitempty" yaml:"hasPrivateKey,omitempty"`
	KeyStatus           string   `xml:"KeyStatus,omitempty" json:"KeyStatus,omitempty" yaml:"KeyStatus,omitempty"`
	ExternallyGenerated bool     `xml:"externallyGenerated,omitempty" json:"externallyGenerated,omitempty" yaml:"externallyGenerated,omitempty"`
	SecurelyStored      bool     `xml:"securelyStored,omitempty" json:"securelyStored,omitempty" yaml:"securelyStored,omitempty"`
	Extension           []string `xml:"Extension>Extension,omitempty" json:"Extension>Extension,omitempty" yaml:"Extension>Extension,omitempty"`
}

type KeystoreCapabilities struct {
	SignatureAlgorithms                                []AlgorithmIdentifier `xml:"SignatureAlgorithms,omitempty" json:"SignatureAlgorithms,omitempty" yaml:"SignatureAlgorithms,omitempty"`
	AnyElement                                         []string              `xml:"anyElement>anyElement,omitempty" json:"anyElement>anyElement,omitempty" yaml:"anyElement>anyElement,omitempty"`
	MaximumNumberOfKeys                                uint64                `xml:"MaximumNumberOfKeys,attr,omitempty" json:"MaximumNumberOfKeys,attr,omitempty" yaml:"MaximumNumberOfKeys,attr,omitempty"`
	MaximumNumberOfCertificates                        uint64                `xml:"MaximumNumberOfCertificates,attr,omitempty" json:"MaximumNumberOfCertificates,attr,omitempty" yaml:"MaximumNumberOfCertificates,attr,omitempty"`
	MaximumNumberOfCertificationPaths                  uint64                `xml:"MaximumNumberOfCertificationPaths,attr,omitempty" json:"MaximumNumberOfCertificationPaths,attr,omitempty" yaml:"MaximumNumberOfCertificationPaths,attr,omitempty"`
	RSAKeyPairGeneration                               bool                  `xml:"RSAKeyPairGeneration,attr,omitempty" json:"RSAKeyPairGeneration,attr,omitempty" yaml:"RSAKeyPairGeneration,attr,omitempty"`
	RSAKeyLengths                                      int                   `xml:"RSAKeyLengths,attr,omitempty" json:"RSAKeyLengths,attr,omitempty" yaml:"RSAKeyLengths,attr,omitempty"`
	PKCS10ExternalCertificationWithRSA                 bool                  `xml:"PKCS10ExternalCertificationWithRSA,attr,omitempty" json:"PKCS10ExternalCertificationWithRSA,attr,omitempty" yaml:"PKCS10ExternalCertificationWithRSA,attr,omitempty"`
	SelfSignedCertificateCreationWithRSA               bool                  `xml:"SelfSignedCertificateCreationWithRSA,attr,omitempty" json:"SelfSignedCertificateCreationWithRSA,attr,omitempty" yaml:"SelfSignedCertificateCreationWithRSA,attr,omitempty"`
	X509Versions                                       string                `xml:"X509Versions,attr,omitempty" json:"X509Versions,attr,omitempty" yaml:"X509Versions,attr,omitempty"`
	MaximumNumberOfPassphrases                         uint                  `xml:"MaximumNumberOfPassphrases,attr,omitempty" json:"MaximumNumberOfPassphrases,attr,omitempty" yaml:"MaximumNumberOfPassphrases,attr,omitempty"`
	PKCS8RSAKeyPairUpload                              bool                  `xml:"PKCS8RSAKeyPairUpload,attr,omitempty" json:"PKCS8RSAKeyPairUpload,attr,omitempty" yaml:"PKCS8RSAKeyPairUpload,attr,omitempty"`
	PKCS12CertificateWithRSAPrivateKeyUpload           bool                  `xml:"PKCS12CertificateWithRSAPrivateKeyUpload,attr,omitempty" json:"PKCS12CertificateWithRSAPrivateKeyUpload,attr,omitempty" yaml:"PKCS12CertificateWithRSAPrivateKeyUpload,attr,omitempty"`
	PasswordBasedEncryptionAlgorithms                  string                `xml:"PasswordBasedEncryptionAlgorithms,attr,omitempty" json:"PasswordBasedEncryptionAlgorithms,attr,omitempty" yaml:"PasswordBasedEncryptionAlgorithms,attr,omitempty"`
	PasswordBasedMACAlgorithms                         string                `xml:"PasswordBasedMACAlgorithms,attr,omitempty" json:"PasswordBasedMACAlgorithms,attr,omitempty" yaml:"PasswordBasedMACAlgorithms,attr,omitempty"`
	MaximumNumberOfCRLs                                uint                  `xml:"MaximumNumberOfCRLs,attr,omitempty" json:"MaximumNumberOfCRLs,attr,omitempty" yaml:"MaximumNumberOfCRLs,attr,omitempty"`
	MaximumNumberOfCertificationPathValidationPolicies uint                  `xml:"MaximumNumberOfCertificationPathValidationPolicies,attr,omitempty" json:"MaximumNumberOfCertificationPathValidationPolicies,attr,omitempty" yaml:"MaximumNumberOfCertificationPathValidationPolicies,attr,omitempty"`
	EnforceTLSWebClientAuthExtKeyUsage                 bool                  `xml:"EnforceTLSWebClientAuthExtKeyUsage,attr,omitempty" json:"EnforceTLSWebClientAuthExtKeyUsage,attr,omitempty" yaml:"EnforceTLSWebClientAuthExtKeyUsage,attr,omitempty"`
	NoPrivateKeySharing                                bool                  `xml:"NoPrivateKeySharing,attr,omitempty" json:"NoPrivateKeySharing,attr,omitempty" yaml:"NoPrivateKeySharing,attr,omitempty"`
}

type MultiValuedRDN struct {
	Attribute []DNAttributeTypeAndValue `xml:"Attribute,omitempty" json:"Attribute,omitempty" yaml:"Attribute,omitempty"`
}

type PassphraseAttribute struct {
	PassphraseID PassphraseID `xml:"PassphraseID,omitempty" json:"PassphraseID,omitempty" yaml:"PassphraseID,omitempty"`
	Alias        string       `xml:"Alias,omitempty" json:"Alias,omitempty" yaml:"Alias,omitempty"`
}

type RemoveCertPathValidationPolicyAssignment struct {
	CertPathValidationPolicyID CertPathValidationPolicyID `xml:"CertPathValidationPolicyID,omitempty" json:"CertPathValidationPolicyID,omitempty" yaml:"CertPathValidationPolicyID,omitempty"`
}

type RemoveCertPathValidationPolicyAssignmentResponse struct {
}

type RemoveServerCertificateAssignment struct {
	CertificationPathID CertificationPathID `xml:"CertificationPathID,omitempty" json:"CertificationPathID,omitempty" yaml:"CertificationPathID,omitempty"`
}

type RemoveServerCertificateAssignmentResponse struct {
}

type ReplaceCertPathValidationPolicyAssignment struct {
	OldCertPathValidationPolicyID CertPathValidationPolicyID `xml:"OldCertPathValidationPolicyID,omitempty" json:"OldCertPathValidationPolicyID,omitempty" yaml:"OldCertPathValidationPolicyID,omitempty"`
	NewCertPathValidationPolicyID CertPathValidationPolicyID `xml:"NewCertPathValidationPolicyID,omitempty" json:"NewCertPathValidationPolicyID,omitempty" yaml:"NewCertPathValidationPolicyID,omitempty"`
}

type ReplaceCertPathValidationPolicyAssignmentResponse struct {
}

type ReplaceServerCertificateAssignment struct {
	OldCertificationPathID CertificationPathID `xml:"OldCertificationPathID,omitempty" json:"OldCertificationPathID,omitempty" yaml:"OldCertificationPathID,omitempty"`
	NewCertificationPathID CertificationPathID `xml:"NewCertificationPathID,omitempty" json:"NewCertificationPathID,omitempty" yaml:"NewCertificationPathID,omitempty"`
}

type ReplaceServerCertificateAssignmentResponse struct {
}

type SetClientAuthenticationRequired struct {
	ClientAuthenticationRequired bool `xml:"clientAuthenticationRequired,omitempty" json:"clientAuthenticationRequired,omitempty" yaml:"clientAuthenticationRequired,omitempty"`
}

type SetClientAuthenticationRequiredResponse struct {
}

type SetCnMapsToUser struct {
	CnMapsToUser bool `xml:"cnMapsToUser,omitempty" json:"cnMapsToUser,omitempty" yaml:"cnMapsToUser,omitempty"`
}

type SetCnMapsToUserResponse struct {
}

type SetEnabledTLSVersions struct {
	Versions TLSVersions `xml:"Versions,omitempty" json:"Versions,omitempty" yaml:"Versions,omitempty"`
}

type SetEnabledTLSVersionsResponse struct {
}

type SetNetworkInterfaceDot1XConfiguration struct {
	Token   string  `xml:"token,omitempty" json:"token,omitempty" yaml:"token,omitempty"`
	Dot1XID Dot1XID `xml:"Dot1XID,omitempty" json:"Dot1XID,omitempty" yaml:"Dot1XID,omitempty"`
}

type SetNetworkInterfaceDot1XConfigurationResponse struct {
	RebootNeeded bool `xml:"RebootNeeded,omitempty" json:"RebootNeeded,omitempty" yaml:"RebootNeeded,omitempty"`
}

type TLSServerCapabilities []interface{}

type TrustAnchor struct {
	CertificateID CertificateID `xml:"CertificateID,omitempty" json:"CertificateID,omitempty" yaml:"CertificateID,omitempty"`
}

type UploadCRL struct {
	Crl           Base64DERencodedASN1Value `xml:"Crl,omitempty" json:"Crl,omitempty" yaml:"Crl,omitempty"`
	Alias         string                    `xml:"Alias,omitempty" json:"Alias,omitempty" yaml:"Alias,omitempty"`
	AnyParameters []string                  `xml:"anyParameters>anyParameters,omitempty" json:"anyParameters>anyParameters,omitempty" yaml:"anyParameters>anyParameters,omitempty"`
}

type UploadCRLResponse struct {
	CrlID CRLID `xml:"CrlID,omitempty" json:"CrlID,omitempty" yaml:"CrlID,omitempty"`
}

type UploadCertificate struct {
	Certificate        Base64DERencodedASN1Value `xml:"Certificate,omitempty" json:"Certificate,omitempty" yaml:"Certificate,omitempty"`
	Alias              string                    `xml:"Alias,omitempty" json:"Alias,omitempty" yaml:"Alias,omitempty"`
	KeyAlias           string                    `xml:"KeyAlias,omitempty" json:"KeyAlias,omitempty" yaml:"KeyAlias,omitempty"`
	PrivateKeyRequired bool                      `xml:"PrivateKeyRequired,omitempty" json:"PrivateKeyRequired,omitempty" yaml:"PrivateKeyRequired,omitempty"`
}

type UploadCertificateResponse struct {
	CertificateID CertificateID `xml:"CertificateID,omitempty" json:"CertificateID,omitempty" yaml:"CertificateID,omitempty"`
	KeyID         KeyID         `xml:"KeyID,omitempty" json:"KeyID,omitempty" yaml:"KeyID,omitempty"`
}

type UploadCertificateWithPrivateKeyInPKCS12 struct {
	CertWithPrivateKey           Base64DERencodedASN1Value `xml:"CertWithPrivateKey,omitempty" json:"CertWithPrivateKey,omitempty" yaml:"CertWithPrivateKey,omitempty"`
	CertificationPathAlias       string                    `xml:"CertificationPathAlias,omitempty" json:"CertificationPathAlias,omitempty" yaml:"CertificationPathAlias,omitempty"`
	KeyAlias                     string                    `xml:"KeyAlias,omitempty" json:"KeyAlias,omitempty" yaml:"KeyAlias,omitempty"`
	IgnoreAdditionalCertificates bool                      `xml:"IgnoreAdditionalCertificates,omitempty" json:"IgnoreAdditionalCertificates,omitempty" yaml:"IgnoreAdditionalCertificates,omitempty"`
	IntegrityPassphraseID        PassphraseID              `xml:"IntegrityPassphraseID,omitempty" json:"IntegrityPassphraseID,omitempty" yaml:"IntegrityPassphraseID,omitempty"`
	EncryptionPassphraseID       PassphraseID              `xml:"EncryptionPassphraseID,omitempty" json:"EncryptionPassphraseID,omitempty" yaml:"EncryptionPassphraseID,omitempty"`
	Passphrase                   string                    `xml:"Passphrase,omitempty" json:"Passphrase,omitempty" yaml:"Passphrase,omitempty"`
}

type UploadCertificateWithPrivateKeyInPKCS12Response struct {
	CertificationPathID CertificationPathID `xml:"CertificationPathID,omitempty" json:"CertificationPathID,omitempty" yaml:"CertificationPathID,omitempty"`
	KeyID               KeyID               `xml:"KeyID,omitempty" json:"KeyID,omitempty" yaml:"KeyID,omitempty"`
}

type UploadKeyPairInPKCS8 struct {
	KeyPair                Base64DERencodedASN1Value `xml:"KeyPair,omitempty" json:"KeyPair,omitempty" yaml:"KeyPair,omitempty"`
	Alias                  string                    `xml:"Alias,omitempty" json:"Alias,omitempty" yaml:"Alias,omitempty"`
	EncryptionPassphraseID PassphraseID              `xml:"EncryptionPassphraseID,omitempty" json:"EncryptionPassphraseID,omitempty" yaml:"EncryptionPassphraseID,omitempty"`
	EncryptionPassphrase   string                    `xml:"EncryptionPassphrase,omitempty" json:"EncryptionPassphrase,omitempty" yaml:"EncryptionPassphrase,omitempty"`
}

type UploadKeyPairInPKCS8Response struct {
	KeyID KeyID `xml:"KeyID,omitempty" json:"KeyID,omitempty" yaml:"KeyID,omitempty"`
}

type UploadPassphrase struct {
	Passphrase      string `xml:"Passphrase,omitempty" json:"Passphrase,omitempty" yaml:"Passphrase,omitempty"`
	PassphraseAlias string `xml:"PassphraseAlias,omitempty" json:"PassphraseAlias,omitempty" yaml:"PassphraseAlias,omitempty"`
}

type UploadPassphraseResponse struct {
	PassphraseID PassphraseID `xml:"PassphraseID,omitempty" json:"PassphraseID,omitempty" yaml:"PassphraseID,omitempty"`
}

type X509Certificate struct {
	CertificateID      CertificateID             `xml:"CertificateID,omitempty" json:"CertificateID,omitempty" yaml:"CertificateID,omitempty"`
	KeyID              KeyID                     `xml:"KeyID,omitempty" json:"KeyID,omitempty" yaml:"KeyID,omitempty"`
	Alias              string                    `xml:"Alias,omitempty" json:"Alias,omitempty" yaml:"Alias,omitempty"`
	CertificateContent Base64DERencodedASN1Value `xml:"CertificateContent,omitempty" json:"CertificateContent,omitempty" yaml:"CertificateContent,omitempty"`
}

type X509v3Extension struct {
	ExtnOID   DotDecimalOID             `xml:"extnOID,omitempty" json:"extnOID,omitempty" yaml:"extnOID,omitempty"`
	Critical  bool                      `xml:"critical,omitempty" json:"critical,omitempty" yaml:"critical,omitempty"`
	ExtnValue Base64DERencodedASN1Value `xml:"extnValue,omitempty" json:"extnValue,omitempty" yaml:"extnValue,omitempty"`
}

// dot1X implements the Dot1X interface.
type dot1X struct {
	cli      common.Client
	Endpoint string
}

func (p *dot1X) OptAddCertPathValidationPolicyAssignment(args AddCertPathValidationPolicyAssignment) (*AddCertPathValidationPolicyAssignmentResponse, *common.Fault) {
	req := struct {
		XMLName                               string `xml:"tas:AddCertPathValidationPolicyAssignment"`
		AddCertPathValidationPolicyAssignment AddCertPathValidationPolicyAssignment
	}{
		AddCertPathValidationPolicyAssignment: args,
	}

	resp := AddCertPathValidationPolicyAssignmentResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *dot1X) OptAddDot1XConfiguration(args AddDot1XConfiguration) (*AddDot1XConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName               string `xml:"tas:AddDot1XConfiguration"`
		AddDot1XConfiguration AddDot1XConfiguration
	}{
		AddDot1XConfiguration: args,
	}

	resp := AddDot1XConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *dot1X) OptAddServerCertificateAssignment(args AddServerCertificateAssignment) (*AddServerCertificateAssignmentResponse, *common.Fault) {
	req := struct {
		XMLName                        string `xml:"tas:AddServerCertificateAssignment"`
		AddServerCertificateAssignment AddServerCertificateAssignment
	}{
		AddServerCertificateAssignment: args,
	}

	resp := AddServerCertificateAssignmentResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *dot1X) OptCreateCertPathValidationPolicy(args CreateCertPathValidationPolicy) (*CreateCertPathValidationPolicyResponse, *common.Fault) {
	req := struct {
		XMLName                        string `xml:"tas:CreateCertPathValidationPolicy"`
		CreateCertPathValidationPolicy CreateCertPathValidationPolicy
	}{
		CreateCertPathValidationPolicy: args,
	}

	resp := CreateCertPathValidationPolicyResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *dot1X) OptCreateCertificationPath(args CreateCertificationPath) (*CreateCertificationPathResponse, *common.Fault) {
	req := struct {
		XMLName                 string `xml:"tas:CreateCertificationPath"`
		CreateCertificationPath CreateCertificationPath
	}{
		CreateCertificationPath: args,
	}

	resp := CreateCertificationPathResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *dot1X) OptCreatePKCS10CSR(args CreatePKCS10CSR) (*CreatePKCS10CSRResponse, *common.Fault) {
	req := struct {
		XMLName         string `xml:"tas:CreatePKCS10CSR"`
		CreatePKCS10CSR CreatePKCS10CSR
	}{
		CreatePKCS10CSR: args,
	}

	resp := CreatePKCS10CSRResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *dot1X) OptCreateRSAKeyPair(args CreateRSAKeyPair) (*CreateRSAKeyPairResponse, *common.Fault) {
	req := struct {
		XMLName          string `xml:"tas:CreateRSAKeyPair"`
		CreateRSAKeyPair CreateRSAKeyPair
	}{
		CreateRSAKeyPair: args,
	}

	resp := CreateRSAKeyPairResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *dot1X) OptCreateSelfSignedCertificate(args CreateSelfSignedCertificate) (*CreateSelfSignedCertificateResponse, *common.Fault) {
	req := struct {
		XMLName                     string `xml:"tas:CreateSelfSignedCertificate"`
		CreateSelfSignedCertificate CreateSelfSignedCertificate
	}{
		CreateSelfSignedCertificate: args,
	}

	resp := CreateSelfSignedCertificateResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *dot1X) OptDeleteCRL(args DeleteCRL) (*DeleteCRLResponse, *common.Fault) {
	req := struct {
		XMLName   string `xml:"tas:DeleteCRL"`
		DeleteCRL DeleteCRL
	}{
		DeleteCRL: args,
	}

	resp := DeleteCRLResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *dot1X) OptDeleteCertPathValidationPolicy(args DeleteCertPathValidationPolicy) (*DeleteCertPathValidationPolicyResponse, *common.Fault) {
	req := struct {
		XMLName                        string `xml:"tas:DeleteCertPathValidationPolicy"`
		DeleteCertPathValidationPolicy DeleteCertPathValidationPolicy
	}{
		DeleteCertPathValidationPolicy: args,
	}

	resp := DeleteCertPathValidationPolicyResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *dot1X) OptDeleteCertificate(args DeleteCertificate) (*DeleteCertificateResponse, *common.Fault) {
	req := struct {
		XMLName           string `xml:"tas:DeleteCertificate"`
		DeleteCertificate DeleteCertificate
	}{
		DeleteCertificate: args,
	}

	resp := DeleteCertificateResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *dot1X) OptDeleteCertificationPath(args DeleteCertificationPath) (*DeleteCertificationPathResponse, *common.Fault) {
	req := struct {
		XMLName                 string `xml:"tas:DeleteCertificationPath"`
		DeleteCertificationPath DeleteCertificationPath
	}{
		DeleteCertificationPath: args,
	}

	resp := DeleteCertificationPathResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *dot1X) OptDeleteDot1XConfiguration(args DeleteDot1XConfiguration) (*DeleteDot1XConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                  string `xml:"tas:DeleteDot1XConfiguration"`
		DeleteDot1XConfiguration DeleteDot1XConfiguration
	}{
		DeleteDot1XConfiguration: args,
	}

	resp := DeleteDot1XConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *dot1X) OptDeleteKey(args DeleteKey) (*DeleteKeyResponse, *common.Fault) {
	req := struct {
		XMLName   string `xml:"tas:DeleteKey"`
		DeleteKey DeleteKey
	}{
		DeleteKey: args,
	}

	resp := DeleteKeyResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *dot1X) OptDeleteNetworkInterfaceDot1XConfiguration(args DeleteNetworkInterfaceDot1XConfiguration) (*DeleteNetworkInterfaceDot1XConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                                  string `xml:"tas:DeleteNetworkInterfaceDot1XConfiguration"`
		DeleteNetworkInterfaceDot1XConfiguration DeleteNetworkInterfaceDot1XConfiguration
	}{
		DeleteNetworkInterfaceDot1XConfiguration: args,
	}

	resp := DeleteNetworkInterfaceDot1XConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *dot1X) OptDeletePassphrase(args DeletePassphrase) (*DeletePassphraseResponse, *common.Fault) {
	req := struct {
		XMLName          string `xml:"tas:DeletePassphrase"`
		DeletePassphrase DeletePassphrase
	}{
		DeletePassphrase: args,
	}

	resp := DeletePassphraseResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *dot1X) OptGetAllCRLs(args GetAllCRLs) (*GetAllCRLsResponse, *common.Fault) {
	req := struct {
		XMLName    string `xml:"tas:GetAllCRLs"`
		GetAllCRLs GetAllCRLs
	}{
		GetAllCRLs: args,
	}

	resp := GetAllCRLsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *dot1X) OptGetAllCertPathValidationPolicies(args GetAllCertPathValidationPolicies) (*GetAllCertPathValidationPoliciesResponse, *common.Fault) {
	req := struct {
		XMLName                          string `xml:"tas:GetAllCertPathValidationPolicies"`
		GetAllCertPathValidationPolicies GetAllCertPathValidationPolicies
	}{
		GetAllCertPathValidationPolicies: args,
	}

	resp := GetAllCertPathValidationPoliciesResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *dot1X) OptGetAllCertificates(args GetAllCertificates) (*GetAllCertificatesResponse, *common.Fault) {
	req := struct {
		XMLName            string `xml:"tas:GetAllCertificates"`
		GetAllCertificates GetAllCertificates
	}{
		GetAllCertificates: args,
	}

	resp := GetAllCertificatesResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *dot1X) OptGetAllCertificationPaths(args GetAllCertificationPaths) (*GetAllCertificationPathsResponse, *common.Fault) {
	req := struct {
		XMLName                  string `xml:"tas:GetAllCertificationPaths"`
		GetAllCertificationPaths GetAllCertificationPaths
	}{
		GetAllCertificationPaths: args,
	}

	resp := GetAllCertificationPathsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *dot1X) OptGetAllDot1XConfigurations(args GetAllDot1XConfigurations) (*GetAllDot1XConfigurationsResponse, *common.Fault) {
	req := struct {
		XMLName                   string `xml:"tas:GetAllDot1XConfigurations"`
		GetAllDot1XConfigurations GetAllDot1XConfigurations
	}{
		GetAllDot1XConfigurations: args,
	}

	resp := GetAllDot1XConfigurationsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *dot1X) OptGetAllKeys(args GetAllKeys) (*GetAllKeysResponse, *common.Fault) {
	req := struct {
		XMLName    string `xml:"tas:GetAllKeys"`
		GetAllKeys GetAllKeys
	}{
		GetAllKeys: args,
	}

	resp := GetAllKeysResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *dot1X) OptGetAllPassphrases(args GetAllPassphrases) (*GetAllPassphrasesResponse, *common.Fault) {
	req := struct {
		XMLName           string `xml:"tas:GetAllPassphrases"`
		GetAllPassphrases GetAllPassphrases
	}{
		GetAllPassphrases: args,
	}

	resp := GetAllPassphrasesResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *dot1X) OptGetAssignedCertPathValidationPolicies(args GetAssignedCertPathValidationPolicies) (*GetAssignedCertPathValidationPoliciesResponse, *common.Fault) {
	req := struct {
		XMLName                               string `xml:"tas:GetAssignedCertPathValidationPolicies"`
		GetAssignedCertPathValidationPolicies GetAssignedCertPathValidationPolicies
	}{
		GetAssignedCertPathValidationPolicies: args,
	}

	resp := GetAssignedCertPathValidationPoliciesResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *dot1X) OptGetAssignedServerCertificates(args GetAssignedServerCertificates) (*GetAssignedServerCertificatesResponse, *common.Fault) {
	req := struct {
		XMLName                       string `xml:"tas:GetAssignedServerCertificates"`
		GetAssignedServerCertificates GetAssignedServerCertificates
	}{
		GetAssignedServerCertificates: args,
	}

	resp := GetAssignedServerCertificatesResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *dot1X) OptGetCRL(args GetCRL) (*GetCRLResponse, *common.Fault) {
	req := struct {
		XMLName string `xml:"tas:GetCRL"`
		GetCRL  GetCRL
	}{
		GetCRL: args,
	}

	resp := GetCRLResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *dot1X) OptGetCertPathValidationPolicy(args GetCertPathValidationPolicy) (*GetCertPathValidationPolicyResponse, *common.Fault) {
	req := struct {
		XMLName                     string `xml:"tas:GetCertPathValidationPolicy"`
		GetCertPathValidationPolicy GetCertPathValidationPolicy
	}{
		GetCertPathValidationPolicy: args,
	}

	resp := GetCertPathValidationPolicyResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *dot1X) OptGetCertificate(args GetCertificate) (*GetCertificateResponse, *common.Fault) {
	req := struct {
		XMLName        string `xml:"tas:GetCertificate"`
		GetCertificate GetCertificate
	}{
		GetCertificate: args,
	}

	resp := GetCertificateResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *dot1X) OptGetCertificationPath(args GetCertificationPath) (*GetCertificationPathResponse, *common.Fault) {
	req := struct {
		XMLName              string `xml:"tas:GetCertificationPath"`
		GetCertificationPath GetCertificationPath
	}{
		GetCertificationPath: args,
	}

	resp := GetCertificationPathResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *dot1X) OptGetClientAuthenticationRequired(args GetClientAuthenticationRequired) (*GetClientAuthenticationRequiredResponse, *common.Fault) {
	req := struct {
		XMLName                         string `xml:"tas:GetClientAuthenticationRequired"`
		GetClientAuthenticationRequired GetClientAuthenticationRequired
	}{
		GetClientAuthenticationRequired: args,
	}

	resp := GetClientAuthenticationRequiredResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *dot1X) OptGetCnMapsToUser(args GetCnMapsToUser) (*GetCnMapsToUserResponse, *common.Fault) {
	req := struct {
		XMLName         string `xml:"tas:GetCnMapsToUser"`
		GetCnMapsToUser GetCnMapsToUser
	}{
		GetCnMapsToUser: args,
	}

	resp := GetCnMapsToUserResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *dot1X) OptGetDot1XConfiguration(args GetDot1XConfiguration) (*GetDot1XConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName               string `xml:"tas:GetDot1XConfiguration"`
		GetDot1XConfiguration GetDot1XConfiguration
	}{
		GetDot1XConfiguration: args,
	}

	resp := GetDot1XConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *dot1X) OptGetEnabledTLSVersions(args GetEnabledTLSVersions) (*GetEnabledTLSVersionsResponse, *common.Fault) {
	req := struct {
		XMLName               string `xml:"tas:GetEnabledTLSVersions"`
		GetEnabledTLSVersions GetEnabledTLSVersions
	}{
		GetEnabledTLSVersions: args,
	}

	resp := GetEnabledTLSVersionsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *dot1X) OptGetKeyStatus(args GetKeyStatus) (*GetKeyStatusResponse, *common.Fault) {
	req := struct {
		XMLName      string `xml:"tas:GetKeyStatus"`
		GetKeyStatus GetKeyStatus
	}{
		GetKeyStatus: args,
	}

	resp := GetKeyStatusResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *dot1X) OptGetNetworkInterfaceDot1XConfiguration(args GetNetworkInterfaceDot1XConfiguration) (*GetNetworkInterfaceDot1XConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                               string `xml:"tas:GetNetworkInterfaceDot1XConfiguration"`
		GetNetworkInterfaceDot1XConfiguration GetNetworkInterfaceDot1XConfiguration
	}{
		GetNetworkInterfaceDot1XConfiguration: args,
	}

	resp := GetNetworkInterfaceDot1XConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *dot1X) OptGetPrivateKeyStatus(args GetPrivateKeyStatus) (*GetPrivateKeyStatusResponse, *common.Fault) {
	req := struct {
		XMLName             string `xml:"tas:GetPrivateKeyStatus"`
		GetPrivateKeyStatus GetPrivateKeyStatus
	}{
		GetPrivateKeyStatus: args,
	}

	resp := GetPrivateKeyStatusResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *dot1X) OptGetServiceCapabilities(args GetServiceCapabilities) (*GetServiceCapabilitiesResponse, *common.Fault) {
	req := struct {
		XMLName                string `xml:"tas:GetServiceCapabilities"`
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

func (p *dot1X) OptRemoveCertPathValidationPolicyAssignment(args RemoveCertPathValidationPolicyAssignment) (*RemoveCertPathValidationPolicyAssignmentResponse, *common.Fault) {
	req := struct {
		XMLName                                  string `xml:"tas:RemoveCertPathValidationPolicyAssignment"`
		RemoveCertPathValidationPolicyAssignment RemoveCertPathValidationPolicyAssignment
	}{
		RemoveCertPathValidationPolicyAssignment: args,
	}

	resp := RemoveCertPathValidationPolicyAssignmentResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *dot1X) OptRemoveServerCertificateAssignment(args RemoveServerCertificateAssignment) (*RemoveServerCertificateAssignmentResponse, *common.Fault) {
	req := struct {
		XMLName                           string `xml:"tas:RemoveServerCertificateAssignment"`
		RemoveServerCertificateAssignment RemoveServerCertificateAssignment
	}{
		RemoveServerCertificateAssignment: args,
	}

	resp := RemoveServerCertificateAssignmentResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *dot1X) OptReplaceCertPathValidationPolicyAssignment(args ReplaceCertPathValidationPolicyAssignment) (*ReplaceCertPathValidationPolicyAssignmentResponse, *common.Fault) {
	req := struct {
		XMLName                                   string `xml:"tas:ReplaceCertPathValidationPolicyAssignment"`
		ReplaceCertPathValidationPolicyAssignment ReplaceCertPathValidationPolicyAssignment
	}{
		ReplaceCertPathValidationPolicyAssignment: args,
	}

	resp := ReplaceCertPathValidationPolicyAssignmentResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *dot1X) OptReplaceServerCertificateAssignment(args ReplaceServerCertificateAssignment) (*ReplaceServerCertificateAssignmentResponse, *common.Fault) {
	req := struct {
		XMLName                            string `xml:"tas:ReplaceServerCertificateAssignment"`
		ReplaceServerCertificateAssignment ReplaceServerCertificateAssignment
	}{
		ReplaceServerCertificateAssignment: args,
	}

	resp := ReplaceServerCertificateAssignmentResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *dot1X) OptSetClientAuthenticationRequired(args SetClientAuthenticationRequired) (*SetClientAuthenticationRequiredResponse, *common.Fault) {
	req := struct {
		XMLName                         string `xml:"tas:SetClientAuthenticationRequired"`
		SetClientAuthenticationRequired SetClientAuthenticationRequired
	}{
		SetClientAuthenticationRequired: args,
	}

	resp := SetClientAuthenticationRequiredResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *dot1X) OptSetCnMapsToUser(args SetCnMapsToUser) (*SetCnMapsToUserResponse, *common.Fault) {
	req := struct {
		XMLName         string `xml:"tas:SetCnMapsToUser"`
		SetCnMapsToUser SetCnMapsToUser
	}{
		SetCnMapsToUser: args,
	}

	resp := SetCnMapsToUserResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *dot1X) OptSetEnabledTLSVersions(args SetEnabledTLSVersions) (*SetEnabledTLSVersionsResponse, *common.Fault) {
	req := struct {
		XMLName               string `xml:"tas:SetEnabledTLSVersions"`
		SetEnabledTLSVersions SetEnabledTLSVersions
	}{
		SetEnabledTLSVersions: args,
	}

	resp := SetEnabledTLSVersionsResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *dot1X) OptSetNetworkInterfaceDot1XConfiguration(args SetNetworkInterfaceDot1XConfiguration) (*SetNetworkInterfaceDot1XConfigurationResponse, *common.Fault) {
	req := struct {
		XMLName                               string `xml:"tas:SetNetworkInterfaceDot1XConfiguration"`
		SetNetworkInterfaceDot1XConfiguration SetNetworkInterfaceDot1XConfiguration
	}{
		SetNetworkInterfaceDot1XConfiguration: args,
	}

	resp := SetNetworkInterfaceDot1XConfigurationResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *dot1X) OptUploadCRL(args UploadCRL) (*UploadCRLResponse, *common.Fault) {
	req := struct {
		XMLName   string `xml:"tas:UploadCRL"`
		UploadCRL UploadCRL
	}{
		UploadCRL: args,
	}

	resp := UploadCRLResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *dot1X) OptUploadCertificate(args UploadCertificate) (*UploadCertificateResponse, *common.Fault) {
	req := struct {
		XMLName           string `xml:"tas:UploadCertificate"`
		UploadCertificate UploadCertificate
	}{
		UploadCertificate: args,
	}

	resp := UploadCertificateResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *dot1X) OptUploadCertificateWithPrivateKeyInPKCS12(args UploadCertificateWithPrivateKeyInPKCS12) (*UploadCertificateWithPrivateKeyInPKCS12Response, *common.Fault) {
	req := struct {
		XMLName                                 string `xml:"tas:UploadCertificateWithPrivateKeyInPKCS12"`
		UploadCertificateWithPrivateKeyInPKCS12 UploadCertificateWithPrivateKeyInPKCS12
	}{
		UploadCertificateWithPrivateKeyInPKCS12: args,
	}

	resp := UploadCertificateWithPrivateKeyInPKCS12Response{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *dot1X) OptUploadKeyPairInPKCS8(args UploadKeyPairInPKCS8) (*UploadKeyPairInPKCS8Response, *common.Fault) {
	req := struct {
		XMLName              string `xml:"tas:UploadKeyPairInPKCS8"`
		UploadKeyPairInPKCS8 UploadKeyPairInPKCS8
	}{
		UploadKeyPairInPKCS8: args,
	}

	resp := UploadKeyPairInPKCS8Response{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}

func (p *dot1X) OptUploadPassphrase(args UploadPassphrase) (*UploadPassphraseResponse, *common.Fault) {
	req := struct {
		XMLName          string `xml:"tas:UploadPassphrase"`
		UploadPassphrase UploadPassphrase
	}{
		UploadPassphrase: args,
	}

	resp := UploadPassphraseResponse{}

	if fault := p.cli.CallMethodUnmarshal(p.Endpoint, req, &resp); fault != nil {
		return nil, fault
	}
	return &resp, nil
}
