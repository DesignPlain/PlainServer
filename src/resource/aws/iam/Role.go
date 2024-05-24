package iam

import types "DesignSphere_Server/src/resource/aws/types"

type Role struct {
	// Maximum session duration (in seconds) that you want to set for the specified role. If you do not specify a value for this setting, the default maximum of one hour is applied. This setting can have a value from 1 hour to 12 hours.
	MaxSessionDuration int `json:"maxSessionDuration,omitempty" yaml:"maxSessionDuration,omitempty"`

	// Creates a unique friendly name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix string `json:"namePrefix,omitempty" yaml:"namePrefix,omitempty"`

	// Path to the role. See [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) for more information.
	Path string `json:"path,omitempty" yaml:"path,omitempty"`

	/*
	   Policy that grants an entity permission to assume the role.

	   > --NOTE:-- The `assume_role_policy` is very similar to but slightly different than a standard IAM policy and cannot use an `aws.iam.Policy` resource.  However, it _can_ use an `aws.iam.getPolicyDocument` data source. See the example above of how this works.

	   The following arguments are optional:
	*/
	AssumeRolePolicy string `json:"assumeRolePolicy,omitempty" yaml:"assumeRolePolicy,omitempty"`

	// Description of the role.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Whether to force detaching any policies the role has before destroying it. Defaults to `false`.
	ForceDetachPolicies bool `json:"forceDetachPolicies,omitempty" yaml:"forceDetachPolicies,omitempty"`

	//
	ManagedPolicyArns []string `json:"managedPolicyArns,omitempty" yaml:"managedPolicyArns,omitempty"`

	// Configuration block defining an exclusive set of IAM inline policies associated with the IAM role. See below. If no blocks are configured, the provider will not manage any inline policies in this resource. Configuring one empty block (i.e., `inline_policy {}`) will cause the provider to remove _all_ inline policies added out of band on `apply`.
	InlinePolicies []types.Iam_RoleInlinePolicy `json:"inlinePolicies,omitempty" yaml:"inlinePolicies,omitempty"`

	// Friendly name of the role. If omitted, the provider will assign a random, unique name. See [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) for more information.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// ARN of the policy that is used to set the permissions boundary for the role.
	PermissionsBoundary string `json:"permissionsBoundary,omitempty" yaml:"permissionsBoundary,omitempty"`

	// Key-value mapping of tags for the IAM role. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
