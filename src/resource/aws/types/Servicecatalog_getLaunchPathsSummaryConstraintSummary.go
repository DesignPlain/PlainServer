package types

type Servicecatalog_getLaunchPathsSummaryConstraintSummary struct {
	// Description of the constraint.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Type of constraint. Valid values are `LAUNCH`, `NOTIFICATION`, `STACKSET`, and `TEMPLATE`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
