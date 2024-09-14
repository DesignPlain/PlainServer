package types

type Vertex_AiFeatureOnlineStoreFeatureviewFeatureRegistrySource struct {
	/*
	   List of features that need to be synced to Online Store.
	   Structure is documented below.
	*/
	FeatureGroups []Vertex_AiFeatureOnlineStoreFeatureviewFeatureRegistrySourceFeatureGroup `json:"featureGroups,omitempty" yaml:"featureGroups,omitempty"`
}
