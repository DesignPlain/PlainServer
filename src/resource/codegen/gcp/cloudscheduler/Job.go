package cloudscheduler

import types "libds/gcp/types"

type Job struct {
	/*
	   App Engine HTTP target.
	   If the job providers a App Engine HTTP target the cron will
	   send a request to the service instance
	   Structure is documented below.
	*/
	AppEngineHttpTarget types.Cloudscheduler_JobAppEngineHttpTarget `json:"appEngineHttpTarget,omitempty" yaml:"appEngineHttpTarget,omitempty"`

	/*
	   A human-readable description for the job.
	   This string must not contain more than 500 characters.
	*/
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   The name of the job.


	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   By default, if a job does not complete successfully,
	   meaning that an acknowledgement is not received from the handler,
	   then it will be retried with exponential backoff according to the settings
	   Structure is documented below.
	*/
	RetryConfig types.Cloudscheduler_JobRetryConfig `json:"retryConfig,omitempty" yaml:"retryConfig,omitempty"`

	/*
	   Specifies the time zone to be used in interpreting schedule.
	   The value of this field must be a time zone name from the tz database.
	*/
	TimeZone string `json:"timeZone,omitempty" yaml:"timeZone,omitempty"`

	/*
	   The deadline for job attempts. If the request handler does not respond by this deadline then the request is
	   cancelled and the attempt is marked as a DEADLINE_EXCEEDED failure. The failed attempt can be viewed in
	   execution logs. Cloud Scheduler will retry the job according to the RetryConfig.
	   The allowed duration for this deadline is:
	   - For HTTP targets, between 15 seconds and 30 minutes.
	   - For App Engine HTTP targets, between 15 seconds and 24 hours.
	   - --Note--: For PubSub targets, this field is ignored - setting it will introduce an unresolvable diff.
	   A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s"
	*/
	AttemptDeadline string `json:"attemptDeadline,omitempty" yaml:"attemptDeadline,omitempty"`

	/*
	   HTTP target.
	   If the job providers a http_target the cron will
	   send a request to the targeted url
	   Structure is documented below.
	*/
	HttpTarget types.Cloudscheduler_JobHttpTarget `json:"httpTarget,omitempty" yaml:"httpTarget,omitempty"`

	// Sets the job to a paused state. Jobs default to being enabled when this property is not set.
	Paused bool `json:"paused,omitempty" yaml:"paused,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Pub/Sub target
	   If the job providers a Pub/Sub target the cron will publish
	   a message to the provided topic
	   Structure is documented below.
	*/
	PubsubTarget types.Cloudscheduler_JobPubsubTarget `json:"pubsubTarget,omitempty" yaml:"pubsubTarget,omitempty"`

	// Region where the scheduler job resides. If it is not provided, this provider will use the provider default.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	// Describes the schedule on which the job will be executed.
	Schedule string `json:"schedule,omitempty" yaml:"schedule,omitempty"`
}
