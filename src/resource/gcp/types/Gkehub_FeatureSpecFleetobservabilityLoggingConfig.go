package types

type Gkehub_FeatureSpecFleetobservabilityLoggingConfig struct {
	/*
	   Specified if applying the routing config to all logs for all fleet scopes.
	   Structure is documented below.
	*/
	FleetScopeLogsConfig Gkehub_FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig `json:"fleetScopeLogsConfig,omitempty" yaml:"fleetScopeLogsConfig,omitempty"`

	/*
	   Specified if applying the default routing config to logs not specified in other configs.
	   Structure is documented below.
	*/
	DefaultConfig Gkehub_FeatureSpecFleetobservabilityLoggingConfigDefaultConfig `json:"defaultConfig,omitempty" yaml:"defaultConfig,omitempty"`
}
