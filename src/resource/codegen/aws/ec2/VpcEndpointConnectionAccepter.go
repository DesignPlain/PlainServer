package ec2

type VpcEndpointConnectionAccepter struct {
	// AWS VPC Endpoint ID.
	VpcEndpointId string `json:"vpcEndpointId,omitempty" yaml:"vpcEndpointId,omitempty"`

	// AWS VPC Endpoint Service ID.
	VpcEndpointServiceId string `json:"vpcEndpointServiceId,omitempty" yaml:"vpcEndpointServiceId,omitempty"`
}
