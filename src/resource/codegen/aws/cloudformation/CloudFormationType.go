package cloudformation

import types "libds/aws/types"

type CloudFormationType struct {
	// Amazon Resource Name (ARN) of the IAM Role for CloudFormation to assume when invoking the extension. If your extension calls AWS APIs in any of its handlers, you must create an IAM execution role that includes the necessary permissions to call those AWS APIs, and provision that execution role in your account. When CloudFormation needs to invoke the extension handler, CloudFormation assumes this execution role to create a temporary session token, which it then passes to the extension handler, thereby supplying your extension with the appropriate credentials.
	ExecutionRoleArn string `json:"executionRoleArn,omitempty" yaml:"executionRoleArn,omitempty"`

	// Configuration block containing logging configuration.
	LoggingConfig types.Cloudformation_CloudFormationTypeLoggingConfig `json:"loggingConfig,omitempty" yaml:"loggingConfig,omitempty"`

	// URL to the S3 bucket containing the extension project package that contains the necessary files for the extension you want to register. Must begin with `s3://` or `https://`. For example, `s3://example-bucket/example-object`.
	SchemaHandlerPackage string `json:"schemaHandlerPackage,omitempty" yaml:"schemaHandlerPackage,omitempty"`

	// CloudFormation Registry Type. For example, `RESOURCE` or `MODULE`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// CloudFormation Type name. For example, `ExampleCompany::ExampleService::ExampleResource`.
	TypeName string `json:"typeName,omitempty" yaml:"typeName,omitempty"`
}
