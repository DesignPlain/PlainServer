package types

type Kinesisanalyticsv2_ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputStartingPositionConfiguration struct {
	// The starting position on the stream. Valid values: `LAST_STOPPED_POINT`, `NOW`, `TRIM_HORIZON`.
	InputStartingPosition string `json:"inputStartingPosition,omitempty" yaml:"inputStartingPosition,omitempty"`
}
