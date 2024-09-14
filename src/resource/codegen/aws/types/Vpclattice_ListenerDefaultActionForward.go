package types

type Vpclattice_ListenerDefaultActionForward struct {
	// One or more target group blocks.
	TargetGroups []Vpclattice_ListenerDefaultActionForwardTargetGroup `json:"targetGroups,omitempty" yaml:"targetGroups,omitempty"`
}
