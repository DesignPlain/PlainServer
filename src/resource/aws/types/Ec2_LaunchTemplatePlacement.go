package types

type Ec2_LaunchTemplatePlacement struct {
	// Reserved for future use.
	SpreadDomain string `json:"spreadDomain,omitempty" yaml:"spreadDomain,omitempty"`

	// The tenancy of the instance (if the instance is running in a VPC). Can be `default`, `dedicated`, or `host`.
	Tenancy string `json:"tenancy,omitempty" yaml:"tenancy,omitempty"`

	// The affinity setting for an instance on a Dedicated Host.
	Affinity string `json:"affinity,omitempty" yaml:"affinity,omitempty"`

	// The Availability Zone for the instance.
	AvailabilityZone string `json:"availabilityZone,omitempty" yaml:"availabilityZone,omitempty"`

	// The name of the placement group for the instance.
	GroupName string `json:"groupName,omitempty" yaml:"groupName,omitempty"`

	// The ID of the Dedicated Host for the instance.
	HostId string `json:"hostId,omitempty" yaml:"hostId,omitempty"`

	// The ARN of the Host Resource Group in which to launch instances.
	HostResourceGroupArn string `json:"hostResourceGroupArn,omitempty" yaml:"hostResourceGroupArn,omitempty"`

	// The number of the partition the instance should launch in. Valid only if the placement group strategy is set to partition.
	PartitionNumber int `json:"partitionNumber,omitempty" yaml:"partitionNumber,omitempty"`
}
