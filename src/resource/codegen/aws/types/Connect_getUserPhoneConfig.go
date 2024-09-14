package types

type Connect_getUserPhoneConfig struct {
	// When Auto-Accept Call is enabled for an available agent, the agent connects to contacts automatically.
	AutoAccept bool `json:"autoAccept,omitempty" yaml:"autoAccept,omitempty"`

	// The phone number for the user's desk phone.
	DeskPhoneNumber string `json:"deskPhoneNumber,omitempty" yaml:"deskPhoneNumber,omitempty"`

	// The phone type. Valid values are `DESK_PHONE` and `SOFT_PHONE`.
	PhoneType string `json:"phoneType,omitempty" yaml:"phoneType,omitempty"`

	// The After Call Work (ACW) timeout setting, in seconds.
	AfterContactWorkTimeLimit int `json:"afterContactWorkTimeLimit,omitempty" yaml:"afterContactWorkTimeLimit,omitempty"`
}
