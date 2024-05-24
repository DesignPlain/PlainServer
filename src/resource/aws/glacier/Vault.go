package glacier

import types "DesignSphere_Server/src/resource/aws/types"

type Vault struct {
	/*
	   The policy document. This is a JSON formatted string.
	   The heredoc syntax or `file` function is helpful here. Use the [Glacier Developer Guide](https://docs.aws.amazon.com/amazonglacier/latest/dev/vault-access-policy.html) for more information on Glacier Vault Policy
	*/
	AccessPolicy string `json:"accessPolicy,omitempty" yaml:"accessPolicy,omitempty"`

	// The name of the Vault. Names can be between 1 and 255 characters long and the valid characters are a-z, A-Z, 0-9, '_' (underscore), '-' (hyphen), and '.' (period).
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The notifications for the Vault. Fields documented below.
	Notification types.Glacier_VaultNotification `json:"notification,omitempty" yaml:"notification,omitempty"`

	// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
