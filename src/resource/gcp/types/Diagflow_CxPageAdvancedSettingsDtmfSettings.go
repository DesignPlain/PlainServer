package types

type Diagflow_CxPageAdvancedSettingsDtmfSettings struct {
	// If true, incoming audio is processed for DTMF (dual tone multi frequency) events. For example, if the caller presses a button on their telephone keypad and DTMF processing is enabled, Dialogflow will detect the event (e.g. a "3" was pressed) in the incoming audio and pass the event to the bot to drive business logic (e.g. when 3 is pressed, return the account balance).
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// The digit that terminates a DTMF digit sequence.
	FinishDigit string `json:"finishDigit,omitempty" yaml:"finishDigit,omitempty"`

	// Max length of DTMF digits.
	MaxDigits int `json:"maxDigits,omitempty" yaml:"maxDigits,omitempty"`
}
