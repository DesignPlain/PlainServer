package ec2

type AvailabilityZoneGroup struct {
	// Name of the Availability Zone Group.
	GroupName string `json:"groupName,omitempty" yaml:"groupName,omitempty"`

	// Indicates whether to enable or disable Availability Zone Group. Valid values: `opted-in` or `not-opted-in`.
	OptInStatus string `json:"optInStatus,omitempty" yaml:"optInStatus,omitempty"`
}
