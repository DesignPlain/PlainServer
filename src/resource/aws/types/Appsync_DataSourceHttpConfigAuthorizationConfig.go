package types

type Appsync_DataSourceHttpConfigAuthorizationConfig struct {
	// Identity and Access Management (IAM) settings. See AWS IAM Config.
	AwsIamConfig Appsync_DataSourceHttpConfigAuthorizationConfigAwsIamConfig `json:"awsIamConfig,omitempty" yaml:"awsIamConfig,omitempty"`

	// Authorization type that the HTTP endpoint requires. Default values is `AWS_IAM`.
	AuthorizationType string `json:"authorizationType,omitempty" yaml:"authorizationType,omitempty"`
}
