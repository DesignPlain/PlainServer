package types

type Connect_UserPhoneConfig struct {
	// The After Call Work (ACW) timeout setting, in seconds. Minimum value of 0.
	AfterContactWorkTimeLimit int `json:"afterContactWorkTimeLimit,omitempty" yaml:"afterContactWorkTimeLimit,omitempty"`

	// When Auto-Accept Call is enabled for an available agent, the agent connects to contacts automatically.
	AutoAccept bool `json:"autoAccept,omitempty" yaml:"autoAccept,omitempty"`

	// The phone number for the user's desk phone. Required if `phone_type` is set as `DESK_PHONE`.
	DeskPhoneNumber string `json:"deskPhoneNumber,omitempty" yaml:"deskPhoneNumber,omitempty"`

	// The phone type. Valid values are `DESK_PHONE` and `SOFT_PHONE`.
	PhoneType string `json:"phoneType,omitempty" yaml:"phoneType,omitempty"`
}
