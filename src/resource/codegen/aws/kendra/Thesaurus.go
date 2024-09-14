package kendra

import types "libds/aws/types"

type Thesaurus struct {
	//
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	//
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The identifier of the index for a thesaurus.
	IndexId string `json:"indexId,omitempty" yaml:"indexId,omitempty"`

	// The name for the thesaurus.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The IAM (Identity and Access Management) role used to access the thesaurus file in S3.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// The S3 path where your thesaurus file sits in S3. Detailed below.
	SourceS3Path types.Kendra_ThesaurusSourceS3Path `json:"sourceS3Path,omitempty" yaml:"sourceS3Path,omitempty"`
}
