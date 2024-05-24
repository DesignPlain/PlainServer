package types

type Appstream_FleetVpcConfig struct {
	// Identifiers of the security groups for the fleet or image builder.
	SecurityGroupIds []string `json:"securityGroupIds,omitempty" yaml:"securityGroupIds,omitempty"`

	// Identifiers of the subnets to which a network interface is attached from the fleet instance or image builder instance.
	SubnetIds []string `json:"subnetIds,omitempty" yaml:"subnetIds,omitempty"`
}
