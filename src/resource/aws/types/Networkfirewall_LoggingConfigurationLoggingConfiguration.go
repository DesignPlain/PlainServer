package types

type Networkfirewall_LoggingConfigurationLoggingConfiguration struct {
	// Set of configuration blocks describing the logging details for a firewall. See Log Destination Config below for details. At most, only two blocks can be specified; one for `FLOW` logs and one for `ALERT` logs.
	LogDestinationConfigs []Networkfirewall_LoggingConfigurationLoggingConfigurationLogDestinationConfig `json:"logDestinationConfigs,omitempty" yaml:"logDestinationConfigs,omitempty"`
}
