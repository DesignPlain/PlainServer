package types

type Cloudscheduler_JobRetryConfig struct {
	/*
	   The maximum amount of time to wait before retrying a job after it fails.
	   A duration in seconds with up to nine fractional digits, terminated by 's'.
	*/
	MaxBackoffDuration string `json:"maxBackoffDuration,omitempty" yaml:"maxBackoffDuration,omitempty"`

	/*
	   The time between retries will double maxDoublings times.
	   A job's retry interval starts at minBackoffDuration,
	   then doubles maxDoublings times, then increases linearly,
	   and finally retries retries at intervals of maxBackoffDuration up to retryCount times.
	*/
	MaxDoublings int `json:"maxDoublings,omitempty" yaml:"maxDoublings,omitempty"`

	/*
	   The time limit for retrying a failed job, measured from time when an execution was first attempted.
	   If specified with retryCount, the job will be retried until both limits are reached.
	   A duration in seconds with up to nine fractional digits, terminated by 's'.
	*/
	MaxRetryDuration string `json:"maxRetryDuration,omitempty" yaml:"maxRetryDuration,omitempty"`

	/*
	   The minimum amount of time to wait before retrying a job after it fails.
	   A duration in seconds with up to nine fractional digits, terminated by 's'.
	*/
	MinBackoffDuration string `json:"minBackoffDuration,omitempty" yaml:"minBackoffDuration,omitempty"`

	/*
	   The number of attempts that the system will make to run a
	   job using the exponential backoff procedure described by maxDoublings.
	   Values greater than 5 and negative values are not allowed.
	*/
	RetryCount int `json:"retryCount,omitempty" yaml:"retryCount,omitempty"`
}
