package appsync

type ApiKey struct {
	// RFC3339 string representation of the expiry date. Rounded down to nearest hour. By default, it is 7 days from the date of creation.
	Expires string `json:"expires,omitempty" yaml:"expires,omitempty"`

	// ID of the associated AppSync API
	ApiId string `json:"apiId,omitempty" yaml:"apiId,omitempty"`

	// API key description. Defaults to "Managed by Pulumi".
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
