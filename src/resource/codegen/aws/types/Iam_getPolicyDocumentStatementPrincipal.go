package types

type Iam_getPolicyDocumentStatementPrincipal struct {
	// List of identifiers for principals. When `type` is `AWS`, these are IAM principal ARNs, e.g., `arn:aws:iam::12345678901:role/yak-role`.  When `type` is `Service`, these are AWS Service roles, e.g., `lambda.amazonaws.com`. When `type` is `Federated`, these are web identity users or SAML provider ARNs, e.g., `accounts.google.com` or `arn:aws:iam::12345678901:saml-provider/yak-saml-provider`. When `type` is `CanonicalUser`, these are [canonical user IDs](https://docs.aws.amazon.com/general/latest/gr/acct-identifiers.html#FindingCanonicalId), e.g., `79a59df900b949e55d96a1e698fbacedfd6e09d98eacf8f8d5218e7cd47ef2be`.
	Identifiers []string `json:"identifiers,omitempty" yaml:"identifiers,omitempty"`

	// Type of principal. Valid values include `AWS`, `Service`, `Federated`, `CanonicalUser` and `-`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
