package types

type Integrationconnectors_ConnectionEventingConfig struct {
	/*
	   registrationDestinationConfig
	   Structure is documented below.
	*/
	RegistrationDestinationConfig Integrationconnectors_ConnectionEventingConfigRegistrationDestinationConfig `json:"registrationDestinationConfig,omitempty" yaml:"registrationDestinationConfig,omitempty"`

	/*
	   List containing additional auth configs.
	   Structure is documented below.
	*/
	AdditionalVariables []Integrationconnectors_ConnectionEventingConfigAdditionalVariable `json:"additionalVariables,omitempty" yaml:"additionalVariables,omitempty"`

	/*
	   authConfig for Eventing Configuration.
	   Structure is documented below.
	*/
	AuthConfig Integrationconnectors_ConnectionEventingConfigAuthConfig `json:"authConfig,omitempty" yaml:"authConfig,omitempty"`

	// Enrichment Enabled.
	EnrichmentEnabled bool `json:"enrichmentEnabled,omitempty" yaml:"enrichmentEnabled,omitempty"`
}
