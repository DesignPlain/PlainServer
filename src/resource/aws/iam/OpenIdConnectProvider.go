package iam

type OpenIdConnectProvider struct {
	// Map of resource tags for the IAM OIDC provider. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// A list of server certificate thumbprints for the OpenID Connect (OIDC) identity provider's server certificate(s).
	ThumbprintLists []string `json:"thumbprintLists,omitempty" yaml:"thumbprintLists,omitempty"`

	// The URL of the identity provider. Corresponds to the _iss_ claim.
	Url string `json:"url,omitempty" yaml:"url,omitempty"`

	// A list of client IDs (also known as audiences). When a mobile or web app registers with an OpenID Connect provider, they establish a value that identifies the application. (This is the value that's sent as the client_id parameter on OAuth requests.)
	ClientIdLists []string `json:"clientIdLists,omitempty" yaml:"clientIdLists,omitempty"`
}
