package types

type Directoryservice_ServiceRegionVpcSettings struct {
	// The identifiers of the subnets for the directory servers.
	SubnetIds []string `json:"subnetIds,omitempty" yaml:"subnetIds,omitempty"`

	// The identifier of the VPC in which to create the directory.
	VpcId string `json:"vpcId,omitempty" yaml:"vpcId,omitempty"`
}
