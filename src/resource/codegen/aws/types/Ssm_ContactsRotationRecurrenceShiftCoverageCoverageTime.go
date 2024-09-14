package types

type Ssm_ContactsRotationRecurrenceShiftCoverageCoverageTime struct {
	// (Required) The end time of the on-call shift. See Hand Off Time for more details.
	End Ssm_ContactsRotationRecurrenceShiftCoverageCoverageTimeEnd `json:"end,omitempty" yaml:"end,omitempty"`

	// (Required) The start time of the on-call shift. See Hand Off Time for more details.
	Start Ssm_ContactsRotationRecurrenceShiftCoverageCoverageTimeStart `json:"start,omitempty" yaml:"start,omitempty"`
}
