package types

type S3_BucketReplicationConfigRuleFilterTag struct {
	// Value of the tag.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`

	// Name of the object key.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`
}
