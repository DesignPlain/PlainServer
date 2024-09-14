package types

type Pipes_PipeTargetParametersCloudwatchLogsParameters struct {
	// The name of the log stream.
	LogStreamName string `json:"logStreamName,omitempty" yaml:"logStreamName,omitempty"`

	// The time the event occurred, expressed as the number of milliseconds after Jan 1, 1970 00:00:00 UTC. This is the JSON path to the field in the event e.g. $.detail.timestamp
	Timestamp string `json:"timestamp,omitempty" yaml:"timestamp,omitempty"`
}
