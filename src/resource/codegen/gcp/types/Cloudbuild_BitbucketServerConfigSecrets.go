package types

type Cloudbuild_BitbucketServerConfigSecrets struct {
	// The resource name for the admin access token's secret version.
	AdminAccessTokenVersionName string `json:"adminAccessTokenVersionName,omitempty" yaml:"adminAccessTokenVersionName,omitempty"`

	// The resource name for the read access token's secret version.
	ReadAccessTokenVersionName string `json:"readAccessTokenVersionName,omitempty" yaml:"readAccessTokenVersionName,omitempty"`

	/*
	   Immutable. The resource name for the webhook secret's secret version. Once this field has been set, it cannot be changed.
	   Changing this field will result in deleting/ recreating the resource.

	   - - -
	*/
	WebhookSecretVersionName string `json:"webhookSecretVersionName,omitempty" yaml:"webhookSecretVersionName,omitempty"`
}
