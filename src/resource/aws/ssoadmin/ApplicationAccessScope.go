package ssoadmin

type ApplicationAccessScope struct {
	// Specifies the ARN of the application with the access scope with the targets to add or update.
	ApplicationArn string `json:"applicationArn,omitempty" yaml:"applicationArn,omitempty"`

	// Specifies an array list of ARNs that represent the authorized targets for this access scope.
	AuthorizedTargets []string `json:"authorizedTargets,omitempty" yaml:"authorizedTargets,omitempty"`

	/*
	   Specifies the name of the access scope to be associated with the specified targets.

	   The following arguments are optional:
	*/
	Scope string `json:"scope,omitempty" yaml:"scope,omitempty"`
}
