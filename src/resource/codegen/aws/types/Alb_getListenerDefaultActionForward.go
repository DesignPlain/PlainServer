package types

type Alb_getListenerDefaultActionForward struct {
	//
	Stickinesses []Alb_getListenerDefaultActionForwardStickiness `json:"stickinesses,omitempty" yaml:"stickinesses,omitempty"`

	//
	TargetGroups []Alb_getListenerDefaultActionForwardTargetGroup `json:"targetGroups,omitempty" yaml:"targetGroups,omitempty"`
}
