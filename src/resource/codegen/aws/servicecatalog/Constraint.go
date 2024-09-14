package servicecatalog

type Constraint struct {
	// Description of the constraint.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Constraint parameters in JSON format. The syntax depends on the constraint type. See details below.
	Parameters string `json:"parameters,omitempty" yaml:"parameters,omitempty"`

	// Portfolio identifier.
	PortfolioId string `json:"portfolioId,omitempty" yaml:"portfolioId,omitempty"`

	// Product identifier.
	ProductId string `json:"productId,omitempty" yaml:"productId,omitempty"`

	/*
	   Type of constraint. Valid values are `LAUNCH`, `NOTIFICATION`, `RESOURCE_UPDATE`, `STACKSET`, and `TEMPLATE`.

	   The following arguments are optional:
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Language code. Valid values: `en` (English), `jp` (Japanese), `zh` (Chinese). Default value is `en`.
	AcceptLanguage string `json:"acceptLanguage,omitempty" yaml:"acceptLanguage,omitempty"`
}
