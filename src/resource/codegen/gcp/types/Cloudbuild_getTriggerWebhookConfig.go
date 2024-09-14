package types

type Cloudbuild_getTriggerWebhookConfig struct {
	/*
	   Potential issues with the underlying Pub/Sub subscription configuration.
	   Only populated on get requests.
	*/
	State string `json:"state,omitempty" yaml:"state,omitempty"`

	// Resource name for the secret required as a URL parameter.
	Secret string `json:"secret,omitempty" yaml:"secret,omitempty"`
}
