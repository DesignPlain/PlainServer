package cfg

type RetentionConfiguration struct {
	// The number of days AWS Config stores historical information.
	RetentionPeriodInDays int `json:"retentionPeriodInDays,omitempty" yaml:"retentionPeriodInDays,omitempty"`
}
