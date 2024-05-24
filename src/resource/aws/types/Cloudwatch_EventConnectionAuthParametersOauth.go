package types

type Cloudwatch_EventConnectionAuthParametersOauth struct {
	// OAuth Http Parameters are additional credentials used to sign the request to the authorization endpoint to exchange the OAuth Client information for an access token. Secret values are stored and managed by AWS Secrets Manager. A maximum of 1 are allowed. Documented below.
	OauthHttpParameters Cloudwatch_EventConnectionAuthParametersOauthOauthHttpParameters `json:"oauthHttpParameters,omitempty" yaml:"oauthHttpParameters,omitempty"`

	// The URL to the authorization endpoint.
	AuthorizationEndpoint string `json:"authorizationEndpoint,omitempty" yaml:"authorizationEndpoint,omitempty"`

	// Contains the client parameters for OAuth authorization. Contains the following two parameters.
	ClientParameters Cloudwatch_EventConnectionAuthParametersOauthClientParameters `json:"clientParameters,omitempty" yaml:"clientParameters,omitempty"`

	// A password for the authorization. Created and stored in AWS Secrets Manager.
	HttpMethod string `json:"httpMethod,omitempty" yaml:"httpMethod,omitempty"`
}
