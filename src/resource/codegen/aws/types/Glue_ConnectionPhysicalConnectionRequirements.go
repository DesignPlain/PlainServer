package types

type Glue_ConnectionPhysicalConnectionRequirements struct {
	// The availability zone of the connection. This field is redundant and implied by `subnet_id`, but is currently an api requirement.
	AvailabilityZone string `json:"availabilityZone,omitempty" yaml:"availabilityZone,omitempty"`

	// The security group ID list used by the connection.
	SecurityGroupIdLists []string `json:"securityGroupIdLists,omitempty" yaml:"securityGroupIdLists,omitempty"`

	// The subnet ID used by the connection.
	SubnetId string `json:"subnetId,omitempty" yaml:"subnetId,omitempty"`
}
