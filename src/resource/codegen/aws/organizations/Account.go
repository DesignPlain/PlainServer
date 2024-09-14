package organizations

type Account struct {
	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// If true, a deletion event will close the account. Otherwise, it will only remove from the organization. This is not supported for GovCloud accounts.
	CloseOnDeletion bool `json:"closeOnDeletion,omitempty" yaml:"closeOnDeletion,omitempty"`

	// Whether to also create a GovCloud account. The GovCloud account is tied to the main (commercial) account this resource creates. If `true`, the GovCloud account ID is available in the `govcloud_id` attribute. The only way to manage the GovCloud account with the provider is to subsequently import the account using this resource.
	CreateGovcloud bool `json:"createGovcloud,omitempty" yaml:"createGovcloud,omitempty"`

	// Email address of the owner to assign to the new member account. This email address must not already be associated with another AWS account.
	Email string `json:"email,omitempty" yaml:"email,omitempty"`

	// If set to `ALLOW`, the new account enables IAM users and roles to access account billing information if they have the required permissions. If set to `DENY`, then only the root user (and no roles) of the new account can access account billing information. If this is unset, the AWS API will default this to `ALLOW`. If the resource is created and this option is changed, it will try to recreate the account.
	IamUserAccessToBilling string `json:"iamUserAccessToBilling,omitempty" yaml:"iamUserAccessToBilling,omitempty"`

	/*
	   Friendly name for the member account.

	   The following arguments are optional:
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Parent Organizational Unit ID or Root ID for the account. Defaults to the Organization default Root ID. A configuration must be present for this argument to perform drift detection.
	ParentId string `json:"parentId,omitempty" yaml:"parentId,omitempty"`

	// The name of an IAM role that Organizations automatically preconfigures in the new member account. This role trusts the root account, allowing users in the root account to assume the role, as permitted by the root account administrator. The role has administrator permissions in the new member account. The Organizations API provides no method for reading this information after account creation, so the provider cannot perform drift detection on its value and will always show a difference for a configured value after import unless `ignoreChanges` is used.
	RoleName string `json:"roleName,omitempty" yaml:"roleName,omitempty"`
}
