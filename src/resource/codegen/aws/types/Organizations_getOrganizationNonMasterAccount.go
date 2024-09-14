package types

type Organizations_getOrganizationNonMasterAccount struct {
	// ARN of the root
	Arn string `json:"arn,omitempty" yaml:"arn,omitempty"`

	// Email of the account
	Email string `json:"email,omitempty" yaml:"email,omitempty"`

	// Identifier of the root
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	// The name of the policy type
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The status of the policy type as it relates to the associated root
	Status string `json:"status,omitempty" yaml:"status,omitempty"`
}
