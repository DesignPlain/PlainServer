package types

type Iam_getPolicyDocumentStatement struct {
	// List of actions that this statement either allows or denies. For example, `["ec2:RunInstances", "s3:-"]`.
	Actions []string `json:"actions,omitempty" yaml:"actions,omitempty"`

	// Whether this statement allows or denies the given actions. Valid values are `Allow` and `Deny`. Defaults to `Allow`.
	Effect string `json:"effect,omitempty" yaml:"effect,omitempty"`

	// Like `principals` except these are principals that the statement does -not- apply to.
	NotPrincipals []Iam_getPolicyDocumentStatementNotPrincipal `json:"notPrincipals,omitempty" yaml:"notPrincipals,omitempty"`

	// Configuration block for principals. Detailed below.
	Principals []Iam_getPolicyDocumentStatementPrincipal `json:"principals,omitempty" yaml:"principals,omitempty"`

	// List of resource ARNs that this statement applies to. This is required by AWS if used for an IAM policy. Conflicts with `not_resources`.
	Resources []string `json:"resources,omitempty" yaml:"resources,omitempty"`

	// Configuration block for a condition. Detailed below.
	Conditions []Iam_getPolicyDocumentStatementCondition `json:"conditions,omitempty" yaml:"conditions,omitempty"`

	// List of actions that this statement does -not- apply to. Use to apply a policy statement to all actions -except- those listed.
	NotActions []string `json:"notActions,omitempty" yaml:"notActions,omitempty"`

	// List of resource ARNs that this statement does -not- apply to. Use to apply a policy statement to all resources -except- those listed. Conflicts with `resources`.
	NotResources []string `json:"notResources,omitempty" yaml:"notResources,omitempty"`

	// Sid (statement ID) is an identifier for a policy statement.
	Sid string `json:"sid,omitempty" yaml:"sid,omitempty"`
}
