package types

type Codepipeline_WebhookAuthenticationConfiguration struct {
	// A valid CIDR block for `IP` filtering. Required for `IP`.
	AllowedIpRange string `json:"allowedIpRange,omitempty" yaml:"allowedIpRange,omitempty"`

	// The shared secret for the GitHub repository webhook. Set this as `secret` in your `github_repository_webhook`'s `configuration` block. Required for `GITHUB_HMAC`.
	SecretToken string `json:"secretToken,omitempty" yaml:"secretToken,omitempty"`
}
