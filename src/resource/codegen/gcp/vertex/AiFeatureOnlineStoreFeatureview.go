package vertex

import types "libds/gcp/types"

type AiFeatureOnlineStoreFeatureview struct {
	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Configuration for vector search. It contains the required configurations to create an index from source data, so that approximate nearest neighbor (a.k.a ANN) algorithms search can be performed during online serving.
	   Structure is documented below.
	*/
	VectorSearchConfig types.Vertex_AiFeatureOnlineStoreFeatureviewVectorSearchConfig `json:"vectorSearchConfig,omitempty" yaml:"vectorSearchConfig,omitempty"`

	/*
	   Configures how data is supposed to be extracted from a BigQuery source to be loaded onto the FeatureOnlineStore.
	   Structure is documented below.
	*/
	BigQuerySource types.Vertex_AiFeatureOnlineStoreFeatureviewBigQuerySource `json:"bigQuerySource,omitempty" yaml:"bigQuerySource,omitempty"`

	// The name of the FeatureOnlineStore to use for the featureview.
	FeatureOnlineStore string `json:"featureOnlineStore,omitempty" yaml:"featureOnlineStore,omitempty"`

	/*
	   A set of key/value label pairs to assign to this FeatureView.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// Name of the FeatureView. This value may be up to 60 characters, and valid characters are [a-z0-9_]. The first character cannot be a number.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   Configures the features from a Feature Registry source that need to be loaded onto the FeatureOnlineStore.
	   Structure is documented below.
	*/
	FeatureRegistrySource types.Vertex_AiFeatureOnlineStoreFeatureviewFeatureRegistrySource `json:"featureRegistrySource,omitempty" yaml:"featureRegistrySource,omitempty"`

	/*
	   The region for the resource. It should be the same as the featureonlinestore region.


	   - - -
	*/
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	/*
	   Configures when data is to be synced/updated for this FeatureView. At the end of the sync the latest featureValues for each entityId of this FeatureView are made ready for online serving.
	   Structure is documented below.
	*/
	SyncConfig types.Vertex_AiFeatureOnlineStoreFeatureviewSyncConfig `json:"syncConfig,omitempty" yaml:"syncConfig,omitempty"`
}
