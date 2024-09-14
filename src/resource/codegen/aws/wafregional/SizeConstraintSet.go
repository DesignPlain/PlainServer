package wafregional

import types "libds/aws/types"

type SizeConstraintSet struct {
	// The name or description of the Size Constraint Set.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Specifies the parts of web requests that you want to inspect the size of.
	SizeConstraints []types.Wafregional_SizeConstraintSetSizeConstraint `json:"sizeConstraints,omitempty" yaml:"sizeConstraints,omitempty"`
}
