package types

type Autoscaling_getGroupInstanceMaintenancePolicy struct {
	// Specifies the upper limit on the number of instances that are in the InService or Pending state with a healthy status during an instance replacement activity.
	MaxHealthyPercentage int `json:"maxHealthyPercentage,omitempty" yaml:"maxHealthyPercentage,omitempty"`

	// Specifies the lower limit on the number of instances that must be in the InService state with a healthy status during an instance replacement activity.
	MinHealthyPercentage int `json:"minHealthyPercentage,omitempty" yaml:"minHealthyPercentage,omitempty"`
}
