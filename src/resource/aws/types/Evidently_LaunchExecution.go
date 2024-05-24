package types

type Evidently_LaunchExecution struct {
	// The date and time that the launch ended.
	EndedTime string `json:"endedTime,omitempty" yaml:"endedTime,omitempty"`

	// The date and time that the launch started.
	StartedTime string `json:"startedTime,omitempty" yaml:"startedTime,omitempty"`
}
