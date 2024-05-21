package types

type Pubsub_getSubscriptionRetryPolicy struct {
	/*
	   The maximum delay between consecutive deliveries of a given message. Value should be between 0 and 600 seconds. Defaults to 600 seconds.
	   A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
	*/
	MaximumBackoff string `json:"maximumBackoff,omitempty" yaml:"maximumBackoff,omitempty"`

	/*
	   The minimum delay between consecutive deliveries of a given message. Value should be between 0 and 600 seconds. Defaults to 10 seconds.
	   A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
	*/
	MinimumBackoff string `json:"minimumBackoff,omitempty" yaml:"minimumBackoff,omitempty"`
}
