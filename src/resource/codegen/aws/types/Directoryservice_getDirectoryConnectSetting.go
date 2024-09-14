package types

type Directoryservice_getDirectoryConnectSetting struct {
	//
	AvailabilityZones []string `json:"availabilityZones,omitempty" yaml:"availabilityZones,omitempty"`

	// IP addresses of the AD Connector servers.
	ConnectIps []string `json:"connectIps,omitempty" yaml:"connectIps,omitempty"`

	// DNS IP addresses of the domain to connect to.
	CustomerDnsIps []string `json:"customerDnsIps,omitempty" yaml:"customerDnsIps,omitempty"`

	// Username corresponding to the password provided.
	CustomerUsername string `json:"customerUsername,omitempty" yaml:"customerUsername,omitempty"`

	// Identifiers of the subnets for the connector servers (2 subnets in 2 different AZs).
	SubnetIds []string `json:"subnetIds,omitempty" yaml:"subnetIds,omitempty"`

	// ID of the VPC that the connector is in.
	VpcId string `json:"vpcId,omitempty" yaml:"vpcId,omitempty"`
}
