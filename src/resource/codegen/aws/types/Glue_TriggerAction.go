package types

type Glue_TriggerAction struct {
	// Arguments to be passed to the job. You can specify arguments here that your own job-execution script consumes, as well as arguments that AWS Glue itself consumes.
	Arguments map[string]string `json:"arguments,omitempty" yaml:"arguments,omitempty"`

	// The name of the crawler to be executed. Conflicts with `job_name`.
	CrawlerName string `json:"crawlerName,omitempty" yaml:"crawlerName,omitempty"`

	// The name of a job to be executed. Conflicts with `crawler_name`.
	JobName string `json:"jobName,omitempty" yaml:"jobName,omitempty"`

	// Specifies configuration properties of a job run notification. See Notification Property details below.
	NotificationProperty Glue_TriggerActionNotificationProperty `json:"notificationProperty,omitempty" yaml:"notificationProperty,omitempty"`

	// The name of the Security Configuration structure to be used with this action.
	SecurityConfiguration string `json:"securityConfiguration,omitempty" yaml:"securityConfiguration,omitempty"`

	// The job run timeout in minutes. It overrides the timeout value of the job.
	Timeout int `json:"timeout,omitempty" yaml:"timeout,omitempty"`
}
