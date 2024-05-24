package types

type Inspector2_OrganizationConfigurationAutoEnable struct {
	// Whether Amazon EC2 scans are automatically enabled for new members of your Amazon Inspector organization.
	Ec2 bool `json:"ec2,omitempty" yaml:"ec2,omitempty"`

	// Whether Amazon ECR scans are automatically enabled for new members of your Amazon Inspector organization.
	Ecr bool `json:"ecr,omitempty" yaml:"ecr,omitempty"`

	// Whether Lambda Function scans are automatically enabled for new members of your Amazon Inspector organization.
	Lambda bool `json:"lambda,omitempty" yaml:"lambda,omitempty"`

	// Whether AWS Lambda code scans are automatically enabled for new members of your Amazon Inspector organization. --Note:-- Lambda code scanning requires Lambda standard scanning to be activated. Consequently, if you are setting this argument to `true`, you must also set the `lambda` argument to `true`. See [Scanning AWS Lambda functions with Amazon Inspector](https://docs.aws.amazon.com/inspector/latest/user/scanning-lambda.html#lambda-code-scans) for more information.
	LambdaCode bool `json:"lambdaCode,omitempty" yaml:"lambdaCode,omitempty"`
}
