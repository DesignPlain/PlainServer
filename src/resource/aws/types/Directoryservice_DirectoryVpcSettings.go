package types

type Directoryservice_DirectoryVpcSettings struct {
	//
	AvailabilityZones []string `json:"availabilityZones,omitempty" yaml:"availabilityZones,omitempty"`

	// The identifiers of the subnets for the directory servers (2 subnets in 2 different AZs).
	SubnetIds []string `json:"subnetIds,omitempty" yaml:"subnetIds,omitempty"`

	// The identifier of the VPC that the directory is in.
	VpcId string `json:"vpcId,omitempty" yaml:"vpcId,omitempty"`
}
