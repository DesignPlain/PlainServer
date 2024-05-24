package types

type Ssm_MaintenanceWindowTaskTaskInvocationParametersLambdaParameters struct {
	// Pass client-specific information to the Lambda function that you are invoking.
	ClientContext string `json:"clientContext,omitempty" yaml:"clientContext,omitempty"`

	// JSON to provide to your Lambda function as input.
	Payload string `json:"payload,omitempty" yaml:"payload,omitempty"`

	// Specify a Lambda function version or alias name.
	Qualifier string `json:"qualifier,omitempty" yaml:"qualifier,omitempty"`
}
