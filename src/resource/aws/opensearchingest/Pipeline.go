package opensearchingest

import types "DesignSphere_Server/src/resource/aws/types"

type Pipeline struct {
	// The pipeline configuration in YAML format. This argument accepts the pipeline configuration as a string or within a .yaml file. If you provide the configuration as a string, each new line must be escaped with \n.
	PipelineConfigurationBody string `json:"pipelineConfigurationBody,omitempty" yaml:"pipelineConfigurationBody,omitempty"`

	/*
	   The name of the OpenSearch Ingestion pipeline to create. Pipeline names are unique across the pipelines owned by an account within an AWS Region.

	   The following arguments are optional:
	*/
	PipelineName string `json:"pipelineName,omitempty" yaml:"pipelineName,omitempty"`

	// Container for the values required to configure VPC access for the pipeline. If you don't specify these values, OpenSearch Ingestion creates the pipeline with a public endpoint. See `vpc_options` below.
	VpcOptions types.Opensearchingest_PipelineVpcOptions `json:"vpcOptions,omitempty" yaml:"vpcOptions,omitempty"`

	// Key-value pairs to configure persistent buffering for the pipeline. See `buffer_options` below.
	BufferOptions types.Opensearchingest_PipelineBufferOptions `json:"bufferOptions,omitempty" yaml:"bufferOptions,omitempty"`

	// Key-value pairs to configure log publishing. See `log_publishing_options` below.
	LogPublishingOptions types.Opensearchingest_PipelineLogPublishingOptions `json:"logPublishingOptions,omitempty" yaml:"logPublishingOptions,omitempty"`

	// The minimum pipeline capacity, in Ingestion Compute Units (ICUs).
	MinUnits int `json:"minUnits,omitempty" yaml:"minUnits,omitempty"`

	// A map of tags to assign to the pipeline. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	//
	Timeouts types.Opensearchingest_PipelineTimeouts `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`

	// Key-value pairs to configure encryption for data that is written to a persistent buffer. See `encryption_at_rest_options` below.
	EncryptionAtRestOptions types.Opensearchingest_PipelineEncryptionAtRestOptions `json:"encryptionAtRestOptions,omitempty" yaml:"encryptionAtRestOptions,omitempty"`

	// The maximum pipeline capacity, in Ingestion Compute Units (ICUs).
	MaxUnits int `json:"maxUnits,omitempty" yaml:"maxUnits,omitempty"`
}
