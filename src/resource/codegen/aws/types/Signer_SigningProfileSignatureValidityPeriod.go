package types

type Signer_SigningProfileSignatureValidityPeriod struct {
	// The time unit for signature validity. Valid values: `DAYS`, `MONTHS`, `YEARS`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// The numerical value of the time unit for signature validity.
	Value int `json:"value,omitempty" yaml:"value,omitempty"`
}
