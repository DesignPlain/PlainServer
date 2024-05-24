package controltower

type ControlTowerControl struct {
	// The ARN of the control. Only Strongly recommended and Elective controls are permitted, with the exception of the Region deny guardrail.
	ControlIdentifier string `json:"controlIdentifier,omitempty" yaml:"controlIdentifier,omitempty"`

	// The ARN of the organizational unit.
	TargetIdentifier string `json:"targetIdentifier,omitempty" yaml:"targetIdentifier,omitempty"`
}
