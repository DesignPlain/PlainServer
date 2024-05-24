package types

type Servicecatalog_getPortfolioConstraintsDetail struct {
	// Product identifier.
	ProductId string `json:"productId,omitempty" yaml:"productId,omitempty"`

	// Type of constraint. Valid values are `LAUNCH`, `NOTIFICATION`, `STACKSET`, and `TEMPLATE`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Identifier of the constraint.
	ConstraintId string `json:"constraintId,omitempty" yaml:"constraintId,omitempty"`

	// Description of the constraint.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	//
	Owner string `json:"owner,omitempty" yaml:"owner,omitempty"`

	/*
	   Portfolio identifier.

	   The following arguments are optional:
	*/
	PortfolioId string `json:"portfolioId,omitempty" yaml:"portfolioId,omitempty"`
}
