package appconfig

type ExtensionAssociation struct {
	// The ARN of the extension defined in the association.
	ExtensionArn string `json:"extensionArn,omitempty" yaml:"extensionArn,omitempty"`

	// The parameter names and values defined for the association.
	Parameters map[string]string `json:"parameters,omitempty" yaml:"parameters,omitempty"`

	// The ARN of the application, configuration profile, or environment to associate with the extension.
	ResourceArn string `json:"resourceArn,omitempty" yaml:"resourceArn,omitempty"`
}
