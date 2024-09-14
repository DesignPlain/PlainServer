package types

type Cognito_IdentityPoolRoleAttachmentRoleMappingMappingRule struct {
	// The claim name that must be present in the token, for example, "isAdmin" or "paid".
	Claim string `json:"claim,omitempty" yaml:"claim,omitempty"`

	// The match condition that specifies how closely the claim value in the IdP token must match Value.
	MatchType string `json:"matchType,omitempty" yaml:"matchType,omitempty"`

	// The role ARN.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// A brief string that the claim must match, for example, "paid" or "yes".
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
