package verifiedpermissions

type PolicyTemplate struct {
	// Provides a description for the policy template.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The ID of the Policy Store.
	PolicyStoreId string `json:"policyStoreId,omitempty" yaml:"policyStoreId,omitempty"`

	/*
	   Defines the content of the statement, written in Cedar policy language.

	   The following arguments are optional:
	*/
	Statement string `json:"statement,omitempty" yaml:"statement,omitempty"`
}
