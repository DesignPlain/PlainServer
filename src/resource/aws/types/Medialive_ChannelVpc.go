package types

type Medialive_ChannelVpc struct {
	//
	AvailabilityZones []string `json:"availabilityZones,omitempty" yaml:"availabilityZones,omitempty"`

	//
	NetworkInterfaceIds []string `json:"networkInterfaceIds,omitempty" yaml:"networkInterfaceIds,omitempty"`

	// List of public address allocation ids to associate with ENIs that will be created in Output VPC. Must specify one for SINGLE_PIPELINE, two for STANDARD channels.
	PublicAddressAllocationIds []string `json:"publicAddressAllocationIds,omitempty" yaml:"publicAddressAllocationIds,omitempty"`

	// A list of up to 5 EC2 VPC security group IDs to attach to the Output VPC network interfaces. If none are specified then the VPC default security group will be used.
	SecurityGroupIds []string `json:"securityGroupIds,omitempty" yaml:"securityGroupIds,omitempty"`

	// A list of VPC subnet IDs from the same VPC. If STANDARD channel, subnet IDs must be mapped to two unique availability zones (AZ).
	SubnetIds []string `json:"subnetIds,omitempty" yaml:"subnetIds,omitempty"`
}
