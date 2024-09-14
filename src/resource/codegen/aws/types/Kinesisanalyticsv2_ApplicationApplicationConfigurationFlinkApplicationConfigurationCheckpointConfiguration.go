package types

type Kinesisanalyticsv2_ApplicationApplicationConfigurationFlinkApplicationConfigurationCheckpointConfiguration struct {
	// Describes the minimum time in milliseconds after a checkpoint operation completes that a new checkpoint operation can start.
	MinPauseBetweenCheckpoints int `json:"minPauseBetweenCheckpoints,omitempty" yaml:"minPauseBetweenCheckpoints,omitempty"`

	// Describes the interval in milliseconds between checkpoint operations.
	CheckpointInterval int `json:"checkpointInterval,omitempty" yaml:"checkpointInterval,omitempty"`

	// Describes whether checkpointing is enabled for a Flink-based Kinesis Data Analytics application.
	CheckpointingEnabled bool `json:"checkpointingEnabled,omitempty" yaml:"checkpointingEnabled,omitempty"`

	/*
	   Describes whether the application uses Kinesis Data Analytics' default checkpointing behavior. Valid values: `CUSTOM`, `DEFAULT`. Set this attribute to `CUSTOM` in order for any specified `checkpointing_enabled`, `checkpoint_interval`, or `min_pause_between_checkpoints` attribute values to be effective. If this attribute is set to `DEFAULT`, the application will always use the following values:
	   - `checkpointing_enabled = true`
	   - `checkpoint_interval = 60000`
	   - `min_pause_between_checkpoints = 5000`
	*/
	ConfigurationType string `json:"configurationType,omitempty" yaml:"configurationType,omitempty"`
}
