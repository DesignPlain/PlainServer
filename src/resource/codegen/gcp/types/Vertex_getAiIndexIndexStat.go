package types

type Vertex_getAiIndexIndexStat struct {
	// The number of shards in the Index.
	ShardsCount int `json:"shardsCount,omitempty" yaml:"shardsCount,omitempty"`

	// The number of vectors in the Index.
	VectorsCount string `json:"vectorsCount,omitempty" yaml:"vectorsCount,omitempty"`
}
