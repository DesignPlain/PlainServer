package types

type Kinesisanalyticsv2_ApplicationApplicationConfigurationFlinkApplicationConfigurationParallelismConfiguration struct {
	// Describes the number of parallel tasks that a Flink-based Kinesis Data Analytics application can perform per Kinesis Processing Unit (KPU) used by the application.
	ParallelismPerKpu int `json:"parallelismPerKpu,omitempty" yaml:"parallelismPerKpu,omitempty"`

	// Describes whether the Kinesis Data Analytics service can increase the parallelism of the application in response to increased throughput.
	AutoScalingEnabled bool `json:"autoScalingEnabled,omitempty" yaml:"autoScalingEnabled,omitempty"`

	// Describes whether the application uses the default parallelism for the Kinesis Data Analytics service. Valid values: `CUSTOM`, `DEFAULT`. Set this attribute to `CUSTOM` in order for any specified `auto_scaling_enabled`, `parallelism`, or `parallelism_per_kpu` attribute values to be effective.
	ConfigurationType string `json:"configurationType,omitempty" yaml:"configurationType,omitempty"`

	// Describes the initial number of parallel tasks that a Flink-based Kinesis Data Analytics application can perform.
	Parallelism int `json:"parallelism,omitempty" yaml:"parallelism,omitempty"`
}
