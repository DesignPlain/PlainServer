package types

type Opsworks_StackCustomCookbooksSource struct {
	// Username to use when authenticating to the source.
	Username string `json:"username,omitempty" yaml:"username,omitempty"`

	// Password to use when authenticating to the source. The provider cannot perform drift detection of this configuration.
	Password string `json:"password,omitempty" yaml:"password,omitempty"`

	// For sources that are version-aware, the revision to use.
	Revision string `json:"revision,omitempty" yaml:"revision,omitempty"`

	// SSH key to use when authenticating to the source. This provider cannot perform drift detection of this configuration.
	SshKey string `json:"sshKey,omitempty" yaml:"sshKey,omitempty"`

	// The type of source to use. For example, "archive".
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// The URL where the cookbooks resource can be found.
	Url string `json:"url,omitempty" yaml:"url,omitempty"`
}
