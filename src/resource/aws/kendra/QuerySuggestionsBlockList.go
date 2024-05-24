package kendra

import types "DesignSphere_Server/src/resource/aws/types"

type QuerySuggestionsBlockList struct {
	// Identifier of the index for a block list.
	IndexId string `json:"indexId,omitempty" yaml:"indexId,omitempty"`

	// Name for the block list.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// IAM (Identity and Access Management) role used to access the block list text file in S3.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	/*
	   S3 path where your block list text file is located. See details below.

	   The `source_s3_path` configuration block supports the following arguments:
	*/
	SourceS3Path types.Kendra_QuerySuggestionsBlockListSourceS3Path `json:"sourceS3Path,omitempty" yaml:"sourceS3Path,omitempty"`

	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Description for a block list.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
