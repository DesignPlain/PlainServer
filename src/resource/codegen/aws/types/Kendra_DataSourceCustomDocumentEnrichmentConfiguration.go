package types

type Kendra_DataSourceCustomDocumentEnrichmentConfiguration struct {
	// Configuration information to alter document attributes or metadata fields and content when ingesting documents into Amazon Kendra. Minimum number of `0` items. Maximum number of `100` items. Detailed below.
	InlineConfigurations []Kendra_DataSourceCustomDocumentEnrichmentConfigurationInlineConfiguration `json:"inlineConfigurations,omitempty" yaml:"inlineConfigurations,omitempty"`

	// A block that specifies the configuration information for invoking a Lambda function in AWS Lambda on the structured documents with their metadata and text extracted. You can use a Lambda function to apply advanced logic for creating, modifying, or deleting document metadata and content. For more information, see [Advanced data manipulation](https://docs.aws.amazon.com/kendra/latest/dg/custom-document-enrichment.html#advanced-data-manipulation). Detailed below.
	PostExtractionHookConfiguration Kendra_DataSourceCustomDocumentEnrichmentConfigurationPostExtractionHookConfiguration `json:"postExtractionHookConfiguration,omitempty" yaml:"postExtractionHookConfiguration,omitempty"`

	// Configuration information for invoking a Lambda function in AWS Lambda on the original or raw documents before extracting their metadata and text. You can use a Lambda function to apply advanced logic for creating, modifying, or deleting document metadata and content. For more information, see [Advanced data manipulation](https://docs.aws.amazon.com/kendra/latest/dg/custom-document-enrichment.html#advanced-data-manipulation). Detailed below.
	PreExtractionHookConfiguration Kendra_DataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfiguration `json:"preExtractionHookConfiguration,omitempty" yaml:"preExtractionHookConfiguration,omitempty"`

	// The Amazon Resource Name (ARN) of a role with permission to run `pre_extraction_hook_configuration` and `post_extraction_hook_configuration` for altering document metadata and content during the document ingestion process. For more information, see [IAM roles for Amazon Kendra](https://docs.aws.amazon.com/kendra/latest/dg/iam-roles.html).
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`
}
