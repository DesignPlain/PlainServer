package types



type Compute_AutoscalerAutoscalingPolicyScaleInControl struct {
	/*
	   How long back autoscaling should look when computing recommendations
	   to include directives regarding slower scale down, as described above.
	*/
	TimeWindowSec int `json:"timeWindowSec,omitempty" yaml:"timeWindowSec,omitempty"`

	/*
	   A nested object resource
	   Structure is documented below.
	*/
	MaxScaledInReplicas Compute_AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas `json:"maxScaledInReplicas,omitempty" yaml:"maxScaledInReplicas,omitempty"`
}