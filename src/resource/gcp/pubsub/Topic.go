package pubsub

import types "DesignSphere_Server/src/resource/gcp/types"

type Topic struct {
	/*
	   Settings for validating messages published against a schema.
	   Structure is documented below.
	*/
	SchemaSettings types.Pubsub_TopicSchemaSettings `json:"schemaSettings,omitempty" yaml:"schemaSettings,omitempty"`

	/*
	   The resource name of the Cloud KMS CryptoKey to be used to protect access
	   to messages published on this topic. Your project's PubSub service account
	   (`service-{{PROJECT_NUMBER}}@gcp-sa-pubsub.iam.gserviceaccount.com`) must have
	   `roles/cloudkms.cryptoKeyEncrypterDecrypter` to use this feature.
	   The expected format is `projects/-/locations/-/keyRings/-/cryptoKeys/-`
	*/
	KmsKeyName string `json:"kmsKeyName,omitempty" yaml:"kmsKeyName,omitempty"`

	/*
	   A set of key/value label pairs to assign to this Topic.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   Indicates the minimum duration to retain a message after it is published
	   to the topic. If this field is set, messages published to the topic in
	   the last messageRetentionDuration are always available to subscribers.
	   For instance, it allows any attached subscription to seek to a timestamp
	   that is up to messageRetentionDuration in the past. If this field is not
	   set, message retention is controlled by settings on individual subscriptions.
	   The rotation period has the format of a decimal number, followed by the
	   letter `s` (seconds). Cannot be more than 31 days or less than 10 minutes.
	*/
	MessageRetentionDuration string `json:"messageRetentionDuration,omitempty" yaml:"messageRetentionDuration,omitempty"`

	/*
	   Policy constraining the set of Google Cloud Platform regions where
	   messages published to the topic may be stored. If not present, then no
	   constraints are in effect.
	   Structure is documented below.
	*/
	MessageStoragePolicy types.Pubsub_TopicMessageStoragePolicy `json:"messageStoragePolicy,omitempty" yaml:"messageStoragePolicy,omitempty"`

	/*
	   Name of the topic.


	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
