package servicecatalog

type PortfolioShare struct {
	// Language code. Valid values: `en` (English), `jp` (Japanese), `zh` (Chinese). Default value is `en`.
	AcceptLanguage string `json:"acceptLanguage,omitempty" yaml:"acceptLanguage,omitempty"`

	// Portfolio identifier.
	PortfolioId string `json:"portfolioId,omitempty" yaml:"portfolioId,omitempty"`

	// Identifier of the principal with whom you will share the portfolio. Valid values AWS account IDs and ARNs of AWS Organizations and organizational units.
	PrincipalId string `json:"principalId,omitempty" yaml:"principalId,omitempty"`

	// Enables or disables Principal sharing when creating the portfolio share. If this flag is not provided, principal sharing is disabled.
	SharePrincipals bool `json:"sharePrincipals,omitempty" yaml:"sharePrincipals,omitempty"`

	// Whether to enable sharing of `aws.servicecatalog.TagOption` resources when creating the portfolio share.
	ShareTagOptions bool `json:"shareTagOptions,omitempty" yaml:"shareTagOptions,omitempty"`

	/*
	   Type of portfolio share. Valid values are `ACCOUNT` (an external account), `ORGANIZATION` (a share to every account in an organization), `ORGANIZATIONAL_UNIT`, `ORGANIZATION_MEMBER_ACCOUNT` (a share to an account in an organization).

	   The following arguments are optional:
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Whether to wait (up to the timeout) for the share to be accepted. Organizational shares are automatically accepted.
	WaitForAcceptance bool `json:"waitForAcceptance,omitempty" yaml:"waitForAcceptance,omitempty"`
}
