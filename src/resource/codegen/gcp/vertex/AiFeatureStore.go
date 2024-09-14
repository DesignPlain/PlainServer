package vertex

import types "libds/gcp/types"

type AiFeatureStore struct {
	// TTL in days for feature values that will be stored in online serving storage. The Feature Store online storage periodically removes obsolete feature values older than onlineStorageTtlDays since the feature generation time. Note that onlineStorageTtlDays should be less than or equal to offlineStorageTtlDays for each EntityType under a featurestore. If not set, default to 4000 days
	OnlineStorageTtlDays int `json:"onlineStorageTtlDays,omitempty" yaml:"onlineStorageTtlDays,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// The region of the dataset. eg us-central1
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	/*
	   If set, both of the online and offline data storage will be secured by this key.
	   Structure is documented below.
	*/
	EncryptionSpec types.Vertex_AiFeatureStoreEncryptionSpec `json:"encryptionSpec,omitempty" yaml:"encryptionSpec,omitempty"`

	// If set to true, any EntityTypes and Features for this Featurestore will also be deleted
	ForceDestroy bool `json:"forceDestroy,omitempty" yaml:"forceDestroy,omitempty"`

	/*
	   A set of key/value label pairs to assign to this Featurestore.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// The name of the Featurestore. This value may be up to 60 characters, and valid characters are [a-z0-9_]. The first character cannot be a number.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   Config for online serving resources.
	   Structure is documented below.
	*/
	OnlineServingConfig types.Vertex_AiFeatureStoreOnlineServingConfig `json:"onlineServingConfig,omitempty" yaml:"onlineServingConfig,omitempty"`
}
