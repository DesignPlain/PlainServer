package kendra

import types "libds/aws/types"

type Faq struct {
	//
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	//
	FileFormat string `json:"fileFormat,omitempty" yaml:"fileFormat,omitempty"`

	// The identifier of the index for a FAQ.
	IndexId string `json:"indexId,omitempty" yaml:"indexId,omitempty"`

	//
	LanguageCode string `json:"languageCode,omitempty" yaml:"languageCode,omitempty"`

	// The name that should be associated with the FAQ.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The Amazon Resource Name (ARN) of a role with permission to access the S3 bucket that contains the FAQs. For more information, see [IAM Roles for Amazon Kendra](https://docs.aws.amazon.com/kendra/latest/dg/iam-roles.html).
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// The S3 location of the FAQ input data. Detailed below.
	S3Path types.Kendra_FaqS3Path `json:"s3Path,omitempty" yaml:"s3Path,omitempty"`

	//
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
