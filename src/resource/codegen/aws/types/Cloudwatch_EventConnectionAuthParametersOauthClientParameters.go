package types

type Cloudwatch_EventConnectionAuthParametersOauthClientParameters struct {
	// The client ID for the credentials to use for authorization. Created and stored in AWS Secrets Manager.
	ClientId string `json:"clientId,omitempty" yaml:"clientId,omitempty"`

	// The client secret for the credentials to use for authorization. Created and stored in AWS Secrets Manager.
	ClientSecret string `json:"clientSecret,omitempty" yaml:"clientSecret,omitempty"`
}
