package types

type S3_BucketObjectv2OverrideProvider struct {
	// Override the provider `default_tags` configuration block.
	DefaultTags S3_BucketObjectv2OverrideProviderDefaultTags `json:"defaultTags,omitempty" yaml:"defaultTags,omitempty"`
}
