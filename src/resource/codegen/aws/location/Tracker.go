package location

type Tracker struct {
	// The optional description for the tracker resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// A key identifier for an AWS KMS customer managed key assigned to the Amazon Location resource.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	// The position filtering method of the tracker resource. Valid values: `TimeBased`, `DistanceBased`, `AccuracyBased`. Default: `TimeBased`.
	PositionFiltering string `json:"positionFiltering,omitempty" yaml:"positionFiltering,omitempty"`

	// Key-value tags for the tracker. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	/*
	   The name of the tracker resource.

	   The following arguments are optional:
	*/
	TrackerName string `json:"trackerName,omitempty" yaml:"trackerName,omitempty"`
}
