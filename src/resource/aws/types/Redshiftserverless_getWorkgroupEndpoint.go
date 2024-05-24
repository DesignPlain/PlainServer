package types

type Redshiftserverless_getWorkgroupEndpoint struct {
	// The port that Amazon Redshift Serverless listens on.
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	// The VPC endpoint or the Redshift Serverless workgroup. See `VPC Endpoint` below.
	VpcEndpoints []Redshiftserverless_getWorkgroupEndpointVpcEndpoint `json:"vpcEndpoints,omitempty" yaml:"vpcEndpoints,omitempty"`

	// The DNS address of the VPC endpoint.
	Address string `json:"address,omitempty" yaml:"address,omitempty"`
}
