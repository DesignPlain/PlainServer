package types

type Emr_InstanceFleetLaunchSpecificationsOnDemandSpecification struct {
	// Specifies one of the following strategies to launch Spot Instance fleets: `price-capacity-optimized`, `capacity-optimized`, `lowest-price`, or `diversified`. For more information on the provisioning strategies, see [Allocation strategies for Spot Instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-fleet-allocation-strategy.html).
	AllocationStrategy string `json:"allocationStrategy,omitempty" yaml:"allocationStrategy,omitempty"`
}
