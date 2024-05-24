package securityhub

type ActionTarget struct {
	// The name of the custom action target.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The ID for the custom action target.
	Identifier string `json:"identifier,omitempty" yaml:"identifier,omitempty"`

	// The description for the custom action target.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
