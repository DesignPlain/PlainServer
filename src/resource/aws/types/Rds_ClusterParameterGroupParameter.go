package types

type Rds_ClusterParameterGroupParameter struct {
	// The value of the DB parameter.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`

	/*
	   "immediate" (default), or "pending-reboot". Some
	   engines can't apply some parameters without a reboot, and you will need to
	   specify "pending-reboot" here.
	*/
	ApplyMethod string `json:"applyMethod,omitempty" yaml:"applyMethod,omitempty"`

	// The name of the DB parameter.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
