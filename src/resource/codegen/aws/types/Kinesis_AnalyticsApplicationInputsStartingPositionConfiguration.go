package types

type Kinesis_AnalyticsApplicationInputsStartingPositionConfiguration struct {
	// The starting position on the stream. Valid values: `LAST_STOPPED_POINT`, `NOW`, `TRIM_HORIZON`.
	StartingPosition string `json:"startingPosition,omitempty" yaml:"startingPosition,omitempty"`
}
