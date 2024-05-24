package types

type Ssm_ContactsRotationRecurrenceShiftCoverage struct {
	// (Required) Information about when an on-call shift begins and ends. See Coverage Times for more details.
	CoverageTimes []Ssm_ContactsRotationRecurrenceShiftCoverageCoverageTime `json:"coverageTimes,omitempty" yaml:"coverageTimes,omitempty"`

	//
	MapBlockKey string `json:"mapBlockKey,omitempty" yaml:"mapBlockKey,omitempty"`
}
