package sesv2

type DedicatedIpPool struct {
	// A map of tags to assign to the pool. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	/*
	   Name of the dedicated IP pool.

	   The following arguments are optional:
	*/
	PoolName string `json:"poolName,omitempty" yaml:"poolName,omitempty"`

	// IP pool scaling mode. Valid values: `STANDARD`, `MANAGED`. If omitted, the AWS API will default to a standard pool.
	ScalingMode string `json:"scalingMode,omitempty" yaml:"scalingMode,omitempty"`
}
