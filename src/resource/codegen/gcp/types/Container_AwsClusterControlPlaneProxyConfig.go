package types

type Container_AwsClusterControlPlaneProxyConfig struct {
	// The ARN of the AWS Secret Manager secret that contains the HTTP(S) proxy configuration.
	SecretArn string `json:"secretArn,omitempty" yaml:"secretArn,omitempty"`

	// The version string of the AWS Secret Manager secret that contains the HTTP(S) proxy configuration.
	SecretVersion string `json:"secretVersion,omitempty" yaml:"secretVersion,omitempty"`
}
