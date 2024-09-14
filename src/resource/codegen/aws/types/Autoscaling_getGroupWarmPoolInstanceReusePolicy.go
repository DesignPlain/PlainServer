package types

type Autoscaling_getGroupWarmPoolInstanceReusePolicy struct {
	// Indicates whether instances in the Auto Scaling group can be returned to the warm pool on scale in.
	ReuseOnScaleIn bool `json:"reuseOnScaleIn,omitempty" yaml:"reuseOnScaleIn,omitempty"`
}
