package kendra

import types "libds/aws/types"

type DataSource struct {
	// A block with the configuration information to connect to your Data Source repository. You can't specify the `configuration` block when the `type` parameter is set to `CUSTOM`. Detailed below.
	Configuration types.Kendra_DataSourceConfiguration `json:"configuration,omitempty" yaml:"configuration,omitempty"`

	// A description for the Data Source connector.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The Amazon Resource Name (ARN) of a role with permission to access the data source connector. For more information, see [IAM roles for Amazon Kendra](https://docs.aws.amazon.com/kendra/latest/dg/iam-roles.html). You can't specify the `role_arn` parameter when the `type` parameter is set to `CUSTOM`. The `role_arn` parameter is required for all other data sources.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// Sets the frequency for Amazon Kendra to check the documents in your Data Source repository and update the index. If you don't set a schedule Amazon Kendra will not periodically update the index. You can call the `StartDataSourceSyncJob` API to update the index.
	Schedule string `json:"schedule,omitempty" yaml:"schedule,omitempty"`

	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// A block with the configuration information for altering document metadata and content during the document ingestion process. For more information on how to create, modify and delete document metadata, or make other content alterations when you ingest documents into Amazon Kendra, see [Customizing document metadata during the ingestion process](https://docs.aws.amazon.com/kendra/latest/dg/custom-document-enrichment.html). Detailed below.
	CustomDocumentEnrichmentConfiguration types.Kendra_DataSourceCustomDocumentEnrichmentConfiguration `json:"customDocumentEnrichmentConfiguration,omitempty" yaml:"customDocumentEnrichmentConfiguration,omitempty"`

	// The identifier of the index for your Amazon Kendra data source.
	IndexId string `json:"indexId,omitempty" yaml:"indexId,omitempty"`

	// The code for a language. This allows you to support a language for all documents when creating the Data Source connector. English is supported by default. For more information on supported languages, including their codes, see [Adding documents in languages other than English](https://docs.aws.amazon.com/kendra/latest/dg/in-adding-languages.html).
	LanguageCode string `json:"languageCode,omitempty" yaml:"languageCode,omitempty"`

	// A name for your data source connector.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The type of data source repository. For an updated list of values, refer to [Valid Values for Type](https://docs.aws.amazon.com/kendra/latest/dg/API_CreateDataSource.html#Kendra-CreateDataSource-request-Type).

	   The following arguments are optional:
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
