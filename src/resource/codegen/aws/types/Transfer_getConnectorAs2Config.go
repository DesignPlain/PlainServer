package types

type Transfer_getConnectorAs2Config struct {
	// Subject HTTP header attribute in outbound AS2 messages to the connector.
	MessageSubject string `json:"messageSubject,omitempty" yaml:"messageSubject,omitempty"`

	//
	SingingAlgorithm string `json:"singingAlgorithm,omitempty" yaml:"singingAlgorithm,omitempty"`

	// Specifies whether AS2 file is compressed. Will be ZLIB or DISABLED
	Compression string `json:"compression,omitempty" yaml:"compression,omitempty"`

	// Algorithm used to encrypt file. Will be AES128_CBC or AES192_CBC or AES256_CBC or DES_EDE3_CBC or NONE.
	EncryptionAlgorithm string `json:"encryptionAlgorithm,omitempty" yaml:"encryptionAlgorithm,omitempty"`

	// Unique identifier for AS2 local profile.
	LocalProfileId string `json:"localProfileId,omitempty" yaml:"localProfileId,omitempty"`

	// Signing algorithm for MDN response. Will be SHA256 or SHA384 or SHA512 or SHA1 or NONE or DEFAULT.
	MdnSigningAlgorithm string `json:"mdnSigningAlgorithm,omitempty" yaml:"mdnSigningAlgorithm,omitempty"`

	// Basic authentication for AS2 connector API. Returns a null value if not set.
	BasicAuthSecretId string `json:"basicAuthSecretId,omitempty" yaml:"basicAuthSecretId,omitempty"`

	// Used for outbound requests to tell if response is asynchronous or not. Will be either SYNC or NONE.
	MdnResponse string `json:"mdnResponse,omitempty" yaml:"mdnResponse,omitempty"`

	// Unique identifier used by connector for partner profile.
	PartnerProfileId string `json:"partnerProfileId,omitempty" yaml:"partnerProfileId,omitempty"`
}
