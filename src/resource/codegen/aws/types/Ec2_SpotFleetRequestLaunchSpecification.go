package types

type Ec2_SpotFleetRequestLaunchSpecification struct {
	//
	RootBlockDevices []Ec2_SpotFleetRequestLaunchSpecificationRootBlockDevice `json:"rootBlockDevices,omitempty" yaml:"rootBlockDevices,omitempty"`

	// The maximum bid price per unit hour.
	SpotPrice string `json:"spotPrice,omitempty" yaml:"spotPrice,omitempty"`

	//
	IamInstanceProfileArn string `json:"iamInstanceProfileArn,omitempty" yaml:"iamInstanceProfileArn,omitempty"`

	//
	AssociatePublicIpAddress bool `json:"associatePublicIpAddress,omitempty" yaml:"associatePublicIpAddress,omitempty"`

	//
	EbsOptimized bool `json:"ebsOptimized,omitempty" yaml:"ebsOptimized,omitempty"`

	//
	IamInstanceProfile string `json:"iamInstanceProfile,omitempty" yaml:"iamInstanceProfile,omitempty"`

	//
	Ami string `json:"ami,omitempty" yaml:"ami,omitempty"`

	// The type of instance to request.
	InstanceType string `json:"instanceType,omitempty" yaml:"instanceType,omitempty"`

	//
	KeyName string `json:"keyName,omitempty" yaml:"keyName,omitempty"`

	//
	Monitoring bool `json:"monitoring,omitempty" yaml:"monitoring,omitempty"`

	//
	PlacementGroup string `json:"placementGroup,omitempty" yaml:"placementGroup,omitempty"`

	//
	UserData string `json:"userData,omitempty" yaml:"userData,omitempty"`

	//
	VpcSecurityGroupIds []string `json:"vpcSecurityGroupIds,omitempty" yaml:"vpcSecurityGroupIds,omitempty"`

	//
	EphemeralBlockDevices []Ec2_SpotFleetRequestLaunchSpecificationEphemeralBlockDevice `json:"ephemeralBlockDevices,omitempty" yaml:"ephemeralBlockDevices,omitempty"`

	//
	EbsBlockDevices []Ec2_SpotFleetRequestLaunchSpecificationEbsBlockDevice `json:"ebsBlockDevices,omitempty" yaml:"ebsBlockDevices,omitempty"`

	//
	PlacementTenancy string `json:"placementTenancy,omitempty" yaml:"placementTenancy,omitempty"`

	// The subnet in which to launch the requested instance.
	SubnetId string `json:"subnetId,omitempty" yaml:"subnetId,omitempty"`

	// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The capacity added to the fleet by a fulfilled request.
	WeightedCapacity string `json:"weightedCapacity,omitempty" yaml:"weightedCapacity,omitempty"`

	// The availability zone in which to place the request.
	AvailabilityZone string `json:"availabilityZone,omitempty" yaml:"availabilityZone,omitempty"`
}
