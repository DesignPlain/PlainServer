package types

type Appsync_DataSourceHttpConfigAuthorizationConfig struct {
	// Authorization type that the HTTP endpoint requires. Default values is `AWS_IAM`.
	AuthorizationType string `json:"authorizationType,omitempty" yaml:"authorizationType,omitempty"`

	// Identity and Access Management (IAM) settings. See `aws_iam_config` Block for details.
	AwsIamConfig Appsync_DataSourceHttpConfigAuthorizationConfigAwsIamConfig `json:"awsIamConfig,omitempty" yaml:"awsIamConfig,omitempty"`
}
