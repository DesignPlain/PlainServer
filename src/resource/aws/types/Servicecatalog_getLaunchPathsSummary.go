package types

type Servicecatalog_getLaunchPathsSummary struct {
	// Block for constraints on the portfolio-product relationship. See details below.
	ConstraintSummaries []Servicecatalog_getLaunchPathsSummaryConstraintSummary `json:"constraintSummaries,omitempty" yaml:"constraintSummaries,omitempty"`

	// Name of the portfolio to which the path was assigned.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Identifier of the product path.
	PathId string `json:"pathId,omitempty" yaml:"pathId,omitempty"`

	// Tags associated with this product path.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
