package iot

type Authorizer struct {
	// The name of the token key used to extract the token from the HTTP headers. This value is required if signing is enabled in your authorizer.
	TokenKeyName string `json:"tokenKeyName,omitempty" yaml:"tokenKeyName,omitempty"`

	// The public keys used to verify the digital signature returned by your custom authentication service. This value is required if signing is enabled in your authorizer.
	TokenSigningPublicKeys map[string]string `json:"tokenSigningPublicKeys,omitempty" yaml:"tokenSigningPublicKeys,omitempty"`

	// The ARN of the authorizer's Lambda function.
	AuthorizerFunctionArn string `json:"authorizerFunctionArn,omitempty" yaml:"authorizerFunctionArn,omitempty"`

	// Specifies whether the HTTP caching is enabled or not. Default: `false`.
	EnableCachingForHttp bool `json:"enableCachingForHttp,omitempty" yaml:"enableCachingForHttp,omitempty"`

	// The name of the authorizer.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Specifies whether AWS IoT validates the token signature in an authorization request. Default: `false`.
	SigningDisabled bool `json:"signingDisabled,omitempty" yaml:"signingDisabled,omitempty"`

	// The status of Authorizer request at creation. Valid values: `ACTIVE`, `INACTIVE`. Default: `ACTIVE`.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`
}
