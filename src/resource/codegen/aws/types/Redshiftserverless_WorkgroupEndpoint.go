package types

type Redshiftserverless_WorkgroupEndpoint struct {
	// The DNS address of the VPC endpoint.
	Address string `json:"address,omitempty" yaml:"address,omitempty"`

	// The port number on which the cluster accepts incoming connections.
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	// The VPC endpoint or the Redshift Serverless workgroup. See `VPC Endpoint` below.
	VpcEndpoints []Redshiftserverless_WorkgroupEndpointVpcEndpoint `json:"vpcEndpoints,omitempty" yaml:"vpcEndpoints,omitempty"`
}
