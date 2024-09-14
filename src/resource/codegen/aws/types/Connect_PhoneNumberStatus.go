package types

type Connect_PhoneNumberStatus struct {
	// The status message.
	Message string `json:"message,omitempty" yaml:"message,omitempty"`

	// The status of the phone number. Valid Values: `CLAIMED` | `IN_PROGRESS` | `FAILED`.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`
}
