package types

type Cloudwatch_EventConnectionAuthParametersBasic struct {
	// A password for the authorization. Created and stored in AWS Secrets Manager.
	Password string `json:"password,omitempty" yaml:"password,omitempty"`

	// A username for the authorization.
	Username string `json:"username,omitempty" yaml:"username,omitempty"`
}
