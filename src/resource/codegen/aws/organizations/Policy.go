package organizations

type Policy struct {
	// A description to assign to the policy.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The friendly name to assign to the policy.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// If set to `true`, destroy will --not-- delete the policy and instead just remove the resource from state. This can be useful in situations where the policies (and the associated attachment) must be preserved to meet the AWS minimum requirement of 1 attached policy.
	SkipDestroy bool `json:"skipDestroy,omitempty" yaml:"skipDestroy,omitempty"`

	// Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The type of policy to create. Valid values are `AISERVICES_OPT_OUT_POLICY`, `BACKUP_POLICY`, `SERVICE_CONTROL_POLICY` (SCP), and `TAG_POLICY`. Defaults to `SERVICE_CONTROL_POLICY`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// The policy content to add to the new policy. For example, if you create a [service control policy (SCP)](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_scp.html), this string must be JSON text that specifies the permissions that admins in attached accounts can delegate to their users, groups, and roles. For more information about the SCP syntax, see the [Service Control Policy Syntax documentation](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_reference_scp-syntax.html) and for more information on the Tag Policy syntax, see the [Tag Policy Syntax documentation](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_example-tag-policies.html).
	Content string `json:"content,omitempty" yaml:"content,omitempty"`
}
