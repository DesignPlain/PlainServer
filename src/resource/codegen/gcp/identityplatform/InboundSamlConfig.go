package identityplatform

import types "libds/gcp/types"

type InboundSamlConfig struct {
	/*
	   SAML SP (Service Provider) configuration when the project acts as the relying party to receive
	   and accept an authentication assertion issued by a SAML identity provider.
	   Structure is documented below.
	*/
	SpConfig types.Identityplatform_InboundSamlConfigSpConfig `json:"spConfig,omitempty" yaml:"spConfig,omitempty"`

	// Human friendly display name.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	// If this config allows users to sign in with the provider.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	/*
	   SAML IdP configuration when the project acts as the relying party
	   Structure is documented below.
	*/
	IdpConfig types.Identityplatform_InboundSamlConfigIdpConfig `json:"idpConfig,omitempty" yaml:"idpConfig,omitempty"`

	/*
	   The name of the InboundSamlConfig resource. Must start with 'saml.' and can only have alphanumeric characters,
	   hyphens, underscores or periods. The part after 'saml.' must also start with a lowercase letter, end with an
	   alphanumeric character, and have at least 2 characters.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
