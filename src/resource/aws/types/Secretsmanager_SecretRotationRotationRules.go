package types

type Secretsmanager_SecretRotationRotationRules struct {
	// Specifies the number of days between automatic scheduled rotations of the secret. Either `automatically_after_days` or `schedule_expression` must be specified.
	AutomaticallyAfterDays int `json:"automaticallyAfterDays,omitempty" yaml:"automaticallyAfterDays,omitempty"`

	// The length of the rotation window in hours. For example, `3h` for a three hour window.
	Duration string `json:"duration,omitempty" yaml:"duration,omitempty"`

	// A `cron()` or `rate()` expression that defines the schedule for rotating your secret. Either `automatically_after_days` or `schedule_expression` must be specified.
	ScheduleExpression string `json:"scheduleExpression,omitempty" yaml:"scheduleExpression,omitempty"`
}
