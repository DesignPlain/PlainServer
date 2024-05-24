package servicecatalog

type PrincipalPortfolioAssociation struct {
	// Language code. Valid values: `en` (English), `jp` (Japanese), `zh` (Chinese). Default value is `en`.
	AcceptLanguage string `json:"acceptLanguage,omitempty" yaml:"acceptLanguage,omitempty"`

	// Portfolio identifier.
	PortfolioId string `json:"portfolioId,omitempty" yaml:"portfolioId,omitempty"`

	/*
	   Principal ARN.

	   The following arguments are optional:
	*/
	PrincipalArn string `json:"principalArn,omitempty" yaml:"principalArn,omitempty"`

	// Principal type. Setting this argument empty (e.g., `principal_type = ""`) will result in an error. Valid values are `IAM` and `IAM_PATTERN`. Default is `IAM`.
	PrincipalType string `json:"principalType,omitempty" yaml:"principalType,omitempty"`
}
