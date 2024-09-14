package vertex

import types "libds/gcp/types"

type AiFeatureOnlineStore struct {
	/*
	   The resource name of the Feature Online Store. This value may be up to 60 characters, and valid characters are [a-z0-9_]. The first character cannot be a number.


	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The region of feature online store. eg us-central1
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	// If set to true, any FeatureViews and Features for this FeatureOnlineStore will also be deleted.
	ForceDestroy bool `json:"forceDestroy,omitempty" yaml:"forceDestroy,omitempty"`

	/*
	   The dedicated serving endpoint for this FeatureOnlineStore, which is different from common vertex service endpoint. Only need to set when you choose Optimized storage type or enable EmbeddingManagement. Will use public endpoint by default.
	   Structure is documented below.
	*/
	DedicatedServingEndpoint types.Vertex_AiFeatureOnlineStoreDedicatedServingEndpoint `json:"dedicatedServingEndpoint,omitempty" yaml:"dedicatedServingEndpoint,omitempty"`

	/*
	   The settings for embedding management in FeatureOnlineStore. Embedding management can only be used with BigTable.
	   Structure is documented below.
	*/
	EmbeddingManagement types.Vertex_AiFeatureOnlineStoreEmbeddingManagement `json:"embeddingManagement,omitempty" yaml:"embeddingManagement,omitempty"`

	/*
	   The labels with user-defined metadata to organize your feature online stores.
	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// Settings for the Optimized store that will be created to serve featureValues for all FeatureViews under this FeatureOnlineStore
	Optimized types.Vertex_AiFeatureOnlineStoreOptimized `json:"optimized,omitempty" yaml:"optimized,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Settings for Cloud Bigtable instance that will be created to serve featureValues for all FeatureViews under this FeatureOnlineStore.
	   Structure is documented below.
	*/
	Bigtable types.Vertex_AiFeatureOnlineStoreBigtable `json:"bigtable,omitempty" yaml:"bigtable,omitempty"`
}
