package types

type Batch_getJobDefinitionNodePropertyNodeRangePropertyContainerUlimit struct {
	// The hard limit for the ulimit type.
	HardLimit int `json:"hardLimit,omitempty" yaml:"hardLimit,omitempty"`

	// The name of the job definition to register. It can be up to 128 letters long. It can contain uppercase and lowercase letters, numbers, hyphens (-), and underscores (_).
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The soft limit for the ulimit type.
	SoftLimit int `json:"softLimit,omitempty" yaml:"softLimit,omitempty"`
}
