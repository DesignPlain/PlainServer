package types

type Kinesisanalyticsv2_ApplicationApplicationConfigurationRunConfiguration struct {
	// The restore behavior of a restarting application.
	ApplicationRestoreConfiguration Kinesisanalyticsv2_ApplicationApplicationConfigurationRunConfigurationApplicationRestoreConfiguration `json:"applicationRestoreConfiguration,omitempty" yaml:"applicationRestoreConfiguration,omitempty"`

	// The starting parameters for a Flink-based Kinesis Data Analytics application.
	FlinkRunConfiguration Kinesisanalyticsv2_ApplicationApplicationConfigurationRunConfigurationFlinkRunConfiguration `json:"flinkRunConfiguration,omitempty" yaml:"flinkRunConfiguration,omitempty"`
}
