package types

type Cloudbuild_getTriggerPubsubConfig struct {
	// Service account that will make the push request.
	ServiceAccountEmail string `json:"serviceAccountEmail,omitempty" yaml:"serviceAccountEmail,omitempty"`

	/*
	   Potential issues with the underlying Pub/Sub subscription configuration.
	   Only populated on get requests.
	*/
	State string `json:"state,omitempty" yaml:"state,omitempty"`

	// Output only. Name of the subscription.
	Subscription string `json:"subscription,omitempty" yaml:"subscription,omitempty"`

	// The name of the topic from which this subscription is receiving messages.
	Topic string `json:"topic,omitempty" yaml:"topic,omitempty"`
}
