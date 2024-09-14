package evidently

type Segment struct {
	// A name for the segment.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The pattern to use for the segment. For more information about pattern syntax, see [Segment rule pattern syntax](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Evidently-segments.html#CloudWatch-Evidently-segments-syntax.html).
	Pattern string `json:"pattern,omitempty" yaml:"pattern,omitempty"`

	// Tags to apply to the segment. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Specifies the description of the segment.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
