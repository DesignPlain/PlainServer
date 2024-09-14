package types

type Ec2_getLaunchTemplatePlacement struct {
	//
	PartitionNumber int `json:"partitionNumber,omitempty" yaml:"partitionNumber,omitempty"`

	//
	SpreadDomain string `json:"spreadDomain,omitempty" yaml:"spreadDomain,omitempty"`

	//
	Tenancy string `json:"tenancy,omitempty" yaml:"tenancy,omitempty"`

	//
	Affinity string `json:"affinity,omitempty" yaml:"affinity,omitempty"`

	//
	AvailabilityZone string `json:"availabilityZone,omitempty" yaml:"availabilityZone,omitempty"`

	//
	GroupName string `json:"groupName,omitempty" yaml:"groupName,omitempty"`

	//
	HostId string `json:"hostId,omitempty" yaml:"hostId,omitempty"`

	//
	HostResourceGroupArn string `json:"hostResourceGroupArn,omitempty" yaml:"hostResourceGroupArn,omitempty"`
}
