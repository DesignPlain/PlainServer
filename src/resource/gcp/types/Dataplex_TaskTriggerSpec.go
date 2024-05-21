package types

type Dataplex_TaskTriggerSpec struct {
	/*
	   Trigger type of the user-specified Task
	   Possible values are: `ON_DEMAND`, `RECURRING`.
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Prevent the task from executing. This does not cancel already running tasks. It is intended to temporarily disable RECURRING tasks.
	Disabled bool `json:"disabled,omitempty" yaml:"disabled,omitempty"`

	// Number of retry attempts before aborting. Set to zero to never attempt to retry a failed task.
	MaxRetries int `json:"maxRetries,omitempty" yaml:"maxRetries,omitempty"`

	// Cron schedule (https://en.wikipedia.org/wiki/Cron) for running tasks periodically. To explicitly set a timezone to the cron tab, apply a prefix in the cron tab: 'CRON_TZ=${IANA_TIME_ZONE}' or 'TZ=${IANA_TIME_ZONE}'. The ${IANA_TIME_ZONE} may only be a valid string from IANA time zone database. For example, CRON_TZ=America/New_York 1 - - - -, or TZ=America/New_York 1 - - - -. This field is required for RECURRING tasks.
	Schedule string `json:"schedule,omitempty" yaml:"schedule,omitempty"`

	// The first run of the task will be after this time. If not specified, the task will run shortly after being submitted if ON_DEMAND and based on the schedule if RECURRING.
	StartTime string `json:"startTime,omitempty" yaml:"startTime,omitempty"`
}
