package kendra

import types "libds/aws/types"

type QuerySuggestionsBlockList struct {
	//
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Identifier of the index for a block list.
	IndexId string `json:"indexId,omitempty" yaml:"indexId,omitempty"`

	// Name for the block list.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// IAM (Identity and Access Management) role used to access the block list text file in S3.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// S3 path where your block list text file is located. See details below.
	SourceS3Path types.Kendra_QuerySuggestionsBlockListSourceS3Path `json:"sourceS3Path,omitempty" yaml:"sourceS3Path,omitempty"`

	//
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
