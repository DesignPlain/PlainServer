package ecr

type LifecyclePolicy struct {
	// The policy document. This is a JSON formatted string. See more details about [Policy Parameters](http://docs.aws.amazon.com/AmazonECR/latest/userguide/LifecyclePolicies.html#lifecycle_policy_parameters) in the official AWS docs. Consider using the `aws.ecr.getLifecyclePolicyDocument` data_source to generate/manage the JSON document used for the `policy` argument.
	Policy string `json:"policy,omitempty" yaml:"policy,omitempty"`

	// Name of the repository to apply the policy.
	Repository string `json:"repository,omitempty" yaml:"repository,omitempty"`
}
