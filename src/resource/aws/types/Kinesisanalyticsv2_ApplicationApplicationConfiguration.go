package types

type Kinesisanalyticsv2_ApplicationApplicationConfiguration struct {
	// The code location and type parameters for the application.
	ApplicationCodeConfiguration Kinesisanalyticsv2_ApplicationApplicationConfigurationApplicationCodeConfiguration `json:"applicationCodeConfiguration,omitempty" yaml:"applicationCodeConfiguration,omitempty"`

	// Describes whether snapshots are enabled for a Flink-based application.
	ApplicationSnapshotConfiguration Kinesisanalyticsv2_ApplicationApplicationConfigurationApplicationSnapshotConfiguration `json:"applicationSnapshotConfiguration,omitempty" yaml:"applicationSnapshotConfiguration,omitempty"`

	// Describes execution properties for a Flink-based application.
	EnvironmentProperties Kinesisanalyticsv2_ApplicationApplicationConfigurationEnvironmentProperties `json:"environmentProperties,omitempty" yaml:"environmentProperties,omitempty"`

	// The configuration of a Flink-based application.
	FlinkApplicationConfiguration Kinesisanalyticsv2_ApplicationApplicationConfigurationFlinkApplicationConfiguration `json:"flinkApplicationConfiguration,omitempty" yaml:"flinkApplicationConfiguration,omitempty"`

	// Describes the starting properties for a Flink-based application.
	RunConfiguration Kinesisanalyticsv2_ApplicationApplicationConfigurationRunConfiguration `json:"runConfiguration,omitempty" yaml:"runConfiguration,omitempty"`

	// The configuration of a SQL-based application.
	SqlApplicationConfiguration Kinesisanalyticsv2_ApplicationApplicationConfigurationSqlApplicationConfiguration `json:"sqlApplicationConfiguration,omitempty" yaml:"sqlApplicationConfiguration,omitempty"`

	// The VPC configuration of a Flink-based application.
	VpcConfiguration Kinesisanalyticsv2_ApplicationApplicationConfigurationVpcConfiguration `json:"vpcConfiguration,omitempty" yaml:"vpcConfiguration,omitempty"`
}
