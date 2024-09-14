package types

type Compute_AutoscalerAutoscalingPolicyScaleDownControl struct {
	/*
	   A nested object resource
	   Structure is documented below.
	*/
	MaxScaledDownReplicas Compute_AutoscalerAutoscalingPolicyScaleDownControlMaxScaledDownReplicas `json:"maxScaledDownReplicas,omitempty" yaml:"maxScaledDownReplicas,omitempty"`

	/*
	   How long back autoscaling should look when computing recommendations
	   to include directives regarding slower scale down, as described above.
	*/
	TimeWindowSec int `json:"timeWindowSec,omitempty" yaml:"timeWindowSec,omitempty"`
}
