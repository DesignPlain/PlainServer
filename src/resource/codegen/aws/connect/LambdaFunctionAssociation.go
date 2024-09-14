package connect

type LambdaFunctionAssociation struct {
	// Amazon Resource Name (ARN) of the Lambda Function, omitting any version or alias qualifier.
	FunctionArn string `json:"functionArn,omitempty" yaml:"functionArn,omitempty"`

	// The identifier of the Amazon Connect instance. You can find the instanceId in the ARN of the instance.
	InstanceId string `json:"instanceId,omitempty" yaml:"instanceId,omitempty"`
}
