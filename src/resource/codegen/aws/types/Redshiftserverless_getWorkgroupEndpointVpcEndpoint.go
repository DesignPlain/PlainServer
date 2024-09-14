package types

type Redshiftserverless_getWorkgroupEndpointVpcEndpoint struct {
	// The network interfaces of the endpoint.. See `Network Interface` below.
	NetworkInterfaces []Redshiftserverless_getWorkgroupEndpointVpcEndpointNetworkInterface `json:"networkInterfaces,omitempty" yaml:"networkInterfaces,omitempty"`

	// The DNS address of the VPC endpoint.
	VpcEndpointId string `json:"vpcEndpointId,omitempty" yaml:"vpcEndpointId,omitempty"`

	// The port that Amazon Redshift Serverless listens on.
	VpcId string `json:"vpcId,omitempty" yaml:"vpcId,omitempty"`
}
