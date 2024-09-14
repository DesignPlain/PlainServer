package types

type Cloudformation_StackSetInstanceStackInstanceSummary struct {
	// Organizational unit ID in which the stack is deployed.
	OrganizationalUnitId string `json:"organizationalUnitId,omitempty" yaml:"organizationalUnitId,omitempty"`

	// Stack identifier.
	StackId string `json:"stackId,omitempty" yaml:"stackId,omitempty"`

	// Target AWS Account ID to create a Stack based on the StackSet. Defaults to current account.
	AccountId string `json:"accountId,omitempty" yaml:"accountId,omitempty"`
}
