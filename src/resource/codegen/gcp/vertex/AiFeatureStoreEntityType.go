package vertex

import types "libds/gcp/types"

type AiFeatureStoreEntityType struct {
	/*
	   A set of key/value label pairs to assign to this EntityType.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   The default monitoring configuration for all Features under this EntityType.
	   If this is populated with [FeaturestoreMonitoringConfig.monitoring_interval] specified, snapshot analysis monitoring is enabled. Otherwise, snapshot analysis monitoring is disabled.
	   Structure is documented below.
	*/
	MonitoringConfig types.Vertex_AiFeatureStoreEntityTypeMonitoringConfig `json:"monitoringConfig,omitempty" yaml:"monitoringConfig,omitempty"`

	// The name of the EntityType. This value may be up to 60 characters, and valid characters are [a-z0-9_]. The first character cannot be a number.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Config for data retention policy in offline storage. TTL in days for feature values that will be stored in offline storage. The Feature Store offline storage periodically removes obsolete feature values older than offlineStorageTtlDays since the feature generation time. If unset (or explicitly set to 0), default to 4000 days TTL.
	OfflineStorageTtlDays int `json:"offlineStorageTtlDays,omitempty" yaml:"offlineStorageTtlDays,omitempty"`

	// Optional. Description of the EntityType.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   The name of the Featurestore to use, in the format projects/{project}/locations/{location}/featurestores/{featurestore}.


	   - - -
	*/
	Featurestore string `json:"featurestore,omitempty" yaml:"featurestore,omitempty"`
}
