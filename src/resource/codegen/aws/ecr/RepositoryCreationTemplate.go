package ecr

import types "libds/aws/types"

type RepositoryCreationTemplate struct {
	// Which features this template applies to. Must contain one or more of `PULL_THROUGH_CACHE` or `REPLICATION`.
	AppliedFors []string `json:"appliedFors,omitempty" yaml:"appliedFors,omitempty"`

	// A custom IAM role to use for repository creation. Required if using repository tags or KMS encryption.
	CustomRoleArn string `json:"customRoleArn,omitempty" yaml:"customRoleArn,omitempty"`

	// The description for this template.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Encryption configuration for any created repositories. See below for schema.
	EncryptionConfigurations []types.Ecr_RepositoryCreationTemplateEncryptionConfiguration `json:"encryptionConfigurations,omitempty" yaml:"encryptionConfigurations,omitempty"`

	// The tag mutability setting for any created repositories. Must be one of: `MUTABLE` or `IMMUTABLE`. Defaults to `MUTABLE`.
	ImageTagMutability string `json:"imageTagMutability,omitempty" yaml:"imageTagMutability,omitempty"`

	// The repository name prefix to match against. Use `ROOT` to match any prefix that doesn't explicitly match another template.
	Prefix string `json:"prefix,omitempty" yaml:"prefix,omitempty"`

	// A map of tags to assign to any created repositories.
	ResourceTags map[string]string `json:"resourceTags,omitempty" yaml:"resourceTags,omitempty"`

	// The lifecycle policy document to apply to any created repositories. See more details about [Policy Parameters](http://docs.aws.amazon.com/AmazonECR/latest/userguide/LifecyclePolicies.html#lifecycle_policy_parameters) in the official AWS docs. Consider using the `aws.ecr.getLifecyclePolicyDocument` data_source to generate/manage the JSON document used for the `lifecycle_policy` argument.
	LifecyclePolicy string `json:"lifecyclePolicy,omitempty" yaml:"lifecyclePolicy,omitempty"`

	//
	RepositoryPolicy string `json:"repositoryPolicy,omitempty" yaml:"repositoryPolicy,omitempty"`
}
