package types

type Scheduler_ScheduleTargetRetryPolicy struct {
	// Maximum amount of time, in seconds, to continue to make retry attempts. Ranges from `60` to `86400` (default).
	MaximumEventAgeInSeconds int `json:"maximumEventAgeInSeconds,omitempty" yaml:"maximumEventAgeInSeconds,omitempty"`

	// Maximum number of retry attempts to make before the request fails. Ranges from `0` to `185` (default).
	MaximumRetryAttempts int `json:"maximumRetryAttempts,omitempty" yaml:"maximumRetryAttempts,omitempty"`
}
