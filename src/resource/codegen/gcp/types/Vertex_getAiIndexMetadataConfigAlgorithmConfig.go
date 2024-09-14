package types

type Vertex_getAiIndexMetadataConfigAlgorithmConfig struct {
	/*
	   Configuration options for using brute force search, which simply implements the
	   standard linear search in the database for each query.
	*/
	BruteForceConfigs []Vertex_getAiIndexMetadataConfigAlgorithmConfigBruteForceConfig `json:"bruteForceConfigs,omitempty" yaml:"bruteForceConfigs,omitempty"`

	/*
	   Configuration options for using the tree-AH algorithm (Shallow tree + Asymmetric Hashing).
	   Please refer to this paper for more details: https://arxiv.org/abs/1908.10396
	*/
	TreeAhConfigs []Vertex_getAiIndexMetadataConfigAlgorithmConfigTreeAhConfig `json:"treeAhConfigs,omitempty" yaml:"treeAhConfigs,omitempty"`
}
