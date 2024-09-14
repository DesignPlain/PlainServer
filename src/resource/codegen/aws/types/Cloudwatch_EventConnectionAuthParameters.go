package types

type Cloudwatch_EventConnectionAuthParameters struct {
	// Parameters used for API_KEY authorization. An API key to include in the header for each authentication request. A maximum of 1 are allowed. Conflicts with `basic` and `oauth`. Documented below.
	ApiKey Cloudwatch_EventConnectionAuthParametersApiKey `json:"apiKey,omitempty" yaml:"apiKey,omitempty"`

	// Parameters used for BASIC authorization. A maximum of 1 are allowed. Conflicts with `api_key` and `oauth`. Documented below.
	Basic Cloudwatch_EventConnectionAuthParametersBasic `json:"basic,omitempty" yaml:"basic,omitempty"`

	// Invocation Http Parameters are additional credentials used to sign each Invocation of the ApiDestination created from this Connection. If the ApiDestination Rule Target has additional HttpParameters, the values will be merged together, with the Connection Invocation Http Parameters taking precedence. Secret values are stored and managed by AWS Secrets Manager. A maximum of 1 are allowed. Documented below.
	InvocationHttpParameters Cloudwatch_EventConnectionAuthParametersInvocationHttpParameters `json:"invocationHttpParameters,omitempty" yaml:"invocationHttpParameters,omitempty"`

	// Parameters used for OAUTH_CLIENT_CREDENTIALS authorization. A maximum of 1 are allowed. Conflicts with `basic` and `api_key`. Documented below.
	Oauth Cloudwatch_EventConnectionAuthParametersOauth `json:"oauth,omitempty" yaml:"oauth,omitempty"`
}
