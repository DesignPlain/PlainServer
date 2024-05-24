package types

type Elasticache_ParameterGroupParameter struct {
	// The value of the ElastiCache parameter.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`

	// The name of the ElastiCache parameter.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
