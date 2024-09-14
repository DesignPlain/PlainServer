package glue

import types "libds/aws/types"

type MLTransform struct {
	// A list of AWS Glue table definitions used by the transform. see Input Record Tables.
	InputRecordTables []types.Glue_MLTransformInputRecordTable `json:"inputRecordTables,omitempty" yaml:"inputRecordTables,omitempty"`

	// The number of workers of a defined `worker_type` that are allocated when an ML Transform runs. Required with `worker_type`.
	NumberOfWorkers int `json:"numberOfWorkers,omitempty" yaml:"numberOfWorkers,omitempty"`

	// The ML Transform timeout in minutes. The default is 2880 minutes (48 hours).
	Timeout int `json:"timeout,omitempty" yaml:"timeout,omitempty"`

	// The type of predefined worker that is allocated when an ML Transform runs. Accepts a value of `Standard`, `G.1X`, or `G.2X`. Required with `number_of_workers`.
	WorkerType string `json:"workerType,omitempty" yaml:"workerType,omitempty"`

	// The name you assign to this ML Transform. It must be unique in your account.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The algorithmic parameters that are specific to the transform type used. Conditionally dependent on the transform type. see Parameters.
	Parameters types.Glue_MLTransformParameters `json:"parameters,omitempty" yaml:"parameters,omitempty"`

	// The ARN of the IAM role associated with this ML Transform.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Description of the ML Transform.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The version of glue to use, for example "1.0". For information about available versions, see the [AWS Glue Release Notes](https://docs.aws.amazon.com/glue/latest/dg/release-notes.html).
	GlueVersion string `json:"glueVersion,omitempty" yaml:"glueVersion,omitempty"`

	// The number of AWS Glue data processing units (DPUs) that are allocated to task runs for this transform. You can allocate from `2` to `100` DPUs; the default is `10`. `max_capacity` is a mutually exclusive option with `number_of_workers` and `worker_type`.
	MaxCapacity float64 `json:"maxCapacity,omitempty" yaml:"maxCapacity,omitempty"`

	// The maximum number of times to retry this ML Transform if it fails.
	MaxRetries int `json:"maxRetries,omitempty" yaml:"maxRetries,omitempty"`
}
