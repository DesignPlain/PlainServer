package vpc

type EndpointPrivateDns struct {
	// Indicates whether a private hosted zone is associated with the VPC. Only applicable for `Interface` endpoints.
	PrivateDnsEnabled bool `json:"privateDnsEnabled,omitempty" yaml:"privateDnsEnabled,omitempty"`

	// VPC endpoint identifier.
	VpcEndpointId string `json:"vpcEndpointId,omitempty" yaml:"vpcEndpointId,omitempty"`
}
