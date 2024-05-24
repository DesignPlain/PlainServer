package types

type Ssm_getPatchBaselineSource struct {
	// Value of the yum repo configuration.
	Configuration string `json:"configuration,omitempty" yaml:"configuration,omitempty"`

	// Name specified to identify the patch source.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Specific operating system versions a patch repository applies to.
	Products []string `json:"products,omitempty" yaml:"products,omitempty"`
}
