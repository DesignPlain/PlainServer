package batch

import types "libds/aws/types"

type JobDefinition struct {
	/*
	   Type of job definition. Must be `container` or `multinode`.

	   The following arguments are optional:
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Valid [container properties](http://docs.aws.amazon.com/batch/latest/APIReference/API_RegisterJobDefinition.html) provided as a single valid JSON document. This parameter is only valid if the `type` parameter is `container`.
	ContainerProperties string `json:"containerProperties,omitempty" yaml:"containerProperties,omitempty"`

	// Valid eks properties. This parameter is only valid if the `type` parameter is `container`.
	EksProperties types.Batch_JobDefinitionEksProperties `json:"eksProperties,omitempty" yaml:"eksProperties,omitempty"`

	// Parameter substitution placeholders to set in the job definition.
	Parameters map[string]string `json:"parameters,omitempty" yaml:"parameters,omitempty"`

	// Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Timeout for jobs so that if a job runs longer, AWS Batch terminates the job. Maximum number of `timeout` is `1`. Defined below.
	Timeout types.Batch_JobDefinitionTimeout `json:"timeout,omitempty" yaml:"timeout,omitempty"`

	// Valid [node properties](http://docs.aws.amazon.com/batch/latest/APIReference/API_RegisterJobDefinition.html) provided as a single valid JSON document. This parameter is required if the `type` parameter is `multinode`.
	NodeProperties string `json:"nodeProperties,omitempty" yaml:"nodeProperties,omitempty"`

	// Platform capabilities required by the job definition. If no value is specified, it defaults to `EC2`. To run the job on Fargate resources, specify `FARGATE`.
	PlatformCapabilities []string `json:"platformCapabilities,omitempty" yaml:"platformCapabilities,omitempty"`

	// Scheduling priority of the job definition. This only affects jobs in job queues with a fair share policy. Jobs with a higher scheduling priority are scheduled before jobs with a lower scheduling priority. Allowed values `0` through `9999`.
	SchedulingPriority int `json:"schedulingPriority,omitempty" yaml:"schedulingPriority,omitempty"`

	// Name of the job definition.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Whether to propagate the tags from the job definition to the corresponding Amazon ECS task. Default is `false`.
	PropagateTags bool `json:"propagateTags,omitempty" yaml:"propagateTags,omitempty"`

	// Retry strategy to use for failed jobs that are submitted with this job definition. Maximum number of `retry_strategy` is `1`.  Defined below.
	RetryStrategy types.Batch_JobDefinitionRetryStrategy `json:"retryStrategy,omitempty" yaml:"retryStrategy,omitempty"`

	// When updating a job definition a new revision is created. This parameter determines if the previous version is `deregistered` (`INACTIVE`) or left  `ACTIVE`. Defaults to `true`.
	DeregisterOnNewRevision bool `json:"deregisterOnNewRevision,omitempty" yaml:"deregisterOnNewRevision,omitempty"`

	// Valid [ECS properties](http://docs.aws.amazon.com/batch/latest/APIReference/API_RegisterJobDefinition.html) provided as a single valid JSON document. This parameter is only valid if the `type` parameter is `container`.
	EcsProperties string `json:"ecsProperties,omitempty" yaml:"ecsProperties,omitempty"`
}
