package types

type Glue_getConnectionPhysicalConnectionRequirement struct {
	//
	AvailabilityZone string `json:"availabilityZone,omitempty" yaml:"availabilityZone,omitempty"`

	//
	SecurityGroupIdLists []string `json:"securityGroupIdLists,omitempty" yaml:"securityGroupIdLists,omitempty"`

	//
	SubnetId string `json:"subnetId,omitempty" yaml:"subnetId,omitempty"`
}
