package types

type Looker_InstanceOauthConfig struct {
	// The client ID for the Oauth config.
	ClientId string `json:"clientId,omitempty" yaml:"clientId,omitempty"`

	// The client secret for the Oauth config.
	ClientSecret string `json:"clientSecret,omitempty" yaml:"clientSecret,omitempty"`
}
