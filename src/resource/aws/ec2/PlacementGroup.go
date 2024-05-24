package ec2

type PlacementGroup struct {
	// The name of the placement group.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The number of partitions to create in the
	   placement group.  Can only be specified when the `strategy` is set to
	   `partition`.  Valid values are 1 - 7 (default is `2`).
	*/
	PartitionCount int `json:"partitionCount,omitempty" yaml:"partitionCount,omitempty"`

	/*
	   Determines how placement groups spread instances. Can only be used
	   when the `strategy` is set to `spread`. Can be `host` or `rack`. `host` can only be used for Outpost placement groups. Defaults to `rack`.
	*/
	SpreadLevel string `json:"spreadLevel,omitempty" yaml:"spreadLevel,omitempty"`

	// The placement strategy. Can be `cluster`, `partition` or `spread`.
	Strategy string `json:"strategy,omitempty" yaml:"strategy,omitempty"`

	// Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
