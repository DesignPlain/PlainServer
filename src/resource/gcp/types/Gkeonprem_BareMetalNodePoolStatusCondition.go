package types

type Gkeonprem_BareMetalNodePoolStatusCondition struct {
	/*
	   (Output)
	   Last time the condition transit from one status to another.
	*/
	LastTransitionTime string `json:"lastTransitionTime,omitempty" yaml:"lastTransitionTime,omitempty"`

	// Human-readable message indicating details about last transition.
	Message string `json:"message,omitempty" yaml:"message,omitempty"`

	// Machine-readable message indicating details about last transition.
	Reason string `json:"reason,omitempty" yaml:"reason,omitempty"`

	/*
	   (Output)
	   The lifecycle state of the condition.
	*/
	State string `json:"state,omitempty" yaml:"state,omitempty"`

	/*
	   Type of the condition.
	   (e.g., ClusterRunning, NodePoolRunning or ServerSidePreflightReady)
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
