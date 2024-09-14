package types

type Iam_getPrincipalPolicySimulationResult struct {
	// A map of arbitrary metadata entries returned by the policy simulator for this request.
	DecisionDetails map[string]string `json:"decisionDetails,omitempty" yaml:"decisionDetails,omitempty"`

	// A nested set of objects describing which policies contained statements that were relevant to this simulation request. Each object has attributes `source_policy_id` and `source_policy_type` to identify one of the policies.
	MatchedStatements []Iam_getPrincipalPolicySimulationResultMatchedStatement `json:"matchedStatements,omitempty" yaml:"matchedStatements,omitempty"`

	// A set of context keys (or condition keys) that were needed by some of the policies contributing to this result but not specified using a `context` block in the configuration. Missing or incorrect context keys will typically cause a simulated request to be disallowed.
	MissingContextKeys []string `json:"missingContextKeys,omitempty" yaml:"missingContextKeys,omitempty"`

	// ARN of the resource that was used for this particular request. When you specify multiple actions and multiple resource ARNs, that causes a separate policy request for each combination of unique action and resource.
	ResourceArn string `json:"resourceArn,omitempty" yaml:"resourceArn,omitempty"`

	// The name of the single IAM action used for this particular request.
	ActionName string `json:"actionName,omitempty" yaml:"actionName,omitempty"`

	// `true` if `decision` is "allowed", and `false` otherwise.
	Allowed bool `json:"allowed,omitempty" yaml:"allowed,omitempty"`

	// The raw decision determined from all of the policies in scope; either "allowed", "explicitDeny", or "implicitDeny".
	Decision string `json:"decision,omitempty" yaml:"decision,omitempty"`
}
