package types

type Lex_V2modelsIntentOutputContext struct {
	// Number of conversation turns that the output context should remain active. The number of turns is counted from the first time that the context is sent to the user.
	TurnsToLive int `json:"turnsToLive,omitempty" yaml:"turnsToLive,omitempty"`

	// Name of the output context.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Amount of time, in seconds, that the output context should remain active. The time is figured from the first time the context is sent to the user.
	TimeToLiveInSeconds int `json:"timeToLiveInSeconds,omitempty" yaml:"timeToLiveInSeconds,omitempty"`
}
