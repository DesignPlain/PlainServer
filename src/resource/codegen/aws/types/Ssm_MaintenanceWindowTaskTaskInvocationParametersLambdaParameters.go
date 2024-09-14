package types

type Ssm_MaintenanceWindowTaskTaskInvocationParametersLambdaParameters struct {
	// JSON to provide to your Lambda function as input.
	Payload string `json:"payload,omitempty" yaml:"payload,omitempty"`

	// Specify a Lambda function version or alias name.
	Qualifier string `json:"qualifier,omitempty" yaml:"qualifier,omitempty"`

	// Pass client-specific information to the Lambda function that you are invoking.
	ClientContext string `json:"clientContext,omitempty" yaml:"clientContext,omitempty"`
}
