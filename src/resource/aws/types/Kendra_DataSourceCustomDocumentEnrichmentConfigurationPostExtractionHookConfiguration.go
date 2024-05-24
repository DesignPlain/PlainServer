package types

type Kendra_DataSourceCustomDocumentEnrichmentConfigurationPostExtractionHookConfiguration struct {
	// The Amazon Resource Name (ARN) of a Lambda Function that can manipulate your document metadata fields or attributes and content.
	LambdaArn string `json:"lambdaArn,omitempty" yaml:"lambdaArn,omitempty"`

	// Stores the original, raw documents or the structured, parsed documents before and after altering them. For more information, see [Data contracts for Lambda functions](https://docs.aws.amazon.com/kendra/latest/dg/custom-document-enrichment.html#cde-data-contracts-lambda).
	S3Bucket string `json:"s3Bucket,omitempty" yaml:"s3Bucket,omitempty"`

	// A block that specifies the condition used for when a Lambda function should be invoked. For example, you can specify a condition that if there are empty date-time values, then Amazon Kendra should invoke a function that inserts the current date-time. See invocation_condition.
	InvocationCondition Kendra_DataSourceCustomDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationCondition `json:"invocationCondition,omitempty" yaml:"invocationCondition,omitempty"`
}
