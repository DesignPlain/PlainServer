package identityplatform

import types "libds/gcp/types"

type TenantInboundSamlConfig struct {
	// The name of the tenant where this inbound SAML config resource exists
	Tenant string `json:"tenant,omitempty" yaml:"tenant,omitempty"`

	// Human friendly display name.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	// If this config allows users to sign in with the provider.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	/*
	   SAML IdP configuration when the project acts as the relying party
	   Structure is documented below.
	*/
	IdpConfig types.Identityplatform_TenantInboundSamlConfigIdpConfig `json:"idpConfig,omitempty" yaml:"idpConfig,omitempty"`

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

	/*
	   SAML SP (Service Provider) configuration when the project acts as the relying party to receive
	   and accept an authentication assertion issued by a SAML identity provider.
	   Structure is documented below.
	*/
	SpConfig types.Identityplatform_TenantInboundSamlConfigSpConfig `json:"spConfig,omitempty" yaml:"spConfig,omitempty"`
}
