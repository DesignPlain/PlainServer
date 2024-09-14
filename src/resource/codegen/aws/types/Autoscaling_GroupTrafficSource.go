package types

type Autoscaling_GroupTrafficSource struct {
	// Identifies the traffic source. For Application Load Balancers, Gateway Load Balancers, Network Load Balancers, and VPC Lattice, this will be the Amazon Resource Name (ARN) for a target group in this account and Region. For Classic Load Balancers, this will be the name of the Classic Load Balancer in this account and Region.
	Identifier string `json:"identifier,omitempty" yaml:"identifier,omitempty"`

	/*
	   Provides additional context for the value of Identifier.
	   The following lists the valid values:
	   `elb` if `identifier` is the name of a Classic Load Balancer.
	   `elbv2` if `identifier` is the ARN of an Application Load Balancer, Gateway Load Balancer, or Network Load Balancer target group.
	   `vpc-lattice` if `identifier` is the ARN of a VPC Lattice target group.
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
