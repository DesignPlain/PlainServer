package types

type Ec2_VpcEndpointServicePrivateDnsNameConfiguration struct {
	// Name of the record subdomain the service provider needs to create.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Verification state of the VPC endpoint service. Consumers of the endpoint service can use the private name only when the state is `verified`.
	State string `json:"state,omitempty" yaml:"state,omitempty"`

	// Endpoint service verification type, for example `TXT`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Value the service provider adds to the private DNS name domain record before verification.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
