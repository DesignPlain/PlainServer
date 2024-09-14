package macie2

import types "libds/aws/types"

type ClassificationJob struct {
	// The status for the job. Valid values are: `CANCELLED`, `RUNNING` and `USER_PAUSED`
	JobStatus string `json:"jobStatus,omitempty" yaml:"jobStatus,omitempty"`

	// The schedule for running the job. Valid values are: `ONE_TIME` - Run the job only once. If you specify this value, don't specify a value for the `schedule_frequency` property. `SCHEDULED` - Run the job on a daily, weekly, or monthly basis. If you specify this value, use the `schedule_frequency` property to define the recurrence pattern for the job.
	JobType string `json:"jobType,omitempty" yaml:"jobType,omitempty"`

	// A custom name for the job. The name can contain as many as 500 characters. If omitted, the provider will assign a random, unique name. Conflicts with `name_prefix`.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix string `json:"namePrefix,omitempty" yaml:"namePrefix,omitempty"`

	// A map of key-value pairs that specifies the tags to associate with the job. A job can have a maximum of 50 tags. Each tag consists of a tag key and an associated tag value. The maximum length of a tag key is 128 characters. The maximum length of a tag value is 256 characters.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The custom data identifiers to use for data analysis and classification.
	CustomDataIdentifierIds []string `json:"customDataIdentifierIds,omitempty" yaml:"customDataIdentifierIds,omitempty"`

	// A custom description of the job. The description can contain as many as 200 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Specifies whether to analyze all existing, eligible objects immediately after the job is created.
	InitialRun bool `json:"initialRun,omitempty" yaml:"initialRun,omitempty"`

	// The S3 buckets that contain the objects to analyze, and the scope of that analysis. (documented below)
	S3JobDefinition types.Macie2_ClassificationJobS3JobDefinition `json:"s3JobDefinition,omitempty" yaml:"s3JobDefinition,omitempty"`

	// The sampling depth, as a percentage, to apply when processing objects. This value determines the percentage of eligible objects that the job analyzes. If this value is less than 100, Amazon Macie selects the objects to analyze at random, up to the specified percentage, and analyzes all the data in those objects.
	SamplingPercentage int `json:"samplingPercentage,omitempty" yaml:"samplingPercentage,omitempty"`

	// The recurrence pattern for running the job. To run the job only once, don't specify a value for this property and set the value for the `job_type` property to `ONE_TIME`. (documented below)
	ScheduleFrequency types.Macie2_ClassificationJobScheduleFrequency `json:"scheduleFrequency,omitempty" yaml:"scheduleFrequency,omitempty"`
}
