package types

type Worklink_FleetIdentityProvider struct {
	// The SAML metadata document provided by the customer’s identity provider.
	SamlMetadata string `json:"samlMetadata,omitempty" yaml:"samlMetadata,omitempty"`

	// The type of identity provider.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
