package types

type Opensearch_ServerlessSecurityConfigSamlOptions struct {
	// Group attribute for this SAML integration.
	GroupAttribute string `json:"groupAttribute,omitempty" yaml:"groupAttribute,omitempty"`

	// The XML IdP metadata file generated from your identity provider.
	Metadata string `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	// Session timeout, in minutes. Minimum is 5 minutes and maximum is 720 minutes (12 hours). Default is 60 minutes.
	SessionTimeout int `json:"sessionTimeout,omitempty" yaml:"sessionTimeout,omitempty"`

	// User attribute for this SAML integration.
	UserAttribute string `json:"userAttribute,omitempty" yaml:"userAttribute,omitempty"`
}
