package types

type Cloudfront_FieldLevelEncryptionConfigQueryArgProfileConfig struct {
	// Flag to set if you want a request to be forwarded to the origin even if the profile specified by the field-level encryption query argument, fle-profile, is unknown.
	ForwardWhenQueryArgProfileIsUnknown bool `json:"forwardWhenQueryArgProfileIsUnknown,omitempty" yaml:"forwardWhenQueryArgProfileIsUnknown,omitempty"`

	// Object that contains an attribute `items` that contains the list ofrofiles specified for query argument-profile mapping for field-level encryption. see Query Arg Profile.
	QueryArgProfiles Cloudfront_FieldLevelEncryptionConfigQueryArgProfileConfigQueryArgProfiles `json:"queryArgProfiles,omitempty" yaml:"queryArgProfiles,omitempty"`
}
