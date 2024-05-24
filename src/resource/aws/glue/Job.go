package glue

import types "DesignSphere_Server/src/resource/aws/types"

type Job struct {
	// Indicates whether the job is run with a standard or flexible execution class. The standard execution class is ideal for time-sensitive workloads that require fast job startup and dedicated resources. Valid value: `FLEX`, `STANDARD`.
	ExecutionClass string `json:"executionClass,omitempty" yaml:"executionClass,omitempty"`

	// The maximum number of times to retry this job if it fails.
	MaxRetries int `json:"maxRetries,omitempty" yaml:"maxRetries,omitempty"`

	// Non-overridable arguments for this job, specified as name-value pairs.
	NonOverridableArguments map[string]string `json:"nonOverridableArguments,omitempty" yaml:"nonOverridableArguments,omitempty"`

	// The name you assign to this job. It must be unique in your account.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The ARN of the IAM role associated with this job.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The job timeout in minutes. The default is 2880 minutes (48 hours) for `glueetl` and `pythonshell` jobs, and null (unlimited) for `gluestreaming` jobs.
	Timeout int `json:"timeout,omitempty" yaml:"timeout,omitempty"`

	// The command of the job. Defined below.
	Command types.Glue_JobCommand `json:"command,omitempty" yaml:"command,omitempty"`

	// The list of connections used for this job.
	Connections []string `json:"connections,omitempty" yaml:"connections,omitempty"`

	// The maximum number of AWS Glue data processing units (DPUs) that can be allocated when this job runs. `Required` when `pythonshell` is set, accept either `0.0625` or `1.0`. Use `number_of_workers` and `worker_type` arguments instead with `glue_version` `2.0` and above.
	MaxCapacity float64 `json:"maxCapacity,omitempty" yaml:"maxCapacity,omitempty"`

	// The version of glue to use, for example "1.0". Ray jobs should set this to 4.0 or greater. For information about available versions, see the [AWS Glue Release Notes](https://docs.aws.amazon.com/glue/latest/dg/release-notes.html).
	GlueVersion string `json:"glueVersion,omitempty" yaml:"glueVersion,omitempty"`

	// The name of the Security Configuration to be associated with the job.
	SecurityConfiguration string `json:"securityConfiguration,omitempty" yaml:"securityConfiguration,omitempty"`

	/*
	   The type of predefined worker that is allocated when a job runs. Accepts a value of Standard, G.1X, G.2X, or G.025X for Spark jobs. Accepts the value Z.2X for Ray jobs.
	   - For the Standard worker type, each worker provides 4 vCPU, 16 GB of memory and a 50GB disk, and 2 executors per worker.
	   - For the G.1X worker type, each worker maps to 1 DPU (4 vCPU, 16 GB of memory, 64 GB disk), and provides 1 executor per worker. Recommended for memory-intensive jobs.
	   - For the G.2X worker type, each worker maps to 2 DPU (8 vCPU, 32 GB of memory, 128 GB disk), and provides 1 executor per worker. Recommended for memory-intensive jobs.
	   - For the G.025X worker type, each worker maps to 0.25 DPU (2 vCPU, 4GB of memory, 64 GB disk), and provides 1 executor per worker. Recommended for low volume streaming jobs. Only available for Glue version 3.0.
	   - For the Z.2X worker type, each worker maps to 2 M-DPU (8vCPU, 64 GB of m emory, 128 GB disk), and provides up to 8 Ray workers based on the autoscaler.
	*/
	WorkerType string `json:"workerType,omitempty" yaml:"workerType,omitempty"`

	// The map of default arguments for this job. You can specify arguments here that your own job-execution script consumes, as well as arguments that AWS Glue itself consumes. For information about how to specify and consume your own Job arguments, see the [Calling AWS Glue APIs in Python](http://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-calling.html) topic in the developer guide. For information about the key-value pairs that AWS Glue consumes to set up your job, see the [Special Parameters Used by AWS Glue](http://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-glue-arguments.html) topic in the developer guide.
	DefaultArguments map[string]string `json:"defaultArguments,omitempty" yaml:"defaultArguments,omitempty"`

	// Description of the job.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Execution property of the job. Defined below.
	ExecutionProperty types.Glue_JobExecutionProperty `json:"executionProperty,omitempty" yaml:"executionProperty,omitempty"`

	// Notification property of the job. Defined below.
	NotificationProperty types.Glue_JobNotificationProperty `json:"notificationProperty,omitempty" yaml:"notificationProperty,omitempty"`

	// The number of workers of a defined workerType that are allocated when a job runs.
	NumberOfWorkers int `json:"numberOfWorkers,omitempty" yaml:"numberOfWorkers,omitempty"`
}
