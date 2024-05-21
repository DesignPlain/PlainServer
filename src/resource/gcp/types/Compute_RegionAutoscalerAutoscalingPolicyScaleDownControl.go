package types

type Compute_RegionAutoscalerAutoscalingPolicyScaleDownControl struct {
	/*
	   A nested object resource
	   Structure is documented below.
	*/
	MaxScaledDownReplicas Compute_RegionAutoscalerAutoscalingPolicyScaleDownControlMaxScaledDownReplicas `json:"maxScaledDownReplicas,omitempty" yaml:"maxScaledDownReplicas,omitempty"`

	/*
	   How long back autoscaling should look when computing recommendations
	   to include directives regarding slower scale down, as described above.
	*/
	TimeWindowSec int `json:"timeWindowSec,omitempty" yaml:"timeWindowSec,omitempty"`
}
