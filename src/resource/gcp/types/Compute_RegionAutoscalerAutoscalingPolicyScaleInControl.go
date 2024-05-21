package types

type Compute_RegionAutoscalerAutoscalingPolicyScaleInControl struct {
	/*
	   A nested object resource
	   Structure is documented below.
	*/
	MaxScaledInReplicas Compute_RegionAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas `json:"maxScaledInReplicas,omitempty" yaml:"maxScaledInReplicas,omitempty"`

	/*
	   How long back autoscaling should look when computing recommendations
	   to include directives regarding slower scale down, as described above.
	*/
	TimeWindowSec int `json:"timeWindowSec,omitempty" yaml:"timeWindowSec,omitempty"`
}
