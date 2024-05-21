package cloudtasks

import types "DesignSphere_Server/src/resource/gcp/types"

type Queue struct {
	/*
	   Settings that determine the retry behavior.
	   Structure is documented below.
	*/
	RetryConfig types.Cloudtasks_QueueRetryConfig `json:"retryConfig,omitempty" yaml:"retryConfig,omitempty"`

	/*
	   Configuration options for writing logs to Stackdriver Logging.
	   Structure is documented below.
	*/
	StackdriverLoggingConfig types.Cloudtasks_QueueStackdriverLoggingConfig `json:"stackdriverLoggingConfig,omitempty" yaml:"stackdriverLoggingConfig,omitempty"`

	/*
	   Overrides for task-level appEngineRouting. These settings apply only
	   to App Engine tasks in this queue
	   Structure is documented below.
	*/
	AppEngineRoutingOverride types.Cloudtasks_QueueAppEngineRoutingOverride `json:"appEngineRoutingOverride,omitempty" yaml:"appEngineRoutingOverride,omitempty"`

	/*
	   The location of the queue


	   - - -
	*/
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// The queue name.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Rate limits for task dispatches.
	   The queue's actual dispatch rate is the result of:
	   - Number of tasks in the queue
	   - User-specified throttling: rateLimits, retryConfig, and the queue's state.
	   - System throttling due to 429 (Too Many Requests) or 503 (Service
	   Unavailable) responses from the worker, high error rates, or to
	   smooth sudden large traffic spikes.
	   Structure is documented below.
	*/
	RateLimits types.Cloudtasks_QueueRateLimits `json:"rateLimits,omitempty" yaml:"rateLimits,omitempty"`
}
