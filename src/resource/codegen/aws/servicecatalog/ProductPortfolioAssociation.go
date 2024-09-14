package servicecatalog

type ProductPortfolioAssociation struct {
	// Language code. Valid values: `en` (English), `jp` (Japanese), `zh` (Chinese). Default value is `en`.
	AcceptLanguage string `json:"acceptLanguage,omitempty" yaml:"acceptLanguage,omitempty"`

	// Portfolio identifier.
	PortfolioId string `json:"portfolioId,omitempty" yaml:"portfolioId,omitempty"`

	/*
	   Product identifier.

	   The following arguments are optional:
	*/
	ProductId string `json:"productId,omitempty" yaml:"productId,omitempty"`

	// Identifier of the source portfolio.
	SourcePortfolioId string `json:"sourcePortfolioId,omitempty" yaml:"sourcePortfolioId,omitempty"`
}
