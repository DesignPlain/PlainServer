package lambda

import types "libds/aws/types"

type Function struct {
	// Amazon Resource Name (ARN) of the AWS Key Management Service (KMS) key that is used to encrypt environment variables. If this configuration is not provided when environment variables are in use, AWS Lambda uses a default service key. If this configuration is provided when environment variables are not in use, the AWS Lambda API does not save this configuration and the provider will show a perpetual difference of adding the key. To fix the perpetual difference, remove this configuration.
	KmsKeyArn string `json:"kmsKeyArn,omitempty" yaml:"kmsKeyArn,omitempty"`

	// Identifier of the function's runtime. See [Runtimes](https://docs.aws.amazon.com/lambda/latest/dg/API_CreateFunction.html#SSS-CreateFunction-request-Runtime) for valid values.
	Runtime string `json:"runtime,omitempty" yaml:"runtime,omitempty"`

	// S3 bucket location containing the function's deployment package. This bucket must reside in the same AWS region where you are creating the Lambda function. Exactly one of `filename`, `image_uri`, or `s3_bucket` must be specified. When `s3_bucket` is set, `s3_key` is required.
	S3Bucket string `json:"s3Bucket,omitempty" yaml:"s3Bucket,omitempty"`

	// The amount of Ephemeral storage(`/tmp`) to allocate for the Lambda Function in MB. This parameter is used to expand the total amount of Ephemeral storage available, beyond the default amount of `512`MB. Detailed below.
	EphemeralStorage types.Lambda_FunctionEphemeralStorage `json:"ephemeralStorage,omitempty" yaml:"ephemeralStorage,omitempty"`

	// Description of what your Lambda Function does.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Path to the function's deployment package within the local filesystem. Exactly one of `filename`, `image_uri`, or `s3_bucket` must be specified.
	Code string `json:"code,omitempty" yaml:"code,omitempty"`

	// Configuration block used to specify advanced logging settings. Detailed below.
	LoggingConfig types.Lambda_FunctionLoggingConfig `json:"loggingConfig,omitempty" yaml:"loggingConfig,omitempty"`

	// Unique name for your Lambda Function.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Snap start settings block. Detailed below.
	SnapStart types.Lambda_FunctionSnapStart `json:"snapStart,omitempty" yaml:"snapStart,omitempty"`

	// Configuration block. Detailed below.
	TracingConfig types.Lambda_FunctionTracingConfig `json:"tracingConfig,omitempty" yaml:"tracingConfig,omitempty"`

	// List of Lambda Layer Version ARNs (maximum of 5) to attach to your Lambda Function. See [Lambda Layers](https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html)
	Layers []string `json:"layers,omitempty" yaml:"layers,omitempty"`

	// Amount of memory in MB your Lambda Function can use at runtime. Defaults to `128`. See [Limits](https://docs.aws.amazon.com/lambda/latest/dg/limits.html)
	MemorySize int `json:"memorySize,omitempty" yaml:"memorySize,omitempty"`

	/*
	   List of security group IDs to assign to the function's VPC configuration prior to destruction.
	   `replace_security_groups_on_destroy` must be set to `true` to use this attribute.
	*/
	ReplacementSecurityGroupIds []string `json:"replacementSecurityGroupIds,omitempty" yaml:"replacementSecurityGroupIds,omitempty"`

	// Set to true if you do not wish the function to be deleted at destroy time, and instead just remove the function from the Pulumi state.
	SkipDestroy bool `json:"skipDestroy,omitempty" yaml:"skipDestroy,omitempty"`

	// Configuration block. Detailed below.
	VpcConfig types.Lambda_FunctionVpcConfig `json:"vpcConfig,omitempty" yaml:"vpcConfig,omitempty"`

	// Configuration block. Detailed below.
	FileSystemConfig types.Lambda_FunctionFileSystemConfig `json:"fileSystemConfig,omitempty" yaml:"fileSystemConfig,omitempty"`

	// Function [entrypoint](https://docs.aws.amazon.com/lambda/latest/dg/walkthrough-custom-events-create-test-function.html) in your code.
	Handler string `json:"handler,omitempty" yaml:"handler,omitempty"`

	// Configuration block. Detailed below.
	ImageConfig types.Lambda_FunctionImageConfig `json:"imageConfig,omitempty" yaml:"imageConfig,omitempty"`

	// Whether to publish creation/change as new Lambda Function Version. Defaults to `false`.
	Publish bool `json:"publish,omitempty" yaml:"publish,omitempty"`

	// Amount of reserved concurrent executions for this lambda function. A value of `0` disables lambda from being triggered and `-1` removes any concurrency limitations. Defaults to Unreserved Concurrency Limits `-1`. See [Managing Concurrency](https://docs.aws.amazon.com/lambda/latest/dg/concurrent-executions.html)
	ReservedConcurrentExecutions int `json:"reservedConcurrentExecutions,omitempty" yaml:"reservedConcurrentExecutions,omitempty"`

	// S3 key of an object containing the function's deployment package. When `s3_bucket` is set, `s3_key` is required.
	S3Key string `json:"s3Key,omitempty" yaml:"s3Key,omitempty"`

	// To enable code signing for this function, specify the ARN of a code-signing configuration. A code-signing configuration includes a set of signing profiles, which define the trusted publishers for this function.
	CodeSigningConfigArn string `json:"codeSigningConfigArn,omitempty" yaml:"codeSigningConfigArn,omitempty"`

	// Lambda deployment package type. Valid values are `Zip` and `Image`. Defaults to `Zip`.
	PackageType string `json:"packageType,omitempty" yaml:"packageType,omitempty"`

	// Virtual attribute used to trigger replacement when source code changes. Must be set to a base64-encoded SHA256 hash of the package file specified with either `filename` or `s3_key`.
	SourceCodeHash string `json:"sourceCodeHash,omitempty" yaml:"sourceCodeHash,omitempty"`

	// ECR image URI containing the function's deployment package. Exactly one of `filename`, `image_uri`,  or `s3_bucket` must be specified.
	ImageUri string `json:"imageUri,omitempty" yaml:"imageUri,omitempty"`

	// Configuration block. Detailed below.
	DeadLetterConfig types.Lambda_FunctionDeadLetterConfig `json:"deadLetterConfig,omitempty" yaml:"deadLetterConfig,omitempty"`

	// Configuration block. Detailed below.
	Environment types.Lambda_FunctionEnvironment `json:"environment,omitempty" yaml:"environment,omitempty"`

	/*
	   Whether to replace the security groups on the function's VPC configuration prior to destruction.
	   Removing these security group associations prior to function destruction can speed up security group deletion times of AWS's internal cleanup operations.
	   By default, the security groups will be replaced with the `default` security group in the function's configured VPC.
	   Set the `replacement_security_group_ids` attribute to use a custom list of security groups for replacement.
	*/
	ReplaceSecurityGroupsOnDestroy bool `json:"replaceSecurityGroupsOnDestroy,omitempty" yaml:"replaceSecurityGroupsOnDestroy,omitempty"`

	/*
	   Amazon Resource Name (ARN) of the function's execution role. The role provides the function's identity and access to AWS services and resources.

	   The following arguments are optional:
	*/
	Role string `json:"role,omitempty" yaml:"role,omitempty"`

	// Object version containing the function's deployment package. Conflicts with `filename` and `image_uri`.
	S3ObjectVersion string `json:"s3ObjectVersion,omitempty" yaml:"s3ObjectVersion,omitempty"`

	// Instruction set architecture for your Lambda function. Valid values are `["x86_64"]` and `["arm64"]`. Default is `["x86_64"]`. Removing this attribute, function's architecture stay the same.
	Architectures []string `json:"architectures,omitempty" yaml:"architectures,omitempty"`

	// Amount of time your Lambda Function has to run in seconds. Defaults to `3`. See [Limits](https://docs.aws.amazon.com/lambda/latest/dg/limits.html).
	Timeout int `json:"timeout,omitempty" yaml:"timeout,omitempty"`

	// Map of tags to assign to the object. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
