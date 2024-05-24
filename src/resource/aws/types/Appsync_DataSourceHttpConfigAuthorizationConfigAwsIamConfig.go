package types

type Appsync_DataSourceHttpConfigAuthorizationConfigAwsIamConfig struct {
	// Signing Amazon Web Services Region for IAM authorization.
	SigningRegion string `json:"signingRegion,omitempty" yaml:"signingRegion,omitempty"`

	// Signing service name for IAM authorization.
	SigningServiceName string `json:"signingServiceName,omitempty" yaml:"signingServiceName,omitempty"`
}
