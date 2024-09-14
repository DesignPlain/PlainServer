package types

type Emrserverless_ApplicationNetworkConfiguration struct {
	// The array of subnet Ids for customer VPC connectivity.
	SubnetIds []string `json:"subnetIds,omitempty" yaml:"subnetIds,omitempty"`

	// The array of security group Ids for customer VPC connectivity.
	SecurityGroupIds []string `json:"securityGroupIds,omitempty" yaml:"securityGroupIds,omitempty"`
}
