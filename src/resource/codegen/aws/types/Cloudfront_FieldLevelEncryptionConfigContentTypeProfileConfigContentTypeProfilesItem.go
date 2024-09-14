package types

type Cloudfront_FieldLevelEncryptionConfigContentTypeProfileConfigContentTypeProfilesItem struct {
	// The format for a field-level encryption content type-profile mapping. Valid value is `URLEncoded`.
	Format string `json:"format,omitempty" yaml:"format,omitempty"`

	//
	ProfileId string `json:"profileId,omitempty" yaml:"profileId,omitempty"`

	// he content type for a field-level encryption content type-profile mapping. Valid value is `application/x-www-form-urlencoded`.
	ContentType string `json:"contentType,omitempty" yaml:"contentType,omitempty"`
}
