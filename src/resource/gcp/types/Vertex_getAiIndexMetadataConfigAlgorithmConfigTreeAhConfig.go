package types

type Vertex_getAiIndexMetadataConfigAlgorithmConfigTreeAhConfig struct {
	/*
	   The default percentage of leaf nodes that any query may be searched. Must be in
	   range 1-100, inclusive. The default value is 10 (means 10%!!(MISSING))(MISSING) if not set.
	*/
	LeafNodesToSearchPercent int `json:"leafNodesToSearchPercent,omitempty" yaml:"leafNodesToSearchPercent,omitempty"`

	// Number of embeddings on each leaf node. The default value is 1000 if not set.
	LeafNodeEmbeddingCount int `json:"leafNodeEmbeddingCount,omitempty" yaml:"leafNodeEmbeddingCount,omitempty"`
}
