package types

type Gkebackup_RestorePlanRestoreConfigTransformationRule struct {
	/*
	   The description is a user specified string description
	   of the transformation rule.
	*/
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   A list of transformation rule actions to take against candidate
	   resources. Actions are executed in order defined - this order
	   matters, as they could potentially interfere with each other and
	   the first operation could affect the outcome of the second operation.
	   Structure is documented below.
	*/
	FieldActions []Gkebackup_RestorePlanRestoreConfigTransformationRuleFieldAction `json:"fieldActions,omitempty" yaml:"fieldActions,omitempty"`

	/*
	   This field is used to specify a set of fields that should be used to
	   determine which resources in backup should be acted upon by the
	   supplied transformation rule actions, and this will ensure that only
	   specific resources are affected by transformation rule actions.
	   Structure is documented below.
	*/
	ResourceFilter Gkebackup_RestorePlanRestoreConfigTransformationRuleResourceFilter `json:"resourceFilter,omitempty" yaml:"resourceFilter,omitempty"`
}
