package types

type Autoscaling_GroupWarmPoolInstanceReusePolicy struct {
	// Whether instances in the Auto Scaling group can be returned to the warm pool on scale in.
	ReuseOnScaleIn bool `json:"reuseOnScaleIn,omitempty" yaml:"reuseOnScaleIn,omitempty"`
}
