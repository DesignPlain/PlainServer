package elasticache

import types "libds/aws/types"

type ParameterGroup struct {
	// The description of the ElastiCache parameter group. Defaults to "Managed by Pulumi".
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The family of the ElastiCache parameter group.
	Family string `json:"family,omitempty" yaml:"family,omitempty"`

	// The name of the ElastiCache parameter.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// A list of ElastiCache parameters to apply.
	Parameters []types.Elasticache_ParameterGroupParameter `json:"parameters,omitempty" yaml:"parameters,omitempty"`

	// Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
