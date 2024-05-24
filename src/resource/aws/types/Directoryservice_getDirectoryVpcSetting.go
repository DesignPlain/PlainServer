package types

type Directoryservice_getDirectoryVpcSetting struct {
	//
	AvailabilityZones []string `json:"availabilityZones,omitempty" yaml:"availabilityZones,omitempty"`

	// Identifiers of the subnets for the connector servers (2 subnets in 2 different AZs).
	SubnetIds []string `json:"subnetIds,omitempty" yaml:"subnetIds,omitempty"`

	// ID of the VPC that the connector is in.
	VpcId string `json:"vpcId,omitempty" yaml:"vpcId,omitempty"`
}
