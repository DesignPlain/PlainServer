package securityposture

import types "libds/gcp/types"

type Posture struct {
	// The parent of the resource, an organization. Format should be `organizations/{organization_id}`.
	Parent string `json:"parent,omitempty" yaml:"parent,omitempty"`

	/*
	   List of policy sets for the posture.
	   Structure is documented below.
	*/
	PolicySets []types.Securityposture_PosturePolicySet `json:"policySets,omitempty" yaml:"policySets,omitempty"`

	// Id of the posture. It is an immutable field.
	PostureId string `json:"postureId,omitempty" yaml:"postureId,omitempty"`

	/*
	   State of the posture. Update to state field should not be triggered along with
	   with other field updates.
	   Possible values are: `DEPRECATED`, `DRAFT`, `ACTIVE`.
	*/
	State string `json:"state,omitempty" yaml:"state,omitempty"`

	// Description of the expression
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Location of the resource, eg: global.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
}
