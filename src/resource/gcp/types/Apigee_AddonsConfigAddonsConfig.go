package types

type Apigee_AddonsConfigAddonsConfig struct {
	/*
	   Configuration for the Monetization add-on.
	   Structure is documented below.
	*/
	AdvancedApiOpsConfig Apigee_AddonsConfigAddonsConfigAdvancedApiOpsConfig `json:"advancedApiOpsConfig,omitempty" yaml:"advancedApiOpsConfig,omitempty"`

	/*
	   Configuration for the Monetization add-on.
	   Structure is documented below.
	*/
	ApiSecurityConfig Apigee_AddonsConfigAddonsConfigApiSecurityConfig `json:"apiSecurityConfig,omitempty" yaml:"apiSecurityConfig,omitempty"`

	/*
	   Configuration for the Monetization add-on.
	   Structure is documented below.
	*/
	ConnectorsPlatformConfig Apigee_AddonsConfigAddonsConfigConnectorsPlatformConfig `json:"connectorsPlatformConfig,omitempty" yaml:"connectorsPlatformConfig,omitempty"`

	/*
	   Configuration for the Monetization add-on.
	   Structure is documented below.
	*/
	IntegrationConfig Apigee_AddonsConfigAddonsConfigIntegrationConfig `json:"integrationConfig,omitempty" yaml:"integrationConfig,omitempty"`

	/*
	   Configuration for the Monetization add-on.
	   Structure is documented below.
	*/
	MonetizationConfig Apigee_AddonsConfigAddonsConfigMonetizationConfig `json:"monetizationConfig,omitempty" yaml:"monetizationConfig,omitempty"`
}
