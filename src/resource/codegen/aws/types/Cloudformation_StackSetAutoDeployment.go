package types

type Cloudformation_StackSetAutoDeployment struct {
	// Whether or not to retain stacks when the account is removed.
	RetainStacksOnAccountRemoval bool `json:"retainStacksOnAccountRemoval,omitempty" yaml:"retainStacksOnAccountRemoval,omitempty"`

	// Whether or not auto-deployment is enabled.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
