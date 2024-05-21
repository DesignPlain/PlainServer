package types

type Vertex_AiIndexIndexStat struct {
	/*
	   (Output)
	   The number of shards in the Index.
	*/
	ShardsCount int `json:"shardsCount,omitempty" yaml:"shardsCount,omitempty"`

	/*
	   (Output)
	   The number of vectors in the Index.
	*/
	VectorsCount string `json:"vectorsCount,omitempty" yaml:"vectorsCount,omitempty"`
}
