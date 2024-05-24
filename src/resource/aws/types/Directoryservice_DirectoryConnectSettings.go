package types

type Directoryservice_DirectoryConnectSettings struct {
	// The username corresponding to the password provided.
	CustomerUsername string `json:"customerUsername,omitempty" yaml:"customerUsername,omitempty"`

	// The identifiers of the subnets for the directory servers (2 subnets in 2 different AZs).
	SubnetIds []string `json:"subnetIds,omitempty" yaml:"subnetIds,omitempty"`

	// The identifier of the VPC that the directory is in.
	VpcId string `json:"vpcId,omitempty" yaml:"vpcId,omitempty"`

	//
	AvailabilityZones []string `json:"availabilityZones,omitempty" yaml:"availabilityZones,omitempty"`

	// The IP addresses of the AD Connector servers.
	ConnectIps []string `json:"connectIps,omitempty" yaml:"connectIps,omitempty"`

	// The DNS IP addresses of the domain to connect to.
	CustomerDnsIps []string `json:"customerDnsIps,omitempty" yaml:"customerDnsIps,omitempty"`
}
