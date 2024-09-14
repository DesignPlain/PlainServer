package types

type Appstream_ImageBuilderVpcConfig struct {
	// Identifiers of the security groups for the image builder or image builder.
	SecurityGroupIds []string `json:"securityGroupIds,omitempty" yaml:"securityGroupIds,omitempty"`

	// Identifier of the subnet to which a network interface is attached from the image builder instance.
	SubnetIds []string `json:"subnetIds,omitempty" yaml:"subnetIds,omitempty"`
}
