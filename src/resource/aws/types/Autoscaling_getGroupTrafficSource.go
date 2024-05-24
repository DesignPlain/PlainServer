package types

type Autoscaling_getGroupTrafficSource struct {
	// Identifies the traffic source. For Application Load Balancers, Gateway Load Balancers, Network Load Balancers, and VPC Lattice, this will be the Amazon Resource Name (ARN) for a target group in this account and Region. For Classic Load Balancers, this will be the name of the Classic Load Balancer in this account and Region.
	Identifier string `json:"identifier,omitempty" yaml:"identifier,omitempty"`

	// Traffic source type.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
