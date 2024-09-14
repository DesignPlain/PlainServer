package types

type Cloudfront_FieldLevelEncryptionProfileEncryptionEntitiesItem struct {
	// The provider associated with the public key being used for encryption.
	ProviderId string `json:"providerId,omitempty" yaml:"providerId,omitempty"`

	// The public key associated with a set of field-level encryption patterns, to be used when encrypting the fields that match the patterns.
	PublicKeyId string `json:"publicKeyId,omitempty" yaml:"publicKeyId,omitempty"`

	// Object that contains an attribute `items` that contains the list of field patterns in a field-level encryption content type profile specify the fields that you want to be encrypted.
	FieldPatterns Cloudfront_FieldLevelEncryptionProfileEncryptionEntitiesItemFieldPatterns `json:"fieldPatterns,omitempty" yaml:"fieldPatterns,omitempty"`
}
