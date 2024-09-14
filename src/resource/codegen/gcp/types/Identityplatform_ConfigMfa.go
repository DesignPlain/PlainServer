package types

type Identityplatform_ConfigMfa struct {
	/*
	   A list of usable second factors for this project.
	   Each value may be one of: `PHONE_SMS`.
	*/
	EnabledProviders []string `json:"enabledProviders,omitempty" yaml:"enabledProviders,omitempty"`

	/*
	   A list of usable second factors for this project along with their configurations.
	   This field does not support phone based MFA, for that use the 'enabledProviders' field.
	   Structure is documented below.
	*/
	ProviderConfigs []Identityplatform_ConfigMfaProviderConfig `json:"providerConfigs,omitempty" yaml:"providerConfigs,omitempty"`

	/*
	   Whether MultiFactor Authentication has been enabled for this project.
	   Possible values are: `DISABLED`, `ENABLED`, `MANDATORY`.
	*/
	State string `json:"state,omitempty" yaml:"state,omitempty"`
}
