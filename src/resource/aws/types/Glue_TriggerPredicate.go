package types

type Glue_TriggerPredicate struct {
	// How to handle multiple conditions. Defaults to `AND`. Valid values are `AND` or `ANY`.
	Logical string `json:"logical,omitempty" yaml:"logical,omitempty"`

	// A list of the conditions that determine when the trigger will fire. See Conditions.
	Conditions []Glue_TriggerPredicateCondition `json:"conditions,omitempty" yaml:"conditions,omitempty"`
}
