package types

type Cognito_getUserPoolLambdaConfigCustomEmailSender struct {
	// - ARN of the Lambda function.
	LambdaArn string `json:"lambdaArn,omitempty" yaml:"lambdaArn,omitempty"`

	// - Version of the Lambda function.
	LambdaVersion string `json:"lambdaVersion,omitempty" yaml:"lambdaVersion,omitempty"`
}
