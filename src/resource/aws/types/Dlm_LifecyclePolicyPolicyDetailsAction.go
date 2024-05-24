package types

type Dlm_LifecyclePolicyPolicyDetailsAction struct {
	// The rule for copying shared snapshots across Regions. See the `cross_region_copy` configuration block.
	CrossRegionCopies []Dlm_LifecyclePolicyPolicyDetailsActionCrossRegionCopy `json:"crossRegionCopies,omitempty" yaml:"crossRegionCopies,omitempty"`

	// A descriptive name for the action.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
