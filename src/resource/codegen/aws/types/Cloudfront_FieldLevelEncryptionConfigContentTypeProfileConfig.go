package types

type Cloudfront_FieldLevelEncryptionConfigContentTypeProfileConfig struct {
	// Object that contains an attribute `items` that contains the list of configurations for a field-level encryption content type-profile. See Content Type Profile.
	ContentTypeProfiles Cloudfront_FieldLevelEncryptionConfigContentTypeProfileConfigContentTypeProfiles `json:"contentTypeProfiles,omitempty" yaml:"contentTypeProfiles,omitempty"`

	// specifies what to do when an unknown content type is provided for the profile. If true, content is forwarded without being encrypted when the content type is unknown. If false (the default), an error is returned when the content type is unknown.
	ForwardWhenContentTypeIsUnknown bool `json:"forwardWhenContentTypeIsUnknown,omitempty" yaml:"forwardWhenContentTypeIsUnknown,omitempty"`
}
