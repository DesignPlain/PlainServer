package types

type Cognito_UserPoolLambdaConfigCustomSmsSender struct {
	// The Lambda Amazon Resource Name of the Lambda function that Amazon Cognito triggers to send SMS notifications to users.
	LambdaArn string `json:"lambdaArn,omitempty" yaml:"lambdaArn,omitempty"`

	// The Lambda version represents the signature of the "request" attribute in the "event" information Amazon Cognito passes to your custom SMS Lambda function. The only supported value is `V1_0`.
	LambdaVersion string `json:"lambdaVersion,omitempty" yaml:"lambdaVersion,omitempty"`
}
