package types

type Cloudwatch_EventTargetRetryPolicy struct {
	// maximum number of retry attempts to make before the request fails
	MaximumRetryAttempts int `json:"maximumRetryAttempts,omitempty" yaml:"maximumRetryAttempts,omitempty"`

	// The age in seconds to continue to make retry attempts.
	MaximumEventAgeInSeconds int `json:"maximumEventAgeInSeconds,omitempty" yaml:"maximumEventAgeInSeconds,omitempty"`
}
