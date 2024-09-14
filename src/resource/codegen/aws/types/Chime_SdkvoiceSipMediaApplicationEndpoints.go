package types

type Chime_SdkvoiceSipMediaApplicationEndpoints struct {
	// Valid Amazon Resource Name (ARN) of the Lambda function, version, or alias. The function must be created in the same AWS Region as the SIP media application.
	LambdaArn string `json:"lambdaArn,omitempty" yaml:"lambdaArn,omitempty"`
}
