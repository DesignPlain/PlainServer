package ssoadmin

type PermissionSet struct {
	// The description of the Permission Set.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The Amazon Resource Name (ARN) of the SSO Instance under which the operation will be executed.
	InstanceArn string `json:"instanceArn,omitempty" yaml:"instanceArn,omitempty"`

	// The name of the Permission Set.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The relay state URL used to redirect users within the application during the federation authentication process.
	RelayState string `json:"relayState,omitempty" yaml:"relayState,omitempty"`

	// The length of time that the application user sessions are valid in the ISO-8601 standard. Default: `PT1H`.
	SessionDuration string `json:"sessionDuration,omitempty" yaml:"sessionDuration,omitempty"`

	// Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
