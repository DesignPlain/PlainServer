package types

type S3control_ObjectLambdaAccessPointConfigurationTransformationConfigurationContentTransformationAwsLambda struct {
	// The Amazon Resource Name (ARN) of the AWS Lambda function.
	FunctionArn string `json:"functionArn,omitempty" yaml:"functionArn,omitempty"`

	// Additional JSON that provides supplemental data to the Lambda function used to transform objects.
	FunctionPayload string `json:"functionPayload,omitempty" yaml:"functionPayload,omitempty"`
}
