package types

type Appautoscaling_PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimension struct {
	// Name of the policy. Must be between 1 and 255 characters in length.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Value of the dimension.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
