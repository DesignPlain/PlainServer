package apigatewayv2

type VpcLink struct {
	// Name of the VPC Link. Must be between 1 and 128 characters in length.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Security group IDs for the VPC Link.
	SecurityGroupIds []string `json:"securityGroupIds,omitempty" yaml:"securityGroupIds,omitempty"`

	// Subnet IDs for the VPC Link.
	SubnetIds []string `json:"subnetIds,omitempty" yaml:"subnetIds,omitempty"`

	// Map of tags to assign to the VPC Link. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
