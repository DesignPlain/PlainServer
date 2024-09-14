package controltower

import types "libds/aws/types"

type ControlTowerControl struct {
	// The ARN of the control. Only Strongly recommended and Elective controls are permitted, with the exception of the Region deny guardrail.
	ControlIdentifier string `json:"controlIdentifier,omitempty" yaml:"controlIdentifier,omitempty"`

	// Parameter values which are specified to configure the control when you enable it. See Parameters for more details.
	Parameters []types.Controltower_ControlTowerControlParameter `json:"parameters,omitempty" yaml:"parameters,omitempty"`

	/*
	   The ARN of the organizational unit.

	   The following arguments are optional:
	*/
	TargetIdentifier string `json:"targetIdentifier,omitempty" yaml:"targetIdentifier,omitempty"`
}
