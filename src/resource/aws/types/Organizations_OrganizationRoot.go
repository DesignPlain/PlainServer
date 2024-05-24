package types

type Organizations_OrganizationRoot struct {
	// ARN of the root
	Arn string `json:"arn,omitempty" yaml:"arn,omitempty"`

	// Identifier of the root
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	// The name of the policy type
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// List of policy types enabled for this root. All elements have these attributes:
	PolicyTypes []Organizations_OrganizationRootPolicyType `json:"policyTypes,omitempty" yaml:"policyTypes,omitempty"`
}
