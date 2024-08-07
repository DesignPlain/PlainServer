package opensearch

type ServerlessAccessPolicy struct {
	// Description of the policy. Typically used to store information about the permissions defined in the policy.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Name of the policy.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// JSON policy document to use as the content for the new policy
	Policy string `json:"policy,omitempty" yaml:"policy,omitempty"`

	/*
	   Type of access policy. Must be `data`.

	   The following arguments are optional:
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
