package types

type Ec2_FleetFleetInstanceSet struct {
	// The IDs of the instances.
	InstanceIds []string `json:"instanceIds,omitempty" yaml:"instanceIds,omitempty"`

	// Instance type.
	InstanceType string `json:"instanceType,omitempty" yaml:"instanceType,omitempty"`

	// Indicates if the instance that was launched is a Spot Instance or On-Demand Instance.
	Lifecycle string `json:"lifecycle,omitempty" yaml:"lifecycle,omitempty"`

	// The value is `Windows` for Windows instances. Otherwise, the value is blank.
	Platform string `json:"platform,omitempty" yaml:"platform,omitempty"`
}
