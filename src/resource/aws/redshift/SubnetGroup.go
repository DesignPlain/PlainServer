package redshift

type SubnetGroup struct {
	// The description of the Redshift Subnet group. Defaults to "Managed by Pulumi".
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The name of the Redshift Subnet group.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// An array of VPC subnet IDs.
	SubnetIds []string `json:"subnetIds,omitempty" yaml:"subnetIds,omitempty"`

	// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
