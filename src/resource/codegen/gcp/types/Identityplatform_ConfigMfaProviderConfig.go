package types

type Identityplatform_ConfigMfaProviderConfig struct {
	/*
	   Whether MultiFactor Authentication has been enabled for this project.
	   Possible values are: `DISABLED`, `ENABLED`, `MANDATORY`.
	*/
	State string `json:"state,omitempty" yaml:"state,omitempty"`

	/*
	   TOTP MFA provider config for this project.
	   Structure is documented below.
	*/
	TotpProviderConfig Identityplatform_ConfigMfaProviderConfigTotpProviderConfig `json:"totpProviderConfig,omitempty" yaml:"totpProviderConfig,omitempty"`
}
