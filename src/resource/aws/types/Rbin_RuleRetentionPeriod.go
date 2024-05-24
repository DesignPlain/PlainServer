package types

type Rbin_RuleRetentionPeriod struct {
	// The period value for which the retention rule is to retain resources. The period is measured using the unit specified for RetentionPeriodUnit.
	RetentionPeriodValue int `json:"retentionPeriodValue,omitempty" yaml:"retentionPeriodValue,omitempty"`

	// The unit of time in which the retention period is measured. Currently, only DAYS is supported.
	RetentionPeriodUnit string `json:"retentionPeriodUnit,omitempty" yaml:"retentionPeriodUnit,omitempty"`
}
