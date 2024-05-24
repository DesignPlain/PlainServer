package opensearch

type ServerlessLifecyclePolicy struct {
	// JSON policy document to use as the content for the new policy.
	Policy string `json:"policy,omitempty" yaml:"policy,omitempty"`

	/*
	   Type of lifecycle policy. Must be `retention`.

	   The following arguments are optional:
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Description of the policy.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Name of the policy.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
