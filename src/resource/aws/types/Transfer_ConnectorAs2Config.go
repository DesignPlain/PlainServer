package types

type Transfer_ConnectorAs2Config struct {
	// The signing algorithm for the Mdn response. The valid values are SHA256 | SHA384 | SHA512 | SHA1 | NONE | DEFAULT.
	MdnSigningAlgorithm string `json:"mdnSigningAlgorithm,omitempty" yaml:"mdnSigningAlgorithm,omitempty"`

	// Used as the subject HTTP header attribute in AS2 messages that are being sent with the connector.
	MessageSubject string `json:"messageSubject,omitempty" yaml:"messageSubject,omitempty"`

	// The unique identifier for the AS2 partner profile.
	PartnerProfileId string `json:"partnerProfileId,omitempty" yaml:"partnerProfileId,omitempty"`

	// The algorithm that is used to sign AS2 messages sent with the connector. The valid values are SHA256 | SHA384 | SHA512 | SHA1 | NONE .
	SigningAlgorithm string `json:"signingAlgorithm,omitempty" yaml:"signingAlgorithm,omitempty"`

	// Specifies weather AS2 file is compressed. The valud values are ZLIB and  DISABLED.
	Compression string `json:"compression,omitempty" yaml:"compression,omitempty"`

	// The algorithm that is used to encrypt the file. The valid values are AES128_CBC | AES192_CBC | AES256_CBC | NONE.
	EncryptionAlgorithm string `json:"encryptionAlgorithm,omitempty" yaml:"encryptionAlgorithm,omitempty"`

	// The unique identifier for the AS2 local profile.
	LocalProfileId string `json:"localProfileId,omitempty" yaml:"localProfileId,omitempty"`

	// Used for outbound requests to determine if a partner response for transfers is synchronous or asynchronous. The valid values are SYNC and NONE.
	MdnResponse string `json:"mdnResponse,omitempty" yaml:"mdnResponse,omitempty"`
}
