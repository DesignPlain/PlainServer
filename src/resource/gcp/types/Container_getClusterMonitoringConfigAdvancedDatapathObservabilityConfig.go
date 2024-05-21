package types

type Container_getClusterMonitoringConfigAdvancedDatapathObservabilityConfig struct {
	// Whether or not the advanced datapath metrics are enabled.
	EnableMetrics bool `json:"enableMetrics,omitempty" yaml:"enableMetrics,omitempty"`

	// Whether or not Relay is enabled.
	EnableRelay bool `json:"enableRelay,omitempty" yaml:"enableRelay,omitempty"`

	// Mode used to make Relay available.
	RelayMode string `json:"relayMode,omitempty" yaml:"relayMode,omitempty"`
}
