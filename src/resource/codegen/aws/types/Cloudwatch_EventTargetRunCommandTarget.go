package types

type Cloudwatch_EventTargetRunCommandTarget struct {
	// Can be either `tag:tag-key` or `InstanceIds`.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// If Key is `tag:tag-key`, Values is a list of tag values. If Key is `InstanceIds`, Values is a list of Amazon EC2 instance IDs.
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`
}
