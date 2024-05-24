package types

type Redshift_EndpointAccessVpcEndpoint struct {
	// One or more network interfaces of the endpoint. Also known as an interface endpoint. See details below.
	NetworkInterfaces []Redshift_EndpointAccessVpcEndpointNetworkInterface `json:"networkInterfaces,omitempty" yaml:"networkInterfaces,omitempty"`

	// The connection endpoint ID for connecting an Amazon Redshift cluster through the proxy.
	VpcEndpointId string `json:"vpcEndpointId,omitempty" yaml:"vpcEndpointId,omitempty"`

	// The VPC identifier that the endpoint is associated.
	VpcId string `json:"vpcId,omitempty" yaml:"vpcId,omitempty"`
}
