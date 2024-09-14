package pinpoint

type BaiduChannel struct {
	// Platform credential API key from Baidu.
	ApiKey string `json:"apiKey,omitempty" yaml:"apiKey,omitempty"`

	// The application ID.
	ApplicationId string `json:"applicationId,omitempty" yaml:"applicationId,omitempty"`

	// Specifies whether to enable the channel. Defaults to `true`.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// Platform credential Secret key from Baidu.
	SecretKey string `json:"secretKey,omitempty" yaml:"secretKey,omitempty"`
}
