package pipes

import types "libds/aws/types"

type Pipe struct {
	// Parameters to configure a source for the pipe. Detailed below.
	SourceParameters types.Pipes_PipeSourceParameters `json:"sourceParameters,omitempty" yaml:"sourceParameters,omitempty"`

	// Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Parameters to configure a target for your pipe. Detailed below.
	TargetParameters types.Pipes_PipeTargetParameters `json:"targetParameters,omitempty" yaml:"targetParameters,omitempty"`

	// The state the pipe should be in. One of: `RUNNING`, `STOPPED`.
	DesiredState string `json:"desiredState,omitempty" yaml:"desiredState,omitempty"`

	// Enrichment resource of the pipe (typically an ARN). Read more about enrichment in the [User Guide](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-pipes.html#pipes-enrichment).
	Enrichment string `json:"enrichment,omitempty" yaml:"enrichment,omitempty"`

	// Parameters to configure enrichment for your pipe. Detailed below.
	EnrichmentParameters types.Pipes_PipeEnrichmentParameters `json:"enrichmentParameters,omitempty" yaml:"enrichmentParameters,omitempty"`

	// Logging configuration settings for the pipe. Detailed below.
	LogConfiguration types.Pipes_PipeLogConfiguration `json:"logConfiguration,omitempty" yaml:"logConfiguration,omitempty"`

	// ARN of the role that allows the pipe to send data to the target.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// A description of the pipe. At most 512 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Name of the pipe. If omitted, the provider will assign a random, unique name. Conflicts with `name_prefix`.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix string `json:"namePrefix,omitempty" yaml:"namePrefix,omitempty"`

	// Source resource of the pipe. This field typically requires an ARN (Amazon Resource Name). However, when using a self-managed Kafka cluster, you should use a different format. Instead of an ARN, use 'smk://' followed by the bootstrap server's address.
	Source string `json:"source,omitempty" yaml:"source,omitempty"`

	/*
	   Target resource of the pipe (typically an ARN).

	   The following arguments are optional:
	*/
	Target string `json:"target,omitempty" yaml:"target,omitempty"`
}
