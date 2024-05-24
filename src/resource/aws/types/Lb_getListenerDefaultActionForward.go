package types

type Lb_getListenerDefaultActionForward struct {
	//
	Stickinesses []Lb_getListenerDefaultActionForwardStickiness `json:"stickinesses,omitempty" yaml:"stickinesses,omitempty"`

	//
	TargetGroups []Lb_getListenerDefaultActionForwardTargetGroup `json:"targetGroups,omitempty" yaml:"targetGroups,omitempty"`
}
