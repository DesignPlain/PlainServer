package kendra

import types "DesignSphere_Server/src/resource/aws/types"

type Faq struct {
	// The Amazon Resource Name (ARN) of a role with permission to access the S3 bucket that contains the FAQs. For more information, see [IAM Roles for Amazon Kendra](https://docs.aws.amazon.com/kendra/latest/dg/iam-roles.html).
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	/*
	   The S3 location of the FAQ input data. Detailed below.

	   The `s3_path` configuration block supports the following arguments:
	*/
	S3Path types.Kendra_FaqS3Path `json:"s3Path,omitempty" yaml:"s3Path,omitempty"`

	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The description for a FAQ.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The file format used by the input files for the FAQ. Valid Values are `CSV`, `CSV_WITH_HEADER`, `JSON`.
	FileFormat string `json:"fileFormat,omitempty" yaml:"fileFormat,omitempty"`

	// The identifier of the index for a FAQ.
	IndexId string `json:"indexId,omitempty" yaml:"indexId,omitempty"`

	// The code for a language. This shows a supported language for the FAQ document. English is supported by default. For more information on supported languages, including their codes, see [Adding documents in languages other than English](https://docs.aws.amazon.com/kendra/latest/dg/in-adding-languages.html).
	LanguageCode string `json:"languageCode,omitempty" yaml:"languageCode,omitempty"`

	// The name that should be associated with the FAQ.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
