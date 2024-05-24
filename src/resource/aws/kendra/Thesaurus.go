package kendra

import types "DesignSphere_Server/src/resource/aws/types"

type Thesaurus struct {
	// The description for a thesaurus.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The identifier of the index for a thesaurus.
	IndexId string `json:"indexId,omitempty" yaml:"indexId,omitempty"`

	// The name for the thesaurus.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The IAM (Identity and Access Management) role used to access the thesaurus file in S3.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	/*
	   The S3 path where your thesaurus file sits in S3. Detailed below.

	   The `source_s3_path` configuration block supports the following arguments:
	*/
	SourceS3Path types.Kendra_ThesaurusSourceS3Path `json:"sourceS3Path,omitempty" yaml:"sourceS3Path,omitempty"`

	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
