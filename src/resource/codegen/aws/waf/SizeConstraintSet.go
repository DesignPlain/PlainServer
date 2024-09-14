package waf

import types "libds/aws/types"

type SizeConstraintSet struct {
	// Name or description of the Size Constraint Set.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Parts of web requests that you want to inspect the size of.
	SizeConstraints []types.Waf_SizeConstraintSetSizeConstraint `json:"sizeConstraints,omitempty" yaml:"sizeConstraints,omitempty"`
}
