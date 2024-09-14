package elasticache

type SubnetGroup struct {
	// Name for the cache subnet group. ElastiCache converts this name to lowercase.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// List of VPC Subnet IDs for the cache subnet group
	SubnetIds []string `json:"subnetIds,omitempty" yaml:"subnetIds,omitempty"`

	// Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Description for the cache subnet group. Defaults to "Managed by Pulumi".
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
